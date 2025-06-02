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

type NetscalerServiceIface interface {
	AddNetscalerLoadBalancer(p *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error)
	NewAddNetscalerLoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddNetscalerLoadBalancerParams
	ConfigureNetscalerLoadBalancer(p *ConfigureNetscalerLoadBalancerParams) (*NetscalerLoadBalancerResponse, error)
	NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams
	DeleteNetscalerControlCenter(p *DeleteNetscalerControlCenterParams) (*DeleteNetscalerControlCenterResponse, error)
	NewDeleteNetscalerControlCenterParams(id string) *DeleteNetscalerControlCenterParams
	DeleteNetscalerLoadBalancer(p *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error)
	NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams
	ListNetscalerControlCenter(p *ListNetscalerControlCenterParams) (*ListNetscalerControlCenterResponse, error)
	NewListNetscalerControlCenterParams() *ListNetscalerControlCenterParams
	ListNetscalerLoadBalancerNetworks(p *ListNetscalerLoadBalancerNetworksParams) (*ListNetscalerLoadBalancerNetworksResponse, error)
	NewListNetscalerLoadBalancerNetworksParams(lbdeviceid string) *ListNetscalerLoadBalancerNetworksParams
	GetNetscalerLoadBalancerNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error)
	ListNetscalerLoadBalancers(p *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error)
	NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams
	RegisterNetscalerControlCenter(p *RegisterNetscalerControlCenterParams) (*RegisterNetscalerControlCenterResponse, error)
	NewRegisterNetscalerControlCenterParams(ipaddress string, numretries int, password string, username string) *RegisterNetscalerControlCenterParams
	RegisterNetscalerServicePackage(p *RegisterNetscalerServicePackageParams) (*RegisterNetscalerServicePackageResponse, error)
	NewRegisterNetscalerServicePackageParams(description string, name string) *RegisterNetscalerServicePackageParams
}

type AddNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *AddNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["gslbprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("gslbprovider", vv)
	}
	if v, found := p.p["gslbproviderprivateip"]; found {
		u.Set("gslbproviderprivateip", v.(string))
	}
	if v, found := p.p["gslbproviderpublicip"]; found {
		u.Set("gslbproviderpublicip", v.(string))
	}
	if v, found := p.p["isexclusivegslbprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isexclusivegslbprovider", vv)
	}
	if v, found := p.p["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddNetscalerLoadBalancerParams) SetGslbprovider(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbprovider"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetGslbprovider() {
	if p.p != nil && p.p["gslbprovider"] != nil {
		delete(p.p, "gslbprovider")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetGslbprovider() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslbprovider"].(bool)
	return value, ok
}

func (p *AddNetscalerLoadBalancerParams) SetGslbproviderprivateip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbproviderprivateip"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetGslbproviderprivateip() {
	if p.p != nil && p.p["gslbproviderprivateip"] != nil {
		delete(p.p, "gslbproviderprivateip")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetGslbproviderprivateip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslbproviderprivateip"].(string)
	return value, ok
}

func (p *AddNetscalerLoadBalancerParams) SetGslbproviderpublicip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbproviderpublicip"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetGslbproviderpublicip() {
	if p.p != nil && p.p["gslbproviderpublicip"] != nil {
		delete(p.p, "gslbproviderpublicip")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetGslbproviderpublicip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslbproviderpublicip"].(string)
	return value, ok
}

func (p *AddNetscalerLoadBalancerParams) SetIsexclusivegslbprovider(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isexclusivegslbprovider"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetIsexclusivegslbprovider() {
	if p.p != nil && p.p["isexclusivegslbprovider"] != nil {
		delete(p.p, "isexclusivegslbprovider")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetIsexclusivegslbprovider() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isexclusivegslbprovider"].(bool)
	return value, ok
}

func (p *AddNetscalerLoadBalancerParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetNetworkdevicetype() {
	if p.p != nil && p.p["networkdevicetype"] != nil {
		delete(p.p, "networkdevicetype")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetNetworkdevicetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdevicetype"].(string)
	return value, ok
}

func (p *AddNetscalerLoadBalancerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddNetscalerLoadBalancerParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddNetscalerLoadBalancerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddNetscalerLoadBalancerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddNetscalerLoadBalancerParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddNetscalerLoadBalancerParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewAddNetscalerLoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddNetscalerLoadBalancerParams {
	p := &AddNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["networkdevicetype"] = networkdevicetype
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds a netscaler load balancer device
func (s *NetscalerService) AddNetscalerLoadBalancer(p *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newPostRequest("addNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetscalerLoadBalancerResponse
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

type AddNetscalerLoadBalancerResponse struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	JobID                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type ConfigureNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *ConfigureNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["inline"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("inline", vv)
	}
	if v, found := p.p["lbdevicecapacity"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("lbdevicecapacity", vv)
	}
	if v, found := p.p["lbdevicededicated"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lbdevicededicated", vv)
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["podids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("podids", vv)
	}
	return u
}

func (p *ConfigureNetscalerLoadBalancerParams) SetInline(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["inline"] = v
}

func (p *ConfigureNetscalerLoadBalancerParams) ResetInline() {
	if p.p != nil && p.p["inline"] != nil {
		delete(p.p, "inline")
	}
}

func (p *ConfigureNetscalerLoadBalancerParams) GetInline() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["inline"].(bool)
	return value, ok
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdevicecapacity(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdevicecapacity"] = v
}

func (p *ConfigureNetscalerLoadBalancerParams) ResetLbdevicecapacity() {
	if p.p != nil && p.p["lbdevicecapacity"] != nil {
		delete(p.p, "lbdevicecapacity")
	}
}

func (p *ConfigureNetscalerLoadBalancerParams) GetLbdevicecapacity() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbdevicecapacity"].(int64)
	return value, ok
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdevicededicated(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdevicededicated"] = v
}

func (p *ConfigureNetscalerLoadBalancerParams) ResetLbdevicededicated() {
	if p.p != nil && p.p["lbdevicededicated"] != nil {
		delete(p.p, "lbdevicededicated")
	}
}

func (p *ConfigureNetscalerLoadBalancerParams) GetLbdevicededicated() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbdevicededicated"].(bool)
	return value, ok
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
}

func (p *ConfigureNetscalerLoadBalancerParams) ResetLbdeviceid() {
	if p.p != nil && p.p["lbdeviceid"] != nil {
		delete(p.p, "lbdeviceid")
	}
}

func (p *ConfigureNetscalerLoadBalancerParams) GetLbdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbdeviceid"].(string)
	return value, ok
}

func (p *ConfigureNetscalerLoadBalancerParams) SetPodids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podids"] = v
}

func (p *ConfigureNetscalerLoadBalancerParams) ResetPodids() {
	if p.p != nil && p.p["podids"] != nil {
		delete(p.p, "podids")
	}
}

func (p *ConfigureNetscalerLoadBalancerParams) GetPodids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podids"].([]string)
	return value, ok
}

// You should always use this function to get a new ConfigureNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams {
	p := &ConfigureNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// configures a netscaler load balancer device
func (s *NetscalerService) ConfigureNetscalerLoadBalancer(p *ConfigureNetscalerLoadBalancerParams) (*NetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newPostRequest("configureNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r NetscalerLoadBalancerResponse
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

type NetscalerLoadBalancerResponse struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	JobID                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type DeleteNetscalerControlCenterParams struct {
	p map[string]interface{}
}

func (p *DeleteNetscalerControlCenterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetscalerControlCenterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteNetscalerControlCenterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteNetscalerControlCenterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetscalerControlCenterParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewDeleteNetscalerControlCenterParams(id string) *DeleteNetscalerControlCenterParams {
	p := &DeleteNetscalerControlCenterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete Netscaler Control Center
func (s *NetscalerService) DeleteNetscalerControlCenter(p *DeleteNetscalerControlCenterParams) (*DeleteNetscalerControlCenterResponse, error) {
	resp, err := s.cs.newPostRequest("deleteNetscalerControlCenter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetscalerControlCenterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteNetscalerControlCenterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteNetscalerControlCenterResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteNetscalerControlCenterResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	return u
}

func (p *DeleteNetscalerLoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
}

func (p *DeleteNetscalerLoadBalancerParams) ResetLbdeviceid() {
	if p.p != nil && p.p["lbdeviceid"] != nil {
		delete(p.p, "lbdeviceid")
	}
}

func (p *DeleteNetscalerLoadBalancerParams) GetLbdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams {
	p := &DeleteNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// delete a netscaler load balancer device
func (s *NetscalerService) DeleteNetscalerLoadBalancer(p *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newPostRequest("deleteNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetscalerLoadBalancerResponse
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

type DeleteNetscalerLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListNetscalerControlCenterParams struct {
	p map[string]interface{}
}

func (p *ListNetscalerControlCenterParams) toURLValues() url.Values {
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
	return u
}

func (p *ListNetscalerControlCenterParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetscalerControlCenterParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNetscalerControlCenterParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetscalerControlCenterParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetscalerControlCenterParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNetscalerControlCenterParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetscalerControlCenterParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetscalerControlCenterParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNetscalerControlCenterParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNetscalerControlCenterParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewListNetscalerControlCenterParams() *ListNetscalerControlCenterParams {
	p := &ListNetscalerControlCenterParams{}
	p.p = make(map[string]interface{})
	return p
}

// list control center
func (s *NetscalerService) ListNetscalerControlCenter(p *ListNetscalerControlCenterParams) (*ListNetscalerControlCenterResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerControlCenter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerControlCenterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetscalerControlCenterResponse struct {
	Count                  int                       `json:"count"`
	NetscalerControlCenter []*NetscalerControlCenter `json:"netscalercontrolcenter"`
}

type NetscalerControlCenter struct {
	Id         string `json:"id"`
	Ipaddress  string `json:"ipaddress"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Numretries string `json:"numretries"`
	Username   string `json:"username"`
	Uuid       string `json:"uuid"`
}

type ListNetscalerLoadBalancerNetworksParams struct {
	p map[string]interface{}
}

func (p *ListNetscalerLoadBalancerNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
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

func (p *ListNetscalerLoadBalancerNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetscalerLoadBalancerNetworksParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNetscalerLoadBalancerNetworksParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
}

func (p *ListNetscalerLoadBalancerNetworksParams) ResetLbdeviceid() {
	if p.p != nil && p.p["lbdeviceid"] != nil {
		delete(p.p, "lbdeviceid")
	}
}

func (p *ListNetscalerLoadBalancerNetworksParams) GetLbdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbdeviceid"].(string)
	return value, ok
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetscalerLoadBalancerNetworksParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNetscalerLoadBalancerNetworksParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetscalerLoadBalancerNetworksParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNetscalerLoadBalancerNetworksParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNetscalerLoadBalancerNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewListNetscalerLoadBalancerNetworksParams(lbdeviceid string) *ListNetscalerLoadBalancerNetworksParams {
	p := &ListNetscalerLoadBalancerNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetscalerService) GetNetscalerLoadBalancerNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error) {
	p := &ListNetscalerLoadBalancerNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["lbdeviceid"] = lbdeviceid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetscalerLoadBalancerNetworks(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.NetscalerLoadBalancerNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetscalerLoadBalancerNetworks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// lists network that are using a netscaler load balancer device
func (s *NetscalerService) ListNetscalerLoadBalancerNetworks(p *ListNetscalerLoadBalancerNetworksParams) (*ListNetscalerLoadBalancerNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerLoadBalancerNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerLoadBalancerNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetscalerLoadBalancerNetworksResponse struct {
	Count                         int                             `json:"count"`
	NetscalerLoadBalancerNetworks []*NetscalerLoadBalancerNetwork `json:"netscalerloadbalancernetwork"`
}

type NetscalerLoadBalancerNetwork struct {
	Account                     string                                `json:"account"`
	Aclid                       string                                `json:"aclid"`
	Aclname                     string                                `json:"aclname"`
	Acltype                     string                                `json:"acltype"`
	Asnumber                    int64                                 `json:"asnumber"`
	Asnumberid                  string                                `json:"asnumberid"`
	Associatednetwork           string                                `json:"associatednetwork"`
	Associatednetworkid         string                                `json:"associatednetworkid"`
	Bgppeers                    []interface{}                         `json:"bgppeers"`
	Broadcastdomaintype         string                                `json:"broadcastdomaintype"`
	Broadcasturi                string                                `json:"broadcasturi"`
	Canusefordeploy             bool                                  `json:"canusefordeploy"`
	Cidr                        string                                `json:"cidr"`
	Created                     string                                `json:"created"`
	Details                     map[string]string                     `json:"details"`
	Displaynetwork              bool                                  `json:"displaynetwork"`
	Displaytext                 string                                `json:"displaytext"`
	Dns1                        string                                `json:"dns1"`
	Dns2                        string                                `json:"dns2"`
	Domain                      string                                `json:"domain"`
	Domainid                    string                                `json:"domainid"`
	Domainpath                  string                                `json:"domainpath"`
	Egressdefaultpolicy         bool                                  `json:"egressdefaultpolicy"`
	Externalid                  string                                `json:"externalid"`
	Gateway                     string                                `json:"gateway"`
	Hasannotations              bool                                  `json:"hasannotations"`
	Icon                        interface{}                           `json:"icon"`
	Id                          string                                `json:"id"`
	Internetprotocol            string                                `json:"internetprotocol"`
	Ip4routes                   []interface{}                         `json:"ip4routes"`
	Ip4routing                  string                                `json:"ip4routing"`
	Ip6cidr                     string                                `json:"ip6cidr"`
	Ip6dns1                     string                                `json:"ip6dns1"`
	Ip6dns2                     string                                `json:"ip6dns2"`
	Ip6gateway                  string                                `json:"ip6gateway"`
	Ip6routes                   []interface{}                         `json:"ip6routes"`
	Ip6routing                  string                                `json:"ip6routing"`
	Isdefault                   bool                                  `json:"isdefault"`
	Ispersistent                bool                                  `json:"ispersistent"`
	Issystem                    bool                                  `json:"issystem"`
	JobID                       string                                `json:"jobid"`
	Jobstatus                   int                                   `json:"jobstatus"`
	Name                        string                                `json:"name"`
	Netmask                     string                                `json:"netmask"`
	Networkcidr                 string                                `json:"networkcidr"`
	Networkdomain               string                                `json:"networkdomain"`
	Networkofferingavailability string                                `json:"networkofferingavailability"`
	Networkofferingconservemode bool                                  `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                                `json:"networkofferingdisplaytext"`
	Networkofferingid           string                                `json:"networkofferingid"`
	Networkofferingname         string                                `json:"networkofferingname"`
	Physicalnetworkid           string                                `json:"physicalnetworkid"`
	Privatemtu                  int                                   `json:"privatemtu"`
	Project                     string                                `json:"project"`
	Projectid                   string                                `json:"projectid"`
	Publicmtu                   int                                   `json:"publicmtu"`
	Receivedbytes               int64                                 `json:"receivedbytes"`
	Redundantrouter             bool                                  `json:"redundantrouter"`
	Related                     string                                `json:"related"`
	Reservediprange             string                                `json:"reservediprange"`
	Restartrequired             bool                                  `json:"restartrequired"`
	Sentbytes                   int64                                 `json:"sentbytes"`
	Service                     []NetscalerLoadBalancerNetworkService `json:"service"`
	Specifyipranges             bool                                  `json:"specifyipranges"`
	State                       string                                `json:"state"`
	Strechedl2subnet            bool                                  `json:"strechedl2subnet"`
	Subdomainaccess             bool                                  `json:"subdomainaccess"`
	Supportsvmautoscaling       bool                                  `json:"supportsvmautoscaling"`
	Tags                        []Tags                                `json:"tags"`
	Traffictype                 string                                `json:"traffictype"`
	Tungstenvirtualrouteruuid   string                                `json:"tungstenvirtualrouteruuid"`
	Type                        string                                `json:"type"`
	Vlan                        string                                `json:"vlan"`
	Vpcid                       string                                `json:"vpcid"`
	Vpcname                     string                                `json:"vpcname"`
	Zoneid                      string                                `json:"zoneid"`
	Zonename                    string                                `json:"zonename"`
	Zonesnetworkspans           []interface{}                         `json:"zonesnetworkspans"`
}

type NetscalerLoadBalancerNetworkService struct {
	Capability []NetscalerLoadBalancerNetworkServiceCapability `json:"capability"`
	Name       string                                          `json:"name"`
	Provider   []NetscalerLoadBalancerNetworkServiceProvider   `json:"provider"`
}

type NetscalerLoadBalancerNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NetscalerLoadBalancerNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListNetscalerLoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListNetscalerLoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
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

func (p *ListNetscalerLoadBalancersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetscalerLoadBalancersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListNetscalerLoadBalancersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetscalerLoadBalancersParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
}

func (p *ListNetscalerLoadBalancersParams) ResetLbdeviceid() {
	if p.p != nil && p.p["lbdeviceid"] != nil {
		delete(p.p, "lbdeviceid")
	}
}

func (p *ListNetscalerLoadBalancersParams) GetLbdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbdeviceid"].(string)
	return value, ok
}

func (p *ListNetscalerLoadBalancersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetscalerLoadBalancersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListNetscalerLoadBalancersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetscalerLoadBalancersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetscalerLoadBalancersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNetscalerLoadBalancersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNetscalerLoadBalancersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListNetscalerLoadBalancersParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *ListNetscalerLoadBalancersParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetscalerLoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams {
	p := &ListNetscalerLoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists netscaler load balancer devices
func (s *NetscalerService) ListNetscalerLoadBalancers(p *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetscalerLoadBalancersResponse struct {
	Count                  int                      `json:"count"`
	NetscalerLoadBalancers []*NetscalerLoadBalancer `json:"netscalerloadbalancer"`
}

type NetscalerLoadBalancer struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	JobID                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type RegisterNetscalerControlCenterParams struct {
	p map[string]interface{}
}

func (p *RegisterNetscalerControlCenterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["numretries"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("numretries", vv)
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *RegisterNetscalerControlCenterParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *RegisterNetscalerControlCenterParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *RegisterNetscalerControlCenterParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *RegisterNetscalerControlCenterParams) SetNumretries(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["numretries"] = v
}

func (p *RegisterNetscalerControlCenterParams) ResetNumretries() {
	if p.p != nil && p.p["numretries"] != nil {
		delete(p.p, "numretries")
	}
}

func (p *RegisterNetscalerControlCenterParams) GetNumretries() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["numretries"].(int)
	return value, ok
}

func (p *RegisterNetscalerControlCenterParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *RegisterNetscalerControlCenterParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *RegisterNetscalerControlCenterParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *RegisterNetscalerControlCenterParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *RegisterNetscalerControlCenterParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *RegisterNetscalerControlCenterParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterNetscalerControlCenterParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewRegisterNetscalerControlCenterParams(ipaddress string, numretries int, password string, username string) *RegisterNetscalerControlCenterParams {
	p := &RegisterNetscalerControlCenterParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddress"] = ipaddress
	p.p["numretries"] = numretries
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Adds a netscaler control center device
func (s *NetscalerService) RegisterNetscalerControlCenter(p *RegisterNetscalerControlCenterParams) (*RegisterNetscalerControlCenterResponse, error) {
	resp, err := s.cs.newPostRequest("registerNetscalerControlCenter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterNetscalerControlCenterResponse
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

type RegisterNetscalerControlCenterResponse struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	JobID                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type RegisterNetscalerServicePackageParams struct {
	p map[string]interface{}
}

func (p *RegisterNetscalerServicePackageParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *RegisterNetscalerServicePackageParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *RegisterNetscalerServicePackageParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *RegisterNetscalerServicePackageParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *RegisterNetscalerServicePackageParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *RegisterNetscalerServicePackageParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *RegisterNetscalerServicePackageParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterNetscalerServicePackageParams instance,
// as then you are sure you have configured all required params
func (s *NetscalerService) NewRegisterNetscalerServicePackageParams(description string, name string) *RegisterNetscalerServicePackageParams {
	p := &RegisterNetscalerServicePackageParams{}
	p.p = make(map[string]interface{})
	p.p["description"] = description
	p.p["name"] = name
	return p
}

// Registers NCC Service Package
func (s *NetscalerService) RegisterNetscalerServicePackage(p *RegisterNetscalerServicePackageParams) (*RegisterNetscalerServicePackageResponse, error) {
	resp, err := s.cs.newPostRequest("registerNetscalerServicePackage", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterNetscalerServicePackageResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterNetscalerServicePackageResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
}
