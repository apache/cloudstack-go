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

type AuthenticationServiceIface interface {
	Login(p *LoginParams) (*LoginResponse, error)
	NewLoginParams(password string, username string) *LoginParams
	Logout(p *LogoutParams) (*LogoutResponse, error)
	NewLogoutParams() *LogoutParams
	Oauthlogin(p *OauthloginParams) (*OauthloginResponse, error)
	NewOauthloginParams(email string, provider string) *OauthloginParams
}

type LoginParams struct {
	p map[string]interface{}
}

func (p *LoginParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["domainId"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("domainId", vv)
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *LoginParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *LoginParams) ResetDomain() {
	if p.p != nil && p.p["domain"] != nil {
		delete(p.p, "domain")
	}
}

func (p *LoginParams) GetDomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domain"].(string)
	return value, ok
}

func (p *LoginParams) SetDomainId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainId"] = v
}

func (p *LoginParams) ResetDomainId() {
	if p.p != nil && p.p["domainId"] != nil {
		delete(p.p, "domainId")
	}
}

func (p *LoginParams) GetDomainId() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainId"].(int64)
	return value, ok
}

func (p *LoginParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *LoginParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *LoginParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *LoginParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *LoginParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *LoginParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new LoginParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLoginParams(password string, username string) *LoginParams {
	p := &LoginParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Logs a user into the CloudStack. A successful login attempt will generate a JSESSIONID cookie value that can be passed in subsequent Query command calls until the "logout" command has been issued or the session has expired.
func (s *AuthenticationService) Login(p *LoginParams) (*LoginResponse, error) {
	resp, err := s.cs.newPostRequest("login", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LoginResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LoginResponse struct {
	Account        string `json:"account"`
	Domainid       string `json:"domainid"`
	Firstname      string `json:"firstname"`
	Is2faenabled   string `json:"is2faenabled"`
	Is2faverified  string `json:"is2faverified"`
	Issuerfor2fa   string `json:"issuerfor2fa"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Lastname       string `json:"lastname"`
	Providerfor2fa string `json:"providerfor2fa"`
	Registered     string `json:"registered"`
	Sessionkey     string `json:"sessionkey"`
	Timeout        int    `json:"timeout"`
	Timezone       string `json:"timezone"`
	Timezoneoffset string `json:"timezoneoffset"`
	Type           string `json:"type"`
	Userid         string `json:"userid"`
	Username       string `json:"username"`
}

type LogoutParams struct {
	p map[string]interface{}
}

func (p *LogoutParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new LogoutParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLogoutParams() *LogoutParams {
	p := &LogoutParams{}
	p.p = make(map[string]interface{})
	return p
}

// Logs out the user
func (s *AuthenticationService) Logout(p *LogoutParams) (*LogoutResponse, error) {
	resp, err := s.cs.newPostRequest("logout", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LogoutResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LogoutResponse struct {
	Description string `json:"description"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
}

type OauthloginParams struct {
	p map[string]interface{}
}

func (p *OauthloginParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["domainId"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("domainId", vv)
	}
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["secretcode"]; found {
		u.Set("secretcode", v.(string))
	}
	return u
}

func (p *OauthloginParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *OauthloginParams) ResetDomain() {
	if p.p != nil && p.p["domain"] != nil {
		delete(p.p, "domain")
	}
}

func (p *OauthloginParams) GetDomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domain"].(string)
	return value, ok
}

func (p *OauthloginParams) SetDomainId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainId"] = v
}

func (p *OauthloginParams) ResetDomainId() {
	if p.p != nil && p.p["domainId"] != nil {
		delete(p.p, "domainId")
	}
}

func (p *OauthloginParams) GetDomainId() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainId"].(int64)
	return value, ok
}

func (p *OauthloginParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
}

func (p *OauthloginParams) ResetEmail() {
	if p.p != nil && p.p["email"] != nil {
		delete(p.p, "email")
	}
}

func (p *OauthloginParams) GetEmail() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["email"].(string)
	return value, ok
}

func (p *OauthloginParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *OauthloginParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *OauthloginParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *OauthloginParams) SetSecretcode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secretcode"] = v
}

func (p *OauthloginParams) ResetSecretcode() {
	if p.p != nil && p.p["secretcode"] != nil {
		delete(p.p, "secretcode")
	}
}

func (p *OauthloginParams) GetSecretcode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["secretcode"].(string)
	return value, ok
}

// You should always use this function to get a new OauthloginParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewOauthloginParams(email string, provider string) *OauthloginParams {
	p := &OauthloginParams{}
	p.p = make(map[string]interface{})
	p.p["email"] = email
	p.p["provider"] = provider
	return p
}

// Logs a user into the CloudStack after successful verification of OAuth secret code from the particular provider.A successful login attempt will generate a JSESSIONID cookie value that can be passed in subsequent Query command calls until the "logout" command has been issued or the session has expired.
func (s *AuthenticationService) Oauthlogin(p *OauthloginParams) (*OauthloginResponse, error) {
	resp, err := s.cs.newPostRequest("oauthlogin", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r OauthloginResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type OauthloginResponse struct {
	Account        string `json:"account"`
	Domainid       string `json:"domainid"`
	Firstname      string `json:"firstname"`
	Is2faenabled   string `json:"is2faenabled"`
	Is2faverified  string `json:"is2faverified"`
	Issuerfor2fa   string `json:"issuerfor2fa"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Lastname       string `json:"lastname"`
	Providerfor2fa string `json:"providerfor2fa"`
	Registered     string `json:"registered"`
	Sessionkey     string `json:"sessionkey"`
	Timeout        int    `json:"timeout"`
	Timezone       string `json:"timezone"`
	Timezoneoffset string `json:"timezoneoffset"`
	Type           string `json:"type"`
	Userid         string `json:"userid"`
	Username       string `json:"username"`
}
