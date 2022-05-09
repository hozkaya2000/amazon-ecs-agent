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
// Source: github.com/aws/amazon-ecs-agent/agent/engine/service_connect (interfaces: Manager)

// Package mock_service_connect is a generated GoMock package.
package mock_service_connect

import (
	reflect "reflect"

	container "github.com/aws/amazon-ecs-agent/agent/api/container"
	task "github.com/aws/amazon-ecs-agent/agent/api/task"
	container0 "github.com/docker/docker/api/types/container"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AugmentTaskContainer mocks base method
func (m *MockManager) AugmentTaskContainer(arg0 *task.Task, arg1 *container.Container, arg2 *container0.HostConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AugmentTaskContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AugmentTaskContainer indicates an expected call of AugmentTaskContainer
func (mr *MockManagerMockRecorder) AugmentTaskContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AugmentTaskContainer", reflect.TypeOf((*MockManager)(nil).AugmentTaskContainer), arg0, arg1, arg2)
}