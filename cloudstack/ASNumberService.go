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

type ASNumberServiceIface interface {
	ListASNumbers(p *ListASNumbersParams) (*ListASNumbersResponse, error)
	NewListASNumbersParams() *ListASNumbersParams
	ReleaseASNumber(p *ReleaseASNumberParams) (*ReleaseASNumberResponse, error)
	NewReleaseASNumberParams(asnumber int64, zoneid string) *ReleaseASNumberParams
}

type ListASNumbersParams struct {
	p map[string]interface{}
}

func (p *ListASNumbersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["asnrangeid"]; found {
		u.Set("asnrangeid", v.(string))
	}
	if v, found := p.p["asnumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("asnumber", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["isallocated"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isallocated", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListASNumbersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListASNumbersParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListASNumbersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListASNumbersParams) SetAsnrangeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asnrangeid"] = v
}

func (p *ListASNumbersParams) ResetAsnrangeid() {
	if p.p != nil && p.p["asnrangeid"] != nil {
		delete(p.p, "asnrangeid")
	}
}

func (p *ListASNumbersParams) GetAsnrangeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asnrangeid"].(string)
	return value, ok
}

func (p *ListASNumbersParams) SetAsnumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asnumber"] = v
}

func (p *ListASNumbersParams) ResetAsnumber() {
	if p.p != nil && p.p["asnumber"] != nil {
		delete(p.p, "asnumber")
	}
}

func (p *ListASNumbersParams) GetAsnumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asnumber"].(int)
	return value, ok
}

func (p *ListASNumbersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListASNumbersParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListASNumbersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListASNumbersParams) SetIsallocated(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isallocated"] = v
}

func (p *ListASNumbersParams) ResetIsallocated() {
	if p.p != nil && p.p["isallocated"] != nil {
		delete(p.p, "isallocated")
	}
}

func (p *ListASNumbersParams) GetIsallocated() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isallocated"].(bool)
	return value, ok
}

func (p *ListASNumbersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListASNumbersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListASNumbersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListASNumbersParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListASNumbersParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListASNumbersParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListASNumbersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListASNumbersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListASNumbersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListASNumbersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListASNumbersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListASNumbersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListASNumbersParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListASNumbersParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListASNumbersParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *ListASNumbersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListASNumbersParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListASNumbersParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListASNumbersParams instance,
// as then you are sure you have configured all required params
func (s *ASNumberService) NewListASNumbersParams() *ListASNumbersParams {
	p := &ListASNumbersParams{}
	p.p = make(map[string]interface{})
	return p
}

// List Autonomous Systems Numbers
func (s *ASNumberService) ListASNumbers(p *ListASNumbersParams) (*ListASNumbersResponse, error) {
	resp, err := s.cs.newRequest("listASNumbers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListASNumbersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListASNumbersResponse struct {
	Count     int         `json:"count"`
	ASNumbers []*ASNumber `json:"asnumber"`
}

type ASNumber struct {
	Account               string `json:"account"`
	Accountid             string `json:"accountid"`
	Allocated             string `json:"allocated"`
	Allocationstate       string `json:"allocationstate"`
	Asnrange              string `json:"asnrange"`
	Asnrangeid            string `json:"asnrangeid"`
	Asnumber              int64  `json:"asnumber"`
	Associatednetworkid   string `json:"associatednetworkid"`
	Associatednetworkname string `json:"associatednetworkname"`
	Created               string `json:"created"`
	Domain                string `json:"domain"`
	Domainid              string `json:"domainid"`
	Id                    string `json:"id"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Vpcid                 string `json:"vpcid"`
	Vpcname               string `json:"vpcname"`
	Zoneid                string `json:"zoneid"`
	Zonename              string `json:"zonename"`
}

type ReleaseASNumberParams struct {
	p map[string]interface{}
}

func (p *ReleaseASNumberParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["asnumber"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("asnumber", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ReleaseASNumberParams) SetAsnumber(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asnumber"] = v
}

func (p *ReleaseASNumberParams) ResetAsnumber() {
	if p.p != nil && p.p["asnumber"] != nil {
		delete(p.p, "asnumber")
	}
}

func (p *ReleaseASNumberParams) GetAsnumber() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asnumber"].(int64)
	return value, ok
}

func (p *ReleaseASNumberParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ReleaseASNumberParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ReleaseASNumberParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseASNumberParams instance,
// as then you are sure you have configured all required params
func (s *ASNumberService) NewReleaseASNumberParams(asnumber int64, zoneid string) *ReleaseASNumberParams {
	p := &ReleaseASNumberParams{}
	p.p = make(map[string]interface{})
	p.p["asnumber"] = asnumber
	p.p["zoneid"] = zoneid
	return p
}

// Releases an AS Number back to the pool
func (s *ASNumberService) ReleaseASNumber(p *ReleaseASNumberParams) (*ReleaseASNumberResponse, error) {
	resp, err := s.cs.newPostRequest("releaseASNumber", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseASNumberResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ReleaseASNumberResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ReleaseASNumberResponse) UnmarshalJSON(b []byte) error {
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

	type alias ReleaseASNumberResponse
	return json.Unmarshal(b, (*alias)(r))
}
