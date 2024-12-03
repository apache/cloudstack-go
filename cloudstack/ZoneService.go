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

type ZoneServiceIface interface {
	CreateZone(p *CreateZoneParams) (*CreateZoneResponse, error)
	NewCreateZoneParams(dns1 string, internaldns1 string, name string, networktype string) *CreateZoneParams
	DedicateZone(p *DedicateZoneParams) (*DedicateZoneResponse, error)
	NewDedicateZoneParams(domainid string, zoneid string) *DedicateZoneParams
	DeleteZone(p *DeleteZoneParams) (*DeleteZoneResponse, error)
	NewDeleteZoneParams(id string) *DeleteZoneParams
	DisableOutOfBandManagementForZone(p *DisableOutOfBandManagementForZoneParams) (*DisableOutOfBandManagementForZoneResponse, error)
	NewDisableOutOfBandManagementForZoneParams(zoneid string) *DisableOutOfBandManagementForZoneParams
	EnableOutOfBandManagementForZone(p *EnableOutOfBandManagementForZoneParams) (*EnableOutOfBandManagementForZoneResponse, error)
	NewEnableOutOfBandManagementForZoneParams(zoneid string) *EnableOutOfBandManagementForZoneParams
	DisableHAForZone(p *DisableHAForZoneParams) (*DisableHAForZoneResponse, error)
	NewDisableHAForZoneParams(zoneid string) *DisableHAForZoneParams
	EnableHAForZone(p *EnableHAForZoneParams) (*EnableHAForZoneResponse, error)
	NewEnableHAForZoneParams(zoneid string) *EnableHAForZoneParams
	ListDedicatedZones(p *ListDedicatedZonesParams) (*ListDedicatedZonesResponse, error)
	NewListDedicatedZonesParams() *ListDedicatedZonesParams
	ListZones(p *ListZonesParams) (*ListZonesResponse, error)
	NewListZonesParams() *ListZonesParams
	GetZoneID(name string, opts ...OptionFunc) (string, int, error)
	GetZoneByName(name string, opts ...OptionFunc) (*Zone, int, error)
	GetZoneByID(id string, opts ...OptionFunc) (*Zone, int, error)
	ListZonesMetrics(p *ListZonesMetricsParams) (*ListZonesMetricsResponse, error)
	NewListZonesMetricsParams() *ListZonesMetricsParams
	GetZonesMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetZonesMetricByName(name string, opts ...OptionFunc) (*ZonesMetric, int, error)
	GetZonesMetricByID(id string, opts ...OptionFunc) (*ZonesMetric, int, error)
	ReleaseDedicatedZone(p *ReleaseDedicatedZoneParams) (*ReleaseDedicatedZoneResponse, error)
	NewReleaseDedicatedZoneParams(zoneid string) *ReleaseDedicatedZoneParams
	UpdateZone(p *UpdateZoneParams) (*UpdateZoneResponse, error)
	NewUpdateZoneParams(id string) *UpdateZoneParams
}

type CreateZoneParams struct {
	p map[string]interface{}
}

func (p *CreateZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["dns1"]; found {
		u.Set("dns1", v.(string))
	}
	if v, found := p.p["dns2"]; found {
		u.Set("dns2", v.(string))
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["guestcidraddress"]; found {
		u.Set("guestcidraddress", v.(string))
	}
	if v, found := p.p["internaldns1"]; found {
		u.Set("internaldns1", v.(string))
	}
	if v, found := p.p["internaldns2"]; found {
		u.Set("internaldns2", v.(string))
	}
	if v, found := p.p["ip6dns1"]; found {
		u.Set("ip6dns1", v.(string))
	}
	if v, found := p.p["ip6dns2"]; found {
		u.Set("ip6dns2", v.(string))
	}
	if v, found := p.p["isedge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isedge", vv)
	}
	if v, found := p.p["localstorageenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("localstorageenabled", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networktype"]; found {
		u.Set("networktype", v.(string))
	}
	if v, found := p.p["securitygroupenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("securitygroupenabled", vv)
	}
	return u
}

func (p *CreateZoneParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *CreateZoneParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *CreateZoneParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetDns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns1"] = v
}

func (p *CreateZoneParams) ResetDns1() {
	if p.p != nil && p.p["dns1"] != nil {
		delete(p.p, "dns1")
	}
}

func (p *CreateZoneParams) GetDns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns1"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetDns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns2"] = v
}

