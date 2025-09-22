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

type NetrisServiceIface interface {
	AddNetrisProvider(p *AddNetrisProviderParams) (*AddNetrisProviderResponse, error)
	NewAddNetrisProviderParams(name string, netristag string, netrisurl string, password string, sitename string, tenantname string, username string, zoneid string) *AddNetrisProviderParams
	DeleteNetrisProvider(p *DeleteNetrisProviderParams) (*DeleteNetrisProviderResponse, error)
	NewDeleteNetrisProviderParams(id string) *DeleteNetrisProviderParams
	ListNetrisProviders(p *ListNetrisProvidersParams) (*ListNetrisProvidersResponse, error)
	NewListNetrisProvidersParams() *ListNetrisProvidersParams
}

type AddNetrisProviderParams struct {
	p map[string]interface{}
}

func (p *AddNetrisProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["netristag"]; found {
		u.Set("netristag", v.(string))
	}
	if v, found := p.p["netrisurl"]; found {
		u.Set("netrisurl", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["sitename"]; found {
		u.Set("sitename", v.(string))
	}
	if v, found := p.p["tenantname"]; found {
		u.Set("tenantname", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddNetrisProviderParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddNetrisProviderParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddNetrisProviderParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddNetrisProviderParams) SetNetristag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netristag"] = v
}

func (p *AddNetrisProviderParams) ResetNetristag() {
	if p.p != nil && p.p["netristag"] != nil {
		delete(p.p, "netristag")
	}
}

func (p *AddNetrisProviderParams) GetNetristag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["netristag"].(string)
	return value, ok
}

func (p *AddNetrisProviderParams) SetNetrisurl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netrisurl"] = v
}

func (p *AddNetrisProviderParams) ResetNetrisurl() {
	if p.p != nil && p.p["netrisurl"] != nil {
		delete(p.p, "netrisurl")
	}
}

func (p *AddNetrisProviderParams) GetNetrisurl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["netrisurl"].(string)
	return value, ok
}

func (p *AddNetrisProviderParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddNetrisProviderParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddNetrisProviderParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddNetrisProviderParams) SetSitename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sitename"] = v
}

func (p *AddNetrisProviderParams) ResetSitename() {
	if p.p != nil && p.p["sitename"] != nil {
		delete(p.p, "sitename")
	}
}

func (p *AddNetrisProviderParams) GetSitename() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sitename"].(string)
	return value, ok
}

func (p *AddNetrisProviderParams) SetTenantname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tenantname"] = v
}

func (p *AddNetrisProviderParams) ResetTenantname() {
	if p.p != nil && p.p["tenantname"] != nil {
		delete(p.p, "tenantname")
	}
}

func (p *AddNetrisProviderParams) GetTenantname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tenantname"].(string)
	return value, ok
}

func (p *AddNetrisProviderParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddNetrisProviderParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddNetrisProviderParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *AddNetrisProviderParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddNetrisProviderParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddNetrisProviderParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddNetrisProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetrisService) NewAddNetrisProviderParams(name string, netristag string, netrisurl string, password string, sitename string, tenantname string, username string, zoneid string) *AddNetrisProviderParams {
	p := &AddNetrisProviderParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["netristag"] = netristag
	p.p["netrisurl"] = netrisurl
	p.p["password"] = password
	p.p["sitename"] = sitename
	p.p["tenantname"] = tenantname
	p.p["username"] = username
	p.p["zoneid"] = zoneid
	return p
}

// Add Netris Provider to CloudStack
func (s *NetrisService) AddNetrisProvider(p *AddNetrisProviderParams) (*AddNetrisProviderResponse, error) {
	resp, err := s.cs.newPostRequest("addNetrisProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetrisProviderResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddNetrisProviderResponse struct {
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Name       string `json:"name"`
	Netristag  string `json:"netristag"`
	Netrisurl  string `json:"netrisurl"`
	Sitename   string `json:"sitename"`
	Tenantname string `json:"tenantname"`
	Uuid       string `json:"uuid"`
	Zoneid     string `json:"zoneid"`
	Zonename   string `json:"zonename"`
}

type DeleteNetrisProviderParams struct {
	p map[string]interface{}
}

func (p *DeleteNetrisProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetrisProviderParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteNetrisProviderParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteNetrisProviderParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetrisProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetrisService) NewDeleteNetrisProviderParams(id string) *DeleteNetrisProviderParams {
	p := &DeleteNetrisProviderParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// delete Netris Provider to CloudStack
func (s *NetrisService) DeleteNetrisProvider(p *DeleteNetrisProviderParams) (*DeleteNetrisProviderResponse, error) {
	resp, err := s.cs.newPostRequest("deleteNetrisProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetrisProviderResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteNetrisProviderResponse struct {
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Name       string `json:"name"`
	Netristag  string `json:"netristag"`
	Netrisurl  string `json:"netrisurl"`
	Sitename   string `json:"sitename"`
	Tenantname string `json:"tenantname"`
	Uuid       string `json:"uuid"`
	Zoneid     string `json:"zoneid"`
	Zonename   string `json:"zonename"`
}

type ListNetrisProvidersParams struct {
	p map[string]interface{}
}

func (p *ListNetrisProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListNetrisProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetrisProvidersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNetrisProvidersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetrisProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetrisProvidersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNetrisProvidersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetrisProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetrisProvidersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNetrisProvidersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNetrisProvidersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListNetrisProvidersParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListNetrisProvidersParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetrisProvidersParams instance,
// as then you are sure you have configured all required params
func (s *NetrisService) NewListNetrisProvidersParams() *ListNetrisProvidersParams {
	p := &ListNetrisProvidersParams{}
	p.p = make(map[string]interface{})
	return p
}

// list all Netris providers added to CloudStack
func (s *NetrisService) ListNetrisProviders(p *ListNetrisProvidersParams) (*ListNetrisProvidersResponse, error) {
	resp, err := s.cs.newRequest("listNetrisProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetrisProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetrisProvidersResponse struct {
	Count           int               `json:"count"`
	NetrisProviders []*NetrisProvider `json:"netrisprovider"`
}

type NetrisProvider struct {
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Name       string `json:"name"`
	Netristag  string `json:"netristag"`
	Netrisurl  string `json:"netrisurl"`
	Sitename   string `json:"sitename"`
	Tenantname string `json:"tenantname"`
	Uuid       string `json:"uuid"`
	Zoneid     string `json:"zoneid"`
	Zonename   string `json:"zonename"`
}
