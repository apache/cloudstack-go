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

//  delete a Palo Alto firewall device
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
