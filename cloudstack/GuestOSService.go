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

type GuestOSServiceIface interface {
	AddGuestOs(p *AddGuestOsParams) (*AddGuestOsResponse, error)
	NewAddGuestOsParams(oscategoryid string, osdisplayname string) *AddGuestOsParams
	AddGuestOsMapping(p *AddGuestOsMappingParams) (*AddGuestOsMappingResponse, error)
	NewAddGuestOsMappingParams(hypervisor string, hypervisorversion string, osnameforhypervisor string) *AddGuestOsMappingParams
	ListGuestOsMapping(p *ListGuestOsMappingParams) (*ListGuestOsMappingResponse, error)
	NewListGuestOsMappingParams() *ListGuestOsMappingParams
	GetGuestOsMappingByID(id string, opts ...OptionFunc) (*GuestOsMapping, int, error)
	ListOsCategories(p *ListOsCategoriesParams) (*ListOsCategoriesResponse, error)
	NewListOsCategoriesParams() *ListOsCategoriesParams
	GetOsCategoryID(name string, opts ...OptionFunc) (string, int, error)
	GetOsCategoryByName(name string, opts ...OptionFunc) (*OsCategory, int, error)
	GetOsCategoryByID(id string, opts ...OptionFunc) (*OsCategory, int, error)
	ListOsTypes(p *ListOsTypesParams) (*ListOsTypesResponse, error)
	NewListOsTypesParams() *ListOsTypesParams
	GetOsTypeID(keyword string, opts ...OptionFunc) (string, int, error)
	GetOsTypeByName(name string, opts ...OptionFunc) (*OsType, int, error)
	GetOsTypeByID(id string, opts ...OptionFunc) (*OsType, int, error)
	RemoveGuestOs(p *RemoveGuestOsParams) (*RemoveGuestOsResponse, error)
	NewRemoveGuestOsParams(id string) *RemoveGuestOsParams
	RemoveGuestOsMapping(p *RemoveGuestOsMappingParams) (*RemoveGuestOsMappingResponse, error)
	NewRemoveGuestOsMappingParams(id string) *RemoveGuestOsMappingParams
	UpdateGuestOs(p *UpdateGuestOsParams) (*UpdateGuestOsResponse, error)
	NewUpdateGuestOsParams(id string, osdisplayname string) *UpdateGuestOsParams
	UpdateGuestOsMapping(p *UpdateGuestOsMappingParams) (*UpdateGuestOsMappingResponse, error)
	NewUpdateGuestOsMappingParams(id string, osnameforhypervisor string) *UpdateGuestOsMappingParams
	GetHypervisorGuestOsNames(p *GetHypervisorGuestOsNamesParams) (*GetHypervisorGuestOsNamesResponse, error)
	NewGetHypervisorGuestOsNamesParams(hypervisor string, hypervisorversion string) *GetHypervisorGuestOsNamesParams
}

type AddGuestOsParams struct {
	p map[string]interface{}
}

func (p *AddGuestOsParams) toURLValues() url.Values {
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
	if v, found := p.p["forDisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forDisplay", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
	}
	if v, found := p.p["osdisplayname"]; found {
		u.Set("osdisplayname", v.(string))
	}
	return u
}

func (p *AddGuestOsParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *AddGuestOsParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *AddGuestOsParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *AddGuestOsParams) SetForDisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forDisplay"] = v
}

func (p *AddGuestOsParams) ResetForDisplay() {
	if p.p != nil && p.p["forDisplay"] != nil {
		delete(p.p, "forDisplay")
	}
}

func (p *AddGuestOsParams) GetForDisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forDisplay"].(bool)
	return value, ok
}

func (p *AddGuestOsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddGuestOsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddGuestOsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddGuestOsParams) SetOscategoryid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["oscategoryid"] = v
}

func (p *AddGuestOsParams) ResetOscategoryid() {
	if p.p != nil && p.p["oscategoryid"] != nil {
		delete(p.p, "oscategoryid")
	}
}

func (p *AddGuestOsParams) GetOscategoryid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["oscategoryid"].(string)
	return value, ok
}

func (p *AddGuestOsParams) SetOsdisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osdisplayname"] = v
}

