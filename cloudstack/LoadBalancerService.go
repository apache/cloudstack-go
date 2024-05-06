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

type LoadBalancerServiceIface interface {
	AddNetscalerLoadBalancer(p *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error)
	NewAddNetscalerLoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddNetscalerLoadBalancerParams
	AssignCertToLoadBalancer(p *AssignCertToLoadBalancerParams) (*AssignCertToLoadBalancerResponse, error)
	NewAssignCertToLoadBalancerParams(certid string, lbruleid string) *AssignCertToLoadBalancerParams
	AssignToGlobalLoadBalancerRule(p *AssignToGlobalLoadBalancerRuleParams) (*AssignToGlobalLoadBalancerRuleResponse, error)
	NewAssignToGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *AssignToGlobalLoadBalancerRuleParams
	AssignToLoadBalancerRule(p *AssignToLoadBalancerRuleParams) (*AssignToLoadBalancerRuleResponse, error)
	NewAssignToLoadBalancerRuleParams(id string) *AssignToLoadBalancerRuleParams
	ConfigureNetscalerLoadBalancer(p *ConfigureNetscalerLoadBalancerParams) (*NetscalerLoadBalancerResponse, error)
	NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams
	CreateGlobalLoadBalancerRule(p *CreateGlobalLoadBalancerRuleParams) (*CreateGlobalLoadBalancerRuleResponse, error)
	NewCreateGlobalLoadBalancerRuleParams(gslbdomainname string, gslbservicetype string, name string, regionid int) *CreateGlobalLoadBalancerRuleParams
	CreateLBHealthCheckPolicy(p *CreateLBHealthCheckPolicyParams) (*CreateLBHealthCheckPolicyResponse, error)
	NewCreateLBHealthCheckPolicyParams(lbruleid string) *CreateLBHealthCheckPolicyParams
	CreateLBStickinessPolicy(p *CreateLBStickinessPolicyParams) (*CreateLBStickinessPolicyResponse, error)
	NewCreateLBStickinessPolicyParams(lbruleid string, methodname string, name string) *CreateLBStickinessPolicyParams
	CreateLoadBalancer(p *CreateLoadBalancerParams) (*CreateLoadBalancerResponse, error)
	NewCreateLoadBalancerParams(algorithm string, instanceport int, name string, networkid string, scheme string, sourceipaddressnetworkid string, sourceport int) *CreateLoadBalancerParams
	CreateLoadBalancerRule(p *CreateLoadBalancerRuleParams) (*CreateLoadBalancerRuleResponse, error)
	NewCreateLoadBalancerRuleParams(algorithm string, name string, privateport int, publicport int) *CreateLoadBalancerRuleParams
	DeleteGlobalLoadBalancerRule(p *DeleteGlobalLoadBalancerRuleParams) (*DeleteGlobalLoadBalancerRuleResponse, error)
	NewDeleteGlobalLoadBalancerRuleParams(id string) *DeleteGlobalLoadBalancerRuleParams
	DeleteLBHealthCheckPolicy(p *DeleteLBHealthCheckPolicyParams) (*DeleteLBHealthCheckPolicyResponse, error)
	NewDeleteLBHealthCheckPolicyParams(id string) *DeleteLBHealthCheckPolicyParams
	DeleteLBStickinessPolicy(p *DeleteLBStickinessPolicyParams) (*DeleteLBStickinessPolicyResponse, error)
	NewDeleteLBStickinessPolicyParams(id string) *DeleteLBStickinessPolicyParams
	DeleteLoadBalancer(p *DeleteLoadBalancerParams) (*DeleteLoadBalancerResponse, error)
	NewDeleteLoadBalancerParams(id string) *DeleteLoadBalancerParams
	DeleteLoadBalancerRule(p *DeleteLoadBalancerRuleParams) (*DeleteLoadBalancerRuleResponse, error)
	NewDeleteLoadBalancerRuleParams(id string) *DeleteLoadBalancerRuleParams
	DeleteNetscalerLoadBalancer(p *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error)
	NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams
	DeleteSslCert(p *DeleteSslCertParams) (*DeleteSslCertResponse, error)
	NewDeleteSslCertParams(id string) *DeleteSslCertParams
	ListGlobalLoadBalancerRules(p *ListGlobalLoadBalancerRulesParams) (*ListGlobalLoadBalancerRulesResponse, error)
	NewListGlobalLoadBalancerRulesParams() *ListGlobalLoadBalancerRulesParams
	GetGlobalLoadBalancerRuleID(keyword string, opts ...OptionFunc) (string, int, error)
	GetGlobalLoadBalancerRuleByName(name string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error)
	GetGlobalLoadBalancerRuleByID(id string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error)
	ListLBHealthCheckPolicies(p *ListLBHealthCheckPoliciesParams) (*ListLBHealthCheckPoliciesResponse, error)
	NewListLBHealthCheckPoliciesParams() *ListLBHealthCheckPoliciesParams
	GetLBHealthCheckPolicyByID(id string, opts ...OptionFunc) (*LBHealthCheckPolicy, int, error)
	ListLBStickinessPolicies(p *ListLBStickinessPoliciesParams) (*ListLBStickinessPoliciesResponse, error)
	NewListLBStickinessPoliciesParams() *ListLBStickinessPoliciesParams
	GetLBStickinessPolicyByID(id string, opts ...OptionFunc) (*LBStickinessPolicy, int, error)
	ListLoadBalancerRuleInstances(p *ListLoadBalancerRuleInstancesParams) (*ListLoadBalancerRuleInstancesResponse, error)
	NewListLoadBalancerRuleInstancesParams(id string) *ListLoadBalancerRuleInstancesParams
	GetLoadBalancerRuleInstanceByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error)
	ListLoadBalancerRules(p *ListLoadBalancerRulesParams) (*ListLoadBalancerRulesResponse, error)
	NewListLoadBalancerRulesParams() *ListLoadBalancerRulesParams
	GetLoadBalancerRuleID(name string, opts ...OptionFunc) (string, int, error)
	GetLoadBalancerRuleByName(name string, opts ...OptionFunc) (*LoadBalancerRule, int, error)
	GetLoadBalancerRuleByID(id string, opts ...OptionFunc) (*LoadBalancerRule, int, error)
	ListLoadBalancers(p *ListLoadBalancersParams) (*ListLoadBalancersResponse, error)
	NewListLoadBalancersParams() *ListLoadBalancersParams
	GetLoadBalancerID(name string, opts ...OptionFunc) (string, int, error)
	GetLoadBalancerByName(name string, opts ...OptionFunc) (*LoadBalancer, int, error)
	GetLoadBalancerByID(id string, opts ...OptionFunc) (*LoadBalancer, int, error)
	ListNetscalerLoadBalancers(p *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error)
	NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams
	ListSslCerts(p *ListSslCertsParams) (*ListSslCertsResponse, error)
	NewListSslCertsParams() *ListSslCertsParams
	RemoveCertFromLoadBalancer(p *RemoveCertFromLoadBalancerParams) (*RemoveCertFromLoadBalancerResponse, error)
	NewRemoveCertFromLoadBalancerParams(lbruleid string) *RemoveCertFromLoadBalancerParams
	RemoveFromGlobalLoadBalancerRule(p *RemoveFromGlobalLoadBalancerRuleParams) (*RemoveFromGlobalLoadBalancerRuleResponse, error)
	NewRemoveFromGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *RemoveFromGlobalLoadBalancerRuleParams
	RemoveFromLoadBalancerRule(p *RemoveFromLoadBalancerRuleParams) (*RemoveFromLoadBalancerRuleResponse, error)
	NewRemoveFromLoadBalancerRuleParams(id string) *RemoveFromLoadBalancerRuleParams
	UpdateGlobalLoadBalancerRule(p *UpdateGlobalLoadBalancerRuleParams) (*UpdateGlobalLoadBalancerRuleResponse, error)
	NewUpdateGlobalLoadBalancerRuleParams(id string) *UpdateGlobalLoadBalancerRuleParams
	UpdateLBHealthCheckPolicy(p *UpdateLBHealthCheckPolicyParams) (*UpdateLBHealthCheckPolicyResponse, error)
	NewUpdateLBHealthCheckPolicyParams(id string) *UpdateLBHealthCheckPolicyParams
	UpdateLBStickinessPolicy(p *UpdateLBStickinessPolicyParams) (*UpdateLBStickinessPolicyResponse, error)
	NewUpdateLBStickinessPolicyParams(id string) *UpdateLBStickinessPolicyParams
	UpdateLoadBalancer(p *UpdateLoadBalancerParams) (*UpdateLoadBalancerResponse, error)
	NewUpdateLoadBalancerParams(id string) *UpdateLoadBalancerParams
	UpdateLoadBalancerRule(p *UpdateLoadBalancerRuleParams) (*UpdateLoadBalancerRuleResponse, error)
	NewUpdateLoadBalancerRuleParams(id string) *UpdateLoadBalancerRuleParams
	UploadSslCert(p *UploadSslCertParams) (*UploadSslCertResponse, error)
	NewUploadSslCertParams(certificate string, name string, privatekey string) *UploadSslCertParams
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
func (s *LoadBalancerService) NewAddNetscalerLoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddNetscalerLoadBalancerParams {
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
func (s *LoadBalancerService) AddNetscalerLoadBalancer(p *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("addNetscalerLoadBalancer", p.toURLValues())
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

type AssignCertToLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *AssignCertToLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["certid"]; found {
		u.Set("certid", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (p *AssignCertToLoadBalancerParams) SetCertid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certid"] = v
}

func (p *AssignCertToLoadBalancerParams) ResetCertid() {
	if p.p != nil && p.p["certid"] != nil {
		delete(p.p, "certid")
	}
}

func (p *AssignCertToLoadBalancerParams) GetCertid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["certid"].(string)
	return value, ok
}

func (p *AssignCertToLoadBalancerParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *AssignCertToLoadBalancerParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *AssignCertToLoadBalancerParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

// You should always use this function to get a new AssignCertToLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignCertToLoadBalancerParams(certid string, lbruleid string) *AssignCertToLoadBalancerParams {
	p := &AssignCertToLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["certid"] = certid
	p.p["lbruleid"] = lbruleid
	return p
}

// Assigns a certificate to a load balancer rule
func (s *LoadBalancerService) AssignCertToLoadBalancer(p *AssignCertToLoadBalancerParams) (*AssignCertToLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("assignCertToLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignCertToLoadBalancerResponse
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

type AssignCertToLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AssignToGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *AssignToGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["gslblbruleweightsmap"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("gslblbruleweightsmap[%d].key", i), k)
			u.Set(fmt.Sprintf("gslblbruleweightsmap[%d].value", i), m[k])
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["loadbalancerrulelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("loadbalancerrulelist", vv)
	}
	return u
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetGslblbruleweightsmap(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbruleweightsmap"] = v
}

func (p *AssignToGlobalLoadBalancerRuleParams) ResetGslblbruleweightsmap() {
	if p.p != nil && p.p["gslblbruleweightsmap"] != nil {
		delete(p.p, "gslblbruleweightsmap")
	}
}

func (p *AssignToGlobalLoadBalancerRuleParams) GetGslblbruleweightsmap() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslblbruleweightsmap"].(map[string]string)
	return value, ok
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *AssignToGlobalLoadBalancerRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *AssignToGlobalLoadBalancerRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetLoadbalancerrulelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerrulelist"] = v
}

func (p *AssignToGlobalLoadBalancerRuleParams) ResetLoadbalancerrulelist() {
	if p.p != nil && p.p["loadbalancerrulelist"] != nil {
		delete(p.p, "loadbalancerrulelist")
	}
}

func (p *AssignToGlobalLoadBalancerRuleParams) GetLoadbalancerrulelist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["loadbalancerrulelist"].([]string)
	return value, ok
}

// You should always use this function to get a new AssignToGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignToGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *AssignToGlobalLoadBalancerRuleParams {
	p := &AssignToGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["loadbalancerrulelist"] = loadbalancerrulelist
	return p
}

// Assign load balancer rule or list of load balancer rules to a global load balancer rules.
func (s *LoadBalancerService) AssignToGlobalLoadBalancerRule(p *AssignToGlobalLoadBalancerRuleParams) (*AssignToGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("assignToGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignToGlobalLoadBalancerRuleResponse
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

type AssignToGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AssignToLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *AssignToLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("virtualmachineids", vv)
	}
	if v, found := p.p["vmidipmap"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("vmidipmap[%d].key", i), k)
			u.Set(fmt.Sprintf("vmidipmap[%d].value", i), m[k])
		}
	}
	return u
}

