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
// Source: ./cloudstack/VLANService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/VLANService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/VLANService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockVLANServiceIface is a mock of VLANServiceIface interface.
type MockVLANServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockVLANServiceIfaceMockRecorder
	isgomock struct{}
}

// MockVLANServiceIfaceMockRecorder is the mock recorder for MockVLANServiceIface.
type MockVLANServiceIfaceMockRecorder struct {
	mock *MockVLANServiceIface
}

// NewMockVLANServiceIface creates a new mock instance.
func NewMockVLANServiceIface(ctrl *gomock.Controller) *MockVLANServiceIface {
	mock := &MockVLANServiceIface{ctrl: ctrl}
	mock.recorder = &MockVLANServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVLANServiceIface) EXPECT() *MockVLANServiceIfaceMockRecorder {
	return m.recorder
}

// CreateVlanIpRange mocks base method.
func (m *MockVLANServiceIface) CreateVlanIpRange(p *CreateVlanIpRangeParams) (*CreateVlanIpRangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVlanIpRange", p)
	ret0, _ := ret[0].(*CreateVlanIpRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVlanIpRange indicates an expected call of CreateVlanIpRange.
func (mr *MockVLANServiceIfaceMockRecorder) CreateVlanIpRange(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVlanIpRange", reflect.TypeOf((*MockVLANServiceIface)(nil).CreateVlanIpRange), p)
}

// DedicateGuestVlanRange mocks base method.
func (m *MockVLANServiceIface) DedicateGuestVlanRange(p *DedicateGuestVlanRangeParams) (*DedicateGuestVlanRangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicateGuestVlanRange", p)
	ret0, _ := ret[0].(*DedicateGuestVlanRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DedicateGuestVlanRange indicates an expected call of DedicateGuestVlanRange.
func (mr *MockVLANServiceIfaceMockRecorder) DedicateGuestVlanRange(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicateGuestVlanRange", reflect.TypeOf((*MockVLANServiceIface)(nil).DedicateGuestVlanRange), p)
}

// DeleteVlanIpRange mocks base method.
func (m *MockVLANServiceIface) DeleteVlanIpRange(p *DeleteVlanIpRangeParams) (*DeleteVlanIpRangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVlanIpRange", p)
	ret0, _ := ret[0].(*DeleteVlanIpRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVlanIpRange indicates an expected call of DeleteVlanIpRange.
func (mr *MockVLANServiceIfaceMockRecorder) DeleteVlanIpRange(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVlanIpRange", reflect.TypeOf((*MockVLANServiceIface)(nil).DeleteVlanIpRange), p)
}

// GetDedicatedGuestVlanRangeByID mocks base method.
func (m *MockVLANServiceIface) GetDedicatedGuestVlanRangeByID(id string, opts ...OptionFunc) (*DedicatedGuestVlanRange, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDedicatedGuestVlanRangeByID", varargs...)
	ret0, _ := ret[0].(*DedicatedGuestVlanRange)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDedicatedGuestVlanRangeByID indicates an expected call of GetDedicatedGuestVlanRangeByID.
func (mr *MockVLANServiceIfaceMockRecorder) GetDedicatedGuestVlanRangeByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDedicatedGuestVlanRangeByID", reflect.TypeOf((*MockVLANServiceIface)(nil).GetDedicatedGuestVlanRangeByID), varargs...)
}

// GetVlanIpRangeByID mocks base method.
func (m *MockVLANServiceIface) GetVlanIpRangeByID(id string, opts ...OptionFunc) (*VlanIpRange, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVlanIpRangeByID", varargs...)
	ret0, _ := ret[0].(*VlanIpRange)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetVlanIpRangeByID indicates an expected call of GetVlanIpRangeByID.
func (mr *MockVLANServiceIfaceMockRecorder) GetVlanIpRangeByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVlanIpRangeByID", reflect.TypeOf((*MockVLANServiceIface)(nil).GetVlanIpRangeByID), varargs...)
}

// ListDedicatedGuestVlanRanges mocks base method.
func (m *MockVLANServiceIface) ListDedicatedGuestVlanRanges(p *ListDedicatedGuestVlanRangesParams) (*ListDedicatedGuestVlanRangesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDedicatedGuestVlanRanges", p)
	ret0, _ := ret[0].(*ListDedicatedGuestVlanRangesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDedicatedGuestVlanRanges indicates an expected call of ListDedicatedGuestVlanRanges.
func (mr *MockVLANServiceIfaceMockRecorder) ListDedicatedGuestVlanRanges(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDedicatedGuestVlanRanges", reflect.TypeOf((*MockVLANServiceIface)(nil).ListDedicatedGuestVlanRanges), p)
}

// ListGuestVlans mocks base method.
func (m *MockVLANServiceIface) ListGuestVlans(p *ListGuestVlansParams) (*ListGuestVlansResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGuestVlans", p)
	ret0, _ := ret[0].(*ListGuestVlansResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGuestVlans indicates an expected call of ListGuestVlans.
func (mr *MockVLANServiceIfaceMockRecorder) ListGuestVlans(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGuestVlans", reflect.TypeOf((*MockVLANServiceIface)(nil).ListGuestVlans), p)
}

// ListVlanIpRanges mocks base method.
func (m *MockVLANServiceIface) ListVlanIpRanges(p *ListVlanIpRangesParams) (*ListVlanIpRangesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVlanIpRanges", p)
	ret0, _ := ret[0].(*ListVlanIpRangesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVlanIpRanges indicates an expected call of ListVlanIpRanges.
func (mr *MockVLANServiceIfaceMockRecorder) ListVlanIpRanges(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVlanIpRanges", reflect.TypeOf((*MockVLANServiceIface)(nil).ListVlanIpRanges), p)
}

// NewCreateVlanIpRangeParams mocks base method.
func (m *MockVLANServiceIface) NewCreateVlanIpRangeParams() *CreateVlanIpRangeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateVlanIpRangeParams")
	ret0, _ := ret[0].(*CreateVlanIpRangeParams)
	return ret0
}

// NewCreateVlanIpRangeParams indicates an expected call of NewCreateVlanIpRangeParams.
func (mr *MockVLANServiceIfaceMockRecorder) NewCreateVlanIpRangeParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateVlanIpRangeParams", reflect.TypeOf((*MockVLANServiceIface)(nil).NewCreateVlanIpRangeParams))
}

// NewDedicateGuestVlanRangeParams mocks base method.
func (m *MockVLANServiceIface) NewDedicateGuestVlanRangeParams(physicalnetworkid, vlanrange string) *DedicateGuestVlanRangeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDedicateGuestVlanRangeParams", physicalnetworkid, vlanrange)
	ret0, _ := ret[0].(*DedicateGuestVlanRangeParams)
	return ret0
}

// NewDedicateGuestVlanRangeParams indicates an expected call of NewDedicateGuestVlanRangeParams.
func (mr *MockVLANServiceIfaceMockRecorder) NewDedicateGuestVlanRangeParams(physicalnetworkid, vlanrange any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDedicateGuestVlanRangeParams", reflect.TypeOf((*MockVLANServiceIface)(nil).NewDedicateGuestVlanRangeParams), physicalnetworkid, vlanrange)
}

// NewDeleteVlanIpRangeParams mocks base method.
func (m *MockVLANServiceIface) NewDeleteVlanIpRangeParams(id string) *DeleteVlanIpRangeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteVlanIpRangeParams", id)
	ret0, _ := ret[0].(*DeleteVlanIpRangeParams)
	return ret0
}

// NewDeleteVlanIpRangeParams indicates an expected call of NewDeleteVlanIpRangeParams.
func (mr *MockVLANServiceIfaceMockRecorder) NewDeleteVlanIpRangeParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteVlanIpRangeParams", reflect.TypeOf((*MockVLANServiceIface)(nil).NewDeleteVlanIpRangeParams), id)
}

// NewListDedicatedGuestVlanRangesParams mocks base method.
func (m *MockVLANServiceIface) NewListDedicatedGuestVlanRangesParams() *ListDedicatedGuestVlanRangesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListDedicatedGuestVlanRangesParams")
	ret0, _ := ret[0].(*ListDedicatedGuestVlanRangesParams)
	return ret0
}

// NewListDedicatedGuestVlanRangesParams indicates an expected call of NewListDedicatedGuestVlanRangesParams.
func (mr *MockVLANServiceIfaceMockRecorder) NewListDedicatedGuestVlanRangesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListDedicatedGuestVlanRangesParams", reflect.TypeOf((*MockVLANServiceIface)(nil).NewListDedicatedGuestVlanRangesParams))
}

// NewListGuestVlansParams mocks base method.
func (m *MockVLANServiceIface) NewListGuestVlansParams() *ListGuestVlansParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListGuestVlansParams")
	ret0, _ := ret[0].(*ListGuestVlansParams)
	return ret0
}

// NewListGuestVlansParams indicates an expected call of NewListGuestVlansParams.
func (mr *MockVLANServiceIfaceMockRecorder) NewListGuestVlansParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListGuestVlansParams", reflect.TypeOf((*MockVLANServiceIface)(nil).NewListGuestVlansParams))
}

// NewListVlanIpRangesParams mocks base method.
func (m *MockVLANServiceIface) NewListVlanIpRangesParams() *ListVlanIpRangesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListVlanIpRangesParams")
	ret0, _ := ret[0].(*ListVlanIpRangesParams)
	return ret0
}

// NewListVlanIpRangesParams indicates an expected call of NewListVlanIpRangesParams.
func (mr *MockVLANServiceIfaceMockRecorder) NewListVlanIpRangesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListVlanIpRangesParams", reflect.TypeOf((*MockVLANServiceIface)(nil).NewListVlanIpRangesParams))
}

// NewReleaseDedicatedGuestVlanRangeParams mocks base method.
func (m *MockVLANServiceIface) NewReleaseDedicatedGuestVlanRangeParams(id string) *ReleaseDedicatedGuestVlanRangeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReleaseDedicatedGuestVlanRangeParams", id)
	ret0, _ := ret[0].(*ReleaseDedicatedGuestVlanRangeParams)
	return ret0
}

