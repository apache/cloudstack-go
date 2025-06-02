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

type VPCServiceIface interface {
	CreatePrivateGateway(p *CreatePrivateGatewayParams) (*CreatePrivateGatewayResponse, error)
	NewCreatePrivateGatewayParams(gateway string, ipaddress string, netmask string, vpcid string) *CreatePrivateGatewayParams
	CreateStaticRoute(p *CreateStaticRouteParams) (*CreateStaticRouteResponse, error)
	NewCreateStaticRouteParams(cidr string, gatewayid string) *CreateStaticRouteParams
	CreateVPC(p *CreateVPCParams) (*CreateVPCResponse, error)
	NewCreateVPCParams(displaytext string, name string, vpcofferingid string, zoneid string) *CreateVPCParams
	CreateVPCOffering(p *CreateVPCOfferingParams) (*CreateVPCOfferingResponse, error)
	NewCreateVPCOfferingParams(displaytext string, name string) *CreateVPCOfferingParams
	DeletePrivateGateway(p *DeletePrivateGatewayParams) (*DeletePrivateGatewayResponse, error)
	NewDeletePrivateGatewayParams(id string) *DeletePrivateGatewayParams
	DeleteStaticRoute(p *DeleteStaticRouteParams) (*DeleteStaticRouteResponse, error)
	NewDeleteStaticRouteParams(id string) *DeleteStaticRouteParams
	DeleteVPC(p *DeleteVPCParams) (*DeleteVPCResponse, error)
	NewDeleteVPCParams(id string) *DeleteVPCParams
	DeleteVPCOffering(p *DeleteVPCOfferingParams) (*DeleteVPCOfferingResponse, error)
	NewDeleteVPCOfferingParams(id string) *DeleteVPCOfferingParams
	ListPrivateGateways(p *ListPrivateGatewaysParams) (*ListPrivateGatewaysResponse, error)
	NewListPrivateGatewaysParams() *ListPrivateGatewaysParams
	GetPrivateGatewayByID(id string, opts ...OptionFunc) (*PrivateGateway, int, error)
	ListStaticRoutes(p *ListStaticRoutesParams) (*ListStaticRoutesResponse, error)
	NewListStaticRoutesParams() *ListStaticRoutesParams
	GetStaticRouteByID(id string, opts ...OptionFunc) (*StaticRoute, int, error)
	ListVPCOfferings(p *ListVPCOfferingsParams) (*ListVPCOfferingsResponse, error)
	NewListVPCOfferingsParams() *ListVPCOfferingsParams
	GetVPCOfferingID(name string, opts ...OptionFunc) (string, int, error)
	GetVPCOfferingByName(name string, opts ...OptionFunc) (*VPCOffering, int, error)
	GetVPCOfferingByID(id string, opts ...OptionFunc) (*VPCOffering, int, error)
	ListVPCs(p *ListVPCsParams) (*ListVPCsResponse, error)
	NewListVPCsParams() *ListVPCsParams
	GetVPCID(name string, opts ...OptionFunc) (string, int, error)
	GetVPCByName(name string, opts ...OptionFunc) (*VPC, int, error)
	GetVPCByID(id string, opts ...OptionFunc) (*VPC, int, error)
	MigrateVPC(p *MigrateVPCParams) (*MigrateVPCResponse, error)
	NewMigrateVPCParams(vpcid string, vpcofferingid string) *MigrateVPCParams
	RestartVPC(p *RestartVPCParams) (*RestartVPCResponse, error)
	NewRestartVPCParams(id string) *RestartVPCParams
	UpdateVPC(p *UpdateVPCParams) (*UpdateVPCResponse, error)
	NewUpdateVPCParams(id string) *UpdateVPCParams
	UpdateVPCOffering(p *UpdateVPCOfferingParams) (*UpdateVPCOfferingResponse, error)
	NewUpdateVPCOfferingParams(id string) *UpdateVPCOfferingParams
}

type CreatePrivateGatewayParams struct {
	p map[string]interface{}
}

func (p *CreatePrivateGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["associatednetworkid"]; found {
		u.Set("associatednetworkid", v.(string))
	}
	if v, found := p.p["bypassvlanoverlapcheck"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bypassvlanoverlapcheck", vv)
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["sourcenatsupported"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sourcenatsupported", vv)
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *CreatePrivateGatewayParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
}

func (p *CreatePrivateGatewayParams) ResetAclid() {
	if p.p != nil && p.p["aclid"] != nil {
		delete(p.p, "aclid")
	}
}

func (p *CreatePrivateGatewayParams) GetAclid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["aclid"].(string)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetAssociatednetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["associatednetworkid"] = v
}

func (p *CreatePrivateGatewayParams) ResetAssociatednetworkid() {
	if p.p != nil && p.p["associatednetworkid"] != nil {
		delete(p.p, "associatednetworkid")
	}
}

func (p *CreatePrivateGatewayParams) GetAssociatednetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["associatednetworkid"].(string)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetBypassvlanoverlapcheck(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bypassvlanoverlapcheck"] = v
}

func (p *CreatePrivateGatewayParams) ResetBypassvlanoverlapcheck() {
	if p.p != nil && p.p["bypassvlanoverlapcheck"] != nil {
		delete(p.p, "bypassvlanoverlapcheck")
	}
}

func (p *CreatePrivateGatewayParams) GetBypassvlanoverlapcheck() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bypassvlanoverlapcheck"].(bool)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
}

func (p *CreatePrivateGatewayParams) ResetGateway() {
	if p.p != nil && p.p["gateway"] != nil {
		delete(p.p, "gateway")
	}
}

func (p *CreatePrivateGatewayParams) GetGateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gateway"].(string)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *CreatePrivateGatewayParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *CreatePrivateGatewayParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
}

func (p *CreatePrivateGatewayParams) ResetNetmask() {
	if p.p != nil && p.p["netmask"] != nil {
		delete(p.p, "netmask")
	}
}

func (p *CreatePrivateGatewayParams) GetNetmask() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["netmask"].(string)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
}

func (p *CreatePrivateGatewayParams) ResetNetworkofferingid() {
	if p.p != nil && p.p["networkofferingid"] != nil {
		delete(p.p, "networkofferingid")
	}
}

func (p *CreatePrivateGatewayParams) GetNetworkofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkofferingid"].(string)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *CreatePrivateGatewayParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *CreatePrivateGatewayParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetSourcenatsupported(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatsupported"] = v
}

func (p *CreatePrivateGatewayParams) ResetSourcenatsupported() {
	if p.p != nil && p.p["sourcenatsupported"] != nil {
		delete(p.p, "sourcenatsupported")
	}
}

func (p *CreatePrivateGatewayParams) GetSourcenatsupported() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcenatsupported"].(bool)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *CreatePrivateGatewayParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
}