func (p *AssignToLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *AssignToLoadBalancerRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *AssignToLoadBalancerRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *AssignToLoadBalancerRuleParams) SetVirtualmachineids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineids"] = v
}

func (p *AssignToLoadBalancerRuleParams) ResetVirtualmachineids() {
	if p.p != nil && p.p["virtualmachineids"] != nil {
		delete(p.p, "virtualmachineids")
	}
}

func (p *AssignToLoadBalancerRuleParams) GetVirtualmachineids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineids"].([]string)
	return value, ok
}

func (p *AssignToLoadBalancerRuleParams) SetVmidipmap(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmidipmap"] = v
}

func (p *AssignToLoadBalancerRuleParams) ResetVmidipmap() {
	if p.p != nil && p.p["vmidipmap"] != nil {
		delete(p.p, "vmidipmap")
	}
}

func (p *AssignToLoadBalancerRuleParams) GetVmidipmap() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmidipmap"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new AssignToLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignToLoadBalancerRuleParams(id string) *AssignToLoadBalancerRuleParams {
	p := &AssignToLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Assigns virtual machine or a list of virtual machines to a load balancer rule.
func (s *LoadBalancerService) AssignToLoadBalancerRule(p *AssignToLoadBalancerRuleParams) (*AssignToLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("assignToLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignToLoadBalancerRuleResponse
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

type AssignToLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
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
func (s *LoadBalancerService) NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams {
	p := &ConfigureNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// configures a netscaler load balancer device
func (s *LoadBalancerService) ConfigureNetscalerLoadBalancer(p *ConfigureNetscalerLoadBalancerParams) (*NetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("configureNetscalerLoadBalancer", p.toURLValues())
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

type CreateGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *CreateGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["gslbdomainname"]; found {
		u.Set("gslbdomainname", v.(string))
	}
	if v, found := p.p["gslblbmethod"]; found {
		u.Set("gslblbmethod", v.(string))
	}
	if v, found := p.p["gslbservicetype"]; found {
		u.Set("gslbservicetype", v.(string))
	}
	if v, found := p.p["gslbstickysessionmethodname"]; found {
		u.Set("gslbstickysessionmethodname", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	return u
}

func (p *CreateGlobalLoadBalancerRuleParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateGlobalLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateGlobalLoadBalancerRuleParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbdomainname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbdomainname"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetGslbdomainname() {
	if p.p != nil && p.p["gslbdomainname"] != nil {
		delete(p.p, "gslbdomainname")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetGslbdomainname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslbdomainname"].(string)
	return value, ok
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslblbmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbmethod"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetGslblbmethod() {
	if p.p != nil && p.p["gslblbmethod"] != nil {
		delete(p.p, "gslblbmethod")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetGslblbmethod() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslblbmethod"].(string)
	return value, ok
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbservicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbservicetype"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetGslbservicetype() {
	if p.p != nil && p.p["gslbservicetype"] != nil {
		delete(p.p, "gslbservicetype")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetGslbservicetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslbservicetype"].(string)
	return value, ok
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbstickysessionmethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbstickysessionmethodname"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetGslbstickysessionmethodname() {
	if p.p != nil && p.p["gslbstickysessionmethodname"] != nil {
		delete(p.p, "gslbstickysessionmethodname")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetGslbstickysessionmethodname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslbstickysessionmethodname"].(string)
	return value, ok
}

func (p *CreateGlobalLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateGlobalLoadBalancerRuleParams) SetRegionid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["regionid"] = v
}

func (p *CreateGlobalLoadBalancerRuleParams) ResetRegionid() {
	if p.p != nil && p.p["regionid"] != nil {
		delete(p.p, "regionid")
	}
}

func (p *CreateGlobalLoadBalancerRuleParams) GetRegionid() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["regionid"].(int)
	return value, ok
}

// You should always use this function to get a new CreateGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateGlobalLoadBalancerRuleParams(gslbdomainname string, gslbservicetype string, name string, regionid int) *CreateGlobalLoadBalancerRuleParams {
	p := &CreateGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["gslbdomainname"] = gslbdomainname
	p.p["gslbservicetype"] = gslbservicetype
	p.p["name"] = name
	p.p["regionid"] = regionid
	return p
}

// Creates a global load balancer rule
func (s *LoadBalancerService) CreateGlobalLoadBalancerRule(p *CreateGlobalLoadBalancerRuleParams) (*CreateGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("createGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateGlobalLoadBalancerRuleResponse
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

type CreateGlobalLoadBalancerRuleResponse struct {
	Account                     string                                                 `json:"account"`
	Description                 string                                                 `json:"description"`
	Domain                      string                                                 `json:"domain"`
	Domainid                    string                                                 `json:"domainid"`
	Gslbdomainname              string                                                 `json:"gslbdomainname"`
	Gslblbmethod                string                                                 `json:"gslblbmethod"`
	Gslbservicetype             string                                                 `json:"gslbservicetype"`
	Gslbstickysessionmethodname string                                                 `json:"gslbstickysessionmethodname"`
	Id                          string                                                 `json:"id"`
	JobID                       string                                                 `json:"jobid"`
	Jobstatus                   int                                                    `json:"jobstatus"`
	Loadbalancerrule            []CreateGlobalLoadBalancerRuleResponseLoadbalancerrule `json:"loadbalancerrule"`
	Name                        string                                                 `json:"name"`
	Project                     string                                                 `json:"project"`
	Projectid                   string                                                 `json:"projectid"`
	Regionid                    int                                                    `json:"regionid"`
}

type CreateGlobalLoadBalancerRuleResponseLoadbalancerrule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type CreateLBHealthCheckPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["healthythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("healthythreshold", vv)
	}
	if v, found := p.p["intervaltime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("intervaltime", vv)
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["pingpath"]; found {
		u.Set("pingpath", v.(string))
	}
	if v, found := p.p["responsetimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("responsetimeout", vv)
	}
	if v, found := p.p["unhealthythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("unhealthythreshold", vv)
	}
	return u
}

func (p *CreateLBHealthCheckPolicyParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateLBHealthCheckPolicyParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateLBHealthCheckPolicyParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateLBHealthCheckPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateLBHealthCheckPolicyParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateLBHealthCheckPolicyParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateLBHealthCheckPolicyParams) SetHealthythreshold(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["healthythreshold"] = v
}

func (p *CreateLBHealthCheckPolicyParams) ResetHealthythreshold() {
	if p.p != nil && p.p["healthythreshold"] != nil {
		delete(p.p, "healthythreshold")
	}
}

func (p *CreateLBHealthCheckPolicyParams) GetHealthythreshold() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["healthythreshold"].(int)
	return value, ok
}

func (p *CreateLBHealthCheckPolicyParams) SetIntervaltime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltime"] = v
}

func (p *CreateLBHealthCheckPolicyParams) ResetIntervaltime() {
	if p.p != nil && p.p["intervaltime"] != nil {
		delete(p.p, "intervaltime")
	}
}

func (p *CreateLBHealthCheckPolicyParams) GetIntervaltime() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["intervaltime"].(int)
	return value, ok
}

func (p *CreateLBHealthCheckPolicyParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *CreateLBHealthCheckPolicyParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *CreateLBHealthCheckPolicyParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *CreateLBHealthCheckPolicyParams) SetPingpath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pingpath"] = v
}

func (p *CreateLBHealthCheckPolicyParams) ResetPingpath() {
	if p.p != nil && p.p["pingpath"] != nil {
		delete(p.p, "pingpath")
	}
}

func (p *CreateLBHealthCheckPolicyParams) GetPingpath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pingpath"].(string)
	return value, ok
}

func (p *CreateLBHealthCheckPolicyParams) SetResponsetimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["responsetimeout"] = v
}

func (p *CreateLBHealthCheckPolicyParams) ResetResponsetimeout() {
	if p.p != nil && p.p["responsetimeout"] != nil {
		delete(p.p, "responsetimeout")
	}
}

func (p *CreateLBHealthCheckPolicyParams) GetResponsetimeout() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["responsetimeout"].(int)
	return value, ok
}

func (p *CreateLBHealthCheckPolicyParams) SetUnhealthythreshold(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["unhealthythreshold"] = v
}

func (p *CreateLBHealthCheckPolicyParams) ResetUnhealthythreshold() {
	if p.p != nil && p.p["unhealthythreshold"] != nil {
		delete(p.p, "unhealthythreshold")
	}
}

func (p *CreateLBHealthCheckPolicyParams) GetUnhealthythreshold() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["unhealthythreshold"].(int)
	return value, ok
}

