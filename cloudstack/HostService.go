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

type HostServiceIface interface {
	AddBaremetalHost(p *AddBaremetalHostParams) (*AddBaremetalHostResponse, error)
	NewAddBaremetalHostParams(hypervisor string, podid string, url string, zoneid string) *AddBaremetalHostParams
	AddGloboDnsHost(p *AddGloboDnsHostParams) (*AddGloboDnsHostResponse, error)
	NewAddGloboDnsHostParams(password string, physicalnetworkid string, url string, username string) *AddGloboDnsHostParams
	AddHost(p *AddHostParams) (*AddHostResponse, error)
	NewAddHostParams(hypervisor string, podid string, url string, zoneid string) *AddHostParams
	AddSecondaryStorage(p *AddSecondaryStorageParams) (*AddSecondaryStorageResponse, error)
	NewAddSecondaryStorageParams(url string) *AddSecondaryStorageParams
	CancelHostMaintenance(p *CancelHostMaintenanceParams) (*CancelHostMaintenanceResponse, error)
	NewCancelHostMaintenanceParams(id string) *CancelHostMaintenanceParams
	ConfigureHAForHost(p *ConfigureHAForHostParams) (*HAForHostResponse, error)
	NewConfigureHAForHostParams(hostid string, provider string) *ConfigureHAForHostParams
	EnableHAForHost(p *EnableHAForHostParams) (*EnableHAForHostResponse, error)
	NewEnableHAForHostParams(hostid string) *EnableHAForHostParams
	DedicateHost(p *DedicateHostParams) (*DedicateHostResponse, error)
	NewDedicateHostParams(domainid string, hostid string) *DedicateHostParams
	DeleteHost(p *DeleteHostParams) (*DeleteHostResponse, error)
	NewDeleteHostParams(id string) *DeleteHostParams
	DisableOutOfBandManagementForHost(p *DisableOutOfBandManagementForHostParams) (*DisableOutOfBandManagementForHostResponse, error)
	NewDisableOutOfBandManagementForHostParams(hostid string) *DisableOutOfBandManagementForHostParams
	EnableOutOfBandManagementForHost(p *EnableOutOfBandManagementForHostParams) (*EnableOutOfBandManagementForHostResponse, error)
	NewEnableOutOfBandManagementForHostParams(hostid string) *EnableOutOfBandManagementForHostParams
	FindHostsForMigration(p *FindHostsForMigrationParams) (*FindHostsForMigrationResponse, error)
	NewFindHostsForMigrationParams(virtualmachineid string) *FindHostsForMigrationParams
	ListDedicatedHosts(p *ListDedicatedHostsParams) (*ListDedicatedHostsResponse, error)
	NewListDedicatedHostsParams() *ListDedicatedHostsParams
	ListHostTags(p *ListHostTagsParams) (*ListHostTagsResponse, error)
	NewListHostTagsParams() *ListHostTagsParams
	GetHostTagID(keyword string, opts ...OptionFunc) (string, int, error)
	ListHosts(p *ListHostsParams) (*ListHostsResponse, error)
	NewListHostsParams() *ListHostsParams
	GetHostID(name string, opts ...OptionFunc) (string, int, error)
	GetHostByName(name string, opts ...OptionFunc) (*Host, int, error)
	GetHostByID(id string, opts ...OptionFunc) (*Host, int, error)
	ListHostsMetrics(p *ListHostsMetricsParams) (*ListHostsMetricsResponse, error)
	NewListHostsMetricsParams() *ListHostsMetricsParams
	GetHostsMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetHostsMetricByName(name string, opts ...OptionFunc) (*HostsMetric, int, error)
	GetHostsMetricByID(id string, opts ...OptionFunc) (*HostsMetric, int, error)
	PrepareHostForMaintenance(p *PrepareHostForMaintenanceParams) (*PrepareHostForMaintenanceResponse, error)
	NewPrepareHostForMaintenanceParams(id string) *PrepareHostForMaintenanceParams
	ReconnectHost(p *ReconnectHostParams) (*ReconnectHostResponse, error)
	NewReconnectHostParams(id string) *ReconnectHostParams
	ReleaseDedicatedHost(p *ReleaseDedicatedHostParams) (*ReleaseDedicatedHostResponse, error)
	NewReleaseDedicatedHostParams(hostid string) *ReleaseDedicatedHostParams
	ReleaseHostReservation(p *ReleaseHostReservationParams) (*ReleaseHostReservationResponse, error)
	NewReleaseHostReservationParams(id string) *ReleaseHostReservationParams
	UpdateHost(p *UpdateHostParams) (*UpdateHostResponse, error)
	NewUpdateHostParams(id string) *UpdateHostParams
	UpdateHostPassword(p *UpdateHostPasswordParams) (*UpdateHostPasswordResponse, error)
	NewUpdateHostPasswordParams(password string, username string) *UpdateHostPasswordParams
}

type AddBaremetalHostParams struct {
	p map[string]interface{}
}

func (p *AddBaremetalHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["hosttags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("hosttags", vv)
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddBaremetalHostParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *AddBaremetalHostParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *AddBaremetalHostParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
}

func (p *AddBaremetalHostParams) GetClustername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustername"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetHosttags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
}

