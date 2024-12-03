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

// Helper function for maintaining backwards compatibility
func convertFirewallServiceResponse(b []byte) ([]byte, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if _, ok := raw["firewallrule"]; ok {
		return convertFirewallServiceListResponse(b)
	}

	for _, k := range []string{"endport", "startport"} {
		if sVal, ok := raw[k].(string); ok {
			iVal, err := strconv.Atoi(sVal)
			if err != nil {
				return nil, err
			}
			raw[k] = iVal
		}
	}

	return json.Marshal(raw)
}

// Helper function for maintaining backwards compatibility
func convertFirewallServiceListResponse(b []byte) ([]byte, error) {
	var rawList struct {
		Count         int                      `json:"count"`
		FirewallRules []map[string]interface{} `json:"firewallrule"`
	}

	if err := json.Unmarshal(b, &rawList); err != nil {
		return nil, err
	}

	for _, r := range rawList.FirewallRules {
		for _, k := range []string{"endport", "startport"} {
			if sVal, ok := r[k].(string); ok {
				iVal, err := strconv.Atoi(sVal)
				if err != nil {
					return nil, err
				}
				r[k] = iVal
			}
		}
	}

	return json.Marshal(rawList)
}

type FirewallServiceIface interface {
	AddPaloAltoFirewall(p *AddPaloAltoFirewallParams) (*AddPaloAltoFirewallResponse, error)
	NewAddPaloAltoFirewallParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddPaloAltoFirewallParams
	ConfigurePaloAltoFirewall(p *ConfigurePaloAltoFirewallParams) (*PaloAltoFirewallResponse, error)
	NewConfigurePaloAltoFirewallParams(fwdeviceid string) *ConfigurePaloAltoFirewallParams
	CreateEgressFirewallRule(p *CreateEgressFirewallRuleParams) (*CreateEgressFirewallRuleResponse, error)
	NewCreateEgressFirewallRuleParams(networkid string, protocol string) *CreateEgressFirewallRuleParams
	CreateFirewallRule(p *CreateFirewallRuleParams) (*CreateFirewallRuleResponse, error)
	NewCreateFirewallRuleParams(ipaddressid string, protocol string) *CreateFirewallRuleParams
	CreatePortForwardingRule(p *CreatePortForwardingRuleParams) (*CreatePortForwardingRuleResponse, error)
	NewCreatePortForwardingRuleParams(ipaddressid string, privateport int, protocol string, publicport int, virtualmachineid string) *CreatePortForwardingRuleParams
	DeleteEgressFirewallRule(p *DeleteEgressFirewallRuleParams) (*DeleteEgressFirewallRuleResponse, error)
	NewDeleteEgressFirewallRuleParams(id string) *DeleteEgressFirewallRuleParams
	DeleteFirewallRule(p *DeleteFirewallRuleParams) (*DeleteFirewallRuleResponse, error)
	NewDeleteFirewallRuleParams(id string) *DeleteFirewallRuleParams
	DeletePaloAltoFirewall(p *DeletePaloAltoFirewallParams) (*DeletePaloAltoFirewallResponse, error)
	NewDeletePaloAltoFirewallParams(fwdeviceid string) *DeletePaloAltoFirewallParams
	DeletePortForwardingRule(p *DeletePortForwardingRuleParams) (*DeletePortForwardingRuleResponse, error)
	NewDeletePortForwardingRuleParams(id string) *DeletePortForwardingRuleParams
	ListEgressFirewallRules(p *ListEgressFirewallRulesParams) (*ListEgressFirewallRulesResponse, error)
	NewListEgressFirewallRulesParams() *ListEgressFirewallRulesParams
	GetEgressFirewallRuleByID(id string, opts ...OptionFunc) (*EgressFirewallRule, int, error)
	ListFirewallRules(p *ListFirewallRulesParams) (*ListFirewallRulesResponse, error)
	NewListFirewallRulesParams() *ListFirewallRulesParams
	GetFirewallRuleByID(id string, opts ...OptionFunc) (*FirewallRule, int, error)
	ListPaloAltoFirewalls(p *ListPaloAltoFirewallsParams) (*ListPaloAltoFirewallsResponse, error)
	NewListPaloAltoFirewallsParams() *ListPaloAltoFirewallsParams
	ListPortForwardingRules(p *ListPortForwardingRulesParams) (*ListPortForwardingRulesResponse, error)
	NewListPortForwardingRulesParams() *ListPortForwardingRulesParams
	GetPortForwardingRuleByID(id string, opts ...OptionFunc) (*PortForwardingRule, int, error)
	UpdateEgressFirewallRule(p *UpdateEgressFirewallRuleParams) (*UpdateEgressFirewallRuleResponse, error)
	NewUpdateEgressFirewallRuleParams(id string) *UpdateEgressFirewallRuleParams
	UpdateFirewallRule(p *UpdateFirewallRuleParams) (*UpdateFirewallRuleResponse, error)
	NewUpdateFirewallRuleParams(id string) *UpdateFirewallRuleParams
	UpdatePortForwardingRule(p *UpdatePortForwardingRuleParams) (*UpdatePortForwardingRuleResponse, error)
	NewUpdatePortForwardingRuleParams(id string) *UpdatePortForwardingRuleParams
	ListIpv6FirewallRules(p *ListIpv6FirewallRulesParams) (*ListIpv6FirewallRulesResponse, error)
	NewListIpv6FirewallRulesParams() *ListIpv6FirewallRulesParams
	GetIpv6FirewallRuleByID(id string, opts ...OptionFunc) (*Ipv6FirewallRule, int, error)
	CreateIpv6FirewallRule(p *CreateIpv6FirewallRuleParams) (*CreateIpv6FirewallRuleResponse, error)
	NewCreateIpv6FirewallRuleParams(networkid string, protocol string) *CreateIpv6FirewallRuleParams
	UpdateIpv6FirewallRule(p *UpdateIpv6FirewallRuleParams) (*UpdateIpv6FirewallRuleResponse, error)
	NewUpdateIpv6FirewallRuleParams(id string) *UpdateIpv6FirewallRuleParams
	DeleteIpv6FirewallRule(p *DeleteIpv6FirewallRuleParams) (*DeleteIpv6FirewallRuleResponse, error)
	NewDeleteIpv6FirewallRuleParams(id string) *DeleteIpv6FirewallRuleParams
}

