// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/k8s (interfaces: Controller)

// Package k8s is a generated GoMock package.
package k8s

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/openservicemesh/osm/pkg/apis/config/v1alpha2"
	v1alpha1 "github.com/openservicemesh/osm/pkg/apis/policy/v1alpha1"
	envoy "github.com/openservicemesh/osm/pkg/envoy"
	identity "github.com/openservicemesh/osm/pkg/identity"
	service "github.com/openservicemesh/osm/pkg/service"
	v1 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// GetEndpoints mocks base method.
func (m *MockController) GetEndpoints(arg0 service.MeshService) (*v1.Endpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEndpoints", arg0)
	ret0, _ := ret[0].(*v1.Endpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEndpoints indicates an expected call of GetEndpoints.
func (mr *MockControllerMockRecorder) GetEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndpoints", reflect.TypeOf((*MockController)(nil).GetEndpoints), arg0)
}

// GetMeshConfig mocks base method.
func (m *MockController) GetMeshConfig() v1alpha2.MeshConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshConfig")
	ret0, _ := ret[0].(v1alpha2.MeshConfig)
	return ret0
}

// GetMeshConfig indicates an expected call of GetMeshConfig.
func (mr *MockControllerMockRecorder) GetMeshConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshConfig", reflect.TypeOf((*MockController)(nil).GetMeshConfig))
}

// GetNamespace mocks base method.
func (m *MockController) GetNamespace(arg0 string) *v1.Namespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", arg0)
	ret0, _ := ret[0].(*v1.Namespace)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockControllerMockRecorder) GetNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockController)(nil).GetNamespace), arg0)
}

// GetOSMNamespace mocks base method.
func (m *MockController) GetOSMNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOSMNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOSMNamespace indicates an expected call of GetOSMNamespace.
func (mr *MockControllerMockRecorder) GetOSMNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOSMNamespace", reflect.TypeOf((*MockController)(nil).GetOSMNamespace))
}

// GetPodForProxy mocks base method.
func (m *MockController) GetPodForProxy(arg0 *envoy.Proxy) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodForProxy", arg0)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodForProxy indicates an expected call of GetPodForProxy.
func (mr *MockControllerMockRecorder) GetPodForProxy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodForProxy", reflect.TypeOf((*MockController)(nil).GetPodForProxy), arg0)
}

// GetService mocks base method.
func (m *MockController) GetService(arg0 service.MeshService) *v1.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0)
	ret0, _ := ret[0].(*v1.Service)
	return ret0
}

// GetService indicates an expected call of GetService.
func (mr *MockControllerMockRecorder) GetService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockController)(nil).GetService), arg0)
}

// GetTargetPortForServicePort mocks base method.
func (m *MockController) GetTargetPortForServicePort(arg0 types.NamespacedName, arg1 uint16) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetPortForServicePort", arg0, arg1)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetPortForServicePort indicates an expected call of GetTargetPortForServicePort.
func (mr *MockControllerMockRecorder) GetTargetPortForServicePort(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetPortForServicePort", reflect.TypeOf((*MockController)(nil).GetTargetPortForServicePort), arg0, arg1)
}

// GetUpstreamTrafficSetting mocks base method.
func (m *MockController) GetUpstreamTrafficSetting(arg0 *types.NamespacedName) *v1alpha1.UpstreamTrafficSetting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpstreamTrafficSetting", arg0)
	ret0, _ := ret[0].(*v1alpha1.UpstreamTrafficSetting)
	return ret0
}

// GetUpstreamTrafficSetting indicates an expected call of GetUpstreamTrafficSetting.
func (mr *MockControllerMockRecorder) GetUpstreamTrafficSetting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpstreamTrafficSetting", reflect.TypeOf((*MockController)(nil).GetUpstreamTrafficSetting), arg0)
}

// IsMonitoredNamespace mocks base method.
func (m *MockController) IsMonitoredNamespace(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMonitoredNamespace", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMonitoredNamespace indicates an expected call of IsMonitoredNamespace.
func (mr *MockControllerMockRecorder) IsMonitoredNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMonitoredNamespace", reflect.TypeOf((*MockController)(nil).IsMonitoredNamespace), arg0)
}

// ListEgressPolicies mocks base method.
func (m *MockController) ListEgressPolicies() []*v1alpha1.Egress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEgressPolicies")
	ret0, _ := ret[0].([]*v1alpha1.Egress)
	return ret0
}

// ListEgressPolicies indicates an expected call of ListEgressPolicies.
func (mr *MockControllerMockRecorder) ListEgressPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEgressPolicies", reflect.TypeOf((*MockController)(nil).ListEgressPolicies))
}

// ListIngressBackendPolicies mocks base method.
func (m *MockController) ListIngressBackendPolicies() []*v1alpha1.IngressBackend {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIngressBackendPolicies")
	ret0, _ := ret[0].([]*v1alpha1.IngressBackend)
	return ret0
}

// ListIngressBackendPolicies indicates an expected call of ListIngressBackendPolicies.
func (mr *MockControllerMockRecorder) ListIngressBackendPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIngressBackendPolicies", reflect.TypeOf((*MockController)(nil).ListIngressBackendPolicies))
}

