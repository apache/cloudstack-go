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
// Source: ./cloudstack/AlertService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/AlertService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/AlertService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAlertServiceIface is a mock of AlertServiceIface interface.
type MockAlertServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockAlertServiceIfaceMockRecorder
	isgomock struct{}
}

// MockAlertServiceIfaceMockRecorder is the mock recorder for MockAlertServiceIface.
type MockAlertServiceIfaceMockRecorder struct {
	mock *MockAlertServiceIface
}

// NewMockAlertServiceIface creates a new mock instance.
func NewMockAlertServiceIface(ctrl *gomock.Controller) *MockAlertServiceIface {
	mock := &MockAlertServiceIface{ctrl: ctrl}
	mock.recorder = &MockAlertServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertServiceIface) EXPECT() *MockAlertServiceIfaceMockRecorder {
	return m.recorder
}

// ArchiveAlerts mocks base method.
func (m *MockAlertServiceIface) ArchiveAlerts(p *ArchiveAlertsParams) (*ArchiveAlertsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArchiveAlerts", p)
	ret0, _ := ret[0].(*ArchiveAlertsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArchiveAlerts indicates an expected call of ArchiveAlerts.
func (mr *MockAlertServiceIfaceMockRecorder) ArchiveAlerts(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveAlerts", reflect.TypeOf((*MockAlertServiceIface)(nil).ArchiveAlerts), p)
}

// DeleteAlerts mocks base method.
func (m *MockAlertServiceIface) DeleteAlerts(p *DeleteAlertsParams) (*DeleteAlertsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlerts", p)
	ret0, _ := ret[0].(*DeleteAlertsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAlerts indicates an expected call of DeleteAlerts.
func (mr *MockAlertServiceIfaceMockRecorder) DeleteAlerts(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlerts", reflect.TypeOf((*MockAlertServiceIface)(nil).DeleteAlerts), p)
}

// GenerateAlert mocks base method.
func (m *MockAlertServiceIface) GenerateAlert(p *GenerateAlertParams) (*GenerateAlertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAlert", p)
	ret0, _ := ret[0].(*GenerateAlertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAlert indicates an expected call of GenerateAlert.
func (mr *MockAlertServiceIfaceMockRecorder) GenerateAlert(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAlert", reflect.TypeOf((*MockAlertServiceIface)(nil).GenerateAlert), p)
}

// GetAlertByID mocks base method.
func (m *MockAlertServiceIface) GetAlertByID(id string, opts ...OptionFunc) (*Alert, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAlertByID", varargs...)
	ret0, _ := ret[0].(*Alert)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlertByID indicates an expected call of GetAlertByID.
func (mr *MockAlertServiceIfaceMockRecorder) GetAlertByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertByID", reflect.TypeOf((*MockAlertServiceIface)(nil).GetAlertByID), varargs...)
}

// GetAlertByName mocks base method.
func (m *MockAlertServiceIface) GetAlertByName(name string, opts ...OptionFunc) (*Alert, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAlertByName", varargs...)
	ret0, _ := ret[0].(*Alert)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlertByName indicates an expected call of GetAlertByName.
func (mr *MockAlertServiceIfaceMockRecorder) GetAlertByName(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertByName", reflect.TypeOf((*MockAlertServiceIface)(nil).GetAlertByName), varargs...)
}

// GetAlertID mocks base method.
func (m *MockAlertServiceIface) GetAlertID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAlertID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlertID indicates an expected call of GetAlertID.
func (mr *MockAlertServiceIfaceMockRecorder) GetAlertID(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertID", reflect.TypeOf((*MockAlertServiceIface)(nil).GetAlertID), varargs...)
}

// ListAlertTypes mocks base method.
func (m *MockAlertServiceIface) ListAlertTypes(p *ListAlertTypesParams) (*ListAlertTypesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlertTypes", p)
	ret0, _ := ret[0].(*ListAlertTypesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlertTypes indicates an expected call of ListAlertTypes.
func (mr *MockAlertServiceIfaceMockRecorder) ListAlertTypes(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlertTypes", reflect.TypeOf((*MockAlertServiceIface)(nil).ListAlertTypes), p)
}

// ListAlerts mocks base method.
func (m *MockAlertServiceIface) ListAlerts(p *ListAlertsParams) (*ListAlertsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlerts", p)
	ret0, _ := ret[0].(*ListAlertsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlerts indicates an expected call of ListAlerts.
func (mr *MockAlertServiceIfaceMockRecorder) ListAlerts(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlerts", reflect.TypeOf((*MockAlertServiceIface)(nil).ListAlerts), p)
}

// NewArchiveAlertsParams mocks base method.
func (m *MockAlertServiceIface) NewArchiveAlertsParams() *ArchiveAlertsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewArchiveAlertsParams")
	ret0, _ := ret[0].(*ArchiveAlertsParams)
	return ret0
}

// NewArchiveAlertsParams indicates an expected call of NewArchiveAlertsParams.
func (mr *MockAlertServiceIfaceMockRecorder) NewArchiveAlertsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewArchiveAlertsParams", reflect.TypeOf((*MockAlertServiceIface)(nil).NewArchiveAlertsParams))
}

// NewDeleteAlertsParams mocks base method.
func (m *MockAlertServiceIface) NewDeleteAlertsParams() *DeleteAlertsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteAlertsParams")
	ret0, _ := ret[0].(*DeleteAlertsParams)
	return ret0
}

// NewDeleteAlertsParams indicates an expected call of NewDeleteAlertsParams.
func (mr *MockAlertServiceIfaceMockRecorder) NewDeleteAlertsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteAlertsParams", reflect.TypeOf((*MockAlertServiceIface)(nil).NewDeleteAlertsParams))
}

// NewGenerateAlertParams mocks base method.
func (m *MockAlertServiceIface) NewGenerateAlertParams(description, name string, alertType int) *GenerateAlertParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGenerateAlertParams", description, name, alertType)
	ret0, _ := ret[0].(*GenerateAlertParams)
	return ret0
}

// NewGenerateAlertParams indicates an expected call of NewGenerateAlertParams.
func (mr *MockAlertServiceIfaceMockRecorder) NewGenerateAlertParams(description, name, alertType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGenerateAlertParams", reflect.TypeOf((*MockAlertServiceIface)(nil).NewGenerateAlertParams), description, name, alertType)
}

// NewListAlertTypesParams mocks base method.
func (m *MockAlertServiceIface) NewListAlertTypesParams() *ListAlertTypesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListAlertTypesParams")
	ret0, _ := ret[0].(*ListAlertTypesParams)
	return ret0
}

// NewListAlertTypesParams indicates an expected call of NewListAlertTypesParams.
func (mr *MockAlertServiceIfaceMockRecorder) NewListAlertTypesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListAlertTypesParams", reflect.TypeOf((*MockAlertServiceIface)(nil).NewListAlertTypesParams))
}

// NewListAlertsParams mocks base method.
func (m *MockAlertServiceIface) NewListAlertsParams() *ListAlertsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListAlertsParams")
	ret0, _ := ret[0].(*ListAlertsParams)
	return ret0
}

// NewListAlertsParams indicates an expected call of NewListAlertsParams.
func (mr *MockAlertServiceIfaceMockRecorder) NewListAlertsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListAlertsParams", reflect.TypeOf((*MockAlertServiceIface)(nil).NewListAlertsParams))
}
