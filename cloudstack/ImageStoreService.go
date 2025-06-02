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

type ImageStoreServiceIface interface {
	AddImageStore(p *AddImageStoreParams) (*AddImageStoreResponse, error)
	NewAddImageStoreParams(provider string) *AddImageStoreParams
	AddImageStoreS3(p *AddImageStoreS3Params) (*AddImageStoreS3Response, error)
	NewAddImageStoreS3Params(accesskey string, bucket string, endpoint string, secretkey string) *AddImageStoreS3Params
	CreateSecondaryStagingStore(p *CreateSecondaryStagingStoreParams) (*CreateSecondaryStagingStoreResponse, error)
	NewCreateSecondaryStagingStoreParams(url string) *CreateSecondaryStagingStoreParams
	DeleteImageStore(p *DeleteImageStoreParams) (*DeleteImageStoreResponse, error)
	NewDeleteImageStoreParams(id string) *DeleteImageStoreParams
	DeleteSecondaryStagingStore(p *DeleteSecondaryStagingStoreParams) (*DeleteSecondaryStagingStoreResponse, error)
	NewDeleteSecondaryStagingStoreParams(id string) *DeleteSecondaryStagingStoreParams
	ListImageStores(p *ListImageStoresParams) (*ListImageStoresResponse, error)
	NewListImageStoresParams() *ListImageStoresParams
	GetImageStoreID(name string, opts ...OptionFunc) (string, int, error)
	GetImageStoreByName(name string, opts ...OptionFunc) (*ImageStore, int, error)
	GetImageStoreByID(id string, opts ...OptionFunc) (*ImageStore, int, error)
	ListSecondaryStagingStores(p *ListSecondaryStagingStoresParams) (*ListSecondaryStagingStoresResponse, error)
	NewListSecondaryStagingStoresParams() *ListSecondaryStagingStoresParams
	GetSecondaryStagingStoreID(name string, opts ...OptionFunc) (string, int, error)
	GetSecondaryStagingStoreByName(name string, opts ...OptionFunc) (*SecondaryStagingStore, int, error)
	GetSecondaryStagingStoreByID(id string, opts ...OptionFunc) (*SecondaryStagingStore, int, error)
	MigrateResourceToAnotherSecondaryStorage(p *MigrateResourceToAnotherSecondaryStorageParams) (*MigrateResourceToAnotherSecondaryStorageResponse, error)
	NewMigrateResourceToAnotherSecondaryStorageParams(destpool string, srcpool string) *MigrateResourceToAnotherSecondaryStorageParams
	UpdateCloudToUseObjectStore(p *UpdateCloudToUseObjectStoreParams) (*UpdateCloudToUseObjectStoreResponse, error)
	NewUpdateCloudToUseObjectStoreParams(provider string) *UpdateCloudToUseObjectStoreParams
	ListImageStoreObjects(p *ListImageStoreObjectsParams) (*ListImageStoreObjectsResponse, error)
	NewListImageStoreObjectsParams(id string) *ListImageStoreObjectsParams
	GetImageStoreObjectByID(id string, opts ...OptionFunc) (*ImageStoreObject, int, error)
	UpdateImageStore(p *UpdateImageStoreParams) (*UpdateImageStoreResponse, error)
	NewUpdateImageStoreParams(id string) *UpdateImageStoreParams
	DownloadImageStoreObject(p *DownloadImageStoreObjectParams) (*DownloadImageStoreObjectResponse, error)
	NewDownloadImageStoreObjectParams(id string) *DownloadImageStoreObjectParams
}

type AddImageStoreParams struct {
	p map[string]interface{}
}