// You should always use this function to get a new CreateLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLBHealthCheckPolicyParams(lbruleid string) *CreateLBHealthCheckPolicyParams {
	p := &CreateLBHealthCheckPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// Creates a load balancer health check policy
func (s *LoadBalancerService) CreateLBHealthCheckPolicy(p *CreateLBHealthCheckPolicyParams) (*CreateLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("createLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLBHealthCheckPolicyResponse
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

type CreateLBHealthCheckPolicyResponse struct {
	Account           string                                               `json:"account"`
	Domain            string                                               `json:"domain"`
	Domainid          string                                               `json:"domainid"`
	Healthcheckpolicy []CreateLBHealthCheckPolicyResponseHealthcheckpolicy `json:"healthcheckpolicy"`
	JobID             string                                               `json:"jobid"`
	Jobstatus         int                                                  `json:"jobstatus"`
	Lbruleid          string                                               `json:"lbruleid"`
	Zoneid            string                                               `json:"zoneid"`
}

type CreateLBHealthCheckPolicyResponseHealthcheckpolicy struct {
	Description             string `json:"description"`
	Fordisplay              bool   `json:"fordisplay"`
	Healthcheckinterval     int    `json:"healthcheckinterval"`
	Healthcheckthresshold   int    `json:"healthcheckthresshold"`
	Id                      string `json:"id"`
	Pingpath                string `json:"pingpath"`
	Responsetime            int    `json:"responsetime"`
	State                   string `json:"state"`
	Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
}

type CreateLBStickinessPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["methodname"]; found {
		u.Set("methodname", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["param"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("param[%d].key", i), k)
			u.Set(fmt.Sprintf("param[%d].value", i), m[k])
		}
	}
	return u
}

func (p *CreateLBStickinessPolicyParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateLBStickinessPolicyParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateLBStickinessPolicyParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateLBStickinessPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateLBStickinessPolicyParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateLBStickinessPolicyParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateLBStickinessPolicyParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *CreateLBStickinessPolicyParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *CreateLBStickinessPolicyParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *CreateLBStickinessPolicyParams) SetMethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["methodname"] = v
}

func (p *CreateLBStickinessPolicyParams) ResetMethodname() {
	if p.p != nil && p.p["methodname"] != nil {
		delete(p.p, "methodname")
	}
}

func (p *CreateLBStickinessPolicyParams) GetMethodname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["methodname"].(string)
	return value, ok
}

func (p *CreateLBStickinessPolicyParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateLBStickinessPolicyParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateLBStickinessPolicyParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateLBStickinessPolicyParams) SetParam(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["param"] = v
}

func (p *CreateLBStickinessPolicyParams) ResetParam() {
	if p.p != nil && p.p["param"] != nil {
		delete(p.p, "param")
	}
}

func (p *CreateLBStickinessPolicyParams) GetParam() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["param"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new CreateLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLBStickinessPolicyParams(lbruleid string, methodname string, name string) *CreateLBStickinessPolicyParams {
	p := &CreateLBStickinessPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	p.p["methodname"] = methodname
	p.p["name"] = name
	return p
}

// Creates a load balancer stickiness policy
func (s *LoadBalancerService) CreateLBStickinessPolicy(p *CreateLBStickinessPolicyParams) (*CreateLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("createLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLBStickinessPolicyResponse
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

type CreateLBStickinessPolicyResponse struct {
	Account          string                                             `json:"account"`
	Description      string                                             `json:"description"`
	Domain           string                                             `json:"domain"`
	Domainid         string                                             `json:"domainid"`
	JobID            string                                             `json:"jobid"`
	Jobstatus        int                                                `json:"jobstatus"`
	Lbruleid         string                                             `json:"lbruleid"`
	Name             string                                             `json:"name"`
	State            string                                             `json:"state"`
	Stickinesspolicy []CreateLBStickinessPolicyResponseStickinesspolicy `json:"stickinesspolicy"`
	Zoneid           string                                             `json:"zoneid"`
}

type CreateLBStickinessPolicyResponseStickinesspolicy struct {
	Description string            `json:"description"`
	Fordisplay  bool              `json:"fordisplay"`
	Id          string            `json:"id"`
	Methodname  string            `json:"methodname"`
	Name        string            `json:"name"`
	Params      map[string]string `json:"params"`
	State       string            `json:"state"`
}

type CreateLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *CreateLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["instanceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("instanceport", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["scheme"]; found {
		u.Set("scheme", v.(string))
	}
	if v, found := p.p["sourceipaddress"]; found {
		u.Set("sourceipaddress", v.(string))
	}
	if v, found := p.p["sourceipaddressnetworkid"]; found {
		u.Set("sourceipaddressnetworkid", v.(string))
	}
	if v, found := p.p["sourceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sourceport", vv)
	}
	return u
}

func (p *CreateLoadBalancerParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
}

func (p *CreateLoadBalancerParams) ResetAlgorithm() {
	if p.p != nil && p.p["algorithm"] != nil {
		delete(p.p, "algorithm")
	}
}

func (p *CreateLoadBalancerParams) GetAlgorithm() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["algorithm"].(string)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateLoadBalancerParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateLoadBalancerParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateLoadBalancerParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateLoadBalancerParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetInstanceport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["instanceport"] = v
}

func (p *CreateLoadBalancerParams) ResetInstanceport() {
	if p.p != nil && p.p["instanceport"] != nil {
		delete(p.p, "instanceport")
	}
}

func (p *CreateLoadBalancerParams) GetInstanceport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["instanceport"].(int)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateLoadBalancerParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateLoadBalancerParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreateLoadBalancerParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreateLoadBalancerParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetScheme(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scheme"] = v
}

