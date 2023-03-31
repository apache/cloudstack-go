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

type LDAPServiceIface interface {
	AddLdapConfiguration(p *AddLdapConfigurationParams) (*AddLdapConfigurationResponse, error)
	NewAddLdapConfigurationParams(hostname string, port int) *AddLdapConfigurationParams
	DeleteLdapConfiguration(p *DeleteLdapConfigurationParams) (*DeleteLdapConfigurationResponse, error)
	NewDeleteLdapConfigurationParams(hostname string) *DeleteLdapConfigurationParams
	ImportLdapUsers(p *ImportLdapUsersParams) (*ImportLdapUsersResponse, error)
	NewImportLdapUsersParams() *ImportLdapUsersParams
	LdapConfig(p *LdapConfigParams) (*LdapConfigResponse, error)
	NewLdapConfigParams() *LdapConfigParams
	LdapCreateAccount(p *LdapCreateAccountParams) (*LdapCreateAccountResponse, error)
	NewLdapCreateAccountParams(username string) *LdapCreateAccountParams
	LdapRemove(p *LdapRemoveParams) (*LdapRemoveResponse, error)
	NewLdapRemoveParams() *LdapRemoveParams
	LinkDomainToLdap(p *LinkDomainToLdapParams) (*LinkDomainToLdapResponse, error)
	NewLinkDomainToLdapParams(accounttype int, domainid string, lDAPType string) *LinkDomainToLdapParams
	ListLdapConfigurations(p *ListLdapConfigurationsParams) (*ListLdapConfigurationsResponse, error)
	NewListLdapConfigurationsParams() *ListLdapConfigurationsParams
	ListLdapUsers(p *ListLdapUsersParams) (*ListLdapUsersResponse, error)
	NewListLdapUsersParams() *ListLdapUsersParams
	SearchLdap(p *SearchLdapParams) (*SearchLdapResponse, error)
	NewSearchLdapParams(query string) *SearchLdapParams
}

type AddLdapConfigurationParams struct {
	p map[string]interface{}
}

func (p *AddLdapConfigurationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (p *AddLdapConfigurationParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *AddLdapConfigurationParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *AddLdapConfigurationParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *AddLdapConfigurationParams) GetHostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostname"].(string)
	return value, ok
}

func (p *AddLdapConfigurationParams) SetPort(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["port"] = v
}

func (p *AddLdapConfigurationParams) GetPort() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["port"].(int)
	return value, ok
}

// You should always use this function to get a new AddLdapConfigurationParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewAddLdapConfigurationParams(hostname string, port int) *AddLdapConfigurationParams {
	p := &AddLdapConfigurationParams{}
	p.p = make(map[string]interface{})
	p.p["hostname"] = hostname
	p.p["port"] = port
	return p
}

