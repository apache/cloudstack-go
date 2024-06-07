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

type StoragePoolServiceIface interface {
	CancelStorageMaintenance(p *CancelStorageMaintenanceParams) (*CancelStorageMaintenanceResponse, error)
	NewCancelStorageMaintenanceParams(id string) *CancelStorageMaintenanceParams
	EnableStorageMaintenance(p *EnableStorageMaintenanceParams) (*EnableStorageMaintenanceResponse, error)
	NewEnableStorageMaintenanceParams(id string) *EnableStorageMaintenanceParams
	ListStorageProviders(p *ListStorageProvidersParams) (*ListStorageProvidersResponse, error)
	NewListStorageProvidersParams(storagePoolType string) *ListStorageProvidersParams
	ListObjectStoragePools(p *ListObjectStoragePoolsParams) (*ListObjectStoragePoolsResponse, error)
	NewListObjectStoragePoolsParams() *ListObjectStoragePoolsParams
	GetObjectStoragePoolID(name string, opts ...OptionFunc) (string, int, error)
	GetObjectStoragePoolByName(name string, opts ...OptionFunc) (*ObjectStoragePool, int, error)
	GetObjectStoragePoolByID(id string, opts ...OptionFunc) (*ObjectStoragePool, int, error)
	ListStoragePoolObjects(p *ListStoragePoolObjectsParams) (*ListStoragePoolObjectsResponse, error)
	NewListStoragePoolObjectsParams(id string) *ListStoragePoolObjectsParams
	GetStoragePoolObjectByID(id string, opts ...OptionFunc) (*StoragePoolObject, int, error)
}

type CancelStorageMaintenanceParams struct {
	p map[string]interface{}
}