func (p *CreatePrivateGatewayParams) GetVlan() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(string)
	return value, ok
}

func (p *CreatePrivateGatewayParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *CreatePrivateGatewayParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *CreatePrivateGatewayParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new CreatePrivateGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreatePrivateGatewayParams(gateway string, ipaddress string, netmask string, vpcid string) *CreatePrivateGatewayParams {
	p := &CreatePrivateGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["gateway"] = gateway
	p.p["ipaddress"] = ipaddress
	p.p["netmask"] = netmask
	p.p["vpcid"] = vpcid
	return p
}

// Creates a private gateway
func (s *VPCService) CreatePrivateGateway(p *CreatePrivateGatewayParams) (*CreatePrivateGatewayResponse, error) {
	resp, err := s.cs.newPostRequest("createPrivateGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePrivateGatewayResponse
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

type CreatePrivateGatewayResponse struct {
	Account             string `json:"account"`
	Aclid               string `json:"aclid"`
	Aclname             string `json:"aclname"`
	Associatednetwork   string `json:"associatednetwork"`
	Associatednetworkid string `json:"associatednetworkid"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Domainpath          string `json:"domainpath"`
	Gateway             string `json:"gateway"`
	Hasannotations      bool   `json:"hasannotations"`
	Id                  string `json:"id"`
	Ipaddress           string `json:"ipaddress"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Netmask             string `json:"netmask"`
	Physicalnetworkid   string `json:"physicalnetworkid"`
	Project             string `json:"project"`
	Projectid           string `json:"projectid"`
	Sourcenatsupported  bool   `json:"sourcenatsupported"`
	State               string `json:"state"`
	Vlan                string `json:"vlan"`
	Vpcid               string `json:"vpcid"`
	Vpcname             string `json:"vpcname"`
	Zoneid              string `json:"zoneid"`
	Zonename            string `json:"zonename"`
}

type CreateStaticRouteParams struct {
	p map[string]interface{}
}

func (p *CreateStaticRouteParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
	}
	return u
}

func (p *CreateStaticRouteParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
}

func (p *CreateStaticRouteParams) ResetCidr() {
	if p.p != nil && p.p["cidr"] != nil {
		delete(p.p, "cidr")
	}
}

func (p *CreateStaticRouteParams) GetCidr() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidr"].(string)
	return value, ok
}

func (p *CreateStaticRouteParams) SetGatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gatewayid"] = v
}

func (p *CreateStaticRouteParams) ResetGatewayid() {
	if p.p != nil && p.p["gatewayid"] != nil {
		delete(p.p, "gatewayid")
	}
}

func (p *CreateStaticRouteParams) GetGatewayid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gatewayid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateStaticRouteParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateStaticRouteParams(cidr string, gatewayid string) *CreateStaticRouteParams {
	p := &CreateStaticRouteParams{}
	p.p = make(map[string]interface{})
	p.p["cidr"] = cidr
	p.p["gatewayid"] = gatewayid
	return p
}

// Creates a static route
func (s *VPCService) CreateStaticRoute(p *CreateStaticRouteParams) (*CreateStaticRouteResponse, error) {
	resp, err := s.cs.newPostRequest("createStaticRoute", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateStaticRouteResponse
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

type CreateStaticRouteResponse struct {
	Account    string `json:"account"`
	Cidr       string `json:"cidr"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Gatewayid  string `json:"gatewayid"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	State      string `json:"state"`
	Tags       []Tags `json:"tags"`
	Vpcid      string `json:"vpcid"`
}

type CreateVPCParams struct {
	p map[string]interface{}
}

func (p *CreateVPCParams) toURLValues() url.Values {
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
	if v, found := p.p["bgppeerids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("bgppeerids", vv)
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["cidrsize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cidrsize", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["dns1"]; found {
		u.Set("dns1", v.(string))
	}
	if v, found := p.p["dns2"]; found {
		u.Set("dns2", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["ip6dns1"]; found {
		u.Set("ip6dns1", v.(string))
	}
	if v, found := p.p["ip6dns2"]; found {
		u.Set("ip6dns2", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["publicmtu"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicmtu", vv)
	}
	if v, found := p.p["sourcenatipaddress"]; found {
		u.Set("sourcenatipaddress", v.(string))
	}
	if v, found := p.p["start"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("start", vv)
	}
	if v, found := p.p["vpcofferingid"]; found {
		u.Set("vpcofferingid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateVPCParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateVPCParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateVPCParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetAsnumber(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asnumber"] = v
}

func (p *CreateVPCParams) ResetAsnumber() {
	if p.p != nil && p.p["asnumber"] != nil {
		delete(p.p, "asnumber")
	}
}

func (p *CreateVPCParams) GetAsnumber() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asnumber"].(int64)
	return value, ok
}

func (p *CreateVPCParams) SetBgppeerids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bgppeerids"] = v
}

func (p *CreateVPCParams) ResetBgppeerids() {
	if p.p != nil && p.p["bgppeerids"] != nil {
		delete(p.p, "bgppeerids")
	}
}

func (p *CreateVPCParams) GetBgppeerids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bgppeerids"].([]string)
	return value, ok
}

func (p *CreateVPCParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
}

func (p *CreateVPCParams) ResetCidr() {
	if p.p != nil && p.p["cidr"] != nil {
		delete(p.p, "cidr")
	}
}

func (p *CreateVPCParams) GetCidr() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidr"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetCidrsize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrsize"] = v
}

func (p *CreateVPCParams) ResetCidrsize() {
	if p.p != nil && p.p["cidrsize"] != nil {
		delete(p.p, "cidrsize")
	}
}

func (p *CreateVPCParams) GetCidrsize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrsize"].(int)
	return value, ok
}

func (p *CreateVPCParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateVPCParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *CreateVPCParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetDns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns1"] = v
}

func (p *CreateVPCParams) ResetDns1() {
	if p.p != nil && p.p["dns1"] != nil {
		delete(p.p, "dns1")
	}
}

func (p *CreateVPCParams) GetDns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns1"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetDns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns2"] = v
}

func (p *CreateVPCParams) ResetDns2() {
	if p.p != nil && p.p["dns2"] != nil {
		delete(p.p, "dns2")
	}
}

func (p *CreateVPCParams) GetDns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns2"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateVPCParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateVPCParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateVPCParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateVPCParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateVPCParams) SetIp6dns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns1"] = v
}

func (p *CreateVPCParams) ResetIp6dns1() {
	if p.p != nil && p.p["ip6dns1"] != nil {
		delete(p.p, "ip6dns1")
	}
}

func (p *CreateVPCParams) GetIp6dns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns1"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetIp6dns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns2"] = v
}

