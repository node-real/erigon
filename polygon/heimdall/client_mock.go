// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/polygon/heimdall (interfaces: HeimdallClient)
//
// Generated by this command:
//
//	mockgen -destination=./client_mock.go -package=heimdall . HeimdallClient
//

// Package heimdall is a generated GoMock package.
package heimdall

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockHeimdallClient is a mock of HeimdallClient interface.
type MockHeimdallClient struct {
	ctrl     *gomock.Controller
	recorder *MockHeimdallClientMockRecorder
}

// MockHeimdallClientMockRecorder is the mock recorder for MockHeimdallClient.
type MockHeimdallClientMockRecorder struct {
	mock *MockHeimdallClient
}

// NewMockHeimdallClient creates a new mock instance.
func NewMockHeimdallClient(ctrl *gomock.Controller) *MockHeimdallClient {
	mock := &MockHeimdallClient{ctrl: ctrl}
	mock.recorder = &MockHeimdallClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeimdallClient) EXPECT() *MockHeimdallClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockHeimdallClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockHeimdallClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHeimdallClient)(nil).Close))
}

// FetchCheckpoint mocks base method.
func (m *MockHeimdallClient) FetchCheckpoint(arg0 context.Context, arg1 int64) (*Checkpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpoint", arg0, arg1)
	ret0, _ := ret[0].(*Checkpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpoint indicates an expected call of FetchCheckpoint.
func (mr *MockHeimdallClientMockRecorder) FetchCheckpoint(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpoint", reflect.TypeOf((*MockHeimdallClient)(nil).FetchCheckpoint), arg0, arg1)
}

// FetchCheckpointCount mocks base method.
func (m *MockHeimdallClient) FetchCheckpointCount(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpointCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpointCount indicates an expected call of FetchCheckpointCount.
func (mr *MockHeimdallClientMockRecorder) FetchCheckpointCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpointCount", reflect.TypeOf((*MockHeimdallClient)(nil).FetchCheckpointCount), arg0)
}

// FetchLastNoAckMilestone mocks base method.
func (m *MockHeimdallClient) FetchLastNoAckMilestone(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchLastNoAckMilestone", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchLastNoAckMilestone indicates an expected call of FetchLastNoAckMilestone.
func (mr *MockHeimdallClientMockRecorder) FetchLastNoAckMilestone(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchLastNoAckMilestone", reflect.TypeOf((*MockHeimdallClient)(nil).FetchLastNoAckMilestone), arg0)
}

// FetchLatestSpan mocks base method.
func (m *MockHeimdallClient) FetchLatestSpan(arg0 context.Context) (*Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchLatestSpan", arg0)
	ret0, _ := ret[0].(*Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchLatestSpan indicates an expected call of FetchLatestSpan.
func (mr *MockHeimdallClientMockRecorder) FetchLatestSpan(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchLatestSpan", reflect.TypeOf((*MockHeimdallClient)(nil).FetchLatestSpan), arg0)
}

// FetchMilestone mocks base method.
func (m *MockHeimdallClient) FetchMilestone(arg0 context.Context, arg1 int64) (*Milestone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMilestone", arg0, arg1)
	ret0, _ := ret[0].(*Milestone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMilestone indicates an expected call of FetchMilestone.
func (mr *MockHeimdallClientMockRecorder) FetchMilestone(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMilestone", reflect.TypeOf((*MockHeimdallClient)(nil).FetchMilestone), arg0, arg1)
}

// FetchMilestoneCount mocks base method.
func (m *MockHeimdallClient) FetchMilestoneCount(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMilestoneCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMilestoneCount indicates an expected call of FetchMilestoneCount.
func (mr *MockHeimdallClientMockRecorder) FetchMilestoneCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMilestoneCount", reflect.TypeOf((*MockHeimdallClient)(nil).FetchMilestoneCount), arg0)
}

// FetchMilestoneID mocks base method.
func (m *MockHeimdallClient) FetchMilestoneID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMilestoneID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchMilestoneID indicates an expected call of FetchMilestoneID.
func (mr *MockHeimdallClientMockRecorder) FetchMilestoneID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMilestoneID", reflect.TypeOf((*MockHeimdallClient)(nil).FetchMilestoneID), arg0, arg1)
}

// FetchNoAckMilestone mocks base method.
func (m *MockHeimdallClient) FetchNoAckMilestone(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchNoAckMilestone", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchNoAckMilestone indicates an expected call of FetchNoAckMilestone.
func (mr *MockHeimdallClientMockRecorder) FetchNoAckMilestone(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNoAckMilestone", reflect.TypeOf((*MockHeimdallClient)(nil).FetchNoAckMilestone), arg0, arg1)
}

// FetchSpan mocks base method.
func (m *MockHeimdallClient) FetchSpan(arg0 context.Context, arg1 uint64) (*Span, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSpan", arg0, arg1)
	ret0, _ := ret[0].(*Span)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSpan indicates an expected call of FetchSpan.
func (mr *MockHeimdallClientMockRecorder) FetchSpan(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSpan", reflect.TypeOf((*MockHeimdallClient)(nil).FetchSpan), arg0, arg1)
}

// FetchStateSyncEvents mocks base method.
func (m *MockHeimdallClient) FetchStateSyncEvents(arg0 context.Context, arg1 uint64, arg2 time.Time, arg3 int) ([]*EventRecordWithTime, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchStateSyncEvents", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*EventRecordWithTime)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchStateSyncEvents indicates an expected call of FetchStateSyncEvents.
func (mr *MockHeimdallClientMockRecorder) FetchStateSyncEvents(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchStateSyncEvents", reflect.TypeOf((*MockHeimdallClient)(nil).FetchStateSyncEvents), arg0, arg1, arg2, arg3)
}