// Add a new Ldap Configuration
func (s *LDAPService) AddLdapConfiguration(p *AddLdapConfigurationParams) (*AddLdapConfigurationResponse, error) {
	resp, err := s.cs.newRequest("addLdapConfiguration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddLdapConfigurationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddLdapConfigurationResponse struct {
	Domainid  string `json:"domainid"`
	Hostname  string `json:"hostname"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Port      int    `json:"port"`
}

type DeleteLdapConfigurationParams struct {
	p map[string]interface{}
}

func (p *DeleteLdapConfigurationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (p *DeleteLdapConfigurationParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DeleteLdapConfigurationParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DeleteLdapConfigurationParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *DeleteLdapConfigurationParams) GetHostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostname"].(string)
	return value, ok
}

func (p *DeleteLdapConfigurationParams) SetPort(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["port"] = v
}

func (p *DeleteLdapConfigurationParams) GetPort() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["port"].(int)
	return value, ok
}

// You should always use this function to get a new DeleteLdapConfigurationParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewDeleteLdapConfigurationParams(hostname string) *DeleteLdapConfigurationParams {
	p := &DeleteLdapConfigurationParams{}
	p.p = make(map[string]interface{})
	p.p["hostname"] = hostname
	return p
}

// Remove an Ldap Configuration
func (s *LDAPService) DeleteLdapConfiguration(p *DeleteLdapConfigurationParams) (*DeleteLdapConfigurationResponse, error) {
	resp, err := s.cs.newRequest("deleteLdapConfiguration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLdapConfigurationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteLdapConfigurationResponse struct {
	Domainid  string `json:"domainid"`
	Hostname  string `json:"hostname"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Port      int    `json:"port"`
}

type ImportLdapUsersParams struct {
	p map[string]interface{}
}

func (p *ImportLdapUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("accountdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("accountdetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["group"]; found {
		u.Set("group", v.(string))
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
	if v, found := p.p["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	return u
}

func (p *ImportLdapUsersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ImportLdapUsersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ImportLdapUsersParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
}

func (p *ImportLdapUsersParams) GetAccountdetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountdetails"].(map[string]string)
	return value, ok
}

func (p *ImportLdapUsersParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *ImportLdapUsersParams) GetAccounttype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounttype"].(int)
	return value, ok
}

func (p *ImportLdapUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ImportLdapUsersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ImportLdapUsersParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
}

func (p *ImportLdapUsersParams) GetGroup() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["group"].(string)
	return value, ok
}

func (p *ImportLdapUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ImportLdapUsersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ImportLdapUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ImportLdapUsersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ImportLdapUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ImportLdapUsersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ImportLdapUsersParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *ImportLdapUsersParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

func (p *ImportLdapUsersParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *ImportLdapUsersParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

// You should always use this function to get a new ImportLdapUsersParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewImportLdapUsersParams() *ImportLdapUsersParams {
	p := &ImportLdapUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Import LDAP users
func (s *LDAPService) ImportLdapUsers(p *ImportLdapUsersParams) (*ImportLdapUsersResponse, error) {
	resp, err := s.cs.newRequest("importLdapUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportLdapUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ImportLdapUsersResponse struct {
	Conflictingusersource string `json:"conflictingusersource"`
	Domain                string `json:"domain"`
	Email                 string `json:"email"`
	Firstname             string `json:"firstname"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Lastname              string `json:"lastname"`
	Principal             string `json:"principal"`
	Username              string `json:"username"`
}

type LdapConfigParams struct {
	p map[string]interface{}
}

func (p *LdapConfigParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["binddn"]; found {
		u.Set("binddn", v.(string))
	}
	if v, found := p.p["bindpass"]; found {
		u.Set("bindpass", v.(string))
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	if v, found := p.p["queryfilter"]; found {
		u.Set("queryfilter", v.(string))
	}
	if v, found := p.p["searchbase"]; found {
		u.Set("searchbase", v.(string))
	}
	if v, found := p.p["ssl"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ssl", vv)
	}
	if v, found := p.p["truststore"]; found {
		u.Set("truststore", v.(string))
	}
	if v, found := p.p["truststorepass"]; found {
		u.Set("truststorepass", v.(string))
	}
	return u
}

func (p *LdapConfigParams) SetBinddn(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["binddn"] = v
}

func (p *LdapConfigParams) GetBinddn() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["binddn"].(string)
	return value, ok
}

func (p *LdapConfigParams) SetBindpass(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bindpass"] = v
}

func (p *LdapConfigParams) GetBindpass() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bindpass"].(string)
	return value, ok
}

func (p *LdapConfigParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *LdapConfigParams) GetHostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostname"].(string)
	return value, ok
}

func (p *LdapConfigParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *LdapConfigParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *LdapConfigParams) SetPort(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["port"] = v
}

func (p *LdapConfigParams) GetPort() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["port"].(int)
	return value, ok
}

func (p *LdapConfigParams) SetQueryfilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["queryfilter"] = v
}

func (p *LdapConfigParams) GetQueryfilter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["queryfilter"].(string)
	return value, ok
}

func (p *LdapConfigParams) SetSearchbase(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["searchbase"] = v
}

func (p *LdapConfigParams) GetSearchbase() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["searchbase"].(string)
	return value, ok
}

func (p *LdapConfigParams) SetSsl(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ssl"] = v
}

func (p *LdapConfigParams) GetSsl() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ssl"].(bool)
	return value, ok
}

func (p *LdapConfigParams) SetTruststore(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["truststore"] = v
}

func (p *LdapConfigParams) GetTruststore() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["truststore"].(string)
	return value, ok
}

func (p *LdapConfigParams) SetTruststorepass(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["truststorepass"] = v
}

func (p *LdapConfigParams) GetTruststorepass() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["truststorepass"].(string)
	return value, ok
}

// You should always use this function to get a new LdapConfigParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLdapConfigParams() *LdapConfigParams {
	p := &LdapConfigParams{}
	p.p = make(map[string]interface{})
	return p
}

// (Deprecated, use addLdapConfiguration) Configure the LDAP context for this site.
func (s *LDAPService) LdapConfig(p *LdapConfigParams) (*LdapConfigResponse, error) {
	resp, err := s.cs.newRequest("ldapConfig", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LdapConfigResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LdapConfigResponse struct {
	Binddn      string `json:"binddn"`
	Bindpass    string `json:"bindpass"`
	Hostname    string `json:"hostname"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Port        string `json:"port"`
	Queryfilter string `json:"queryfilter"`
	Searchbase  string `json:"searchbase"`
	Ssl         string `json:"ssl"`
}

type LdapCreateAccountParams struct {
	p map[string]interface{}
}

func (p *LdapCreateAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("accountdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("accountdetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *LdapCreateAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *LdapCreateAccountParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *LdapCreateAccountParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
}

func (p *LdapCreateAccountParams) GetAccountdetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountdetails"].(map[string]string)
	return value, ok
}

func (p *LdapCreateAccountParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *LdapCreateAccountParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *LdapCreateAccountParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *LdapCreateAccountParams) GetAccounttype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounttype"].(int)
	return value, ok
}

func (p *LdapCreateAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *LdapCreateAccountParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *LdapCreateAccountParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *LdapCreateAccountParams) GetNetworkdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdomain"].(string)
	return value, ok
}

func (p *LdapCreateAccountParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *LdapCreateAccountParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

func (p *LdapCreateAccountParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *LdapCreateAccountParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

func (p *LdapCreateAccountParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *LdapCreateAccountParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

func (p *LdapCreateAccountParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *LdapCreateAccountParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new LdapCreateAccountParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLdapCreateAccountParams(username string) *LdapCreateAccountParams {
	p := &LdapCreateAccountParams{}
	p.p = make(map[string]interface{})
	p.p["username"] = username
	return p
}

// Creates an account from an LDAP user
func (s *LDAPService) LdapCreateAccount(p *LdapCreateAccountParams) (*LdapCreateAccountResponse, error) {
	resp, err := s.cs.newRequest("ldapCreateAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LdapCreateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LdapCreateAccountResponse struct {
	Accountdetails            map[string]string               `json:"accountdetails"`
	Accounttype               int                             `json:"accounttype"`
	Cpuavailable              string                          `json:"cpuavailable"`
	Cpulimit                  string                          `json:"cpulimit"`
	Cputotal                  int64                           `json:"cputotal"`
	Created                   string                          `json:"created"`
	Defaultzoneid             string                          `json:"defaultzoneid"`
	Domain                    string                          `json:"domain"`
	Domainid                  string                          `json:"domainid"`
	Domainpath                string                          `json:"domainpath"`
	Groups                    []string                        `json:"groups"`
	Icon                      interface{}                     `json:"icon"`
	Id                        string                          `json:"id"`
	Ipavailable               string                          `json:"ipavailable"`
	Iplimit                   string                          `json:"iplimit"`
	Iptotal                   int64                           `json:"iptotal"`
	Iscleanuprequired         bool                            `json:"iscleanuprequired"`
	Isdefault                 bool                            `json:"isdefault"`
	JobID                     string                          `json:"jobid"`
	Jobstatus                 int                             `json:"jobstatus"`
	Memoryavailable           string                          `json:"memoryavailable"`
	Memorylimit               string                          `json:"memorylimit"`
	Memorytotal               int64                           `json:"memorytotal"`
	Name                      string                          `json:"name"`
	Networkavailable          string                          `json:"networkavailable"`
	Networkdomain             string                          `json:"networkdomain"`
	Networklimit              string                          `json:"networklimit"`
	Networktotal              int64                           `json:"networktotal"`
	Primarystorageavailable   string                          `json:"primarystorageavailable"`
	Primarystoragelimit       string                          `json:"primarystoragelimit"`
	Primarystoragetotal       int64                           `json:"primarystoragetotal"`
	Projectavailable          string                          `json:"projectavailable"`
	Projectlimit              string                          `json:"projectlimit"`
	Projecttotal              int64                           `json:"projecttotal"`
	Receivedbytes             int64                           `json:"receivedbytes"`
	Roleid                    string                          `json:"roleid"`
	Rolename                  string                          `json:"rolename"`
	Roletype                  string                          `json:"roletype"`
	Secondarystorageavailable string                          `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                          `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                         `json:"secondarystoragetotal"`
	Sentbytes                 int64                           `json:"sentbytes"`
	Snapshotavailable         string                          `json:"snapshotavailable"`
	Snapshotlimit             string                          `json:"snapshotlimit"`
	Snapshottotal             int64                           `json:"snapshottotal"`
	State                     string                          `json:"state"`
	Templateavailable         string                          `json:"templateavailable"`
	Templatelimit             string                          `json:"templatelimit"`
	Templatetotal             int64                           `json:"templatetotal"`
	User                      []LdapCreateAccountResponseUser `json:"user"`
	Vmavailable               string                          `json:"vmavailable"`
	Vmlimit                   string                          `json:"vmlimit"`
	Vmrunning                 int                             `json:"vmrunning"`
	Vmstopped                 int                             `json:"vmstopped"`
	Vmtotal                   int64                           `json:"vmtotal"`
	Volumeavailable           string                          `json:"volumeavailable"`
	Volumelimit               string                          `json:"volumelimit"`
	Volumetotal               int64                           `json:"volumetotal"`
	Vpcavailable              string                          `json:"vpcavailable"`
	Vpclimit                  string                          `json:"vpclimit"`
	Vpctotal                  int64                           `json:"vpctotal"`
}

type LdapCreateAccountResponseUser struct {
	Account             string      `json:"account"`
	Accountid           string      `json:"accountid"`
	Accounttype         int         `json:"accounttype"`
	Apikey              string      `json:"apikey"`
	Created             string      `json:"created"`
	Domain              string      `json:"domain"`
	Domainid            string      `json:"domainid"`
	Email               string      `json:"email"`
	Firstname           string      `json:"firstname"`
	Icon                interface{} `json:"icon"`
	Id                  string      `json:"id"`
	Is2faenabled        bool        `json:"is2faenabled"`
	Is2famandated       bool        `json:"is2famandated"`
	Iscallerchilddomain bool        `json:"iscallerchilddomain"`
	Isdefault           bool        `json:"isdefault"`
	Lastname            string      `json:"lastname"`
	Roleid              string      `json:"roleid"`
	Rolename            string      `json:"rolename"`
	Roletype            string      `json:"roletype"`
	Secretkey           string      `json:"secretkey"`
	State               string      `json:"state"`
	Timezone            string      `json:"timezone"`
	Username            string      `json:"username"`
	Usersource          string      `json:"usersource"`
}

type LdapRemoveParams struct {
	p map[string]interface{}
}

func (p *LdapRemoveParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new LdapRemoveParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLdapRemoveParams() *LdapRemoveParams {
	p := &LdapRemoveParams{}
	p.p = make(map[string]interface{})
	return p
}

// (Deprecated , use deleteLdapConfiguration) Remove the LDAP context for this site.
func (s *LDAPService) LdapRemove(p *LdapRemoveParams) (*LdapRemoveResponse, error) {
	resp, err := s.cs.newRequest("ldapRemove", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LdapRemoveResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LdapRemoveResponse struct {
	Binddn      string `json:"binddn"`
	Bindpass    string `json:"bindpass"`
	Hostname    string `json:"hostname"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Port        string `json:"port"`
	Queryfilter string `json:"queryfilter"`
	Searchbase  string `json:"searchbase"`
	Ssl         string `json:"ssl"`
}

type LinkDomainToLdapParams struct {
	p map[string]interface{}
}

func (p *LinkDomainToLdapParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["admin"]; found {
		u.Set("admin", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["ldapdomain"]; found {
		u.Set("ldapdomain", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *LinkDomainToLdapParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *LinkDomainToLdapParams) GetAccounttype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounttype"].(int)
	return value, ok
}

func (p *LinkDomainToLdapParams) SetAdmin(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["admin"] = v
}

func (p *LinkDomainToLdapParams) GetAdmin() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["admin"].(string)
	return value, ok
}

func (p *LinkDomainToLdapParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *LinkDomainToLdapParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *LinkDomainToLdapParams) SetLdapdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ldapdomain"] = v
}

func (p *LinkDomainToLdapParams) GetLdapdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ldapdomain"].(string)
	return value, ok
}

func (p *LinkDomainToLdapParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *LinkDomainToLdapParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *LinkDomainToLdapParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *LinkDomainToLdapParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new LinkDomainToLdapParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLinkDomainToLdapParams(accounttype int, domainid string, lDAPType string) *LinkDomainToLdapParams {
	p := &LinkDomainToLdapParams{}
	p.p = make(map[string]interface{})
	p.p["accounttype"] = accounttype
	p.p["domainid"] = domainid
	p.p["type"] = lDAPType
	return p
}

// link an existing cloudstack domain to group or OU in ldap
func (s *LDAPService) LinkDomainToLdap(p *LinkDomainToLdapParams) (*LinkDomainToLdapResponse, error) {
	resp, err := s.cs.newRequest("linkDomainToLdap", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LinkDomainToLdapResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LinkDomainToLdapResponse struct {
	Accountid   string `json:"accountid"`
	Accounttype int    `json:"accounttype"`
	Domainid    string `json:"domainid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Ldapdomain  string `json:"ldapdomain"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type ListLdapConfigurationsParams struct {
	p map[string]interface{}
}

func (p *ListLdapConfigurationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
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
	if v, found := p.p["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (p *ListLdapConfigurationsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListLdapConfigurationsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListLdapConfigurationsParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *ListLdapConfigurationsParams) GetHostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostname"].(string)
	return value, ok
}

func (p *ListLdapConfigurationsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLdapConfigurationsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListLdapConfigurationsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListLdapConfigurationsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListLdapConfigurationsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLdapConfigurationsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListLdapConfigurationsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLdapConfigurationsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListLdapConfigurationsParams) SetPort(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["port"] = v
}

func (p *ListLdapConfigurationsParams) GetPort() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["port"].(int)
	return value, ok
}

// You should always use this function to get a new ListLdapConfigurationsParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewListLdapConfigurationsParams() *ListLdapConfigurationsParams {
	p := &ListLdapConfigurationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all LDAP configurations
func (s *LDAPService) ListLdapConfigurations(p *ListLdapConfigurationsParams) (*ListLdapConfigurationsResponse, error) {
	resp, err := s.cs.newRequest("listLdapConfigurations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLdapConfigurationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLdapConfigurationsResponse struct {
	Count              int                  `json:"count"`
	LdapConfigurations []*LdapConfiguration `json:"ldapconfiguration"`
}

type LdapConfiguration struct {
	Domainid  string `json:"domainid"`
	Hostname  string `json:"hostname"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Port      int    `json:"port"`
}

type ListLdapUsersParams struct {
	p map[string]interface{}
}

func (p *ListLdapUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listtype"]; found {
		u.Set("listtype", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["userfilter"]; found {
		u.Set("userfilter", v.(string))
	}
	return u
}

func (p *ListLdapUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListLdapUsersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListLdapUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLdapUsersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListLdapUsersParams) SetListtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listtype"] = v
}

func (p *ListLdapUsersParams) GetListtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listtype"].(string)
	return value, ok
}

func (p *ListLdapUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLdapUsersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListLdapUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLdapUsersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListLdapUsersParams) SetUserfilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userfilter"] = v
}

func (p *ListLdapUsersParams) GetUserfilter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userfilter"].(string)
	return value, ok
}

// You should always use this function to get a new ListLdapUsersParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewListLdapUsersParams() *ListLdapUsersParams {
	p := &ListLdapUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists LDAP Users according to the specifications from the user request.
func (s *LDAPService) ListLdapUsers(p *ListLdapUsersParams) (*ListLdapUsersResponse, error) {
	resp, err := s.cs.newRequest("listLdapUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLdapUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLdapUsersResponse struct {
	Count     int         `json:"count"`
	LdapUsers []*LdapUser `json:"ldapuser"`
}

type LdapUser struct {
	Conflictingusersource string `json:"conflictingusersource"`
	Domain                string `json:"domain"`
	Email                 string `json:"email"`
	Firstname             string `json:"firstname"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Lastname              string `json:"lastname"`
	Principal             string `json:"principal"`
	Username              string `json:"username"`
}

type SearchLdapParams struct {
	p map[string]interface{}
}

func (p *SearchLdapParams) toURLValues() url.Values {
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
	if v, found := p.p["query"]; found {
		u.Set("query", v.(string))
	}
	return u
}

func (p *SearchLdapParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *SearchLdapParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *SearchLdapParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *SearchLdapParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *SearchLdapParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *SearchLdapParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *SearchLdapParams) SetQuery(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["query"] = v
}

func (p *SearchLdapParams) GetQuery() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["query"].(string)
	return value, ok
}

// You should always use this function to get a new SearchLdapParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewSearchLdapParams(query string) *SearchLdapParams {
	p := &SearchLdapParams{}
	p.p = make(map[string]interface{})
	p.p["query"] = query
	return p
}

// Searches LDAP based on the username attribute
func (s *LDAPService) SearchLdap(p *SearchLdapParams) (*SearchLdapResponse, error) {
	resp, err := s.cs.newRequest("searchLdap", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r SearchLdapResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type SearchLdapResponse struct {
	Conflictingusersource string `json:"conflictingusersource"`
	Domain                string `json:"domain"`
	Email                 string `json:"email"`
	Firstname             string `json:"firstname"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Lastname              string `json:"lastname"`
	Principal             string `json:"principal"`
	Username              string `json:"username"`
}