func (p *CreateVPCParams) ResetIp6dns2() {
	if p.p != nil && p.p["ip6dns2"] != nil {
		delete(p.p, "ip6dns2")
	}
}

func (p *CreateVPCParams) GetIp6dns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns2"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVPCParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateVPCParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *CreateVPCParams) ResetNetworkdomain() {
	if p.p != nil && p.p["networkdomain"] != nil {
		delete(p.p, "networkdomain")
	}
}

func (p *CreateVPCParams) GetNetworkdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdomain"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateVPCParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateVPCParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetPublicmtu(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicmtu"] = v
}

func (p *CreateVPCParams) ResetPublicmtu() {
	if p.p != nil && p.p["publicmtu"] != nil {
		delete(p.p, "publicmtu")
	}
}

func (p *CreateVPCParams) GetPublicmtu() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicmtu"].(int)
	return value, ok
}

func (p *CreateVPCParams) SetSourcenatipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatipaddress"] = v
}

func (p *CreateVPCParams) ResetSourcenatipaddress() {
	if p.p != nil && p.p["sourcenatipaddress"] != nil {
		delete(p.p, "sourcenatipaddress")
	}
}

func (p *CreateVPCParams) GetSourcenatipaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcenatipaddress"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetStart(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["start"] = v
}

func (p *CreateVPCParams) ResetStart() {
	if p.p != nil && p.p["start"] != nil {
		delete(p.p, "start")
	}
}

func (p *CreateVPCParams) GetStart() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["start"].(bool)
	return value, ok
}

func (p *CreateVPCParams) SetVpcofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcofferingid"] = v
}

func (p *CreateVPCParams) ResetVpcofferingid() {
	if p.p != nil && p.p["vpcofferingid"] != nil {
		delete(p.p, "vpcofferingid")
	}
}

func (p *CreateVPCParams) GetVpcofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcofferingid"].(string)
	return value, ok
}

func (p *CreateVPCParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateVPCParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateVPCParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateVPCParams(displaytext string, name string, vpcofferingid string, zoneid string) *CreateVPCParams {
	p := &CreateVPCParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["vpcofferingid"] = vpcofferingid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a VPC
func (s *VPCService) CreateVPC(p *CreateVPCParams) (*CreateVPCResponse, error) {
	resp, err := s.cs.newPostRequest("createVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVPCResponse
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

type CreateVPCResponse struct {
	Account              string                     `json:"account"`
	Asnumber             int64                      `json:"asnumber"`
	Asnumberid           string                     `json:"asnumberid"`
	Bgppeers             []interface{}              `json:"bgppeers"`
	Cidr                 string                     `json:"cidr"`
	Created              string                     `json:"created"`
	Displaytext          string                     `json:"displaytext"`
	Distributedvpcrouter bool                       `json:"distributedvpcrouter"`
	Dns1                 string                     `json:"dns1"`
	Dns2                 string                     `json:"dns2"`
	Domain               string                     `json:"domain"`
	Domainid             string                     `json:"domainid"`
	Domainpath           string                     `json:"domainpath"`
	Fordisplay           bool                       `json:"fordisplay"`
	Hasannotations       bool                       `json:"hasannotations"`
	Icon                 interface{}                `json:"icon"`
	Id                   string                     `json:"id"`
	Ip4routes            []interface{}              `json:"ip4routes"`
	Ip4routing           string                     `json:"ip4routing"`
	Ip6dns1              string                     `json:"ip6dns1"`
	Ip6dns2              string                     `json:"ip6dns2"`
	Ip6routes            []interface{}              `json:"ip6routes"`
	JobID                string                     `json:"jobid"`
	Jobstatus            int                        `json:"jobstatus"`
	Name                 string                     `json:"name"`
	Network              []*Network                 `json:"network"`
	Networkdomain        string                     `json:"networkdomain"`
	Project              string                     `json:"project"`
	Projectid            string                     `json:"projectid"`
	Publicmtu            int                        `json:"publicmtu"`
	Redundantvpcrouter   bool                       `json:"redundantvpcrouter"`
	Regionlevelvpc       bool                       `json:"regionlevelvpc"`
	Restartrequired      bool                       `json:"restartrequired"`
	Service              []CreateVPCResponseService `json:"service"`
	State                string                     `json:"state"`
	Tags                 []Tags                     `json:"tags"`
	Vpcofferingid        string                     `json:"vpcofferingid"`
	Vpcofferingname      string                     `json:"vpcofferingname"`
	Zoneid               string                     `json:"zoneid"`
	Zonename             string                     `json:"zonename"`
}

type CreateVPCResponseService struct {
	Capability []CreateVPCResponseServiceCapability `json:"capability"`
	Name       string                               `json:"name"`
	Provider   []CreateVPCResponseServiceProvider   `json:"provider"`
}

type CreateVPCResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type CreateVPCResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type CreateVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("domainid", vv)
	}
	if v, found := p.p["enable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enable", vv)
	}
	if v, found := p.p["fornsx"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fornsx", vv)
	}
	if v, found := p.p["internetprotocol"]; found {
		u.Set("internetprotocol", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkmode"]; found {
		u.Set("networkmode", v.(string))
	}
	if v, found := p.p["nsxsupportlb"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("nsxsupportlb", vv)
	}
	if v, found := p.p["routingmode"]; found {
		u.Set("routingmode", v.(string))
	}
	if v, found := p.p["servicecapabilitylist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].key", i), k)
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].value", i), m[k])
		}
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["serviceproviderlist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("serviceproviderlist[%d].service", i), k)
			u.Set(fmt.Sprintf("serviceproviderlist[%d].provider", i), m[k])
		}
	}
	if v, found := p.p["specifyasnumber"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyasnumber", vv)
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["zoneid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneid", vv)
	}
	return u
}

func (p *CreateVPCOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateVPCOfferingParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *CreateVPCOfferingParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetDomainid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateVPCOfferingParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateVPCOfferingParams) GetDomainid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].([]string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetEnable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enable"] = v
}

func (p *CreateVPCOfferingParams) ResetEnable() {
	if p.p != nil && p.p["enable"] != nil {
		delete(p.p, "enable")
	}
}

func (p *CreateVPCOfferingParams) GetEnable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enable"].(bool)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetFornsx(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fornsx"] = v
}

func (p *CreateVPCOfferingParams) ResetFornsx() {
	if p.p != nil && p.p["fornsx"] != nil {
		delete(p.p, "fornsx")
	}
}

func (p *CreateVPCOfferingParams) GetFornsx() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fornsx"].(bool)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetInternetprotocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internetprotocol"] = v
}

func (p *CreateVPCOfferingParams) ResetInternetprotocol() {
	if p.p != nil && p.p["internetprotocol"] != nil {
		delete(p.p, "internetprotocol")
	}
}

