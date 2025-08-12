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
	"net/url"
	"strconv"
)

type NiciraNVPServiceIface interface {
	AddNiciraNvpDevice(p *AddNiciraNvpDeviceParams) (*AddNiciraNvpDeviceResponse, error)
	NewAddNiciraNvpDeviceParams(hostname string, password string, physicalnetworkid string, transportzoneuuid string, username string) *AddNiciraNvpDeviceParams
	DeleteNiciraNvpDevice(p *DeleteNiciraNvpDeviceParams) (*DeleteNiciraNvpDeviceResponse, error)
	NewDeleteNiciraNvpDeviceParams(nvpdeviceid string) *DeleteNiciraNvpDeviceParams
	ListNiciraNvpDevices(p *ListNiciraNvpDevicesParams) (*ListNiciraNvpDevicesResponse, error)
	NewListNiciraNvpDevicesParams() *ListNiciraNvpDevicesParams
}

type AddNiciraNvpDeviceParams struct {
	p map[string]interface{}
}

func (p *AddNiciraNvpDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["l2gatewayserviceuuid"]; found {
		u.Set("l2gatewayserviceuuid", v.(string))
	}
	if v, found := p.p["l3gatewayserviceuuid"]; found {
		u.Set("l3gatewayserviceuuid", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["transportzoneuuid"]; found {
		u.Set("transportzoneuuid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddNiciraNvpDeviceParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *AddNiciraNvpDeviceParams) ResetHostname() {
	if p.p != nil && p.p["hostname"] != nil {
		delete(p.p, "hostname")
	}
}

func (p *AddNiciraNvpDeviceParams) GetHostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostname"].(string)
	return value, ok
}

func (p *AddNiciraNvpDeviceParams) SetL2gatewayserviceuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["l2gatewayserviceuuid"] = v
}

func (p *AddNiciraNvpDeviceParams) ResetL2gatewayserviceuuid() {
	if p.p != nil && p.p["l2gatewayserviceuuid"] != nil {
		delete(p.p, "l2gatewayserviceuuid")
	}
}

func (p *AddNiciraNvpDeviceParams) GetL2gatewayserviceuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["l2gatewayserviceuuid"].(string)
	return value, ok
}

func (p *AddNiciraNvpDeviceParams) SetL3gatewayserviceuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["l3gatewayserviceuuid"] = v
}

func (p *AddNiciraNvpDeviceParams) ResetL3gatewayserviceuuid() {
	if p.p != nil && p.p["l3gatewayserviceuuid"] != nil {
		delete(p.p, "l3gatewayserviceuuid")
	}
}

func (p *AddNiciraNvpDeviceParams) GetL3gatewayserviceuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["l3gatewayserviceuuid"].(string)
	return value, ok
}

func (p *AddNiciraNvpDeviceParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddNiciraNvpDeviceParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddNiciraNvpDeviceParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddNiciraNvpDeviceParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddNiciraNvpDeviceParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *AddNiciraNvpDeviceParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddNiciraNvpDeviceParams) SetTransportzoneuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["transportzoneuuid"] = v
}

func (p *AddNiciraNvpDeviceParams) ResetTransportzoneuuid() {
	if p.p != nil && p.p["transportzoneuuid"] != nil {
		delete(p.p, "transportzoneuuid")
	}
}

func (p *AddNiciraNvpDeviceParams) GetTransportzoneuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["transportzoneuuid"].(string)
	return value, ok
}

func (p *AddNiciraNvpDeviceParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddNiciraNvpDeviceParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddNiciraNvpDeviceParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddNiciraNvpDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewAddNiciraNvpDeviceParams(hostname string, password string, physicalnetworkid string, transportzoneuuid string, username string) *AddNiciraNvpDeviceParams {
	p := &AddNiciraNvpDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["hostname"] = hostname
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["transportzoneuuid"] = transportzoneuuid
	p.p["username"] = username
	return p
}

// Adds a Nicira NVP device
func (s *NiciraNVPService) AddNiciraNvpDevice(p *AddNiciraNvpDeviceParams) (*AddNiciraNvpDeviceResponse, error) {
	resp, err := s.cs.newPostRequest("addNiciraNvpDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNiciraNvpDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AddNiciraNvpDeviceResponse struct {
	Hostname             string `json:"hostname"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	L2gatewayserviceuuid string `json:"l2gatewayserviceuuid"`
	L3gatewayserviceuuid string `json:"l3gatewayserviceuuid"`
	Niciradevicename     string `json:"niciradevicename"`
	Nvpdeviceid          string `json:"nvpdeviceid"`
	Physicalnetworkid    string `json:"physicalnetworkid"`
	Provider             string `json:"provider"`
	Transportzoneuuid    string `json:"transportzoneuuid"`
}

type DeleteNiciraNvpDeviceParams struct {
	p map[string]interface{}
}

func (p *DeleteNiciraNvpDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
	}
	return u
}

func (p *DeleteNiciraNvpDeviceParams) SetNvpdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nvpdeviceid"] = v
}

func (p *DeleteNiciraNvpDeviceParams) ResetNvpdeviceid() {
	if p.p != nil && p.p["nvpdeviceid"] != nil {
		delete(p.p, "nvpdeviceid")
	}
}

func (p *DeleteNiciraNvpDeviceParams) GetNvpdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nvpdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNiciraNvpDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewDeleteNiciraNvpDeviceParams(nvpdeviceid string) *DeleteNiciraNvpDeviceParams {
	p := &DeleteNiciraNvpDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["nvpdeviceid"] = nvpdeviceid
	return p
}

// delete a nicira nvp device
func (s *NiciraNVPService) DeleteNiciraNvpDevice(p *DeleteNiciraNvpDeviceParams) (*DeleteNiciraNvpDeviceResponse, error) {
	resp, err := s.cs.newPostRequest("deleteNiciraNvpDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNiciraNvpDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteNiciraNvpDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListNiciraNvpDevicesParams struct {
	p map[string]interface{}
}

func (p *ListNiciraNvpDevicesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListNiciraNvpDevicesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNiciraNvpDevicesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNiciraNvpDevicesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNiciraNvpDevicesParams) SetNvpdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nvpdeviceid"] = v
}

func (p *ListNiciraNvpDevicesParams) ResetNvpdeviceid() {
	if p.p != nil && p.p["nvpdeviceid"] != nil {
		delete(p.p, "nvpdeviceid")
	}
}

func (p *ListNiciraNvpDevicesParams) GetNvpdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nvpdeviceid"].(string)
	return value, ok
}

func (p *ListNiciraNvpDevicesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNiciraNvpDevicesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNiciraNvpDevicesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNiciraNvpDevicesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNiciraNvpDevicesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNiciraNvpDevicesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNiciraNvpDevicesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListNiciraNvpDevicesParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *ListNiciraNvpDevicesParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNiciraNvpDevicesParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewListNiciraNvpDevicesParams() *ListNiciraNvpDevicesParams {
	p := &ListNiciraNvpDevicesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Nicira NVP devices
func (s *NiciraNVPService) ListNiciraNvpDevices(p *ListNiciraNvpDevicesParams) (*ListNiciraNvpDevicesResponse, error) {
	resp, err := s.cs.newRequest("listNiciraNvpDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNiciraNvpDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNiciraNvpDevicesResponse struct {
	Count            int                `json:"count"`
	NiciraNvpDevices []*NiciraNvpDevice `json:"niciranvpdevice"`
}

type NiciraNvpDevice struct {
	Hostname             string `json:"hostname"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	L2gatewayserviceuuid string `json:"l2gatewayserviceuuid"`
	L3gatewayserviceuuid string `json:"l3gatewayserviceuuid"`
	Niciradevicename     string `json:"niciradevicename"`
	Nvpdeviceid          string `json:"nvpdeviceid"`
	Physicalnetworkid    string `json:"physicalnetworkid"`
	Provider             string `json:"provider"`
	Transportzoneuuid    string `json:"transportzoneuuid"`
}