// NewReleaseDedicatedGuestVlanRangeParams indicates an expected call of NewReleaseDedicatedGuestVlanRangeParams.
func (mr *MockVLANServiceIfaceMockRecorder) NewReleaseDedicatedGuestVlanRangeParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReleaseDedicatedGuestVlanRangeParams", reflect.TypeOf((*MockVLANServiceIface)(nil).NewReleaseDedicatedGuestVlanRangeParams), id)
}

// NewUpdateVlanIpRangeParams mocks base method.
func (m *MockVLANServiceIface) NewUpdateVlanIpRangeParams(id string) *UpdateVlanIpRangeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateVlanIpRangeParams", id)
	ret0, _ := ret[0].(*UpdateVlanIpRangeParams)
	return ret0
}

// NewUpdateVlanIpRangeParams indicates an expected call of NewUpdateVlanIpRangeParams.
func (mr *MockVLANServiceIfaceMockRecorder) NewUpdateVlanIpRangeParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateVlanIpRangeParams", reflect.TypeOf((*MockVLANServiceIface)(nil).NewUpdateVlanIpRangeParams), id)
}

// ReleaseDedicatedGuestVlanRange mocks base method.
func (m *MockVLANServiceIface) ReleaseDedicatedGuestVlanRange(p *ReleaseDedicatedGuestVlanRangeParams) (*ReleaseDedicatedGuestVlanRangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseDedicatedGuestVlanRange", p)
	ret0, _ := ret[0].(*ReleaseDedicatedGuestVlanRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseDedicatedGuestVlanRange indicates an expected call of ReleaseDedicatedGuestVlanRange.
func (mr *MockVLANServiceIfaceMockRecorder) ReleaseDedicatedGuestVlanRange(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseDedicatedGuestVlanRange", reflect.TypeOf((*MockVLANServiceIface)(nil).ReleaseDedicatedGuestVlanRange), p)
}

// UpdateVlanIpRange mocks base method.
func (m *MockVLANServiceIface) UpdateVlanIpRange(p *UpdateVlanIpRangeParams) (*UpdateVlanIpRangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVlanIpRange", p)
	ret0, _ := ret[0].(*UpdateVlanIpRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVlanIpRange indicates an expected call of UpdateVlanIpRange.
func (mr *MockVLANServiceIfaceMockRecorder) UpdateVlanIpRange(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVlanIpRange", reflect.TypeOf((*MockVLANServiceIface)(nil).UpdateVlanIpRange), p)
}
