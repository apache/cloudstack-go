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

type ClusterServiceIface interface {
	AddCluster(p *AddClusterParams) (*AddClusterResponse, error)
	NewAddClusterParams(clustername string, clustertype string, hypervisor string, podid string, zoneid string) *AddClusterParams
	DedicateCluster(p *DedicateClusterParams) (*DedicateClusterResponse, error)
	NewDedicateClusterParams(clusterid string, domainid string) *DedicateClusterParams
	DeleteCluster(p *DeleteClusterParams) (*DeleteClusterResponse, error)
	NewDeleteClusterParams(id string) *DeleteClusterParams
	DisableOutOfBandManagementForCluster(p *DisableOutOfBandManagementForClusterParams) (*DisableOutOfBandManagementForClusterResponse, error)
	NewDisableOutOfBandManagementForClusterParams(clusterid string) *DisableOutOfBandManagementForClusterParams
	EnableOutOfBandManagementForCluster(p *EnableOutOfBandManagementForClusterParams) (*EnableOutOfBandManagementForClusterResponse, error)
	NewEnableOutOfBandManagementForClusterParams(clusterid string) *EnableOutOfBandManagementForClusterParams
	EnableHAForCluster(p *EnableHAForClusterParams) (*EnableHAForClusterResponse, error)
	NewEnableHAForClusterParams(clusterid string) *EnableHAForClusterParams
	ExecuteClusterDrsPlan(p *ExecuteClusterDrsPlanParams) (*ExecuteClusterDrsPlanResponse, error)
	NewExecuteClusterDrsPlanParams(id string) *ExecuteClusterDrsPlanParams
	GenerateClusterDrsPlan(p *GenerateClusterDrsPlanParams) (*GenerateClusterDrsPlanResponse, error)
	NewGenerateClusterDrsPlanParams(id string) *GenerateClusterDrsPlanParams
	DisableHAForCluster(p *DisableHAForClusterParams) (*DisableHAForClusterResponse, error)
	NewDisableHAForClusterParams(clusterid string) *DisableHAForClusterParams
	ListClusters(p *ListClustersParams) (*ListClustersResponse, error)
	NewListClustersParams() *ListClustersParams
	GetClusterID(name string, opts ...OptionFunc) (string, int, error)
	GetClusterByName(name string, opts ...OptionFunc) (*Cluster, int, error)
	GetClusterByID(id string, opts ...OptionFunc) (*Cluster, int, error)
	ListClusterDrsPlan(p *ListClusterDrsPlanParams) (*ListClusterDrsPlanResponse, error)
	NewListClusterDrsPlanParams() *ListClusterDrsPlanParams
	GetClusterDrsPlanByID(id string, opts ...OptionFunc) (*ClusterDrsPlan, int, error)
	ListClustersMetrics(p *ListClustersMetricsParams) (*ListClustersMetricsResponse, error)
	NewListClustersMetricsParams() *ListClustersMetricsParams
	GetClustersMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetClustersMetricByName(name string, opts ...OptionFunc) (*ClustersMetric, int, error)
	GetClustersMetricByID(id string, opts ...OptionFunc) (*ClustersMetric, int, error)
	ListDedicatedClusters(p *ListDedicatedClustersParams) (*ListDedicatedClustersResponse, error)
	NewListDedicatedClustersParams() *ListDedicatedClustersParams
	ReleaseDedicatedCluster(p *ReleaseDedicatedClusterParams) (*ReleaseDedicatedClusterResponse, error)
	NewReleaseDedicatedClusterParams(clusterid string) *ReleaseDedicatedClusterParams
	UpdateCluster(p *UpdateClusterParams) (*UpdateClusterResponse, error)
	NewUpdateClusterParams(id string) *UpdateClusterParams
}

type AddClusterParams struct {
	p map[string]interface{}
}