// ListMonitoredNamespaces mocks base method.
func (m *MockController) ListMonitoredNamespaces() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMonitoredNamespaces")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMonitoredNamespaces indicates an expected call of ListMonitoredNamespaces.
func (mr *MockControllerMockRecorder) ListMonitoredNamespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMonitoredNamespaces", reflect.TypeOf((*MockController)(nil).ListMonitoredNamespaces))
}

// ListPods mocks base method.
func (m *MockController) ListPods() []*v1.Pod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPods")
	ret0, _ := ret[0].([]*v1.Pod)
	return ret0
}

// ListPods indicates an expected call of ListPods.
func (mr *MockControllerMockRecorder) ListPods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPods", reflect.TypeOf((*MockController)(nil).ListPods))
}

// ListRetryPolicies mocks base method.
func (m *MockController) ListRetryPolicies() []*v1alpha1.Retry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRetryPolicies")
	ret0, _ := ret[0].([]*v1alpha1.Retry)
	return ret0
}

// ListRetryPolicies indicates an expected call of ListRetryPolicies.
func (mr *MockControllerMockRecorder) ListRetryPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRetryPolicies", reflect.TypeOf((*MockController)(nil).ListRetryPolicies))
}

// ListServiceAccounts mocks base method.
func (m *MockController) ListServiceAccounts() []*v1.ServiceAccount {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServiceAccounts")
	ret0, _ := ret[0].([]*v1.ServiceAccount)
	return ret0
}

// ListServiceAccounts indicates an expected call of ListServiceAccounts.
func (mr *MockControllerMockRecorder) ListServiceAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceAccounts", reflect.TypeOf((*MockController)(nil).ListServiceAccounts))
}

// ListServiceIdentitiesForService mocks base method.
func (m *MockController) ListServiceIdentitiesForService(arg0 service.MeshService) ([]identity.K8sServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServiceIdentitiesForService", arg0)
	ret0, _ := ret[0].([]identity.K8sServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServiceIdentitiesForService indicates an expected call of ListServiceIdentitiesForService.
func (mr *MockControllerMockRecorder) ListServiceIdentitiesForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceIdentitiesForService", reflect.TypeOf((*MockController)(nil).ListServiceIdentitiesForService), arg0)
}

// ListServices mocks base method.
func (m *MockController) ListServices() []*v1.Service {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices")
	ret0, _ := ret[0].([]*v1.Service)
	return ret0
}

// ListServices indicates an expected call of ListServices.
func (mr *MockControllerMockRecorder) ListServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockController)(nil).ListServices))
}

// ListUpstreamTrafficSettings mocks base method.
func (m *MockController) ListUpstreamTrafficSettings() []*v1alpha1.UpstreamTrafficSetting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUpstreamTrafficSettings")
	ret0, _ := ret[0].([]*v1alpha1.UpstreamTrafficSetting)
	return ret0
}

// ListUpstreamTrafficSettings indicates an expected call of ListUpstreamTrafficSettings.
func (mr *MockControllerMockRecorder) ListUpstreamTrafficSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUpstreamTrafficSettings", reflect.TypeOf((*MockController)(nil).ListUpstreamTrafficSettings))
}

// ServiceToMeshServices mocks base method.
func (m *MockController) ServiceToMeshServices(arg0 v1.Service) []service.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceToMeshServices", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	return ret0
}

// ServiceToMeshServices indicates an expected call of ServiceToMeshServices.
func (mr *MockControllerMockRecorder) ServiceToMeshServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceToMeshServices", reflect.TypeOf((*MockController)(nil).ServiceToMeshServices), arg0)
}

// UpdateIngressBackendStatus mocks base method.
func (m *MockController) UpdateIngressBackendStatus(arg0 *v1alpha1.IngressBackend) (*v1alpha1.IngressBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIngressBackendStatus", arg0)
	ret0, _ := ret[0].(*v1alpha1.IngressBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIngressBackendStatus indicates an expected call of UpdateIngressBackendStatus.
func (mr *MockControllerMockRecorder) UpdateIngressBackendStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIngressBackendStatus", reflect.TypeOf((*MockController)(nil).UpdateIngressBackendStatus), arg0)
}

// UpdateUpstreamTrafficSettingStatus mocks base method.
func (m *MockController) UpdateUpstreamTrafficSettingStatus(arg0 *v1alpha1.UpstreamTrafficSetting) (*v1alpha1.UpstreamTrafficSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUpstreamTrafficSettingStatus", arg0)
	ret0, _ := ret[0].(*v1alpha1.UpstreamTrafficSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUpstreamTrafficSettingStatus indicates an expected call of UpdateUpstreamTrafficSettingStatus.
func (mr *MockControllerMockRecorder) UpdateUpstreamTrafficSettingStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUpstreamTrafficSettingStatus", reflect.TypeOf((*MockController)(nil).UpdateUpstreamTrafficSettingStatus), arg0)
}
