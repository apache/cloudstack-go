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

type ObjectStoreServiceIface interface {
	CreateBucket(p *CreateBucketParams) (*CreateBucketResponse, error)
	NewCreateBucketParams(name string, objectstorageid string) *CreateBucketParams
	DeleteBucket(p *DeleteBucketParams) (*DeleteBucketResponse, error)
	NewDeleteBucketParams(id string) *DeleteBucketParams
	UpdateBucket(p *UpdateBucketParams) (*UpdateBucketResponse, error)
	NewUpdateBucketParams(id string) *UpdateBucketParams
	ListBuckets(p *ListBucketsParams) (*ListBucketsResponse, error)
	NewListBucketsParams() *ListBucketsParams
	GetBucketID(name string, opts ...OptionFunc) (string, int, error)
	GetBucketByName(name string, opts ...OptionFunc) (*Bucket, int, error)
	GetBucketByID(id string, opts ...OptionFunc) (*Bucket, int, error)
}

type CreateBucketParams struct {
	p map[string]interface{}
}

func (p *CreateBucketParams) toURLValues() url.Values {
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
	if v, found := p.p["encryption"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("encryption", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["objectlocking"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("objectlocking", vv)
	}
	if v, found := p.p["objectstorageid"]; found {
		u.Set("objectstorageid", v.(string))
	}
	if v, found := p.p["policy"]; found {
		u.Set("policy", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["quota"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("quota", vv)
	}
	if v, found := p.p["versioning"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("versioning", vv)
	}
	return u
}

func (p *CreateBucketParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateBucketParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateBucketParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateBucketParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateBucketParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateBucketParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateBucketParams) SetEncryption(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["encryption"] = v
}

func (p *CreateBucketParams) ResetEncryption() {
	if p.p != nil && p.p["encryption"] != nil {
		delete(p.p, "encryption")
	}
}

func (p *CreateBucketParams) GetEncryption() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["encryption"].(bool)
	return value, ok
}

func (p *CreateBucketParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateBucketParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateBucketParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateBucketParams) SetObjectlocking(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["objectlocking"] = v
}

func (p *CreateBucketParams) ResetObjectlocking() {
	if p.p != nil && p.p["objectlocking"] != nil {
		delete(p.p, "objectlocking")
	}
}

func (p *CreateBucketParams) GetObjectlocking() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["objectlocking"].(bool)
	return value, ok
}

func (p *CreateBucketParams) SetObjectstorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["objectstorageid"] = v
}

func (p *CreateBucketParams) ResetObjectstorageid() {
	if p.p != nil && p.p["objectstorageid"] != nil {
		delete(p.p, "objectstorageid")
	}
}

func (p *CreateBucketParams) GetObjectstorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["objectstorageid"].(string)
	return value, ok
}

func (p *CreateBucketParams) SetPolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policy"] = v
}

func (p *CreateBucketParams) ResetPolicy() {
	if p.p != nil && p.p["policy"] != nil {
		delete(p.p, "policy")
	}
}

func (p *CreateBucketParams) GetPolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policy"].(string)
	return value, ok
}

func (p *CreateBucketParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateBucketParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateBucketParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateBucketParams) SetQuota(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quota"] = v
}

func (p *CreateBucketParams) ResetQuota() {
	if p.p != nil && p.p["quota"] != nil {
		delete(p.p, "quota")
	}
}

func (p *CreateBucketParams) GetQuota() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quota"].(int)
	return value, ok
}

func (p *CreateBucketParams) SetVersioning(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["versioning"] = v
}

func (p *CreateBucketParams) ResetVersioning() {
	if p.p != nil && p.p["versioning"] != nil {
		delete(p.p, "versioning")
	}
}

func (p *CreateBucketParams) GetVersioning() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["versioning"].(bool)
	return value, ok
}

// You should always use this function to get a new CreateBucketParams instance,
// as then you are sure you have configured all required params
func (s *ObjectStoreService) NewCreateBucketParams(name string, objectstorageid string) *CreateBucketParams {
	p := &CreateBucketParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["objectstorageid"] = objectstorageid
	return p
}

