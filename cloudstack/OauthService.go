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

type OauthServiceIface interface {
	ListOauthProvider(p *ListOauthProviderParams) (*ListOauthProviderResponse, error)
	NewListOauthProviderParams() *ListOauthProviderParams
	GetOauthProviderID(keyword string, opts ...OptionFunc) (string, int, error)
	GetOauthProviderByName(name string, opts ...OptionFunc) (*OauthProvider, int, error)
	GetOauthProviderByID(id string, opts ...OptionFunc) (*OauthProvider, int, error)
	UpdateOauthProvider(p *UpdateOauthProviderParams) (*UpdateOauthProviderResponse, error)
	NewUpdateOauthProviderParams(id string) *UpdateOauthProviderParams
	DeleteOauthProvider(p *DeleteOauthProviderParams) (*DeleteOauthProviderResponse, error)
	NewDeleteOauthProviderParams(id string) *DeleteOauthProviderParams
}

type ListOauthProviderParams struct {
	p map[string]interface{}
}

func (p *ListOauthProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (p *ListOauthProviderParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListOauthProviderParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListOauthProviderParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListOauthProviderParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListOauthProviderParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListOauthProviderParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListOauthProviderParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListOauthProviderParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListOauthProviderParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListOauthProviderParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListOauthProviderParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListOauthProviderParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListOauthProviderParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListOauthProviderParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListOauthProviderParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

// You should always use this function to get a new ListOauthProviderParams instance,
// as then you are sure you have configured all required params
func (s *OauthService) NewListOauthProviderParams() *ListOauthProviderParams {
	p := &ListOauthProviderParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *OauthService) GetOauthProviderID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListOauthProviderParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListOauthProvider(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.OauthProvider[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.OauthProvider {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *OauthService) GetOauthProviderByName(name string, opts ...OptionFunc) (*OauthProvider, int, error) {
	id, count, err := s.GetOauthProviderID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetOauthProviderByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *OauthService) GetOauthProviderByID(id string, opts ...OptionFunc) (*OauthProvider, int, error) {
	p := &ListOauthProviderParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOauthProvider(p)
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
		return l.OauthProvider[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for OauthProvider UUID: %s!", id)
}

// List OAuth providers registered
func (s *OauthService) ListOauthProvider(p *ListOauthProviderParams) (*ListOauthProviderResponse, error) {
	resp, err := s.cs.newRequest("listOauthProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListOauthProviderResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListOauthProviderResponse struct {
	Count         int              `json:"count"`
	OauthProvider []*OauthProvider `json:"oauthprovider"`
}

type OauthProvider struct {
	Clientid    string `json:"clientid"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	Redirecturi string `json:"redirecturi"`
	Secretkey   string `json:"secretkey"`
}

type UpdateOauthProviderParams struct {
	p map[string]interface{}
}

func (p *UpdateOauthProviderParams) toURLValues() url.Values {
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
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["redirecturi"]; found {
		u.Set("redirecturi", v.(string))
	}
	if v, found := p.p["secretkey"]; found {
		u.Set("secretkey", v.(string))
	}
	return u
}

func (p *UpdateOauthProviderParams) SetClientid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clientid"] = v
}

func (p *UpdateOauthProviderParams) ResetClientid() {
	if p.p != nil && p.p["clientid"] != nil {
		delete(p.p, "clientid")
	}
}

func (p *UpdateOauthProviderParams) GetClientid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clientid"].(string)
	return value, ok
}

func (p *UpdateOauthProviderParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateOauthProviderParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateOauthProviderParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateOauthProviderParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *UpdateOauthProviderParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *UpdateOauthProviderParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *UpdateOauthProviderParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateOauthProviderParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateOauthProviderParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateOauthProviderParams) SetRedirecturi(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["redirecturi"] = v
}

func (p *UpdateOauthProviderParams) ResetRedirecturi() {
	if p.p != nil && p.p["redirecturi"] != nil {
		delete(p.p, "redirecturi")
	}
}

func (p *UpdateOauthProviderParams) GetRedirecturi() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["redirecturi"].(string)
	return value, ok
}

func (p *UpdateOauthProviderParams) SetSecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secretkey"] = v
}

func (p *UpdateOauthProviderParams) ResetSecretkey() {
	if p.p != nil && p.p["secretkey"] != nil {
		delete(p.p, "secretkey")
	}
}

func (p *UpdateOauthProviderParams) GetSecretkey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["secretkey"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateOauthProviderParams instance,
// as then you are sure you have configured all required params
func (s *OauthService) NewUpdateOauthProviderParams(id string) *UpdateOauthProviderParams {
	p := &UpdateOauthProviderParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates the registered OAuth provider details
func (s *OauthService) UpdateOauthProvider(p *UpdateOauthProviderParams) (*UpdateOauthProviderResponse, error) {
	resp, err := s.cs.newPostRequest("updateOauthProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response UpdateOauthProviderResponse `json:"oauthprovider"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type UpdateOauthProviderResponse struct {
	Clientid    string `json:"clientid"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	Redirecturi string `json:"redirecturi"`
	Secretkey   string `json:"secretkey"`
}

type DeleteOauthProviderParams struct {
	p map[string]interface{}
}

func (p *DeleteOauthProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteOauthProviderParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteOauthProviderParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteOauthProviderParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteOauthProviderParams instance,
// as then you are sure you have configured all required params
func (s *OauthService) NewDeleteOauthProviderParams(id string) *DeleteOauthProviderParams {
	p := &DeleteOauthProviderParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes the registered OAuth provider
func (s *OauthService) DeleteOauthProvider(p *DeleteOauthProviderParams) (*DeleteOauthProviderResponse, error) {
	resp, err := s.cs.newPostRequest("deleteOauthProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteOauthProviderResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteOauthProviderResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteOauthProviderResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteOauthProviderResponse
	return json.Unmarshal(b, (*alias)(r))
}
