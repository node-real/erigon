// Copyright 2024 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package ethconsensusconfig

import (
	"context"
	"github.com/ledgerwatch/erigon/core/blob_storage"
	"github.com/spf13/afero"
	"math"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"

	"github.com/ledgerwatch/erigon-lib/log/v3"

	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/ledgerwatch/erigon/polygon/bor/borcfg"
	"github.com/ledgerwatch/erigon/polygon/bridge"

	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon/consensus"
	"github.com/ledgerwatch/erigon/consensus/aura"
	"github.com/ledgerwatch/erigon/consensus/clique"
	"github.com/ledgerwatch/erigon/consensus/ethash"
	"github.com/ledgerwatch/erigon/consensus/ethash/ethashcfg"
	"github.com/ledgerwatch/erigon/consensus/merge"
	"github.com/ledgerwatch/erigon/consensus/parlia"
	"github.com/ledgerwatch/erigon/node"
	"github.com/ledgerwatch/erigon/node/nodecfg"
	"github.com/ledgerwatch/erigon/params"
	"github.com/ledgerwatch/erigon/polygon/bor"
	"github.com/ledgerwatch/erigon/polygon/heimdall"
	"github.com/ledgerwatch/erigon/turbo/services"
)

func CreateConsensusEngine(ctx context.Context, nodeConfig *nodecfg.Config, chainConfig *chain.Config, config interface{}, notify []string, noVerify bool,
	heimdallClient heimdall.HeimdallClient, withoutHeimdall bool, blockReader services.FullBlockReader, readonly bool,
	logger log.Logger, chainDb kv.RwDB, polygonBridge bridge.Service,
) consensus.Engine {
	var eng consensus.Engine

	switch consensusCfg := config.(type) {
	case *ethashcfg.Config:
		switch consensusCfg.PowMode {
		case ethashcfg.ModeFake:
			logger.Warn("Ethash used in fake mode")
			eng = ethash.NewFaker()
		case ethashcfg.ModeTest:
			logger.Warn("Ethash used in test mode")
			eng = ethash.NewTester(nil, noVerify)
		case ethashcfg.ModeShared:
			logger.Warn("Ethash used in shared mode")
			eng = ethash.NewShared()
		default:
			eng = ethash.New(ethashcfg.Config{
				CachesInMem:      consensusCfg.CachesInMem,
				CachesLockMmap:   consensusCfg.CachesLockMmap,
				DatasetDir:       consensusCfg.DatasetDir,
				DatasetsInMem:    consensusCfg.DatasetsInMem,
				DatasetsOnDisk:   consensusCfg.DatasetsOnDisk,
				DatasetsLockMmap: consensusCfg.DatasetsLockMmap,
			}, notify, noVerify)
		}
	case *params.ConsensusSnapshotConfig:
		if chainConfig.Clique != nil {
			if consensusCfg.InMemory {
				nodeConfig.Dirs.DataDir = ""
			} else {
				if consensusCfg.DBPath != "" {
					if filepath.Base(consensusCfg.DBPath) == "clique" {
						nodeConfig.Dirs.DataDir = filepath.Dir(consensusCfg.DBPath)
					} else {
						nodeConfig.Dirs.DataDir = consensusCfg.DBPath
					}
				}
			}

			var err error
			var db kv.RwDB

			db, err = node.OpenDatabase(ctx, nodeConfig, kv.ConsensusDB, "clique", readonly, logger)

			if err != nil {
				panic(err)
			}

			eng = clique.New(chainConfig, consensusCfg, db, logger)
		}
	case *chain.AuRaConfig:
		if chainConfig.Aura != nil {
			var err error
			var db kv.RwDB

			db, err = node.OpenDatabase(ctx, nodeConfig, kv.ConsensusDB, "aura", readonly, logger)

			if err != nil {
				panic(err)
			}

			eng, err = aura.NewAuRa(chainConfig.Aura, db)
			if err != nil {
				panic(err)
			}
		}
	case *chain.ParliaConfig:
		if chainConfig.Parlia != nil {
			var err error
			var db kv.RwDB

			if consensusCfg.DBPath != "" {
				if filepath.Base(consensusCfg.DBPath) == "parlia" {
					nodeConfig.Dirs.DataDir = filepath.Dir(consensusCfg.DBPath)
				} else {
					nodeConfig.Dirs.DataDir = consensusCfg.DBPath
				}
			}

			db, err = node.OpenDatabase(ctx, nodeConfig, kv.ConsensusDB, "parlia", readonly, logger)
			if err != nil {
				panic(err)
			}
			nodeConfig.Dirs.DataDir = filepath.Join(nodeConfig.Dirs.DataDir, "blobs")
			blobDb, err := node.OpenDatabase(ctx, nodeConfig, kv.BlobDb, "", false, logger)
			if err != nil {
				panic(err)
			}
			blobStore := blob_storage.NewBlobStore(blobDb, afero.NewBasePathFs(afero.NewOsFs(), nodeConfig.Dirs.DataDir), math.MaxUint64, chainConfig, blockReader)

			eng = parlia.New(chainConfig, db, blobStore, blockReader, chainDb, logger)
		}
	case *borcfg.BorConfig:
		// If Matic bor consensus is requested, set it up
		// In order to pass the ethereum transaction tests, we need to set the burn contract which is in the bor config
		// Then, bor != nil will also be enabled for ethash and clique. Only enable Bor for real if there is a validator contract present.
		if chainConfig.Bor != nil && consensusCfg.ValidatorContract != "" {
			genesisContractsClient := bor.NewGenesisContractsClient(chainConfig, consensusCfg.ValidatorContract, consensusCfg.StateReceiverContract, logger)

			spanner := bor.NewChainSpanner(bor.GenesisContractValidatorSetABI(), chainConfig, withoutHeimdall, logger)

			var err error
			var db kv.RwDB

			db, err = node.OpenDatabase(ctx, nodeConfig, kv.ConsensusDB, "bor", readonly, logger)
			if err != nil {
				panic(err)
			}

			eng = bor.New(chainConfig, db, blockReader, spanner, heimdallClient, genesisContractsClient, logger, polygonBridge)
		}
	}

	if eng == nil {
		panic("unknown config" + spew.Sdump(config))
	}

	if chainConfig.TerminalTotalDifficulty == nil {
		return eng
	} else {
		return merge.New(eng) // the Merge
	}
}

func CreateConsensusEngineBareBones(ctx context.Context, chainConfig *chain.Config, logger log.Logger) consensus.Engine {
	var consensusConfig interface{}

	if chainConfig.Clique != nil {
		consensusConfig = params.CliqueSnapshot
	} else if chainConfig.Aura != nil {
		consensusConfig = chainConfig.Aura
	} else if chainConfig.Bor != nil {
		consensusConfig = chainConfig.Bor
	} else if chainConfig.Parlia != nil {
		consensusConfig = chainConfig.Parlia
	} else {
		var ethashCfg ethashcfg.Config
		ethashCfg.PowMode = ethashcfg.ModeFake
		consensusConfig = &ethashCfg
	}

	return CreateConsensusEngine(ctx, &nodecfg.Config{}, chainConfig, consensusConfig, nil /* notify */, true, /* noVerify */
		nil /* heimdallClient */, true /* withoutHeimdall */, nil /* blockReader */, false /* readonly */, logger, nil, nil)
}
