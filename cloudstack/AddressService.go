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
	AcquirePodIpAddress(p *AcquirePodIpAddressParams) (*AcquirePodIpAddressResponse, error)
	NewAcquirePodIpAddressParams(zoneid string) *AcquirePodIpAddressParams
	AssociateIpAddress(p *AssociateIpAddressParams) (*AssociateIpAddressResponse, error)
	NewAssociateIpAddressParams() *AssociateIpAddressParams
	DisassociateIpAddress(p *DisassociateIpAddressParams) (*DisassociateIpAddressResponse, error)
	NewDisassociateIpAddressParams(id string) *DisassociateIpAddressParams
	ListPublicIpAddresses(p *ListPublicIpAddressesParams) (*ListPublicIpAddressesResponse, error)
	NewListPublicIpAddressesParams() *ListPublicIpAddressesParams
	GetPublicIpAddressByID(id string, opts ...OptionFunc) (*PublicIpAddress, int, error)
	UpdateIpAddress(p *UpdateIpAddressParams) (*UpdateIpAddressResponse, error)
	NewUpdateIpAddressParams(id string) *UpdateIpAddressParams
	ReleaseIpAddress(p *ReleaseIpAddressParams) (*ReleaseIpAddressResponse, error)
	NewReleaseIpAddressParams(id string) *ReleaseIpAddressParams
	ReleasePodIpAddress(p *ReleasePodIpAddressParams) (*ReleasePodIpAddressResponse, error)
	NewReleasePodIpAddressParams(id int64) *ReleasePodIpAddressParams
	ReserveIpAddress(p *ReserveIpAddressParams) (*ReserveIpAddressResponse, error)
	NewReserveIpAddressParams(id string) *ReserveIpAddressParams
}

type AcquirePodIpAddressParams struct {
	p map[string]interface{}
}

