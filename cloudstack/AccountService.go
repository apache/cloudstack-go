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

type AccountServiceIface interface {
	CreateAccount(p *CreateAccountParams) (*CreateAccountResponse, error)
	NewCreateAccountParams(email string, firstname string, lastname string, password string, username string) *CreateAccountParams
	DeleteAccount(p *DeleteAccountParams) (*DeleteAccountResponse, error)
	NewDeleteAccountParams(id string) *DeleteAccountParams
	DisableAccount(p *DisableAccountParams) (*DisableAccountResponse, error)
	NewDisableAccountParams(lock bool) *DisableAccountParams
	EnableAccount(p *EnableAccountParams) (*EnableAccountResponse, error)
	NewEnableAccountParams() *EnableAccountParams
	IsAccountAllowedToCreateOfferingsWithTags(p *IsAccountAllowedToCreateOfferingsWithTagsParams) (*IsAccountAllowedToCreateOfferingsWithTagsResponse, error)
	NewIsAccountAllowedToCreateOfferingsWithTagsParams(id string) *IsAccountAllowedToCreateOfferingsWithTagsParams
	LinkAccountToLdap(p *LinkAccountToLdapParams) (*LinkAccountToLdapResponse, error)
	NewLinkAccountToLdapParams(account string, domainid string, ldapdomain string) *LinkAccountToLdapParams
	ListAccounts(p *ListAccountsParams) (*ListAccountsResponse, error)
	NewListAccountsParams() *ListAccountsParams
	GetAccountID(name string, opts ...OptionFunc) (string, int, error)
	GetAccountByName(name string, opts ...OptionFunc) (*Account, int, error)
	GetAccountByID(id string, opts ...OptionFunc) (*Account, int, error)
	ListProjectAccounts(p *ListProjectAccountsParams) (*ListProjectAccountsResponse, error)
	NewListProjectAccountsParams(projectid string) *ListProjectAccountsParams
	GetProjectAccountID(keyword string, projectid string, opts ...OptionFunc) (string, int, error)
	LockAccount(p *LockAccountParams) (*LockAccountResponse, error)
	NewLockAccountParams(account string, domainid string) *LockAccountParams
	MarkDefaultZoneForAccount(p *MarkDefaultZoneForAccountParams) (*MarkDefaultZoneForAccountResponse, error)
	NewMarkDefaultZoneForAccountParams(account string, domainid string, zoneid string) *MarkDefaultZoneForAccountParams
	UpdateAccount(p *UpdateAccountParams) (*UpdateAccountResponse, error)
	NewUpdateAccountParams() *UpdateAccountParams
}

type CreateAccountParams struct {
	p map[string]interface{}
}

func (p *CreateAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountdetails"]; found {
		m := v.(map[string]string)
		for _, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("accountdetails[0].%s", k), m[k])
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
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := p.p["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
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

func (p *CreateAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateAccountParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateAccountParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
}

func (p *CreateAccountParams) ResetAccountdetails() {
	if p.p != nil && p.p["accountdetails"] != nil {
		delete(p.p, "accountdetails")
	}
}

func (p *CreateAccountParams) GetAccountdetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountdetails"].(map[string]string)
	return value, ok
}

func (p *CreateAccountParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *CreateAccountParams) ResetAccountid() {
	if p.p != nil && p.p["accountid"] != nil {
		delete(p.p, "accountid")
	}
}

func (p *CreateAccountParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *CreateAccountParams) ResetAccounttype() {
	if p.p != nil && p.p["accounttype"] != nil {
		delete(p.p, "accounttype")
	}
}

func (p *CreateAccountParams) GetAccounttype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounttype"].(int)
	return value, ok
}

func (p *CreateAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateAccountParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateAccountParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
}

func (p *CreateAccountParams) ResetEmail() {
	if p.p != nil && p.p["email"] != nil {
		delete(p.p, "email")
	}
}

func (p *CreateAccountParams) GetEmail() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["email"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetFirstname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firstname"] = v
}

func (p *CreateAccountParams) ResetFirstname() {
	if p.p != nil && p.p["firstname"] != nil {
		delete(p.p, "firstname")
	}
}

func (p *CreateAccountParams) GetFirstname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["firstname"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetLastname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lastname"] = v
}

func (p *CreateAccountParams) ResetLastname() {
	if p.p != nil && p.p["lastname"] != nil {
		delete(p.p, "lastname")
	}
}

func (p *CreateAccountParams) GetLastname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lastname"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *CreateAccountParams) ResetNetworkdomain() {
	if p.p != nil && p.p["networkdomain"] != nil {
		delete(p.p, "networkdomain")
	}
}

func (p *CreateAccountParams) GetNetworkdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdomain"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *CreateAccountParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *CreateAccountParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *CreateAccountParams) ResetRoleid() {
	if p.p != nil && p.p["roleid"] != nil {
		delete(p.p, "roleid")
	}
}

func (p *CreateAccountParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *CreateAccountParams) ResetTimezone() {
	if p.p != nil && p.p["timezone"] != nil {
		delete(p.p, "timezone")
	}
}

