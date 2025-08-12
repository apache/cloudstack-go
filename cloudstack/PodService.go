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

type PodServiceIface interface {
	CreatePod(p *CreatePodParams) (*CreatePodResponse, error)
	NewCreatePodParams(name string, zoneid string) *CreatePodParams
	DedicatePod(p *DedicatePodParams) (*DedicatePodResponse, error)
	NewDedicatePodParams(domainid string, podid string) *DedicatePodParams
	DeletePod(p *DeletePodParams) (*DeletePodResponse, error)
	NewDeletePodParams(id string) *DeletePodParams
	ListDedicatedPods(p *ListDedicatedPodsParams) (*ListDedicatedPodsResponse, error)
	NewListDedicatedPodsParams() *ListDedicatedPodsParams
	ListPods(p *ListPodsParams) (*ListPodsResponse, error)
	NewListPodsParams() *ListPodsParams
	GetPodID(name string, opts ...OptionFunc) (string, int, error)
	GetPodByName(name string, opts ...OptionFunc) (*Pod, int, error)
	GetPodByID(id string, opts ...OptionFunc) (*Pod, int, error)
	ReleaseDedicatedPod(p *ReleaseDedicatedPodParams) (*ReleaseDedicatedPodResponse, error)
	NewReleaseDedicatedPodParams(podid string) *ReleaseDedicatedPodParams
	UpdatePod(p *UpdatePodParams) (*UpdatePodResponse, error)
	NewUpdatePodParams(id string) *UpdatePodParams
}

type CreatePodParams struct {
	p map[string]interface{}
}

func (p *CreatePodParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreatePodParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *CreatePodParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *CreatePodParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *CreatePodParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
}

func (p *CreatePodParams) ResetEndip() {
	if p.p != nil && p.p["endip"] != nil {
		delete(p.p, "endip")
	}
}

func (p *CreatePodParams) GetEndip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endip"].(string)
	return value, ok
}

func (p *CreatePodParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
}

func (p *CreatePodParams) ResetGateway() {
	if p.p != nil && p.p["gateway"] != nil {
		delete(p.p, "gateway")
	}
}

func (p *CreatePodParams) GetGateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gateway"].(string)
	return value, ok
}

func (p *CreatePodParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreatePodParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreatePodParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreatePodParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
}

func (p *CreatePodParams) ResetNetmask() {
	if p.p != nil && p.p["netmask"] != nil {
		delete(p.p, "netmask")
	}
}

func (p *CreatePodParams) GetNetmask() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["netmask"].(string)
	return value, ok
}

func (p *CreatePodParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
}

func (p *CreatePodParams) ResetStartip() {
	if p.p != nil && p.p["startip"] != nil {
		delete(p.p, "startip")
	}
}

func (p *CreatePodParams) GetStartip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startip"].(string)
	return value, ok
}

