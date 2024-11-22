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

package diagnostics

import (
	"encoding/json"
	"net/http"

	"github.com/ledgerwatch/erigon/turbo/node"
)

func SetupNodeInfoAccess(metricsMux *http.ServeMux, node *node.ErigonNode) {
	if metricsMux == nil {
		return
	}

	metricsMux.HandleFunc("/nodeinfo", func(w http.ResponseWriter, r *http.Request) {
		writeNodeInfo(w, node)
	})
}

func writeNodeInfo(w http.ResponseWriter, node *node.ErigonNode) {
	reply, err := node.Backend().NodesInfo(0)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reply)
}
