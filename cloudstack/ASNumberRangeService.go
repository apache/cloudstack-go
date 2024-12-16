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

type ASNumberRangeServiceIface interface {
	CreateASNRange(p *CreateASNRangeParams) (*CreateASNRangeResponse, error)
	NewCreateASNRangeParams(endasn int64, startasn int64, zoneid string) *CreateASNRangeParams
	DeleteASNRange(p *DeleteASNRangeParams) (*DeleteASNRangeResponse, error)
	NewDeleteASNRangeParams(id string) *DeleteASNRangeParams
	ListASNRanges(p *ListASNRangesParams) (*ListASNRangesResponse, error)
	NewListASNRangesParams() *ListASNRangesParams
}

type CreateASNRangeParams struct {
	p map[string]interface{}
}

func (p *CreateASNRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["endasn"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("endasn", vv)
	}
	if v, found := p.p["startasn"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("startasn", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateASNRangeParams) SetEndasn(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endasn"] = v
}

func (p *CreateASNRangeParams) ResetEndasn() {
	if p.p != nil && p.p["endasn"] != nil {
		delete(p.p, "endasn")
	}
}

func (p *CreateASNRangeParams) GetEndasn() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endasn"].(int64)
	return value, ok
}

func (p *CreateASNRangeParams) SetStartasn(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startasn"] = v
}

func (p *CreateASNRangeParams) ResetStartasn() {
	if p.p != nil && p.p["startasn"] != nil {
		delete(p.p, "startasn")
	}
}

func (p *CreateASNRangeParams) GetStartasn() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startasn"].(int64)
	return value, ok
}

func (p *CreateASNRangeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateASNRangeParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateASNRangeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateASNRangeParams instance,
// as then you are sure you have configured all required params
func (s *ASNumberRangeService) NewCreateASNRangeParams(endasn int64, startasn int64, zoneid string) *CreateASNRangeParams {
	p := &CreateASNRangeParams{}
	p.p = make(map[string]interface{})
	p.p["endasn"] = endasn
	p.p["startasn"] = startasn
	p.p["zoneid"] = zoneid
	return p
}

// Creates a range of Autonomous Systems for BGP Dynamic Routing
func (s *ASNumberRangeService) CreateASNRange(p *CreateASNRangeParams) (*CreateASNRangeResponse, error) {
	resp, err := s.cs.newRequest("createASNRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateASNRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateASNRangeResponse struct {
	Created   string `json:"created"`
	Endasn    int64  `json:"endasn"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Startasn  int64  `json:"startasn"`
	Zoneid    string `json:"zoneid"`
}

type DeleteASNRangeParams struct {
	p map[string]interface{}
}

func (p *DeleteASNRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteASNRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteASNRangeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteASNRangeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteASNRangeParams instance,
// as then you are sure you have configured all required params
func (s *ASNumberRangeService) NewDeleteASNRangeParams(id string) *DeleteASNRangeParams {
	p := &DeleteASNRangeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// deletes a range of Autonomous Systems for BGP Dynamic Routing
func (s *ASNumberRangeService) DeleteASNRange(p *DeleteASNRangeParams) (*DeleteASNRangeResponse, error) {
	resp, err := s.cs.newRequest("deleteASNRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteASNRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteASNRangeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteASNRangeResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteASNRangeResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListASNRangesParams struct {
	p map[string]interface{}
}

func (p *ListASNRangesParams) toURLValues() url.Values {
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListASNRangesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListASNRangesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListASNRangesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListASNRangesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListASNRangesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListASNRangesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListASNRangesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListASNRangesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListASNRangesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListASNRangesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListASNRangesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListASNRangesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListASNRangesParams instance,
// as then you are sure you have configured all required params
func (s *ASNumberRangeService) NewListASNRangesParams() *ListASNRangesParams {
	p := &ListASNRangesParams{}
	p.p = make(map[string]interface{})
	return p
}

// List Autonomous Systems Number Ranges
func (s *ASNumberRangeService) ListASNRanges(p *ListASNRangesParams) (*ListASNRangesResponse, error) {
	resp, err := s.cs.newRequest("listASNRanges", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListASNRangesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListASNRangesResponse struct {
	Count     int         `json:"count"`
	ASNRanges []*ASNRange `json:"asnrange"`
}

type ASNRange struct {
	Created   string `json:"created"`
	Endasn    int64  `json:"endasn"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Startasn  int64  `json:"startasn"`
	Zoneid    string `json:"zoneid"`
}
