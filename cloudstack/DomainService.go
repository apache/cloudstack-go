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

type DomainServiceIface interface {
	CreateDomain(p *CreateDomainParams) (*CreateDomainResponse, error)
	NewCreateDomainParams(name string) *CreateDomainParams
	DeleteDomain(p *DeleteDomainParams) (*DeleteDomainResponse, error)
	NewDeleteDomainParams(id string) *DeleteDomainParams
	ListDomainChildren(p *ListDomainChildrenParams) (*ListDomainChildrenResponse, error)
	NewListDomainChildrenParams() *ListDomainChildrenParams
	GetDomainChildrenID(name string, opts ...OptionFunc) (string, int, error)
	GetDomainChildrenByName(name string, opts ...OptionFunc) (*DomainChildren, int, error)
	GetDomainChildrenByID(id string, opts ...OptionFunc) (*DomainChildren, int, error)
	ListDomains(p *ListDomainsParams) (*ListDomainsResponse, error)
	NewListDomainsParams() *ListDomainsParams
	GetDomainID(name string, opts ...OptionFunc) (string, int, error)
	GetDomainByName(name string, opts ...OptionFunc) (*Domain, int, error)
	GetDomainByID(id string, opts ...OptionFunc) (*Domain, int, error)
	MoveDomain(p *MoveDomainParams) (*MoveDomainResponse, error)
	NewMoveDomainParams(domainid string, parentdomainid string) *MoveDomainParams
	UpdateDomain(p *UpdateDomainParams) (*UpdateDomainResponse, error)
	NewUpdateDomainParams(id string) *UpdateDomainParams
}

type CreateDomainParams struct {
	p map[string]interface{}
}

func (p *CreateDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["parentdomainid"]; found {
		u.Set("parentdomainid", v.(string))
	}
	return u
}

func (p *CreateDomainParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateDomainParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateDomainParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateDomainParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateDomainParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateDomainParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateDomainParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *CreateDomainParams) ResetNetworkdomain() {
	if p.p != nil && p.p["networkdomain"] != nil {
		delete(p.p, "networkdomain")
	}
}

func (p *CreateDomainParams) GetNetworkdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdomain"].(string)
	return value, ok
}

func (p *CreateDomainParams) SetParentdomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parentdomainid"] = v
}

func (p *CreateDomainParams) ResetParentdomainid() {
	if p.p != nil && p.p["parentdomainid"] != nil {
		delete(p.p, "parentdomainid")
	}
}

