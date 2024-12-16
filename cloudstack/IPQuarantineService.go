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

type IPQuarantineServiceIface interface {
	ListQuarantinedIps(p *ListQuarantinedIpsParams) (*ListQuarantinedIpsResponse, error)
	NewListQuarantinedIpsParams() *ListQuarantinedIpsParams
	RemoveQuarantinedIp(p *RemoveQuarantinedIpParams) (*RemoveQuarantinedIpResponse, error)
	NewRemoveQuarantinedIpParams(removalreason string) *RemoveQuarantinedIpParams
	UpdateQuarantinedIp(p *UpdateQuarantinedIpParams) (*UpdateQuarantinedIpResponse, error)
	NewUpdateQuarantinedIpParams(enddate string) *UpdateQuarantinedIpParams
}

type ListQuarantinedIpsParams struct {
	p map[string]interface{}
}

func (p *ListQuarantinedIpsParams) toURLValues() url.Values {
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
	if v, found := p.p["showinactive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showinactive", vv)
	}
	if v, found := p.p["showremoved"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showremoved", vv)
	}
	return u
}

func (p *ListQuarantinedIpsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListQuarantinedIpsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListQuarantinedIpsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListQuarantinedIpsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListQuarantinedIpsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListQuarantinedIpsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListQuarantinedIpsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListQuarantinedIpsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListQuarantinedIpsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListQuarantinedIpsParams) SetShowinactive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showinactive"] = v
}

func (p *ListQuarantinedIpsParams) ResetShowinactive() {
	if p.p != nil && p.p["showinactive"] != nil {
		delete(p.p, "showinactive")
	}
}

func (p *ListQuarantinedIpsParams) GetShowinactive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showinactive"].(bool)
	return value, ok
}

func (p *ListQuarantinedIpsParams) SetShowremoved(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showremoved"] = v
}

func (p *ListQuarantinedIpsParams) ResetShowremoved() {
	if p.p != nil && p.p["showremoved"] != nil {
		delete(p.p, "showremoved")
	}
}

func (p *ListQuarantinedIpsParams) GetShowremoved() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showremoved"].(bool)
	return value, ok
}

// You should always use this function to get a new ListQuarantinedIpsParams instance,
// as then you are sure you have configured all required params
func (s *IPQuarantineService) NewListQuarantinedIpsParams() *ListQuarantinedIpsParams {
	p := &ListQuarantinedIpsParams{}
	p.p = make(map[string]interface{})
	return p
}

// List public IP addresses in quarantine.
func (s *IPQuarantineService) ListQuarantinedIps(p *ListQuarantinedIpsParams) (*ListQuarantinedIpsResponse, error) {
	resp, err := s.cs.newRequest("listQuarantinedIps", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListQuarantinedIpsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListQuarantinedIpsResponse struct {
	Count          int              `json:"count"`
	QuarantinedIps []*QuarantinedIp `json:"quarantinedip"`
}

type QuarantinedIp struct {
	Created           string `json:"created"`
	Enddate           string `json:"enddate"`
	Id                string `json:"id"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Previousownerid   string `json:"previousownerid"`
	Previousownername string `json:"previousownername"`
	Removalreason     string `json:"removalreason"`
	Removed           string `json:"removed"`
	Removeraccountid  string `json:"removeraccountid"`
}

type RemoveQuarantinedIpParams struct {
	p map[string]interface{}
}

func (p *RemoveQuarantinedIpParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["removalreason"]; found {
		u.Set("removalreason", v.(string))
	}
	return u
}

func (p *RemoveQuarantinedIpParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RemoveQuarantinedIpParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RemoveQuarantinedIpParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *RemoveQuarantinedIpParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *RemoveQuarantinedIpParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *RemoveQuarantinedIpParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *RemoveQuarantinedIpParams) SetRemovalreason(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["removalreason"] = v
}

func (p *RemoveQuarantinedIpParams) ResetRemovalreason() {
	if p.p != nil && p.p["removalreason"] != nil {
		delete(p.p, "removalreason")
	}
}

func (p *RemoveQuarantinedIpParams) GetRemovalreason() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["removalreason"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveQuarantinedIpParams instance,
// as then you are sure you have configured all required params
func (s *IPQuarantineService) NewRemoveQuarantinedIpParams(removalreason string) *RemoveQuarantinedIpParams {
	p := &RemoveQuarantinedIpParams{}
	p.p = make(map[string]interface{})
	p.p["removalreason"] = removalreason
	return p
}

// Removes a public IP address from quarantine. Only IPs in active quarantine can be removed.
func (s *IPQuarantineService) RemoveQuarantinedIp(p *RemoveQuarantinedIpParams) (*RemoveQuarantinedIpResponse, error) {
	resp, err := s.cs.newRequest("removeQuarantinedIp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveQuarantinedIpResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveQuarantinedIpResponse struct {
	Created           string `json:"created"`
	Enddate           string `json:"enddate"`
	Id                string `json:"id"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Previousownerid   string `json:"previousownerid"`
	Previousownername string `json:"previousownername"`
	Removalreason     string `json:"removalreason"`
	Removed           string `json:"removed"`
	Removeraccountid  string `json:"removeraccountid"`
}

type UpdateQuarantinedIpParams struct {
	p map[string]interface{}
}

func (p *UpdateQuarantinedIpParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	return u
}

func (p *UpdateQuarantinedIpParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *UpdateQuarantinedIpParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *UpdateQuarantinedIpParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *UpdateQuarantinedIpParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateQuarantinedIpParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateQuarantinedIpParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateQuarantinedIpParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *UpdateQuarantinedIpParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *UpdateQuarantinedIpParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateQuarantinedIpParams instance,
// as then you are sure you have configured all required params
func (s *IPQuarantineService) NewUpdateQuarantinedIpParams(enddate string) *UpdateQuarantinedIpParams {
	p := &UpdateQuarantinedIpParams{}
	p.p = make(map[string]interface{})
	p.p["enddate"] = enddate
	return p
}

// Updates the quarantine end date for the given public IP address.
func (s *IPQuarantineService) UpdateQuarantinedIp(p *UpdateQuarantinedIpParams) (*UpdateQuarantinedIpResponse, error) {
	resp, err := s.cs.newRequest("updateQuarantinedIp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateQuarantinedIpResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateQuarantinedIpResponse struct {
	Created           string `json:"created"`
	Enddate           string `json:"enddate"`
	Id                string `json:"id"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Previousownerid   string `json:"previousownerid"`
	Previousownername string `json:"previousownername"`
	Removalreason     string `json:"removalreason"`
	Removed           string `json:"removed"`
	Removeraccountid  string `json:"removeraccountid"`
}
