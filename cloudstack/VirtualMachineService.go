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

type VirtualMachineServiceIface interface {
	AddNicToVirtualMachine(p *AddNicToVirtualMachineParams) (*AddNicToVirtualMachineResponse, error)
	NewAddNicToVirtualMachineParams(networkid string, virtualmachineid string) *AddNicToVirtualMachineParams
	AssignVirtualMachine(p *AssignVirtualMachineParams) (*AssignVirtualMachineResponse, error)
	NewAssignVirtualMachineParams(virtualmachineid string) *AssignVirtualMachineParams
	ChangeServiceForVirtualMachine(p *ChangeServiceForVirtualMachineParams) (*ChangeServiceForVirtualMachineResponse, error)
	NewChangeServiceForVirtualMachineParams(id string, serviceofferingid string) *ChangeServiceForVirtualMachineParams
	CleanVMReservations(p *CleanVMReservationsParams) (*CleanVMReservationsResponse, error)
	NewCleanVMReservationsParams() *CleanVMReservationsParams
	DeployVirtualMachine(p *DeployVirtualMachineParams) (*DeployVirtualMachineResponse, error)
	NewDeployVirtualMachineParams(serviceofferingid string, templateid string, zoneid string) *DeployVirtualMachineParams
	DestroyVirtualMachine(p *DestroyVirtualMachineParams) (*DestroyVirtualMachineResponse, error)
	NewDestroyVirtualMachineParams(id string) *DestroyVirtualMachineParams
	ExpungeVirtualMachine(p *ExpungeVirtualMachineParams) (*ExpungeVirtualMachineResponse, error)
	NewExpungeVirtualMachineParams(id string) *ExpungeVirtualMachineParams
	GetVMPassword(p *GetVMPasswordParams) (*GetVMPasswordResponse, error)
	NewGetVMPasswordParams(id string) *GetVMPasswordParams
	ListVirtualMachines(p *ListVirtualMachinesParams) (*ListVirtualMachinesResponse, error)
	NewListVirtualMachinesParams() *ListVirtualMachinesParams
	GetVirtualMachineID(name string, opts ...OptionFunc) (string, int, error)
	GetVirtualMachineByName(name string, opts ...OptionFunc) (*VirtualMachine, int, error)
	GetVirtualMachineByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error)
	ListVirtualMachinesMetrics(p *ListVirtualMachinesMetricsParams) (*ListVirtualMachinesMetricsResponse, error)
	NewListVirtualMachinesMetricsParams() *ListVirtualMachinesMetricsParams
	GetVirtualMachinesMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetVirtualMachinesMetricByName(name string, opts ...OptionFunc) (*VirtualMachinesMetric, int, error)
	GetVirtualMachinesMetricByID(id string, opts ...OptionFunc) (*VirtualMachinesMetric, int, error)
	ListVmsForImport(p *ListVmsForImportParams) (*ListVmsForImportResponse, error)
	NewListVmsForImportParams(host string, hypervisor string, zoneid string) *ListVmsForImportParams
	MigrateVirtualMachine(p *MigrateVirtualMachineParams) (*MigrateVirtualMachineResponse, error)
	NewMigrateVirtualMachineParams(virtualmachineid string) *MigrateVirtualMachineParams
	MigrateVirtualMachineWithVolume(p *MigrateVirtualMachineWithVolumeParams) (*MigrateVirtualMachineWithVolumeResponse, error)
	NewMigrateVirtualMachineWithVolumeParams(virtualmachineid string) *MigrateVirtualMachineWithVolumeParams
	RebootVirtualMachine(p *RebootVirtualMachineParams) (*RebootVirtualMachineResponse, error)
	NewRebootVirtualMachineParams(id string) *RebootVirtualMachineParams
	RecoverVirtualMachine(p *RecoverVirtualMachineParams) (*RecoverVirtualMachineResponse, error)
	NewRecoverVirtualMachineParams(id string) *RecoverVirtualMachineParams
	RemoveNicFromVirtualMachine(p *RemoveNicFromVirtualMachineParams) (*RemoveNicFromVirtualMachineResponse, error)
	NewRemoveNicFromVirtualMachineParams(nicid string, virtualmachineid string) *RemoveNicFromVirtualMachineParams
	ResetPasswordForVirtualMachine(p *ResetPasswordForVirtualMachineParams) (*ResetPasswordForVirtualMachineResponse, error)
	NewResetPasswordForVirtualMachineParams(id string) *ResetPasswordForVirtualMachineParams
	ResetUserDataForVirtualMachine(p *ResetUserDataForVirtualMachineParams) (*ResetUserDataForVirtualMachineResponse, error)
	NewResetUserDataForVirtualMachineParams(id string) *ResetUserDataForVirtualMachineParams
	RestoreVirtualMachine(p *RestoreVirtualMachineParams) (*RestoreVirtualMachineResponse, error)
	NewRestoreVirtualMachineParams(virtualmachineid string) *RestoreVirtualMachineParams
	ScaleVirtualMachine(p *ScaleVirtualMachineParams) (*ScaleVirtualMachineResponse, error)
	NewScaleVirtualMachineParams(id string, serviceofferingid string) *ScaleVirtualMachineParams
	StartVirtualMachine(p *StartVirtualMachineParams) (*StartVirtualMachineResponse, error)
	NewStartVirtualMachineParams(id string) *StartVirtualMachineParams
	StopVirtualMachine(p *StopVirtualMachineParams) (*StopVirtualMachineResponse, error)
	NewStopVirtualMachineParams(id string) *StopVirtualMachineParams
	UpdateDefaultNicForVirtualMachine(p *UpdateDefaultNicForVirtualMachineParams) (*UpdateDefaultNicForVirtualMachineResponse, error)
	NewUpdateDefaultNicForVirtualMachineParams(nicid string, virtualmachineid string) *UpdateDefaultNicForVirtualMachineParams
	UpdateVirtualMachine(p *UpdateVirtualMachineParams) (*UpdateVirtualMachineResponse, error)
	NewUpdateVirtualMachineParams(id string) *UpdateVirtualMachineParams
	ListVirtualMachinesUsageHistory(p *ListVirtualMachinesUsageHistoryParams) (*ListVirtualMachinesUsageHistoryResponse, error)
	NewListVirtualMachinesUsageHistoryParams() *ListVirtualMachinesUsageHistoryParams
	GetVirtualMachinesUsageHistoryID(name string, opts ...OptionFunc) (string, int, error)
	GetVirtualMachinesUsageHistoryByName(name string, opts ...OptionFunc) (*VirtualMachinesUsageHistory, int, error)
	GetVirtualMachinesUsageHistoryByID(id string, opts ...OptionFunc) (*VirtualMachinesUsageHistory, int, error)
	ImportVm(p *ImportVmParams) (*ImportVmResponse, error)
	NewImportVmParams(clusterid string, hypervisor string, importsource string, name string, serviceofferingid string, zoneid string) *ImportVmParams
	UnmanageVirtualMachine(p *UnmanageVirtualMachineParams) (*UnmanageVirtualMachineResponse, error)
	NewUnmanageVirtualMachineParams(id string) *UnmanageVirtualMachineParams
	ListUnmanagedInstances(p *ListUnmanagedInstancesParams) (*ListUnmanagedInstancesResponse, error)
	NewListUnmanagedInstancesParams(clusterid string) *ListUnmanagedInstancesParams
	ImportUnmanagedInstance(p *ImportUnmanagedInstanceParams) (*ImportUnmanagedInstanceResponse, error)
	NewImportUnmanagedInstanceParams(clusterid string, name string, serviceofferingid string) *ImportUnmanagedInstanceParams
	CreateVMSchedule(p *CreateVMScheduleParams) (*CreateVMScheduleResponse, error)
	NewCreateVMScheduleParams(action string, schedule string, timezone string, virtualmachineid string) *CreateVMScheduleParams
	UpdateVMSchedule(p *UpdateVMScheduleParams) (*UpdateVMScheduleResponse, error)
	NewUpdateVMScheduleParams(id string) *UpdateVMScheduleParams
	ListVMSchedule(p *ListVMScheduleParams) (*ListVMScheduleResponse, error)
	NewListVMScheduleParams(virtualmachineid string) *ListVMScheduleParams
	GetVMScheduleByID(id string, virtualmachineid string, opts ...OptionFunc) (*VMSchedule, int, error)
	DeleteVMSchedule(p *DeleteVMScheduleParams) (*DeleteVMScheduleResponse, error)
	NewDeleteVMScheduleParams(virtualmachineid string) *DeleteVMScheduleParams
}

type AddNicToVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *AddNicToVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["dhcpoptions"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("dhcpoptions[%d].key", i), k)
			u.Set(fmt.Sprintf("dhcpoptions[%d].value", i), m[k])
		}
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["macaddress"]; found {
		u.Set("macaddress", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AddNicToVirtualMachineParams) SetDhcpoptions(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpoptions"] = v
}

func (p *AddNicToVirtualMachineParams) ResetDhcpoptions() {
	if p.p != nil && p.p["dhcpoptions"] != nil {
		delete(p.p, "dhcpoptions")
	}
}

func (p *AddNicToVirtualMachineParams) GetDhcpoptions() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dhcpoptions"].(map[string]string)
	return value, ok
}

func (p *AddNicToVirtualMachineParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *AddNicToVirtualMachineParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *AddNicToVirtualMachineParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *AddNicToVirtualMachineParams) SetMacaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["macaddress"] = v
}

func (p *AddNicToVirtualMachineParams) ResetMacaddress() {
	if p.p != nil && p.p["macaddress"] != nil {
		delete(p.p, "macaddress")
	}
}

func (p *AddNicToVirtualMachineParams) GetMacaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["macaddress"].(string)
	return value, ok
}

func (p *AddNicToVirtualMachineParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *AddNicToVirtualMachineParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *AddNicToVirtualMachineParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *AddNicToVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *AddNicToVirtualMachineParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *AddNicToVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new AddNicToVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewAddNicToVirtualMachineParams(networkid string, virtualmachineid string) *AddNicToVirtualMachineParams {
	p := &AddNicToVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["networkid"] = networkid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Adds VM to specified network by creating a NIC
func (s *VirtualMachineService) AddNicToVirtualMachine(p *AddNicToVirtualMachineParams) (*AddNicToVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("addNicToVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNicToVirtualMachineResponse
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

type AddNicToVirtualMachineResponse struct {
	Account               string                                        `json:"account"`
	Affinitygroup         []AddNicToVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                        `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                        `json:"autoscalevmgroupname"`
	Backupofferingid      string                                        `json:"backupofferingid"`
	Backupofferingname    string                                        `json:"backupofferingname"`
	Bootmode              string                                        `json:"bootmode"`
	Boottype              string                                        `json:"boottype"`
	Cpunumber             int                                           `json:"cpunumber"`
	Cpuspeed              int                                           `json:"cpuspeed"`
	Cpuused               string                                        `json:"cpuused"`
	Created               string                                        `json:"created"`
	Deleteprotection      bool                                          `json:"deleteprotection"`
	Details               map[string]string                             `json:"details"`
	Diskioread            int64                                         `json:"diskioread"`
	Diskiowrite           int64                                         `json:"diskiowrite"`
	Diskkbsread           int64                                         `json:"diskkbsread"`
	Diskkbswrite          int64                                         `json:"diskkbswrite"`
	Diskofferingid        string                                        `json:"diskofferingid"`
	Diskofferingname      string                                        `json:"diskofferingname"`
	Displayname           string                                        `json:"displayname"`
	Displayvm             bool                                          `json:"displayvm"`
	Domain                string                                        `json:"domain"`
	Domainid              string                                        `json:"domainid"`
	Domainpath            string                                        `json:"domainpath"`
	Forvirtualnetwork     bool                                          `json:"forvirtualnetwork"`
	Group                 string                                        `json:"group"`
	Groupid               string                                        `json:"groupid"`
	Guestosid             string                                        `json:"guestosid"`
	Haenable              bool                                          `json:"haenable"`
	Hasannotations        bool                                          `json:"hasannotations"`
	Hostcontrolstate      string                                        `json:"hostcontrolstate"`
	Hostid                string                                        `json:"hostid"`
	Hostname              string                                        `json:"hostname"`
	Hypervisor            string                                        `json:"hypervisor"`
	Icon                  interface{}                                   `json:"icon"`
	Id                    string                                        `json:"id"`
	Instancename          string                                        `json:"instancename"`
	Ipaddress             string                                        `json:"ipaddress"`
	Isdynamicallyscalable bool                                          `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                        `json:"isodisplaytext"`
	Isoid                 string                                        `json:"isoid"`
	Isoname               string                                        `json:"isoname"`
	JobID                 string                                        `json:"jobid"`
	Jobstatus             int                                           `json:"jobstatus"`
	Keypairs              string                                        `json:"keypairs"`
	Lastupdated           string                                        `json:"lastupdated"`
	Memory                int                                           `json:"memory"`
	Memoryintfreekbs      int64                                         `json:"memoryintfreekbs"`
	Memorykbs             int64                                         `json:"memorykbs"`
	Memorytargetkbs       int64                                         `json:"memorytargetkbs"`
	Name                  string                                        `json:"name"`
	Networkkbsread        int64                                         `json:"networkkbsread"`
	Networkkbswrite       int64                                         `json:"networkkbswrite"`
	Nic                   []Nic                                         `json:"nic"`
	Osdisplayname         string                                        `json:"osdisplayname"`
	Ostypeid              string                                        `json:"ostypeid"`
	Password              string                                        `json:"password"`
	Passwordenabled       bool                                          `json:"passwordenabled"`
	Pooltype              string                                        `json:"pooltype"`
	Project               string                                        `json:"project"`
	Projectid             string                                        `json:"projectid"`
	Publicip              string                                        `json:"publicip"`
	Publicipid            string                                        `json:"publicipid"`
	Readonlydetails       string                                        `json:"readonlydetails"`
	Receivedbytes         int64                                         `json:"receivedbytes"`
	Rootdeviceid          int64                                         `json:"rootdeviceid"`
	Rootdevicetype        string                                        `json:"rootdevicetype"`
	Securitygroup         []AddNicToVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                         `json:"sentbytes"`
	Serviceofferingid     string                                        `json:"serviceofferingid"`
	Serviceofferingname   string                                        `json:"serviceofferingname"`
	Servicestate          string                                        `json:"servicestate"`
	State                 string                                        `json:"state"`
	Tags                  []Tags                                        `json:"tags"`
	Templatedisplaytext   string                                        `json:"templatedisplaytext"`
	Templateformat        string                                        `json:"templateformat"`
	Templateid            string                                        `json:"templateid"`
	Templatename          string                                        `json:"templatename"`
	Templatetype          string                                        `json:"templatetype"`
	Userdata              string                                        `json:"userdata"`
	Userdatadetails       string                                        `json:"userdatadetails"`
	Userdataid            string                                        `json:"userdataid"`
	Userdataname          string                                        `json:"userdataname"`
	Userdatapolicy        string                                        `json:"userdatapolicy"`
	Userid                string                                        `json:"userid"`
	Username              string                                        `json:"username"`
	Vgpu                  string                                        `json:"vgpu"`
	Vmtype                string                                        `json:"vmtype"`
	Vnfdetails            map[string]string                             `json:"vnfdetails"`
	Vnfnics               []string                                      `json:"vnfnics"`
	Zoneid                string                                        `json:"zoneid"`
	Zonename              string                                        `json:"zonename"`
}

type AddNicToVirtualMachineResponseSecuritygroup struct {
	Account             string                                            `json:"account"`
	Description         string                                            `json:"description"`
	Domain              string                                            `json:"domain"`
	Domainid            string                                            `json:"domainid"`
	Domainpath          string                                            `json:"domainpath"`
	Egressrule          []AddNicToVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                            `json:"id"`
	Ingressrule         []AddNicToVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                            `json:"name"`
	Project             string                                            `json:"project"`
	Projectid           string                                            `json:"projectid"`
	Tags                []Tags                                            `json:"tags"`
	Virtualmachinecount int                                               `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                     `json:"virtualmachineids"`
}

type AddNicToVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type AddNicToVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *AddNicToVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias AddNicToVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type AssignVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *AssignVirtualMachineParams) toURLValues() url.Values {
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
	if v, found := p.p["networkids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("networkids", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["securitygroupids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupids", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AssignVirtualMachineParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AssignVirtualMachineParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *AssignVirtualMachineParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *AssignVirtualMachineParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *AssignVirtualMachineParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *AssignVirtualMachineParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *AssignVirtualMachineParams) SetNetworkids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
}

func (p *AssignVirtualMachineParams) ResetNetworkids() {
	if p.p != nil && p.p["networkids"] != nil {
		delete(p.p, "networkids")
	}
}

func (p *AssignVirtualMachineParams) GetNetworkids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkids"].([]string)
	return value, ok
}

func (p *AssignVirtualMachineParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AssignVirtualMachineParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *AssignVirtualMachineParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *AssignVirtualMachineParams) SetSecuritygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupids"] = v
}

func (p *AssignVirtualMachineParams) ResetSecuritygroupids() {
	if p.p != nil && p.p["securitygroupids"] != nil {
		delete(p.p, "securitygroupids")
	}
}

func (p *AssignVirtualMachineParams) GetSecuritygroupids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupids"].([]string)
	return value, ok
}

func (p *AssignVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *AssignVirtualMachineParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *AssignVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new AssignVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewAssignVirtualMachineParams(virtualmachineid string) *AssignVirtualMachineParams {
	p := &AssignVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Change ownership of a VM from one account to another. This API is available for Basic zones with security groups and Advanced zones with guest networks. A root administrator can reassign a VM from any account to any other account in any domain. A domain administrator can reassign a VM to any account in the same domain.
func (s *VirtualMachineService) AssignVirtualMachine(p *AssignVirtualMachineParams) (*AssignVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("assignVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AssignVirtualMachineResponse struct {
	Account               string                                      `json:"account"`
	Affinitygroup         []AssignVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                      `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                      `json:"autoscalevmgroupname"`
	Backupofferingid      string                                      `json:"backupofferingid"`
	Backupofferingname    string                                      `json:"backupofferingname"`
	Bootmode              string                                      `json:"bootmode"`
	Boottype              string                                      `json:"boottype"`
	Cpunumber             int                                         `json:"cpunumber"`
	Cpuspeed              int                                         `json:"cpuspeed"`
	Cpuused               string                                      `json:"cpuused"`
	Created               string                                      `json:"created"`
	Deleteprotection      bool                                        `json:"deleteprotection"`
	Details               map[string]string                           `json:"details"`
	Diskioread            int64                                       `json:"diskioread"`
	Diskiowrite           int64                                       `json:"diskiowrite"`
	Diskkbsread           int64                                       `json:"diskkbsread"`
	Diskkbswrite          int64                                       `json:"diskkbswrite"`
	Diskofferingid        string                                      `json:"diskofferingid"`
	Diskofferingname      string                                      `json:"diskofferingname"`
	Displayname           string                                      `json:"displayname"`
	Displayvm             bool                                        `json:"displayvm"`
	Domain                string                                      `json:"domain"`
	Domainid              string                                      `json:"domainid"`
	Domainpath            string                                      `json:"domainpath"`
	Forvirtualnetwork     bool                                        `json:"forvirtualnetwork"`
	Group                 string                                      `json:"group"`
	Groupid               string                                      `json:"groupid"`
	Guestosid             string                                      `json:"guestosid"`
	Haenable              bool                                        `json:"haenable"`
	Hasannotations        bool                                        `json:"hasannotations"`
	Hostcontrolstate      string                                      `json:"hostcontrolstate"`
	Hostid                string                                      `json:"hostid"`
	Hostname              string                                      `json:"hostname"`
	Hypervisor            string                                      `json:"hypervisor"`
	Icon                  interface{}                                 `json:"icon"`
	Id                    string                                      `json:"id"`
	Instancename          string                                      `json:"instancename"`
	Ipaddress             string                                      `json:"ipaddress"`
	Isdynamicallyscalable bool                                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                      `json:"isodisplaytext"`
	Isoid                 string                                      `json:"isoid"`
	Isoname               string                                      `json:"isoname"`
	JobID                 string                                      `json:"jobid"`
	Jobstatus             int                                         `json:"jobstatus"`
	Keypairs              string                                      `json:"keypairs"`
	Lastupdated           string                                      `json:"lastupdated"`
	Memory                int                                         `json:"memory"`
	Memoryintfreekbs      int64                                       `json:"memoryintfreekbs"`
	Memorykbs             int64                                       `json:"memorykbs"`
	Memorytargetkbs       int64                                       `json:"memorytargetkbs"`
	Name                  string                                      `json:"name"`
	Networkkbsread        int64                                       `json:"networkkbsread"`
	Networkkbswrite       int64                                       `json:"networkkbswrite"`
	Nic                   []Nic                                       `json:"nic"`
	Osdisplayname         string                                      `json:"osdisplayname"`
	Ostypeid              string                                      `json:"ostypeid"`
	Password              string                                      `json:"password"`
	Passwordenabled       bool                                        `json:"passwordenabled"`
	Pooltype              string                                      `json:"pooltype"`
	Project               string                                      `json:"project"`
	Projectid             string                                      `json:"projectid"`
	Publicip              string                                      `json:"publicip"`
	Publicipid            string                                      `json:"publicipid"`
	Readonlydetails       string                                      `json:"readonlydetails"`
	Receivedbytes         int64                                       `json:"receivedbytes"`
	Rootdeviceid          int64                                       `json:"rootdeviceid"`
	Rootdevicetype        string                                      `json:"rootdevicetype"`
	Securitygroup         []AssignVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                       `json:"sentbytes"`
	Serviceofferingid     string                                      `json:"serviceofferingid"`
	Serviceofferingname   string                                      `json:"serviceofferingname"`
	Servicestate          string                                      `json:"servicestate"`
	State                 string                                      `json:"state"`
	Tags                  []Tags                                      `json:"tags"`
	Templatedisplaytext   string                                      `json:"templatedisplaytext"`
	Templateformat        string                                      `json:"templateformat"`
	Templateid            string                                      `json:"templateid"`
	Templatename          string                                      `json:"templatename"`
	Templatetype          string                                      `json:"templatetype"`
	Userdata              string                                      `json:"userdata"`
	Userdatadetails       string                                      `json:"userdatadetails"`
	Userdataid            string                                      `json:"userdataid"`
	Userdataname          string                                      `json:"userdataname"`
	Userdatapolicy        string                                      `json:"userdatapolicy"`
	Userid                string                                      `json:"userid"`
	Username              string                                      `json:"username"`
	Vgpu                  string                                      `json:"vgpu"`
	Vmtype                string                                      `json:"vmtype"`
	Vnfdetails            map[string]string                           `json:"vnfdetails"`
	Vnfnics               []string                                    `json:"vnfnics"`
	Zoneid                string                                      `json:"zoneid"`
	Zonename              string                                      `json:"zonename"`
}

type AssignVirtualMachineResponseSecuritygroup struct {
	Account             string                                          `json:"account"`
	Description         string                                          `json:"description"`
	Domain              string                                          `json:"domain"`
	Domainid            string                                          `json:"domainid"`
	Domainpath          string                                          `json:"domainpath"`
	Egressrule          []AssignVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                          `json:"id"`
	Ingressrule         []AssignVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Projectid           string                                          `json:"projectid"`
	Tags                []Tags                                          `json:"tags"`
	Virtualmachinecount int                                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                   `json:"virtualmachineids"`
}

type AssignVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type AssignVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *AssignVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias AssignVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ChangeServiceForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ChangeServiceForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["automigrate"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("automigrate", vv)
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
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["shrinkok"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("shrinkok", vv)
	}
	return u
}

func (p *ChangeServiceForVirtualMachineParams) SetAutomigrate(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["automigrate"] = v
}

func (p *ChangeServiceForVirtualMachineParams) ResetAutomigrate() {
	if p.p != nil && p.p["automigrate"] != nil {
		delete(p.p, "automigrate")
	}
}

func (p *ChangeServiceForVirtualMachineParams) GetAutomigrate() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["automigrate"].(bool)
	return value, ok
}

func (p *ChangeServiceForVirtualMachineParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ChangeServiceForVirtualMachineParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ChangeServiceForVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *ChangeServiceForVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ChangeServiceForVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ChangeServiceForVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ChangeServiceForVirtualMachineParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *ChangeServiceForVirtualMachineParams) ResetMaxiops() {
	if p.p != nil && p.p["maxiops"] != nil {
		delete(p.p, "maxiops")
	}
}

func (p *ChangeServiceForVirtualMachineParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *ChangeServiceForVirtualMachineParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *ChangeServiceForVirtualMachineParams) ResetMiniops() {
	if p.p != nil && p.p["miniops"] != nil {
		delete(p.p, "miniops")
	}
}

func (p *ChangeServiceForVirtualMachineParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *ChangeServiceForVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ChangeServiceForVirtualMachineParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ChangeServiceForVirtualMachineParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ChangeServiceForVirtualMachineParams) SetShrinkok(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["shrinkok"] = v
}

func (p *ChangeServiceForVirtualMachineParams) ResetShrinkok() {
	if p.p != nil && p.p["shrinkok"] != nil {
		delete(p.p, "shrinkok")
	}
}

func (p *ChangeServiceForVirtualMachineParams) GetShrinkok() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["shrinkok"].(bool)
	return value, ok
}

// You should always use this function to get a new ChangeServiceForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewChangeServiceForVirtualMachineParams(id string, serviceofferingid string) *ChangeServiceForVirtualMachineParams {
	p := &ChangeServiceForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// (This API is deprecated, use scaleVirtualMachine API)Changes the service offering for a virtual machine. The virtual machine must be in a "Stopped" state for this command to take effect.
func (s *VirtualMachineService) ChangeServiceForVirtualMachine(p *ChangeServiceForVirtualMachineParams) (*ChangeServiceForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ChangeServiceForVirtualMachineResponse struct {
	Account               string                                                `json:"account"`
	Affinitygroup         []ChangeServiceForVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                                `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                                `json:"autoscalevmgroupname"`
	Backupofferingid      string                                                `json:"backupofferingid"`
	Backupofferingname    string                                                `json:"backupofferingname"`
	Bootmode              string                                                `json:"bootmode"`
	Boottype              string                                                `json:"boottype"`
	Cpunumber             int                                                   `json:"cpunumber"`
	Cpuspeed              int                                                   `json:"cpuspeed"`
	Cpuused               string                                                `json:"cpuused"`
	Created               string                                                `json:"created"`
	Deleteprotection      bool                                                  `json:"deleteprotection"`
	Details               map[string]string                                     `json:"details"`
	Diskioread            int64                                                 `json:"diskioread"`
	Diskiowrite           int64                                                 `json:"diskiowrite"`
	Diskkbsread           int64                                                 `json:"diskkbsread"`
	Diskkbswrite          int64                                                 `json:"diskkbswrite"`
	Diskofferingid        string                                                `json:"diskofferingid"`
	Diskofferingname      string                                                `json:"diskofferingname"`
	Displayname           string                                                `json:"displayname"`
	Displayvm             bool                                                  `json:"displayvm"`
	Domain                string                                                `json:"domain"`
	Domainid              string                                                `json:"domainid"`
	Domainpath            string                                                `json:"domainpath"`
	Forvirtualnetwork     bool                                                  `json:"forvirtualnetwork"`
	Group                 string                                                `json:"group"`
	Groupid               string                                                `json:"groupid"`
	Guestosid             string                                                `json:"guestosid"`
	Haenable              bool                                                  `json:"haenable"`
	Hasannotations        bool                                                  `json:"hasannotations"`
	Hostcontrolstate      string                                                `json:"hostcontrolstate"`
	Hostid                string                                                `json:"hostid"`
	Hostname              string                                                `json:"hostname"`
	Hypervisor            string                                                `json:"hypervisor"`
	Icon                  interface{}                                           `json:"icon"`
	Id                    string                                                `json:"id"`
	Instancename          string                                                `json:"instancename"`
	Ipaddress             string                                                `json:"ipaddress"`
	Isdynamicallyscalable bool                                                  `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                `json:"isodisplaytext"`
	Isoid                 string                                                `json:"isoid"`
	Isoname               string                                                `json:"isoname"`
	JobID                 string                                                `json:"jobid"`
	Jobstatus             int                                                   `json:"jobstatus"`
	Keypairs              string                                                `json:"keypairs"`
	Lastupdated           string                                                `json:"lastupdated"`
	Memory                int                                                   `json:"memory"`
	Memoryintfreekbs      int64                                                 `json:"memoryintfreekbs"`
	Memorykbs             int64                                                 `json:"memorykbs"`
	Memorytargetkbs       int64                                                 `json:"memorytargetkbs"`
	Name                  string                                                `json:"name"`
	Networkkbsread        int64                                                 `json:"networkkbsread"`
	Networkkbswrite       int64                                                 `json:"networkkbswrite"`
	Nic                   []Nic                                                 `json:"nic"`
	Osdisplayname         string                                                `json:"osdisplayname"`
	Ostypeid              string                                                `json:"ostypeid"`
	Password              string                                                `json:"password"`
	Passwordenabled       bool                                                  `json:"passwordenabled"`
	Pooltype              string                                                `json:"pooltype"`
	Project               string                                                `json:"project"`
	Projectid             string                                                `json:"projectid"`
	Publicip              string                                                `json:"publicip"`
	Publicipid            string                                                `json:"publicipid"`
	Readonlydetails       string                                                `json:"readonlydetails"`
	Receivedbytes         int64                                                 `json:"receivedbytes"`
	Rootdeviceid          int64                                                 `json:"rootdeviceid"`
	Rootdevicetype        string                                                `json:"rootdevicetype"`
	Securitygroup         []ChangeServiceForVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                 `json:"sentbytes"`
	Serviceofferingid     string                                                `json:"serviceofferingid"`
	Serviceofferingname   string                                                `json:"serviceofferingname"`
	Servicestate          string                                                `json:"servicestate"`
	State                 string                                                `json:"state"`
	Tags                  []Tags                                                `json:"tags"`
	Templatedisplaytext   string                                                `json:"templatedisplaytext"`
	Templateformat        string                                                `json:"templateformat"`
	Templateid            string                                                `json:"templateid"`
	Templatename          string                                                `json:"templatename"`
	Templatetype          string                                                `json:"templatetype"`
	Userdata              string                                                `json:"userdata"`
	Userdatadetails       string                                                `json:"userdatadetails"`
	Userdataid            string                                                `json:"userdataid"`
	Userdataname          string                                                `json:"userdataname"`
	Userdatapolicy        string                                                `json:"userdatapolicy"`
	Userid                string                                                `json:"userid"`
	Username              string                                                `json:"username"`
	Vgpu                  string                                                `json:"vgpu"`
	Vmtype                string                                                `json:"vmtype"`
	Vnfdetails            map[string]string                                     `json:"vnfdetails"`
	Vnfnics               []string                                              `json:"vnfnics"`
	Zoneid                string                                                `json:"zoneid"`
	Zonename              string                                                `json:"zonename"`
}

type ChangeServiceForVirtualMachineResponseSecuritygroup struct {
	Account             string                                                    `json:"account"`
	Description         string                                                    `json:"description"`
	Domain              string                                                    `json:"domain"`
	Domainid            string                                                    `json:"domainid"`
	Domainpath          string                                                    `json:"domainpath"`
	Egressrule          []ChangeServiceForVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                    `json:"id"`
	Ingressrule         []ChangeServiceForVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                    `json:"name"`
	Project             string                                                    `json:"project"`
	Projectid           string                                                    `json:"projectid"`
	Tags                []Tags                                                    `json:"tags"`
	Virtualmachinecount int                                                       `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                             `json:"virtualmachineids"`
}

type ChangeServiceForVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type ChangeServiceForVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *ChangeServiceForVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias ChangeServiceForVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CleanVMReservationsParams struct {
	p map[string]interface{}
}

func (p *CleanVMReservationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new CleanVMReservationsParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewCleanVMReservationsParams() *CleanVMReservationsParams {
	p := &CleanVMReservationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Cleanups VM reservations in the database.
func (s *VirtualMachineService) CleanVMReservations(p *CleanVMReservationsParams) (*CleanVMReservationsResponse, error) {
	resp, err := s.cs.newRequest("cleanVMReservations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CleanVMReservationsResponse
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

type CleanVMReservationsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeployVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *DeployVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["affinitygroupids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("affinitygroupids", vv)
	}
	if v, found := p.p["affinitygroupnames"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("affinitygroupnames", vv)
	}
	if v, found := p.p["bootintosetup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootintosetup", vv)
	}
	if v, found := p.p["bootmode"]; found {
		u.Set("bootmode", v.(string))
	}
	if v, found := p.p["boottype"]; found {
		u.Set("boottype", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["copyimagetags"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("copyimagetags", vv)
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["datadiskofferinglist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("datadiskofferinglist[%d].disk", i), k)
			u.Set(fmt.Sprintf("datadiskofferinglist[%d].diskOffering", i), m[k])
		}
	}
	if v, found := p.p["deploymentplanner"]; found {
		u.Set("deploymentplanner", v.(string))
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["dhcpoptionsnetworklist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("dhcpoptionsnetworklist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["displayname"]; found {
		u.Set("displayname", v.(string))
	}
	if v, found := p.p["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["dynamicscalingenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dynamicscalingenabled", vv)
	}
	if v, found := p.p["extraconfig"]; found {
		u.Set("extraconfig", v.(string))
	}
	if v, found := p.p["group"]; found {
		u.Set("group", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["iodriverpolicy"]; found {
		u.Set("iodriverpolicy", v.(string))
	}
	if v, found := p.p["iothreadsenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("iothreadsenabled", vv)
	}
	if v, found := p.p["ip6address"]; found {
		u.Set("ip6address", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["iptonetworklist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("iptonetworklist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := p.p["keyboard"]; found {
		u.Set("keyboard", v.(string))
	}
	if v, found := p.p["keypair"]; found {
		u.Set("keypair", v.(string))
	}
	if v, found := p.p["keypairs"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("keypairs", vv)
	}
	if v, found := p.p["macaddress"]; found {
		u.Set("macaddress", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("networkids", vv)
	}
	if v, found := p.p["nicmultiqueuenumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("nicmultiqueuenumber", vv)
	}
	if v, found := p.p["nicnetworklist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("nicnetworklist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := p.p["nicpackedvirtqueuesenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("nicpackedvirtqueuesenabled", vv)
	}
	if v, found := p.p["overridediskofferingid"]; found {
		u.Set("overridediskofferingid", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["properties"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("properties[%d].key", i), k)
			u.Set(fmt.Sprintf("properties[%d].value", i), m[k])
		}
	}
	if v, found := p.p["rootdisksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("rootdisksize", vv)
	}
	if v, found := p.p["securitygroupids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupids", vv)
	}
	if v, found := p.p["securitygroupnames"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupnames", vv)
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := p.p["startvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("startvm", vv)
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	if v, found := p.p["userdatadetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("userdatadetails[%d].key", i), k)
			u.Set(fmt.Sprintf("userdatadetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["userdataid"]; found {
		u.Set("userdataid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeployVirtualMachineParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DeployVirtualMachineParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DeployVirtualMachineParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetAffinitygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupids"] = v
}

func (p *DeployVirtualMachineParams) ResetAffinitygroupids() {
	if p.p != nil && p.p["affinitygroupids"] != nil {
		delete(p.p, "affinitygroupids")
	}
}

func (p *DeployVirtualMachineParams) GetAffinitygroupids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupids"].([]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetAffinitygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupnames"] = v
}

func (p *DeployVirtualMachineParams) ResetAffinitygroupnames() {
	if p.p != nil && p.p["affinitygroupnames"] != nil {
		delete(p.p, "affinitygroupnames")
	}
}

func (p *DeployVirtualMachineParams) GetAffinitygroupnames() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupnames"].([]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetBootintosetup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootintosetup"] = v
}

func (p *DeployVirtualMachineParams) ResetBootintosetup() {
	if p.p != nil && p.p["bootintosetup"] != nil {
		delete(p.p, "bootintosetup")
	}
}

func (p *DeployVirtualMachineParams) GetBootintosetup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootintosetup"].(bool)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetBootmode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootmode"] = v
}

func (p *DeployVirtualMachineParams) ResetBootmode() {
	if p.p != nil && p.p["bootmode"] != nil {
		delete(p.p, "bootmode")
	}
}

func (p *DeployVirtualMachineParams) GetBootmode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootmode"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetBoottype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["boottype"] = v
}

func (p *DeployVirtualMachineParams) ResetBoottype() {
	if p.p != nil && p.p["boottype"] != nil {
		delete(p.p, "boottype")
	}
}

func (p *DeployVirtualMachineParams) GetBoottype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["boottype"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *DeployVirtualMachineParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *DeployVirtualMachineParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetCopyimagetags(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["copyimagetags"] = v
}

func (p *DeployVirtualMachineParams) ResetCopyimagetags() {
	if p.p != nil && p.p["copyimagetags"] != nil {
		delete(p.p, "copyimagetags")
	}
}

func (p *DeployVirtualMachineParams) GetCopyimagetags() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["copyimagetags"].(bool)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *DeployVirtualMachineParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *DeployVirtualMachineParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetDatadiskofferinglist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["datadiskofferinglist"] = v
}

func (p *DeployVirtualMachineParams) ResetDatadiskofferinglist() {
	if p.p != nil && p.p["datadiskofferinglist"] != nil {
		delete(p.p, "datadiskofferinglist")
	}
}

func (p *DeployVirtualMachineParams) GetDatadiskofferinglist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["datadiskofferinglist"].(map[string]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetDeploymentplanner(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deploymentplanner"] = v
}

func (p *DeployVirtualMachineParams) ResetDeploymentplanner() {
	if p.p != nil && p.p["deploymentplanner"] != nil {
		delete(p.p, "deploymentplanner")
	}
}

func (p *DeployVirtualMachineParams) GetDeploymentplanner() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deploymentplanner"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *DeployVirtualMachineParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *DeployVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetDhcpoptionsnetworklist(v []map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpoptionsnetworklist"] = v
}

func (p *DeployVirtualMachineParams) ResetDhcpoptionsnetworklist() {
	if p.p != nil && p.p["dhcpoptionsnetworklist"] != nil {
		delete(p.p, "dhcpoptionsnetworklist")
	}
}

func (p *DeployVirtualMachineParams) GetDhcpoptionsnetworklist() ([]map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dhcpoptionsnetworklist"].([]map[string]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) AddDhcpoptionsnetworklist(item map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	val, found := p.p["dhcpoptionsnetworklist"]
	if !found {
		p.p["dhcpoptionsnetworklist"] = []map[string]string{}
		val = p.p["dhcpoptionsnetworklist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	p.p["dhcpoptionsnetworklist"] = l
}

func (p *DeployVirtualMachineParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *DeployVirtualMachineParams) ResetDiskofferingid() {
	if p.p != nil && p.p["diskofferingid"] != nil {
		delete(p.p, "diskofferingid")
	}
}

func (p *DeployVirtualMachineParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetDisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayname"] = v
}

func (p *DeployVirtualMachineParams) ResetDisplayname() {
	if p.p != nil && p.p["displayname"] != nil {
		delete(p.p, "displayname")
	}
}

func (p *DeployVirtualMachineParams) GetDisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayname"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
}

func (p *DeployVirtualMachineParams) ResetDisplayvm() {
	if p.p != nil && p.p["displayvm"] != nil {
		delete(p.p, "displayvm")
	}
}

func (p *DeployVirtualMachineParams) GetDisplayvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvm"].(bool)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DeployVirtualMachineParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DeployVirtualMachineParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetDynamicscalingenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dynamicscalingenabled"] = v
}

func (p *DeployVirtualMachineParams) ResetDynamicscalingenabled() {
	if p.p != nil && p.p["dynamicscalingenabled"] != nil {
		delete(p.p, "dynamicscalingenabled")
	}
}

func (p *DeployVirtualMachineParams) GetDynamicscalingenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dynamicscalingenabled"].(bool)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetExtraconfig(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extraconfig"] = v
}

func (p *DeployVirtualMachineParams) ResetExtraconfig() {
	if p.p != nil && p.p["extraconfig"] != nil {
		delete(p.p, "extraconfig")
	}
}

func (p *DeployVirtualMachineParams) GetExtraconfig() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extraconfig"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
}

func (p *DeployVirtualMachineParams) ResetGroup() {
	if p.p != nil && p.p["group"] != nil {
		delete(p.p, "group")
	}
}

func (p *DeployVirtualMachineParams) GetGroup() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["group"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *DeployVirtualMachineParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *DeployVirtualMachineParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *DeployVirtualMachineParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *DeployVirtualMachineParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetIodriverpolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iodriverpolicy"] = v
}

func (p *DeployVirtualMachineParams) ResetIodriverpolicy() {
	if p.p != nil && p.p["iodriverpolicy"] != nil {
		delete(p.p, "iodriverpolicy")
	}
}

func (p *DeployVirtualMachineParams) GetIodriverpolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iodriverpolicy"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetIothreadsenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iothreadsenabled"] = v
}

func (p *DeployVirtualMachineParams) ResetIothreadsenabled() {
	if p.p != nil && p.p["iothreadsenabled"] != nil {
		delete(p.p, "iothreadsenabled")
	}
}

func (p *DeployVirtualMachineParams) GetIothreadsenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iothreadsenabled"].(bool)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetIp6address(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6address"] = v
}

func (p *DeployVirtualMachineParams) ResetIp6address() {
	if p.p != nil && p.p["ip6address"] != nil {
		delete(p.p, "ip6address")
	}
}

func (p *DeployVirtualMachineParams) GetIp6address() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6address"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *DeployVirtualMachineParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *DeployVirtualMachineParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetIptonetworklist(v []map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iptonetworklist"] = v
}

func (p *DeployVirtualMachineParams) ResetIptonetworklist() {
	if p.p != nil && p.p["iptonetworklist"] != nil {
		delete(p.p, "iptonetworklist")
	}
}

func (p *DeployVirtualMachineParams) GetIptonetworklist() ([]map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iptonetworklist"].([]map[string]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) AddIptonetworklist(item map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	val, found := p.p["iptonetworklist"]
	if !found {
		p.p["iptonetworklist"] = []map[string]string{}
		val = p.p["iptonetworklist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	p.p["iptonetworklist"] = l
}

func (p *DeployVirtualMachineParams) SetKeyboard(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyboard"] = v
}

func (p *DeployVirtualMachineParams) ResetKeyboard() {
	if p.p != nil && p.p["keyboard"] != nil {
		delete(p.p, "keyboard")
	}
}

func (p *DeployVirtualMachineParams) GetKeyboard() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyboard"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
}

func (p *DeployVirtualMachineParams) ResetKeypair() {
	if p.p != nil && p.p["keypair"] != nil {
		delete(p.p, "keypair")
	}
}

func (p *DeployVirtualMachineParams) GetKeypair() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypair"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetKeypairs(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypairs"] = v
}

func (p *DeployVirtualMachineParams) ResetKeypairs() {
	if p.p != nil && p.p["keypairs"] != nil {
		delete(p.p, "keypairs")
	}
}

func (p *DeployVirtualMachineParams) GetKeypairs() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypairs"].([]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetMacaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["macaddress"] = v
}

func (p *DeployVirtualMachineParams) ResetMacaddress() {
	if p.p != nil && p.p["macaddress"] != nil {
		delete(p.p, "macaddress")
	}
}

func (p *DeployVirtualMachineParams) GetMacaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["macaddress"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *DeployVirtualMachineParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *DeployVirtualMachineParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetNetworkids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
}

func (p *DeployVirtualMachineParams) ResetNetworkids() {
	if p.p != nil && p.p["networkids"] != nil {
		delete(p.p, "networkids")
	}
}

func (p *DeployVirtualMachineParams) GetNetworkids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkids"].([]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetNicmultiqueuenumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicmultiqueuenumber"] = v
}

func (p *DeployVirtualMachineParams) ResetNicmultiqueuenumber() {
	if p.p != nil && p.p["nicmultiqueuenumber"] != nil {
		delete(p.p, "nicmultiqueuenumber")
	}
}

func (p *DeployVirtualMachineParams) GetNicmultiqueuenumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicmultiqueuenumber"].(int)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetNicnetworklist(v []map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicnetworklist"] = v
}

func (p *DeployVirtualMachineParams) ResetNicnetworklist() {
	if p.p != nil && p.p["nicnetworklist"] != nil {
		delete(p.p, "nicnetworklist")
	}
}

func (p *DeployVirtualMachineParams) GetNicnetworklist() ([]map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicnetworklist"].([]map[string]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) AddNicnetworklist(item map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	val, found := p.p["nicnetworklist"]
	if !found {
		p.p["nicnetworklist"] = []map[string]string{}
		val = p.p["nicnetworklist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	p.p["nicnetworklist"] = l
}

func (p *DeployVirtualMachineParams) SetNicpackedvirtqueuesenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicpackedvirtqueuesenabled"] = v
}

func (p *DeployVirtualMachineParams) ResetNicpackedvirtqueuesenabled() {
	if p.p != nil && p.p["nicpackedvirtqueuesenabled"] != nil {
		delete(p.p, "nicpackedvirtqueuesenabled")
	}
}

func (p *DeployVirtualMachineParams) GetNicpackedvirtqueuesenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicpackedvirtqueuesenabled"].(bool)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetOverridediskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["overridediskofferingid"] = v
}

func (p *DeployVirtualMachineParams) ResetOverridediskofferingid() {
	if p.p != nil && p.p["overridediskofferingid"] != nil {
		delete(p.p, "overridediskofferingid")
	}
}

func (p *DeployVirtualMachineParams) GetOverridediskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["overridediskofferingid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *DeployVirtualMachineParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *DeployVirtualMachineParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *DeployVirtualMachineParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *DeployVirtualMachineParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DeployVirtualMachineParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *DeployVirtualMachineParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetProperties(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["properties"] = v
}

func (p *DeployVirtualMachineParams) ResetProperties() {
	if p.p != nil && p.p["properties"] != nil {
		delete(p.p, "properties")
	}
}

func (p *DeployVirtualMachineParams) GetProperties() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["properties"].(map[string]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetRootdisksize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rootdisksize"] = v
}

func (p *DeployVirtualMachineParams) ResetRootdisksize() {
	if p.p != nil && p.p["rootdisksize"] != nil {
		delete(p.p, "rootdisksize")
	}
}

func (p *DeployVirtualMachineParams) GetRootdisksize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rootdisksize"].(int64)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetSecuritygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupids"] = v
}

func (p *DeployVirtualMachineParams) ResetSecuritygroupids() {
	if p.p != nil && p.p["securitygroupids"] != nil {
		delete(p.p, "securitygroupids")
	}
}

func (p *DeployVirtualMachineParams) GetSecuritygroupids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupids"].([]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetSecuritygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupnames"] = v
}

func (p *DeployVirtualMachineParams) ResetSecuritygroupnames() {
	if p.p != nil && p.p["securitygroupnames"] != nil {
		delete(p.p, "securitygroupnames")
	}
}

func (p *DeployVirtualMachineParams) GetSecuritygroupnames() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupnames"].([]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *DeployVirtualMachineParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *DeployVirtualMachineParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *DeployVirtualMachineParams) ResetSize() {
	if p.p != nil && p.p["size"] != nil {
		delete(p.p, "size")
	}
}

func (p *DeployVirtualMachineParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetStartvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startvm"] = v
}

func (p *DeployVirtualMachineParams) ResetStartvm() {
	if p.p != nil && p.p["startvm"] != nil {
		delete(p.p, "startvm")
	}
}

func (p *DeployVirtualMachineParams) GetStartvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startvm"].(bool)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *DeployVirtualMachineParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *DeployVirtualMachineParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *DeployVirtualMachineParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *DeployVirtualMachineParams) GetUserdata() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetUserdatadetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdatadetails"] = v
}

func (p *DeployVirtualMachineParams) ResetUserdatadetails() {
	if p.p != nil && p.p["userdatadetails"] != nil {
		delete(p.p, "userdatadetails")
	}
}

func (p *DeployVirtualMachineParams) GetUserdatadetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdatadetails"].(map[string]string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetUserdataid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdataid"] = v
}

func (p *DeployVirtualMachineParams) ResetUserdataid() {
	if p.p != nil && p.p["userdataid"] != nil {
		delete(p.p, "userdataid")
	}
}

func (p *DeployVirtualMachineParams) GetUserdataid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdataid"].(string)
	return value, ok
}

func (p *DeployVirtualMachineParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeployVirtualMachineParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeployVirtualMachineParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeployVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDeployVirtualMachineParams(serviceofferingid string, templateid string, zoneid string) *DeployVirtualMachineParams {
	p := &DeployVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["serviceofferingid"] = serviceofferingid
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	return p
}

// Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
func (s *VirtualMachineService) DeployVirtualMachine(p *DeployVirtualMachineParams) (*DeployVirtualMachineResponse, error) {
	resp, err := s.cs.newPostRequest("deployVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeployVirtualMachineResponse
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

type DeployVirtualMachineResponse struct {
	Account               string                                      `json:"account"`
	Affinitygroup         []DeployVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                      `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                      `json:"autoscalevmgroupname"`
	Backupofferingid      string                                      `json:"backupofferingid"`
	Backupofferingname    string                                      `json:"backupofferingname"`
	Bootmode              string                                      `json:"bootmode"`
	Boottype              string                                      `json:"boottype"`
	Cpunumber             int                                         `json:"cpunumber"`
	Cpuspeed              int                                         `json:"cpuspeed"`
	Cpuused               string                                      `json:"cpuused"`
	Created               string                                      `json:"created"`
	Deleteprotection      bool                                        `json:"deleteprotection"`
	Details               map[string]string                           `json:"details"`
	Diskioread            int64                                       `json:"diskioread"`
	Diskiowrite           int64                                       `json:"diskiowrite"`
	Diskkbsread           int64                                       `json:"diskkbsread"`
	Diskkbswrite          int64                                       `json:"diskkbswrite"`
	Diskofferingid        string                                      `json:"diskofferingid"`
	Diskofferingname      string                                      `json:"diskofferingname"`
	Displayname           string                                      `json:"displayname"`
	Displayvm             bool                                        `json:"displayvm"`
	Domain                string                                      `json:"domain"`
	Domainid              string                                      `json:"domainid"`
	Domainpath            string                                      `json:"domainpath"`
	Forvirtualnetwork     bool                                        `json:"forvirtualnetwork"`
	Group                 string                                      `json:"group"`
	Groupid               string                                      `json:"groupid"`
	Guestosid             string                                      `json:"guestosid"`
	Haenable              bool                                        `json:"haenable"`
	Hasannotations        bool                                        `json:"hasannotations"`
	Hostcontrolstate      string                                      `json:"hostcontrolstate"`
	Hostid                string                                      `json:"hostid"`
	Hostname              string                                      `json:"hostname"`
	Hypervisor            string                                      `json:"hypervisor"`
	Icon                  interface{}                                 `json:"icon"`
	Id                    string                                      `json:"id"`
	Instancename          string                                      `json:"instancename"`
	Ipaddress             string                                      `json:"ipaddress"`
	Isdynamicallyscalable bool                                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                      `json:"isodisplaytext"`
	Isoid                 string                                      `json:"isoid"`
	Isoname               string                                      `json:"isoname"`
	JobID                 string                                      `json:"jobid"`
	Jobstatus             int                                         `json:"jobstatus"`
	Keypairs              string                                      `json:"keypairs"`
	Lastupdated           string                                      `json:"lastupdated"`
	Memory                int                                         `json:"memory"`
	Memoryintfreekbs      int64                                       `json:"memoryintfreekbs"`
	Memorykbs             int64                                       `json:"memorykbs"`
	Memorytargetkbs       int64                                       `json:"memorytargetkbs"`
	Name                  string                                      `json:"name"`
	Networkkbsread        int64                                       `json:"networkkbsread"`
	Networkkbswrite       int64                                       `json:"networkkbswrite"`
	Nic                   []Nic                                       `json:"nic"`
	Osdisplayname         string                                      `json:"osdisplayname"`
	Ostypeid              string                                      `json:"ostypeid"`
	Password              string                                      `json:"password"`
	Passwordenabled       bool                                        `json:"passwordenabled"`
	Pooltype              string                                      `json:"pooltype"`
	Project               string                                      `json:"project"`
	Projectid             string                                      `json:"projectid"`
	Publicip              string                                      `json:"publicip"`
	Publicipid            string                                      `json:"publicipid"`
	Readonlydetails       string                                      `json:"readonlydetails"`
	Receivedbytes         int64                                       `json:"receivedbytes"`
	Rootdeviceid          int64                                       `json:"rootdeviceid"`
	Rootdevicetype        string                                      `json:"rootdevicetype"`
	Securitygroup         []DeployVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                       `json:"sentbytes"`
	Serviceofferingid     string                                      `json:"serviceofferingid"`
	Serviceofferingname   string                                      `json:"serviceofferingname"`
	Servicestate          string                                      `json:"servicestate"`
	State                 string                                      `json:"state"`
	Tags                  []Tags                                      `json:"tags"`
	Templatedisplaytext   string                                      `json:"templatedisplaytext"`
	Templateformat        string                                      `json:"templateformat"`
	Templateid            string                                      `json:"templateid"`
	Templatename          string                                      `json:"templatename"`
	Templatetype          string                                      `json:"templatetype"`
	Userdata              string                                      `json:"userdata"`
	Userdatadetails       string                                      `json:"userdatadetails"`
	Userdataid            string                                      `json:"userdataid"`
	Userdataname          string                                      `json:"userdataname"`
	Userdatapolicy        string                                      `json:"userdatapolicy"`
	Userid                string                                      `json:"userid"`
	Username              string                                      `json:"username"`
	Vgpu                  string                                      `json:"vgpu"`
	Vmtype                string                                      `json:"vmtype"`
	Vnfdetails            map[string]string                           `json:"vnfdetails"`
	Vnfnics               []string                                    `json:"vnfnics"`
	Zoneid                string                                      `json:"zoneid"`
	Zonename              string                                      `json:"zonename"`
}

type DeployVirtualMachineResponseSecuritygroup struct {
	Account             string                                          `json:"account"`
	Description         string                                          `json:"description"`
	Domain              string                                          `json:"domain"`
	Domainid            string                                          `json:"domainid"`
	Domainpath          string                                          `json:"domainpath"`
	Egressrule          []DeployVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                          `json:"id"`
	Ingressrule         []DeployVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Projectid           string                                          `json:"projectid"`
	Tags                []Tags                                          `json:"tags"`
	Virtualmachinecount int                                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                   `json:"virtualmachineids"`
}

type DeployVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type DeployVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *DeployVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeployVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DestroyVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *DestroyVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["expunge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("expunge", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["volumeids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("volumeids", vv)
	}
	return u
}

func (p *DestroyVirtualMachineParams) SetExpunge(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expunge"] = v
}

func (p *DestroyVirtualMachineParams) ResetExpunge() {
	if p.p != nil && p.p["expunge"] != nil {
		delete(p.p, "expunge")
	}
}

func (p *DestroyVirtualMachineParams) GetExpunge() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["expunge"].(bool)
	return value, ok
}

func (p *DestroyVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DestroyVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DestroyVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DestroyVirtualMachineParams) SetVolumeids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeids"] = v
}

func (p *DestroyVirtualMachineParams) ResetVolumeids() {
	if p.p != nil && p.p["volumeids"] != nil {
		delete(p.p, "volumeids")
	}
}