func (p *AddImageStoreParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddImageStoreParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *AddImageStoreParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *AddImageStoreParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *AddImageStoreParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddImageStoreParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddImageStoreParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddImageStoreParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *AddImageStoreParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *AddImageStoreParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *AddImageStoreParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddImageStoreParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddImageStoreParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddImageStoreParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddImageStoreParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddImageStoreParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddImageStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewAddImageStoreParams(provider string) *AddImageStoreParams {
	p := &AddImageStoreParams{}
	p.p = make(map[string]interface{})
	p.p["provider"] = provider
	return p
}

// Adds backup image store.
func (s *ImageStoreService) AddImageStore(p *AddImageStoreParams) (*AddImageStoreResponse, error) {
	resp, err := s.cs.newPostRequest("addImageStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r AddImageStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddImageStoreResponse struct {
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

type AddImageStoreS3Params struct {
	p map[string]interface{}
}

func (p *AddImageStoreS3Params) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accesskey"]; found {
		u.Set("accesskey", v.(string))
	}
	if v, found := p.p["bucket"]; found {
		u.Set("bucket", v.(string))
	}
	if v, found := p.p["connectiontimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("connectiontimeout", vv)
	}
	if v, found := p.p["connectionttl"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("connectionttl", vv)
	}
	if v, found := p.p["endpoint"]; found {
		u.Set("endpoint", v.(string))
	}
	if v, found := p.p["maxerrorretry"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxerrorretry", vv)
	}
	if v, found := p.p["s3signer"]; found {
		u.Set("s3signer", v.(string))
	}
	if v, found := p.p["secretkey"]; found {
		u.Set("secretkey", v.(string))
	}
	if v, found := p.p["sockettimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sockettimeout", vv)
	}
	if v, found := p.p["usehttps"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("usehttps", vv)
	}
	if v, found := p.p["usetcpkeepalive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("usetcpkeepalive", vv)
	}
	return u
}

func (p *AddImageStoreS3Params) SetAccesskey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accesskey"] = v
}

func (p *AddImageStoreS3Params) ResetAccesskey() {
	if p.p != nil && p.p["accesskey"] != nil {
		delete(p.p, "accesskey")
	}
}

func (p *AddImageStoreS3Params) GetAccesskey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accesskey"].(string)
	return value, ok
}

func (p *AddImageStoreS3Params) SetBucket(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bucket"] = v
}

func (p *AddImageStoreS3Params) ResetBucket() {
	if p.p != nil && p.p["bucket"] != nil {
		delete(p.p, "bucket")
	}
}

func (p *AddImageStoreS3Params) GetBucket() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bucket"].(string)
	return value, ok
}

func (p *AddImageStoreS3Params) SetConnectiontimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["connectiontimeout"] = v
}

func (p *AddImageStoreS3Params) ResetConnectiontimeout() {
	if p.p != nil && p.p["connectiontimeout"] != nil {
		delete(p.p, "connectiontimeout")
	}
}

func (p *AddImageStoreS3Params) GetConnectiontimeout() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["connectiontimeout"].(int)
	return value, ok
}

func (p *AddImageStoreS3Params) SetConnectionttl(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["connectionttl"] = v
}

func (p *AddImageStoreS3Params) ResetConnectionttl() {
	if p.p != nil && p.p["connectionttl"] != nil {
		delete(p.p, "connectionttl")
	}
}

func (p *AddImageStoreS3Params) GetConnectionttl() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["connectionttl"].(int)
	return value, ok
}

func (p *AddImageStoreS3Params) SetEndpoint(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endpoint"] = v
}

func (p *AddImageStoreS3Params) ResetEndpoint() {
	if p.p != nil && p.p["endpoint"] != nil {
		delete(p.p, "endpoint")
	}
}

func (p *AddImageStoreS3Params) GetEndpoint() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endpoint"].(string)
	return value, ok
}

func (p *AddImageStoreS3Params) SetMaxerrorretry(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxerrorretry"] = v
}

func (p *AddImageStoreS3Params) ResetMaxerrorretry() {
	if p.p != nil && p.p["maxerrorretry"] != nil {
		delete(p.p, "maxerrorretry")
	}
}

func (p *AddImageStoreS3Params) GetMaxerrorretry() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxerrorretry"].(int)
	return value, ok
}

