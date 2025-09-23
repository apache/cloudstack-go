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

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type GPUServiceIface interface {
	CreateGpuCard(p *CreateGpuCardParams) (*CreateGpuCardResponse, error)
	NewCreateGpuCardParams(deviceid string, devicename string, name string, vendorid string, vendorname string) *CreateGpuCardParams
	CreateGpuDevice(p *CreateGpuDeviceParams) (*CreateGpuDeviceResponse, error)
	NewCreateGpuDeviceParams(busaddress string, gpucardid string, hostid string, vgpuprofileid string) *CreateGpuDeviceParams
	CreateVgpuProfile(p *CreateVgpuProfileParams) (*CreateVgpuProfileResponse, error)
	NewCreateVgpuProfileParams(gpucardid string, name string) *CreateVgpuProfileParams
	DeleteGpuCard(p *DeleteGpuCardParams) (*DeleteGpuCardResponse, error)
	NewDeleteGpuCardParams(id string) *DeleteGpuCardParams
	DeleteGpuDevice(p *DeleteGpuDeviceParams) (*DeleteGpuDeviceResponse, error)
	NewDeleteGpuDeviceParams(ids []string) *DeleteGpuDeviceParams
	DeleteVgpuProfile(p *DeleteVgpuProfileParams) (*DeleteVgpuProfileResponse, error)
	NewDeleteVgpuProfileParams(id string) *DeleteVgpuProfileParams
	DiscoverGpuDevices(p *DiscoverGpuDevicesParams) (*DiscoverGpuDevicesResponse, error)
	NewDiscoverGpuDevicesParams(id string) *DiscoverGpuDevicesParams
	ListGpuCards(p *ListGpuCardsParams) (*ListGpuCardsResponse, error)
	NewListGpuCardsParams() *ListGpuCardsParams
	GetGpuCardID(keyword string, opts ...OptionFunc) (string, int, error)
	GetGpuCardByName(name string, opts ...OptionFunc) (*GpuCard, int, error)
	GetGpuCardByID(id string, opts ...OptionFunc) (*GpuCard, int, error)
	ListGpuDevices(p *ListGpuDevicesParams) (*ListGpuDevicesResponse, error)
	NewListGpuDevicesParams() *ListGpuDevicesParams
	GetGpuDeviceByID(id string, opts ...OptionFunc) (*GpuDevice, int, error)
	ListVgpuProfiles(p *ListVgpuProfilesParams) (*ListVgpuProfilesResponse, error)
	NewListVgpuProfilesParams() *ListVgpuProfilesParams
	GetVgpuProfileID(name string, opts ...OptionFunc) (string, int, error)
	GetVgpuProfileByName(name string, opts ...OptionFunc) (*VgpuProfile, int, error)
	GetVgpuProfileByID(id string, opts ...OptionFunc) (*VgpuProfile, int, error)
	ManageGpuDevice(p *ManageGpuDeviceParams) (*ManageGpuDeviceResponse, error)
	NewManageGpuDeviceParams(ids []string) *ManageGpuDeviceParams
	UnmanageGpuDevice(p *UnmanageGpuDeviceParams) (*UnmanageGpuDeviceResponse, error)
	NewUnmanageGpuDeviceParams(ids []string) *UnmanageGpuDeviceParams
	UpdateGpuCard(p *UpdateGpuCardParams) (*UpdateGpuCardResponse, error)
	NewUpdateGpuCardParams(id string) *UpdateGpuCardParams
	UpdateGpuDevice(p *UpdateGpuDeviceParams) (*UpdateGpuDeviceResponse, error)
	NewUpdateGpuDeviceParams(id string) *UpdateGpuDeviceParams
	UpdateVgpuProfile(p *UpdateVgpuProfileParams) (*UpdateVgpuProfileResponse, error)
	NewUpdateVgpuProfileParams(id string) *UpdateVgpuProfileParams
}

type CreateGpuCardParams struct {
	p map[string]interface{}
}

func (p *CreateGpuCardParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["deviceid"]; found {
		u.Set("deviceid", v.(string))
	}
	if v, found := p.p["devicename"]; found {
		u.Set("devicename", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["vendorid"]; found {
		u.Set("vendorid", v.(string))
	}
	if v, found := p.p["vendorname"]; found {
		u.Set("vendorname", v.(string))
	}
	if v, found := p.p["videoram"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("videoram", vv)
	}
	return u
}

func (p *CreateGpuCardParams) SetDeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deviceid"] = v
}

func (p *CreateGpuCardParams) ResetDeviceid() {
	if p.p != nil && p.p["deviceid"] != nil {
		delete(p.p, "deviceid")
	}
}

func (p *CreateGpuCardParams) GetDeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deviceid"].(string)
	return value, ok
}

func (p *CreateGpuCardParams) SetDevicename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["devicename"] = v
}

func (p *CreateGpuCardParams) ResetDevicename() {
	if p.p != nil && p.p["devicename"] != nil {
		delete(p.p, "devicename")
	}
}

func (p *CreateGpuCardParams) GetDevicename() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["devicename"].(string)
	return value, ok
}

