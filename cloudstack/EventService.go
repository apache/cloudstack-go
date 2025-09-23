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

type EventServiceIface interface {
	ArchiveEvents(p *ArchiveEventsParams) (*ArchiveEventsResponse, error)
	NewArchiveEventsParams() *ArchiveEventsParams
	DeleteEvents(p *DeleteEventsParams) (*DeleteEventsResponse, error)
	NewDeleteEventsParams() *DeleteEventsParams
	ListEventTypes(p *ListEventTypesParams) (*ListEventTypesResponse, error)
	NewListEventTypesParams() *ListEventTypesParams
	ListEvents(p *ListEventsParams) (*ListEventsResponse, error)
	NewListEventsParams() *ListEventsParams
	GetEventByID(id string, opts ...OptionFunc) (*Event, int, error)
}

type ArchiveEventsParams struct {
	p map[string]interface{}
}

func (p *ArchiveEventsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *ArchiveEventsParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *ArchiveEventsParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *ArchiveEventsParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *ArchiveEventsParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ArchiveEventsParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ArchiveEventsParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ArchiveEventsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *ArchiveEventsParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *ArchiveEventsParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *ArchiveEventsParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ArchiveEventsParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ArchiveEventsParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new ArchiveEventsParams instance,
// as then you are sure you have configured all required params
func (s *EventService) NewArchiveEventsParams() *ArchiveEventsParams {
	p := &ArchiveEventsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Archive one or more events.
func (s *EventService) ArchiveEvents(p *ArchiveEventsParams) (*ArchiveEventsResponse, error) {
	resp, err := s.cs.newPostRequest("archiveEvents", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ArchiveEventsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ArchiveEventsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ArchiveEventsResponse) UnmarshalJSON(b []byte) error {
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

	type alias ArchiveEventsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteEventsParams struct {
	p map[string]interface{}
}

func (p *DeleteEventsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *DeleteEventsParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *DeleteEventsParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *DeleteEventsParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *DeleteEventsParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *DeleteEventsParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *DeleteEventsParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *DeleteEventsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *DeleteEventsParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *DeleteEventsParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *DeleteEventsParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *DeleteEventsParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *DeleteEventsParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteEventsParams instance,
// as then you are sure you have configured all required params
func (s *EventService) NewDeleteEventsParams() *DeleteEventsParams {
	p := &DeleteEventsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Delete one or more events.
func (s *EventService) DeleteEvents(p *DeleteEventsParams) (*DeleteEventsResponse, error) {
	resp, err := s.cs.newPostRequest("deleteEvents", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteEventsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteEventsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteEventsResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteEventsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListEventTypesParams struct {
	p map[string]interface{}
}

func (p *ListEventTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListEventTypesParams instance,
// as then you are sure you have configured all required params
func (s *EventService) NewListEventTypesParams() *ListEventTypesParams {
	p := &ListEventTypesParams{}
	p.p = make(map[string]interface{})
	return p
}

// List Event Types
func (s *EventService) ListEventTypes(p *ListEventTypesParams) (*ListEventTypesResponse, error) {
	resp, err := s.cs.newRequest("listEventTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListEventTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListEventTypesResponse struct {
	Count      int          `json:"count"`
	EventTypes []*EventType `json:"eventtype"`
}

type EventType struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListEventsParams struct {
	p map[string]interface{}
}

func (p *ListEventsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["archived"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("archived", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["entrytime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("entrytime", vv)
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
	if v, found := p.p["level"]; found {
		u.Set("level", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["startid"]; found {
		u.Set("startid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *ListEventsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListEventsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListEventsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListEventsParams) SetArchived(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["archived"] = v
}

func (p *ListEventsParams) ResetArchived() {
	if p.p != nil && p.p["archived"] != nil {
		delete(p.p, "archived")
	}
}

func (p *ListEventsParams) GetArchived() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["archived"].(bool)
	return value, ok
}

func (p *ListEventsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListEventsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListEventsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListEventsParams) SetDuration(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["duration"] = v
}

func (p *ListEventsParams) ResetDuration() {
	if p.p != nil && p.p["duration"] != nil {
		delete(p.p, "duration")
	}
}

func (p *ListEventsParams) GetDuration() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["duration"].(int)
	return value, ok
}

func (p *ListEventsParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *ListEventsParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *ListEventsParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *ListEventsParams) SetEntrytime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["entrytime"] = v
}

func (p *ListEventsParams) ResetEntrytime() {
	if p.p != nil && p.p["entrytime"] != nil {
		delete(p.p, "entrytime")
	}
}

func (p *ListEventsParams) GetEntrytime() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["entrytime"].(int)
	return value, ok
}

func (p *ListEventsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListEventsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListEventsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListEventsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListEventsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListEventsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListEventsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListEventsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListEventsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListEventsParams) SetLevel(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["level"] = v
}

func (p *ListEventsParams) ResetLevel() {
	if p.p != nil && p.p["level"] != nil {
		delete(p.p, "level")
	}
}

func (p *ListEventsParams) GetLevel() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["level"].(string)
	return value, ok
}

func (p *ListEventsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListEventsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListEventsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListEventsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListEventsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListEventsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListEventsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListEventsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListEventsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListEventsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListEventsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListEventsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListEventsParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *ListEventsParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *ListEventsParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *ListEventsParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *ListEventsParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *ListEventsParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

func (p *ListEventsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *ListEventsParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *ListEventsParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *ListEventsParams) SetStartid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startid"] = v
}

func (p *ListEventsParams) ResetStartid() {
	if p.p != nil && p.p["startid"] != nil {
		delete(p.p, "startid")
	}
}

func (p *ListEventsParams) GetStartid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startid"].(string)
	return value, ok
}

func (p *ListEventsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListEventsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListEventsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListEventsParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListEventsParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ListEventsParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new ListEventsParams instance,
// as then you are sure you have configured all required params
func (s *EventService) NewListEventsParams() *ListEventsParams {
	p := &ListEventsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *EventService) GetEventByID(id string, opts ...OptionFunc) (*Event, int, error) {
	p := &ListEventsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListEvents(p)
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
		return l.Events[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Event UUID: %s!", id)
}

// A command to list events.
func (s *EventService) ListEvents(p *ListEventsParams) (*ListEventsResponse, error) {
	resp, err := s.cs.newRequest("listEvents", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListEventsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListEventsResponse struct {
	Count  int      `json:"count"`
	Events []*Event `json:"event"`
}

type Event struct {
	Account      string `json:"account"`
	Archived     bool   `json:"archived"`
	Created      string `json:"created"`
	Description  string `json:"description"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Domainpath   string `json:"domainpath"`
	Id           string `json:"id"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Level        string `json:"level"`
	Parentid     string `json:"parentid"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Resourceid   string `json:"resourceid"`
	Resourcename string `json:"resourcename"`
	Resourcetype string `json:"resourcetype"`
	State        string `json:"state"`
	Type         string `json:"type"`
	Username     string `json:"username"`
}
