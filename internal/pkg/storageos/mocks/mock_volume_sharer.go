// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/storageos/api-manager/internal/pkg/storageos (interfaces: VolumeSharer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	storageos "github.com/storageos/api-manager/internal/pkg/storageos"
	reflect "reflect"
)

// MockVolumeSharer is a mock of VolumeSharer interface
type MockVolumeSharer struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeSharerMockRecorder
}

// MockVolumeSharerMockRecorder is the mock recorder for MockVolumeSharer
type MockVolumeSharerMockRecorder struct {
	mock *MockVolumeSharer
}

// NewMockVolumeSharer creates a new mock instance
func NewMockVolumeSharer(ctrl *gomock.Controller) *MockVolumeSharer {
	mock := &MockVolumeSharer{ctrl: ctrl}
	mock.recorder = &MockVolumeSharerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVolumeSharer) EXPECT() *MockVolumeSharerMockRecorder {
	return m.recorder
}

// ListSharedVolumes mocks base method
func (m *MockVolumeSharer) ListSharedVolumes(arg0 context.Context) (storageos.SharedVolumeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSharedVolumes", arg0)
	ret0, _ := ret[0].(storageos.SharedVolumeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSharedVolumes indicates an expected call of ListSharedVolumes
func (mr *MockVolumeSharerMockRecorder) ListSharedVolumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSharedVolumes", reflect.TypeOf((*MockVolumeSharer)(nil).ListSharedVolumes), arg0)
}

// SetExternalEndpoint mocks base method
func (m *MockVolumeSharer) SetExternalEndpoint(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetExternalEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetExternalEndpoint indicates an expected call of SetExternalEndpoint
func (mr *MockVolumeSharerMockRecorder) SetExternalEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExternalEndpoint", reflect.TypeOf((*MockVolumeSharer)(nil).SetExternalEndpoint), arg0, arg1, arg2, arg3)
}