func (p *AddImageStoreS3Params) SetS3signer(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["s3signer"] = v
}

func (p *AddImageStoreS3Params) ResetS3signer() {
	if p.p != nil && p.p["s3signer"] != nil {
		delete(p.p, "s3signer")
	}
}

func (p *AddImageStoreS3Params) GetS3signer() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["s3signer"].(string)
	return value, ok
}

func (p *AddImageStoreS3Params) SetSecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secretkey"] = v
}

func (p *AddImageStoreS3Params) ResetSecretkey() {
	if p.p != nil && p.p["secretkey"] != nil {
		delete(p.p, "secretkey")
	}
}

func (p *AddImageStoreS3Params) GetSecretkey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["secretkey"].(string)
	return value, ok
}

func (p *AddImageStoreS3Params) SetSockettimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sockettimeout"] = v
}

func (p *AddImageStoreS3Params) ResetSockettimeout() {
	if p.p != nil && p.p["sockettimeout"] != nil {
		delete(p.p, "sockettimeout")
	}
}

func (p *AddImageStoreS3Params) GetSockettimeout() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sockettimeout"].(int)
	return value, ok
}

func (p *AddImageStoreS3Params) SetUsehttps(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usehttps"] = v
}

func (p *AddImageStoreS3Params) ResetUsehttps() {
	if p.p != nil && p.p["usehttps"] != nil {
		delete(p.p, "usehttps")
	}
}

func (p *AddImageStoreS3Params) GetUsehttps() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["usehttps"].(bool)
	return value, ok
}

func (p *AddImageStoreS3Params) SetUsetcpkeepalive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usetcpkeepalive"] = v
}

func (p *AddImageStoreS3Params) ResetUsetcpkeepalive() {
	if p.p != nil && p.p["usetcpkeepalive"] != nil {
		delete(p.p, "usetcpkeepalive")
	}
}

func (p *AddImageStoreS3Params) GetUsetcpkeepalive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["usetcpkeepalive"].(bool)
	return value, ok
}

// You should always use this function to get a new AddImageStoreS3Params instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewAddImageStoreS3Params(accesskey string, bucket string, endpoint string, secretkey string) *AddImageStoreS3Params {
	p := &AddImageStoreS3Params{}
	p.p = make(map[string]interface{})
	p.p["accesskey"] = accesskey
	p.p["bucket"] = bucket
	p.p["endpoint"] = endpoint
	p.p["secretkey"] = secretkey
	return p
}