func (p *AddGuestOsParams) ResetOsdisplayname() {
	if p.p != nil && p.p["osdisplayname"] != nil {
		delete(p.p, "osdisplayname")
	}
}

func (p *AddGuestOsParams) GetOsdisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osdisplayname"].(string)
	return value, ok
}

// You should always use this function to get a new AddGuestOsParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewAddGuestOsParams(oscategoryid string, osdisplayname string) *AddGuestOsParams {
	p := &AddGuestOsParams{}
	p.p = make(map[string]interface{})
	p.p["oscategoryid"] = oscategoryid
	p.p["osdisplayname"] = osdisplayname
	return p
}

// Add a new guest OS type
func (s *GuestOSService) AddGuestOs(p *AddGuestOsParams) (*AddGuestOsResponse, error) {
	resp, err := s.cs.newRequest("addGuestOs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddGuestOsResponse
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

type AddGuestOsResponse struct {
	Description    string `json:"description"`
	Fordisplay     bool   `json:"fordisplay"`
	Id             string `json:"id"`
	Isuserdefined  bool   `json:"isuserdefined"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Oscategoryid   string `json:"oscategoryid"`
	Oscategoryname string `json:"oscategoryname"`
}

type AddGuestOsMappingParams struct {
	p map[string]interface{}
}

func (p *AddGuestOsMappingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["hypervisorversion"]; found {
		u.Set("hypervisorversion", v.(string))
	}
	if v, found := p.p["osdisplayname"]; found {
		u.Set("osdisplayname", v.(string))
	}
	if v, found := p.p["osmappingcheckenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("osmappingcheckenabled", vv)
	}
	if v, found := p.p["osnameforhypervisor"]; found {
		u.Set("osnameforhypervisor", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	return u
}

func (p *AddGuestOsMappingParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *AddGuestOsMappingParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *AddGuestOsMappingParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *AddGuestOsMappingParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *AddGuestOsMappingParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *AddGuestOsMappingParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *AddGuestOsMappingParams) SetHypervisorversion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisorversion"] = v
}

func (p *AddGuestOsMappingParams) ResetHypervisorversion() {
	if p.p != nil && p.p["hypervisorversion"] != nil {
		delete(p.p, "hypervisorversion")
	}
}

func (p *AddGuestOsMappingParams) GetHypervisorversion() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisorversion"].(string)
	return value, ok
}

func (p *AddGuestOsMappingParams) SetOsdisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osdisplayname"] = v
}

func (p *AddGuestOsMappingParams) ResetOsdisplayname() {
	if p.p != nil && p.p["osdisplayname"] != nil {
		delete(p.p, "osdisplayname")
	}
}

func (p *AddGuestOsMappingParams) GetOsdisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osdisplayname"].(string)
	return value, ok
}

func (p *AddGuestOsMappingParams) SetOsmappingcheckenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osmappingcheckenabled"] = v
}

func (p *AddGuestOsMappingParams) ResetOsmappingcheckenabled() {
	if p.p != nil && p.p["osmappingcheckenabled"] != nil {
		delete(p.p, "osmappingcheckenabled")
	}
}

func (p *AddGuestOsMappingParams) GetOsmappingcheckenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osmappingcheckenabled"].(bool)
	return value, ok
}

func (p *AddGuestOsMappingParams) SetOsnameforhypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osnameforhypervisor"] = v
}

func (p *AddGuestOsMappingParams) ResetOsnameforhypervisor() {
	if p.p != nil && p.p["osnameforhypervisor"] != nil {
		delete(p.p, "osnameforhypervisor")
	}
}

func (p *AddGuestOsMappingParams) GetOsnameforhypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osnameforhypervisor"].(string)
	return value, ok
}

func (p *AddGuestOsMappingParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *AddGuestOsMappingParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *AddGuestOsMappingParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

// You should always use this function to get a new AddGuestOsMappingParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewAddGuestOsMappingParams(hypervisor string, hypervisorversion string, osnameforhypervisor string) *AddGuestOsMappingParams {
	p := &AddGuestOsMappingParams{}
	p.p = make(map[string]interface{})
	p.p["hypervisor"] = hypervisor
	p.p["hypervisorversion"] = hypervisorversion
	p.p["osnameforhypervisor"] = osnameforhypervisor
	return p
}

// Adds a guest OS name to hypervisor OS name mapping
func (s *GuestOSService) AddGuestOsMapping(p *AddGuestOsMappingParams) (*AddGuestOsMappingResponse, error) {
	resp, err := s.cs.newRequest("addGuestOsMapping", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddGuestOsMappingResponse
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

type AddGuestOsMappingResponse struct {
	Hypervisor          string `json:"hypervisor"`
	Hypervisorversion   string `json:"hypervisorversion"`
	Id                  string `json:"id"`
	Isuserdefined       string `json:"isuserdefined"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Osdisplayname       string `json:"osdisplayname"`
	Osnameforhypervisor string `json:"osnameforhypervisor"`
	Ostypeid            string `json:"ostypeid"`
}

func (r *AddGuestOsMappingResponse) UnmarshalJSON(b []byte) error {
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

	type alias AddGuestOsMappingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListGuestOsMappingParams struct {
	p map[string]interface{}
}

func (p *ListGuestOsMappingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["hypervisorversion"]; found {
		u.Set("hypervisorversion", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["osdisplayname"]; found {
		u.Set("osdisplayname", v.(string))
	}
	if v, found := p.p["osnameforhypervisor"]; found {
		u.Set("osnameforhypervisor", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
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

func (p *ListGuestOsMappingParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListGuestOsMappingParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListGuestOsMappingParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListGuestOsMappingParams) SetHypervisorversion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisorversion"] = v
}

func (p *ListGuestOsMappingParams) ResetHypervisorversion() {
	if p.p != nil && p.p["hypervisorversion"] != nil {
		delete(p.p, "hypervisorversion")
	}
}

func (p *ListGuestOsMappingParams) GetHypervisorversion() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisorversion"].(string)
	return value, ok
}

func (p *ListGuestOsMappingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListGuestOsMappingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListGuestOsMappingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListGuestOsMappingParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListGuestOsMappingParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListGuestOsMappingParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListGuestOsMappingParams) SetOsdisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osdisplayname"] = v
}

func (p *ListGuestOsMappingParams) ResetOsdisplayname() {
	if p.p != nil && p.p["osdisplayname"] != nil {
		delete(p.p, "osdisplayname")
	}
}

func (p *ListGuestOsMappingParams) GetOsdisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osdisplayname"].(string)
	return value, ok
}

func (p *ListGuestOsMappingParams) SetOsnameforhypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osnameforhypervisor"] = v
}

func (p *ListGuestOsMappingParams) ResetOsnameforhypervisor() {
	if p.p != nil && p.p["osnameforhypervisor"] != nil {
		delete(p.p, "osnameforhypervisor")
	}
}

func (p *ListGuestOsMappingParams) GetOsnameforhypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osnameforhypervisor"].(string)
	return value, ok
}

func (p *ListGuestOsMappingParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *ListGuestOsMappingParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *ListGuestOsMappingParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *ListGuestOsMappingParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListGuestOsMappingParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListGuestOsMappingParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListGuestOsMappingParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListGuestOsMappingParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListGuestOsMappingParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListGuestOsMappingParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewListGuestOsMappingParams() *ListGuestOsMappingParams {
	p := &ListGuestOsMappingParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetGuestOsMappingByID(id string, opts ...OptionFunc) (*GuestOsMapping, int, error) {
	p := &ListGuestOsMappingParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListGuestOsMapping(p)
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
		return l.GuestOsMapping[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GuestOsMapping UUID: %s!", id)
}

// Lists all available OS mappings for given hypervisor
func (s *GuestOSService) ListGuestOsMapping(p *ListGuestOsMappingParams) (*ListGuestOsMappingResponse, error) {
	resp, err := s.cs.newRequest("listGuestOsMapping", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGuestOsMappingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListGuestOsMappingResponse struct {
	Count          int               `json:"count"`
	GuestOsMapping []*GuestOsMapping `json:"guestosmapping"`
}

type GuestOsMapping struct {
	Hypervisor          string `json:"hypervisor"`
	Hypervisorversion   string `json:"hypervisorversion"`
	Id                  string `json:"id"`
	Isuserdefined       string `json:"isuserdefined"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Osdisplayname       string `json:"osdisplayname"`
	Osnameforhypervisor string `json:"osnameforhypervisor"`
	Ostypeid            string `json:"ostypeid"`
}

func (r *GuestOsMapping) UnmarshalJSON(b []byte) error {
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

	type alias GuestOsMapping
	return json.Unmarshal(b, (*alias)(r))
}

type ListOsCategoriesParams struct {
	p map[string]interface{}
}

func (p *ListOsCategoriesParams) toURLValues() url.Values {
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
	return u
}

func (p *ListOsCategoriesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListOsCategoriesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListOsCategoriesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListOsCategoriesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListOsCategoriesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListOsCategoriesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListOsCategoriesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListOsCategoriesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListOsCategoriesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListOsCategoriesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListOsCategoriesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListOsCategoriesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListOsCategoriesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListOsCategoriesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListOsCategoriesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListOsCategoriesParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewListOsCategoriesParams() *ListOsCategoriesParams {
	p := &ListOsCategoriesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListOsCategoriesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListOsCategories(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.OsCategories[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.OsCategories {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryByName(name string, opts ...OptionFunc) (*OsCategory, int, error) {
	id, count, err := s.GetOsCategoryID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetOsCategoryByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryByID(id string, opts ...OptionFunc) (*OsCategory, int, error) {
	p := &ListOsCategoriesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOsCategories(p)
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
		return l.OsCategories[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for OsCategory UUID: %s!", id)
}

// Lists all supported OS categories for this cloud.
func (s *GuestOSService) ListOsCategories(p *ListOsCategoriesParams) (*ListOsCategoriesResponse, error) {
	resp, err := s.cs.newRequest("listOsCategories", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListOsCategoriesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListOsCategoriesResponse struct {
	Count        int           `json:"count"`
	OsCategories []*OsCategory `json:"oscategory"`
}

type OsCategory struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListOsTypesParams struct {
	p map[string]interface{}
}

func (p *ListOsTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
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

func (p *ListOsTypesParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *ListOsTypesParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *ListOsTypesParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *ListOsTypesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListOsTypesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListOsTypesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListOsTypesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListOsTypesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListOsTypesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListOsTypesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListOsTypesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListOsTypesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListOsTypesParams) SetOscategoryid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["oscategoryid"] = v
}

func (p *ListOsTypesParams) ResetOscategoryid() {
	if p.p != nil && p.p["oscategoryid"] != nil {
		delete(p.p, "oscategoryid")
	}
}

func (p *ListOsTypesParams) GetOscategoryid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["oscategoryid"].(string)
	return value, ok
}

func (p *ListOsTypesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListOsTypesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListOsTypesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListOsTypesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListOsTypesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListOsTypesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListOsTypesParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewListOsTypesParams() *ListOsTypesParams {
	p := &ListOsTypesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsTypeID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListOsTypesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListOsTypes(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.OsTypes[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.OsTypes {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsTypeByName(name string, opts ...OptionFunc) (*OsType, int, error) {
	id, count, err := s.GetOsTypeID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetOsTypeByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsTypeByID(id string, opts ...OptionFunc) (*OsType, int, error) {
	p := &ListOsTypesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOsTypes(p)
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
		return l.OsTypes[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for OsType UUID: %s!", id)
}

// Lists all supported OS types for this cloud.
func (s *GuestOSService) ListOsTypes(p *ListOsTypesParams) (*ListOsTypesResponse, error) {
	resp, err := s.cs.newRequest("listOsTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListOsTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListOsTypesResponse struct {
	Count   int       `json:"count"`
	OsTypes []*OsType `json:"ostype"`
}

type OsType struct {
	Description    string `json:"description"`
	Fordisplay     bool   `json:"fordisplay"`
	Id             string `json:"id"`
	Isuserdefined  bool   `json:"isuserdefined"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Oscategoryid   string `json:"oscategoryid"`
	Oscategoryname string `json:"oscategoryname"`
}

type RemoveGuestOsParams struct {
	p map[string]interface{}
}

func (p *RemoveGuestOsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RemoveGuestOsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RemoveGuestOsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RemoveGuestOsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveGuestOsParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewRemoveGuestOsParams(id string) *RemoveGuestOsParams {
	p := &RemoveGuestOsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes a Guest OS from listing.
func (s *GuestOSService) RemoveGuestOs(p *RemoveGuestOsParams) (*RemoveGuestOsResponse, error) {
	resp, err := s.cs.newRequest("removeGuestOs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveGuestOsResponse
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

type RemoveGuestOsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RemoveGuestOsMappingParams struct {
	p map[string]interface{}
}

func (p *RemoveGuestOsMappingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RemoveGuestOsMappingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RemoveGuestOsMappingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RemoveGuestOsMappingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveGuestOsMappingParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewRemoveGuestOsMappingParams(id string) *RemoveGuestOsMappingParams {
	p := &RemoveGuestOsMappingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes a Guest OS Mapping.
func (s *GuestOSService) RemoveGuestOsMapping(p *RemoveGuestOsMappingParams) (*RemoveGuestOsMappingResponse, error) {
	resp, err := s.cs.newRequest("removeGuestOsMapping", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveGuestOsMappingResponse
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

type RemoveGuestOsMappingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateGuestOsParams struct {
	p map[string]interface{}
}

func (p *UpdateGuestOsParams) toURLValues() url.Values {
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
	if v, found := p.p["forDisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forDisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["osdisplayname"]; found {
		u.Set("osdisplayname", v.(string))
	}
	return u
}

func (p *UpdateGuestOsParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateGuestOsParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateGuestOsParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateGuestOsParams) SetForDisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forDisplay"] = v
}

func (p *UpdateGuestOsParams) ResetForDisplay() {
	if p.p != nil && p.p["forDisplay"] != nil {
		delete(p.p, "forDisplay")
	}
}

func (p *UpdateGuestOsParams) GetForDisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forDisplay"].(bool)
	return value, ok
}

func (p *UpdateGuestOsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateGuestOsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateGuestOsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateGuestOsParams) SetOsdisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osdisplayname"] = v
}

func (p *UpdateGuestOsParams) ResetOsdisplayname() {
	if p.p != nil && p.p["osdisplayname"] != nil {
		delete(p.p, "osdisplayname")
	}
}

func (p *UpdateGuestOsParams) GetOsdisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osdisplayname"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateGuestOsParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewUpdateGuestOsParams(id string, osdisplayname string) *UpdateGuestOsParams {
	p := &UpdateGuestOsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["osdisplayname"] = osdisplayname
	return p
}

// Updates the information about Guest OS
func (s *GuestOSService) UpdateGuestOs(p *UpdateGuestOsParams) (*UpdateGuestOsResponse, error) {
	resp, err := s.cs.newRequest("updateGuestOs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGuestOsResponse
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

type UpdateGuestOsResponse struct {
	Description    string `json:"description"`
	Fordisplay     bool   `json:"fordisplay"`
	Id             string `json:"id"`
	Isuserdefined  bool   `json:"isuserdefined"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Oscategoryid   string `json:"oscategoryid"`
	Oscategoryname string `json:"oscategoryname"`
}

type UpdateGuestOsMappingParams struct {
	p map[string]interface{}
}

func (p *UpdateGuestOsMappingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["osmappingcheckenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("osmappingcheckenabled", vv)
	}
	if v, found := p.p["osnameforhypervisor"]; found {
		u.Set("osnameforhypervisor", v.(string))
	}
	return u
}

func (p *UpdateGuestOsMappingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateGuestOsMappingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateGuestOsMappingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateGuestOsMappingParams) SetOsmappingcheckenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osmappingcheckenabled"] = v
}

func (p *UpdateGuestOsMappingParams) ResetOsmappingcheckenabled() {
	if p.p != nil && p.p["osmappingcheckenabled"] != nil {
		delete(p.p, "osmappingcheckenabled")
	}
}

func (p *UpdateGuestOsMappingParams) GetOsmappingcheckenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osmappingcheckenabled"].(bool)
	return value, ok
}

func (p *UpdateGuestOsMappingParams) SetOsnameforhypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["osnameforhypervisor"] = v
}

func (p *UpdateGuestOsMappingParams) ResetOsnameforhypervisor() {
	if p.p != nil && p.p["osnameforhypervisor"] != nil {
		delete(p.p, "osnameforhypervisor")
	}
}

func (p *UpdateGuestOsMappingParams) GetOsnameforhypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["osnameforhypervisor"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateGuestOsMappingParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewUpdateGuestOsMappingParams(id string, osnameforhypervisor string) *UpdateGuestOsMappingParams {
	p := &UpdateGuestOsMappingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["osnameforhypervisor"] = osnameforhypervisor
	return p
}

// Updates the information about Guest OS to Hypervisor specific name mapping
func (s *GuestOSService) UpdateGuestOsMapping(p *UpdateGuestOsMappingParams) (*UpdateGuestOsMappingResponse, error) {
	resp, err := s.cs.newRequest("updateGuestOsMapping", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGuestOsMappingResponse
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

type UpdateGuestOsMappingResponse struct {
	Hypervisor          string `json:"hypervisor"`
	Hypervisorversion   string `json:"hypervisorversion"`
	Id                  string `json:"id"`
	Isuserdefined       string `json:"isuserdefined"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Osdisplayname       string `json:"osdisplayname"`
	Osnameforhypervisor string `json:"osnameforhypervisor"`
	Ostypeid            string `json:"ostypeid"`
}

func (r *UpdateGuestOsMappingResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateGuestOsMappingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type GetHypervisorGuestOsNamesParams struct {
	p map[string]interface{}
}

func (p *GetHypervisorGuestOsNamesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["hypervisorversion"]; found {
		u.Set("hypervisorversion", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	return u
}

func (p *GetHypervisorGuestOsNamesParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *GetHypervisorGuestOsNamesParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *GetHypervisorGuestOsNamesParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *GetHypervisorGuestOsNamesParams) SetHypervisorversion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisorversion"] = v
}

func (p *GetHypervisorGuestOsNamesParams) ResetHypervisorversion() {
	if p.p != nil && p.p["hypervisorversion"] != nil {
		delete(p.p, "hypervisorversion")
	}
}

func (p *GetHypervisorGuestOsNamesParams) GetHypervisorversion() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisorversion"].(string)
	return value, ok
}

func (p *GetHypervisorGuestOsNamesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *GetHypervisorGuestOsNamesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *GetHypervisorGuestOsNamesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

// You should always use this function to get a new GetHypervisorGuestOsNamesParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewGetHypervisorGuestOsNamesParams(hypervisor string, hypervisorversion string) *GetHypervisorGuestOsNamesParams {
	p := &GetHypervisorGuestOsNamesParams{}
	p.p = make(map[string]interface{})
	p.p["hypervisor"] = hypervisor
	p.p["hypervisorversion"] = hypervisorversion
	return p
}

// Gets the guest OS names in the hypervisor
func (s *GuestOSService) GetHypervisorGuestOsNames(p *GetHypervisorGuestOsNamesParams) (*GetHypervisorGuestOsNamesResponse, error) {
	resp, err := s.cs.newRequest("getHypervisorGuestOsNames", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetHypervisorGuestOsNamesResponse
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

type GetHypervisorGuestOsNamesResponse struct {
	Guestoscount      int                                            `json:"guestoscount"`
	Guestoslist       []GetHypervisorGuestOsNamesResponseGuestoslist `json:"guestoslist"`
	Hypervisor        string                                         `json:"hypervisor"`
	Hypervisorversion string                                         `json:"hypervisorversion"`
	JobID             string                                         `json:"jobid"`
	Jobstatus         int                                            `json:"jobstatus"`
}

type GetHypervisorGuestOsNamesResponseGuestoslist struct {
	Osdisplayname       string `json:"osdisplayname"`
	Osnameforhypervisor string `json:"osnameforhypervisor"`
}
