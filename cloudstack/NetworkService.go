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

type NetworkServiceIface interface {
	AddNetworkServiceProvider(p *AddNetworkServiceProviderParams) (*AddNetworkServiceProviderResponse, error)
	NewAddNetworkServiceProviderParams(name string, physicalnetworkid string) *AddNetworkServiceProviderParams
	AddOpenDaylightController(p *AddOpenDaylightControllerParams) (*AddOpenDaylightControllerResponse, error)
	NewAddOpenDaylightControllerParams(password string, physicalnetworkid string, url string, username string) *AddOpenDaylightControllerParams
	CreateNetwork(p *CreateNetworkParams) (*CreateNetworkResponse, error)
	NewCreateNetworkParams(displaytext string, name string, networkofferingid string, zoneid string) *CreateNetworkParams
	CreatePhysicalNetwork(p *CreatePhysicalNetworkParams) (*CreatePhysicalNetworkResponse, error)
	NewCreatePhysicalNetworkParams(name string, zoneid string) *CreatePhysicalNetworkParams
	CreateServiceInstance(p *CreateServiceInstanceParams) (*CreateServiceInstanceResponse, error)
	NewCreateServiceInstanceParams(leftnetworkid string, name string, rightnetworkid string, serviceofferingid string, templateid string, zoneid string) *CreateServiceInstanceParams
	CreateStorageNetworkIpRange(p *CreateStorageNetworkIpRangeParams) (*CreateStorageNetworkIpRangeResponse, error)
	NewCreateStorageNetworkIpRangeParams(gateway string, netmask string, podid string, startip string) *CreateStorageNetworkIpRangeParams
	DedicatePublicIpRange(p *DedicatePublicIpRangeParams) (*DedicatePublicIpRangeResponse, error)
	NewDedicatePublicIpRangeParams(domainid string, id string) *DedicatePublicIpRangeParams
	DeleteNetwork(p *DeleteNetworkParams) (*DeleteNetworkResponse, error)
	NewDeleteNetworkParams(id string) *DeleteNetworkParams
	DeleteNetworkServiceProvider(p *DeleteNetworkServiceProviderParams) (*DeleteNetworkServiceProviderResponse, error)
	NewDeleteNetworkServiceProviderParams(id string) *DeleteNetworkServiceProviderParams
	DeleteOpenDaylightController(p *DeleteOpenDaylightControllerParams) (*DeleteOpenDaylightControllerResponse, error)
	NewDeleteOpenDaylightControllerParams(id string) *DeleteOpenDaylightControllerParams
	DeletePhysicalNetwork(p *DeletePhysicalNetworkParams) (*DeletePhysicalNetworkResponse, error)
	NewDeletePhysicalNetworkParams(id string) *DeletePhysicalNetworkParams
	DeleteStorageNetworkIpRange(p *DeleteStorageNetworkIpRangeParams) (*DeleteStorageNetworkIpRangeResponse, error)
	NewDeleteStorageNetworkIpRangeParams(id string) *DeleteStorageNetworkIpRangeParams
	ListNetscalerLoadBalancerNetworks(p *ListNetscalerLoadBalancerNetworksParams) (*ListNetscalerLoadBalancerNetworksResponse, error)
	NewListNetscalerLoadBalancerNetworksParams(lbdeviceid string) *ListNetscalerLoadBalancerNetworksParams
	GetNetscalerLoadBalancerNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error)
	ListNetworkIsolationMethods(p *ListNetworkIsolationMethodsParams) (*ListNetworkIsolationMethodsResponse, error)
	NewListNetworkIsolationMethodsParams() *ListNetworkIsolationMethodsParams
	ListNetworkServiceProviders(p *ListNetworkServiceProvidersParams) (*ListNetworkServiceProvidersResponse, error)
	NewListNetworkServiceProvidersParams() *ListNetworkServiceProvidersParams
	GetNetworkServiceProviderID(name string, opts ...OptionFunc) (string, int, error)
	ListNetworks(p *ListNetworksParams) (*ListNetworksResponse, error)
	NewListNetworksParams() *ListNetworksParams
	GetNetworkID(keyword string, opts ...OptionFunc) (string, int, error)
	GetNetworkByName(name string, opts ...OptionFunc) (*Network, int, error)
	GetNetworkByID(id string, opts ...OptionFunc) (*Network, int, error)
	ListNiciraNvpDeviceNetworks(p *ListNiciraNvpDeviceNetworksParams) (*ListNiciraNvpDeviceNetworksResponse, error)
	NewListNiciraNvpDeviceNetworksParams(nvpdeviceid string) *ListNiciraNvpDeviceNetworksParams
	GetNiciraNvpDeviceNetworkID(keyword string, nvpdeviceid string, opts ...OptionFunc) (string, int, error)
	ListOpenDaylightControllers(p *ListOpenDaylightControllersParams) (*ListOpenDaylightControllersResponse, error)
	NewListOpenDaylightControllersParams() *ListOpenDaylightControllersParams
	GetOpenDaylightControllerByID(id string, opts ...OptionFunc) (*OpenDaylightController, int, error)
	ListPaloAltoFirewallNetworks(p *ListPaloAltoFirewallNetworksParams) (*ListPaloAltoFirewallNetworksResponse, error)
	NewListPaloAltoFirewallNetworksParams(lbdeviceid string) *ListPaloAltoFirewallNetworksParams
	GetPaloAltoFirewallNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error)
	ListPhysicalNetworks(p *ListPhysicalNetworksParams) (*ListPhysicalNetworksResponse, error)
	NewListPhysicalNetworksParams() *ListPhysicalNetworksParams
	GetPhysicalNetworkID(name string, opts ...OptionFunc) (string, int, error)
	GetPhysicalNetworkByName(name string, opts ...OptionFunc) (*PhysicalNetwork, int, error)
	GetPhysicalNetworkByID(id string, opts ...OptionFunc) (*PhysicalNetwork, int, error)
	ListStorageNetworkIpRange(p *ListStorageNetworkIpRangeParams) (*ListStorageNetworkIpRangeResponse, error)
	NewListStorageNetworkIpRangeParams() *ListStorageNetworkIpRangeParams
	GetStorageNetworkIpRangeByID(id string, opts ...OptionFunc) (*StorageNetworkIpRange, int, error)
	ListSupportedNetworkServices(p *ListSupportedNetworkServicesParams) (*ListSupportedNetworkServicesResponse, error)
	NewListSupportedNetworkServicesParams() *ListSupportedNetworkServicesParams
	ReleasePublicIpRange(p *ReleasePublicIpRangeParams) (*ReleasePublicIpRangeResponse, error)
	NewReleasePublicIpRangeParams(id string) *ReleasePublicIpRangeParams
	RestartNetwork(p *RestartNetworkParams) (*RestartNetworkResponse, error)
	NewRestartNetworkParams(id string) *RestartNetworkParams
	UpdateNetwork(p *UpdateNetworkParams) (*UpdateNetworkResponse, error)
	NewUpdateNetworkParams(id string) *UpdateNetworkParams
	UpdateNetworkServiceProvider(p *UpdateNetworkServiceProviderParams) (*UpdateNetworkServiceProviderResponse, error)
	NewUpdateNetworkServiceProviderParams(id string) *UpdateNetworkServiceProviderParams
	UpdatePhysicalNetwork(p *UpdatePhysicalNetworkParams) (*UpdatePhysicalNetworkResponse, error)
	NewUpdatePhysicalNetworkParams(id string) *UpdatePhysicalNetworkParams
	UpdateStorageNetworkIpRange(p *UpdateStorageNetworkIpRangeParams) (*UpdateStorageNetworkIpRangeResponse, error)
	NewUpdateStorageNetworkIpRangeParams(id string) *UpdateStorageNetworkIpRangeParams
}

type AddNetworkServiceProviderParams struct {
	p map[string]interface{}
}

func (p *AddNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["destinationphysicalnetworkid"]; found {
		u.Set("destinationphysicalnetworkid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["servicelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("servicelist", vv)
	}
	return u
}

func (p *AddNetworkServiceProviderParams) SetDestinationphysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destinationphysicalnetworkid"] = v
}

func (p *AddNetworkServiceProviderParams) GetDestinationphysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destinationphysicalnetworkid"].(string)
	return value, ok
}

func (p *AddNetworkServiceProviderParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddNetworkServiceProviderParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddNetworkServiceProviderParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddNetworkServiceProviderParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddNetworkServiceProviderParams) SetServicelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicelist"] = v
}

func (p *AddNetworkServiceProviderParams) GetServicelist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["servicelist"].([]string)
	return value, ok
}

// You should always use this function to get a new AddNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewAddNetworkServiceProviderParams(name string, physicalnetworkid string) *AddNetworkServiceProviderParams {
	p := &AddNetworkServiceProviderParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["physicalnetworkid"] = physicalnetworkid
	return p
}

