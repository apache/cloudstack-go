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

type KubernetesServiceIface interface {
	AddKubernetesSupportedVersion(p *AddKubernetesSupportedVersionParams) (*AddKubernetesSupportedVersionResponse, error)
	NewAddKubernetesSupportedVersionParams(mincpunumber int, minmemory int, semanticversion string) *AddKubernetesSupportedVersionParams
	CreateKubernetesCluster(p *CreateKubernetesClusterParams) (*CreateKubernetesClusterResponse, error)
	NewCreateKubernetesClusterParams(description string, kubernetesversionid string, name string, serviceofferingid string, size int64, zoneid string) *CreateKubernetesClusterParams
	DeleteKubernetesCluster(p *DeleteKubernetesClusterParams) (*DeleteKubernetesClusterResponse, error)
	NewDeleteKubernetesClusterParams(id string) *DeleteKubernetesClusterParams
	DeleteKubernetesSupportedVersion(p *DeleteKubernetesSupportedVersionParams) (*DeleteKubernetesSupportedVersionResponse, error)
	NewDeleteKubernetesSupportedVersionParams(id string) *DeleteKubernetesSupportedVersionParams
	GetKubernetesClusterConfig(p *GetKubernetesClusterConfigParams) (*GetKubernetesClusterConfigResponse, error)
	NewGetKubernetesClusterConfigParams() *GetKubernetesClusterConfigParams
	ListKubernetesClusters(p *ListKubernetesClustersParams) (*ListKubernetesClustersResponse, error)
	NewListKubernetesClustersParams() *ListKubernetesClustersParams
	GetKubernetesClusterID(name string, opts ...OptionFunc) (string, int, error)
	GetKubernetesClusterByName(name string, opts ...OptionFunc) (*KubernetesCluster, int, error)
	GetKubernetesClusterByID(id string, opts ...OptionFunc) (*KubernetesCluster, int, error)
	ListKubernetesSupportedVersions(p *ListKubernetesSupportedVersionsParams) (*ListKubernetesSupportedVersionsResponse, error)
	NewListKubernetesSupportedVersionsParams() *ListKubernetesSupportedVersionsParams
	GetKubernetesSupportedVersionID(keyword string, opts ...OptionFunc) (string, int, error)
	GetKubernetesSupportedVersionByName(name string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error)
	GetKubernetesSupportedVersionByID(id string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error)
	ScaleKubernetesCluster(p *ScaleKubernetesClusterParams) (*ScaleKubernetesClusterResponse, error)
	NewScaleKubernetesClusterParams(id string) *ScaleKubernetesClusterParams
	StartKubernetesCluster(p *StartKubernetesClusterParams) (*StartKubernetesClusterResponse, error)
	NewStartKubernetesClusterParams(id string) *StartKubernetesClusterParams
	StopKubernetesCluster(p *StopKubernetesClusterParams) (*StopKubernetesClusterResponse, error)
	NewStopKubernetesClusterParams(id string) *StopKubernetesClusterParams
	UpdateKubernetesSupportedVersion(p *UpdateKubernetesSupportedVersionParams) (*UpdateKubernetesSupportedVersionResponse, error)
	NewUpdateKubernetesSupportedVersionParams(id string, state string) *UpdateKubernetesSupportedVersionParams
	UpgradeKubernetesCluster(p *UpgradeKubernetesClusterParams) (*UpgradeKubernetesClusterResponse, error)
	NewUpgradeKubernetesClusterParams(id string, kubernetesversionid string) *UpgradeKubernetesClusterParams
	AddVirtualMachinesToKubernetesCluster(p *AddVirtualMachinesToKubernetesClusterParams) (*AddVirtualMachinesToKubernetesClusterResponse, error)
	NewAddVirtualMachinesToKubernetesClusterParams(id string, virtualmachineids []string) *AddVirtualMachinesToKubernetesClusterParams
	RemoveVirtualMachinesFromKubernetesCluster(p *RemoveVirtualMachinesFromKubernetesClusterParams) (*RemoveVirtualMachinesFromKubernetesClusterResponse, error)
	NewRemoveVirtualMachinesFromKubernetesClusterParams(id string, virtualmachineids []string) *RemoveVirtualMachinesFromKubernetesClusterParams
}

type AddKubernetesSupportedVersionParams struct {
	p map[string]interface{}
}

func (p *AddKubernetesSupportedVersionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := p.p["directdownload"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("directdownload", vv)
	}
	if v, found := p.p["mincpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("mincpunumber", vv)
	}
	if v, found := p.p["minmemory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmemory", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["semanticversion"]; found {
		u.Set("semanticversion", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddKubernetesSupportedVersionParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *AddKubernetesSupportedVersionParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetChecksum() {
	if p.p != nil && p.p["checksum"] != nil {
		delete(p.p, "checksum")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetChecksum() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["checksum"].(string)
	return value, ok
}

func (p *AddKubernetesSupportedVersionParams) SetDirectdownload(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["directdownload"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetDirectdownload() {
	if p.p != nil && p.p["directdownload"] != nil {
		delete(p.p, "directdownload")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetDirectdownload() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["directdownload"].(bool)
	return value, ok
}

func (p *AddKubernetesSupportedVersionParams) SetMincpunumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mincpunumber"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetMincpunumber() {
	if p.p != nil && p.p["mincpunumber"] != nil {
		delete(p.p, "mincpunumber")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetMincpunumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mincpunumber"].(int)
	return value, ok
}

func (p *AddKubernetesSupportedVersionParams) SetMinmemory(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minmemory"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetMinmemory() {
	if p.p != nil && p.p["minmemory"] != nil {
		delete(p.p, "minmemory")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetMinmemory() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["minmemory"].(int)
	return value, ok
}

func (p *AddKubernetesSupportedVersionParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddKubernetesSupportedVersionParams) SetSemanticversion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["semanticversion"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetSemanticversion() {
	if p.p != nil && p.p["semanticversion"] != nil {
		delete(p.p, "semanticversion")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetSemanticversion() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["semanticversion"].(string)
	return value, ok
}

func (p *AddKubernetesSupportedVersionParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *AddKubernetesSupportedVersionParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddKubernetesSupportedVersionParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddKubernetesSupportedVersionParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddKubernetesSupportedVersionParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewAddKubernetesSupportedVersionParams(mincpunumber int, minmemory int, semanticversion string) *AddKubernetesSupportedVersionParams {
	p := &AddKubernetesSupportedVersionParams{}
	p.p = make(map[string]interface{})
	p.p["mincpunumber"] = mincpunumber
	p.p["minmemory"] = minmemory
	p.p["semanticversion"] = semanticversion
	return p
}

// Add a supported Kubernetes version
func (s *KubernetesService) AddKubernetesSupportedVersion(p *AddKubernetesSupportedVersionParams) (*AddKubernetesSupportedVersionResponse, error) {
	resp, err := s.cs.newRequest("addKubernetesSupportedVersion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r AddKubernetesSupportedVersionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddKubernetesSupportedVersionResponse struct {
	Created             string `json:"created"`
	Directdownload      bool   `json:"directdownload"`
	Id                  string `json:"id"`
	Isoid               string `json:"isoid"`
	Isoname             string `json:"isoname"`
	Isostate            string `json:"isostate"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Mincpunumber        int    `json:"mincpunumber"`
	Minmemory           int    `json:"minmemory"`
	Name                string `json:"name"`
	Semanticversion     string `json:"semanticversion"`
	State               string `json:"state"`
	Supportsautoscaling bool   `json:"supportsautoscaling"`
	Supportsha          bool   `json:"supportsha"`
	Zoneid              string `json:"zoneid"`
	Zonename            string `json:"zonename"`
}

type CreateKubernetesClusterParams struct {
	p map[string]interface{}
}

func (p *CreateKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["controlnodes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("controlnodes", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["dockerregistrypassword"]; found {
		u.Set("dockerregistrypassword", v.(string))
	}
	if v, found := p.p["dockerregistryurl"]; found {
		u.Set("dockerregistryurl", v.(string))
	}
	if v, found := p.p["dockerregistryusername"]; found {
		u.Set("dockerregistryusername", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["externalloadbalanceripaddress"]; found {
		u.Set("externalloadbalanceripaddress", v.(string))
	}
	if v, found := p.p["keypair"]; found {
		u.Set("keypair", v.(string))
	}
	if v, found := p.p["kubernetesversionid"]; found {
		u.Set("kubernetesversionid", v.(string))
	}
	if v, found := p.p["masternodes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("masternodes", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["noderootdisksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("noderootdisksize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateKubernetesClusterParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateKubernetesClusterParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateKubernetesClusterParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *CreateKubernetesClusterParams) ResetClustertype() {
	if p.p != nil && p.p["clustertype"] != nil {
		delete(p.p, "clustertype")
	}
}

func (p *CreateKubernetesClusterParams) GetClustertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustertype"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetControlnodes(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["controlnodes"] = v
}

func (p *CreateKubernetesClusterParams) ResetControlnodes() {
	if p.p != nil && p.p["controlnodes"] != nil {
		delete(p.p, "controlnodes")
	}
}

func (p *CreateKubernetesClusterParams) GetControlnodes() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["controlnodes"].(int64)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateKubernetesClusterParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateKubernetesClusterParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetDockerregistrypassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dockerregistrypassword"] = v
}

func (p *CreateKubernetesClusterParams) ResetDockerregistrypassword() {
	if p.p != nil && p.p["dockerregistrypassword"] != nil {
		delete(p.p, "dockerregistrypassword")
	}
}

func (p *CreateKubernetesClusterParams) GetDockerregistrypassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dockerregistrypassword"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetDockerregistryurl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dockerregistryurl"] = v
}

func (p *CreateKubernetesClusterParams) ResetDockerregistryurl() {
	if p.p != nil && p.p["dockerregistryurl"] != nil {
		delete(p.p, "dockerregistryurl")
	}
}

func (p *CreateKubernetesClusterParams) GetDockerregistryurl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dockerregistryurl"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetDockerregistryusername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dockerregistryusername"] = v
}

func (p *CreateKubernetesClusterParams) ResetDockerregistryusername() {
	if p.p != nil && p.p["dockerregistryusername"] != nil {
		delete(p.p, "dockerregistryusername")
	}
}

func (p *CreateKubernetesClusterParams) GetDockerregistryusername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dockerregistryusername"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateKubernetesClusterParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateKubernetesClusterParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetExternalloadbalanceripaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["externalloadbalanceripaddress"] = v
}

func (p *CreateKubernetesClusterParams) ResetExternalloadbalanceripaddress() {
	if p.p != nil && p.p["externalloadbalanceripaddress"] != nil {
		delete(p.p, "externalloadbalanceripaddress")
	}
}

func (p *CreateKubernetesClusterParams) GetExternalloadbalanceripaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["externalloadbalanceripaddress"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
}

func (p *CreateKubernetesClusterParams) ResetKeypair() {
	if p.p != nil && p.p["keypair"] != nil {
		delete(p.p, "keypair")
	}
}

func (p *CreateKubernetesClusterParams) GetKeypair() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypair"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetKubernetesversionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["kubernetesversionid"] = v
}

func (p *CreateKubernetesClusterParams) ResetKubernetesversionid() {
	if p.p != nil && p.p["kubernetesversionid"] != nil {
		delete(p.p, "kubernetesversionid")
	}
}

func (p *CreateKubernetesClusterParams) GetKubernetesversionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["kubernetesversionid"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetMasternodes(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["masternodes"] = v
}

func (p *CreateKubernetesClusterParams) ResetMasternodes() {
	if p.p != nil && p.p["masternodes"] != nil {
		delete(p.p, "masternodes")
	}
}

func (p *CreateKubernetesClusterParams) GetMasternodes() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["masternodes"].(int64)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateKubernetesClusterParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateKubernetesClusterParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *CreateKubernetesClusterParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *CreateKubernetesClusterParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetNoderootdisksize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["noderootdisksize"] = v
}

func (p *CreateKubernetesClusterParams) ResetNoderootdisksize() {
	if p.p != nil && p.p["noderootdisksize"] != nil {
		delete(p.p, "noderootdisksize")
	}
}

func (p *CreateKubernetesClusterParams) GetNoderootdisksize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["noderootdisksize"].(int64)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateKubernetesClusterParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateKubernetesClusterParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateKubernetesClusterParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *CreateKubernetesClusterParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *CreateKubernetesClusterParams) ResetSize() {
	if p.p != nil && p.p["size"] != nil {
		delete(p.p, "size")
	}
}

func (p *CreateKubernetesClusterParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

func (p *CreateKubernetesClusterParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateKubernetesClusterParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateKubernetesClusterParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewCreateKubernetesClusterParams(description string, kubernetesversionid string, name string, serviceofferingid string, size int64, zoneid string) *CreateKubernetesClusterParams {
	p := &CreateKubernetesClusterParams{}
	p.p = make(map[string]interface{})
	p.p["description"] = description
	p.p["kubernetesversionid"] = kubernetesversionid
	p.p["name"] = name
	p.p["serviceofferingid"] = serviceofferingid
	p.p["size"] = size
	p.p["zoneid"] = zoneid
	return p
}

// Creates a Kubernetes cluster
func (s *KubernetesService) CreateKubernetesCluster(p *CreateKubernetesClusterParams) (*CreateKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("createKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateKubernetesClusterResponse
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

type CreateKubernetesClusterResponse struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Clustertype           string            `json:"clustertype"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Created               string            `json:"created"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainpath            string            `json:"domainpath"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type DeleteKubernetesClusterParams struct {
	p map[string]interface{}
}

func (p *DeleteKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := p.p["expunge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("expunge", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteKubernetesClusterParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *DeleteKubernetesClusterParams) ResetCleanup() {
	if p.p != nil && p.p["cleanup"] != nil {
		delete(p.p, "cleanup")
	}
}

func (p *DeleteKubernetesClusterParams) GetCleanup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanup"].(bool)
	return value, ok
}

func (p *DeleteKubernetesClusterParams) SetExpunge(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expunge"] = v
}

func (p *DeleteKubernetesClusterParams) ResetExpunge() {
	if p.p != nil && p.p["expunge"] != nil {
		delete(p.p, "expunge")
	}
}

func (p *DeleteKubernetesClusterParams) GetExpunge() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["expunge"].(bool)
	return value, ok
}

func (p *DeleteKubernetesClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteKubernetesClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteKubernetesClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewDeleteKubernetesClusterParams(id string) *DeleteKubernetesClusterParams {
	p := &DeleteKubernetesClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Kubernetes cluster
func (s *KubernetesService) DeleteKubernetesCluster(p *DeleteKubernetesClusterParams) (*DeleteKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("deleteKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteKubernetesClusterResponse
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

type DeleteKubernetesClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteKubernetesSupportedVersionParams struct {
	p map[string]interface{}
}

func (p *DeleteKubernetesSupportedVersionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteKubernetesSupportedVersionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteKubernetesSupportedVersionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteKubernetesSupportedVersionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteKubernetesSupportedVersionParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewDeleteKubernetesSupportedVersionParams(id string) *DeleteKubernetesSupportedVersionParams {
	p := &DeleteKubernetesSupportedVersionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Kubernetes cluster
func (s *KubernetesService) DeleteKubernetesSupportedVersion(p *DeleteKubernetesSupportedVersionParams) (*DeleteKubernetesSupportedVersionResponse, error) {
	resp, err := s.cs.newRequest("deleteKubernetesSupportedVersion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteKubernetesSupportedVersionResponse
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

type DeleteKubernetesSupportedVersionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type GetKubernetesClusterConfigParams struct {
	p map[string]interface{}
}

func (p *GetKubernetesClusterConfigParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *GetKubernetesClusterConfigParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *GetKubernetesClusterConfigParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *GetKubernetesClusterConfigParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new GetKubernetesClusterConfigParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewGetKubernetesClusterConfigParams() *GetKubernetesClusterConfigParams {
	p := &GetKubernetesClusterConfigParams{}
	p.p = make(map[string]interface{})
	return p
}

// Get Kubernetes cluster config
func (s *KubernetesService) GetKubernetesClusterConfig(p *GetKubernetesClusterConfigParams) (*GetKubernetesClusterConfigResponse, error) {
	resp, err := s.cs.newRequest("getKubernetesClusterConfig", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response GetKubernetesClusterConfigResponse `json:"clusterconfig"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type GetKubernetesClusterConfigResponse struct {
	Configdata string `json:"configdata"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Name       string `json:"name"`
}

type ListKubernetesClustersParams struct {
	p map[string]interface{}
}

func (p *ListKubernetesClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *ListKubernetesClustersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListKubernetesClustersParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListKubernetesClustersParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetClustertype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustertype"] = v
}

func (p *ListKubernetesClustersParams) ResetClustertype() {
	if p.p != nil && p.p["clustertype"] != nil {
		delete(p.p, "clustertype")
	}
}

func (p *ListKubernetesClustersParams) GetClustertype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustertype"].(string)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListKubernetesClustersParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListKubernetesClustersParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListKubernetesClustersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListKubernetesClustersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListKubernetesClustersParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListKubernetesClustersParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListKubernetesClustersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListKubernetesClustersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListKubernetesClustersParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListKubernetesClustersParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListKubernetesClustersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListKubernetesClustersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListKubernetesClustersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListKubernetesClustersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListKubernetesClustersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListKubernetesClustersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListKubernetesClustersParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListKubernetesClustersParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListKubernetesClustersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListKubernetesClustersParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListKubernetesClustersParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

// You should always use this function to get a new ListKubernetesClustersParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewListKubernetesClustersParams() *ListKubernetesClustersParams {
	p := &ListKubernetesClustersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesClusterID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListKubernetesClustersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListKubernetesClusters(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.KubernetesClusters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.KubernetesClusters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesClusterByName(name string, opts ...OptionFunc) (*KubernetesCluster, int, error) {
	id, count, err := s.GetKubernetesClusterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetKubernetesClusterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesClusterByID(id string, opts ...OptionFunc) (*KubernetesCluster, int, error) {
	p := &ListKubernetesClustersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListKubernetesClusters(p)
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
		return l.KubernetesClusters[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for KubernetesCluster UUID: %s!", id)
}

// Lists Kubernetes clusters
func (s *KubernetesService) ListKubernetesClusters(p *ListKubernetesClustersParams) (*ListKubernetesClustersResponse, error) {
	resp, err := s.cs.newRequest("listKubernetesClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListKubernetesClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListKubernetesClustersResponse struct {
	Count              int                  `json:"count"`
	KubernetesClusters []*KubernetesCluster `json:"kubernetescluster"`
}

type KubernetesCluster struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Clustertype           string            `json:"clustertype"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Created               string            `json:"created"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainpath            string            `json:"domainpath"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type ListKubernetesSupportedVersionsParams struct {
	p map[string]interface{}
}

func (p *ListKubernetesSupportedVersionsParams) toURLValues() url.Values {
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
	if v, found := p.p["minimumkubernetesversionid"]; found {
		u.Set("minimumkubernetesversionid", v.(string))
	}
	if v, found := p.p["minimumsemanticversion"]; found {
		u.Set("minimumsemanticversion", v.(string))
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

func (p *ListKubernetesSupportedVersionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListKubernetesSupportedVersionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListKubernetesSupportedVersionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListKubernetesSupportedVersionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListKubernetesSupportedVersionsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListKubernetesSupportedVersionsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListKubernetesSupportedVersionsParams) SetMinimumkubernetesversionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minimumkubernetesversionid"] = v
}

func (p *ListKubernetesSupportedVersionsParams) ResetMinimumkubernetesversionid() {
	if p.p != nil && p.p["minimumkubernetesversionid"] != nil {
		delete(p.p, "minimumkubernetesversionid")
	}
}

func (p *ListKubernetesSupportedVersionsParams) GetMinimumkubernetesversionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["minimumkubernetesversionid"].(string)
	return value, ok
}

func (p *ListKubernetesSupportedVersionsParams) SetMinimumsemanticversion(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minimumsemanticversion"] = v
}

func (p *ListKubernetesSupportedVersionsParams) ResetMinimumsemanticversion() {
	if p.p != nil && p.p["minimumsemanticversion"] != nil {
		delete(p.p, "minimumsemanticversion")
	}
}

func (p *ListKubernetesSupportedVersionsParams) GetMinimumsemanticversion() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["minimumsemanticversion"].(string)
	return value, ok
}

func (p *ListKubernetesSupportedVersionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListKubernetesSupportedVersionsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListKubernetesSupportedVersionsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListKubernetesSupportedVersionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListKubernetesSupportedVersionsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListKubernetesSupportedVersionsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListKubernetesSupportedVersionsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListKubernetesSupportedVersionsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListKubernetesSupportedVersionsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListKubernetesSupportedVersionsParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewListKubernetesSupportedVersionsParams() *ListKubernetesSupportedVersionsParams {
	p := &ListKubernetesSupportedVersionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesSupportedVersionID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListKubernetesSupportedVersionsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListKubernetesSupportedVersions(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.KubernetesSupportedVersions[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.KubernetesSupportedVersions {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesSupportedVersionByName(name string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error) {
	id, count, err := s.GetKubernetesSupportedVersionID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetKubernetesSupportedVersionByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesSupportedVersionByID(id string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error) {
	p := &ListKubernetesSupportedVersionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListKubernetesSupportedVersions(p)
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
		return l.KubernetesSupportedVersions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for KubernetesSupportedVersion UUID: %s!", id)
}

// Lists supported Kubernetes version
func (s *KubernetesService) ListKubernetesSupportedVersions(p *ListKubernetesSupportedVersionsParams) (*ListKubernetesSupportedVersionsResponse, error) {
	resp, err := s.cs.newRequest("listKubernetesSupportedVersions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListKubernetesSupportedVersionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListKubernetesSupportedVersionsResponse struct {
	Count                       int                           `json:"count"`
	KubernetesSupportedVersions []*KubernetesSupportedVersion `json:"kubernetessupportedversion"`
}

type KubernetesSupportedVersion struct {
	Created             string `json:"created"`
	Directdownload      bool   `json:"directdownload"`
	Id                  string `json:"id"`
	Isoid               string `json:"isoid"`
	Isoname             string `json:"isoname"`
	Isostate            string `json:"isostate"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Mincpunumber        int    `json:"mincpunumber"`
	Minmemory           int    `json:"minmemory"`
	Name                string `json:"name"`
	Semanticversion     string `json:"semanticversion"`
	State               string `json:"state"`
	Supportsautoscaling bool   `json:"supportsautoscaling"`
	Supportsha          bool   `json:"supportsha"`
	Zoneid              string `json:"zoneid"`
	Zonename            string `json:"zonename"`
}

type ScaleKubernetesClusterParams struct {
	p map[string]interface{}
}

func (p *ScaleKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["autoscalingenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("autoscalingenabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["maxsize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxsize", vv)
	}
	if v, found := p.p["minsize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("minsize", vv)
	}
	if v, found := p.p["nodeids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("nodeids", vv)
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	return u
}

func (p *ScaleKubernetesClusterParams) SetAutoscalingenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoscalingenabled"] = v
}

func (p *ScaleKubernetesClusterParams) ResetAutoscalingenabled() {
	if p.p != nil && p.p["autoscalingenabled"] != nil {
		delete(p.p, "autoscalingenabled")
	}
}

func (p *ScaleKubernetesClusterParams) GetAutoscalingenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoscalingenabled"].(bool)
	return value, ok
}

func (p *ScaleKubernetesClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ScaleKubernetesClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ScaleKubernetesClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ScaleKubernetesClusterParams) SetMaxsize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxsize"] = v
}

func (p *ScaleKubernetesClusterParams) ResetMaxsize() {
	if p.p != nil && p.p["maxsize"] != nil {
		delete(p.p, "maxsize")
	}
}

func (p *ScaleKubernetesClusterParams) GetMaxsize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxsize"].(int64)
	return value, ok
}

func (p *ScaleKubernetesClusterParams) SetMinsize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minsize"] = v
}

func (p *ScaleKubernetesClusterParams) ResetMinsize() {
	if p.p != nil && p.p["minsize"] != nil {
		delete(p.p, "minsize")
	}
}

func (p *ScaleKubernetesClusterParams) GetMinsize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["minsize"].(int64)
	return value, ok
}

func (p *ScaleKubernetesClusterParams) SetNodeids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nodeids"] = v
}

func (p *ScaleKubernetesClusterParams) ResetNodeids() {
	if p.p != nil && p.p["nodeids"] != nil {
		delete(p.p, "nodeids")
	}
}

func (p *ScaleKubernetesClusterParams) GetNodeids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nodeids"].([]string)
	return value, ok
}

func (p *ScaleKubernetesClusterParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ScaleKubernetesClusterParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ScaleKubernetesClusterParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ScaleKubernetesClusterParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *ScaleKubernetesClusterParams) ResetSize() {
	if p.p != nil && p.p["size"] != nil {
		delete(p.p, "size")
	}
}

func (p *ScaleKubernetesClusterParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

// You should always use this function to get a new ScaleKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewScaleKubernetesClusterParams(id string) *ScaleKubernetesClusterParams {
	p := &ScaleKubernetesClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Scales a created, running or stopped CloudManaged Kubernetes cluster
func (s *KubernetesService) ScaleKubernetesCluster(p *ScaleKubernetesClusterParams) (*ScaleKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("scaleKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleKubernetesClusterResponse
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

type ScaleKubernetesClusterResponse struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Clustertype           string            `json:"clustertype"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Created               string            `json:"created"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainpath            string            `json:"domainpath"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type StartKubernetesClusterParams struct {
	p map[string]interface{}
}

func (p *StartKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartKubernetesClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StartKubernetesClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StartKubernetesClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StartKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewStartKubernetesClusterParams(id string) *StartKubernetesClusterParams {
	p := &StartKubernetesClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a stopped CloudManaged Kubernetes cluster
func (s *KubernetesService) StartKubernetesCluster(p *StartKubernetesClusterParams) (*StartKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("startKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartKubernetesClusterResponse
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

type StartKubernetesClusterResponse struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Clustertype           string            `json:"clustertype"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Created               string            `json:"created"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainpath            string            `json:"domainpath"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type StopKubernetesClusterParams struct {
	p map[string]interface{}
}

func (p *StopKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StopKubernetesClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StopKubernetesClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StopKubernetesClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewStopKubernetesClusterParams(id string) *StopKubernetesClusterParams {
	p := &StopKubernetesClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a running CloudManaged Kubernetes cluster
func (s *KubernetesService) StopKubernetesCluster(p *StopKubernetesClusterParams) (*StopKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("stopKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopKubernetesClusterResponse
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

type StopKubernetesClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateKubernetesSupportedVersionParams struct {
	p map[string]interface{}
}

func (p *UpdateKubernetesSupportedVersionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *UpdateKubernetesSupportedVersionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateKubernetesSupportedVersionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateKubernetesSupportedVersionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateKubernetesSupportedVersionParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdateKubernetesSupportedVersionParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *UpdateKubernetesSupportedVersionParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateKubernetesSupportedVersionParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewUpdateKubernetesSupportedVersionParams(id string, state string) *UpdateKubernetesSupportedVersionParams {
	p := &UpdateKubernetesSupportedVersionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["state"] = state
	return p
}

// Update a supported Kubernetes version
func (s *KubernetesService) UpdateKubernetesSupportedVersion(p *UpdateKubernetesSupportedVersionParams) (*UpdateKubernetesSupportedVersionResponse, error) {
	resp, err := s.cs.newRequest("updateKubernetesSupportedVersion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateKubernetesSupportedVersionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateKubernetesSupportedVersionResponse struct {
	Created             string `json:"created"`
	Directdownload      bool   `json:"directdownload"`
	Id                  string `json:"id"`
	Isoid               string `json:"isoid"`
	Isoname             string `json:"isoname"`
	Isostate            string `json:"isostate"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Mincpunumber        int    `json:"mincpunumber"`
	Minmemory           int    `json:"minmemory"`
	Name                string `json:"name"`
	Semanticversion     string `json:"semanticversion"`
	State               string `json:"state"`
	Supportsautoscaling bool   `json:"supportsautoscaling"`
	Supportsha          bool   `json:"supportsha"`
	Zoneid              string `json:"zoneid"`
	Zonename            string `json:"zonename"`
}

type UpgradeKubernetesClusterParams struct {
	p map[string]interface{}
}

func (p *UpgradeKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["kubernetesversionid"]; found {
		u.Set("kubernetesversionid", v.(string))
	}
	return u
}

func (p *UpgradeKubernetesClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpgradeKubernetesClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpgradeKubernetesClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpgradeKubernetesClusterParams) SetKubernetesversionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["kubernetesversionid"] = v
}

func (p *UpgradeKubernetesClusterParams) ResetKubernetesversionid() {
	if p.p != nil && p.p["kubernetesversionid"] != nil {
		delete(p.p, "kubernetesversionid")
	}
}

func (p *UpgradeKubernetesClusterParams) GetKubernetesversionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["kubernetesversionid"].(string)
	return value, ok
}

// You should always use this function to get a new UpgradeKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewUpgradeKubernetesClusterParams(id string, kubernetesversionid string) *UpgradeKubernetesClusterParams {
	p := &UpgradeKubernetesClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["kubernetesversionid"] = kubernetesversionid
	return p
}

// Upgrades a running CloudManaged Kubernetes cluster
func (s *KubernetesService) UpgradeKubernetesCluster(p *UpgradeKubernetesClusterParams) (*UpgradeKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("upgradeKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpgradeKubernetesClusterResponse
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

type UpgradeKubernetesClusterResponse struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Clustertype           string            `json:"clustertype"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Created               string            `json:"created"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainpath            string            `json:"domainpath"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type AddVirtualMachinesToKubernetesClusterParams struct {
	p map[string]interface{}
}

func (p *AddVirtualMachinesToKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["iscontrolnode"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("iscontrolnode", vv)
	}
	if v, found := p.p["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("virtualmachineids", vv)
	}
	return u
}

func (p *AddVirtualMachinesToKubernetesClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *AddVirtualMachinesToKubernetesClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *AddVirtualMachinesToKubernetesClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *AddVirtualMachinesToKubernetesClusterParams) SetIscontrolnode(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iscontrolnode"] = v
}

func (p *AddVirtualMachinesToKubernetesClusterParams) ResetIscontrolnode() {
	if p.p != nil && p.p["iscontrolnode"] != nil {
		delete(p.p, "iscontrolnode")
	}
}

func (p *AddVirtualMachinesToKubernetesClusterParams) GetIscontrolnode() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iscontrolnode"].(bool)
	return value, ok
}

func (p *AddVirtualMachinesToKubernetesClusterParams) SetVirtualmachineids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineids"] = v
}

func (p *AddVirtualMachinesToKubernetesClusterParams) ResetVirtualmachineids() {
	if p.p != nil && p.p["virtualmachineids"] != nil {
		delete(p.p, "virtualmachineids")
	}
}

func (p *AddVirtualMachinesToKubernetesClusterParams) GetVirtualmachineids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineids"].([]string)
	return value, ok
}

// You should always use this function to get a new AddVirtualMachinesToKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewAddVirtualMachinesToKubernetesClusterParams(id string, virtualmachineids []string) *AddVirtualMachinesToKubernetesClusterParams {
	p := &AddVirtualMachinesToKubernetesClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineids"] = virtualmachineids
	return p
}

// Add VMs to an ExternalManaged kubernetes cluster. Not applicable for CloudManaged kubernetes clusters.
func (s *KubernetesService) AddVirtualMachinesToKubernetesCluster(p *AddVirtualMachinesToKubernetesClusterParams) (*AddVirtualMachinesToKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("addVirtualMachinesToKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddVirtualMachinesToKubernetesClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddVirtualMachinesToKubernetesClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *AddVirtualMachinesToKubernetesClusterResponse) UnmarshalJSON(b []byte) error {
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

	type alias AddVirtualMachinesToKubernetesClusterResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RemoveVirtualMachinesFromKubernetesClusterParams struct {
	p map[string]interface{}
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) toURLValues() url.Values {
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
	if v, found := p.p["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("virtualmachineids", vv)
	}
	return u
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) SetVirtualmachineids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineids"] = v
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) ResetVirtualmachineids() {
	if p.p != nil && p.p["virtualmachineids"] != nil {
		delete(p.p, "virtualmachineids")
	}
}

func (p *RemoveVirtualMachinesFromKubernetesClusterParams) GetVirtualmachineids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineids"].([]string)
	return value, ok
}

// You should always use this function to get a new RemoveVirtualMachinesFromKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewRemoveVirtualMachinesFromKubernetesClusterParams(id string, virtualmachineids []string) *RemoveVirtualMachinesFromKubernetesClusterParams {
	p := &RemoveVirtualMachinesFromKubernetesClusterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineids"] = virtualmachineids
	return p
}

// Remove VMs from an ExternalManaged kubernetes cluster. Not applicable for CloudManaged kubernetes clusters.
func (s *KubernetesService) RemoveVirtualMachinesFromKubernetesCluster(p *RemoveVirtualMachinesFromKubernetesClusterParams) (*RemoveVirtualMachinesFromKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("removeVirtualMachinesFromKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveVirtualMachinesFromKubernetesClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveVirtualMachinesFromKubernetesClusterResponse struct {
	Displaytext string `json:"displaytext"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RemoveVirtualMachinesFromKubernetesClusterResponse) UnmarshalJSON(b []byte) error {
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

	type alias RemoveVirtualMachinesFromKubernetesClusterResponse
	return json.Unmarshal(b, (*alias)(r))
}
