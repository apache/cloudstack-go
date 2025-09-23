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
	ChangeBgpPeersForNetwork(p *ChangeBgpPeersForNetworkParams) (*ChangeBgpPeersForNetworkResponse, error)
	NewChangeBgpPeersForNetworkParams(networkid string) *ChangeBgpPeersForNetworkParams
	CreateIpv4SubnetForGuestNetwork(p *CreateIpv4SubnetForGuestNetworkParams) (*CreateIpv4SubnetForGuestNetworkResponse, error)
	NewCreateIpv4SubnetForGuestNetworkParams(parentid string) *CreateIpv4SubnetForGuestNetworkParams
	CreateNetwork(p *CreateNetworkParams) (*CreateNetworkResponse, error)
	NewCreateNetworkParams(name string, networkofferingid string, zoneid string) *CreateNetworkParams
	CreatePhysicalNetwork(p *CreatePhysicalNetworkParams) (*CreatePhysicalNetworkResponse, error)
	NewCreatePhysicalNetworkParams(name string, zoneid string) *CreatePhysicalNetworkParams
	CreateServiceInstance(p *CreateServiceInstanceParams) (*CreateServiceInstanceResponse, error)
	NewCreateServiceInstanceParams(leftnetworkid string, name string, rightnetworkid string, serviceofferingid string, templateid string, zoneid string) *CreateServiceInstanceParams
	CreateStorageNetworkIpRange(p *CreateStorageNetworkIpRangeParams) (*CreateStorageNetworkIpRangeResponse, error)
	NewCreateStorageNetworkIpRangeParams(gateway string, netmask string, podid string, startip string) *CreateStorageNetworkIpRangeParams
	DedicatePublicIpRange(p *DedicatePublicIpRangeParams) (*DedicatePublicIpRangeResponse, error)
	NewDedicatePublicIpRangeParams(domainid string, id string) *DedicatePublicIpRangeParams
	DeleteIpv4SubnetForGuestNetwork(p *DeleteIpv4SubnetForGuestNetworkParams) (*DeleteIpv4SubnetForGuestNetworkResponse, error)
	NewDeleteIpv4SubnetForGuestNetworkParams(id string) *DeleteIpv4SubnetForGuestNetworkParams
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
	ListIpv4SubnetsForGuestNetwork(p *ListIpv4SubnetsForGuestNetworkParams) (*ListIpv4SubnetsForGuestNetworkResponse, error)
	NewListIpv4SubnetsForGuestNetworkParams() *ListIpv4SubnetsForGuestNetworkParams
	GetIpv4SubnetsForGuestNetworkByID(id string, opts ...OptionFunc) (*Ipv4SubnetsForGuestNetwork, int, error)
	ListNetworkIsolationMethods(p *ListNetworkIsolationMethodsParams) (*ListNetworkIsolationMethodsResponse, error)
	NewListNetworkIsolationMethodsParams() *ListNetworkIsolationMethodsParams
	ListNetworkProtocols(p *ListNetworkProtocolsParams) (*ListNetworkProtocolsResponse, error)
	NewListNetworkProtocolsParams(option string) *ListNetworkProtocolsParams
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
	MigrateNetwork(p *MigrateNetworkParams) (*MigrateNetworkResponse, error)
	NewMigrateNetworkParams(networkid string, networkofferingid string) *MigrateNetworkParams
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
	DeleteGuestNetworkIpv6Prefix(p *DeleteGuestNetworkIpv6PrefixParams) (*DeleteGuestNetworkIpv6PrefixResponse, error)
	NewDeleteGuestNetworkIpv6PrefixParams(id string) *DeleteGuestNetworkIpv6PrefixParams
	CreateGuestNetworkIpv6Prefix(p *CreateGuestNetworkIpv6PrefixParams) (*CreateGuestNetworkIpv6PrefixResponse, error)
	NewCreateGuestNetworkIpv6PrefixParams(prefix string, zoneid string) *CreateGuestNetworkIpv6PrefixParams
	ListGuestNetworkIpv6Prefixes(p *ListGuestNetworkIpv6PrefixesParams) (*ListGuestNetworkIpv6PrefixesResponse, error)
	NewListGuestNetworkIpv6PrefixesParams() *ListGuestNetworkIpv6PrefixesParams
	GetGuestNetworkIpv6PrefixeByID(id string, opts ...OptionFunc) (*GuestNetworkIpv6Prefixe, int, error)
	CreateNetworkPermissions(p *CreateNetworkPermissionsParams) (*CreateNetworkPermissionsResponse, error)
	NewCreateNetworkPermissionsParams(networkid string) *CreateNetworkPermissionsParams
	ResetNetworkPermissions(p *ResetNetworkPermissionsParams) (*ResetNetworkPermissionsResponse, error)
	NewResetNetworkPermissionsParams(networkid string) *ResetNetworkPermissionsParams
	ListNetworkPermissions(p *ListNetworkPermissionsParams) (*ListNetworkPermissionsResponse, error)
	NewListNetworkPermissionsParams(networkid string) *ListNetworkPermissionsParams
	RemoveNetworkPermissions(p *RemoveNetworkPermissionsParams) (*RemoveNetworkPermissionsResponse, error)
	NewRemoveNetworkPermissionsParams(networkid string) *RemoveNetworkPermissionsParams
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