type AddPaloAltoFirewallParams struct {
	p map[string]interface{}
}

func (p *AddPaloAltoFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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

func (p *AddPaloAltoFirewallParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
}

func (p *AddPaloAltoFirewallParams) ResetNetworkdevicetype() {
	if p.p != nil && p.p["networkdevicetype"] != nil {
		delete(p.p, "networkdevicetype")
	}
}

func (p *AddPaloAltoFirewallParams) GetNetworkdevicetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdevicetype"].(string)
	return value, ok
}

func (p *AddPaloAltoFirewallParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddPaloAltoFirewallParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddPaloAltoFirewallParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddPaloAltoFirewallParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddPaloAltoFirewallParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *AddPaloAltoFirewallParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddPaloAltoFirewallParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddPaloAltoFirewallParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddPaloAltoFirewallParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddPaloAltoFirewallParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddPaloAltoFirewallParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddPaloAltoFirewallParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddPaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewAddPaloAltoFirewallParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddPaloAltoFirewallParams {
	p := &AddPaloAltoFirewallParams{}
	p.p = make(map[string]interface{})
	p.p["networkdevicetype"] = networkdevicetype
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds a Palo Alto firewall device
func (s *FirewallService) AddPaloAltoFirewall(p *AddPaloAltoFirewallParams) (*AddPaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("addPaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddPaloAltoFirewallResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AddPaloAltoFirewallResponse struct {
	Fwdevicecapacity  int64  `json:"fwdevicecapacity"`
	Fwdeviceid        string `json:"fwdeviceid"`
	Fwdevicename      string `json:"fwdevicename"`
	Fwdevicestate     string `json:"fwdevicestate"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Numretries        string `json:"numretries"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Privatezone       string `json:"privatezone"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
	Publiczone        string `json:"publiczone"`
	Timeout           string `json:"timeout"`
	Usageinterface    string `json:"usageinterface"`
	Username          string `json:"username"`
	Zoneid            string `json:"zoneid"`
}

type ConfigurePaloAltoFirewallParams struct {
	p map[string]interface{}
}

func (p *ConfigurePaloAltoFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fwdevicecapacity"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("fwdevicecapacity", vv)
	}
	if v, found := p.p["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
	}
	return u
}

func (p *ConfigurePaloAltoFirewallParams) SetFwdevicecapacity(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fwdevicecapacity"] = v
}

func (p *ConfigurePaloAltoFirewallParams) ResetFwdevicecapacity() {
	if p.p != nil && p.p["fwdevicecapacity"] != nil {
		delete(p.p, "fwdevicecapacity")
	}
}

func (p *ConfigurePaloAltoFirewallParams) GetFwdevicecapacity() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fwdevicecapacity"].(int64)
	return value, ok
}

func (p *ConfigurePaloAltoFirewallParams) SetFwdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fwdeviceid"] = v
}

func (p *ConfigurePaloAltoFirewallParams) ResetFwdeviceid() {
	if p.p != nil && p.p["fwdeviceid"] != nil {
		delete(p.p, "fwdeviceid")
	}
}

func (p *ConfigurePaloAltoFirewallParams) GetFwdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fwdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigurePaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewConfigurePaloAltoFirewallParams(fwdeviceid string) *ConfigurePaloAltoFirewallParams {
	p := &ConfigurePaloAltoFirewallParams{}
	p.p = make(map[string]interface{})
	p.p["fwdeviceid"] = fwdeviceid
	return p
}

// Configures a Palo Alto firewall device
func (s *FirewallService) ConfigurePaloAltoFirewall(p *ConfigurePaloAltoFirewallParams) (*PaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("configurePaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PaloAltoFirewallResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type PaloAltoFirewallResponse struct {
	Fwdevicecapacity  int64  `json:"fwdevicecapacity"`
	Fwdeviceid        string `json:"fwdeviceid"`
	Fwdevicename      string `json:"fwdevicename"`
	Fwdevicestate     string `json:"fwdevicestate"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Numretries        string `json:"numretries"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Privatezone       string `json:"privatezone"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
	Publiczone        string `json:"publiczone"`
	Timeout           string `json:"timeout"`
	Usageinterface    string `json:"usageinterface"`
	Username          string `json:"username"`
	Zoneid            string `json:"zoneid"`
}

type CreateEgressFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *CreateEgressFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["destcidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("destcidrlist", vv)
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *CreateEgressFirewallRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *CreateEgressFirewallRuleParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetDestcidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destcidrlist"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetDestcidrlist() {
	if p.p != nil && p.p["destcidrlist"] != nil {
		delete(p.p, "destcidrlist")
	}
}

func (p *CreateEgressFirewallRuleParams) GetDestcidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destcidrlist"].([]string)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *CreateEgressFirewallRuleParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateEgressFirewallRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetIcmpcode() {
	if p.p != nil && p.p["icmpcode"] != nil {
		delete(p.p, "icmpcode")
	}
}

func (p *CreateEgressFirewallRuleParams) GetIcmpcode() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmpcode"].(int)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetIcmptype() {
	if p.p != nil && p.p["icmptype"] != nil {
		delete(p.p, "icmptype")
	}
}

func (p *CreateEgressFirewallRuleParams) GetIcmptype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmptype"].(int)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreateEgressFirewallRuleParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *CreateEgressFirewallRuleParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *CreateEgressFirewallRuleParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *CreateEgressFirewallRuleParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *CreateEgressFirewallRuleParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *CreateEgressFirewallRuleParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new CreateEgressFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreateEgressFirewallRuleParams(networkid string, protocol string) *CreateEgressFirewallRuleParams {
	p := &CreateEgressFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	p.p["protocol"] = protocol
	return p
}

