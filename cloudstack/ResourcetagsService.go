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

type ResourcetagsServiceIface interface {
	CreateTags(p *CreateTagsParams) (*CreateTagsResponse, error)
	NewCreateTagsParams(resourceids []string, resourcetype string, tags map[string]string) *CreateTagsParams
	DeleteTags(p *DeleteTagsParams) (*DeleteTagsResponse, error)
	NewDeleteTagsParams(resourceids []string, resourcetype string) *DeleteTagsParams
	ListStorageTags(p *ListStorageTagsParams) (*ListStorageTagsResponse, error)
	NewListStorageTagsParams() *ListStorageTagsParams
	GetStorageTagID(keyword string, opts ...OptionFunc) (string, int, error)
	ListTags(p *ListTagsParams) (*ListTagsResponse, error)
	NewListTagsParams() *ListTagsParams
}

type CreateTagsParams struct {
	p map[string]interface{}
}

func (p *CreateTagsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customer"]; found {
		u.Set("customer", v.(string))
	}
	if v, found := p.p["resourceids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("resourceids", vv)
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *CreateTagsParams) SetCustomer(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customer"] = v
}

func (p *CreateTagsParams) ResetCustomer() {
	if p.p != nil && p.p["customer"] != nil {
		delete(p.p, "customer")
	}
}

func (p *CreateTagsParams) GetCustomer() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customer"].(string)
	return value, ok
}

func (p *CreateTagsParams) SetResourceids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceids"] = v
}

func (p *CreateTagsParams) ResetResourceids() {
	if p.p != nil && p.p["resourceids"] != nil {
		delete(p.p, "resourceids")
	}
}

func (p *CreateTagsParams) GetResourceids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceids"].([]string)
	return value, ok
}

func (p *CreateTagsParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *CreateTagsParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *CreateTagsParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

func (p *CreateTagsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreateTagsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *CreateTagsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new CreateTagsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcetagsService) NewCreateTagsParams(resourceids []string, resourcetype string, tags map[string]string) *CreateTagsParams {
	p := &CreateTagsParams{}
	p.p = make(map[string]interface{})
	p.p["resourceids"] = resourceids
	p.p["resourcetype"] = resourcetype
	p.p["tags"] = tags
	return p
}

// Creates resource tag(s)
func (s *ResourcetagsService) CreateTags(p *CreateTagsParams) (*CreateTagsResponse, error) {
	resp, err := s.cs.newRequest("createTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTagsResponse
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

type CreateTagsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteTagsParams struct {
	p map[string]interface{}
}

func (p *DeleteTagsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["resourceids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("resourceids", vv)
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			if m[k] != "" {
				u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
			}
		}
	}
	return u
}

func (p *DeleteTagsParams) SetResourceids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceids"] = v
}

func (p *DeleteTagsParams) ResetResourceids() {
	if p.p != nil && p.p["resourceids"] != nil {
		delete(p.p, "resourceids")
	}
}

func (p *DeleteTagsParams) GetResourceids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceids"].([]string)
	return value, ok
}

func (p *DeleteTagsParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *DeleteTagsParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *DeleteTagsParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

func (p *DeleteTagsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *DeleteTagsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *DeleteTagsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new DeleteTagsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcetagsService) NewDeleteTagsParams(resourceids []string, resourcetype string) *DeleteTagsParams {
	p := &DeleteTagsParams{}
	p.p = make(map[string]interface{})
	p.p["resourceids"] = resourceids
	p.p["resourcetype"] = resourcetype
	return p
}

// Deleting resource tag(s)
func (s *ResourcetagsService) DeleteTags(p *DeleteTagsParams) (*DeleteTagsResponse, error) {
	resp, err := s.cs.newRequest("deleteTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTagsResponse
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

type DeleteTagsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListStorageTagsParams struct {
	p map[string]interface{}
}

func (p *ListStorageTagsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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

func (p *ListStorageTagsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStorageTagsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListStorageTagsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStorageTagsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStorageTagsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListStorageTagsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStorageTagsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStorageTagsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListStorageTagsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListStorageTagsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcetagsService) NewListStorageTagsParams() *ListStorageTagsParams {
	p := &ListStorageTagsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ResourcetagsService) GetStorageTagID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListStorageTagsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListStorageTags(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.StorageTags[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.StorageTags {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists storage tags
func (s *ResourcetagsService) ListStorageTags(p *ListStorageTagsParams) (*ListStorageTagsResponse, error) {
	resp, err := s.cs.newRequest("listStorageTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStorageTagsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStorageTagsResponse struct {
	Count       int           `json:"count"`
	StorageTags []*StorageTag `json:"storagetag"`
}

type StorageTag struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Poolid    int64  `json:"poolid"`
}

type ListTagsParams struct {
	p map[string]interface{}
}

func (p *ListTagsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["customer"]; found {
		u.Set("customer", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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

func (p *ListTagsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListTagsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListTagsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListTagsParams) SetCustomer(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customer"] = v
}

func (p *ListTagsParams) ResetCustomer() {
	if p.p != nil && p.p["customer"] != nil {
		delete(p.p, "customer")
	}
}

func (p *ListTagsParams) GetCustomer() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customer"].(string)
	return value, ok
}

func (p *ListTagsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListTagsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListTagsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListTagsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListTagsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListTagsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListTagsParams) SetKey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["key"] = v
}

func (p *ListTagsParams) ResetKey() {
	if p.p != nil && p.p["key"] != nil {
		delete(p.p, "key")
	}
}

func (p *ListTagsParams) GetKey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["key"].(string)
	return value, ok
}

func (p *ListTagsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTagsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTagsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTagsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListTagsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListTagsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListTagsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTagsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTagsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTagsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTagsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTagsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTagsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListTagsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListTagsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListTagsParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *ListTagsParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *ListTagsParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *ListTagsParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *ListTagsParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *ListTagsParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

func (p *ListTagsParams) SetValue(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["value"] = v
}

func (p *ListTagsParams) ResetValue() {
	if p.p != nil && p.p["value"] != nil {
		delete(p.p, "value")
	}
}

func (p *ListTagsParams) GetValue() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["value"].(string)
	return value, ok
}

// You should always use this function to get a new ListTagsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcetagsService) NewListTagsParams() *ListTagsParams {
	p := &ListTagsParams{}
	p.p = make(map[string]interface{})
	return p
}

// List resource tag(s)
func (s *ResourcetagsService) ListTags(p *ListTagsParams) (*ListTagsResponse, error) {
	resp, err := s.cs.newRequest("listTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTagsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTagsResponse struct {
	Count int    `json:"count"`
	Tags  []*Tag `json:"tag"`
}

type Tag struct {
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
