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

type RegistrationServiceIface interface {
	RegisterOauthProvider(p *RegisterOauthProviderParams) (*RegisterOauthProviderResponse, error)
	NewRegisterOauthProviderParams(clientid string, description string, provider string, redirecturi string, secretkey string) *RegisterOauthProviderParams
}

type RegisterOauthProviderParams struct {
	p map[string]interface{}
}

func (p *RegisterOauthProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clientid"]; found {
		u.Set("clientid", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["details"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("details[%d].%s", i, key), val)
			}
		}
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["redirecturi"]; found {
		u.Set("redirecturi", v.(string))
	}
	if v, found := p.p["secretkey"]; found {
		u.Set("secretkey", v.(string))
	}
	return u
}

func (p *RegisterOauthProviderParams) SetClientid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clientid"] = v
}

func (p *RegisterOauthProviderParams) ResetClientid() {
	if p.p != nil && p.p["clientid"] != nil {
		delete(p.p, "clientid")
	}
}

func (p *RegisterOauthProviderParams) GetClientid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clientid"].(string)
	return value, ok
}

func (p *RegisterOauthProviderParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *RegisterOauthProviderParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *RegisterOauthProviderParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *RegisterOauthProviderParams) SetDetails(v []map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *RegisterOauthProviderParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *RegisterOauthProviderParams) GetDetails() ([]map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]map[string]string)
	return value, ok
}

func (p *RegisterOauthProviderParams) AddDetails(item map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	val, found := p.p["details"]
	if !found {
		p.p["details"] = []map[string]string{}
		val = p.p["details"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	p.p["details"] = l
}

func (p *RegisterOauthProviderParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *RegisterOauthProviderParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *RegisterOauthProviderParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *RegisterOauthProviderParams) SetRedirecturi(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["redirecturi"] = v
}

func (p *RegisterOauthProviderParams) ResetRedirecturi() {
	if p.p != nil && p.p["redirecturi"] != nil {
		delete(p.p, "redirecturi")
	}
}

func (p *RegisterOauthProviderParams) GetRedirecturi() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["redirecturi"].(string)
	return value, ok
}

func (p *RegisterOauthProviderParams) SetSecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secretkey"] = v
}

func (p *RegisterOauthProviderParams) ResetSecretkey() {
	if p.p != nil && p.p["secretkey"] != nil {
		delete(p.p, "secretkey")
	}
}

func (p *RegisterOauthProviderParams) GetSecretkey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["secretkey"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterOauthProviderParams instance,
// as then you are sure you have configured all required params
func (s *RegistrationService) NewRegisterOauthProviderParams(clientid string, description string, provider string, redirecturi string, secretkey string) *RegisterOauthProviderParams {
	p := &RegisterOauthProviderParams{}
	p.p = make(map[string]interface{})
	p.p["clientid"] = clientid
	p.p["description"] = description
	p.p["provider"] = provider
	p.p["redirecturi"] = redirecturi
	p.p["secretkey"] = secretkey
	return p
}

// Register the OAuth2 provider in CloudStack
func (s *RegistrationService) RegisterOauthProvider(p *RegisterOauthProviderParams) (*RegisterOauthProviderResponse, error) {
	resp, err := s.cs.newRequest("registerOauthProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterOauthProviderResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterOauthProviderResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RegisterOauthProviderResponse) UnmarshalJSON(b []byte) error {
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

	type alias RegisterOauthProviderResponse
	return json.Unmarshal(b, (*alias)(r))
}
