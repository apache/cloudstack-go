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

type BGPPeerServiceIface interface {
	ChangeBgpPeersForVpc(p *ChangeBgpPeersForVpcParams) (*ChangeBgpPeersForVpcResponse, error)
	NewChangeBgpPeersForVpcParams(vpcid string) *ChangeBgpPeersForVpcParams
	CreateBgpPeer(p *CreateBgpPeerParams) (*CreateBgpPeerResponse, error)
	NewCreateBgpPeerParams(asnumber int64, zoneid string) *CreateBgpPeerParams
	DedicateBgpPeer(p *DedicateBgpPeerParams) (*DedicateBgpPeerResponse, error)
	NewDedicateBgpPeerParams(id string) *DedicateBgpPeerParams
	DeleteBgpPeer(p *DeleteBgpPeerParams) (*DeleteBgpPeerResponse, error)
	NewDeleteBgpPeerParams(id string) *DeleteBgpPeerParams
	ListBgpPeers(p *ListBgpPeersParams) (*ListBgpPeersResponse, error)
	NewListBgpPeersParams() *ListBgpPeersParams
	GetBgpPeerByID(id string, opts ...OptionFunc) (*BgpPeer, int, error)
	ReleaseBgpPeer(p *ReleaseBgpPeerParams) (*ReleaseBgpPeerResponse, error)
	NewReleaseBgpPeerParams(id string) *ReleaseBgpPeerParams
	UpdateBgpPeer(p *UpdateBgpPeerParams) (*UpdateBgpPeerResponse, error)
	NewUpdateBgpPeerParams(id string) *UpdateBgpPeerParams
}

type ChangeBgpPeersForVpcParams struct {
	p map[string]interface{}
}

func (p *ChangeBgpPeersForVpcParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bgppeerids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("bgppeerids", vv)
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ChangeBgpPeersForVpcParams) SetBgppeerids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bgppeerids"] = v
}

func (p *ChangeBgpPeersForVpcParams) ResetBgppeerids() {
	if p.p != nil && p.p["bgppeerids"] != nil {
		delete(p.p, "bgppeerids")
	}
}

func (p *ChangeBgpPeersForVpcParams) GetBgppeerids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bgppeerids"].([]string)
	return value, ok
}

func (p *ChangeBgpPeersForVpcParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ChangeBgpPeersForVpcParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ChangeBgpPeersForVpcParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeBgpPeersForVpcParams instance,
// as then you are sure you have configured all required params
func (s *BGPPeerService) NewChangeBgpPeersForVpcParams(vpcid string) *ChangeBgpPeersForVpcParams {
	p := &ChangeBgpPeersForVpcParams{}
	p.p = make(map[string]interface{})
	p.p["vpcid"] = vpcid
	return p
}

// Change the BGP peers for a VPC.
func (s *BGPPeerService) ChangeBgpPeersForVpc(p *ChangeBgpPeersForVpcParams) (*ChangeBgpPeersForVpcResponse, error) {
	resp, err := s.cs.newRequest("changeBgpPeersForVpc", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeBgpPeersForVpcResponse
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

type ChangeBgpPeersForVpcResponse struct {
	Account    string            `json:"account"`
	Asnumber   int64             `json:"asnumber"`
	Created    string            `json:"created"`
	Details    map[string]string `json:"details"`
	Domain     string            `json:"domain"`
	Domainid   string            `json:"domainid"`
	Id         string            `json:"id"`
	Ip6address string            `json:"ip6address"`
	Ipaddress  string            `json:"ipaddress"`
	JobID      string            `json:"jobid"`
	Jobstatus  int               `json:"jobstatus"`
	Password   string            `json:"password"`
	Project    string            `json:"project"`
	Projectid  string            `json:"projectid"`
	Zoneid     string            `json:"zoneid"`
	Zonename   string            `json:"zonename"`
}

type CreateBgpPeerParams struct {
	p map[string]interface{}
}

func (p *CreateBgpPeerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["asnumber"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("asnumber", vv)
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["ip6address"]; found {
		u.Set("ip6address", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateBgpPeerParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateBgpPeerParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateBgpPeerParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateBgpPeerParams) SetAsnumber(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asnumber"] = v
}

func (p *CreateBgpPeerParams) ResetAsnumber() {
	if p.p != nil && p.p["asnumber"] != nil {
		delete(p.p, "asnumber")
	}
}

func (p *CreateBgpPeerParams) GetAsnumber() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asnumber"].(int64)
	return value, ok
}

func (p *CreateBgpPeerParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateBgpPeerParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *CreateBgpPeerParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *CreateBgpPeerParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateBgpPeerParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateBgpPeerParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateBgpPeerParams) SetIp6address(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6address"] = v
}

func (p *CreateBgpPeerParams) ResetIp6address() {
	if p.p != nil && p.p["ip6address"] != nil {
		delete(p.p, "ip6address")
	}
}

func (p *CreateBgpPeerParams) GetIp6address() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6address"].(string)
	return value, ok
}

func (p *CreateBgpPeerParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *CreateBgpPeerParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *CreateBgpPeerParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *CreateBgpPeerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *CreateBgpPeerParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *CreateBgpPeerParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *CreateBgpPeerParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateBgpPeerParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateBgpPeerParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateBgpPeerParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateBgpPeerParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateBgpPeerParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateBgpPeerParams instance,
// as then you are sure you have configured all required params
func (s *BGPPeerService) NewCreateBgpPeerParams(asnumber int64, zoneid string) *CreateBgpPeerParams {
	p := &CreateBgpPeerParams{}
	p.p = make(map[string]interface{})
	p.p["asnumber"] = asnumber
	p.p["zoneid"] = zoneid
	return p
}

// Creates a Bgp Peer for a zone.
func (s *BGPPeerService) CreateBgpPeer(p *CreateBgpPeerParams) (*CreateBgpPeerResponse, error) {
	resp, err := s.cs.newRequest("createBgpPeer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateBgpPeerResponse
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

type CreateBgpPeerResponse struct {
	Account    string            `json:"account"`
	Asnumber   int64             `json:"asnumber"`
	Created    string            `json:"created"`
	Details    map[string]string `json:"details"`
	Domain     string            `json:"domain"`
	Domainid   string            `json:"domainid"`
	Id         string            `json:"id"`
	Ip6address string            `json:"ip6address"`
	Ipaddress  string            `json:"ipaddress"`
	JobID      string            `json:"jobid"`
	Jobstatus  int               `json:"jobstatus"`
	Password   string            `json:"password"`
	Project    string            `json:"project"`
	Projectid  string            `json:"projectid"`
	Zoneid     string            `json:"zoneid"`
	Zonename   string            `json:"zonename"`
}

type DedicateBgpPeerParams struct {
	p map[string]interface{}
}

func (p *DedicateBgpPeerParams) toURLValues() url.Values {
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
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DedicateBgpPeerParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DedicateBgpPeerParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DedicateBgpPeerParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DedicateBgpPeerParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DedicateBgpPeerParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DedicateBgpPeerParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DedicateBgpPeerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DedicateBgpPeerParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DedicateBgpPeerParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DedicateBgpPeerParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DedicateBgpPeerParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *DedicateBgpPeerParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicateBgpPeerParams instance,
// as then you are sure you have configured all required params
func (s *BGPPeerService) NewDedicateBgpPeerParams(id string) *DedicateBgpPeerParams {
	p := &DedicateBgpPeerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Dedicates an existing Bgp Peer to an account or a domain.
func (s *BGPPeerService) DedicateBgpPeer(p *DedicateBgpPeerParams) (*DedicateBgpPeerResponse, error) {
	resp, err := s.cs.newRequest("dedicateBgpPeer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateBgpPeerResponse
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

type DedicateBgpPeerResponse struct {
	Account    string            `json:"account"`
	Asnumber   int64             `json:"asnumber"`
	Created    string            `json:"created"`
	Details    map[string]string `json:"details"`
	Domain     string            `json:"domain"`
	Domainid   string            `json:"domainid"`
	Id         string            `json:"id"`
	Ip6address string            `json:"ip6address"`
	Ipaddress  string            `json:"ipaddress"`
	JobID      string            `json:"jobid"`
	Jobstatus  int               `json:"jobstatus"`
	Password   string            `json:"password"`
	Project    string            `json:"project"`
	Projectid  string            `json:"projectid"`
	Zoneid     string            `json:"zoneid"`
	Zonename   string            `json:"zonename"`
}

type DeleteBgpPeerParams struct {
	p map[string]interface{}
}

func (p *DeleteBgpPeerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteBgpPeerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteBgpPeerParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteBgpPeerParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBgpPeerParams instance,
// as then you are sure you have configured all required params
func (s *BGPPeerService) NewDeleteBgpPeerParams(id string) *DeleteBgpPeerParams {
	p := &DeleteBgpPeerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an existing Bgp Peer.
func (s *BGPPeerService) DeleteBgpPeer(p *DeleteBgpPeerParams) (*DeleteBgpPeerResponse, error) {
	resp, err := s.cs.newRequest("deleteBgpPeer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBgpPeerResponse
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

type DeleteBgpPeerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListBgpPeersParams struct {
	p map[string]interface{}
}

func (p *ListBgpPeersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["asnumber"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("asnumber", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdedicated"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdedicated", vv)
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListBgpPeersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListBgpPeersParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListBgpPeersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListBgpPeersParams) SetAsnumber(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asnumber"] = v
}

func (p *ListBgpPeersParams) ResetAsnumber() {
	if p.p != nil && p.p["asnumber"] != nil {
		delete(p.p, "asnumber")
	}
}

func (p *ListBgpPeersParams) GetAsnumber() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asnumber"].(int64)
	return value, ok
}

func (p *ListBgpPeersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListBgpPeersParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListBgpPeersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListBgpPeersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListBgpPeersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListBgpPeersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListBgpPeersParams) SetIsdedicated(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdedicated"] = v
}

func (p *ListBgpPeersParams) ResetIsdedicated() {
	if p.p != nil && p.p["isdedicated"] != nil {
		delete(p.p, "isdedicated")
	}
}

func (p *ListBgpPeersParams) GetIsdedicated() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdedicated"].(bool)
	return value, ok
}

func (p *ListBgpPeersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBgpPeersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBgpPeersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBgpPeersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBgpPeersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBgpPeersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBgpPeersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBgpPeersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBgpPeersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBgpPeersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListBgpPeersParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListBgpPeersParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListBgpPeersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListBgpPeersParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListBgpPeersParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBgpPeersParams instance,
// as then you are sure you have configured all required params
func (s *BGPPeerService) NewListBgpPeersParams() *ListBgpPeersParams {
	p := &ListBgpPeersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BGPPeerService) GetBgpPeerByID(id string, opts ...OptionFunc) (*BgpPeer, int, error) {
	p := &ListBgpPeersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListBgpPeers(p)
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
		return l.BgpPeers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for BgpPeer UUID: %s!", id)
}

// Lists Bgp Peers.
func (s *BGPPeerService) ListBgpPeers(p *ListBgpPeersParams) (*ListBgpPeersResponse, error) {
	resp, err := s.cs.newRequest("listBgpPeers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBgpPeersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBgpPeersResponse struct {
	Count    int        `json:"count"`
	BgpPeers []*BgpPeer `json:"bgppeer"`
}

type BgpPeer struct {
	Account    string            `json:"account"`
	Asnumber   int64             `json:"asnumber"`
	Created    string            `json:"created"`
	Details    map[string]string `json:"details"`
	Domain     string            `json:"domain"`
	Domainid   string            `json:"domainid"`
	Id         string            `json:"id"`
	Ip6address string            `json:"ip6address"`
	Ipaddress  string            `json:"ipaddress"`
	JobID      string            `json:"jobid"`
	Jobstatus  int               `json:"jobstatus"`
	Password   string            `json:"password"`
	Project    string            `json:"project"`
	Projectid  string            `json:"projectid"`
	Zoneid     string            `json:"zoneid"`
	Zonename   string            `json:"zonename"`
}

type ReleaseBgpPeerParams struct {
	p map[string]interface{}
}

func (p *ReleaseBgpPeerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ReleaseBgpPeerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ReleaseBgpPeerParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ReleaseBgpPeerParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseBgpPeerParams instance,
// as then you are sure you have configured all required params
func (s *BGPPeerService) NewReleaseBgpPeerParams(id string) *ReleaseBgpPeerParams {
	p := &ReleaseBgpPeerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Releases an existing dedicated Bgp Peer.
func (s *BGPPeerService) ReleaseBgpPeer(p *ReleaseBgpPeerParams) (*ReleaseBgpPeerResponse, error) {
	resp, err := s.cs.newRequest("releaseBgpPeer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseBgpPeerResponse
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

type ReleaseBgpPeerResponse struct {
	Account    string            `json:"account"`
	Asnumber   int64             `json:"asnumber"`
	Created    string            `json:"created"`
	Details    map[string]string `json:"details"`
	Domain     string            `json:"domain"`
	Domainid   string            `json:"domainid"`
	Id         string            `json:"id"`
	Ip6address string            `json:"ip6address"`
	Ipaddress  string            `json:"ipaddress"`
	JobID      string            `json:"jobid"`
	Jobstatus  int               `json:"jobstatus"`
	Password   string            `json:"password"`
	Project    string            `json:"project"`
	Projectid  string            `json:"projectid"`
	Zoneid     string            `json:"zoneid"`
	Zonename   string            `json:"zonename"`
}

type UpdateBgpPeerParams struct {
	p map[string]interface{}
}

func (p *UpdateBgpPeerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["asnumber"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("asnumber", vv)
	}
	if v, found := p.p["cleanupdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupdetails", vv)
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ip6address"]; found {
		u.Set("ip6address", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	return u
}

func (p *UpdateBgpPeerParams) SetAsnumber(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asnumber"] = v
}

func (p *UpdateBgpPeerParams) ResetAsnumber() {
	if p.p != nil && p.p["asnumber"] != nil {
		delete(p.p, "asnumber")
	}
}

func (p *UpdateBgpPeerParams) GetAsnumber() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asnumber"].(int64)
	return value, ok
}

func (p *UpdateBgpPeerParams) SetCleanupdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupdetails"] = v
}

func (p *UpdateBgpPeerParams) ResetCleanupdetails() {
	if p.p != nil && p.p["cleanupdetails"] != nil {
		delete(p.p, "cleanupdetails")
	}
}

func (p *UpdateBgpPeerParams) GetCleanupdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupdetails"].(bool)
	return value, ok
}

func (p *UpdateBgpPeerParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateBgpPeerParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateBgpPeerParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateBgpPeerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateBgpPeerParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateBgpPeerParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateBgpPeerParams) SetIp6address(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6address"] = v
}

func (p *UpdateBgpPeerParams) ResetIp6address() {
	if p.p != nil && p.p["ip6address"] != nil {
		delete(p.p, "ip6address")
	}
}

func (p *UpdateBgpPeerParams) GetIp6address() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6address"].(string)
	return value, ok
}

func (p *UpdateBgpPeerParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *UpdateBgpPeerParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *UpdateBgpPeerParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *UpdateBgpPeerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *UpdateBgpPeerParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *UpdateBgpPeerParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateBgpPeerParams instance,
// as then you are sure you have configured all required params
func (s *BGPPeerService) NewUpdateBgpPeerParams(id string) *UpdateBgpPeerParams {
	p := &UpdateBgpPeerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing Bgp Peer.
func (s *BGPPeerService) UpdateBgpPeer(p *UpdateBgpPeerParams) (*UpdateBgpPeerResponse, error) {
	resp, err := s.cs.newRequest("updateBgpPeer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateBgpPeerResponse
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

type UpdateBgpPeerResponse struct {
	Account    string            `json:"account"`
	Asnumber   int64             `json:"asnumber"`
	Created    string            `json:"created"`
	Details    map[string]string `json:"details"`
	Domain     string            `json:"domain"`
	Domainid   string            `json:"domainid"`
	Id         string            `json:"id"`
	Ip6address string            `json:"ip6address"`
	Ipaddress  string            `json:"ipaddress"`
	JobID      string            `json:"jobid"`
	Jobstatus  int               `json:"jobstatus"`
	Password   string            `json:"password"`
	Project    string            `json:"project"`
	Projectid  string            `json:"projectid"`
	Zoneid     string            `json:"zoneid"`
	Zonename   string            `json:"zonename"`
}
