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

type ResourcemetadataServiceIface interface {
	AddResourceDetail(p *AddResourceDetailParams) (*AddResourceDetailResponse, error)
	NewAddResourceDetailParams(details map[string]string, resourceid string, resourcetype string) *AddResourceDetailParams
	ListDetailOptions(p *ListDetailOptionsParams) (*ListDetailOptionsResponse, error)
	NewListDetailOptionsParams(resourcetype string) *ListDetailOptionsParams
	GetVolumeSnapshotDetails(p *GetVolumeSnapshotDetailsParams) (*GetVolumeSnapshotDetailsResponse, error)
	NewGetVolumeSnapshotDetailsParams(snapshotid string) *GetVolumeSnapshotDetailsParams
	ListResourceDetails(p *ListResourceDetailsParams) (*ListResourceDetailsResponse, error)
	NewListResourceDetailsParams(resourcetype string) *ListResourceDetailsParams
	RemoveResourceDetail(p *RemoveResourceDetailParams) (*RemoveResourceDetailResponse, error)
	NewRemoveResourceDetailParams(resourceid string, resourcetype string) *RemoveResourceDetailParams
}

type AddResourceDetailParams struct {
	p map[string]interface{}
}

func (p *AddResourceDetailParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), m[k])
		}
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *AddResourceDetailParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *AddResourceDetailParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *AddResourceDetailParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *AddResourceDetailParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *AddResourceDetailParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *AddResourceDetailParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *AddResourceDetailParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *AddResourceDetailParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *AddResourceDetailParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *AddResourceDetailParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *AddResourceDetailParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *AddResourceDetailParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new AddResourceDetailParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewAddResourceDetailParams(details map[string]string, resourceid string, resourcetype string) *AddResourceDetailParams {
	p := &AddResourceDetailParams{}
	p.p = make(map[string]interface{})
	p.p["details"] = details
	p.p["resourceid"] = resourceid
	p.p["resourcetype"] = resourcetype
	return p
}