// Adds S3 Image Store
func (s *ImageStoreService) AddImageStoreS3(p *AddImageStoreS3Params) (*AddImageStoreS3Response, error) {
	resp, err := s.cs.newPostRequest("addImageStoreS3", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddImageStoreS3Response
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddImageStoreS3Response struct {
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

type CreateSecondaryStagingStoreParams struct {
	p map[string]interface{}
}

func (p *CreateSecondaryStagingStoreParams) toURLValues() url.Values {
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
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateSecondaryStagingStoreParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateSecondaryStagingStoreParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *CreateSecondaryStagingStoreParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *CreateSecondaryStagingStoreParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *CreateSecondaryStagingStoreParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *CreateSecondaryStagingStoreParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *CreateSecondaryStagingStoreParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *CreateSecondaryStagingStoreParams) ResetScope() {
	if p.p != nil && p.p["scope"] != nil {
		delete(p.p, "scope")
	}
}

func (p *CreateSecondaryStagingStoreParams) GetScope() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scope"].(string)
	return value, ok
}

func (p *CreateSecondaryStagingStoreParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *CreateSecondaryStagingStoreParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *CreateSecondaryStagingStoreParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *CreateSecondaryStagingStoreParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateSecondaryStagingStoreParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateSecondaryStagingStoreParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSecondaryStagingStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewCreateSecondaryStagingStoreParams(url string) *CreateSecondaryStagingStoreParams {
	p := &CreateSecondaryStagingStoreParams{}
	p.p = make(map[string]interface{})
	p.p["url"] = url
	return p
}

// create secondary staging store.
func (s *ImageStoreService) CreateSecondaryStagingStore(p *CreateSecondaryStagingStoreParams) (*CreateSecondaryStagingStoreResponse, error) {
	resp, err := s.cs.newPostRequest("createSecondaryStagingStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSecondaryStagingStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateSecondaryStagingStoreResponse struct {
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

type DeleteImageStoreParams struct {
	p map[string]interface{}
}

func (p *DeleteImageStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteImageStoreParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteImageStoreParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteImageStoreParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteImageStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewDeleteImageStoreParams(id string) *DeleteImageStoreParams {
	p := &DeleteImageStoreParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an image store or Secondary Storage.
func (s *ImageStoreService) DeleteImageStore(p *DeleteImageStoreParams) (*DeleteImageStoreResponse, error) {
	resp, err := s.cs.newPostRequest("deleteImageStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteImageStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteImageStoreResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteImageStoreResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteImageStoreResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteSecondaryStagingStoreParams struct {
	p map[string]interface{}
}

func (p *DeleteSecondaryStagingStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteSecondaryStagingStoreParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteSecondaryStagingStoreParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteSecondaryStagingStoreParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSecondaryStagingStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewDeleteSecondaryStagingStoreParams(id string) *DeleteSecondaryStagingStoreParams {
	p := &DeleteSecondaryStagingStoreParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a secondary staging store .
func (s *ImageStoreService) DeleteSecondaryStagingStore(p *DeleteSecondaryStagingStoreParams) (*DeleteSecondaryStagingStoreResponse, error) {
	resp, err := s.cs.newPostRequest("deleteSecondaryStagingStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSecondaryStagingStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSecondaryStagingStoreResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSecondaryStagingStoreResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSecondaryStagingStoreResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListImageStoresParams struct {
	p map[string]interface{}
}

func (p *ListImageStoresParams) toURLValues() url.Values {
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
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["readonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("readonly", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListImageStoresParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListImageStoresParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListImageStoresParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListImageStoresParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListImageStoresParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListImageStoresParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListImageStoresParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListImageStoresParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListImageStoresParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListImageStoresParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListImageStoresParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListImageStoresParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListImageStoresParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListImageStoresParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListImageStoresParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListImageStoresParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *ListImageStoresParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *ListImageStoresParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *ListImageStoresParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListImageStoresParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListImageStoresParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *ListImageStoresParams) SetReadonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["readonly"] = v
}

func (p *ListImageStoresParams) ResetReadonly() {
	if p.p != nil && p.p["readonly"] != nil {
		delete(p.p, "readonly")
	}
}

func (p *ListImageStoresParams) GetReadonly() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["readonly"].(bool)
	return value, ok
}

func (p *ListImageStoresParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListImageStoresParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListImageStoresParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListImageStoresParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewListImageStoresParams() *ListImageStoresParams {
	p := &ListImageStoresParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListImageStoresParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListImageStores(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ImageStores[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ImageStores {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreByName(name string, opts ...OptionFunc) (*ImageStore, int, error) {
	id, count, err := s.GetImageStoreID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetImageStoreByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreByID(id string, opts ...OptionFunc) (*ImageStore, int, error) {
	p := &ListImageStoresParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListImageStores(p)
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
		return l.ImageStores[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ImageStore UUID: %s!", id)
}

// Lists image stores.
func (s *ImageStoreService) ListImageStores(p *ListImageStoresParams) (*ListImageStoresResponse, error) {
	resp, err := s.cs.newRequest("listImageStores", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListImageStoresResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListImageStoresResponse struct {
	Count       int           `json:"count"`
	ImageStores []*ImageStore `json:"imagestore"`
}

type ImageStore struct {
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

type ListSecondaryStagingStoresParams struct {
	p map[string]interface{}
}

func (p *ListSecondaryStagingStoresParams) toURLValues() url.Values {
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
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListSecondaryStagingStoresParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSecondaryStagingStoresParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListSecondaryStagingStoresParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListSecondaryStagingStoresParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSecondaryStagingStoresParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSecondaryStagingStoresParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSecondaryStagingStoresParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListSecondaryStagingStoresParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListSecondaryStagingStoresParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListSecondaryStagingStoresParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSecondaryStagingStoresParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSecondaryStagingStoresParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSecondaryStagingStoresParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSecondaryStagingStoresParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSecondaryStagingStoresParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListSecondaryStagingStoresParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *ListSecondaryStagingStoresParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *ListSecondaryStagingStoresParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *ListSecondaryStagingStoresParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListSecondaryStagingStoresParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListSecondaryStagingStoresParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *ListSecondaryStagingStoresParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListSecondaryStagingStoresParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListSecondaryStagingStoresParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSecondaryStagingStoresParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewListSecondaryStagingStoresParams() *ListSecondaryStagingStoresParams {
	p := &ListSecondaryStagingStoresParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListSecondaryStagingStoresParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSecondaryStagingStores(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SecondaryStagingStores[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SecondaryStagingStores {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreByName(name string, opts ...OptionFunc) (*SecondaryStagingStore, int, error) {
	id, count, err := s.GetSecondaryStagingStoreID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSecondaryStagingStoreByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreByID(id string, opts ...OptionFunc) (*SecondaryStagingStore, int, error) {
	p := &ListSecondaryStagingStoresParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSecondaryStagingStores(p)
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
		return l.SecondaryStagingStores[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SecondaryStagingStore UUID: %s!", id)
}

// Lists secondary staging stores.
func (s *ImageStoreService) ListSecondaryStagingStores(p *ListSecondaryStagingStoresParams) (*ListSecondaryStagingStoresResponse, error) {
	resp, err := s.cs.newRequest("listSecondaryStagingStores", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSecondaryStagingStoresResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSecondaryStagingStoresResponse struct {
	Count                  int                      `json:"count"`
	SecondaryStagingStores []*SecondaryStagingStore `json:"secondarystagingstore"`
}

type SecondaryStagingStore struct {
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

type MigrateResourceToAnotherSecondaryStorageParams struct {
	p map[string]interface{}
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["destpool"]; found {
		u.Set("destpool", v.(string))
	}
	if v, found := p.p["snapshots"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("snapshots", vv)
	}
	if v, found := p.p["srcpool"]; found {
		u.Set("srcpool", v.(string))
	}
	if v, found := p.p["templates"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("templates", vv)
	}
	return u
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) SetDestpool(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destpool"] = v
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) ResetDestpool() {
	if p.p != nil && p.p["destpool"] != nil {
		delete(p.p, "destpool")
	}
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) GetDestpool() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destpool"].(string)
	return value, ok
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) SetSnapshots(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshots"] = v
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) ResetSnapshots() {
	if p.p != nil && p.p["snapshots"] != nil {
		delete(p.p, "snapshots")
	}
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) GetSnapshots() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["snapshots"].([]string)
	return value, ok
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) SetSrcpool(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["srcpool"] = v
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) ResetSrcpool() {
	if p.p != nil && p.p["srcpool"] != nil {
		delete(p.p, "srcpool")
	}
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) GetSrcpool() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["srcpool"].(string)
	return value, ok
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) SetTemplates(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templates"] = v
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) ResetTemplates() {
	if p.p != nil && p.p["templates"] != nil {
		delete(p.p, "templates")
	}
}

func (p *MigrateResourceToAnotherSecondaryStorageParams) GetTemplates() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templates"].([]string)
	return value, ok
}

// You should always use this function to get a new MigrateResourceToAnotherSecondaryStorageParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewMigrateResourceToAnotherSecondaryStorageParams(destpool string, srcpool string) *MigrateResourceToAnotherSecondaryStorageParams {
	p := &MigrateResourceToAnotherSecondaryStorageParams{}
	p.p = make(map[string]interface{})
	p.p["destpool"] = destpool
	p.p["srcpool"] = srcpool
	return p
}

// migrates resources from one secondary storage to destination image store
func (s *ImageStoreService) MigrateResourceToAnotherSecondaryStorage(p *MigrateResourceToAnotherSecondaryStorageParams) (*MigrateResourceToAnotherSecondaryStorageResponse, error) {
	resp, err := s.cs.newPostRequest("migrateResourceToAnotherSecondaryStorage", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateResourceToAnotherSecondaryStorageResponse
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

type MigrateResourceToAnotherSecondaryStorageResponse struct {
	JobID         string `json:"jobid"`
	Jobstatus     int    `json:"jobstatus"`
	Message       string `json:"message"`
	Migrationtype string `json:"migrationtype"`
	Success       bool   `json:"success"`
}

type UpdateCloudToUseObjectStoreParams struct {
	p map[string]interface{}
}

func (p *UpdateCloudToUseObjectStoreParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	return u
}

func (p *UpdateCloudToUseObjectStoreParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateCloudToUseObjectStoreParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateCloudToUseObjectStoreParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateCloudToUseObjectStoreParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateCloudToUseObjectStoreParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateCloudToUseObjectStoreParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateCloudToUseObjectStoreParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *UpdateCloudToUseObjectStoreParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *UpdateCloudToUseObjectStoreParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *UpdateCloudToUseObjectStoreParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *UpdateCloudToUseObjectStoreParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *UpdateCloudToUseObjectStoreParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateCloudToUseObjectStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewUpdateCloudToUseObjectStoreParams(provider string) *UpdateCloudToUseObjectStoreParams {
	p := &UpdateCloudToUseObjectStoreParams{}
	p.p = make(map[string]interface{})
	p.p["provider"] = provider
	return p
}

// Migrate current NFS secondary storages to use object store.
func (s *ImageStoreService) UpdateCloudToUseObjectStore(p *UpdateCloudToUseObjectStoreParams) (*UpdateCloudToUseObjectStoreResponse, error) {
	resp, err := s.cs.newPostRequest("updateCloudToUseObjectStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateCloudToUseObjectStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateCloudToUseObjectStoreResponse struct {
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

type ListImageStoreObjectsParams struct {
	p map[string]interface{}
}

func (p *ListImageStoreObjectsParams) toURLValues() url.Values {
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

func (p *ListImageStoreObjectsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListImageStoreObjectsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListImageStoreObjectsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListImageStoreObjectsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListImageStoreObjectsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListImageStoreObjectsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListImageStoreObjectsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListImageStoreObjectsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListImageStoreObjectsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListImageStoreObjectsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListImageStoreObjectsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListImageStoreObjectsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListImageStoreObjectsParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *ListImageStoreObjectsParams) ResetPath() {
	if p.p != nil && p.p["path"] != nil {
		delete(p.p, "path")
	}
}

func (p *ListImageStoreObjectsParams) GetPath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["path"].(string)
	return value, ok
}

// You should always use this function to get a new ListImageStoreObjectsParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewListImageStoreObjectsParams(id string) *ListImageStoreObjectsParams {
	p := &ListImageStoreObjectsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreObjectByID(id string, opts ...OptionFunc) (*ImageStoreObject, int, error) {
	p := &ListImageStoreObjectsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListImageStoreObjects(p)
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
		return l.ImageStoreObjects[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ImageStoreObject UUID: %s!", id)
}

// Lists objects at specified path on an image store.
func (s *ImageStoreService) ListImageStoreObjects(p *ListImageStoreObjectsParams) (*ListImageStoreObjectsResponse, error) {
	resp, err := s.cs.newRequest("listImageStoreObjects", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListImageStoreObjectsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListImageStoreObjectsResponse struct {
	Count             int                 `json:"count"`
	ImageStoreObjects []*ImageStoreObject `json:"datastoreobject"`
}

type ImageStoreObject struct {
	Format       string `json:"format"`
	Isdirectory  bool   `json:"isdirectory"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Lastupdated  string `json:"lastupdated"`
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	Snapshotid   string `json:"snapshotid"`
	Snapshotname string `json:"snapshotname"`
	Templateid   string `json:"templateid"`
	Templatename string `json:"templatename"`
	Volumeid     string `json:"volumeid"`
	Volumename   string `json:"volumename"`
}

type UpdateImageStoreParams struct {
	p map[string]interface{}
}

func (p *UpdateImageStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["capacitybytes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacitybytes", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["readonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("readonly", vv)
	}
	return u
}

func (p *UpdateImageStoreParams) SetCapacitybytes(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacitybytes"] = v
}

func (p *UpdateImageStoreParams) ResetCapacitybytes() {
	if p.p != nil && p.p["capacitybytes"] != nil {
		delete(p.p, "capacitybytes")
	}
}

func (p *UpdateImageStoreParams) GetCapacitybytes() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["capacitybytes"].(int64)
	return value, ok
}

func (p *UpdateImageStoreParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateImageStoreParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateImageStoreParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateImageStoreParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateImageStoreParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateImageStoreParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateImageStoreParams) SetReadonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["readonly"] = v
}

func (p *UpdateImageStoreParams) ResetReadonly() {
	if p.p != nil && p.p["readonly"] != nil {
		delete(p.p, "readonly")
	}
}

func (p *UpdateImageStoreParams) GetReadonly() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["readonly"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateImageStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewUpdateImageStoreParams(id string) *UpdateImageStoreParams {
	p := &UpdateImageStoreParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates image store read-only status
func (s *ImageStoreService) UpdateImageStore(p *UpdateImageStoreParams) (*UpdateImageStoreResponse, error) {
	resp, err := s.cs.newPostRequest("updateImageStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response UpdateImageStoreResponse `json:"imagestore"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type UpdateImageStoreResponse struct {
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

type DownloadImageStoreObjectParams struct {
	p map[string]interface{}
}

func (p *DownloadImageStoreObjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	return u
}

func (p *DownloadImageStoreObjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DownloadImageStoreObjectParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DownloadImageStoreObjectParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DownloadImageStoreObjectParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *DownloadImageStoreObjectParams) ResetPath() {
	if p.p != nil && p.p["path"] != nil {
		delete(p.p, "path")
	}
}

func (p *DownloadImageStoreObjectParams) GetPath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["path"].(string)
	return value, ok
}

// You should always use this function to get a new DownloadImageStoreObjectParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewDownloadImageStoreObjectParams(id string) *DownloadImageStoreObjectParams {
	p := &DownloadImageStoreObjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Download object at a specified path on an image store.
func (s *ImageStoreService) DownloadImageStoreObject(p *DownloadImageStoreObjectParams) (*DownloadImageStoreObjectResponse, error) {
	resp, err := s.cs.newPostRequest("downloadImageStoreObject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DownloadImageStoreObjectResponse
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

type DownloadImageStoreObjectResponse struct {
	Accountid        string `json:"accountid"`
	Created          string `json:"created"`
	ExtractId        string `json:"extractId"`
	ExtractMode      string `json:"extractMode"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Resultstring     string `json:"resultstring"`
	State            string `json:"state"`
	Status           string `json:"status"`
	Storagetype      string `json:"storagetype"`
	Uploadpercentage int    `json:"uploadpercentage"`
	Url              string `json:"url"`
	Zoneid           string `json:"zoneid"`
	Zonename         string `json:"zonename"`
}
