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

type ManagementServiceIface interface {
	ListManagementServers(p *ListManagementServersParams) (*ListManagementServersResponse, error)
	NewListManagementServersParams() *ListManagementServersParams
	GetManagementServerID(name string, opts ...OptionFunc) (string, int, error)
	GetManagementServerByName(name string, opts ...OptionFunc) (*ManagementServer, int, error)
	GetManagementServerByID(id string, opts ...OptionFunc) (*ManagementServer, int, error)
}

type ListManagementServersParams struct {
	p map[string]interface{}
}

func (p *ListManagementServersParams) toURLValues() url.Values {
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
	return u
}

func (p *ListManagementServersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListManagementServersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListManagementServersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListManagementServersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListManagementServersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListManagementServersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListManagementServersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListManagementServersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListManagementServersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListManagementServersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListManagementServersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListManagementServersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListManagementServersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListManagementServersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListManagementServersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListManagementServersParams instance,
// as then you are sure you have configured all required params
func (s *ManagementService) NewListManagementServersParams() *ListManagementServersParams {
	p := &ListManagementServersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ManagementService) GetManagementServerID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListManagementServersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListManagementServers(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ManagementServers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ManagementServers {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ManagementService) GetManagementServerByName(name string, opts ...OptionFunc) (*ManagementServer, int, error) {
	id, count, err := s.GetManagementServerID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetManagementServerByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ManagementService) GetManagementServerByID(id string, opts ...OptionFunc) (*ManagementServer, int, error) {
	p := &ListManagementServersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListManagementServers(p)
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
		return l.ManagementServers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ManagementServer UUID: %s!", id)
}

// Lists management servers.
func (s *ManagementService) ListManagementServers(p *ListManagementServersParams) (*ListManagementServersResponse, error) {
	resp, err := s.cs.newRequest("listManagementServers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListManagementServersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListManagementServersResponse struct {
	Count             int                 `json:"count"`
	ManagementServers []*ManagementServer `json:"managementserver"`
}

type ManagementServer struct {
	Id               string `json:"id"`
	Javadistribution string `json:"javadistribution"`
	Javaversion      string `json:"javaversion"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Kernelversion    string `json:"kernelversion"`
	Lastboottime     string `json:"lastboottime"`
	Lastserverstart  string `json:"lastserverstart"`
	Lastserverstop   string `json:"lastserverstop"`
	Name             string `json:"name"`
	Osdistribution   string `json:"osdistribution"`
	Serviceip        string `json:"serviceip"`
	State            string `json:"state"`
	Version          string `json:"version"`
}
