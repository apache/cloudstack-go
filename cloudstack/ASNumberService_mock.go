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
// Source: ./cloudstack/ASNumberService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockASNumberServiceIface is a mock of ASNumberServiceIface interface.
type MockASNumberServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockASNumberServiceIfaceMockRecorder
}

// MockASNumberServiceIfaceMockRecorder is the mock recorder for MockASNumberServiceIface.
type MockASNumberServiceIfaceMockRecorder struct {
	mock *MockASNumberServiceIface
}

// NewMockASNumberServiceIface creates a new mock instance.
func NewMockASNumberServiceIface(ctrl *gomock.Controller) *MockASNumberServiceIface {
	mock := &MockASNumberServiceIface{ctrl: ctrl}
	mock.recorder = &MockASNumberServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockASNumberServiceIface) EXPECT() *MockASNumberServiceIfaceMockRecorder {
	return m.recorder
}

// ListASNumbers mocks base method.
func (m *MockASNumberServiceIface) ListASNumbers(p *ListASNumbersParams) (*ListASNumbersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListASNumbers", p)
	ret0, _ := ret[0].(*ListASNumbersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListASNumbers indicates an expected call of ListASNumbers.
func (mr *MockASNumberServiceIfaceMockRecorder) ListASNumbers(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListASNumbers", reflect.TypeOf((*MockASNumberServiceIface)(nil).ListASNumbers), p)
}

// NewListASNumbersParams mocks base method.
func (m *MockASNumberServiceIface) NewListASNumbersParams() *ListASNumbersParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListASNumbersParams")
	ret0, _ := ret[0].(*ListASNumbersParams)
	return ret0
}

// NewListASNumbersParams indicates an expected call of NewListASNumbersParams.
func (mr *MockASNumberServiceIfaceMockRecorder) NewListASNumbersParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListASNumbersParams", reflect.TypeOf((*MockASNumberServiceIface)(nil).NewListASNumbersParams))
}

// NewReleaseASNumberParams mocks base method.
func (m *MockASNumberServiceIface) NewReleaseASNumberParams(asnumber int64, zoneid string) *ReleaseASNumberParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReleaseASNumberParams", asnumber, zoneid)
	ret0, _ := ret[0].(*ReleaseASNumberParams)
	return ret0
}

// NewReleaseASNumberParams indicates an expected call of NewReleaseASNumberParams.
func (mr *MockASNumberServiceIfaceMockRecorder) NewReleaseASNumberParams(asnumber, zoneid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReleaseASNumberParams", reflect.TypeOf((*MockASNumberServiceIface)(nil).NewReleaseASNumberParams), asnumber, zoneid)
}

// ReleaseASNumber mocks base method.
func (m *MockASNumberServiceIface) ReleaseASNumber(p *ReleaseASNumberParams) (*ReleaseASNumberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseASNumber", p)
	ret0, _ := ret[0].(*ReleaseASNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseASNumber indicates an expected call of ReleaseASNumber.
func (mr *MockASNumberServiceIfaceMockRecorder) ReleaseASNumber(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseASNumber", reflect.TypeOf((*MockASNumberServiceIface)(nil).ReleaseASNumber), p)
}