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

type SharedFileSystemServiceIface interface {
	ChangeSharedFileSystemDiskOffering(p *ChangeSharedFileSystemDiskOfferingParams) (*ChangeSharedFileSystemDiskOfferingResponse, error)
	NewChangeSharedFileSystemDiskOfferingParams(id string) *ChangeSharedFileSystemDiskOfferingParams
	ChangeSharedFileSystemServiceOffering(p *ChangeSharedFileSystemServiceOfferingParams) (*ChangeSharedFileSystemServiceOfferingResponse, error)
	NewChangeSharedFileSystemServiceOfferingParams(id string, serviceofferingid string) *ChangeSharedFileSystemServiceOfferingParams
	CreateSharedFileSystem(p *CreateSharedFileSystemParams) (*CreateSharedFileSystemResponse, error)
	NewCreateSharedFileSystemParams(diskofferingid string, filesystem string, name string, networkid string, serviceofferingid string, zoneid string) *CreateSharedFileSystemParams
	DestroySharedFileSystem(p *DestroySharedFileSystemParams) (*DestroySharedFileSystemResponse, error)
	NewDestroySharedFileSystemParams() *DestroySharedFileSystemParams
	ExpungeSharedFileSystem(p *ExpungeSharedFileSystemParams) (*ExpungeSharedFileSystemResponse, error)
	NewExpungeSharedFileSystemParams() *ExpungeSharedFileSystemParams
	ListSharedFileSystemProviders(p *ListSharedFileSystemProvidersParams) (*ListSharedFileSystemProvidersResponse, error)
	NewListSharedFileSystemProvidersParams() *ListSharedFileSystemProvidersParams
	ListSharedFileSystems(p *ListSharedFileSystemsParams) (*ListSharedFileSystemsResponse, error)
	NewListSharedFileSystemsParams() *ListSharedFileSystemsParams
	GetSharedFileSystemID(name string, opts ...OptionFunc) (string, int, error)
	GetSharedFileSystemByName(name string, opts ...OptionFunc) (*SharedFileSystem, int, error)
	GetSharedFileSystemByID(id string, opts ...OptionFunc) (*SharedFileSystem, int, error)
	RecoverSharedFileSystem(p *RecoverSharedFileSystemParams) (*RecoverSharedFileSystemResponse, error)
	NewRecoverSharedFileSystemParams() *RecoverSharedFileSystemParams
	RestartSharedFileSystem(p *RestartSharedFileSystemParams) (*RestartSharedFileSystemResponse, error)
	NewRestartSharedFileSystemParams(id string) *RestartSharedFileSystemParams
	StartSharedFileSystem(p *StartSharedFileSystemParams) (*StartSharedFileSystemResponse, error)
	NewStartSharedFileSystemParams(id string) *StartSharedFileSystemParams
	StopSharedFileSystem(p *StopSharedFileSystemParams) (*StopSharedFileSystemResponse, error)
	NewStopSharedFileSystemParams(id string) *StopSharedFileSystemParams
	UpdateSharedFileSystem(p *UpdateSharedFileSystemParams) (*UpdateSharedFileSystemResponse, error)
	NewUpdateSharedFileSystemParams(id string) *UpdateSharedFileSystemParams
}

type ChangeSharedFileSystemDiskOfferingParams struct {
	p map[string]interface{}
}

func (p *ChangeSharedFileSystemDiskOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	return u
}

func (p *ChangeSharedFileSystemDiskOfferingParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *ChangeSharedFileSystemDiskOfferingParams) ResetDiskofferingid() {
	if p.p != nil && p.p["diskofferingid"] != nil {
		delete(p.p, "diskofferingid")
	}
}

func (p *ChangeSharedFileSystemDiskOfferingParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *ChangeSharedFileSystemDiskOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ChangeSharedFileSystemDiskOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ChangeSharedFileSystemDiskOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ChangeSharedFileSystemDiskOfferingParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *ChangeSharedFileSystemDiskOfferingParams) ResetMaxiops() {
	if p.p != nil && p.p["maxiops"] != nil {
		delete(p.p, "maxiops")
	}
}

func (p *ChangeSharedFileSystemDiskOfferingParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *ChangeSharedFileSystemDiskOfferingParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *ChangeSharedFileSystemDiskOfferingParams) ResetMiniops() {
	if p.p != nil && p.p["miniops"] != nil {
		delete(p.p, "miniops")
	}
}

