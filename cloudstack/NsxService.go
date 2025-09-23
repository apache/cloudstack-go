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

type NsxServiceIface interface {
	AddNsxController(p *AddNsxControllerParams) (*AddNsxControllerResponse, error)
	NewAddNsxControllerParams(edgecluster string, name string, nsxproviderhostname string, password string, tier0gateway string, transportzone string, username string, zoneid string) *AddNsxControllerParams
	DeleteNsxController(p *DeleteNsxControllerParams) (*DeleteNsxControllerResponse, error)
	NewDeleteNsxControllerParams(nsxcontrollerid string) *DeleteNsxControllerParams
	ListNsxControllers(p *ListNsxControllersParams) (*ListNsxControllersResponse, error)
	NewListNsxControllersParams() *ListNsxControllersParams
}

type AddNsxControllerParams struct {
	p map[string]interface{}
}

func (p *AddNsxControllerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["edgecluster"]; found {
		u.Set("edgecluster", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["nsxproviderhostname"]; found {
		u.Set("nsxproviderhostname", v.(string))
	}
	if v, found := p.p["nsxproviderport"]; found {
		u.Set("nsxproviderport", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["tier0gateway"]; found {
		u.Set("tier0gateway", v.(string))
	}
	if v, found := p.p["transportzone"]; found {
		u.Set("transportzone", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddNsxControllerParams) SetEdgecluster(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["edgecluster"] = v
}

func (p *AddNsxControllerParams) ResetEdgecluster() {
	if p.p != nil && p.p["edgecluster"] != nil {
		delete(p.p, "edgecluster")
	}
}

func (p *AddNsxControllerParams) GetEdgecluster() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["edgecluster"].(string)
	return value, ok
}

func (p *AddNsxControllerParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddNsxControllerParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddNsxControllerParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddNsxControllerParams) SetNsxproviderhostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nsxproviderhostname"] = v
}

func (p *AddNsxControllerParams) ResetNsxproviderhostname() {
	if p.p != nil && p.p["nsxproviderhostname"] != nil {
		delete(p.p, "nsxproviderhostname")
	}
}

func (p *AddNsxControllerParams) GetNsxproviderhostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nsxproviderhostname"].(string)
	return value, ok
}

func (p *AddNsxControllerParams) SetNsxproviderport(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nsxproviderport"] = v
}

func (p *AddNsxControllerParams) ResetNsxproviderport() {
	if p.p != nil && p.p["nsxproviderport"] != nil {
		delete(p.p, "nsxproviderport")
	}
}

func (p *AddNsxControllerParams) GetNsxproviderport() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nsxproviderport"].(string)
	return value, ok
}

func (p *AddNsxControllerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddNsxControllerParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddNsxControllerParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddNsxControllerParams) SetTier0gateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tier0gateway"] = v
}

func (p *AddNsxControllerParams) ResetTier0gateway() {
	if p.p != nil && p.p["tier0gateway"] != nil {
		delete(p.p, "tier0gateway")
	}
}

func (p *AddNsxControllerParams) GetTier0gateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tier0gateway"].(string)
	return value, ok
}

func (p *AddNsxControllerParams) SetTransportzone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["transportzone"] = v
}

func (p *AddNsxControllerParams) ResetTransportzone() {
	if p.p != nil && p.p["transportzone"] != nil {
		delete(p.p, "transportzone")
	}
}

func (p *AddNsxControllerParams) GetTransportzone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["transportzone"].(string)
	return value, ok
}

