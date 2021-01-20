// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/storageos/api-manager/internal/pkg/storageos (interfaces: ControlPlane)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/storageos/go-api/v2"
	http "net/http"
	reflect "reflect"
)

// MockControlPlane is a mock of ControlPlane interface
type MockControlPlane struct {
	ctrl     *gomock.Controller
	recorder *MockControlPlaneMockRecorder
}

// MockControlPlaneMockRecorder is the mock recorder for MockControlPlane
type MockControlPlaneMockRecorder struct {
	mock *MockControlPlane
}

// NewMockControlPlane creates a new mock instance
func NewMockControlPlane(ctrl *gomock.Controller) *MockControlPlane {
	mock := &MockControlPlane{ctrl: ctrl}
	mock.recorder = &MockControlPlaneMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockControlPlane) EXPECT() *MockControlPlaneMockRecorder {
	return m.recorder
}

// DeleteNamespace mocks base method
func (m *MockControlPlane) DeleteNamespace(arg0 context.Context, arg1, arg2 string, arg3 *api.DeleteNamespaceOpts) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNamespace indicates an expected call of DeleteNamespace
func (mr *MockControlPlaneMockRecorder) DeleteNamespace(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockControlPlane)(nil).DeleteNamespace), arg0, arg1, arg2, arg3)
}

// DeleteNode mocks base method
func (m *MockControlPlane) DeleteNode(arg0 context.Context, arg1, arg2 string, arg3 *api.DeleteNodeOpts) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNode", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNode indicates an expected call of DeleteNode
func (mr *MockControlPlaneMockRecorder) DeleteNode(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockControlPlane)(nil).DeleteNode), arg0, arg1, arg2, arg3)
}

// GetVolume mocks base method
func (m *MockControlPlane) GetVolume(arg0 context.Context, arg1, arg2 string) (api.Volume, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.Volume)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetVolume indicates an expected call of GetVolume
func (mr *MockControlPlaneMockRecorder) GetVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolume", reflect.TypeOf((*MockControlPlane)(nil).GetVolume), arg0, arg1, arg2)
}

// ListNamespaces mocks base method
func (m *MockControlPlane) ListNamespaces(arg0 context.Context) ([]api.Namespace, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaces", arg0)
	ret0, _ := ret[0].([]api.Namespace)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListNamespaces indicates an expected call of ListNamespaces
func (mr *MockControlPlaneMockRecorder) ListNamespaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockControlPlane)(nil).ListNamespaces), arg0)
}

// ListNodes mocks base method
func (m *MockControlPlane) ListNodes(arg0 context.Context) ([]api.Node, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes", arg0)
	ret0, _ := ret[0].([]api.Node)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListNodes indicates an expected call of ListNodes
func (mr *MockControlPlaneMockRecorder) ListNodes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockControlPlane)(nil).ListNodes), arg0)
}

// ListVolumes mocks base method
func (m *MockControlPlane) ListVolumes(arg0 context.Context, arg1 string) ([]api.Volume, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumes", arg0, arg1)
	ret0, _ := ret[0].([]api.Volume)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListVolumes indicates an expected call of ListVolumes
func (mr *MockControlPlaneMockRecorder) ListVolumes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumes", reflect.TypeOf((*MockControlPlane)(nil).ListVolumes), arg0, arg1)
}

// RefreshJwt mocks base method
func (m *MockControlPlane) RefreshJwt(arg0 context.Context) (api.UserSession, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshJwt", arg0)
	ret0, _ := ret[0].(api.UserSession)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RefreshJwt indicates an expected call of RefreshJwt
func (mr *MockControlPlaneMockRecorder) RefreshJwt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshJwt", reflect.TypeOf((*MockControlPlane)(nil).RefreshJwt), arg0)
}

// SetComputeOnly mocks base method
func (m *MockControlPlane) SetComputeOnly(arg0 context.Context, arg1 string, arg2 api.SetComputeOnlyNodeData, arg3 *api.SetComputeOnlyOpts) (api.Node, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetComputeOnly", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(api.Node)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SetComputeOnly indicates an expected call of SetComputeOnly
func (mr *MockControlPlaneMockRecorder) SetComputeOnly(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetComputeOnly", reflect.TypeOf((*MockControlPlane)(nil).SetComputeOnly), arg0, arg1, arg2, arg3)
}

// UpdateNFSVolumeMountEndpoint mocks base method
func (m *MockControlPlane) UpdateNFSVolumeMountEndpoint(arg0 context.Context, arg1, arg2 string, arg3 api.NfsVolumeMountEndpoint, arg4 *api.UpdateNFSVolumeMountEndpointOpts) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNFSVolumeMountEndpoint", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNFSVolumeMountEndpoint indicates an expected call of UpdateNFSVolumeMountEndpoint
func (mr *MockControlPlaneMockRecorder) UpdateNFSVolumeMountEndpoint(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNFSVolumeMountEndpoint", reflect.TypeOf((*MockControlPlane)(nil).UpdateNFSVolumeMountEndpoint), arg0, arg1, arg2, arg3, arg4)
}

// UpdateNode mocks base method
func (m *MockControlPlane) UpdateNode(arg0 context.Context, arg1 string, arg2 api.UpdateNodeData) (api.Node, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNode", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.Node)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateNode indicates an expected call of UpdateNode
func (mr *MockControlPlaneMockRecorder) UpdateNode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNode", reflect.TypeOf((*MockControlPlane)(nil).UpdateNode), arg0, arg1, arg2)
}