func (p *CreateGpuCardParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateGpuCardParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateGpuCardParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateGpuCardParams) SetVendorid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vendorid"] = v
}

func (p *CreateGpuCardParams) ResetVendorid() {
	if p.p != nil && p.p["vendorid"] != nil {
		delete(p.p, "vendorid")
	}
}

func (p *CreateGpuCardParams) GetVendorid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vendorid"].(string)
	return value, ok
}

func (p *CreateGpuCardParams) SetVendorname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vendorname"] = v
}

func (p *CreateGpuCardParams) ResetVendorname() {
	if p.p != nil && p.p["vendorname"] != nil {
		delete(p.p, "vendorname")
	}
}

func (p *CreateGpuCardParams) GetVendorname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vendorname"].(string)
	return value, ok
}

func (p *CreateGpuCardParams) SetVideoram(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["videoram"] = v
}

func (p *CreateGpuCardParams) ResetVideoram() {
	if p.p != nil && p.p["videoram"] != nil {
		delete(p.p, "videoram")
	}
}

func (p *CreateGpuCardParams) GetVideoram() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["videoram"].(int64)
	return value, ok
}

// You should always use this function to get a new CreateGpuCardParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewCreateGpuCardParams(deviceid string, devicename string, name string, vendorid string, vendorname string) *CreateGpuCardParams {
	p := &CreateGpuCardParams{}
	p.p = make(map[string]interface{})
	p.p["deviceid"] = deviceid
	p.p["devicename"] = devicename
	p.p["name"] = name
	p.p["vendorid"] = vendorid
	p.p["vendorname"] = vendorname
	return p
}

