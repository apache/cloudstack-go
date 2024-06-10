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

type UserServiceIface interface {
	CreateUser(p *CreateUserParams) (*CreateUserResponse, error)
	NewCreateUserParams(account string, email string, firstname string, lastname string, password string, username string) *CreateUserParams
	DeleteUser(p *DeleteUserParams) (*DeleteUserResponse, error)
	NewDeleteUserParams(id string) *DeleteUserParams
	DisableUser(p *DisableUserParams) (*DisableUserResponse, error)
	NewDisableUserParams(id string) *DisableUserParams
	EnableUser(p *EnableUserParams) (*EnableUserResponse, error)
	NewEnableUserParams(id string) *EnableUserParams
	GetUser(p *GetUserParams) (*GetUserResponse, error)
	NewGetUserParams(userapikey string) *GetUserParams
	GetUserKeys(p *GetUserKeysParams) (*GetUserKeysResponse, error)
	NewGetUserKeysParams(id string) *GetUserKeysParams
	GetVirtualMachineUserData(p *GetVirtualMachineUserDataParams) (*GetVirtualMachineUserDataResponse, error)
	NewGetVirtualMachineUserDataParams(virtualmachineid string) *GetVirtualMachineUserDataParams
	ListUsers(p *ListUsersParams) (*ListUsersResponse, error)
	NewListUsersParams() *ListUsersParams
	GetUserByID(id string, opts ...OptionFunc) (*User, int, error)
	LockUser(p *LockUserParams) (*LockUserResponse, error)
	NewLockUserParams(id string) *LockUserParams
	RegisterUserKeys(p *RegisterUserKeysParams) (*RegisterUserKeysResponse, error)
	NewRegisterUserKeysParams(id string) *RegisterUserKeysParams
	UpdateUser(p *UpdateUserParams) (*UpdateUserResponse, error)
	NewUpdateUserParams(id string) *UpdateUserParams
	ListUserData(p *ListUserDataParams) (*ListUserDataResponse, error)
	NewListUserDataParams() *ListUserDataParams
	GetUserDataID(name string, opts ...OptionFunc) (string, int, error)
	GetUserDataByName(name string, opts ...OptionFunc) (*UserData, int, error)
	GetUserDataByID(id string, opts ...OptionFunc) (*UserData, int, error)
	DeleteUserData(p *DeleteUserDataParams) (*DeleteUserDataResponse, error)
	NewDeleteUserDataParams(id string) *DeleteUserDataParams
	RegisterUserData(p *RegisterUserDataParams) (*RegisterUserDataResponse, error)
	NewRegisterUserDataParams(name string, userdata string) *RegisterUserDataParams
	MoveUser(p *MoveUserParams) (*MoveUserResponse, error)
	NewMoveUserParams(id string) *MoveUserParams
}

type CreateUserParams struct {
	p map[string]interface{}
}

func (p *CreateUserParams) toURLValues() url.Values {
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
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := p.p["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
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

func (p *CreateUserParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateUserParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateUserParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateUserParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateUserParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateUserParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateUserParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
}

func (p *CreateUserParams) ResetEmail() {
	if p.p != nil && p.p["email"] != nil {
		delete(p.p, "email")
	}
}

func (p *CreateUserParams) GetEmail() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["email"].(string)
	return value, ok
}

func (p *CreateUserParams) SetFirstname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firstname"] = v
}

func (p *CreateUserParams) ResetFirstname() {
	if p.p != nil && p.p["firstname"] != nil {
		delete(p.p, "firstname")
	}
}

func (p *CreateUserParams) GetFirstname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["firstname"].(string)
	return value, ok
}

func (p *CreateUserParams) SetLastname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lastname"] = v
}

func (p *CreateUserParams) ResetLastname() {
	if p.p != nil && p.p["lastname"] != nil {
		delete(p.p, "lastname")
	}
}

func (p *CreateUserParams) GetLastname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lastname"].(string)
	return value, ok
}

func (p *CreateUserParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *CreateUserParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *CreateUserParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *CreateUserParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *CreateUserParams) ResetTimezone() {
	if p.p != nil && p.p["timezone"] != nil {
		delete(p.p, "timezone")
	}
}

