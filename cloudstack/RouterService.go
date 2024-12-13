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

type RouterServiceIface interface {
	ChangeServiceForRouter(p *ChangeServiceForRouterParams) (*ChangeServiceForRouterResponse, error)
	NewChangeServiceForRouterParams(id string, serviceofferingid string) *ChangeServiceForRouterParams
	ConfigureVirtualRouterElement(p *ConfigureVirtualRouterElementParams) (*VirtualRouterElementResponse, error)
	NewConfigureVirtualRouterElementParams(enabled bool, id string) *ConfigureVirtualRouterElementParams
	CreateVirtualRouterElement(p *CreateVirtualRouterElementParams) (*CreateVirtualRouterElementResponse, error)
	NewCreateVirtualRouterElementParams(nspid string) *CreateVirtualRouterElementParams
	DestroyRouter(p *DestroyRouterParams) (*DestroyRouterResponse, error)
	NewDestroyRouterParams(id string) *DestroyRouterParams
	ListRouters(p *ListRoutersParams) (*ListRoutersResponse, error)
	NewListRoutersParams() *ListRoutersParams
	GetRouterID(name string, opts ...OptionFunc) (string, int, error)
	GetRouterByName(name string, opts ...OptionFunc) (*Router, int, error)
	GetRouterByID(id string, opts ...OptionFunc) (*Router, int, error)
	ListVirtualRouterElements(p *ListVirtualRouterElementsParams) (*ListVirtualRouterElementsResponse, error)
	NewListVirtualRouterElementsParams() *ListVirtualRouterElementsParams
	GetVirtualRouterElementByID(id string, opts ...OptionFunc) (*VirtualRouterElement, int, error)
	RebootRouter(p *RebootRouterParams) (*RebootRouterResponse, error)
	NewRebootRouterParams(id string) *RebootRouterParams
	StartRouter(p *StartRouterParams) (*StartRouterResponse, error)
	NewStartRouterParams(id string) *StartRouterParams
	StopRouter(p *StopRouterParams) (*StopRouterResponse, error)
	NewStopRouterParams(id string) *StopRouterParams
}

type ChangeServiceForRouterParams struct {
	p map[string]interface{}
}