// Creates a egress firewall rule for a given network
func (s *FirewallService) CreateEgressFirewallRule(p *CreateEgressFirewallRuleParams) (*CreateEgressFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("createEgressFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateEgressFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateEgressFirewallRuleResponse struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
	Traffictype  string `json:"traffictype"`
}

type CreateFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *CreateFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *CreateFirewallRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *CreateFirewallRuleParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *CreateFirewallRuleParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *CreateFirewallRuleParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *CreateFirewallRuleParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *CreateFirewallRuleParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *CreateFirewallRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateFirewallRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateFirewallRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateFirewallRuleParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *CreateFirewallRuleParams) ResetIcmpcode() {
	if p.p != nil && p.p["icmpcode"] != nil {
		delete(p.p, "icmpcode")
	}
}

func (p *CreateFirewallRuleParams) GetIcmpcode() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmpcode"].(int)
	return value, ok
}

func (p *CreateFirewallRuleParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *CreateFirewallRuleParams) ResetIcmptype() {
	if p.p != nil && p.p["icmptype"] != nil {
		delete(p.p, "icmptype")
	}
}

func (p *CreateFirewallRuleParams) GetIcmptype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmptype"].(int)
	return value, ok
}

func (p *CreateFirewallRuleParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *CreateFirewallRuleParams) ResetIpaddressid() {
	if p.p != nil && p.p["ipaddressid"] != nil {
		delete(p.p, "ipaddressid")
	}
}

func (p *CreateFirewallRuleParams) GetIpaddressid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddressid"].(string)
	return value, ok
}

