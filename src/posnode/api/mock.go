// Code generated by MockGen. DO NOT EDIT.
// Source: service.pb.go

// Package api is a generated GoMock package.
package api

import (
	wire "github.com/Fantom-foundation/go-lachesis/src/inter/wire"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockNodeClient is a mock of NodeClient interface
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// SyncEvents mocks base method
func (m *MockNodeClient) SyncEvents(ctx context.Context, in *KnownEvents, opts ...grpc.CallOption) (*KnownEvents, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncEvents", varargs...)
	ret0, _ := ret[0].(*KnownEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncEvents indicates an expected call of SyncEvents
func (mr *MockNodeClientMockRecorder) SyncEvents(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncEvents", reflect.TypeOf((*MockNodeClient)(nil).SyncEvents), varargs...)
}

// GetEvent mocks base method
func (m *MockNodeClient) GetEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*wire.Event, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEvent", varargs...)
	ret0, _ := ret[0].(*wire.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent
func (mr *MockNodeClientMockRecorder) GetEvent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockNodeClient)(nil).GetEvent), varargs...)
}

// GetPeerInfo mocks base method
func (m *MockNodeClient) GetPeerInfo(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPeerInfo", varargs...)
	ret0, _ := ret[0].(*PeerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeerInfo indicates an expected call of GetPeerInfo
func (mr *MockNodeClientMockRecorder) GetPeerInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerInfo", reflect.TypeOf((*MockNodeClient)(nil).GetPeerInfo), varargs...)
}

// MockNodeServer is a mock of NodeServer interface
type MockNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeServerMockRecorder
}

// MockNodeServerMockRecorder is the mock recorder for MockNodeServer
type MockNodeServerMockRecorder struct {
	mock *MockNodeServer
}

// NewMockNodeServer creates a new mock instance
func NewMockNodeServer(ctrl *gomock.Controller) *MockNodeServer {
	mock := &MockNodeServer{ctrl: ctrl}
	mock.recorder = &MockNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeServer) EXPECT() *MockNodeServerMockRecorder {
	return m.recorder
}

// SyncEvents mocks base method
func (m *MockNodeServer) SyncEvents(arg0 context.Context, arg1 *KnownEvents) (*KnownEvents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncEvents", arg0, arg1)
	ret0, _ := ret[0].(*KnownEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncEvents indicates an expected call of SyncEvents
func (mr *MockNodeServerMockRecorder) SyncEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncEvents", reflect.TypeOf((*MockNodeServer)(nil).SyncEvents), arg0, arg1)
}

// GetEvent mocks base method
func (m *MockNodeServer) GetEvent(arg0 context.Context, arg1 *EventRequest) (*wire.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", arg0, arg1)
	ret0, _ := ret[0].(*wire.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent
func (mr *MockNodeServerMockRecorder) GetEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockNodeServer)(nil).GetEvent), arg0, arg1)
}

// GetPeerInfo mocks base method
func (m *MockNodeServer) GetPeerInfo(arg0 context.Context, arg1 *PeerRequest) (*PeerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeerInfo", arg0, arg1)
	ret0, _ := ret[0].(*PeerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeerInfo indicates an expected call of GetPeerInfo
func (mr *MockNodeServerMockRecorder) GetPeerInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerInfo", reflect.TypeOf((*MockNodeServer)(nil).GetPeerInfo), arg0, arg1)
}