func (p *AddBaremetalHostParams) GetHosttags() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hosttags"].([]string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *AddBaremetalHostParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *AddBaremetalHostParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddBaremetalHostParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *AddBaremetalHostParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddBaremetalHostParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddBaremetalHostParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *AddBaremetalHostParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddBaremetalHostParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddBaremetalHostParams(hypervisor string, podid string, url string, zoneid string) *AddBaremetalHostParams {
	p := &AddBaremetalHostParams{}
	p.p = make(map[string]interface{})
	p.p["hypervisor"] = hypervisor
	p.p["podid"] = podid
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// add a baremetal host
func (s *HostService) AddBaremetalHost(p *AddBaremetalHostParams) (*AddBaremetalHostResponse, error) {
	resp, err := s.cs.newRequest("addBaremetalHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddBaremetalHostResponse struct {
	Annotation                       string                             `json:"annotation"`
	Capabilities                     string                             `json:"capabilities"`
	Clusterid                        string                             `json:"clusterid"`
	Clustername                      string                             `json:"clustername"`
	Clustertype                      string                             `json:"clustertype"`
	Cpuallocated                     string                             `json:"cpuallocated"`
	Cpuallocatedpercentage           string                             `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                              `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                             `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                            `json:"cpuloadaverage"`
	Cpunumber                        int                                `json:"cpunumber"`
	Cpusockets                       int                                `json:"cpusockets"`
	Cpuspeed                         int64                              `json:"cpuspeed"`
	Cpuused                          string                             `json:"cpuused"`
	Cpuwithoverprovisioning          string                             `json:"cpuwithoverprovisioning"`
	Created                          string                             `json:"created"`
	Details                          map[string]string                  `json:"details"`
	Disconnected                     string                             `json:"disconnected"`
	Disksizeallocated                int64                              `json:"disksizeallocated"`
	Disksizetotal                    int64                              `json:"disksizetotal"`
	Encryptionsupported              bool                               `json:"encryptionsupported"`
	Events                           string                             `json:"events"`
	Gpugroup                         []AddBaremetalHostResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                               `json:"hahost"`
	Hasannotations                   bool                               `json:"hasannotations"`
	Hasenoughcapacity                bool                               `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse                  `json:"hostha"`
	Hosttags                         string                             `json:"hosttags"`
	Hypervisor                       string                             `json:"hypervisor"`
	Hypervisorversion                string                             `json:"hypervisorversion"`
	Id                               string                             `json:"id"`
	Ipaddress                        string                             `json:"ipaddress"`
	Islocalstorageactive             bool                               `json:"islocalstorageactive"`
	JobID                            string                             `json:"jobid"`
	Jobstatus                        int                                `json:"jobstatus"`
	Lastannotated                    string                             `json:"lastannotated"`
	Lastpinged                       string                             `json:"lastpinged"`
	Managementserverid               UUID                               `json:"managementserverid"`
	Memoryallocated                  int64                              `json:"memoryallocated"`
	Memoryallocatedbytes             int64                              `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                             `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                              `json:"memorytotal"`
	Memoryused                       int64                              `json:"memoryused"`
	Memorywithoverprovisioning       string                             `json:"memorywithoverprovisioning"`
	Name                             string                             `json:"name"`
	Networkkbsread                   int64                              `json:"networkkbsread"`
	Networkkbswrite                  int64                              `json:"networkkbswrite"`
	Oscategoryid                     string                             `json:"oscategoryid"`
	Oscategoryname                   string                             `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse        `json:"outofbandmanagement"`
	Podid                            string                             `json:"podid"`
	Podname                          string                             `json:"podname"`
	Removed                          string                             `json:"removed"`
	Resourcestate                    string                             `json:"resourcestate"`
	State                            string                             `json:"state"`
	Suitableformigration             bool                               `json:"suitableformigration"`
	Type                             string                             `json:"type"`
	Ueficapability                   bool                               `json:"ueficapability"`
	Username                         string                             `json:"username"`
	Version                          string                             `json:"version"`
	Zoneid                           string                             `json:"zoneid"`
	Zonename                         string                             `json:"zonename"`
}

type AddBaremetalHostResponseGpugroup struct {
	Gpugroupname string                                 `json:"gpugroupname"`
	Vgpu         []AddBaremetalHostResponseGpugroupVgpu `json:"vgpu"`
}

type AddBaremetalHostResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type AddGloboDnsHostParams struct {
	p map[string]interface{}
}

func (p *AddGloboDnsHostParams) toURLValues() url.Values {
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

func (p *AddGloboDnsHostParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddGloboDnsHostParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddGloboDnsHostParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *AddGloboDnsHostParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *AddGloboDnsHostParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddGloboDnsHostParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddGloboDnsHostParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddGloboDnsHostParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddGloboDnsHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddGloboDnsHostParams(password string, physicalnetworkid string, url string, username string) *AddGloboDnsHostParams {
	p := &AddGloboDnsHostParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds the GloboDNS external host
func (s *HostService) AddGloboDnsHost(p *AddGloboDnsHostParams) (*AddGloboDnsHostResponse, error) {
	resp, err := s.cs.newRequest("addGloboDnsHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddGloboDnsHostResponse
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

type AddGloboDnsHostResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AddHostParams struct {
	p map[string]interface{}
}

func (p *AddHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["hosttags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("hosttags", vv)
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddHostParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *AddHostParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *AddHostParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *AddHostParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *AddHostParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
}

func (p *AddHostParams) GetClustername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustername"].(string)
	return value, ok
}

func (p *AddHostParams) SetHosttags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
}

func (p *AddHostParams) GetHosttags() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hosttags"].([]string)
	return value, ok
}

func (p *AddHostParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *AddHostParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *AddHostParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddHostParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddHostParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *AddHostParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *AddHostParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddHostParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddHostParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddHostParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *AddHostParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddHostParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddHostParams(hypervisor string, podid string, url string, zoneid string) *AddHostParams {
	p := &AddHostParams{}
	p.p = make(map[string]interface{})
	p.p["hypervisor"] = hypervisor
	p.p["podid"] = podid
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Adds a new host.
func (s *HostService) AddHost(p *AddHostParams) (*AddHostResponse, error) {
	resp, err := s.cs.newRequest("addHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r AddHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddHostResponse struct {
	Annotation                       string                      `json:"annotation"`
	Capabilities                     string                      `json:"capabilities"`
	Clusterid                        string                      `json:"clusterid"`
	Clustername                      string                      `json:"clustername"`
	Clustertype                      string                      `json:"clustertype"`
	Cpuallocated                     string                      `json:"cpuallocated"`
	Cpuallocatedpercentage           string                      `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                       `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                      `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                     `json:"cpuloadaverage"`
	Cpunumber                        int                         `json:"cpunumber"`
	Cpusockets                       int                         `json:"cpusockets"`
	Cpuspeed                         int64                       `json:"cpuspeed"`
	Cpuused                          string                      `json:"cpuused"`
	Cpuwithoverprovisioning          string                      `json:"cpuwithoverprovisioning"`
	Created                          string                      `json:"created"`
	Details                          map[string]string           `json:"details"`
	Disconnected                     string                      `json:"disconnected"`
	Disksizeallocated                int64                       `json:"disksizeallocated"`
	Disksizetotal                    int64                       `json:"disksizetotal"`
	Encryptionsupported              bool                        `json:"encryptionsupported"`
	Events                           string                      `json:"events"`
	Gpugroup                         []AddHostResponseGpugroup   `json:"gpugroup"`
	Hahost                           bool                        `json:"hahost"`
	Hasannotations                   bool                        `json:"hasannotations"`
	Hasenoughcapacity                bool                        `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse           `json:"hostha"`
	Hosttags                         string                      `json:"hosttags"`
	Hypervisor                       string                      `json:"hypervisor"`
	Hypervisorversion                string                      `json:"hypervisorversion"`
	Id                               string                      `json:"id"`
	Ipaddress                        string                      `json:"ipaddress"`
	Islocalstorageactive             bool                        `json:"islocalstorageactive"`
	JobID                            string                      `json:"jobid"`
	Jobstatus                        int                         `json:"jobstatus"`
	Lastannotated                    string                      `json:"lastannotated"`
	Lastpinged                       string                      `json:"lastpinged"`
	Managementserverid               UUID                        `json:"managementserverid"`
	Memoryallocated                  int64                       `json:"memoryallocated"`
	Memoryallocatedbytes             int64                       `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                      `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                       `json:"memorytotal"`
	Memoryused                       int64                       `json:"memoryused"`
	Memorywithoverprovisioning       string                      `json:"memorywithoverprovisioning"`
	Name                             string                      `json:"name"`
	Networkkbsread                   int64                       `json:"networkkbsread"`
	Networkkbswrite                  int64                       `json:"networkkbswrite"`
	Oscategoryid                     string                      `json:"oscategoryid"`
	Oscategoryname                   string                      `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse `json:"outofbandmanagement"`
	Podid                            string                      `json:"podid"`
	Podname                          string                      `json:"podname"`
	Removed                          string                      `json:"removed"`
	Resourcestate                    string                      `json:"resourcestate"`
	State                            string                      `json:"state"`
	Suitableformigration             bool                        `json:"suitableformigration"`
	Type                             string                      `json:"type"`
	Ueficapability                   bool                        `json:"ueficapability"`
	Username                         string                      `json:"username"`
	Version                          string                      `json:"version"`
	Zoneid                           string                      `json:"zoneid"`
	Zonename                         string                      `json:"zonename"`
}

type AddHostResponseGpugroup struct {
	Gpugroupname string                        `json:"gpugroupname"`
	Vgpu         []AddHostResponseGpugroupVgpu `json:"vgpu"`
}

type AddHostResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type AddSecondaryStorageParams struct {
	p map[string]interface{}
}

func (p *AddSecondaryStorageParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddSecondaryStorageParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddSecondaryStorageParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddSecondaryStorageParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddSecondaryStorageParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddSecondaryStorageParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddSecondaryStorageParams(url string) *AddSecondaryStorageParams {
	p := &AddSecondaryStorageParams{}
	p.p = make(map[string]interface{})
	p.p["url"] = url
	return p
}

// Adds secondary storage.
func (s *HostService) AddSecondaryStorage(p *AddSecondaryStorageParams) (*AddSecondaryStorageResponse, error) {
	resp, err := s.cs.newRequest("addSecondaryStorage", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddSecondaryStorageResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddSecondaryStorageResponse struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}

type CancelHostMaintenanceParams struct {
	p map[string]interface{}
}

func (p *CancelHostMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *CancelHostMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *CancelHostMaintenanceParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new CancelHostMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewCancelHostMaintenanceParams(id string) *CancelHostMaintenanceParams {
	p := &CancelHostMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Cancels host maintenance.
func (s *HostService) CancelHostMaintenance(p *CancelHostMaintenanceParams) (*CancelHostMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("cancelHostMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelHostMaintenanceResponse
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

type CancelHostMaintenanceResponse struct {
	Annotation                       string                                  `json:"annotation"`
	Capabilities                     string                                  `json:"capabilities"`
	Clusterid                        string                                  `json:"clusterid"`
	Clustername                      string                                  `json:"clustername"`
	Clustertype                      string                                  `json:"clustertype"`
	Cpuallocated                     string                                  `json:"cpuallocated"`
	Cpuallocatedpercentage           string                                  `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                                   `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                                  `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                                 `json:"cpuloadaverage"`
	Cpunumber                        int                                     `json:"cpunumber"`
	Cpusockets                       int                                     `json:"cpusockets"`
	Cpuspeed                         int64                                   `json:"cpuspeed"`
	Cpuused                          string                                  `json:"cpuused"`
	Cpuwithoverprovisioning          string                                  `json:"cpuwithoverprovisioning"`
	Created                          string                                  `json:"created"`
	Details                          map[string]string                       `json:"details"`
	Disconnected                     string                                  `json:"disconnected"`
	Disksizeallocated                int64                                   `json:"disksizeallocated"`
	Disksizetotal                    int64                                   `json:"disksizetotal"`
	Encryptionsupported              bool                                    `json:"encryptionsupported"`
	Events                           string                                  `json:"events"`
	Gpugroup                         []CancelHostMaintenanceResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                                    `json:"hahost"`
	Hasannotations                   bool                                    `json:"hasannotations"`
	Hasenoughcapacity                bool                                    `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse                       `json:"hostha"`
	Hosttags                         string                                  `json:"hosttags"`
	Hypervisor                       string                                  `json:"hypervisor"`
	Hypervisorversion                string                                  `json:"hypervisorversion"`
	Id                               string                                  `json:"id"`
	Ipaddress                        string                                  `json:"ipaddress"`
	Islocalstorageactive             bool                                    `json:"islocalstorageactive"`
	JobID                            string                                  `json:"jobid"`
	Jobstatus                        int                                     `json:"jobstatus"`
	Lastannotated                    string                                  `json:"lastannotated"`
	Lastpinged                       string                                  `json:"lastpinged"`
	Managementserverid               UUID                                    `json:"managementserverid"`
	Memoryallocated                  int64                                   `json:"memoryallocated"`
	Memoryallocatedbytes             int64                                   `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                                  `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                                   `json:"memorytotal"`
	Memoryused                       int64                                   `json:"memoryused"`
	Memorywithoverprovisioning       string                                  `json:"memorywithoverprovisioning"`
	Name                             string                                  `json:"name"`
	Networkkbsread                   int64                                   `json:"networkkbsread"`
	Networkkbswrite                  int64                                   `json:"networkkbswrite"`
	Oscategoryid                     string                                  `json:"oscategoryid"`
	Oscategoryname                   string                                  `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse             `json:"outofbandmanagement"`
	Podid                            string                                  `json:"podid"`
	Podname                          string                                  `json:"podname"`
	Removed                          string                                  `json:"removed"`
	Resourcestate                    string                                  `json:"resourcestate"`
	State                            string                                  `json:"state"`
	Suitableformigration             bool                                    `json:"suitableformigration"`
	Type                             string                                  `json:"type"`
	Ueficapability                   bool                                    `json:"ueficapability"`
	Username                         string                                  `json:"username"`
	Version                          string                                  `json:"version"`
	Zoneid                           string                                  `json:"zoneid"`
	Zonename                         string                                  `json:"zonename"`
}

type CancelHostMaintenanceResponseGpugroup struct {
	Gpugroupname string                                      `json:"gpugroupname"`
	Vgpu         []CancelHostMaintenanceResponseGpugroupVgpu `json:"vgpu"`
}

type CancelHostMaintenanceResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type ConfigureHAForHostParams struct {
	p map[string]interface{}
}

func (p *ConfigureHAForHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (p *ConfigureHAForHostParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ConfigureHAForHostParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ConfigureHAForHostParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ConfigureHAForHostParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigureHAForHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewConfigureHAForHostParams(hostid string, provider string) *ConfigureHAForHostParams {
	p := &ConfigureHAForHostParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	p.p["provider"] = provider
	return p
}

// Configures HA for a host
func (s *HostService) ConfigureHAForHost(p *ConfigureHAForHostParams) (*HAForHostResponse, error) {
	resp, err := s.cs.newRequest("configureHAForHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r HAForHostResponse
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

type HAForHostResponse struct {
	Haenable   bool   `json:"haenable"`
	Haprovider string `json:"haprovider"`
	Hastate    string `json:"hastate"`
	Hostid     string `json:"hostid"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Status     bool   `json:"status"`
}

type EnableHAForHostParams struct {
	p map[string]interface{}
}

func (p *EnableHAForHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (p *EnableHAForHostParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *EnableHAForHostParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableHAForHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewEnableHAForHostParams(hostid string) *EnableHAForHostParams {
	p := &EnableHAForHostParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	return p
}

// Enables HA for a host
func (s *HostService) EnableHAForHost(p *EnableHAForHostParams) (*EnableHAForHostResponse, error) {
	resp, err := s.cs.newRequest("enableHAForHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableHAForHostResponse
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

type EnableHAForHostResponse struct {
	Haenable   bool   `json:"haenable"`
	Haprovider string `json:"haprovider"`
	Hastate    string `json:"hastate"`
	Hostid     string `json:"hostid"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Status     bool   `json:"status"`
}

type DedicateHostParams struct {
	p map[string]interface{}
}

func (p *DedicateHostParams) toURLValues() url.Values {
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
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (p *DedicateHostParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DedicateHostParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DedicateHostParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DedicateHostParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DedicateHostParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *DedicateHostParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicateHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewDedicateHostParams(domainid string, hostid string) *DedicateHostParams {
	p := &DedicateHostParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	p.p["hostid"] = hostid
	return p
}

// Dedicates a host.
func (s *HostService) DedicateHost(p *DedicateHostParams) (*DedicateHostResponse, error) {
	resp, err := s.cs.newRequest("dedicateHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateHostResponse
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

type DedicateHostResponse struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Hostid          string `json:"hostid"`
	Hostname        string `json:"hostname"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
}

type DeleteHostParams struct {
	p map[string]interface{}
}

func (p *DeleteHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["forcedestroylocalstorage"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forcedestroylocalstorage", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteHostParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DeleteHostParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *DeleteHostParams) SetForcedestroylocalstorage(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forcedestroylocalstorage"] = v
}

func (p *DeleteHostParams) GetForcedestroylocalstorage() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forcedestroylocalstorage"].(bool)
	return value, ok
}

func (p *DeleteHostParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteHostParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewDeleteHostParams(id string) *DeleteHostParams {
	p := &DeleteHostParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a host.
func (s *HostService) DeleteHost(p *DeleteHostParams) (*DeleteHostResponse, error) {
	resp, err := s.cs.newRequest("deleteHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteHostResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteHostResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteHostResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableOutOfBandManagementForHostParams struct {
	p map[string]interface{}
}

func (p *DisableOutOfBandManagementForHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (p *DisableOutOfBandManagementForHostParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *DisableOutOfBandManagementForHostParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableOutOfBandManagementForHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewDisableOutOfBandManagementForHostParams(hostid string) *DisableOutOfBandManagementForHostParams {
	p := &DisableOutOfBandManagementForHostParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	return p
}

// Disables out-of-band management for a host
func (s *HostService) DisableOutOfBandManagementForHost(p *DisableOutOfBandManagementForHostParams) (*DisableOutOfBandManagementForHostResponse, error) {
	resp, err := s.cs.newRequest("disableOutOfBandManagementForHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableOutOfBandManagementForHostResponse
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

type DisableOutOfBandManagementForHostResponse struct {
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

type EnableOutOfBandManagementForHostParams struct {
	p map[string]interface{}
}

func (p *EnableOutOfBandManagementForHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (p *EnableOutOfBandManagementForHostParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *EnableOutOfBandManagementForHostParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableOutOfBandManagementForHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewEnableOutOfBandManagementForHostParams(hostid string) *EnableOutOfBandManagementForHostParams {
	p := &EnableOutOfBandManagementForHostParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	return p
}

// Enables out-of-band management for a host
func (s *HostService) EnableOutOfBandManagementForHost(p *EnableOutOfBandManagementForHostParams) (*EnableOutOfBandManagementForHostResponse, error) {
	resp, err := s.cs.newRequest("enableOutOfBandManagementForHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableOutOfBandManagementForHostResponse
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

type EnableOutOfBandManagementForHostResponse struct {
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

type FindHostsForMigrationParams struct {
	p map[string]interface{}
}

func (p *FindHostsForMigrationParams) toURLValues() url.Values {
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
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *FindHostsForMigrationParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *FindHostsForMigrationParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *FindHostsForMigrationParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *FindHostsForMigrationParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *FindHostsForMigrationParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *FindHostsForMigrationParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *FindHostsForMigrationParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *FindHostsForMigrationParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new FindHostsForMigrationParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewFindHostsForMigrationParams(virtualmachineid string) *FindHostsForMigrationParams {
	p := &FindHostsForMigrationParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Find hosts suitable for migrating a virtual machine.
func (s *HostService) FindHostsForMigration(p *FindHostsForMigrationParams) (*FindHostsForMigrationResponse, error) {
	resp, err := s.cs.newRequest("findHostsForMigration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r FindHostsForMigrationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type FindHostsForMigrationResponse struct {
	Count int                 `json:"count"`
	Host  []*HostForMigration `json:"host"`
}

type HostForMigration struct {
	Averageload                      int64  `json:"averageload"`
	Capabilities                     string `json:"capabilities"`
	Clusterid                        string `json:"clusterid"`
	Clustername                      string `json:"clustername"`
	Clustertype                      string `json:"clustertype"`
	Cpuallocated                     string `json:"cpuallocated"`
	Cpuallocatedpercentage           string `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64  `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string `json:"cpuallocatedwithoverprovisioning"`
	Cpunumber                        int    `json:"cpunumber"`
	Cpuspeed                         int64  `json:"cpuspeed"`
	Cpuused                          string `json:"cpuused"`
	Cpuwithoverprovisioning          string `json:"cpuwithoverprovisioning"`
	Created                          string `json:"created"`
	Disconnected                     string `json:"disconnected"`
	Disksizeallocated                int64  `json:"disksizeallocated"`
	Disksizetotal                    int64  `json:"disksizetotal"`
	Events                           string `json:"events"`
	Hahost                           bool   `json:"hahost"`
	Hasenoughcapacity                bool   `json:"hasenoughcapacity"`
	Hosttags                         string `json:"hosttags"`
	Hypervisor                       string `json:"hypervisor"`
	Hypervisorversion                string `json:"hypervisorversion"`
	Id                               string `json:"id"`
	Ipaddress                        string `json:"ipaddress"`
	Islocalstorageactive             bool   `json:"islocalstorageactive"`
	JobID                            string `json:"jobid"`
	Jobstatus                        int    `json:"jobstatus"`
	Lastpinged                       string `json:"lastpinged"`
	Managementserverid               UUID   `json:"managementserverid"`
	Memoryallocated                  string `json:"memoryallocated"`
	Memoryallocatedbytes             int64  `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string `json:"memoryallocatedpercentage"`
	Memorytotal                      int64  `json:"memorytotal"`
	Memoryused                       int64  `json:"memoryused"`
	Memorywithoverprovisioning       string `json:"memorywithoverprovisioning"`
	Name                             string `json:"name"`
	Networkkbsread                   int64  `json:"networkkbsread"`
	Networkkbswrite                  int64  `json:"networkkbswrite"`
	Oscategoryid                     string `json:"oscategoryid"`
	Oscategoryname                   string `json:"oscategoryname"`
	Podid                            string `json:"podid"`
	Podname                          string `json:"podname"`
	Removed                          string `json:"removed"`
	RequiresStorageMotion            bool   `json:"requiresStorageMotion"`
	Resourcestate                    string `json:"resourcestate"`
	State                            string `json:"state"`
	Suitableformigration             bool   `json:"suitableformigration"`
	Type                             string `json:"type"`
	Version                          string `json:"version"`
	Zoneid                           string `json:"zoneid"`
	Zonename                         string `json:"zonename"`
}

type ListDedicatedHostsParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedHostsParams) toURLValues() url.Values {
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
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
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

func (p *ListDedicatedHostsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListDedicatedHostsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListDedicatedHostsParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
}

func (p *ListDedicatedHostsParams) GetAffinitygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupid"].(string)
	return value, ok
}

func (p *ListDedicatedHostsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListDedicatedHostsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListDedicatedHostsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListDedicatedHostsParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListDedicatedHostsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListDedicatedHostsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListDedicatedHostsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListDedicatedHostsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListDedicatedHostsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListDedicatedHostsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListDedicatedHostsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListDedicatedHostsParams() *ListDedicatedHostsParams {
	p := &ListDedicatedHostsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists dedicated hosts.
func (s *HostService) ListDedicatedHosts(p *ListDedicatedHostsParams) (*ListDedicatedHostsResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedHosts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedHostsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedHostsResponse struct {
	Count          int              `json:"count"`
	DedicatedHosts []*DedicatedHost `json:"dedicatedhost"`
}

type DedicatedHost struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Hostid          string `json:"hostid"`
	Hostname        string `json:"hostname"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
}

type ListHostTagsParams struct {
	p map[string]interface{}
}

func (p *ListHostTagsParams) toURLValues() url.Values {
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

func (p *ListHostTagsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListHostTagsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListHostTagsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListHostTagsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListHostTagsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListHostTagsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListHostTagsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListHostTagsParams() *ListHostTagsParams {
	p := &ListHostTagsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostTagID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListHostTagsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListHostTags(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.HostTags[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.HostTags {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists host tags
func (s *HostService) ListHostTags(p *ListHostTagsParams) (*ListHostTagsResponse, error) {
	resp, err := s.cs.newRequest("listHostTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHostTagsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHostTagsResponse struct {
	Count    int        `json:"count"`
	HostTags []*HostTag `json:"hosttag"`
}

type HostTag struct {
	Hostid    int64  `json:"hostid"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListHostsParams struct {
	p map[string]interface{}
}

func (p *ListHostsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := p.p["hahost"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hahost", vv)
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	if v, found := p.p["outofbandmanagementenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("outofbandmanagementenabled", vv)
	}
	if v, found := p.p["outofbandmanagementpowerstate"]; found {
		u.Set("outofbandmanagementpowerstate", v.(string))
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
	if v, found := p.p["resourcestate"]; found {
		u.Set("resourcestate", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListHostsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListHostsParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListHostsParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListHostsParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListHostsParams) SetHahost(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hahost"] = v
}

func (p *ListHostsParams) GetHahost() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hahost"].(bool)
	return value, ok
}

func (p *ListHostsParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListHostsParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListHostsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListHostsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListHostsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListHostsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListHostsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListHostsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListHostsParams) SetOutofbandmanagementenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["outofbandmanagementenabled"] = v
}

func (p *ListHostsParams) GetOutofbandmanagementenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["outofbandmanagementenabled"].(bool)
	return value, ok
}

func (p *ListHostsParams) SetOutofbandmanagementpowerstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["outofbandmanagementpowerstate"] = v
}

func (p *ListHostsParams) GetOutofbandmanagementpowerstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["outofbandmanagementpowerstate"].(string)
	return value, ok
}

func (p *ListHostsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListHostsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListHostsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListHostsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListHostsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListHostsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListHostsParams) SetResourcestate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcestate"] = v
}

func (p *ListHostsParams) GetResourcestate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcestate"].(string)
	return value, ok
}

func (p *ListHostsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListHostsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListHostsParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListHostsParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *ListHostsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListHostsParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListHostsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListHostsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListHostsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListHostsParams() *ListHostsParams {
	p := &ListHostsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListHostsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListHosts(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Hosts[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Hosts {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostByName(name string, opts ...OptionFunc) (*Host, int, error) {
	id, count, err := s.GetHostID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetHostByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostByID(id string, opts ...OptionFunc) (*Host, int, error) {
	p := &ListHostsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListHosts(p)
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
		return l.Hosts[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Host UUID: %s!", id)
}

// Lists hosts.
func (s *HostService) ListHosts(p *ListHostsParams) (*ListHostsResponse, error) {
	resp, err := s.cs.newRequest("listHosts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHostsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHostsResponse struct {
	Count int     `json:"count"`
	Hosts []*Host `json:"host"`
}

type Host struct {
	Annotation                       string                      `json:"annotation"`
	Capabilities                     string                      `json:"capabilities"`
	Clusterid                        string                      `json:"clusterid"`
	Clustername                      string                      `json:"clustername"`
	Clustertype                      string                      `json:"clustertype"`
	Cpuallocated                     string                      `json:"cpuallocated"`
	Cpuallocatedpercentage           string                      `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                       `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                      `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                     `json:"cpuloadaverage"`
	Cpunumber                        int                         `json:"cpunumber"`
	Cpusockets                       int                         `json:"cpusockets"`
	Cpuspeed                         int64                       `json:"cpuspeed"`
	Cpuused                          string                      `json:"cpuused"`
	Cpuwithoverprovisioning          string                      `json:"cpuwithoverprovisioning"`
	Created                          string                      `json:"created"`
	Details                          map[string]string           `json:"details"`
	Disconnected                     string                      `json:"disconnected"`
	Disksizeallocated                int64                       `json:"disksizeallocated"`
	Disksizetotal                    int64                       `json:"disksizetotal"`
	Encryptionsupported              bool                        `json:"encryptionsupported"`
	Events                           string                      `json:"events"`
	Gpugroup                         []HostGpugroup              `json:"gpugroup"`
	Hahost                           bool                        `json:"hahost"`
	Hasannotations                   bool                        `json:"hasannotations"`
	Hasenoughcapacity                bool                        `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse           `json:"hostha"`
	Hosttags                         string                      `json:"hosttags"`
	Hypervisor                       string                      `json:"hypervisor"`
	Hypervisorversion                string                      `json:"hypervisorversion"`
	Id                               string                      `json:"id"`
	Ipaddress                        string                      `json:"ipaddress"`
	Islocalstorageactive             bool                        `json:"islocalstorageactive"`
	JobID                            string                      `json:"jobid"`
	Jobstatus                        int                         `json:"jobstatus"`
	Lastannotated                    string                      `json:"lastannotated"`
	Lastpinged                       string                      `json:"lastpinged"`
	Managementserverid               UUID                        `json:"managementserverid"`
	Memoryallocated                  int64                       `json:"memoryallocated"`
	Memoryallocatedbytes             int64                       `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                      `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                       `json:"memorytotal"`
	Memoryused                       int64                       `json:"memoryused"`
	Memorywithoverprovisioning       string                      `json:"memorywithoverprovisioning"`
	Name                             string                      `json:"name"`
	Networkkbsread                   int64                       `json:"networkkbsread"`
	Networkkbswrite                  int64                       `json:"networkkbswrite"`
	Oscategoryid                     string                      `json:"oscategoryid"`
	Oscategoryname                   string                      `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse `json:"outofbandmanagement"`
	Podid                            string                      `json:"podid"`
	Podname                          string                      `json:"podname"`
	Removed                          string                      `json:"removed"`
	Resourcestate                    string                      `json:"resourcestate"`
	State                            string                      `json:"state"`
	Suitableformigration             bool                        `json:"suitableformigration"`
	Type                             string                      `json:"type"`
	Ueficapability                   bool                        `json:"ueficapability"`
	Username                         string                      `json:"username"`
	Version                          string                      `json:"version"`
	Zoneid                           string                      `json:"zoneid"`
	Zonename                         string                      `json:"zonename"`
}

type HostGpugroup struct {
	Gpugroupname string             `json:"gpugroupname"`
	Vgpu         []HostGpugroupVgpu `json:"vgpu"`
}

type HostGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type ListHostsMetricsParams struct {
	p map[string]interface{}
}

func (p *ListHostsMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := p.p["hahost"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hahost", vv)
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	if v, found := p.p["outofbandmanagementenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("outofbandmanagementenabled", vv)
	}
	if v, found := p.p["outofbandmanagementpowerstate"]; found {
		u.Set("outofbandmanagementpowerstate", v.(string))
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
	if v, found := p.p["resourcestate"]; found {
		u.Set("resourcestate", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListHostsMetricsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListHostsMetricsParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListHostsMetricsParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetHahost(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hahost"] = v
}

func (p *ListHostsMetricsParams) GetHahost() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hahost"].(bool)
	return value, ok
}

func (p *ListHostsMetricsParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListHostsMetricsParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListHostsMetricsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListHostsMetricsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListHostsMetricsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetOutofbandmanagementenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["outofbandmanagementenabled"] = v
}

func (p *ListHostsMetricsParams) GetOutofbandmanagementenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["outofbandmanagementenabled"].(bool)
	return value, ok
}

func (p *ListHostsMetricsParams) SetOutofbandmanagementpowerstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["outofbandmanagementpowerstate"] = v
}

func (p *ListHostsMetricsParams) GetOutofbandmanagementpowerstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["outofbandmanagementpowerstate"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListHostsMetricsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListHostsMetricsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListHostsMetricsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListHostsMetricsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListHostsMetricsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetResourcestate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcestate"] = v
}

func (p *ListHostsMetricsParams) GetResourcestate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcestate"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListHostsMetricsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListHostsMetricsParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListHostsMetricsParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListHostsMetricsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListHostsMetricsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListHostsMetricsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListHostsMetricsParams() *ListHostsMetricsParams {
	p := &ListHostsMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostsMetricID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListHostsMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListHostsMetrics(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.HostsMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.HostsMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostsMetricByName(name string, opts ...OptionFunc) (*HostsMetric, int, error) {
	id, count, err := s.GetHostsMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetHostsMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostsMetricByID(id string, opts ...OptionFunc) (*HostsMetric, int, error) {
	p := &ListHostsMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListHostsMetrics(p)
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
		return l.HostsMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for HostsMetric UUID: %s!", id)
}

// Lists hosts metrics
func (s *HostService) ListHostsMetrics(p *ListHostsMetricsParams) (*ListHostsMetricsResponse, error) {
	resp, err := s.cs.newRequest("listHostsMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHostsMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHostsMetricsResponse struct {
	Count        int            `json:"count"`
	HostsMetrics []*HostsMetric `json:"hostsmetric"`
}

type HostsMetric struct {
	Annotation                       string                      `json:"annotation"`
	Capabilities                     string                      `json:"capabilities"`
	Clusterid                        string                      `json:"clusterid"`
	Clustername                      string                      `json:"clustername"`
	Clustertype                      string                      `json:"clustertype"`
	Cpuallocated                     string                      `json:"cpuallocated"`
	Cpuallocateddisablethreshold     bool                        `json:"cpuallocateddisablethreshold"`
	Cpuallocatedghz                  string                      `json:"cpuallocatedghz"`
	Cpuallocatedpercentage           string                      `json:"cpuallocatedpercentage"`
	Cpuallocatedthreshold            bool                        `json:"cpuallocatedthreshold"`
	Cpuallocatedvalue                int64                       `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                      `json:"cpuallocatedwithoverprovisioning"`
	Cpudisablethreshold              bool                        `json:"cpudisablethreshold"`
	Cpuloadaverage                   float64                     `json:"cpuloadaverage"`
	Cpunumber                        int                         `json:"cpunumber"`
	Cpusockets                       int                         `json:"cpusockets"`
	Cpuspeed                         int64                       `json:"cpuspeed"`
	Cputhreshold                     bool                        `json:"cputhreshold"`
	Cputotalghz                      string                      `json:"cputotalghz"`
	Cpuused                          string                      `json:"cpuused"`
	Cpuusedghz                       string                      `json:"cpuusedghz"`
	Cpuwithoverprovisioning          string                      `json:"cpuwithoverprovisioning"`
	Created                          string                      `json:"created"`
	Details                          map[string]string           `json:"details"`
	Disconnected                     string                      `json:"disconnected"`
	Disksizeallocated                int64                       `json:"disksizeallocated"`
	Disksizetotal                    int64                       `json:"disksizetotal"`
	Encryptionsupported              bool                        `json:"encryptionsupported"`
	Events                           string                      `json:"events"`
	Gpugroup                         []HostsMetricGpugroup       `json:"gpugroup"`
	Hahost                           bool                        `json:"hahost"`
	Hasannotations                   bool                        `json:"hasannotations"`
	Hasenoughcapacity                bool                        `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse           `json:"hostha"`
	Hosttags                         string                      `json:"hosttags"`
	Hypervisor                       string                      `json:"hypervisor"`
	Hypervisorversion                string                      `json:"hypervisorversion"`
	Id                               string                      `json:"id"`
	Instances                        string                      `json:"instances"`
	Ipaddress                        string                      `json:"ipaddress"`
	Islocalstorageactive             bool                        `json:"islocalstorageactive"`
	JobID                            string                      `json:"jobid"`
	Jobstatus                        int                         `json:"jobstatus"`
	Lastannotated                    string                      `json:"lastannotated"`
	Lastpinged                       string                      `json:"lastpinged"`
	Managementserverid               UUID                        `json:"managementserverid"`
	Memoryallocated                  int64                       `json:"memoryallocated"`
	Memoryallocatedbytes             int64                       `json:"memoryallocatedbytes"`
	Memoryallocateddisablethreshold  bool                        `json:"memoryallocateddisablethreshold"`
	Memoryallocatedgb                string                      `json:"memoryallocatedgb"`
	Memoryallocatedpercentage        string                      `json:"memoryallocatedpercentage"`
	Memoryallocatedthreshold         bool                        `json:"memoryallocatedthreshold"`
	Memorydisablethreshold           bool                        `json:"memorydisablethreshold"`
	Memorythreshold                  bool                        `json:"memorythreshold"`
	Memorytotal                      int64                       `json:"memorytotal"`
	Memorytotalgb                    string                      `json:"memorytotalgb"`
	Memoryused                       int64                       `json:"memoryused"`
	Memoryusedgb                     string                      `json:"memoryusedgb"`
	Memorywithoverprovisioning       string                      `json:"memorywithoverprovisioning"`
	Name                             string                      `json:"name"`
	Networkkbsread                   int64                       `json:"networkkbsread"`
	Networkkbswrite                  int64                       `json:"networkkbswrite"`
	Networkread                      string                      `json:"networkread"`
	Networkwrite                     string                      `json:"networkwrite"`
	Oscategoryid                     string                      `json:"oscategoryid"`
	Oscategoryname                   string                      `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse `json:"outofbandmanagement"`
	Podid                            string                      `json:"podid"`
	Podname                          string                      `json:"podname"`
	Powerstate                       string                      `json:"powerstate"`
	Removed                          string                      `json:"removed"`
	Resourcestate                    string                      `json:"resourcestate"`
	State                            string                      `json:"state"`
	Suitableformigration             bool                        `json:"suitableformigration"`
	Type                             string                      `json:"type"`
	Ueficapability                   bool                        `json:"ueficapability"`
	Username                         string                      `json:"username"`
	Version                          string                      `json:"version"`
	Zoneid                           string                      `json:"zoneid"`
	Zonename                         string                      `json:"zonename"`
}

type HostsMetricGpugroup struct {
	Gpugroupname string                    `json:"gpugroupname"`
	Vgpu         []HostsMetricGpugroupVgpu `json:"vgpu"`
}

type HostsMetricGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type PrepareHostForMaintenanceParams struct {
	p map[string]interface{}
}

func (p *PrepareHostForMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *PrepareHostForMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *PrepareHostForMaintenanceParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new PrepareHostForMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewPrepareHostForMaintenanceParams(id string) *PrepareHostForMaintenanceParams {
	p := &PrepareHostForMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Prepares a host for maintenance.
func (s *HostService) PrepareHostForMaintenance(p *PrepareHostForMaintenanceParams) (*PrepareHostForMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("prepareHostForMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PrepareHostForMaintenanceResponse
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

type PrepareHostForMaintenanceResponse struct {
	Annotation                       string                                      `json:"annotation"`
	Capabilities                     string                                      `json:"capabilities"`
	Clusterid                        string                                      `json:"clusterid"`
	Clustername                      string                                      `json:"clustername"`
	Clustertype                      string                                      `json:"clustertype"`
	Cpuallocated                     string                                      `json:"cpuallocated"`
	Cpuallocatedpercentage           string                                      `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                                       `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                                      `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                                     `json:"cpuloadaverage"`
	Cpunumber                        int                                         `json:"cpunumber"`
	Cpusockets                       int                                         `json:"cpusockets"`
	Cpuspeed                         int64                                       `json:"cpuspeed"`
	Cpuused                          string                                      `json:"cpuused"`
	Cpuwithoverprovisioning          string                                      `json:"cpuwithoverprovisioning"`
	Created                          string                                      `json:"created"`
	Details                          map[string]string                           `json:"details"`
	Disconnected                     string                                      `json:"disconnected"`
	Disksizeallocated                int64                                       `json:"disksizeallocated"`
	Disksizetotal                    int64                                       `json:"disksizetotal"`
	Encryptionsupported              bool                                        `json:"encryptionsupported"`
	Events                           string                                      `json:"events"`
	Gpugroup                         []PrepareHostForMaintenanceResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                                        `json:"hahost"`
	Hasannotations                   bool                                        `json:"hasannotations"`
	Hasenoughcapacity                bool                                        `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse                           `json:"hostha"`
	Hosttags                         string                                      `json:"hosttags"`
	Hypervisor                       string                                      `json:"hypervisor"`
	Hypervisorversion                string                                      `json:"hypervisorversion"`
	Id                               string                                      `json:"id"`
	Ipaddress                        string                                      `json:"ipaddress"`
	Islocalstorageactive             bool                                        `json:"islocalstorageactive"`
	JobID                            string                                      `json:"jobid"`
	Jobstatus                        int                                         `json:"jobstatus"`
	Lastannotated                    string                                      `json:"lastannotated"`
	Lastpinged                       string                                      `json:"lastpinged"`
	Managementserverid               UUID                                        `json:"managementserverid"`
	Memoryallocated                  int64                                       `json:"memoryallocated"`
	Memoryallocatedbytes             int64                                       `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                                      `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                                       `json:"memorytotal"`
	Memoryused                       int64                                       `json:"memoryused"`
	Memorywithoverprovisioning       string                                      `json:"memorywithoverprovisioning"`
	Name                             string                                      `json:"name"`
	Networkkbsread                   int64                                       `json:"networkkbsread"`
	Networkkbswrite                  int64                                       `json:"networkkbswrite"`
	Oscategoryid                     string                                      `json:"oscategoryid"`
	Oscategoryname                   string                                      `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse                 `json:"outofbandmanagement"`
	Podid                            string                                      `json:"podid"`
	Podname                          string                                      `json:"podname"`
	Removed                          string                                      `json:"removed"`
	Resourcestate                    string                                      `json:"resourcestate"`
	State                            string                                      `json:"state"`
	Suitableformigration             bool                                        `json:"suitableformigration"`
	Type                             string                                      `json:"type"`
	Ueficapability                   bool                                        `json:"ueficapability"`
	Username                         string                                      `json:"username"`
	Version                          string                                      `json:"version"`
	Zoneid                           string                                      `json:"zoneid"`
	Zonename                         string                                      `json:"zonename"`
}

type PrepareHostForMaintenanceResponseGpugroup struct {
	Gpugroupname string                                          `json:"gpugroupname"`
	Vgpu         []PrepareHostForMaintenanceResponseGpugroupVgpu `json:"vgpu"`
}

type PrepareHostForMaintenanceResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type ReconnectHostParams struct {
	p map[string]interface{}
}

func (p *ReconnectHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ReconnectHostParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ReconnectHostParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReconnectHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReconnectHostParams(id string) *ReconnectHostParams {
	p := &ReconnectHostParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reconnects a host.
func (s *HostService) ReconnectHost(p *ReconnectHostParams) (*ReconnectHostResponse, error) {
	resp, err := s.cs.newRequest("reconnectHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReconnectHostResponse
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

type ReconnectHostResponse struct {
	Annotation                       string                          `json:"annotation"`
	Capabilities                     string                          `json:"capabilities"`
	Clusterid                        string                          `json:"clusterid"`
	Clustername                      string                          `json:"clustername"`
	Clustertype                      string                          `json:"clustertype"`
	Cpuallocated                     string                          `json:"cpuallocated"`
	Cpuallocatedpercentage           string                          `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                           `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                          `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                         `json:"cpuloadaverage"`
	Cpunumber                        int                             `json:"cpunumber"`
	Cpusockets                       int                             `json:"cpusockets"`
	Cpuspeed                         int64                           `json:"cpuspeed"`
	Cpuused                          string                          `json:"cpuused"`
	Cpuwithoverprovisioning          string                          `json:"cpuwithoverprovisioning"`
	Created                          string                          `json:"created"`
	Details                          map[string]string               `json:"details"`
	Disconnected                     string                          `json:"disconnected"`
	Disksizeallocated                int64                           `json:"disksizeallocated"`
	Disksizetotal                    int64                           `json:"disksizetotal"`
	Encryptionsupported              bool                            `json:"encryptionsupported"`
	Events                           string                          `json:"events"`
	Gpugroup                         []ReconnectHostResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                            `json:"hahost"`
	Hasannotations                   bool                            `json:"hasannotations"`
	Hasenoughcapacity                bool                            `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse               `json:"hostha"`
	Hosttags                         string                          `json:"hosttags"`
	Hypervisor                       string                          `json:"hypervisor"`
	Hypervisorversion                string                          `json:"hypervisorversion"`
	Id                               string                          `json:"id"`
	Ipaddress                        string                          `json:"ipaddress"`
	Islocalstorageactive             bool                            `json:"islocalstorageactive"`
	JobID                            string                          `json:"jobid"`
	Jobstatus                        int                             `json:"jobstatus"`
	Lastannotated                    string                          `json:"lastannotated"`
	Lastpinged                       string                          `json:"lastpinged"`
	Managementserverid               UUID                            `json:"managementserverid"`
	Memoryallocated                  int64                           `json:"memoryallocated"`
	Memoryallocatedbytes             int64                           `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                          `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                           `json:"memorytotal"`
	Memoryused                       int64                           `json:"memoryused"`
	Memorywithoverprovisioning       string                          `json:"memorywithoverprovisioning"`
	Name                             string                          `json:"name"`
	Networkkbsread                   int64                           `json:"networkkbsread"`
	Networkkbswrite                  int64                           `json:"networkkbswrite"`
	Oscategoryid                     string                          `json:"oscategoryid"`
	Oscategoryname                   string                          `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse     `json:"outofbandmanagement"`
	Podid                            string                          `json:"podid"`
	Podname                          string                          `json:"podname"`
	Removed                          string                          `json:"removed"`
	Resourcestate                    string                          `json:"resourcestate"`
	State                            string                          `json:"state"`
	Suitableformigration             bool                            `json:"suitableformigration"`
	Type                             string                          `json:"type"`
	Ueficapability                   bool                            `json:"ueficapability"`
	Username                         string                          `json:"username"`
	Version                          string                          `json:"version"`
	Zoneid                           string                          `json:"zoneid"`
	Zonename                         string                          `json:"zonename"`
}

type ReconnectHostResponseGpugroup struct {
	Gpugroupname string                              `json:"gpugroupname"`
	Vgpu         []ReconnectHostResponseGpugroupVgpu `json:"vgpu"`
}

type ReconnectHostResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type ReleaseDedicatedHostParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedHostParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ReleaseDedicatedHostParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReleaseDedicatedHostParams(hostid string) *ReleaseDedicatedHostParams {
	p := &ReleaseDedicatedHostParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	return p
}

// Release the dedication for host
func (s *HostService) ReleaseDedicatedHost(p *ReleaseDedicatedHostParams) (*ReleaseDedicatedHostResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedHostResponse
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

type ReleaseDedicatedHostResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ReleaseHostReservationParams struct {
	p map[string]interface{}
}

func (p *ReleaseHostReservationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ReleaseHostReservationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ReleaseHostReservationParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseHostReservationParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReleaseHostReservationParams(id string) *ReleaseHostReservationParams {
	p := &ReleaseHostReservationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Releases host reservation.
func (s *HostService) ReleaseHostReservation(p *ReleaseHostReservationParams) (*ReleaseHostReservationResponse, error) {
	resp, err := s.cs.newRequest("releaseHostReservation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseHostReservationResponse
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

type ReleaseHostReservationResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateHostParams struct {
	p map[string]interface{}
}

func (p *UpdateHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["annotation"]; found {
		u.Set("annotation", v.(string))
	}
	if v, found := p.p["hosttags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("hosttags", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	return u
}

func (p *UpdateHostParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *UpdateHostParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *UpdateHostParams) SetAnnotation(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["annotation"] = v
}

func (p *UpdateHostParams) GetAnnotation() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["annotation"].(string)
	return value, ok
}

func (p *UpdateHostParams) SetHosttags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
}

func (p *UpdateHostParams) GetHosttags() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hosttags"].([]string)
	return value, ok
}

func (p *UpdateHostParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateHostParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateHostParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateHostParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateHostParams) SetOscategoryid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["oscategoryid"] = v
}

func (p *UpdateHostParams) GetOscategoryid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["oscategoryid"].(string)
	return value, ok
}

func (p *UpdateHostParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *UpdateHostParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewUpdateHostParams(id string) *UpdateHostParams {
	p := &UpdateHostParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a host.
func (s *HostService) UpdateHost(p *UpdateHostParams) (*UpdateHostResponse, error) {
	resp, err := s.cs.newRequest("updateHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateHostResponse struct {
	Annotation                       string                       `json:"annotation"`
	Capabilities                     string                       `json:"capabilities"`
	Clusterid                        string                       `json:"clusterid"`
	Clustername                      string                       `json:"clustername"`
	Clustertype                      string                       `json:"clustertype"`
	Cpuallocated                     string                       `json:"cpuallocated"`
	Cpuallocatedpercentage           string                       `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                        `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                       `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                      `json:"cpuloadaverage"`
	Cpunumber                        int                          `json:"cpunumber"`
	Cpusockets                       int                          `json:"cpusockets"`
	Cpuspeed                         int64                        `json:"cpuspeed"`
	Cpuused                          string                       `json:"cpuused"`
	Cpuwithoverprovisioning          string                       `json:"cpuwithoverprovisioning"`
	Created                          string                       `json:"created"`
	Details                          map[string]string            `json:"details"`
	Disconnected                     string                       `json:"disconnected"`
	Disksizeallocated                int64                        `json:"disksizeallocated"`
	Disksizetotal                    int64                        `json:"disksizetotal"`
	Encryptionsupported              bool                         `json:"encryptionsupported"`
	Events                           string                       `json:"events"`
	Gpugroup                         []UpdateHostResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                         `json:"hahost"`
	Hasannotations                   bool                         `json:"hasannotations"`
	Hasenoughcapacity                bool                         `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse            `json:"hostha"`
	Hosttags                         string                       `json:"hosttags"`
	Hypervisor                       string                       `json:"hypervisor"`
	Hypervisorversion                string                       `json:"hypervisorversion"`
	Id                               string                       `json:"id"`
	Ipaddress                        string                       `json:"ipaddress"`
	Islocalstorageactive             bool                         `json:"islocalstorageactive"`
	JobID                            string                       `json:"jobid"`
	Jobstatus                        int                          `json:"jobstatus"`
	Lastannotated                    string                       `json:"lastannotated"`
	Lastpinged                       string                       `json:"lastpinged"`
	Managementserverid               UUID                         `json:"managementserverid"`
	Memoryallocated                  int64                        `json:"memoryallocated"`
	Memoryallocatedbytes             int64                        `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                       `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                        `json:"memorytotal"`
	Memoryused                       int64                        `json:"memoryused"`
	Memorywithoverprovisioning       string                       `json:"memorywithoverprovisioning"`
	Name                             string                       `json:"name"`
	Networkkbsread                   int64                        `json:"networkkbsread"`
	Networkkbswrite                  int64                        `json:"networkkbswrite"`
	Oscategoryid                     string                       `json:"oscategoryid"`
	Oscategoryname                   string                       `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse  `json:"outofbandmanagement"`
	Podid                            string                       `json:"podid"`
	Podname                          string                       `json:"podname"`
	Removed                          string                       `json:"removed"`
	Resourcestate                    string                       `json:"resourcestate"`
	State                            string                       `json:"state"`
	Suitableformigration             bool                         `json:"suitableformigration"`
	Type                             string                       `json:"type"`
	Ueficapability                   bool                         `json:"ueficapability"`
	Username                         string                       `json:"username"`
	Version                          string                       `json:"version"`
	Zoneid                           string                       `json:"zoneid"`
	Zonename                         string                       `json:"zonename"`
}

type UpdateHostResponseGpugroup struct {
	Gpugroupname string                           `json:"gpugroupname"`
	Vgpu         []UpdateHostResponseGpugroupVgpu `json:"vgpu"`
}

type UpdateHostResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type UpdateHostPasswordParams struct {
	p map[string]interface{}
}

func (p *UpdateHostPasswordParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["update_passwd_on_host"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("update_passwd_on_host", vv)
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *UpdateHostPasswordParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *UpdateHostPasswordParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *UpdateHostPasswordParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *UpdateHostPasswordParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *UpdateHostPasswordParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *UpdateHostPasswordParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *UpdateHostPasswordParams) SetUpdate_passwd_on_host(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["update_passwd_on_host"] = v
}

func (p *UpdateHostPasswordParams) GetUpdate_passwd_on_host() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["update_passwd_on_host"].(bool)
	return value, ok
}

func (p *UpdateHostPasswordParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *UpdateHostPasswordParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateHostPasswordParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewUpdateHostPasswordParams(password string, username string) *UpdateHostPasswordParams {
	p := &UpdateHostPasswordParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Update password of a host/pool on management server.
func (s *HostService) UpdateHostPassword(p *UpdateHostPasswordParams) (*UpdateHostPasswordResponse, error) {
	resp, err := s.cs.newRequest("updateHostPassword", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateHostPasswordResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateHostPasswordResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateHostPasswordResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateHostPasswordResponse
	return json.Unmarshal(b, (*alias)(r))
}