func (p *CreateAccountParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *CreateAccountParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *CreateAccountParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

func (p *CreateAccountParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *CreateAccountParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *CreateAccountParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new CreateAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewCreateAccountParams(email string, firstname string, lastname string, password string, username string) *CreateAccountParams {
	p := &CreateAccountParams{}
	p.p = make(map[string]interface{})
	p.p["email"] = email
	p.p["firstname"] = firstname
	p.p["lastname"] = lastname
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Creates an account
func (s *AccountService) CreateAccount(p *CreateAccountParams) (*CreateAccountResponse, error) {
	resp, err := s.cs.newPostRequest("createAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateAccountResponse struct {
	Accountdetails            map[string]string           `json:"accountdetails"`
	Accounttype               int                         `json:"accounttype"`
	Apikeyaccess              string                      `json:"apikeyaccess"`
	Cpuavailable              string                      `json:"cpuavailable"`
	Cpulimit                  string                      `json:"cpulimit"`
	Cputotal                  int64                       `json:"cputotal"`
	Created                   string                      `json:"created"`
	Defaultzoneid             string                      `json:"defaultzoneid"`
	Domain                    string                      `json:"domain"`
	Domainid                  string                      `json:"domainid"`
	Domainpath                string                      `json:"domainpath"`
	Groups                    []string                    `json:"groups"`
	Icon                      interface{}                 `json:"icon"`
	Id                        string                      `json:"id"`
	Ipavailable               string                      `json:"ipavailable"`
	Iplimit                   string                      `json:"iplimit"`
	Iptotal                   int64                       `json:"iptotal"`
	Iscleanuprequired         bool                        `json:"iscleanuprequired"`
	Isdefault                 bool                        `json:"isdefault"`
	JobID                     string                      `json:"jobid"`
	Jobstatus                 int                         `json:"jobstatus"`
	Memoryavailable           string                      `json:"memoryavailable"`
	Memorylimit               string                      `json:"memorylimit"`
	Memorytotal               int64                       `json:"memorytotal"`
	Name                      string                      `json:"name"`
	Networkavailable          string                      `json:"networkavailable"`
	Networkdomain             string                      `json:"networkdomain"`
	Networklimit              string                      `json:"networklimit"`
	Networktotal              int64                       `json:"networktotal"`
	Primarystorageavailable   string                      `json:"primarystorageavailable"`
	Primarystoragelimit       string                      `json:"primarystoragelimit"`
	Primarystoragetotal       int64                       `json:"primarystoragetotal"`
	Projectavailable          string                      `json:"projectavailable"`
	Projectlimit              string                      `json:"projectlimit"`
	Projecttotal              int64                       `json:"projecttotal"`
	Receivedbytes             int64                       `json:"receivedbytes"`
	Roleid                    string                      `json:"roleid"`
	Rolename                  string                      `json:"rolename"`
	Roletype                  string                      `json:"roletype"`
	Secondarystorageavailable string                      `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                      `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                     `json:"secondarystoragetotal"`
	Sentbytes                 int64                       `json:"sentbytes"`
	Snapshotavailable         string                      `json:"snapshotavailable"`
	Snapshotlimit             string                      `json:"snapshotlimit"`
	Snapshottotal             int64                       `json:"snapshottotal"`
	State                     string                      `json:"state"`
	Taggedresources           []string                    `json:"taggedresources"`
	Templateavailable         string                      `json:"templateavailable"`
	Templatelimit             string                      `json:"templatelimit"`
	Templatetotal             int64                       `json:"templatetotal"`
	User                      []CreateAccountResponseUser `json:"user"`
	Vmavailable               string                      `json:"vmavailable"`
	Vmlimit                   string                      `json:"vmlimit"`
	Vmrunning                 int                         `json:"vmrunning"`
	Vmstopped                 int                         `json:"vmstopped"`
	Vmtotal                   int64                       `json:"vmtotal"`
	Volumeavailable           string                      `json:"volumeavailable"`
	Volumelimit               string                      `json:"volumelimit"`
	Volumetotal               int64                       `json:"volumetotal"`
	Vpcavailable              string                      `json:"vpcavailable"`
	Vpclimit                  string                      `json:"vpclimit"`
	Vpctotal                  int64                       `json:"vpctotal"`
}

type CreateAccountResponseUser struct {
	Account             string      `json:"account"`
	Accountid           string      `json:"accountid"`
	Accounttype         int         `json:"accounttype"`
	Apikey              string      `json:"apikey"`
	Apikeyaccess        string      `json:"apikeyaccess"`
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

type DeleteAccountParams struct {
	p map[string]interface{}
}

func (p *DeleteAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteAccountParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteAccountParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteAccountParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewDeleteAccountParams(id string) *DeleteAccountParams {
	p := &DeleteAccountParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a account, and all users associated with this account
func (s *AccountService) DeleteAccount(p *DeleteAccountParams) (*DeleteAccountResponse, error) {
	resp, err := s.cs.newPostRequest("deleteAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAccountResponse
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

type DeleteAccountResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DisableAccountParams struct {
	p map[string]interface{}
}

func (p *DisableAccountParams) toURLValues() url.Values {
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
	if v, found := p.p["lock"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lock", vv)
	}
	return u
}

func (p *DisableAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DisableAccountParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DisableAccountParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DisableAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DisableAccountParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DisableAccountParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DisableAccountParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DisableAccountParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DisableAccountParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DisableAccountParams) SetLock(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lock"] = v
}

func (p *DisableAccountParams) ResetLock() {
	if p.p != nil && p.p["lock"] != nil {
		delete(p.p, "lock")
	}
}

func (p *DisableAccountParams) GetLock() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lock"].(bool)
	return value, ok
}

// You should always use this function to get a new DisableAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewDisableAccountParams(lock bool) *DisableAccountParams {
	p := &DisableAccountParams{}
	p.p = make(map[string]interface{})
	p.p["lock"] = lock
	return p
}

// Disables an account
func (s *AccountService) DisableAccount(p *DisableAccountParams) (*DisableAccountResponse, error) {
	resp, err := s.cs.newPostRequest("disableAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableAccountResponse
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

type DisableAccountResponse struct {
	Accountdetails            map[string]string            `json:"accountdetails"`
	Accounttype               int                          `json:"accounttype"`
	Apikeyaccess              string                       `json:"apikeyaccess"`
	Cpuavailable              string                       `json:"cpuavailable"`
	Cpulimit                  string                       `json:"cpulimit"`
	Cputotal                  int64                        `json:"cputotal"`
	Created                   string                       `json:"created"`
	Defaultzoneid             string                       `json:"defaultzoneid"`
	Domain                    string                       `json:"domain"`
	Domainid                  string                       `json:"domainid"`
	Domainpath                string                       `json:"domainpath"`
	Groups                    []string                     `json:"groups"`
	Icon                      interface{}                  `json:"icon"`
	Id                        string                       `json:"id"`
	Ipavailable               string                       `json:"ipavailable"`
	Iplimit                   string                       `json:"iplimit"`
	Iptotal                   int64                        `json:"iptotal"`
	Iscleanuprequired         bool                         `json:"iscleanuprequired"`
	Isdefault                 bool                         `json:"isdefault"`
	JobID                     string                       `json:"jobid"`
	Jobstatus                 int                          `json:"jobstatus"`
	Memoryavailable           string                       `json:"memoryavailable"`
	Memorylimit               string                       `json:"memorylimit"`
	Memorytotal               int64                        `json:"memorytotal"`
	Name                      string                       `json:"name"`
	Networkavailable          string                       `json:"networkavailable"`
	Networkdomain             string                       `json:"networkdomain"`
	Networklimit              string                       `json:"networklimit"`
	Networktotal              int64                        `json:"networktotal"`
	Primarystorageavailable   string                       `json:"primarystorageavailable"`
	Primarystoragelimit       string                       `json:"primarystoragelimit"`
	Primarystoragetotal       int64                        `json:"primarystoragetotal"`
	Projectavailable          string                       `json:"projectavailable"`
	Projectlimit              string                       `json:"projectlimit"`
	Projecttotal              int64                        `json:"projecttotal"`
	Receivedbytes             int64                        `json:"receivedbytes"`
	Roleid                    string                       `json:"roleid"`
	Rolename                  string                       `json:"rolename"`
	Roletype                  string                       `json:"roletype"`
	Secondarystorageavailable string                       `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                       `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                      `json:"secondarystoragetotal"`
	Sentbytes                 int64                        `json:"sentbytes"`
	Snapshotavailable         string                       `json:"snapshotavailable"`
	Snapshotlimit             string                       `json:"snapshotlimit"`
	Snapshottotal             int64                        `json:"snapshottotal"`
	State                     string                       `json:"state"`
	Taggedresources           []string                     `json:"taggedresources"`
	Templateavailable         string                       `json:"templateavailable"`
	Templatelimit             string                       `json:"templatelimit"`
	Templatetotal             int64                        `json:"templatetotal"`
	User                      []DisableAccountResponseUser `json:"user"`
	Vmavailable               string                       `json:"vmavailable"`
	Vmlimit                   string                       `json:"vmlimit"`
	Vmrunning                 int                          `json:"vmrunning"`
	Vmstopped                 int                          `json:"vmstopped"`
	Vmtotal                   int64                        `json:"vmtotal"`
	Volumeavailable           string                       `json:"volumeavailable"`
	Volumelimit               string                       `json:"volumelimit"`
	Volumetotal               int64                        `json:"volumetotal"`
	Vpcavailable              string                       `json:"vpcavailable"`
	Vpclimit                  string                       `json:"vpclimit"`
	Vpctotal                  int64                        `json:"vpctotal"`
}

type DisableAccountResponseUser struct {
	Account             string      `json:"account"`
	Accountid           string      `json:"accountid"`
	Accounttype         int         `json:"accounttype"`
	Apikey              string      `json:"apikey"`
	Apikeyaccess        string      `json:"apikeyaccess"`
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

type EnableAccountParams struct {
	p map[string]interface{}
}

func (p *EnableAccountParams) toURLValues() url.Values {
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
	return u
}

func (p *EnableAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *EnableAccountParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *EnableAccountParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *EnableAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *EnableAccountParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *EnableAccountParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *EnableAccountParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *EnableAccountParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *EnableAccountParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new EnableAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewEnableAccountParams() *EnableAccountParams {
	p := &EnableAccountParams{}
	p.p = make(map[string]interface{})
	return p
}

// Enables an account
func (s *AccountService) EnableAccount(p *EnableAccountParams) (*EnableAccountResponse, error) {
	resp, err := s.cs.newPostRequest("enableAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type EnableAccountResponse struct {
	Accountdetails            map[string]string           `json:"accountdetails"`
	Accounttype               int                         `json:"accounttype"`
	Apikeyaccess              string                      `json:"apikeyaccess"`
	Cpuavailable              string                      `json:"cpuavailable"`
	Cpulimit                  string                      `json:"cpulimit"`
	Cputotal                  int64                       `json:"cputotal"`
	Created                   string                      `json:"created"`
	Defaultzoneid             string                      `json:"defaultzoneid"`
	Domain                    string                      `json:"domain"`
	Domainid                  string                      `json:"domainid"`
	Domainpath                string                      `json:"domainpath"`
	Groups                    []string                    `json:"groups"`
	Icon                      interface{}                 `json:"icon"`
	Id                        string                      `json:"id"`
	Ipavailable               string                      `json:"ipavailable"`
	Iplimit                   string                      `json:"iplimit"`
	Iptotal                   int64                       `json:"iptotal"`
	Iscleanuprequired         bool                        `json:"iscleanuprequired"`
	Isdefault                 bool                        `json:"isdefault"`
	JobID                     string                      `json:"jobid"`
	Jobstatus                 int                         `json:"jobstatus"`
	Memoryavailable           string                      `json:"memoryavailable"`
	Memorylimit               string                      `json:"memorylimit"`
	Memorytotal               int64                       `json:"memorytotal"`
	Name                      string                      `json:"name"`
	Networkavailable          string                      `json:"networkavailable"`
	Networkdomain             string                      `json:"networkdomain"`
	Networklimit              string                      `json:"networklimit"`
	Networktotal              int64                       `json:"networktotal"`
	Primarystorageavailable   string                      `json:"primarystorageavailable"`
	Primarystoragelimit       string                      `json:"primarystoragelimit"`
	Primarystoragetotal       int64                       `json:"primarystoragetotal"`
	Projectavailable          string                      `json:"projectavailable"`
	Projectlimit              string                      `json:"projectlimit"`
	Projecttotal              int64                       `json:"projecttotal"`
	Receivedbytes             int64                       `json:"receivedbytes"`
	Roleid                    string                      `json:"roleid"`
	Rolename                  string                      `json:"rolename"`
	Roletype                  string                      `json:"roletype"`
	Secondarystorageavailable string                      `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                      `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                     `json:"secondarystoragetotal"`
	Sentbytes                 int64                       `json:"sentbytes"`
	Snapshotavailable         string                      `json:"snapshotavailable"`
	Snapshotlimit             string                      `json:"snapshotlimit"`
	Snapshottotal             int64                       `json:"snapshottotal"`
	State                     string                      `json:"state"`
	Taggedresources           []string                    `json:"taggedresources"`
	Templateavailable         string                      `json:"templateavailable"`
	Templatelimit             string                      `json:"templatelimit"`
	Templatetotal             int64                       `json:"templatetotal"`
	User                      []EnableAccountResponseUser `json:"user"`
	Vmavailable               string                      `json:"vmavailable"`
	Vmlimit                   string                      `json:"vmlimit"`
	Vmrunning                 int                         `json:"vmrunning"`
	Vmstopped                 int                         `json:"vmstopped"`
	Vmtotal                   int64                       `json:"vmtotal"`
	Volumeavailable           string                      `json:"volumeavailable"`
	Volumelimit               string                      `json:"volumelimit"`
	Volumetotal               int64                       `json:"volumetotal"`
	Vpcavailable              string                      `json:"vpcavailable"`
	Vpclimit                  string                      `json:"vpclimit"`
	Vpctotal                  int64                       `json:"vpctotal"`
}

type EnableAccountResponseUser struct {
	Account             string      `json:"account"`
	Accountid           string      `json:"accountid"`
	Accounttype         int         `json:"accounttype"`
	Apikey              string      `json:"apikey"`
	Apikeyaccess        string      `json:"apikeyaccess"`
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

type IsAccountAllowedToCreateOfferingsWithTagsParams struct {
	p map[string]interface{}
}

func (p *IsAccountAllowedToCreateOfferingsWithTagsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *IsAccountAllowedToCreateOfferingsWithTagsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *IsAccountAllowedToCreateOfferingsWithTagsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *IsAccountAllowedToCreateOfferingsWithTagsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new IsAccountAllowedToCreateOfferingsWithTagsParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewIsAccountAllowedToCreateOfferingsWithTagsParams(id string) *IsAccountAllowedToCreateOfferingsWithTagsParams {
	p := &IsAccountAllowedToCreateOfferingsWithTagsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Return true if the specified account is allowed to create offerings with tags.
func (s *AccountService) IsAccountAllowedToCreateOfferingsWithTags(p *IsAccountAllowedToCreateOfferingsWithTagsParams) (*IsAccountAllowedToCreateOfferingsWithTagsResponse, error) {
	resp, err := s.cs.newRequest("isAccountAllowedToCreateOfferingsWithTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r IsAccountAllowedToCreateOfferingsWithTagsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type IsAccountAllowedToCreateOfferingsWithTagsResponse struct {
	Isallowed bool   `json:"isallowed"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
}

type LinkAccountToLdapParams struct {
	p map[string]interface{}
}

func (p *LinkAccountToLdapParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
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
	if v, found := p.p["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *LinkAccountToLdapParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *LinkAccountToLdapParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *LinkAccountToLdapParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *LinkAccountToLdapParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *LinkAccountToLdapParams) ResetAccounttype() {
	if p.p != nil && p.p["accounttype"] != nil {
		delete(p.p, "accounttype")
	}
}

func (p *LinkAccountToLdapParams) GetAccounttype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounttype"].(int)
	return value, ok
}

func (p *LinkAccountToLdapParams) SetAdmin(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["admin"] = v
}

func (p *LinkAccountToLdapParams) ResetAdmin() {
	if p.p != nil && p.p["admin"] != nil {
		delete(p.p, "admin")
	}
}

func (p *LinkAccountToLdapParams) GetAdmin() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["admin"].(string)
	return value, ok
}

func (p *LinkAccountToLdapParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *LinkAccountToLdapParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *LinkAccountToLdapParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *LinkAccountToLdapParams) SetLdapdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ldapdomain"] = v
}

func (p *LinkAccountToLdapParams) ResetLdapdomain() {
	if p.p != nil && p.p["ldapdomain"] != nil {
		delete(p.p, "ldapdomain")
	}
}

func (p *LinkAccountToLdapParams) GetLdapdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ldapdomain"].(string)
	return value, ok
}

func (p *LinkAccountToLdapParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *LinkAccountToLdapParams) ResetRoleid() {
	if p.p != nil && p.p["roleid"] != nil {
		delete(p.p, "roleid")
	}
}

func (p *LinkAccountToLdapParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

func (p *LinkAccountToLdapParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *LinkAccountToLdapParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *LinkAccountToLdapParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new LinkAccountToLdapParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewLinkAccountToLdapParams(account string, domainid string, ldapdomain string) *LinkAccountToLdapParams {
	p := &LinkAccountToLdapParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	p.p["ldapdomain"] = ldapdomain
	return p
}

// link a cloudstack account to a group or OU in ldap
func (s *AccountService) LinkAccountToLdap(p *LinkAccountToLdapParams) (*LinkAccountToLdapResponse, error) {
	resp, err := s.cs.newPostRequest("linkAccountToLdap", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LinkAccountToLdapResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LinkAccountToLdapResponse struct {
	Accountid   string `json:"accountid"`
	Accounttype int    `json:"accounttype"`
	Domainid    string `json:"domainid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Ldapdomain  string `json:"ldapdomain"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type ListAccountsParams struct {
	p map[string]interface{}
}

func (p *ListAccountsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := p.p["apikeyaccess"]; found {
		u.Set("apikeyaccess", v.(string))
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["iscleanuprequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("iscleanuprequired", vv)
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
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["tag"]; found {
		u.Set("tag", v.(string))
	}
	return u
}

func (p *ListAccountsParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *ListAccountsParams) ResetAccounttype() {
	if p.p != nil && p.p["accounttype"] != nil {
		delete(p.p, "accounttype")
	}
}

func (p *ListAccountsParams) GetAccounttype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounttype"].(int)
	return value, ok
}

func (p *ListAccountsParams) SetApikeyaccess(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["apikeyaccess"] = v
}

func (p *ListAccountsParams) ResetApikeyaccess() {
	if p.p != nil && p.p["apikeyaccess"] != nil {
		delete(p.p, "apikeyaccess")
	}
}

func (p *ListAccountsParams) GetApikeyaccess() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["apikeyaccess"].(string)
	return value, ok
}

func (p *ListAccountsParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListAccountsParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListAccountsParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListAccountsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListAccountsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListAccountsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListAccountsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListAccountsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListAccountsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListAccountsParams) SetIscleanuprequired(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iscleanuprequired"] = v
}

func (p *ListAccountsParams) ResetIscleanuprequired() {
	if p.p != nil && p.p["iscleanuprequired"] != nil {
		delete(p.p, "iscleanuprequired")
	}
}

func (p *ListAccountsParams) GetIscleanuprequired() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iscleanuprequired"].(bool)
	return value, ok
}

func (p *ListAccountsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListAccountsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListAccountsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListAccountsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAccountsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListAccountsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListAccountsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListAccountsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListAccountsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListAccountsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListAccountsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListAccountsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListAccountsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAccountsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListAccountsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListAccountsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAccountsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListAccountsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListAccountsParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListAccountsParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListAccountsParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListAccountsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListAccountsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListAccountsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListAccountsParams) SetTag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tag"] = v
}

func (p *ListAccountsParams) ResetTag() {
	if p.p != nil && p.p["tag"] != nil {
		delete(p.p, "tag")
	}
}

func (p *ListAccountsParams) GetTag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tag"].(string)
	return value, ok
}

// You should always use this function to get a new ListAccountsParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewListAccountsParams() *ListAccountsParams {
	p := &ListAccountsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListAccountsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListAccounts(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Accounts[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Accounts {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountByName(name string, opts ...OptionFunc) (*Account, int, error) {
	id, count, err := s.GetAccountID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetAccountByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountByID(id string, opts ...OptionFunc) (*Account, int, error) {
	p := &ListAccountsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAccounts(p)
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
		return l.Accounts[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Account UUID: %s!", id)
}

// Lists accounts and provides detailed account information for listed accounts
func (s *AccountService) ListAccounts(p *ListAccountsParams) (*ListAccountsResponse, error) {
	resp, err := s.cs.newRequest("listAccounts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAccountsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAccountsResponse struct {
	Count    int        `json:"count"`
	Accounts []*Account `json:"account"`
}

type Account struct {
	Accountdetails            map[string]string `json:"accountdetails"`
	Accounttype               int               `json:"accounttype"`
	Apikeyaccess              string            `json:"apikeyaccess"`
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Defaultzoneid             string            `json:"defaultzoneid"`
	Domain                    string            `json:"domain"`
	Domainid                  string            `json:"domainid"`
	Domainpath                string            `json:"domainpath"`
	Groups                    []string          `json:"groups"`
	Icon                      interface{}       `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	Iscleanuprequired         bool              `json:"iscleanuprequired"`
	Isdefault                 bool              `json:"isdefault"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Receivedbytes             int64             `json:"receivedbytes"`
	Roleid                    string            `json:"roleid"`
	Rolename                  string            `json:"rolename"`
	Roletype                  string            `json:"roletype"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Sentbytes                 int64             `json:"sentbytes"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Taggedresources           []string          `json:"taggedresources"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	User                      []AccountUser     `json:"user"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmrunning                 int               `json:"vmrunning"`
	Vmstopped                 int               `json:"vmstopped"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type AccountUser struct {
	Account             string      `json:"account"`
	Accountid           string      `json:"accountid"`
	Accounttype         int         `json:"accounttype"`
	Apikey              string      `json:"apikey"`
	Apikeyaccess        string      `json:"apikeyaccess"`
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

type ListProjectAccountsParams struct {
	p map[string]interface{}
}

func (p *ListProjectAccountsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := p.p["role"]; found {
		u.Set("role", v.(string))
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (p *ListProjectAccountsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListProjectAccountsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListProjectAccountsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListProjectAccountsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListProjectAccountsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListProjectAccountsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListProjectAccountsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListProjectAccountsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListProjectAccountsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListProjectAccountsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListProjectAccountsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListProjectAccountsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListProjectAccountsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListProjectAccountsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListProjectAccountsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListProjectAccountsParams) SetProjectroleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectroleid"] = v
}

func (p *ListProjectAccountsParams) ResetProjectroleid() {
	if p.p != nil && p.p["projectroleid"] != nil {
		delete(p.p, "projectroleid")
	}
}

func (p *ListProjectAccountsParams) GetProjectroleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectroleid"].(string)
	return value, ok
}

func (p *ListProjectAccountsParams) SetRole(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["role"] = v
}

func (p *ListProjectAccountsParams) ResetRole() {
	if p.p != nil && p.p["role"] != nil {
		delete(p.p, "role")
	}
}

func (p *ListProjectAccountsParams) GetRole() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["role"].(string)
	return value, ok
}

func (p *ListProjectAccountsParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *ListProjectAccountsParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *ListProjectAccountsParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

// You should always use this function to get a new ListProjectAccountsParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewListProjectAccountsParams(projectid string) *ListProjectAccountsParams {
	p := &ListProjectAccountsParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetProjectAccountID(keyword string, projectid string, opts ...OptionFunc) (string, int, error) {
	p := &ListProjectAccountsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["projectid"] = projectid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListProjectAccounts(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.ProjectAccounts[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ProjectAccounts {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists project's accounts
func (s *AccountService) ListProjectAccounts(p *ListProjectAccountsParams) (*ListProjectAccountsResponse, error) {
	resp, err := s.cs.newRequest("listProjectAccounts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectAccountsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectAccountsResponse struct {
	Count           int               `json:"count"`
	ProjectAccounts []*ProjectAccount `json:"projectaccount"`
}

type ProjectAccount struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      interface{}         `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Taggedresources           []string            `json:"taggedresources"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type Tags struct {
	Account      string `json:"account"`
	Customer     string `json:"customer"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Domainpath   string `json:"domainpath"`
	Key          string `json:"key"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Resourceid   string `json:"resourceid"`
	Resourcetype string `json:"resourcetype"`
	Value        string `json:"value"`
}

type LockAccountParams struct {
	p map[string]interface{}
}

func (p *LockAccountParams) toURLValues() url.Values {
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
	return u
}

func (p *LockAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *LockAccountParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *LockAccountParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *LockAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *LockAccountParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *LockAccountParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

// You should always use this function to get a new LockAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewLockAccountParams(account string, domainid string) *LockAccountParams {
	p := &LockAccountParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	return p
}

// This deprecated function used to locks an account. Look for the API DisableAccount instead
func (s *AccountService) LockAccount(p *LockAccountParams) (*LockAccountResponse, error) {
	resp, err := s.cs.newPostRequest("lockAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LockAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LockAccountResponse struct {
	Accountdetails            map[string]string         `json:"accountdetails"`
	Accounttype               int                       `json:"accounttype"`
	Apikeyaccess              string                    `json:"apikeyaccess"`
	Cpuavailable              string                    `json:"cpuavailable"`
	Cpulimit                  string                    `json:"cpulimit"`
	Cputotal                  int64                     `json:"cputotal"`
	Created                   string                    `json:"created"`
	Defaultzoneid             string                    `json:"defaultzoneid"`
	Domain                    string                    `json:"domain"`
	Domainid                  string                    `json:"domainid"`
	Domainpath                string                    `json:"domainpath"`
	Groups                    []string                  `json:"groups"`
	Icon                      interface{}               `json:"icon"`
	Id                        string                    `json:"id"`
	Ipavailable               string                    `json:"ipavailable"`
	Iplimit                   string                    `json:"iplimit"`
	Iptotal                   int64                     `json:"iptotal"`
	Iscleanuprequired         bool                      `json:"iscleanuprequired"`
	Isdefault                 bool                      `json:"isdefault"`
	JobID                     string                    `json:"jobid"`
	Jobstatus                 int                       `json:"jobstatus"`
	Memoryavailable           string                    `json:"memoryavailable"`
	Memorylimit               string                    `json:"memorylimit"`
	Memorytotal               int64                     `json:"memorytotal"`
	Name                      string                    `json:"name"`
	Networkavailable          string                    `json:"networkavailable"`
	Networkdomain             string                    `json:"networkdomain"`
	Networklimit              string                    `json:"networklimit"`
	Networktotal              int64                     `json:"networktotal"`
	Primarystorageavailable   string                    `json:"primarystorageavailable"`
	Primarystoragelimit       string                    `json:"primarystoragelimit"`
	Primarystoragetotal       int64                     `json:"primarystoragetotal"`
	Projectavailable          string                    `json:"projectavailable"`
	Projectlimit              string                    `json:"projectlimit"`
	Projecttotal              int64                     `json:"projecttotal"`
	Receivedbytes             int64                     `json:"receivedbytes"`
	Roleid                    string                    `json:"roleid"`
	Rolename                  string                    `json:"rolename"`
	Roletype                  string                    `json:"roletype"`
	Secondarystorageavailable string                    `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                    `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                   `json:"secondarystoragetotal"`
	Sentbytes                 int64                     `json:"sentbytes"`
	Snapshotavailable         string                    `json:"snapshotavailable"`
	Snapshotlimit             string                    `json:"snapshotlimit"`
	Snapshottotal             int64                     `json:"snapshottotal"`
	State                     string                    `json:"state"`
	Taggedresources           []string                  `json:"taggedresources"`
	Templateavailable         string                    `json:"templateavailable"`
	Templatelimit             string                    `json:"templatelimit"`
	Templatetotal             int64                     `json:"templatetotal"`
	User                      []LockAccountResponseUser `json:"user"`
	Vmavailable               string                    `json:"vmavailable"`
	Vmlimit                   string                    `json:"vmlimit"`
	Vmrunning                 int                       `json:"vmrunning"`
	Vmstopped                 int                       `json:"vmstopped"`
	Vmtotal                   int64                     `json:"vmtotal"`
	Volumeavailable           string                    `json:"volumeavailable"`
	Volumelimit               string                    `json:"volumelimit"`
	Volumetotal               int64                     `json:"volumetotal"`
	Vpcavailable              string                    `json:"vpcavailable"`
	Vpclimit                  string                    `json:"vpclimit"`
	Vpctotal                  int64                     `json:"vpctotal"`
}

type LockAccountResponseUser struct {
	Account             string      `json:"account"`
	Accountid           string      `json:"accountid"`
	Accounttype         int         `json:"accounttype"`
	Apikey              string      `json:"apikey"`
	Apikeyaccess        string      `json:"apikeyaccess"`
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

type MarkDefaultZoneForAccountParams struct {
	p map[string]interface{}
}

func (p *MarkDefaultZoneForAccountParams) toURLValues() url.Values {
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *MarkDefaultZoneForAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *MarkDefaultZoneForAccountParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *MarkDefaultZoneForAccountParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *MarkDefaultZoneForAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *MarkDefaultZoneForAccountParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *MarkDefaultZoneForAccountParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *MarkDefaultZoneForAccountParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *MarkDefaultZoneForAccountParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *MarkDefaultZoneForAccountParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new MarkDefaultZoneForAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewMarkDefaultZoneForAccountParams(account string, domainid string, zoneid string) *MarkDefaultZoneForAccountParams {
	p := &MarkDefaultZoneForAccountParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	p.p["zoneid"] = zoneid
	return p
}

// Marks a default zone for this account
func (s *AccountService) MarkDefaultZoneForAccount(p *MarkDefaultZoneForAccountParams) (*MarkDefaultZoneForAccountResponse, error) {
	resp, err := s.cs.newPostRequest("markDefaultZoneForAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MarkDefaultZoneForAccountResponse
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

type MarkDefaultZoneForAccountResponse struct {
	Accountdetails            map[string]string                       `json:"accountdetails"`
	Accounttype               int                                     `json:"accounttype"`
	Apikeyaccess              string                                  `json:"apikeyaccess"`
	Cpuavailable              string                                  `json:"cpuavailable"`
	Cpulimit                  string                                  `json:"cpulimit"`
	Cputotal                  int64                                   `json:"cputotal"`
	Created                   string                                  `json:"created"`
	Defaultzoneid             string                                  `json:"defaultzoneid"`
	Domain                    string                                  `json:"domain"`
	Domainid                  string                                  `json:"domainid"`
	Domainpath                string                                  `json:"domainpath"`
	Groups                    []string                                `json:"groups"`
	Icon                      interface{}                             `json:"icon"`
	Id                        string                                  `json:"id"`
	Ipavailable               string                                  `json:"ipavailable"`
	Iplimit                   string                                  `json:"iplimit"`
	Iptotal                   int64                                   `json:"iptotal"`
	Iscleanuprequired         bool                                    `json:"iscleanuprequired"`
	Isdefault                 bool                                    `json:"isdefault"`
	JobID                     string                                  `json:"jobid"`
	Jobstatus                 int                                     `json:"jobstatus"`
	Memoryavailable           string                                  `json:"memoryavailable"`
	Memorylimit               string                                  `json:"memorylimit"`
	Memorytotal               int64                                   `json:"memorytotal"`
	Name                      string                                  `json:"name"`
	Networkavailable          string                                  `json:"networkavailable"`
	Networkdomain             string                                  `json:"networkdomain"`
	Networklimit              string                                  `json:"networklimit"`
	Networktotal              int64                                   `json:"networktotal"`
	Primarystorageavailable   string                                  `json:"primarystorageavailable"`
	Primarystoragelimit       string                                  `json:"primarystoragelimit"`
	Primarystoragetotal       int64                                   `json:"primarystoragetotal"`
	Projectavailable          string                                  `json:"projectavailable"`
	Projectlimit              string                                  `json:"projectlimit"`
	Projecttotal              int64                                   `json:"projecttotal"`
	Receivedbytes             int64                                   `json:"receivedbytes"`
	Roleid                    string                                  `json:"roleid"`
	Rolename                  string                                  `json:"rolename"`
	Roletype                  string                                  `json:"roletype"`
	Secondarystorageavailable string                                  `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                                  `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                                 `json:"secondarystoragetotal"`
	Sentbytes                 int64                                   `json:"sentbytes"`
	Snapshotavailable         string                                  `json:"snapshotavailable"`
	Snapshotlimit             string                                  `json:"snapshotlimit"`
	Snapshottotal             int64                                   `json:"snapshottotal"`
	State                     string                                  `json:"state"`
	Taggedresources           []string                                `json:"taggedresources"`
	Templateavailable         string                                  `json:"templateavailable"`
	Templatelimit             string                                  `json:"templatelimit"`
	Templatetotal             int64                                   `json:"templatetotal"`
	User                      []MarkDefaultZoneForAccountResponseUser `json:"user"`
	Vmavailable               string                                  `json:"vmavailable"`
	Vmlimit                   string                                  `json:"vmlimit"`
	Vmrunning                 int                                     `json:"vmrunning"`
	Vmstopped                 int                                     `json:"vmstopped"`
	Vmtotal                   int64                                   `json:"vmtotal"`
	Volumeavailable           string                                  `json:"volumeavailable"`
	Volumelimit               string                                  `json:"volumelimit"`
	Volumetotal               int64                                   `json:"volumetotal"`
	Vpcavailable              string                                  `json:"vpcavailable"`
	Vpclimit                  string                                  `json:"vpclimit"`
	Vpctotal                  int64                                   `json:"vpctotal"`
}

type MarkDefaultZoneForAccountResponseUser struct {
	Account             string      `json:"account"`
	Accountid           string      `json:"accountid"`
	Accounttype         int         `json:"accounttype"`
	Apikey              string      `json:"apikey"`
	Apikeyaccess        string      `json:"apikeyaccess"`
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

type UpdateAccountParams struct {
	p map[string]interface{}
}

func (p *UpdateAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountdetails"]; found {
		m := v.(map[string]string)
		for _, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("accountdetails[0].%s", k), m[k])
		}
	}
	if v, found := p.p["apikeyaccess"]; found {
		u.Set("apikeyaccess", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["newname"]; found {
		u.Set("newname", v.(string))
	}
	if v, found := p.p["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	return u
}

func (p *UpdateAccountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpdateAccountParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *UpdateAccountParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UpdateAccountParams) SetAccountdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountdetails"] = v
}

func (p *UpdateAccountParams) ResetAccountdetails() {
	if p.p != nil && p.p["accountdetails"] != nil {
		delete(p.p, "accountdetails")
	}
}

func (p *UpdateAccountParams) GetAccountdetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountdetails"].(map[string]string)
	return value, ok
}

func (p *UpdateAccountParams) SetApikeyaccess(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["apikeyaccess"] = v
}

func (p *UpdateAccountParams) ResetApikeyaccess() {
	if p.p != nil && p.p["apikeyaccess"] != nil {
		delete(p.p, "apikeyaccess")
	}
}

func (p *UpdateAccountParams) GetApikeyaccess() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["apikeyaccess"].(string)
	return value, ok
}

func (p *UpdateAccountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpdateAccountParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *UpdateAccountParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpdateAccountParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateAccountParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateAccountParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateAccountParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *UpdateAccountParams) ResetNetworkdomain() {
	if p.p != nil && p.p["networkdomain"] != nil {
		delete(p.p, "networkdomain")
	}
}

func (p *UpdateAccountParams) GetNetworkdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdomain"].(string)
	return value, ok
}

func (p *UpdateAccountParams) SetNewname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["newname"] = v
}

func (p *UpdateAccountParams) ResetNewname() {
	if p.p != nil && p.p["newname"] != nil {
		delete(p.p, "newname")
	}
}

func (p *UpdateAccountParams) GetNewname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["newname"].(string)
	return value, ok
}

func (p *UpdateAccountParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *UpdateAccountParams) ResetRoleid() {
	if p.p != nil && p.p["roleid"] != nil {
		delete(p.p, "roleid")
	}
}

func (p *UpdateAccountParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewUpdateAccountParams() *UpdateAccountParams {
	p := &UpdateAccountParams{}
	p.p = make(map[string]interface{})
	return p
}

// Updates account information for the authenticated user
func (s *AccountService) UpdateAccount(p *UpdateAccountParams) (*UpdateAccountResponse, error) {
	resp, err := s.cs.newPostRequest("updateAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateAccountResponse struct {
	Accountdetails            map[string]string           `json:"accountdetails"`
	Accounttype               int                         `json:"accounttype"`
	Apikeyaccess              string                      `json:"apikeyaccess"`
	Cpuavailable              string                      `json:"cpuavailable"`
	Cpulimit                  string                      `json:"cpulimit"`
	Cputotal                  int64                       `json:"cputotal"`
	Created                   string                      `json:"created"`
	Defaultzoneid             string                      `json:"defaultzoneid"`
	Domain                    string                      `json:"domain"`
	Domainid                  string                      `json:"domainid"`
	Domainpath                string                      `json:"domainpath"`
	Groups                    []string                    `json:"groups"`
	Icon                      interface{}                 `json:"icon"`
	Id                        string                      `json:"id"`
	Ipavailable               string                      `json:"ipavailable"`
	Iplimit                   string                      `json:"iplimit"`
	Iptotal                   int64                       `json:"iptotal"`
	Iscleanuprequired         bool                        `json:"iscleanuprequired"`
	Isdefault                 bool                        `json:"isdefault"`
	JobID                     string                      `json:"jobid"`
	Jobstatus                 int                         `json:"jobstatus"`
	Memoryavailable           string                      `json:"memoryavailable"`
	Memorylimit               string                      `json:"memorylimit"`
	Memorytotal               int64                       `json:"memorytotal"`
	Name                      string                      `json:"name"`
	Networkavailable          string                      `json:"networkavailable"`
	Networkdomain             string                      `json:"networkdomain"`
	Networklimit              string                      `json:"networklimit"`
	Networktotal              int64                       `json:"networktotal"`
	Primarystorageavailable   string                      `json:"primarystorageavailable"`
	Primarystoragelimit       string                      `json:"primarystoragelimit"`
	Primarystoragetotal       int64                       `json:"primarystoragetotal"`
	Projectavailable          string                      `json:"projectavailable"`
	Projectlimit              string                      `json:"projectlimit"`
	Projecttotal              int64                       `json:"projecttotal"`
	Receivedbytes             int64                       `json:"receivedbytes"`
	Roleid                    string                      `json:"roleid"`
	Rolename                  string                      `json:"rolename"`
	Roletype                  string                      `json:"roletype"`
	Secondarystorageavailable string                      `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                      `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                     `json:"secondarystoragetotal"`
	Sentbytes                 int64                       `json:"sentbytes"`
	Snapshotavailable         string                      `json:"snapshotavailable"`
	Snapshotlimit             string                      `json:"snapshotlimit"`
	Snapshottotal             int64                       `json:"snapshottotal"`
	State                     string                      `json:"state"`
	Taggedresources           []string                    `json:"taggedresources"`
	Templateavailable         string                      `json:"templateavailable"`
	Templatelimit             string                      `json:"templatelimit"`
	Templatetotal             int64                       `json:"templatetotal"`
	User                      []UpdateAccountResponseUser `json:"user"`
	Vmavailable               string                      `json:"vmavailable"`
	Vmlimit                   string                      `json:"vmlimit"`
	Vmrunning                 int                         `json:"vmrunning"`
	Vmstopped                 int                         `json:"vmstopped"`
	Vmtotal                   int64                       `json:"vmtotal"`
	Volumeavailable           string                      `json:"volumeavailable"`
	Volumelimit               string                      `json:"volumelimit"`
	Volumetotal               int64                       `json:"volumetotal"`
	Vpcavailable              string                      `json:"vpcavailable"`
	Vpclimit                  string                      `json:"vpclimit"`
	Vpctotal                  int64                       `json:"vpctotal"`
}

type UpdateAccountResponseUser struct {
	Account             string      `json:"account"`
	Accountid           string      `json:"accountid"`
	Accounttype         int         `json:"accounttype"`
	Apikey              string      `json:"apikey"`
	Apikeyaccess        string      `json:"apikeyaccess"`
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