func (p *CreateFirewallRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *CreateFirewallRuleParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *CreateFirewallRuleParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *CreateFirewallRuleParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *CreateFirewallRuleParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *CreateFirewallRuleParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *CreateFirewallRuleParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *CreateFirewallRuleParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *CreateFirewallRuleParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new CreateFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreateFirewallRuleParams(ipaddressid string, protocol string) *CreateFirewallRuleParams {
	p := &CreateFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["protocol"] = protocol
	return p
}

// Creates a firewall rule for a given IP address
func (s *FirewallService) CreateFirewallRule(p *CreateFirewallRuleParams) (*CreateFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("createFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateFirewallRuleResponse struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
	Traffictype  string `json:"traffictype"`
}

type CreatePortForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *CreatePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := p.p["privateendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateendport", vv)
	}
	if v, found := p.p["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["publicendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicendport", vv)
	}
	if v, found := p.p["publicport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicport", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (p *CreatePortForwardingRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *CreatePortForwardingRuleParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *CreatePortForwardingRuleParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreatePortForwardingRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreatePortForwardingRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *CreatePortForwardingRuleParams) ResetIpaddressid() {
	if p.p != nil && p.p["ipaddressid"] != nil {
		delete(p.p, "ipaddressid")
	}
}

func (p *CreatePortForwardingRuleParams) GetIpaddressid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddressid"].(string)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreatePortForwardingRuleParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreatePortForwardingRuleParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetOpenfirewall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["openfirewall"] = v
}

func (p *CreatePortForwardingRuleParams) ResetOpenfirewall() {
	if p.p != nil && p.p["openfirewall"] != nil {
		delete(p.p, "openfirewall")
	}
}

func (p *CreatePortForwardingRuleParams) GetOpenfirewall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["openfirewall"].(bool)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetPrivateendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateendport"] = v
}

func (p *CreatePortForwardingRuleParams) ResetPrivateendport() {
	if p.p != nil && p.p["privateendport"] != nil {
		delete(p.p, "privateendport")
	}
}

func (p *CreatePortForwardingRuleParams) GetPrivateendport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privateendport"].(int)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetPrivateport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateport"] = v
}

func (p *CreatePortForwardingRuleParams) ResetPrivateport() {
	if p.p != nil && p.p["privateport"] != nil {
		delete(p.p, "privateport")
	}
}

func (p *CreatePortForwardingRuleParams) GetPrivateport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privateport"].(int)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *CreatePortForwardingRuleParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *CreatePortForwardingRuleParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetPublicendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicendport"] = v
}

func (p *CreatePortForwardingRuleParams) ResetPublicendport() {
	if p.p != nil && p.p["publicendport"] != nil {
		delete(p.p, "publicendport")
	}
}

func (p *CreatePortForwardingRuleParams) GetPublicendport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicendport"].(int)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetPublicport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicport"] = v
}

func (p *CreatePortForwardingRuleParams) ResetPublicport() {
	if p.p != nil && p.p["publicport"] != nil {
		delete(p.p, "publicport")
	}
}

func (p *CreatePortForwardingRuleParams) GetPublicport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicport"].(int)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *CreatePortForwardingRuleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *CreatePortForwardingRuleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *CreatePortForwardingRuleParams) SetVmguestip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmguestip"] = v
}

func (p *CreatePortForwardingRuleParams) ResetVmguestip() {
	if p.p != nil && p.p["vmguestip"] != nil {
		delete(p.p, "vmguestip")
	}
}

func (p *CreatePortForwardingRuleParams) GetVmguestip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmguestip"].(string)
	return value, ok
}

