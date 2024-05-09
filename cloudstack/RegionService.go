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
	"net/url"
	"strconv"
)

type RegionServiceIface interface {
	AddRegion(p *AddRegionParams) (*AddRegionResponse, error)
	NewAddRegionParams(endpoint string, id int, name string) *AddRegionParams
	ListRegions(p *ListRegionsParams) (*ListRegionsResponse, error)
	NewListRegionsParams() *ListRegionsParams
	RemoveRegion(p *RemoveRegionParams) (*RemoveRegionResponse, error)
	NewRemoveRegionParams(id int) *RemoveRegionParams
	UpdateRegion(p *UpdateRegionParams) (*UpdateRegionResponse, error)
	NewUpdateRegionParams(id int) *UpdateRegionParams
}

type AddRegionParams struct {
	p map[string]interface{}
}

func (p *AddRegionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["endpoint"]; found {
		u.Set("endpoint", v.(string))
	}
	if v, found := p.p["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *AddRegionParams) SetEndpoint(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endpoint"] = v
}

func (p *AddRegionParams) ResetEndpoint() {
	if p.p != nil && p.p["endpoint"] != nil {
		delete(p.p, "endpoint")
	}
}

func (p *AddRegionParams) GetEndpoint() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endpoint"].(string)
	return value, ok
}

func (p *AddRegionParams) SetId(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *AddRegionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *AddRegionParams) GetId() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int)
	return value, ok
}

func (p *AddRegionParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddRegionParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddRegionParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new AddRegionParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewAddRegionParams(endpoint string, id int, name string) *AddRegionParams {
	p := &AddRegionParams{}
	p.p = make(map[string]interface{})
	p.p["endpoint"] = endpoint
	p.p["id"] = id
	p.p["name"] = name
	return p
}

// Adds a Region
func (s *RegionService) AddRegion(p *AddRegionParams) (*AddRegionResponse, error) {
	resp, err := s.cs.newRequest("addRegion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddRegionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddRegionResponse struct {
	Endpoint                 string `json:"endpoint"`
	Gslbserviceenabled       bool   `json:"gslbserviceenabled"`
	Id                       int    `json:"id"`
	JobID                    string `json:"jobid"`
	Jobstatus                int    `json:"jobstatus"`
	Name                     string `json:"name"`
	Portableipserviceenabled bool   `json:"portableipserviceenabled"`
}

type ListRegionsParams struct {
	p map[string]interface{}
}

func (p *ListRegionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
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

func (p *ListRegionsParams) SetId(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListRegionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListRegionsParams) GetId() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int)
	return value, ok
}

func (p *ListRegionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListRegionsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListRegionsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListRegionsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListRegionsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListRegionsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListRegionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListRegionsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListRegionsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListRegionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListRegionsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListRegionsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListRegionsParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewListRegionsParams() *ListRegionsParams {
	p := &ListRegionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Regions
func (s *RegionService) ListRegions(p *ListRegionsParams) (*ListRegionsResponse, error) {
	resp, err := s.cs.newRequest("listRegions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRegionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRegionsResponse struct {
	Count   int       `json:"count"`
	Regions []*Region `json:"region"`
}

type Region struct {
	Endpoint                 string `json:"endpoint"`
	Gslbserviceenabled       bool   `json:"gslbserviceenabled"`
	Id                       int    `json:"id"`
	JobID                    string `json:"jobid"`
	Jobstatus                int    `json:"jobstatus"`
	Name                     string `json:"name"`
	Portableipserviceenabled bool   `json:"portableipserviceenabled"`
}

type RemoveRegionParams struct {
	p map[string]interface{}
}

func (p *RemoveRegionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	return u
}

func (p *RemoveRegionParams) SetId(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RemoveRegionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RemoveRegionParams) GetId() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int)
	return value, ok
}

// You should always use this function to get a new RemoveRegionParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewRemoveRegionParams(id int) *RemoveRegionParams {
	p := &RemoveRegionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes specified region
func (s *RegionService) RemoveRegion(p *RemoveRegionParams) (*RemoveRegionResponse, error) {
	resp, err := s.cs.newRequest("removeRegion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveRegionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveRegionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RemoveRegionResponse) UnmarshalJSON(b []byte) error {
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

	type alias RemoveRegionResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateRegionParams struct {
	p map[string]interface{}
}

func (p *UpdateRegionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["endpoint"]; found {
		u.Set("endpoint", v.(string))
	}
	if v, found := p.p["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateRegionParams) SetEndpoint(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endpoint"] = v
}

func (p *UpdateRegionParams) ResetEndpoint() {
	if p.p != nil && p.p["endpoint"] != nil {
		delete(p.p, "endpoint")
	}
}

func (p *UpdateRegionParams) GetEndpoint() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endpoint"].(string)
	return value, ok
}

func (p *UpdateRegionParams) SetId(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateRegionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateRegionParams) GetId() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int)
	return value, ok
}

func (p *UpdateRegionParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateRegionParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateRegionParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateRegionParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewUpdateRegionParams(id int) *UpdateRegionParams {
	p := &UpdateRegionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a region
func (s *RegionService) UpdateRegion(p *UpdateRegionParams) (*UpdateRegionResponse, error) {
	resp, err := s.cs.newRequest("updateRegion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateRegionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateRegionResponse struct {
	Endpoint                 string `json:"endpoint"`
	Gslbserviceenabled       bool   `json:"gslbserviceenabled"`
	Id                       int    `json:"id"`
	JobID                    string `json:"jobid"`
	Jobstatus                int    `json:"jobstatus"`
	Name                     string `json:"name"`
	Portableipserviceenabled bool   `json:"portableipserviceenabled"`
}
