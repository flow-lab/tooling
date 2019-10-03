// Code generated by MockGen. DO NOT EDIT.
// Source: msk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFlowMSK is a mock of FlowMSK interface
type MockFlowMSK struct {
	ctrl     *gomock.Controller
	recorder *MockFlowMSKMockRecorder
}

// MockFlowMSKMockRecorder is the mock recorder for MockFlowMSK
type MockFlowMSKMockRecorder struct {
	mock *MockFlowMSK
}

// NewMockFlowMSK creates a new mock instance
func NewMockFlowMSK(ctrl *gomock.Controller) *MockFlowMSK {
	mock := &MockFlowMSK{ctrl: ctrl}
	mock.recorder = &MockFlowMSKMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlowMSK) EXPECT() *MockFlowMSKMockRecorder {
	return m.recorder
}

// GetBootstrapBrokers mocks base method
func (m *MockFlowMSK) GetBootstrapBrokers(clusterName string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBootstrapBrokers", clusterName)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBootstrapBrokers indicates an expected call of GetBootstrapBrokers
func (mr *MockFlowMSKMockRecorder) GetBootstrapBrokers(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBootstrapBrokers", reflect.TypeOf((*MockFlowMSK)(nil).GetBootstrapBrokers), clusterName)
}

// GetClusterArn mocks base method
func (m *MockFlowMSK) GetClusterArn(clusterName string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterArn", clusterName)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterArn indicates an expected call of GetClusterArn
func (mr *MockFlowMSKMockRecorder) GetClusterArn(clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterArn", reflect.TypeOf((*MockFlowMSK)(nil).GetClusterArn), clusterName)
}