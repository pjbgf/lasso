// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/lasso/pkg/client (interfaces: SharedClientFactory)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// MockSharedClientFactory is a mock of SharedClientFactory interface.
type MockSharedClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSharedClientFactoryMockRecorder
}

// MockSharedClientFactoryMockRecorder is the mock recorder for MockSharedClientFactory.
type MockSharedClientFactoryMockRecorder struct {
	mock *MockSharedClientFactory
}

// NewMockSharedClientFactory creates a new mock instance.
func NewMockSharedClientFactory(ctrl *gomock.Controller) *MockSharedClientFactory {
	mock := &MockSharedClientFactory{ctrl: ctrl}
	mock.recorder = &MockSharedClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharedClientFactory) EXPECT() *MockSharedClientFactoryMockRecorder {
	return m.recorder
}

// ForKind mocks base method.
func (m *MockSharedClientFactory) ForKind(arg0 schema.GroupVersionKind) (*Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForKind", arg0)
	ret0, _ := ret[0].(*Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForKind indicates an expected call of ForKind.
func (mr *MockSharedClientFactoryMockRecorder) ForKind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForKind", reflect.TypeOf((*MockSharedClientFactory)(nil).ForKind), arg0)
}

// ForResource mocks base method.
func (m *MockSharedClientFactory) ForResource(arg0 schema.GroupVersionResource, arg1 bool) (*Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForResource", arg0, arg1)
	ret0, _ := ret[0].(*Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForResource indicates an expected call of ForResource.
func (mr *MockSharedClientFactoryMockRecorder) ForResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForResource", reflect.TypeOf((*MockSharedClientFactory)(nil).ForResource), arg0, arg1)
}

// ForResourceKind mocks base method.
func (m *MockSharedClientFactory) ForResourceKind(arg0 schema.GroupVersionResource, arg1 string, arg2 bool) *Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForResourceKind", arg0, arg1, arg2)
	ret0, _ := ret[0].(*Client)
	return ret0
}

// ForResourceKind indicates an expected call of ForResourceKind.
func (mr *MockSharedClientFactoryMockRecorder) ForResourceKind(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForResourceKind", reflect.TypeOf((*MockSharedClientFactory)(nil).ForResourceKind), arg0, arg1, arg2)
}

// GVKForObject mocks base method.
func (m *MockSharedClientFactory) GVKForObject(arg0 runtime.Object) (schema.GroupVersionKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GVKForObject", arg0)
	ret0, _ := ret[0].(schema.GroupVersionKind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GVKForObject indicates an expected call of GVKForObject.
func (mr *MockSharedClientFactoryMockRecorder) GVKForObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GVKForObject", reflect.TypeOf((*MockSharedClientFactory)(nil).GVKForObject), arg0)
}

// GVKForResource mocks base method.
func (m *MockSharedClientFactory) GVKForResource(arg0 schema.GroupVersionResource) (schema.GroupVersionKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GVKForResource", arg0)
	ret0, _ := ret[0].(schema.GroupVersionKind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GVKForResource indicates an expected call of GVKForResource.
func (mr *MockSharedClientFactoryMockRecorder) GVKForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GVKForResource", reflect.TypeOf((*MockSharedClientFactory)(nil).GVKForResource), arg0)
}

// IsHealthy mocks base method.
func (m *MockSharedClientFactory) IsHealthy(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHealthy", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsHealthy indicates an expected call of IsHealthy.
func (mr *MockSharedClientFactoryMockRecorder) IsHealthy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHealthy", reflect.TypeOf((*MockSharedClientFactory)(nil).IsHealthy), arg0)
}

// IsNamespaced mocks base method.
func (m *MockSharedClientFactory) IsNamespaced(arg0 schema.GroupVersionKind) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNamespaced", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsNamespaced indicates an expected call of IsNamespaced.
func (mr *MockSharedClientFactoryMockRecorder) IsNamespaced(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNamespaced", reflect.TypeOf((*MockSharedClientFactory)(nil).IsNamespaced), arg0)
}

// NewObjects mocks base method.
func (m *MockSharedClientFactory) NewObjects(arg0 schema.GroupVersionKind) (runtime.Object, runtime.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewObjects", arg0)
	ret0, _ := ret[0].(runtime.Object)
	ret1, _ := ret[1].(runtime.Object)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewObjects indicates an expected call of NewObjects.
func (mr *MockSharedClientFactoryMockRecorder) NewObjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewObjects", reflect.TypeOf((*MockSharedClientFactory)(nil).NewObjects), arg0)
}

// ResourceForGVK mocks base method.
func (m *MockSharedClientFactory) ResourceForGVK(arg0 schema.GroupVersionKind) (schema.GroupVersionResource, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceForGVK", arg0)
	ret0, _ := ret[0].(schema.GroupVersionResource)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResourceForGVK indicates an expected call of ResourceForGVK.
func (mr *MockSharedClientFactoryMockRecorder) ResourceForGVK(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceForGVK", reflect.TypeOf((*MockSharedClientFactory)(nil).ResourceForGVK), arg0)
}