// Creates a bucket in the specified object storage pool.
func (s *ObjectStoreService) CreateBucket(p *CreateBucketParams) (*CreateBucketResponse, error) {
	resp, err := s.cs.newPostRequest("createBucket", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateBucketResponse
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

type CreateBucketResponse struct {
	Accesskey       string `json:"accesskey"`
	Account         string `json:"account"`
	Created         string `json:"created"`
	Domain          string `json:"domain"`
	Domainid        string `json:"domainid"`
	Domainpath      string `json:"domainpath"`
	Encryption      bool   `json:"encryption"`
	Hasannotations  bool   `json:"hasannotations"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Name            string `json:"name"`
	Objectlocking   bool   `json:"objectlocking"`
	Objectstorageid string `json:"objectstorageid"`
	Objectstore     string `json:"objectstore"`
	Policy          string `json:"policy"`
	Project         string `json:"project"`
	Projectid       string `json:"projectid"`
	Provider        string `json:"provider"`
	Quota           int    `json:"quota"`
	Size            int64  `json:"size"`
	State           string `json:"state"`
	Tags            []Tags `json:"tags"`
	Url             string `json:"url"`
	Usersecretkey   string `json:"usersecretkey"`
	Versioning      bool   `json:"versioning"`
}

type DeleteBucketParams struct {
	p map[string]interface{}
}

func (p *DeleteBucketParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteBucketParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteBucketParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteBucketParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBucketParams instance,
// as then you are sure you have configured all required params
func (s *ObjectStoreService) NewDeleteBucketParams(id string) *DeleteBucketParams {
	p := &DeleteBucketParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an empty Bucket.
func (s *ObjectStoreService) DeleteBucket(p *DeleteBucketParams) (*DeleteBucketResponse, error) {
	resp, err := s.cs.newPostRequest("deleteBucket", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBucketResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteBucketResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteBucketResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteBucketResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateBucketParams struct {
	p map[string]interface{}
}

func (p *UpdateBucketParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["encryption"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("encryption", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["policy"]; found {
		u.Set("policy", v.(string))
	}
	if v, found := p.p["quota"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("quota", vv)
	}
	if v, found := p.p["versioning"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("versioning", vv)
	}
	return u
}

func (p *UpdateBucketParams) SetEncryption(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["encryption"] = v
}

func (p *UpdateBucketParams) ResetEncryption() {
	if p.p != nil && p.p["encryption"] != nil {
		delete(p.p, "encryption")
	}
}

func (p *UpdateBucketParams) GetEncryption() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["encryption"].(bool)
	return value, ok
}

func (p *UpdateBucketParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateBucketParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateBucketParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateBucketParams) SetPolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policy"] = v
}

func (p *UpdateBucketParams) ResetPolicy() {
	if p.p != nil && p.p["policy"] != nil {
		delete(p.p, "policy")
	}
}

func (p *UpdateBucketParams) GetPolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policy"].(string)
	return value, ok
}

func (p *UpdateBucketParams) SetQuota(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quota"] = v
}

func (p *UpdateBucketParams) ResetQuota() {
	if p.p != nil && p.p["quota"] != nil {
		delete(p.p, "quota")
	}
}

func (p *UpdateBucketParams) GetQuota() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quota"].(int)
	return value, ok
}

func (p *UpdateBucketParams) SetVersioning(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["versioning"] = v
}

func (p *UpdateBucketParams) ResetVersioning() {
	if p.p != nil && p.p["versioning"] != nil {
		delete(p.p, "versioning")
	}
}

func (p *UpdateBucketParams) GetVersioning() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["versioning"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateBucketParams instance,
// as then you are sure you have configured all required params
func (s *ObjectStoreService) NewUpdateBucketParams(id string) *UpdateBucketParams {
	p := &UpdateBucketParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates Bucket properties
func (s *ObjectStoreService) UpdateBucket(p *UpdateBucketParams) (*UpdateBucketResponse, error) {
	resp, err := s.cs.newPostRequest("updateBucket", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateBucketResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateBucketResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateBucketResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateBucketResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListBucketsParams struct {
	p map[string]interface{}
}

func (p *ListBucketsParams) toURLValues() url.Values {
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
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
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
	if v, found := p.p["objectstorageid"]; found {
		u.Set("objectstorageid", v.(string))
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
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListBucketsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListBucketsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListBucketsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListBucketsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListBucketsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListBucketsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListBucketsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListBucketsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListBucketsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListBucketsParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListBucketsParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListBucketsParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListBucketsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListBucketsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListBucketsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListBucketsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBucketsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBucketsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBucketsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListBucketsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListBucketsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListBucketsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListBucketsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListBucketsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListBucketsParams) SetObjectstorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["objectstorageid"] = v
}

func (p *ListBucketsParams) ResetObjectstorageid() {
	if p.p != nil && p.p["objectstorageid"] != nil {
		delete(p.p, "objectstorageid")
	}
}

func (p *ListBucketsParams) GetObjectstorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["objectstorageid"].(string)
	return value, ok
}

func (p *ListBucketsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBucketsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBucketsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBucketsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBucketsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBucketsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBucketsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListBucketsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListBucketsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListBucketsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListBucketsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListBucketsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListBucketsParams instance,
// as then you are sure you have configured all required params
func (s *ObjectStoreService) NewListBucketsParams() *ListBucketsParams {
	p := &ListBucketsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ObjectStoreService) GetBucketID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListBucketsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListBuckets(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Buckets[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Buckets {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ObjectStoreService) GetBucketByName(name string, opts ...OptionFunc) (*Bucket, int, error) {
	id, count, err := s.GetBucketID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetBucketByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ObjectStoreService) GetBucketByID(id string, opts ...OptionFunc) (*Bucket, int, error) {
	p := &ListBucketsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListBuckets(p)
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
		return l.Buckets[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Bucket UUID: %s!", id)
}

// Lists all Buckets.
func (s *ObjectStoreService) ListBuckets(p *ListBucketsParams) (*ListBucketsResponse, error) {
	resp, err := s.cs.newRequest("listBuckets", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBucketsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBucketsResponse struct {
	Count   int       `json:"count"`
	Buckets []*Bucket `json:"bucket"`
}

type Bucket struct {
	Accesskey       string `json:"accesskey"`
	Account         string `json:"account"`
	Created         string `json:"created"`
	Domain          string `json:"domain"`
	Domainid        string `json:"domainid"`
	Domainpath      string `json:"domainpath"`
	Encryption      bool   `json:"encryption"`
	Hasannotations  bool   `json:"hasannotations"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Name            string `json:"name"`
	Objectlocking   bool   `json:"objectlocking"`
	Objectstorageid string `json:"objectstorageid"`
	Objectstore     string `json:"objectstore"`
	Policy          string `json:"policy"`
	Project         string `json:"project"`
	Projectid       string `json:"projectid"`
	Provider        string `json:"provider"`
	Quota           int    `json:"quota"`
	Size            int64  `json:"size"`
	State           string `json:"state"`
	Tags            []Tags `json:"tags"`
	Url             string `json:"url"`
	Usersecretkey   string `json:"usersecretkey"`
	Versioning      bool   `json:"versioning"`
}