func (p *CancelStorageMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *CancelStorageMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *CancelStorageMaintenanceParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *CancelStorageMaintenanceParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new CancelStorageMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewCancelStorageMaintenanceParams(id string) *CancelStorageMaintenanceParams {
	p := &CancelStorageMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Cancels maintenance for primary storage
func (s *StoragePoolService) CancelStorageMaintenance(p *CancelStorageMaintenanceParams) (*CancelStorageMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("cancelStorageMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelStorageMaintenanceResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CancelStorageMaintenanceResponse struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}

type EnableStorageMaintenanceParams struct {
	p map[string]interface{}
}

func (p *EnableStorageMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableStorageMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *EnableStorageMaintenanceParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *EnableStorageMaintenanceParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new EnableStorageMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewEnableStorageMaintenanceParams(id string) *EnableStorageMaintenanceParams {
	p := &EnableStorageMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Puts storage pool into maintenance state
func (s *StoragePoolService) EnableStorageMaintenance(p *EnableStorageMaintenanceParams) (*EnableStorageMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("enableStorageMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableStorageMaintenanceResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type EnableStorageMaintenanceResponse struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}

type ListStorageProvidersParams struct {
	p map[string]interface{}
}

func (p *ListStorageProvidersParams) toURLValues() url.Values {
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
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *ListStorageProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStorageProvidersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListStorageProvidersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStorageProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStorageProvidersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListStorageProvidersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStorageProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStorageProvidersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListStorageProvidersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListStorageProvidersParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListStorageProvidersParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ListStorageProvidersParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new ListStorageProvidersParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStorageProvidersParams(storagePoolType string) *ListStorageProvidersParams {
	p := &ListStorageProvidersParams{}
	p.p = make(map[string]interface{})
	p.p["type"] = storagePoolType
	return p
}

// Lists storage providers.
func (s *StoragePoolService) ListStorageProviders(p *ListStorageProvidersParams) (*ListStorageProvidersResponse, error) {
	resp, err := s.cs.newRequest("listStorageProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStorageProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStorageProvidersResponse struct {
	Count            int                `json:"count"`
	StorageProviders []*StorageProvider `json:"storageprovider"`
}

type StorageProvider struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type ListObjectStoragePoolsParams struct {
	p map[string]interface{}
}

func (p *ListObjectStoragePoolsParams) toURLValues() url.Values {
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
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (p *ListObjectStoragePoolsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListObjectStoragePoolsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListObjectStoragePoolsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListObjectStoragePoolsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListObjectStoragePoolsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListObjectStoragePoolsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListObjectStoragePoolsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListObjectStoragePoolsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListObjectStoragePoolsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListObjectStoragePoolsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListObjectStoragePoolsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListObjectStoragePoolsParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListObjectStoragePoolsParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

// You should always use this function to get a new ListObjectStoragePoolsParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListObjectStoragePoolsParams() *ListObjectStoragePoolsParams {
	p := &ListObjectStoragePoolsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetObjectStoragePoolID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListObjectStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListObjectStoragePools(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ObjectStoragePools[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ObjectStoragePools {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetObjectStoragePoolByName(name string, opts ...OptionFunc) (*ObjectStoragePool, int, error) {
	id, count, err := s.GetObjectStoragePoolID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetObjectStoragePoolByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetObjectStoragePoolByID(id string, opts ...OptionFunc) (*ObjectStoragePool, int, error) {
	p := &ListObjectStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListObjectStoragePools(p)
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
		return l.ObjectStoragePools[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ObjectStoragePool UUID: %s!", id)
}

// Lists object storage pools.
func (s *StoragePoolService) ListObjectStoragePools(p *ListObjectStoragePoolsParams) (*ListObjectStoragePoolsResponse, error) {
	resp, err := s.cs.newRequest("listObjectStoragePools", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListObjectStoragePoolsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListObjectStoragePoolsResponse struct {
	Count              int                  `json:"count"`
	ObjectStoragePools []*ObjectStoragePool `json:"objectstore"`
}

type ObjectStoragePool struct {
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Providername   string `json:"providername"`
	Storagetotal   int64  `json:"storagetotal"`
	Storageused    int64  `json:"storageused"`
	Url            string `json:"url"`
}

type ListStoragePoolObjectsParams struct {
	p map[string]interface{}
}

func (p *ListStoragePoolObjectsParams) toURLValues() url.Values {
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
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	return u
}

func (p *ListStoragePoolObjectsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListStoragePoolObjectsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListStoragePoolObjectsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListStoragePoolObjectsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStoragePoolObjectsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListStoragePoolObjectsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStoragePoolObjectsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStoragePoolObjectsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListStoragePoolObjectsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStoragePoolObjectsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStoragePoolObjectsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListStoragePoolObjectsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListStoragePoolObjectsParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *ListStoragePoolObjectsParams) ResetPath() {
	if p.p != nil && p.p["path"] != nil {
		delete(p.p, "path")
	}
}

func (p *ListStoragePoolObjectsParams) GetPath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["path"].(string)
	return value, ok
}

// You should always use this function to get a new ListStoragePoolObjectsParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStoragePoolObjectsParams(id string) *ListStoragePoolObjectsParams {
	p := &ListStoragePoolObjectsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolObjectByID(id string, opts ...OptionFunc) (*StoragePoolObject, int, error) {
	p := &ListStoragePoolObjectsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStoragePoolObjects(p)
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
		return l.StoragePoolObjects[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for StoragePoolObject UUID: %s!", id)
}

// Lists objects at specified path on a storage pool.
func (s *StoragePoolService) ListStoragePoolObjects(p *ListStoragePoolObjectsParams) (*ListStoragePoolObjectsResponse, error) {
	resp, err := s.cs.newRequest("listStoragePoolObjects", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStoragePoolObjectsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStoragePoolObjectsResponse struct {
	Count              int                  `json:"count"`
	StoragePoolObjects []*StoragePoolObject `json:"datastoreobject"`
}

type StoragePoolObject struct {
	Format      string `json:"format"`
	Isdirectory bool   `json:"isdirectory"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Lastupdated string `json:"lastupdated"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	Snapshotid  string `json:"snapshotid"`
	Templateid  string `json:"templateid"`
	Volumeid    string `json:"volumeid"`
}