func (p *CreateVPCOfferingParams) GetInternetprotocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["internetprotocol"].(string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVPCOfferingParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateVPCOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetNetworkmode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkmode"] = v
}

func (p *CreateVPCOfferingParams) ResetNetworkmode() {
	if p.p != nil && p.p["networkmode"] != nil {
		delete(p.p, "networkmode")
	}
}

func (p *CreateVPCOfferingParams) GetNetworkmode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkmode"].(string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetNsxsupportlb(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nsxsupportlb"] = v
}

func (p *CreateVPCOfferingParams) ResetNsxsupportlb() {
	if p.p != nil && p.p["nsxsupportlb"] != nil {
		delete(p.p, "nsxsupportlb")
	}
}

func (p *CreateVPCOfferingParams) GetNsxsupportlb() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nsxsupportlb"].(bool)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetRoutingmode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["routingmode"] = v
}

func (p *CreateVPCOfferingParams) ResetRoutingmode() {
	if p.p != nil && p.p["routingmode"] != nil {
		delete(p.p, "routingmode")
	}
}

func (p *CreateVPCOfferingParams) GetRoutingmode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["routingmode"].(string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetServicecapabilitylist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicecapabilitylist"] = v
}

func (p *CreateVPCOfferingParams) ResetServicecapabilitylist() {
	if p.p != nil && p.p["servicecapabilitylist"] != nil {
		delete(p.p, "servicecapabilitylist")
	}
}

func (p *CreateVPCOfferingParams) GetServicecapabilitylist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["servicecapabilitylist"].(map[string]string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateVPCOfferingParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *CreateVPCOfferingParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetServiceproviderlist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceproviderlist"] = v
}

func (p *CreateVPCOfferingParams) ResetServiceproviderlist() {
	if p.p != nil && p.p["serviceproviderlist"] != nil {
		delete(p.p, "serviceproviderlist")
	}
}

func (p *CreateVPCOfferingParams) GetServiceproviderlist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceproviderlist"].(map[string]string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetSpecifyasnumber(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyasnumber"] = v
}

func (p *CreateVPCOfferingParams) ResetSpecifyasnumber() {
	if p.p != nil && p.p["specifyasnumber"] != nil {
		delete(p.p, "specifyasnumber")
	}
}

func (p *CreateVPCOfferingParams) GetSpecifyasnumber() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["specifyasnumber"].(bool)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

func (p *CreateVPCOfferingParams) ResetSupportedservices() {
	if p.p != nil && p.p["supportedservices"] != nil {
		delete(p.p, "supportedservices")
	}
}

func (p *CreateVPCOfferingParams) GetSupportedservices() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["supportedservices"].([]string)
	return value, ok
}

func (p *CreateVPCOfferingParams) SetZoneid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateVPCOfferingParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateVPCOfferingParams) GetZoneid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewCreateVPCOfferingParams(displaytext string, name string) *CreateVPCOfferingParams {
	p := &CreateVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	return p
}

