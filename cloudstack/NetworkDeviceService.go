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
)

type NetworkDeviceServiceIface interface {
	AddNetworkDevice(p *AddNetworkDeviceParams) (*AddNetworkDeviceResponse, error)
	NewAddNetworkDeviceParams() *AddNetworkDeviceParams
	DeleteNetworkDevice(p *DeleteNetworkDeviceParams) (*DeleteNetworkDeviceResponse, error)
	NewDeleteNetworkDeviceParams(id string) *DeleteNetworkDeviceParams
	ListNetworkDevice(p *ListNetworkDeviceParams) (*ListNetworkDeviceResponse, error)
	NewListNetworkDeviceParams() *ListNetworkDeviceParams
}

type AddNetworkDeviceParams struct {
	p map[string]interface{}
}

func (p *AddNetworkDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["networkdeviceparameterlist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("networkdeviceparameterlist[%d].key", i), k)
			u.Set(fmt.Sprintf("networkdeviceparameterlist[%d].value", i), m[k])
		}
	}
	if v, found := p.p["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
	}
	return u
}

func (p *AddNetworkDeviceParams) SetNetworkdeviceparameterlist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdeviceparameterlist"] = v
}

func (p *AddNetworkDeviceParams) ResetNetworkdeviceparameterlist() {
	if p.p != nil && p.p["networkdeviceparameterlist"] != nil {
		delete(p.p, "networkdeviceparameterlist")
	}
}

func (p *AddNetworkDeviceParams) GetNetworkdeviceparameterlist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdeviceparameterlist"].(map[string]string)
	return value, ok
}

func (p *AddNetworkDeviceParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
}

func (p *AddNetworkDeviceParams) ResetNetworkdevicetype() {
	if p.p != nil && p.p["networkdevicetype"] != nil {
		delete(p.p, "networkdevicetype")
	}
}

func (p *AddNetworkDeviceParams) GetNetworkdevicetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdevicetype"].(string)
	return value, ok
}

// You should always use this function to get a new AddNetworkDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NetworkDeviceService) NewAddNetworkDeviceParams() *AddNetworkDeviceParams {
	p := &AddNetworkDeviceParams{}
	p.p = make(map[string]interface{})
	return p
}

// Adds a network device of one of the following types: ExternalDhcp, ExternalFirewall, ExternalLoadBalancer, PxeServer
func (s *NetworkDeviceService) AddNetworkDevice(p *AddNetworkDeviceParams) (*AddNetworkDeviceResponse, error) {
	resp, err := s.cs.newRequest("addNetworkDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetworkDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddNetworkDeviceResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
}

type DeleteNetworkDeviceParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkDeviceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteNetworkDeviceParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteNetworkDeviceParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NetworkDeviceService) NewDeleteNetworkDeviceParams(id string) *DeleteNetworkDeviceParams {
	p := &DeleteNetworkDeviceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes network device.
func (s *NetworkDeviceService) DeleteNetworkDevice(p *DeleteNetworkDeviceParams) (*DeleteNetworkDeviceResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteNetworkDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteNetworkDeviceResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteNetworkDeviceResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListNetworkDeviceParams struct {
	p map[string]interface{}
}

func (p *ListNetworkDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["networkdeviceparameterlist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("networkdeviceparameterlist[%d].key", i), k)
			u.Set(fmt.Sprintf("networkdeviceparameterlist[%d].value", i), m[k])
		}
	}
	if v, found := p.p["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
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

func (p *ListNetworkDeviceParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetworkDeviceParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNetworkDeviceParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetworkDeviceParams) SetNetworkdeviceparameterlist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdeviceparameterlist"] = v
}

func (p *ListNetworkDeviceParams) ResetNetworkdeviceparameterlist() {
	if p.p != nil && p.p["networkdeviceparameterlist"] != nil {
		delete(p.p, "networkdeviceparameterlist")
	}
}

func (p *ListNetworkDeviceParams) GetNetworkdeviceparameterlist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdeviceparameterlist"].(map[string]string)
	return value, ok
}

func (p *ListNetworkDeviceParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
}

func (p *ListNetworkDeviceParams) ResetNetworkdevicetype() {
	if p.p != nil && p.p["networkdevicetype"] != nil {
		delete(p.p, "networkdevicetype")
	}
}

func (p *ListNetworkDeviceParams) GetNetworkdevicetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdevicetype"].(string)
	return value, ok
}

func (p *ListNetworkDeviceParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetworkDeviceParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNetworkDeviceParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetworkDeviceParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetworkDeviceParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNetworkDeviceParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNetworkDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NetworkDeviceService) NewListNetworkDeviceParams() *ListNetworkDeviceParams {
	p := &ListNetworkDeviceParams{}
	p.p = make(map[string]interface{})
	return p
}

// List network devices
func (s *NetworkDeviceService) ListNetworkDevice(p *ListNetworkDeviceParams) (*ListNetworkDeviceResponse, error) {
	resp, err := s.cs.newRequest("listNetworkDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkDeviceResponse struct {
	Count         int              `json:"count"`
	NetworkDevice []*NetworkDevice `json:"networkdevice"`
}

type NetworkDevice struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
}