// Adds detail for the Resource.
func (s *ResourcemetadataService) AddResourceDetail(p *AddResourceDetailParams) (*AddResourceDetailResponse, error) {
	resp, err := s.cs.newPostRequest("addResourceDetail", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddResourceDetailResponse
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

type AddResourceDetailResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListDetailOptionsParams struct {
	p map[string]interface{}
}

func (p *ListDetailOptionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *ListDetailOptionsParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *ListDetailOptionsParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *ListDetailOptionsParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *ListDetailOptionsParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *ListDetailOptionsParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *ListDetailOptionsParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new ListDetailOptionsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewListDetailOptionsParams(resourcetype string) *ListDetailOptionsParams {
	p := &ListDetailOptionsParams{}
	p.p = make(map[string]interface{})
	p.p["resourcetype"] = resourcetype
	return p
}

// Lists all possible details and their options for a resource type such as a VM or a template
func (s *ResourcemetadataService) ListDetailOptions(p *ListDetailOptionsParams) (*ListDetailOptionsResponse, error) {
	resp, err := s.cs.newRequest("listDetailOptions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDetailOptionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDetailOptionsResponse struct {
	Count         int             `json:"count"`
	DetailOptions []*DetailOption `json:"detailoption"`
}

type DetailOption struct {
	Details   map[string]string `json:"details"`
	JobID     string            `json:"jobid"`
	Jobstatus int               `json:"jobstatus"`
}

type GetVolumeSnapshotDetailsParams struct {
	p map[string]interface{}
}

func (p *GetVolumeSnapshotDetailsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	return u
}

func (p *GetVolumeSnapshotDetailsParams) SetSnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshotid"] = v
}

func (p *GetVolumeSnapshotDetailsParams) ResetSnapshotid() {
	if p.p != nil && p.p["snapshotid"] != nil {
		delete(p.p, "snapshotid")
	}
}

func (p *GetVolumeSnapshotDetailsParams) GetSnapshotid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["snapshotid"].(string)
	return value, ok
}

// You should always use this function to get a new GetVolumeSnapshotDetailsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewGetVolumeSnapshotDetailsParams(snapshotid string) *GetVolumeSnapshotDetailsParams {
	p := &GetVolumeSnapshotDetailsParams{}
	p.p = make(map[string]interface{})
	p.p["snapshotid"] = snapshotid
	return p
}

// Get Volume Snapshot Details
func (s *ResourcemetadataService) GetVolumeSnapshotDetails(p *GetVolumeSnapshotDetailsParams) (*GetVolumeSnapshotDetailsResponse, error) {
	resp, err := s.cs.newRequest("getVolumeSnapshotDetails", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetVolumeSnapshotDetailsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetVolumeSnapshotDetailsResponse struct {
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	VolumeiScsiName string `json:"volumeiScsiName"`
}

type ListResourceDetailsParams struct {
	p map[string]interface{}
}

func (p *ListResourceDetailsParams) toURLValues() url.Values {
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
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["key"]; found {
		u.Set("key", v.(string))
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
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := p.p["value"]; found {
		u.Set("value", v.(string))
	}
	return u
}

func (p *ListResourceDetailsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListResourceDetailsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListResourceDetailsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListResourceDetailsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListResourceDetailsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListResourceDetailsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListResourceDetailsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListResourceDetailsParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListResourceDetailsParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListResourceDetailsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListResourceDetailsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListResourceDetailsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListResourceDetailsParams) SetKey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["key"] = v
}

func (p *ListResourceDetailsParams) ResetKey() {
	if p.p != nil && p.p["key"] != nil {
		delete(p.p, "key")
	}
}

func (p *ListResourceDetailsParams) GetKey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["key"].(string)
	return value, ok
}

func (p *ListResourceDetailsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListResourceDetailsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListResourceDetailsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListResourceDetailsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListResourceDetailsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListResourceDetailsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListResourceDetailsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListResourceDetailsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListResourceDetailsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListResourceDetailsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListResourceDetailsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListResourceDetailsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListResourceDetailsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListResourceDetailsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListResourceDetailsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListResourceDetailsParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *ListResourceDetailsParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *ListResourceDetailsParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *ListResourceDetailsParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *ListResourceDetailsParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *ListResourceDetailsParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

func (p *ListResourceDetailsParams) SetValue(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["value"] = v
}

func (p *ListResourceDetailsParams) ResetValue() {
	if p.p != nil && p.p["value"] != nil {
		delete(p.p, "value")
	}
}

func (p *ListResourceDetailsParams) GetValue() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["value"].(string)
	return value, ok
}

// You should always use this function to get a new ListResourceDetailsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewListResourceDetailsParams(resourcetype string) *ListResourceDetailsParams {
	p := &ListResourceDetailsParams{}
	p.p = make(map[string]interface{})
	p.p["resourcetype"] = resourcetype
	return p
}

// List resource detail(s)
func (s *ResourcemetadataService) ListResourceDetails(p *ListResourceDetailsParams) (*ListResourceDetailsResponse, error) {
	resp, err := s.cs.newRequest("listResourceDetails", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListResourceDetailsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListResourceDetailsResponse struct {
	Count           int               `json:"count"`
	ResourceDetails []*ResourceDetail `json:"resourcedetail"`
}

type ResourceDetail struct {
	Account      string `json:"account"`
	Customer     string `json:"customer"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Domainpath   string `json:"domainpath"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Key          string `json:"key"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Resourceid   string `json:"resourceid"`
	Resourcetype string `json:"resourcetype"`
	Value        string `json:"value"`
}

type RemoveResourceDetailParams struct {
	p map[string]interface{}
}

func (p *RemoveResourceDetailParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["key"]; found {
		u.Set("key", v.(string))
	}
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *RemoveResourceDetailParams) SetKey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["key"] = v
}

func (p *RemoveResourceDetailParams) ResetKey() {
	if p.p != nil && p.p["key"] != nil {
		delete(p.p, "key")
	}
}

func (p *RemoveResourceDetailParams) GetKey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["key"].(string)
	return value, ok
}

func (p *RemoveResourceDetailParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *RemoveResourceDetailParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *RemoveResourceDetailParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *RemoveResourceDetailParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *RemoveResourceDetailParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *RemoveResourceDetailParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveResourceDetailParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewRemoveResourceDetailParams(resourceid string, resourcetype string) *RemoveResourceDetailParams {
	p := &RemoveResourceDetailParams{}
	p.p = make(map[string]interface{})
	p.p["resourceid"] = resourceid
	p.p["resourcetype"] = resourcetype
	return p
}

// Removes detail for the Resource.
func (s *ResourcemetadataService) RemoveResourceDetail(p *RemoveResourceDetailParams) (*RemoveResourceDetailResponse, error) {
	resp, err := s.cs.newPostRequest("removeResourceDetail", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveResourceDetailResponse
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

type RemoveResourceDetailResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