func (p *AddClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["extensionid"]; found {
		u.Set("extensionid", v.(string))
	}
	if v, found := p.p["externaldetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("externaldetails[%d].key", i), k)
			u.Set(fmt.Sprintf("externaldetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["guestvswitchname"]; found {
		u.Set("guestvswitchname", v.(string))
	}
	if v, found := p.p["guestvswitchtype"]; found {
		u.Set("guestvswitchtype", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["ovm3cluster"]; found {
		u.Set("ovm3cluster", v.(string))
	}
	if v, found := p.p["ovm3pool"]; found {
		u.Set("ovm3pool", v.(string))
	}
	if v, found := p.p["ovm3vip"]; found {
		u.Set("ovm3vip", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["publicvswitchname"]; found {
		u.Set("publicvswitchname", v.(string))
	}
	if v, found := p.p["publicvswitchtype"]; found {
		u.Set("publicvswitchtype", v.(string))
	}
	if v, found := p.p["storageaccessgroups"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("storageaccessgroups", vv)
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["vsmipaddress"]; found {
		u.Set("vsmipaddress", v.(string))
	}
	if v, found := p.p["vsmpassword"]; found {
		u.Set("vsmpassword", v.(string))
	}
	if v, found := p.p["vsmusername"]; found {
		u.Set("vsmusername", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddClusterParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *AddClusterParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *AddClusterParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *AddClusterParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *AddClusterParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *AddClusterParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *AddClusterParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
}

func (p *AddClusterParams) ResetClustername() {
	if p.p != nil && p.p["clustername"] != nil {
		delete(p.p, "clustername")
	}
}

func (p *AddClusterParams) GetClustername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustername"].(string)
	return value, ok
}

func (p *AddClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *AddClusterParams) ResetClustertype() {
	if p.p != nil && p.p["clustertype"] != nil {
		delete(p.p, "clustertype")
	}
}

func (p *AddClusterParams) GetClustertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustertype"].(string)
	return value, ok
}

func (p *AddClusterParams) SetExtensionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extensionid"] = v
}

func (p *AddClusterParams) ResetExtensionid() {
	if p.p != nil && p.p["extensionid"] != nil {
		delete(p.p, "extensionid")
	}
}

func (p *AddClusterParams) GetExtensionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extensionid"].(string)
	return value, ok
}

func (p *AddClusterParams) SetExternaldetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["externaldetails"] = v
}

func (p *AddClusterParams) ResetExternaldetails() {
	if p.p != nil && p.p["externaldetails"] != nil {
		delete(p.p, "externaldetails")
	}
}

func (p *AddClusterParams) GetExternaldetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["externaldetails"].(map[string]string)
	return value, ok
}

func (p *AddClusterParams) SetGuestvswitchname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestvswitchname"] = v
}

func (p *AddClusterParams) ResetGuestvswitchname() {
	if p.p != nil && p.p["guestvswitchname"] != nil {
		delete(p.p, "guestvswitchname")
	}
}

func (p *AddClusterParams) GetGuestvswitchname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["guestvswitchname"].(string)
	return value, ok
}

func (p *AddClusterParams) SetGuestvswitchtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestvswitchtype"] = v
}

func (p *AddClusterParams) ResetGuestvswitchtype() {
	if p.p != nil && p.p["guestvswitchtype"] != nil {
		delete(p.p, "guestvswitchtype")
	}
}

func (p *AddClusterParams) GetGuestvswitchtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["guestvswitchtype"].(string)
	return value, ok
}

func (p *AddClusterParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *AddClusterParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *AddClusterParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *AddClusterParams) SetOvm3cluster(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ovm3cluster"] = v
}

func (p *AddClusterParams) ResetOvm3cluster() {
	if p.p != nil && p.p["ovm3cluster"] != nil {
		delete(p.p, "ovm3cluster")
	}
}

func (p *AddClusterParams) GetOvm3cluster() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ovm3cluster"].(string)
	return value, ok
}

func (p *AddClusterParams) SetOvm3pool(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ovm3pool"] = v
}

func (p *AddClusterParams) ResetOvm3pool() {
	if p.p != nil && p.p["ovm3pool"] != nil {
		delete(p.p, "ovm3pool")
	}
}

func (p *AddClusterParams) GetOvm3pool() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ovm3pool"].(string)
	return value, ok
}

func (p *AddClusterParams) SetOvm3vip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ovm3vip"] = v
}

func (p *AddClusterParams) ResetOvm3vip() {
	if p.p != nil && p.p["ovm3vip"] != nil {
		delete(p.p, "ovm3vip")
	}
}

func (p *AddClusterParams) GetOvm3vip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ovm3vip"].(string)
	return value, ok
}

func (p *AddClusterParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *AddClusterParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *AddClusterParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *AddClusterParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *AddClusterParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *AddClusterParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *AddClusterParams) SetPublicvswitchname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicvswitchname"] = v
}

func (p *AddClusterParams) ResetPublicvswitchname() {
	if p.p != nil && p.p["publicvswitchname"] != nil {
		delete(p.p, "publicvswitchname")
	}
}

func (p *AddClusterParams) GetPublicvswitchname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicvswitchname"].(string)
	return value, ok
}

func (p *AddClusterParams) SetPublicvswitchtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicvswitchtype"] = v
}

func (p *AddClusterParams) ResetPublicvswitchtype() {
	if p.p != nil && p.p["publicvswitchtype"] != nil {
		delete(p.p, "publicvswitchtype")
	}
}

func (p *AddClusterParams) GetPublicvswitchtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["publicvswitchtype"].(string)
	return value, ok
}

func (p *AddClusterParams) SetStorageaccessgroups(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageaccessgroups"] = v
}

func (p *AddClusterParams) ResetStorageaccessgroups() {
	if p.p != nil && p.p["storageaccessgroups"] != nil {
		delete(p.p, "storageaccessgroups")
	}
}

