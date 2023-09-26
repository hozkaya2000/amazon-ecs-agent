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
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/api/resource (interfaces: EBSDiscovery)

// Package mock_resource is a generated GoMock package.
package mock_resource

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEBSDiscovery is a mock of EBSDiscovery interface.
type MockEBSDiscovery struct {
	ctrl     *gomock.Controller
	recorder *MockEBSDiscoveryMockRecorder
}

// MockEBSDiscoveryMockRecorder is the mock recorder for MockEBSDiscovery.
type MockEBSDiscoveryMockRecorder struct {
	mock *MockEBSDiscovery
}

// NewMockEBSDiscovery creates a new mock instance.
func NewMockEBSDiscovery(ctrl *gomock.Controller) *MockEBSDiscovery {
	mock := &MockEBSDiscovery{ctrl: ctrl}
	mock.recorder = &MockEBSDiscoveryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEBSDiscovery) EXPECT() *MockEBSDiscoveryMockRecorder {
	return m.recorder
}

// ConfirmEBSVolumeIsAttached mocks base method.
func (m *MockEBSDiscovery) ConfirmEBSVolumeIsAttached(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmEBSVolumeIsAttached", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmEBSVolumeIsAttached indicates an expected call of ConfirmEBSVolumeIsAttached.
func (mr *MockEBSDiscoveryMockRecorder) ConfirmEBSVolumeIsAttached(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmEBSVolumeIsAttached", reflect.TypeOf((*MockEBSDiscovery)(nil).ConfirmEBSVolumeIsAttached), arg0, arg1)
}