func (p *CreatePodParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreatePodParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreatePodParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewCreatePodParams(name string, zoneid string) *CreatePodParams {
	p := &CreatePodParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// Creates a new Pod.
func (s *PodService) CreatePod(p *CreatePodParams) (*CreatePodResponse, error) {
	resp, err := s.cs.newRequest("createPod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreatePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreatePodResponse struct {
	Allocationstate string                      `json:"allocationstate"`
	Capacity        []CreatePodResponseCapacity `json:"capacity"`
	Endip           []string                    `json:"endip"`
	Forsystemvms    []string                    `json:"forsystemvms"`
	Gateway         string                      `json:"gateway"`
	Hasannotations  bool                        `json:"hasannotations"`
	Id              string                      `json:"id"`
	Ipranges        []CreatePodResponseIpranges `json:"ipranges"`
	JobID           string                      `json:"jobid"`
	Jobstatus       int                         `json:"jobstatus"`
	Name            string                      `json:"name"`
	Netmask         string                      `json:"netmask"`
	Startip         []string                    `json:"startip"`
	Vlanid          []string                    `json:"vlanid"`
	Zoneid          string                      `json:"zoneid"`
	Zonename        string                      `json:"zonename"`
}

type CreatePodResponseIpranges struct {
	Cidr         string `json:"cidr"`
	Endip        string `json:"endip"`
	Forsystemvms string `json:"forsystemvms"`
	Gateway      string `json:"gateway"`
	Startip      string `json:"startip"`
	Vlanid       string `json:"vlanid"`
}

type CreatePodResponseCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type DedicatePodParams struct {
	p map[string]interface{}
}

func (p *DedicatePodParams) toURLValues() url.Values {
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (p *DedicatePodParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DedicatePodParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DedicatePodParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DedicatePodParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DedicatePodParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DedicatePodParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DedicatePodParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *DedicatePodParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *DedicatePodParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewDedicatePodParams(domainid string, podid string) *DedicatePodParams {
	p := &DedicatePodParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	p.p["podid"] = podid
	return p
}

// Dedicates a Pod.
func (s *PodService) DedicatePod(p *DedicatePodParams) (*DedicatePodResponse, error) {
	resp, err := s.cs.newRequest("dedicatePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicatePodResponse
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

type DedicatePodResponse struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Podid           string `json:"podid"`
	Podname         string `json:"podname"`
}

type DeletePodParams struct {
	p map[string]interface{}
}

func (p *DeletePodParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePodParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeletePodParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeletePodParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewDeletePodParams(id string) *DeletePodParams {
	p := &DeletePodParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Pod.
func (s *PodService) DeletePod(p *DeletePodParams) (*DeletePodResponse, error) {
	resp, err := s.cs.newRequest("deletePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeletePodResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeletePodResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeletePodResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListDedicatedPodsParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedPodsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (p *ListDedicatedPodsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListDedicatedPodsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListDedicatedPodsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListDedicatedPodsParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
}

func (p *ListDedicatedPodsParams) ResetAffinitygroupid() {
	if p.p != nil && p.p["affinitygroupid"] != nil {
		delete(p.p, "affinitygroupid")
	}
}

func (p *ListDedicatedPodsParams) GetAffinitygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupid"].(string)
	return value, ok
}

func (p *ListDedicatedPodsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListDedicatedPodsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListDedicatedPodsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListDedicatedPodsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListDedicatedPodsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListDedicatedPodsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListDedicatedPodsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListDedicatedPodsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListDedicatedPodsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListDedicatedPodsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListDedicatedPodsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListDedicatedPodsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListDedicatedPodsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListDedicatedPodsParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListDedicatedPodsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

// You should always use this function to get a new ListDedicatedPodsParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewListDedicatedPodsParams() *ListDedicatedPodsParams {
	p := &ListDedicatedPodsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists dedicated pods.
func (s *PodService) ListDedicatedPods(p *ListDedicatedPodsParams) (*ListDedicatedPodsResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedPods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedPodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedPodsResponse struct {
	Count         int             `json:"count"`
	DedicatedPods []*DedicatedPod `json:"dedicatedpod"`
}

type DedicatedPod struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Podid           string `json:"podid"`
	Podname         string `json:"podname"`
}

type ListPodsParams struct {
	p map[string]interface{}
}

func (p *ListPodsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
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
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListPodsParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *ListPodsParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *ListPodsParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *ListPodsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListPodsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListPodsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListPodsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListPodsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListPodsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListPodsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListPodsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListPodsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListPodsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListPodsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListPodsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListPodsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListPodsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListPodsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListPodsParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
}

func (p *ListPodsParams) ResetShowcapacities() {
	if p.p != nil && p.p["showcapacities"] != nil {
		delete(p.p, "showcapacities")
	}
}

func (p *ListPodsParams) GetShowcapacities() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showcapacities"].(bool)
	return value, ok
}

func (p *ListPodsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListPodsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListPodsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPodsParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewListPodsParams() *ListPodsParams {
	p := &ListPodsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PodService) GetPodID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListPodsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListPods(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Pods[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Pods {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PodService) GetPodByName(name string, opts ...OptionFunc) (*Pod, int, error) {
	id, count, err := s.GetPodID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetPodByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PodService) GetPodByID(id string, opts ...OptionFunc) (*Pod, int, error) {
	p := &ListPodsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPods(p)
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
		return l.Pods[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Pod UUID: %s!", id)
}

// Lists all Pods.
func (s *PodService) ListPods(p *ListPodsParams) (*ListPodsResponse, error) {
	resp, err := s.cs.newRequest("listPods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPodsResponse struct {
	Count int    `json:"count"`
	Pods  []*Pod `json:"pod"`
}

type Pod struct {
	Allocationstate string        `json:"allocationstate"`
	Capacity        []PodCapacity `json:"capacity"`
	Endip           []string      `json:"endip"`
	Forsystemvms    []string      `json:"forsystemvms"`
	Gateway         string        `json:"gateway"`
	Hasannotations  bool          `json:"hasannotations"`
	Id              string        `json:"id"`
	Ipranges        []PodIpranges `json:"ipranges"`
	JobID           string        `json:"jobid"`
	Jobstatus       int           `json:"jobstatus"`
	Name            string        `json:"name"`
	Netmask         string        `json:"netmask"`
	Startip         []string      `json:"startip"`
	Vlanid          []string      `json:"vlanid"`
	Zoneid          string        `json:"zoneid"`
	Zonename        string        `json:"zonename"`
}

type PodIpranges struct {
	Cidr         string `json:"cidr"`
	Endip        string `json:"endip"`
	Forsystemvms string `json:"forsystemvms"`
	Gateway      string `json:"gateway"`
	Startip      string `json:"startip"`
	Vlanid       string `json:"vlanid"`
}

type PodCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type ReleaseDedicatedPodParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedPodParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedPodParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ReleaseDedicatedPodParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ReleaseDedicatedPodParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedPodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewReleaseDedicatedPodParams(podid string) *ReleaseDedicatedPodParams {
	p := &ReleaseDedicatedPodParams{}
	p.p = make(map[string]interface{})
	p.p["podid"] = podid
	return p
}

// Release the dedication for the pod
func (s *PodService) ReleaseDedicatedPod(p *ReleaseDedicatedPodParams) (*ReleaseDedicatedPodResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedPod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedPodResponse
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

type ReleaseDedicatedPodResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdatePodParams struct {
	p map[string]interface{}
}

func (p *UpdatePodParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	return u
}

func (p *UpdatePodParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *UpdatePodParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *UpdatePodParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *UpdatePodParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
}

func (p *UpdatePodParams) ResetEndip() {
	if p.p != nil && p.p["endip"] != nil {
		delete(p.p, "endip")
	}
}

func (p *UpdatePodParams) GetEndip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endip"].(string)
	return value, ok
}

func (p *UpdatePodParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
}

func (p *UpdatePodParams) ResetGateway() {
	if p.p != nil && p.p["gateway"] != nil {
		delete(p.p, "gateway")
	}
}

func (p *UpdatePodParams) GetGateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gateway"].(string)
	return value, ok
}

func (p *UpdatePodParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdatePodParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdatePodParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdatePodParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdatePodParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdatePodParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdatePodParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
}

func (p *UpdatePodParams) ResetNetmask() {
	if p.p != nil && p.p["netmask"] != nil {
		delete(p.p, "netmask")
	}
}

func (p *UpdatePodParams) GetNetmask() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["netmask"].(string)
	return value, ok
}

func (p *UpdatePodParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
}

func (p *UpdatePodParams) ResetStartip() {
	if p.p != nil && p.p["startip"] != nil {
		delete(p.p, "startip")
	}
}

func (p *UpdatePodParams) GetStartip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startip"].(string)
	return value, ok
}

// You should always use this function to get a new UpdatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewUpdatePodParams(id string) *UpdatePodParams {
	p := &UpdatePodParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a Pod.
func (s *PodService) UpdatePod(p *UpdatePodParams) (*UpdatePodResponse, error) {
	resp, err := s.cs.newRequest("updatePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdatePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdatePodResponse struct {
	Allocationstate string                      `json:"allocationstate"`
	Capacity        []UpdatePodResponseCapacity `json:"capacity"`
	Endip           []string                    `json:"endip"`
	Forsystemvms    []string                    `json:"forsystemvms"`
	Gateway         string                      `json:"gateway"`
	Hasannotations  bool                        `json:"hasannotations"`
	Id              string                      `json:"id"`
	Ipranges        []UpdatePodResponseIpranges `json:"ipranges"`
	JobID           string                      `json:"jobid"`
	Jobstatus       int                         `json:"jobstatus"`
	Name            string                      `json:"name"`
	Netmask         string                      `json:"netmask"`
	Startip         []string                    `json:"startip"`
	Vlanid          []string                    `json:"vlanid"`
	Zoneid          string                      `json:"zoneid"`
	Zonename        string                      `json:"zonename"`
}

type UpdatePodResponseIpranges struct {
	Cidr         string `json:"cidr"`
	Endip        string `json:"endip"`
	Forsystemvms string `json:"forsystemvms"`
	Gateway      string `json:"gateway"`
	Startip      string `json:"startip"`
	Vlanid       string `json:"vlanid"`
}

type UpdatePodResponseCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}