func (p *AddClusterParams) GetStorageaccessgroups() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageaccessgroups"].([]string)
	return value, ok
}

func (p *AddClusterParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddClusterParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddClusterParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddClusterParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddClusterParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddClusterParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *AddClusterParams) SetVsmipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmipaddress"] = v
}

func (p *AddClusterParams) ResetVsmipaddress() {
	if p.p != nil && p.p["vsmipaddress"] != nil {
		delete(p.p, "vsmipaddress")
	}
}

func (p *AddClusterParams) GetVsmipaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vsmipaddress"].(string)
	return value, ok
}

func (p *AddClusterParams) SetVsmpassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmpassword"] = v
}

func (p *AddClusterParams) ResetVsmpassword() {
	if p.p != nil && p.p["vsmpassword"] != nil {
		delete(p.p, "vsmpassword")
	}
}

func (p *AddClusterParams) GetVsmpassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vsmpassword"].(string)
	return value, ok
}

func (p *AddClusterParams) SetVsmusername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vsmusername"] = v
}

func (p *AddClusterParams) ResetVsmusername() {
	if p.p != nil && p.p["vsmusername"] != nil {
		delete(p.p, "vsmusername")
	}
}

func (p *AddClusterParams) GetVsmusername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vsmusername"].(string)
	return value, ok
}

