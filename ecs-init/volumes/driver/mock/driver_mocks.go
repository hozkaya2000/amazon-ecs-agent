// Copyright 2015-2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
// Source: driver.go in package 
// Code generated by MockGen. DO NOT EDIT.

// Package mock_driver is a generated GoMock package.
package mock_driver

import (
	reflect "reflect"

	driver "github.com/aws/amazon-ecs-agent/ecs-init/volumes/driver"
	types "github.com/aws/amazon-ecs-agent/ecs-init/volumes/types"
	gomock "github.com/golang/mock/gomock"
)

// MockVolumeDriver is a mock of VolumeDriver interface.
type MockVolumeDriver struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeDriverMockRecorder
}

// MockVolumeDriverMockRecorder is the mock recorder for MockVolumeDriver.
type MockVolumeDriverMockRecorder struct {
	mock *MockVolumeDriver
}

// NewMockVolumeDriver creates a new mock instance.
func NewMockVolumeDriver(ctrl *gomock.Controller) *MockVolumeDriver {
	mock := &MockVolumeDriver{ctrl: ctrl}
	mock.recorder = &MockVolumeDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumeDriver) EXPECT() *MockVolumeDriverMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVolumeDriver) Create(createRequest *driver.CreateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", createRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockVolumeDriverMockRecorder) Create(createRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVolumeDriver)(nil).Create), createRequest)
}

// IsMounted mocks base method.
func (m *MockVolumeDriver) IsMounted(volumeName string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMounted", volumeName)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMounted indicates an expected call of IsMounted.
func (mr *MockVolumeDriverMockRecorder) IsMounted(volumeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMounted", reflect.TypeOf((*MockVolumeDriver)(nil).IsMounted), volumeName)
}

// Remove mocks base method.
func (m *MockVolumeDriver) Remove(removeRequest *driver.RemoveRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", removeRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockVolumeDriverMockRecorder) Remove(removeRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockVolumeDriver)(nil).Remove), removeRequest)
}

// Setup mocks base method.
func (m *MockVolumeDriver) Setup(volumeName string, volume *types.Volume) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Setup", volumeName, volume)
}

// Setup indicates an expected call of Setup.
func (mr *MockVolumeDriverMockRecorder) Setup(volumeName, volume interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockVolumeDriver)(nil).Setup), volumeName, volume)
}