func (p *AddNetworkServiceProviderParams) ResetDestinationphysicalnetworkid() {
	if p.p != nil && p.p["destinationphysicalnetworkid"] != nil {
		delete(p.p, "destinationphysicalnetworkid")
	}
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

func (p *AddNetworkServiceProviderParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *AddNetworkServiceProviderParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
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

func (p *AddNetworkServiceProviderParams) ResetServicelist() {
	if p.p != nil && p.p["servicelist"] != nil {
		delete(p.p, "servicelist")
	}
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
	resp, err := s.cs.newPostRequest("addNetworkServiceProvider", p.toURLValues())
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

func (p *AddOpenDaylightControllerParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
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

func (p *AddOpenDaylightControllerParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
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

func (p *AddOpenDaylightControllerParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
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

func (p *AddOpenDaylightControllerParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
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
	resp, err := s.cs.newPostRequest("addOpenDaylightController", p.toURLValues())
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

type ChangeBgpPeersForNetworkParams struct {
	p map[string]interface{}
}

func (p *ChangeBgpPeersForNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bgppeerids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("bgppeerids", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	return u
}

func (p *ChangeBgpPeersForNetworkParams) SetBgppeerids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bgppeerids"] = v
}

func (p *ChangeBgpPeersForNetworkParams) ResetBgppeerids() {
	if p.p != nil && p.p["bgppeerids"] != nil {
		delete(p.p, "bgppeerids")
	}
}

func (p *ChangeBgpPeersForNetworkParams) GetBgppeerids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bgppeerids"].([]string)
	return value, ok
}

func (p *ChangeBgpPeersForNetworkParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ChangeBgpPeersForNetworkParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ChangeBgpPeersForNetworkParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeBgpPeersForNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewChangeBgpPeersForNetworkParams(networkid string) *ChangeBgpPeersForNetworkParams {
	p := &ChangeBgpPeersForNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	return p
}

// Change the BGP peers for a network.
func (s *NetworkService) ChangeBgpPeersForNetwork(p *ChangeBgpPeersForNetworkParams) (*ChangeBgpPeersForNetworkResponse, error) {
	resp, err := s.cs.newPostRequest("changeBgpPeersForNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeBgpPeersForNetworkResponse
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

type ChangeBgpPeersForNetworkResponse struct {
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

type CreateIpv4SubnetForGuestNetworkParams struct {
	p map[string]interface{}
}

func (p *CreateIpv4SubnetForGuestNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cidrsize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cidrsize", vv)
	}
	if v, found := p.p["parentid"]; found {
		u.Set("parentid", v.(string))
	}
	if v, found := p.p["subnet"]; found {
		u.Set("subnet", v.(string))
	}
	return u
}

func (p *CreateIpv4SubnetForGuestNetworkParams) SetCidrsize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrsize"] = v
}

func (p *CreateIpv4SubnetForGuestNetworkParams) ResetCidrsize() {
	if p.p != nil && p.p["cidrsize"] != nil {
		delete(p.p, "cidrsize")
	}
}

func (p *CreateIpv4SubnetForGuestNetworkParams) GetCidrsize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrsize"].(int)
	return value, ok
}

func (p *CreateIpv4SubnetForGuestNetworkParams) SetParentid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parentid"] = v
}

func (p *CreateIpv4SubnetForGuestNetworkParams) ResetParentid() {
	if p.p != nil && p.p["parentid"] != nil {
		delete(p.p, "parentid")
	}
}

func (p *CreateIpv4SubnetForGuestNetworkParams) GetParentid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parentid"].(string)
	return value, ok
}

func (p *CreateIpv4SubnetForGuestNetworkParams) SetSubnet(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["subnet"] = v
}

func (p *CreateIpv4SubnetForGuestNetworkParams) ResetSubnet() {
	if p.p != nil && p.p["subnet"] != nil {
		delete(p.p, "subnet")
	}
}

func (p *CreateIpv4SubnetForGuestNetworkParams) GetSubnet() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["subnet"].(string)
	return value, ok
}

// You should always use this function to get a new CreateIpv4SubnetForGuestNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateIpv4SubnetForGuestNetworkParams(parentid string) *CreateIpv4SubnetForGuestNetworkParams {
	p := &CreateIpv4SubnetForGuestNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["parentid"] = parentid
	return p
}

// Creates a IPv4 subnet for guest networks.
func (s *NetworkService) CreateIpv4SubnetForGuestNetwork(p *CreateIpv4SubnetForGuestNetworkParams) (*CreateIpv4SubnetForGuestNetworkResponse, error) {
	resp, err := s.cs.newPostRequest("createIpv4SubnetForGuestNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateIpv4SubnetForGuestNetworkResponse
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

type CreateIpv4SubnetForGuestNetworkResponse struct {
	Allocated    string `json:"allocated"`
	Created      string `json:"created"`
	Id           string `json:"id"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Networkname  string `json:"networkname"`
	Parentid     string `json:"parentid"`
	Parentsubnet string `json:"parentsubnet"`
	Removed      string `json:"removed"`
	State        string `json:"state"`
	Subnet       string `json:"subnet"`
	Vpcid        string `json:"vpcid"`
	Vpcname      string `json:"vpcname"`
	Zoneid       string `json:"zoneid"`
	Zonename     string `json:"zonename"`
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
	if v, found := p.p["asnumber"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("asnumber", vv)
	}
	if v, found := p.p["associatednetworkid"]; found {
		u.Set("associatednetworkid", v.(string))
	}
	if v, found := p.p["bgppeerids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("bgppeerids", vv)
	}
	if v, found := p.p["bypassvlanoverlapcheck"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bypassvlanoverlapcheck", vv)
	}
	if v, found := p.p["cidrsize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cidrsize", vv)
	}
	if v, found := p.p["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
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
	if v, found := p.p["ip6dns1"]; found {
		u.Set("ip6dns1", v.(string))
	}
	if v, found := p.p["ip6dns2"]; found {
		u.Set("ip6dns2", v.(string))
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
	if v, found := p.p["privatemtu"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privatemtu", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["publicmtu"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicmtu", vv)
	}
	if v, found := p.p["routerip"]; found {
		u.Set("routerip", v.(string))
	}
	if v, found := p.p["routeripv6"]; found {
		u.Set("routeripv6", v.(string))
	}
	if v, found := p.p["sourcenatipaddress"]; found {
		u.Set("sourcenatipaddress", v.(string))
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
	if v, found := p.p["tungstenvirtualrouteruuid"]; found {
		u.Set("tungstenvirtualrouteruuid", v.(string))
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

func (p *CreateNetworkParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
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

func (p *CreateNetworkParams) ResetAclid() {
	if p.p != nil && p.p["aclid"] != nil {
		delete(p.p, "aclid")
	}
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

func (p *CreateNetworkParams) ResetAcltype() {
	if p.p != nil && p.p["acltype"] != nil {
		delete(p.p, "acltype")
	}
}

func (p *CreateNetworkParams) GetAcltype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["acltype"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetAsnumber(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asnumber"] = v
}

func (p *CreateNetworkParams) ResetAsnumber() {
	if p.p != nil && p.p["asnumber"] != nil {
		delete(p.p, "asnumber")
	}
}

func (p *CreateNetworkParams) GetAsnumber() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asnumber"].(int64)
	return value, ok
}

func (p *CreateNetworkParams) SetAssociatednetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["associatednetworkid"] = v
}

func (p *CreateNetworkParams) ResetAssociatednetworkid() {
	if p.p != nil && p.p["associatednetworkid"] != nil {
		delete(p.p, "associatednetworkid")
	}
}

func (p *CreateNetworkParams) GetAssociatednetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["associatednetworkid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetBgppeerids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bgppeerids"] = v
}

func (p *CreateNetworkParams) ResetBgppeerids() {
	if p.p != nil && p.p["bgppeerids"] != nil {
		delete(p.p, "bgppeerids")
	}
}

func (p *CreateNetworkParams) GetBgppeerids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bgppeerids"].([]string)
	return value, ok
}

func (p *CreateNetworkParams) SetBypassvlanoverlapcheck(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bypassvlanoverlapcheck"] = v
}

func (p *CreateNetworkParams) ResetBypassvlanoverlapcheck() {
	if p.p != nil && p.p["bypassvlanoverlapcheck"] != nil {
		delete(p.p, "bypassvlanoverlapcheck")
	}
}

func (p *CreateNetworkParams) GetBypassvlanoverlapcheck() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bypassvlanoverlapcheck"].(bool)
	return value, ok
}

func (p *CreateNetworkParams) SetCidrsize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrsize"] = v
}

func (p *CreateNetworkParams) ResetCidrsize() {
	if p.p != nil && p.p["cidrsize"] != nil {
		delete(p.p, "cidrsize")
	}
}

func (p *CreateNetworkParams) GetCidrsize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cidrsize"].(int)
	return value, ok
}

func (p *CreateNetworkParams) SetDisplaynetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaynetwork"] = v
}

func (p *CreateNetworkParams) ResetDisplaynetwork() {
	if p.p != nil && p.p["displaynetwork"] != nil {
		delete(p.p, "displaynetwork")
	}
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

func (p *CreateNetworkParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *CreateNetworkParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetDns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns1"] = v
}

func (p *CreateNetworkParams) ResetDns1() {
	if p.p != nil && p.p["dns1"] != nil {
		delete(p.p, "dns1")
	}
}

func (p *CreateNetworkParams) GetDns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns1"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetDns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns2"] = v
}

func (p *CreateNetworkParams) ResetDns2() {
	if p.p != nil && p.p["dns2"] != nil {
		delete(p.p, "dns2")
	}
}

func (p *CreateNetworkParams) GetDns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns2"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateNetworkParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *CreateNetworkParams) ResetEndip() {
	if p.p != nil && p.p["endip"] != nil {
		delete(p.p, "endip")
	}
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

func (p *CreateNetworkParams) ResetEndipv6() {
	if p.p != nil && p.p["endipv6"] != nil {
		delete(p.p, "endipv6")
	}
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

func (p *CreateNetworkParams) ResetExternalid() {
	if p.p != nil && p.p["externalid"] != nil {
		delete(p.p, "externalid")
	}
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

func (p *CreateNetworkParams) ResetGateway() {
	if p.p != nil && p.p["gateway"] != nil {
		delete(p.p, "gateway")
	}
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

func (p *CreateNetworkParams) ResetHideipaddressusage() {
	if p.p != nil && p.p["hideipaddressusage"] != nil {
		delete(p.p, "hideipaddressusage")
	}
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

func (p *CreateNetworkParams) ResetIp6cidr() {
	if p.p != nil && p.p["ip6cidr"] != nil {
		delete(p.p, "ip6cidr")
	}
}

func (p *CreateNetworkParams) GetIp6cidr() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6cidr"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetIp6dns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns1"] = v
}

func (p *CreateNetworkParams) ResetIp6dns1() {
	if p.p != nil && p.p["ip6dns1"] != nil {
		delete(p.p, "ip6dns1")
	}
}

func (p *CreateNetworkParams) GetIp6dns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns1"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetIp6dns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns2"] = v
}

func (p *CreateNetworkParams) ResetIp6dns2() {
	if p.p != nil && p.p["ip6dns2"] != nil {
		delete(p.p, "ip6dns2")
	}
}

func (p *CreateNetworkParams) GetIp6dns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns2"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetIp6gateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6gateway"] = v
}

func (p *CreateNetworkParams) ResetIp6gateway() {
	if p.p != nil && p.p["ip6gateway"] != nil {
		delete(p.p, "ip6gateway")
	}
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

func (p *CreateNetworkParams) ResetIsolatedpvlan() {
	if p.p != nil && p.p["isolatedpvlan"] != nil {
		delete(p.p, "isolatedpvlan")
	}
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

func (p *CreateNetworkParams) ResetIsolatedpvlantype() {
	if p.p != nil && p.p["isolatedpvlantype"] != nil {
		delete(p.p, "isolatedpvlantype")
	}
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

func (p *CreateNetworkParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *CreateNetworkParams) ResetNetmask() {
	if p.p != nil && p.p["netmask"] != nil {
		delete(p.p, "netmask")
	}
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

func (p *CreateNetworkParams) ResetNetworkdomain() {
	if p.p != nil && p.p["networkdomain"] != nil {
		delete(p.p, "networkdomain")
	}
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

func (p *CreateNetworkParams) ResetNetworkofferingid() {
	if p.p != nil && p.p["networkofferingid"] != nil {
		delete(p.p, "networkofferingid")
	}
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

func (p *CreateNetworkParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *CreateNetworkParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetPrivatemtu(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privatemtu"] = v
}

func (p *CreateNetworkParams) ResetPrivatemtu() {
	if p.p != nil && p.p["privatemtu"] != nil {
		delete(p.p, "privatemtu")
	}
}

func (p *CreateNetworkParams) GetPrivatemtu() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privatemtu"].(int)
	return value, ok
}

func (p *CreateNetworkParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateNetworkParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateNetworkParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetPublicmtu(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicmtu"] = v
}

func (p *CreateNetworkParams) ResetPublicmtu() {
	if p.p != nil && p.p["publicmtu"] != nil {
		delete(p.p, "publicmtu")
	}
}

func (p *CreateNetworkParams) GetPublicmtu() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicmtu"].(int)
	return value, ok
}

func (p *CreateNetworkParams) SetRouterip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["routerip"] = v
}

func (p *CreateNetworkParams) ResetRouterip() {
	if p.p != nil && p.p["routerip"] != nil {
		delete(p.p, "routerip")
	}
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

func (p *CreateNetworkParams) ResetRouteripv6() {
	if p.p != nil && p.p["routeripv6"] != nil {
		delete(p.p, "routeripv6")
	}
}

func (p *CreateNetworkParams) GetRouteripv6() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["routeripv6"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetSourcenatipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatipaddress"] = v
}

func (p *CreateNetworkParams) ResetSourcenatipaddress() {
	if p.p != nil && p.p["sourcenatipaddress"] != nil {
		delete(p.p, "sourcenatipaddress")
	}
}

func (p *CreateNetworkParams) GetSourcenatipaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcenatipaddress"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetStartip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startip"] = v
}

func (p *CreateNetworkParams) ResetStartip() {
	if p.p != nil && p.p["startip"] != nil {
		delete(p.p, "startip")
	}
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

func (p *CreateNetworkParams) ResetStartipv6() {
	if p.p != nil && p.p["startipv6"] != nil {
		delete(p.p, "startipv6")
	}
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

func (p *CreateNetworkParams) ResetSubdomainaccess() {
	if p.p != nil && p.p["subdomainaccess"] != nil {
		delete(p.p, "subdomainaccess")
	}
}

func (p *CreateNetworkParams) GetSubdomainaccess() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["subdomainaccess"].(bool)
	return value, ok
}

func (p *CreateNetworkParams) SetTungstenvirtualrouteruuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tungstenvirtualrouteruuid"] = v
}

func (p *CreateNetworkParams) ResetTungstenvirtualrouteruuid() {
	if p.p != nil && p.p["tungstenvirtualrouteruuid"] != nil {
		delete(p.p, "tungstenvirtualrouteruuid")
	}
}

func (p *CreateNetworkParams) GetTungstenvirtualrouteruuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tungstenvirtualrouteruuid"].(string)
	return value, ok
}

func (p *CreateNetworkParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *CreateNetworkParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
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

func (p *CreateNetworkParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
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

func (p *CreateNetworkParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
func (s *NetworkService) NewCreateNetworkParams(name string, networkofferingid string, zoneid string) *CreateNetworkParams {
	p := &CreateNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["networkofferingid"] = networkofferingid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a network
func (s *NetworkService) CreateNetwork(p *CreateNetworkParams) (*CreateNetworkResponse, error) {
	resp, err := s.cs.newPostRequest("createNetwork", p.toURLValues())
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
	Asnumber                    int64                          `json:"asnumber"`
	Asnumberid                  string                         `json:"asnumberid"`
	Associatednetwork           string                         `json:"associatednetwork"`
	Associatednetworkid         string                         `json:"associatednetworkid"`
	Bgppeers                    []interface{}                  `json:"bgppeers"`
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
	Domainpath                  string                         `json:"domainpath"`
	Egressdefaultpolicy         bool                           `json:"egressdefaultpolicy"`
	Externalid                  string                         `json:"externalid"`
	Gateway                     string                         `json:"gateway"`
	Hasannotations              bool                           `json:"hasannotations"`
	Icon                        interface{}                    `json:"icon"`
	Id                          string                         `json:"id"`
	Internetprotocol            string                         `json:"internetprotocol"`
	Ip4routes                   []interface{}                  `json:"ip4routes"`
	Ip4routing                  string                         `json:"ip4routing"`
	Ip6cidr                     string                         `json:"ip6cidr"`
	Ip6dns1                     string                         `json:"ip6dns1"`
	Ip6dns2                     string                         `json:"ip6dns2"`
	Ip6gateway                  string                         `json:"ip6gateway"`
	Ip6routes                   []interface{}                  `json:"ip6routes"`
	Ip6routing                  string                         `json:"ip6routing"`
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
	Privatemtu                  int                            `json:"privatemtu"`
	Project                     string                         `json:"project"`
	Projectid                   string                         `json:"projectid"`
	Publicmtu                   int                            `json:"publicmtu"`
	Receivedbytes               int64                          `json:"receivedbytes"`
	Redundantrouter             bool                           `json:"redundantrouter"`
	Related                     string                         `json:"related"`
	Reservediprange             string                         `json:"reservediprange"`
	Restartrequired             bool                           `json:"restartrequired"`
	Sentbytes                   int64                          `json:"sentbytes"`
	Service                     []CreateNetworkResponseService `json:"service"`
	Specifyipranges             bool                           `json:"specifyipranges"`
	Specifyvlan                 bool                           `json:"specifyvlan"`
	State                       string                         `json:"state"`
	Strechedl2subnet            bool                           `json:"strechedl2subnet"`
	Subdomainaccess             bool                           `json:"subdomainaccess"`
	Supportsvmautoscaling       bool                           `json:"supportsvmautoscaling"`
	Tags                        []Tags                         `json:"tags"`
	Traffictype                 string                         `json:"traffictype"`
	Tungstenvirtualrouteruuid   string                         `json:"tungstenvirtualrouteruuid"`
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

func (p *CreatePhysicalNetworkParams) ResetBroadcastdomainrange() {
	if p.p != nil && p.p["broadcastdomainrange"] != nil {
		delete(p.p, "broadcastdomainrange")
	}
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

func (p *CreatePhysicalNetworkParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *CreatePhysicalNetworkParams) ResetIsolationmethods() {
	if p.p != nil && p.p["isolationmethods"] != nil {
		delete(p.p, "isolationmethods")
	}
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

func (p *CreatePhysicalNetworkParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *CreatePhysicalNetworkParams) ResetNetworkspeed() {
	if p.p != nil && p.p["networkspeed"] != nil {
		delete(p.p, "networkspeed")
	}
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

func (p *CreatePhysicalNetworkParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
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

func (p *CreatePhysicalNetworkParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
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

func (p *CreatePhysicalNetworkParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
	resp, err := s.cs.newPostRequest("createPhysicalNetwork", p.toURLValues())
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

func (p *CreateServiceInstanceParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
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

func (p *CreateServiceInstanceParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *CreateServiceInstanceParams) ResetLeftnetworkid() {
	if p.p != nil && p.p["leftnetworkid"] != nil {
		delete(p.p, "leftnetworkid")
	}
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

func (p *CreateServiceInstanceParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *CreateServiceInstanceParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
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

func (p *CreateServiceInstanceParams) ResetRightnetworkid() {
	if p.p != nil && p.p["rightnetworkid"] != nil {
		delete(p.p, "rightnetworkid")
	}
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

func (p *CreateServiceInstanceParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
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

func (p *CreateServiceInstanceParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
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

func (p *CreateServiceInstanceParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
	resp, err := s.cs.newPostRequest("createServiceInstance", p.toURLValues())
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
	Domainpath  string `json:"domainpath"`
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

func (p *CreateStorageNetworkIpRangeParams) ResetEndip() {
	if p.p != nil && p.p["endip"] != nil {
		delete(p.p, "endip")
	}
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

func (p *CreateStorageNetworkIpRangeParams) ResetGateway() {
	if p.p != nil && p.p["gateway"] != nil {
		delete(p.p, "gateway")
	}
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

func (p *CreateStorageNetworkIpRangeParams) ResetNetmask() {
	if p.p != nil && p.p["netmask"] != nil {
		delete(p.p, "netmask")
	}
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

func (p *CreateStorageNetworkIpRangeParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
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

func (p *CreateStorageNetworkIpRangeParams) ResetStartip() {
	if p.p != nil && p.p["startip"] != nil {
		delete(p.p, "startip")
	}
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

func (p *CreateStorageNetworkIpRangeParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
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
	resp, err := s.cs.newPostRequest("createStorageNetworkIpRange", p.toURLValues())
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

func (p *DedicatePublicIpRangeParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
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

func (p *DedicatePublicIpRangeParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *DedicatePublicIpRangeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *DedicatePublicIpRangeParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
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
	resp, err := s.cs.newPostRequest("dedicatePublicIpRange", p.toURLValues())
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
	Cidr              string `json:"cidr"`
	Description       string `json:"description"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Domainpath        string `json:"domainpath"`
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
	Provider          string `json:"provider"`
	Startip           string `json:"startip"`
	Startipv6         string `json:"startipv6"`
	Vlan              string `json:"vlan"`
	Zoneid            string `json:"zoneid"`
}

type DeleteIpv4SubnetForGuestNetworkParams struct {
	p map[string]interface{}
}

func (p *DeleteIpv4SubnetForGuestNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteIpv4SubnetForGuestNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteIpv4SubnetForGuestNetworkParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteIpv4SubnetForGuestNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteIpv4SubnetForGuestNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteIpv4SubnetForGuestNetworkParams(id string) *DeleteIpv4SubnetForGuestNetworkParams {
	p := &DeleteIpv4SubnetForGuestNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an existing IPv4 subnet for guest network.
func (s *NetworkService) DeleteIpv4SubnetForGuestNetwork(p *DeleteIpv4SubnetForGuestNetworkParams) (*DeleteIpv4SubnetForGuestNetworkResponse, error) {
	resp, err := s.cs.newPostRequest("deleteIpv4SubnetForGuestNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteIpv4SubnetForGuestNetworkResponse
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

type DeleteIpv4SubnetForGuestNetworkResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
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

func (p *DeleteNetworkParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
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

func (p *DeleteNetworkParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
	resp, err := s.cs.newPostRequest("deleteNetwork", p.toURLValues())
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

func (p *DeleteNetworkServiceProviderParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
	resp, err := s.cs.newPostRequest("deleteNetworkServiceProvider", p.toURLValues())
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

func (p *DeleteOpenDaylightControllerParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
	resp, err := s.cs.newPostRequest("deleteOpenDaylightController", p.toURLValues())
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

func (p *DeletePhysicalNetworkParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
	resp, err := s.cs.newPostRequest("deletePhysicalNetwork", p.toURLValues())
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

func (p *DeleteStorageNetworkIpRangeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
	resp, err := s.cs.newPostRequest("deleteStorageNetworkIpRange", p.toURLValues())
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

type ListIpv4SubnetsForGuestNetworkParams struct {
	p map[string]interface{}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) toURLValues() url.Values {
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
	if v, found := p.p["parentid"]; found {
		u.Set("parentid", v.(string))
	}
	if v, found := p.p["subnet"]; found {
		u.Set("subnet", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetParentid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parentid"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetParentid() {
	if p.p != nil && p.p["parentid"] != nil {
		delete(p.p, "parentid")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetParentid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parentid"].(string)
	return value, ok
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetSubnet(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["subnet"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetSubnet() {
	if p.p != nil && p.p["subnet"] != nil {
		delete(p.p, "subnet")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetSubnet() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["subnet"].(string)
	return value, ok
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *ListIpv4SubnetsForGuestNetworkParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListIpv4SubnetsForGuestNetworkParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListIpv4SubnetsForGuestNetworkParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListIpv4SubnetsForGuestNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListIpv4SubnetsForGuestNetworkParams() *ListIpv4SubnetsForGuestNetworkParams {
	p := &ListIpv4SubnetsForGuestNetworkParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetIpv4SubnetsForGuestNetworkByID(id string, opts ...OptionFunc) (*Ipv4SubnetsForGuestNetwork, int, error) {
	p := &ListIpv4SubnetsForGuestNetworkParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListIpv4SubnetsForGuestNetwork(p)
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
		return l.Ipv4SubnetsForGuestNetwork[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Ipv4SubnetsForGuestNetwork UUID: %s!", id)
}

// Lists IPv4 subnets for guest networks.
func (s *NetworkService) ListIpv4SubnetsForGuestNetwork(p *ListIpv4SubnetsForGuestNetworkParams) (*ListIpv4SubnetsForGuestNetworkResponse, error) {
	resp, err := s.cs.newRequest("listIpv4SubnetsForGuestNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIpv4SubnetsForGuestNetworkResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListIpv4SubnetsForGuestNetworkResponse struct {
	Count                      int                           `json:"count"`
	Ipv4SubnetsForGuestNetwork []*Ipv4SubnetsForGuestNetwork `json:"ipv4subnetsforguestnetwork"`
}

type Ipv4SubnetsForGuestNetwork struct {
	Allocated    string `json:"allocated"`
	Created      string `json:"created"`
	Id           string `json:"id"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Networkname  string `json:"networkname"`
	Parentid     string `json:"parentid"`
	Parentsubnet string `json:"parentsubnet"`
	Removed      string `json:"removed"`
	State        string `json:"state"`
	Subnet       string `json:"subnet"`
	Vpcid        string `json:"vpcid"`
	Vpcname      string `json:"vpcname"`
	Zoneid       string `json:"zoneid"`
	Zonename     string `json:"zonename"`
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

func (p *ListNetworkIsolationMethodsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListNetworkIsolationMethodsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListNetworkIsolationMethodsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

type ListNetworkProtocolsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkProtocolsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["option"]; found {
		u.Set("option", v.(string))
	}
	return u
}

func (p *ListNetworkProtocolsParams) SetOption(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["option"] = v
}

func (p *ListNetworkProtocolsParams) ResetOption() {
	if p.p != nil && p.p["option"] != nil {
		delete(p.p, "option")
	}
}

func (p *ListNetworkProtocolsParams) GetOption() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["option"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkProtocolsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworkProtocolsParams(option string) *ListNetworkProtocolsParams {
	p := &ListNetworkProtocolsParams{}
	p.p = make(map[string]interface{})
	p.p["option"] = option
	return p
}

// Lists details of network protocols
func (s *NetworkService) ListNetworkProtocols(p *ListNetworkProtocolsParams) (*ListNetworkProtocolsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkProtocols", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkProtocolsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkProtocolsResponse struct {
	Count            int                `json:"count"`
	NetworkProtocols []*NetworkProtocol `json:"networkprotocol"`
}

type NetworkProtocol struct {
	Description string            `json:"description"`
	Details     map[string]string `json:"details"`
	Index       int               `json:"index"`
	JobID       string            `json:"jobid"`
	Jobstatus   int               `json:"jobstatus"`
	Name        string            `json:"name"`
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

func (p *ListNetworkServiceProvidersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListNetworkServiceProvidersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *ListNetworkServiceProvidersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListNetworkServiceProvidersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

func (p *ListNetworkServiceProvidersParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
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

func (p *ListNetworkServiceProvidersParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
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
	if v, found := p.p["associatednetworkid"]; found {
		u.Set("associatednetworkid", v.(string))
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
	if v, found := p.p["networkfilter"]; found {
		u.Set("networkfilter", v.(string))
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
	if v, found := p.p["retrieveonlyresourcecount"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("retrieveonlyresourcecount", vv)
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

func (p *ListNetworksParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListNetworksParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
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

func (p *ListNetworksParams) ResetAcltype() {
	if p.p != nil && p.p["acltype"] != nil {
		delete(p.p, "acltype")
	}
}

func (p *ListNetworksParams) GetAcltype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["acltype"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetAssociatednetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["associatednetworkid"] = v
}

func (p *ListNetworksParams) ResetAssociatednetworkid() {
	if p.p != nil && p.p["associatednetworkid"] != nil {
		delete(p.p, "associatednetworkid")
	}
}

func (p *ListNetworksParams) GetAssociatednetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["associatednetworkid"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetCanusefordeploy(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["canusefordeploy"] = v
}

func (p *ListNetworksParams) ResetCanusefordeploy() {
	if p.p != nil && p.p["canusefordeploy"] != nil {
		delete(p.p, "canusefordeploy")
	}
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

func (p *ListNetworksParams) ResetDisplaynetwork() {
	if p.p != nil && p.p["displaynetwork"] != nil {
		delete(p.p, "displaynetwork")
	}
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

func (p *ListNetworksParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *ListNetworksParams) ResetForvpc() {
	if p.p != nil && p.p["forvpc"] != nil {
		delete(p.p, "forvpc")
	}
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

func (p *ListNetworksParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *ListNetworksParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
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

func (p *ListNetworksParams) ResetIssystem() {
	if p.p != nil && p.p["issystem"] != nil {
		delete(p.p, "issystem")
	}
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

func (p *ListNetworksParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListNetworksParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListNetworksParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetNetworkfilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkfilter"] = v
}

func (p *ListNetworksParams) ResetNetworkfilter() {
	if p.p != nil && p.p["networkfilter"] != nil {
		delete(p.p, "networkfilter")
	}
}

func (p *ListNetworksParams) GetNetworkfilter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkfilter"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
}

func (p *ListNetworksParams) ResetNetworkofferingid() {
	if p.p != nil && p.p["networkofferingid"] != nil {
		delete(p.p, "networkofferingid")
	}
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

func (p *ListNetworksParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListNetworksParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

func (p *ListNetworksParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
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

func (p *ListNetworksParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
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

func (p *ListNetworksParams) ResetRestartrequired() {
	if p.p != nil && p.p["restartrequired"] != nil {
		delete(p.p, "restartrequired")
	}
}

func (p *ListNetworksParams) GetRestartrequired() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["restartrequired"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetRetrieveonlyresourcecount(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["retrieveonlyresourcecount"] = v
}

func (p *ListNetworksParams) ResetRetrieveonlyresourcecount() {
	if p.p != nil && p.p["retrieveonlyresourcecount"] != nil {
		delete(p.p, "retrieveonlyresourcecount")
	}
}

func (p *ListNetworksParams) GetRetrieveonlyresourcecount() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["retrieveonlyresourcecount"].(bool)
	return value, ok
}

func (p *ListNetworksParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListNetworksParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
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

func (p *ListNetworksParams) ResetSpecifyipranges() {
	if p.p != nil && p.p["specifyipranges"] != nil {
		delete(p.p, "specifyipranges")
	}
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

func (p *ListNetworksParams) ResetSupportedservices() {
	if p.p != nil && p.p["supportedservices"] != nil {
		delete(p.p, "supportedservices")
	}
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

func (p *ListNetworksParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
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

func (p *ListNetworksParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
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

func (p *ListNetworksParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ListNetworksParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetVlan(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vlan"] = v
}

func (p *ListNetworksParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
}

func (p *ListNetworksParams) GetVlan() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vlan"].(string)
	return value, ok
}

func (p *ListNetworksParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListNetworksParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
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

func (p *ListNetworksParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
	Asnumber                    int64                    `json:"asnumber"`
	Asnumberid                  string                   `json:"asnumberid"`
	Associatednetwork           string                   `json:"associatednetwork"`
	Associatednetworkid         string                   `json:"associatednetworkid"`
	Bgppeers                    []interface{}            `json:"bgppeers"`
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
	Domainpath                  string                   `json:"domainpath"`
	Egressdefaultpolicy         bool                     `json:"egressdefaultpolicy"`
	Externalid                  string                   `json:"externalid"`
	Gateway                     string                   `json:"gateway"`
	Hasannotations              bool                     `json:"hasannotations"`
	Icon                        interface{}              `json:"icon"`
	Id                          string                   `json:"id"`
	Internetprotocol            string                   `json:"internetprotocol"`
	Ip4routes                   []interface{}            `json:"ip4routes"`
	Ip4routing                  string                   `json:"ip4routing"`
	Ip6cidr                     string                   `json:"ip6cidr"`
	Ip6dns1                     string                   `json:"ip6dns1"`
	Ip6dns2                     string                   `json:"ip6dns2"`
	Ip6gateway                  string                   `json:"ip6gateway"`
	Ip6routes                   []interface{}            `json:"ip6routes"`
	Ip6routing                  string                   `json:"ip6routing"`
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
	Privatemtu                  int                      `json:"privatemtu"`
	Project                     string                   `json:"project"`
	Projectid                   string                   `json:"projectid"`
	Publicmtu                   int                      `json:"publicmtu"`
	Receivedbytes               int64                    `json:"receivedbytes"`
	Redundantrouter             bool                     `json:"redundantrouter"`
	Related                     string                   `json:"related"`
	Reservediprange             string                   `json:"reservediprange"`
	Restartrequired             bool                     `json:"restartrequired"`
	Sentbytes                   int64                    `json:"sentbytes"`
	Service                     []NetworkServiceInternal `json:"service"`
	Specifyipranges             bool                     `json:"specifyipranges"`
	Specifyvlan                 bool                     `json:"specifyvlan"`
	State                       string                   `json:"state"`
	Strechedl2subnet            bool                     `json:"strechedl2subnet"`
	Subdomainaccess             bool                     `json:"subdomainaccess"`
	Supportsvmautoscaling       bool                     `json:"supportsvmautoscaling"`
	Tags                        []Tags                   `json:"tags"`
	Traffictype                 string                   `json:"traffictype"`
	Tungstenvirtualrouteruuid   string                   `json:"tungstenvirtualrouteruuid"`
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

func (p *ListNiciraNvpDeviceNetworksParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListNiciraNvpDeviceNetworksParams) ResetNvpdeviceid() {
	if p.p != nil && p.p["nvpdeviceid"] != nil {
		delete(p.p, "nvpdeviceid")
	}
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

func (p *ListNiciraNvpDeviceNetworksParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListNiciraNvpDeviceNetworksParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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
	Asnumber                    int64                           `json:"asnumber"`
	Asnumberid                  string                          `json:"asnumberid"`
	Associatednetwork           string                          `json:"associatednetwork"`
	Associatednetworkid         string                          `json:"associatednetworkid"`
	Bgppeers                    []interface{}                   `json:"bgppeers"`
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
	Domainpath                  string                          `json:"domainpath"`
	Egressdefaultpolicy         bool                            `json:"egressdefaultpolicy"`
	Externalid                  string                          `json:"externalid"`
	Gateway                     string                          `json:"gateway"`
	Hasannotations              bool                            `json:"hasannotations"`
	Icon                        interface{}                     `json:"icon"`
	Id                          string                          `json:"id"`
	Internetprotocol            string                          `json:"internetprotocol"`
	Ip4routes                   []interface{}                   `json:"ip4routes"`
	Ip4routing                  string                          `json:"ip4routing"`
	Ip6cidr                     string                          `json:"ip6cidr"`
	Ip6dns1                     string                          `json:"ip6dns1"`
	Ip6dns2                     string                          `json:"ip6dns2"`
	Ip6gateway                  string                          `json:"ip6gateway"`
	Ip6routes                   []interface{}                   `json:"ip6routes"`
	Ip6routing                  string                          `json:"ip6routing"`
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
	Privatemtu                  int                             `json:"privatemtu"`
	Project                     string                          `json:"project"`
	Projectid                   string                          `json:"projectid"`
	Publicmtu                   int                             `json:"publicmtu"`
	Receivedbytes               int64                           `json:"receivedbytes"`
	Redundantrouter             bool                            `json:"redundantrouter"`
	Related                     string                          `json:"related"`
	Reservediprange             string                          `json:"reservediprange"`
	Restartrequired             bool                            `json:"restartrequired"`
	Sentbytes                   int64                           `json:"sentbytes"`
	Service                     []NiciraNvpDeviceNetworkService `json:"service"`
	Specifyipranges             bool                            `json:"specifyipranges"`
	Specifyvlan                 bool                            `json:"specifyvlan"`
	State                       string                          `json:"state"`
	Strechedl2subnet            bool                            `json:"strechedl2subnet"`
	Subdomainaccess             bool                            `json:"subdomainaccess"`
	Supportsvmautoscaling       bool                            `json:"supportsvmautoscaling"`
	Tags                        []Tags                          `json:"tags"`
	Traffictype                 string                          `json:"traffictype"`
	Tungstenvirtualrouteruuid   string                          `json:"tungstenvirtualrouteruuid"`
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

func (p *ListOpenDaylightControllersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *ListOpenDaylightControllersParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
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

func (p *ListPaloAltoFirewallNetworksParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListPaloAltoFirewallNetworksParams) ResetLbdeviceid() {
	if p.p != nil && p.p["lbdeviceid"] != nil {
		delete(p.p, "lbdeviceid")
	}
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

func (p *ListPaloAltoFirewallNetworksParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListPaloAltoFirewallNetworksParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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
	Asnumber                    int64                            `json:"asnumber"`
	Asnumberid                  string                           `json:"asnumberid"`
	Associatednetwork           string                           `json:"associatednetwork"`
	Associatednetworkid         string                           `json:"associatednetworkid"`
	Bgppeers                    []interface{}                    `json:"bgppeers"`
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
	Domainpath                  string                           `json:"domainpath"`
	Egressdefaultpolicy         bool                             `json:"egressdefaultpolicy"`
	Externalid                  string                           `json:"externalid"`
	Gateway                     string                           `json:"gateway"`
	Hasannotations              bool                             `json:"hasannotations"`
	Icon                        interface{}                      `json:"icon"`
	Id                          string                           `json:"id"`
	Internetprotocol            string                           `json:"internetprotocol"`
	Ip4routes                   []interface{}                    `json:"ip4routes"`
	Ip4routing                  string                           `json:"ip4routing"`
	Ip6cidr                     string                           `json:"ip6cidr"`
	Ip6dns1                     string                           `json:"ip6dns1"`
	Ip6dns2                     string                           `json:"ip6dns2"`
	Ip6gateway                  string                           `json:"ip6gateway"`
	Ip6routes                   []interface{}                    `json:"ip6routes"`
	Ip6routing                  string                           `json:"ip6routing"`
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
	Privatemtu                  int                              `json:"privatemtu"`
	Project                     string                           `json:"project"`
	Projectid                   string                           `json:"projectid"`
	Publicmtu                   int                              `json:"publicmtu"`
	Receivedbytes               int64                            `json:"receivedbytes"`
	Redundantrouter             bool                             `json:"redundantrouter"`
	Related                     string                           `json:"related"`
	Reservediprange             string                           `json:"reservediprange"`
	Restartrequired             bool                             `json:"restartrequired"`
	Sentbytes                   int64                            `json:"sentbytes"`
	Service                     []PaloAltoFirewallNetworkService `json:"service"`
	Specifyipranges             bool                             `json:"specifyipranges"`
	Specifyvlan                 bool                             `json:"specifyvlan"`
	State                       string                           `json:"state"`
	Strechedl2subnet            bool                             `json:"strechedl2subnet"`
	Subdomainaccess             bool                             `json:"subdomainaccess"`
	Supportsvmautoscaling       bool                             `json:"supportsvmautoscaling"`
	Tags                        []Tags                           `json:"tags"`
	Traffictype                 string                           `json:"traffictype"`
	Tungstenvirtualrouteruuid   string                           `json:"tungstenvirtualrouteruuid"`
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

func (p *ListPhysicalNetworksParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *ListPhysicalNetworksParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListPhysicalNetworksParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *ListPhysicalNetworksParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListPhysicalNetworksParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

func (p *ListPhysicalNetworksParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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

func (p *ListStorageNetworkIpRangeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *ListStorageNetworkIpRangeParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListStorageNetworkIpRangeParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListStorageNetworkIpRangeParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

func (p *ListStorageNetworkIpRangeParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
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

func (p *ListStorageNetworkIpRangeParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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

func (p *ListSupportedNetworkServicesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListSupportedNetworkServicesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListSupportedNetworkServicesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

func (p *ListSupportedNetworkServicesParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
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

func (p *ListSupportedNetworkServicesParams) ResetService() {
	if p.p != nil && p.p["service"] != nil {
		delete(p.p, "service")
	}
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

type MigrateNetworkParams struct {
	p map[string]interface{}
}

func (p *MigrateNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := p.p["resume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("resume", vv)
	}
	return u
}

func (p *MigrateNetworkParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *MigrateNetworkParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *MigrateNetworkParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *MigrateNetworkParams) SetNetworkofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkofferingid"] = v
}

func (p *MigrateNetworkParams) ResetNetworkofferingid() {
	if p.p != nil && p.p["networkofferingid"] != nil {
		delete(p.p, "networkofferingid")
	}
}

func (p *MigrateNetworkParams) GetNetworkofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkofferingid"].(string)
	return value, ok
}

func (p *MigrateNetworkParams) SetResume(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resume"] = v
}

func (p *MigrateNetworkParams) ResetResume() {
	if p.p != nil && p.p["resume"] != nil {
		delete(p.p, "resume")
	}
}

func (p *MigrateNetworkParams) GetResume() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resume"].(bool)
	return value, ok
}

// You should always use this function to get a new MigrateNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewMigrateNetworkParams(networkid string, networkofferingid string) *MigrateNetworkParams {
	p := &MigrateNetworkParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	p.p["networkofferingid"] = networkofferingid
	return p
}

// moves a network to another physical network
func (s *NetworkService) MigrateNetwork(p *MigrateNetworkParams) (*MigrateNetworkResponse, error) {
	resp, err := s.cs.newPostRequest("migrateNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateNetworkResponse
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

type MigrateNetworkResponse struct {
	Account                     string                          `json:"account"`
	Aclid                       string                          `json:"aclid"`
	Aclname                     string                          `json:"aclname"`
	Acltype                     string                          `json:"acltype"`
	Asnumber                    int64                           `json:"asnumber"`
	Asnumberid                  string                          `json:"asnumberid"`
	Associatednetwork           string                          `json:"associatednetwork"`
	Associatednetworkid         string                          `json:"associatednetworkid"`
	Bgppeers                    []interface{}                   `json:"bgppeers"`
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
	Domainpath                  string                          `json:"domainpath"`
	Egressdefaultpolicy         bool                            `json:"egressdefaultpolicy"`
	Externalid                  string                          `json:"externalid"`
	Gateway                     string                          `json:"gateway"`
	Hasannotations              bool                            `json:"hasannotations"`
	Icon                        interface{}                     `json:"icon"`
	Id                          string                          `json:"id"`
	Internetprotocol            string                          `json:"internetprotocol"`
	Ip4routes                   []interface{}                   `json:"ip4routes"`
	Ip4routing                  string                          `json:"ip4routing"`
	Ip6cidr                     string                          `json:"ip6cidr"`
	Ip6dns1                     string                          `json:"ip6dns1"`
	Ip6dns2                     string                          `json:"ip6dns2"`
	Ip6gateway                  string                          `json:"ip6gateway"`
	Ip6routes                   []interface{}                   `json:"ip6routes"`
	Ip6routing                  string                          `json:"ip6routing"`
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
	Privatemtu                  int                             `json:"privatemtu"`
	Project                     string                          `json:"project"`
	Projectid                   string                          `json:"projectid"`
	Publicmtu                   int                             `json:"publicmtu"`
	Receivedbytes               int64                           `json:"receivedbytes"`
	Redundantrouter             bool                            `json:"redundantrouter"`
	Related                     string                          `json:"related"`
	Reservediprange             string                          `json:"reservediprange"`
	Restartrequired             bool                            `json:"restartrequired"`
	Sentbytes                   int64                           `json:"sentbytes"`
	Service                     []MigrateNetworkResponseService `json:"service"`
	Specifyipranges             bool                            `json:"specifyipranges"`
	Specifyvlan                 bool                            `json:"specifyvlan"`
	State                       string                          `json:"state"`
	Strechedl2subnet            bool                            `json:"strechedl2subnet"`
	Subdomainaccess             bool                            `json:"subdomainaccess"`
	Supportsvmautoscaling       bool                            `json:"supportsvmautoscaling"`
	Tags                        []Tags                          `json:"tags"`
	Traffictype                 string                          `json:"traffictype"`
	Tungstenvirtualrouteruuid   string                          `json:"tungstenvirtualrouteruuid"`
	Type                        string                          `json:"type"`
	Vlan                        string                          `json:"vlan"`
	Vpcid                       string                          `json:"vpcid"`
	Vpcname                     string                          `json:"vpcname"`
	Zoneid                      string                          `json:"zoneid"`
	Zonename                    string                          `json:"zonename"`
	Zonesnetworkspans           []interface{}                   `json:"zonesnetworkspans"`
}

type MigrateNetworkResponseService struct {
	Capability []MigrateNetworkResponseServiceCapability `json:"capability"`
	Name       string                                    `json:"name"`
	Provider   []MigrateNetworkResponseServiceProvider   `json:"provider"`
}

type MigrateNetworkResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type MigrateNetworkResponseServiceCapability struct {
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

func (p *ReleasePublicIpRangeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
	resp, err := s.cs.newPostRequest("releasePublicIpRange", p.toURLValues())
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

func (p *RestartNetworkParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *RestartNetworkParams) ResetCleanup() {
	if p.p != nil && p.p["cleanup"] != nil {
		delete(p.p, "cleanup")
	}
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

func (p *RestartNetworkParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RestartNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *RestartNetworkParams) SetLivepatch(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["livepatch"] = v
}

func (p *RestartNetworkParams) ResetLivepatch() {
	if p.p != nil && p.p["livepatch"] != nil {
		delete(p.p, "livepatch")
	}
}

func (p *RestartNetworkParams) GetLivepatch() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["livepatch"].(bool)
	return value, ok
}

func (p *RestartNetworkParams) SetMakeredundant(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["makeredundant"] = v
}

func (p *RestartNetworkParams) ResetMakeredundant() {
	if p.p != nil && p.p["makeredundant"] != nil {
		delete(p.p, "makeredundant")
	}
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
	resp, err := s.cs.newPostRequest("restartNetwork", p.toURLValues())
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
	if v, found := p.p["dns1"]; found {
		u.Set("dns1", v.(string))
	}
	if v, found := p.p["dns2"]; found {
		u.Set("dns2", v.(string))
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
	if v, found := p.p["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := p.p["privatemtu"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privatemtu", vv)
	}
	if v, found := p.p["publicmtu"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicmtu", vv)
	}
	if v, found := p.p["sourcenatipaddress"]; found {
		u.Set("sourcenatipaddress", v.(string))
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

func (p *UpdateNetworkParams) ResetChangecidr() {
	if p.p != nil && p.p["changecidr"] != nil {
		delete(p.p, "changecidr")
	}
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

func (p *UpdateNetworkParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
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

func (p *UpdateNetworkParams) ResetDisplaynetwork() {
	if p.p != nil && p.p["displaynetwork"] != nil {
		delete(p.p, "displaynetwork")
	}
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

func (p *UpdateNetworkParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *UpdateNetworkParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetDns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns1"] = v
}

func (p *UpdateNetworkParams) ResetDns1() {
	if p.p != nil && p.p["dns1"] != nil {
		delete(p.p, "dns1")
	}
}

func (p *UpdateNetworkParams) GetDns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns1"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetDns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dns2"] = v
}

func (p *UpdateNetworkParams) ResetDns2() {
	if p.p != nil && p.p["dns2"] != nil {
		delete(p.p, "dns2")
	}
}

func (p *UpdateNetworkParams) GetDns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dns2"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *UpdateNetworkParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
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

func (p *UpdateNetworkParams) ResetGuestvmcidr() {
	if p.p != nil && p.p["guestvmcidr"] != nil {
		delete(p.p, "guestvmcidr")
	}
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

func (p *UpdateNetworkParams) ResetHideipaddressusage() {
	if p.p != nil && p.p["hideipaddressusage"] != nil {
		delete(p.p, "hideipaddressusage")
	}
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

func (p *UpdateNetworkParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateNetworkParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetIp6dns1(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns1"] = v
}

func (p *UpdateNetworkParams) ResetIp6dns1() {
	if p.p != nil && p.p["ip6dns1"] != nil {
		delete(p.p, "ip6dns1")
	}
}

func (p *UpdateNetworkParams) GetIp6dns1() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns1"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetIp6dns2(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6dns2"] = v
}

func (p *UpdateNetworkParams) ResetIp6dns2() {
	if p.p != nil && p.p["ip6dns2"] != nil {
		delete(p.p, "ip6dns2")
	}
}

func (p *UpdateNetworkParams) GetIp6dns2() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6dns2"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateNetworkParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *UpdateNetworkParams) ResetNetworkdomain() {
	if p.p != nil && p.p["networkdomain"] != nil {
		delete(p.p, "networkdomain")
	}
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

func (p *UpdateNetworkParams) ResetNetworkofferingid() {
	if p.p != nil && p.p["networkofferingid"] != nil {
		delete(p.p, "networkofferingid")
	}
}

func (p *UpdateNetworkParams) GetNetworkofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkofferingid"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetPrivatemtu(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privatemtu"] = v
}

func (p *UpdateNetworkParams) ResetPrivatemtu() {
	if p.p != nil && p.p["privatemtu"] != nil {
		delete(p.p, "privatemtu")
	}
}

func (p *UpdateNetworkParams) GetPrivatemtu() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privatemtu"].(int)
	return value, ok
}

func (p *UpdateNetworkParams) SetPublicmtu(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicmtu"] = v
}

func (p *UpdateNetworkParams) ResetPublicmtu() {
	if p.p != nil && p.p["publicmtu"] != nil {
		delete(p.p, "publicmtu")
	}
}

func (p *UpdateNetworkParams) GetPublicmtu() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicmtu"].(int)
	return value, ok
}

func (p *UpdateNetworkParams) SetSourcenatipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatipaddress"] = v
}

func (p *UpdateNetworkParams) ResetSourcenatipaddress() {
	if p.p != nil && p.p["sourcenatipaddress"] != nil {
		delete(p.p, "sourcenatipaddress")
	}
}

func (p *UpdateNetworkParams) GetSourcenatipaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcenatipaddress"].(string)
	return value, ok
}

func (p *UpdateNetworkParams) SetUpdateinsequence(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["updateinsequence"] = v
}

func (p *UpdateNetworkParams) ResetUpdateinsequence() {
	if p.p != nil && p.p["updateinsequence"] != nil {
		delete(p.p, "updateinsequence")
	}
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
	resp, err := s.cs.newPostRequest("updateNetwork", p.toURLValues())
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
	Asnumber                    int64                          `json:"asnumber"`
	Asnumberid                  string                         `json:"asnumberid"`
	Associatednetwork           string                         `json:"associatednetwork"`
	Associatednetworkid         string                         `json:"associatednetworkid"`
	Bgppeers                    []interface{}                  `json:"bgppeers"`
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
	Domainpath                  string                         `json:"domainpath"`
	Egressdefaultpolicy         bool                           `json:"egressdefaultpolicy"`
	Externalid                  string                         `json:"externalid"`
	Gateway                     string                         `json:"gateway"`
	Hasannotations              bool                           `json:"hasannotations"`
	Icon                        interface{}                    `json:"icon"`
	Id                          string                         `json:"id"`
	Internetprotocol            string                         `json:"internetprotocol"`
	Ip4routes                   []interface{}                  `json:"ip4routes"`
	Ip4routing                  string                         `json:"ip4routing"`
	Ip6cidr                     string                         `json:"ip6cidr"`
	Ip6dns1                     string                         `json:"ip6dns1"`
	Ip6dns2                     string                         `json:"ip6dns2"`
	Ip6gateway                  string                         `json:"ip6gateway"`
	Ip6routes                   []interface{}                  `json:"ip6routes"`
	Ip6routing                  string                         `json:"ip6routing"`
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
	Privatemtu                  int                            `json:"privatemtu"`
	Project                     string                         `json:"project"`
	Projectid                   string                         `json:"projectid"`
	Publicmtu                   int                            `json:"publicmtu"`
	Receivedbytes               int64                          `json:"receivedbytes"`
	Redundantrouter             bool                           `json:"redundantrouter"`
	Related                     string                         `json:"related"`
	Reservediprange             string                         `json:"reservediprange"`
	Restartrequired             bool                           `json:"restartrequired"`
	Sentbytes                   int64                          `json:"sentbytes"`
	Service                     []UpdateNetworkResponseService `json:"service"`
	Specifyipranges             bool                           `json:"specifyipranges"`
	Specifyvlan                 bool                           `json:"specifyvlan"`
	State                       string                         `json:"state"`
	Strechedl2subnet            bool                           `json:"strechedl2subnet"`
	Subdomainaccess             bool                           `json:"subdomainaccess"`
	Supportsvmautoscaling       bool                           `json:"supportsvmautoscaling"`
	Tags                        []Tags                         `json:"tags"`
	Traffictype                 string                         `json:"traffictype"`
	Tungstenvirtualrouteruuid   string                         `json:"tungstenvirtualrouteruuid"`
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

func (p *UpdateNetworkServiceProviderParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *UpdateNetworkServiceProviderParams) ResetServicelist() {
	if p.p != nil && p.p["servicelist"] != nil {
		delete(p.p, "servicelist")
	}
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

func (p *UpdateNetworkServiceProviderParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
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
	resp, err := s.cs.newPostRequest("updateNetworkServiceProvider", p.toURLValues())
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

func (p *UpdatePhysicalNetworkParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *UpdatePhysicalNetworkParams) ResetNetworkspeed() {
	if p.p != nil && p.p["networkspeed"] != nil {
		delete(p.p, "networkspeed")
	}
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

func (p *UpdatePhysicalNetworkParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
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

func (p *UpdatePhysicalNetworkParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
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

func (p *UpdatePhysicalNetworkParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
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
	resp, err := s.cs.newPostRequest("updatePhysicalNetwork", p.toURLValues())
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

func (p *UpdateStorageNetworkIpRangeParams) ResetEndip() {
	if p.p != nil && p.p["endip"] != nil {
		delete(p.p, "endip")
	}
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

func (p *UpdateStorageNetworkIpRangeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *UpdateStorageNetworkIpRangeParams) ResetNetmask() {
	if p.p != nil && p.p["netmask"] != nil {
		delete(p.p, "netmask")
	}
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

func (p *UpdateStorageNetworkIpRangeParams) ResetStartip() {
	if p.p != nil && p.p["startip"] != nil {
		delete(p.p, "startip")
	}
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

func (p *UpdateStorageNetworkIpRangeParams) ResetVlan() {
	if p.p != nil && p.p["vlan"] != nil {
		delete(p.p, "vlan")
	}
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
	resp, err := s.cs.newPostRequest("updateStorageNetworkIpRange", p.toURLValues())
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

type DeleteGuestNetworkIpv6PrefixParams struct {
	p map[string]interface{}
}

func (p *DeleteGuestNetworkIpv6PrefixParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteGuestNetworkIpv6PrefixParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteGuestNetworkIpv6PrefixParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteGuestNetworkIpv6PrefixParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteGuestNetworkIpv6PrefixParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteGuestNetworkIpv6PrefixParams(id string) *DeleteGuestNetworkIpv6PrefixParams {
	p := &DeleteGuestNetworkIpv6PrefixParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an existing guest network IPv6 prefix.
func (s *NetworkService) DeleteGuestNetworkIpv6Prefix(p *DeleteGuestNetworkIpv6PrefixParams) (*DeleteGuestNetworkIpv6PrefixResponse, error) {
	resp, err := s.cs.newPostRequest("deleteGuestNetworkIpv6Prefix", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteGuestNetworkIpv6PrefixResponse
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

type DeleteGuestNetworkIpv6PrefixResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type CreateGuestNetworkIpv6PrefixParams struct {
	p map[string]interface{}
}

func (p *CreateGuestNetworkIpv6PrefixParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["prefix"]; found {
		u.Set("prefix", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateGuestNetworkIpv6PrefixParams) SetPrefix(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["prefix"] = v
}

func (p *CreateGuestNetworkIpv6PrefixParams) ResetPrefix() {
	if p.p != nil && p.p["prefix"] != nil {
		delete(p.p, "prefix")
	}
}

func (p *CreateGuestNetworkIpv6PrefixParams) GetPrefix() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["prefix"].(string)
	return value, ok
}

func (p *CreateGuestNetworkIpv6PrefixParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateGuestNetworkIpv6PrefixParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateGuestNetworkIpv6PrefixParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateGuestNetworkIpv6PrefixParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateGuestNetworkIpv6PrefixParams(prefix string, zoneid string) *CreateGuestNetworkIpv6PrefixParams {
	p := &CreateGuestNetworkIpv6PrefixParams{}
	p.p = make(map[string]interface{})
	p.p["prefix"] = prefix
	p.p["zoneid"] = zoneid
	return p
}

// Creates a guest network IPv6 prefix.
func (s *NetworkService) CreateGuestNetworkIpv6Prefix(p *CreateGuestNetworkIpv6PrefixParams) (*CreateGuestNetworkIpv6PrefixResponse, error) {
	resp, err := s.cs.newPostRequest("createGuestNetworkIpv6Prefix", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateGuestNetworkIpv6PrefixResponse
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

type CreateGuestNetworkIpv6PrefixResponse struct {
	Availablesubnets int    `json:"availablesubnets"`
	Created          string `json:"created"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Prefix           string `json:"prefix"`
	Totalsubnets     int    `json:"totalsubnets"`
	Usedsubnets      int    `json:"usedsubnets"`
	Zoneid           string `json:"zoneid"`
}

type ListGuestNetworkIpv6PrefixesParams struct {
	p map[string]interface{}
}

func (p *ListGuestNetworkIpv6PrefixesParams) toURLValues() url.Values {
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListGuestNetworkIpv6PrefixesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListGuestNetworkIpv6PrefixesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListGuestNetworkIpv6PrefixesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListGuestNetworkIpv6PrefixesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListGuestNetworkIpv6PrefixesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListGuestNetworkIpv6PrefixesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListGuestNetworkIpv6PrefixesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListGuestNetworkIpv6PrefixesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListGuestNetworkIpv6PrefixesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListGuestNetworkIpv6PrefixesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListGuestNetworkIpv6PrefixesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListGuestNetworkIpv6PrefixesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListGuestNetworkIpv6PrefixesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListGuestNetworkIpv6PrefixesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListGuestNetworkIpv6PrefixesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListGuestNetworkIpv6PrefixesParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListGuestNetworkIpv6PrefixesParams() *ListGuestNetworkIpv6PrefixesParams {
	p := &ListGuestNetworkIpv6PrefixesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetGuestNetworkIpv6PrefixeByID(id string, opts ...OptionFunc) (*GuestNetworkIpv6Prefixe, int, error) {
	p := &ListGuestNetworkIpv6PrefixesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListGuestNetworkIpv6Prefixes(p)
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
		return l.GuestNetworkIpv6Prefixes[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GuestNetworkIpv6Prefixe UUID: %s!", id)
}

// Lists guest network IPv6 prefixes
func (s *NetworkService) ListGuestNetworkIpv6Prefixes(p *ListGuestNetworkIpv6PrefixesParams) (*ListGuestNetworkIpv6PrefixesResponse, error) {
	resp, err := s.cs.newRequest("listGuestNetworkIpv6Prefixes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGuestNetworkIpv6PrefixesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListGuestNetworkIpv6PrefixesResponse struct {
	Count                    int                        `json:"count"`
	GuestNetworkIpv6Prefixes []*GuestNetworkIpv6Prefixe `json:"guestnetworkipv6prefixe"`
}

type GuestNetworkIpv6Prefixe struct {
	Availablesubnets int    `json:"availablesubnets"`
	Created          string `json:"created"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Prefix           string `json:"prefix"`
	Totalsubnets     int    `json:"totalsubnets"`
	Usedsubnets      int    `json:"usedsubnets"`
	Zoneid           string `json:"zoneid"`
}

type CreateNetworkPermissionsParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accountids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("accountids", vv)
	}
	if v, found := p.p["accounts"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("accounts", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["projectids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("projectids", vv)
	}
	return u
}

func (p *CreateNetworkPermissionsParams) SetAccountids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountids"] = v
}

func (p *CreateNetworkPermissionsParams) ResetAccountids() {
	if p.p != nil && p.p["accountids"] != nil {
		delete(p.p, "accountids")
	}
}

func (p *CreateNetworkPermissionsParams) GetAccountids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountids"].([]string)
	return value, ok
}

func (p *CreateNetworkPermissionsParams) SetAccounts(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounts"] = v
}

func (p *CreateNetworkPermissionsParams) ResetAccounts() {
	if p.p != nil && p.p["accounts"] != nil {
		delete(p.p, "accounts")
	}
}

func (p *CreateNetworkPermissionsParams) GetAccounts() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounts"].([]string)
	return value, ok
}

func (p *CreateNetworkPermissionsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreateNetworkPermissionsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreateNetworkPermissionsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreateNetworkPermissionsParams) SetProjectids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectids"] = v
}

func (p *CreateNetworkPermissionsParams) ResetProjectids() {
	if p.p != nil && p.p["projectids"] != nil {
		delete(p.p, "projectids")
	}
}

func (p *CreateNetworkPermissionsParams) GetProjectids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectids"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateNetworkPermissionsParams(networkid string) *CreateNetworkPermissionsParams {
	p := &CreateNetworkPermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	return p
}

// Updates network permissions.
func (s *NetworkService) CreateNetworkPermissions(p *CreateNetworkPermissionsParams) (*CreateNetworkPermissionsResponse, error) {
	resp, err := s.cs.newPostRequest("createNetworkPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateNetworkPermissionsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *CreateNetworkPermissionsResponse) UnmarshalJSON(b []byte) error {
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

	type alias CreateNetworkPermissionsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ResetNetworkPermissionsParams struct {
	p map[string]interface{}
}

func (p *ResetNetworkPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	return u
}

func (p *ResetNetworkPermissionsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ResetNetworkPermissionsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ResetNetworkPermissionsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

// You should always use this function to get a new ResetNetworkPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewResetNetworkPermissionsParams(networkid string) *ResetNetworkPermissionsParams {
	p := &ResetNetworkPermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	return p
}

// Resets network permissions.
func (s *NetworkService) ResetNetworkPermissions(p *ResetNetworkPermissionsParams) (*ResetNetworkPermissionsResponse, error) {
	resp, err := s.cs.newPostRequest("resetNetworkPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetNetworkPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ResetNetworkPermissionsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ResetNetworkPermissionsResponse) UnmarshalJSON(b []byte) error {
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

	type alias ResetNetworkPermissionsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListNetworkPermissionsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	return u
}

func (p *ListNetworkPermissionsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListNetworkPermissionsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListNetworkPermissionsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworkPermissionsParams(networkid string) *ListNetworkPermissionsParams {
	p := &ListNetworkPermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	return p
}

// List network visibility and all accounts that have permissions to view this network.
func (s *NetworkService) ListNetworkPermissions(p *ListNetworkPermissionsParams) (*ListNetworkPermissionsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkPermissionsResponse struct {
	Count              int                  `json:"count"`
	NetworkPermissions []*NetworkPermission `json:"networkpermission"`
}

type NetworkPermission struct {
	Account   string `json:"account"`
	Accountid string `json:"accountid"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Networkid string `json:"networkid"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
}

type RemoveNetworkPermissionsParams struct {
	p map[string]interface{}
}

func (p *RemoveNetworkPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accountids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("accountids", vv)
	}
	if v, found := p.p["accounts"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("accounts", vv)
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["projectids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("projectids", vv)
	}
	return u
}

func (p *RemoveNetworkPermissionsParams) SetAccountids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountids"] = v
}

func (p *RemoveNetworkPermissionsParams) ResetAccountids() {
	if p.p != nil && p.p["accountids"] != nil {
		delete(p.p, "accountids")
	}
}

func (p *RemoveNetworkPermissionsParams) GetAccountids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountids"].([]string)
	return value, ok
}

func (p *RemoveNetworkPermissionsParams) SetAccounts(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounts"] = v
}

func (p *RemoveNetworkPermissionsParams) ResetAccounts() {
	if p.p != nil && p.p["accounts"] != nil {
		delete(p.p, "accounts")
	}
}

func (p *RemoveNetworkPermissionsParams) GetAccounts() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounts"].([]string)
	return value, ok
}

func (p *RemoveNetworkPermissionsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *RemoveNetworkPermissionsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *RemoveNetworkPermissionsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *RemoveNetworkPermissionsParams) SetProjectids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectids"] = v
}

func (p *RemoveNetworkPermissionsParams) ResetProjectids() {
	if p.p != nil && p.p["projectids"] != nil {
		delete(p.p, "projectids")
	}
}

func (p *RemoveNetworkPermissionsParams) GetProjectids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectids"].([]string)
	return value, ok
}

// You should always use this function to get a new RemoveNetworkPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewRemoveNetworkPermissionsParams(networkid string) *RemoveNetworkPermissionsParams {
	p := &RemoveNetworkPermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	return p
}

// Removes network permissions.
func (s *NetworkService) RemoveNetworkPermissions(p *RemoveNetworkPermissionsParams) (*RemoveNetworkPermissionsResponse, error) {
	resp, err := s.cs.newPostRequest("removeNetworkPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveNetworkPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveNetworkPermissionsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RemoveNetworkPermissionsResponse) UnmarshalJSON(b []byte) error {
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

	type alias RemoveNetworkPermissionsResponse
	return json.Unmarshal(b, (*alias)(r))
}