// Creates VPC offering
func (s *VPCService) CreateVPCOffering(p *CreateVPCOfferingParams) (*CreateVPCOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("createVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVPCOfferingResponse
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

type CreateVPCOfferingResponse struct {
	Created                string                             `json:"created"`
	Displaytext            string                             `json:"displaytext"`
	Distributedvpcrouter   bool                               `json:"distributedvpcrouter"`
	Domain                 string                             `json:"domain"`
	Domainid               string                             `json:"domainid"`
	Fornsx                 bool                               `json:"fornsx"`
	Id                     string                             `json:"id"`
	Internetprotocol       string                             `json:"internetprotocol"`
	Isdefault              bool                               `json:"isdefault"`
	JobID                  string                             `json:"jobid"`
	Jobstatus              int                                `json:"jobstatus"`
	Name                   string                             `json:"name"`
	Networkmode            string                             `json:"networkmode"`
	Routingmode            string                             `json:"routingmode"`
	Service                []CreateVPCOfferingResponseService `json:"service"`
	Specifyasnumber        bool                               `json:"specifyasnumber"`
	State                  string                             `json:"state"`
	SupportsregionLevelvpc bool                               `json:"supportsregionLevelvpc"`
	Zone                   string                             `json:"zone"`
	Zoneid                 string                             `json:"zoneid"`
}

type CreateVPCOfferingResponseService struct {
	Capability []CreateVPCOfferingResponseServiceCapability `json:"capability"`
	Name       string                                       `json:"name"`
	Provider   []CreateVPCOfferingResponseServiceProvider   `json:"provider"`
}

type CreateVPCOfferingResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type CreateVPCOfferingResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type DeletePrivateGatewayParams struct {
	p map[string]interface{}
}

func (p *DeletePrivateGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePrivateGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeletePrivateGatewayParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeletePrivateGatewayParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePrivateGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeletePrivateGatewayParams(id string) *DeletePrivateGatewayParams {
	p := &DeletePrivateGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Private gateway
func (s *VPCService) DeletePrivateGateway(p *DeletePrivateGatewayParams) (*DeletePrivateGatewayResponse, error) {
	resp, err := s.cs.newPostRequest("deletePrivateGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePrivateGatewayResponse
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

type DeletePrivateGatewayResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteStaticRouteParams struct {
	p map[string]interface{}
}

func (p *DeleteStaticRouteParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteStaticRouteParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteStaticRouteParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteStaticRouteParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteStaticRouteParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteStaticRouteParams(id string) *DeleteStaticRouteParams {
	p := &DeleteStaticRouteParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a static route
func (s *VPCService) DeleteStaticRoute(p *DeleteStaticRouteParams) (*DeleteStaticRouteResponse, error) {
	resp, err := s.cs.newPostRequest("deleteStaticRoute", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStaticRouteResponse
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

type DeleteStaticRouteResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteVPCParams struct {
	p map[string]interface{}
}

func (p *DeleteVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVPCParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteVPCParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteVPCParams(id string) *DeleteVPCParams {
	p := &DeleteVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a VPC
func (s *VPCService) DeleteVPC(p *DeleteVPCParams) (*DeleteVPCResponse, error) {
	resp, err := s.cs.newPostRequest("deleteVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVPCResponse
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

type DeleteVPCResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVPCOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVPCOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteVPCOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewDeleteVPCOfferingParams(id string) *DeleteVPCOfferingParams {
	p := &DeleteVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes VPC offering
func (s *VPCService) DeleteVPCOffering(p *DeleteVPCOfferingParams) (*DeleteVPCOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("deleteVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVPCOfferingResponse
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

type DeleteVPCOfferingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListPrivateGatewaysParams struct {
	p map[string]interface{}
}

func (p *ListPrivateGatewaysParams) toURLValues() url.Values {
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
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListPrivateGatewaysParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListPrivateGatewaysParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListPrivateGatewaysParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListPrivateGatewaysParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListPrivateGatewaysParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListPrivateGatewaysParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListPrivateGatewaysParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *ListPrivateGatewaysParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *ListPrivateGatewaysParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListPrivateGatewaysParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListPrivateGatewaysParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListPrivateGatewaysParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListPrivateGatewaysParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListPrivateGatewaysParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListPrivateGatewaysParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListPrivateGatewaysParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListPrivateGatewaysParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListPrivateGatewaysParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListPrivateGatewaysParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListPrivateGatewaysParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListPrivateGatewaysParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListPrivateGatewaysParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListPrivateGatewaysParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *ListPrivateGatewaysParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
}

func (p *ListPrivateGatewaysParams) GetVlan() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(string)
	return value, ok
}

func (p *ListPrivateGatewaysParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListPrivateGatewaysParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListPrivateGatewaysParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPrivateGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListPrivateGatewaysParams() *ListPrivateGatewaysParams {
	p := &ListPrivateGatewaysParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetPrivateGatewayByID(id string, opts ...OptionFunc) (*PrivateGateway, int, error) {
	p := &ListPrivateGatewaysParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPrivateGateways(p)
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
		return l.PrivateGateways[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for PrivateGateway UUID: %s!", id)
}

// List private gateways
func (s *VPCService) ListPrivateGateways(p *ListPrivateGatewaysParams) (*ListPrivateGatewaysResponse, error) {
	resp, err := s.cs.newRequest("listPrivateGateways", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPrivateGatewaysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPrivateGatewaysResponse struct {
	Count           int               `json:"count"`
	PrivateGateways []*PrivateGateway `json:"privategateway"`
}

type PrivateGateway struct {
	Account             string `json:"account"`
	Aclid               string `json:"aclid"`
	Aclname             string `json:"aclname"`
	Associatednetwork   string `json:"associatednetwork"`
	Associatednetworkid string `json:"associatednetworkid"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Domainpath          string `json:"domainpath"`
	Gateway             string `json:"gateway"`
	Hasannotations      bool   `json:"hasannotations"`
	Id                  string `json:"id"`
	Ipaddress           string `json:"ipaddress"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Netmask             string `json:"netmask"`
	Physicalnetworkid   string `json:"physicalnetworkid"`
	Project             string `json:"project"`
	Projectid           string `json:"projectid"`
	Sourcenatsupported  bool   `json:"sourcenatsupported"`
	State               string `json:"state"`
	Vlan                string `json:"vlan"`
	Vpcid               string `json:"vpcid"`
	Vpcname             string `json:"vpcname"`
	Zoneid              string `json:"zoneid"`
	Zonename            string `json:"zonename"`
}

type ListStaticRoutesParams struct {
	p map[string]interface{}
}

func (p *ListStaticRoutesParams) toURLValues() url.Values {
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
	if v, found := p.p["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
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
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListStaticRoutesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListStaticRoutesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListStaticRoutesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListStaticRoutesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListStaticRoutesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListStaticRoutesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListStaticRoutesParams) SetGatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gatewayid"] = v
}

func (p *ListStaticRoutesParams) ResetGatewayid() {
	if p.p != nil && p.p["gatewayid"] != nil {
		delete(p.p, "gatewayid")
	}
}

func (p *ListStaticRoutesParams) GetGatewayid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gatewayid"].(string)
	return value, ok
}

func (p *ListStaticRoutesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListStaticRoutesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListStaticRoutesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListStaticRoutesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListStaticRoutesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListStaticRoutesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListStaticRoutesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStaticRoutesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListStaticRoutesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStaticRoutesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListStaticRoutesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListStaticRoutesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListStaticRoutesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStaticRoutesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListStaticRoutesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStaticRoutesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStaticRoutesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListStaticRoutesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListStaticRoutesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListStaticRoutesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListStaticRoutesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListStaticRoutesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListStaticRoutesParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListStaticRoutesParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListStaticRoutesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListStaticRoutesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListStaticRoutesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListStaticRoutesParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListStaticRoutesParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListStaticRoutesParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ListStaticRoutesParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListStaticRoutesParams() *ListStaticRoutesParams {
	p := &ListStaticRoutesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetStaticRouteByID(id string, opts ...OptionFunc) (*StaticRoute, int, error) {
	p := &ListStaticRoutesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStaticRoutes(p)
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
		return l.StaticRoutes[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for StaticRoute UUID: %s!", id)
}

// Lists all static routes
func (s *VPCService) ListStaticRoutes(p *ListStaticRoutesParams) (*ListStaticRoutesResponse, error) {
	resp, err := s.cs.newRequest("listStaticRoutes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStaticRoutesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStaticRoutesResponse struct {
	Count        int            `json:"count"`
	StaticRoutes []*StaticRoute `json:"staticroute"`
}

type StaticRoute struct {
	Account    string `json:"account"`
	Cidr       string `json:"cidr"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Gatewayid  string `json:"gatewayid"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	State      string `json:"state"`
	Tags       []Tags `json:"tags"`
	Vpcid      string `json:"vpcid"`
}

type ListVPCOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListVPCOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdefault"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdefault", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVPCOfferingsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *ListVPCOfferingsParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *ListVPCOfferingsParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVPCOfferingsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVPCOfferingsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVPCOfferingsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVPCOfferingsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetIsdefault(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdefault"] = v
}

func (p *ListVPCOfferingsParams) ResetIsdefault() {
	if p.p != nil && p.p["isdefault"] != nil {
		delete(p.p, "isdefault")
	}
}

func (p *ListVPCOfferingsParams) GetIsdefault() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdefault"].(bool)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVPCOfferingsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVPCOfferingsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVPCOfferingsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVPCOfferingsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVPCOfferingsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVPCOfferingsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVPCOfferingsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVPCOfferingsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVPCOfferingsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListVPCOfferingsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

func (p *ListVPCOfferingsParams) ResetSupportedservices() {
	if p.p != nil && p.p["supportedservices"] != nil {
		delete(p.p, "supportedservices")
	}
}

func (p *ListVPCOfferingsParams) GetSupportedservices() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["supportedservices"].([]string)
	return value, ok
}

func (p *ListVPCOfferingsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVPCOfferingsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVPCOfferingsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVPCOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListVPCOfferingsParams() *ListVPCOfferingsParams {
	p := &ListVPCOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVPCOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVPCOfferings(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VPCOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VPCOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCOfferingByName(name string, opts ...OptionFunc) (*VPCOffering, int, error) {
	id, count, err := s.GetVPCOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVPCOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCOfferingByID(id string, opts ...OptionFunc) (*VPCOffering, int, error) {
	p := &ListVPCOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVPCOfferings(p)
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
		return l.VPCOfferings[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VPCOffering UUID: %s!", id)
}

// Lists VPC offerings
func (s *VPCService) ListVPCOfferings(p *ListVPCOfferingsParams) (*ListVPCOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listVPCOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVPCOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVPCOfferingsResponse struct {
	Count        int            `json:"count"`
	VPCOfferings []*VPCOffering `json:"vpcoffering"`
}

type VPCOffering struct {
	Created                string               `json:"created"`
	Displaytext            string               `json:"displaytext"`
	Distributedvpcrouter   bool                 `json:"distributedvpcrouter"`
	Domain                 string               `json:"domain"`
	Domainid               string               `json:"domainid"`
	Fornsx                 bool                 `json:"fornsx"`
	Id                     string               `json:"id"`
	Internetprotocol       string               `json:"internetprotocol"`
	Isdefault              bool                 `json:"isdefault"`
	JobID                  string               `json:"jobid"`
	Jobstatus              int                  `json:"jobstatus"`
	Name                   string               `json:"name"`
	Networkmode            string               `json:"networkmode"`
	Routingmode            string               `json:"routingmode"`
	Service                []VPCOfferingService `json:"service"`
	Specifyasnumber        bool                 `json:"specifyasnumber"`
	State                  string               `json:"state"`
	SupportsregionLevelvpc bool                 `json:"supportsregionLevelvpc"`
	Zone                   string               `json:"zone"`
	Zoneid                 string               `json:"zoneid"`
}

type VPCOfferingService struct {
	Capability []VPCOfferingServiceCapability `json:"capability"`
	Name       string                         `json:"name"`
	Provider   []VPCOfferingServiceProvider   `json:"provider"`
}

type VPCOfferingServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type VPCOfferingServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListVPCsParams struct {
	p map[string]interface{}
}

func (p *ListVPCsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["restartrequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("restartrequired", vv)
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["vpcofferingid"]; found {
		u.Set("vpcofferingid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVPCsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVPCsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVPCsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
}

func (p *ListVPCsParams) ResetCidr() {
	if p.p != nil && p.p["cidr"] != nil {
		delete(p.p, "cidr")
	}
}

func (p *ListVPCsParams) GetCidr() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidr"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *ListVPCsParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *ListVPCsParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVPCsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVPCsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListVPCsParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListVPCsParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListVPCsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVPCsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVPCsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVPCsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVPCsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVPCsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVPCsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVPCsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVPCsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVPCsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVPCsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVPCsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVPCsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVPCsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVPCsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVPCsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVPCsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVPCsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVPCsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVPCsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVPCsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetRestartrequired(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["restartrequired"] = v
}

func (p *ListVPCsParams) ResetRestartrequired() {
	if p.p != nil && p.p["restartrequired"] != nil {
		delete(p.p, "restartrequired")
	}
}

func (p *ListVPCsParams) GetRestartrequired() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["restartrequired"].(bool)
	return value, ok
}

func (p *ListVPCsParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListVPCsParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListVPCsParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListVPCsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVPCsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListVPCsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

func (p *ListVPCsParams) ResetSupportedservices() {
	if p.p != nil && p.p["supportedservices"] != nil {
		delete(p.p, "supportedservices")
	}
}

func (p *ListVPCsParams) GetSupportedservices() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["supportedservices"].([]string)
	return value, ok
}

func (p *ListVPCsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVPCsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListVPCsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListVPCsParams) SetVpcofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcofferingid"] = v
}

func (p *ListVPCsParams) ResetVpcofferingid() {
	if p.p != nil && p.p["vpcofferingid"] != nil {
		delete(p.p, "vpcofferingid")
	}
}

func (p *ListVPCsParams) GetVpcofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcofferingid"].(string)
	return value, ok
}

func (p *ListVPCsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVPCsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVPCsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVPCsParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewListVPCsParams() *ListVPCsParams {
	p := &ListVPCsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVPCsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVPCs(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VPCs[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VPCs {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCByName(name string, opts ...OptionFunc) (*VPC, int, error) {
	id, count, err := s.GetVPCID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVPCByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPCService) GetVPCByID(id string, opts ...OptionFunc) (*VPC, int, error) {
	p := &ListVPCsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVPCs(p)
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
		return l.VPCs[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VPC UUID: %s!", id)
}

// Lists VPCs
func (s *VPCService) ListVPCs(p *ListVPCsParams) (*ListVPCsResponse, error) {
	resp, err := s.cs.newRequest("listVPCs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVPCsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVPCsResponse struct {
	Count int    `json:"count"`
	VPCs  []*VPC `json:"vpc"`
}

type VPC struct {
	Account              string               `json:"account"`
	Asnumber             int64                `json:"asnumber"`
	Asnumberid           string               `json:"asnumberid"`
	Bgppeers             []interface{}        `json:"bgppeers"`
	Cidr                 string               `json:"cidr"`
	Created              string               `json:"created"`
	Displaytext          string               `json:"displaytext"`
	Distributedvpcrouter bool                 `json:"distributedvpcrouter"`
	Dns1                 string               `json:"dns1"`
	Dns2                 string               `json:"dns2"`
	Domain               string               `json:"domain"`
	Domainid             string               `json:"domainid"`
	Domainpath           string               `json:"domainpath"`
	Fordisplay           bool                 `json:"fordisplay"`
	Hasannotations       bool                 `json:"hasannotations"`
	Icon                 interface{}          `json:"icon"`
	Id                   string               `json:"id"`
	Ip4routes            []interface{}        `json:"ip4routes"`
	Ip4routing           string               `json:"ip4routing"`
	Ip6dns1              string               `json:"ip6dns1"`
	Ip6dns2              string               `json:"ip6dns2"`
	Ip6routes            []interface{}        `json:"ip6routes"`
	JobID                string               `json:"jobid"`
	Jobstatus            int                  `json:"jobstatus"`
	Name                 string               `json:"name"`
	Network              []*Network           `json:"network"`
	Networkdomain        string               `json:"networkdomain"`
	Project              string               `json:"project"`
	Projectid            string               `json:"projectid"`
	Publicmtu            int                  `json:"publicmtu"`
	Redundantvpcrouter   bool                 `json:"redundantvpcrouter"`
	Regionlevelvpc       bool                 `json:"regionlevelvpc"`
	Restartrequired      bool                 `json:"restartrequired"`
	Service              []VPCServiceInternal `json:"service"`
	State                string               `json:"state"`
	Tags                 []Tags               `json:"tags"`
	Vpcofferingid        string               `json:"vpcofferingid"`
	Vpcofferingname      string               `json:"vpcofferingname"`
	Zoneid               string               `json:"zoneid"`
	Zonename             string               `json:"zonename"`
}

type VPCServiceInternal struct {
	Capability []VPCServiceInternalCapability `json:"capability"`
	Name       string                         `json:"name"`
	Provider   []VPCServiceInternalProvider   `json:"provider"`
}

type VPCServiceInternalProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type VPCServiceInternalCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type MigrateVPCParams struct {
	p map[string]interface{}
}

func (p *MigrateVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["resume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("resume", vv)
	}
	if v, found := p.p["tiernetworkofferings"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tiernetworkofferings[%d].key", i), k)
			u.Set(fmt.Sprintf("tiernetworkofferings[%d].value", i), m[k])
		}
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["vpcofferingid"]; found {
		u.Set("vpcofferingid", v.(string))
	}
	return u
}

func (p *MigrateVPCParams) SetResume(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resume"] = v
}

func (p *MigrateVPCParams) ResetResume() {
	if p.p != nil && p.p["resume"] != nil {
		delete(p.p, "resume")
	}
}

func (p *MigrateVPCParams) GetResume() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resume"].(bool)
	return value, ok
}

func (p *MigrateVPCParams) SetTiernetworkofferings(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tiernetworkofferings"] = v
}

func (p *MigrateVPCParams) ResetTiernetworkofferings() {
	if p.p != nil && p.p["tiernetworkofferings"] != nil {
		delete(p.p, "tiernetworkofferings")
	}
}

func (p *MigrateVPCParams) GetTiernetworkofferings() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tiernetworkofferings"].(map[string]string)
	return value, ok
}

func (p *MigrateVPCParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *MigrateVPCParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *MigrateVPCParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *MigrateVPCParams) SetVpcofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcofferingid"] = v
}

func (p *MigrateVPCParams) ResetVpcofferingid() {
	if p.p != nil && p.p["vpcofferingid"] != nil {
		delete(p.p, "vpcofferingid")
	}
}

func (p *MigrateVPCParams) GetVpcofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewMigrateVPCParams(vpcid string, vpcofferingid string) *MigrateVPCParams {
	p := &MigrateVPCParams{}
	p.p = make(map[string]interface{})
	p.p["vpcid"] = vpcid
	p.p["vpcofferingid"] = vpcofferingid
	return p
}

// moves a vpc to another physical network
func (s *VPCService) MigrateVPC(p *MigrateVPCParams) (*MigrateVPCResponse, error) {
	resp, err := s.cs.newPostRequest("migrateVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVPCResponse
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

type MigrateVPCResponse struct {
	Account              string                      `json:"account"`
	Asnumber             int64                       `json:"asnumber"`
	Asnumberid           string                      `json:"asnumberid"`
	Bgppeers             []interface{}               `json:"bgppeers"`
	Cidr                 string                      `json:"cidr"`
	Created              string                      `json:"created"`
	Displaytext          string                      `json:"displaytext"`
	Distributedvpcrouter bool                        `json:"distributedvpcrouter"`
	Dns1                 string                      `json:"dns1"`
	Dns2                 string                      `json:"dns2"`
	Domain               string                      `json:"domain"`
	Domainid             string                      `json:"domainid"`
	Domainpath           string                      `json:"domainpath"`
	Fordisplay           bool                        `json:"fordisplay"`
	Hasannotations       bool                        `json:"hasannotations"`
	Icon                 interface{}                 `json:"icon"`
	Id                   string                      `json:"id"`
	Ip4routes            []interface{}               `json:"ip4routes"`
	Ip4routing           string                      `json:"ip4routing"`
	Ip6dns1              string                      `json:"ip6dns1"`
	Ip6dns2              string                      `json:"ip6dns2"`
	Ip6routes            []interface{}               `json:"ip6routes"`
	JobID                string                      `json:"jobid"`
	Jobstatus            int                         `json:"jobstatus"`
	Name                 string                      `json:"name"`
	Network              []*Network                  `json:"network"`
	Networkdomain        string                      `json:"networkdomain"`
	Project              string                      `json:"project"`
	Projectid            string                      `json:"projectid"`
	Publicmtu            int                         `json:"publicmtu"`
	Redundantvpcrouter   bool                        `json:"redundantvpcrouter"`
	Regionlevelvpc       bool                        `json:"regionlevelvpc"`
	Restartrequired      bool                        `json:"restartrequired"`
	Service              []MigrateVPCResponseService `json:"service"`
	State                string                      `json:"state"`
	Tags                 []Tags                      `json:"tags"`
	Vpcofferingid        string                      `json:"vpcofferingid"`
	Vpcofferingname      string                      `json:"vpcofferingname"`
	Zoneid               string                      `json:"zoneid"`
	Zonename             string                      `json:"zonename"`
}

type MigrateVPCResponseService struct {
	Capability []MigrateVPCResponseServiceCapability `json:"capability"`
	Name       string                                `json:"name"`
	Provider   []MigrateVPCResponseServiceProvider   `json:"provider"`
}

type MigrateVPCResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type MigrateVPCResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type RestartVPCParams struct {
	p map[string]interface{}
}

func (p *RestartVPCParams) toURLValues() url.Values {
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
	if v, found := p.p["livepatch"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("livepatch", vv)
	}
	if v, found := p.p["makeredundant"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("makeredundant", vv)
	}
	return u
}

func (p *RestartVPCParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *RestartVPCParams) ResetCleanup() {
	if p.p != nil && p.p["cleanup"] != nil {
		delete(p.p, "cleanup")
	}
}

func (p *RestartVPCParams) GetCleanup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanup"].(bool)
	return value, ok
}

func (p *RestartVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RestartVPCParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RestartVPCParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *RestartVPCParams) SetLivepatch(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["livepatch"] = v
}

func (p *RestartVPCParams) ResetLivepatch() {
	if p.p != nil && p.p["livepatch"] != nil {
		delete(p.p, "livepatch")
	}
}

func (p *RestartVPCParams) GetLivepatch() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["livepatch"].(bool)
	return value, ok
}

func (p *RestartVPCParams) SetMakeredundant(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["makeredundant"] = v
}

func (p *RestartVPCParams) ResetMakeredundant() {
	if p.p != nil && p.p["makeredundant"] != nil {
		delete(p.p, "makeredundant")
	}
}

func (p *RestartVPCParams) GetMakeredundant() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["makeredundant"].(bool)
	return value, ok
}

// You should always use this function to get a new RestartVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewRestartVPCParams(id string) *RestartVPCParams {
	p := &RestartVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Restarts a VPC
func (s *VPCService) RestartVPC(p *RestartVPCParams) (*RestartVPCResponse, error) {
	resp, err := s.cs.newPostRequest("restartVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestartVPCResponse
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

type RestartVPCResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateVPCParams struct {
	p map[string]interface{}
}

func (p *UpdateVPCParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["publicmtu"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicmtu", vv)
	}
	if v, found := p.p["sourcenatipaddress"]; found {
		u.Set("sourcenatipaddress", v.(string))
	}
	return u
}

func (p *UpdateVPCParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateVPCParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateVPCParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateVPCParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateVPCParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *UpdateVPCParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateVPCParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateVPCParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateVPCParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateVPCParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVPCParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVPCParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateVPCParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVPCParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateVPCParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateVPCParams) SetPublicmtu(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicmtu"] = v
}

func (p *UpdateVPCParams) ResetPublicmtu() {
	if p.p != nil && p.p["publicmtu"] != nil {
		delete(p.p, "publicmtu")
	}
}

func (p *UpdateVPCParams) GetPublicmtu() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicmtu"].(int)
	return value, ok
}

func (p *UpdateVPCParams) SetSourcenatipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatipaddress"] = v
}

func (p *UpdateVPCParams) ResetSourcenatipaddress() {
	if p.p != nil && p.p["sourcenatipaddress"] != nil {
		delete(p.p, "sourcenatipaddress")
	}
}

func (p *UpdateVPCParams) GetSourcenatipaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcenatipaddress"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVPCParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewUpdateVPCParams(id string) *UpdateVPCParams {
	p := &UpdateVPCParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a VPC
func (s *VPCService) UpdateVPC(p *UpdateVPCParams) (*UpdateVPCResponse, error) {
	resp, err := s.cs.newPostRequest("updateVPC", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVPCResponse
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

type UpdateVPCResponse struct {
	Account              string                     `json:"account"`
	Asnumber             int64                      `json:"asnumber"`
	Asnumberid           string                     `json:"asnumberid"`
	Bgppeers             []interface{}              `json:"bgppeers"`
	Cidr                 string                     `json:"cidr"`
	Created              string                     `json:"created"`
	Displaytext          string                     `json:"displaytext"`
	Distributedvpcrouter bool                       `json:"distributedvpcrouter"`
	Dns1                 string                     `json:"dns1"`
	Dns2                 string                     `json:"dns2"`
	Domain               string                     `json:"domain"`
	Domainid             string                     `json:"domainid"`
	Domainpath           string                     `json:"domainpath"`
	Fordisplay           bool                       `json:"fordisplay"`
	Hasannotations       bool                       `json:"hasannotations"`
	Icon                 interface{}                `json:"icon"`
	Id                   string                     `json:"id"`
	Ip4routes            []interface{}              `json:"ip4routes"`
	Ip4routing           string                     `json:"ip4routing"`
	Ip6dns1              string                     `json:"ip6dns1"`
	Ip6dns2              string                     `json:"ip6dns2"`
	Ip6routes            []interface{}              `json:"ip6routes"`
	JobID                string                     `json:"jobid"`
	Jobstatus            int                        `json:"jobstatus"`
	Name                 string                     `json:"name"`
	Network              []*Network                 `json:"network"`
	Networkdomain        string                     `json:"networkdomain"`
	Project              string                     `json:"project"`
	Projectid            string                     `json:"projectid"`
	Publicmtu            int                        `json:"publicmtu"`
	Redundantvpcrouter   bool                       `json:"redundantvpcrouter"`
	Regionlevelvpc       bool                       `json:"regionlevelvpc"`
	Restartrequired      bool                       `json:"restartrequired"`
	Service              []UpdateVPCResponseService `json:"service"`
	State                string                     `json:"state"`
	Tags                 []Tags                     `json:"tags"`
	Vpcofferingid        string                     `json:"vpcofferingid"`
	Vpcofferingname      string                     `json:"vpcofferingname"`
	Zoneid               string                     `json:"zoneid"`
	Zonename             string                     `json:"zonename"`
}

type UpdateVPCResponseService struct {
	Capability []UpdateVPCResponseServiceCapability `json:"capability"`
	Name       string                               `json:"name"`
	Provider   []UpdateVPCResponseServiceProvider   `json:"provider"`
}

type UpdateVPCResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type UpdateVPCResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type UpdateVPCOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateVPCOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *UpdateVPCOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateVPCOfferingParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *UpdateVPCOfferingParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateVPCOfferingParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpdateVPCOfferingParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *UpdateVPCOfferingParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpdateVPCOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVPCOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVPCOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateVPCOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVPCOfferingParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateVPCOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateVPCOfferingParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

func (p *UpdateVPCOfferingParams) ResetSortkey() {
	if p.p != nil && p.p["sortkey"] != nil {
		delete(p.p, "sortkey")
	}
}

func (p *UpdateVPCOfferingParams) GetSortkey() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortkey"].(int)
	return value, ok
}

func (p *UpdateVPCOfferingParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdateVPCOfferingParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *UpdateVPCOfferingParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *UpdateVPCOfferingParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *UpdateVPCOfferingParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *UpdateVPCOfferingParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVPCOfferingParams instance,
// as then you are sure you have configured all required params
func (s *VPCService) NewUpdateVPCOfferingParams(id string) *UpdateVPCOfferingParams {
	p := &UpdateVPCOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates VPC offering
func (s *VPCService) UpdateVPCOffering(p *UpdateVPCOfferingParams) (*UpdateVPCOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("updateVPCOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVPCOfferingResponse
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

type UpdateVPCOfferingResponse struct {
	Created                string                             `json:"created"`
	Displaytext            string                             `json:"displaytext"`
	Distributedvpcrouter   bool                               `json:"distributedvpcrouter"`
	Domain                 string                             `json:"domain"`
	Domainid               string                             `json:"domainid"`
	Fornsx                 bool                               `json:"fornsx"`
	Id                     string                             `json:"id"`
	Internetprotocol       string                             `json:"internetprotocol"`
	Isdefault              bool                               `json:"isdefault"`
	JobID                  string                             `json:"jobid"`
	Jobstatus              int                                `json:"jobstatus"`
	Name                   string                             `json:"name"`
	Networkmode            string                             `json:"networkmode"`
	Routingmode            string                             `json:"routingmode"`
	Service                []UpdateVPCOfferingResponseService `json:"service"`
	Specifyasnumber        bool                               `json:"specifyasnumber"`
	State                  string                             `json:"state"`
	SupportsregionLevelvpc bool                               `json:"supportsregionLevelvpc"`
	Zone                   string                             `json:"zone"`
	Zoneid                 string                             `json:"zoneid"`
}

type UpdateVPCOfferingResponseService struct {
	Capability []UpdateVPCOfferingResponseServiceCapability `json:"capability"`
	Name       string                                       `json:"name"`
	Provider   []UpdateVPCOfferingResponseServiceProvider   `json:"provider"`
}

type UpdateVPCOfferingResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type UpdateVPCOfferingResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}