func (p *CreateZoneParams) ResetDns2() {
	if p.p != nil && p.p["dns2"] != nil {
		delete(p.p, "dns2")
	}
}

func (p *CreateZoneParams) GetDns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns2"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *CreateZoneParams) ResetDomain() {
	if p.p != nil && p.p["domain"] != nil {
		delete(p.p, "domain")
	}
}

func (p *CreateZoneParams) GetDomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domain"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateZoneParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateZoneParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetGuestcidraddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestcidraddress"] = v
}

func (p *CreateZoneParams) ResetGuestcidraddress() {
	if p.p != nil && p.p["guestcidraddress"] != nil {
		delete(p.p, "guestcidraddress")
	}
}

func (p *CreateZoneParams) GetGuestcidraddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["guestcidraddress"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetInternaldns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internaldns1"] = v
}

func (p *CreateZoneParams) ResetInternaldns1() {
	if p.p != nil && p.p["internaldns1"] != nil {
		delete(p.p, "internaldns1")
	}
}

func (p *CreateZoneParams) GetInternaldns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["internaldns1"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetInternaldns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internaldns2"] = v
}

func (p *CreateZoneParams) ResetInternaldns2() {
	if p.p != nil && p.p["internaldns2"] != nil {
		delete(p.p, "internaldns2")
	}
}

func (p *CreateZoneParams) GetInternaldns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["internaldns2"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetIp6dns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns1"] = v
}

func (p *CreateZoneParams) ResetIp6dns1() {
	if p.p != nil && p.p["ip6dns1"] != nil {
		delete(p.p, "ip6dns1")
	}
}

func (p *CreateZoneParams) GetIp6dns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns1"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetIp6dns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns2"] = v
}

func (p *CreateZoneParams) ResetIp6dns2() {
	if p.p != nil && p.p["ip6dns2"] != nil {
		delete(p.p, "ip6dns2")
	}
}

func (p *CreateZoneParams) GetIp6dns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns2"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetIsedge(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isedge"] = v
}

func (p *CreateZoneParams) ResetIsedge() {
	if p.p != nil && p.p["isedge"] != nil {
		delete(p.p, "isedge")
	}
}

func (p *CreateZoneParams) GetIsedge() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isedge"].(bool)
	return value, ok
}

func (p *CreateZoneParams) SetLocalstorageenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["localstorageenabled"] = v
}

func (p *CreateZoneParams) ResetLocalstorageenabled() {
	if p.p != nil && p.p["localstorageenabled"] != nil {
		delete(p.p, "localstorageenabled")
	}
}

func (p *CreateZoneParams) GetLocalstorageenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["localstorageenabled"].(bool)
	return value, ok
}

func (p *CreateZoneParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateZoneParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateZoneParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetNetworktype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networktype"] = v
}

func (p *CreateZoneParams) ResetNetworktype() {
	if p.p != nil && p.p["networktype"] != nil {
		delete(p.p, "networktype")
	}
}

func (p *CreateZoneParams) GetNetworktype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networktype"].(string)
	return value, ok
}

func (p *CreateZoneParams) SetSecuritygroupenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupenabled"] = v
}

func (p *CreateZoneParams) ResetSecuritygroupenabled() {
	if p.p != nil && p.p["securitygroupenabled"] != nil {
		delete(p.p, "securitygroupenabled")
	}
}

func (p *CreateZoneParams) GetSecuritygroupenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupenabled"].(bool)
	return value, ok
}

// You should always use this function to get a new CreateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewCreateZoneParams(dns1 string, internaldns1 string, name string, networktype string) *CreateZoneParams {
	p := &CreateZoneParams{}
	p.p = make(map[string]interface{})
	p.p["dns1"] = dns1
	p.p["internaldns1"] = internaldns1
	p.p["name"] = name
	p.p["networktype"] = networktype
	return p
}