func (p *ChangeSharedFileSystemDiskOfferingParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *ChangeSharedFileSystemDiskOfferingParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *ChangeSharedFileSystemDiskOfferingParams) ResetSize() {
	if p.p != nil && p.p["size"] != nil {
		delete(p.p, "size")
	}
}

func (p *ChangeSharedFileSystemDiskOfferingParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

// You should always use this function to get a new ChangeSharedFileSystemDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewChangeSharedFileSystemDiskOfferingParams(id string) *ChangeSharedFileSystemDiskOfferingParams {
	p := &ChangeSharedFileSystemDiskOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Change Disk offering of a Shared FileSystem
func (s *SharedFileSystemService) ChangeSharedFileSystemDiskOffering(p *ChangeSharedFileSystemDiskOfferingParams) (*ChangeSharedFileSystemDiskOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("changeSharedFileSystemDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeSharedFileSystemDiskOfferingResponse
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

type ChangeSharedFileSystemDiskOfferingResponse struct {
	Account                 string `json:"account"`
	Description             string `json:"description"`
	Diskioread              int64  `json:"diskioread"`
	Diskiowrite             int64  `json:"diskiowrite"`
	Diskkbsread             int64  `json:"diskkbsread"`
	Diskkbswrite            int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext"`
	Diskofferingid          string `json:"diskofferingid"`
	Diskofferingname        string `json:"diskofferingname"`
	Domain                  string `json:"domain"`
	Domainid                string `json:"domainid"`
	Domainpath              string `json:"domainpath"`
	Filesystem              string `json:"filesystem"`
	Hasannotations          bool   `json:"hasannotations"`
	Id                      string `json:"id"`
	Iscustomdiskoffering    bool   `json:"iscustomdiskoffering"`
	JobID                   string `json:"jobid"`
	Jobstatus               int    `json:"jobstatus"`
	Name                    string `json:"name"`
	Networkid               string `json:"networkid"`
	Networkname             string `json:"networkname"`
	Nic                     []Nic  `json:"nic"`
	Path                    string `json:"path"`
	Physicalsize            int64  `json:"physicalsize"`
	Project                 string `json:"project"`
	Projectid               string `json:"projectid"`
	Provider                string `json:"provider"`
	Provisioningtype        string `json:"provisioningtype"`
	Serviceofferingid       string `json:"serviceofferingid"`
	Serviceofferingname     string `json:"serviceofferingname"`
	Size                    int64  `json:"size"`
	Sizegb                  string `json:"sizegb"`
	State                   string `json:"state"`
	Storage                 string `json:"storage"`
	Storageid               string `json:"storageid"`
	Tags                    []Tags `json:"tags"`
	Utilization             string `json:"utilization"`
	Virtualmachineid        string `json:"virtualmachineid"`
	Virtualsize             int64  `json:"virtualsize"`
	Vmstate                 string `json:"vmstate"`
	Volumeid                string `json:"volumeid"`
	Volumename              string `json:"volumename"`
	Zoneid                  string `json:"zoneid"`
	Zonename                string `json:"zonename"`
}

type ChangeSharedFileSystemServiceOfferingParams struct {
	p map[string]interface{}
}

func (p *ChangeSharedFileSystemServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ChangeSharedFileSystemServiceOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ChangeSharedFileSystemServiceOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ChangeSharedFileSystemServiceOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ChangeSharedFileSystemServiceOfferingParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ChangeSharedFileSystemServiceOfferingParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ChangeSharedFileSystemServiceOfferingParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeSharedFileSystemServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewChangeSharedFileSystemServiceOfferingParams(id string, serviceofferingid string) *ChangeSharedFileSystemServiceOfferingParams {
	p := &ChangeSharedFileSystemServiceOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Change Service offering of a Shared FileSystem
func (s *SharedFileSystemService) ChangeSharedFileSystemServiceOffering(p *ChangeSharedFileSystemServiceOfferingParams) (*ChangeSharedFileSystemServiceOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("changeSharedFileSystemServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeSharedFileSystemServiceOfferingResponse
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

type ChangeSharedFileSystemServiceOfferingResponse struct {
	Account                 string `json:"account"`
	Description             string `json:"description"`
	Diskioread              int64  `json:"diskioread"`
	Diskiowrite             int64  `json:"diskiowrite"`
	Diskkbsread             int64  `json:"diskkbsread"`
	Diskkbswrite            int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext"`
	Diskofferingid          string `json:"diskofferingid"`
	Diskofferingname        string `json:"diskofferingname"`
	Domain                  string `json:"domain"`
	Domainid                string `json:"domainid"`
	Domainpath              string `json:"domainpath"`
	Filesystem              string `json:"filesystem"`
	Hasannotations          bool   `json:"hasannotations"`
	Id                      string `json:"id"`
	Iscustomdiskoffering    bool   `json:"iscustomdiskoffering"`
	JobID                   string `json:"jobid"`
	Jobstatus               int    `json:"jobstatus"`
	Name                    string `json:"name"`
	Networkid               string `json:"networkid"`
	Networkname             string `json:"networkname"`
	Nic                     []Nic  `json:"nic"`
	Path                    string `json:"path"`
	Physicalsize            int64  `json:"physicalsize"`
	Project                 string `json:"project"`
	Projectid               string `json:"projectid"`
	Provider                string `json:"provider"`
	Provisioningtype        string `json:"provisioningtype"`
	Serviceofferingid       string `json:"serviceofferingid"`
	Serviceofferingname     string `json:"serviceofferingname"`
	Size                    int64  `json:"size"`
	Sizegb                  string `json:"sizegb"`
	State                   string `json:"state"`
	Storage                 string `json:"storage"`
	Storageid               string `json:"storageid"`
	Tags                    []Tags `json:"tags"`
	Utilization             string `json:"utilization"`
	Virtualmachineid        string `json:"virtualmachineid"`
	Virtualsize             int64  `json:"virtualsize"`
	Vmstate                 string `json:"vmstate"`
	Volumeid                string `json:"volumeid"`
	Volumename              string `json:"volumename"`
	Zoneid                  string `json:"zoneid"`
	Zonename                string `json:"zonename"`
}

type CreateSharedFileSystemParams struct {
	p map[string]interface{}
}

func (p *CreateSharedFileSystemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["filesystem"]; found {
		u.Set("filesystem", v.(string))
	}
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateSharedFileSystemParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateSharedFileSystemParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateSharedFileSystemParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateSharedFileSystemParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateSharedFileSystemParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *CreateSharedFileSystemParams) ResetDiskofferingid() {
	if p.p != nil && p.p["diskofferingid"] != nil {
		delete(p.p, "diskofferingid")
	}
}

func (p *CreateSharedFileSystemParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateSharedFileSystemParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateSharedFileSystemParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetFilesystem(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["filesystem"] = v
}

func (p *CreateSharedFileSystemParams) ResetFilesystem() {
	if p.p != nil && p.p["filesystem"] != nil {
		delete(p.p, "filesystem")
	}
}

func (p *CreateSharedFileSystemParams) GetFilesystem() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["filesystem"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *CreateSharedFileSystemParams) ResetMaxiops() {
	if p.p != nil && p.p["maxiops"] != nil {
		delete(p.p, "maxiops")
	}
}

func (p *CreateSharedFileSystemParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *CreateSharedFileSystemParams) ResetMiniops() {
	if p.p != nil && p.p["miniops"] != nil {
		delete(p.p, "miniops")
	}
}

func (p *CreateSharedFileSystemParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateSharedFileSystemParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateSharedFileSystemParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreateSharedFileSystemParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreateSharedFileSystemParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateSharedFileSystemParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateSharedFileSystemParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *CreateSharedFileSystemParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *CreateSharedFileSystemParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateSharedFileSystemParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *CreateSharedFileSystemParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *CreateSharedFileSystemParams) ResetSize() {
	if p.p != nil && p.p["size"] != nil {
		delete(p.p, "size")
	}
}

func (p *CreateSharedFileSystemParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

func (p *CreateSharedFileSystemParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateSharedFileSystemParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateSharedFileSystemParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSharedFileSystemParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewCreateSharedFileSystemParams(diskofferingid string, filesystem string, name string, networkid string, serviceofferingid string, zoneid string) *CreateSharedFileSystemParams {
	p := &CreateSharedFileSystemParams{}
	p.p = make(map[string]interface{})
	p.p["diskofferingid"] = diskofferingid
	p.p["filesystem"] = filesystem
	p.p["name"] = name
	p.p["networkid"] = networkid
	p.p["serviceofferingid"] = serviceofferingid
	p.p["zoneid"] = zoneid
	return p
}

// Create a new Shared File System of specified size and disk offering, attached to the given network
func (s *SharedFileSystemService) CreateSharedFileSystem(p *CreateSharedFileSystemParams) (*CreateSharedFileSystemResponse, error) {
	resp, err := s.cs.newPostRequest("createSharedFileSystem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSharedFileSystemResponse
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

type CreateSharedFileSystemResponse struct {
	Account                 string `json:"account"`
	Description             string `json:"description"`
	Diskioread              int64  `json:"diskioread"`
	Diskiowrite             int64  `json:"diskiowrite"`
	Diskkbsread             int64  `json:"diskkbsread"`
	Diskkbswrite            int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext"`
	Diskofferingid          string `json:"diskofferingid"`
	Diskofferingname        string `json:"diskofferingname"`
	Domain                  string `json:"domain"`
	Domainid                string `json:"domainid"`
	Domainpath              string `json:"domainpath"`
	Filesystem              string `json:"filesystem"`
	Hasannotations          bool   `json:"hasannotations"`
	Id                      string `json:"id"`
	Iscustomdiskoffering    bool   `json:"iscustomdiskoffering"`
	JobID                   string `json:"jobid"`
	Jobstatus               int    `json:"jobstatus"`
	Name                    string `json:"name"`
	Networkid               string `json:"networkid"`
	Networkname             string `json:"networkname"`
	Nic                     []Nic  `json:"nic"`
	Path                    string `json:"path"`
	Physicalsize            int64  `json:"physicalsize"`
	Project                 string `json:"project"`
	Projectid               string `json:"projectid"`
	Provider                string `json:"provider"`
	Provisioningtype        string `json:"provisioningtype"`
	Serviceofferingid       string `json:"serviceofferingid"`
	Serviceofferingname     string `json:"serviceofferingname"`
	Size                    int64  `json:"size"`
	Sizegb                  string `json:"sizegb"`
	State                   string `json:"state"`
	Storage                 string `json:"storage"`
	Storageid               string `json:"storageid"`
	Tags                    []Tags `json:"tags"`
	Utilization             string `json:"utilization"`
	Virtualmachineid        string `json:"virtualmachineid"`
	Virtualsize             int64  `json:"virtualsize"`
	Vmstate                 string `json:"vmstate"`
	Volumeid                string `json:"volumeid"`
	Volumename              string `json:"volumename"`
	Zoneid                  string `json:"zoneid"`
	Zonename                string `json:"zonename"`
}

type DestroySharedFileSystemParams struct {
	p map[string]interface{}
}

func (p *DestroySharedFileSystemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["expunge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("expunge", vv)
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DestroySharedFileSystemParams) SetExpunge(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expunge"] = v
}

func (p *DestroySharedFileSystemParams) ResetExpunge() {
	if p.p != nil && p.p["expunge"] != nil {
		delete(p.p, "expunge")
	}
}

func (p *DestroySharedFileSystemParams) GetExpunge() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["expunge"].(bool)
	return value, ok
}

func (p *DestroySharedFileSystemParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DestroySharedFileSystemParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *DestroySharedFileSystemParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *DestroySharedFileSystemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DestroySharedFileSystemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DestroySharedFileSystemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DestroySharedFileSystemParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewDestroySharedFileSystemParams() *DestroySharedFileSystemParams {
	p := &DestroySharedFileSystemParams{}
	p.p = make(map[string]interface{})
	return p
}

// Destroy a Shared FileSystem by id
func (s *SharedFileSystemService) DestroySharedFileSystem(p *DestroySharedFileSystemParams) (*DestroySharedFileSystemResponse, error) {
	resp, err := s.cs.newPostRequest("destroySharedFileSystem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroySharedFileSystemResponse
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

type DestroySharedFileSystemResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ExpungeSharedFileSystemParams struct {
	p map[string]interface{}
}

func (p *ExpungeSharedFileSystemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ExpungeSharedFileSystemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ExpungeSharedFileSystemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ExpungeSharedFileSystemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ExpungeSharedFileSystemParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewExpungeSharedFileSystemParams() *ExpungeSharedFileSystemParams {
	p := &ExpungeSharedFileSystemParams{}
	p.p = make(map[string]interface{})
	return p
}

// Expunge a Shared FileSystem by id
func (s *SharedFileSystemService) ExpungeSharedFileSystem(p *ExpungeSharedFileSystemParams) (*ExpungeSharedFileSystemResponse, error) {
	resp, err := s.cs.newPostRequest("expungeSharedFileSystem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExpungeSharedFileSystemResponse
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

type ExpungeSharedFileSystemResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListSharedFileSystemProvidersParams struct {
	p map[string]interface{}
}

func (p *ListSharedFileSystemProvidersParams) toURLValues() url.Values {
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

func (p *ListSharedFileSystemProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSharedFileSystemProvidersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSharedFileSystemProvidersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSharedFileSystemProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSharedFileSystemProvidersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSharedFileSystemProvidersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSharedFileSystemProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSharedFileSystemProvidersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSharedFileSystemProvidersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListSharedFileSystemProvidersParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewListSharedFileSystemProvidersParams() *ListSharedFileSystemProvidersParams {
	p := &ListSharedFileSystemProvidersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all available shared filesystem providers.
func (s *SharedFileSystemService) ListSharedFileSystemProviders(p *ListSharedFileSystemProvidersParams) (*ListSharedFileSystemProvidersResponse, error) {
	resp, err := s.cs.newRequest("listSharedFileSystemProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSharedFileSystemProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSharedFileSystemProvidersResponse struct {
	Count                     int                         `json:"count"`
	SharedFileSystemProviders []*SharedFileSystemProvider `json:"sharedfilesystemprovider"`
}

type SharedFileSystemProvider struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListSharedFileSystemsParams struct {
	p map[string]interface{}
}

func (p *ListSharedFileSystemsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
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
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
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
	if v, found := p.p["retrieveonlyresourcecount"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("retrieveonlyresourcecount", vv)
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListSharedFileSystemsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListSharedFileSystemsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListSharedFileSystemsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *ListSharedFileSystemsParams) ResetDiskofferingid() {
	if p.p != nil && p.p["diskofferingid"] != nil {
		delete(p.p, "diskofferingid")
	}
}

func (p *ListSharedFileSystemsParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListSharedFileSystemsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListSharedFileSystemsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSharedFileSystemsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListSharedFileSystemsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListSharedFileSystemsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListSharedFileSystemsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSharedFileSystemsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSharedFileSystemsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListSharedFileSystemsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListSharedFileSystemsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListSharedFileSystemsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListSharedFileSystemsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListSharedFileSystemsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListSharedFileSystemsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSharedFileSystemsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSharedFileSystemsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSharedFileSystemsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSharedFileSystemsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListSharedFileSystemsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListSharedFileSystemsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetRetrieveonlyresourcecount(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["retrieveonlyresourcecount"] = v
}

func (p *ListSharedFileSystemsParams) ResetRetrieveonlyresourcecount() {
	if p.p != nil && p.p["retrieveonlyresourcecount"] != nil {
		delete(p.p, "retrieveonlyresourcecount")
	}
}

func (p *ListSharedFileSystemsParams) GetRetrieveonlyresourcecount() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["retrieveonlyresourcecount"].(bool)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ListSharedFileSystemsParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ListSharedFileSystemsParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListSharedFileSystemsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListSharedFileSystemsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListSharedFileSystemsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListSharedFileSystemsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListSharedFileSystemsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSharedFileSystemsParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewListSharedFileSystemsParams() *ListSharedFileSystemsParams {
	p := &ListSharedFileSystemsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SharedFileSystemService) GetSharedFileSystemID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListSharedFileSystemsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSharedFileSystems(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SharedFileSystems[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SharedFileSystems {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SharedFileSystemService) GetSharedFileSystemByName(name string, opts ...OptionFunc) (*SharedFileSystem, int, error) {
	id, count, err := s.GetSharedFileSystemID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSharedFileSystemByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SharedFileSystemService) GetSharedFileSystemByID(id string, opts ...OptionFunc) (*SharedFileSystem, int, error) {
	p := &ListSharedFileSystemsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSharedFileSystems(p)
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
		return l.SharedFileSystems[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SharedFileSystem UUID: %s!", id)
}

// List Shared FileSystems
func (s *SharedFileSystemService) ListSharedFileSystems(p *ListSharedFileSystemsParams) (*ListSharedFileSystemsResponse, error) {
	resp, err := s.cs.newRequest("listSharedFileSystems", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSharedFileSystemsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSharedFileSystemsResponse struct {
	Count             int                 `json:"count"`
	SharedFileSystems []*SharedFileSystem `json:"sharedfilesystem"`
}

type SharedFileSystem struct {
	Account                 string `json:"account"`
	Description             string `json:"description"`
	Diskioread              int64  `json:"diskioread"`
	Diskiowrite             int64  `json:"diskiowrite"`
	Diskkbsread             int64  `json:"diskkbsread"`
	Diskkbswrite            int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext"`
	Diskofferingid          string `json:"diskofferingid"`
	Diskofferingname        string `json:"diskofferingname"`
	Domain                  string `json:"domain"`
	Domainid                string `json:"domainid"`
	Domainpath              string `json:"domainpath"`
	Filesystem              string `json:"filesystem"`
	Hasannotations          bool   `json:"hasannotations"`
	Id                      string `json:"id"`
	Iscustomdiskoffering    bool   `json:"iscustomdiskoffering"`
	JobID                   string `json:"jobid"`
	Jobstatus               int    `json:"jobstatus"`
	Name                    string `json:"name"`
	Networkid               string `json:"networkid"`
	Networkname             string `json:"networkname"`
	Nic                     []Nic  `json:"nic"`
	Path                    string `json:"path"`
	Physicalsize            int64  `json:"physicalsize"`
	Project                 string `json:"project"`
	Projectid               string `json:"projectid"`
	Provider                string `json:"provider"`
	Provisioningtype        string `json:"provisioningtype"`
	Serviceofferingid       string `json:"serviceofferingid"`
	Serviceofferingname     string `json:"serviceofferingname"`
	Size                    int64  `json:"size"`
	Sizegb                  string `json:"sizegb"`
	State                   string `json:"state"`
	Storage                 string `json:"storage"`
	Storageid               string `json:"storageid"`
	Tags                    []Tags `json:"tags"`
	Utilization             string `json:"utilization"`
	Virtualmachineid        string `json:"virtualmachineid"`
	Virtualsize             int64  `json:"virtualsize"`
	Vmstate                 string `json:"vmstate"`
	Volumeid                string `json:"volumeid"`
	Volumename              string `json:"volumename"`
	Zoneid                  string `json:"zoneid"`
	Zonename                string `json:"zonename"`
}

type RecoverSharedFileSystemParams struct {
	p map[string]interface{}
}

func (p *RecoverSharedFileSystemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RecoverSharedFileSystemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RecoverSharedFileSystemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RecoverSharedFileSystemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RecoverSharedFileSystemParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewRecoverSharedFileSystemParams() *RecoverSharedFileSystemParams {
	p := &RecoverSharedFileSystemParams{}
	p.p = make(map[string]interface{})
	return p
}

// Recover a Shared FileSystem by id
func (s *SharedFileSystemService) RecoverSharedFileSystem(p *RecoverSharedFileSystemParams) (*RecoverSharedFileSystemResponse, error) {
	resp, err := s.cs.newPostRequest("recoverSharedFileSystem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RecoverSharedFileSystemResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RecoverSharedFileSystemResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RecoverSharedFileSystemResponse) UnmarshalJSON(b []byte) error {
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

	type alias RecoverSharedFileSystemResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RestartSharedFileSystemParams struct {
	p map[string]interface{}
}

func (p *RestartSharedFileSystemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RestartSharedFileSystemParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *RestartSharedFileSystemParams) ResetCleanup() {
	if p.p != nil && p.p["cleanup"] != nil {
		delete(p.p, "cleanup")
	}
}

func (p *RestartSharedFileSystemParams) GetCleanup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanup"].(bool)
	return value, ok
}

func (p *RestartSharedFileSystemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RestartSharedFileSystemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RestartSharedFileSystemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RestartSharedFileSystemParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewRestartSharedFileSystemParams(id string) *RestartSharedFileSystemParams {
	p := &RestartSharedFileSystemParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Restart a Shared FileSystem
func (s *SharedFileSystemService) RestartSharedFileSystem(p *RestartSharedFileSystemParams) (*RestartSharedFileSystemResponse, error) {
	resp, err := s.cs.newPostRequest("restartSharedFileSystem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestartSharedFileSystemResponse
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

type RestartSharedFileSystemResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type StartSharedFileSystemParams struct {
	p map[string]interface{}
}

func (p *StartSharedFileSystemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartSharedFileSystemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StartSharedFileSystemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StartSharedFileSystemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StartSharedFileSystemParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewStartSharedFileSystemParams(id string) *StartSharedFileSystemParams {
	p := &StartSharedFileSystemParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Start a Shared FileSystem
func (s *SharedFileSystemService) StartSharedFileSystem(p *StartSharedFileSystemParams) (*StartSharedFileSystemResponse, error) {
	resp, err := s.cs.newPostRequest("startSharedFileSystem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartSharedFileSystemResponse
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

type StartSharedFileSystemResponse struct {
	Account                 string `json:"account"`
	Description             string `json:"description"`
	Diskioread              int64  `json:"diskioread"`
	Diskiowrite             int64  `json:"diskiowrite"`
	Diskkbsread             int64  `json:"diskkbsread"`
	Diskkbswrite            int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext"`
	Diskofferingid          string `json:"diskofferingid"`
	Diskofferingname        string `json:"diskofferingname"`
	Domain                  string `json:"domain"`
	Domainid                string `json:"domainid"`
	Domainpath              string `json:"domainpath"`
	Filesystem              string `json:"filesystem"`
	Hasannotations          bool   `json:"hasannotations"`
	Id                      string `json:"id"`
	Iscustomdiskoffering    bool   `json:"iscustomdiskoffering"`
	JobID                   string `json:"jobid"`
	Jobstatus               int    `json:"jobstatus"`
	Name                    string `json:"name"`
	Networkid               string `json:"networkid"`
	Networkname             string `json:"networkname"`
	Nic                     []Nic  `json:"nic"`
	Path                    string `json:"path"`
	Physicalsize            int64  `json:"physicalsize"`
	Project                 string `json:"project"`
	Projectid               string `json:"projectid"`
	Provider                string `json:"provider"`
	Provisioningtype        string `json:"provisioningtype"`
	Serviceofferingid       string `json:"serviceofferingid"`
	Serviceofferingname     string `json:"serviceofferingname"`
	Size                    int64  `json:"size"`
	Sizegb                  string `json:"sizegb"`
	State                   string `json:"state"`
	Storage                 string `json:"storage"`
	Storageid               string `json:"storageid"`
	Tags                    []Tags `json:"tags"`
	Utilization             string `json:"utilization"`
	Virtualmachineid        string `json:"virtualmachineid"`
	Virtualsize             int64  `json:"virtualsize"`
	Vmstate                 string `json:"vmstate"`
	Volumeid                string `json:"volumeid"`
	Volumename              string `json:"volumename"`
	Zoneid                  string `json:"zoneid"`
	Zonename                string `json:"zonename"`
}

type StopSharedFileSystemParams struct {
	p map[string]interface{}
}

func (p *StopSharedFileSystemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StopSharedFileSystemParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *StopSharedFileSystemParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *StopSharedFileSystemParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *StopSharedFileSystemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StopSharedFileSystemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StopSharedFileSystemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopSharedFileSystemParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewStopSharedFileSystemParams(id string) *StopSharedFileSystemParams {
	p := &StopSharedFileSystemParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stop a Shared FileSystem
func (s *SharedFileSystemService) StopSharedFileSystem(p *StopSharedFileSystemParams) (*StopSharedFileSystemResponse, error) {
	resp, err := s.cs.newPostRequest("stopSharedFileSystem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopSharedFileSystemResponse
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

type StopSharedFileSystemResponse struct {
	Account                 string `json:"account"`
	Description             string `json:"description"`
	Diskioread              int64  `json:"diskioread"`
	Diskiowrite             int64  `json:"diskiowrite"`
	Diskkbsread             int64  `json:"diskkbsread"`
	Diskkbswrite            int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext"`
	Diskofferingid          string `json:"diskofferingid"`
	Diskofferingname        string `json:"diskofferingname"`
	Domain                  string `json:"domain"`
	Domainid                string `json:"domainid"`
	Domainpath              string `json:"domainpath"`
	Filesystem              string `json:"filesystem"`
	Hasannotations          bool   `json:"hasannotations"`
	Id                      string `json:"id"`
	Iscustomdiskoffering    bool   `json:"iscustomdiskoffering"`
	JobID                   string `json:"jobid"`
	Jobstatus               int    `json:"jobstatus"`
	Name                    string `json:"name"`
	Networkid               string `json:"networkid"`
	Networkname             string `json:"networkname"`
	Nic                     []Nic  `json:"nic"`
	Path                    string `json:"path"`
	Physicalsize            int64  `json:"physicalsize"`
	Project                 string `json:"project"`
	Projectid               string `json:"projectid"`
	Provider                string `json:"provider"`
	Provisioningtype        string `json:"provisioningtype"`
	Serviceofferingid       string `json:"serviceofferingid"`
	Serviceofferingname     string `json:"serviceofferingname"`
	Size                    int64  `json:"size"`
	Sizegb                  string `json:"sizegb"`
	State                   string `json:"state"`
	Storage                 string `json:"storage"`
	Storageid               string `json:"storageid"`
	Tags                    []Tags `json:"tags"`
	Utilization             string `json:"utilization"`
	Virtualmachineid        string `json:"virtualmachineid"`
	Virtualsize             int64  `json:"virtualsize"`
	Vmstate                 string `json:"vmstate"`
	Volumeid                string `json:"volumeid"`
	Volumename              string `json:"volumename"`
	Zoneid                  string `json:"zoneid"`
	Zonename                string `json:"zonename"`
}

type UpdateSharedFileSystemParams struct {
	p map[string]interface{}
}

func (p *UpdateSharedFileSystemParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateSharedFileSystemParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateSharedFileSystemParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateSharedFileSystemParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateSharedFileSystemParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateSharedFileSystemParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateSharedFileSystemParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateSharedFileSystemParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateSharedFileSystemParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateSharedFileSystemParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateSharedFileSystemParams instance,
// as then you are sure you have configured all required params
func (s *SharedFileSystemService) NewUpdateSharedFileSystemParams(id string) *UpdateSharedFileSystemParams {
	p := &UpdateSharedFileSystemParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Update a Shared FileSystem
func (s *SharedFileSystemService) UpdateSharedFileSystem(p *UpdateSharedFileSystemParams) (*UpdateSharedFileSystemResponse, error) {
	resp, err := s.cs.newPostRequest("updateSharedFileSystem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateSharedFileSystemResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateSharedFileSystemResponse struct {
	Account                 string `json:"account"`
	Description             string `json:"description"`
	Diskioread              int64  `json:"diskioread"`
	Diskiowrite             int64  `json:"diskiowrite"`
	Diskkbsread             int64  `json:"diskkbsread"`
	Diskkbswrite            int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext string `json:"diskofferingdisplaytext"`
	Diskofferingid          string `json:"diskofferingid"`
	Diskofferingname        string `json:"diskofferingname"`
	Domain                  string `json:"domain"`
	Domainid                string `json:"domainid"`
	Domainpath              string `json:"domainpath"`
	Filesystem              string `json:"filesystem"`
	Hasannotations          bool   `json:"hasannotations"`
	Id                      string `json:"id"`
	Iscustomdiskoffering    bool   `json:"iscustomdiskoffering"`
	JobID                   string `json:"jobid"`
	Jobstatus               int    `json:"jobstatus"`
	Name                    string `json:"name"`
	Networkid               string `json:"networkid"`
	Networkname             string `json:"networkname"`
	Nic                     []Nic  `json:"nic"`
	Path                    string `json:"path"`
	Physicalsize            int64  `json:"physicalsize"`
	Project                 string `json:"project"`
	Projectid               string `json:"projectid"`
	Provider                string `json:"provider"`
	Provisioningtype        string `json:"provisioningtype"`
	Serviceofferingid       string `json:"serviceofferingid"`
	Serviceofferingname     string `json:"serviceofferingname"`
	Size                    int64  `json:"size"`
	Sizegb                  string `json:"sizegb"`
	State                   string `json:"state"`
	Storage                 string `json:"storage"`
	Storageid               string `json:"storageid"`
	Tags                    []Tags `json:"tags"`
	Utilization             string `json:"utilization"`
	Virtualmachineid        string `json:"virtualmachineid"`
	Virtualsize             int64  `json:"virtualsize"`
	Vmstate                 string `json:"vmstate"`
	Volumeid                string `json:"volumeid"`
	Volumename              string `json:"volumename"`
	Zoneid                  string `json:"zoneid"`
	Zonename                string `json:"zonename"`
}