func (p *AcquirePodIpAddressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AcquirePodIpAddressParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *AcquirePodIpAddressParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *AcquirePodIpAddressParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *AcquirePodIpAddressParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AcquirePodIpAddressParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AcquirePodIpAddressParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AcquirePodIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewAcquirePodIpAddressParams(zoneid string) *AcquirePodIpAddressParams {
	p := &AcquirePodIpAddressParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Allocates IP addresses in respective Pod of a Zone
func (s *AddressService) AcquirePodIpAddress(p *AcquirePodIpAddressParams) (*AcquirePodIpAddressResponse, error) {
	resp, err := s.cs.newPostRequest("acquirePodIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AcquirePodIpAddressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AcquirePodIpAddressResponse struct {
	Cidr      string `json:"cidr"`
	Gateway   string `json:"gateway"`
	Hostmac   int64  `json:"hostmac"`
	Id        int64  `json:"id"`
	Ipaddress string `json:"ipaddress"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nicid     int64  `json:"nicid"`
	Podid     int64  `json:"podid"`
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

func (p *AssociateIpAddressParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
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

func (p *AssociateIpAddressParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *AssociateIpAddressParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
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

func (p *AssociateIpAddressParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
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

func (p *AssociateIpAddressParams) ResetIsportable() {
	if p.p != nil && p.p["isportable"] != nil {
		delete(p.p, "isportable")
	}
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

func (p *AssociateIpAddressParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
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

func (p *AssociateIpAddressParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
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

func (p *AssociateIpAddressParams) ResetRegionid() {
	if p.p != nil && p.p["regionid"] != nil {
		delete(p.p, "regionid")
	}
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

func (p *AssociateIpAddressParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
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

func (p *AssociateIpAddressParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
	resp, err := s.cs.newPostRequest("associateIpAddress", p.toURLValues())
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
	Domainpath                string `json:"domainpath"`
	Fordisplay                bool   `json:"fordisplay"`
	Forprovider               bool   `json:"forprovider"`
	Forsystemvms              bool   `json:"forsystemvms"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Hasannotations            bool   `json:"hasannotations"`
	Hasrules                  bool   `json:"hasrules"`
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
	Virtualmachinetype        string `json:"virtualmachinetype"`
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
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	return u
}

func (p *DisassociateIpAddressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DisassociateIpAddressParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DisassociateIpAddressParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DisassociateIpAddressParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *DisassociateIpAddressParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *DisassociateIpAddressParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
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
	resp, err := s.cs.newPostRequest("disassociateIpAddress", p.toURLValues())
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
	if v, found := p.p["forprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forprovider", vv)
	}
	if v, found := p.p["forsystemvms"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forsystemvms", vv)
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
	if v, found := p.p["retrieveonlyresourcecount"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("retrieveonlyresourcecount", vv)
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

func (p *ListPublicIpAddressesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
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

func (p *ListPublicIpAddressesParams) ResetAllocatedonly() {
	if p.p != nil && p.p["allocatedonly"] != nil {
		delete(p.p, "allocatedonly")
	}
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

func (p *ListPublicIpAddressesParams) ResetAssociatednetworkid() {
	if p.p != nil && p.p["associatednetworkid"] != nil {
		delete(p.p, "associatednetworkid")
	}
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

func (p *ListPublicIpAddressesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *ListPublicIpAddressesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
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

func (p *ListPublicIpAddressesParams) ResetForloadbalancing() {
	if p.p != nil && p.p["forloadbalancing"] != nil {
		delete(p.p, "forloadbalancing")
	}
}

func (p *ListPublicIpAddressesParams) GetForloadbalancing() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forloadbalancing"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetForprovider(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forprovider"] = v
}

func (p *ListPublicIpAddressesParams) ResetForprovider() {
	if p.p != nil && p.p["forprovider"] != nil {
		delete(p.p, "forprovider")
	}
}

func (p *ListPublicIpAddressesParams) GetForprovider() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forprovider"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetForsystemvms(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forsystemvms"] = v
}

func (p *ListPublicIpAddressesParams) ResetForsystemvms() {
	if p.p != nil && p.p["forsystemvms"] != nil {
		delete(p.p, "forsystemvms")
	}
}

func (p *ListPublicIpAddressesParams) GetForsystemvms() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forsystemvms"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetForvirtualnetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvirtualnetwork"] = v
}

func (p *ListPublicIpAddressesParams) ResetForvirtualnetwork() {
	if p.p != nil && p.p["forvirtualnetwork"] != nil {
		delete(p.p, "forvirtualnetwork")
	}
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

func (p *ListPublicIpAddressesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *ListPublicIpAddressesParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
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

func (p *ListPublicIpAddressesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
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

func (p *ListPublicIpAddressesParams) ResetIssourcenat() {
	if p.p != nil && p.p["issourcenat"] != nil {
		delete(p.p, "issourcenat")
	}
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

func (p *ListPublicIpAddressesParams) ResetIsstaticnat() {
	if p.p != nil && p.p["isstaticnat"] != nil {
		delete(p.p, "isstaticnat")
	}
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

func (p *ListPublicIpAddressesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListPublicIpAddressesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
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

func (p *ListPublicIpAddressesParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
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

func (p *ListPublicIpAddressesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListPublicIpAddressesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

func (p *ListPublicIpAddressesParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
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

func (p *ListPublicIpAddressesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListPublicIpAddressesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetRetrieveonlyresourcecount(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["retrieveonlyresourcecount"] = v
}

func (p *ListPublicIpAddressesParams) ResetRetrieveonlyresourcecount() {
	if p.p != nil && p.p["retrieveonlyresourcecount"] != nil {
		delete(p.p, "retrieveonlyresourcecount")
	}
}

func (p *ListPublicIpAddressesParams) GetRetrieveonlyresourcecount() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["retrieveonlyresourcecount"].(bool)
	return value, ok
}

func (p *ListPublicIpAddressesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListPublicIpAddressesParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
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

func (p *ListPublicIpAddressesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
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

func (p *ListPublicIpAddressesParams) ResetVlanid() {
	if p.p != nil && p.p["vlanid"] != nil {
		delete(p.p, "vlanid")
	}
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

func (p *ListPublicIpAddressesParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
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

func (p *ListPublicIpAddressesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
	Domainpath                string `json:"domainpath"`
	Fordisplay                bool   `json:"fordisplay"`
	Forprovider               bool   `json:"forprovider"`
	Forsystemvms              bool   `json:"forsystemvms"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Hasannotations            bool   `json:"hasannotations"`
	Hasrules                  bool   `json:"hasrules"`
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
	Virtualmachinetype        string `json:"virtualmachinetype"`
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

func (p *UpdateIpAddressParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
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

func (p *UpdateIpAddressParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
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

func (p *UpdateIpAddressParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
	resp, err := s.cs.newPostRequest("updateIpAddress", p.toURLValues())
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
	Domainpath                string `json:"domainpath"`
	Fordisplay                bool   `json:"fordisplay"`
	Forprovider               bool   `json:"forprovider"`
	Forsystemvms              bool   `json:"forsystemvms"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Hasannotations            bool   `json:"hasannotations"`
	Hasrules                  bool   `json:"hasrules"`
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
	Virtualmachinetype        string `json:"virtualmachinetype"`
	Vlanid                    string `json:"vlanid"`
	Vlanname                  string `json:"vlanname"`
	Vmipaddress               string `json:"vmipaddress"`
	Vpcid                     string `json:"vpcid"`
	Vpcname                   string `json:"vpcname"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type ReleaseIpAddressParams struct {
	p map[string]interface{}
}

func (p *ReleaseIpAddressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ReleaseIpAddressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ReleaseIpAddressParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ReleaseIpAddressParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewReleaseIpAddressParams(id string) *ReleaseIpAddressParams {
	p := &ReleaseIpAddressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Releases an IP address from the account.
func (s *AddressService) ReleaseIpAddress(p *ReleaseIpAddressParams) (*ReleaseIpAddressResponse, error) {
	resp, err := s.cs.newPostRequest("releaseIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseIpAddressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ReleaseIpAddressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ReleaseIpAddressResponse) UnmarshalJSON(b []byte) error {
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

	type alias ReleaseIpAddressResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ReleasePodIpAddressParams struct {
	p map[string]interface{}
}

func (p *ReleasePodIpAddressParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
	}
	return u
}

func (p *ReleasePodIpAddressParams) SetId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ReleasePodIpAddressParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ReleasePodIpAddressParams) GetId() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int64)
	return value, ok
}

// You should always use this function to get a new ReleasePodIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewReleasePodIpAddressParams(id int64) *ReleasePodIpAddressParams {
	p := &ReleasePodIpAddressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Releases a Pod IP back to the Pod
func (s *AddressService) ReleasePodIpAddress(p *ReleasePodIpAddressParams) (*ReleasePodIpAddressResponse, error) {
	resp, err := s.cs.newPostRequest("releasePodIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleasePodIpAddressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ReleasePodIpAddressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ReleasePodIpAddressResponse) UnmarshalJSON(b []byte) error {
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

	type alias ReleasePodIpAddressResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ReserveIpAddressParams struct {
	p map[string]interface{}
}

func (p *ReserveIpAddressParams) toURLValues() url.Values {
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
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *ReserveIpAddressParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ReserveIpAddressParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ReserveIpAddressParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ReserveIpAddressParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ReserveIpAddressParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ReserveIpAddressParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ReserveIpAddressParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ReserveIpAddressParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ReserveIpAddressParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ReserveIpAddressParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ReserveIpAddressParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ReserveIpAddressParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ReserveIpAddressParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ReserveIpAddressParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ReserveIpAddressParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ReserveIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewReserveIpAddressParams(id string) *ReserveIpAddressParams {
	p := &ReserveIpAddressParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reserve a public IP to an account.
func (s *AddressService) ReserveIpAddress(p *ReserveIpAddressParams) (*ReserveIpAddressResponse, error) {
	resp, err := s.cs.newPostRequest("reserveIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReserveIpAddressResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ReserveIpAddressResponse struct {
	Account                   string `json:"account"`
	Allocated                 string `json:"allocated"`
	Associatednetworkid       string `json:"associatednetworkid"`
	Associatednetworkname     string `json:"associatednetworkname"`
	Domain                    string `json:"domain"`
	Domainid                  string `json:"domainid"`
	Domainpath                string `json:"domainpath"`
	Fordisplay                bool   `json:"fordisplay"`
	Forprovider               bool   `json:"forprovider"`
	Forsystemvms              bool   `json:"forsystemvms"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Hasannotations            bool   `json:"hasannotations"`
	Hasrules                  bool   `json:"hasrules"`
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
	Virtualmachinetype        string `json:"virtualmachinetype"`
	Vlanid                    string `json:"vlanid"`
	Vlanname                  string `json:"vlanname"`
	Vmipaddress               string `json:"vmipaddress"`
	Vpcid                     string `json:"vpcid"`
	Vpcname                   string `json:"vpcname"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}
