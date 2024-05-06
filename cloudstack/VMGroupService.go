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

type VMGroupServiceIface interface {
	CreateInstanceGroup(p *CreateInstanceGroupParams) (*CreateInstanceGroupResponse, error)
	NewCreateInstanceGroupParams(name string) *CreateInstanceGroupParams
	DeleteInstanceGroup(p *DeleteInstanceGroupParams) (*DeleteInstanceGroupResponse, error)
	NewDeleteInstanceGroupParams(id string) *DeleteInstanceGroupParams
	ListInstanceGroups(p *ListInstanceGroupsParams) (*ListInstanceGroupsResponse, error)
	NewListInstanceGroupsParams() *ListInstanceGroupsParams
	GetInstanceGroupID(name string, opts ...OptionFunc) (string, int, error)
	GetInstanceGroupByName(name string, opts ...OptionFunc) (*InstanceGroup, int, error)
	GetInstanceGroupByID(id string, opts ...OptionFunc) (*InstanceGroup, int, error)
	UpdateInstanceGroup(p *UpdateInstanceGroupParams) (*UpdateInstanceGroupResponse, error)
	NewUpdateInstanceGroupParams(id string) *UpdateInstanceGroupParams
}

type CreateInstanceGroupParams struct {
	p map[string]interface{}
}

func (p *CreateInstanceGroupParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *CreateInstanceGroupParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateInstanceGroupParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateInstanceGroupParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateInstanceGroupParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateInstanceGroupParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateInstanceGroupParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateInstanceGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateInstanceGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateInstanceGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateInstanceGroupParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateInstanceGroupParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateInstanceGroupParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateInstanceGroupParams instance,
// as then you are sure you have configured all required params
func (s *VMGroupService) NewCreateInstanceGroupParams(name string) *CreateInstanceGroupParams {
	p := &CreateInstanceGroupParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Creates a vm group
func (s *VMGroupService) CreateInstanceGroup(p *CreateInstanceGroupParams) (*CreateInstanceGroupResponse, error) {
	resp, err := s.cs.newRequest("createInstanceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateInstanceGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateInstanceGroupResponse struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
}

type DeleteInstanceGroupParams struct {
	p map[string]interface{}
}

func (p *DeleteInstanceGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteInstanceGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteInstanceGroupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteInstanceGroupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteInstanceGroupParams instance,
// as then you are sure you have configured all required params
func (s *VMGroupService) NewDeleteInstanceGroupParams(id string) *DeleteInstanceGroupParams {
	p := &DeleteInstanceGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a vm group
func (s *VMGroupService) DeleteInstanceGroup(p *DeleteInstanceGroupParams) (*DeleteInstanceGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteInstanceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteInstanceGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteInstanceGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteInstanceGroupResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteInstanceGroupResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListInstanceGroupsParams struct {
	p map[string]interface{}
}

func (p *ListInstanceGroupsParams) toURLValues() url.Values {
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

func (p *ListInstanceGroupsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListInstanceGroupsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListInstanceGroupsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListInstanceGroupsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListInstanceGroupsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListInstanceGroupsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListInstanceGroupsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListInstanceGroupsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListInstanceGroupsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListInstanceGroupsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListInstanceGroupsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListInstanceGroupsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListInstanceGroupsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListInstanceGroupsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListInstanceGroupsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListInstanceGroupsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListInstanceGroupsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListInstanceGroupsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListInstanceGroupsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListInstanceGroupsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListInstanceGroupsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListInstanceGroupsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListInstanceGroupsParams instance,
// as then you are sure you have configured all required params
func (s *VMGroupService) NewListInstanceGroupsParams() *ListInstanceGroupsParams {
	p := &ListInstanceGroupsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VMGroupService) GetInstanceGroupID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListInstanceGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListInstanceGroups(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.InstanceGroups[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.InstanceGroups {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VMGroupService) GetInstanceGroupByName(name string, opts ...OptionFunc) (*InstanceGroup, int, error) {
	id, count, err := s.GetInstanceGroupID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetInstanceGroupByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VMGroupService) GetInstanceGroupByID(id string, opts ...OptionFunc) (*InstanceGroup, int, error) {
	p := &ListInstanceGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListInstanceGroups(p)
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
		return l.InstanceGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for InstanceGroup UUID: %s!", id)
}

// Lists vm groups
func (s *VMGroupService) ListInstanceGroups(p *ListInstanceGroupsParams) (*ListInstanceGroupsResponse, error) {
	resp, err := s.cs.newRequest("listInstanceGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListInstanceGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListInstanceGroupsResponse struct {
	Count          int              `json:"count"`
	InstanceGroups []*InstanceGroup `json:"instancegroup"`
}

type InstanceGroup struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
}

type UpdateInstanceGroupParams struct {
	p map[string]interface{}
}

func (p *UpdateInstanceGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateInstanceGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateInstanceGroupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateInstanceGroupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateInstanceGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateInstanceGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateInstanceGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateInstanceGroupParams instance,
// as then you are sure you have configured all required params
func (s *VMGroupService) NewUpdateInstanceGroupParams(id string) *UpdateInstanceGroupParams {
	p := &UpdateInstanceGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a vm group
func (s *VMGroupService) UpdateInstanceGroup(p *UpdateInstanceGroupParams) (*UpdateInstanceGroupResponse, error) {
	resp, err := s.cs.newRequest("updateInstanceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateInstanceGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateInstanceGroupResponse struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
}