// Adds a network serviceProvider to a physical network
func (s *NetworkService) AddNetworkServiceProvider(p *AddNetworkServiceProviderParams) (*AddNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("addNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetworkServiceProviderResponse
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

type AddNetworkServiceProviderResponse struct {
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

type AddOpenDaylightControllerParams struct {
	p map[string]interface{}
}

func (p *AddOpenDaylightControllerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddOpenDaylightControllerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddOpenDaylightControllerParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddOpenDaylightControllerParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddOpenDaylightControllerParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddOpenDaylightControllerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddOpenDaylightControllerParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddOpenDaylightControllerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddOpenDaylightControllerParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddOpenDaylightControllerParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewAddOpenDaylightControllerParams(password string, physicalnetworkid string, url string, username string) *AddOpenDaylightControllerParams {
	p := &AddOpenDaylightControllerParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds an OpenDyalight controler
func (s *NetworkService) AddOpenDaylightController(p *AddOpenDaylightControllerParams) (*AddOpenDaylightControllerResponse, error) {
	resp, err := s.cs.newRequest("addOpenDaylightController", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddOpenDaylightControllerResponse
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

type AddOpenDaylightControllerResponse struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Name              string `json:"name"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Url               string `json:"url"`
	Username          string `json:"username"`
}

type CreateNetworkParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := p.p["acltype"]; found {
		u.Set("acltype", v.(string))
	}
	if v, found := p.p["bypassvlanoverlapcheck"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bypassvlanoverlapcheck", vv)
	}
	if v, found := p.p["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["endipv6"]; found {
		u.Set("endipv6", v.(string))
	}
	if v, found := p.p["externalid"]; found {
		u.Set("externalid", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["hideipaddressusage"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hideipaddressusage", vv)
	}
	if v, found := p.p["ip6cidr"]; found {
		u.Set("ip6cidr", v.(string))
	}
	if v, found := p.p["ip6gateway"]; found {
		u.Set("ip6gateway", v.(string))
	}
	if v, found := p.p["isolatedpvlan"]; found {
		u.Set("isolatedpvlan", v.(string))
	}
	if v, found := p.p["isolatedpvlantype"]; found {
		u.Set("isolatedpvlantype", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["routerip"]; found {
		u.Set("routerip", v.(string))
	}
	if v, found := p.p["routeripv6"]; found {
		u.Set("routeripv6", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := p.p["startipv6"]; found {
		u.Set("startipv6", v.(string))
	}
	if v, found := p.p["subdomainaccess"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("subdomainaccess", vv)
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateNetworkParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateNetworkParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetAclid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["aclid"] = v
}

func (p *CreateNetworkParams) GetAclid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["aclid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetAcltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["acltype"] = v
}

func (p *CreateNetworkParams) GetAcltype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["acltype"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetBypassvlanoverlapcheck(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bypassvlanoverlapcheck"] = v
}

func (p *CreateNetworkParams) GetBypassvlanoverlapcheck() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bypassvlanoverlapcheck"].(bool)
	return value, ok
}

func (p *CreateNetworkParams) SetDisplaynetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaynetwork"] = v
}

func (p *CreateNetworkParams) GetDisplaynetwork() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaynetwork"].(bool)
	return value, ok
}

func (p *CreateNetworkParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateNetworkParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateNetworkParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
}

func (p *CreateNetworkParams) GetEndip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endip"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetEndipv6(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endipv6"] = v
}

func (p *CreateNetworkParams) GetEndipv6() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endipv6"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetExternalid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["externalid"] = v
}

func (p *CreateNetworkParams) GetExternalid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["externalid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
}

func (p *CreateNetworkParams) GetGateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gateway"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetHideipaddressusage(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hideipaddressusage"] = v
}

func (p *CreateNetworkParams) GetHideipaddressusage() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hideipaddressusage"].(bool)
	return value, ok
}

func (p *CreateNetworkParams) SetIp6cidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6cidr"] = v
}

func (p *CreateNetworkParams) GetIp6cidr() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6cidr"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetIp6gateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6gateway"] = v
}

func (p *CreateNetworkParams) GetIp6gateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6gateway"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetIsolatedpvlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isolatedpvlan"] = v
}

func (p *CreateNetworkParams) GetIsolatedpvlan() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isolatedpvlan"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetIsolatedpvlantype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isolatedpvlantype"] = v
}

func (p *CreateNetworkParams) GetIsolatedpvlantype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isolatedpvlantype"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateNetworkParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
}

func (p *CreateNetworkParams) GetNetmask() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["netmask"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *CreateNetworkParams) GetNetworkdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdomain"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
}

func (p *CreateNetworkParams) GetNetworkofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkofferingid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *CreateNetworkParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateNetworkParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetRouterip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["routerip"] = v
}

func (p *CreateNetworkParams) GetRouterip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["routerip"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetRouteripv6(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["routeripv6"] = v
}

func (p *CreateNetworkParams) GetRouteripv6() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["routeripv6"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
}

func (p *CreateNetworkParams) GetStartip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startip"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetStartipv6(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startipv6"] = v
}

func (p *CreateNetworkParams) GetStartipv6() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startipv6"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetSubdomainaccess(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["subdomainaccess"] = v
}

func (p *CreateNetworkParams) GetSubdomainaccess() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["subdomainaccess"].(bool)
	return value, ok
}

func (p *CreateNetworkParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *CreateNetworkParams) GetVlan() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *CreateNetworkParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateNetworkParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateNetworkParams(displaytext string, name string, networkofferingid string, zoneid string) *CreateNetworkParams {
	p := &CreateNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["networkofferingid"] = networkofferingid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a network
func (s *NetworkService) CreateNetwork(p *CreateNetworkParams) (*CreateNetworkResponse, error) {
	resp, err := s.cs.newRequest("createNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateNetworkResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateNetworkResponse struct {
	Account                     string                         `json:"account"`
	Aclid                       string                         `json:"aclid"`
	Aclname                     string                         `json:"aclname"`
	Acltype                     string                         `json:"acltype"`
	Broadcastdomaintype         string                         `json:"broadcastdomaintype"`
	Broadcasturi                string                         `json:"broadcasturi"`
	Canusefordeploy             bool                           `json:"canusefordeploy"`
	Cidr                        string                         `json:"cidr"`
	Created                     string                         `json:"created"`
	Details                     map[string]string              `json:"details"`
	Displaynetwork              bool                           `json:"displaynetwork"`
	Displaytext                 string                         `json:"displaytext"`
	Dns1                        string                         `json:"dns1"`
	Dns2                        string                         `json:"dns2"`
	Domain                      string                         `json:"domain"`
	Domainid                    string                         `json:"domainid"`
	Externalid                  string                         `json:"externalid"`
	Gateway                     string                         `json:"gateway"`
	Hasannotations              bool                           `json:"hasannotations"`
	Icon                        string                         `json:"icon"`
	Id                          string                         `json:"id"`
	Ip6cidr                     string                         `json:"ip6cidr"`
	Ip6gateway                  string                         `json:"ip6gateway"`
	Isdefault                   bool                           `json:"isdefault"`
	Ispersistent                bool                           `json:"ispersistent"`
	Issystem                    bool                           `json:"issystem"`
	JobID                       string                         `json:"jobid"`
	Jobstatus                   int                            `json:"jobstatus"`
	Name                        string                         `json:"name"`
	Netmask                     string                         `json:"netmask"`
	Networkcidr                 string                         `json:"networkcidr"`
	Networkdomain               string                         `json:"networkdomain"`
	Networkofferingavailability string                         `json:"networkofferingavailability"`
	Networkofferingconservemode bool                           `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                         `json:"networkofferingdisplaytext"`
	Networkofferingid           string                         `json:"networkofferingid"`
	Networkofferingname         string                         `json:"networkofferingname"`
	Physicalnetworkid           string                         `json:"physicalnetworkid"`
	Project                     string                         `json:"project"`
	Projectid                   string                         `json:"projectid"`
	Receivedbytes               int64                          `json:"receivedbytes"`
	Redundantrouter             bool                           `json:"redundantrouter"`
	Related                     string                         `json:"related"`
	Reservediprange             string                         `json:"reservediprange"`
	Restartrequired             bool                           `json:"restartrequired"`
	Sentbytes                   int64                          `json:"sentbytes"`
	Service                     []CreateNetworkResponseService `json:"service"`
	Specifyipranges             bool                           `json:"specifyipranges"`
	State                       string                         `json:"state"`
	Strechedl2subnet            bool                           `json:"strechedl2subnet"`
	Subdomainaccess             bool                           `json:"subdomainaccess"`
	Tags                        []Tags                         `json:"tags"`
	Traffictype                 string                         `json:"traffictype"`
	Type                        string                         `json:"type"`
	Vlan                        string                         `json:"vlan"`
	Vpcid                       string                         `json:"vpcid"`
	Vpcname                     string                         `json:"vpcname"`
	Zoneid                      string                         `json:"zoneid"`
	Zonename                    string                         `json:"zonename"`
	Zonesnetworkspans           []interface{}                  `json:"zonesnetworkspans"`
}

type CreateNetworkResponseService struct {
	Capability []CreateNetworkResponseServiceCapability `json:"capability"`
	Name       string                                   `json:"name"`
	Provider   []CreateNetworkResponseServiceProvider   `json:"provider"`
}

type CreateNetworkResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type CreateNetworkResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type CreatePhysicalNetworkParams struct {
	p map[string]interface{}
}

func (p *CreatePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["broadcastdomainrange"]; found {
		u.Set("broadcastdomainrange", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["isolationmethods"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("isolationmethods", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkspeed"]; found {
		u.Set("networkspeed", v.(string))
	}
	if v, found := p.p["tags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("tags", vv)
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreatePhysicalNetworkParams) SetBroadcastdomainrange(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["broadcastdomainrange"] = v
}

func (p *CreatePhysicalNetworkParams) GetBroadcastdomainrange() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["broadcastdomainrange"].(string)
	return value, ok
}

func (p *CreatePhysicalNetworkParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreatePhysicalNetworkParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreatePhysicalNetworkParams) SetIsolationmethods(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isolationmethods"] = v
}

func (p *CreatePhysicalNetworkParams) GetIsolationmethods() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isolationmethods"].([]string)
	return value, ok
}

func (p *CreatePhysicalNetworkParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreatePhysicalNetworkParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreatePhysicalNetworkParams) SetNetworkspeed(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkspeed"] = v
}

func (p *CreatePhysicalNetworkParams) GetNetworkspeed() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkspeed"].(string)
	return value, ok
}

func (p *CreatePhysicalNetworkParams) SetTags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreatePhysicalNetworkParams) GetTags() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].([]string)
	return value, ok
}

func (p *CreatePhysicalNetworkParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *CreatePhysicalNetworkParams) GetVlan() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(string)
	return value, ok
}

func (p *CreatePhysicalNetworkParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreatePhysicalNetworkParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreatePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreatePhysicalNetworkParams(name string, zoneid string) *CreatePhysicalNetworkParams {
	p := &CreatePhysicalNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// Creates a physical network
func (s *NetworkService) CreatePhysicalNetwork(p *CreatePhysicalNetworkParams) (*CreatePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("createPhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePhysicalNetworkResponse
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

type CreatePhysicalNetworkResponse struct {
	Broadcastdomainrange string `json:"broadcastdomainrange"`
	Domainid             string `json:"domainid"`
	Id                   string `json:"id"`
	Isolationmethods     string `json:"isolationmethods"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Name                 string `json:"name"`
	Networkspeed         string `json:"networkspeed"`
	State                string `json:"state"`
	Tags                 string `json:"tags"`
	Vlan                 string `json:"vlan"`
	Zoneid               string `json:"zoneid"`
	Zonename             string `json:"zonename"`
}

type CreateServiceInstanceParams struct {
	p map[string]interface{}
}

func (p *CreateServiceInstanceParams) toURLValues() url.Values {
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
	if v, found := p.p["leftnetworkid"]; found {
		u.Set("leftnetworkid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["rightnetworkid"]; found {
		u.Set("rightnetworkid", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateServiceInstanceParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateServiceInstanceParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateServiceInstanceParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateServiceInstanceParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateServiceInstanceParams) SetLeftnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["leftnetworkid"] = v
}

func (p *CreateServiceInstanceParams) GetLeftnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["leftnetworkid"].(string)
	return value, ok
}

func (p *CreateServiceInstanceParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateServiceInstanceParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateServiceInstanceParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateServiceInstanceParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateServiceInstanceParams) SetRightnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rightnetworkid"] = v
}

func (p *CreateServiceInstanceParams) GetRightnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rightnetworkid"].(string)
	return value, ok
}

func (p *CreateServiceInstanceParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateServiceInstanceParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *CreateServiceInstanceParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *CreateServiceInstanceParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *CreateServiceInstanceParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateServiceInstanceParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateServiceInstanceParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateServiceInstanceParams(leftnetworkid string, name string, rightnetworkid string, serviceofferingid string, templateid string, zoneid string) *CreateServiceInstanceParams {
	p := &CreateServiceInstanceParams{}
	p.p = make(map[string]interface{})
	p.p["leftnetworkid"] = leftnetworkid
	p.p["name"] = name
	p.p["rightnetworkid"] = rightnetworkid
	p.p["serviceofferingid"] = serviceofferingid
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a system virtual-machine that implements network services
func (s *NetworkService) CreateServiceInstance(p *CreateServiceInstanceParams) (*CreateServiceInstanceResponse, error) {
	resp, err := s.cs.newRequest("createServiceInstance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateServiceInstanceResponse
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

type CreateServiceInstanceResponse struct {
	Account     string `json:"account"`
	Displayname string `json:"displayname"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
}

type CreateStorageNetworkIpRangeParams struct {
	p map[string]interface{}
}

func (p *CreateStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("vlan", vv)
	}
	return u
}

func (p *CreateStorageNetworkIpRangeParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
}

func (p *CreateStorageNetworkIpRangeParams) GetEndip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endip"].(string)
	return value, ok
}

func (p *CreateStorageNetworkIpRangeParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
}

func (p *CreateStorageNetworkIpRangeParams) GetGateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["gateway"].(string)
	return value, ok
}

func (p *CreateStorageNetworkIpRangeParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
}

func (p *CreateStorageNetworkIpRangeParams) GetNetmask() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["netmask"].(string)
	return value, ok
}

func (p *CreateStorageNetworkIpRangeParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *CreateStorageNetworkIpRangeParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *CreateStorageNetworkIpRangeParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
}

func (p *CreateStorageNetworkIpRangeParams) GetStartip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startip"].(string)
	return value, ok
}

func (p *CreateStorageNetworkIpRangeParams) SetVlan(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *CreateStorageNetworkIpRangeParams) GetVlan() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(int)
	return value, ok
}

// You should always use this function to get a new CreateStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateStorageNetworkIpRangeParams(gateway string, netmask string, podid string, startip string) *CreateStorageNetworkIpRangeParams {
	p := &CreateStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["gateway"] = gateway
	p.p["netmask"] = netmask
	p.p["podid"] = podid
	p.p["startip"] = startip
	return p
}