func (p *DestroyVirtualMachineParams) GetVolumeids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeids"].([]string)
	return value, ok
}

// You should always use this function to get a new DestroyVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDestroyVirtualMachineParams(id string) *DestroyVirtualMachineParams {
	p := &DestroyVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroys a virtual machine. Once destroyed, only the administrator can recover it.
func (s *VirtualMachineService) DestroyVirtualMachine(p *DestroyVirtualMachineParams) (*DestroyVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("destroyVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroyVirtualMachineResponse
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

type DestroyVirtualMachineResponse struct {
	Account               string                                       `json:"account"`
	Affinitygroup         []DestroyVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                       `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                       `json:"autoscalevmgroupname"`
	Backupofferingid      string                                       `json:"backupofferingid"`
	Backupofferingname    string                                       `json:"backupofferingname"`
	Bootmode              string                                       `json:"bootmode"`
	Boottype              string                                       `json:"boottype"`
	Cpunumber             int                                          `json:"cpunumber"`
	Cpuspeed              int                                          `json:"cpuspeed"`
	Cpuused               string                                       `json:"cpuused"`
	Created               string                                       `json:"created"`
	Deleteprotection      bool                                         `json:"deleteprotection"`
	Details               map[string]string                            `json:"details"`
	Diskioread            int64                                        `json:"diskioread"`
	Diskiowrite           int64                                        `json:"diskiowrite"`
	Diskkbsread           int64                                        `json:"diskkbsread"`
	Diskkbswrite          int64                                        `json:"diskkbswrite"`
	Diskofferingid        string                                       `json:"diskofferingid"`
	Diskofferingname      string                                       `json:"diskofferingname"`
	Displayname           string                                       `json:"displayname"`
	Displayvm             bool                                         `json:"displayvm"`
	Domain                string                                       `json:"domain"`
	Domainid              string                                       `json:"domainid"`
	Domainpath            string                                       `json:"domainpath"`
	Forvirtualnetwork     bool                                         `json:"forvirtualnetwork"`
	Group                 string                                       `json:"group"`
	Groupid               string                                       `json:"groupid"`
	Guestosid             string                                       `json:"guestosid"`
	Haenable              bool                                         `json:"haenable"`
	Hasannotations        bool                                         `json:"hasannotations"`
	Hostcontrolstate      string                                       `json:"hostcontrolstate"`
	Hostid                string                                       `json:"hostid"`
	Hostname              string                                       `json:"hostname"`
	Hypervisor            string                                       `json:"hypervisor"`
	Icon                  interface{}                                  `json:"icon"`
	Id                    string                                       `json:"id"`
	Instancename          string                                       `json:"instancename"`
	Ipaddress             string                                       `json:"ipaddress"`
	Isdynamicallyscalable bool                                         `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                       `json:"isodisplaytext"`
	Isoid                 string                                       `json:"isoid"`
	Isoname               string                                       `json:"isoname"`
	JobID                 string                                       `json:"jobid"`
	Jobstatus             int                                          `json:"jobstatus"`
	Keypairs              string                                       `json:"keypairs"`
	Lastupdated           string                                       `json:"lastupdated"`
	Memory                int                                          `json:"memory"`
	Memoryintfreekbs      int64                                        `json:"memoryintfreekbs"`
	Memorykbs             int64                                        `json:"memorykbs"`
	Memorytargetkbs       int64                                        `json:"memorytargetkbs"`
	Name                  string                                       `json:"name"`
	Networkkbsread        int64                                        `json:"networkkbsread"`
	Networkkbswrite       int64                                        `json:"networkkbswrite"`
	Nic                   []Nic                                        `json:"nic"`
	Osdisplayname         string                                       `json:"osdisplayname"`
	Ostypeid              string                                       `json:"ostypeid"`
	Password              string                                       `json:"password"`
	Passwordenabled       bool                                         `json:"passwordenabled"`
	Pooltype              string                                       `json:"pooltype"`
	Project               string                                       `json:"project"`
	Projectid             string                                       `json:"projectid"`
	Publicip              string                                       `json:"publicip"`
	Publicipid            string                                       `json:"publicipid"`
	Readonlydetails       string                                       `json:"readonlydetails"`
	Receivedbytes         int64                                        `json:"receivedbytes"`
	Rootdeviceid          int64                                        `json:"rootdeviceid"`
	Rootdevicetype        string                                       `json:"rootdevicetype"`
	Securitygroup         []DestroyVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                        `json:"sentbytes"`
	Serviceofferingid     string                                       `json:"serviceofferingid"`
	Serviceofferingname   string                                       `json:"serviceofferingname"`
	Servicestate          string                                       `json:"servicestate"`
	State                 string                                       `json:"state"`
	Tags                  []Tags                                       `json:"tags"`
	Templatedisplaytext   string                                       `json:"templatedisplaytext"`
	Templateformat        string                                       `json:"templateformat"`
	Templateid            string                                       `json:"templateid"`
	Templatename          string                                       `json:"templatename"`
	Templatetype          string                                       `json:"templatetype"`
	Userdata              string                                       `json:"userdata"`
	Userdatadetails       string                                       `json:"userdatadetails"`
	Userdataid            string                                       `json:"userdataid"`
	Userdataname          string                                       `json:"userdataname"`
	Userdatapolicy        string                                       `json:"userdatapolicy"`
	Userid                string                                       `json:"userid"`
	Username              string                                       `json:"username"`
	Vgpu                  string                                       `json:"vgpu"`
	Vmtype                string                                       `json:"vmtype"`
	Vnfdetails            map[string]string                            `json:"vnfdetails"`
	Vnfnics               []string                                     `json:"vnfnics"`
	Zoneid                string                                       `json:"zoneid"`
	Zonename              string                                       `json:"zonename"`
}

type DestroyVirtualMachineResponseSecuritygroup struct {
	Account             string                                           `json:"account"`
	Description         string                                           `json:"description"`
	Domain              string                                           `json:"domain"`
	Domainid            string                                           `json:"domainid"`
	Domainpath          string                                           `json:"domainpath"`
	Egressrule          []DestroyVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                           `json:"id"`
	Ingressrule         []DestroyVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	Projectid           string                                           `json:"projectid"`
	Tags                []Tags                                           `json:"tags"`
	Virtualmachinecount int                                              `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                    `json:"virtualmachineids"`
}

type DestroyVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type DestroyVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *DestroyVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias DestroyVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ExpungeVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ExpungeVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ExpungeVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ExpungeVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ExpungeVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ExpungeVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewExpungeVirtualMachineParams(id string) *ExpungeVirtualMachineParams {
	p := &ExpungeVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Expunge a virtual machine. Once expunged, it cannot be recoverd.
func (s *VirtualMachineService) ExpungeVirtualMachine(p *ExpungeVirtualMachineParams) (*ExpungeVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("expungeVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExpungeVirtualMachineResponse
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

type ExpungeVirtualMachineResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type GetVMPasswordParams struct {
	p map[string]interface{}
}

func (p *GetVMPasswordParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *GetVMPasswordParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *GetVMPasswordParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *GetVMPasswordParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new GetVMPasswordParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewGetVMPasswordParams(id string) *GetVMPasswordParams {
	p := &GetVMPasswordParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Returns an encrypted password for the VM
func (s *VirtualMachineService) GetVMPassword(p *GetVMPasswordParams) (*GetVMPasswordResponse, error) {
	resp, err := s.cs.newRequest("getVMPassword", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetVMPasswordResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetVMPasswordResponse struct {
	Encryptedpassword string `json:"encryptedpassword"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
}

type ListVirtualMachinesParams struct {
	p map[string]interface{}
}

func (p *ListVirtualMachinesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accumulate"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("accumulate", vv)
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["autoscalevmgroupid"]; found {
		u.Set("autoscalevmgroupid", v.(string))
	}
	if v, found := p.p["backupofferingid"]; found {
		u.Set("backupofferingid", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := p.p["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := p.p["groupid"]; found {
		u.Set("groupid", v.(string))
	}
	if v, found := p.p["haenable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("haenable", vv)
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["isoid"]; found {
		u.Set("isoid", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["isvnf"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isvnf", vv)
	}
	if v, found := p.p["keypair"]; found {
		u.Set("keypair", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["retrieveonlyresourcecount"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("retrieveonlyresourcecount", vv)
	}
	if v, found := p.p["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["userdata"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("userdata", vv)
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVirtualMachinesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVirtualMachinesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVirtualMachinesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetAccumulate(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accumulate"] = v
}

func (p *ListVirtualMachinesParams) ResetAccumulate() {
	if p.p != nil && p.p["accumulate"] != nil {
		delete(p.p, "accumulate")
	}
}

func (p *ListVirtualMachinesParams) GetAccumulate() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accumulate"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
}

func (p *ListVirtualMachinesParams) ResetAffinitygroupid() {
	if p.p != nil && p.p["affinitygroupid"] != nil {
		delete(p.p, "affinitygroupid")
	}
}

func (p *ListVirtualMachinesParams) GetAffinitygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetAutoscalevmgroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoscalevmgroupid"] = v
}

func (p *ListVirtualMachinesParams) ResetAutoscalevmgroupid() {
	if p.p != nil && p.p["autoscalevmgroupid"] != nil {
		delete(p.p, "autoscalevmgroupid")
	}
}

func (p *ListVirtualMachinesParams) GetAutoscalevmgroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoscalevmgroupid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetBackupofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["backupofferingid"] = v
}

func (p *ListVirtualMachinesParams) ResetBackupofferingid() {
	if p.p != nil && p.p["backupofferingid"] != nil {
		delete(p.p, "backupofferingid")
	}
}

func (p *ListVirtualMachinesParams) GetBackupofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["backupofferingid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListVirtualMachinesParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListVirtualMachinesParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListVirtualMachinesParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListVirtualMachinesParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
}

func (p *ListVirtualMachinesParams) ResetDisplayvm() {
	if p.p != nil && p.p["displayvm"] != nil {
		delete(p.p, "displayvm")
	}
}

func (p *ListVirtualMachinesParams) GetDisplayvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvm"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVirtualMachinesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVirtualMachinesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetForvirtualnetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvirtualnetwork"] = v
}

func (p *ListVirtualMachinesParams) ResetForvirtualnetwork() {
	if p.p != nil && p.p["forvirtualnetwork"] != nil {
		delete(p.p, "forvirtualnetwork")
	}
}

func (p *ListVirtualMachinesParams) GetForvirtualnetwork() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forvirtualnetwork"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetGroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["groupid"] = v
}

func (p *ListVirtualMachinesParams) ResetGroupid() {
	if p.p != nil && p.p["groupid"] != nil {
		delete(p.p, "groupid")
	}
}

func (p *ListVirtualMachinesParams) GetGroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["groupid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetHaenable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["haenable"] = v
}

func (p *ListVirtualMachinesParams) ResetHaenable() {
	if p.p != nil && p.p["haenable"] != nil {
		delete(p.p, "haenable")
	}
}

func (p *ListVirtualMachinesParams) GetHaenable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["haenable"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListVirtualMachinesParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ListVirtualMachinesParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListVirtualMachinesParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListVirtualMachinesParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVirtualMachinesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVirtualMachinesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListVirtualMachinesParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListVirtualMachinesParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetIsoid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isoid"] = v
}

func (p *ListVirtualMachinesParams) ResetIsoid() {
	if p.p != nil && p.p["isoid"] != nil {
		delete(p.p, "isoid")
	}
}

func (p *ListVirtualMachinesParams) GetIsoid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isoid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVirtualMachinesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVirtualMachinesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetIsvnf(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isvnf"] = v
}

func (p *ListVirtualMachinesParams) ResetIsvnf() {
	if p.p != nil && p.p["isvnf"] != nil {
		delete(p.p, "isvnf")
	}
}

func (p *ListVirtualMachinesParams) GetIsvnf() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isvnf"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
}

func (p *ListVirtualMachinesParams) ResetKeypair() {
	if p.p != nil && p.p["keypair"] != nil {
		delete(p.p, "keypair")
	}
}

func (p *ListVirtualMachinesParams) GetKeypair() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypair"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVirtualMachinesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVirtualMachinesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVirtualMachinesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVirtualMachinesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVirtualMachinesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVirtualMachinesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListVirtualMachinesParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListVirtualMachinesParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVirtualMachinesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVirtualMachinesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVirtualMachinesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVirtualMachinesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListVirtualMachinesParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListVirtualMachinesParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVirtualMachinesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVirtualMachinesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetRetrieveonlyresourcecount(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["retrieveonlyresourcecount"] = v
}

func (p *ListVirtualMachinesParams) ResetRetrieveonlyresourcecount() {
	if p.p != nil && p.p["retrieveonlyresourcecount"] != nil {
		delete(p.p, "retrieveonlyresourcecount")
	}
}

func (p *ListVirtualMachinesParams) GetRetrieveonlyresourcecount() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["retrieveonlyresourcecount"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetSecuritygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupid"] = v
}

func (p *ListVirtualMachinesParams) ResetSecuritygroupid() {
	if p.p != nil && p.p["securitygroupid"] != nil {
		delete(p.p, "securitygroupid")
	}
}

func (p *ListVirtualMachinesParams) GetSecuritygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ListVirtualMachinesParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ListVirtualMachinesParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListVirtualMachinesParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListVirtualMachinesParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVirtualMachinesParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListVirtualMachinesParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListVirtualMachinesParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *ListVirtualMachinesParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVirtualMachinesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListVirtualMachinesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *ListVirtualMachinesParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *ListVirtualMachinesParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetUserdata(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *ListVirtualMachinesParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *ListVirtualMachinesParams) GetUserdata() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *ListVirtualMachinesParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *ListVirtualMachinesParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListVirtualMachinesParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListVirtualMachinesParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVirtualMachinesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVirtualMachinesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVirtualMachinesParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVirtualMachinesParams() *ListVirtualMachinesParams {
	p := &ListVirtualMachinesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVirtualMachinesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVirtualMachines(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VirtualMachines[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VirtualMachines {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineByName(name string, opts ...OptionFunc) (*VirtualMachine, int, error) {
	id, count, err := s.GetVirtualMachineID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVirtualMachineByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error) {
	p := &ListVirtualMachinesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVirtualMachines(p)
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
		return l.VirtualMachines[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualMachine UUID: %s!", id)
}

// List the virtual machines owned by the account.
func (s *VirtualMachineService) ListVirtualMachines(p *ListVirtualMachinesParams) (*ListVirtualMachinesResponse, error) {
	resp, err := s.cs.newRequest("listVirtualMachines", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualMachinesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVirtualMachinesResponse struct {
	Count           int               `json:"count"`
	VirtualMachines []*VirtualMachine `json:"virtualmachine"`
}

type VirtualMachine struct {
	Account               string                        `json:"account"`
	Affinitygroup         []VirtualMachineAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                        `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                        `json:"autoscalevmgroupname"`
	Backupofferingid      string                        `json:"backupofferingid"`
	Backupofferingname    string                        `json:"backupofferingname"`
	Bootmode              string                        `json:"bootmode"`
	Boottype              string                        `json:"boottype"`
	Cpunumber             int                           `json:"cpunumber"`
	Cpuspeed              int                           `json:"cpuspeed"`
	Cpuused               string                        `json:"cpuused"`
	Created               string                        `json:"created"`
	Deleteprotection      bool                          `json:"deleteprotection"`
	Details               map[string]string             `json:"details"`
	Diskioread            int64                         `json:"diskioread"`
	Diskiowrite           int64                         `json:"diskiowrite"`
	Diskkbsread           int64                         `json:"diskkbsread"`
	Diskkbswrite          int64                         `json:"diskkbswrite"`
	Diskofferingid        string                        `json:"diskofferingid"`
	Diskofferingname      string                        `json:"diskofferingname"`
	Displayname           string                        `json:"displayname"`
	Displayvm             bool                          `json:"displayvm"`
	Domain                string                        `json:"domain"`
	Domainid              string                        `json:"domainid"`
	Domainpath            string                        `json:"domainpath"`
	Forvirtualnetwork     bool                          `json:"forvirtualnetwork"`
	Group                 string                        `json:"group"`
	Groupid               string                        `json:"groupid"`
	Guestosid             string                        `json:"guestosid"`
	Haenable              bool                          `json:"haenable"`
	Hasannotations        bool                          `json:"hasannotations"`
	Hostcontrolstate      string                        `json:"hostcontrolstate"`
	Hostid                string                        `json:"hostid"`
	Hostname              string                        `json:"hostname"`
	Hypervisor            string                        `json:"hypervisor"`
	Icon                  interface{}                   `json:"icon"`
	Id                    string                        `json:"id"`
	Instancename          string                        `json:"instancename"`
	Ipaddress             string                        `json:"ipaddress"`
	Isdynamicallyscalable bool                          `json:"isdynamicallyscalable"`
	Isodisplaytext        string                        `json:"isodisplaytext"`
	Isoid                 string                        `json:"isoid"`
	Isoname               string                        `json:"isoname"`
	JobID                 string                        `json:"jobid"`
	Jobstatus             int                           `json:"jobstatus"`
	Keypairs              string                        `json:"keypairs"`
	Lastupdated           string                        `json:"lastupdated"`
	Memory                int                           `json:"memory"`
	Memoryintfreekbs      int64                         `json:"memoryintfreekbs"`
	Memorykbs             int64                         `json:"memorykbs"`
	Memorytargetkbs       int64                         `json:"memorytargetkbs"`
	Name                  string                        `json:"name"`
	Networkkbsread        int64                         `json:"networkkbsread"`
	Networkkbswrite       int64                         `json:"networkkbswrite"`
	Nic                   []Nic                         `json:"nic"`
	Osdisplayname         string                        `json:"osdisplayname"`
	Ostypeid              string                        `json:"ostypeid"`
	Password              string                        `json:"password"`
	Passwordenabled       bool                          `json:"passwordenabled"`
	Pooltype              string                        `json:"pooltype"`
	Project               string                        `json:"project"`
	Projectid             string                        `json:"projectid"`
	Publicip              string                        `json:"publicip"`
	Publicipid            string                        `json:"publicipid"`
	Readonlydetails       string                        `json:"readonlydetails"`
	Receivedbytes         int64                         `json:"receivedbytes"`
	Rootdeviceid          int64                         `json:"rootdeviceid"`
	Rootdevicetype        string                        `json:"rootdevicetype"`
	Securitygroup         []VirtualMachineSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                         `json:"sentbytes"`
	Serviceofferingid     string                        `json:"serviceofferingid"`
	Serviceofferingname   string                        `json:"serviceofferingname"`
	Servicestate          string                        `json:"servicestate"`
	State                 string                        `json:"state"`
	Tags                  []Tags                        `json:"tags"`
	Templatedisplaytext   string                        `json:"templatedisplaytext"`
	Templateformat        string                        `json:"templateformat"`
	Templateid            string                        `json:"templateid"`
	Templatename          string                        `json:"templatename"`
	Templatetype          string                        `json:"templatetype"`
	Userdata              string                        `json:"userdata"`
	Userdatadetails       string                        `json:"userdatadetails"`
	Userdataid            string                        `json:"userdataid"`
	Userdataname          string                        `json:"userdataname"`
	Userdatapolicy        string                        `json:"userdatapolicy"`
	Userid                string                        `json:"userid"`
	Username              string                        `json:"username"`
	Vgpu                  string                        `json:"vgpu"`
	Vmtype                string                        `json:"vmtype"`
	Vnfdetails            map[string]string             `json:"vnfdetails"`
	Vnfnics               []string                      `json:"vnfnics"`
	Zoneid                string                        `json:"zoneid"`
	Zonename              string                        `json:"zonename"`
}

type VirtualMachineSecuritygroup struct {
	Account             string                            `json:"account"`
	Description         string                            `json:"description"`
	Domain              string                            `json:"domain"`
	Domainid            string                            `json:"domainid"`
	Domainpath          string                            `json:"domainpath"`
	Egressrule          []VirtualMachineSecuritygroupRule `json:"egressrule"`
	Id                  string                            `json:"id"`
	Ingressrule         []VirtualMachineSecuritygroupRule `json:"ingressrule"`
	Name                string                            `json:"name"`
	Project             string                            `json:"project"`
	Projectid           string                            `json:"projectid"`
	Tags                []Tags                            `json:"tags"`
	Virtualmachinecount int                               `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                     `json:"virtualmachineids"`
}

type VirtualMachineSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type VirtualMachineAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *VirtualMachine) UnmarshalJSON(b []byte) error {
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

	type alias VirtualMachine
	return json.Unmarshal(b, (*alias)(r))
}

type ListVirtualMachinesMetricsParams struct {
	p map[string]interface{}
}

func (p *ListVirtualMachinesMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accumulate"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("accumulate", vv)
	}
	if v, found := p.p["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := p.p["autoscalevmgroupid"]; found {
		u.Set("autoscalevmgroupid", v.(string))
	}
	if v, found := p.p["backupofferingid"]; found {
		u.Set("backupofferingid", v.(string))
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := p.p["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := p.p["groupid"]; found {
		u.Set("groupid", v.(string))
	}
	if v, found := p.p["haenable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("haenable", vv)
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["isoid"]; found {
		u.Set("isoid", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["isvnf"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isvnf", vv)
	}
	if v, found := p.p["keypair"]; found {
		u.Set("keypair", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["retrieveonlyresourcecount"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("retrieveonlyresourcecount", vv)
	}
	if v, found := p.p["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["userdata"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("userdata", vv)
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVirtualMachinesMetricsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetAccumulate(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accumulate"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetAccumulate() {
	if p.p != nil && p.p["accumulate"] != nil {
		delete(p.p, "accumulate")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetAccumulate() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accumulate"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetAffinitygroupid() {
	if p.p != nil && p.p["affinitygroupid"] != nil {
		delete(p.p, "affinitygroupid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetAffinitygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetAutoscalevmgroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoscalevmgroupid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetAutoscalevmgroupid() {
	if p.p != nil && p.p["autoscalevmgroupid"] != nil {
		delete(p.p, "autoscalevmgroupid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetAutoscalevmgroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoscalevmgroupid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetBackupofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["backupofferingid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetBackupofferingid() {
	if p.p != nil && p.p["backupofferingid"] != nil {
		delete(p.p, "backupofferingid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetBackupofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["backupofferingid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetDisplayvm() {
	if p.p != nil && p.p["displayvm"] != nil {
		delete(p.p, "displayvm")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetDisplayvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvm"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetForvirtualnetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvirtualnetwork"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetForvirtualnetwork() {
	if p.p != nil && p.p["forvirtualnetwork"] != nil {
		delete(p.p, "forvirtualnetwork")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetForvirtualnetwork() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forvirtualnetwork"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetGroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["groupid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetGroupid() {
	if p.p != nil && p.p["groupid"] != nil {
		delete(p.p, "groupid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetGroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["groupid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetHaenable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["haenable"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetHaenable() {
	if p.p != nil && p.p["haenable"] != nil {
		delete(p.p, "haenable")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetHaenable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["haenable"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetIsoid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isoid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetIsoid() {
	if p.p != nil && p.p["isoid"] != nil {
		delete(p.p, "isoid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetIsoid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isoid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetIsvnf(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isvnf"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetIsvnf() {
	if p.p != nil && p.p["isvnf"] != nil {
		delete(p.p, "isvnf")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetIsvnf() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isvnf"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetKeypair() {
	if p.p != nil && p.p["keypair"] != nil {
		delete(p.p, "keypair")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetKeypair() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypair"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetRetrieveonlyresourcecount(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["retrieveonlyresourcecount"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetRetrieveonlyresourcecount() {
	if p.p != nil && p.p["retrieveonlyresourcecount"] != nil {
		delete(p.p, "retrieveonlyresourcecount")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetRetrieveonlyresourcecount() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["retrieveonlyresourcecount"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetSecuritygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetSecuritygroupid() {
	if p.p != nil && p.p["securitygroupid"] != nil {
		delete(p.p, "securitygroupid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetSecuritygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetUserdata(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetUserdata() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(bool)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *ListVirtualMachinesMetricsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVirtualMachinesMetricsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVirtualMachinesMetricsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVirtualMachinesMetricsParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVirtualMachinesMetricsParams() *ListVirtualMachinesMetricsParams {
	p := &ListVirtualMachinesMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesMetricID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVirtualMachinesMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVirtualMachinesMetrics(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VirtualMachinesMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VirtualMachinesMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesMetricByName(name string, opts ...OptionFunc) (*VirtualMachinesMetric, int, error) {
	id, count, err := s.GetVirtualMachinesMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVirtualMachinesMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesMetricByID(id string, opts ...OptionFunc) (*VirtualMachinesMetric, int, error) {
	p := &ListVirtualMachinesMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVirtualMachinesMetrics(p)
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
		return l.VirtualMachinesMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualMachinesMetric UUID: %s!", id)
}

// Lists VM metrics
func (s *VirtualMachineService) ListVirtualMachinesMetrics(p *ListVirtualMachinesMetricsParams) (*ListVirtualMachinesMetricsResponse, error) {
	resp, err := s.cs.newRequest("listVirtualMachinesMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualMachinesMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVirtualMachinesMetricsResponse struct {
	Count                  int                      `json:"count"`
	VirtualMachinesMetrics []*VirtualMachinesMetric `json:"virtualmachine"`
}

type VirtualMachinesMetric struct {
	Account               string                               `json:"account"`
	Affinitygroup         []VirtualMachinesMetricAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                               `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                               `json:"autoscalevmgroupname"`
	Backupofferingid      string                               `json:"backupofferingid"`
	Backupofferingname    string                               `json:"backupofferingname"`
	Bootmode              string                               `json:"bootmode"`
	Boottype              string                               `json:"boottype"`
	Cpunumber             int                                  `json:"cpunumber"`
	Cpuspeed              int                                  `json:"cpuspeed"`
	Cputotal              string                               `json:"cputotal"`
	Cpuused               string                               `json:"cpuused"`
	Created               string                               `json:"created"`
	Deleteprotection      bool                                 `json:"deleteprotection"`
	Details               map[string]string                    `json:"details"`
	Diskiopstotal         int64                                `json:"diskiopstotal"`
	Diskioread            int64                                `json:"diskioread"`
	Diskiowrite           int64                                `json:"diskiowrite"`
	Diskkbsread           int64                                `json:"diskkbsread"`
	Diskkbswrite          int64                                `json:"diskkbswrite"`
	Diskofferingid        string                               `json:"diskofferingid"`
	Diskofferingname      string                               `json:"diskofferingname"`
	Diskread              string                               `json:"diskread"`
	Diskwrite             string                               `json:"diskwrite"`
	Displayname           string                               `json:"displayname"`
	Displayvm             bool                                 `json:"displayvm"`
	Domain                string                               `json:"domain"`
	Domainid              string                               `json:"domainid"`
	Domainpath            string                               `json:"domainpath"`
	Forvirtualnetwork     bool                                 `json:"forvirtualnetwork"`
	Group                 string                               `json:"group"`
	Groupid               string                               `json:"groupid"`
	Guestosid             string                               `json:"guestosid"`
	Haenable              bool                                 `json:"haenable"`
	Hasannotations        bool                                 `json:"hasannotations"`
	Hostcontrolstate      string                               `json:"hostcontrolstate"`
	Hostid                string                               `json:"hostid"`
	Hostname              string                               `json:"hostname"`
	Hypervisor            string                               `json:"hypervisor"`
	Icon                  interface{}                          `json:"icon"`
	Id                    string                               `json:"id"`
	Instancename          string                               `json:"instancename"`
	Ipaddress             string                               `json:"ipaddress"`
	Isdynamicallyscalable bool                                 `json:"isdynamicallyscalable"`
	Isodisplaytext        string                               `json:"isodisplaytext"`
	Isoid                 string                               `json:"isoid"`
	Isoname               string                               `json:"isoname"`
	JobID                 string                               `json:"jobid"`
	Jobstatus             int                                  `json:"jobstatus"`
	Keypairs              string                               `json:"keypairs"`
	Lastupdated           string                               `json:"lastupdated"`
	Memory                int                                  `json:"memory"`
	Memoryintfreekbs      int64                                `json:"memoryintfreekbs"`
	Memorykbs             int64                                `json:"memorykbs"`
	Memorytargetkbs       int64                                `json:"memorytargetkbs"`
	Memorytotal           string                               `json:"memorytotal"`
	Name                  string                               `json:"name"`
	Networkkbsread        int64                                `json:"networkkbsread"`
	Networkkbswrite       int64                                `json:"networkkbswrite"`
	Networkread           string                               `json:"networkread"`
	Networkwrite          string                               `json:"networkwrite"`
	Nic                   []Nic                                `json:"nic"`
	Osdisplayname         string                               `json:"osdisplayname"`
	Ostypeid              string                               `json:"ostypeid"`
	Password              string                               `json:"password"`
	Passwordenabled       bool                                 `json:"passwordenabled"`
	Pooltype              string                               `json:"pooltype"`
	Project               string                               `json:"project"`
	Projectid             string                               `json:"projectid"`
	Publicip              string                               `json:"publicip"`
	Publicipid            string                               `json:"publicipid"`
	Readonlydetails       string                               `json:"readonlydetails"`
	Receivedbytes         int64                                `json:"receivedbytes"`
	Rootdeviceid          int64                                `json:"rootdeviceid"`
	Rootdevicetype        string                               `json:"rootdevicetype"`
	Securitygroup         []VirtualMachinesMetricSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                `json:"sentbytes"`
	Serviceofferingid     string                               `json:"serviceofferingid"`
	Serviceofferingname   string                               `json:"serviceofferingname"`
	Servicestate          string                               `json:"servicestate"`
	State                 string                               `json:"state"`
	Tags                  []Tags                               `json:"tags"`
	Templatedisplaytext   string                               `json:"templatedisplaytext"`
	Templateformat        string                               `json:"templateformat"`
	Templateid            string                               `json:"templateid"`
	Templatename          string                               `json:"templatename"`
	Templatetype          string                               `json:"templatetype"`
	Userdata              string                               `json:"userdata"`
	Userdatadetails       string                               `json:"userdatadetails"`
	Userdataid            string                               `json:"userdataid"`
	Userdataname          string                               `json:"userdataname"`
	Userdatapolicy        string                               `json:"userdatapolicy"`
	Userid                string                               `json:"userid"`
	Username              string                               `json:"username"`
	Vgpu                  string                               `json:"vgpu"`
	Vmtype                string                               `json:"vmtype"`
	Vnfdetails            map[string]string                    `json:"vnfdetails"`
	Vnfnics               []string                             `json:"vnfnics"`
	Zoneid                string                               `json:"zoneid"`
	Zonename              string                               `json:"zonename"`
}

type VirtualMachinesMetricSecuritygroup struct {
	Account             string                                   `json:"account"`
	Description         string                                   `json:"description"`
	Domain              string                                   `json:"domain"`
	Domainid            string                                   `json:"domainid"`
	Domainpath          string                                   `json:"domainpath"`
	Egressrule          []VirtualMachinesMetricSecuritygroupRule `json:"egressrule"`
	Id                  string                                   `json:"id"`
	Ingressrule         []VirtualMachinesMetricSecuritygroupRule `json:"ingressrule"`
	Name                string                                   `json:"name"`
	Project             string                                   `json:"project"`
	Projectid           string                                   `json:"projectid"`
	Tags                []Tags                                   `json:"tags"`
	Virtualmachinecount int                                      `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                            `json:"virtualmachineids"`
}

type VirtualMachinesMetricSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type VirtualMachinesMetricAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *VirtualMachinesMetric) UnmarshalJSON(b []byte) error {
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

	type alias VirtualMachinesMetric
	return json.Unmarshal(b, (*alias)(r))
}

type ListVmsForImportParams struct {
	p map[string]interface{}
}

func (p *ListVmsForImportParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["host"]; found {
		u.Set("host", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVmsForImportParams) SetHost(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["host"] = v
}

func (p *ListVmsForImportParams) ResetHost() {
	if p.p != nil && p.p["host"] != nil {
		delete(p.p, "host")
	}
}

func (p *ListVmsForImportParams) GetHost() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["host"].(string)
	return value, ok
}

func (p *ListVmsForImportParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListVmsForImportParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListVmsForImportParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListVmsForImportParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVmsForImportParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVmsForImportParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVmsForImportParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVmsForImportParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVmsForImportParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVmsForImportParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVmsForImportParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVmsForImportParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVmsForImportParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *ListVmsForImportParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *ListVmsForImportParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *ListVmsForImportParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *ListVmsForImportParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *ListVmsForImportParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *ListVmsForImportParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVmsForImportParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVmsForImportParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVmsForImportParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVmsForImportParams(host string, hypervisor string, zoneid string) *ListVmsForImportParams {
	p := &ListVmsForImportParams{}
	p.p = make(map[string]interface{})
	p.p["host"] = host
	p.p["hypervisor"] = hypervisor
	p.p["zoneid"] = zoneid
	return p
}

// Lists virtual machines on a unmanaged host
func (s *VirtualMachineService) ListVmsForImport(p *ListVmsForImportParams) (*ListVmsForImportResponse, error) {
	resp, err := s.cs.newRequest("listVmsForImport", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVmsForImportResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVmsForImportResponse struct {
	Count        int             `json:"count"`
	VmsForImport []*VmsForImport `json:"vmsforimport"`
}

type VmsForImport struct {
	Clusterid        string             `json:"clusterid"`
	Clustername      string             `json:"clustername"`
	Cpucorepersocket int                `json:"cpucorepersocket"`
	Cpunumber        int                `json:"cpunumber"`
	Cpuspeed         int                `json:"cpuspeed"`
	Disk             []VmsForImportDisk `json:"disk"`
	Hostid           string             `json:"hostid"`
	Hostname         string             `json:"hostname"`
	JobID            string             `json:"jobid"`
	Jobstatus        int                `json:"jobstatus"`
	Memory           int                `json:"memory"`
	Name             string             `json:"name"`
	Nic              []Nic              `json:"nic"`
	Osdisplayname    string             `json:"osdisplayname"`
	Osid             string             `json:"osid"`
	Powerstate       string             `json:"powerstate"`
}

type VmsForImportDisk struct {
	Capacity       int64  `json:"capacity"`
	Controller     string `json:"controller"`
	Controllerunit int    `json:"controllerunit"`
	Datastorehost  string `json:"datastorehost"`
	Datastorename  string `json:"datastorename"`
	Datastorepath  string `json:"datastorepath"`
	Datastoretype  string `json:"datastoretype"`
	Id             string `json:"id"`
	Imagepath      string `json:"imagepath"`
	Label          string `json:"label"`
	Position       int    `json:"position"`
}

type MigrateVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *MigrateVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["autoselect"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("autoselect", vv)
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *MigrateVirtualMachineParams) SetAutoselect(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoselect"] = v
}

func (p *MigrateVirtualMachineParams) ResetAutoselect() {
	if p.p != nil && p.p["autoselect"] != nil {
		delete(p.p, "autoselect")
	}
}

func (p *MigrateVirtualMachineParams) GetAutoselect() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoselect"].(bool)
	return value, ok
}

func (p *MigrateVirtualMachineParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *MigrateVirtualMachineParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *MigrateVirtualMachineParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *MigrateVirtualMachineParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *MigrateVirtualMachineParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *MigrateVirtualMachineParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *MigrateVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *MigrateVirtualMachineParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *MigrateVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewMigrateVirtualMachineParams(virtualmachineid string) *MigrateVirtualMachineParams {
	p := &MigrateVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attempts Migration of a VM to a different host or Root volume of the vm to a different storage pool
func (s *VirtualMachineService) MigrateVirtualMachine(p *MigrateVirtualMachineParams) (*MigrateVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("migrateVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVirtualMachineResponse
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

type MigrateVirtualMachineResponse struct {
	Account               string                                       `json:"account"`
	Affinitygroup         []MigrateVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                       `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                       `json:"autoscalevmgroupname"`
	Backupofferingid      string                                       `json:"backupofferingid"`
	Backupofferingname    string                                       `json:"backupofferingname"`
	Bootmode              string                                       `json:"bootmode"`
	Boottype              string                                       `json:"boottype"`
	Cpunumber             int                                          `json:"cpunumber"`
	Cpuspeed              int                                          `json:"cpuspeed"`
	Cpuused               string                                       `json:"cpuused"`
	Created               string                                       `json:"created"`
	Deleteprotection      bool                                         `json:"deleteprotection"`
	Details               map[string]string                            `json:"details"`
	Diskioread            int64                                        `json:"diskioread"`
	Diskiowrite           int64                                        `json:"diskiowrite"`
	Diskkbsread           int64                                        `json:"diskkbsread"`
	Diskkbswrite          int64                                        `json:"diskkbswrite"`
	Diskofferingid        string                                       `json:"diskofferingid"`
	Diskofferingname      string                                       `json:"diskofferingname"`
	Displayname           string                                       `json:"displayname"`
	Displayvm             bool                                         `json:"displayvm"`
	Domain                string                                       `json:"domain"`
	Domainid              string                                       `json:"domainid"`
	Domainpath            string                                       `json:"domainpath"`
	Forvirtualnetwork     bool                                         `json:"forvirtualnetwork"`
	Group                 string                                       `json:"group"`
	Groupid               string                                       `json:"groupid"`
	Guestosid             string                                       `json:"guestosid"`
	Haenable              bool                                         `json:"haenable"`
	Hasannotations        bool                                         `json:"hasannotations"`
	Hostcontrolstate      string                                       `json:"hostcontrolstate"`
	Hostid                string                                       `json:"hostid"`
	Hostname              string                                       `json:"hostname"`
	Hypervisor            string                                       `json:"hypervisor"`
	Icon                  interface{}                                  `json:"icon"`
	Id                    string                                       `json:"id"`
	Instancename          string                                       `json:"instancename"`
	Ipaddress             string                                       `json:"ipaddress"`
	Isdynamicallyscalable bool                                         `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                       `json:"isodisplaytext"`
	Isoid                 string                                       `json:"isoid"`
	Isoname               string                                       `json:"isoname"`
	JobID                 string                                       `json:"jobid"`
	Jobstatus             int                                          `json:"jobstatus"`
	Keypairs              string                                       `json:"keypairs"`
	Lastupdated           string                                       `json:"lastupdated"`
	Memory                int                                          `json:"memory"`
	Memoryintfreekbs      int64                                        `json:"memoryintfreekbs"`
	Memorykbs             int64                                        `json:"memorykbs"`
	Memorytargetkbs       int64                                        `json:"memorytargetkbs"`
	Name                  string                                       `json:"name"`
	Networkkbsread        int64                                        `json:"networkkbsread"`
	Networkkbswrite       int64                                        `json:"networkkbswrite"`
	Nic                   []Nic                                        `json:"nic"`
	Osdisplayname         string                                       `json:"osdisplayname"`
	Ostypeid              string                                       `json:"ostypeid"`
	Password              string                                       `json:"password"`
	Passwordenabled       bool                                         `json:"passwordenabled"`
	Pooltype              string                                       `json:"pooltype"`
	Project               string                                       `json:"project"`
	Projectid             string                                       `json:"projectid"`
	Publicip              string                                       `json:"publicip"`
	Publicipid            string                                       `json:"publicipid"`
	Readonlydetails       string                                       `json:"readonlydetails"`
	Receivedbytes         int64                                        `json:"receivedbytes"`
	Rootdeviceid          int64                                        `json:"rootdeviceid"`
	Rootdevicetype        string                                       `json:"rootdevicetype"`
	Securitygroup         []MigrateVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                        `json:"sentbytes"`
	Serviceofferingid     string                                       `json:"serviceofferingid"`
	Serviceofferingname   string                                       `json:"serviceofferingname"`
	Servicestate          string                                       `json:"servicestate"`
	State                 string                                       `json:"state"`
	Tags                  []Tags                                       `json:"tags"`
	Templatedisplaytext   string                                       `json:"templatedisplaytext"`
	Templateformat        string                                       `json:"templateformat"`
	Templateid            string                                       `json:"templateid"`
	Templatename          string                                       `json:"templatename"`
	Templatetype          string                                       `json:"templatetype"`
	Userdata              string                                       `json:"userdata"`
	Userdatadetails       string                                       `json:"userdatadetails"`
	Userdataid            string                                       `json:"userdataid"`
	Userdataname          string                                       `json:"userdataname"`
	Userdatapolicy        string                                       `json:"userdatapolicy"`
	Userid                string                                       `json:"userid"`
	Username              string                                       `json:"username"`
	Vgpu                  string                                       `json:"vgpu"`
	Vmtype                string                                       `json:"vmtype"`
	Vnfdetails            map[string]string                            `json:"vnfdetails"`
	Vnfnics               []string                                     `json:"vnfnics"`
	Zoneid                string                                       `json:"zoneid"`
	Zonename              string                                       `json:"zonename"`
}

type MigrateVirtualMachineResponseSecuritygroup struct {
	Account             string                                           `json:"account"`
	Description         string                                           `json:"description"`
	Domain              string                                           `json:"domain"`
	Domainid            string                                           `json:"domainid"`
	Domainpath          string                                           `json:"domainpath"`
	Egressrule          []MigrateVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                           `json:"id"`
	Ingressrule         []MigrateVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	Projectid           string                                           `json:"projectid"`
	Tags                []Tags                                           `json:"tags"`
	Virtualmachinecount int                                              `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                    `json:"virtualmachineids"`
}

type MigrateVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type MigrateVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *MigrateVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias MigrateVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type MigrateVirtualMachineWithVolumeParams struct {
	p map[string]interface{}
}

func (p *MigrateVirtualMachineWithVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["autoselect"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("autoselect", vv)
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["migrateto"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("migrateto[%d].%s", i, key), val)
			}
		}
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *MigrateVirtualMachineWithVolumeParams) SetAutoselect(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoselect"] = v
}

func (p *MigrateVirtualMachineWithVolumeParams) ResetAutoselect() {
	if p.p != nil && p.p["autoselect"] != nil {
		delete(p.p, "autoselect")
	}
}

func (p *MigrateVirtualMachineWithVolumeParams) GetAutoselect() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoselect"].(bool)
	return value, ok
}

func (p *MigrateVirtualMachineWithVolumeParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *MigrateVirtualMachineWithVolumeParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *MigrateVirtualMachineWithVolumeParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *MigrateVirtualMachineWithVolumeParams) SetMigrateto(v []map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["migrateto"] = v
}

func (p *MigrateVirtualMachineWithVolumeParams) ResetMigrateto() {
	if p.p != nil && p.p["migrateto"] != nil {
		delete(p.p, "migrateto")
	}
}

func (p *MigrateVirtualMachineWithVolumeParams) GetMigrateto() ([]map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["migrateto"].([]map[string]string)
	return value, ok
}

func (p *MigrateVirtualMachineWithVolumeParams) AddMigrateto(item map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	val, found := p.p["migrateto"]
	if !found {
		p.p["migrateto"] = []map[string]string{}
		val = p.p["migrateto"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	p.p["migrateto"] = l
}

func (p *MigrateVirtualMachineWithVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *MigrateVirtualMachineWithVolumeParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *MigrateVirtualMachineWithVolumeParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateVirtualMachineWithVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewMigrateVirtualMachineWithVolumeParams(virtualmachineid string) *MigrateVirtualMachineWithVolumeParams {
	p := &MigrateVirtualMachineWithVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attempts Migration of a VM with its volumes to a different host
func (s *VirtualMachineService) MigrateVirtualMachineWithVolume(p *MigrateVirtualMachineWithVolumeParams) (*MigrateVirtualMachineWithVolumeResponse, error) {
	resp, err := s.cs.newRequest("migrateVirtualMachineWithVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVirtualMachineWithVolumeResponse
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

type MigrateVirtualMachineWithVolumeResponse struct {
	Account               string                                                 `json:"account"`
	Affinitygroup         []MigrateVirtualMachineWithVolumeResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                                 `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                                 `json:"autoscalevmgroupname"`
	Backupofferingid      string                                                 `json:"backupofferingid"`
	Backupofferingname    string                                                 `json:"backupofferingname"`
	Bootmode              string                                                 `json:"bootmode"`
	Boottype              string                                                 `json:"boottype"`
	Cpunumber             int                                                    `json:"cpunumber"`
	Cpuspeed              int                                                    `json:"cpuspeed"`
	Cpuused               string                                                 `json:"cpuused"`
	Created               string                                                 `json:"created"`
	Deleteprotection      bool                                                   `json:"deleteprotection"`
	Details               map[string]string                                      `json:"details"`
	Diskioread            int64                                                  `json:"diskioread"`
	Diskiowrite           int64                                                  `json:"diskiowrite"`
	Diskkbsread           int64                                                  `json:"diskkbsread"`
	Diskkbswrite          int64                                                  `json:"diskkbswrite"`
	Diskofferingid        string                                                 `json:"diskofferingid"`
	Diskofferingname      string                                                 `json:"diskofferingname"`
	Displayname           string                                                 `json:"displayname"`
	Displayvm             bool                                                   `json:"displayvm"`
	Domain                string                                                 `json:"domain"`
	Domainid              string                                                 `json:"domainid"`
	Domainpath            string                                                 `json:"domainpath"`
	Forvirtualnetwork     bool                                                   `json:"forvirtualnetwork"`
	Group                 string                                                 `json:"group"`
	Groupid               string                                                 `json:"groupid"`
	Guestosid             string                                                 `json:"guestosid"`
	Haenable              bool                                                   `json:"haenable"`
	Hasannotations        bool                                                   `json:"hasannotations"`
	Hostcontrolstate      string                                                 `json:"hostcontrolstate"`
	Hostid                string                                                 `json:"hostid"`
	Hostname              string                                                 `json:"hostname"`
	Hypervisor            string                                                 `json:"hypervisor"`
	Icon                  interface{}                                            `json:"icon"`
	Id                    string                                                 `json:"id"`
	Instancename          string                                                 `json:"instancename"`
	Ipaddress             string                                                 `json:"ipaddress"`
	Isdynamicallyscalable bool                                                   `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                 `json:"isodisplaytext"`
	Isoid                 string                                                 `json:"isoid"`
	Isoname               string                                                 `json:"isoname"`
	JobID                 string                                                 `json:"jobid"`
	Jobstatus             int                                                    `json:"jobstatus"`
	Keypairs              string                                                 `json:"keypairs"`
	Lastupdated           string                                                 `json:"lastupdated"`
	Memory                int                                                    `json:"memory"`
	Memoryintfreekbs      int64                                                  `json:"memoryintfreekbs"`
	Memorykbs             int64                                                  `json:"memorykbs"`
	Memorytargetkbs       int64                                                  `json:"memorytargetkbs"`
	Name                  string                                                 `json:"name"`
	Networkkbsread        int64                                                  `json:"networkkbsread"`
	Networkkbswrite       int64                                                  `json:"networkkbswrite"`
	Nic                   []Nic                                                  `json:"nic"`
	Osdisplayname         string                                                 `json:"osdisplayname"`
	Ostypeid              string                                                 `json:"ostypeid"`
	Password              string                                                 `json:"password"`
	Passwordenabled       bool                                                   `json:"passwordenabled"`
	Pooltype              string                                                 `json:"pooltype"`
	Project               string                                                 `json:"project"`
	Projectid             string                                                 `json:"projectid"`
	Publicip              string                                                 `json:"publicip"`
	Publicipid            string                                                 `json:"publicipid"`
	Readonlydetails       string                                                 `json:"readonlydetails"`
	Receivedbytes         int64                                                  `json:"receivedbytes"`
	Rootdeviceid          int64                                                  `json:"rootdeviceid"`
	Rootdevicetype        string                                                 `json:"rootdevicetype"`
	Securitygroup         []MigrateVirtualMachineWithVolumeResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                  `json:"sentbytes"`
	Serviceofferingid     string                                                 `json:"serviceofferingid"`
	Serviceofferingname   string                                                 `json:"serviceofferingname"`
	Servicestate          string                                                 `json:"servicestate"`
	State                 string                                                 `json:"state"`
	Tags                  []Tags                                                 `json:"tags"`
	Templatedisplaytext   string                                                 `json:"templatedisplaytext"`
	Templateformat        string                                                 `json:"templateformat"`
	Templateid            string                                                 `json:"templateid"`
	Templatename          string                                                 `json:"templatename"`
	Templatetype          string                                                 `json:"templatetype"`
	Userdata              string                                                 `json:"userdata"`
	Userdatadetails       string                                                 `json:"userdatadetails"`
	Userdataid            string                                                 `json:"userdataid"`
	Userdataname          string                                                 `json:"userdataname"`
	Userdatapolicy        string                                                 `json:"userdatapolicy"`
	Userid                string                                                 `json:"userid"`
	Username              string                                                 `json:"username"`
	Vgpu                  string                                                 `json:"vgpu"`
	Vmtype                string                                                 `json:"vmtype"`
	Vnfdetails            map[string]string                                      `json:"vnfdetails"`
	Vnfnics               []string                                               `json:"vnfnics"`
	Zoneid                string                                                 `json:"zoneid"`
	Zonename              string                                                 `json:"zonename"`
}

type MigrateVirtualMachineWithVolumeResponseSecuritygroup struct {
	Account             string                                                     `json:"account"`
	Description         string                                                     `json:"description"`
	Domain              string                                                     `json:"domain"`
	Domainid            string                                                     `json:"domainid"`
	Domainpath          string                                                     `json:"domainpath"`
	Egressrule          []MigrateVirtualMachineWithVolumeResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                     `json:"id"`
	Ingressrule         []MigrateVirtualMachineWithVolumeResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                     `json:"name"`
	Project             string                                                     `json:"project"`
	Projectid           string                                                     `json:"projectid"`
	Tags                []Tags                                                     `json:"tags"`
	Virtualmachinecount int                                                        `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                              `json:"virtualmachineids"`
}

type MigrateVirtualMachineWithVolumeResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type MigrateVirtualMachineWithVolumeResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *MigrateVirtualMachineWithVolumeResponse) UnmarshalJSON(b []byte) error {
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

	type alias MigrateVirtualMachineWithVolumeResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RebootVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RebootVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bootintosetup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootintosetup", vv)
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

func (p *RebootVirtualMachineParams) SetBootintosetup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootintosetup"] = v
}

func (p *RebootVirtualMachineParams) ResetBootintosetup() {
	if p.p != nil && p.p["bootintosetup"] != nil {
		delete(p.p, "bootintosetup")
	}
}

func (p *RebootVirtualMachineParams) GetBootintosetup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootintosetup"].(bool)
	return value, ok
}

func (p *RebootVirtualMachineParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *RebootVirtualMachineParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *RebootVirtualMachineParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *RebootVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RebootVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RebootVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RebootVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRebootVirtualMachineParams(id string) *RebootVirtualMachineParams {
	p := &RebootVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reboots a virtual machine.
func (s *VirtualMachineService) RebootVirtualMachine(p *RebootVirtualMachineParams) (*RebootVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("rebootVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootVirtualMachineResponse
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

type RebootVirtualMachineResponse struct {
	Account               string                                      `json:"account"`
	Affinitygroup         []RebootVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                      `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                      `json:"autoscalevmgroupname"`
	Backupofferingid      string                                      `json:"backupofferingid"`
	Backupofferingname    string                                      `json:"backupofferingname"`
	Bootmode              string                                      `json:"bootmode"`
	Boottype              string                                      `json:"boottype"`
	Cpunumber             int                                         `json:"cpunumber"`
	Cpuspeed              int                                         `json:"cpuspeed"`
	Cpuused               string                                      `json:"cpuused"`
	Created               string                                      `json:"created"`
	Deleteprotection      bool                                        `json:"deleteprotection"`
	Details               map[string]string                           `json:"details"`
	Diskioread            int64                                       `json:"diskioread"`
	Diskiowrite           int64                                       `json:"diskiowrite"`
	Diskkbsread           int64                                       `json:"diskkbsread"`
	Diskkbswrite          int64                                       `json:"diskkbswrite"`
	Diskofferingid        string                                      `json:"diskofferingid"`
	Diskofferingname      string                                      `json:"diskofferingname"`
	Displayname           string                                      `json:"displayname"`
	Displayvm             bool                                        `json:"displayvm"`
	Domain                string                                      `json:"domain"`
	Domainid              string                                      `json:"domainid"`
	Domainpath            string                                      `json:"domainpath"`
	Forvirtualnetwork     bool                                        `json:"forvirtualnetwork"`
	Group                 string                                      `json:"group"`
	Groupid               string                                      `json:"groupid"`
	Guestosid             string                                      `json:"guestosid"`
	Haenable              bool                                        `json:"haenable"`
	Hasannotations        bool                                        `json:"hasannotations"`
	Hostcontrolstate      string                                      `json:"hostcontrolstate"`
	Hostid                string                                      `json:"hostid"`
	Hostname              string                                      `json:"hostname"`
	Hypervisor            string                                      `json:"hypervisor"`
	Icon                  interface{}                                 `json:"icon"`
	Id                    string                                      `json:"id"`
	Instancename          string                                      `json:"instancename"`
	Ipaddress             string                                      `json:"ipaddress"`
	Isdynamicallyscalable bool                                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                      `json:"isodisplaytext"`
	Isoid                 string                                      `json:"isoid"`
	Isoname               string                                      `json:"isoname"`
	JobID                 string                                      `json:"jobid"`
	Jobstatus             int                                         `json:"jobstatus"`
	Keypairs              string                                      `json:"keypairs"`
	Lastupdated           string                                      `json:"lastupdated"`
	Memory                int                                         `json:"memory"`
	Memoryintfreekbs      int64                                       `json:"memoryintfreekbs"`
	Memorykbs             int64                                       `json:"memorykbs"`
	Memorytargetkbs       int64                                       `json:"memorytargetkbs"`
	Name                  string                                      `json:"name"`
	Networkkbsread        int64                                       `json:"networkkbsread"`
	Networkkbswrite       int64                                       `json:"networkkbswrite"`
	Nic                   []Nic                                       `json:"nic"`
	Osdisplayname         string                                      `json:"osdisplayname"`
	Ostypeid              string                                      `json:"ostypeid"`
	Password              string                                      `json:"password"`
	Passwordenabled       bool                                        `json:"passwordenabled"`
	Pooltype              string                                      `json:"pooltype"`
	Project               string                                      `json:"project"`
	Projectid             string                                      `json:"projectid"`
	Publicip              string                                      `json:"publicip"`
	Publicipid            string                                      `json:"publicipid"`
	Readonlydetails       string                                      `json:"readonlydetails"`
	Receivedbytes         int64                                       `json:"receivedbytes"`
	Rootdeviceid          int64                                       `json:"rootdeviceid"`
	Rootdevicetype        string                                      `json:"rootdevicetype"`
	Securitygroup         []RebootVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                       `json:"sentbytes"`
	Serviceofferingid     string                                      `json:"serviceofferingid"`
	Serviceofferingname   string                                      `json:"serviceofferingname"`
	Servicestate          string                                      `json:"servicestate"`
	State                 string                                      `json:"state"`
	Tags                  []Tags                                      `json:"tags"`
	Templatedisplaytext   string                                      `json:"templatedisplaytext"`
	Templateformat        string                                      `json:"templateformat"`
	Templateid            string                                      `json:"templateid"`
	Templatename          string                                      `json:"templatename"`
	Templatetype          string                                      `json:"templatetype"`
	Userdata              string                                      `json:"userdata"`
	Userdatadetails       string                                      `json:"userdatadetails"`
	Userdataid            string                                      `json:"userdataid"`
	Userdataname          string                                      `json:"userdataname"`
	Userdatapolicy        string                                      `json:"userdatapolicy"`
	Userid                string                                      `json:"userid"`
	Username              string                                      `json:"username"`
	Vgpu                  string                                      `json:"vgpu"`
	Vmtype                string                                      `json:"vmtype"`
	Vnfdetails            map[string]string                           `json:"vnfdetails"`
	Vnfnics               []string                                    `json:"vnfnics"`
	Zoneid                string                                      `json:"zoneid"`
	Zonename              string                                      `json:"zonename"`
}

type RebootVirtualMachineResponseSecuritygroup struct {
	Account             string                                          `json:"account"`
	Description         string                                          `json:"description"`
	Domain              string                                          `json:"domain"`
	Domainid            string                                          `json:"domainid"`
	Domainpath          string                                          `json:"domainpath"`
	Egressrule          []RebootVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                          `json:"id"`
	Ingressrule         []RebootVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Projectid           string                                          `json:"projectid"`
	Tags                []Tags                                          `json:"tags"`
	Virtualmachinecount int                                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                   `json:"virtualmachineids"`
}

type RebootVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type RebootVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *RebootVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias RebootVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RecoverVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RecoverVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RecoverVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RecoverVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RecoverVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RecoverVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRecoverVirtualMachineParams(id string) *RecoverVirtualMachineParams {
	p := &RecoverVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Recovers a virtual machine.
func (s *VirtualMachineService) RecoverVirtualMachine(p *RecoverVirtualMachineParams) (*RecoverVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("recoverVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RecoverVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RecoverVirtualMachineResponse struct {
	Account               string                                       `json:"account"`
	Affinitygroup         []RecoverVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                       `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                       `json:"autoscalevmgroupname"`
	Backupofferingid      string                                       `json:"backupofferingid"`
	Backupofferingname    string                                       `json:"backupofferingname"`
	Bootmode              string                                       `json:"bootmode"`
	Boottype              string                                       `json:"boottype"`
	Cpunumber             int                                          `json:"cpunumber"`
	Cpuspeed              int                                          `json:"cpuspeed"`
	Cpuused               string                                       `json:"cpuused"`
	Created               string                                       `json:"created"`
	Deleteprotection      bool                                         `json:"deleteprotection"`
	Details               map[string]string                            `json:"details"`
	Diskioread            int64                                        `json:"diskioread"`
	Diskiowrite           int64                                        `json:"diskiowrite"`
	Diskkbsread           int64                                        `json:"diskkbsread"`
	Diskkbswrite          int64                                        `json:"diskkbswrite"`
	Diskofferingid        string                                       `json:"diskofferingid"`
	Diskofferingname      string                                       `json:"diskofferingname"`
	Displayname           string                                       `json:"displayname"`
	Displayvm             bool                                         `json:"displayvm"`
	Domain                string                                       `json:"domain"`
	Domainid              string                                       `json:"domainid"`
	Domainpath            string                                       `json:"domainpath"`
	Forvirtualnetwork     bool                                         `json:"forvirtualnetwork"`
	Group                 string                                       `json:"group"`
	Groupid               string                                       `json:"groupid"`
	Guestosid             string                                       `json:"guestosid"`
	Haenable              bool                                         `json:"haenable"`
	Hasannotations        bool                                         `json:"hasannotations"`
	Hostcontrolstate      string                                       `json:"hostcontrolstate"`
	Hostid                string                                       `json:"hostid"`
	Hostname              string                                       `json:"hostname"`
	Hypervisor            string                                       `json:"hypervisor"`
	Icon                  interface{}                                  `json:"icon"`
	Id                    string                                       `json:"id"`
	Instancename          string                                       `json:"instancename"`
	Ipaddress             string                                       `json:"ipaddress"`
	Isdynamicallyscalable bool                                         `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                       `json:"isodisplaytext"`
	Isoid                 string                                       `json:"isoid"`
	Isoname               string                                       `json:"isoname"`
	JobID                 string                                       `json:"jobid"`
	Jobstatus             int                                          `json:"jobstatus"`
	Keypairs              string                                       `json:"keypairs"`
	Lastupdated           string                                       `json:"lastupdated"`
	Memory                int                                          `json:"memory"`
	Memoryintfreekbs      int64                                        `json:"memoryintfreekbs"`
	Memorykbs             int64                                        `json:"memorykbs"`
	Memorytargetkbs       int64                                        `json:"memorytargetkbs"`
	Name                  string                                       `json:"name"`
	Networkkbsread        int64                                        `json:"networkkbsread"`
	Networkkbswrite       int64                                        `json:"networkkbswrite"`
	Nic                   []Nic                                        `json:"nic"`
	Osdisplayname         string                                       `json:"osdisplayname"`
	Ostypeid              string                                       `json:"ostypeid"`
	Password              string                                       `json:"password"`
	Passwordenabled       bool                                         `json:"passwordenabled"`
	Pooltype              string                                       `json:"pooltype"`
	Project               string                                       `json:"project"`
	Projectid             string                                       `json:"projectid"`
	Publicip              string                                       `json:"publicip"`
	Publicipid            string                                       `json:"publicipid"`
	Readonlydetails       string                                       `json:"readonlydetails"`
	Receivedbytes         int64                                        `json:"receivedbytes"`
	Rootdeviceid          int64                                        `json:"rootdeviceid"`
	Rootdevicetype        string                                       `json:"rootdevicetype"`
	Securitygroup         []RecoverVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                        `json:"sentbytes"`
	Serviceofferingid     string                                       `json:"serviceofferingid"`
	Serviceofferingname   string                                       `json:"serviceofferingname"`
	Servicestate          string                                       `json:"servicestate"`
	State                 string                                       `json:"state"`
	Tags                  []Tags                                       `json:"tags"`
	Templatedisplaytext   string                                       `json:"templatedisplaytext"`
	Templateformat        string                                       `json:"templateformat"`
	Templateid            string                                       `json:"templateid"`
	Templatename          string                                       `json:"templatename"`
	Templatetype          string                                       `json:"templatetype"`
	Userdata              string                                       `json:"userdata"`
	Userdatadetails       string                                       `json:"userdatadetails"`
	Userdataid            string                                       `json:"userdataid"`
	Userdataname          string                                       `json:"userdataname"`
	Userdatapolicy        string                                       `json:"userdatapolicy"`
	Userid                string                                       `json:"userid"`
	Username              string                                       `json:"username"`
	Vgpu                  string                                       `json:"vgpu"`
	Vmtype                string                                       `json:"vmtype"`
	Vnfdetails            map[string]string                            `json:"vnfdetails"`
	Vnfnics               []string                                     `json:"vnfnics"`
	Zoneid                string                                       `json:"zoneid"`
	Zonename              string                                       `json:"zonename"`
}

type RecoverVirtualMachineResponseSecuritygroup struct {
	Account             string                                           `json:"account"`
	Description         string                                           `json:"description"`
	Domain              string                                           `json:"domain"`
	Domainid            string                                           `json:"domainid"`
	Domainpath          string                                           `json:"domainpath"`
	Egressrule          []RecoverVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                           `json:"id"`
	Ingressrule         []RecoverVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	Projectid           string                                           `json:"projectid"`
	Tags                []Tags                                           `json:"tags"`
	Virtualmachinecount int                                              `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                    `json:"virtualmachineids"`
}

type RecoverVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type RecoverVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *RecoverVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias RecoverVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RemoveNicFromVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RemoveNicFromVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *RemoveNicFromVirtualMachineParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
}

func (p *RemoveNicFromVirtualMachineParams) ResetNicid() {
	if p.p != nil && p.p["nicid"] != nil {
		delete(p.p, "nicid")
	}
}

func (p *RemoveNicFromVirtualMachineParams) GetNicid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicid"].(string)
	return value, ok
}

func (p *RemoveNicFromVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *RemoveNicFromVirtualMachineParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *RemoveNicFromVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveNicFromVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRemoveNicFromVirtualMachineParams(nicid string, virtualmachineid string) *RemoveNicFromVirtualMachineParams {
	p := &RemoveNicFromVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["nicid"] = nicid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Removes VM from specified network by deleting a NIC
func (s *VirtualMachineService) RemoveNicFromVirtualMachine(p *RemoveNicFromVirtualMachineParams) (*RemoveNicFromVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("removeNicFromVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveNicFromVirtualMachineResponse
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

type RemoveNicFromVirtualMachineResponse struct {
	Account               string                                             `json:"account"`
	Affinitygroup         []RemoveNicFromVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                             `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                             `json:"autoscalevmgroupname"`
	Backupofferingid      string                                             `json:"backupofferingid"`
	Backupofferingname    string                                             `json:"backupofferingname"`
	Bootmode              string                                             `json:"bootmode"`
	Boottype              string                                             `json:"boottype"`
	Cpunumber             int                                                `json:"cpunumber"`
	Cpuspeed              int                                                `json:"cpuspeed"`
	Cpuused               string                                             `json:"cpuused"`
	Created               string                                             `json:"created"`
	Deleteprotection      bool                                               `json:"deleteprotection"`
	Details               map[string]string                                  `json:"details"`
	Diskioread            int64                                              `json:"diskioread"`
	Diskiowrite           int64                                              `json:"diskiowrite"`
	Diskkbsread           int64                                              `json:"diskkbsread"`
	Diskkbswrite          int64                                              `json:"diskkbswrite"`
	Diskofferingid        string                                             `json:"diskofferingid"`
	Diskofferingname      string                                             `json:"diskofferingname"`
	Displayname           string                                             `json:"displayname"`
	Displayvm             bool                                               `json:"displayvm"`
	Domain                string                                             `json:"domain"`
	Domainid              string                                             `json:"domainid"`
	Domainpath            string                                             `json:"domainpath"`
	Forvirtualnetwork     bool                                               `json:"forvirtualnetwork"`
	Group                 string                                             `json:"group"`
	Groupid               string                                             `json:"groupid"`
	Guestosid             string                                             `json:"guestosid"`
	Haenable              bool                                               `json:"haenable"`
	Hasannotations        bool                                               `json:"hasannotations"`
	Hostcontrolstate      string                                             `json:"hostcontrolstate"`
	Hostid                string                                             `json:"hostid"`
	Hostname              string                                             `json:"hostname"`
	Hypervisor            string                                             `json:"hypervisor"`
	Icon                  interface{}                                        `json:"icon"`
	Id                    string                                             `json:"id"`
	Instancename          string                                             `json:"instancename"`
	Ipaddress             string                                             `json:"ipaddress"`
	Isdynamicallyscalable bool                                               `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                             `json:"isodisplaytext"`
	Isoid                 string                                             `json:"isoid"`
	Isoname               string                                             `json:"isoname"`
	JobID                 string                                             `json:"jobid"`
	Jobstatus             int                                                `json:"jobstatus"`
	Keypairs              string                                             `json:"keypairs"`
	Lastupdated           string                                             `json:"lastupdated"`
	Memory                int                                                `json:"memory"`
	Memoryintfreekbs      int64                                              `json:"memoryintfreekbs"`
	Memorykbs             int64                                              `json:"memorykbs"`
	Memorytargetkbs       int64                                              `json:"memorytargetkbs"`
	Name                  string                                             `json:"name"`
	Networkkbsread        int64                                              `json:"networkkbsread"`
	Networkkbswrite       int64                                              `json:"networkkbswrite"`
	Nic                   []Nic                                              `json:"nic"`
	Osdisplayname         string                                             `json:"osdisplayname"`
	Ostypeid              string                                             `json:"ostypeid"`
	Password              string                                             `json:"password"`
	Passwordenabled       bool                                               `json:"passwordenabled"`
	Pooltype              string                                             `json:"pooltype"`
	Project               string                                             `json:"project"`
	Projectid             string                                             `json:"projectid"`
	Publicip              string                                             `json:"publicip"`
	Publicipid            string                                             `json:"publicipid"`
	Readonlydetails       string                                             `json:"readonlydetails"`
	Receivedbytes         int64                                              `json:"receivedbytes"`
	Rootdeviceid          int64                                              `json:"rootdeviceid"`
	Rootdevicetype        string                                             `json:"rootdevicetype"`
	Securitygroup         []RemoveNicFromVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                              `json:"sentbytes"`
	Serviceofferingid     string                                             `json:"serviceofferingid"`
	Serviceofferingname   string                                             `json:"serviceofferingname"`
	Servicestate          string                                             `json:"servicestate"`
	State                 string                                             `json:"state"`
	Tags                  []Tags                                             `json:"tags"`
	Templatedisplaytext   string                                             `json:"templatedisplaytext"`
	Templateformat        string                                             `json:"templateformat"`
	Templateid            string                                             `json:"templateid"`
	Templatename          string                                             `json:"templatename"`
	Templatetype          string                                             `json:"templatetype"`
	Userdata              string                                             `json:"userdata"`
	Userdatadetails       string                                             `json:"userdatadetails"`
	Userdataid            string                                             `json:"userdataid"`
	Userdataname          string                                             `json:"userdataname"`
	Userdatapolicy        string                                             `json:"userdatapolicy"`
	Userid                string                                             `json:"userid"`
	Username              string                                             `json:"username"`
	Vgpu                  string                                             `json:"vgpu"`
	Vmtype                string                                             `json:"vmtype"`
	Vnfdetails            map[string]string                                  `json:"vnfdetails"`
	Vnfnics               []string                                           `json:"vnfnics"`
	Zoneid                string                                             `json:"zoneid"`
	Zonename              string                                             `json:"zonename"`
}

type RemoveNicFromVirtualMachineResponseSecuritygroup struct {
	Account             string                                                 `json:"account"`
	Description         string                                                 `json:"description"`
	Domain              string                                                 `json:"domain"`
	Domainid            string                                                 `json:"domainid"`
	Domainpath          string                                                 `json:"domainpath"`
	Egressrule          []RemoveNicFromVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                 `json:"id"`
	Ingressrule         []RemoveNicFromVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                 `json:"name"`
	Project             string                                                 `json:"project"`
	Projectid           string                                                 `json:"projectid"`
	Tags                []Tags                                                 `json:"tags"`
	Virtualmachinecount int                                                    `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                          `json:"virtualmachineids"`
}

type RemoveNicFromVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type RemoveNicFromVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *RemoveNicFromVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias RemoveNicFromVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ResetPasswordForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ResetPasswordForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	return u
}

func (p *ResetPasswordForVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ResetPasswordForVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ResetPasswordForVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ResetPasswordForVirtualMachineParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *ResetPasswordForVirtualMachineParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *ResetPasswordForVirtualMachineParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

// You should always use this function to get a new ResetPasswordForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewResetPasswordForVirtualMachineParams(id string) *ResetPasswordForVirtualMachineParams {
	p := &ResetPasswordForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Resets the password for virtual machine. The virtual machine must be in a "Stopped" state and the template must already support this feature for this command to take effect. [async]
func (s *VirtualMachineService) ResetPasswordForVirtualMachine(p *ResetPasswordForVirtualMachineParams) (*ResetPasswordForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("resetPasswordForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetPasswordForVirtualMachineResponse
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

type ResetPasswordForVirtualMachineResponse struct {
	Account               string                                                `json:"account"`
	Affinitygroup         []ResetPasswordForVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                                `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                                `json:"autoscalevmgroupname"`
	Backupofferingid      string                                                `json:"backupofferingid"`
	Backupofferingname    string                                                `json:"backupofferingname"`
	Bootmode              string                                                `json:"bootmode"`
	Boottype              string                                                `json:"boottype"`
	Cpunumber             int                                                   `json:"cpunumber"`
	Cpuspeed              int                                                   `json:"cpuspeed"`
	Cpuused               string                                                `json:"cpuused"`
	Created               string                                                `json:"created"`
	Deleteprotection      bool                                                  `json:"deleteprotection"`
	Details               map[string]string                                     `json:"details"`
	Diskioread            int64                                                 `json:"diskioread"`
	Diskiowrite           int64                                                 `json:"diskiowrite"`
	Diskkbsread           int64                                                 `json:"diskkbsread"`
	Diskkbswrite          int64                                                 `json:"diskkbswrite"`
	Diskofferingid        string                                                `json:"diskofferingid"`
	Diskofferingname      string                                                `json:"diskofferingname"`
	Displayname           string                                                `json:"displayname"`
	Displayvm             bool                                                  `json:"displayvm"`
	Domain                string                                                `json:"domain"`
	Domainid              string                                                `json:"domainid"`
	Domainpath            string                                                `json:"domainpath"`
	Forvirtualnetwork     bool                                                  `json:"forvirtualnetwork"`
	Group                 string                                                `json:"group"`
	Groupid               string                                                `json:"groupid"`
	Guestosid             string                                                `json:"guestosid"`
	Haenable              bool                                                  `json:"haenable"`
	Hasannotations        bool                                                  `json:"hasannotations"`
	Hostcontrolstate      string                                                `json:"hostcontrolstate"`
	Hostid                string                                                `json:"hostid"`
	Hostname              string                                                `json:"hostname"`
	Hypervisor            string                                                `json:"hypervisor"`
	Icon                  interface{}                                           `json:"icon"`
	Id                    string                                                `json:"id"`
	Instancename          string                                                `json:"instancename"`
	Ipaddress             string                                                `json:"ipaddress"`
	Isdynamicallyscalable bool                                                  `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                `json:"isodisplaytext"`
	Isoid                 string                                                `json:"isoid"`
	Isoname               string                                                `json:"isoname"`
	JobID                 string                                                `json:"jobid"`
	Jobstatus             int                                                   `json:"jobstatus"`
	Keypairs              string                                                `json:"keypairs"`
	Lastupdated           string                                                `json:"lastupdated"`
	Memory                int                                                   `json:"memory"`
	Memoryintfreekbs      int64                                                 `json:"memoryintfreekbs"`
	Memorykbs             int64                                                 `json:"memorykbs"`
	Memorytargetkbs       int64                                                 `json:"memorytargetkbs"`
	Name                  string                                                `json:"name"`
	Networkkbsread        int64                                                 `json:"networkkbsread"`
	Networkkbswrite       int64                                                 `json:"networkkbswrite"`
	Nic                   []Nic                                                 `json:"nic"`
	Osdisplayname         string                                                `json:"osdisplayname"`
	Ostypeid              string                                                `json:"ostypeid"`
	Password              string                                                `json:"password"`
	Passwordenabled       bool                                                  `json:"passwordenabled"`
	Pooltype              string                                                `json:"pooltype"`
	Project               string                                                `json:"project"`
	Projectid             string                                                `json:"projectid"`
	Publicip              string                                                `json:"publicip"`
	Publicipid            string                                                `json:"publicipid"`
	Readonlydetails       string                                                `json:"readonlydetails"`
	Receivedbytes         int64                                                 `json:"receivedbytes"`
	Rootdeviceid          int64                                                 `json:"rootdeviceid"`
	Rootdevicetype        string                                                `json:"rootdevicetype"`
	Securitygroup         []ResetPasswordForVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                 `json:"sentbytes"`
	Serviceofferingid     string                                                `json:"serviceofferingid"`
	Serviceofferingname   string                                                `json:"serviceofferingname"`
	Servicestate          string                                                `json:"servicestate"`
	State                 string                                                `json:"state"`
	Tags                  []Tags                                                `json:"tags"`
	Templatedisplaytext   string                                                `json:"templatedisplaytext"`
	Templateformat        string                                                `json:"templateformat"`
	Templateid            string                                                `json:"templateid"`
	Templatename          string                                                `json:"templatename"`
	Templatetype          string                                                `json:"templatetype"`
	Userdata              string                                                `json:"userdata"`
	Userdatadetails       string                                                `json:"userdatadetails"`
	Userdataid            string                                                `json:"userdataid"`
	Userdataname          string                                                `json:"userdataname"`
	Userdatapolicy        string                                                `json:"userdatapolicy"`
	Userid                string                                                `json:"userid"`
	Username              string                                                `json:"username"`
	Vgpu                  string                                                `json:"vgpu"`
	Vmtype                string                                                `json:"vmtype"`
	Vnfdetails            map[string]string                                     `json:"vnfdetails"`
	Vnfnics               []string                                              `json:"vnfnics"`
	Zoneid                string                                                `json:"zoneid"`
	Zonename              string                                                `json:"zonename"`
}

type ResetPasswordForVirtualMachineResponseSecuritygroup struct {
	Account             string                                                    `json:"account"`
	Description         string                                                    `json:"description"`
	Domain              string                                                    `json:"domain"`
	Domainid            string                                                    `json:"domainid"`
	Domainpath          string                                                    `json:"domainpath"`
	Egressrule          []ResetPasswordForVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                    `json:"id"`
	Ingressrule         []ResetPasswordForVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                    `json:"name"`
	Project             string                                                    `json:"project"`
	Projectid           string                                                    `json:"projectid"`
	Tags                []Tags                                                    `json:"tags"`
	Virtualmachinecount int                                                       `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                             `json:"virtualmachineids"`
}

type ResetPasswordForVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type ResetPasswordForVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *ResetPasswordForVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias ResetPasswordForVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ResetUserDataForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ResetUserDataForVirtualMachineParams) toURLValues() url.Values {
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
	if v, found := p.p["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	if v, found := p.p["userdatadetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("userdatadetails[%d].key", i), k)
			u.Set(fmt.Sprintf("userdatadetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["userdataid"]; found {
		u.Set("userdataid", v.(string))
	}
	return u
}

func (p *ResetUserDataForVirtualMachineParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ResetUserDataForVirtualMachineParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ResetUserDataForVirtualMachineParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ResetUserDataForVirtualMachineParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ResetUserDataForVirtualMachineParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ResetUserDataForVirtualMachineParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ResetUserDataForVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ResetUserDataForVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ResetUserDataForVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ResetUserDataForVirtualMachineParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ResetUserDataForVirtualMachineParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ResetUserDataForVirtualMachineParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ResetUserDataForVirtualMachineParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *ResetUserDataForVirtualMachineParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *ResetUserDataForVirtualMachineParams) GetUserdata() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(string)
	return value, ok
}

func (p *ResetUserDataForVirtualMachineParams) SetUserdatadetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdatadetails"] = v
}

func (p *ResetUserDataForVirtualMachineParams) ResetUserdatadetails() {
	if p.p != nil && p.p["userdatadetails"] != nil {
		delete(p.p, "userdatadetails")
	}
}

func (p *ResetUserDataForVirtualMachineParams) GetUserdatadetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdatadetails"].(map[string]string)
	return value, ok
}

func (p *ResetUserDataForVirtualMachineParams) SetUserdataid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdataid"] = v
}

func (p *ResetUserDataForVirtualMachineParams) ResetUserdataid() {
	if p.p != nil && p.p["userdataid"] != nil {
		delete(p.p, "userdataid")
	}
}

func (p *ResetUserDataForVirtualMachineParams) GetUserdataid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdataid"].(string)
	return value, ok
}

// You should always use this function to get a new ResetUserDataForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewResetUserDataForVirtualMachineParams(id string) *ResetUserDataForVirtualMachineParams {
	p := &ResetUserDataForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Resets the UserData for virtual machine. The virtual machine must be in a "Stopped" state.
func (s *VirtualMachineService) ResetUserDataForVirtualMachine(p *ResetUserDataForVirtualMachineParams) (*ResetUserDataForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("resetUserDataForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetUserDataForVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ResetUserDataForVirtualMachineResponse struct {
	Account               string                                                `json:"account"`
	Affinitygroup         []ResetUserDataForVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                                `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                                `json:"autoscalevmgroupname"`
	Backupofferingid      string                                                `json:"backupofferingid"`
	Backupofferingname    string                                                `json:"backupofferingname"`
	Bootmode              string                                                `json:"bootmode"`
	Boottype              string                                                `json:"boottype"`
	Cpunumber             int                                                   `json:"cpunumber"`
	Cpuspeed              int                                                   `json:"cpuspeed"`
	Cpuused               string                                                `json:"cpuused"`
	Created               string                                                `json:"created"`
	Deleteprotection      bool                                                  `json:"deleteprotection"`
	Details               map[string]string                                     `json:"details"`
	Diskioread            int64                                                 `json:"diskioread"`
	Diskiowrite           int64                                                 `json:"diskiowrite"`
	Diskkbsread           int64                                                 `json:"diskkbsread"`
	Diskkbswrite          int64                                                 `json:"diskkbswrite"`
	Diskofferingid        string                                                `json:"diskofferingid"`
	Diskofferingname      string                                                `json:"diskofferingname"`
	Displayname           string                                                `json:"displayname"`
	Displayvm             bool                                                  `json:"displayvm"`
	Domain                string                                                `json:"domain"`
	Domainid              string                                                `json:"domainid"`
	Domainpath            string                                                `json:"domainpath"`
	Forvirtualnetwork     bool                                                  `json:"forvirtualnetwork"`
	Group                 string                                                `json:"group"`
	Groupid               string                                                `json:"groupid"`
	Guestosid             string                                                `json:"guestosid"`
	Haenable              bool                                                  `json:"haenable"`
	Hasannotations        bool                                                  `json:"hasannotations"`
	Hostcontrolstate      string                                                `json:"hostcontrolstate"`
	Hostid                string                                                `json:"hostid"`
	Hostname              string                                                `json:"hostname"`
	Hypervisor            string                                                `json:"hypervisor"`
	Icon                  interface{}                                           `json:"icon"`
	Id                    string                                                `json:"id"`
	Instancename          string                                                `json:"instancename"`
	Ipaddress             string                                                `json:"ipaddress"`
	Isdynamicallyscalable bool                                                  `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                `json:"isodisplaytext"`
	Isoid                 string                                                `json:"isoid"`
	Isoname               string                                                `json:"isoname"`
	JobID                 string                                                `json:"jobid"`
	Jobstatus             int                                                   `json:"jobstatus"`
	Keypairs              string                                                `json:"keypairs"`
	Lastupdated           string                                                `json:"lastupdated"`
	Memory                int                                                   `json:"memory"`
	Memoryintfreekbs      int64                                                 `json:"memoryintfreekbs"`
	Memorykbs             int64                                                 `json:"memorykbs"`
	Memorytargetkbs       int64                                                 `json:"memorytargetkbs"`
	Name                  string                                                `json:"name"`
	Networkkbsread        int64                                                 `json:"networkkbsread"`
	Networkkbswrite       int64                                                 `json:"networkkbswrite"`
	Nic                   []Nic                                                 `json:"nic"`
	Osdisplayname         string                                                `json:"osdisplayname"`
	Ostypeid              string                                                `json:"ostypeid"`
	Password              string                                                `json:"password"`
	Passwordenabled       bool                                                  `json:"passwordenabled"`
	Pooltype              string                                                `json:"pooltype"`
	Project               string                                                `json:"project"`
	Projectid             string                                                `json:"projectid"`
	Publicip              string                                                `json:"publicip"`
	Publicipid            string                                                `json:"publicipid"`
	Readonlydetails       string                                                `json:"readonlydetails"`
	Receivedbytes         int64                                                 `json:"receivedbytes"`
	Rootdeviceid          int64                                                 `json:"rootdeviceid"`
	Rootdevicetype        string                                                `json:"rootdevicetype"`
	Securitygroup         []ResetUserDataForVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                 `json:"sentbytes"`
	Serviceofferingid     string                                                `json:"serviceofferingid"`
	Serviceofferingname   string                                                `json:"serviceofferingname"`
	Servicestate          string                                                `json:"servicestate"`
	State                 string                                                `json:"state"`
	Tags                  []Tags                                                `json:"tags"`
	Templatedisplaytext   string                                                `json:"templatedisplaytext"`
	Templateformat        string                                                `json:"templateformat"`
	Templateid            string                                                `json:"templateid"`
	Templatename          string                                                `json:"templatename"`
	Templatetype          string                                                `json:"templatetype"`
	Userdata              string                                                `json:"userdata"`
	Userdatadetails       string                                                `json:"userdatadetails"`
	Userdataid            string                                                `json:"userdataid"`
	Userdataname          string                                                `json:"userdataname"`
	Userdatapolicy        string                                                `json:"userdatapolicy"`
	Userid                string                                                `json:"userid"`
	Username              string                                                `json:"username"`
	Vgpu                  string                                                `json:"vgpu"`
	Vmtype                string                                                `json:"vmtype"`
	Vnfdetails            map[string]string                                     `json:"vnfdetails"`
	Vnfnics               []string                                              `json:"vnfnics"`
	Zoneid                string                                                `json:"zoneid"`
	Zonename              string                                                `json:"zonename"`
}

type ResetUserDataForVirtualMachineResponseSecuritygroup struct {
	Account             string                                                    `json:"account"`
	Description         string                                                    `json:"description"`
	Domain              string                                                    `json:"domain"`
	Domainid            string                                                    `json:"domainid"`
	Domainpath          string                                                    `json:"domainpath"`
	Egressrule          []ResetUserDataForVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                    `json:"id"`
	Ingressrule         []ResetUserDataForVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                    `json:"name"`
	Project             string                                                    `json:"project"`
	Projectid           string                                                    `json:"projectid"`
	Tags                []Tags                                                    `json:"tags"`
	Virtualmachinecount int                                                       `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                             `json:"virtualmachineids"`
}

type ResetUserDataForVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type ResetUserDataForVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *ResetUserDataForVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias ResetUserDataForVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RestoreVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RestoreVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["expunge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("expunge", vv)
	}
	if v, found := p.p["rootdisksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("rootdisksize", vv)
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *RestoreVirtualMachineParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *RestoreVirtualMachineParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *RestoreVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *RestoreVirtualMachineParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *RestoreVirtualMachineParams) ResetDiskofferingid() {
	if p.p != nil && p.p["diskofferingid"] != nil {
		delete(p.p, "diskofferingid")
	}
}

func (p *RestoreVirtualMachineParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *RestoreVirtualMachineParams) SetExpunge(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expunge"] = v
}

func (p *RestoreVirtualMachineParams) ResetExpunge() {
	if p.p != nil && p.p["expunge"] != nil {
		delete(p.p, "expunge")
	}
}

func (p *RestoreVirtualMachineParams) GetExpunge() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["expunge"].(bool)
	return value, ok
}

func (p *RestoreVirtualMachineParams) SetRootdisksize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rootdisksize"] = v
}

func (p *RestoreVirtualMachineParams) ResetRootdisksize() {
	if p.p != nil && p.p["rootdisksize"] != nil {
		delete(p.p, "rootdisksize")
	}
}

func (p *RestoreVirtualMachineParams) GetRootdisksize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rootdisksize"].(int64)
	return value, ok
}

func (p *RestoreVirtualMachineParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *RestoreVirtualMachineParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *RestoreVirtualMachineParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *RestoreVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *RestoreVirtualMachineParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *RestoreVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new RestoreVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRestoreVirtualMachineParams(virtualmachineid string) *RestoreVirtualMachineParams {
	p := &RestoreVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Restore a VM to original template/ISO or new template/ISO
func (s *VirtualMachineService) RestoreVirtualMachine(p *RestoreVirtualMachineParams) (*RestoreVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("restoreVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestoreVirtualMachineResponse
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

type RestoreVirtualMachineResponse struct {
	Account               string                                       `json:"account"`
	Affinitygroup         []RestoreVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                       `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                       `json:"autoscalevmgroupname"`
	Backupofferingid      string                                       `json:"backupofferingid"`
	Backupofferingname    string                                       `json:"backupofferingname"`
	Bootmode              string                                       `json:"bootmode"`
	Boottype              string                                       `json:"boottype"`
	Cpunumber             int                                          `json:"cpunumber"`
	Cpuspeed              int                                          `json:"cpuspeed"`
	Cpuused               string                                       `json:"cpuused"`
	Created               string                                       `json:"created"`
	Deleteprotection      bool                                         `json:"deleteprotection"`
	Details               map[string]string                            `json:"details"`
	Diskioread            int64                                        `json:"diskioread"`
	Diskiowrite           int64                                        `json:"diskiowrite"`
	Diskkbsread           int64                                        `json:"diskkbsread"`
	Diskkbswrite          int64                                        `json:"diskkbswrite"`
	Diskofferingid        string                                       `json:"diskofferingid"`
	Diskofferingname      string                                       `json:"diskofferingname"`
	Displayname           string                                       `json:"displayname"`
	Displayvm             bool                                         `json:"displayvm"`
	Domain                string                                       `json:"domain"`
	Domainid              string                                       `json:"domainid"`
	Domainpath            string                                       `json:"domainpath"`
	Forvirtualnetwork     bool                                         `json:"forvirtualnetwork"`
	Group                 string                                       `json:"group"`
	Groupid               string                                       `json:"groupid"`
	Guestosid             string                                       `json:"guestosid"`
	Haenable              bool                                         `json:"haenable"`
	Hasannotations        bool                                         `json:"hasannotations"`
	Hostcontrolstate      string                                       `json:"hostcontrolstate"`
	Hostid                string                                       `json:"hostid"`
	Hostname              string                                       `json:"hostname"`
	Hypervisor            string                                       `json:"hypervisor"`
	Icon                  interface{}                                  `json:"icon"`
	Id                    string                                       `json:"id"`
	Instancename          string                                       `json:"instancename"`
	Ipaddress             string                                       `json:"ipaddress"`
	Isdynamicallyscalable bool                                         `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                       `json:"isodisplaytext"`
	Isoid                 string                                       `json:"isoid"`
	Isoname               string                                       `json:"isoname"`
	JobID                 string                                       `json:"jobid"`
	Jobstatus             int                                          `json:"jobstatus"`
	Keypairs              string                                       `json:"keypairs"`
	Lastupdated           string                                       `json:"lastupdated"`
	Memory                int                                          `json:"memory"`
	Memoryintfreekbs      int64                                        `json:"memoryintfreekbs"`
	Memorykbs             int64                                        `json:"memorykbs"`
	Memorytargetkbs       int64                                        `json:"memorytargetkbs"`
	Name                  string                                       `json:"name"`
	Networkkbsread        int64                                        `json:"networkkbsread"`
	Networkkbswrite       int64                                        `json:"networkkbswrite"`
	Nic                   []Nic                                        `json:"nic"`
	Osdisplayname         string                                       `json:"osdisplayname"`
	Ostypeid              string                                       `json:"ostypeid"`
	Password              string                                       `json:"password"`
	Passwordenabled       bool                                         `json:"passwordenabled"`
	Pooltype              string                                       `json:"pooltype"`
	Project               string                                       `json:"project"`
	Projectid             string                                       `json:"projectid"`
	Publicip              string                                       `json:"publicip"`
	Publicipid            string                                       `json:"publicipid"`
	Readonlydetails       string                                       `json:"readonlydetails"`
	Receivedbytes         int64                                        `json:"receivedbytes"`
	Rootdeviceid          int64                                        `json:"rootdeviceid"`
	Rootdevicetype        string                                       `json:"rootdevicetype"`
	Securitygroup         []RestoreVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                        `json:"sentbytes"`
	Serviceofferingid     string                                       `json:"serviceofferingid"`
	Serviceofferingname   string                                       `json:"serviceofferingname"`
	Servicestate          string                                       `json:"servicestate"`
	State                 string                                       `json:"state"`
	Tags                  []Tags                                       `json:"tags"`
	Templatedisplaytext   string                                       `json:"templatedisplaytext"`
	Templateformat        string                                       `json:"templateformat"`
	Templateid            string                                       `json:"templateid"`
	Templatename          string                                       `json:"templatename"`
	Templatetype          string                                       `json:"templatetype"`
	Userdata              string                                       `json:"userdata"`
	Userdatadetails       string                                       `json:"userdatadetails"`
	Userdataid            string                                       `json:"userdataid"`
	Userdataname          string                                       `json:"userdataname"`
	Userdatapolicy        string                                       `json:"userdatapolicy"`
	Userid                string                                       `json:"userid"`
	Username              string                                       `json:"username"`
	Vgpu                  string                                       `json:"vgpu"`
	Vmtype                string                                       `json:"vmtype"`
	Vnfdetails            map[string]string                            `json:"vnfdetails"`
	Vnfnics               []string                                     `json:"vnfnics"`
	Zoneid                string                                       `json:"zoneid"`
	Zonename              string                                       `json:"zonename"`
}

type RestoreVirtualMachineResponseSecuritygroup struct {
	Account             string                                           `json:"account"`
	Description         string                                           `json:"description"`
	Domain              string                                           `json:"domain"`
	Domainid            string                                           `json:"domainid"`
	Domainpath          string                                           `json:"domainpath"`
	Egressrule          []RestoreVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                           `json:"id"`
	Ingressrule         []RestoreVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	Projectid           string                                           `json:"projectid"`
	Tags                []Tags                                           `json:"tags"`
	Virtualmachinecount int                                              `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                    `json:"virtualmachineids"`
}

type RestoreVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type RestoreVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *RestoreVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias RestoreVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ScaleVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ScaleVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["automigrate"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("automigrate", vv)
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
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["shrinkok"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("shrinkok", vv)
	}
	return u
}

func (p *ScaleVirtualMachineParams) SetAutomigrate(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["automigrate"] = v
}

func (p *ScaleVirtualMachineParams) ResetAutomigrate() {
	if p.p != nil && p.p["automigrate"] != nil {
		delete(p.p, "automigrate")
	}
}

func (p *ScaleVirtualMachineParams) GetAutomigrate() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["automigrate"].(bool)
	return value, ok
}

func (p *ScaleVirtualMachineParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ScaleVirtualMachineParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ScaleVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *ScaleVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ScaleVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ScaleVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ScaleVirtualMachineParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *ScaleVirtualMachineParams) ResetMaxiops() {
	if p.p != nil && p.p["maxiops"] != nil {
		delete(p.p, "maxiops")
	}
}

func (p *ScaleVirtualMachineParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *ScaleVirtualMachineParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *ScaleVirtualMachineParams) ResetMiniops() {
	if p.p != nil && p.p["miniops"] != nil {
		delete(p.p, "miniops")
	}
}

func (p *ScaleVirtualMachineParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *ScaleVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ScaleVirtualMachineParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ScaleVirtualMachineParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ScaleVirtualMachineParams) SetShrinkok(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["shrinkok"] = v
}

func (p *ScaleVirtualMachineParams) ResetShrinkok() {
	if p.p != nil && p.p["shrinkok"] != nil {
		delete(p.p, "shrinkok")
	}
}

func (p *ScaleVirtualMachineParams) GetShrinkok() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["shrinkok"].(bool)
	return value, ok
}

// You should always use this function to get a new ScaleVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewScaleVirtualMachineParams(id string, serviceofferingid string) *ScaleVirtualMachineParams {
	p := &ScaleVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Scales the virtual machine to a new service offering. This command also considers the volume size in the service offering or disk offering linked to the new service offering and apply all characteristics to the root volume.
func (s *VirtualMachineService) ScaleVirtualMachine(p *ScaleVirtualMachineParams) (*ScaleVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("scaleVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleVirtualMachineResponse
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

type ScaleVirtualMachineResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type StartVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *StartVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bootintosetup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootintosetup", vv)
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["considerlasthost"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("considerlasthost", vv)
	}
	if v, found := p.p["deploymentplanner"]; found {
		u.Set("deploymentplanner", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (p *StartVirtualMachineParams) SetBootintosetup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootintosetup"] = v
}

func (p *StartVirtualMachineParams) ResetBootintosetup() {
	if p.p != nil && p.p["bootintosetup"] != nil {
		delete(p.p, "bootintosetup")
	}
}

func (p *StartVirtualMachineParams) GetBootintosetup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootintosetup"].(bool)
	return value, ok
}

func (p *StartVirtualMachineParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *StartVirtualMachineParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *StartVirtualMachineParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *StartVirtualMachineParams) SetConsiderlasthost(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["considerlasthost"] = v
}

func (p *StartVirtualMachineParams) ResetConsiderlasthost() {
	if p.p != nil && p.p["considerlasthost"] != nil {
		delete(p.p, "considerlasthost")
	}
}

func (p *StartVirtualMachineParams) GetConsiderlasthost() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["considerlasthost"].(bool)
	return value, ok
}

func (p *StartVirtualMachineParams) SetDeploymentplanner(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deploymentplanner"] = v
}

func (p *StartVirtualMachineParams) ResetDeploymentplanner() {
	if p.p != nil && p.p["deploymentplanner"] != nil {
		delete(p.p, "deploymentplanner")
	}
}

func (p *StartVirtualMachineParams) GetDeploymentplanner() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deploymentplanner"].(string)
	return value, ok
}

func (p *StartVirtualMachineParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *StartVirtualMachineParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *StartVirtualMachineParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *StartVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StartVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StartVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *StartVirtualMachineParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *StartVirtualMachineParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *StartVirtualMachineParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

// You should always use this function to get a new StartVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewStartVirtualMachineParams(id string) *StartVirtualMachineParams {
	p := &StartVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a virtual machine.
func (s *VirtualMachineService) StartVirtualMachine(p *StartVirtualMachineParams) (*StartVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("startVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartVirtualMachineResponse
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

type StartVirtualMachineResponse struct {
	Account               string                                     `json:"account"`
	Affinitygroup         []StartVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                     `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                     `json:"autoscalevmgroupname"`
	Backupofferingid      string                                     `json:"backupofferingid"`
	Backupofferingname    string                                     `json:"backupofferingname"`
	Bootmode              string                                     `json:"bootmode"`
	Boottype              string                                     `json:"boottype"`
	Cpunumber             int                                        `json:"cpunumber"`
	Cpuspeed              int                                        `json:"cpuspeed"`
	Cpuused               string                                     `json:"cpuused"`
	Created               string                                     `json:"created"`
	Deleteprotection      bool                                       `json:"deleteprotection"`
	Details               map[string]string                          `json:"details"`
	Diskioread            int64                                      `json:"diskioread"`
	Diskiowrite           int64                                      `json:"diskiowrite"`
	Diskkbsread           int64                                      `json:"diskkbsread"`
	Diskkbswrite          int64                                      `json:"diskkbswrite"`
	Diskofferingid        string                                     `json:"diskofferingid"`
	Diskofferingname      string                                     `json:"diskofferingname"`
	Displayname           string                                     `json:"displayname"`
	Displayvm             bool                                       `json:"displayvm"`
	Domain                string                                     `json:"domain"`
	Domainid              string                                     `json:"domainid"`
	Domainpath            string                                     `json:"domainpath"`
	Forvirtualnetwork     bool                                       `json:"forvirtualnetwork"`
	Group                 string                                     `json:"group"`
	Groupid               string                                     `json:"groupid"`
	Guestosid             string                                     `json:"guestosid"`
	Haenable              bool                                       `json:"haenable"`
	Hasannotations        bool                                       `json:"hasannotations"`
	Hostcontrolstate      string                                     `json:"hostcontrolstate"`
	Hostid                string                                     `json:"hostid"`
	Hostname              string                                     `json:"hostname"`
	Hypervisor            string                                     `json:"hypervisor"`
	Icon                  interface{}                                `json:"icon"`
	Id                    string                                     `json:"id"`
	Instancename          string                                     `json:"instancename"`
	Ipaddress             string                                     `json:"ipaddress"`
	Isdynamicallyscalable bool                                       `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                     `json:"isodisplaytext"`
	Isoid                 string                                     `json:"isoid"`
	Isoname               string                                     `json:"isoname"`
	JobID                 string                                     `json:"jobid"`
	Jobstatus             int                                        `json:"jobstatus"`
	Keypairs              string                                     `json:"keypairs"`
	Lastupdated           string                                     `json:"lastupdated"`
	Memory                int                                        `json:"memory"`
	Memoryintfreekbs      int64                                      `json:"memoryintfreekbs"`
	Memorykbs             int64                                      `json:"memorykbs"`
	Memorytargetkbs       int64                                      `json:"memorytargetkbs"`
	Name                  string                                     `json:"name"`
	Networkkbsread        int64                                      `json:"networkkbsread"`
	Networkkbswrite       int64                                      `json:"networkkbswrite"`
	Nic                   []Nic                                      `json:"nic"`
	Osdisplayname         string                                     `json:"osdisplayname"`
	Ostypeid              string                                     `json:"ostypeid"`
	Password              string                                     `json:"password"`
	Passwordenabled       bool                                       `json:"passwordenabled"`
	Pooltype              string                                     `json:"pooltype"`
	Project               string                                     `json:"project"`
	Projectid             string                                     `json:"projectid"`
	Publicip              string                                     `json:"publicip"`
	Publicipid            string                                     `json:"publicipid"`
	Readonlydetails       string                                     `json:"readonlydetails"`
	Receivedbytes         int64                                      `json:"receivedbytes"`
	Rootdeviceid          int64                                      `json:"rootdeviceid"`
	Rootdevicetype        string                                     `json:"rootdevicetype"`
	Securitygroup         []StartVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                      `json:"sentbytes"`
	Serviceofferingid     string                                     `json:"serviceofferingid"`
	Serviceofferingname   string                                     `json:"serviceofferingname"`
	Servicestate          string                                     `json:"servicestate"`
	State                 string                                     `json:"state"`
	Tags                  []Tags                                     `json:"tags"`
	Templatedisplaytext   string                                     `json:"templatedisplaytext"`
	Templateformat        string                                     `json:"templateformat"`
	Templateid            string                                     `json:"templateid"`
	Templatename          string                                     `json:"templatename"`
	Templatetype          string                                     `json:"templatetype"`
	Userdata              string                                     `json:"userdata"`
	Userdatadetails       string                                     `json:"userdatadetails"`
	Userdataid            string                                     `json:"userdataid"`
	Userdataname          string                                     `json:"userdataname"`
	Userdatapolicy        string                                     `json:"userdatapolicy"`
	Userid                string                                     `json:"userid"`
	Username              string                                     `json:"username"`
	Vgpu                  string                                     `json:"vgpu"`
	Vmtype                string                                     `json:"vmtype"`
	Vnfdetails            map[string]string                          `json:"vnfdetails"`
	Vnfnics               []string                                   `json:"vnfnics"`
	Zoneid                string                                     `json:"zoneid"`
	Zonename              string                                     `json:"zonename"`
}

type StartVirtualMachineResponseSecuritygroup struct {
	Account             string                                         `json:"account"`
	Description         string                                         `json:"description"`
	Domain              string                                         `json:"domain"`
	Domainid            string                                         `json:"domainid"`
	Domainpath          string                                         `json:"domainpath"`
	Egressrule          []StartVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                         `json:"id"`
	Ingressrule         []StartVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                         `json:"name"`
	Project             string                                         `json:"project"`
	Projectid           string                                         `json:"projectid"`
	Tags                []Tags                                         `json:"tags"`
	Virtualmachinecount int                                            `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                  `json:"virtualmachineids"`
}

type StartVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type StartVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *StartVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias StartVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type StopVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *StopVirtualMachineParams) toURLValues() url.Values {
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

func (p *StopVirtualMachineParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *StopVirtualMachineParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *StopVirtualMachineParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *StopVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *StopVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *StopVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewStopVirtualMachineParams(id string) *StopVirtualMachineParams {
	p := &StopVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a virtual machine.
func (s *VirtualMachineService) StopVirtualMachine(p *StopVirtualMachineParams) (*StopVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("stopVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopVirtualMachineResponse
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

type StopVirtualMachineResponse struct {
	Account               string                                    `json:"account"`
	Affinitygroup         []StopVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                    `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                    `json:"autoscalevmgroupname"`
	Backupofferingid      string                                    `json:"backupofferingid"`
	Backupofferingname    string                                    `json:"backupofferingname"`
	Bootmode              string                                    `json:"bootmode"`
	Boottype              string                                    `json:"boottype"`
	Cpunumber             int                                       `json:"cpunumber"`
	Cpuspeed              int                                       `json:"cpuspeed"`
	Cpuused               string                                    `json:"cpuused"`
	Created               string                                    `json:"created"`
	Deleteprotection      bool                                      `json:"deleteprotection"`
	Details               map[string]string                         `json:"details"`
	Diskioread            int64                                     `json:"diskioread"`
	Diskiowrite           int64                                     `json:"diskiowrite"`
	Diskkbsread           int64                                     `json:"diskkbsread"`
	Diskkbswrite          int64                                     `json:"diskkbswrite"`
	Diskofferingid        string                                    `json:"diskofferingid"`
	Diskofferingname      string                                    `json:"diskofferingname"`
	Displayname           string                                    `json:"displayname"`
	Displayvm             bool                                      `json:"displayvm"`
	Domain                string                                    `json:"domain"`
	Domainid              string                                    `json:"domainid"`
	Domainpath            string                                    `json:"domainpath"`
	Forvirtualnetwork     bool                                      `json:"forvirtualnetwork"`
	Group                 string                                    `json:"group"`
	Groupid               string                                    `json:"groupid"`
	Guestosid             string                                    `json:"guestosid"`
	Haenable              bool                                      `json:"haenable"`
	Hasannotations        bool                                      `json:"hasannotations"`
	Hostcontrolstate      string                                    `json:"hostcontrolstate"`
	Hostid                string                                    `json:"hostid"`
	Hostname              string                                    `json:"hostname"`
	Hypervisor            string                                    `json:"hypervisor"`
	Icon                  interface{}                               `json:"icon"`
	Id                    string                                    `json:"id"`
	Instancename          string                                    `json:"instancename"`
	Ipaddress             string                                    `json:"ipaddress"`
	Isdynamicallyscalable bool                                      `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                    `json:"isodisplaytext"`
	Isoid                 string                                    `json:"isoid"`
	Isoname               string                                    `json:"isoname"`
	JobID                 string                                    `json:"jobid"`
	Jobstatus             int                                       `json:"jobstatus"`
	Keypairs              string                                    `json:"keypairs"`
	Lastupdated           string                                    `json:"lastupdated"`
	Memory                int                                       `json:"memory"`
	Memoryintfreekbs      int64                                     `json:"memoryintfreekbs"`
	Memorykbs             int64                                     `json:"memorykbs"`
	Memorytargetkbs       int64                                     `json:"memorytargetkbs"`
	Name                  string                                    `json:"name"`
	Networkkbsread        int64                                     `json:"networkkbsread"`
	Networkkbswrite       int64                                     `json:"networkkbswrite"`
	Nic                   []Nic                                     `json:"nic"`
	Osdisplayname         string                                    `json:"osdisplayname"`
	Ostypeid              string                                    `json:"ostypeid"`
	Password              string                                    `json:"password"`
	Passwordenabled       bool                                      `json:"passwordenabled"`
	Pooltype              string                                    `json:"pooltype"`
	Project               string                                    `json:"project"`
	Projectid             string                                    `json:"projectid"`
	Publicip              string                                    `json:"publicip"`
	Publicipid            string                                    `json:"publicipid"`
	Readonlydetails       string                                    `json:"readonlydetails"`
	Receivedbytes         int64                                     `json:"receivedbytes"`
	Rootdeviceid          int64                                     `json:"rootdeviceid"`
	Rootdevicetype        string                                    `json:"rootdevicetype"`
	Securitygroup         []StopVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                     `json:"sentbytes"`
	Serviceofferingid     string                                    `json:"serviceofferingid"`
	Serviceofferingname   string                                    `json:"serviceofferingname"`
	Servicestate          string                                    `json:"servicestate"`
	State                 string                                    `json:"state"`
	Tags                  []Tags                                    `json:"tags"`
	Templatedisplaytext   string                                    `json:"templatedisplaytext"`
	Templateformat        string                                    `json:"templateformat"`
	Templateid            string                                    `json:"templateid"`
	Templatename          string                                    `json:"templatename"`
	Templatetype          string                                    `json:"templatetype"`
	Userdata              string                                    `json:"userdata"`
	Userdatadetails       string                                    `json:"userdatadetails"`
	Userdataid            string                                    `json:"userdataid"`
	Userdataname          string                                    `json:"userdataname"`
	Userdatapolicy        string                                    `json:"userdatapolicy"`
	Userid                string                                    `json:"userid"`
	Username              string                                    `json:"username"`
	Vgpu                  string                                    `json:"vgpu"`
	Vmtype                string                                    `json:"vmtype"`
	Vnfdetails            map[string]string                         `json:"vnfdetails"`
	Vnfnics               []string                                  `json:"vnfnics"`
	Zoneid                string                                    `json:"zoneid"`
	Zonename              string                                    `json:"zonename"`
}

type StopVirtualMachineResponseSecuritygroup struct {
	Account             string                                        `json:"account"`
	Description         string                                        `json:"description"`
	Domain              string                                        `json:"domain"`
	Domainid            string                                        `json:"domainid"`
	Domainpath          string                                        `json:"domainpath"`
	Egressrule          []StopVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                        `json:"id"`
	Ingressrule         []StopVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                        `json:"name"`
	Project             string                                        `json:"project"`
	Projectid           string                                        `json:"projectid"`
	Tags                []Tags                                        `json:"tags"`
	Virtualmachinecount int                                           `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                 `json:"virtualmachineids"`
}

type StopVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type StopVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *StopVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias StopVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateDefaultNicForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *UpdateDefaultNicForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *UpdateDefaultNicForVirtualMachineParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
}

func (p *UpdateDefaultNicForVirtualMachineParams) ResetNicid() {
	if p.p != nil && p.p["nicid"] != nil {
		delete(p.p, "nicid")
	}
}

func (p *UpdateDefaultNicForVirtualMachineParams) GetNicid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicid"].(string)
	return value, ok
}

func (p *UpdateDefaultNicForVirtualMachineParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *UpdateDefaultNicForVirtualMachineParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *UpdateDefaultNicForVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateDefaultNicForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewUpdateDefaultNicForVirtualMachineParams(nicid string, virtualmachineid string) *UpdateDefaultNicForVirtualMachineParams {
	p := &UpdateDefaultNicForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["nicid"] = nicid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Changes the default NIC on a VM
func (s *VirtualMachineService) UpdateDefaultNicForVirtualMachine(p *UpdateDefaultNicForVirtualMachineParams) (*UpdateDefaultNicForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("updateDefaultNicForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateDefaultNicForVirtualMachineResponse
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

type UpdateDefaultNicForVirtualMachineResponse struct {
	Account               string                                                   `json:"account"`
	Affinitygroup         []UpdateDefaultNicForVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                                   `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                                   `json:"autoscalevmgroupname"`
	Backupofferingid      string                                                   `json:"backupofferingid"`
	Backupofferingname    string                                                   `json:"backupofferingname"`
	Bootmode              string                                                   `json:"bootmode"`
	Boottype              string                                                   `json:"boottype"`
	Cpunumber             int                                                      `json:"cpunumber"`
	Cpuspeed              int                                                      `json:"cpuspeed"`
	Cpuused               string                                                   `json:"cpuused"`
	Created               string                                                   `json:"created"`
	Deleteprotection      bool                                                     `json:"deleteprotection"`
	Details               map[string]string                                        `json:"details"`
	Diskioread            int64                                                    `json:"diskioread"`
	Diskiowrite           int64                                                    `json:"diskiowrite"`
	Diskkbsread           int64                                                    `json:"diskkbsread"`
	Diskkbswrite          int64                                                    `json:"diskkbswrite"`
	Diskofferingid        string                                                   `json:"diskofferingid"`
	Diskofferingname      string                                                   `json:"diskofferingname"`
	Displayname           string                                                   `json:"displayname"`
	Displayvm             bool                                                     `json:"displayvm"`
	Domain                string                                                   `json:"domain"`
	Domainid              string                                                   `json:"domainid"`
	Domainpath            string                                                   `json:"domainpath"`
	Forvirtualnetwork     bool                                                     `json:"forvirtualnetwork"`
	Group                 string                                                   `json:"group"`
	Groupid               string                                                   `json:"groupid"`
	Guestosid             string                                                   `json:"guestosid"`
	Haenable              bool                                                     `json:"haenable"`
	Hasannotations        bool                                                     `json:"hasannotations"`
	Hostcontrolstate      string                                                   `json:"hostcontrolstate"`
	Hostid                string                                                   `json:"hostid"`
	Hostname              string                                                   `json:"hostname"`
	Hypervisor            string                                                   `json:"hypervisor"`
	Icon                  interface{}                                              `json:"icon"`
	Id                    string                                                   `json:"id"`
	Instancename          string                                                   `json:"instancename"`
	Ipaddress             string                                                   `json:"ipaddress"`
	Isdynamicallyscalable bool                                                     `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                   `json:"isodisplaytext"`
	Isoid                 string                                                   `json:"isoid"`
	Isoname               string                                                   `json:"isoname"`
	JobID                 string                                                   `json:"jobid"`
	Jobstatus             int                                                      `json:"jobstatus"`
	Keypairs              string                                                   `json:"keypairs"`
	Lastupdated           string                                                   `json:"lastupdated"`
	Memory                int                                                      `json:"memory"`
	Memoryintfreekbs      int64                                                    `json:"memoryintfreekbs"`
	Memorykbs             int64                                                    `json:"memorykbs"`
	Memorytargetkbs       int64                                                    `json:"memorytargetkbs"`
	Name                  string                                                   `json:"name"`
	Networkkbsread        int64                                                    `json:"networkkbsread"`
	Networkkbswrite       int64                                                    `json:"networkkbswrite"`
	Nic                   []Nic                                                    `json:"nic"`
	Osdisplayname         string                                                   `json:"osdisplayname"`
	Ostypeid              string                                                   `json:"ostypeid"`
	Password              string                                                   `json:"password"`
	Passwordenabled       bool                                                     `json:"passwordenabled"`
	Pooltype              string                                                   `json:"pooltype"`
	Project               string                                                   `json:"project"`
	Projectid             string                                                   `json:"projectid"`
	Publicip              string                                                   `json:"publicip"`
	Publicipid            string                                                   `json:"publicipid"`
	Readonlydetails       string                                                   `json:"readonlydetails"`
	Receivedbytes         int64                                                    `json:"receivedbytes"`
	Rootdeviceid          int64                                                    `json:"rootdeviceid"`
	Rootdevicetype        string                                                   `json:"rootdevicetype"`
	Securitygroup         []UpdateDefaultNicForVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                    `json:"sentbytes"`
	Serviceofferingid     string                                                   `json:"serviceofferingid"`
	Serviceofferingname   string                                                   `json:"serviceofferingname"`
	Servicestate          string                                                   `json:"servicestate"`
	State                 string                                                   `json:"state"`
	Tags                  []Tags                                                   `json:"tags"`
	Templatedisplaytext   string                                                   `json:"templatedisplaytext"`
	Templateformat        string                                                   `json:"templateformat"`
	Templateid            string                                                   `json:"templateid"`
	Templatename          string                                                   `json:"templatename"`
	Templatetype          string                                                   `json:"templatetype"`
	Userdata              string                                                   `json:"userdata"`
	Userdatadetails       string                                                   `json:"userdatadetails"`
	Userdataid            string                                                   `json:"userdataid"`
	Userdataname          string                                                   `json:"userdataname"`
	Userdatapolicy        string                                                   `json:"userdatapolicy"`
	Userid                string                                                   `json:"userid"`
	Username              string                                                   `json:"username"`
	Vgpu                  string                                                   `json:"vgpu"`
	Vmtype                string                                                   `json:"vmtype"`
	Vnfdetails            map[string]string                                        `json:"vnfdetails"`
	Vnfnics               []string                                                 `json:"vnfnics"`
	Zoneid                string                                                   `json:"zoneid"`
	Zonename              string                                                   `json:"zonename"`
}

type UpdateDefaultNicForVirtualMachineResponseSecuritygroup struct {
	Account             string                                                       `json:"account"`
	Description         string                                                       `json:"description"`
	Domain              string                                                       `json:"domain"`
	Domainid            string                                                       `json:"domainid"`
	Domainpath          string                                                       `json:"domainpath"`
	Egressrule          []UpdateDefaultNicForVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                       `json:"id"`
	Ingressrule         []UpdateDefaultNicForVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                       `json:"name"`
	Project             string                                                       `json:"project"`
	Projectid           string                                                       `json:"projectid"`
	Tags                []Tags                                                       `json:"tags"`
	Virtualmachinecount int                                                          `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                                `json:"virtualmachineids"`
}

type UpdateDefaultNicForVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type UpdateDefaultNicForVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *UpdateDefaultNicForVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateDefaultNicForVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *UpdateVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanupdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupdetails", vv)
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["deleteprotection"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("deleteprotection", vv)
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["dhcpoptionsnetworklist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("dhcpoptionsnetworklist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := p.p["displayname"]; found {
		u.Set("displayname", v.(string))
	}
	if v, found := p.p["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := p.p["extraconfig"]; found {
		u.Set("extraconfig", v.(string))
	}
	if v, found := p.p["group"]; found {
		u.Set("group", v.(string))
	}
	if v, found := p.p["haenable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("haenable", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["instancename"]; found {
		u.Set("instancename", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["securitygroupids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupids", vv)
	}
	if v, found := p.p["securitygroupnames"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupnames", vv)
	}
	if v, found := p.p["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	if v, found := p.p["userdatadetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("userdatadetails[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["userdataid"]; found {
		u.Set("userdataid", v.(string))
	}
	return u
}

func (p *UpdateVirtualMachineParams) SetCleanupdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupdetails"] = v
}

func (p *UpdateVirtualMachineParams) ResetCleanupdetails() {
	if p.p != nil && p.p["cleanupdetails"] != nil {
		delete(p.p, "cleanupdetails")
	}
}

func (p *UpdateVirtualMachineParams) GetCleanupdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupdetails"].(bool)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateVirtualMachineParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateVirtualMachineParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetDeleteprotection(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deleteprotection"] = v
}

func (p *UpdateVirtualMachineParams) ResetDeleteprotection() {
	if p.p != nil && p.p["deleteprotection"] != nil {
		delete(p.p, "deleteprotection")
	}
}

func (p *UpdateVirtualMachineParams) GetDeleteprotection() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deleteprotection"].(bool)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateVirtualMachineParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetDhcpoptionsnetworklist(v []map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpoptionsnetworklist"] = v
}

func (p *UpdateVirtualMachineParams) ResetDhcpoptionsnetworklist() {
	if p.p != nil && p.p["dhcpoptionsnetworklist"] != nil {
		delete(p.p, "dhcpoptionsnetworklist")
	}
}

func (p *UpdateVirtualMachineParams) GetDhcpoptionsnetworklist() ([]map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dhcpoptionsnetworklist"].([]map[string]string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) AddDhcpoptionsnetworklist(item map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	val, found := p.p["dhcpoptionsnetworklist"]
	if !found {
		p.p["dhcpoptionsnetworklist"] = []map[string]string{}
		val = p.p["dhcpoptionsnetworklist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	p.p["dhcpoptionsnetworklist"] = l
}

func (p *UpdateVirtualMachineParams) SetDisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayname"] = v
}

func (p *UpdateVirtualMachineParams) ResetDisplayname() {
	if p.p != nil && p.p["displayname"] != nil {
		delete(p.p, "displayname")
	}
}

func (p *UpdateVirtualMachineParams) GetDisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayname"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
}

func (p *UpdateVirtualMachineParams) ResetDisplayvm() {
	if p.p != nil && p.p["displayvm"] != nil {
		delete(p.p, "displayvm")
	}
}

func (p *UpdateVirtualMachineParams) GetDisplayvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvm"].(bool)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetExtraconfig(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extraconfig"] = v
}

func (p *UpdateVirtualMachineParams) ResetExtraconfig() {
	if p.p != nil && p.p["extraconfig"] != nil {
		delete(p.p, "extraconfig")
	}
}

func (p *UpdateVirtualMachineParams) GetExtraconfig() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extraconfig"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
}

func (p *UpdateVirtualMachineParams) ResetGroup() {
	if p.p != nil && p.p["group"] != nil {
		delete(p.p, "group")
	}
}

func (p *UpdateVirtualMachineParams) GetGroup() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["group"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetHaenable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["haenable"] = v
}

func (p *UpdateVirtualMachineParams) ResetHaenable() {
	if p.p != nil && p.p["haenable"] != nil {
		delete(p.p, "haenable")
	}
}

func (p *UpdateVirtualMachineParams) GetHaenable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["haenable"].(bool)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetInstancename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["instancename"] = v
}

func (p *UpdateVirtualMachineParams) ResetInstancename() {
	if p.p != nil && p.p["instancename"] != nil {
		delete(p.p, "instancename")
	}
}

func (p *UpdateVirtualMachineParams) GetInstancename() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["instancename"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *UpdateVirtualMachineParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *UpdateVirtualMachineParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVirtualMachineParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateVirtualMachineParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *UpdateVirtualMachineParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *UpdateVirtualMachineParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetSecuritygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupids"] = v
}

func (p *UpdateVirtualMachineParams) ResetSecuritygroupids() {
	if p.p != nil && p.p["securitygroupids"] != nil {
		delete(p.p, "securitygroupids")
	}
}

func (p *UpdateVirtualMachineParams) GetSecuritygroupids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupids"].([]string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetSecuritygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupnames"] = v
}

func (p *UpdateVirtualMachineParams) ResetSecuritygroupnames() {
	if p.p != nil && p.p["securitygroupnames"] != nil {
		delete(p.p, "securitygroupnames")
	}
}

func (p *UpdateVirtualMachineParams) GetSecuritygroupnames() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupnames"].([]string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *UpdateVirtualMachineParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *UpdateVirtualMachineParams) GetUserdata() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetUserdatadetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdatadetails"] = v
}

func (p *UpdateVirtualMachineParams) ResetUserdatadetails() {
	if p.p != nil && p.p["userdatadetails"] != nil {
		delete(p.p, "userdatadetails")
	}
}

func (p *UpdateVirtualMachineParams) GetUserdatadetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdatadetails"].(map[string]string)
	return value, ok
}

func (p *UpdateVirtualMachineParams) SetUserdataid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdataid"] = v
}

func (p *UpdateVirtualMachineParams) ResetUserdataid() {
	if p.p != nil && p.p["userdataid"] != nil {
		delete(p.p, "userdataid")
	}
}

func (p *UpdateVirtualMachineParams) GetUserdataid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdataid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewUpdateVirtualMachineParams(id string) *UpdateVirtualMachineParams {
	p := &UpdateVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates properties of a virtual machine. The VM has to be stopped and restarted for the new properties to take effect. UpdateVirtualMachine does not first check whether the VM is stopped. Therefore, stop the VM manually before issuing this call.
func (s *VirtualMachineService) UpdateVirtualMachine(p *UpdateVirtualMachineParams) (*UpdateVirtualMachineResponse, error) {
	resp, err := s.cs.newPostRequest("updateVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateVirtualMachineResponse struct {
	Account               string                                      `json:"account"`
	Affinitygroup         []UpdateVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                      `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                      `json:"autoscalevmgroupname"`
	Backupofferingid      string                                      `json:"backupofferingid"`
	Backupofferingname    string                                      `json:"backupofferingname"`
	Bootmode              string                                      `json:"bootmode"`
	Boottype              string                                      `json:"boottype"`
	Cpunumber             int                                         `json:"cpunumber"`
	Cpuspeed              int                                         `json:"cpuspeed"`
	Cpuused               string                                      `json:"cpuused"`
	Created               string                                      `json:"created"`
	Deleteprotection      bool                                        `json:"deleteprotection"`
	Details               map[string]string                           `json:"details"`
	Diskioread            int64                                       `json:"diskioread"`
	Diskiowrite           int64                                       `json:"diskiowrite"`
	Diskkbsread           int64                                       `json:"diskkbsread"`
	Diskkbswrite          int64                                       `json:"diskkbswrite"`
	Diskofferingid        string                                      `json:"diskofferingid"`
	Diskofferingname      string                                      `json:"diskofferingname"`
	Displayname           string                                      `json:"displayname"`
	Displayvm             bool                                        `json:"displayvm"`
	Domain                string                                      `json:"domain"`
	Domainid              string                                      `json:"domainid"`
	Domainpath            string                                      `json:"domainpath"`
	Forvirtualnetwork     bool                                        `json:"forvirtualnetwork"`
	Group                 string                                      `json:"group"`
	Groupid               string                                      `json:"groupid"`
	Guestosid             string                                      `json:"guestosid"`
	Haenable              bool                                        `json:"haenable"`
	Hasannotations        bool                                        `json:"hasannotations"`
	Hostcontrolstate      string                                      `json:"hostcontrolstate"`
	Hostid                string                                      `json:"hostid"`
	Hostname              string                                      `json:"hostname"`
	Hypervisor            string                                      `json:"hypervisor"`
	Icon                  interface{}                                 `json:"icon"`
	Id                    string                                      `json:"id"`
	Instancename          string                                      `json:"instancename"`
	Ipaddress             string                                      `json:"ipaddress"`
	Isdynamicallyscalable bool                                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                      `json:"isodisplaytext"`
	Isoid                 string                                      `json:"isoid"`
	Isoname               string                                      `json:"isoname"`
	JobID                 string                                      `json:"jobid"`
	Jobstatus             int                                         `json:"jobstatus"`
	Keypairs              string                                      `json:"keypairs"`
	Lastupdated           string                                      `json:"lastupdated"`
	Memory                int                                         `json:"memory"`
	Memoryintfreekbs      int64                                       `json:"memoryintfreekbs"`
	Memorykbs             int64                                       `json:"memorykbs"`
	Memorytargetkbs       int64                                       `json:"memorytargetkbs"`
	Name                  string                                      `json:"name"`
	Networkkbsread        int64                                       `json:"networkkbsread"`
	Networkkbswrite       int64                                       `json:"networkkbswrite"`
	Nic                   []Nic                                       `json:"nic"`
	Osdisplayname         string                                      `json:"osdisplayname"`
	Ostypeid              string                                      `json:"ostypeid"`
	Password              string                                      `json:"password"`
	Passwordenabled       bool                                        `json:"passwordenabled"`
	Pooltype              string                                      `json:"pooltype"`
	Project               string                                      `json:"project"`
	Projectid             string                                      `json:"projectid"`
	Publicip              string                                      `json:"publicip"`
	Publicipid            string                                      `json:"publicipid"`
	Readonlydetails       string                                      `json:"readonlydetails"`
	Receivedbytes         int64                                       `json:"receivedbytes"`
	Rootdeviceid          int64                                       `json:"rootdeviceid"`
	Rootdevicetype        string                                      `json:"rootdevicetype"`
	Securitygroup         []UpdateVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                       `json:"sentbytes"`
	Serviceofferingid     string                                      `json:"serviceofferingid"`
	Serviceofferingname   string                                      `json:"serviceofferingname"`
	Servicestate          string                                      `json:"servicestate"`
	State                 string                                      `json:"state"`
	Tags                  []Tags                                      `json:"tags"`
	Templatedisplaytext   string                                      `json:"templatedisplaytext"`
	Templateformat        string                                      `json:"templateformat"`
	Templateid            string                                      `json:"templateid"`
	Templatename          string                                      `json:"templatename"`
	Templatetype          string                                      `json:"templatetype"`
	Userdata              string                                      `json:"userdata"`
	Userdatadetails       string                                      `json:"userdatadetails"`
	Userdataid            string                                      `json:"userdataid"`
	Userdataname          string                                      `json:"userdataname"`
	Userdatapolicy        string                                      `json:"userdatapolicy"`
	Userid                string                                      `json:"userid"`
	Username              string                                      `json:"username"`
	Vgpu                  string                                      `json:"vgpu"`
	Vmtype                string                                      `json:"vmtype"`
	Vnfdetails            map[string]string                           `json:"vnfdetails"`
	Vnfnics               []string                                    `json:"vnfnics"`
	Zoneid                string                                      `json:"zoneid"`
	Zonename              string                                      `json:"zonename"`
}

type UpdateVirtualMachineResponseSecuritygroup struct {
	Account             string                                          `json:"account"`
	Description         string                                          `json:"description"`
	Domain              string                                          `json:"domain"`
	Domainid            string                                          `json:"domainid"`
	Domainpath          string                                          `json:"domainpath"`
	Egressrule          []UpdateVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                          `json:"id"`
	Ingressrule         []UpdateVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Projectid           string                                          `json:"projectid"`
	Tags                []Tags                                          `json:"tags"`
	Virtualmachinecount int                                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                   `json:"virtualmachineids"`
}

type UpdateVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type UpdateVirtualMachineResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *UpdateVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListVirtualMachinesUsageHistoryParams struct {
	p map[string]interface{}
}

func (p *ListVirtualMachinesUsageHistoryParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
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
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *ListVirtualMachinesUsageHistoryParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *ListVirtualMachinesUsageHistoryParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *ListVirtualMachinesUsageHistoryParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *ListVirtualMachinesUsageHistoryParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVirtualMachinesUsageHistoryParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVirtualMachinesUsageHistoryParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVirtualMachinesUsageHistoryParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListVirtualMachinesUsageHistoryParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListVirtualMachinesUsageHistoryParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListVirtualMachinesUsageHistoryParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVirtualMachinesUsageHistoryParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVirtualMachinesUsageHistoryParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVirtualMachinesUsageHistoryParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVirtualMachinesUsageHistoryParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVirtualMachinesUsageHistoryParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVirtualMachinesUsageHistoryParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVirtualMachinesUsageHistoryParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVirtualMachinesUsageHistoryParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVirtualMachinesUsageHistoryParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVirtualMachinesUsageHistoryParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVirtualMachinesUsageHistoryParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVirtualMachinesUsageHistoryParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *ListVirtualMachinesUsageHistoryParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *ListVirtualMachinesUsageHistoryParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new ListVirtualMachinesUsageHistoryParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVirtualMachinesUsageHistoryParams() *ListVirtualMachinesUsageHistoryParams {
	p := &ListVirtualMachinesUsageHistoryParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesUsageHistoryID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVirtualMachinesUsageHistoryParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVirtualMachinesUsageHistory(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VirtualMachinesUsageHistory[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VirtualMachinesUsageHistory {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesUsageHistoryByName(name string, opts ...OptionFunc) (*VirtualMachinesUsageHistory, int, error) {
	id, count, err := s.GetVirtualMachinesUsageHistoryID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVirtualMachinesUsageHistoryByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesUsageHistoryByID(id string, opts ...OptionFunc) (*VirtualMachinesUsageHistory, int, error) {
	p := &ListVirtualMachinesUsageHistoryParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVirtualMachinesUsageHistory(p)
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
		return l.VirtualMachinesUsageHistory[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualMachinesUsageHistory UUID: %s!", id)
}

// Lists VM stats
func (s *VirtualMachineService) ListVirtualMachinesUsageHistory(p *ListVirtualMachinesUsageHistoryParams) (*ListVirtualMachinesUsageHistoryResponse, error) {
	resp, err := s.cs.newRequest("listVirtualMachinesUsageHistory", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualMachinesUsageHistoryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVirtualMachinesUsageHistoryResponse struct {
	Count                       int                            `json:"count"`
	VirtualMachinesUsageHistory []*VirtualMachinesUsageHistory `json:"virtualmachinesusagehistory"`
}

type VirtualMachinesUsageHistory struct {
	Displayname string   `json:"displayname"`
	Id          string   `json:"id"`
	JobID       string   `json:"jobid"`
	Jobstatus   int      `json:"jobstatus"`
	Name        string   `json:"name"`
	Stats       []string `json:"stats"`
}

type ImportVmParams struct {
	p map[string]interface{}
}

func (p *ImportVmParams) toURLValues() url.Values {
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
	if v, found := p.p["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := p.p["convertinstancehostid"]; found {
		u.Set("convertinstancehostid", v.(string))
	}
	if v, found := p.p["convertinstancepoolid"]; found {
		u.Set("convertinstancepoolid", v.(string))
	}
	if v, found := p.p["datacentername"]; found {
		u.Set("datacentername", v.(string))
	}
	if v, found := p.p["datadiskofferinglist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("datadiskofferinglist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["diskpath"]; found {
		u.Set("diskpath", v.(string))
	}
	if v, found := p.p["displayname"]; found {
		u.Set("displayname", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["existingvcenterid"]; found {
		u.Set("existingvcenterid", v.(string))
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["forcemstoimportvmfiles"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forcemstoimportvmfiles", vv)
	}
	if v, found := p.p["host"]; found {
		u.Set("host", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["hostip"]; found {
		u.Set("hostip", v.(string))
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["importsource"]; found {
		u.Set("importsource", v.(string))
	}
	if v, found := p.p["migrateallowed"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("migrateallowed", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["nicipaddresslist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("nicipaddresslist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := p.p["nicnetworklist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("nicnetworklist[%d].nic", i), k)
			u.Set(fmt.Sprintf("nicnetworklist[%d].network", i), m[k])
		}
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["temppath"]; found {
		u.Set("temppath", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["vcenter"]; found {
		u.Set("vcenter", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ImportVmParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ImportVmParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ImportVmParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ImportVmParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ImportVmParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ImportVmParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetClustername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clustername"] = v
}

func (p *ImportVmParams) ResetClustername() {
	if p.p != nil && p.p["clustername"] != nil {
		delete(p.p, "clustername")
	}
}

func (p *ImportVmParams) GetClustername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clustername"].(string)
	return value, ok
}

func (p *ImportVmParams) SetConvertinstancehostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["convertinstancehostid"] = v
}

func (p *ImportVmParams) ResetConvertinstancehostid() {
	if p.p != nil && p.p["convertinstancehostid"] != nil {
		delete(p.p, "convertinstancehostid")
	}
}

func (p *ImportVmParams) GetConvertinstancehostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["convertinstancehostid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetConvertinstancepoolid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["convertinstancepoolid"] = v
}

func (p *ImportVmParams) ResetConvertinstancepoolid() {
	if p.p != nil && p.p["convertinstancepoolid"] != nil {
		delete(p.p, "convertinstancepoolid")
	}
}

func (p *ImportVmParams) GetConvertinstancepoolid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["convertinstancepoolid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetDatacentername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["datacentername"] = v
}

func (p *ImportVmParams) ResetDatacentername() {
	if p.p != nil && p.p["datacentername"] != nil {
		delete(p.p, "datacentername")
	}
}

func (p *ImportVmParams) GetDatacentername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["datacentername"].(string)
	return value, ok
}

func (p *ImportVmParams) SetDatadiskofferinglist(v []map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["datadiskofferinglist"] = v
}

func (p *ImportVmParams) ResetDatadiskofferinglist() {
	if p.p != nil && p.p["datadiskofferinglist"] != nil {
		delete(p.p, "datadiskofferinglist")
	}
}

func (p *ImportVmParams) GetDatadiskofferinglist() ([]map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["datadiskofferinglist"].([]map[string]string)
	return value, ok
}

func (p *ImportVmParams) AddDatadiskofferinglist(item map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	val, found := p.p["datadiskofferinglist"]
	if !found {
		p.p["datadiskofferinglist"] = []map[string]string{}
		val = p.p["datadiskofferinglist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	p.p["datadiskofferinglist"] = l
}

func (p *ImportVmParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ImportVmParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ImportVmParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *ImportVmParams) SetDiskpath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskpath"] = v
}

func (p *ImportVmParams) ResetDiskpath() {
	if p.p != nil && p.p["diskpath"] != nil {
		delete(p.p, "diskpath")
	}
}

func (p *ImportVmParams) GetDiskpath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskpath"].(string)
	return value, ok
}

func (p *ImportVmParams) SetDisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayname"] = v
}

func (p *ImportVmParams) ResetDisplayname() {
	if p.p != nil && p.p["displayname"] != nil {
		delete(p.p, "displayname")
	}
}

func (p *ImportVmParams) GetDisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayname"].(string)
	return value, ok
}

func (p *ImportVmParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ImportVmParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ImportVmParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetExistingvcenterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["existingvcenterid"] = v
}

func (p *ImportVmParams) ResetExistingvcenterid() {
	if p.p != nil && p.p["existingvcenterid"] != nil {
		delete(p.p, "existingvcenterid")
	}
}

func (p *ImportVmParams) GetExistingvcenterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["existingvcenterid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *ImportVmParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *ImportVmParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *ImportVmParams) SetForcemstoimportvmfiles(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forcemstoimportvmfiles"] = v
}

func (p *ImportVmParams) ResetForcemstoimportvmfiles() {
	if p.p != nil && p.p["forcemstoimportvmfiles"] != nil {
		delete(p.p, "forcemstoimportvmfiles")
	}
}

func (p *ImportVmParams) GetForcemstoimportvmfiles() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forcemstoimportvmfiles"].(bool)
	return value, ok
}

func (p *ImportVmParams) SetHost(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["host"] = v
}

func (p *ImportVmParams) ResetHost() {
	if p.p != nil && p.p["host"] != nil {
		delete(p.p, "host")
	}
}

func (p *ImportVmParams) GetHost() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["host"].(string)
	return value, ok
}

func (p *ImportVmParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ImportVmParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ImportVmParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetHostip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostip"] = v
}

func (p *ImportVmParams) ResetHostip() {
	if p.p != nil && p.p["hostip"] != nil {
		delete(p.p, "hostip")
	}
}

func (p *ImportVmParams) GetHostip() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostip"].(string)
	return value, ok
}

func (p *ImportVmParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *ImportVmParams) ResetHostname() {
	if p.p != nil && p.p["hostname"] != nil {
		delete(p.p, "hostname")
	}
}

func (p *ImportVmParams) GetHostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostname"].(string)
	return value, ok
}

func (p *ImportVmParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ImportVmParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ImportVmParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ImportVmParams) SetImportsource(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["importsource"] = v
}

func (p *ImportVmParams) ResetImportsource() {
	if p.p != nil && p.p["importsource"] != nil {
		delete(p.p, "importsource")
	}
}

func (p *ImportVmParams) GetImportsource() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["importsource"].(string)
	return value, ok
}

func (p *ImportVmParams) SetMigrateallowed(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["migrateallowed"] = v
}

func (p *ImportVmParams) ResetMigrateallowed() {
	if p.p != nil && p.p["migrateallowed"] != nil {
		delete(p.p, "migrateallowed")
	}
}

func (p *ImportVmParams) GetMigrateallowed() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["migrateallowed"].(bool)
	return value, ok
}

func (p *ImportVmParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ImportVmParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ImportVmParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ImportVmParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ImportVmParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ImportVmParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetNicipaddresslist(v []map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicipaddresslist"] = v
}

func (p *ImportVmParams) ResetNicipaddresslist() {
	if p.p != nil && p.p["nicipaddresslist"] != nil {
		delete(p.p, "nicipaddresslist")
	}
}

func (p *ImportVmParams) GetNicipaddresslist() ([]map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicipaddresslist"].([]map[string]string)
	return value, ok
}

func (p *ImportVmParams) AddNicipaddresslist(item map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	val, found := p.p["nicipaddresslist"]
	if !found {
		p.p["nicipaddresslist"] = []map[string]string{}
		val = p.p["nicipaddresslist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	p.p["nicipaddresslist"] = l
}

func (p *ImportVmParams) SetNicnetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicnetworklist"] = v
}

func (p *ImportVmParams) ResetNicnetworklist() {
	if p.p != nil && p.p["nicnetworklist"] != nil {
		delete(p.p, "nicnetworklist")
	}
}

func (p *ImportVmParams) GetNicnetworklist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicnetworklist"].(map[string]string)
	return value, ok
}

func (p *ImportVmParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *ImportVmParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *ImportVmParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *ImportVmParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ImportVmParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ImportVmParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ImportVmParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ImportVmParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ImportVmParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *ImportVmParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *ImportVmParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *ImportVmParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *ImportVmParams) SetTemppath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["temppath"] = v
}

func (p *ImportVmParams) ResetTemppath() {
	if p.p != nil && p.p["temppath"] != nil {
		delete(p.p, "temppath")
	}
}

func (p *ImportVmParams) GetTemppath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["temppath"].(string)
	return value, ok
}

func (p *ImportVmParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *ImportVmParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *ImportVmParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

func (p *ImportVmParams) SetVcenter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vcenter"] = v
}

func (p *ImportVmParams) ResetVcenter() {
	if p.p != nil && p.p["vcenter"] != nil {
		delete(p.p, "vcenter")
	}
}

func (p *ImportVmParams) GetVcenter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vcenter"].(string)
	return value, ok
}

func (p *ImportVmParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ImportVmParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ImportVmParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ImportVmParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewImportVmParams(clusterid string, hypervisor string, importsource string, name string, serviceofferingid string, zoneid string) *ImportVmParams {
	p := &ImportVmParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	p.p["hypervisor"] = hypervisor
	p.p["importsource"] = importsource
	p.p["name"] = name
	p.p["serviceofferingid"] = serviceofferingid
	p.p["zoneid"] = zoneid
	return p
}

// Import virtual machine from a unmanaged host into CloudStack
func (s *VirtualMachineService) ImportVm(p *ImportVmParams) (*ImportVmResponse, error) {
	resp, err := s.cs.newRequest("importVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportVmResponse
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

type ImportVmResponse struct {
	Account               string                          `json:"account"`
	Affinitygroup         []ImportVmResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                          `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                          `json:"autoscalevmgroupname"`
	Backupofferingid      string                          `json:"backupofferingid"`
	Backupofferingname    string                          `json:"backupofferingname"`
	Bootmode              string                          `json:"bootmode"`
	Boottype              string                          `json:"boottype"`
	Cpunumber             int                             `json:"cpunumber"`
	Cpuspeed              int                             `json:"cpuspeed"`
	Cpuused               string                          `json:"cpuused"`
	Created               string                          `json:"created"`
	Deleteprotection      bool                            `json:"deleteprotection"`
	Details               map[string]string               `json:"details"`
	Diskioread            int64                           `json:"diskioread"`
	Diskiowrite           int64                           `json:"diskiowrite"`
	Diskkbsread           int64                           `json:"diskkbsread"`
	Diskkbswrite          int64                           `json:"diskkbswrite"`
	Diskofferingid        string                          `json:"diskofferingid"`
	Diskofferingname      string                          `json:"diskofferingname"`
	Displayname           string                          `json:"displayname"`
	Displayvm             bool                            `json:"displayvm"`
	Domain                string                          `json:"domain"`
	Domainid              string                          `json:"domainid"`
	Domainpath            string                          `json:"domainpath"`
	Forvirtualnetwork     bool                            `json:"forvirtualnetwork"`
	Group                 string                          `json:"group"`
	Groupid               string                          `json:"groupid"`
	Guestosid             string                          `json:"guestosid"`
	Haenable              bool                            `json:"haenable"`
	Hasannotations        bool                            `json:"hasannotations"`
	Hostcontrolstate      string                          `json:"hostcontrolstate"`
	Hostid                string                          `json:"hostid"`
	Hostname              string                          `json:"hostname"`
	Hypervisor            string                          `json:"hypervisor"`
	Icon                  interface{}                     `json:"icon"`
	Id                    string                          `json:"id"`
	Instancename          string                          `json:"instancename"`
	Ipaddress             string                          `json:"ipaddress"`
	Isdynamicallyscalable bool                            `json:"isdynamicallyscalable"`
	Isodisplaytext        string                          `json:"isodisplaytext"`
	Isoid                 string                          `json:"isoid"`
	Isoname               string                          `json:"isoname"`
	JobID                 string                          `json:"jobid"`
	Jobstatus             int                             `json:"jobstatus"`
	Keypairs              string                          `json:"keypairs"`
	Lastupdated           string                          `json:"lastupdated"`
	Memory                int                             `json:"memory"`
	Memoryintfreekbs      int64                           `json:"memoryintfreekbs"`
	Memorykbs             int64                           `json:"memorykbs"`
	Memorytargetkbs       int64                           `json:"memorytargetkbs"`
	Name                  string                          `json:"name"`
	Networkkbsread        int64                           `json:"networkkbsread"`
	Networkkbswrite       int64                           `json:"networkkbswrite"`
	Nic                   []Nic                           `json:"nic"`
	Osdisplayname         string                          `json:"osdisplayname"`
	Ostypeid              string                          `json:"ostypeid"`
	Password              string                          `json:"password"`
	Passwordenabled       bool                            `json:"passwordenabled"`
	Pooltype              string                          `json:"pooltype"`
	Project               string                          `json:"project"`
	Projectid             string                          `json:"projectid"`
	Publicip              string                          `json:"publicip"`
	Publicipid            string                          `json:"publicipid"`
	Readonlydetails       string                          `json:"readonlydetails"`
	Receivedbytes         int64                           `json:"receivedbytes"`
	Rootdeviceid          int64                           `json:"rootdeviceid"`
	Rootdevicetype        string                          `json:"rootdevicetype"`
	Securitygroup         []ImportVmResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                           `json:"sentbytes"`
	Serviceofferingid     string                          `json:"serviceofferingid"`
	Serviceofferingname   string                          `json:"serviceofferingname"`
	Servicestate          string                          `json:"servicestate"`
	State                 string                          `json:"state"`
	Tags                  []Tags                          `json:"tags"`
	Templatedisplaytext   string                          `json:"templatedisplaytext"`
	Templateformat        string                          `json:"templateformat"`
	Templateid            string                          `json:"templateid"`
	Templatename          string                          `json:"templatename"`
	Templatetype          string                          `json:"templatetype"`
	Userdata              string                          `json:"userdata"`
	Userdatadetails       string                          `json:"userdatadetails"`
	Userdataid            string                          `json:"userdataid"`
	Userdataname          string                          `json:"userdataname"`
	Userdatapolicy        string                          `json:"userdatapolicy"`
	Userid                string                          `json:"userid"`
	Username              string                          `json:"username"`
	Vgpu                  string                          `json:"vgpu"`
	Vmtype                string                          `json:"vmtype"`
	Vnfdetails            map[string]string               `json:"vnfdetails"`
	Vnfnics               []string                        `json:"vnfnics"`
	Zoneid                string                          `json:"zoneid"`
	Zonename              string                          `json:"zonename"`
}

type ImportVmResponseSecuritygroup struct {
	Account             string                              `json:"account"`
	Description         string                              `json:"description"`
	Domain              string                              `json:"domain"`
	Domainid            string                              `json:"domainid"`
	Domainpath          string                              `json:"domainpath"`
	Egressrule          []ImportVmResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                              `json:"id"`
	Ingressrule         []ImportVmResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                              `json:"name"`
	Project             string                              `json:"project"`
	Projectid           string                              `json:"projectid"`
	Tags                []Tags                              `json:"tags"`
	Virtualmachinecount int                                 `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                       `json:"virtualmachineids"`
}

type ImportVmResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type ImportVmResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *ImportVmResponse) UnmarshalJSON(b []byte) error {
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

	type alias ImportVmResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UnmanageVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *UnmanageVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UnmanageVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UnmanageVirtualMachineParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UnmanageVirtualMachineParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UnmanageVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewUnmanageVirtualMachineParams(id string) *UnmanageVirtualMachineParams {
	p := &UnmanageVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Unmanage a guest virtual machine.
func (s *VirtualMachineService) UnmanageVirtualMachine(p *UnmanageVirtualMachineParams) (*UnmanageVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("unmanageVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UnmanageVirtualMachineResponse
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

type UnmanageVirtualMachineResponse struct {
	Details   string `json:"details"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Success   bool   `json:"success"`
}

type ListUnmanagedInstancesParams struct {
	p map[string]interface{}
}

func (p *ListUnmanagedInstancesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
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
	return u
}

func (p *ListUnmanagedInstancesParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListUnmanagedInstancesParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListUnmanagedInstancesParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListUnmanagedInstancesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListUnmanagedInstancesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListUnmanagedInstancesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListUnmanagedInstancesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListUnmanagedInstancesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListUnmanagedInstancesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListUnmanagedInstancesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListUnmanagedInstancesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListUnmanagedInstancesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListUnmanagedInstancesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListUnmanagedInstancesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListUnmanagedInstancesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListUnmanagedInstancesParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListUnmanagedInstancesParams(clusterid string) *ListUnmanagedInstancesParams {
	p := &ListUnmanagedInstancesParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	return p
}

// Lists unmanaged virtual machines for a given cluster.
func (s *VirtualMachineService) ListUnmanagedInstances(p *ListUnmanagedInstancesParams) (*ListUnmanagedInstancesResponse, error) {
	resp, err := s.cs.newRequest("listUnmanagedInstances", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUnmanagedInstancesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUnmanagedInstancesResponse struct {
	Count              int                  `json:"count"`
	UnmanagedInstances []*UnmanagedInstance `json:"unmanagedinstance"`
}

type UnmanagedInstance struct {
	Clusterid        string                  `json:"clusterid"`
	Clustername      string                  `json:"clustername"`
	Cpucorepersocket int                     `json:"cpucorepersocket"`
	Cpunumber        int                     `json:"cpunumber"`
	Cpuspeed         int                     `json:"cpuspeed"`
	Disk             []UnmanagedInstanceDisk `json:"disk"`
	Hostid           string                  `json:"hostid"`
	Hostname         string                  `json:"hostname"`
	JobID            string                  `json:"jobid"`
	Jobstatus        int                     `json:"jobstatus"`
	Memory           int                     `json:"memory"`
	Name             string                  `json:"name"`
	Nic              []Nic                   `json:"nic"`
	Osdisplayname    string                  `json:"osdisplayname"`
	Osid             string                  `json:"osid"`
	Powerstate       string                  `json:"powerstate"`
}

type UnmanagedInstanceDisk struct {
	Capacity       int64  `json:"capacity"`
	Controller     string `json:"controller"`
	Controllerunit int    `json:"controllerunit"`
	Datastorehost  string `json:"datastorehost"`
	Datastorename  string `json:"datastorename"`
	Datastorepath  string `json:"datastorepath"`
	Datastoretype  string `json:"datastoretype"`
	Id             string `json:"id"`
	Imagepath      string `json:"imagepath"`
	Label          string `json:"label"`
	Position       int    `json:"position"`
}

type ImportUnmanagedInstanceParams struct {
	p map[string]interface{}
}

func (p *ImportUnmanagedInstanceParams) toURLValues() url.Values {
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
	if v, found := p.p["datadiskofferinglist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("datadiskofferinglist[%d].disk", i), k)
			u.Set(fmt.Sprintf("datadiskofferinglist[%d].diskOffering", i), m[k])
		}
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["displayname"]; found {
		u.Set("displayname", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["migrateallowed"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("migrateallowed", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["nicipaddresslist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("nicipaddresslist[%d].nic", i), k)
			u.Set(fmt.Sprintf("nicipaddresslist[%d].ip4Address", i), m[k])
		}
	}
	if v, found := p.p["nicnetworklist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("nicnetworklist[%d].nic", i), k)
			u.Set(fmt.Sprintf("nicnetworklist[%d].network", i), m[k])
		}
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	return u
}

func (p *ImportUnmanagedInstanceParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ImportUnmanagedInstanceParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ImportUnmanagedInstanceParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetDatadiskofferinglist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["datadiskofferinglist"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetDatadiskofferinglist() {
	if p.p != nil && p.p["datadiskofferinglist"] != nil {
		delete(p.p, "datadiskofferinglist")
	}
}

func (p *ImportUnmanagedInstanceParams) GetDatadiskofferinglist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["datadiskofferinglist"].(map[string]string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ImportUnmanagedInstanceParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetDisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayname"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetDisplayname() {
	if p.p != nil && p.p["displayname"] != nil {
		delete(p.p, "displayname")
	}
}

func (p *ImportUnmanagedInstanceParams) GetDisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayname"].(string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ImportUnmanagedInstanceParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *ImportUnmanagedInstanceParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetHostname() {
	if p.p != nil && p.p["hostname"] != nil {
		delete(p.p, "hostname")
	}
}

func (p *ImportUnmanagedInstanceParams) GetHostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostname"].(string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetMigrateallowed(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["migrateallowed"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetMigrateallowed() {
	if p.p != nil && p.p["migrateallowed"] != nil {
		delete(p.p, "migrateallowed")
	}
}

func (p *ImportUnmanagedInstanceParams) GetMigrateallowed() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["migrateallowed"].(bool)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ImportUnmanagedInstanceParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetNicipaddresslist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicipaddresslist"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetNicipaddresslist() {
	if p.p != nil && p.p["nicipaddresslist"] != nil {
		delete(p.p, "nicipaddresslist")
	}
}

func (p *ImportUnmanagedInstanceParams) GetNicipaddresslist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicipaddresslist"].(map[string]string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetNicnetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicnetworklist"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetNicnetworklist() {
	if p.p != nil && p.p["nicnetworklist"] != nil {
		delete(p.p, "nicnetworklist")
	}
}

func (p *ImportUnmanagedInstanceParams) GetNicnetworklist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicnetworklist"].(map[string]string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ImportUnmanagedInstanceParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ImportUnmanagedInstanceParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ImportUnmanagedInstanceParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *ImportUnmanagedInstanceParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *ImportUnmanagedInstanceParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

// You should always use this function to get a new ImportUnmanagedInstanceParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewImportUnmanagedInstanceParams(clusterid string, name string, serviceofferingid string) *ImportUnmanagedInstanceParams {
	p := &ImportUnmanagedInstanceParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	p.p["name"] = name
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Import unmanaged virtual machine from a given cluster.
func (s *VirtualMachineService) ImportUnmanagedInstance(p *ImportUnmanagedInstanceParams) (*ImportUnmanagedInstanceResponse, error) {
	resp, err := s.cs.newRequest("importUnmanagedInstance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportUnmanagedInstanceResponse
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

type ImportUnmanagedInstanceResponse struct {
	Account               string                                         `json:"account"`
	Affinitygroup         []ImportUnmanagedInstanceResponseAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                                         `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                                         `json:"autoscalevmgroupname"`
	Backupofferingid      string                                         `json:"backupofferingid"`
	Backupofferingname    string                                         `json:"backupofferingname"`
	Bootmode              string                                         `json:"bootmode"`
	Boottype              string                                         `json:"boottype"`
	Cpunumber             int                                            `json:"cpunumber"`
	Cpuspeed              int                                            `json:"cpuspeed"`
	Cpuused               string                                         `json:"cpuused"`
	Created               string                                         `json:"created"`
	Deleteprotection      bool                                           `json:"deleteprotection"`
	Details               map[string]string                              `json:"details"`
	Diskioread            int64                                          `json:"diskioread"`
	Diskiowrite           int64                                          `json:"diskiowrite"`
	Diskkbsread           int64                                          `json:"diskkbsread"`
	Diskkbswrite          int64                                          `json:"diskkbswrite"`
	Diskofferingid        string                                         `json:"diskofferingid"`
	Diskofferingname      string                                         `json:"diskofferingname"`
	Displayname           string                                         `json:"displayname"`
	Displayvm             bool                                           `json:"displayvm"`
	Domain                string                                         `json:"domain"`
	Domainid              string                                         `json:"domainid"`
	Domainpath            string                                         `json:"domainpath"`
	Forvirtualnetwork     bool                                           `json:"forvirtualnetwork"`
	Group                 string                                         `json:"group"`
	Groupid               string                                         `json:"groupid"`
	Guestosid             string                                         `json:"guestosid"`
	Haenable              bool                                           `json:"haenable"`
	Hasannotations        bool                                           `json:"hasannotations"`
	Hostcontrolstate      string                                         `json:"hostcontrolstate"`
	Hostid                string                                         `json:"hostid"`
	Hostname              string                                         `json:"hostname"`
	Hypervisor            string                                         `json:"hypervisor"`
	Icon                  interface{}                                    `json:"icon"`
	Id                    string                                         `json:"id"`
	Instancename          string                                         `json:"instancename"`
	Ipaddress             string                                         `json:"ipaddress"`
	Isdynamicallyscalable bool                                           `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                         `json:"isodisplaytext"`
	Isoid                 string                                         `json:"isoid"`
	Isoname               string                                         `json:"isoname"`
	JobID                 string                                         `json:"jobid"`
	Jobstatus             int                                            `json:"jobstatus"`
	Keypairs              string                                         `json:"keypairs"`
	Lastupdated           string                                         `json:"lastupdated"`
	Memory                int                                            `json:"memory"`
	Memoryintfreekbs      int64                                          `json:"memoryintfreekbs"`
	Memorykbs             int64                                          `json:"memorykbs"`
	Memorytargetkbs       int64                                          `json:"memorytargetkbs"`
	Name                  string                                         `json:"name"`
	Networkkbsread        int64                                          `json:"networkkbsread"`
	Networkkbswrite       int64                                          `json:"networkkbswrite"`
	Nic                   []Nic                                          `json:"nic"`
	Osdisplayname         string                                         `json:"osdisplayname"`
	Ostypeid              string                                         `json:"ostypeid"`
	Password              string                                         `json:"password"`
	Passwordenabled       bool                                           `json:"passwordenabled"`
	Pooltype              string                                         `json:"pooltype"`
	Project               string                                         `json:"project"`
	Projectid             string                                         `json:"projectid"`
	Publicip              string                                         `json:"publicip"`
	Publicipid            string                                         `json:"publicipid"`
	Readonlydetails       string                                         `json:"readonlydetails"`
	Receivedbytes         int64                                          `json:"receivedbytes"`
	Rootdeviceid          int64                                          `json:"rootdeviceid"`
	Rootdevicetype        string                                         `json:"rootdevicetype"`
	Securitygroup         []ImportUnmanagedInstanceResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                          `json:"sentbytes"`
	Serviceofferingid     string                                         `json:"serviceofferingid"`
	Serviceofferingname   string                                         `json:"serviceofferingname"`
	Servicestate          string                                         `json:"servicestate"`
	State                 string                                         `json:"state"`
	Tags                  []Tags                                         `json:"tags"`
	Templatedisplaytext   string                                         `json:"templatedisplaytext"`
	Templateformat        string                                         `json:"templateformat"`
	Templateid            string                                         `json:"templateid"`
	Templatename          string                                         `json:"templatename"`
	Templatetype          string                                         `json:"templatetype"`
	Userdata              string                                         `json:"userdata"`
	Userdatadetails       string                                         `json:"userdatadetails"`
	Userdataid            string                                         `json:"userdataid"`
	Userdataname          string                                         `json:"userdataname"`
	Userdatapolicy        string                                         `json:"userdatapolicy"`
	Userid                string                                         `json:"userid"`
	Username              string                                         `json:"username"`
	Vgpu                  string                                         `json:"vgpu"`
	Vmtype                string                                         `json:"vmtype"`
	Vnfdetails            map[string]string                              `json:"vnfdetails"`
	Vnfnics               []string                                       `json:"vnfnics"`
	Zoneid                string                                         `json:"zoneid"`
	Zonename              string                                         `json:"zonename"`
}

type ImportUnmanagedInstanceResponseSecuritygroup struct {
	Account             string                                             `json:"account"`
	Description         string                                             `json:"description"`
	Domain              string                                             `json:"domain"`
	Domainid            string                                             `json:"domainid"`
	Domainpath          string                                             `json:"domainpath"`
	Egressrule          []ImportUnmanagedInstanceResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                             `json:"id"`
	Ingressrule         []ImportUnmanagedInstanceResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                             `json:"name"`
	Project             string                                             `json:"project"`
	Projectid           string                                             `json:"projectid"`
	Tags                []Tags                                             `json:"tags"`
	Virtualmachinecount int                                                `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                      `json:"virtualmachineids"`
}

type ImportUnmanagedInstanceResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type ImportUnmanagedInstanceResponseAffinitygroup struct {
	Account            string   `json:"account"`
	Dedicatedresources []string `json:"dedicatedresources"`
	Description        string   `json:"description"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Type               string   `json:"type"`
	VirtualmachineIds  []string `json:"virtualmachineIds"`
}

func (r *ImportUnmanagedInstanceResponse) UnmarshalJSON(b []byte) error {
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

	type alias ImportUnmanagedInstanceResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CreateVMScheduleParams struct {
	p map[string]interface{}
}

func (p *CreateVMScheduleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["schedule"]; found {
		u.Set("schedule", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *CreateVMScheduleParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *CreateVMScheduleParams) ResetAction() {
	if p.p != nil && p.p["action"] != nil {
		delete(p.p, "action")
	}
}

func (p *CreateVMScheduleParams) GetAction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["action"].(string)
	return value, ok
}

func (p *CreateVMScheduleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateVMScheduleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateVMScheduleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateVMScheduleParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *CreateVMScheduleParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *CreateVMScheduleParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *CreateVMScheduleParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *CreateVMScheduleParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *CreateVMScheduleParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *CreateVMScheduleParams) SetSchedule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["schedule"] = v
}

func (p *CreateVMScheduleParams) ResetSchedule() {
	if p.p != nil && p.p["schedule"] != nil {
		delete(p.p, "schedule")
	}
}

func (p *CreateVMScheduleParams) GetSchedule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["schedule"].(string)
	return value, ok
}

func (p *CreateVMScheduleParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *CreateVMScheduleParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *CreateVMScheduleParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *CreateVMScheduleParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *CreateVMScheduleParams) ResetTimezone() {
	if p.p != nil && p.p["timezone"] != nil {
		delete(p.p, "timezone")
	}
}

func (p *CreateVMScheduleParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

func (p *CreateVMScheduleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *CreateVMScheduleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *CreateVMScheduleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVMScheduleParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewCreateVMScheduleParams(action string, schedule string, timezone string, virtualmachineid string) *CreateVMScheduleParams {
	p := &CreateVMScheduleParams{}
	p.p = make(map[string]interface{})
	p.p["action"] = action
	p.p["schedule"] = schedule
	p.p["timezone"] = timezone
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Create VM Schedule
func (s *VirtualMachineService) CreateVMSchedule(p *CreateVMScheduleParams) (*CreateVMScheduleResponse, error) {
	resp, err := s.cs.newRequest("createVMSchedule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response CreateVMScheduleResponse `json:"vmschedule"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type CreateVMScheduleResponse struct {
	Action           string `json:"action"`
	Created          string `json:"created"`
	Description      string `json:"description"`
	Enabled          bool   `json:"enabled"`
	Enddate          string `json:"enddate"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Schedule         string `json:"schedule"`
	Startdate        string `json:"startdate"`
	Timezone         string `json:"timezone"`
	Virtualmachineid string `json:"virtualmachineid"`
}

type UpdateVMScheduleParams struct {
	p map[string]interface{}
}

func (p *UpdateVMScheduleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["schedule"]; found {
		u.Set("schedule", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	return u
}

func (p *UpdateVMScheduleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateVMScheduleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateVMScheduleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateVMScheduleParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *UpdateVMScheduleParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *UpdateVMScheduleParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *UpdateVMScheduleParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *UpdateVMScheduleParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *UpdateVMScheduleParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *UpdateVMScheduleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVMScheduleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVMScheduleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateVMScheduleParams) SetSchedule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["schedule"] = v
}

func (p *UpdateVMScheduleParams) ResetSchedule() {
	if p.p != nil && p.p["schedule"] != nil {
		delete(p.p, "schedule")
	}
}

func (p *UpdateVMScheduleParams) GetSchedule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["schedule"].(string)
	return value, ok
}

func (p *UpdateVMScheduleParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *UpdateVMScheduleParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *UpdateVMScheduleParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *UpdateVMScheduleParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *UpdateVMScheduleParams) ResetTimezone() {
	if p.p != nil && p.p["timezone"] != nil {
		delete(p.p, "timezone")
	}
}

func (p *UpdateVMScheduleParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVMScheduleParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewUpdateVMScheduleParams(id string) *UpdateVMScheduleParams {
	p := &UpdateVMScheduleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Update VM Schedule.
func (s *VirtualMachineService) UpdateVMSchedule(p *UpdateVMScheduleParams) (*UpdateVMScheduleResponse, error) {
	resp, err := s.cs.newRequest("updateVMSchedule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response UpdateVMScheduleResponse `json:"vmschedule"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type UpdateVMScheduleResponse struct {
	Action           string `json:"action"`
	Created          string `json:"created"`
	Description      string `json:"description"`
	Enabled          bool   `json:"enabled"`
	Enddate          string `json:"enddate"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Schedule         string `json:"schedule"`
	Startdate        string `json:"startdate"`
	Timezone         string `json:"timezone"`
	Virtualmachineid string `json:"virtualmachineid"`
}

type ListVMScheduleParams struct {
	p map[string]interface{}
}

func (p *ListVMScheduleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
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
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListVMScheduleParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *ListVMScheduleParams) ResetAction() {
	if p.p != nil && p.p["action"] != nil {
		delete(p.p, "action")
	}
}

func (p *ListVMScheduleParams) GetAction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["action"].(string)
	return value, ok
}

func (p *ListVMScheduleParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *ListVMScheduleParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *ListVMScheduleParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *ListVMScheduleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVMScheduleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVMScheduleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVMScheduleParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVMScheduleParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVMScheduleParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVMScheduleParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVMScheduleParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVMScheduleParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVMScheduleParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVMScheduleParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVMScheduleParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVMScheduleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListVMScheduleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *ListVMScheduleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVMScheduleParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVMScheduleParams(virtualmachineid string) *ListVMScheduleParams {
	p := &ListVMScheduleParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVMScheduleByID(id string, virtualmachineid string, opts ...OptionFunc) (*VMSchedule, int, error) {
	p := &ListVMScheduleParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id
	p.p["virtualmachineid"] = virtualmachineid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVMSchedule(p)
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
		return l.VMSchedule[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VMSchedule UUID: %s!", id)
}

// List VM Schedules.
func (s *VirtualMachineService) ListVMSchedule(p *ListVMScheduleParams) (*ListVMScheduleResponse, error) {
	resp, err := s.cs.newRequest("listVMSchedule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVMScheduleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVMScheduleResponse struct {
	Count      int           `json:"count"`
	VMSchedule []*VMSchedule `json:"vmschedule"`
}

type VMSchedule struct {
	Action           string `json:"action"`
	Created          string `json:"created"`
	Description      string `json:"description"`
	Enabled          bool   `json:"enabled"`
	Enddate          string `json:"enddate"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Schedule         string `json:"schedule"`
	Startdate        string `json:"startdate"`
	Timezone         string `json:"timezone"`
	Virtualmachineid string `json:"virtualmachineid"`
}

type DeleteVMScheduleParams struct {
	p map[string]interface{}
}

func (p *DeleteVMScheduleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *DeleteVMScheduleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVMScheduleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteVMScheduleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteVMScheduleParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *DeleteVMScheduleParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *DeleteVMScheduleParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *DeleteVMScheduleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *DeleteVMScheduleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *DeleteVMScheduleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVMScheduleParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDeleteVMScheduleParams(virtualmachineid string) *DeleteVMScheduleParams {
	p := &DeleteVMScheduleParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Delete VM Schedule.
func (s *VirtualMachineService) DeleteVMSchedule(p *DeleteVMScheduleParams) (*DeleteVMScheduleResponse, error) {
	resp, err := s.cs.newRequest("deleteVMSchedule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVMScheduleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteVMScheduleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteVMScheduleResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteVMScheduleResponse
	return json.Unmarshal(b, (*alias)(r))
}
