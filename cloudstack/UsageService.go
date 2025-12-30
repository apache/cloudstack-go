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
)

type UsageServiceIface interface {
	AddTrafficMonitor(p *AddTrafficMonitorParams) (*AddTrafficMonitorResponse, error)
	NewAddTrafficMonitorParams(url string, zoneid string) *AddTrafficMonitorParams
	AddTrafficType(p *AddTrafficTypeParams) (*AddTrafficTypeResponse, error)
	NewAddTrafficTypeParams(physicalnetworkid string, traffictype string) *AddTrafficTypeParams
	DeleteTrafficMonitor(p *DeleteTrafficMonitorParams) (*DeleteTrafficMonitorResponse, error)
	NewDeleteTrafficMonitorParams(id string) *DeleteTrafficMonitorParams
	DeleteTrafficType(p *DeleteTrafficTypeParams) (*DeleteTrafficTypeResponse, error)
	NewDeleteTrafficTypeParams(id string) *DeleteTrafficTypeParams
	GenerateUsageRecords(p *GenerateUsageRecordsParams) (*GenerateUsageRecordsResponse, error)
	NewGenerateUsageRecordsParams() *GenerateUsageRecordsParams
	ListTrafficMonitors(p *ListTrafficMonitorsParams) (*ListTrafficMonitorsResponse, error)
	NewListTrafficMonitorsParams(zoneid string) *ListTrafficMonitorsParams
	ListTrafficTypeImplementors(p *ListTrafficTypeImplementorsParams) (*ListTrafficTypeImplementorsResponse, error)
	NewListTrafficTypeImplementorsParams() *ListTrafficTypeImplementorsParams
	ListTrafficTypes(p *ListTrafficTypesParams) (*ListTrafficTypesResponse, error)
	NewListTrafficTypesParams(physicalnetworkid string) *ListTrafficTypesParams
	GetTrafficTypeID(keyword string, physicalnetworkid string, opts ...OptionFunc) (string, int, error)
	ListUsageRecords(p *ListUsageRecordsParams) (*ListUsageRecordsResponse, error)
	NewListUsageRecordsParams(enddate string, startdate string) *ListUsageRecordsParams
	ListUsageTypes(p *ListUsageTypesParams) (*ListUsageTypesResponse, error)
	NewListUsageTypesParams() *ListUsageTypesParams
	RemoveRawUsageRecords(p *RemoveRawUsageRecordsParams) (*RemoveRawUsageRecordsResponse, error)
	NewRemoveRawUsageRecordsParams(interval int) *RemoveRawUsageRecordsParams
	UpdateTrafficType(p *UpdateTrafficTypeParams) (*UpdateTrafficTypeResponse, error)
	NewUpdateTrafficTypeParams(id string) *UpdateTrafficTypeParams
	ListUsageServerMetrics(p *ListUsageServerMetricsParams) (*ListUsageServerMetricsResponse, error)
	NewListUsageServerMetricsParams() *ListUsageServerMetricsParams
}

type AddTrafficMonitorParams struct {
	p map[string]interface{}
}

func (p *AddTrafficMonitorParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["excludezones"]; found {
		u.Set("excludezones", v.(string))
	}
	if v, found := p.p["includezones"]; found {
		u.Set("includezones", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddTrafficMonitorParams) SetExcludezones(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["excludezones"] = v
}

func (p *AddTrafficMonitorParams) ResetExcludezones() {
	if p.p != nil && p.p["excludezones"] != nil {
		delete(p.p, "excludezones")
	}
}

func (p *AddTrafficMonitorParams) GetExcludezones() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["excludezones"].(string)
	return value, ok
}

func (p *AddTrafficMonitorParams) SetIncludezones(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["includezones"] = v
}

func (p *AddTrafficMonitorParams) ResetIncludezones() {
	if p.p != nil && p.p["includezones"] != nil {
		delete(p.p, "includezones")
	}
}

func (p *AddTrafficMonitorParams) GetIncludezones() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["includezones"].(string)
	return value, ok
}