func (p *ChangeServiceForRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ChangeServiceForRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ChangeServiceForRouterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ChangeServiceForRouterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ChangeServiceForRouterParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ChangeServiceForRouterParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ChangeServiceForRouterParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeServiceForRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewChangeServiceForRouterParams(id string, serviceofferingid string) *ChangeServiceForRouterParams {
	p := &ChangeServiceForRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Upgrades domain router to a new service offering
func (s *RouterService) ChangeServiceForRouter(p *ChangeServiceForRouterParams) (*ChangeServiceForRouterResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForRouterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ChangeServiceForRouterResponse struct {
	Account             string                                             `json:"account"`
	Created             string                                             `json:"created"`
	Dns1                string                                             `json:"dns1"`
	Dns2                string                                             `json:"dns2"`
	Domain              string                                             `json:"domain"`
	Domainid            string                                             `json:"domainid"`
	Domainpath          string                                             `json:"domainpath"`
	Gateway             string                                             `json:"gateway"`
	Guestipaddress      string                                             `json:"guestipaddress"`
	Guestmacaddress     string                                             `json:"guestmacaddress"`
	Guestnetmask        string                                             `json:"guestnetmask"`
	Guestnetworkid      string                                             `json:"guestnetworkid"`
	Guestnetworkname    string                                             `json:"guestnetworkname"`
	Hasannotations      bool                                               `json:"hasannotations"`
	Healthcheckresults  []ChangeServiceForRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                               `json:"healthchecksfailed"`
	Hostcontrolstate    string                                             `json:"hostcontrolstate"`
	Hostid              string                                             `json:"hostid"`
	Hostname            string                                             `json:"hostname"`
	Hypervisor          string                                             `json:"hypervisor"`
	Id                  string                                             `json:"id"`
	Ip6dns1             string                                             `json:"ip6dns1"`
	Ip6dns2             string                                             `json:"ip6dns2"`
	Isredundantrouter   bool                                               `json:"isredundantrouter"`
	JobID               string                                             `json:"jobid"`
	Jobstatus           int                                                `json:"jobstatus"`
	Linklocalip         string                                             `json:"linklocalip"`
	Linklocalmacaddress string                                             `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                             `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                             `json:"linklocalnetworkid"`
	Name                string                                             `json:"name"`
	Networkdomain       string                                             `json:"networkdomain"`
	Nic                 []Nic                                              `json:"nic"`
	Podid               string                                             `json:"podid"`
	Podname             string                                             `json:"podname"`
	Project             string                                             `json:"project"`
	Projectid           string                                             `json:"projectid"`
	Publicip            string                                             `json:"publicip"`
	Publicmacaddress    string                                             `json:"publicmacaddress"`
	Publicnetmask       string                                             `json:"publicnetmask"`
	Publicnetworkid     string                                             `json:"publicnetworkid"`
	Redundantstate      string                                             `json:"redundantstate"`
	Requiresupgrade     bool                                               `json:"requiresupgrade"`
	Role                string                                             `json:"role"`
	Scriptsversion      string                                             `json:"scriptsversion"`
	Serviceofferingid   string                                             `json:"serviceofferingid"`
	Serviceofferingname string                                             `json:"serviceofferingname"`
	Softwareversion     string                                             `json:"softwareversion"`
	State               string                                             `json:"state"`
	Templateid          string                                             `json:"templateid"`
	Templatename        string                                             `json:"templatename"`
	Version             string                                             `json:"version"`
	Vpcid               string                                             `json:"vpcid"`
	Vpcname             string                                             `json:"vpcname"`
	Zoneid              string                                             `json:"zoneid"`
	Zonename            string                                             `json:"zonename"`
}

type ChangeServiceForRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type ConfigureVirtualRouterElementParams struct {
	p map[string]interface{}
}

func (p *ConfigureVirtualRouterElementParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ConfigureVirtualRouterElementParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *ConfigureVirtualRouterElementParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *ConfigureVirtualRouterElementParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *ConfigureVirtualRouterElementParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ConfigureVirtualRouterElementParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ConfigureVirtualRouterElementParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigureVirtualRouterElementParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewConfigureVirtualRouterElementParams(enabled bool, id string) *ConfigureVirtualRouterElementParams {
	p := &ConfigureVirtualRouterElementParams{}
	p.p = make(map[string]interface{})
	p.p["enabled"] = enabled
	p.p["id"] = id
	return p
}

// Configures a virtual router element.
func (s *RouterService) ConfigureVirtualRouterElement(p *ConfigureVirtualRouterElementParams) (*VirtualRouterElementResponse, error) {
	resp, err := s.cs.newRequest("configureVirtualRouterElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r VirtualRouterElementResponse
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

type VirtualRouterElementResponse struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Enabled    bool   `json:"enabled"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Nspid      string `json:"nspid"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
}

type CreateVirtualRouterElementParams struct {
	p map[string]interface{}
}

func (p *CreateVirtualRouterElementParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nspid"]; found {
		u.Set("nspid", v.(string))
	}
	if v, found := p.p["providertype"]; found {
		u.Set("providertype", v.(string))
	}
	return u
}

func (p *CreateVirtualRouterElementParams) SetNspid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nspid"] = v
}

func (p *CreateVirtualRouterElementParams) ResetNspid() {
	if p.p != nil && p.p["nspid"] != nil {
		delete(p.p, "nspid")
	}
}

func (p *CreateVirtualRouterElementParams) GetNspid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nspid"].(string)
	return value, ok
}

func (p *CreateVirtualRouterElementParams) SetProvidertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["providertype"] = v
}

func (p *CreateVirtualRouterElementParams) ResetProvidertype() {
	if p.p != nil && p.p["providertype"] != nil {
		delete(p.p, "providertype")
	}
}

func (p *CreateVirtualRouterElementParams) GetProvidertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["providertype"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVirtualRouterElementParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewCreateVirtualRouterElementParams(nspid string) *CreateVirtualRouterElementParams {
	p := &CreateVirtualRouterElementParams{}
	p.p = make(map[string]interface{})
	p.p["nspid"] = nspid
	return p
}

// Create a virtual router element.
func (s *RouterService) CreateVirtualRouterElement(p *CreateVirtualRouterElementParams) (*CreateVirtualRouterElementResponse, error) {
	resp, err := s.cs.newRequest("createVirtualRouterElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVirtualRouterElementResponse
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

type CreateVirtualRouterElementResponse struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Enabled    bool   `json:"enabled"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Nspid      string `json:"nspid"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
}

type DestroyRouterParams struct {
	p map[string]interface{}
}

func (p *DestroyRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DestroyRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DestroyRouterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DestroyRouterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DestroyRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewDestroyRouterParams(id string) *DestroyRouterParams {
	p := &DestroyRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroys a router.
func (s *RouterService) DestroyRouter(p *DestroyRouterParams) (*DestroyRouterResponse, error) {
	resp, err := s.cs.newRequest("destroyRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroyRouterResponse
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

type DestroyRouterResponse struct {
	Account             string                                    `json:"account"`
	Created             string                                    `json:"created"`
	Dns1                string                                    `json:"dns1"`
	Dns2                string                                    `json:"dns2"`
	Domain              string                                    `json:"domain"`
	Domainid            string                                    `json:"domainid"`
	Domainpath          string                                    `json:"domainpath"`
	Gateway             string                                    `json:"gateway"`
	Guestipaddress      string                                    `json:"guestipaddress"`
	Guestmacaddress     string                                    `json:"guestmacaddress"`
	Guestnetmask        string                                    `json:"guestnetmask"`
	Guestnetworkid      string                                    `json:"guestnetworkid"`
	Guestnetworkname    string                                    `json:"guestnetworkname"`
	Hasannotations      bool                                      `json:"hasannotations"`
	Healthcheckresults  []DestroyRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                      `json:"healthchecksfailed"`
	Hostcontrolstate    string                                    `json:"hostcontrolstate"`
	Hostid              string                                    `json:"hostid"`
	Hostname            string                                    `json:"hostname"`
	Hypervisor          string                                    `json:"hypervisor"`
	Id                  string                                    `json:"id"`
	Ip6dns1             string                                    `json:"ip6dns1"`
	Ip6dns2             string                                    `json:"ip6dns2"`
	Isredundantrouter   bool                                      `json:"isredundantrouter"`
	JobID               string                                    `json:"jobid"`
	Jobstatus           int                                       `json:"jobstatus"`
	Linklocalip         string                                    `json:"linklocalip"`
	Linklocalmacaddress string                                    `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                    `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                    `json:"linklocalnetworkid"`
	Name                string                                    `json:"name"`
	Networkdomain       string                                    `json:"networkdomain"`
	Nic                 []Nic                                     `json:"nic"`
	Podid               string                                    `json:"podid"`
	Podname             string                                    `json:"podname"`
	Project             string                                    `json:"project"`
	Projectid           string                                    `json:"projectid"`
	Publicip            string                                    `json:"publicip"`
	Publicmacaddress    string                                    `json:"publicmacaddress"`
	Publicnetmask       string                                    `json:"publicnetmask"`
	Publicnetworkid     string                                    `json:"publicnetworkid"`
	Redundantstate      string                                    `json:"redundantstate"`
	Requiresupgrade     bool                                      `json:"requiresupgrade"`
	Role                string                                    `json:"role"`
	Scriptsversion      string                                    `json:"scriptsversion"`
	Serviceofferingid   string                                    `json:"serviceofferingid"`
	Serviceofferingname string                                    `json:"serviceofferingname"`
	Softwareversion     string                                    `json:"softwareversion"`
	State               string                                    `json:"state"`
	Templateid          string                                    `json:"templateid"`
	Templatename        string                                    `json:"templatename"`
	Version             string                                    `json:"version"`
	Vpcid               string                                    `json:"vpcid"`
	Vpcname             string                                    `json:"vpcname"`
	Zoneid              string                                    `json:"zoneid"`
	Zonename            string                                    `json:"zonename"`
}

type DestroyRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type ListRoutersParams struct {
	p map[string]interface{}
}

func (p *ListRoutersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fetchhealthcheckresults"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fetchhealthcheckresults", vv)
	}
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["healthchecksfailed"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("healthchecksfailed", vv)
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["version"]; found {
		u.Set("version", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListRoutersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListRoutersParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListRoutersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListRoutersParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListRoutersParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListRoutersParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListRoutersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetFetchhealthcheckresults(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fetchhealthcheckresults"] = v
}

func (p *ListRoutersParams) ResetFetchhealthcheckresults() {
	if p.p != nil && p.p["fetchhealthcheckresults"] != nil {
		delete(p.p, "fetchhealthcheckresults")
	}
}

func (p *ListRoutersParams) GetFetchhealthcheckresults() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fetchhealthcheckresults"].(bool)
	return value, ok
}

func (p *ListRoutersParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
}

func (p *ListRoutersParams) ResetForvpc() {
	if p.p != nil && p.p["forvpc"] != nil {
		delete(p.p, "forvpc")
	}
}

func (p *ListRoutersParams) GetForvpc() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forvpc"].(bool)
	return value, ok
}

func (p *ListRoutersParams) SetHealthchecksfailed(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["healthchecksfailed"] = v
}

func (p *ListRoutersParams) ResetHealthchecksfailed() {
	if p.p != nil && p.p["healthchecksfailed"] != nil {
		delete(p.p, "healthchecksfailed")
	}
}

func (p *ListRoutersParams) GetHealthchecksfailed() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["healthchecksfailed"].(bool)
	return value, ok
}

func (p *ListRoutersParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListRoutersParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ListRoutersParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListRoutersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListRoutersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListRoutersParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListRoutersParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListRoutersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListRoutersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListRoutersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListRoutersParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListRoutersParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListRoutersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListRoutersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListRoutersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListRoutersParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListRoutersParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListRoutersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListRoutersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListRoutersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListRoutersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListRoutersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListRoutersParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListRoutersParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListRoutersParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListRoutersParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListRoutersParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListRoutersParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListRoutersParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetVersion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["version"] = v
}

func (p *ListRoutersParams) ResetVersion() {
	if p.p != nil && p.p["version"] != nil {
		delete(p.p, "version")
	}
}

func (p *ListRoutersParams) GetVersion() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["version"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListRoutersParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListRoutersParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *ListRoutersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListRoutersParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListRoutersParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListRoutersParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewListRoutersParams() *ListRoutersParams {
	p := &ListRoutersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListRoutersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListRouters(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Routers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Routers {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterByName(name string, opts ...OptionFunc) (*Router, int, error) {
	id, count, err := s.GetRouterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetRouterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterByID(id string, opts ...OptionFunc) (*Router, int, error) {
	p := &ListRoutersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListRouters(p)
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
		return l.Routers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Router UUID: %s!", id)
}

// List routers.
func (s *RouterService) ListRouters(p *ListRoutersParams) (*ListRoutersResponse, error) {
	resp, err := s.cs.newRequest("listRouters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRoutersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRoutersResponse struct {
	Count   int       `json:"count"`
	Routers []*Router `json:"router"`
}

type Router struct {
	Account             string                     `json:"account"`
	Created             string                     `json:"created"`
	Dns1                string                     `json:"dns1"`
	Dns2                string                     `json:"dns2"`
	Domain              string                     `json:"domain"`
	Domainid            string                     `json:"domainid"`
	Domainpath          string                     `json:"domainpath"`
	Gateway             string                     `json:"gateway"`
	Guestipaddress      string                     `json:"guestipaddress"`
	Guestmacaddress     string                     `json:"guestmacaddress"`
	Guestnetmask        string                     `json:"guestnetmask"`
	Guestnetworkid      string                     `json:"guestnetworkid"`
	Guestnetworkname    string                     `json:"guestnetworkname"`
	Hasannotations      bool                       `json:"hasannotations"`
	Healthcheckresults  []RouterHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                       `json:"healthchecksfailed"`
	Hostcontrolstate    string                     `json:"hostcontrolstate"`
	Hostid              string                     `json:"hostid"`
	Hostname            string                     `json:"hostname"`
	Hypervisor          string                     `json:"hypervisor"`
	Id                  string                     `json:"id"`
	Ip6dns1             string                     `json:"ip6dns1"`
	Ip6dns2             string                     `json:"ip6dns2"`
	Isredundantrouter   bool                       `json:"isredundantrouter"`
	JobID               string                     `json:"jobid"`
	Jobstatus           int                        `json:"jobstatus"`
	Linklocalip         string                     `json:"linklocalip"`
	Linklocalmacaddress string                     `json:"linklocalmacaddress"`
	Linklocalnetmask    string                     `json:"linklocalnetmask"`
	Linklocalnetworkid  string                     `json:"linklocalnetworkid"`
	Name                string                     `json:"name"`
	Networkdomain       string                     `json:"networkdomain"`
	Nic                 []Nic                      `json:"nic"`
	Podid               string                     `json:"podid"`
	Podname             string                     `json:"podname"`
	Project             string                     `json:"project"`
	Projectid           string                     `json:"projectid"`
	Publicip            string                     `json:"publicip"`
	Publicmacaddress    string                     `json:"publicmacaddress"`
	Publicnetmask       string                     `json:"publicnetmask"`
	Publicnetworkid     string                     `json:"publicnetworkid"`
	Redundantstate      string                     `json:"redundantstate"`
	Requiresupgrade     bool                       `json:"requiresupgrade"`
	Role                string                     `json:"role"`
	Scriptsversion      string                     `json:"scriptsversion"`
	Serviceofferingid   string                     `json:"serviceofferingid"`
	Serviceofferingname string                     `json:"serviceofferingname"`
	Softwareversion     string                     `json:"softwareversion"`
	State               string                     `json:"state"`
	Templateid          string                     `json:"templateid"`
	Templatename        string                     `json:"templatename"`
	Version             string                     `json:"version"`
	Vpcid               string                     `json:"vpcid"`
	Vpcname             string                     `json:"vpcname"`
	Zoneid              string                     `json:"zoneid"`
	Zonename            string                     `json:"zonename"`
}

type RouterHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type ListVirtualRouterElementsParams struct {
	p map[string]interface{}
}

func (p *ListVirtualRouterElementsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nspid"]; found {
		u.Set("nspid", v.(string))
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

func (p *ListVirtualRouterElementsParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *ListVirtualRouterElementsParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *ListVirtualRouterElementsParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *ListVirtualRouterElementsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVirtualRouterElementsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVirtualRouterElementsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVirtualRouterElementsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVirtualRouterElementsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVirtualRouterElementsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVirtualRouterElementsParams) SetNspid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nspid"] = v
}

func (p *ListVirtualRouterElementsParams) ResetNspid() {
	if p.p != nil && p.p["nspid"] != nil {
		delete(p.p, "nspid")
	}
}

func (p *ListVirtualRouterElementsParams) GetNspid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nspid"].(string)
	return value, ok
}

func (p *ListVirtualRouterElementsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVirtualRouterElementsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVirtualRouterElementsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVirtualRouterElementsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVirtualRouterElementsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVirtualRouterElementsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListVirtualRouterElementsParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewListVirtualRouterElementsParams() *ListVirtualRouterElementsParams {
	p := &ListVirtualRouterElementsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetVirtualRouterElementByID(id string, opts ...OptionFunc) (*VirtualRouterElement, int, error) {
	p := &ListVirtualRouterElementsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVirtualRouterElements(p)
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
		return l.VirtualRouterElements[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualRouterElement UUID: %s!", id)
}

// Lists all available virtual router elements.
func (s *RouterService) ListVirtualRouterElements(p *ListVirtualRouterElementsParams) (*ListVirtualRouterElementsResponse, error) {
	resp, err := s.cs.newRequest("listVirtualRouterElements", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualRouterElementsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVirtualRouterElementsResponse struct {
	Count                 int                     `json:"count"`
	VirtualRouterElements []*VirtualRouterElement `json:"virtualrouterelement"`
}

type VirtualRouterElement struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Enabled    bool   `json:"enabled"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Nspid      string `json:"nspid"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
}

type RebootRouterParams struct {
	p map[string]interface{}
}

func (p *RebootRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RebootRouterParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *RebootRouterParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *RebootRouterParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *RebootRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RebootRouterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RebootRouterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RebootRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewRebootRouterParams(id string) *RebootRouterParams {
	p := &RebootRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a router.
func (s *RouterService) RebootRouter(p *RebootRouterParams) (*RebootRouterResponse, error) {
	resp, err := s.cs.newRequest("rebootRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootRouterResponse
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

type RebootRouterResponse struct {
	Account             string                                   `json:"account"`
	Created             string                                   `json:"created"`
	Dns1                string                                   `json:"dns1"`
	Dns2                string                                   `json:"dns2"`
	Domain              string                                   `json:"domain"`
	Domainid            string                                   `json:"domainid"`
	Domainpath          string                                   `json:"domainpath"`
	Gateway             string                                   `json:"gateway"`
	Guestipaddress      string                                   `json:"guestipaddress"`
	Guestmacaddress     string                                   `json:"guestmacaddress"`
	Guestnetmask        string                                   `json:"guestnetmask"`
	Guestnetworkid      string                                   `json:"guestnetworkid"`
	Guestnetworkname    string                                   `json:"guestnetworkname"`
	Hasannotations      bool                                     `json:"hasannotations"`
	Healthcheckresults  []RebootRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                     `json:"healthchecksfailed"`
	Hostcontrolstate    string                                   `json:"hostcontrolstate"`
	Hostid              string                                   `json:"hostid"`
	Hostname            string                                   `json:"hostname"`
	Hypervisor          string                                   `json:"hypervisor"`
	Id                  string                                   `json:"id"`
	Ip6dns1             string                                   `json:"ip6dns1"`
	Ip6dns2             string                                   `json:"ip6dns2"`
	Isredundantrouter   bool                                     `json:"isredundantrouter"`
	JobID               string                                   `json:"jobid"`
	Jobstatus           int                                      `json:"jobstatus"`
	Linklocalip         string                                   `json:"linklocalip"`
	Linklocalmacaddress string                                   `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                   `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                   `json:"linklocalnetworkid"`
	Name                string                                   `json:"name"`
	Networkdomain       string                                   `json:"networkdomain"`
	Nic                 []Nic                                    `json:"nic"`
	Podid               string                                   `json:"podid"`
	Podname             string                                   `json:"podname"`
	Project             string                                   `json:"project"`
	Projectid           string                                   `json:"projectid"`
	Publicip            string                                   `json:"publicip"`
	Publicmacaddress    string                                   `json:"publicmacaddress"`
	Publicnetmask       string                                   `json:"publicnetmask"`
	Publicnetworkid     string                                   `json:"publicnetworkid"`
	Redundantstate      string                                   `json:"redundantstate"`
	Requiresupgrade     bool                                     `json:"requiresupgrade"`
	Role                string                                   `json:"role"`
	Scriptsversion      string                                   `json:"scriptsversion"`
	Serviceofferingid   string                                   `json:"serviceofferingid"`
	Serviceofferingname string                                   `json:"serviceofferingname"`
	Softwareversion     string                                   `json:"softwareversion"`
	State               string                                   `json:"state"`
	Templateid          string                                   `json:"templateid"`
	Templatename        string                                   `json:"templatename"`
	Version             string                                   `json:"version"`
	Vpcid               string                                   `json:"vpcid"`
	Vpcname             string                                   `json:"vpcname"`
	Zoneid              string                                   `json:"zoneid"`
	Zonename            string                                   `json:"zonename"`
}

type RebootRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type StartRouterParams struct {
	p map[string]interface{}
}

func (p *StartRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StartRouterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StartRouterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StartRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewStartRouterParams(id string) *StartRouterParams {
	p := &StartRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a router.
func (s *RouterService) StartRouter(p *StartRouterParams) (*StartRouterResponse, error) {
	resp, err := s.cs.newRequest("startRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartRouterResponse
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

type StartRouterResponse struct {
	Account             string                                  `json:"account"`
	Created             string                                  `json:"created"`
	Dns1                string                                  `json:"dns1"`
	Dns2                string                                  `json:"dns2"`
	Domain              string                                  `json:"domain"`
	Domainid            string                                  `json:"domainid"`
	Domainpath          string                                  `json:"domainpath"`
	Gateway             string                                  `json:"gateway"`
	Guestipaddress      string                                  `json:"guestipaddress"`
	Guestmacaddress     string                                  `json:"guestmacaddress"`
	Guestnetmask        string                                  `json:"guestnetmask"`
	Guestnetworkid      string                                  `json:"guestnetworkid"`
	Guestnetworkname    string                                  `json:"guestnetworkname"`
	Hasannotations      bool                                    `json:"hasannotations"`
	Healthcheckresults  []StartRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                    `json:"healthchecksfailed"`
	Hostcontrolstate    string                                  `json:"hostcontrolstate"`
	Hostid              string                                  `json:"hostid"`
	Hostname            string                                  `json:"hostname"`
	Hypervisor          string                                  `json:"hypervisor"`
	Id                  string                                  `json:"id"`
	Ip6dns1             string                                  `json:"ip6dns1"`
	Ip6dns2             string                                  `json:"ip6dns2"`
	Isredundantrouter   bool                                    `json:"isredundantrouter"`
	JobID               string                                  `json:"jobid"`
	Jobstatus           int                                     `json:"jobstatus"`
	Linklocalip         string                                  `json:"linklocalip"`
	Linklocalmacaddress string                                  `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                  `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                  `json:"linklocalnetworkid"`
	Name                string                                  `json:"name"`
	Networkdomain       string                                  `json:"networkdomain"`
	Nic                 []Nic                                   `json:"nic"`
	Podid               string                                  `json:"podid"`
	Podname             string                                  `json:"podname"`
	Project             string                                  `json:"project"`
	Projectid           string                                  `json:"projectid"`
	Publicip            string                                  `json:"publicip"`
	Publicmacaddress    string                                  `json:"publicmacaddress"`
	Publicnetmask       string                                  `json:"publicnetmask"`
	Publicnetworkid     string                                  `json:"publicnetworkid"`
	Redundantstate      string                                  `json:"redundantstate"`
	Requiresupgrade     bool                                    `json:"requiresupgrade"`
	Role                string                                  `json:"role"`
	Scriptsversion      string                                  `json:"scriptsversion"`
	Serviceofferingid   string                                  `json:"serviceofferingid"`
	Serviceofferingname string                                  `json:"serviceofferingname"`
	Softwareversion     string                                  `json:"softwareversion"`
	State               string                                  `json:"state"`
	Templateid          string                                  `json:"templateid"`
	Templatename        string                                  `json:"templatename"`
	Version             string                                  `json:"version"`
	Vpcid               string                                  `json:"vpcid"`
	Vpcname             string                                  `json:"vpcname"`
	Zoneid              string                                  `json:"zoneid"`
	Zonename            string                                  `json:"zonename"`
}

type StartRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type StopRouterParams struct {
	p map[string]interface{}
}

func (p *StopRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StopRouterParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *StopRouterParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *StopRouterParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *StopRouterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StopRouterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StopRouterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewStopRouterParams(id string) *StopRouterParams {
	p := &StopRouterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a router.
func (s *RouterService) StopRouter(p *StopRouterParams) (*StopRouterResponse, error) {
	resp, err := s.cs.newRequest("stopRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopRouterResponse
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

type StopRouterResponse struct {
	Account             string                                 `json:"account"`
	Created             string                                 `json:"created"`
	Dns1                string                                 `json:"dns1"`
	Dns2                string                                 `json:"dns2"`
	Domain              string                                 `json:"domain"`
	Domainid            string                                 `json:"domainid"`
	Domainpath          string                                 `json:"domainpath"`
	Gateway             string                                 `json:"gateway"`
	Guestipaddress      string                                 `json:"guestipaddress"`
	Guestmacaddress     string                                 `json:"guestmacaddress"`
	Guestnetmask        string                                 `json:"guestnetmask"`
	Guestnetworkid      string                                 `json:"guestnetworkid"`
	Guestnetworkname    string                                 `json:"guestnetworkname"`
	Hasannotations      bool                                   `json:"hasannotations"`
	Healthcheckresults  []StopRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                   `json:"healthchecksfailed"`
	Hostcontrolstate    string                                 `json:"hostcontrolstate"`
	Hostid              string                                 `json:"hostid"`
	Hostname            string                                 `json:"hostname"`
	Hypervisor          string                                 `json:"hypervisor"`
	Id                  string                                 `json:"id"`
	Ip6dns1             string                                 `json:"ip6dns1"`
	Ip6dns2             string                                 `json:"ip6dns2"`
	Isredundantrouter   bool                                   `json:"isredundantrouter"`
	JobID               string                                 `json:"jobid"`
	Jobstatus           int                                    `json:"jobstatus"`
	Linklocalip         string                                 `json:"linklocalip"`
	Linklocalmacaddress string                                 `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                 `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                 `json:"linklocalnetworkid"`
	Name                string                                 `json:"name"`
	Networkdomain       string                                 `json:"networkdomain"`
	Nic                 []Nic                                  `json:"nic"`
	Podid               string                                 `json:"podid"`
	Podname             string                                 `json:"podname"`
	Project             string                                 `json:"project"`
	Projectid           string                                 `json:"projectid"`
	Publicip            string                                 `json:"publicip"`
	Publicmacaddress    string                                 `json:"publicmacaddress"`
	Publicnetmask       string                                 `json:"publicnetmask"`
	Publicnetworkid     string                                 `json:"publicnetworkid"`
	Redundantstate      string                                 `json:"redundantstate"`
	Requiresupgrade     bool                                   `json:"requiresupgrade"`
	Role                string                                 `json:"role"`
	Scriptsversion      string                                 `json:"scriptsversion"`
	Serviceofferingid   string                                 `json:"serviceofferingid"`
	Serviceofferingname string                                 `json:"serviceofferingname"`
	Softwareversion     string                                 `json:"softwareversion"`
	State               string                                 `json:"state"`
	Templateid          string                                 `json:"templateid"`
	Templatename        string                                 `json:"templatename"`
	Version             string                                 `json:"version"`
	Vpcid               string                                 `json:"vpcid"`
	Vpcname             string                                 `json:"vpcname"`
	Zoneid              string                                 `json:"zoneid"`
	Zonename            string                                 `json:"zonename"`
}

type StopRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}