// Creates a GPU card definition in the system
func (s *GPUService) CreateGpuCard(p *CreateGpuCardParams) (*CreateGpuCardResponse, error) {
	resp, err := s.cs.newPostRequest("createGpuCard", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateGpuCardResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateGpuCardResponse struct {
	Deviceid   string `json:"deviceid"`
	Devicename string `json:"devicename"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Name       string `json:"name"`
	Vendorid   string `json:"vendorid"`
	Vendorname string `json:"vendorname"`
}

type CreateGpuDeviceParams struct {
	p map[string]interface{}
}

func (p *CreateGpuDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["busaddress"]; found {
		u.Set("busaddress", v.(string))
	}
	if v, found := p.p["gpucardid"]; found {
		u.Set("gpucardid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["numanode"]; found {
		u.Set("numanode", v.(string))
	}
	if v, found := p.p["parentgpudeviceid"]; found {
		u.Set("parentgpudeviceid", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["vgpuprofileid"]; found {
		u.Set("vgpuprofileid", v.(string))
	}
	return u
}

func (p *CreateGpuDeviceParams) SetBusaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["busaddress"] = v
}

func (p *CreateGpuDeviceParams) ResetBusaddress() {
	if p.p != nil && p.p["busaddress"] != nil {
		delete(p.p, "busaddress")
	}
}

func (p *CreateGpuDeviceParams) GetBusaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["busaddress"].(string)
	return value, ok
}

func (p *CreateGpuDeviceParams) SetGpucardid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gpucardid"] = v
}

func (p *CreateGpuDeviceParams) ResetGpucardid() {
	if p.p != nil && p.p["gpucardid"] != nil {
		delete(p.p, "gpucardid")
	}
}

func (p *CreateGpuDeviceParams) GetGpucardid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gpucardid"].(string)
	return value, ok
}

func (p *CreateGpuDeviceParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *CreateGpuDeviceParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *CreateGpuDeviceParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *CreateGpuDeviceParams) SetNumanode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["numanode"] = v
}

func (p *CreateGpuDeviceParams) ResetNumanode() {
	if p.p != nil && p.p["numanode"] != nil {
		delete(p.p, "numanode")
	}
}

func (p *CreateGpuDeviceParams) GetNumanode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["numanode"].(string)
	return value, ok
}

func (p *CreateGpuDeviceParams) SetParentgpudeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parentgpudeviceid"] = v
}

func (p *CreateGpuDeviceParams) ResetParentgpudeviceid() {
	if p.p != nil && p.p["parentgpudeviceid"] != nil {
		delete(p.p, "parentgpudeviceid")
	}
}

func (p *CreateGpuDeviceParams) GetParentgpudeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parentgpudeviceid"].(string)
	return value, ok
}

func (p *CreateGpuDeviceParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *CreateGpuDeviceParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *CreateGpuDeviceParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *CreateGpuDeviceParams) SetVgpuprofileid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vgpuprofileid"] = v
}

func (p *CreateGpuDeviceParams) ResetVgpuprofileid() {
	if p.p != nil && p.p["vgpuprofileid"] != nil {
		delete(p.p, "vgpuprofileid")
	}
}

func (p *CreateGpuDeviceParams) GetVgpuprofileid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vgpuprofileid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateGpuDeviceParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewCreateGpuDeviceParams(busaddress string, gpucardid string, hostid string, vgpuprofileid string) *CreateGpuDeviceParams {
	p := &CreateGpuDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["busaddress"] = busaddress
	p.p["gpucardid"] = gpucardid
	p.p["hostid"] = hostid
	p.p["vgpuprofileid"] = vgpuprofileid
	return p
}

// Creates a GPU device manually on a host
func (s *GPUService) CreateGpuDevice(p *CreateGpuDeviceParams) (*CreateGpuDeviceResponse, error) {
	resp, err := s.cs.newPostRequest("createGpuDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateGpuDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateGpuDeviceResponse struct {
	Busaddress         string `json:"busaddress"`
	Gpucardid          string `json:"gpucardid"`
	Gpucardname        string `json:"gpucardname"`
	Gpudevicetype      string `json:"gpudevicetype"`
	Hostid             string `json:"hostid"`
	Hostname           string `json:"hostname"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managedstate       string `json:"managedstate"`
	Numanode           string `json:"numanode"`
	Parentgpudeviceid  string `json:"parentgpudeviceid"`
	State              string `json:"state"`
	Vgpuprofileid      string `json:"vgpuprofileid"`
	Vgpuprofilename    string `json:"vgpuprofilename"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
	Vmstate            string `json:"vmstate"`
}

type CreateVgpuProfileParams struct {
	p map[string]interface{}
}

func (p *CreateVgpuProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["gpucardid"]; found {
		u.Set("gpucardid", v.(string))
	}
	if v, found := p.p["maxheads"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxheads", vv)
	}
	if v, found := p.p["maxresolutionx"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxresolutionx", vv)
	}
	if v, found := p.p["maxresolutiony"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxresolutiony", vv)
	}
	if v, found := p.p["maxvgpuperphysicalgpu"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxvgpuperphysicalgpu", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["videoram"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("videoram", vv)
	}
	return u
}

func (p *CreateVgpuProfileParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateVgpuProfileParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateVgpuProfileParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateVgpuProfileParams) SetGpucardid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gpucardid"] = v
}

func (p *CreateVgpuProfileParams) ResetGpucardid() {
	if p.p != nil && p.p["gpucardid"] != nil {
		delete(p.p, "gpucardid")
	}
}

func (p *CreateVgpuProfileParams) GetGpucardid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gpucardid"].(string)
	return value, ok
}

func (p *CreateVgpuProfileParams) SetMaxheads(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxheads"] = v
}

func (p *CreateVgpuProfileParams) ResetMaxheads() {
	if p.p != nil && p.p["maxheads"] != nil {
		delete(p.p, "maxheads")
	}
}

func (p *CreateVgpuProfileParams) GetMaxheads() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxheads"].(int64)
	return value, ok
}

func (p *CreateVgpuProfileParams) SetMaxresolutionx(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxresolutionx"] = v
}

func (p *CreateVgpuProfileParams) ResetMaxresolutionx() {
	if p.p != nil && p.p["maxresolutionx"] != nil {
		delete(p.p, "maxresolutionx")
	}
}

func (p *CreateVgpuProfileParams) GetMaxresolutionx() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxresolutionx"].(int64)
	return value, ok
}

func (p *CreateVgpuProfileParams) SetMaxresolutiony(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxresolutiony"] = v
}

func (p *CreateVgpuProfileParams) ResetMaxresolutiony() {
	if p.p != nil && p.p["maxresolutiony"] != nil {
		delete(p.p, "maxresolutiony")
	}
}

func (p *CreateVgpuProfileParams) GetMaxresolutiony() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxresolutiony"].(int64)
	return value, ok
}

func (p *CreateVgpuProfileParams) SetMaxvgpuperphysicalgpu(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxvgpuperphysicalgpu"] = v
}

func (p *CreateVgpuProfileParams) ResetMaxvgpuperphysicalgpu() {
	if p.p != nil && p.p["maxvgpuperphysicalgpu"] != nil {
		delete(p.p, "maxvgpuperphysicalgpu")
	}
}

func (p *CreateVgpuProfileParams) GetMaxvgpuperphysicalgpu() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxvgpuperphysicalgpu"].(int64)
	return value, ok
}

func (p *CreateVgpuProfileParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVgpuProfileParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateVgpuProfileParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateVgpuProfileParams) SetVideoram(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["videoram"] = v
}

func (p *CreateVgpuProfileParams) ResetVideoram() {
	if p.p != nil && p.p["videoram"] != nil {
		delete(p.p, "videoram")
	}
}

func (p *CreateVgpuProfileParams) GetVideoram() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["videoram"].(int64)
	return value, ok
}

// You should always use this function to get a new CreateVgpuProfileParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewCreateVgpuProfileParams(gpucardid string, name string) *CreateVgpuProfileParams {
	p := &CreateVgpuProfileParams{}
	p.p = make(map[string]interface{})
	p.p["gpucardid"] = gpucardid
	p.p["name"] = name
	return p
}

// Creates a vGPU profile in the system
func (s *GPUService) CreateVgpuProfile(p *CreateVgpuProfileParams) (*CreateVgpuProfileResponse, error) {
	resp, err := s.cs.newPostRequest("createVgpuProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVgpuProfileResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateVgpuProfileResponse struct {
	Description           string `json:"description"`
	Deviceid              string `json:"deviceid"`
	Devicename            string `json:"devicename"`
	Gpucardid             string `json:"gpucardid"`
	Gpucardname           string `json:"gpucardname"`
	Id                    string `json:"id"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Maxheads              int64  `json:"maxheads"`
	Maxresolutionx        int64  `json:"maxresolutionx"`
	Maxresolutiony        int64  `json:"maxresolutiony"`
	Maxvgpuperphysicalgpu int64  `json:"maxvgpuperphysicalgpu"`
	Name                  string `json:"name"`
	Vendorid              string `json:"vendorid"`
	Vendorname            string `json:"vendorname"`
	Videoram              int64  `json:"videoram"`
}

type DeleteGpuCardParams struct {
	p map[string]interface{}
}

func (p *DeleteGpuCardParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteGpuCardParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteGpuCardParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteGpuCardParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteGpuCardParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewDeleteGpuCardParams(id string) *DeleteGpuCardParams {
	p := &DeleteGpuCardParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a GPU card definition from the system
func (s *GPUService) DeleteGpuCard(p *DeleteGpuCardParams) (*DeleteGpuCardResponse, error) {
	resp, err := s.cs.newPostRequest("deleteGpuCard", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteGpuCardResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteGpuCardResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteGpuCardResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteGpuCardResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteGpuDeviceParams struct {
	p map[string]interface{}
}

func (p *DeleteGpuDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	return u
}

func (p *DeleteGpuDeviceParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *DeleteGpuDeviceParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *DeleteGpuDeviceParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

// You should always use this function to get a new DeleteGpuDeviceParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewDeleteGpuDeviceParams(ids []string) *DeleteGpuDeviceParams {
	p := &DeleteGpuDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["ids"] = ids
	return p
}

// Deletes a vGPU profile from the system
func (s *GPUService) DeleteGpuDevice(p *DeleteGpuDeviceParams) (*DeleteGpuDeviceResponse, error) {
	resp, err := s.cs.newPostRequest("deleteGpuDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteGpuDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteGpuDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteGpuDeviceResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteGpuDeviceResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteVgpuProfileParams struct {
	p map[string]interface{}
}

func (p *DeleteVgpuProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVgpuProfileParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVgpuProfileParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteVgpuProfileParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVgpuProfileParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewDeleteVgpuProfileParams(id string) *DeleteVgpuProfileParams {
	p := &DeleteVgpuProfileParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a vGPU profile from the system
func (s *GPUService) DeleteVgpuProfile(p *DeleteVgpuProfileParams) (*DeleteVgpuProfileResponse, error) {
	resp, err := s.cs.newPostRequest("deleteVgpuProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVgpuProfileResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteVgpuProfileResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteVgpuProfileResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteVgpuProfileResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DiscoverGpuDevicesParams struct {
	p map[string]interface{}
}

func (p *DiscoverGpuDevicesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *DiscoverGpuDevicesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DiscoverGpuDevicesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DiscoverGpuDevicesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DiscoverGpuDevicesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *DiscoverGpuDevicesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *DiscoverGpuDevicesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *DiscoverGpuDevicesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *DiscoverGpuDevicesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *DiscoverGpuDevicesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *DiscoverGpuDevicesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *DiscoverGpuDevicesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *DiscoverGpuDevicesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new DiscoverGpuDevicesParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewDiscoverGpuDevicesParams(id string) *DiscoverGpuDevicesParams {
	p := &DiscoverGpuDevicesParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Discovers available GPU devices on a host
func (s *GPUService) DiscoverGpuDevices(p *DiscoverGpuDevicesParams) (*DiscoverGpuDevicesResponse, error) {
	resp, err := s.cs.newPostRequest("discoverGpuDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DiscoverGpuDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DiscoverGpuDevicesResponse struct {
	Busaddress         string `json:"busaddress"`
	Gpucardid          string `json:"gpucardid"`
	Gpucardname        string `json:"gpucardname"`
	Gpudevicetype      string `json:"gpudevicetype"`
	Hostid             string `json:"hostid"`
	Hostname           string `json:"hostname"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managedstate       string `json:"managedstate"`
	Numanode           string `json:"numanode"`
	Parentgpudeviceid  string `json:"parentgpudeviceid"`
	State              string `json:"state"`
	Vgpuprofileid      string `json:"vgpuprofileid"`
	Vgpuprofilename    string `json:"vgpuprofilename"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
	Vmstate            string `json:"vmstate"`
}

type ListGpuCardsParams struct {
	p map[string]interface{}
}

func (p *ListGpuCardsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["activeonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("activeonly", vv)
	}
	if v, found := p.p["deviceid"]; found {
		u.Set("deviceid", v.(string))
	}
	if v, found := p.p["devicename"]; found {
		u.Set("devicename", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["vendorid"]; found {
		u.Set("vendorid", v.(string))
	}
	if v, found := p.p["vendorname"]; found {
		u.Set("vendorname", v.(string))
	}
	return u
}

func (p *ListGpuCardsParams) SetActiveonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["activeonly"] = v
}

func (p *ListGpuCardsParams) ResetActiveonly() {
	if p.p != nil && p.p["activeonly"] != nil {
		delete(p.p, "activeonly")
	}
}

func (p *ListGpuCardsParams) GetActiveonly() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["activeonly"].(bool)
	return value, ok
}

func (p *ListGpuCardsParams) SetDeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deviceid"] = v
}

func (p *ListGpuCardsParams) ResetDeviceid() {
	if p.p != nil && p.p["deviceid"] != nil {
		delete(p.p, "deviceid")
	}
}

func (p *ListGpuCardsParams) GetDeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deviceid"].(string)
	return value, ok
}

func (p *ListGpuCardsParams) SetDevicename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["devicename"] = v
}

func (p *ListGpuCardsParams) ResetDevicename() {
	if p.p != nil && p.p["devicename"] != nil {
		delete(p.p, "devicename")
	}
}

func (p *ListGpuCardsParams) GetDevicename() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["devicename"].(string)
	return value, ok
}

func (p *ListGpuCardsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListGpuCardsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListGpuCardsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListGpuCardsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListGpuCardsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListGpuCardsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListGpuCardsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListGpuCardsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListGpuCardsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListGpuCardsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListGpuCardsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListGpuCardsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListGpuCardsParams) SetVendorid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vendorid"] = v
}

func (p *ListGpuCardsParams) ResetVendorid() {
	if p.p != nil && p.p["vendorid"] != nil {
		delete(p.p, "vendorid")
	}
}

func (p *ListGpuCardsParams) GetVendorid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vendorid"].(string)
	return value, ok
}

func (p *ListGpuCardsParams) SetVendorname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vendorname"] = v
}

func (p *ListGpuCardsParams) ResetVendorname() {
	if p.p != nil && p.p["vendorname"] != nil {
		delete(p.p, "vendorname")
	}
}

func (p *ListGpuCardsParams) GetVendorname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vendorname"].(string)
	return value, ok
}

// You should always use this function to get a new ListGpuCardsParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewListGpuCardsParams() *ListGpuCardsParams {
	p := &ListGpuCardsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GPUService) GetGpuCardID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListGpuCardsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListGpuCards(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.GpuCards[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.GpuCards {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GPUService) GetGpuCardByName(name string, opts ...OptionFunc) (*GpuCard, int, error) {
	id, count, err := s.GetGpuCardID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetGpuCardByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GPUService) GetGpuCardByID(id string, opts ...OptionFunc) (*GpuCard, int, error) {
	p := &ListGpuCardsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListGpuCards(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.GpuCards[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GpuCard UUID: %s!", id)
}

// Lists all available GPU cards
func (s *GPUService) ListGpuCards(p *ListGpuCardsParams) (*ListGpuCardsResponse, error) {
	resp, err := s.cs.newRequest("listGpuCards", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGpuCardsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListGpuCardsResponse struct {
	Count    int        `json:"count"`
	GpuCards []*GpuCard `json:"gpucard"`
}

type GpuCard struct {
	Deviceid   string `json:"deviceid"`
	Devicename string `json:"devicename"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Name       string `json:"name"`
	Vendorid   string `json:"vendorid"`
	Vendorname string `json:"vendorname"`
}

type ListGpuDevicesParams struct {
	p map[string]interface{}
}

func (p *ListGpuDevicesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["gpucardid"]; found {
		u.Set("gpucardid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["vgpuprofileid"]; found {
		u.Set("vgpuprofileid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListGpuDevicesParams) SetGpucardid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gpucardid"] = v
}

func (p *ListGpuDevicesParams) ResetGpucardid() {
	if p.p != nil && p.p["gpucardid"] != nil {
		delete(p.p, "gpucardid")
	}
}

func (p *ListGpuDevicesParams) GetGpucardid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gpucardid"].(string)
	return value, ok
}

func (p *ListGpuDevicesParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListGpuDevicesParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ListGpuDevicesParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListGpuDevicesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListGpuDevicesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListGpuDevicesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListGpuDevicesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListGpuDevicesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListGpuDevicesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListGpuDevicesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListGpuDevicesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListGpuDevicesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListGpuDevicesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListGpuDevicesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListGpuDevicesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListGpuDevicesParams) SetVgpuprofileid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vgpuprofileid"] = v
}

func (p *ListGpuDevicesParams) ResetVgpuprofileid() {
	if p.p != nil && p.p["vgpuprofileid"] != nil {
		delete(p.p, "vgpuprofileid")
	}
}

func (p *ListGpuDevicesParams) GetVgpuprofileid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vgpuprofileid"].(string)
	return value, ok
}

func (p *ListGpuDevicesParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListGpuDevicesParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *ListGpuDevicesParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new ListGpuDevicesParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewListGpuDevicesParams() *ListGpuDevicesParams {
	p := &ListGpuDevicesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GPUService) GetGpuDeviceByID(id string, opts ...OptionFunc) (*GpuDevice, int, error) {
	p := &ListGpuDevicesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListGpuDevices(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.GpuDevices[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GpuDevice UUID: %s!", id)
}

// Lists all available GPU devices
func (s *GPUService) ListGpuDevices(p *ListGpuDevicesParams) (*ListGpuDevicesResponse, error) {
	resp, err := s.cs.newRequest("listGpuDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGpuDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListGpuDevicesResponse struct {
	Count      int          `json:"count"`
	GpuDevices []*GpuDevice `json:"gpudevice"`
}

type GpuDevice struct {
	Busaddress         string `json:"busaddress"`
	Gpucardid          string `json:"gpucardid"`
	Gpucardname        string `json:"gpucardname"`
	Gpudevicetype      string `json:"gpudevicetype"`
	Hostid             string `json:"hostid"`
	Hostname           string `json:"hostname"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managedstate       string `json:"managedstate"`
	Numanode           string `json:"numanode"`
	Parentgpudeviceid  string `json:"parentgpudeviceid"`
	State              string `json:"state"`
	Vgpuprofileid      string `json:"vgpuprofileid"`
	Vgpuprofilename    string `json:"vgpuprofilename"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
	Vmstate            string `json:"vmstate"`
}

type ListVgpuProfilesParams struct {
	p map[string]interface{}
}

func (p *ListVgpuProfilesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["activeonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("activeonly", vv)
	}
	if v, found := p.p["gpucardid"]; found {
		u.Set("gpucardid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListVgpuProfilesParams) SetActiveonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["activeonly"] = v
}

func (p *ListVgpuProfilesParams) ResetActiveonly() {
	if p.p != nil && p.p["activeonly"] != nil {
		delete(p.p, "activeonly")
	}
}

func (p *ListVgpuProfilesParams) GetActiveonly() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["activeonly"].(bool)
	return value, ok
}

func (p *ListVgpuProfilesParams) SetGpucardid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gpucardid"] = v
}

func (p *ListVgpuProfilesParams) ResetGpucardid() {
	if p.p != nil && p.p["gpucardid"] != nil {
		delete(p.p, "gpucardid")
	}
}

func (p *ListVgpuProfilesParams) GetGpucardid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gpucardid"].(string)
	return value, ok
}

func (p *ListVgpuProfilesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVgpuProfilesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVgpuProfilesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVgpuProfilesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVgpuProfilesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVgpuProfilesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVgpuProfilesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVgpuProfilesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVgpuProfilesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVgpuProfilesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVgpuProfilesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVgpuProfilesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVgpuProfilesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVgpuProfilesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVgpuProfilesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListVgpuProfilesParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewListVgpuProfilesParams() *ListVgpuProfilesParams {
	p := &ListVgpuProfilesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GPUService) GetVgpuProfileID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVgpuProfilesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVgpuProfiles(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VgpuProfiles[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VgpuProfiles {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GPUService) GetVgpuProfileByName(name string, opts ...OptionFunc) (*VgpuProfile, int, error) {
	id, count, err := s.GetVgpuProfileID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVgpuProfileByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GPUService) GetVgpuProfileByID(id string, opts ...OptionFunc) (*VgpuProfile, int, error) {
	p := &ListVgpuProfilesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVgpuProfiles(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.VgpuProfiles[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VgpuProfile UUID: %s!", id)
}

// Lists all available vGPU profiles
func (s *GPUService) ListVgpuProfiles(p *ListVgpuProfilesParams) (*ListVgpuProfilesResponse, error) {
	resp, err := s.cs.newRequest("listVgpuProfiles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVgpuProfilesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVgpuProfilesResponse struct {
	Count        int            `json:"count"`
	VgpuProfiles []*VgpuProfile `json:"vgpuprofile"`
}

type VgpuProfile struct {
	Description           string `json:"description"`
	Deviceid              string `json:"deviceid"`
	Devicename            string `json:"devicename"`
	Gpucardid             string `json:"gpucardid"`
	Gpucardname           string `json:"gpucardname"`
	Id                    string `json:"id"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Maxheads              int64  `json:"maxheads"`
	Maxresolutionx        int64  `json:"maxresolutionx"`
	Maxresolutiony        int64  `json:"maxresolutiony"`
	Maxvgpuperphysicalgpu int64  `json:"maxvgpuperphysicalgpu"`
	Name                  string `json:"name"`
	Vendorid              string `json:"vendorid"`
	Vendorname            string `json:"vendorname"`
	Videoram              int64  `json:"videoram"`
}

type ManageGpuDeviceParams struct {
	p map[string]interface{}
}

func (p *ManageGpuDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	return u
}

func (p *ManageGpuDeviceParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ManageGpuDeviceParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ManageGpuDeviceParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

// You should always use this function to get a new ManageGpuDeviceParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewManageGpuDeviceParams(ids []string) *ManageGpuDeviceParams {
	p := &ManageGpuDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["ids"] = ids
	return p
}

// Manages a GPU device
func (s *GPUService) ManageGpuDevice(p *ManageGpuDeviceParams) (*ManageGpuDeviceResponse, error) {
	resp, err := s.cs.newPostRequest("manageGpuDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ManageGpuDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ManageGpuDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ManageGpuDeviceResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias ManageGpuDeviceResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UnmanageGpuDeviceParams struct {
	p map[string]interface{}
}

func (p *UnmanageGpuDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	return u
}

func (p *UnmanageGpuDeviceParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *UnmanageGpuDeviceParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *UnmanageGpuDeviceParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

// You should always use this function to get a new UnmanageGpuDeviceParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewUnmanageGpuDeviceParams(ids []string) *UnmanageGpuDeviceParams {
	p := &UnmanageGpuDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["ids"] = ids
	return p
}

// Unmanage a GPU device
func (s *GPUService) UnmanageGpuDevice(p *UnmanageGpuDeviceParams) (*UnmanageGpuDeviceResponse, error) {
	resp, err := s.cs.newPostRequest("unmanageGpuDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UnmanageGpuDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UnmanageGpuDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UnmanageGpuDeviceResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias UnmanageGpuDeviceResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateGpuCardParams struct {
	p map[string]interface{}
}

func (p *UpdateGpuCardParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["devicename"]; found {
		u.Set("devicename", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["vendorname"]; found {
		u.Set("vendorname", v.(string))
	}
	return u
}

func (p *UpdateGpuCardParams) SetDevicename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["devicename"] = v
}

func (p *UpdateGpuCardParams) ResetDevicename() {
	if p.p != nil && p.p["devicename"] != nil {
		delete(p.p, "devicename")
	}
}

func (p *UpdateGpuCardParams) GetDevicename() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["devicename"].(string)
	return value, ok
}

func (p *UpdateGpuCardParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateGpuCardParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateGpuCardParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateGpuCardParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateGpuCardParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateGpuCardParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateGpuCardParams) SetVendorname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vendorname"] = v
}

func (p *UpdateGpuCardParams) ResetVendorname() {
	if p.p != nil && p.p["vendorname"] != nil {
		delete(p.p, "vendorname")
	}
}

func (p *UpdateGpuCardParams) GetVendorname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vendorname"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateGpuCardParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewUpdateGpuCardParams(id string) *UpdateGpuCardParams {
	p := &UpdateGpuCardParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a GPU card definition in the system
func (s *GPUService) UpdateGpuCard(p *UpdateGpuCardParams) (*UpdateGpuCardResponse, error) {
	resp, err := s.cs.newPostRequest("updateGpuCard", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGpuCardResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateGpuCardResponse struct {
	Deviceid   string `json:"deviceid"`
	Devicename string `json:"devicename"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Name       string `json:"name"`
	Vendorid   string `json:"vendorid"`
	Vendorname string `json:"vendorname"`
}

type UpdateGpuDeviceParams struct {
	p map[string]interface{}
}

func (p *UpdateGpuDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["gpucardid"]; found {
		u.Set("gpucardid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["numanode"]; found {
		u.Set("numanode", v.(string))
	}
	if v, found := p.p["parentgpudeviceid"]; found {
		u.Set("parentgpudeviceid", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["vgpuprofileid"]; found {
		u.Set("vgpuprofileid", v.(string))
	}
	return u
}

func (p *UpdateGpuDeviceParams) SetGpucardid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gpucardid"] = v
}

func (p *UpdateGpuDeviceParams) ResetGpucardid() {
	if p.p != nil && p.p["gpucardid"] != nil {
		delete(p.p, "gpucardid")
	}
}

func (p *UpdateGpuDeviceParams) GetGpucardid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gpucardid"].(string)
	return value, ok
}

func (p *UpdateGpuDeviceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateGpuDeviceParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateGpuDeviceParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateGpuDeviceParams) SetNumanode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["numanode"] = v
}

func (p *UpdateGpuDeviceParams) ResetNumanode() {
	if p.p != nil && p.p["numanode"] != nil {
		delete(p.p, "numanode")
	}
}

func (p *UpdateGpuDeviceParams) GetNumanode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["numanode"].(string)
	return value, ok
}

func (p *UpdateGpuDeviceParams) SetParentgpudeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parentgpudeviceid"] = v
}

func (p *UpdateGpuDeviceParams) ResetParentgpudeviceid() {
	if p.p != nil && p.p["parentgpudeviceid"] != nil {
		delete(p.p, "parentgpudeviceid")
	}
}

func (p *UpdateGpuDeviceParams) GetParentgpudeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parentgpudeviceid"].(string)
	return value, ok
}

func (p *UpdateGpuDeviceParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *UpdateGpuDeviceParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *UpdateGpuDeviceParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *UpdateGpuDeviceParams) SetVgpuprofileid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vgpuprofileid"] = v
}

func (p *UpdateGpuDeviceParams) ResetVgpuprofileid() {
	if p.p != nil && p.p["vgpuprofileid"] != nil {
		delete(p.p, "vgpuprofileid")
	}
}

func (p *UpdateGpuDeviceParams) GetVgpuprofileid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vgpuprofileid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateGpuDeviceParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewUpdateGpuDeviceParams(id string) *UpdateGpuDeviceParams {
	p := &UpdateGpuDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing GPU device
func (s *GPUService) UpdateGpuDevice(p *UpdateGpuDeviceParams) (*UpdateGpuDeviceResponse, error) {
	resp, err := s.cs.newPostRequest("updateGpuDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGpuDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateGpuDeviceResponse struct {
	Busaddress         string `json:"busaddress"`
	Gpucardid          string `json:"gpucardid"`
	Gpucardname        string `json:"gpucardname"`
	Gpudevicetype      string `json:"gpudevicetype"`
	Hostid             string `json:"hostid"`
	Hostname           string `json:"hostname"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managedstate       string `json:"managedstate"`
	Numanode           string `json:"numanode"`
	Parentgpudeviceid  string `json:"parentgpudeviceid"`
	State              string `json:"state"`
	Vgpuprofileid      string `json:"vgpuprofileid"`
	Vgpuprofilename    string `json:"vgpuprofilename"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
	Vmstate            string `json:"vmstate"`
}

type UpdateVgpuProfileParams struct {
	p map[string]interface{}
}

func (p *UpdateVgpuProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["maxheads"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxheads", vv)
	}
	if v, found := p.p["maxresolutionx"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxresolutionx", vv)
	}
	if v, found := p.p["maxresolutiony"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxresolutiony", vv)
	}
	if v, found := p.p["maxvgpuperphysicalgpu"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxvgpuperphysicalgpu", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["videoram"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("videoram", vv)
	}
	return u
}

func (p *UpdateVgpuProfileParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateVgpuProfileParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateVgpuProfileParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateVgpuProfileParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVgpuProfileParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVgpuProfileParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateVgpuProfileParams) SetMaxheads(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxheads"] = v
}

func (p *UpdateVgpuProfileParams) ResetMaxheads() {
	if p.p != nil && p.p["maxheads"] != nil {
		delete(p.p, "maxheads")
	}
}

func (p *UpdateVgpuProfileParams) GetMaxheads() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxheads"].(int64)
	return value, ok
}

func (p *UpdateVgpuProfileParams) SetMaxresolutionx(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxresolutionx"] = v
}

func (p *UpdateVgpuProfileParams) ResetMaxresolutionx() {
	if p.p != nil && p.p["maxresolutionx"] != nil {
		delete(p.p, "maxresolutionx")
	}
}

func (p *UpdateVgpuProfileParams) GetMaxresolutionx() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxresolutionx"].(int64)
	return value, ok
}

func (p *UpdateVgpuProfileParams) SetMaxresolutiony(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxresolutiony"] = v
}

func (p *UpdateVgpuProfileParams) ResetMaxresolutiony() {
	if p.p != nil && p.p["maxresolutiony"] != nil {
		delete(p.p, "maxresolutiony")
	}
}

func (p *UpdateVgpuProfileParams) GetMaxresolutiony() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxresolutiony"].(int64)
	return value, ok
}

func (p *UpdateVgpuProfileParams) SetMaxvgpuperphysicalgpu(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxvgpuperphysicalgpu"] = v
}

func (p *UpdateVgpuProfileParams) ResetMaxvgpuperphysicalgpu() {
	if p.p != nil && p.p["maxvgpuperphysicalgpu"] != nil {
		delete(p.p, "maxvgpuperphysicalgpu")
	}
}

func (p *UpdateVgpuProfileParams) GetMaxvgpuperphysicalgpu() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxvgpuperphysicalgpu"].(int64)
	return value, ok
}

func (p *UpdateVgpuProfileParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVgpuProfileParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateVgpuProfileParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateVgpuProfileParams) SetVideoram(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["videoram"] = v
}

func (p *UpdateVgpuProfileParams) ResetVideoram() {
	if p.p != nil && p.p["videoram"] != nil {
		delete(p.p, "videoram")
	}
}

func (p *UpdateVgpuProfileParams) GetVideoram() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["videoram"].(int64)
	return value, ok
}

// You should always use this function to get a new UpdateVgpuProfileParams instance,
// as then you are sure you have configured all required params
func (s *GPUService) NewUpdateVgpuProfileParams(id string) *UpdateVgpuProfileParams {
	p := &UpdateVgpuProfileParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a vGPU profile in the system
func (s *GPUService) UpdateVgpuProfile(p *UpdateVgpuProfileParams) (*UpdateVgpuProfileResponse, error) {
	resp, err := s.cs.newPostRequest("updateVgpuProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVgpuProfileResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateVgpuProfileResponse struct {
	Description           string `json:"description"`
	Deviceid              string `json:"deviceid"`
	Devicename            string `json:"devicename"`
	Gpucardid             string `json:"gpucardid"`
	Gpucardname           string `json:"gpucardname"`
	Id                    string `json:"id"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Maxheads              int64  `json:"maxheads"`
	Maxresolutionx        int64  `json:"maxresolutionx"`
	Maxresolutiony        int64  `json:"maxresolutiony"`
	Maxvgpuperphysicalgpu int64  `json:"maxvgpuperphysicalgpu"`
	Name                  string `json:"name"`
	Vendorid              string `json:"vendorid"`
	Vendorname            string `json:"vendorname"`
	Videoram              int64  `json:"videoram"`
}