func (p *AddTrafficMonitorParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddTrafficMonitorParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddTrafficMonitorParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddTrafficMonitorParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddTrafficMonitorParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddTrafficMonitorParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddTrafficMonitorParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewAddTrafficMonitorParams(url string, zoneid string) *AddTrafficMonitorParams {
	p := &AddTrafficMonitorParams{}
	p.p = make(map[string]interface{})
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Adds Traffic Monitor Host for Direct Network Usage
func (s *UsageService) AddTrafficMonitor(p *AddTrafficMonitorParams) (*AddTrafficMonitorResponse, error) {
	resp, err := s.cs.newPostRequest("addTrafficMonitor", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddTrafficMonitorResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddTrafficMonitorResponse struct {
	Id         string `json:"id"`
	Ipaddress  string `json:"ipaddress"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Numretries string `json:"numretries"`
	Timeout    string `json:"timeout"`
	Zoneid     string `json:"zoneid"`
}

type AddTrafficTypeParams struct {
	p map[string]interface{}
}

func (p *AddTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hypervnetworklabel"]; found {
		u.Set("hypervnetworklabel", v.(string))
	}
	if v, found := p.p["isolationmethod"]; found {
		u.Set("isolationmethod", v.(string))
	}
	if v, found := p.p["kvmnetworklabel"]; found {
		u.Set("kvmnetworklabel", v.(string))
	}
	if v, found := p.p["ovm3networklabel"]; found {
		u.Set("ovm3networklabel", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["vmwarenetworklabel"]; found {
		u.Set("vmwarenetworklabel", v.(string))
	}
	if v, found := p.p["xennetworklabel"]; found {
		u.Set("xennetworklabel", v.(string))
	}
	return u
}

func (p *AddTrafficTypeParams) SetHypervnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervnetworklabel"] = v
}

func (p *AddTrafficTypeParams) ResetHypervnetworklabel() {
	if p.p != nil && p.p["hypervnetworklabel"] != nil {
		delete(p.p, "hypervnetworklabel")
	}
}

func (p *AddTrafficTypeParams) GetHypervnetworklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervnetworklabel"].(string)
	return value, ok
}

func (p *AddTrafficTypeParams) SetIsolationmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isolationmethod"] = v
}

func (p *AddTrafficTypeParams) ResetIsolationmethod() {
	if p.p != nil && p.p["isolationmethod"] != nil {
		delete(p.p, "isolationmethod")
	}
}

func (p *AddTrafficTypeParams) GetIsolationmethod() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isolationmethod"].(string)
	return value, ok
}

func (p *AddTrafficTypeParams) SetKvmnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["kvmnetworklabel"] = v
}

func (p *AddTrafficTypeParams) ResetKvmnetworklabel() {
	if p.p != nil && p.p["kvmnetworklabel"] != nil {
		delete(p.p, "kvmnetworklabel")
	}
}

func (p *AddTrafficTypeParams) GetKvmnetworklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["kvmnetworklabel"].(string)
	return value, ok
}

func (p *AddTrafficTypeParams) SetOvm3networklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ovm3networklabel"] = v
}

func (p *AddTrafficTypeParams) ResetOvm3networklabel() {
	if p.p != nil && p.p["ovm3networklabel"] != nil {
		delete(p.p, "ovm3networklabel")
	}
}

func (p *AddTrafficTypeParams) GetOvm3networklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ovm3networklabel"].(string)
	return value, ok
}

func (p *AddTrafficTypeParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddTrafficTypeParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *AddTrafficTypeParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddTrafficTypeParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *AddTrafficTypeParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
}

func (p *AddTrafficTypeParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

func (p *AddTrafficTypeParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *AddTrafficTypeParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
}

func (p *AddTrafficTypeParams) GetVlan() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(string)
	return value, ok
}

func (p *AddTrafficTypeParams) SetVmwarenetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmwarenetworklabel"] = v
}

func (p *AddTrafficTypeParams) ResetVmwarenetworklabel() {
	if p.p != nil && p.p["vmwarenetworklabel"] != nil {
		delete(p.p, "vmwarenetworklabel")
	}
}

func (p *AddTrafficTypeParams) GetVmwarenetworklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmwarenetworklabel"].(string)
	return value, ok
}

func (p *AddTrafficTypeParams) SetXennetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["xennetworklabel"] = v
}

func (p *AddTrafficTypeParams) ResetXennetworklabel() {
	if p.p != nil && p.p["xennetworklabel"] != nil {
		delete(p.p, "xennetworklabel")
	}
}

func (p *AddTrafficTypeParams) GetXennetworklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["xennetworklabel"].(string)
	return value, ok
}

// You should always use this function to get a new AddTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewAddTrafficTypeParams(physicalnetworkid string, traffictype string) *AddTrafficTypeParams {
	p := &AddTrafficTypeParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["traffictype"] = traffictype
	return p
}

// Adds traffic type to a physical network
func (s *UsageService) AddTrafficType(p *AddTrafficTypeParams) (*AddTrafficTypeResponse, error) {
	resp, err := s.cs.newPostRequest("addTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddTrafficTypeResponse
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

type AddTrafficTypeResponse struct {
	Hypervnetworklabel string `json:"hypervnetworklabel"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Kvmnetworklabel    string `json:"kvmnetworklabel"`
	Ovm3networklabel   string `json:"ovm3networklabel"`
	Physicalnetworkid  string `json:"physicalnetworkid"`
	Traffictype        string `json:"traffictype"`
	Vmwarenetworklabel string `json:"vmwarenetworklabel"`
	Xennetworklabel    string `json:"xennetworklabel"`
}

type DeleteTrafficMonitorParams struct {
	p map[string]interface{}
}

func (p *DeleteTrafficMonitorParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteTrafficMonitorParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteTrafficMonitorParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteTrafficMonitorParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTrafficMonitorParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewDeleteTrafficMonitorParams(id string) *DeleteTrafficMonitorParams {
	p := &DeleteTrafficMonitorParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an traffic monitor host.
func (s *UsageService) DeleteTrafficMonitor(p *DeleteTrafficMonitorParams) (*DeleteTrafficMonitorResponse, error) {
	resp, err := s.cs.newPostRequest("deleteTrafficMonitor", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTrafficMonitorResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteTrafficMonitorResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteTrafficMonitorResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteTrafficMonitorResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteTrafficTypeParams struct {
	p map[string]interface{}
}

func (p *DeleteTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteTrafficTypeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteTrafficTypeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteTrafficTypeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewDeleteTrafficTypeParams(id string) *DeleteTrafficTypeParams {
	p := &DeleteTrafficTypeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes traffic type of a physical network
func (s *UsageService) DeleteTrafficType(p *DeleteTrafficTypeParams) (*DeleteTrafficTypeResponse, error) {
	resp, err := s.cs.newPostRequest("deleteTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTrafficTypeResponse
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

type DeleteTrafficTypeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type GenerateUsageRecordsParams struct {
	p map[string]interface{}
}

func (p *GenerateUsageRecordsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *GenerateUsageRecordsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *GenerateUsageRecordsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *GenerateUsageRecordsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *GenerateUsageRecordsParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *GenerateUsageRecordsParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *GenerateUsageRecordsParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *GenerateUsageRecordsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *GenerateUsageRecordsParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *GenerateUsageRecordsParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new GenerateUsageRecordsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewGenerateUsageRecordsParams() *GenerateUsageRecordsParams {
	p := &GenerateUsageRecordsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Generates usage records. This will generate records only if there any records to be generated, i.e if the scheduled usage job was not run or failed
func (s *UsageService) GenerateUsageRecords(p *GenerateUsageRecordsParams) (*GenerateUsageRecordsResponse, error) {
	resp, err := s.cs.newPostRequest("generateUsageRecords", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GenerateUsageRecordsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GenerateUsageRecordsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *GenerateUsageRecordsResponse) UnmarshalJSON(b []byte) error {
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

	type alias GenerateUsageRecordsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListTrafficMonitorsParams struct {
	p map[string]interface{}
}

func (p *ListTrafficMonitorsParams) toURLValues() url.Values {
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

func (p *ListTrafficMonitorsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTrafficMonitorsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTrafficMonitorsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTrafficMonitorsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTrafficMonitorsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTrafficMonitorsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTrafficMonitorsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTrafficMonitorsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTrafficMonitorsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTrafficMonitorsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTrafficMonitorsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTrafficMonitorsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTrafficMonitorsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficMonitorsParams(zoneid string) *ListTrafficMonitorsParams {
	p := &ListTrafficMonitorsParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// List traffic monitor Hosts.
func (s *UsageService) ListTrafficMonitors(p *ListTrafficMonitorsParams) (*ListTrafficMonitorsResponse, error) {
	resp, err := s.cs.newRequest("listTrafficMonitors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficMonitorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTrafficMonitorsResponse struct {
	Count           int               `json:"count"`
	TrafficMonitors []*TrafficMonitor `json:"trafficmonitor"`
}

type TrafficMonitor struct {
	Id         string `json:"id"`
	Ipaddress  string `json:"ipaddress"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Numretries string `json:"numretries"`
	Timeout    string `json:"timeout"`
	Zoneid     string `json:"zoneid"`
}

type ListTrafficTypeImplementorsParams struct {
	p map[string]interface{}
}

func (p *ListTrafficTypeImplementorsParams) toURLValues() url.Values {
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
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *ListTrafficTypeImplementorsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTrafficTypeImplementorsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTrafficTypeImplementorsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTrafficTypeImplementorsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTrafficTypeImplementorsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTrafficTypeImplementorsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTrafficTypeImplementorsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTrafficTypeImplementorsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTrafficTypeImplementorsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTrafficTypeImplementorsParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *ListTrafficTypeImplementorsParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
}

func (p *ListTrafficTypeImplementorsParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new ListTrafficTypeImplementorsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficTypeImplementorsParams() *ListTrafficTypeImplementorsParams {
	p := &ListTrafficTypeImplementorsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists implementors of implementor of a network traffic type or implementors of all network traffic types
func (s *UsageService) ListTrafficTypeImplementors(p *ListTrafficTypeImplementorsParams) (*ListTrafficTypeImplementorsResponse, error) {
	resp, err := s.cs.newRequest("listTrafficTypeImplementors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficTypeImplementorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTrafficTypeImplementorsResponse struct {
	Count                   int                       `json:"count"`
	TrafficTypeImplementors []*TrafficTypeImplementor `json:"traffictypeimplementor"`
}

type TrafficTypeImplementor struct {
	JobID                  string `json:"jobid"`
	Jobstatus              int    `json:"jobstatus"`
	Traffictype            string `json:"traffictype"`
	Traffictypeimplementor string `json:"traffictypeimplementor"`
}

type ListTrafficTypesParams struct {
	p map[string]interface{}
}

func (p *ListTrafficTypesParams) toURLValues() url.Values {
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
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListTrafficTypesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTrafficTypesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTrafficTypesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTrafficTypesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTrafficTypesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTrafficTypesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTrafficTypesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTrafficTypesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTrafficTypesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTrafficTypesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListTrafficTypesParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *ListTrafficTypesParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTrafficTypesParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficTypesParams(physicalnetworkid string) *ListTrafficTypesParams {
	p := &ListTrafficTypesParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UsageService) GetTrafficTypeID(keyword string, physicalnetworkid string, opts ...OptionFunc) (string, int, error) {
	p := &ListTrafficTypesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["physicalnetworkid"] = physicalnetworkid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListTrafficTypes(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.TrafficTypes[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.TrafficTypes {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists traffic types of a given physical network.
func (s *UsageService) ListTrafficTypes(p *ListTrafficTypesParams) (*ListTrafficTypesResponse, error) {
	resp, err := s.cs.newRequest("listTrafficTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTrafficTypesResponse struct {
	Count        int            `json:"count"`
	TrafficTypes []*TrafficType `json:"traffictype"`
}

type TrafficType struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	JobID                        string   `json:"jobid"`
	Jobstatus                    int      `json:"jobstatus"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type ListUsageRecordsParams struct {
	p map[string]interface{}
}

func (p *ListUsageRecordsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["includetags"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("includetags", vv)
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["oldformat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("oldformat", vv)
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
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["type"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("type", vv)
	}
	if v, found := p.p["usageid"]; found {
		u.Set("usageid", v.(string))
	}
	return u
}

func (p *ListUsageRecordsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListUsageRecordsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListUsageRecordsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListUsageRecordsParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *ListUsageRecordsParams) ResetAccountid() {
	if p.p != nil && p.p["accountid"] != nil {
		delete(p.p, "accountid")
	}
}

func (p *ListUsageRecordsParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *ListUsageRecordsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListUsageRecordsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListUsageRecordsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListUsageRecordsParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *ListUsageRecordsParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *ListUsageRecordsParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *ListUsageRecordsParams) SetIncludetags(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["includetags"] = v
}

func (p *ListUsageRecordsParams) ResetIncludetags() {
	if p.p != nil && p.p["includetags"] != nil {
		delete(p.p, "includetags")
	}
}

func (p *ListUsageRecordsParams) GetIncludetags() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["includetags"].(bool)
	return value, ok
}

func (p *ListUsageRecordsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListUsageRecordsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListUsageRecordsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListUsageRecordsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListUsageRecordsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListUsageRecordsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListUsageRecordsParams) SetOldformat(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["oldformat"] = v
}

func (p *ListUsageRecordsParams) ResetOldformat() {
	if p.p != nil && p.p["oldformat"] != nil {
		delete(p.p, "oldformat")
	}
}

func (p *ListUsageRecordsParams) GetOldformat() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["oldformat"].(bool)
	return value, ok
}

func (p *ListUsageRecordsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListUsageRecordsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListUsageRecordsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListUsageRecordsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListUsageRecordsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListUsageRecordsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListUsageRecordsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListUsageRecordsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListUsageRecordsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListUsageRecordsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *ListUsageRecordsParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *ListUsageRecordsParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *ListUsageRecordsParams) SetType(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListUsageRecordsParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ListUsageRecordsParams) GetType() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(int64)
	return value, ok
}

func (p *ListUsageRecordsParams) SetUsageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usageid"] = v
}

func (p *ListUsageRecordsParams) ResetUsageid() {
	if p.p != nil && p.p["usageid"] != nil {
		delete(p.p, "usageid")
	}
}

func (p *ListUsageRecordsParams) GetUsageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["usageid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUsageRecordsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListUsageRecordsParams(enddate string, startdate string) *ListUsageRecordsParams {
	p := &ListUsageRecordsParams{}
	p.p = make(map[string]interface{})
	p.p["enddate"] = enddate
	p.p["startdate"] = startdate
	return p
}

// Lists usage records for accounts
func (s *UsageService) ListUsageRecords(p *ListUsageRecordsParams) (*ListUsageRecordsResponse, error) {
	resp, err := s.cs.newRequest("listUsageRecords", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsageRecordsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUsageRecordsResponse struct {
	Count        int            `json:"count"`
	UsageRecords []*UsageRecord `json:"usagerecord"`
}

type UsageRecord struct {
	Account          string `json:"account"`
	Accountid        string `json:"accountid"`
	Cpunumber        int64  `json:"cpunumber"`
	Cpuspeed         int64  `json:"cpuspeed"`
	Description      string `json:"description"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Domainpath       string `json:"domainpath"`
	Enddate          string `json:"enddate"`
	Hasannotations   bool   `json:"hasannotations"`
	Isdefault        bool   `json:"isdefault"`
	Issourcenat      bool   `json:"issourcenat"`
	Issystem         bool   `json:"issystem"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Memory           int64  `json:"memory"`
	Name             string `json:"name"`
	Networkid        string `json:"networkid"`
	Offeringid       string `json:"offeringid"`
	Oscategoryid     string `json:"oscategoryid"`
	Oscategoryname   string `json:"oscategoryname"`
	Osdisplayname    string `json:"osdisplayname"`
	Ostypeid         string `json:"ostypeid"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Rawusage         string `json:"rawusage"`
	Size             int64  `json:"size"`
	Startdate        string `json:"startdate"`
	Tags             []Tags `json:"tags"`
	Templateid       string `json:"templateid"`
	Type             string `json:"type"`
	Usage            string `json:"usage"`
	Usageid          string `json:"usageid"`
	Usagetype        int    `json:"usagetype"`
	Virtualmachineid string `json:"virtualmachineid"`
	Virtualsize      int64  `json:"virtualsize"`
	Vpcid            string `json:"vpcid"`
	Zoneid           string `json:"zoneid"`
}

func (r *UsageRecord) UnmarshalJSON(b []byte) error {
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

	type alias UsageRecord
	return json.Unmarshal(b, (*alias)(r))
}

type ListUsageTypesParams struct {
	p map[string]interface{}
}

func (p *ListUsageTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListUsageTypesParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListUsageTypesParams() *ListUsageTypesParams {
	p := &ListUsageTypesParams{}
	p.p = make(map[string]interface{})
	return p
}

// List Usage Types
func (s *UsageService) ListUsageTypes(p *ListUsageTypesParams) (*ListUsageTypesResponse, error) {
	resp, err := s.cs.newRequest("listUsageTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsageTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUsageTypesResponse struct {
	Count      int          `json:"count"`
	UsageTypes []*UsageType `json:"usagetype"`
}

type UsageType struct {
	Description string `json:"description"`
	Id          int    `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
}

type RemoveRawUsageRecordsParams struct {
	p map[string]interface{}
}

func (p *RemoveRawUsageRecordsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	return u
}

func (p *RemoveRawUsageRecordsParams) SetInterval(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["interval"] = v
}

func (p *RemoveRawUsageRecordsParams) ResetInterval() {
	if p.p != nil && p.p["interval"] != nil {
		delete(p.p, "interval")
	}
}

func (p *RemoveRawUsageRecordsParams) GetInterval() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["interval"].(int)
	return value, ok
}

// You should always use this function to get a new RemoveRawUsageRecordsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewRemoveRawUsageRecordsParams(interval int) *RemoveRawUsageRecordsParams {
	p := &RemoveRawUsageRecordsParams{}
	p.p = make(map[string]interface{})
	p.p["interval"] = interval
	return p
}

// Safely removes raw records from cloud_usage table
func (s *UsageService) RemoveRawUsageRecords(p *RemoveRawUsageRecordsParams) (*RemoveRawUsageRecordsResponse, error) {
	resp, err := s.cs.newPostRequest("removeRawUsageRecords", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveRawUsageRecordsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveRawUsageRecordsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RemoveRawUsageRecordsResponse) UnmarshalJSON(b []byte) error {
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

	type alias RemoveRawUsageRecordsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateTrafficTypeParams struct {
	p map[string]interface{}
}

func (p *UpdateTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hypervnetworklabel"]; found {
		u.Set("hypervnetworklabel", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["kvmnetworklabel"]; found {
		u.Set("kvmnetworklabel", v.(string))
	}
	if v, found := p.p["ovm3networklabel"]; found {
		u.Set("ovm3networklabel", v.(string))
	}
	if v, found := p.p["vmwarenetworklabel"]; found {
		u.Set("vmwarenetworklabel", v.(string))
	}
	if v, found := p.p["xennetworklabel"]; found {
		u.Set("xennetworklabel", v.(string))
	}
	return u
}

func (p *UpdateTrafficTypeParams) SetHypervnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervnetworklabel"] = v
}

func (p *UpdateTrafficTypeParams) ResetHypervnetworklabel() {
	if p.p != nil && p.p["hypervnetworklabel"] != nil {
		delete(p.p, "hypervnetworklabel")
	}
}

func (p *UpdateTrafficTypeParams) GetHypervnetworklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervnetworklabel"].(string)
	return value, ok
}

func (p *UpdateTrafficTypeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateTrafficTypeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateTrafficTypeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateTrafficTypeParams) SetKvmnetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["kvmnetworklabel"] = v
}

func (p *UpdateTrafficTypeParams) ResetKvmnetworklabel() {
	if p.p != nil && p.p["kvmnetworklabel"] != nil {
		delete(p.p, "kvmnetworklabel")
	}
}

func (p *UpdateTrafficTypeParams) GetKvmnetworklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["kvmnetworklabel"].(string)
	return value, ok
}

func (p *UpdateTrafficTypeParams) SetOvm3networklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ovm3networklabel"] = v
}

func (p *UpdateTrafficTypeParams) ResetOvm3networklabel() {
	if p.p != nil && p.p["ovm3networklabel"] != nil {
		delete(p.p, "ovm3networklabel")
	}
}

func (p *UpdateTrafficTypeParams) GetOvm3networklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ovm3networklabel"].(string)
	return value, ok
}

func (p *UpdateTrafficTypeParams) SetVmwarenetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmwarenetworklabel"] = v
}

func (p *UpdateTrafficTypeParams) ResetVmwarenetworklabel() {
	if p.p != nil && p.p["vmwarenetworklabel"] != nil {
		delete(p.p, "vmwarenetworklabel")
	}
}

func (p *UpdateTrafficTypeParams) GetVmwarenetworklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmwarenetworklabel"].(string)
	return value, ok
}

func (p *UpdateTrafficTypeParams) SetXennetworklabel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["xennetworklabel"] = v
}

func (p *UpdateTrafficTypeParams) ResetXennetworklabel() {
	if p.p != nil && p.p["xennetworklabel"] != nil {
		delete(p.p, "xennetworklabel")
	}
}

func (p *UpdateTrafficTypeParams) GetXennetworklabel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["xennetworklabel"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewUpdateTrafficTypeParams(id string) *UpdateTrafficTypeParams {
	p := &UpdateTrafficTypeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates traffic type of a physical network
func (s *UsageService) UpdateTrafficType(p *UpdateTrafficTypeParams) (*UpdateTrafficTypeResponse, error) {
	resp, err := s.cs.newPostRequest("updateTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTrafficTypeResponse
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

type UpdateTrafficTypeResponse struct {
	Hypervnetworklabel string `json:"hypervnetworklabel"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Kvmnetworklabel    string `json:"kvmnetworklabel"`
	Ovm3networklabel   string `json:"ovm3networklabel"`
	Physicalnetworkid  string `json:"physicalnetworkid"`
	Traffictype        string `json:"traffictype"`
	Vmwarenetworklabel string `json:"vmwarenetworklabel"`
	Xennetworklabel    string `json:"xennetworklabel"`
}

type ListUsageServerMetricsParams struct {
	p map[string]interface{}
}

func (p *ListUsageServerMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListUsageServerMetricsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListUsageServerMetricsParams() *ListUsageServerMetricsParams {
	p := &ListUsageServerMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Usage Server metrics
func (s *UsageService) ListUsageServerMetrics(p *ListUsageServerMetricsParams) (*ListUsageServerMetricsResponse, error) {
	resp, err := s.cs.newRequest("listUsageServerMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsageServerMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUsageServerMetricsResponse struct {
	Count              int                  `json:"count"`
	UsageServerMetrics []*UsageServerMetric `json:"usageserver"`
}

type UsageServerMetric struct {
	Collectiontime    string `json:"collectiontime"`
	Hostname          string `json:"hostname"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Lastheartbeat     string `json:"lastheartbeat"`
	Lastsuccessfuljob string `json:"lastsuccessfuljob"`
	State             string `json:"state"`
}