// Creates a Storage network IP range.
func (s *NetworkService) CreateStorageNetworkIpRange(p *CreateStorageNetworkIpRangeParams) (*CreateStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("createStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateStorageNetworkIpRangeResponse
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

type CreateStorageNetworkIpRangeResponse struct {
	Endip     string `json:"endip"`
	Gateway   string `json:"gateway"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Netmask   string `json:"netmask"`
	Networkid string `json:"networkid"`
	Podid     string `json:"podid"`
	Startip   string `json:"startip"`
	Vlan      int    `json:"vlan"`
	Zoneid    string `json:"zoneid"`
}

type DedicatePublicIpRangeParams struct {
	p map[string]interface{}
}

func (p *DedicatePublicIpRangeParams) toURLValues() url.Values {
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

func (p *DedicatePublicIpRangeParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DedicatePublicIpRangeParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DedicatePublicIpRangeParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DedicatePublicIpRangeParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DedicatePublicIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DedicatePublicIpRangeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DedicatePublicIpRangeParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DedicatePublicIpRangeParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicatePublicIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDedicatePublicIpRangeParams(domainid string, id string) *DedicatePublicIpRangeParams {
	p := &DedicatePublicIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	p.p["id"] = id
	return p
}

// Dedicates a Public IP range to an account
func (s *NetworkService) DedicatePublicIpRange(p *DedicatePublicIpRangeParams) (*DedicatePublicIpRangeResponse, error) {
	resp, err := s.cs.newRequest("dedicatePublicIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicatePublicIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DedicatePublicIpRangeResponse struct {
	Account           string `json:"account"`
	Description       string `json:"description"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Endip             string `json:"endip"`
	Endipv6           string `json:"endipv6"`
	Forsystemvms      bool   `json:"forsystemvms"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork"`
	Gateway           string `json:"gateway"`
	Id                string `json:"id"`
	Ip6cidr           string `json:"ip6cidr"`
	Ip6gateway        string `json:"ip6gateway"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Netmask           string `json:"netmask"`
	Networkid         string `json:"networkid"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Project           string `json:"project"`
	Projectid         string `json:"projectid"`
	Startip           string `json:"startip"`
	Startipv6         string `json:"startipv6"`
	Vlan              string `json:"vlan"`
	Zoneid            string `json:"zoneid"`
}

type DeleteNetworkParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DeleteNetworkParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *DeleteNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteNetworkParams(id string) *DeleteNetworkParams {
	p := &DeleteNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a network
func (s *NetworkService) DeleteNetwork(p *DeleteNetworkParams) (*DeleteNetworkResponse, error) {
	resp, err := s.cs.newRequest("deleteNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkResponse
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

type DeleteNetworkResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteNetworkServiceProviderParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkServiceProviderParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteNetworkServiceProviderParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteNetworkServiceProviderParams(id string) *DeleteNetworkServiceProviderParams {
	p := &DeleteNetworkServiceProviderParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Network Service Provider.
func (s *NetworkService) DeleteNetworkServiceProvider(p *DeleteNetworkServiceProviderParams) (*DeleteNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkServiceProviderResponse
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

type DeleteNetworkServiceProviderResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteOpenDaylightControllerParams struct {
	p map[string]interface{}
}

func (p *DeleteOpenDaylightControllerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteOpenDaylightControllerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteOpenDaylightControllerParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteOpenDaylightControllerParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteOpenDaylightControllerParams(id string) *DeleteOpenDaylightControllerParams {
	p := &DeleteOpenDaylightControllerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes an OpenDyalight controler
func (s *NetworkService) DeleteOpenDaylightController(p *DeleteOpenDaylightControllerParams) (*DeleteOpenDaylightControllerResponse, error) {
	resp, err := s.cs.newRequest("deleteOpenDaylightController", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteOpenDaylightControllerResponse
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

type DeleteOpenDaylightControllerResponse struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Name              string `json:"name"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Url               string `json:"url"`
	Username          string `json:"username"`
}

type DeletePhysicalNetworkParams struct {
	p map[string]interface{}
}

func (p *DeletePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePhysicalNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeletePhysicalNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeletePhysicalNetworkParams(id string) *DeletePhysicalNetworkParams {
	p := &DeletePhysicalNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Physical Network.
func (s *NetworkService) DeletePhysicalNetwork(p *DeletePhysicalNetworkParams) (*DeletePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("deletePhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePhysicalNetworkResponse
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

type DeletePhysicalNetworkResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteStorageNetworkIpRangeParams struct {
	p map[string]interface{}
}

func (p *DeleteStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteStorageNetworkIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteStorageNetworkIpRangeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteStorageNetworkIpRangeParams(id string) *DeleteStorageNetworkIpRangeParams {
	p := &DeleteStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a storage network IP Range.
func (s *NetworkService) DeleteStorageNetworkIpRange(p *DeleteStorageNetworkIpRangeParams) (*DeleteStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("deleteStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStorageNetworkIpRangeResponse
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

type DeleteStorageNetworkIpRangeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListNetscalerLoadBalancerNetworksParams struct {
	p map[string]interface{}
}

func (p *ListNetscalerLoadBalancerNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetscalerLoadBalancerNetworksParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
}

func (p *ListNetscalerLoadBalancerNetworksParams) GetLbdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbdeviceid"].(string)
	return value, ok
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetscalerLoadBalancerNetworksParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetscalerLoadBalancerNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetscalerLoadBalancerNetworksParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNetscalerLoadBalancerNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetscalerLoadBalancerNetworksParams(lbdeviceid string) *ListNetscalerLoadBalancerNetworksParams {
	p := &ListNetscalerLoadBalancerNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetscalerLoadBalancerNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error) {
	p := &ListNetscalerLoadBalancerNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["lbdeviceid"] = lbdeviceid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetscalerLoadBalancerNetworks(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.NetscalerLoadBalancerNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetscalerLoadBalancerNetworks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// lists network that are using a netscaler load balancer device
func (s *NetworkService) ListNetscalerLoadBalancerNetworks(p *ListNetscalerLoadBalancerNetworksParams) (*ListNetscalerLoadBalancerNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerLoadBalancerNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerLoadBalancerNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetscalerLoadBalancerNetworksResponse struct {
	Count                         int                             `json:"count"`
	NetscalerLoadBalancerNetworks []*NetscalerLoadBalancerNetwork `json:"netscalerloadbalancernetwork"`
}

type NetscalerLoadBalancerNetwork struct {
	Account                     string                                `json:"account"`
	Aclid                       string                                `json:"aclid"`
	Aclname                     string                                `json:"aclname"`
	Acltype                     string                                `json:"acltype"`
	Broadcastdomaintype         string                                `json:"broadcastdomaintype"`
	Broadcasturi                string                                `json:"broadcasturi"`
	Canusefordeploy             bool                                  `json:"canusefordeploy"`
	Cidr                        string                                `json:"cidr"`
	Created                     string                                `json:"created"`
	Details                     map[string]string                     `json:"details"`
	Displaynetwork              bool                                  `json:"displaynetwork"`
	Displaytext                 string                                `json:"displaytext"`
	Dns1                        string                                `json:"dns1"`
	Dns2                        string                                `json:"dns2"`
	Domain                      string                                `json:"domain"`
	Domainid                    string                                `json:"domainid"`
	Externalid                  string                                `json:"externalid"`
	Gateway                     string                                `json:"gateway"`
	Hasannotations              bool                                  `json:"hasannotations"`
	Icon                        string                                `json:"icon"`
	Id                          string                                `json:"id"`
	Ip6cidr                     string                                `json:"ip6cidr"`
	Ip6gateway                  string                                `json:"ip6gateway"`
	Isdefault                   bool                                  `json:"isdefault"`
	Ispersistent                bool                                  `json:"ispersistent"`
	Issystem                    bool                                  `json:"issystem"`
	JobID                       string                                `json:"jobid"`
	Jobstatus                   int                                   `json:"jobstatus"`
	Name                        string                                `json:"name"`
	Netmask                     string                                `json:"netmask"`
	Networkcidr                 string                                `json:"networkcidr"`
	Networkdomain               string                                `json:"networkdomain"`
	Networkofferingavailability string                                `json:"networkofferingavailability"`
	Networkofferingconservemode bool                                  `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                                `json:"networkofferingdisplaytext"`
	Networkofferingid           string                                `json:"networkofferingid"`
	Networkofferingname         string                                `json:"networkofferingname"`
	Physicalnetworkid           string                                `json:"physicalnetworkid"`
	Project                     string                                `json:"project"`
	Projectid                   string                                `json:"projectid"`
	Receivedbytes               int64                                 `json:"receivedbytes"`
	Redundantrouter             bool                                  `json:"redundantrouter"`
	Related                     string                                `json:"related"`
	Reservediprange             string                                `json:"reservediprange"`
	Restartrequired             bool                                  `json:"restartrequired"`
	Sentbytes                   int64                                 `json:"sentbytes"`
	Service                     []NetscalerLoadBalancerNetworkService `json:"service"`
	Specifyipranges             bool                                  `json:"specifyipranges"`
	State                       string                                `json:"state"`
	Strechedl2subnet            bool                                  `json:"strechedl2subnet"`
	Subdomainaccess             bool                                  `json:"subdomainaccess"`
	Tags                        []Tags                                `json:"tags"`
	Traffictype                 string                                `json:"traffictype"`
	Type                        string                                `json:"type"`
	Vlan                        string                                `json:"vlan"`
	Vpcid                       string                                `json:"vpcid"`
	Vpcname                     string                                `json:"vpcname"`
	Zoneid                      string                                `json:"zoneid"`
	Zonename                    string                                `json:"zonename"`
	Zonesnetworkspans           []interface{}                         `json:"zonesnetworkspans"`
}

type NetscalerLoadBalancerNetworkService struct {
	Capability []NetscalerLoadBalancerNetworkServiceCapability `json:"capability"`
	Name       string                                          `json:"name"`
	Provider   []NetscalerLoadBalancerNetworkServiceProvider   `json:"provider"`
}

type NetscalerLoadBalancerNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NetscalerLoadBalancerNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListNetworkIsolationMethodsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkIsolationMethodsParams) toURLValues() url.Values {
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
	return u
}

func (p *ListNetworkIsolationMethodsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetworkIsolationMethodsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetworkIsolationMethodsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetworkIsolationMethodsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetworkIsolationMethodsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetworkIsolationMethodsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNetworkIsolationMethodsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworkIsolationMethodsParams() *ListNetworkIsolationMethodsParams {
	p := &ListNetworkIsolationMethodsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists supported methods of network isolation
func (s *NetworkService) ListNetworkIsolationMethods(p *ListNetworkIsolationMethodsParams) (*ListNetworkIsolationMethodsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkIsolationMethods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkIsolationMethodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkIsolationMethodsResponse struct {
	Count                   int                       `json:"count"`
	NetworkIsolationMethods []*NetworkIsolationMethod `json:"networkisolationmethod"`
}

type NetworkIsolationMethod struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListNetworkServiceProvidersParams struct {
	p map[string]interface{}
}

func (p *ListNetworkServiceProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *ListNetworkServiceProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetworkServiceProvidersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetworkServiceProvidersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListNetworkServiceProvidersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListNetworkServiceProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetworkServiceProvidersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetworkServiceProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetworkServiceProvidersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNetworkServiceProvidersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListNetworkServiceProvidersParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *ListNetworkServiceProvidersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListNetworkServiceProvidersParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkServiceProvidersParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworkServiceProvidersParams() *ListNetworkServiceProvidersParams {
	p := &ListNetworkServiceProvidersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkServiceProviderID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListNetworkServiceProvidersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetworkServiceProviders(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.NetworkServiceProviders[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetworkServiceProviders {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// Lists network serviceproviders for a given physical network.
func (s *NetworkService) ListNetworkServiceProviders(p *ListNetworkServiceProvidersParams) (*ListNetworkServiceProvidersResponse, error) {
	resp, err := s.cs.newRequest("listNetworkServiceProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkServiceProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkServiceProvidersResponse struct {
	Count                   int                       `json:"count"`
	NetworkServiceProviders []*NetworkServiceProvider `json:"networkserviceprovider"`
}

type NetworkServiceProvider struct {
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

type ListNetworksParams struct {
	p map[string]interface{}
}

func (p *ListNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["acltype"]; found {
		u.Set("acltype", v.(string))
	}
	if v, found := p.p["canusefordeploy"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("canusefordeploy", vv)
	}
	if v, found := p.p["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
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
	if v, found := p.p["restartrequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("restartrequired", vv)
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
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
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListNetworksParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListNetworksParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetAcltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["acltype"] = v
}

func (p *ListNetworksParams) GetAcltype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["acltype"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetCanusefordeploy(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["canusefordeploy"] = v
}

func (p *ListNetworksParams) GetCanusefordeploy() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["canusefordeploy"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetDisplaynetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaynetwork"] = v
}

func (p *ListNetworksParams) GetDisplaynetwork() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaynetwork"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListNetworksParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
}

func (p *ListNetworksParams) GetForvpc() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forvpc"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListNetworksParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListNetworksParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
}

func (p *ListNetworksParams) GetIssystem() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["issystem"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetworksParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListNetworksParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
}

func (p *ListNetworksParams) GetNetworkofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkofferingid"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetworksParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetworksParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNetworksParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListNetworksParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListNetworksParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetRestartrequired(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["restartrequired"] = v
}

func (p *ListNetworksParams) GetRestartrequired() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["restartrequired"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListNetworksParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetSpecifyipranges(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyipranges"] = v
}

func (p *ListNetworksParams) GetSpecifyipranges() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["specifyipranges"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

func (p *ListNetworksParams) GetSupportedservices() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["supportedservices"].([]string)
	return value, ok
}

func (p *ListNetworksParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListNetworksParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListNetworksParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *ListNetworksParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListNetworksParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListNetworksParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListNetworksParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworksParams() *ListNetworksParams {
	p := &ListNetworksParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetworks(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.Networks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Networks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkByName(name string, opts ...OptionFunc) (*Network, int, error) {
	id, count, err := s.GetNetworkID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetNetworkByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkByID(id string, opts ...OptionFunc) (*Network, int, error) {
	p := &ListNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListNetworks(p)
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
		return l.Networks[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Network UUID: %s!", id)
}

// Lists all available networks.
func (s *NetworkService) ListNetworks(p *ListNetworksParams) (*ListNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworksResponse struct {
	Count    int        `json:"count"`
	Networks []*Network `json:"network"`
}

type Network struct {
	Account                     string                   `json:"account"`
	Aclid                       string                   `json:"aclid"`
	Aclname                     string                   `json:"aclname"`
	Acltype                     string                   `json:"acltype"`
	Broadcastdomaintype         string                   `json:"broadcastdomaintype"`
	Broadcasturi                string                   `json:"broadcasturi"`
	Canusefordeploy             bool                     `json:"canusefordeploy"`
	Cidr                        string                   `json:"cidr"`
	Created                     string                   `json:"created"`
	Details                     map[string]string        `json:"details"`
	Displaynetwork              bool                     `json:"displaynetwork"`
	Displaytext                 string                   `json:"displaytext"`
	Dns1                        string                   `json:"dns1"`
	Dns2                        string                   `json:"dns2"`
	Domain                      string                   `json:"domain"`
	Domainid                    string                   `json:"domainid"`
	Externalid                  string                   `json:"externalid"`
	Gateway                     string                   `json:"gateway"`
	Hasannotations              bool                     `json:"hasannotations"`
	Icon                        string                   `json:"icon"`
	Id                          string                   `json:"id"`
	Ip6cidr                     string                   `json:"ip6cidr"`
	Ip6gateway                  string                   `json:"ip6gateway"`
	Isdefault                   bool                     `json:"isdefault"`
	Ispersistent                bool                     `json:"ispersistent"`
	Issystem                    bool                     `json:"issystem"`
	JobID                       string                   `json:"jobid"`
	Jobstatus                   int                      `json:"jobstatus"`
	Name                        string                   `json:"name"`
	Netmask                     string                   `json:"netmask"`
	Networkcidr                 string                   `json:"networkcidr"`
	Networkdomain               string                   `json:"networkdomain"`
	Networkofferingavailability string                   `json:"networkofferingavailability"`
	Networkofferingconservemode bool                     `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                   `json:"networkofferingdisplaytext"`
	Networkofferingid           string                   `json:"networkofferingid"`
	Networkofferingname         string                   `json:"networkofferingname"`
	Physicalnetworkid           string                   `json:"physicalnetworkid"`
	Project                     string                   `json:"project"`
	Projectid                   string                   `json:"projectid"`
	Receivedbytes               int64                    `json:"receivedbytes"`
	Redundantrouter             bool                     `json:"redundantrouter"`
	Related                     string                   `json:"related"`
	Reservediprange             string                   `json:"reservediprange"`
	Restartrequired             bool                     `json:"restartrequired"`
	Sentbytes                   int64                    `json:"sentbytes"`
	Service                     []NetworkServiceInternal `json:"service"`
	Specifyipranges             bool                     `json:"specifyipranges"`
	State                       string                   `json:"state"`
	Strechedl2subnet            bool                     `json:"strechedl2subnet"`
	Subdomainaccess             bool                     `json:"subdomainaccess"`
	Tags                        []Tags                   `json:"tags"`
	Traffictype                 string                   `json:"traffictype"`
	Type                        string                   `json:"type"`
	Vlan                        string                   `json:"vlan"`
	Vpcid                       string                   `json:"vpcid"`
	Vpcname                     string                   `json:"vpcname"`
	Zoneid                      string                   `json:"zoneid"`
	Zonename                    string                   `json:"zonename"`
	Zonesnetworkspans           []interface{}            `json:"zonesnetworkspans"`
}

type NetworkServiceInternal struct {
	Capability []NetworkServiceInternalCapability `json:"capability"`
	Name       string                             `json:"name"`
	Provider   []NetworkServiceInternalProvider   `json:"provider"`
}

type NetworkServiceInternalProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NetworkServiceInternalCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListNiciraNvpDeviceNetworksParams struct {
	p map[string]interface{}
}

func (p *ListNiciraNvpDeviceNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListNiciraNvpDeviceNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNiciraNvpDeviceNetworksParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNiciraNvpDeviceNetworksParams) SetNvpdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nvpdeviceid"] = v
}

func (p *ListNiciraNvpDeviceNetworksParams) GetNvpdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nvpdeviceid"].(string)
	return value, ok
}

func (p *ListNiciraNvpDeviceNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNiciraNvpDeviceNetworksParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNiciraNvpDeviceNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNiciraNvpDeviceNetworksParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNiciraNvpDeviceNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNiciraNvpDeviceNetworksParams(nvpdeviceid string) *ListNiciraNvpDeviceNetworksParams {
	p := &ListNiciraNvpDeviceNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["nvpdeviceid"] = nvpdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNiciraNvpDeviceNetworkID(keyword string, nvpdeviceid string, opts ...OptionFunc) (string, int, error) {
	p := &ListNiciraNvpDeviceNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["nvpdeviceid"] = nvpdeviceid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNiciraNvpDeviceNetworks(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.NiciraNvpDeviceNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NiciraNvpDeviceNetworks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// lists network that are using a nicira nvp device
func (s *NetworkService) ListNiciraNvpDeviceNetworks(p *ListNiciraNvpDeviceNetworksParams) (*ListNiciraNvpDeviceNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNiciraNvpDeviceNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNiciraNvpDeviceNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNiciraNvpDeviceNetworksResponse struct {
	Count                   int                       `json:"count"`
	NiciraNvpDeviceNetworks []*NiciraNvpDeviceNetwork `json:"niciranvpdevicenetwork"`
}

type NiciraNvpDeviceNetwork struct {
	Account                     string                          `json:"account"`
	Aclid                       string                          `json:"aclid"`
	Aclname                     string                          `json:"aclname"`
	Acltype                     string                          `json:"acltype"`
	Broadcastdomaintype         string                          `json:"broadcastdomaintype"`
	Broadcasturi                string                          `json:"broadcasturi"`
	Canusefordeploy             bool                            `json:"canusefordeploy"`
	Cidr                        string                          `json:"cidr"`
	Created                     string                          `json:"created"`
	Details                     map[string]string               `json:"details"`
	Displaynetwork              bool                            `json:"displaynetwork"`
	Displaytext                 string                          `json:"displaytext"`
	Dns1                        string                          `json:"dns1"`
	Dns2                        string                          `json:"dns2"`
	Domain                      string                          `json:"domain"`
	Domainid                    string                          `json:"domainid"`
	Externalid                  string                          `json:"externalid"`
	Gateway                     string                          `json:"gateway"`
	Hasannotations              bool                            `json:"hasannotations"`
	Icon                        string                          `json:"icon"`
	Id                          string                          `json:"id"`
	Ip6cidr                     string                          `json:"ip6cidr"`
	Ip6gateway                  string                          `json:"ip6gateway"`
	Isdefault                   bool                            `json:"isdefault"`
	Ispersistent                bool                            `json:"ispersistent"`
	Issystem                    bool                            `json:"issystem"`
	JobID                       string                          `json:"jobid"`
	Jobstatus                   int                             `json:"jobstatus"`
	Name                        string                          `json:"name"`
	Netmask                     string                          `json:"netmask"`
	Networkcidr                 string                          `json:"networkcidr"`
	Networkdomain               string                          `json:"networkdomain"`
	Networkofferingavailability string                          `json:"networkofferingavailability"`
	Networkofferingconservemode bool                            `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                          `json:"networkofferingdisplaytext"`
	Networkofferingid           string                          `json:"networkofferingid"`
	Networkofferingname         string                          `json:"networkofferingname"`
	Physicalnetworkid           string                          `json:"physicalnetworkid"`
	Project                     string                          `json:"project"`
	Projectid                   string                          `json:"projectid"`
	Receivedbytes               int64                           `json:"receivedbytes"`
	Redundantrouter             bool                            `json:"redundantrouter"`
	Related                     string                          `json:"related"`
	Reservediprange             string                          `json:"reservediprange"`
	Restartrequired             bool                            `json:"restartrequired"`
	Sentbytes                   int64                           `json:"sentbytes"`
	Service                     []NiciraNvpDeviceNetworkService `json:"service"`
	Specifyipranges             bool                            `json:"specifyipranges"`
	State                       string                          `json:"state"`
	Strechedl2subnet            bool                            `json:"strechedl2subnet"`
	Subdomainaccess             bool                            `json:"subdomainaccess"`
	Tags                        []Tags                          `json:"tags"`
	Traffictype                 string                          `json:"traffictype"`
	Type                        string                          `json:"type"`
	Vlan                        string                          `json:"vlan"`
	Vpcid                       string                          `json:"vpcid"`
	Vpcname                     string                          `json:"vpcname"`
	Zoneid                      string                          `json:"zoneid"`
	Zonename                    string                          `json:"zonename"`
	Zonesnetworkspans           []interface{}                   `json:"zonesnetworkspans"`
}

type NiciraNvpDeviceNetworkService struct {
	Capability []NiciraNvpDeviceNetworkServiceCapability `json:"capability"`
	Name       string                                    `json:"name"`
	Provider   []NiciraNvpDeviceNetworkServiceProvider   `json:"provider"`
}

type NiciraNvpDeviceNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NiciraNvpDeviceNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListOpenDaylightControllersParams struct {
	p map[string]interface{}
}

func (p *ListOpenDaylightControllersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListOpenDaylightControllersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListOpenDaylightControllersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListOpenDaylightControllersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ListOpenDaylightControllersParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListOpenDaylightControllersParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListOpenDaylightControllersParams() *ListOpenDaylightControllersParams {
	p := &ListOpenDaylightControllersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetOpenDaylightControllerByID(id string, opts ...OptionFunc) (*OpenDaylightController, int, error) {
	p := &ListOpenDaylightControllersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOpenDaylightControllers(p)
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
		return l.OpenDaylightControllers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for OpenDaylightController UUID: %s!", id)
}

// Lists OpenDyalight controllers
func (s *NetworkService) ListOpenDaylightControllers(p *ListOpenDaylightControllersParams) (*ListOpenDaylightControllersResponse, error) {
	resp, err := s.cs.newRequest("listOpenDaylightControllers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListOpenDaylightControllersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListOpenDaylightControllersResponse struct {
	Count                   int                       `json:"count"`
	OpenDaylightControllers []*OpenDaylightController `json:"opendaylightcontroller"`
}

type OpenDaylightController struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Name              string `json:"name"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Url               string `json:"url"`
	Username          string `json:"username"`
}

type ListPaloAltoFirewallNetworksParams struct {
	p map[string]interface{}
}

func (p *ListPaloAltoFirewallNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListPaloAltoFirewallNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListPaloAltoFirewallNetworksParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListPaloAltoFirewallNetworksParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
}

func (p *ListPaloAltoFirewallNetworksParams) GetLbdeviceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbdeviceid"].(string)
	return value, ok
}

func (p *ListPaloAltoFirewallNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListPaloAltoFirewallNetworksParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListPaloAltoFirewallNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListPaloAltoFirewallNetworksParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListPaloAltoFirewallNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListPaloAltoFirewallNetworksParams(lbdeviceid string) *ListPaloAltoFirewallNetworksParams {
	p := &ListPaloAltoFirewallNetworksParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPaloAltoFirewallNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error) {
	p := &ListPaloAltoFirewallNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["lbdeviceid"] = lbdeviceid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListPaloAltoFirewallNetworks(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.PaloAltoFirewallNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.PaloAltoFirewallNetworks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// lists network that are using Palo Alto firewall device
func (s *NetworkService) ListPaloAltoFirewallNetworks(p *ListPaloAltoFirewallNetworksParams) (*ListPaloAltoFirewallNetworksResponse, error) {
	resp, err := s.cs.newRequest("listPaloAltoFirewallNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPaloAltoFirewallNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPaloAltoFirewallNetworksResponse struct {
	Count                    int                        `json:"count"`
	PaloAltoFirewallNetworks []*PaloAltoFirewallNetwork `json:"paloaltofirewallnetwork"`
}

type PaloAltoFirewallNetwork struct {
	Account                     string                           `json:"account"`
	Aclid                       string                           `json:"aclid"`
	Aclname                     string                           `json:"aclname"`
	Acltype                     string                           `json:"acltype"`
	Broadcastdomaintype         string                           `json:"broadcastdomaintype"`
	Broadcasturi                string                           `json:"broadcasturi"`
	Canusefordeploy             bool                             `json:"canusefordeploy"`
	Cidr                        string                           `json:"cidr"`
	Created                     string                           `json:"created"`
	Details                     map[string]string                `json:"details"`
	Displaynetwork              bool                             `json:"displaynetwork"`
	Displaytext                 string                           `json:"displaytext"`
	Dns1                        string                           `json:"dns1"`
	Dns2                        string                           `json:"dns2"`
	Domain                      string                           `json:"domain"`
	Domainid                    string                           `json:"domainid"`
	Externalid                  string                           `json:"externalid"`
	Gateway                     string                           `json:"gateway"`
	Hasannotations              bool                             `json:"hasannotations"`
	Icon                        string                           `json:"icon"`
	Id                          string                           `json:"id"`
	Ip6cidr                     string                           `json:"ip6cidr"`
	Ip6gateway                  string                           `json:"ip6gateway"`
	Isdefault                   bool                             `json:"isdefault"`
	Ispersistent                bool                             `json:"ispersistent"`
	Issystem                    bool                             `json:"issystem"`
	JobID                       string                           `json:"jobid"`
	Jobstatus                   int                              `json:"jobstatus"`
	Name                        string                           `json:"name"`
	Netmask                     string                           `json:"netmask"`
	Networkcidr                 string                           `json:"networkcidr"`
	Networkdomain               string                           `json:"networkdomain"`
	Networkofferingavailability string                           `json:"networkofferingavailability"`
	Networkofferingconservemode bool                             `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                           `json:"networkofferingdisplaytext"`
	Networkofferingid           string                           `json:"networkofferingid"`
	Networkofferingname         string                           `json:"networkofferingname"`
	Physicalnetworkid           string                           `json:"physicalnetworkid"`
	Project                     string                           `json:"project"`
	Projectid                   string                           `json:"projectid"`
	Receivedbytes               int64                            `json:"receivedbytes"`
	Redundantrouter             bool                             `json:"redundantrouter"`
	Related                     string                           `json:"related"`
	Reservediprange             string                           `json:"reservediprange"`
	Restartrequired             bool                             `json:"restartrequired"`
	Sentbytes                   int64                            `json:"sentbytes"`
	Service                     []PaloAltoFirewallNetworkService `json:"service"`
	Specifyipranges             bool                             `json:"specifyipranges"`
	State                       string                           `json:"state"`
	Strechedl2subnet            bool                             `json:"strechedl2subnet"`
	Subdomainaccess             bool                             `json:"subdomainaccess"`
	Tags                        []Tags                           `json:"tags"`
	Traffictype                 string                           `json:"traffictype"`
	Type                        string                           `json:"type"`
	Vlan                        string                           `json:"vlan"`
	Vpcid                       string                           `json:"vpcid"`
	Vpcname                     string                           `json:"vpcname"`
	Zoneid                      string                           `json:"zoneid"`
	Zonename                    string                           `json:"zonename"`
	Zonesnetworkspans           []interface{}                    `json:"zonesnetworkspans"`
}

type PaloAltoFirewallNetworkService struct {
	Capability []PaloAltoFirewallNetworkServiceCapability `json:"capability"`
	Name       string                                     `json:"name"`
	Provider   []PaloAltoFirewallNetworkServiceProvider   `json:"provider"`
}

type PaloAltoFirewallNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type PaloAltoFirewallNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListPhysicalNetworksParams struct {
	p map[string]interface{}
}

func (p *ListPhysicalNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListPhysicalNetworksParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListPhysicalNetworksParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListPhysicalNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListPhysicalNetworksParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListPhysicalNetworksParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListPhysicalNetworksParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListPhysicalNetworksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListPhysicalNetworksParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListPhysicalNetworksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListPhysicalNetworksParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListPhysicalNetworksParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListPhysicalNetworksParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPhysicalNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListPhysicalNetworksParams() *ListPhysicalNetworksParams {
	p := &ListPhysicalNetworksParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPhysicalNetworkID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListPhysicalNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListPhysicalNetworks(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.PhysicalNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.PhysicalNetworks {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPhysicalNetworkByName(name string, opts ...OptionFunc) (*PhysicalNetwork, int, error) {
	id, count, err := s.GetPhysicalNetworkID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetPhysicalNetworkByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPhysicalNetworkByID(id string, opts ...OptionFunc) (*PhysicalNetwork, int, error) {
	p := &ListPhysicalNetworksParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPhysicalNetworks(p)
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
		return l.PhysicalNetworks[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for PhysicalNetwork UUID: %s!", id)
}

// Lists physical networks
func (s *NetworkService) ListPhysicalNetworks(p *ListPhysicalNetworksParams) (*ListPhysicalNetworksResponse, error) {
	resp, err := s.cs.newRequest("listPhysicalNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPhysicalNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPhysicalNetworksResponse struct {
	Count            int                `json:"count"`
	PhysicalNetworks []*PhysicalNetwork `json:"physicalnetwork"`
}

type PhysicalNetwork struct {
	Broadcastdomainrange string `json:"broadcastdomainrange"`
	Domainid             string `json:"domainid"`
	Id                   string `json:"id"`
	Isolationmethods     string `json:"isolationmethods"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Name                 string `json:"name"`
	Networkspeed         string `json:"networkspeed"`
	State                string `json:"state"`
	Tags                 string `json:"tags"`
	Vlan                 string `json:"vlan"`
	Zoneid               string `json:"zoneid"`
	Zonename             string `json:"zonename"`
}

type ListStorageNetworkIpRangeParams struct {
	p map[string]interface{}
}

func (p *ListStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListStorageNetworkIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListStorageNetworkIpRangeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListStorageNetworkIpRangeParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStorageNetworkIpRangeParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStorageNetworkIpRangeParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStorageNetworkIpRangeParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStorageNetworkIpRangeParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStorageNetworkIpRangeParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListStorageNetworkIpRangeParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListStorageNetworkIpRangeParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListStorageNetworkIpRangeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListStorageNetworkIpRangeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListStorageNetworkIpRangeParams() *ListStorageNetworkIpRangeParams {
	p := &ListStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetStorageNetworkIpRangeByID(id string, opts ...OptionFunc) (*StorageNetworkIpRange, int, error) {
	p := &ListStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStorageNetworkIpRange(p)
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
		return l.StorageNetworkIpRange[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for StorageNetworkIpRange UUID: %s!", id)
}

// List a storage network IP range.
func (s *NetworkService) ListStorageNetworkIpRange(p *ListStorageNetworkIpRangeParams) (*ListStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("listStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStorageNetworkIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStorageNetworkIpRangeResponse struct {
	Count                 int                      `json:"count"`
	StorageNetworkIpRange []*StorageNetworkIpRange `json:"storagenetworkiprange"`
}

type StorageNetworkIpRange struct {
	Endip     string `json:"endip"`
	Gateway   string `json:"gateway"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Netmask   string `json:"netmask"`
	Networkid string `json:"networkid"`
	Podid     string `json:"podid"`
	Startip   string `json:"startip"`
	Vlan      int    `json:"vlan"`
	Zoneid    string `json:"zoneid"`
}

type ListSupportedNetworkServicesParams struct {
	p map[string]interface{}
}

func (p *ListSupportedNetworkServicesParams) toURLValues() url.Values {
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
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["service"]; found {
		u.Set("service", v.(string))
	}
	return u
}

func (p *ListSupportedNetworkServicesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSupportedNetworkServicesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSupportedNetworkServicesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSupportedNetworkServicesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSupportedNetworkServicesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSupportedNetworkServicesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListSupportedNetworkServicesParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListSupportedNetworkServicesParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *ListSupportedNetworkServicesParams) SetService(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["service"] = v
}

func (p *ListSupportedNetworkServicesParams) GetService() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["service"].(string)
	return value, ok
}

// You should always use this function to get a new ListSupportedNetworkServicesParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListSupportedNetworkServicesParams() *ListSupportedNetworkServicesParams {
	p := &ListSupportedNetworkServicesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all network services provided by CloudStack or for the given Provider.
func (s *NetworkService) ListSupportedNetworkServices(p *ListSupportedNetworkServicesParams) (*ListSupportedNetworkServicesResponse, error) {
	resp, err := s.cs.newRequest("listSupportedNetworkServices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSupportedNetworkServicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSupportedNetworkServicesResponse struct {
	Count                    int                        `json:"count"`
	SupportedNetworkServices []*SupportedNetworkService `json:"supportednetworkservice"`
}

type SupportedNetworkService struct {
	Capability []SupportedNetworkServiceCapability `json:"capability"`
	JobID      string                              `json:"jobid"`
	Jobstatus  int                                 `json:"jobstatus"`
	Name       string                              `json:"name"`
	Provider   []SupportedNetworkServiceProvider   `json:"provider"`
}

type SupportedNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type SupportedNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ReleasePublicIpRangeParams struct {
	p map[string]interface{}
}

func (p *ReleasePublicIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ReleasePublicIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ReleasePublicIpRangeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReleasePublicIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewReleasePublicIpRangeParams(id string) *ReleasePublicIpRangeParams {
	p := &ReleasePublicIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Releases a Public IP range back to the system pool
func (s *NetworkService) ReleasePublicIpRange(p *ReleasePublicIpRangeParams) (*ReleasePublicIpRangeResponse, error) {
	resp, err := s.cs.newRequest("releasePublicIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleasePublicIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ReleasePublicIpRangeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ReleasePublicIpRangeResponse) UnmarshalJSON(b []byte) error {
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

	type alias ReleasePublicIpRangeResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RestartNetworkParams struct {
	p map[string]interface{}
}

func (p *RestartNetworkParams) toURLValues() url.Values {
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
	if v, found := p.p["makeredundant"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("makeredundant", vv)
	}
	return u
}

func (p *RestartNetworkParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *RestartNetworkParams) GetCleanup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanup"].(bool)
	return value, ok
}

func (p *RestartNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RestartNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *RestartNetworkParams) SetMakeredundant(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["makeredundant"] = v
}

func (p *RestartNetworkParams) GetMakeredundant() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["makeredundant"].(bool)
	return value, ok
}

// You should always use this function to get a new RestartNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewRestartNetworkParams(id string) *RestartNetworkParams {
	p := &RestartNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Restarts the network; includes 1) restarting network elements - virtual routers, DHCP servers 2) reapplying all public IPs 3) reapplying loadBalancing/portForwarding rules
func (s *NetworkService) RestartNetwork(p *RestartNetworkParams) (*RestartNetworkResponse, error) {
	resp, err := s.cs.newRequest("restartNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestartNetworkResponse
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

type RestartNetworkResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateNetworkParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["changecidr"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("changecidr", vv)
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["guestvmcidr"]; found {
		u.Set("guestvmcidr", v.(string))
	}
	if v, found := p.p["hideipaddressusage"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hideipaddressusage", vv)
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
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := p.p["updateinsequence"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("updateinsequence", vv)
	}
	return u
}

func (p *UpdateNetworkParams) SetChangecidr(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["changecidr"] = v
}

func (p *UpdateNetworkParams) GetChangecidr() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["changecidr"].(bool)
	return value, ok
}

func (p *UpdateNetworkParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateNetworkParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetDisplaynetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaynetwork"] = v
}

func (p *UpdateNetworkParams) GetDisplaynetwork() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaynetwork"].(bool)
	return value, ok
}

func (p *UpdateNetworkParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateNetworkParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *UpdateNetworkParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *UpdateNetworkParams) SetGuestvmcidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestvmcidr"] = v
}

func (p *UpdateNetworkParams) GetGuestvmcidr() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["guestvmcidr"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetHideipaddressusage(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hideipaddressusage"] = v
}

func (p *UpdateNetworkParams) GetHideipaddressusage() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hideipaddressusage"].(bool)
	return value, ok
}

func (p *UpdateNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateNetworkParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetNetworkdomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdomain"] = v
}

func (p *UpdateNetworkParams) GetNetworkdomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkdomain"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
}

func (p *UpdateNetworkParams) GetNetworkofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkofferingid"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetUpdateinsequence(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["updateinsequence"] = v
}

func (p *UpdateNetworkParams) GetUpdateinsequence() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["updateinsequence"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateNetworkParams(id string) *UpdateNetworkParams {
	p := &UpdateNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a network
func (s *NetworkService) UpdateNetwork(p *UpdateNetworkParams) (*UpdateNetworkResponse, error) {
	resp, err := s.cs.newRequest("updateNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkResponse
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

type UpdateNetworkResponse struct {
	Account                     string                         `json:"account"`
	Aclid                       string                         `json:"aclid"`
	Aclname                     string                         `json:"aclname"`
	Acltype                     string                         `json:"acltype"`
	Broadcastdomaintype         string                         `json:"broadcastdomaintype"`
	Broadcasturi                string                         `json:"broadcasturi"`
	Canusefordeploy             bool                           `json:"canusefordeploy"`
	Cidr                        string                         `json:"cidr"`
	Created                     string                         `json:"created"`
	Details                     map[string]string              `json:"details"`
	Displaynetwork              bool                           `json:"displaynetwork"`
	Displaytext                 string                         `json:"displaytext"`
	Dns1                        string                         `json:"dns1"`
	Dns2                        string                         `json:"dns2"`
	Domain                      string                         `json:"domain"`
	Domainid                    string                         `json:"domainid"`
	Externalid                  string                         `json:"externalid"`
	Gateway                     string                         `json:"gateway"`
	Hasannotations              bool                           `json:"hasannotations"`
	Icon                        string                         `json:"icon"`
	Id                          string                         `json:"id"`
	Ip6cidr                     string                         `json:"ip6cidr"`
	Ip6gateway                  string                         `json:"ip6gateway"`
	Isdefault                   bool                           `json:"isdefault"`
	Ispersistent                bool                           `json:"ispersistent"`
	Issystem                    bool                           `json:"issystem"`
	JobID                       string                         `json:"jobid"`
	Jobstatus                   int                            `json:"jobstatus"`
	Name                        string                         `json:"name"`
	Netmask                     string                         `json:"netmask"`
	Networkcidr                 string                         `json:"networkcidr"`
	Networkdomain               string                         `json:"networkdomain"`
	Networkofferingavailability string                         `json:"networkofferingavailability"`
	Networkofferingconservemode bool                           `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                         `json:"networkofferingdisplaytext"`
	Networkofferingid           string                         `json:"networkofferingid"`
	Networkofferingname         string                         `json:"networkofferingname"`
	Physicalnetworkid           string                         `json:"physicalnetworkid"`
	Project                     string                         `json:"project"`
	Projectid                   string                         `json:"projectid"`
	Receivedbytes               int64                          `json:"receivedbytes"`
	Redundantrouter             bool                           `json:"redundantrouter"`
	Related                     string                         `json:"related"`
	Reservediprange             string                         `json:"reservediprange"`
	Restartrequired             bool                           `json:"restartrequired"`
	Sentbytes                   int64                          `json:"sentbytes"`
	Service                     []UpdateNetworkResponseService `json:"service"`
	Specifyipranges             bool                           `json:"specifyipranges"`
	State                       string                         `json:"state"`
	Strechedl2subnet            bool                           `json:"strechedl2subnet"`
	Subdomainaccess             bool                           `json:"subdomainaccess"`
	Tags                        []Tags                         `json:"tags"`
	Traffictype                 string                         `json:"traffictype"`
	Type                        string                         `json:"type"`
	Vlan                        string                         `json:"vlan"`
	Vpcid                       string                         `json:"vpcid"`
	Vpcname                     string                         `json:"vpcname"`
	Zoneid                      string                         `json:"zoneid"`
	Zonename                    string                         `json:"zonename"`
	Zonesnetworkspans           []interface{}                  `json:"zonesnetworkspans"`
}

type UpdateNetworkResponseService struct {
	Capability []UpdateNetworkResponseServiceCapability `json:"capability"`
	Name       string                                   `json:"name"`
	Provider   []UpdateNetworkResponseServiceProvider   `json:"provider"`
}

type UpdateNetworkResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type UpdateNetworkResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type UpdateNetworkServiceProviderParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["servicelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("servicelist", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *UpdateNetworkServiceProviderParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateNetworkServiceProviderParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateNetworkServiceProviderParams) SetServicelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicelist"] = v
}

func (p *UpdateNetworkServiceProviderParams) GetServicelist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["servicelist"].([]string)
	return value, ok
}

func (p *UpdateNetworkServiceProviderParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdateNetworkServiceProviderParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateNetworkServiceProviderParams(id string) *UpdateNetworkServiceProviderParams {
	p := &UpdateNetworkServiceProviderParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a network serviceProvider of a physical network
func (s *NetworkService) UpdateNetworkServiceProvider(p *UpdateNetworkServiceProviderParams) (*UpdateNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkServiceProviderResponse
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

type UpdateNetworkServiceProviderResponse struct {
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

type UpdatePhysicalNetworkParams struct {
	p map[string]interface{}
}

func (p *UpdatePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["networkspeed"]; found {
		u.Set("networkspeed", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["tags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("tags", vv)
	}
	if v, found := p.p["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	return u
}

func (p *UpdatePhysicalNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdatePhysicalNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdatePhysicalNetworkParams) SetNetworkspeed(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkspeed"] = v
}

func (p *UpdatePhysicalNetworkParams) GetNetworkspeed() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkspeed"].(string)
	return value, ok
}

func (p *UpdatePhysicalNetworkParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdatePhysicalNetworkParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *UpdatePhysicalNetworkParams) SetTags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *UpdatePhysicalNetworkParams) GetTags() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].([]string)
	return value, ok
}

func (p *UpdatePhysicalNetworkParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *UpdatePhysicalNetworkParams) GetVlan() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(string)
	return value, ok
}

// You should always use this function to get a new UpdatePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdatePhysicalNetworkParams(id string) *UpdatePhysicalNetworkParams {
	p := &UpdatePhysicalNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a physical network
func (s *NetworkService) UpdatePhysicalNetwork(p *UpdatePhysicalNetworkParams) (*UpdatePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("updatePhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdatePhysicalNetworkResponse
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

type UpdatePhysicalNetworkResponse struct {
	Broadcastdomainrange string `json:"broadcastdomainrange"`
	Domainid             string `json:"domainid"`
	Id                   string `json:"id"`
	Isolationmethods     string `json:"isolationmethods"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Name                 string `json:"name"`
	Networkspeed         string `json:"networkspeed"`
	State                string `json:"state"`
	Tags                 string `json:"tags"`
	Vlan                 string `json:"vlan"`
	Zoneid               string `json:"zoneid"`
	Zonename             string `json:"zonename"`
}

type UpdateStorageNetworkIpRangeParams struct {
	p map[string]interface{}
}

func (p *UpdateStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := p.p["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := p.p["vlan"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("vlan", vv)
	}
	return u
}

func (p *UpdateStorageNetworkIpRangeParams) SetEndip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endip"] = v
}

func (p *UpdateStorageNetworkIpRangeParams) GetEndip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endip"].(string)
	return value, ok
}

func (p *UpdateStorageNetworkIpRangeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateStorageNetworkIpRangeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateStorageNetworkIpRangeParams) SetNetmask(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["netmask"] = v
}

func (p *UpdateStorageNetworkIpRangeParams) GetNetmask() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["netmask"].(string)
	return value, ok
}

func (p *UpdateStorageNetworkIpRangeParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
}

func (p *UpdateStorageNetworkIpRangeParams) GetStartip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startip"].(string)
	return value, ok
}

func (p *UpdateStorageNetworkIpRangeParams) SetVlan(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *UpdateStorageNetworkIpRangeParams) GetVlan() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateStorageNetworkIpRangeParams(id string) *UpdateStorageNetworkIpRangeParams {
	p := &UpdateStorageNetworkIpRangeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Update a Storage network IP range, only allowed when no IPs in this range have been allocated.
func (s *NetworkService) UpdateStorageNetworkIpRange(p *UpdateStorageNetworkIpRangeParams) (*UpdateStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("updateStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateStorageNetworkIpRangeResponse
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

type UpdateStorageNetworkIpRangeResponse struct {
	Endip     string `json:"endip"`
	Gateway   string `json:"gateway"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Netmask   string `json:"netmask"`
	Networkid string `json:"networkid"`
	Podid     string `json:"podid"`
	Startip   string `json:"startip"`
	Vlan      int    `json:"vlan"`
	Zoneid    string `json:"zoneid"`
}