// Creates a Zone.
func (s *ZoneService) CreateZone(p *CreateZoneParams) (*CreateZoneResponse, error) {
	resp, err := s.cs.newRequest("createZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateZoneResponse struct {
	Allocationstate              string                       `json:"allocationstate"`
	Allowuserspecifyvrmtu        bool                         `json:"allowuserspecifyvrmtu"`
	Asnrange                     string                       `json:"asnrange"`
	Capacity                     []CreateZoneResponseCapacity `json:"capacity"`
	Description                  string                       `json:"description"`
	Dhcpprovider                 string                       `json:"dhcpprovider"`
	Displaytext                  string                       `json:"displaytext"`
	Dns1                         string                       `json:"dns1"`
	Dns2                         string                       `json:"dns2"`
	Domain                       string                       `json:"domain"`
	Domainid                     string                       `json:"domainid"`
	Domainname                   string                       `json:"domainname"`
	Guestcidraddress             string                       `json:"guestcidraddress"`
	Hasannotations               bool                         `json:"hasannotations"`
	Icon                         interface{}                  `json:"icon"`
	Id                           string                       `json:"id"`
	Internaldns1                 string                       `json:"internaldns1"`
	Internaldns2                 string                       `json:"internaldns2"`
	Ip6dns1                      string                       `json:"ip6dns1"`
	Ip6dns2                      string                       `json:"ip6dns2"`
	Ismultiarch                  bool                         `json:"ismultiarch"`
	Isnsxenabled                 bool                         `json:"isnsxenabled"`
	JobID                        string                       `json:"jobid"`
	Jobstatus                    int                          `json:"jobstatus"`
	Localstorageenabled          bool                         `json:"localstorageenabled"`
	Name                         string                       `json:"name"`
	Networktype                  string                       `json:"networktype"`
	Resourcedetails              map[string]string            `json:"resourcedetails"`
	Routerprivateinterfacemaxmtu int                          `json:"routerprivateinterfacemaxmtu"`
	Routerpublicinterfacemaxmtu  int                          `json:"routerpublicinterfacemaxmtu"`
	Securitygroupsenabled        bool                         `json:"securitygroupsenabled"`
	Tags                         []Tags                       `json:"tags"`
	Type                         string                       `json:"type"`
	Zonetoken                    string                       `json:"zonetoken"`
}

type CreateZoneResponseCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Tag               string `json:"tag"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type DedicateZoneParams struct {
	p map[string]interface{}
}

func (p *DedicateZoneParams) toURLValues() url.Values {
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DedicateZoneParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DedicateZoneParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DedicateZoneParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DedicateZoneParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DedicateZoneParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DedicateZoneParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DedicateZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DedicateZoneParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DedicateZoneParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDedicateZoneParams(domainid string, zoneid string) *DedicateZoneParams {
	p := &DedicateZoneParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	p.p["zoneid"] = zoneid
	return p
}

// Dedicates a zones.
func (s *ZoneService) DedicateZone(p *DedicateZoneParams) (*DedicateZoneResponse, error) {
	resp, err := s.cs.newRequest("dedicateZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateZoneResponse
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

type DedicateZoneResponse struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type DeleteZoneParams struct {
	p map[string]interface{}
}

func (p *DeleteZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteZoneParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteZoneParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteZoneParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDeleteZoneParams(id string) *DeleteZoneParams {
	p := &DeleteZoneParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Zone.
func (s *ZoneService) DeleteZone(p *DeleteZoneParams) (*DeleteZoneResponse, error) {
	resp, err := s.cs.newRequest("deleteZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteZoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteZoneResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteZoneResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableOutOfBandManagementForZoneParams struct {
	p map[string]interface{}
}

func (p *DisableOutOfBandManagementForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DisableOutOfBandManagementForZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DisableOutOfBandManagementForZoneParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DisableOutOfBandManagementForZoneParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableOutOfBandManagementForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDisableOutOfBandManagementForZoneParams(zoneid string) *DisableOutOfBandManagementForZoneParams {
	p := &DisableOutOfBandManagementForZoneParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Disables out-of-band management for a zone
func (s *ZoneService) DisableOutOfBandManagementForZone(p *DisableOutOfBandManagementForZoneParams) (*DisableOutOfBandManagementForZoneResponse, error) {
	resp, err := s.cs.newRequest("disableOutOfBandManagementForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableOutOfBandManagementForZoneResponse
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

type DisableOutOfBandManagementForZoneResponse struct {
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type EnableOutOfBandManagementForZoneParams struct {
	p map[string]interface{}
}

func (p *EnableOutOfBandManagementForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *EnableOutOfBandManagementForZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *EnableOutOfBandManagementForZoneParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *EnableOutOfBandManagementForZoneParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableOutOfBandManagementForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewEnableOutOfBandManagementForZoneParams(zoneid string) *EnableOutOfBandManagementForZoneParams {
	p := &EnableOutOfBandManagementForZoneParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Enables out-of-band management for a zone
func (s *ZoneService) EnableOutOfBandManagementForZone(p *EnableOutOfBandManagementForZoneParams) (*EnableOutOfBandManagementForZoneResponse, error) {
	resp, err := s.cs.newRequest("enableOutOfBandManagementForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableOutOfBandManagementForZoneResponse
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

type EnableOutOfBandManagementForZoneResponse struct {
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type DisableHAForZoneParams struct {
	p map[string]interface{}
}

func (p *DisableHAForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DisableHAForZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DisableHAForZoneParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DisableHAForZoneParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableHAForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDisableHAForZoneParams(zoneid string) *DisableHAForZoneParams {
	p := &DisableHAForZoneParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Disables HA for a zone
func (s *ZoneService) DisableHAForZone(p *DisableHAForZoneParams) (*DisableHAForZoneResponse, error) {
	resp, err := s.cs.newRequest("disableHAForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableHAForZoneResponse
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

type DisableHAForZoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type EnableHAForZoneParams struct {
	p map[string]interface{}
}

func (p *EnableHAForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *EnableHAForZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *EnableHAForZoneParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *EnableHAForZoneParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableHAForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewEnableHAForZoneParams(zoneid string) *EnableHAForZoneParams {
	p := &EnableHAForZoneParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Enables HA for a zone
func (s *ZoneService) EnableHAForZone(p *EnableHAForZoneParams) (*EnableHAForZoneResponse, error) {
	resp, err := s.cs.newRequest("enableHAForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableHAForZoneResponse
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

type EnableHAForZoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListDedicatedZonesParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedZonesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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

func (p *ListDedicatedZonesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListDedicatedZonesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListDedicatedZonesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListDedicatedZonesParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
}

func (p *ListDedicatedZonesParams) ResetAffinitygroupid() {
	if p.p != nil && p.p["affinitygroupid"] != nil {
		delete(p.p, "affinitygroupid")
	}
}

func (p *ListDedicatedZonesParams) GetAffinitygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupid"].(string)
	return value, ok
}

func (p *ListDedicatedZonesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListDedicatedZonesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListDedicatedZonesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListDedicatedZonesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListDedicatedZonesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListDedicatedZonesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListDedicatedZonesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListDedicatedZonesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListDedicatedZonesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListDedicatedZonesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListDedicatedZonesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListDedicatedZonesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListDedicatedZonesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListDedicatedZonesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListDedicatedZonesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListDedicatedZonesParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewListDedicatedZonesParams() *ListDedicatedZonesParams {
	p := &ListDedicatedZonesParams{}
	p.p = make(map[string]interface{})
	return p
}

// List dedicated zones.
func (s *ZoneService) ListDedicatedZones(p *ListDedicatedZonesParams) (*ListDedicatedZonesResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedZones", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedZonesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedZonesResponse struct {
	Count          int              `json:"count"`
	DedicatedZones []*DedicatedZone `json:"dedicatedzone"`
}

type DedicatedZone struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type ListZonesParams struct {
	p map[string]interface{}
}

func (p *ListZonesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["available"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("available", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networktype"]; found {
		u.Set("networktype", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListZonesParams) SetAvailable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["available"] = v
}

func (p *ListZonesParams) ResetAvailable() {
	if p.p != nil && p.p["available"] != nil {
		delete(p.p, "available")
	}
}

func (p *ListZonesParams) GetAvailable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["available"].(bool)
	return value, ok
}

func (p *ListZonesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListZonesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListZonesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListZonesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListZonesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListZonesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListZonesParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListZonesParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListZonesParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListZonesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListZonesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListZonesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListZonesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListZonesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListZonesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListZonesParams) SetNetworktype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networktype"] = v
}

func (p *ListZonesParams) ResetNetworktype() {
	if p.p != nil && p.p["networktype"] != nil {
		delete(p.p, "networktype")
	}
}

func (p *ListZonesParams) GetNetworktype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networktype"].(string)
	return value, ok
}

func (p *ListZonesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListZonesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListZonesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListZonesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListZonesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListZonesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListZonesParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
}

func (p *ListZonesParams) ResetShowcapacities() {
	if p.p != nil && p.p["showcapacities"] != nil {
		delete(p.p, "showcapacities")
	}
}

func (p *ListZonesParams) GetShowcapacities() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showcapacities"].(bool)
	return value, ok
}

func (p *ListZonesParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListZonesParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListZonesParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListZonesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListZonesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListZonesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListZonesParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewListZonesParams() *ListZonesParams {
	p := &ListZonesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListZonesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListZones(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Zones[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Zones {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneByName(name string, opts ...OptionFunc) (*Zone, int, error) {
	id, count, err := s.GetZoneID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetZoneByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneByID(id string, opts ...OptionFunc) (*Zone, int, error) {
	p := &ListZonesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListZones(p)
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
		return l.Zones[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Zone UUID: %s!", id)
}

// Lists zones
func (s *ZoneService) ListZones(p *ListZonesParams) (*ListZonesResponse, error) {
	resp, err := s.cs.newRequest("listZones", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListZonesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListZonesResponse struct {
	Count int     `json:"count"`
	Zones []*Zone `json:"zone"`
}

type Zone struct {
	Allocationstate              string            `json:"allocationstate"`
	Allowuserspecifyvrmtu        bool              `json:"allowuserspecifyvrmtu"`
	Asnrange                     string            `json:"asnrange"`
	Capacity                     []ZoneCapacity    `json:"capacity"`
	Description                  string            `json:"description"`
	Dhcpprovider                 string            `json:"dhcpprovider"`
	Displaytext                  string            `json:"displaytext"`
	Dns1                         string            `json:"dns1"`
	Dns2                         string            `json:"dns2"`
	Domain                       string            `json:"domain"`
	Domainid                     string            `json:"domainid"`
	Domainname                   string            `json:"domainname"`
	Guestcidraddress             string            `json:"guestcidraddress"`
	Hasannotations               bool              `json:"hasannotations"`
	Icon                         interface{}       `json:"icon"`
	Id                           string            `json:"id"`
	Internaldns1                 string            `json:"internaldns1"`
	Internaldns2                 string            `json:"internaldns2"`
	Ip6dns1                      string            `json:"ip6dns1"`
	Ip6dns2                      string            `json:"ip6dns2"`
	Ismultiarch                  bool              `json:"ismultiarch"`
	Isnsxenabled                 bool              `json:"isnsxenabled"`
	JobID                        string            `json:"jobid"`
	Jobstatus                    int               `json:"jobstatus"`
	Localstorageenabled          bool              `json:"localstorageenabled"`
	Name                         string            `json:"name"`
	Networktype                  string            `json:"networktype"`
	Resourcedetails              map[string]string `json:"resourcedetails"`
	Routerprivateinterfacemaxmtu int               `json:"routerprivateinterfacemaxmtu"`
	Routerpublicinterfacemaxmtu  int               `json:"routerpublicinterfacemaxmtu"`
	Securitygroupsenabled        bool              `json:"securitygroupsenabled"`
	Tags                         []Tags            `json:"tags"`
	Type                         string            `json:"type"`
	Zonetoken                    string            `json:"zonetoken"`
}

type ZoneCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Tag               string `json:"tag"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type ListZonesMetricsParams struct {
	p map[string]interface{}
}

func (p *ListZonesMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["available"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("available", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networktype"]; found {
		u.Set("networktype", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ListZonesMetricsParams) SetAvailable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["available"] = v
}

func (p *ListZonesMetricsParams) ResetAvailable() {
	if p.p != nil && p.p["available"] != nil {
		delete(p.p, "available")
	}
}

func (p *ListZonesMetricsParams) GetAvailable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["available"].(bool)
	return value, ok
}

func (p *ListZonesMetricsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListZonesMetricsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListZonesMetricsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListZonesMetricsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListZonesMetricsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListZonesMetricsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListZonesMetricsParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListZonesMetricsParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListZonesMetricsParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListZonesMetricsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListZonesMetricsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListZonesMetricsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListZonesMetricsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListZonesMetricsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListZonesMetricsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListZonesMetricsParams) SetNetworktype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networktype"] = v
}

func (p *ListZonesMetricsParams) ResetNetworktype() {
	if p.p != nil && p.p["networktype"] != nil {
		delete(p.p, "networktype")
	}
}

func (p *ListZonesMetricsParams) GetNetworktype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networktype"].(string)
	return value, ok
}

