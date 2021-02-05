// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/storageos/api-manager/controllers/fencer (interfaces: NodeFencer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/storageos/api-manager/api/v1"
	storageos "github.com/storageos/api-manager/internal/pkg/storageos"
	types "k8s.io/apimachinery/pkg/types"
	reflect "reflect"
)

// MockNodeFencer is a mock of NodeFencer interface
type MockNodeFencer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeFencerMockRecorder
}

// MockNodeFencerMockRecorder is the mock recorder for MockNodeFencer
type MockNodeFencerMockRecorder struct {
	mock *MockNodeFencer
}

// NewMockNodeFencer creates a new mock instance
func NewMockNodeFencer(ctrl *gomock.Controller) *MockNodeFencer {
	mock := &MockNodeFencer{ctrl: ctrl}
	mock.recorder = &MockNodeFencerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeFencer) EXPECT() *MockNodeFencerMockRecorder {
	return m.recorder
}

// GetVolume mocks base method
func (m *MockNodeFencer) GetVolume(arg0 context.Context, arg1 types.NamespacedName) (storageos.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolume", arg0, arg1)
	ret0, _ := ret[0].(storageos.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolume indicates an expected call of GetVolume
func (mr *MockNodeFencerMockRecorder) GetVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolume", reflect.TypeOf((*MockNodeFencer)(nil).GetVolume), arg0, arg1)
}

// ListNodes mocks base method
func (m *MockNodeFencer) ListNodes(arg0 context.Context) ([]v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes", arg0)
	ret0, _ := ret[0].([]v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes
func (mr *MockNodeFencerMockRecorder) ListNodes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockNodeFencer)(nil).ListNodes), arg0)
}