func (p *CreateLoadBalancerParams) ResetScheme() {
	if p.p != nil && p.p["scheme"] != nil {
		delete(p.p, "scheme")
	}
}

func (p *CreateLoadBalancerParams) GetScheme() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scheme"].(string)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetSourceipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddress"] = v
}

func (p *CreateLoadBalancerParams) ResetSourceipaddress() {
	if p.p != nil && p.p["sourceipaddress"] != nil {
		delete(p.p, "sourceipaddress")
	}
}

func (p *CreateLoadBalancerParams) GetSourceipaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourceipaddress"].(string)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetSourceipaddressnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddressnetworkid"] = v
}

func (p *CreateLoadBalancerParams) ResetSourceipaddressnetworkid() {
	if p.p != nil && p.p["sourceipaddressnetworkid"] != nil {
		delete(p.p, "sourceipaddressnetworkid")
	}
}

func (p *CreateLoadBalancerParams) GetSourceipaddressnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourceipaddressnetworkid"].(string)
	return value, ok
}

func (p *CreateLoadBalancerParams) SetSourceport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceport"] = v
}

func (p *CreateLoadBalancerParams) ResetSourceport() {
	if p.p != nil && p.p["sourceport"] != nil {
		delete(p.p, "sourceport")
	}
}

func (p *CreateLoadBalancerParams) GetSourceport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourceport"].(int)
	return value, ok
}

// You should always use this function to get a new CreateLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLoadBalancerParams(algorithm string, instanceport int, name string, networkid string, scheme string, sourceipaddressnetworkid string, sourceport int) *CreateLoadBalancerParams {
	p := &CreateLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["algorithm"] = algorithm
	p.p["instanceport"] = instanceport
	p.p["name"] = name
	p.p["networkid"] = networkid
	p.p["scheme"] = scheme
	p.p["sourceipaddressnetworkid"] = sourceipaddressnetworkid
	p.p["sourceport"] = sourceport
	return p
}

// Creates an internal load balancer
func (s *LoadBalancerService) CreateLoadBalancer(p *CreateLoadBalancerParams) (*CreateLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLoadBalancerResponse
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

type CreateLoadBalancerResponse struct {
	Account                  string                                           `json:"account"`
	Algorithm                string                                           `json:"algorithm"`
	Description              string                                           `json:"description"`
	Domain                   string                                           `json:"domain"`
	Domainid                 string                                           `json:"domainid"`
	Fordisplay               bool                                             `json:"fordisplay"`
	Id                       string                                           `json:"id"`
	JobID                    string                                           `json:"jobid"`
	Jobstatus                int                                              `json:"jobstatus"`
	Loadbalancerinstance     []CreateLoadBalancerResponseLoadbalancerinstance `json:"loadbalancerinstance"`
	Loadbalancerrule         []CreateLoadBalancerResponseLoadbalancerrule     `json:"loadbalancerrule"`
	Name                     string                                           `json:"name"`
	Networkid                string                                           `json:"networkid"`
	Project                  string                                           `json:"project"`
	Projectid                string                                           `json:"projectid"`
	Sourceipaddress          string                                           `json:"sourceipaddress"`
	Sourceipaddressnetworkid string                                           `json:"sourceipaddressnetworkid"`
	Tags                     []Tags                                           `json:"tags"`
}

type CreateLoadBalancerResponseLoadbalancerrule struct {
	Instanceport int    `json:"instanceport"`
	Sourceport   int    `json:"sourceport"`
	State        string `json:"state"`
}

type CreateLoadBalancerResponseLoadbalancerinstance struct {
	Id        string `json:"id"`
	Ipaddress string `json:"ipaddress"`
	Name      string `json:"name"`
	State     string `json:"state"`
}

type CreateLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *CreateLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := p.p["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	if v, found := p.p["publicport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicport", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateLoadBalancerRuleParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateLoadBalancerRuleParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetAlgorithm() {
	if p.p != nil && p.p["algorithm"] != nil {
		delete(p.p, "algorithm")
	}
}

func (p *CreateLoadBalancerRuleParams) GetAlgorithm() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["algorithm"].(string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *CreateLoadBalancerRuleParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateLoadBalancerRuleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateLoadBalancerRuleParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateLoadBalancerRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateLoadBalancerRuleParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreateLoadBalancerRuleParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetOpenfirewall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["openfirewall"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetOpenfirewall() {
	if p.p != nil && p.p["openfirewall"] != nil {
		delete(p.p, "openfirewall")
	}
}

func (p *CreateLoadBalancerRuleParams) GetOpenfirewall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["openfirewall"].(bool)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetPrivateport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateport"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetPrivateport() {
	if p.p != nil && p.p["privateport"] != nil {
		delete(p.p, "privateport")
	}
}

func (p *CreateLoadBalancerRuleParams) GetPrivateport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privateport"].(int)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *CreateLoadBalancerRuleParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetPublicipid() {
	if p.p != nil && p.p["publicipid"] != nil {
		delete(p.p, "publicipid")
	}
}

func (p *CreateLoadBalancerRuleParams) GetPublicipid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicipid"].(string)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetPublicport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicport"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetPublicport() {
	if p.p != nil && p.p["publicport"] != nil {
		delete(p.p, "publicport")
	}
}

func (p *CreateLoadBalancerRuleParams) GetPublicport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicport"].(int)
	return value, ok
}

func (p *CreateLoadBalancerRuleParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateLoadBalancerRuleParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateLoadBalancerRuleParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLoadBalancerRuleParams(algorithm string, name string, privateport int, publicport int) *CreateLoadBalancerRuleParams {
	p := &CreateLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["algorithm"] = algorithm
	p.p["name"] = name
	p.p["privateport"] = privateport
	p.p["publicport"] = publicport
	return p
}

// Creates a load balancer rule
func (s *LoadBalancerService) CreateLoadBalancerRule(p *CreateLoadBalancerRuleParams) (*CreateLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLoadBalancerRuleResponse
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

type CreateLoadBalancerRuleResponse struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type DeleteGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteGlobalLoadBalancerRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteGlobalLoadBalancerRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteGlobalLoadBalancerRuleParams(id string) *DeleteGlobalLoadBalancerRuleParams {
	p := &DeleteGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a global load balancer rule.
func (s *LoadBalancerService) DeleteGlobalLoadBalancerRule(p *DeleteGlobalLoadBalancerRuleParams) (*DeleteGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteGlobalLoadBalancerRuleResponse
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

type DeleteGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLBHealthCheckPolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLBHealthCheckPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteLBHealthCheckPolicyParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteLBHealthCheckPolicyParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLBHealthCheckPolicyParams(id string) *DeleteLBHealthCheckPolicyParams {
	p := &DeleteLBHealthCheckPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer health check policy.
func (s *LoadBalancerService) DeleteLBHealthCheckPolicy(p *DeleteLBHealthCheckPolicyParams) (*DeleteLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLBHealthCheckPolicyResponse
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

type DeleteLBHealthCheckPolicyResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLBStickinessPolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLBStickinessPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteLBStickinessPolicyParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteLBStickinessPolicyParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLBStickinessPolicyParams(id string) *DeleteLBStickinessPolicyParams {
	p := &DeleteLBStickinessPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer stickiness policy.
func (s *LoadBalancerService) DeleteLBStickinessPolicy(p *DeleteLBStickinessPolicyParams) (*DeleteLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLBStickinessPolicyResponse
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

type DeleteLBStickinessPolicyResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLoadBalancerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteLoadBalancerParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteLoadBalancerParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLoadBalancerParams(id string) *DeleteLoadBalancerParams {
	p := &DeleteLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an internal load balancer
func (s *LoadBalancerService) DeleteLoadBalancer(p *DeleteLoadBalancerParams) (*DeleteLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLoadBalancerResponse
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

type DeleteLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteLoadBalancerRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteLoadBalancerRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLoadBalancerRuleParams(id string) *DeleteLoadBalancerRuleParams {
	p := &DeleteLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer rule.
func (s *LoadBalancerService) DeleteLoadBalancerRule(p *DeleteLoadBalancerRuleParams) (*DeleteLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLoadBalancerRuleResponse
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

type DeleteLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
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
func (s *LoadBalancerService) NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams {
	p := &DeleteNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// delete a netscaler load balancer device
func (s *LoadBalancerService) DeleteNetscalerLoadBalancer(p *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteNetscalerLoadBalancer", p.toURLValues())
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

type DeleteSslCertParams struct {
	p map[string]interface{}
}

func (p *DeleteSslCertParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteSslCertParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteSslCertParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteSslCertParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSslCertParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteSslCertParams(id string) *DeleteSslCertParams {
	p := &DeleteSslCertParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete a certificate to CloudStack
func (s *LoadBalancerService) DeleteSslCert(p *DeleteSslCertParams) (*DeleteSslCertResponse, error) {
	resp, err := s.cs.newRequest("deleteSslCert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSslCertResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSslCertResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSslCertResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSslCertResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListGlobalLoadBalancerRulesParams struct {
	p map[string]interface{}
}

func (p *ListGlobalLoadBalancerRulesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListGlobalLoadBalancerRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetRegionid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["regionid"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetRegionid() {
	if p.p != nil && p.p["regionid"] != nil {
		delete(p.p, "regionid")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetRegionid() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["regionid"].(int)
	return value, ok
}

func (p *ListGlobalLoadBalancerRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListGlobalLoadBalancerRulesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListGlobalLoadBalancerRulesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListGlobalLoadBalancerRulesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListGlobalLoadBalancerRulesParams() *ListGlobalLoadBalancerRulesParams {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListGlobalLoadBalancerRules(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.GlobalLoadBalancerRules[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.GlobalLoadBalancerRules {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleByName(name string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error) {
	id, count, err := s.GetGlobalLoadBalancerRuleID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetGlobalLoadBalancerRuleByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleByID(id string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error) {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListGlobalLoadBalancerRules(p)
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
		return l.GlobalLoadBalancerRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GlobalLoadBalancerRule UUID: %s!", id)
}

// Lists load balancer rules.
func (s *LoadBalancerService) ListGlobalLoadBalancerRules(p *ListGlobalLoadBalancerRulesParams) (*ListGlobalLoadBalancerRulesResponse, error) {
	resp, err := s.cs.newRequest("listGlobalLoadBalancerRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGlobalLoadBalancerRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListGlobalLoadBalancerRulesResponse struct {
	Count                   int                       `json:"count"`
	GlobalLoadBalancerRules []*GlobalLoadBalancerRule `json:"globalloadbalancerrule"`
}

type GlobalLoadBalancerRule struct {
	Account                     string                                   `json:"account"`
	Description                 string                                   `json:"description"`
	Domain                      string                                   `json:"domain"`
	Domainid                    string                                   `json:"domainid"`
	Gslbdomainname              string                                   `json:"gslbdomainname"`
	Gslblbmethod                string                                   `json:"gslblbmethod"`
	Gslbservicetype             string                                   `json:"gslbservicetype"`
	Gslbstickysessionmethodname string                                   `json:"gslbstickysessionmethodname"`
	Id                          string                                   `json:"id"`
	JobID                       string                                   `json:"jobid"`
	Jobstatus                   int                                      `json:"jobstatus"`
	Loadbalancerrule            []GlobalLoadBalancerRuleLoadbalancerrule `json:"loadbalancerrule"`
	Name                        string                                   `json:"name"`
	Project                     string                                   `json:"project"`
	Projectid                   string                                   `json:"projectid"`
	Regionid                    int                                      `json:"regionid"`
}

type GlobalLoadBalancerRuleLoadbalancerrule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type ListLBHealthCheckPoliciesParams struct {
	p map[string]interface{}
}

func (p *ListLBHealthCheckPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
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

func (p *ListLBHealthCheckPoliciesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListLBHealthCheckPoliciesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListLBHealthCheckPoliciesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListLBHealthCheckPoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListLBHealthCheckPoliciesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListLBHealthCheckPoliciesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListLBHealthCheckPoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLBHealthCheckPoliciesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListLBHealthCheckPoliciesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListLBHealthCheckPoliciesParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *ListLBHealthCheckPoliciesParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *ListLBHealthCheckPoliciesParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *ListLBHealthCheckPoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLBHealthCheckPoliciesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListLBHealthCheckPoliciesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListLBHealthCheckPoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLBHealthCheckPoliciesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListLBHealthCheckPoliciesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListLBHealthCheckPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLBHealthCheckPoliciesParams() *ListLBHealthCheckPoliciesParams {
	p := &ListLBHealthCheckPoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLBHealthCheckPolicyByID(id string, opts ...OptionFunc) (*LBHealthCheckPolicy, int, error) {
	p := &ListLBHealthCheckPoliciesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLBHealthCheckPolicies(p)
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
		return l.LBHealthCheckPolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LBHealthCheckPolicy UUID: %s!", id)
}

// Lists load balancer health check policies.
func (s *LoadBalancerService) ListLBHealthCheckPolicies(p *ListLBHealthCheckPoliciesParams) (*ListLBHealthCheckPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listLBHealthCheckPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLBHealthCheckPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLBHealthCheckPoliciesResponse struct {
	Count                 int                    `json:"count"`
	LBHealthCheckPolicies []*LBHealthCheckPolicy `json:"lbhealthcheckpolicy"`
}

type LBHealthCheckPolicy struct {
	Account           string                                 `json:"account"`
	Domain            string                                 `json:"domain"`
	Domainid          string                                 `json:"domainid"`
	Healthcheckpolicy []LBHealthCheckPolicyHealthcheckpolicy `json:"healthcheckpolicy"`
	JobID             string                                 `json:"jobid"`
	Jobstatus         int                                    `json:"jobstatus"`
	Lbruleid          string                                 `json:"lbruleid"`
	Zoneid            string                                 `json:"zoneid"`
}

type LBHealthCheckPolicyHealthcheckpolicy struct {
	Description             string `json:"description"`
	Fordisplay              bool   `json:"fordisplay"`
	Healthcheckinterval     int    `json:"healthcheckinterval"`
	Healthcheckthresshold   int    `json:"healthcheckthresshold"`
	Id                      string `json:"id"`
	Pingpath                string `json:"pingpath"`
	Responsetime            int    `json:"responsetime"`
	State                   string `json:"state"`
	Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
}

type ListLBStickinessPoliciesParams struct {
	p map[string]interface{}
}

func (p *ListLBStickinessPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
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

func (p *ListLBStickinessPoliciesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListLBStickinessPoliciesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListLBStickinessPoliciesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListLBStickinessPoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListLBStickinessPoliciesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListLBStickinessPoliciesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListLBStickinessPoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLBStickinessPoliciesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListLBStickinessPoliciesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListLBStickinessPoliciesParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *ListLBStickinessPoliciesParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *ListLBStickinessPoliciesParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *ListLBStickinessPoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLBStickinessPoliciesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListLBStickinessPoliciesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListLBStickinessPoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLBStickinessPoliciesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListLBStickinessPoliciesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListLBStickinessPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLBStickinessPoliciesParams() *ListLBStickinessPoliciesParams {
	p := &ListLBStickinessPoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLBStickinessPolicyByID(id string, opts ...OptionFunc) (*LBStickinessPolicy, int, error) {
	p := &ListLBStickinessPoliciesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLBStickinessPolicies(p)
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
		return l.LBStickinessPolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LBStickinessPolicy UUID: %s!", id)
}

// Lists load balancer stickiness policies.
func (s *LoadBalancerService) ListLBStickinessPolicies(p *ListLBStickinessPoliciesParams) (*ListLBStickinessPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listLBStickinessPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLBStickinessPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLBStickinessPoliciesResponse struct {
	Count                int                   `json:"count"`
	LBStickinessPolicies []*LBStickinessPolicy `json:"lbstickinesspolicy"`
}

type LBStickinessPolicy struct {
	Account          string                               `json:"account"`
	Description      string                               `json:"description"`
	Domain           string                               `json:"domain"`
	Domainid         string                               `json:"domainid"`
	JobID            string                               `json:"jobid"`
	Jobstatus        int                                  `json:"jobstatus"`
	Lbruleid         string                               `json:"lbruleid"`
	Name             string                               `json:"name"`
	State            string                               `json:"state"`
	Stickinesspolicy []LBStickinessPolicyStickinesspolicy `json:"stickinesspolicy"`
	Zoneid           string                               `json:"zoneid"`
}

type LBStickinessPolicyStickinesspolicy struct {
	Description string            `json:"description"`
	Fordisplay  bool              `json:"fordisplay"`
	Id          string            `json:"id"`
	Methodname  string            `json:"methodname"`
	Name        string            `json:"name"`
	Params      map[string]string `json:"params"`
	State       string            `json:"state"`
}

type ListLoadBalancerRuleInstancesParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancerRuleInstancesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["applied"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("applied", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbvmips"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lbvmips", vv)
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

func (p *ListLoadBalancerRuleInstancesParams) SetApplied(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["applied"] = v
}

func (p *ListLoadBalancerRuleInstancesParams) ResetApplied() {
	if p.p != nil && p.p["applied"] != nil {
		delete(p.p, "applied")
	}
}

func (p *ListLoadBalancerRuleInstancesParams) GetApplied() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["applied"].(bool)
	return value, ok
}

func (p *ListLoadBalancerRuleInstancesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListLoadBalancerRuleInstancesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListLoadBalancerRuleInstancesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListLoadBalancerRuleInstancesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLoadBalancerRuleInstancesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListLoadBalancerRuleInstancesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListLoadBalancerRuleInstancesParams) SetLbvmips(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbvmips"] = v
}

func (p *ListLoadBalancerRuleInstancesParams) ResetLbvmips() {
	if p.p != nil && p.p["lbvmips"] != nil {
		delete(p.p, "lbvmips")
	}
}

func (p *ListLoadBalancerRuleInstancesParams) GetLbvmips() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbvmips"].(bool)
	return value, ok
}

func (p *ListLoadBalancerRuleInstancesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLoadBalancerRuleInstancesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListLoadBalancerRuleInstancesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListLoadBalancerRuleInstancesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLoadBalancerRuleInstancesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListLoadBalancerRuleInstancesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListLoadBalancerRuleInstancesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancerRuleInstancesParams(id string) *ListLoadBalancerRuleInstancesParams {
	p := &ListLoadBalancerRuleInstancesParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleInstanceByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error) {
	p := &ListLoadBalancerRuleInstancesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancerRuleInstances(p)
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
		return l.LoadBalancerRuleInstances[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancerRuleInstance UUID: %s!", id)
}

// List all virtual machine instances that are assigned to a load balancer rule.
func (s *LoadBalancerService) ListLoadBalancerRuleInstances(p *ListLoadBalancerRuleInstancesParams) (*ListLoadBalancerRuleInstancesResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerRuleInstances", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancerRuleInstancesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancerRuleInstancesResponse struct {
	Count                     int                         `json:"count"`
	LBRuleVMIDIPs             []*LoadBalancerRuleInstance `json:"lbrulevmidip"`
	LoadBalancerRuleInstances []*VirtualMachine           `json:"loadbalancerruleinstance"`
}

type LoadBalancerRuleInstance struct {
	JobID                    string          `json:"jobid"`
	Jobstatus                int             `json:"jobstatus"`
	Lbvmipaddresses          []string        `json:"lbvmipaddresses"`
	Loadbalancerruleinstance *VirtualMachine `json:"loadbalancerruleinstance"`
}

type ListLoadBalancerRulesParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancerRulesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListLoadBalancerRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListLoadBalancerRulesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListLoadBalancerRulesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListLoadBalancerRulesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListLoadBalancerRulesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListLoadBalancerRulesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListLoadBalancerRulesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListLoadBalancerRulesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListLoadBalancerRulesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListLoadBalancerRulesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListLoadBalancerRulesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLoadBalancerRulesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListLoadBalancerRulesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListLoadBalancerRulesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListLoadBalancerRulesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListLoadBalancerRulesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListLoadBalancerRulesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListLoadBalancerRulesParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListLoadBalancerRulesParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLoadBalancerRulesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListLoadBalancerRulesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLoadBalancerRulesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListLoadBalancerRulesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListLoadBalancerRulesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListLoadBalancerRulesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
}

func (p *ListLoadBalancerRulesParams) ResetPublicipid() {
	if p.p != nil && p.p["publicipid"] != nil {
		delete(p.p, "publicipid")
	}
}

func (p *ListLoadBalancerRulesParams) GetPublicipid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicipid"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListLoadBalancerRulesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListLoadBalancerRulesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListLoadBalancerRulesParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *ListLoadBalancerRulesParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListLoadBalancerRulesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListLoadBalancerRulesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListLoadBalancerRulesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListLoadBalancerRulesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancerRulesParams() *ListLoadBalancerRulesParams {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListLoadBalancerRules(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.LoadBalancerRules[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.LoadBalancerRules {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleByName(name string, opts ...OptionFunc) (*LoadBalancerRule, int, error) {
	id, count, err := s.GetLoadBalancerRuleID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetLoadBalancerRuleByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleByID(id string, opts ...OptionFunc) (*LoadBalancerRule, int, error) {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancerRules(p)
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
		return l.LoadBalancerRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancerRule UUID: %s!", id)
}

// Lists load balancer rules.
func (s *LoadBalancerService) ListLoadBalancerRules(p *ListLoadBalancerRulesParams) (*ListLoadBalancerRulesResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancerRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancerRulesResponse struct {
	Count             int                 `json:"count"`
	LoadBalancerRules []*LoadBalancerRule `json:"loadbalancerrule"`
}

type LoadBalancerRule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type ListLoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["scheme"]; found {
		u.Set("scheme", v.(string))
	}
	if v, found := p.p["sourceipaddress"]; found {
		u.Set("sourceipaddress", v.(string))
	}
	if v, found := p.p["sourceipaddressnetworkid"]; found {
		u.Set("sourceipaddressnetworkid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListLoadBalancersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListLoadBalancersParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListLoadBalancersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListLoadBalancersParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListLoadBalancersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListLoadBalancersParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListLoadBalancersParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListLoadBalancersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListLoadBalancersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListLoadBalancersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListLoadBalancersParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListLoadBalancersParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListLoadBalancersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLoadBalancersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListLoadBalancersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListLoadBalancersParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListLoadBalancersParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListLoadBalancersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListLoadBalancersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListLoadBalancersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListLoadBalancersParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListLoadBalancersParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLoadBalancersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListLoadBalancersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListLoadBalancersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLoadBalancersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListLoadBalancersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListLoadBalancersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListLoadBalancersParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListLoadBalancersParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetScheme(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scheme"] = v
}

func (p *ListLoadBalancersParams) ResetScheme() {
	if p.p != nil && p.p["scheme"] != nil {
		delete(p.p, "scheme")
	}
}

func (p *ListLoadBalancersParams) GetScheme() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scheme"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetSourceipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddress"] = v
}

func (p *ListLoadBalancersParams) ResetSourceipaddress() {
	if p.p != nil && p.p["sourceipaddress"] != nil {
		delete(p.p, "sourceipaddress")
	}
}

func (p *ListLoadBalancersParams) GetSourceipaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourceipaddress"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetSourceipaddressnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddressnetworkid"] = v
}

func (p *ListLoadBalancersParams) ResetSourceipaddressnetworkid() {
	if p.p != nil && p.p["sourceipaddressnetworkid"] != nil {
		delete(p.p, "sourceipaddressnetworkid")
	}
}

func (p *ListLoadBalancersParams) GetSourceipaddressnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourceipaddressnetworkid"].(string)
	return value, ok
}

func (p *ListLoadBalancersParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListLoadBalancersParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListLoadBalancersParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListLoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancersParams() *ListLoadBalancersParams {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListLoadBalancers(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.LoadBalancers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.LoadBalancers {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerByName(name string, opts ...OptionFunc) (*LoadBalancer, int, error) {
	id, count, err := s.GetLoadBalancerID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetLoadBalancerByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerByID(id string, opts ...OptionFunc) (*LoadBalancer, int, error) {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancers(p)
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
		return l.LoadBalancers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancer UUID: %s!", id)
}

// Lists internal load balancers
func (s *LoadBalancerService) ListLoadBalancers(p *ListLoadBalancersParams) (*ListLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancersResponse struct {
	Count         int             `json:"count"`
	LoadBalancers []*LoadBalancer `json:"loadbalancer"`
}

type LoadBalancer struct {
	Account                  string                             `json:"account"`
	Algorithm                string                             `json:"algorithm"`
	Description              string                             `json:"description"`
	Domain                   string                             `json:"domain"`
	Domainid                 string                             `json:"domainid"`
	Fordisplay               bool                               `json:"fordisplay"`
	Id                       string                             `json:"id"`
	JobID                    string                             `json:"jobid"`
	Jobstatus                int                                `json:"jobstatus"`
	Loadbalancerinstance     []LoadBalancerLoadbalancerinstance `json:"loadbalancerinstance"`
	Loadbalancerrule         []LoadBalancerLoadbalancerrule     `json:"loadbalancerrule"`
	Name                     string                             `json:"name"`
	Networkid                string                             `json:"networkid"`
	Project                  string                             `json:"project"`
	Projectid                string                             `json:"projectid"`
	Sourceipaddress          string                             `json:"sourceipaddress"`
	Sourceipaddressnetworkid string                             `json:"sourceipaddressnetworkid"`
	Tags                     []Tags                             `json:"tags"`
}

type LoadBalancerLoadbalancerrule struct {
	Instanceport int    `json:"instanceport"`
	Sourceport   int    `json:"sourceport"`
	State        string `json:"state"`
}

type LoadBalancerLoadbalancerinstance struct {
	Id        string `json:"id"`
	Ipaddress string `json:"ipaddress"`
	Name      string `json:"name"`
	State     string `json:"state"`
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
func (s *LoadBalancerService) NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams {
	p := &ListNetscalerLoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists netscaler load balancer devices
func (s *LoadBalancerService) ListNetscalerLoadBalancers(p *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error) {
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

type ListSslCertsParams struct {
	p map[string]interface{}
}

func (p *ListSslCertsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["certid"]; found {
		u.Set("certid", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *ListSslCertsParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *ListSslCertsParams) ResetAccountid() {
	if p.p != nil && p.p["accountid"] != nil {
		delete(p.p, "accountid")
	}
}

func (p *ListSslCertsParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *ListSslCertsParams) SetCertid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certid"] = v
}

func (p *ListSslCertsParams) ResetCertid() {
	if p.p != nil && p.p["certid"] != nil {
		delete(p.p, "certid")
	}
}

func (p *ListSslCertsParams) GetCertid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["certid"].(string)
	return value, ok
}

func (p *ListSslCertsParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *ListSslCertsParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *ListSslCertsParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *ListSslCertsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListSslCertsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListSslCertsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSslCertsParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListSslCertsParams() *ListSslCertsParams {
	p := &ListSslCertsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists SSL certificates
func (s *LoadBalancerService) ListSslCerts(p *ListSslCertsParams) (*ListSslCertsResponse, error) {
	resp, err := s.cs.newRequest("listSslCerts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSslCertsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSslCertsResponse struct {
	Count    int        `json:"count"`
	SslCerts []*SslCert `json:"sslcert"`
}

type SslCert struct {
	Account              string   `json:"account"`
	Certchain            string   `json:"certchain"`
	Certificate          string   `json:"certificate"`
	Domain               string   `json:"domain"`
	Domainid             string   `json:"domainid"`
	Fingerprint          string   `json:"fingerprint"`
	Id                   string   `json:"id"`
	JobID                string   `json:"jobid"`
	Jobstatus            int      `json:"jobstatus"`
	Loadbalancerrulelist []string `json:"loadbalancerrulelist"`
	Name                 string   `json:"name"`
	Project              string   `json:"project"`
	Projectid            string   `json:"projectid"`
}

type RemoveCertFromLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *RemoveCertFromLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (p *RemoveCertFromLoadBalancerParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *RemoveCertFromLoadBalancerParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *RemoveCertFromLoadBalancerParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveCertFromLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveCertFromLoadBalancerParams(lbruleid string) *RemoveCertFromLoadBalancerParams {
	p := &RemoveCertFromLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// Removes a certificate from a load balancer rule
func (s *LoadBalancerService) RemoveCertFromLoadBalancer(p *RemoveCertFromLoadBalancerParams) (*RemoveCertFromLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("removeCertFromLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveCertFromLoadBalancerResponse
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

type RemoveCertFromLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RemoveFromGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["loadbalancerrulelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("loadbalancerrulelist", vv)
	}
	return u
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) SetLoadbalancerrulelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerrulelist"] = v
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) ResetLoadbalancerrulelist() {
	if p.p != nil && p.p["loadbalancerrulelist"] != nil {
		delete(p.p, "loadbalancerrulelist")
	}
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) GetLoadbalancerrulelist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["loadbalancerrulelist"].([]string)
	return value, ok
}

// You should always use this function to get a new RemoveFromGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveFromGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *RemoveFromGlobalLoadBalancerRuleParams {
	p := &RemoveFromGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["loadbalancerrulelist"] = loadbalancerrulelist
	return p
}

// Removes a load balancer rule association with global load balancer rule
func (s *LoadBalancerService) RemoveFromGlobalLoadBalancerRule(p *RemoveFromGlobalLoadBalancerRuleParams) (*RemoveFromGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("removeFromGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveFromGlobalLoadBalancerRuleResponse
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

type RemoveFromGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RemoveFromLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *RemoveFromLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("virtualmachineids", vv)
	}
	if v, found := p.p["vmidipmap"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("vmidipmap[%d].key", i), k)
			u.Set(fmt.Sprintf("vmidipmap[%d].value", i), m[k])
		}
	}
	return u
}

func (p *RemoveFromLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RemoveFromLoadBalancerRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RemoveFromLoadBalancerRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *RemoveFromLoadBalancerRuleParams) SetVirtualmachineids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineids"] = v
}

func (p *RemoveFromLoadBalancerRuleParams) ResetVirtualmachineids() {
	if p.p != nil && p.p["virtualmachineids"] != nil {
		delete(p.p, "virtualmachineids")
	}
}

func (p *RemoveFromLoadBalancerRuleParams) GetVirtualmachineids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineids"].([]string)
	return value, ok
}

func (p *RemoveFromLoadBalancerRuleParams) SetVmidipmap(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmidipmap"] = v
}

func (p *RemoveFromLoadBalancerRuleParams) ResetVmidipmap() {
	if p.p != nil && p.p["vmidipmap"] != nil {
		delete(p.p, "vmidipmap")
	}
}

func (p *RemoveFromLoadBalancerRuleParams) GetVmidipmap() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmidipmap"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new RemoveFromLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveFromLoadBalancerRuleParams(id string) *RemoveFromLoadBalancerRuleParams {
	p := &RemoveFromLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes a virtual machine or a list of virtual machines from a load balancer rule.
func (s *LoadBalancerService) RemoveFromLoadBalancerRule(p *RemoveFromLoadBalancerRuleParams) (*RemoveFromLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("removeFromLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveFromLoadBalancerRuleResponse
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

type RemoveFromLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["gslblbmethod"]; found {
		u.Set("gslblbmethod", v.(string))
	}
	if v, found := p.p["gslbstickysessionmethodname"]; found {
		u.Set("gslbstickysessionmethodname", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateGlobalLoadBalancerRuleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateGlobalLoadBalancerRuleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetGslblbmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbmethod"] = v
}

func (p *UpdateGlobalLoadBalancerRuleParams) ResetGslblbmethod() {
	if p.p != nil && p.p["gslblbmethod"] != nil {
		delete(p.p, "gslblbmethod")
	}
}

func (p *UpdateGlobalLoadBalancerRuleParams) GetGslblbmethod() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslblbmethod"].(string)
	return value, ok
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetGslbstickysessionmethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbstickysessionmethodname"] = v
}

func (p *UpdateGlobalLoadBalancerRuleParams) ResetGslbstickysessionmethodname() {
	if p.p != nil && p.p["gslbstickysessionmethodname"] != nil {
		delete(p.p, "gslbstickysessionmethodname")
	}
}

func (p *UpdateGlobalLoadBalancerRuleParams) GetGslbstickysessionmethodname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gslbstickysessionmethodname"].(string)
	return value, ok
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateGlobalLoadBalancerRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateGlobalLoadBalancerRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateGlobalLoadBalancerRuleParams(id string) *UpdateGlobalLoadBalancerRuleParams {
	p := &UpdateGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// update global load balancer rules.
func (s *LoadBalancerService) UpdateGlobalLoadBalancerRule(p *UpdateGlobalLoadBalancerRuleParams) (*UpdateGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("updateGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGlobalLoadBalancerRuleResponse
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

type UpdateGlobalLoadBalancerRuleResponse struct {
	Account                     string                                                 `json:"account"`
	Description                 string                                                 `json:"description"`
	Domain                      string                                                 `json:"domain"`
	Domainid                    string                                                 `json:"domainid"`
	Gslbdomainname              string                                                 `json:"gslbdomainname"`
	Gslblbmethod                string                                                 `json:"gslblbmethod"`
	Gslbservicetype             string                                                 `json:"gslbservicetype"`
	Gslbstickysessionmethodname string                                                 `json:"gslbstickysessionmethodname"`
	Id                          string                                                 `json:"id"`
	JobID                       string                                                 `json:"jobid"`
	Jobstatus                   int                                                    `json:"jobstatus"`
	Loadbalancerrule            []UpdateGlobalLoadBalancerRuleResponseLoadbalancerrule `json:"loadbalancerrule"`
	Name                        string                                                 `json:"name"`
	Project                     string                                                 `json:"project"`
	Projectid                   string                                                 `json:"projectid"`
	Regionid                    int                                                    `json:"regionid"`
}

type UpdateGlobalLoadBalancerRuleResponseLoadbalancerrule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type UpdateLBHealthCheckPolicyParams struct {
	p map[string]interface{}
}

func (p *UpdateLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateLBHealthCheckPolicyParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateLBHealthCheckPolicyParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateLBHealthCheckPolicyParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateLBHealthCheckPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateLBHealthCheckPolicyParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateLBHealthCheckPolicyParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateLBHealthCheckPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateLBHealthCheckPolicyParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateLBHealthCheckPolicyParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLBHealthCheckPolicyParams(id string) *UpdateLBHealthCheckPolicyParams {
	p := &UpdateLBHealthCheckPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates load balancer health check policy
func (s *LoadBalancerService) UpdateLBHealthCheckPolicy(p *UpdateLBHealthCheckPolicyParams) (*UpdateLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("updateLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLBHealthCheckPolicyResponse
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

type UpdateLBHealthCheckPolicyResponse struct {
	Account           string                                               `json:"account"`
	Domain            string                                               `json:"domain"`
	Domainid          string                                               `json:"domainid"`
	Healthcheckpolicy []UpdateLBHealthCheckPolicyResponseHealthcheckpolicy `json:"healthcheckpolicy"`
	JobID             string                                               `json:"jobid"`
	Jobstatus         int                                                  `json:"jobstatus"`
	Lbruleid          string                                               `json:"lbruleid"`
	Zoneid            string                                               `json:"zoneid"`
}

type UpdateLBHealthCheckPolicyResponseHealthcheckpolicy struct {
	Description             string `json:"description"`
	Fordisplay              bool   `json:"fordisplay"`
	Healthcheckinterval     int    `json:"healthcheckinterval"`
	Healthcheckthresshold   int    `json:"healthcheckthresshold"`
	Id                      string `json:"id"`
	Pingpath                string `json:"pingpath"`
	Responsetime            int    `json:"responsetime"`
	State                   string `json:"state"`
	Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
}

type UpdateLBStickinessPolicyParams struct {
	p map[string]interface{}
}

func (p *UpdateLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateLBStickinessPolicyParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateLBStickinessPolicyParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateLBStickinessPolicyParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateLBStickinessPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateLBStickinessPolicyParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateLBStickinessPolicyParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateLBStickinessPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateLBStickinessPolicyParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateLBStickinessPolicyParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLBStickinessPolicyParams(id string) *UpdateLBStickinessPolicyParams {
	p := &UpdateLBStickinessPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates load balancer stickiness policy
func (s *LoadBalancerService) UpdateLBStickinessPolicy(p *UpdateLBStickinessPolicyParams) (*UpdateLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("updateLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLBStickinessPolicyResponse
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

type UpdateLBStickinessPolicyResponse struct {
	Account          string                                             `json:"account"`
	Description      string                                             `json:"description"`
	Domain           string                                             `json:"domain"`
	Domainid         string                                             `json:"domainid"`
	JobID            string                                             `json:"jobid"`
	Jobstatus        int                                                `json:"jobstatus"`
	Lbruleid         string                                             `json:"lbruleid"`
	Name             string                                             `json:"name"`
	State            string                                             `json:"state"`
	Stickinesspolicy []UpdateLBStickinessPolicyResponseStickinesspolicy `json:"stickinesspolicy"`
	Zoneid           string                                             `json:"zoneid"`
}

type UpdateLBStickinessPolicyResponseStickinesspolicy struct {
	Description string            `json:"description"`
	Fordisplay  bool              `json:"fordisplay"`
	Id          string            `json:"id"`
	Methodname  string            `json:"methodname"`
	Name        string            `json:"name"`
	Params      map[string]string `json:"params"`
	State       string            `json:"state"`
}

type UpdateLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *UpdateLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateLoadBalancerParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateLoadBalancerParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateLoadBalancerParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateLoadBalancerParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateLoadBalancerParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateLoadBalancerParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateLoadBalancerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateLoadBalancerParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateLoadBalancerParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLoadBalancerParams(id string) *UpdateLoadBalancerParams {
	p := &UpdateLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an internal load balancer
func (s *LoadBalancerService) UpdateLoadBalancer(p *UpdateLoadBalancerParams) (*UpdateLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("updateLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLoadBalancerResponse
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

type UpdateLoadBalancerResponse struct {
	Account                  string                                           `json:"account"`
	Algorithm                string                                           `json:"algorithm"`
	Description              string                                           `json:"description"`
	Domain                   string                                           `json:"domain"`
	Domainid                 string                                           `json:"domainid"`
	Fordisplay               bool                                             `json:"fordisplay"`
	Id                       string                                           `json:"id"`
	JobID                    string                                           `json:"jobid"`
	Jobstatus                int                                              `json:"jobstatus"`
	Loadbalancerinstance     []UpdateLoadBalancerResponseLoadbalancerinstance `json:"loadbalancerinstance"`
	Loadbalancerrule         []UpdateLoadBalancerResponseLoadbalancerrule     `json:"loadbalancerrule"`
	Name                     string                                           `json:"name"`
	Networkid                string                                           `json:"networkid"`
	Project                  string                                           `json:"project"`
	Projectid                string                                           `json:"projectid"`
	Sourceipaddress          string                                           `json:"sourceipaddress"`
	Sourceipaddressnetworkid string                                           `json:"sourceipaddressnetworkid"`
	Tags                     []Tags                                           `json:"tags"`
}

type UpdateLoadBalancerResponseLoadbalancerrule struct {
	Instanceport int    `json:"instanceport"`
	Sourceport   int    `json:"sourceport"`
	State        string `json:"state"`
}

type UpdateLoadBalancerResponseLoadbalancerinstance struct {
	Id        string `json:"id"`
	Ipaddress string `json:"ipaddress"`
	Name      string `json:"name"`
	State     string `json:"state"`
}

type UpdateLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	return u
}

func (p *UpdateLoadBalancerRuleParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
}

func (p *UpdateLoadBalancerRuleParams) ResetAlgorithm() {
	if p.p != nil && p.p["algorithm"] != nil {
		delete(p.p, "algorithm")
	}
}

func (p *UpdateLoadBalancerRuleParams) GetAlgorithm() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["algorithm"].(string)
	return value, ok
}

func (p *UpdateLoadBalancerRuleParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateLoadBalancerRuleParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateLoadBalancerRuleParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateLoadBalancerRuleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateLoadBalancerRuleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateLoadBalancerRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateLoadBalancerRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateLoadBalancerRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateLoadBalancerRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateLoadBalancerRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateLoadBalancerRuleParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateLoadBalancerRuleParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateLoadBalancerRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *UpdateLoadBalancerRuleParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *UpdateLoadBalancerRuleParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLoadBalancerRuleParams(id string) *UpdateLoadBalancerRuleParams {
	p := &UpdateLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates load balancer
func (s *LoadBalancerService) UpdateLoadBalancerRule(p *UpdateLoadBalancerRuleParams) (*UpdateLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("updateLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLoadBalancerRuleResponse
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

type UpdateLoadBalancerRuleResponse struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type UploadSslCertParams struct {
	p map[string]interface{}
}

func (p *UploadSslCertParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["certchain"]; found {
		u.Set("certchain", v.(string))
	}
	if v, found := p.p["certificate"]; found {
		u.Set("certificate", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["enabledrevocationcheck"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabledrevocationcheck", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["privatekey"]; found {
		u.Set("privatekey", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *UploadSslCertParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UploadSslCertParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *UploadSslCertParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UploadSslCertParams) SetCertchain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certchain"] = v
}

func (p *UploadSslCertParams) ResetCertchain() {
	if p.p != nil && p.p["certchain"] != nil {
		delete(p.p, "certchain")
	}
}

func (p *UploadSslCertParams) GetCertchain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["certchain"].(string)
	return value, ok
}

func (p *UploadSslCertParams) SetCertificate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certificate"] = v
}

func (p *UploadSslCertParams) ResetCertificate() {
	if p.p != nil && p.p["certificate"] != nil {
		delete(p.p, "certificate")
	}
}

func (p *UploadSslCertParams) GetCertificate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["certificate"].(string)
	return value, ok
}

func (p *UploadSslCertParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UploadSslCertParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *UploadSslCertParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UploadSslCertParams) SetEnabledrevocationcheck(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabledrevocationcheck"] = v
}

func (p *UploadSslCertParams) ResetEnabledrevocationcheck() {
	if p.p != nil && p.p["enabledrevocationcheck"] != nil {
		delete(p.p, "enabledrevocationcheck")
	}
}

func (p *UploadSslCertParams) GetEnabledrevocationcheck() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabledrevocationcheck"].(bool)
	return value, ok
}

func (p *UploadSslCertParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UploadSslCertParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UploadSslCertParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UploadSslCertParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *UploadSslCertParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *UploadSslCertParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *UploadSslCertParams) SetPrivatekey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privatekey"] = v
}

func (p *UploadSslCertParams) ResetPrivatekey() {
	if p.p != nil && p.p["privatekey"] != nil {
		delete(p.p, "privatekey")
	}
}

func (p *UploadSslCertParams) GetPrivatekey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privatekey"].(string)
	return value, ok
}

func (p *UploadSslCertParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *UploadSslCertParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *UploadSslCertParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new UploadSslCertParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUploadSslCertParams(certificate string, name string, privatekey string) *UploadSslCertParams {
	p := &UploadSslCertParams{}
	p.p = make(map[string]interface{})
	p.p["certificate"] = certificate
	p.p["name"] = name
	p.p["privatekey"] = privatekey
	return p
}

// Upload a certificate to CloudStack
func (s *LoadBalancerService) UploadSslCert(p *UploadSslCertParams) (*UploadSslCertResponse, error) {
	resp, err := s.cs.newRequest("uploadSslCert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadSslCertResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UploadSslCertResponse struct {
	Account              string   `json:"account"`
	Certchain            string   `json:"certchain"`
	Certificate          string   `json:"certificate"`
	Domain               string   `json:"domain"`
	Domainid             string   `json:"domainid"`
	Fingerprint          string   `json:"fingerprint"`
	Id                   string   `json:"id"`
	JobID                string   `json:"jobid"`
	Jobstatus            int      `json:"jobstatus"`
	Loadbalancerrulelist []string `json:"loadbalancerrulelist"`
	Name                 string   `json:"name"`
	Project              string   `json:"project"`
	Projectid            string   `json:"projectid"`
}