func (p *ListZonesMetricsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListZonesMetricsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListZonesMetricsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListZonesMetricsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListZonesMetricsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListZonesMetricsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListZonesMetricsParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
}

func (p *ListZonesMetricsParams) ResetShowcapacities() {
	if p.p != nil && p.p["showcapacities"] != nil {
		delete(p.p, "showcapacities")
	}
}

func (p *ListZonesMetricsParams) GetShowcapacities() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showcapacities"].(bool)
	return value, ok
}

func (p *ListZonesMetricsParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListZonesMetricsParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListZonesMetricsParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListZonesMetricsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListZonesMetricsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListZonesMetricsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListZonesMetricsParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewListZonesMetricsParams() *ListZonesMetricsParams {
	p := &ListZonesMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZonesMetricID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListZonesMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListZonesMetrics(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ZonesMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ZonesMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZonesMetricByName(name string, opts ...OptionFunc) (*ZonesMetric, int, error) {
	id, count, err := s.GetZonesMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetZonesMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZonesMetricByID(id string, opts ...OptionFunc) (*ZonesMetric, int, error) {
	p := &ListZonesMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListZonesMetrics(p)
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
		return l.ZonesMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ZonesMetric UUID: %s!", id)
}

// Lists zone metrics
func (s *ZoneService) ListZonesMetrics(p *ListZonesMetricsParams) (*ListZonesMetricsResponse, error) {
	resp, err := s.cs.newRequest("listZonesMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListZonesMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListZonesMetricsResponse struct {
	Count        int            `json:"count"`
	ZonesMetrics []*ZonesMetric `json:"zonesmetric"`
}

type ZonesMetric struct {
	Allocationstate                 string                `json:"allocationstate"`
	Allowuserspecifyvrmtu           bool                  `json:"allowuserspecifyvrmtu"`
	Asnrange                        string                `json:"asnrange"`
	Capacity                        []ZonesMetricCapacity `json:"capacity"`
	Clusters                        string                `json:"clusters"`
	Cpuallocated                    string                `json:"cpuallocated"`
	Cpuallocateddisablethreshold    bool                  `json:"cpuallocateddisablethreshold"`
	Cpuallocatedthreshold           bool                  `json:"cpuallocatedthreshold"`
	Cpudisablethreshold             bool                  `json:"cpudisablethreshold"`
	Cpumaxdeviation                 string                `json:"cpumaxdeviation"`
	Cputhreshold                    bool                  `json:"cputhreshold"`
	Cputotal                        string                `json:"cputotal"`
	Cpuused                         string                `json:"cpuused"`
	Description                     string                `json:"description"`
	Dhcpprovider                    string                `json:"dhcpprovider"`
	Displaytext                     string                `json:"displaytext"`
	Dns1                            string                `json:"dns1"`
	Dns2                            string                `json:"dns2"`
	Domain                          string                `json:"domain"`
	Domainid                        string                `json:"domainid"`
	Domainname                      string                `json:"domainname"`
	Guestcidraddress                string                `json:"guestcidraddress"`
	Hasannotations                  bool                  `json:"hasannotations"`
	Icon                            interface{}           `json:"icon"`
	Id                              string                `json:"id"`
	Internaldns1                    string                `json:"internaldns1"`
	Internaldns2                    string                `json:"internaldns2"`
	Ip6dns1                         string                `json:"ip6dns1"`
	Ip6dns2                         string                `json:"ip6dns2"`
	Ismultiarch                     bool                  `json:"ismultiarch"`
	Isnsxenabled                    bool                  `json:"isnsxenabled"`
	JobID                           string                `json:"jobid"`
	Jobstatus                       int                   `json:"jobstatus"`
	Localstorageenabled             bool                  `json:"localstorageenabled"`
	Memoryallocated                 string                `json:"memoryallocated"`
	Memoryallocateddisablethreshold bool                  `json:"memoryallocateddisablethreshold"`
	Memoryallocatedthreshold        bool                  `json:"memoryallocatedthreshold"`
	Memorydisablethreshold          bool                  `json:"memorydisablethreshold"`
	Memorymaxdeviation              string                `json:"memorymaxdeviation"`
	Memorythreshold                 bool                  `json:"memorythreshold"`
	Memorytotal                     string                `json:"memorytotal"`
	Memoryused                      string                `json:"memoryused"`
	Name                            string                `json:"name"`
	Networktype                     string                `json:"networktype"`
	Resourcedetails                 map[string]string     `json:"resourcedetails"`
	Routerprivateinterfacemaxmtu    int                   `json:"routerprivateinterfacemaxmtu"`
	Routerpublicinterfacemaxmtu     int                   `json:"routerpublicinterfacemaxmtu"`
	Securitygroupsenabled           bool                  `json:"securitygroupsenabled"`
	State                           string                `json:"state"`
	Tags                            []Tags                `json:"tags"`
	Type                            string                `json:"type"`
	Zonetoken                       string                `json:"zonetoken"`
}

type ZonesMetricCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Tag               string `json:"tag"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type ReleaseDedicatedZoneParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedZoneParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ReleaseDedicatedZoneParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ReleaseDedicatedZoneParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewReleaseDedicatedZoneParams(zoneid string) *ReleaseDedicatedZoneParams {
	p := &ReleaseDedicatedZoneParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Release dedication of zone
func (s *ZoneService) ReleaseDedicatedZone(p *ReleaseDedicatedZoneParams) (*ReleaseDedicatedZoneResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedZoneResponse
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

type ReleaseDedicatedZoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateZoneParams struct {
	p map[string]interface{}
}

func (p *UpdateZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), m[k])
		}
	}
	if v, found := p.p["dhcpprovider"]; found {
		u.Set("dhcpprovider", v.(string))
	}
	if v, found := p.p["dns1"]; found {
		u.Set("dns1", v.(string))
	}
	if v, found := p.p["dns2"]; found {
		u.Set("dns2", v.(string))
	}
	if v, found := p.p["dnssearchorder"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("dnssearchorder", vv)
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["guestcidraddress"]; found {
		u.Set("guestcidraddress", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["internaldns1"]; found {
		u.Set("internaldns1", v.(string))
	}
	if v, found := p.p["internaldns2"]; found {
		u.Set("internaldns2", v.(string))
	}
	if v, found := p.p["ip6dns1"]; found {
		u.Set("ip6dns1", v.(string))
	}
	if v, found := p.p["ip6dns2"]; found {
		u.Set("ip6dns2", v.(string))
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["localstorageenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("localstorageenabled", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	return u
}

func (p *UpdateZoneParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *UpdateZoneParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *UpdateZoneParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateZoneParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateZoneParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateZoneParams) SetDhcpprovider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpprovider"] = v
}

func (p *UpdateZoneParams) ResetDhcpprovider() {
	if p.p != nil && p.p["dhcpprovider"] != nil {
		delete(p.p, "dhcpprovider")
	}
}

func (p *UpdateZoneParams) GetDhcpprovider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dhcpprovider"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetDns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns1"] = v
}

func (p *UpdateZoneParams) ResetDns1() {
	if p.p != nil && p.p["dns1"] != nil {
		delete(p.p, "dns1")
	}
}

func (p *UpdateZoneParams) GetDns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns1"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetDns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns2"] = v
}

func (p *UpdateZoneParams) ResetDns2() {
	if p.p != nil && p.p["dns2"] != nil {
		delete(p.p, "dns2")
	}
}

func (p *UpdateZoneParams) GetDns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns2"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetDnssearchorder(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dnssearchorder"] = v
}

func (p *UpdateZoneParams) ResetDnssearchorder() {
	if p.p != nil && p.p["dnssearchorder"] != nil {
		delete(p.p, "dnssearchorder")
	}
}

func (p *UpdateZoneParams) GetDnssearchorder() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dnssearchorder"].([]string)
	return value, ok
}

func (p *UpdateZoneParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *UpdateZoneParams) ResetDomain() {
	if p.p != nil && p.p["domain"] != nil {
		delete(p.p, "domain")
	}
}

func (p *UpdateZoneParams) GetDomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domain"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetGuestcidraddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestcidraddress"] = v
}

func (p *UpdateZoneParams) ResetGuestcidraddress() {
	if p.p != nil && p.p["guestcidraddress"] != nil {
		delete(p.p, "guestcidraddress")
	}
}

func (p *UpdateZoneParams) GetGuestcidraddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["guestcidraddress"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateZoneParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateZoneParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetInternaldns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internaldns1"] = v
}

func (p *UpdateZoneParams) ResetInternaldns1() {
	if p.p != nil && p.p["internaldns1"] != nil {
		delete(p.p, "internaldns1")
	}
}

func (p *UpdateZoneParams) GetInternaldns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["internaldns1"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetInternaldns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internaldns2"] = v
}

func (p *UpdateZoneParams) ResetInternaldns2() {
	if p.p != nil && p.p["internaldns2"] != nil {
		delete(p.p, "internaldns2")
	}
}

func (p *UpdateZoneParams) GetInternaldns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["internaldns2"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetIp6dns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns1"] = v
}

func (p *UpdateZoneParams) ResetIp6dns1() {
	if p.p != nil && p.p["ip6dns1"] != nil {
		delete(p.p, "ip6dns1")
	}
}

func (p *UpdateZoneParams) GetIp6dns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns1"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetIp6dns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns2"] = v
}

func (p *UpdateZoneParams) ResetIp6dns2() {
	if p.p != nil && p.p["ip6dns2"] != nil {
		delete(p.p, "ip6dns2")
	}
}

func (p *UpdateZoneParams) GetIp6dns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns2"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *UpdateZoneParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *UpdateZoneParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *UpdateZoneParams) SetLocalstorageenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["localstorageenabled"] = v
}

func (p *UpdateZoneParams) ResetLocalstorageenabled() {
	if p.p != nil && p.p["localstorageenabled"] != nil {
		delete(p.p, "localstorageenabled")
	}
}

func (p *UpdateZoneParams) GetLocalstorageenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["localstorageenabled"].(bool)
	return value, ok
}

func (p *UpdateZoneParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateZoneParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateZoneParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateZoneParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

func (p *UpdateZoneParams) ResetSortkey() {
	if p.p != nil && p.p["sortkey"] != nil {
		delete(p.p, "sortkey")
	}
}

func (p *UpdateZoneParams) GetSortkey() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortkey"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewUpdateZoneParams(id string) *UpdateZoneParams {
	p := &UpdateZoneParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a Zone.
func (s *ZoneService) UpdateZone(p *UpdateZoneParams) (*UpdateZoneResponse, error) {
	resp, err := s.cs.newRequest("updateZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateZoneResponse struct {
	Allocationstate              string                       `json:"allocationstate"`
	Allowuserspecifyvrmtu        bool                         `json:"allowuserspecifyvrmtu"`
	Asnrange                     string                       `json:"asnrange"`
	Capacity                     []UpdateZoneResponseCapacity `json:"capacity"`
	Description                  string                       `json:"description"`
	Dhcpprovider                 string                       `json:"dhcpprovider"`
	Displaytext                  string                       `json:"displaytext"`
	Dns1                         string                       `json:"dns1"`
	Dns2                         string                       `json:"dns2"`
	Domain                       string                       `json:"domain"`
	Domainid                     string                       `json:"domainid"`
	Domainname                   string                       `json:"domainname"`
	Guestcidraddress             string                       `json:"guestcidraddress"`
	Hasannotations               bool                         `json:"hasannotations"`
	Icon                         interface{}                  `json:"icon"`
	Id                           string                       `json:"id"`
	Internaldns1                 string                       `json:"internaldns1"`
	Internaldns2                 string                       `json:"internaldns2"`
	Ip6dns1                      string                       `json:"ip6dns1"`
	Ip6dns2                      string                       `json:"ip6dns2"`
	Ismultiarch                  bool                         `json:"ismultiarch"`
	Isnsxenabled                 bool                         `json:"isnsxenabled"`
	JobID                        string                       `json:"jobid"`
	Jobstatus                    int                          `json:"jobstatus"`
	Localstorageenabled          bool                         `json:"localstorageenabled"`
	Name                         string                       `json:"name"`
	Networktype                  string                       `json:"networktype"`
	Resourcedetails              map[string]string            `json:"resourcedetails"`
	Routerprivateinterfacemaxmtu int                          `json:"routerprivateinterfacemaxmtu"`
	Routerpublicinterfacemaxmtu  int                          `json:"routerpublicinterfacemaxmtu"`
	Securitygroupsenabled        bool                         `json:"securitygroupsenabled"`
	Tags                         []Tags                       `json:"tags"`
	Type                         string                       `json:"type"`
	Zonetoken                    string                       `json:"zonetoken"`
}

type UpdateZoneResponseCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Tag               string `json:"tag"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}
