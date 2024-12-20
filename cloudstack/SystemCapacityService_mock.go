//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/SystemCapacityService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/SystemCapacityService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/SystemCapacityService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSystemCapacityServiceIface is a mock of SystemCapacityServiceIface interface.
type MockSystemCapacityServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockSystemCapacityServiceIfaceMockRecorder
	isgomock struct{}
}

// MockSystemCapacityServiceIfaceMockRecorder is the mock recorder for MockSystemCapacityServiceIface.
type MockSystemCapacityServiceIfaceMockRecorder struct {
	mock *MockSystemCapacityServiceIface
}

// NewMockSystemCapacityServiceIface creates a new mock instance.
func NewMockSystemCapacityServiceIface(ctrl *gomock.Controller) *MockSystemCapacityServiceIface {
	mock := &MockSystemCapacityServiceIface{ctrl: ctrl}
	mock.recorder = &MockSystemCapacityServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemCapacityServiceIface) EXPECT() *MockSystemCapacityServiceIfaceMockRecorder {
	return m.recorder
}

// ListCapacity mocks base method.
func (m *MockSystemCapacityServiceIface) ListCapacity(p *ListCapacityParams) (*ListCapacityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCapacity", p)
	ret0, _ := ret[0].(*ListCapacityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCapacity indicates an expected call of ListCapacity.
func (mr *MockSystemCapacityServiceIfaceMockRecorder) ListCapacity(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCapacity", reflect.TypeOf((*MockSystemCapacityServiceIface)(nil).ListCapacity), p)
}

// NewListCapacityParams mocks base method.
func (m *MockSystemCapacityServiceIface) NewListCapacityParams() *ListCapacityParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListCapacityParams")
	ret0, _ := ret[0].(*ListCapacityParams)
	return ret0
}

// NewListCapacityParams indicates an expected call of NewListCapacityParams.
func (mr *MockSystemCapacityServiceIfaceMockRecorder) NewListCapacityParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListCapacityParams", reflect.TypeOf((*MockSystemCapacityServiceIface)(nil).NewListCapacityParams))
}
