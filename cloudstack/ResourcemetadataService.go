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
	resp, err := s.cs.newRequest("addResourceDetail", p.toURLValues())
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
	resp, err := s.cs.newRequest("removeResourceDetail", p.toURLValues())
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
