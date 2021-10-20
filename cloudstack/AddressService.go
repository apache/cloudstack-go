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

type AddressServiceIface interface {
	AssociateIpAddress(p *AssociateIpAddressParams) (*AssociateIpAddressResponse, error)
	NewAssociateIpAddressParams() *AssociateIpAddressParams
	DisassociateIpAddress(p *DisassociateIpAddressParams) (*DisassociateIpAddressResponse, error)
	NewDisassociateIpAddressParams(id string) *DisassociateIpAddressParams
	ListPublicIpAddresses(p *ListPublicIpAddressesParams) (*ListPublicIpAddressesResponse, error)
	NewListPublicIpAddressesParams() *ListPublicIpAddressesParams
	GetPublicIpAddressByID(id string, opts ...OptionFunc) (*PublicIpAddress, int, error)
	UpdateIpAddress(p *UpdateIpAddressParams) (*UpdateIpAddressResponse, error)
	NewUpdateIpAddressParams(id string) *UpdateIpAddressParams
}

type AssociateIpAddressParams struct {
	p map[string]interface{}
}

func (p *AssociateIpAddressParams) toURLValues() url.Values {
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
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["isportable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isportable", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AssociateIpAddressParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AssociateIpAddressParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *AssociateIpAddressParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *AssociateIpAddressParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *AssociateIpAddressParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *AssociateIpAddressParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *AssociateIpAddressParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *AssociateIpAddressParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *AssociateIpAddressParams) SetIsportable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isportable"] = v
}

func (p *AssociateIpAddressParams) GetIsportable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isportable"].(bool)
	return value, ok
}

func (p *AssociateIpAddressParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *AssociateIpAddressParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *AssociateIpAddressParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AssociateIpAddressParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *AssociateIpAddressParams) SetRegionid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["regionid"] = v
}

func (p *AssociateIpAddressParams) GetRegionid() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["regionid"].(int)
	return value, ok
}

