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

type HypervisorServiceIface interface {
	ListHypervisorCapabilities(p *ListHypervisorCapabilitiesParams) (*ListHypervisorCapabilitiesResponse, error)
	NewListHypervisorCapabilitiesParams() *ListHypervisorCapabilitiesParams
	GetHypervisorCapabilityByID(id string, opts ...OptionFunc) (*HypervisorCapability, int, error)
	ListHypervisors(p *ListHypervisorsParams) (*ListHypervisorsResponse, error)
	NewListHypervisorsParams() *ListHypervisorsParams
	UpdateHypervisorCapabilities(p *UpdateHypervisorCapabilitiesParams) (*UpdateHypervisorCapabilitiesResponse, error)
	NewUpdateHypervisorCapabilitiesParams() *UpdateHypervisorCapabilitiesParams
}

type ListHypervisorCapabilitiesParams struct {
	p map[string]interface{}
}

func (p *ListHypervisorCapabilitiesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	return u
}

func (p *ListHypervisorCapabilitiesParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListHypervisorCapabilitiesParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListHypervisorCapabilitiesParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListHypervisorCapabilitiesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListHypervisorCapabilitiesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListHypervisorCapabilitiesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListHypervisorCapabilitiesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListHypervisorCapabilitiesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListHypervisorCapabilitiesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListHypervisorCapabilitiesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListHypervisorCapabilitiesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListHypervisorCapabilitiesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListHypervisorCapabilitiesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListHypervisorCapabilitiesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListHypervisorCapabilitiesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListHypervisorCapabilitiesParams instance,
// as then you are sure you have configured all required params
func (s *HypervisorService) NewListHypervisorCapabilitiesParams() *ListHypervisorCapabilitiesParams {
	p := &ListHypervisorCapabilitiesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HypervisorService) GetHypervisorCapabilityByID(id string, opts ...OptionFunc) (*HypervisorCapability, int, error) {
	p := &ListHypervisorCapabilitiesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListHypervisorCapabilities(p)
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
		return l.HypervisorCapabilities[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for HypervisorCapability UUID: %s!", id)
}

// Lists all hypervisor capabilities.
func (s *HypervisorService) ListHypervisorCapabilities(p *ListHypervisorCapabilitiesParams) (*ListHypervisorCapabilitiesResponse, error) {
	resp, err := s.cs.newRequest("listHypervisorCapabilities", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHypervisorCapabilitiesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHypervisorCapabilitiesResponse struct {
	Count                  int                     `json:"count"`
	HypervisorCapabilities []*HypervisorCapability `json:"hypervisorcapability"`
}

type HypervisorCapability struct {
	Hypervisor           string `json:"hypervisor"`
	Hypervisorversion    string `json:"hypervisorversion"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Maxdatavolumeslimit  int    `json:"maxdatavolumeslimit"`
	Maxguestslimit       int64  `json:"maxguestslimit"`
	Maxhostspercluster   int    `json:"maxhostspercluster"`
	Securitygroupenabled bool   `json:"securitygroupenabled"`
	Storagemotionenabled bool   `json:"storagemotionenabled"`
	Vmsnapshotenabled    bool   `json:"vmsnapshotenabled"`
}

type ListHypervisorsParams struct {
	p map[string]interface{}
}

func (p *ListHypervisorsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListHypervisorsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListHypervisorsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListHypervisorsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListHypervisorsParams instance,
// as then you are sure you have configured all required params
func (s *HypervisorService) NewListHypervisorsParams() *ListHypervisorsParams {
	p := &ListHypervisorsParams{}
	p.p = make(map[string]interface{})
	return p
}

// List hypervisors
func (s *HypervisorService) ListHypervisors(p *ListHypervisorsParams) (*ListHypervisorsResponse, error) {
	resp, err := s.cs.newRequest("listHypervisors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHypervisorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHypervisorsResponse struct {
	Count       int           `json:"count"`
	Hypervisors []*Hypervisor `json:"hypervisor"`
}

type Hypervisor struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type UpdateHypervisorCapabilitiesParams struct {
	p map[string]interface{}
}

func (p *UpdateHypervisorCapabilitiesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["maxdatavolumeslimit"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxdatavolumeslimit", vv)
	}
	if v, found := p.p["maxguestslimit"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxguestslimit", vv)
	}
	if v, found := p.p["maxhostspercluster"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxhostspercluster", vv)
	}
	if v, found := p.p["securitygroupenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("securitygroupenabled", vv)
	}
	if v, found := p.p["storagemotionenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("storagemotionenabled", vv)
	}
	if v, found := p.p["vmsnapshotenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("vmsnapshotenabled", vv)
	}
	return u
}

func (p *UpdateHypervisorCapabilitiesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateHypervisorCapabilitiesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateHypervisorCapabilitiesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateHypervisorCapabilitiesParams) SetMaxdatavolumeslimit(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxdatavolumeslimit"] = v
}

func (p *UpdateHypervisorCapabilitiesParams) ResetMaxdatavolumeslimit() {
	if p.p != nil && p.p["maxdatavolumeslimit"] != nil {
		delete(p.p, "maxdatavolumeslimit")
	}
}

func (p *UpdateHypervisorCapabilitiesParams) GetMaxdatavolumeslimit() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxdatavolumeslimit"].(int)
	return value, ok
}

func (p *UpdateHypervisorCapabilitiesParams) SetMaxguestslimit(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxguestslimit"] = v
}

func (p *UpdateHypervisorCapabilitiesParams) ResetMaxguestslimit() {
	if p.p != nil && p.p["maxguestslimit"] != nil {
		delete(p.p, "maxguestslimit")
	}
}

func (p *UpdateHypervisorCapabilitiesParams) GetMaxguestslimit() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxguestslimit"].(int64)
	return value, ok
}

func (p *UpdateHypervisorCapabilitiesParams) SetMaxhostspercluster(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxhostspercluster"] = v
}

func (p *UpdateHypervisorCapabilitiesParams) ResetMaxhostspercluster() {
	if p.p != nil && p.p["maxhostspercluster"] != nil {
		delete(p.p, "maxhostspercluster")
	}
}

func (p *UpdateHypervisorCapabilitiesParams) GetMaxhostspercluster() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxhostspercluster"].(int)
	return value, ok
}

func (p *UpdateHypervisorCapabilitiesParams) SetSecuritygroupenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupenabled"] = v
}

func (p *UpdateHypervisorCapabilitiesParams) ResetSecuritygroupenabled() {
	if p.p != nil && p.p["securitygroupenabled"] != nil {
		delete(p.p, "securitygroupenabled")
	}
}

func (p *UpdateHypervisorCapabilitiesParams) GetSecuritygroupenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupenabled"].(bool)
	return value, ok
}

func (p *UpdateHypervisorCapabilitiesParams) SetStoragemotionenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagemotionenabled"] = v
}

func (p *UpdateHypervisorCapabilitiesParams) ResetStoragemotionenabled() {
	if p.p != nil && p.p["storagemotionenabled"] != nil {
		delete(p.p, "storagemotionenabled")
	}
}

func (p *UpdateHypervisorCapabilitiesParams) GetStoragemotionenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagemotionenabled"].(bool)
	return value, ok
}

func (p *UpdateHypervisorCapabilitiesParams) SetVmsnapshotenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotenabled"] = v
}

func (p *UpdateHypervisorCapabilitiesParams) ResetVmsnapshotenabled() {
	if p.p != nil && p.p["vmsnapshotenabled"] != nil {
		delete(p.p, "vmsnapshotenabled")
	}
}

func (p *UpdateHypervisorCapabilitiesParams) GetVmsnapshotenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmsnapshotenabled"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateHypervisorCapabilitiesParams instance,
// as then you are sure you have configured all required params
func (s *HypervisorService) NewUpdateHypervisorCapabilitiesParams() *UpdateHypervisorCapabilitiesParams {
	p := &UpdateHypervisorCapabilitiesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Updates a hypervisor capabilities.
func (s *HypervisorService) UpdateHypervisorCapabilities(p *UpdateHypervisorCapabilitiesParams) (*UpdateHypervisorCapabilitiesResponse, error) {
	resp, err := s.cs.newRequest("updateHypervisorCapabilities", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateHypervisorCapabilitiesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateHypervisorCapabilitiesResponse struct {
	Hypervisor           string `json:"hypervisor"`
	Hypervisorversion    string `json:"hypervisorversion"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Maxdatavolumeslimit  int    `json:"maxdatavolumeslimit"`
	Maxguestslimit       int64  `json:"maxguestslimit"`
	Maxhostspercluster   int    `json:"maxhostspercluster"`
	Securitygroupenabled bool   `json:"securitygroupenabled"`
	Storagemotionenabled bool   `json:"storagemotionenabled"`
	Vmsnapshotenabled    bool   `json:"vmsnapshotenabled"`
}
