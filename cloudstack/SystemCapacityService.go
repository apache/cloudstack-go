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

type SystemCapacityServiceIface interface {
	ListCapacity(p *ListCapacityParams) (*ListCapacityResponse, error)
	NewListCapacityParams() *ListCapacityParams
}

type ListCapacityParams struct {
	p map[string]interface{}
}

func (p *ListCapacityParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["fetchlatest"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fetchlatest", vv)
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
	if v, found := p.p["sortby"]; found {
		u.Set("sortby", v.(string))
	}
	if v, found := p.p["tag"]; found {
		u.Set("tag", v.(string))
	}
	if v, found := p.p["type"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("type", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListCapacityParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListCapacityParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListCapacityParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListCapacityParams) SetFetchlatest(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fetchlatest"] = v
}

func (p *ListCapacityParams) ResetFetchlatest() {
	if p.p != nil && p.p["fetchlatest"] != nil {
		delete(p.p, "fetchlatest")
	}
}

func (p *ListCapacityParams) GetFetchlatest() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fetchlatest"].(bool)
	return value, ok
}

func (p *ListCapacityParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListCapacityParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListCapacityParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListCapacityParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListCapacityParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListCapacityParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListCapacityParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListCapacityParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListCapacityParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListCapacityParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListCapacityParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListCapacityParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListCapacityParams) SetSortby(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortby"] = v
}

func (p *ListCapacityParams) ResetSortby() {
	if p.p != nil && p.p["sortby"] != nil {
		delete(p.p, "sortby")
	}
}

func (p *ListCapacityParams) GetSortby() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortby"].(string)
	return value, ok
}

func (p *ListCapacityParams) SetTag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tag"] = v
}

func (p *ListCapacityParams) ResetTag() {
	if p.p != nil && p.p["tag"] != nil {
		delete(p.p, "tag")
	}
}

func (p *ListCapacityParams) GetTag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tag"].(string)
	return value, ok
}

func (p *ListCapacityParams) SetType(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListCapacityParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ListCapacityParams) GetType() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(int)
	return value, ok
}

func (p *ListCapacityParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListCapacityParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListCapacityParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListCapacityParams instance,
// as then you are sure you have configured all required params
func (s *SystemCapacityService) NewListCapacityParams() *ListCapacityParams {
	p := &ListCapacityParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all the system wide capacities.
func (s *SystemCapacityService) ListCapacity(p *ListCapacityParams) (*ListCapacityResponse, error) {
	resp, err := s.cs.newRequest("listCapacity", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCapacityResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCapacityResponse struct {
	Count    int         `json:"count"`
	Capacity []*Capacity `json:"capacity"`
}

type Capacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Tag               string `json:"tag"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}