func (p *AssociateIpAddressParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *AssociateIpAddressParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *AssociateIpAddressParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AssociateIpAddressParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AssociateIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewAssociateIpAddressParams() *AssociateIpAddressParams {
	p := &AssociateIpAddressParams{}
	p.p = make(map[string]interface{})
	return p
}

// Acquires and associates a public IP to an account.
func (s *AddressService) AssociateIpAddress(p *AssociateIpAddressParams) (*AssociateIpAddressResponse, error) {
	resp, err := s.cs.newRequest("associateIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssociateIpAddressResponse
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

type AssociateIpAddressResponse struct {
	Account                   string `json:"account"`
	Allocated                 string `json:"allocated"`
	Associatednetworkid       string `json:"associatednetworkid"`
	Associatednetworkname     string `json:"associatednetworkname"`
	Domain                    string `json:"domain"`
	Domainid                  string `json:"domainid"`
	Fordisplay                bool   `json:"fordisplay"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Isportable                bool   `json:"isportable"`
	Issourcenat               bool   `json:"issourcenat"`
	Isstaticnat               bool   `json:"isstaticnat"`
	Issystem                  bool   `json:"issystem"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Networkname               string `json:"networkname"`
	Physicalnetworkid         string `json:"physicalnetworkid"`
	Project                   string `json:"project"`
	Projectid                 string `json:"projectid"`
	Purpose                   string `json:"purpose"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vlanid                    string `json:"vlanid"`
	Vlanname                  string `json:"vlanname"`
	Vmipaddress               string `json:"vmipaddress"`
	Vpcid                     string `json:"vpcid"`
	Vpcname                   string `json:"vpcname"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type DisassociateIpAddressParams struct {
	p map[string]interface{}
}

func (p *DisassociateIpAddressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DisassociateIpAddressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DisassociateIpAddressParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DisassociateIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewDisassociateIpAddressParams(id string) *DisassociateIpAddressParams {
	p := &DisassociateIpAddressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Disassociates an IP address from the account.
func (s *AddressService) DisassociateIpAddress(p *DisassociateIpAddressParams) (*DisassociateIpAddressResponse, error) {
	resp, err := s.cs.newRequest("disassociateIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisassociateIpAddressResponse
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

type DisassociateIpAddressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListPublicIpAddressesParams struct {
	p map[string]interface{}
}

func (p *ListPublicIpAddressesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["allocatedonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("allocatedonly", vv)
	}
	if v, found := p.p["associatednetworkid"]; found {
		u.Set("associatednetworkid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["forloadbalancing"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forloadbalancing", vv)
	}
	if v, found := p.p["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["issourcenat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issourcenat", vv)
	}
	if v, found := p.p["isstaticnat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isstaticnat", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["vlanid"]; found {
		u.Set("vlanid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListPublicIpAddressesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListPublicIpAddressesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetAllocatedonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocatedonly"] = v
}

func (p *ListPublicIpAddressesParams) GetAllocatedonly() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocatedonly"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetAssociatednetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["associatednetworkid"] = v
}

func (p *ListPublicIpAddressesParams) GetAssociatednetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["associatednetworkid"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListPublicIpAddressesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListPublicIpAddressesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetForloadbalancing(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forloadbalancing"] = v
}

func (p *ListPublicIpAddressesParams) GetForloadbalancing() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forloadbalancing"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetForvirtualnetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvirtualnetwork"] = v
}

func (p *ListPublicIpAddressesParams) GetForvirtualnetwork() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forvirtualnetwork"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListPublicIpAddressesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *ListPublicIpAddressesParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListPublicIpAddressesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetIssourcenat(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issourcenat"] = v
}

func (p *ListPublicIpAddressesParams) GetIssourcenat() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["issourcenat"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetIsstaticnat(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isstaticnat"] = v
}

func (p *ListPublicIpAddressesParams) GetIsstaticnat() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isstaticnat"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListPublicIpAddressesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListPublicIpAddressesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListPublicIpAddressesParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListPublicIpAddressesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListPublicIpAddressesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListPublicIpAddressesParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListPublicIpAddressesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListPublicIpAddressesParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListPublicIpAddressesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetVlanid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlanid"] = v
}

func (p *ListPublicIpAddressesParams) GetVlanid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlanid"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListPublicIpAddressesParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListPublicIpAddressesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPublicIpAddressesParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewListPublicIpAddressesParams() *ListPublicIpAddressesParams {
	p := &ListPublicIpAddressesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AddressService) GetPublicIpAddressByID(id string, opts ...OptionFunc) (*PublicIpAddress, int, error) {
	p := &ListPublicIpAddressesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPublicIpAddresses(p)
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
		return l.PublicIpAddresses[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for PublicIpAddress UUID: %s!", id)
}

// Lists all public ip addresses
func (s *AddressService) ListPublicIpAddresses(p *ListPublicIpAddressesParams) (*ListPublicIpAddressesResponse, error) {
	resp, err := s.cs.newRequest("listPublicIpAddresses", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPublicIpAddressesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPublicIpAddressesResponse struct {
	Count             int                `json:"count"`
	PublicIpAddresses []*PublicIpAddress `json:"publicipaddress"`
}

type PublicIpAddress struct {
	Account                   string `json:"account"`
	Allocated                 string `json:"allocated"`
	Associatednetworkid       string `json:"associatednetworkid"`
	Associatednetworkname     string `json:"associatednetworkname"`
	Domain                    string `json:"domain"`
	Domainid                  string `json:"domainid"`
	Fordisplay                bool   `json:"fordisplay"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Isportable                bool   `json:"isportable"`
	Issourcenat               bool   `json:"issourcenat"`
	Isstaticnat               bool   `json:"isstaticnat"`
	Issystem                  bool   `json:"issystem"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Networkname               string `json:"networkname"`
	Physicalnetworkid         string `json:"physicalnetworkid"`
	Project                   string `json:"project"`
	Projectid                 string `json:"projectid"`
	Purpose                   string `json:"purpose"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vlanid                    string `json:"vlanid"`
	Vlanname                  string `json:"vlanname"`
	Vmipaddress               string `json:"vmipaddress"`
	Vpcid                     string `json:"vpcid"`
	Vpcname                   string `json:"vpcname"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type UpdateIpAddressParams struct {
	p map[string]interface{}
}

func (p *UpdateIpAddressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateIpAddressParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateIpAddressParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateIpAddressParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateIpAddressParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateIpAddressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateIpAddressParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewUpdateIpAddressParams(id string) *UpdateIpAddressParams {
	p := &UpdateIpAddressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an IP address
func (s *AddressService) UpdateIpAddress(p *UpdateIpAddressParams) (*UpdateIpAddressResponse, error) {
	resp, err := s.cs.newRequest("updateIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIpAddressResponse
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

type UpdateIpAddressResponse struct {
	Account                   string `json:"account"`
	Allocated                 string `json:"allocated"`
	Associatednetworkid       string `json:"associatednetworkid"`
	Associatednetworkname     string `json:"associatednetworkname"`
	Domain                    string `json:"domain"`
	Domainid                  string `json:"domainid"`
	Fordisplay                bool   `json:"fordisplay"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Isportable                bool   `json:"isportable"`
	Issourcenat               bool   `json:"issourcenat"`
	Isstaticnat               bool   `json:"isstaticnat"`
	Issystem                  bool   `json:"issystem"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Networkname               string `json:"networkname"`
	Physicalnetworkid         string `json:"physicalnetworkid"`
	Project                   string `json:"project"`
	Projectid                 string `json:"projectid"`
	Purpose                   string `json:"purpose"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vlanid                    string `json:"vlanid"`
	Vlanname                  string `json:"vlanname"`
	Vmipaddress               string `json:"vmipaddress"`
	Vpcid                     string `json:"vpcid"`
	Vpcname                   string `json:"vpcname"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}