func (p *AddNsxControllerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddNsxControllerParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddNsxControllerParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *AddNsxControllerParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddNsxControllerParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddNsxControllerParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddNsxControllerParams instance,
// as then you are sure you have configured all required params
func (s *NsxService) NewAddNsxControllerParams(edgecluster string, name string, nsxproviderhostname string, password string, tier0gateway string, transportzone string, username string, zoneid string) *AddNsxControllerParams {
	p := &AddNsxControllerParams{}
	p.p = make(map[string]interface{})
	p.p["edgecluster"] = edgecluster
	p.p["name"] = name
	p.p["nsxproviderhostname"] = nsxproviderhostname
	p.p["password"] = password
	p.p["tier0gateway"] = tier0gateway
	p.p["transportzone"] = transportzone
	p.p["username"] = username
	p.p["zoneid"] = zoneid
	return p
}

// Add NSX Controller to CloudStack
func (s *NsxService) AddNsxController(p *AddNsxControllerParams) (*AddNsxControllerResponse, error) {
	resp, err := s.cs.newPostRequest("addNsxController", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNsxControllerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddNsxControllerResponse struct {
	Edgecluster     string `json:"edgecluster"`
	Hostname        string `json:"hostname"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Name            string `json:"name"`
	Nsxprovideruuid string `json:"nsxprovideruuid"`
	Port            string `json:"port"`
	Tier0gateway    string `json:"tier0gateway"`
	Transportzone   string `json:"transportzone"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type DeleteNsxControllerParams struct {
	p map[string]interface{}
}

func (p *DeleteNsxControllerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nsxcontrollerid"]; found {
		u.Set("nsxcontrollerid", v.(string))
	}
	return u
}

func (p *DeleteNsxControllerParams) SetNsxcontrollerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nsxcontrollerid"] = v
}

func (p *DeleteNsxControllerParams) ResetNsxcontrollerid() {
	if p.p != nil && p.p["nsxcontrollerid"] != nil {
		delete(p.p, "nsxcontrollerid")
	}
}

func (p *DeleteNsxControllerParams) GetNsxcontrollerid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nsxcontrollerid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNsxControllerParams instance,
// as then you are sure you have configured all required params
func (s *NsxService) NewDeleteNsxControllerParams(nsxcontrollerid string) *DeleteNsxControllerParams {
	p := &DeleteNsxControllerParams{}
	p.p = make(map[string]interface{})
	p.p["nsxcontrollerid"] = nsxcontrollerid
	return p
}

// delete NSX Controller to CloudStack
func (s *NsxService) DeleteNsxController(p *DeleteNsxControllerParams) (*DeleteNsxControllerResponse, error) {
	resp, err := s.cs.newPostRequest("deleteNsxController", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNsxControllerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteNsxControllerResponse struct {
	Edgecluster     string `json:"edgecluster"`
	Hostname        string `json:"hostname"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Name            string `json:"name"`
	Nsxprovideruuid string `json:"nsxprovideruuid"`
	Port            string `json:"port"`
	Tier0gateway    string `json:"tier0gateway"`
	Transportzone   string `json:"transportzone"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type ListNsxControllersParams struct {
	p map[string]interface{}
}

func (p *ListNsxControllersParams) toURLValues() url.Values {
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

func (p *ListNsxControllersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNsxControllersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNsxControllersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNsxControllersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNsxControllersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNsxControllersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNsxControllersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNsxControllersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNsxControllersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNsxControllersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListNsxControllersParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListNsxControllersParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNsxControllersParams instance,
// as then you are sure you have configured all required params
func (s *NsxService) NewListNsxControllersParams() *ListNsxControllersParams {
	p := &ListNsxControllersParams{}
	p.p = make(map[string]interface{})
	return p
}

// list all NSX controllers added to CloudStack
func (s *NsxService) ListNsxControllers(p *ListNsxControllersParams) (*ListNsxControllersResponse, error) {
	resp, err := s.cs.newRequest("listNsxControllers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNsxControllersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNsxControllersResponse struct {
	Count          int              `json:"count"`
	NsxControllers []*NsxController `json:"nsxcontroller"`
}

type NsxController struct {
	Edgecluster     string `json:"edgecluster"`
	Hostname        string `json:"hostname"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Name            string `json:"name"`
	Nsxprovideruuid string `json:"nsxprovideruuid"`
	Port            string `json:"port"`
	Tier0gateway    string `json:"tier0gateway"`
	Transportzone   string `json:"transportzone"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}