func (p *CreateDomainParams) GetParentdomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parentdomainid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewCreateDomainParams(name string) *CreateDomainParams {
	p := &CreateDomainParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Creates a domain
func (s *DomainService) CreateDomain(p *CreateDomainParams) (*CreateDomainResponse, error) {
	resp, err := s.cs.newPostRequest("createDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateDomainResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateDomainResponse struct {
	Backupavailable           string            `json:"backupavailable"`
	Backuplimit               string            `json:"backuplimit"`
	Backupstorageavailable    string            `json:"backupstorageavailable"`
	Backupstoragelimit        string            `json:"backupstoragelimit"`
	Backupstoragetotal        int64             `json:"backupstoragetotal"`
	Backuptotal               int64             `json:"backuptotal"`
	Bucketavailable           string            `json:"bucketavailable"`
	Bucketlimit               string            `json:"bucketlimit"`
	Buckettotal               int64             `json:"buckettotal"`
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Gpuavailable              string            `json:"gpuavailable"`
	Gpulimit                  string            `json:"gpulimit"`
	Gputotal                  int64             `json:"gputotal"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      interface{}       `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Objectstorageavailable    string            `json:"objectstorageavailable"`
	Objectstoragelimit        string            `json:"objectstoragelimit"`
	Objectstoragetotal        int64             `json:"objectstoragetotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Taggedresources           []string          `json:"taggedresources"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type DeleteDomainParams struct {
	p map[string]interface{}
}

func (p *DeleteDomainParams) toURLValues() url.Values {
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

func (p *DeleteDomainParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *DeleteDomainParams) ResetCleanup() {
	if p.p != nil && p.p["cleanup"] != nil {
		delete(p.p, "cleanup")
	}
}

func (p *DeleteDomainParams) GetCleanup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanup"].(bool)
	return value, ok
}

func (p *DeleteDomainParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteDomainParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteDomainParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewDeleteDomainParams(id string) *DeleteDomainParams {
	p := &DeleteDomainParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a specified domain
func (s *DomainService) DeleteDomain(p *DeleteDomainParams) (*DeleteDomainResponse, error) {
	resp, err := s.cs.newPostRequest("deleteDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteDomainResponse
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

type DeleteDomainResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListDomainChildrenParams struct {
	p map[string]interface{}
}

func (p *ListDomainChildrenParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	return u
}

func (p *ListDomainChildrenParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListDomainChildrenParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListDomainChildrenParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListDomainChildrenParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListDomainChildrenParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListDomainChildrenParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListDomainChildrenParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListDomainChildrenParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListDomainChildrenParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListDomainChildrenParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListDomainChildrenParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListDomainChildrenParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListDomainChildrenParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListDomainChildrenParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListDomainChildrenParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListDomainChildrenParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListDomainChildrenParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListDomainChildrenParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListDomainChildrenParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListDomainChildrenParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListDomainChildrenParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListDomainChildrenParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListDomainChildrenParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListDomainChildrenParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

// You should always use this function to get a new ListDomainChildrenParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewListDomainChildrenParams() *ListDomainChildrenParams {
	p := &ListDomainChildrenParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainChildrenID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListDomainChildrenParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListDomainChildren(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.DomainChildren[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.DomainChildren {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainChildrenByName(name string, opts ...OptionFunc) (*DomainChildren, int, error) {
	id, count, err := s.GetDomainChildrenID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetDomainChildrenByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainChildrenByID(id string, opts ...OptionFunc) (*DomainChildren, int, error) {
	p := &ListDomainChildrenParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListDomainChildren(p)
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
		return l.DomainChildren[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for DomainChildren UUID: %s!", id)
}

// Lists all children domains belonging to a specified domain
func (s *DomainService) ListDomainChildren(p *ListDomainChildrenParams) (*ListDomainChildrenResponse, error) {
	resp, err := s.cs.newRequest("listDomainChildren", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDomainChildrenResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDomainChildrenResponse struct {
	Count          int               `json:"count"`
	DomainChildren []*DomainChildren `json:"domain"`
}

type DomainChildren struct {
	Backupavailable           string            `json:"backupavailable"`
	Backuplimit               string            `json:"backuplimit"`
	Backupstorageavailable    string            `json:"backupstorageavailable"`
	Backupstoragelimit        string            `json:"backupstoragelimit"`
	Backupstoragetotal        int64             `json:"backupstoragetotal"`
	Backuptotal               int64             `json:"backuptotal"`
	Bucketavailable           string            `json:"bucketavailable"`
	Bucketlimit               string            `json:"bucketlimit"`
	Buckettotal               int64             `json:"buckettotal"`
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Gpuavailable              string            `json:"gpuavailable"`
	Gpulimit                  string            `json:"gpulimit"`
	Gputotal                  int64             `json:"gputotal"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      interface{}       `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Objectstorageavailable    string            `json:"objectstorageavailable"`
	Objectstoragelimit        string            `json:"objectstoragelimit"`
	Objectstoragetotal        int64             `json:"objectstoragetotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Taggedresources           []string          `json:"taggedresources"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type ListDomainsParams struct {
	p map[string]interface{}
}

func (p *ListDomainsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["level"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("level", vv)
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["tag"]; found {
		u.Set("tag", v.(string))
	}
	return u
}

func (p *ListDomainsParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListDomainsParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListDomainsParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListDomainsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListDomainsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListDomainsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListDomainsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListDomainsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListDomainsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListDomainsParams) SetLevel(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["level"] = v
}

func (p *ListDomainsParams) ResetLevel() {
	if p.p != nil && p.p["level"] != nil {
		delete(p.p, "level")
	}
}

func (p *ListDomainsParams) GetLevel() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["level"].(int)
	return value, ok
}

func (p *ListDomainsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListDomainsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListDomainsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListDomainsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListDomainsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListDomainsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListDomainsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListDomainsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListDomainsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListDomainsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListDomainsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListDomainsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListDomainsParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListDomainsParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListDomainsParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListDomainsParams) SetTag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tag"] = v
}

func (p *ListDomainsParams) ResetTag() {
	if p.p != nil && p.p["tag"] != nil {
		delete(p.p, "tag")
	}
}

func (p *ListDomainsParams) GetTag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tag"].(string)
	return value, ok
}

// You should always use this function to get a new ListDomainsParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewListDomainsParams() *ListDomainsParams {
	p := &ListDomainsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListDomainsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListDomains(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Domains[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Domains {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainByName(name string, opts ...OptionFunc) (*Domain, int, error) {
	id, count, err := s.GetDomainID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetDomainByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainByID(id string, opts ...OptionFunc) (*Domain, int, error) {
	p := &ListDomainsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListDomains(p)
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
		return l.Domains[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Domain UUID: %s!", id)
}

// Lists domains and provides detailed information for listed domains
func (s *DomainService) ListDomains(p *ListDomainsParams) (*ListDomainsResponse, error) {
	resp, err := s.cs.newRequest("listDomains", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDomainsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDomainsResponse struct {
	Count   int       `json:"count"`
	Domains []*Domain `json:"domain"`
}

type Domain struct {
	Backupavailable           string            `json:"backupavailable"`
	Backuplimit               string            `json:"backuplimit"`
	Backupstorageavailable    string            `json:"backupstorageavailable"`
	Backupstoragelimit        string            `json:"backupstoragelimit"`
	Backupstoragetotal        int64             `json:"backupstoragetotal"`
	Backuptotal               int64             `json:"backuptotal"`
	Bucketavailable           string            `json:"bucketavailable"`
	Bucketlimit               string            `json:"bucketlimit"`
	Buckettotal               int64             `json:"buckettotal"`
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Gpuavailable              string            `json:"gpuavailable"`
	Gpulimit                  string            `json:"gpulimit"`
	Gputotal                  int64             `json:"gputotal"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      interface{}       `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Objectstorageavailable    string            `json:"objectstorageavailable"`
	Objectstoragelimit        string            `json:"objectstoragelimit"`
	Objectstoragetotal        int64             `json:"objectstoragetotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Taggedresources           []string          `json:"taggedresources"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type MoveDomainParams struct {
	p map[string]interface{}
}

func (p *MoveDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["parentdomainid"]; found {
		u.Set("parentdomainid", v.(string))
	}
	return u
}

func (p *MoveDomainParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *MoveDomainParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *MoveDomainParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *MoveDomainParams) SetParentdomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parentdomainid"] = v
}

func (p *MoveDomainParams) ResetParentdomainid() {
	if p.p != nil && p.p["parentdomainid"] != nil {
		delete(p.p, "parentdomainid")
	}
}

func (p *MoveDomainParams) GetParentdomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parentdomainid"].(string)
	return value, ok
}

// You should always use this function to get a new MoveDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewMoveDomainParams(domainid string, parentdomainid string) *MoveDomainParams {
	p := &MoveDomainParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	p.p["parentdomainid"] = parentdomainid
	return p
}

// Moves a domain and its children to a new parent domain.
func (s *DomainService) MoveDomain(p *MoveDomainParams) (*MoveDomainResponse, error) {
	resp, err := s.cs.newPostRequest("moveDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MoveDomainResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type MoveDomainResponse struct {
	Backupavailable           string            `json:"backupavailable"`
	Backuplimit               string            `json:"backuplimit"`
	Backupstorageavailable    string            `json:"backupstorageavailable"`
	Backupstoragelimit        string            `json:"backupstoragelimit"`
	Backupstoragetotal        int64             `json:"backupstoragetotal"`
	Backuptotal               int64             `json:"backuptotal"`
	Bucketavailable           string            `json:"bucketavailable"`
	Bucketlimit               string            `json:"bucketlimit"`
	Buckettotal               int64             `json:"buckettotal"`
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Gpuavailable              string            `json:"gpuavailable"`
	Gpulimit                  string            `json:"gpulimit"`
	Gputotal                  int64             `json:"gputotal"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      interface{}       `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Objectstorageavailable    string            `json:"objectstorageavailable"`
	Objectstoragelimit        string            `json:"objectstoragelimit"`
	Objectstoragetotal        int64             `json:"objectstoragetotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Taggedresources           []string          `json:"taggedresources"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type UpdateDomainParams struct {
	p map[string]interface{}
}

func (p *UpdateDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	return u
}

func (p *UpdateDomainParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateDomainParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateDomainParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateDomainParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateDomainParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateDomainParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateDomainParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *UpdateDomainParams) ResetNetworkdomain() {
	if p.p != nil && p.p["networkdomain"] != nil {
		delete(p.p, "networkdomain")
	}
}

func (p *UpdateDomainParams) GetNetworkdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdomain"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewUpdateDomainParams(id string) *UpdateDomainParams {
	p := &UpdateDomainParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a domain with a new name
func (s *DomainService) UpdateDomain(p *UpdateDomainParams) (*UpdateDomainResponse, error) {
	resp, err := s.cs.newPostRequest("updateDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateDomainResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateDomainResponse struct {
	Backupavailable           string            `json:"backupavailable"`
	Backuplimit               string            `json:"backuplimit"`
	Backupstorageavailable    string            `json:"backupstorageavailable"`
	Backupstoragelimit        string            `json:"backupstoragelimit"`
	Backupstoragetotal        int64             `json:"backupstoragetotal"`
	Backuptotal               int64             `json:"backuptotal"`
	Bucketavailable           string            `json:"bucketavailable"`
	Bucketlimit               string            `json:"bucketlimit"`
	Buckettotal               int64             `json:"buckettotal"`
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Gpuavailable              string            `json:"gpuavailable"`
	Gpulimit                  string            `json:"gpulimit"`
	Gputotal                  int64             `json:"gputotal"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      interface{}       `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Objectstorageavailable    string            `json:"objectstorageavailable"`
	Objectstoragelimit        string            `json:"objectstoragelimit"`
	Objectstoragetotal        int64             `json:"objectstoragetotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Taggedresources           []string          `json:"taggedresources"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}