func (p *AddClusterParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddClusterParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddClusterParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewAddClusterParams(clustername string, clustertype string, hypervisor string, podid string, zoneid string) *AddClusterParams {
	p := &AddClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clustername"] = clustername
	p.p["clustertype"] = clustertype
	p.p["hypervisor"] = hypervisor
	p.p["podid"] = podid
	p.p["zoneid"] = zoneid
	return p
}

// Adds a new cluster
func (s *ClusterService) AddCluster(p *AddClusterParams) (*AddClusterResponse, error) {
	resp, err := s.cs.newPostRequest("addCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r AddClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddClusterResponse struct {
	Allocationstate         string                       `json:"allocationstate"`
	Arch                    string                       `json:"arch"`
	Capacity                []AddClusterResponseCapacity `json:"capacity"`
	Clustertype             string                       `json:"clustertype"`
	Cpuovercommitratio      string                       `json:"cpuovercommitratio"`
	Extensionid             string                       `json:"extensionid"`
	Extensionname           string                       `json:"extensionname"`
	Hasannotations          bool                         `json:"hasannotations"`
	Hypervisortype          string                       `json:"hypervisortype"`
	Id                      string                       `json:"id"`
	JobID                   string                       `json:"jobid"`
	Jobstatus               int                          `json:"jobstatus"`
	Managedstate            string                       `json:"managedstate"`
	Memoryovercommitratio   string                       `json:"memoryovercommitratio"`
	Name                    string                       `json:"name"`
	Ovm3vip                 string                       `json:"ovm3vip"`
	Podid                   string                       `json:"podid"`
	Podname                 string                       `json:"podname"`
	Podstorageaccessgroups  string                       `json:"podstorageaccessgroups"`
	Resourcedetails         map[string]string            `json:"resourcedetails"`
	Storageaccessgroups     string                       `json:"storageaccessgroups"`
	Zoneid                  string                       `json:"zoneid"`
	Zonename                string                       `json:"zonename"`
	Zonestorageaccessgroups string                       `json:"zonestorageaccessgroups"`
}

type AddClusterResponseCapacity struct {
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

type DedicateClusterParams struct {
	p map[string]interface{}
}

func (p *DedicateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	return u
}

func (p *DedicateClusterParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DedicateClusterParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DedicateClusterParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DedicateClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *DedicateClusterParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *DedicateClusterParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *DedicateClusterParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DedicateClusterParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DedicateClusterParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDedicateClusterParams(clusterid string, domainid string) *DedicateClusterParams {
	p := &DedicateClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	p.p["domainid"] = domainid
	return p
}

// Dedicate an existing cluster
func (s *ClusterService) DedicateCluster(p *DedicateClusterParams) (*DedicateClusterResponse, error) {
	resp, err := s.cs.newPostRequest("dedicateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateClusterResponse
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

type DedicateClusterResponse struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Clusterid       string `json:"clusterid"`
	Clustername     string `json:"clustername"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
}

type DeleteClusterParams struct {
	p map[string]interface{}
}

func (p *DeleteClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDeleteClusterParams(id string) *DeleteClusterParams {
	p := &DeleteClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a cluster.
func (s *ClusterService) DeleteCluster(p *DeleteClusterParams) (*DeleteClusterResponse, error) {
	resp, err := s.cs.newPostRequest("deleteCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteClusterResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteClusterResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableOutOfBandManagementForClusterParams struct {
	p map[string]interface{}
}

func (p *DisableOutOfBandManagementForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *DisableOutOfBandManagementForClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *DisableOutOfBandManagementForClusterParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *DisableOutOfBandManagementForClusterParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableOutOfBandManagementForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDisableOutOfBandManagementForClusterParams(clusterid string) *DisableOutOfBandManagementForClusterParams {
	p := &DisableOutOfBandManagementForClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Disables out-of-band management for a cluster
func (s *ClusterService) DisableOutOfBandManagementForCluster(p *DisableOutOfBandManagementForClusterParams) (*DisableOutOfBandManagementForClusterResponse, error) {
	resp, err := s.cs.newPostRequest("disableOutOfBandManagementForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableOutOfBandManagementForClusterResponse
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

type DisableOutOfBandManagementForClusterResponse struct {
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

type EnableOutOfBandManagementForClusterParams struct {
	p map[string]interface{}
}

func (p *EnableOutOfBandManagementForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *EnableOutOfBandManagementForClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *EnableOutOfBandManagementForClusterParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *EnableOutOfBandManagementForClusterParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableOutOfBandManagementForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewEnableOutOfBandManagementForClusterParams(clusterid string) *EnableOutOfBandManagementForClusterParams {
	p := &EnableOutOfBandManagementForClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Enables out-of-band management for a cluster
func (s *ClusterService) EnableOutOfBandManagementForCluster(p *EnableOutOfBandManagementForClusterParams) (*EnableOutOfBandManagementForClusterResponse, error) {
	resp, err := s.cs.newPostRequest("enableOutOfBandManagementForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableOutOfBandManagementForClusterResponse
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

type EnableOutOfBandManagementForClusterResponse struct {
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

type EnableHAForClusterParams struct {
	p map[string]interface{}
}

func (p *EnableHAForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *EnableHAForClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *EnableHAForClusterParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *EnableHAForClusterParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableHAForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewEnableHAForClusterParams(clusterid string) *EnableHAForClusterParams {
	p := &EnableHAForClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Enables HA cluster-wide
func (s *ClusterService) EnableHAForCluster(p *EnableHAForClusterParams) (*EnableHAForClusterResponse, error) {
	resp, err := s.cs.newPostRequest("enableHAForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableHAForClusterResponse
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

type EnableHAForClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ExecuteClusterDrsPlanParams struct {
	p map[string]interface{}
}

func (p *ExecuteClusterDrsPlanParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["migrateto"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("migrateto[%d].key", i), k)
			u.Set(fmt.Sprintf("migrateto[%d].value", i), m[k])
		}
	}
	return u
}

func (p *ExecuteClusterDrsPlanParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ExecuteClusterDrsPlanParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ExecuteClusterDrsPlanParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ExecuteClusterDrsPlanParams) SetMigrateto(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["migrateto"] = v
}

func (p *ExecuteClusterDrsPlanParams) ResetMigrateto() {
	if p.p != nil && p.p["migrateto"] != nil {
		delete(p.p, "migrateto")
	}
}

func (p *ExecuteClusterDrsPlanParams) GetMigrateto() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["migrateto"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ExecuteClusterDrsPlanParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewExecuteClusterDrsPlanParams(id string) *ExecuteClusterDrsPlanParams {
	p := &ExecuteClusterDrsPlanParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Execute DRS for a cluster. If there is another plan in progress for the same cluster, this command will fail.
func (s *ClusterService) ExecuteClusterDrsPlan(p *ExecuteClusterDrsPlanParams) (*ExecuteClusterDrsPlanResponse, error) {
	resp, err := s.cs.newPostRequest("executeClusterDrsPlan", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExecuteClusterDrsPlanResponse
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

type ExecuteClusterDrsPlanResponse struct {
	Clusterid  string   `json:"clusterid"`
	Eventid    string   `json:"eventid"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Migrations []string `json:"migrations"`
	Status     string   `json:"status"`
	Type       string   `json:"type"`
}

type GenerateClusterDrsPlanParams struct {
	p map[string]interface{}
}

func (p *GenerateClusterDrsPlanParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["migrations"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("migrations", vv)
	}
	return u
}

func (p *GenerateClusterDrsPlanParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *GenerateClusterDrsPlanParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *GenerateClusterDrsPlanParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *GenerateClusterDrsPlanParams) SetMigrations(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["migrations"] = v
}

func (p *GenerateClusterDrsPlanParams) ResetMigrations() {
	if p.p != nil && p.p["migrations"] != nil {
		delete(p.p, "migrations")
	}
}

func (p *GenerateClusterDrsPlanParams) GetMigrations() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["migrations"].(int)
	return value, ok
}

// You should always use this function to get a new GenerateClusterDrsPlanParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewGenerateClusterDrsPlanParams(id string) *GenerateClusterDrsPlanParams {
	p := &GenerateClusterDrsPlanParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Generate DRS plan for a cluster
func (s *ClusterService) GenerateClusterDrsPlan(p *GenerateClusterDrsPlanParams) (*GenerateClusterDrsPlanResponse, error) {
	resp, err := s.cs.newPostRequest("generateClusterDrsPlan", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GenerateClusterDrsPlanResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GenerateClusterDrsPlanResponse struct {
	Clusterid  string   `json:"clusterid"`
	Eventid    string   `json:"eventid"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Migrations []string `json:"migrations"`
	Status     string   `json:"status"`
	Type       string   `json:"type"`
}

type DisableHAForClusterParams struct {
	p map[string]interface{}
}

func (p *DisableHAForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *DisableHAForClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *DisableHAForClusterParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *DisableHAForClusterParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableHAForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDisableHAForClusterParams(clusterid string) *DisableHAForClusterParams {
	p := &DisableHAForClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Disables HA cluster-wide
func (s *ClusterService) DisableHAForCluster(p *DisableHAForClusterParams) (*DisableHAForClusterResponse, error) {
	resp, err := s.cs.newPostRequest("disableHAForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableHAForClusterResponse
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

type DisableHAForClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListClustersParams struct {
	p map[string]interface{}
}

func (p *ListClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
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
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["storageaccessgroup"]; found {
		u.Set("storageaccessgroup", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListClustersParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *ListClustersParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *ListClustersParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *ListClustersParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *ListClustersParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *ListClustersParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *ListClustersParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *ListClustersParams) ResetClustertype() {
	if p.p != nil && p.p["clustertype"] != nil {
		delete(p.p, "clustertype")
	}
}

func (p *ListClustersParams) GetClustertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustertype"].(string)
	return value, ok
}

func (p *ListClustersParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListClustersParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListClustersParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListClustersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListClustersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListClustersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListClustersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListClustersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListClustersParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
}

func (p *ListClustersParams) ResetManagedstate() {
	if p.p != nil && p.p["managedstate"] != nil {
		delete(p.p, "managedstate")
	}
}

func (p *ListClustersParams) GetManagedstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managedstate"].(string)
	return value, ok
}

func (p *ListClustersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListClustersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListClustersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListClustersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListClustersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListClustersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListClustersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListClustersParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListClustersParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListClustersParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListClustersParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
}

func (p *ListClustersParams) ResetShowcapacities() {
	if p.p != nil && p.p["showcapacities"] != nil {
		delete(p.p, "showcapacities")
	}
}

func (p *ListClustersParams) GetShowcapacities() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showcapacities"].(bool)
	return value, ok
}

func (p *ListClustersParams) SetStorageaccessgroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageaccessgroup"] = v
}

func (p *ListClustersParams) ResetStorageaccessgroup() {
	if p.p != nil && p.p["storageaccessgroup"] != nil {
		delete(p.p, "storageaccessgroup")
	}
}

func (p *ListClustersParams) GetStorageaccessgroup() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageaccessgroup"].(string)
	return value, ok
}

func (p *ListClustersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListClustersParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListClustersParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListClustersParams() *ListClustersParams {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListClusters(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Clusters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Clusters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterByName(name string, opts ...OptionFunc) (*Cluster, int, error) {
	id, count, err := s.GetClusterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetClusterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterByID(id string, opts ...OptionFunc) (*Cluster, int, error) {
	p := &ListClustersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListClusters(p)
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
		return l.Clusters[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Cluster UUID: %s!", id)
}

// Lists clusters.
func (s *ClusterService) ListClusters(p *ListClustersParams) (*ListClustersResponse, error) {
	resp, err := s.cs.newRequest("listClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListClustersResponse struct {
	Count    int        `json:"count"`
	Clusters []*Cluster `json:"cluster"`
}

type Cluster struct {
	Allocationstate         string            `json:"allocationstate"`
	Arch                    string            `json:"arch"`
	Capacity                []ClusterCapacity `json:"capacity"`
	Clustertype             string            `json:"clustertype"`
	Cpuovercommitratio      string            `json:"cpuovercommitratio"`
	Extensionid             string            `json:"extensionid"`
	Extensionname           string            `json:"extensionname"`
	Hasannotations          bool              `json:"hasannotations"`
	Hypervisortype          string            `json:"hypervisortype"`
	Id                      string            `json:"id"`
	JobID                   string            `json:"jobid"`
	Jobstatus               int               `json:"jobstatus"`
	Managedstate            string            `json:"managedstate"`
	Memoryovercommitratio   string            `json:"memoryovercommitratio"`
	Name                    string            `json:"name"`
	Ovm3vip                 string            `json:"ovm3vip"`
	Podid                   string            `json:"podid"`
	Podname                 string            `json:"podname"`
	Podstorageaccessgroups  string            `json:"podstorageaccessgroups"`
	Resourcedetails         map[string]string `json:"resourcedetails"`
	Storageaccessgroups     string            `json:"storageaccessgroups"`
	Zoneid                  string            `json:"zoneid"`
	Zonename                string            `json:"zonename"`
	Zonestorageaccessgroups string            `json:"zonestorageaccessgroups"`
}

type ClusterCapacity struct {
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

type ListClusterDrsPlanParams struct {
	p map[string]interface{}
}

func (p *ListClusterDrsPlanParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
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
	return u
}

func (p *ListClusterDrsPlanParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListClusterDrsPlanParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListClusterDrsPlanParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListClusterDrsPlanParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListClusterDrsPlanParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListClusterDrsPlanParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListClusterDrsPlanParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListClusterDrsPlanParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListClusterDrsPlanParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListClusterDrsPlanParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListClusterDrsPlanParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListClusterDrsPlanParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListClusterDrsPlanParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListClusterDrsPlanParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListClusterDrsPlanParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListClusterDrsPlanParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListClusterDrsPlanParams() *ListClusterDrsPlanParams {
	p := &ListClusterDrsPlanParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterDrsPlanByID(id string, opts ...OptionFunc) (*ClusterDrsPlan, int, error) {
	p := &ListClusterDrsPlanParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListClusterDrsPlan(p)
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
		return l.ClusterDrsPlan[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ClusterDrsPlan UUID: %s!", id)
}

// List DRS plans for a clusters
func (s *ClusterService) ListClusterDrsPlan(p *ListClusterDrsPlanParams) (*ListClusterDrsPlanResponse, error) {
	resp, err := s.cs.newRequest("listClusterDrsPlan", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListClusterDrsPlanResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListClusterDrsPlanResponse struct {
	Count          int               `json:"count"`
	ClusterDrsPlan []*ClusterDrsPlan `json:"clusterdrsplan"`
}

type ClusterDrsPlan struct {
	Clusterid  string   `json:"clusterid"`
	Eventid    string   `json:"eventid"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Migrations []string `json:"migrations"`
	Status     string   `json:"status"`
	Type       string   `json:"type"`
}

type ListClustersMetricsParams struct {
	p map[string]interface{}
}

func (p *ListClustersMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
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
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["storageaccessgroup"]; found {
		u.Set("storageaccessgroup", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListClustersMetricsParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *ListClustersMetricsParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *ListClustersMetricsParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *ListClustersMetricsParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *ListClustersMetricsParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *ListClustersMetricsParams) ResetClustertype() {
	if p.p != nil && p.p["clustertype"] != nil {
		delete(p.p, "clustertype")
	}
}

func (p *ListClustersMetricsParams) GetClustertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustertype"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListClustersMetricsParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListClustersMetricsParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListClustersMetricsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListClustersMetricsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListClustersMetricsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListClustersMetricsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
}

func (p *ListClustersMetricsParams) ResetManagedstate() {
	if p.p != nil && p.p["managedstate"] != nil {
		delete(p.p, "managedstate")
	}
}

func (p *ListClustersMetricsParams) GetManagedstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managedstate"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListClustersMetricsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListClustersMetricsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListClustersMetricsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListClustersMetricsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListClustersMetricsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListClustersMetricsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListClustersMetricsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListClustersMetricsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListClustersMetricsParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListClustersMetricsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
}

func (p *ListClustersMetricsParams) ResetShowcapacities() {
	if p.p != nil && p.p["showcapacities"] != nil {
		delete(p.p, "showcapacities")
	}
}

func (p *ListClustersMetricsParams) GetShowcapacities() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showcapacities"].(bool)
	return value, ok
}

func (p *ListClustersMetricsParams) SetStorageaccessgroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageaccessgroup"] = v
}

func (p *ListClustersMetricsParams) ResetStorageaccessgroup() {
	if p.p != nil && p.p["storageaccessgroup"] != nil {
		delete(p.p, "storageaccessgroup")
	}
}

func (p *ListClustersMetricsParams) GetStorageaccessgroup() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageaccessgroup"].(string)
	return value, ok
}

func (p *ListClustersMetricsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListClustersMetricsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListClustersMetricsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListClustersMetricsParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListClustersMetricsParams() *ListClustersMetricsParams {
	p := &ListClustersMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClustersMetricID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListClustersMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListClustersMetrics(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ClustersMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ClustersMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClustersMetricByName(name string, opts ...OptionFunc) (*ClustersMetric, int, error) {
	id, count, err := s.GetClustersMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetClustersMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClustersMetricByID(id string, opts ...OptionFunc) (*ClustersMetric, int, error) {
	p := &ListClustersMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListClustersMetrics(p)
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
		return l.ClustersMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ClustersMetric UUID: %s!", id)
}

// Lists clusters metrics
func (s *ClusterService) ListClustersMetrics(p *ListClustersMetricsParams) (*ListClustersMetricsResponse, error) {
	resp, err := s.cs.newRequest("listClustersMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListClustersMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListClustersMetricsResponse struct {
	Count           int               `json:"count"`
	ClustersMetrics []*ClustersMetric `json:"clustersmetric"`
}

type ClustersMetric struct {
	Allocationstate                 string                   `json:"allocationstate"`
	Arch                            string                   `json:"arch"`
	Capacity                        []ClustersMetricCapacity `json:"capacity"`
	Clustertype                     string                   `json:"clustertype"`
	Cpuallocated                    string                   `json:"cpuallocated"`
	Cpuallocateddisablethreshold    bool                     `json:"cpuallocateddisablethreshold"`
	Cpuallocatedthreshold           bool                     `json:"cpuallocatedthreshold"`
	Cpudisablethreshold             bool                     `json:"cpudisablethreshold"`
	Cpumaxdeviation                 string                   `json:"cpumaxdeviation"`
	Cpuovercommitratio              string                   `json:"cpuovercommitratio"`
	Cputhreshold                    bool                     `json:"cputhreshold"`
	Cputotal                        string                   `json:"cputotal"`
	Cpuused                         string                   `json:"cpuused"`
	Drsimbalance                    string                   `json:"drsimbalance"`
	Extensionid                     string                   `json:"extensionid"`
	Extensionname                   string                   `json:"extensionname"`
	Hasannotations                  bool                     `json:"hasannotations"`
	Hosts                           string                   `json:"hosts"`
	Hypervisortype                  string                   `json:"hypervisortype"`
	Id                              string                   `json:"id"`
	JobID                           string                   `json:"jobid"`
	Jobstatus                       int                      `json:"jobstatus"`
	Managedstate                    string                   `json:"managedstate"`
	Memoryallocated                 string                   `json:"memoryallocated"`
	Memoryallocateddisablethreshold bool                     `json:"memoryallocateddisablethreshold"`
	Memoryallocatedthreshold        bool                     `json:"memoryallocatedthreshold"`
	Memorydisablethreshold          bool                     `json:"memorydisablethreshold"`
	Memorymaxdeviation              string                   `json:"memorymaxdeviation"`
	Memoryovercommitratio           string                   `json:"memoryovercommitratio"`
	Memorythreshold                 bool                     `json:"memorythreshold"`
	Memorytotal                     string                   `json:"memorytotal"`
	Memoryused                      string                   `json:"memoryused"`
	Name                            string                   `json:"name"`
	Ovm3vip                         string                   `json:"ovm3vip"`
	Podid                           string                   `json:"podid"`
	Podname                         string                   `json:"podname"`
	Podstorageaccessgroups          string                   `json:"podstorageaccessgroups"`
	Resourcedetails                 map[string]string        `json:"resourcedetails"`
	State                           string                   `json:"state"`
	Storageaccessgroups             string                   `json:"storageaccessgroups"`
	Zoneid                          string                   `json:"zoneid"`
	Zonename                        string                   `json:"zonename"`
	Zonestorageaccessgroups         string                   `json:"zonestorageaccessgroups"`
}

type ClustersMetricCapacity struct {
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

type ListDedicatedClustersParams struct {
	p map[string]interface{}
}

func (p *ListDedicatedClustersParams) toURLValues() url.Values {
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
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
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
	return u
}

func (p *ListDedicatedClustersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListDedicatedClustersParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListDedicatedClustersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListDedicatedClustersParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
}

func (p *ListDedicatedClustersParams) ResetAffinitygroupid() {
	if p.p != nil && p.p["affinitygroupid"] != nil {
		delete(p.p, "affinitygroupid")
	}
}

func (p *ListDedicatedClustersParams) GetAffinitygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupid"].(string)
	return value, ok
}

func (p *ListDedicatedClustersParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListDedicatedClustersParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListDedicatedClustersParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListDedicatedClustersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListDedicatedClustersParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListDedicatedClustersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListDedicatedClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListDedicatedClustersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListDedicatedClustersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListDedicatedClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListDedicatedClustersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListDedicatedClustersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListDedicatedClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListDedicatedClustersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListDedicatedClustersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListDedicatedClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListDedicatedClustersParams() *ListDedicatedClustersParams {
	p := &ListDedicatedClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists dedicated clusters.
func (s *ClusterService) ListDedicatedClusters(p *ListDedicatedClustersParams) (*ListDedicatedClustersResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedClustersResponse struct {
	Count             int                 `json:"count"`
	DedicatedClusters []*DedicatedCluster `json:"dedicatedcluster"`
}

type DedicatedCluster struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Clusterid       string `json:"clusterid"`
	Clustername     string `json:"clustername"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
}

type ReleaseDedicatedClusterParams struct {
	p map[string]interface{}
}

func (p *ReleaseDedicatedClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (p *ReleaseDedicatedClusterParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ReleaseDedicatedClusterParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ReleaseDedicatedClusterParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewReleaseDedicatedClusterParams(clusterid string) *ReleaseDedicatedClusterParams {
	p := &ReleaseDedicatedClusterParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Release the dedication for cluster
func (s *ClusterService) ReleaseDedicatedCluster(p *ReleaseDedicatedClusterParams) (*ReleaseDedicatedClusterResponse, error) {
	resp, err := s.cs.newPostRequest("releaseDedicatedCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedClusterResponse
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

type ReleaseDedicatedClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateClusterParams struct {
	p map[string]interface{}
}

func (p *UpdateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["externaldetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("externaldetails[%d].key", i), k)
			u.Set(fmt.Sprintf("externaldetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["managedstate"]; found {
		u.Set("managedstate", v.(string))
	}
	return u
}

func (p *UpdateClusterParams) SetAllocationstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allocationstate"] = v
}

func (p *UpdateClusterParams) ResetAllocationstate() {
	if p.p != nil && p.p["allocationstate"] != nil {
		delete(p.p, "allocationstate")
	}
}

func (p *UpdateClusterParams) GetAllocationstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allocationstate"].(string)
	return value, ok
}

func (p *UpdateClusterParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *UpdateClusterParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *UpdateClusterParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *UpdateClusterParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
}

func (p *UpdateClusterParams) ResetClustername() {
	if p.p != nil && p.p["clustername"] != nil {
		delete(p.p, "clustername")
	}
}

func (p *UpdateClusterParams) GetClustername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustername"].(string)
	return value, ok
}

func (p *UpdateClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *UpdateClusterParams) ResetClustertype() {
	if p.p != nil && p.p["clustertype"] != nil {
		delete(p.p, "clustertype")
	}
}

func (p *UpdateClusterParams) GetClustertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustertype"].(string)
	return value, ok
}

func (p *UpdateClusterParams) SetExternaldetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["externaldetails"] = v
}

func (p *UpdateClusterParams) ResetExternaldetails() {
	if p.p != nil && p.p["externaldetails"] != nil {
		delete(p.p, "externaldetails")
	}
}

func (p *UpdateClusterParams) GetExternaldetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["externaldetails"].(map[string]string)
	return value, ok
}

func (p *UpdateClusterParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *UpdateClusterParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *UpdateClusterParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *UpdateClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateClusterParams) SetManagedstate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managedstate"] = v
}

func (p *UpdateClusterParams) ResetManagedstate() {
	if p.p != nil && p.p["managedstate"] != nil {
		delete(p.p, "managedstate")
	}
}

func (p *UpdateClusterParams) GetManagedstate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managedstate"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewUpdateClusterParams(id string) *UpdateClusterParams {
	p := &UpdateClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing cluster
func (s *ClusterService) UpdateCluster(p *UpdateClusterParams) (*UpdateClusterResponse, error) {
	resp, err := s.cs.newPostRequest("updateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateClusterResponse struct {
	Allocationstate         string                          `json:"allocationstate"`
	Arch                    string                          `json:"arch"`
	Capacity                []UpdateClusterResponseCapacity `json:"capacity"`
	Clustertype             string                          `json:"clustertype"`
	Cpuovercommitratio      string                          `json:"cpuovercommitratio"`
	Extensionid             string                          `json:"extensionid"`
	Extensionname           string                          `json:"extensionname"`
	Hasannotations          bool                            `json:"hasannotations"`
	Hypervisortype          string                          `json:"hypervisortype"`
	Id                      string                          `json:"id"`
	JobID                   string                          `json:"jobid"`
	Jobstatus               int                             `json:"jobstatus"`
	Managedstate            string                          `json:"managedstate"`
	Memoryovercommitratio   string                          `json:"memoryovercommitratio"`
	Name                    string                          `json:"name"`
	Ovm3vip                 string                          `json:"ovm3vip"`
	Podid                   string                          `json:"podid"`
	Podname                 string                          `json:"podname"`
	Podstorageaccessgroups  string                          `json:"podstorageaccessgroups"`
	Resourcedetails         map[string]string               `json:"resourcedetails"`
	Storageaccessgroups     string                          `json:"storageaccessgroups"`
	Zoneid                  string                          `json:"zoneid"`
	Zonename                string                          `json:"zonename"`
	Zonestorageaccessgroups string                          `json:"zonestorageaccessgroups"`
}

type UpdateClusterResponseCapacity struct {
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