// You should always use this function to get a new CreatePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreatePortForwardingRuleParams(ipaddressid string, privateport int, protocol string, publicport int, virtualmachineid string) *CreatePortForwardingRuleParams {
	p := &CreatePortForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["privateport"] = privateport
	p.p["protocol"] = protocol
	p.p["publicport"] = publicport
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Creates a port forwarding rule
func (s *FirewallService) CreatePortForwardingRule(p *CreatePortForwardingRuleParams) (*CreatePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("createPortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePortForwardingRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreatePortForwardingRuleResponse struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}

type DeleteEgressFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteEgressFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteEgressFirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteEgressFirewallRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteEgressFirewallRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteEgressFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeleteEgressFirewallRuleParams(id string) *DeleteEgressFirewallRuleParams {
	p := &DeleteEgressFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an egress firewall rule
func (s *FirewallService) DeleteEgressFirewallRule(p *DeleteEgressFirewallRuleParams) (*DeleteEgressFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteEgressFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteEgressFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteEgressFirewallRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteFirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteFirewallRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteFirewallRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeleteFirewallRuleParams(id string) *DeleteFirewallRuleParams {
	p := &DeleteFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a firewall rule
func (s *FirewallService) DeleteFirewallRule(p *DeleteFirewallRuleParams) (*DeleteFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteFirewallRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeletePaloAltoFirewallParams struct {
	p map[string]interface{}
}

func (p *DeletePaloAltoFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
	}
	return u
}

func (p *DeletePaloAltoFirewallParams) SetFwdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fwdeviceid"] = v
}

func (p *DeletePaloAltoFirewallParams) ResetFwdeviceid() {
	if p.p != nil && p.p["fwdeviceid"] != nil {
		delete(p.p, "fwdeviceid")
	}
}

func (p *DeletePaloAltoFirewallParams) GetFwdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fwdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeletePaloAltoFirewallParams(fwdeviceid string) *DeletePaloAltoFirewallParams {
	p := &DeletePaloAltoFirewallParams{}
	p.p = make(map[string]interface{})
	p.p["fwdeviceid"] = fwdeviceid
	return p
}

// delete a Palo Alto firewall device
func (s *FirewallService) DeletePaloAltoFirewall(p *DeletePaloAltoFirewallParams) (*DeletePaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("deletePaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePaloAltoFirewallResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeletePaloAltoFirewallResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeletePortForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *DeletePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePortForwardingRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeletePortForwardingRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeletePortForwardingRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeletePortForwardingRuleParams(id string) *DeletePortForwardingRuleParams {
	p := &DeletePortForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a port forwarding rule
func (s *FirewallService) DeletePortForwardingRule(p *DeletePortForwardingRuleParams) (*DeletePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("deletePortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePortForwardingRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeletePortForwardingRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListEgressFirewallRulesParams struct {
	p map[string]interface{}
}

func (p *ListEgressFirewallRulesParams) toURLValues() url.Values {
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
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
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
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListEgressFirewallRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListEgressFirewallRulesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListEgressFirewallRulesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListEgressFirewallRulesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListEgressFirewallRulesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListEgressFirewallRulesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListEgressFirewallRulesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListEgressFirewallRulesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListEgressFirewallRulesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *ListEgressFirewallRulesParams) ResetIpaddressid() {
	if p.p != nil && p.p["ipaddressid"] != nil {
		delete(p.p, "ipaddressid")
	}
}

func (p *ListEgressFirewallRulesParams) GetIpaddressid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddressid"].(string)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListEgressFirewallRulesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListEgressFirewallRulesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListEgressFirewallRulesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListEgressFirewallRulesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListEgressFirewallRulesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListEgressFirewallRulesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListEgressFirewallRulesParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListEgressFirewallRulesParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListEgressFirewallRulesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListEgressFirewallRulesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListEgressFirewallRulesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListEgressFirewallRulesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListEgressFirewallRulesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListEgressFirewallRulesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListEgressFirewallRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListEgressFirewallRulesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListEgressFirewallRulesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListEgressFirewallRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListEgressFirewallRulesParams() *ListEgressFirewallRulesParams {
	p := &ListEgressFirewallRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetEgressFirewallRuleByID(id string, opts ...OptionFunc) (*EgressFirewallRule, int, error) {
	p := &ListEgressFirewallRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListEgressFirewallRules(p)
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
		return l.EgressFirewallRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for EgressFirewallRule UUID: %s!", id)
}

// Lists all egress firewall rules for network ID.
func (s *FirewallService) ListEgressFirewallRules(p *ListEgressFirewallRulesParams) (*ListEgressFirewallRulesResponse, error) {
	resp, err := s.cs.newRequest("listEgressFirewallRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListEgressFirewallRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListEgressFirewallRulesResponse struct {
	Count               int                   `json:"count"`
	EgressFirewallRules []*EgressFirewallRule `json:"firewallrule"`
}

type EgressFirewallRule struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
	Traffictype  string `json:"traffictype"`
}

type ListFirewallRulesParams struct {
	p map[string]interface{}
}

func (p *ListFirewallRulesParams) toURLValues() url.Values {
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
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
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
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListFirewallRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListFirewallRulesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListFirewallRulesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListFirewallRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListFirewallRulesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListFirewallRulesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListFirewallRulesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListFirewallRulesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListFirewallRulesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListFirewallRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListFirewallRulesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListFirewallRulesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListFirewallRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *ListFirewallRulesParams) ResetIpaddressid() {
	if p.p != nil && p.p["ipaddressid"] != nil {
		delete(p.p, "ipaddressid")
	}
}

func (p *ListFirewallRulesParams) GetIpaddressid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddressid"].(string)
	return value, ok
}

func (p *ListFirewallRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListFirewallRulesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListFirewallRulesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListFirewallRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListFirewallRulesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListFirewallRulesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListFirewallRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListFirewallRulesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListFirewallRulesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListFirewallRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListFirewallRulesParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListFirewallRulesParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListFirewallRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListFirewallRulesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListFirewallRulesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListFirewallRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListFirewallRulesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListFirewallRulesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListFirewallRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListFirewallRulesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListFirewallRulesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListFirewallRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListFirewallRulesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListFirewallRulesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListFirewallRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListFirewallRulesParams() *ListFirewallRulesParams {
	p := &ListFirewallRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetFirewallRuleByID(id string, opts ...OptionFunc) (*FirewallRule, int, error) {
	p := &ListFirewallRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListFirewallRules(p)
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
		return l.FirewallRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for FirewallRule UUID: %s!", id)
}

// Lists all firewall rules for an IP address.
func (s *FirewallService) ListFirewallRules(p *ListFirewallRulesParams) (*ListFirewallRulesResponse, error) {
	resp, err := s.cs.newRequest("listFirewallRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListFirewallRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListFirewallRulesResponse struct {
	Count         int             `json:"count"`
	FirewallRules []*FirewallRule `json:"firewallrule"`
}

type FirewallRule struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
	Traffictype  string `json:"traffictype"`
}

type ListPaloAltoFirewallsParams struct {
	p map[string]interface{}
}

func (p *ListPaloAltoFirewallsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
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

func (p *ListPaloAltoFirewallsParams) SetFwdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fwdeviceid"] = v
}

func (p *ListPaloAltoFirewallsParams) ResetFwdeviceid() {
	if p.p != nil && p.p["fwdeviceid"] != nil {
		delete(p.p, "fwdeviceid")
	}
}

func (p *ListPaloAltoFirewallsParams) GetFwdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fwdeviceid"].(string)
	return value, ok
}

func (p *ListPaloAltoFirewallsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListPaloAltoFirewallsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListPaloAltoFirewallsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListPaloAltoFirewallsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListPaloAltoFirewallsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListPaloAltoFirewallsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListPaloAltoFirewallsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListPaloAltoFirewallsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListPaloAltoFirewallsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListPaloAltoFirewallsParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListPaloAltoFirewallsParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *ListPaloAltoFirewallsParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPaloAltoFirewallsParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListPaloAltoFirewallsParams() *ListPaloAltoFirewallsParams {
	p := &ListPaloAltoFirewallsParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists Palo Alto firewall devices in a physical network
func (s *FirewallService) ListPaloAltoFirewalls(p *ListPaloAltoFirewallsParams) (*ListPaloAltoFirewallsResponse, error) {
	resp, err := s.cs.newRequest("listPaloAltoFirewalls", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListPaloAltoFirewallsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPaloAltoFirewallsResponse struct {
	Count             int                 `json:"count"`
	PaloAltoFirewalls []*PaloAltoFirewall `json:"paloaltofirewall"`
}

type PaloAltoFirewall struct {
	Fwdevicecapacity  int64  `json:"fwdevicecapacity"`
	Fwdeviceid        string `json:"fwdeviceid"`
	Fwdevicename      string `json:"fwdevicename"`
	Fwdevicestate     string `json:"fwdevicestate"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Numretries        string `json:"numretries"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Privatezone       string `json:"privatezone"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
	Publiczone        string `json:"publiczone"`
	Timeout           string `json:"timeout"`
	Usageinterface    string `json:"usageinterface"`
	Username          string `json:"username"`
	Zoneid            string `json:"zoneid"`
}

type ListPortForwardingRulesParams struct {
	p map[string]interface{}
}

func (p *ListPortForwardingRulesParams) toURLValues() url.Values {
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
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
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
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListPortForwardingRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListPortForwardingRulesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListPortForwardingRulesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListPortForwardingRulesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListPortForwardingRulesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListPortForwardingRulesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListPortForwardingRulesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListPortForwardingRulesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListPortForwardingRulesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *ListPortForwardingRulesParams) ResetIpaddressid() {
	if p.p != nil && p.p["ipaddressid"] != nil {
		delete(p.p, "ipaddressid")
	}
}

func (p *ListPortForwardingRulesParams) GetIpaddressid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddressid"].(string)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListPortForwardingRulesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListPortForwardingRulesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListPortForwardingRulesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListPortForwardingRulesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListPortForwardingRulesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListPortForwardingRulesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListPortForwardingRulesParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListPortForwardingRulesParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListPortForwardingRulesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListPortForwardingRulesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListPortForwardingRulesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListPortForwardingRulesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListPortForwardingRulesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListPortForwardingRulesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListPortForwardingRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListPortForwardingRulesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListPortForwardingRulesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListPortForwardingRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListPortForwardingRulesParams() *ListPortForwardingRulesParams {
	p := &ListPortForwardingRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetPortForwardingRuleByID(id string, opts ...OptionFunc) (*PortForwardingRule, int, error) {
	p := &ListPortForwardingRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPortForwardingRules(p)
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
		return l.PortForwardingRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for PortForwardingRule UUID: %s!", id)
}

// Lists all port forwarding rules for an IP address.
func (s *FirewallService) ListPortForwardingRules(p *ListPortForwardingRulesParams) (*ListPortForwardingRulesResponse, error) {
	resp, err := s.cs.newRequest("listPortForwardingRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListPortForwardingRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPortForwardingRulesResponse struct {
	Count               int                   `json:"count"`
	PortForwardingRules []*PortForwardingRule `json:"portforwardingrule"`
}

type PortForwardingRule struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}

type UpdateEgressFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateEgressFirewallRuleParams) toURLValues() url.Values {
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

func (p *UpdateEgressFirewallRuleParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateEgressFirewallRuleParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateEgressFirewallRuleParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateEgressFirewallRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateEgressFirewallRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateEgressFirewallRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateEgressFirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateEgressFirewallRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateEgressFirewallRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateEgressFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewUpdateEgressFirewallRuleParams(id string) *UpdateEgressFirewallRuleParams {
	p := &UpdateEgressFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates egress firewall rule
func (s *FirewallService) UpdateEgressFirewallRule(p *UpdateEgressFirewallRuleParams) (*UpdateEgressFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("updateEgressFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateEgressFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateEgressFirewallRuleResponse struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
	Traffictype  string `json:"traffictype"`
}

type UpdateFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateFirewallRuleParams) toURLValues() url.Values {
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

func (p *UpdateFirewallRuleParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateFirewallRuleParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateFirewallRuleParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateFirewallRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateFirewallRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateFirewallRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateFirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateFirewallRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateFirewallRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewUpdateFirewallRuleParams(id string) *UpdateFirewallRuleParams {
	p := &UpdateFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates firewall rule
func (s *FirewallService) UpdateFirewallRule(p *UpdateFirewallRuleParams) (*UpdateFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("updateFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateFirewallRuleResponse struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
	Traffictype  string `json:"traffictype"`
}

type UpdatePortForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *UpdatePortForwardingRuleParams) toURLValues() url.Values {
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
	if v, found := p.p["privateendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateendport", vv)
	}
	if v, found := p.p["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (p *UpdatePortForwardingRuleParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdatePortForwardingRuleParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdatePortForwardingRuleParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdatePortForwardingRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdatePortForwardingRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdatePortForwardingRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdatePortForwardingRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdatePortForwardingRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdatePortForwardingRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdatePortForwardingRuleParams) SetPrivateendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateendport"] = v
}

func (p *UpdatePortForwardingRuleParams) ResetPrivateendport() {
	if p.p != nil && p.p["privateendport"] != nil {
		delete(p.p, "privateendport")
	}
}

func (p *UpdatePortForwardingRuleParams) GetPrivateendport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privateendport"].(int)
	return value, ok
}

func (p *UpdatePortForwardingRuleParams) SetPrivateport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateport"] = v
}

func (p *UpdatePortForwardingRuleParams) ResetPrivateport() {
	if p.p != nil && p.p["privateport"] != nil {
		delete(p.p, "privateport")
	}
}

func (p *UpdatePortForwardingRuleParams) GetPrivateport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privateport"].(int)
	return value, ok
}

func (p *UpdatePortForwardingRuleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *UpdatePortForwardingRuleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *UpdatePortForwardingRuleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *UpdatePortForwardingRuleParams) SetVmguestip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmguestip"] = v
}

func (p *UpdatePortForwardingRuleParams) ResetVmguestip() {
	if p.p != nil && p.p["vmguestip"] != nil {
		delete(p.p, "vmguestip")
	}
}

func (p *UpdatePortForwardingRuleParams) GetVmguestip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmguestip"].(string)
	return value, ok
}

// You should always use this function to get a new UpdatePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewUpdatePortForwardingRuleParams(id string) *UpdatePortForwardingRuleParams {
	p := &UpdatePortForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a port forwarding rule. Only the private port and the virtual machine can be updated.
func (s *FirewallService) UpdatePortForwardingRule(p *UpdatePortForwardingRuleParams) (*UpdatePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("updatePortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdatePortForwardingRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdatePortForwardingRuleResponse struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}

type ListIpv6FirewallRulesParams struct {
	p map[string]interface{}
}

func (p *ListIpv6FirewallRulesParams) toURLValues() url.Values {
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
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *ListIpv6FirewallRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListIpv6FirewallRulesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListIpv6FirewallRulesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListIpv6FirewallRulesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListIpv6FirewallRulesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListIpv6FirewallRulesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListIpv6FirewallRulesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListIpv6FirewallRulesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListIpv6FirewallRulesParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListIpv6FirewallRulesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListIpv6FirewallRulesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListIpv6FirewallRulesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListIpv6FirewallRulesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListIpv6FirewallRulesParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *ListIpv6FirewallRulesParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
}

func (p *ListIpv6FirewallRulesParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new ListIpv6FirewallRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListIpv6FirewallRulesParams() *ListIpv6FirewallRulesParams {
	p := &ListIpv6FirewallRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetIpv6FirewallRuleByID(id string, opts ...OptionFunc) (*Ipv6FirewallRule, int, error) {
	p := &ListIpv6FirewallRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListIpv6FirewallRules(p)
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
		return l.Ipv6FirewallRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Ipv6FirewallRule UUID: %s!", id)
}

// Lists all IPv6 firewall rules
func (s *FirewallService) ListIpv6FirewallRules(p *ListIpv6FirewallRulesParams) (*ListIpv6FirewallRulesResponse, error) {
	resp, err := s.cs.newRequest("listIpv6FirewallRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListIpv6FirewallRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListIpv6FirewallRulesResponse struct {
	Count             int                 `json:"count"`
	Ipv6FirewallRules []*Ipv6FirewallRule `json:"ipv6firewallrule"`
}

type Ipv6FirewallRule struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}

type CreateIpv6FirewallRuleParams struct {
	p map[string]interface{}
}

func (p *CreateIpv6FirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["destcidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("destcidrlist", vv)
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *CreateIpv6FirewallRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetDestcidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destcidrlist"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetDestcidrlist() {
	if p.p != nil && p.p["destcidrlist"] != nil {
		delete(p.p, "destcidrlist")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetDestcidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destcidrlist"].([]string)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetIcmpcode() {
	if p.p != nil && p.p["icmpcode"] != nil {
		delete(p.p, "icmpcode")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetIcmpcode() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmpcode"].(int)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetIcmptype() {
	if p.p != nil && p.p["icmptype"] != nil {
		delete(p.p, "icmptype")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetIcmptype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmptype"].(int)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *CreateIpv6FirewallRuleParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *CreateIpv6FirewallRuleParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
}

func (p *CreateIpv6FirewallRuleParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new CreateIpv6FirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreateIpv6FirewallRuleParams(networkid string, protocol string) *CreateIpv6FirewallRuleParams {
	p := &CreateIpv6FirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	p.p["protocol"] = protocol
	return p
}

// Creates an Ipv6 firewall rule in the given network (the network must not belong to VPC)
func (s *FirewallService) CreateIpv6FirewallRule(p *CreateIpv6FirewallRuleParams) (*CreateIpv6FirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("createIpv6FirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateIpv6FirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateIpv6FirewallRuleResponse struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}

type UpdateIpv6FirewallRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateIpv6FirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *UpdateIpv6FirewallRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetIcmpcode() {
	if p.p != nil && p.p["icmpcode"] != nil {
		delete(p.p, "icmpcode")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetIcmpcode() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmpcode"].(int)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetIcmptype() {
	if p.p != nil && p.p["icmptype"] != nil {
		delete(p.p, "icmptype")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetIcmptype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmptype"].(int)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *UpdateIpv6FirewallRuleParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *UpdateIpv6FirewallRuleParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
}

func (p *UpdateIpv6FirewallRuleParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateIpv6FirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewUpdateIpv6FirewallRuleParams(id string) *UpdateIpv6FirewallRuleParams {
	p := &UpdateIpv6FirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates Ipv6 firewall rule with specified ID
func (s *FirewallService) UpdateIpv6FirewallRule(p *UpdateIpv6FirewallRuleParams) (*UpdateIpv6FirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("updateIpv6FirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIpv6FirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateIpv6FirewallRuleResponse struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}

type DeleteIpv6FirewallRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteIpv6FirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteIpv6FirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteIpv6FirewallRuleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteIpv6FirewallRuleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteIpv6FirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeleteIpv6FirewallRuleParams(id string) *DeleteIpv6FirewallRuleParams {
	p := &DeleteIpv6FirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a IPv6 firewall rule
func (s *FirewallService) DeleteIpv6FirewallRule(p *DeleteIpv6FirewallRuleParams) (*DeleteIpv6FirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteIpv6FirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteIpv6FirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteIpv6FirewallRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
