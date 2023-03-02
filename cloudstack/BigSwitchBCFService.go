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

type BigSwitchBCFServiceIface interface {
	AddBigSwitchBcfDevice(p *AddBigSwitchBcfDeviceParams) (*AddBigSwitchBcfDeviceResponse, error)
	NewAddBigSwitchBcfDeviceParams(hostname string, nat bool, password string, physicalnetworkid string, username string) *AddBigSwitchBcfDeviceParams
	DeleteBigSwitchBcfDevice(p *DeleteBigSwitchBcfDeviceParams) (*DeleteBigSwitchBcfDeviceResponse, error)
	NewDeleteBigSwitchBcfDeviceParams(bcfdeviceid string) *DeleteBigSwitchBcfDeviceParams
	ListBigSwitchBcfDevices(p *ListBigSwitchBcfDevicesParams) (*ListBigSwitchBcfDevicesResponse, error)
	NewListBigSwitchBcfDevicesParams() *ListBigSwitchBcfDevicesParams
}

type AddBigSwitchBcfDeviceParams struct {
	p map[string]interface{}
}

func (p *AddBigSwitchBcfDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["nat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("nat", vv)
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddBigSwitchBcfDeviceParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *AddBigSwitchBcfDeviceParams) GetHostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostname"].(string)
	return value, ok
}

func (p *AddBigSwitchBcfDeviceParams) SetNat(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nat"] = v
}

func (p *AddBigSwitchBcfDeviceParams) GetNat() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nat"].(bool)
	return value, ok
}

func (p *AddBigSwitchBcfDeviceParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddBigSwitchBcfDeviceParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddBigSwitchBcfDeviceParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddBigSwitchBcfDeviceParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddBigSwitchBcfDeviceParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddBigSwitchBcfDeviceParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBigSwitchBcfDeviceParams instance,
// as then you are sure you have configured all required params
func (s *BigSwitchBCFService) NewAddBigSwitchBcfDeviceParams(hostname string, nat bool, password string, physicalnetworkid string, username string) *AddBigSwitchBcfDeviceParams {
	p := &AddBigSwitchBcfDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["hostname"] = hostname
	p.p["nat"] = nat
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["username"] = username
	return p
}

// Adds a BigSwitch BCF Controller device
func (s *BigSwitchBCFService) AddBigSwitchBcfDevice(p *AddBigSwitchBcfDeviceParams) (*AddBigSwitchBcfDeviceResponse, error) {
	resp, err := s.cs.newRequest("addBigSwitchBcfDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBigSwitchBcfDeviceResponse
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

type AddBigSwitchBcfDeviceResponse struct {
	Bcfdeviceid         string `json:"bcfdeviceid"`
	Bigswitchdevicename string `json:"bigswitchdevicename"`
	Hostname            string `json:"hostname"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Nat                 bool   `json:"nat"`
	Password            string `json:"password"`
	Physicalnetworkid   string `json:"physicalnetworkid"`
	Provider            string `json:"provider"`
	Username            string `json:"username"`
}

type DeleteBigSwitchBcfDeviceParams struct {
	p map[string]interface{}
}

func (p *DeleteBigSwitchBcfDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bcfdeviceid"]; found {
		u.Set("bcfdeviceid", v.(string))
	}
	return u
}

func (p *DeleteBigSwitchBcfDeviceParams) SetBcfdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bcfdeviceid"] = v
}

func (p *DeleteBigSwitchBcfDeviceParams) GetBcfdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bcfdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBigSwitchBcfDeviceParams instance,
// as then you are sure you have configured all required params
func (s *BigSwitchBCFService) NewDeleteBigSwitchBcfDeviceParams(bcfdeviceid string) *DeleteBigSwitchBcfDeviceParams {
	p := &DeleteBigSwitchBcfDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["bcfdeviceid"] = bcfdeviceid
	return p
}

// delete a BigSwitch BCF Controller device
func (s *BigSwitchBCFService) DeleteBigSwitchBcfDevice(p *DeleteBigSwitchBcfDeviceParams) (*DeleteBigSwitchBcfDeviceResponse, error) {
	resp, err := s.cs.newRequest("deleteBigSwitchBcfDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBigSwitchBcfDeviceResponse
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

type DeleteBigSwitchBcfDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListBigSwitchBcfDevicesParams struct {
	p map[string]interface{}
}

func (p *ListBigSwitchBcfDevicesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bcfdeviceid"]; found {
		u.Set("bcfdeviceid", v.(string))
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
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListBigSwitchBcfDevicesParams) SetBcfdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bcfdeviceid"] = v
}

func (p *ListBigSwitchBcfDevicesParams) GetBcfdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bcfdeviceid"].(string)
	return value, ok
}

func (p *ListBigSwitchBcfDevicesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBigSwitchBcfDevicesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBigSwitchBcfDevicesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBigSwitchBcfDevicesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBigSwitchBcfDevicesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBigSwitchBcfDevicesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBigSwitchBcfDevicesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListBigSwitchBcfDevicesParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBigSwitchBcfDevicesParams instance,
// as then you are sure you have configured all required params
func (s *BigSwitchBCFService) NewListBigSwitchBcfDevicesParams() *ListBigSwitchBcfDevicesParams {
	p := &ListBigSwitchBcfDevicesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists BigSwitch BCF Controller devices
func (s *BigSwitchBCFService) ListBigSwitchBcfDevices(p *ListBigSwitchBcfDevicesParams) (*ListBigSwitchBcfDevicesResponse, error) {
	resp, err := s.cs.newRequest("listBigSwitchBcfDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBigSwitchBcfDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBigSwitchBcfDevicesResponse struct {
	Count               int                   `json:"count"`
	BigSwitchBcfDevices []*BigSwitchBcfDevice `json:"bigswitchbcfdevice"`
}

type BigSwitchBcfDevice struct {
	Bcfdeviceid         string `json:"bcfdeviceid"`
	Bigswitchdevicename string `json:"bigswitchdevicename"`
	Hostname            string `json:"hostname"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Nat                 bool   `json:"nat"`
	Password            string `json:"password"`
	Physicalnetworkid   string `json:"physicalnetworkid"`
	Provider            string `json:"provider"`
	Username            string `json:"username"`
}
