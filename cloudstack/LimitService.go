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

type LimitServiceIface interface {
	GetApiLimit(p *GetApiLimitParams) (*GetApiLimitResponse, error)
	NewGetApiLimitParams() *GetApiLimitParams
	ListResourceLimits(p *ListResourceLimitsParams) (*ListResourceLimitsResponse, error)
	NewListResourceLimitsParams() *ListResourceLimitsParams
	ResetApiLimit(p *ResetApiLimitParams) (*ResetApiLimitResponse, error)
	NewResetApiLimitParams() *ResetApiLimitParams
	UpdateResourceCount(p *UpdateResourceCountParams) (*UpdateResourceCountResponse, error)
	NewUpdateResourceCountParams(domainid string) *UpdateResourceCountParams
	UpdateResourceLimit(p *UpdateResourceLimitParams) (*UpdateResourceLimitResponse, error)
	NewUpdateResourceLimitParams(resourcetype int) *UpdateResourceLimitParams
}

type GetApiLimitParams struct {
	p map[string]interface{}
}

func (p *GetApiLimitParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new GetApiLimitParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewGetApiLimitParams() *GetApiLimitParams {
	p := &GetApiLimitParams{}
	p.p = make(map[string]interface{})
	return p
}

// Get API limit count for the caller
func (s *LimitService) GetApiLimit(p *GetApiLimitParams) (*GetApiLimitResponse, error) {
	resp, err := s.cs.newRequest("getApiLimit", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetApiLimitResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetApiLimitResponse struct {
	Account     string `json:"account"`
	Accountid   string `json:"accountid"`
	ApiAllowed  int    `json:"apiAllowed"`
	ApiIssued   int    `json:"apiIssued"`
	ExpireAfter int64  `json:"expireAfter"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
}

type ListResourceLimitsParams struct {
	p map[string]interface{}
}

func (p *ListResourceLimitsParams) toURLValues() url.Values {
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
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
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
	if v, found := p.p["resourcetype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("resourcetype", vv)
	}
	if v, found := p.p["resourcetypename"]; found {
		u.Set("resourcetypename", v.(string))
	}
	if v, found := p.p["tag"]; found {
		u.Set("tag", v.(string))
	}
	return u
}

func (p *ListResourceLimitsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListResourceLimitsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListResourceLimitsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListResourceLimitsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListResourceLimitsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListResourceLimitsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListResourceLimitsParams) SetId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListResourceLimitsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListResourceLimitsParams) GetId() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int64)
	return value, ok
}

func (p *ListResourceLimitsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListResourceLimitsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListResourceLimitsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListResourceLimitsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListResourceLimitsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListResourceLimitsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListResourceLimitsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListResourceLimitsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListResourceLimitsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListResourceLimitsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListResourceLimitsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListResourceLimitsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListResourceLimitsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListResourceLimitsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListResourceLimitsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListResourceLimitsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListResourceLimitsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListResourceLimitsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListResourceLimitsParams) SetResourcetype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *ListResourceLimitsParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *ListResourceLimitsParams) GetResourcetype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(int)
	return value, ok
}

func (p *ListResourceLimitsParams) SetResourcetypename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetypename"] = v
}

func (p *ListResourceLimitsParams) ResetResourcetypename() {
	if p.p != nil && p.p["resourcetypename"] != nil {
		delete(p.p, "resourcetypename")
	}
}

func (p *ListResourceLimitsParams) GetResourcetypename() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetypename"].(string)
	return value, ok
}

func (p *ListResourceLimitsParams) SetTag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tag"] = v
}

func (p *ListResourceLimitsParams) ResetTag() {
	if p.p != nil && p.p["tag"] != nil {
		delete(p.p, "tag")
	}
}

func (p *ListResourceLimitsParams) GetTag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tag"].(string)
	return value, ok
}

// You should always use this function to get a new ListResourceLimitsParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewListResourceLimitsParams() *ListResourceLimitsParams {
	p := &ListResourceLimitsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists resource limits.
func (s *LimitService) ListResourceLimits(p *ListResourceLimitsParams) (*ListResourceLimitsResponse, error) {
	resp, err := s.cs.newRequest("listResourceLimits", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListResourceLimitsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListResourceLimitsResponse struct {
	Count          int              `json:"count"`
	ResourceLimits []*ResourceLimit `json:"resourcelimit"`
}

type ResourceLimit struct {
	Account          string `json:"account"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Domainpath       string `json:"domainpath"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Max              int64  `json:"max"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Resourcetype     string `json:"resourcetype"`
	Resourcetypename string `json:"resourcetypename"`
	Tag              string `json:"tag"`
}

type ResetApiLimitParams struct {
	p map[string]interface{}
}

func (p *ResetApiLimitParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	return u
}

func (p *ResetApiLimitParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ResetApiLimitParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ResetApiLimitParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

// You should always use this function to get a new ResetApiLimitParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewResetApiLimitParams() *ResetApiLimitParams {
	p := &ResetApiLimitParams{}
	p.p = make(map[string]interface{})
	return p
}

// Reset api count
func (s *LimitService) ResetApiLimit(p *ResetApiLimitParams) (*ResetApiLimitResponse, error) {
	resp, err := s.cs.newRequest("resetApiLimit", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetApiLimitResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ResetApiLimitResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ResetApiLimitResponse) UnmarshalJSON(b []byte) error {
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

	type alias ResetApiLimitResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateResourceCountParams struct {
	p map[string]interface{}
}

func (p *UpdateResourceCountParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("resourcetype", vv)
	}
	if v, found := p.p["tag"]; found {
		u.Set("tag", v.(string))
	}
	return u
}

func (p *UpdateResourceCountParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpdateResourceCountParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *UpdateResourceCountParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UpdateResourceCountParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpdateResourceCountParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *UpdateResourceCountParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpdateResourceCountParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *UpdateResourceCountParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *UpdateResourceCountParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *UpdateResourceCountParams) SetResourcetype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *UpdateResourceCountParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *UpdateResourceCountParams) GetResourcetype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(int)
	return value, ok
}

func (p *UpdateResourceCountParams) SetTag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tag"] = v
}

func (p *UpdateResourceCountParams) ResetTag() {
	if p.p != nil && p.p["tag"] != nil {
		delete(p.p, "tag")
	}
}

func (p *UpdateResourceCountParams) GetTag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tag"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateResourceCountParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewUpdateResourceCountParams(domainid string) *UpdateResourceCountParams {
	p := &UpdateResourceCountParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	return p
}