func (p *CreateUserParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

func (p *CreateUserParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *CreateUserParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *CreateUserParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

func (p *CreateUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *CreateUserParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *CreateUserParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new CreateUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewCreateUserParams(account string, email string, firstname string, lastname string, password string, username string) *CreateUserParams {
	p := &CreateUserParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["email"] = email
	p.p["firstname"] = firstname
	p.p["lastname"] = lastname
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Creates a user for an account that already exists
func (s *UserService) CreateUser(p *CreateUserParams) (*CreateUserResponse, error) {
	resp, err := s.cs.newPostRequest("createUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateUserResponse struct {
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
	JobID               string      `json:"jobid"`
	Jobstatus           int         `json:"jobstatus"`
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

type DeleteUserParams struct {
	p map[string]interface{}
}

func (p *DeleteUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteUserParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteUserParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDeleteUserParams(id string) *DeleteUserParams {
	p := &DeleteUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a user for an account
func (s *UserService) DeleteUser(p *DeleteUserParams) (*DeleteUserResponse, error) {
	resp, err := s.cs.newRequest("deleteUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteUserResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteUserResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteUserResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableUserParams struct {
	p map[string]interface{}
}

func (p *DisableUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DisableUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DisableUserParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DisableUserParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DisableUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDisableUserParams(id string) *DisableUserParams {
	p := &DisableUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Disables a user account
func (s *UserService) DisableUser(p *DisableUserParams) (*DisableUserResponse, error) {
	resp, err := s.cs.newRequest("disableUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableUserResponse
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

type DisableUserResponse struct {
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
	JobID               string      `json:"jobid"`
	Jobstatus           int         `json:"jobstatus"`
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

type EnableUserParams struct {
	p map[string]interface{}
}

func (p *EnableUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *EnableUserParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *EnableUserParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new EnableUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewEnableUserParams(id string) *EnableUserParams {
	p := &EnableUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Enables a user account
func (s *UserService) EnableUser(p *EnableUserParams) (*EnableUserResponse, error) {
	resp, err := s.cs.newRequest("enableUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r EnableUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type EnableUserResponse struct {
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
	JobID               string      `json:"jobid"`
	Jobstatus           int         `json:"jobstatus"`
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

type GetUserParams struct {
	p map[string]interface{}
}

func (p *GetUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["userapikey"]; found {
		u.Set("userapikey", v.(string))
	}
	return u
}

func (p *GetUserParams) SetUserapikey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userapikey"] = v
}

func (p *GetUserParams) ResetUserapikey() {
	if p.p != nil && p.p["userapikey"] != nil {
		delete(p.p, "userapikey")
	}
}

func (p *GetUserParams) GetUserapikey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userapikey"].(string)
	return value, ok
}

// You should always use this function to get a new GetUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewGetUserParams(userapikey string) *GetUserParams {
	p := &GetUserParams{}
	p.p = make(map[string]interface{})
	p.p["userapikey"] = userapikey
	return p
}

// Find user account by API key
func (s *UserService) GetUser(p *GetUserParams) (*GetUserResponse, error) {
	resp, err := s.cs.newRequest("getUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetUserResponse struct {
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
	JobID               string      `json:"jobid"`
	Jobstatus           int         `json:"jobstatus"`
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

type GetUserKeysParams struct {
	p map[string]interface{}
}

func (p *GetUserKeysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *GetUserKeysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *GetUserKeysParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *GetUserKeysParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new GetUserKeysParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewGetUserKeysParams(id string) *GetUserKeysParams {
	p := &GetUserKeysParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This command allows the user to query the seceret and API keys for the account
func (s *UserService) GetUserKeys(p *GetUserKeysParams) (*GetUserKeysResponse, error) {
	resp, err := s.cs.newRequest("getUserKeys", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r GetUserKeysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetUserKeysResponse struct {
	Apikey    string `json:"apikey"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Secretkey string `json:"secretkey"`
}

type GetVirtualMachineUserDataParams struct {
	p map[string]interface{}
}

func (p *GetVirtualMachineUserDataParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *GetVirtualMachineUserDataParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *GetVirtualMachineUserDataParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *GetVirtualMachineUserDataParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new GetVirtualMachineUserDataParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewGetVirtualMachineUserDataParams(virtualmachineid string) *GetVirtualMachineUserDataParams {
	p := &GetVirtualMachineUserDataParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Returns user data associated with the VM
func (s *UserService) GetVirtualMachineUserData(p *GetVirtualMachineUserDataParams) (*GetVirtualMachineUserDataResponse, error) {
	resp, err := s.cs.newRequest("getVirtualMachineUserData", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r GetVirtualMachineUserDataResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetVirtualMachineUserDataResponse struct {
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Userdata         string `json:"userdata"`
	Virtualmachineid string `json:"virtualmachineid"`
}

type ListUsersParams struct {
	p map[string]interface{}
}

func (p *ListUsersParams) toURLValues() url.Values {
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
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *ListUsersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListUsersParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListUsersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListUsersParams) SetAccounttype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *ListUsersParams) ResetAccounttype() {
	if p.p != nil && p.p["accounttype"] != nil {
		delete(p.p, "accounttype")
	}
}

func (p *ListUsersParams) GetAccounttype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounttype"].(int)
	return value, ok
}

func (p *ListUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListUsersParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListUsersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListUsersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListUsersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListUsersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListUsersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListUsersParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListUsersParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListUsersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListUsersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListUsersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListUsersParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListUsersParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListUsersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListUsersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListUsersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListUsersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListUsersParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListUsersParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListUsersParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListUsersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListUsersParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListUsersParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListUsersParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *ListUsersParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *ListUsersParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new ListUsersParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewListUsersParams() *ListUsersParams {
	p := &ListUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UserService) GetUserByID(id string, opts ...OptionFunc) (*User, int, error) {
	p := &ListUsersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListUsers(p)
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
		return l.Users[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for User UUID: %s!", id)
}

// Lists user accounts
func (s *UserService) ListUsers(p *ListUsersParams) (*ListUsersResponse, error) {
	resp, err := s.cs.newRequest("listUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUsersResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"user"`
}

type User struct {
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
	JobID               string      `json:"jobid"`
	Jobstatus           int         `json:"jobstatus"`
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

type LockUserParams struct {
	p map[string]interface{}
}

func (p *LockUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *LockUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *LockUserParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *LockUserParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new LockUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewLockUserParams(id string) *LockUserParams {
	p := &LockUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Locks a user account
func (s *UserService) LockUser(p *LockUserParams) (*LockUserResponse, error) {
	resp, err := s.cs.newRequest("lockUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r LockUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LockUserResponse struct {
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
	JobID               string      `json:"jobid"`
	Jobstatus           int         `json:"jobstatus"`
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

type RegisterUserKeysParams struct {
	p map[string]interface{}
}

func (p *RegisterUserKeysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RegisterUserKeysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RegisterUserKeysParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RegisterUserKeysParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterUserKeysParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewRegisterUserKeysParams(id string) *RegisterUserKeysParams {
	p := &RegisterUserKeysParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This command allows a user to register for the developer API, returning a secret key and an API key. This request is made through the integration API port, so it is a privileged command and must be made on behalf of a user. It is up to the implementer just how the username and password are entered, and then how that translates to an integration API request. Both secret key and API key should be returned to the user
func (s *UserService) RegisterUserKeys(p *RegisterUserKeysParams) (*RegisterUserKeysResponse, error) {
	resp, err := s.cs.newRequest("registerUserKeys", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r RegisterUserKeysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterUserKeysResponse struct {
	Apikey    string `json:"apikey"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Secretkey string `json:"secretkey"`
}

type UpdateUserParams struct {
	p map[string]interface{}
}

func (p *UpdateUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["currentpassword"]; found {
		u.Set("currentpassword", v.(string))
	}
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := p.p["mandate2fa"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("mandate2fa", vv)
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["userapikey"]; found {
		u.Set("userapikey", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["usersecretkey"]; found {
		u.Set("usersecretkey", v.(string))
	}
	return u
}

func (p *UpdateUserParams) SetCurrentpassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["currentpassword"] = v
}

func (p *UpdateUserParams) ResetCurrentpassword() {
	if p.p != nil && p.p["currentpassword"] != nil {
		delete(p.p, "currentpassword")
	}
}

func (p *UpdateUserParams) GetCurrentpassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["currentpassword"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
}

func (p *UpdateUserParams) ResetEmail() {
	if p.p != nil && p.p["email"] != nil {
		delete(p.p, "email")
	}
}

func (p *UpdateUserParams) GetEmail() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["email"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetFirstname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firstname"] = v
}

func (p *UpdateUserParams) ResetFirstname() {
	if p.p != nil && p.p["firstname"] != nil {
		delete(p.p, "firstname")
	}
}

func (p *UpdateUserParams) GetFirstname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["firstname"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateUserParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateUserParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetLastname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lastname"] = v
}

func (p *UpdateUserParams) ResetLastname() {
	if p.p != nil && p.p["lastname"] != nil {
		delete(p.p, "lastname")
	}
}

func (p *UpdateUserParams) GetLastname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lastname"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetMandate2fa(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mandate2fa"] = v
}

func (p *UpdateUserParams) ResetMandate2fa() {
	if p.p != nil && p.p["mandate2fa"] != nil {
		delete(p.p, "mandate2fa")
	}
}

func (p *UpdateUserParams) GetMandate2fa() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mandate2fa"].(bool)
	return value, ok
}

func (p *UpdateUserParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *UpdateUserParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *UpdateUserParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *UpdateUserParams) ResetTimezone() {
	if p.p != nil && p.p["timezone"] != nil {
		delete(p.p, "timezone")
	}
}

func (p *UpdateUserParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetUserapikey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userapikey"] = v
}

func (p *UpdateUserParams) ResetUserapikey() {
	if p.p != nil && p.p["userapikey"] != nil {
		delete(p.p, "userapikey")
	}
}

func (p *UpdateUserParams) GetUserapikey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userapikey"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *UpdateUserParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *UpdateUserParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *UpdateUserParams) SetUsersecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usersecretkey"] = v
}

func (p *UpdateUserParams) ResetUsersecretkey() {
	if p.p != nil && p.p["usersecretkey"] != nil {
		delete(p.p, "usersecretkey")
	}
}

func (p *UpdateUserParams) GetUsersecretkey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["usersecretkey"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewUpdateUserParams(id string) *UpdateUserParams {
	p := &UpdateUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a user account
func (s *UserService) UpdateUser(p *UpdateUserParams) (*UpdateUserResponse, error) {
	resp, err := s.cs.newPostRequest("updateUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateUserResponse struct {
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
	JobID               string      `json:"jobid"`
	Jobstatus           int         `json:"jobstatus"`
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

type ListUserDataParams struct {
	p map[string]interface{}
}

func (p *ListUserDataParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *ListUserDataParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListUserDataParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListUserDataParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListUserDataParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListUserDataParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListUserDataParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListUserDataParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListUserDataParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListUserDataParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListUserDataParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListUserDataParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListUserDataParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListUserDataParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListUserDataParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListUserDataParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListUserDataParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListUserDataParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListUserDataParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListUserDataParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListUserDataParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListUserDataParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListUserDataParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListUserDataParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListUserDataParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListUserDataParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListUserDataParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListUserDataParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListUserDataParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListUserDataParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListUserDataParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUserDataParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewListUserDataParams() *ListUserDataParams {
	p := &ListUserDataParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UserService) GetUserDataID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListUserDataParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListUserData(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.UserData[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.UserData {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UserService) GetUserDataByName(name string, opts ...OptionFunc) (*UserData, int, error) {
	id, count, err := s.GetUserDataID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetUserDataByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UserService) GetUserDataByID(id string, opts ...OptionFunc) (*UserData, int, error) {
	p := &ListUserDataParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListUserData(p)
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
		return l.UserData[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for UserData UUID: %s!", id)
}

// List registered userdatas
func (s *UserService) ListUserData(p *ListUserDataParams) (*ListUserDataResponse, error) {
	resp, err := s.cs.newRequest("listUserData", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUserDataResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUserDataResponse struct {
	Count    int         `json:"count"`
	UserData []*UserData `json:"userdata"`
}

type UserData struct {
	Account        string `json:"account"`
	Accountid      string `json:"accountid"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Params         string `json:"params"`
	Userdata       string `json:"userdata"`
}

type DeleteUserDataParams struct {
	p map[string]interface{}
}

func (p *DeleteUserDataParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DeleteUserDataParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DeleteUserDataParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DeleteUserDataParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DeleteUserDataParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DeleteUserDataParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DeleteUserDataParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DeleteUserDataParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteUserDataParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteUserDataParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteUserDataParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DeleteUserDataParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *DeleteUserDataParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteUserDataParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDeleteUserDataParams(id string) *DeleteUserDataParams {
	p := &DeleteUserDataParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a userdata
func (s *UserService) DeleteUserData(p *DeleteUserDataParams) (*DeleteUserDataResponse, error) {
	resp, err := s.cs.newRequest("deleteUserData", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUserDataResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteUserDataResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteUserDataResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteUserDataResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RegisterUserDataParams struct {
	p map[string]interface{}
}

func (p *RegisterUserDataParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["params"]; found {
		u.Set("params", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	return u
}

func (p *RegisterUserDataParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *RegisterUserDataParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *RegisterUserDataParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *RegisterUserDataParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *RegisterUserDataParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *RegisterUserDataParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *RegisterUserDataParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *RegisterUserDataParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *RegisterUserDataParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *RegisterUserDataParams) SetParams(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["params"] = v
}

func (p *RegisterUserDataParams) ResetParams() {
	if p.p != nil && p.p["params"] != nil {
		delete(p.p, "params")
	}
}

func (p *RegisterUserDataParams) GetParams() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["params"].(string)
	return value, ok
}

func (p *RegisterUserDataParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *RegisterUserDataParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *RegisterUserDataParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *RegisterUserDataParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *RegisterUserDataParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *RegisterUserDataParams) GetUserdata() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterUserDataParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewRegisterUserDataParams(name string, userdata string) *RegisterUserDataParams {
	p := &RegisterUserDataParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["userdata"] = userdata
	return p
}

// Register a new userdata.
func (s *UserService) RegisterUserData(p *RegisterUserDataParams) (*RegisterUserDataResponse, error) {
	resp, err := s.cs.newPostRequest("registerUserData", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response RegisterUserDataResponse `json:"userdata"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type RegisterUserDataResponse struct {
	Account        string `json:"account"`
	Accountid      string `json:"accountid"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Params         string `json:"params"`
	Userdata       string `json:"userdata"`
}

type RegisterUserData struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RegisterUserData) UnmarshalJSON(b []byte) error {
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

	type alias RegisterUserData
	return json.Unmarshal(b, (*alias)(r))
}

type MoveUserParams struct {
	p map[string]interface{}
}

func (p *MoveUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *MoveUserParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *MoveUserParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *MoveUserParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *MoveUserParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *MoveUserParams) ResetAccountid() {
	if p.p != nil && p.p["accountid"] != nil {
		delete(p.p, "accountid")
	}
}

func (p *MoveUserParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *MoveUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *MoveUserParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *MoveUserParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new MoveUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewMoveUserParams(id string) *MoveUserParams {
	p := &MoveUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Moves a user to another account
func (s *UserService) MoveUser(p *MoveUserParams) (*MoveUserResponse, error) {
	resp, err := s.cs.newRequest("moveUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MoveUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type MoveUserResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *MoveUserResponse) UnmarshalJSON(b []byte) error {
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

	type alias MoveUserResponse
	return json.Unmarshal(b, (*alias)(r))
}
