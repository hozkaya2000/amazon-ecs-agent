// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/eni/gonetwrapper (interfaces: GolangNetUtils)

// Package mock_gonetwrapper is a generated GoMock package.
package mock_gonetwrapper

import (
	gomock "github.com/golang/mock/gomock"
	net "net"
	reflect "reflect"
)

// MockGolangNetUtils is a mock of GolangNetUtils interface
type MockGolangNetUtils struct {
	ctrl     *gomock.Controller
	recorder *MockGolangNetUtilsMockRecorder
}

// MockGolangNetUtilsMockRecorder is the mock recorder for MockGolangNetUtils
type MockGolangNetUtilsMockRecorder struct {
	mock *MockGolangNetUtils
}

// NewMockGolangNetUtils creates a new mock instance
func NewMockGolangNetUtils(ctrl *gomock.Controller) *MockGolangNetUtils {
	mock := &MockGolangNetUtils{ctrl: ctrl}
	mock.recorder = &MockGolangNetUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGolangNetUtils) EXPECT() *MockGolangNetUtilsMockRecorder {
	return m.recorder
}

// FindInterfaceByIndex mocks base method
func (m *MockGolangNetUtils) FindInterfaceByIndex(arg0 int) (*net.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInterfaceByIndex", arg0)
	ret0, _ := ret[0].(*net.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindInterfaceByIndex indicates an expected call of FindInterfaceByIndex
func (mr *MockGolangNetUtilsMockRecorder) FindInterfaceByIndex(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInterfaceByIndex", reflect.TypeOf((*MockGolangNetUtils)(nil).FindInterfaceByIndex), arg0)
}

// GetAllNetworkInterfaces mocks base method
func (m *MockGolangNetUtils) GetAllNetworkInterfaces() ([]net.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNetworkInterfaces")
	ret0, _ := ret[0].([]net.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNetworkInterfaces indicates an expected call of GetAllNetworkInterfaces
func (mr *MockGolangNetUtilsMockRecorder) GetAllNetworkInterfaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNetworkInterfaces", reflect.TypeOf((*MockGolangNetUtils)(nil).GetAllNetworkInterfaces))
}
