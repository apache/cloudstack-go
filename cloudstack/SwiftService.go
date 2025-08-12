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

type SwiftServiceIface interface {
	AddSwift(p *AddSwiftParams) (*AddSwiftResponse, error)
	NewAddSwiftParams(url string) *AddSwiftParams
	ListSwifts(p *ListSwiftsParams) (*ListSwiftsResponse, error)
	NewListSwiftsParams() *ListSwiftsParams
	GetSwiftID(keyword string, opts ...OptionFunc) (string, int, error)
}

type AddSwiftParams struct {
	p map[string]interface{}
}

func (p *AddSwiftParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["key"]; found {
		u.Set("key", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddSwiftParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AddSwiftParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *AddSwiftParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *AddSwiftParams) SetKey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["key"] = v
}

func (p *AddSwiftParams) ResetKey() {
	if p.p != nil && p.p["key"] != nil {
		delete(p.p, "key")
	}
}

func (p *AddSwiftParams) GetKey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["key"].(string)
	return value, ok
}

func (p *AddSwiftParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddSwiftParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddSwiftParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddSwiftParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddSwiftParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddSwiftParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddSwiftParams instance,
// as then you are sure you have configured all required params
func (s *SwiftService) NewAddSwiftParams(url string) *AddSwiftParams {
	p := &AddSwiftParams{}
	p.p = make(map[string]interface{})
	p.p["url"] = url
	return p
}

// Adds Swift.
func (s *SwiftService) AddSwift(p *AddSwiftParams) (*AddSwiftResponse, error) {
	resp, err := s.cs.newPostRequest("addSwift", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddSwiftResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddSwiftResponse struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}

type ListSwiftsParams struct {
	p map[string]interface{}
}

func (p *ListSwiftsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
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
	return u
}

func (p *ListSwiftsParams) SetId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSwiftsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListSwiftsParams) GetId() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int64)
	return value, ok
}

func (p *ListSwiftsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSwiftsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSwiftsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSwiftsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSwiftsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSwiftsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSwiftsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSwiftsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSwiftsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListSwiftsParams instance,
// as then you are sure you have configured all required params
func (s *SwiftService) NewListSwiftsParams() *ListSwiftsParams {
	p := &ListSwiftsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SwiftService) GetSwiftID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListSwiftsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSwifts(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.Swifts[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Swifts {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// List Swift.
func (s *SwiftService) ListSwifts(p *ListSwiftsParams) (*ListSwiftsResponse, error) {
	resp, err := s.cs.newRequest("listSwifts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSwiftsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSwiftsResponse struct {
	Count  int      `json:"count"`
	Swifts []*Swift `json:"swift"`
}

type Swift struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}
