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
func convertAuthorizeSecurityGroupIngressResponse(b []byte) ([]byte, error) {
	var raw struct {
		Ingressrule []interface{} `json:"ingressrule"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if len(raw.Ingressrule) != 1 {
		return b, nil
	}

	return json.Marshal(raw.Ingressrule[0])
}

// Helper function for maintaining backwards compatibility
func convertAuthorizeSecurityGroupEgressResponse(b []byte) ([]byte, error) {
	var raw struct {
		Egressrule []interface{} `json:"egressrule"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if len(raw.Egressrule) != 1 {
		return b, nil
	}

	return json.Marshal(raw.Egressrule[0])
}

type SecurityGroupServiceIface interface {
	AuthorizeSecurityGroupEgress(p *AuthorizeSecurityGroupEgressParams) (*AuthorizeSecurityGroupEgressResponse, error)
	NewAuthorizeSecurityGroupEgressParams() *AuthorizeSecurityGroupEgressParams
	AuthorizeSecurityGroupIngress(p *AuthorizeSecurityGroupIngressParams) (*AuthorizeSecurityGroupIngressResponse, error)
	NewAuthorizeSecurityGroupIngressParams() *AuthorizeSecurityGroupIngressParams
	CreateSecurityGroup(p *CreateSecurityGroupParams) (*CreateSecurityGroupResponse, error)
	NewCreateSecurityGroupParams(name string) *CreateSecurityGroupParams
	DeleteSecurityGroup(p *DeleteSecurityGroupParams) (*DeleteSecurityGroupResponse, error)
	NewDeleteSecurityGroupParams() *DeleteSecurityGroupParams
	ListSecurityGroups(p *ListSecurityGroupsParams) (*ListSecurityGroupsResponse, error)
	NewListSecurityGroupsParams() *ListSecurityGroupsParams
	GetSecurityGroupID(keyword string, opts ...OptionFunc) (string, int, error)
	GetSecurityGroupByName(name string, opts ...OptionFunc) (*SecurityGroup, int, error)
	GetSecurityGroupByID(id string, opts ...OptionFunc) (*SecurityGroup, int, error)
	RevokeSecurityGroupEgress(p *RevokeSecurityGroupEgressParams) (*RevokeSecurityGroupEgressResponse, error)
	NewRevokeSecurityGroupEgressParams(id string) *RevokeSecurityGroupEgressParams
	RevokeSecurityGroupIngress(p *RevokeSecurityGroupIngressParams) (*RevokeSecurityGroupIngressResponse, error)
	NewRevokeSecurityGroupIngressParams(id string) *RevokeSecurityGroupIngressParams
	UpdateSecurityGroup(p *UpdateSecurityGroupParams) (*UpdateSecurityGroupResponse, error)
	NewUpdateSecurityGroupParams(id string) *UpdateSecurityGroupParams
}

type AuthorizeSecurityGroupEgressParams struct {
	p map[string]interface{}
}

func (p *AuthorizeSecurityGroupEgressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := p.p["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["usersecuritygrouplist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].account", i), k)
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].group", i), m[k])
		}
	}
	return u
}

func (p *AuthorizeSecurityGroupEgressParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetIcmpcode() {
	if p.p != nil && p.p["icmpcode"] != nil {
		delete(p.p, "icmpcode")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetIcmpcode() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmpcode"].(int)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetIcmptype() {
	if p.p != nil && p.p["icmptype"] != nil {
		delete(p.p, "icmptype")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetIcmptype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmptype"].(int)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetSecuritygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupid"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetSecuritygroupid() {
	if p.p != nil && p.p["securitygroupid"] != nil {
		delete(p.p, "securitygroupid")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetSecuritygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupid"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetSecuritygroupname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupname"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetSecuritygroupname() {
	if p.p != nil && p.p["securitygroupname"] != nil {
		delete(p.p, "securitygroupname")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetSecuritygroupname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupname"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *AuthorizeSecurityGroupEgressParams) SetUsersecuritygrouplist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usersecuritygrouplist"] = v
}

func (p *AuthorizeSecurityGroupEgressParams) ResetUsersecuritygrouplist() {
	if p.p != nil && p.p["usersecuritygrouplist"] != nil {
		delete(p.p, "usersecuritygrouplist")
	}
}

func (p *AuthorizeSecurityGroupEgressParams) GetUsersecuritygrouplist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["usersecuritygrouplist"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new AuthorizeSecurityGroupEgressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewAuthorizeSecurityGroupEgressParams() *AuthorizeSecurityGroupEgressParams {
	p := &AuthorizeSecurityGroupEgressParams{}
	p.p = make(map[string]interface{})
	return p
}

// Authorizes a particular egress rule for this security group
func (s *SecurityGroupService) AuthorizeSecurityGroupEgress(p *AuthorizeSecurityGroupEgressParams) (*AuthorizeSecurityGroupEgressResponse, error) {
	resp, err := s.cs.newPostRequest("authorizeSecurityGroupEgress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AuthorizeSecurityGroupEgressResponse
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

		b, err = convertAuthorizeSecurityGroupEgressResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AuthorizeSecurityGroupEgressResponse struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type AuthorizeSecurityGroupIngressParams struct {
	p map[string]interface{}
}

func (p *AuthorizeSecurityGroupIngressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := p.p["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["usersecuritygrouplist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].account", i), k)
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].group", i), m[k])
		}
	}
	return u
}

func (p *AuthorizeSecurityGroupIngressParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetCidrlist() {
	if p.p != nil && p.p["cidrlist"] != nil {
		delete(p.p, "cidrlist")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetCidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrlist"].([]string)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetIcmpcode() {
	if p.p != nil && p.p["icmpcode"] != nil {
		delete(p.p, "icmpcode")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetIcmpcode() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmpcode"].(int)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetIcmptype() {
	if p.p != nil && p.p["icmptype"] != nil {
		delete(p.p, "icmptype")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetIcmptype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["icmptype"].(int)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetSecuritygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupid"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetSecuritygroupid() {
	if p.p != nil && p.p["securitygroupid"] != nil {
		delete(p.p, "securitygroupid")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetSecuritygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupid"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetSecuritygroupname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupname"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetSecuritygroupname() {
	if p.p != nil && p.p["securitygroupname"] != nil {
		delete(p.p, "securitygroupname")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetSecuritygroupname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupname"].(string)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *AuthorizeSecurityGroupIngressParams) SetUsersecuritygrouplist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usersecuritygrouplist"] = v
}

func (p *AuthorizeSecurityGroupIngressParams) ResetUsersecuritygrouplist() {
	if p.p != nil && p.p["usersecuritygrouplist"] != nil {
		delete(p.p, "usersecuritygrouplist")
	}
}

func (p *AuthorizeSecurityGroupIngressParams) GetUsersecuritygrouplist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["usersecuritygrouplist"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new AuthorizeSecurityGroupIngressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewAuthorizeSecurityGroupIngressParams() *AuthorizeSecurityGroupIngressParams {
	p := &AuthorizeSecurityGroupIngressParams{}
	p.p = make(map[string]interface{})
	return p
}

// Authorizes a particular ingress rule for this security group
func (s *SecurityGroupService) AuthorizeSecurityGroupIngress(p *AuthorizeSecurityGroupIngressParams) (*AuthorizeSecurityGroupIngressResponse, error) {
	resp, err := s.cs.newPostRequest("authorizeSecurityGroupIngress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AuthorizeSecurityGroupIngressResponse
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

		b, err = convertAuthorizeSecurityGroupIngressResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AuthorizeSecurityGroupIngressResponse struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type CreateSecurityGroupParams struct {
	p map[string]interface{}
}

func (p *CreateSecurityGroupParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *CreateSecurityGroupParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateSecurityGroupParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateSecurityGroupParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateSecurityGroupParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateSecurityGroupParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateSecurityGroupParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateSecurityGroupParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateSecurityGroupParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateSecurityGroupParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateSecurityGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateSecurityGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateSecurityGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateSecurityGroupParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateSecurityGroupParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateSecurityGroupParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSecurityGroupParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewCreateSecurityGroupParams(name string) *CreateSecurityGroupParams {
	p := &CreateSecurityGroupParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Creates a security group
func (s *SecurityGroupService) CreateSecurityGroup(p *CreateSecurityGroupParams) (*CreateSecurityGroupResponse, error) {
	resp, err := s.cs.newPostRequest("createSecurityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateSecurityGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateSecurityGroupResponse struct {
	Account             string                            `json:"account"`
	Description         string                            `json:"description"`
	Domain              string                            `json:"domain"`
	Domainid            string                            `json:"domainid"`
	Domainpath          string                            `json:"domainpath"`
	Egressrule          []CreateSecurityGroupResponseRule `json:"egressrule"`
	Id                  string                            `json:"id"`
	Ingressrule         []CreateSecurityGroupResponseRule `json:"ingressrule"`
	JobID               string                            `json:"jobid"`
	Jobstatus           int                               `json:"jobstatus"`
	Name                string                            `json:"name"`
	Project             string                            `json:"project"`
	Projectid           string                            `json:"projectid"`
	Tags                []Tags                            `json:"tags"`
	Virtualmachinecount int                               `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                     `json:"virtualmachineids"`
}

type CreateSecurityGroupResponseRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type DeleteSecurityGroupParams struct {
	p map[string]interface{}
}

func (p *DeleteSecurityGroupParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DeleteSecurityGroupParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DeleteSecurityGroupParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DeleteSecurityGroupParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DeleteSecurityGroupParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DeleteSecurityGroupParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DeleteSecurityGroupParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DeleteSecurityGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteSecurityGroupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteSecurityGroupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteSecurityGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *DeleteSecurityGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *DeleteSecurityGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *DeleteSecurityGroupParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DeleteSecurityGroupParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *DeleteSecurityGroupParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSecurityGroupParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewDeleteSecurityGroupParams() *DeleteSecurityGroupParams {
	p := &DeleteSecurityGroupParams{}
	p.p = make(map[string]interface{})
	return p
}

// Deletes security group
func (s *SecurityGroupService) DeleteSecurityGroup(p *DeleteSecurityGroupParams) (*DeleteSecurityGroupResponse, error) {
	resp, err := s.cs.newPostRequest("deleteSecurityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSecurityGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSecurityGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSecurityGroupResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSecurityGroupResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListSecurityGroupsParams struct {
	p map[string]interface{}
}

func (p *ListSecurityGroupsParams) toURLValues() url.Values {
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
	if v, found := p.p["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
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
	return u
}

func (p *ListSecurityGroupsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListSecurityGroupsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListSecurityGroupsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListSecurityGroupsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListSecurityGroupsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSecurityGroupsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListSecurityGroupsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListSecurityGroupsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListSecurityGroupsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSecurityGroupsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSecurityGroupsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListSecurityGroupsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListSecurityGroupsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSecurityGroupsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSecurityGroupsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSecurityGroupsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSecurityGroupsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListSecurityGroupsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListSecurityGroupsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetSecuritygroupname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupname"] = v
}

func (p *ListSecurityGroupsParams) ResetSecuritygroupname() {
	if p.p != nil && p.p["securitygroupname"] != nil {
		delete(p.p, "securitygroupname")
	}
}

func (p *ListSecurityGroupsParams) GetSecuritygroupname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupname"].(string)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListSecurityGroupsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListSecurityGroupsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListSecurityGroupsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListSecurityGroupsParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *ListSecurityGroupsParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSecurityGroupsParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewListSecurityGroupsParams() *ListSecurityGroupsParams {
	p := &ListSecurityGroupsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListSecurityGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSecurityGroups(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.SecurityGroups[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SecurityGroups {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupByName(name string, opts ...OptionFunc) (*SecurityGroup, int, error) {
	id, count, err := s.GetSecurityGroupID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSecurityGroupByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupByID(id string, opts ...OptionFunc) (*SecurityGroup, int, error) {
	p := &ListSecurityGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSecurityGroups(p)
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
		return l.SecurityGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SecurityGroup UUID: %s!", id)
}

// Lists security groups
func (s *SecurityGroupService) ListSecurityGroups(p *ListSecurityGroupsParams) (*ListSecurityGroupsResponse, error) {
	resp, err := s.cs.newRequest("listSecurityGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSecurityGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSecurityGroupsResponse struct {
	Count          int              `json:"count"`
	SecurityGroups []*SecurityGroup `json:"securitygroup"`
}

type SecurityGroup struct {
	Account             string              `json:"account"`
	Description         string              `json:"description"`
	Domain              string              `json:"domain"`
	Domainid            string              `json:"domainid"`
	Domainpath          string              `json:"domainpath"`
	Egressrule          []SecurityGroupRule `json:"egressrule"`
	Id                  string              `json:"id"`
	Ingressrule         []SecurityGroupRule `json:"ingressrule"`
	JobID               string              `json:"jobid"`
	Jobstatus           int                 `json:"jobstatus"`
	Name                string              `json:"name"`
	Project             string              `json:"project"`
	Projectid           string              `json:"projectid"`
	Tags                []Tags              `json:"tags"`
	Virtualmachinecount int                 `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}       `json:"virtualmachineids"`
}

type SecurityGroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type RevokeSecurityGroupEgressParams struct {
	p map[string]interface{}
}

func (p *RevokeSecurityGroupEgressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RevokeSecurityGroupEgressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RevokeSecurityGroupEgressParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RevokeSecurityGroupEgressParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RevokeSecurityGroupEgressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewRevokeSecurityGroupEgressParams(id string) *RevokeSecurityGroupEgressParams {
	p := &RevokeSecurityGroupEgressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a particular egress rule from this security group
func (s *SecurityGroupService) RevokeSecurityGroupEgress(p *RevokeSecurityGroupEgressParams) (*RevokeSecurityGroupEgressResponse, error) {
	resp, err := s.cs.newPostRequest("revokeSecurityGroupEgress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeSecurityGroupEgressResponse
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

type RevokeSecurityGroupEgressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RevokeSecurityGroupIngressParams struct {
	p map[string]interface{}
}

func (p *RevokeSecurityGroupIngressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RevokeSecurityGroupIngressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RevokeSecurityGroupIngressParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RevokeSecurityGroupIngressParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RevokeSecurityGroupIngressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewRevokeSecurityGroupIngressParams(id string) *RevokeSecurityGroupIngressParams {
	p := &RevokeSecurityGroupIngressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a particular ingress rule from this security group
func (s *SecurityGroupService) RevokeSecurityGroupIngress(p *RevokeSecurityGroupIngressParams) (*RevokeSecurityGroupIngressResponse, error) {
	resp, err := s.cs.newPostRequest("revokeSecurityGroupIngress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeSecurityGroupIngressResponse
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

type RevokeSecurityGroupIngressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateSecurityGroupParams struct {
	p map[string]interface{}
}

func (p *UpdateSecurityGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateSecurityGroupParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateSecurityGroupParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateSecurityGroupParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateSecurityGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateSecurityGroupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateSecurityGroupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateSecurityGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateSecurityGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateSecurityGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateSecurityGroupParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewUpdateSecurityGroupParams(id string) *UpdateSecurityGroupParams {
	p := &UpdateSecurityGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a security group
func (s *SecurityGroupService) UpdateSecurityGroup(p *UpdateSecurityGroupParams) (*UpdateSecurityGroupResponse, error) {
	resp, err := s.cs.newPostRequest("updateSecurityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response UpdateSecurityGroupResponse `json:"securitygroup"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type UpdateSecurityGroupResponse struct {
	Account             string                            `json:"account"`
	Description         string                            `json:"description"`
	Domain              string                            `json:"domain"`
	Domainid            string                            `json:"domainid"`
	Domainpath          string                            `json:"domainpath"`
	Egressrule          []UpdateSecurityGroupResponseRule `json:"egressrule"`
	Id                  string                            `json:"id"`
	Ingressrule         []UpdateSecurityGroupResponseRule `json:"ingressrule"`
	JobID               string                            `json:"jobid"`
	Jobstatus           int                               `json:"jobstatus"`
	Name                string                            `json:"name"`
	Project             string                            `json:"project"`
	Projectid           string                            `json:"projectid"`
	Tags                []Tags                            `json:"tags"`
	Virtualmachinecount int                               `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                     `json:"virtualmachineids"`
}

type UpdateSecurityGroupResponseRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}