// Recalculate and update resource count for an account or domain. This also executes some cleanup tasks before calculating resource counts.
func (s *LimitService) UpdateResourceCount(p *UpdateResourceCountParams) (*UpdateResourceCountResponse, error) {
	resp, err := s.cs.newRequest("updateResourceCount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateResourceCountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateResourceCountResponse struct {
	Account          string `json:"account"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Domainpath       string `json:"domainpath"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Resourcecount    int64  `json:"resourcecount"`
	Resourcetype     string `json:"resourcetype"`
	Resourcetypename string `json:"resourcetypename"`
	Tag              string `json:"tag"`
}

type UpdateResourceLimitParams struct {
	p map[string]interface{}
}

func (p *UpdateResourceLimitParams) toURLValues() url.Values {
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
	if v, found := p.p["max"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("max", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("resourcetype", vv)
	}
	if v, found := p.p["tag"]; found {
		u.Set("tag", v.(string))
	}
	return u
}

func (p *UpdateResourceLimitParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpdateResourceLimitParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *UpdateResourceLimitParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UpdateResourceLimitParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpdateResourceLimitParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *UpdateResourceLimitParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpdateResourceLimitParams) SetMax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["max"] = v
}

func (p *UpdateResourceLimitParams) ResetMax() {
	if p.p != nil && p.p["max"] != nil {
		delete(p.p, "max")
	}
}

func (p *UpdateResourceLimitParams) GetMax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["max"].(int64)
	return value, ok
}

func (p *UpdateResourceLimitParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *UpdateResourceLimitParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *UpdateResourceLimitParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *UpdateResourceLimitParams) SetResourcetype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *UpdateResourceLimitParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *UpdateResourceLimitParams) GetResourcetype() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(int)
	return value, ok
}

func (p *UpdateResourceLimitParams) SetTag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tag"] = v
}

func (p *UpdateResourceLimitParams) ResetTag() {
	if p.p != nil && p.p["tag"] != nil {
		delete(p.p, "tag")
	}
}

func (p *UpdateResourceLimitParams) GetTag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tag"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateResourceLimitParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewUpdateResourceLimitParams(resourcetype int) *UpdateResourceLimitParams {
	p := &UpdateResourceLimitParams{}
	p.p = make(map[string]interface{})
	p.p["resourcetype"] = resourcetype
	return p
}

// Updates resource limits for an account or domain.
func (s *LimitService) UpdateResourceLimit(p *UpdateResourceLimitParams) (*UpdateResourceLimitResponse, error) {
	resp, err := s.cs.newRequest("updateResourceLimit", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateResourceLimitResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateResourceLimitResponse struct {
	Account          string `json:"account"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Domainpath       string `json:"domainpath"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Max              int64  `json:"max"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Resourcetype     string `json:"resourcetype"`
	Resourcetypename string `json:"resourcetypename"`
	Tag              string `json:"tag"`
}
