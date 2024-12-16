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

type VirtualNetworkFunctionsServiceIface interface {
	DeleteVnfTemplate(p *DeleteVnfTemplateParams) (*DeleteVnfTemplateResponse, error)
	NewDeleteVnfTemplateParams(id string) *DeleteVnfTemplateParams
	DeployVnfAppliance(p *DeployVnfApplianceParams) (*DeployVnfApplianceResponse, error)
	NewDeployVnfApplianceParams(serviceofferingid string, templateid string, zoneid string) *DeployVnfApplianceParams
	ListVnfAppliances(p *ListVnfAppliancesParams) (*ListVnfAppliancesResponse, error)
	NewListVnfAppliancesParams() *ListVnfAppliancesParams
	GetVnfApplianceID(name string, opts ...OptionFunc) (string, int, error)
	GetVnfApplianceByName(name string, opts ...OptionFunc) (*VnfAppliance, int, error)
	GetVnfApplianceByID(id string, opts ...OptionFunc) (*VnfAppliance, int, error)
	ListVnfTemplates(p *ListVnfTemplatesParams) (*ListVnfTemplatesResponse, error)
	NewListVnfTemplatesParams(templatefilter string) *ListVnfTemplatesParams
	GetVnfTemplateID(name string, templatefilter string, opts ...OptionFunc) (string, int, error)
	GetVnfTemplateByName(name string, templatefilter string, opts ...OptionFunc) (*VnfTemplate, int, error)
	GetVnfTemplateByID(id string, templatefilter string, opts ...OptionFunc) (*VnfTemplate, int, error)
	RegisterVnfTemplate(p *RegisterVnfTemplateParams) (*RegisterVnfTemplateResponse, error)
	NewRegisterVnfTemplateParams(format string, hypervisor string, name string, url string) *RegisterVnfTemplateParams
	UpdateVnfTemplate(p *UpdateVnfTemplateParams) (*UpdateVnfTemplateResponse, error)
	NewUpdateVnfTemplateParams(id string) *UpdateVnfTemplateParams
}

type DeleteVnfTemplateParams struct {
	p map[string]interface{}
}

func (p *DeleteVnfTemplateParams) toURLValues() url.Values {
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
	if v, found := p.p["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteVnfTemplateParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DeleteVnfTemplateParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *DeleteVnfTemplateParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *DeleteVnfTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVnfTemplateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteVnfTemplateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteVnfTemplateParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
}

func (p *DeleteVnfTemplateParams) ResetIssystem() {
	if p.p != nil && p.p["issystem"] != nil {
		delete(p.p, "issystem")
	}
}

func (p *DeleteVnfTemplateParams) GetIssystem() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["issystem"].(bool)
	return value, ok
}

func (p *DeleteVnfTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteVnfTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteVnfTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVnfTemplateParams instance,
// as then you are sure you have configured all required params
func (s *VirtualNetworkFunctionsService) NewDeleteVnfTemplateParams(id string) *DeleteVnfTemplateParams {
	p := &DeleteVnfTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a VNF template from the system. All virtual machines using the deleted template will not be affected.
func (s *VirtualNetworkFunctionsService) DeleteVnfTemplate(p *DeleteVnfTemplateParams) (*DeleteVnfTemplateResponse, error) {
	resp, err := s.cs.newRequest("deleteVnfTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVnfTemplateResponse
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

type DeleteVnfTemplateResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeployVnfApplianceParams struct {
	p map[string]interface{}
}

func (p *DeployVnfApplianceParams) toURLValues() url.Values {
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
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("dhcpoptionsnetworklist[%d].key", i), k)
			u.Set(fmt.Sprintf("dhcpoptionsnetworklist[%d].value", i), m[k])
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
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("iptonetworklist[%d].key", i), k)
			u.Set(fmt.Sprintf("iptonetworklist[%d].value", i), m[k])
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
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("nicnetworklist[%d].nic", i), k)
			u.Set(fmt.Sprintf("nicnetworklist[%d].network", i), m[k])
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
	if v, found := p.p["vnfcidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("vnfcidrlist", vv)
	}
	if v, found := p.p["vnfconfiguremanagement"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("vnfconfiguremanagement", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeployVnfApplianceParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DeployVnfApplianceParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DeployVnfApplianceParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetAffinitygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupids"] = v
}

func (p *DeployVnfApplianceParams) ResetAffinitygroupids() {
	if p.p != nil && p.p["affinitygroupids"] != nil {
		delete(p.p, "affinitygroupids")
	}
}

func (p *DeployVnfApplianceParams) GetAffinitygroupids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupids"].([]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetAffinitygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupnames"] = v
}

func (p *DeployVnfApplianceParams) ResetAffinitygroupnames() {
	if p.p != nil && p.p["affinitygroupnames"] != nil {
		delete(p.p, "affinitygroupnames")
	}
}

func (p *DeployVnfApplianceParams) GetAffinitygroupnames() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupnames"].([]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetBootintosetup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootintosetup"] = v
}

func (p *DeployVnfApplianceParams) ResetBootintosetup() {
	if p.p != nil && p.p["bootintosetup"] != nil {
		delete(p.p, "bootintosetup")
	}
}

func (p *DeployVnfApplianceParams) GetBootintosetup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootintosetup"].(bool)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetBootmode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootmode"] = v
}

func (p *DeployVnfApplianceParams) ResetBootmode() {
	if p.p != nil && p.p["bootmode"] != nil {
		delete(p.p, "bootmode")
	}
}

func (p *DeployVnfApplianceParams) GetBootmode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootmode"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetBoottype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["boottype"] = v
}

func (p *DeployVnfApplianceParams) ResetBoottype() {
	if p.p != nil && p.p["boottype"] != nil {
		delete(p.p, "boottype")
	}
}

func (p *DeployVnfApplianceParams) GetBoottype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["boottype"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetCopyimagetags(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["copyimagetags"] = v
}

func (p *DeployVnfApplianceParams) ResetCopyimagetags() {
	if p.p != nil && p.p["copyimagetags"] != nil {
		delete(p.p, "copyimagetags")
	}
}

func (p *DeployVnfApplianceParams) GetCopyimagetags() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["copyimagetags"].(bool)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *DeployVnfApplianceParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *DeployVnfApplianceParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDatadiskofferinglist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["datadiskofferinglist"] = v
}

func (p *DeployVnfApplianceParams) ResetDatadiskofferinglist() {
	if p.p != nil && p.p["datadiskofferinglist"] != nil {
		delete(p.p, "datadiskofferinglist")
	}
}

func (p *DeployVnfApplianceParams) GetDatadiskofferinglist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["datadiskofferinglist"].(map[string]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDeploymentplanner(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deploymentplanner"] = v
}

func (p *DeployVnfApplianceParams) ResetDeploymentplanner() {
	if p.p != nil && p.p["deploymentplanner"] != nil {
		delete(p.p, "deploymentplanner")
	}
}

func (p *DeployVnfApplianceParams) GetDeploymentplanner() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deploymentplanner"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *DeployVnfApplianceParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *DeployVnfApplianceParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDhcpoptionsnetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpoptionsnetworklist"] = v
}

func (p *DeployVnfApplianceParams) ResetDhcpoptionsnetworklist() {
	if p.p != nil && p.p["dhcpoptionsnetworklist"] != nil {
		delete(p.p, "dhcpoptionsnetworklist")
	}
}

func (p *DeployVnfApplianceParams) GetDhcpoptionsnetworklist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dhcpoptionsnetworklist"].(map[string]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *DeployVnfApplianceParams) ResetDiskofferingid() {
	if p.p != nil && p.p["diskofferingid"] != nil {
		delete(p.p, "diskofferingid")
	}
}

func (p *DeployVnfApplianceParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayname"] = v
}

func (p *DeployVnfApplianceParams) ResetDisplayname() {
	if p.p != nil && p.p["displayname"] != nil {
		delete(p.p, "displayname")
	}
}

func (p *DeployVnfApplianceParams) GetDisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayname"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
}

func (p *DeployVnfApplianceParams) ResetDisplayvm() {
	if p.p != nil && p.p["displayvm"] != nil {
		delete(p.p, "displayvm")
	}
}

func (p *DeployVnfApplianceParams) GetDisplayvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvm"].(bool)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *DeployVnfApplianceParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *DeployVnfApplianceParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetDynamicscalingenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dynamicscalingenabled"] = v
}

func (p *DeployVnfApplianceParams) ResetDynamicscalingenabled() {
	if p.p != nil && p.p["dynamicscalingenabled"] != nil {
		delete(p.p, "dynamicscalingenabled")
	}
}

func (p *DeployVnfApplianceParams) GetDynamicscalingenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dynamicscalingenabled"].(bool)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetExtraconfig(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extraconfig"] = v
}

func (p *DeployVnfApplianceParams) ResetExtraconfig() {
	if p.p != nil && p.p["extraconfig"] != nil {
		delete(p.p, "extraconfig")
	}
}

func (p *DeployVnfApplianceParams) GetExtraconfig() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extraconfig"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
}

func (p *DeployVnfApplianceParams) ResetGroup() {
	if p.p != nil && p.p["group"] != nil {
		delete(p.p, "group")
	}
}

func (p *DeployVnfApplianceParams) GetGroup() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["group"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *DeployVnfApplianceParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *DeployVnfApplianceParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *DeployVnfApplianceParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *DeployVnfApplianceParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetIodriverpolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iodriverpolicy"] = v
}

func (p *DeployVnfApplianceParams) ResetIodriverpolicy() {
	if p.p != nil && p.p["iodriverpolicy"] != nil {
		delete(p.p, "iodriverpolicy")
	}
}

func (p *DeployVnfApplianceParams) GetIodriverpolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iodriverpolicy"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetIothreadsenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iothreadsenabled"] = v
}

func (p *DeployVnfApplianceParams) ResetIothreadsenabled() {
	if p.p != nil && p.p["iothreadsenabled"] != nil {
		delete(p.p, "iothreadsenabled")
	}
}

func (p *DeployVnfApplianceParams) GetIothreadsenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iothreadsenabled"].(bool)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetIp6address(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6address"] = v
}

func (p *DeployVnfApplianceParams) ResetIp6address() {
	if p.p != nil && p.p["ip6address"] != nil {
		delete(p.p, "ip6address")
	}
}

func (p *DeployVnfApplianceParams) GetIp6address() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6address"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *DeployVnfApplianceParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *DeployVnfApplianceParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetIptonetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iptonetworklist"] = v
}

func (p *DeployVnfApplianceParams) ResetIptonetworklist() {
	if p.p != nil && p.p["iptonetworklist"] != nil {
		delete(p.p, "iptonetworklist")
	}
}

func (p *DeployVnfApplianceParams) GetIptonetworklist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iptonetworklist"].(map[string]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetKeyboard(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyboard"] = v
}

func (p *DeployVnfApplianceParams) ResetKeyboard() {
	if p.p != nil && p.p["keyboard"] != nil {
		delete(p.p, "keyboard")
	}
}

func (p *DeployVnfApplianceParams) GetKeyboard() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyboard"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
}

func (p *DeployVnfApplianceParams) ResetKeypair() {
	if p.p != nil && p.p["keypair"] != nil {
		delete(p.p, "keypair")
	}
}

func (p *DeployVnfApplianceParams) GetKeypair() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypair"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetKeypairs(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypairs"] = v
}

func (p *DeployVnfApplianceParams) ResetKeypairs() {
	if p.p != nil && p.p["keypairs"] != nil {
		delete(p.p, "keypairs")
	}
}

func (p *DeployVnfApplianceParams) GetKeypairs() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypairs"].([]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetMacaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["macaddress"] = v
}

func (p *DeployVnfApplianceParams) ResetMacaddress() {
	if p.p != nil && p.p["macaddress"] != nil {
		delete(p.p, "macaddress")
	}
}

func (p *DeployVnfApplianceParams) GetMacaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["macaddress"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *DeployVnfApplianceParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *DeployVnfApplianceParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetNetworkids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
}

func (p *DeployVnfApplianceParams) ResetNetworkids() {
	if p.p != nil && p.p["networkids"] != nil {
		delete(p.p, "networkids")
	}
}

func (p *DeployVnfApplianceParams) GetNetworkids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkids"].([]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetNicmultiqueuenumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicmultiqueuenumber"] = v
}

func (p *DeployVnfApplianceParams) ResetNicmultiqueuenumber() {
	if p.p != nil && p.p["nicmultiqueuenumber"] != nil {
		delete(p.p, "nicmultiqueuenumber")
	}
}

func (p *DeployVnfApplianceParams) GetNicmultiqueuenumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicmultiqueuenumber"].(int)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetNicnetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicnetworklist"] = v
}

func (p *DeployVnfApplianceParams) ResetNicnetworklist() {
	if p.p != nil && p.p["nicnetworklist"] != nil {
		delete(p.p, "nicnetworklist")
	}
}

func (p *DeployVnfApplianceParams) GetNicnetworklist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicnetworklist"].(map[string]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetNicpackedvirtqueuesenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicpackedvirtqueuesenabled"] = v
}

func (p *DeployVnfApplianceParams) ResetNicpackedvirtqueuesenabled() {
	if p.p != nil && p.p["nicpackedvirtqueuesenabled"] != nil {
		delete(p.p, "nicpackedvirtqueuesenabled")
	}
}

func (p *DeployVnfApplianceParams) GetNicpackedvirtqueuesenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicpackedvirtqueuesenabled"].(bool)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetOverridediskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["overridediskofferingid"] = v
}

func (p *DeployVnfApplianceParams) ResetOverridediskofferingid() {
	if p.p != nil && p.p["overridediskofferingid"] != nil {
		delete(p.p, "overridediskofferingid")
	}
}

func (p *DeployVnfApplianceParams) GetOverridediskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["overridediskofferingid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *DeployVnfApplianceParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *DeployVnfApplianceParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DeployVnfApplianceParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *DeployVnfApplianceParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetProperties(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["properties"] = v
}

func (p *DeployVnfApplianceParams) ResetProperties() {
	if p.p != nil && p.p["properties"] != nil {
		delete(p.p, "properties")
	}
}

func (p *DeployVnfApplianceParams) GetProperties() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["properties"].(map[string]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetRootdisksize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rootdisksize"] = v
}

func (p *DeployVnfApplianceParams) ResetRootdisksize() {
	if p.p != nil && p.p["rootdisksize"] != nil {
		delete(p.p, "rootdisksize")
	}
}

func (p *DeployVnfApplianceParams) GetRootdisksize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rootdisksize"].(int64)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetSecuritygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupids"] = v
}

func (p *DeployVnfApplianceParams) ResetSecuritygroupids() {
	if p.p != nil && p.p["securitygroupids"] != nil {
		delete(p.p, "securitygroupids")
	}
}

func (p *DeployVnfApplianceParams) GetSecuritygroupids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupids"].([]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetSecuritygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupnames"] = v
}

func (p *DeployVnfApplianceParams) ResetSecuritygroupnames() {
	if p.p != nil && p.p["securitygroupnames"] != nil {
		delete(p.p, "securitygroupnames")
	}
}

func (p *DeployVnfApplianceParams) GetSecuritygroupnames() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupnames"].([]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *DeployVnfApplianceParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *DeployVnfApplianceParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *DeployVnfApplianceParams) ResetSize() {
	if p.p != nil && p.p["size"] != nil {
		delete(p.p, "size")
	}
}

func (p *DeployVnfApplianceParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetStartvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startvm"] = v
}

func (p *DeployVnfApplianceParams) ResetStartvm() {
	if p.p != nil && p.p["startvm"] != nil {
		delete(p.p, "startvm")
	}
}

func (p *DeployVnfApplianceParams) GetStartvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startvm"].(bool)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *DeployVnfApplianceParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *DeployVnfApplianceParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *DeployVnfApplianceParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *DeployVnfApplianceParams) GetUserdata() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetUserdatadetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdatadetails"] = v
}

func (p *DeployVnfApplianceParams) ResetUserdatadetails() {
	if p.p != nil && p.p["userdatadetails"] != nil {
		delete(p.p, "userdatadetails")
	}
}

func (p *DeployVnfApplianceParams) GetUserdatadetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdatadetails"].(map[string]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetUserdataid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdataid"] = v
}

func (p *DeployVnfApplianceParams) ResetUserdataid() {
	if p.p != nil && p.p["userdataid"] != nil {
		delete(p.p, "userdataid")
	}
}

func (p *DeployVnfApplianceParams) GetUserdataid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdataid"].(string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetVnfcidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vnfcidrlist"] = v
}

func (p *DeployVnfApplianceParams) ResetVnfcidrlist() {
	if p.p != nil && p.p["vnfcidrlist"] != nil {
		delete(p.p, "vnfcidrlist")
	}
}

func (p *DeployVnfApplianceParams) GetVnfcidrlist() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vnfcidrlist"].([]string)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetVnfconfiguremanagement(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vnfconfiguremanagement"] = v
}

func (p *DeployVnfApplianceParams) ResetVnfconfiguremanagement() {
	if p.p != nil && p.p["vnfconfiguremanagement"] != nil {
		delete(p.p, "vnfconfiguremanagement")
	}
}

func (p *DeployVnfApplianceParams) GetVnfconfiguremanagement() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vnfconfiguremanagement"].(bool)
	return value, ok
}

func (p *DeployVnfApplianceParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeployVnfApplianceParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeployVnfApplianceParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeployVnfApplianceParams instance,
// as then you are sure you have configured all required params
func (s *VirtualNetworkFunctionsService) NewDeployVnfApplianceParams(serviceofferingid string, templateid string, zoneid string) *DeployVnfApplianceParams {
	p := &DeployVnfApplianceParams{}
	p.p = make(map[string]interface{})
	p.p["serviceofferingid"] = serviceofferingid
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	return p
}

// Creates and automatically starts a VNF appliance based on a service offering, disk offering, and template.
func (s *VirtualNetworkFunctionsService) DeployVnfAppliance(p *DeployVnfApplianceParams) (*DeployVnfApplianceResponse, error) {
	resp, err := s.cs.newRequest("deployVnfAppliance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeployVnfApplianceResponse
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

type DeployVnfApplianceResponse struct {
	Account               string                                    `json:"account"`
	Affinitygroup         []DeployVnfApplianceResponseAffinitygroup `json:"affinitygroup"`
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
	Securitygroup         []DeployVnfApplianceResponseSecuritygroup `json:"securitygroup"`
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

type DeployVnfApplianceResponseSecuritygroup struct {
	Account             string                                        `json:"account"`
	Description         string                                        `json:"description"`
	Domain              string                                        `json:"domain"`
	Domainid            string                                        `json:"domainid"`
	Domainpath          string                                        `json:"domainpath"`
	Egressrule          []DeployVnfApplianceResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                        `json:"id"`
	Ingressrule         []DeployVnfApplianceResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                        `json:"name"`
	Project             string                                        `json:"project"`
	Projectid           string                                        `json:"projectid"`
	Tags                []Tags                                        `json:"tags"`
	Virtualmachinecount int                                           `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                 `json:"virtualmachineids"`
}

type DeployVnfApplianceResponseSecuritygroupRule struct {
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

type DeployVnfApplianceResponseAffinitygroup struct {
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

func (r *DeployVnfApplianceResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeployVnfApplianceResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListVnfAppliancesParams struct {
	p map[string]interface{}
}

func (p *ListVnfAppliancesParams) toURLValues() url.Values {
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

func (p *ListVnfAppliancesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVnfAppliancesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVnfAppliancesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetAccumulate(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accumulate"] = v
}

func (p *ListVnfAppliancesParams) ResetAccumulate() {
	if p.p != nil && p.p["accumulate"] != nil {
		delete(p.p, "accumulate")
	}
}

func (p *ListVnfAppliancesParams) GetAccumulate() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accumulate"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetAffinitygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupid"] = v
}

func (p *ListVnfAppliancesParams) ResetAffinitygroupid() {
	if p.p != nil && p.p["affinitygroupid"] != nil {
		delete(p.p, "affinitygroupid")
	}
}

func (p *ListVnfAppliancesParams) GetAffinitygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetAutoscalevmgroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoscalevmgroupid"] = v
}

func (p *ListVnfAppliancesParams) ResetAutoscalevmgroupid() {
	if p.p != nil && p.p["autoscalevmgroupid"] != nil {
		delete(p.p, "autoscalevmgroupid")
	}
}

func (p *ListVnfAppliancesParams) GetAutoscalevmgroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoscalevmgroupid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetBackupofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["backupofferingid"] = v
}

func (p *ListVnfAppliancesParams) ResetBackupofferingid() {
	if p.p != nil && p.p["backupofferingid"] != nil {
		delete(p.p, "backupofferingid")
	}
}

func (p *ListVnfAppliancesParams) GetBackupofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["backupofferingid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListVnfAppliancesParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListVnfAppliancesParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
}

func (p *ListVnfAppliancesParams) ResetDisplayvm() {
	if p.p != nil && p.p["displayvm"] != nil {
		delete(p.p, "displayvm")
	}
}

func (p *ListVnfAppliancesParams) GetDisplayvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvm"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVnfAppliancesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVnfAppliancesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetForvirtualnetwork(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvirtualnetwork"] = v
}

func (p *ListVnfAppliancesParams) ResetForvirtualnetwork() {
	if p.p != nil && p.p["forvirtualnetwork"] != nil {
		delete(p.p, "forvirtualnetwork")
	}
}

func (p *ListVnfAppliancesParams) GetForvirtualnetwork() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forvirtualnetwork"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetGroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["groupid"] = v
}

func (p *ListVnfAppliancesParams) ResetGroupid() {
	if p.p != nil && p.p["groupid"] != nil {
		delete(p.p, "groupid")
	}
}

func (p *ListVnfAppliancesParams) GetGroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["groupid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetHaenable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["haenable"] = v
}

func (p *ListVnfAppliancesParams) ResetHaenable() {
	if p.p != nil && p.p["haenable"] != nil {
		delete(p.p, "haenable")
	}
}

func (p *ListVnfAppliancesParams) GetHaenable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["haenable"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListVnfAppliancesParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListVnfAppliancesParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVnfAppliancesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVnfAppliancesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListVnfAppliancesParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListVnfAppliancesParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetIsoid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isoid"] = v
}

func (p *ListVnfAppliancesParams) ResetIsoid() {
	if p.p != nil && p.p["isoid"] != nil {
		delete(p.p, "isoid")
	}
}

func (p *ListVnfAppliancesParams) GetIsoid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isoid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVnfAppliancesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVnfAppliancesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetIsvnf(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isvnf"] = v
}

func (p *ListVnfAppliancesParams) ResetIsvnf() {
	if p.p != nil && p.p["isvnf"] != nil {
		delete(p.p, "isvnf")
	}
}

func (p *ListVnfAppliancesParams) GetIsvnf() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isvnf"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
}

func (p *ListVnfAppliancesParams) ResetKeypair() {
	if p.p != nil && p.p["keypair"] != nil {
		delete(p.p, "keypair")
	}
}

func (p *ListVnfAppliancesParams) GetKeypair() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypair"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVnfAppliancesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVnfAppliancesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVnfAppliancesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVnfAppliancesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVnfAppliancesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVnfAppliancesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListVnfAppliancesParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListVnfAppliancesParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVnfAppliancesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVnfAppliancesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVnfAppliancesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVnfAppliancesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVnfAppliancesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVnfAppliancesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetRetrieveonlyresourcecount(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["retrieveonlyresourcecount"] = v
}

func (p *ListVnfAppliancesParams) ResetRetrieveonlyresourcecount() {
	if p.p != nil && p.p["retrieveonlyresourcecount"] != nil {
		delete(p.p, "retrieveonlyresourcecount")
	}
}

func (p *ListVnfAppliancesParams) GetRetrieveonlyresourcecount() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["retrieveonlyresourcecount"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetSecuritygroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupid"] = v
}

func (p *ListVnfAppliancesParams) ResetSecuritygroupid() {
	if p.p != nil && p.p["securitygroupid"] != nil {
		delete(p.p, "securitygroupid")
	}
}

func (p *ListVnfAppliancesParams) GetSecuritygroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ListVnfAppliancesParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ListVnfAppliancesParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListVnfAppliancesParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListVnfAppliancesParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVnfAppliancesParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListVnfAppliancesParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVnfAppliancesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListVnfAppliancesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *ListVnfAppliancesParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *ListVnfAppliancesParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetUserdata(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *ListVnfAppliancesParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *ListVnfAppliancesParams) GetUserdata() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(bool)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *ListVnfAppliancesParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *ListVnfAppliancesParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
}

func (p *ListVnfAppliancesParams) ResetVpcid() {
	if p.p != nil && p.p["vpcid"] != nil {
		delete(p.p, "vpcid")
	}
}

func (p *ListVnfAppliancesParams) GetVpcid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vpcid"].(string)
	return value, ok
}

func (p *ListVnfAppliancesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVnfAppliancesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVnfAppliancesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVnfAppliancesParams instance,
// as then you are sure you have configured all required params
func (s *VirtualNetworkFunctionsService) NewListVnfAppliancesParams() *ListVnfAppliancesParams {
	p := &ListVnfAppliancesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualNetworkFunctionsService) GetVnfApplianceID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVnfAppliancesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVnfAppliances(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VnfAppliances[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VnfAppliances {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualNetworkFunctionsService) GetVnfApplianceByName(name string, opts ...OptionFunc) (*VnfAppliance, int, error) {
	id, count, err := s.GetVnfApplianceID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVnfApplianceByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualNetworkFunctionsService) GetVnfApplianceByID(id string, opts ...OptionFunc) (*VnfAppliance, int, error) {
	p := &ListVnfAppliancesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVnfAppliances(p)
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
		return l.VnfAppliances[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VnfAppliance UUID: %s!", id)
}

// List VNF appliance owned by the account.
func (s *VirtualNetworkFunctionsService) ListVnfAppliances(p *ListVnfAppliancesParams) (*ListVnfAppliancesResponse, error) {
	resp, err := s.cs.newRequest("listVnfAppliances", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVnfAppliancesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVnfAppliancesResponse struct {
	Count         int             `json:"count"`
	VnfAppliances []*VnfAppliance `json:"vnfappliance"`
}

type VnfAppliance struct {
	Account               string                      `json:"account"`
	Affinitygroup         []VnfApplianceAffinitygroup `json:"affinitygroup"`
	Autoscalevmgroupid    string                      `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                      `json:"autoscalevmgroupname"`
	Backupofferingid      string                      `json:"backupofferingid"`
	Backupofferingname    string                      `json:"backupofferingname"`
	Bootmode              string                      `json:"bootmode"`
	Boottype              string                      `json:"boottype"`
	Cpunumber             int                         `json:"cpunumber"`
	Cpuspeed              int                         `json:"cpuspeed"`
	Cpuused               string                      `json:"cpuused"`
	Created               string                      `json:"created"`
	Deleteprotection      bool                        `json:"deleteprotection"`
	Details               map[string]string           `json:"details"`
	Diskioread            int64                       `json:"diskioread"`
	Diskiowrite           int64                       `json:"diskiowrite"`
	Diskkbsread           int64                       `json:"diskkbsread"`
	Diskkbswrite          int64                       `json:"diskkbswrite"`
	Diskofferingid        string                      `json:"diskofferingid"`
	Diskofferingname      string                      `json:"diskofferingname"`
	Displayname           string                      `json:"displayname"`
	Displayvm             bool                        `json:"displayvm"`
	Domain                string                      `json:"domain"`
	Domainid              string                      `json:"domainid"`
	Domainpath            string                      `json:"domainpath"`
	Forvirtualnetwork     bool                        `json:"forvirtualnetwork"`
	Group                 string                      `json:"group"`
	Groupid               string                      `json:"groupid"`
	Guestosid             string                      `json:"guestosid"`
	Haenable              bool                        `json:"haenable"`
	Hasannotations        bool                        `json:"hasannotations"`
	Hostcontrolstate      string                      `json:"hostcontrolstate"`
	Hostid                string                      `json:"hostid"`
	Hostname              string                      `json:"hostname"`
	Hypervisor            string                      `json:"hypervisor"`
	Icon                  interface{}                 `json:"icon"`
	Id                    string                      `json:"id"`
	Instancename          string                      `json:"instancename"`
	Ipaddress             string                      `json:"ipaddress"`
	Isdynamicallyscalable bool                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                      `json:"isodisplaytext"`
	Isoid                 string                      `json:"isoid"`
	Isoname               string                      `json:"isoname"`
	JobID                 string                      `json:"jobid"`
	Jobstatus             int                         `json:"jobstatus"`
	Keypairs              string                      `json:"keypairs"`
	Lastupdated           string                      `json:"lastupdated"`
	Memory                int                         `json:"memory"`
	Memoryintfreekbs      int64                       `json:"memoryintfreekbs"`
	Memorykbs             int64                       `json:"memorykbs"`
	Memorytargetkbs       int64                       `json:"memorytargetkbs"`
	Name                  string                      `json:"name"`
	Networkkbsread        int64                       `json:"networkkbsread"`
	Networkkbswrite       int64                       `json:"networkkbswrite"`
	Nic                   []Nic                       `json:"nic"`
	Osdisplayname         string                      `json:"osdisplayname"`
	Ostypeid              string                      `json:"ostypeid"`
	Password              string                      `json:"password"`
	Passwordenabled       bool                        `json:"passwordenabled"`
	Pooltype              string                      `json:"pooltype"`
	Project               string                      `json:"project"`
	Projectid             string                      `json:"projectid"`
	Publicip              string                      `json:"publicip"`
	Publicipid            string                      `json:"publicipid"`
	Readonlydetails       string                      `json:"readonlydetails"`
	Receivedbytes         int64                       `json:"receivedbytes"`
	Rootdeviceid          int64                       `json:"rootdeviceid"`
	Rootdevicetype        string                      `json:"rootdevicetype"`
	Securitygroup         []VnfApplianceSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                       `json:"sentbytes"`
	Serviceofferingid     string                      `json:"serviceofferingid"`
	Serviceofferingname   string                      `json:"serviceofferingname"`
	Servicestate          string                      `json:"servicestate"`
	State                 string                      `json:"state"`
	Tags                  []Tags                      `json:"tags"`
	Templatedisplaytext   string                      `json:"templatedisplaytext"`
	Templateformat        string                      `json:"templateformat"`
	Templateid            string                      `json:"templateid"`
	Templatename          string                      `json:"templatename"`
	Templatetype          string                      `json:"templatetype"`
	Userdata              string                      `json:"userdata"`
	Userdatadetails       string                      `json:"userdatadetails"`
	Userdataid            string                      `json:"userdataid"`
	Userdataname          string                      `json:"userdataname"`
	Userdatapolicy        string                      `json:"userdatapolicy"`
	Userid                string                      `json:"userid"`
	Username              string                      `json:"username"`
	Vgpu                  string                      `json:"vgpu"`
	Vmtype                string                      `json:"vmtype"`
	Vnfdetails            map[string]string           `json:"vnfdetails"`
	Vnfnics               []string                    `json:"vnfnics"`
	Zoneid                string                      `json:"zoneid"`
	Zonename              string                      `json:"zonename"`
}

type VnfApplianceSecuritygroup struct {
	Account             string                          `json:"account"`
	Description         string                          `json:"description"`
	Domain              string                          `json:"domain"`
	Domainid            string                          `json:"domainid"`
	Domainpath          string                          `json:"domainpath"`
	Egressrule          []VnfApplianceSecuritygroupRule `json:"egressrule"`
	Id                  string                          `json:"id"`
	Ingressrule         []VnfApplianceSecuritygroupRule `json:"ingressrule"`
	Name                string                          `json:"name"`
	Project             string                          `json:"project"`
	Projectid           string                          `json:"projectid"`
	Tags                []Tags                          `json:"tags"`
	Virtualmachinecount int                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                   `json:"virtualmachineids"`
}

type VnfApplianceSecuritygroupRule struct {
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

type VnfApplianceAffinitygroup struct {
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

func (r *VnfAppliance) UnmarshalJSON(b []byte) error {
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

	type alias VnfAppliance
	return json.Unmarshal(b, (*alias)(r))
}

type ListVnfTemplatesParams struct {
	p map[string]interface{}
}

func (p *ListVnfTemplatesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["isvnf"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isvnf", vv)
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
	if v, found := p.p["parenttemplateid"]; found {
		u.Set("parenttemplateid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := p.p["showremoved"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showremoved", vv)
	}
	if v, found := p.p["showunique"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showunique", vv)
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["templatefilter"]; found {
		u.Set("templatefilter", v.(string))
	}
	if v, found := p.p["templatetype"]; found {
		u.Set("templatetype", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVnfTemplatesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVnfTemplatesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVnfTemplatesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *ListVnfTemplatesParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *ListVnfTemplatesParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListVnfTemplatesParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListVnfTemplatesParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVnfTemplatesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVnfTemplatesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListVnfTemplatesParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListVnfTemplatesParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVnfTemplatesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListVnfTemplatesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListVnfTemplatesParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListVnfTemplatesParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVnfTemplatesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVnfTemplatesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetIsvnf(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isvnf"] = v
}

func (p *ListVnfTemplatesParams) ResetIsvnf() {
	if p.p != nil && p.p["isvnf"] != nil {
		delete(p.p, "isvnf")
	}
}

func (p *ListVnfTemplatesParams) GetIsvnf() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isvnf"].(bool)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVnfTemplatesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVnfTemplatesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVnfTemplatesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVnfTemplatesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVnfTemplatesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVnfTemplatesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVnfTemplatesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVnfTemplatesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVnfTemplatesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVnfTemplatesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetParenttemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parenttemplateid"] = v
}

func (p *ListVnfTemplatesParams) ResetParenttemplateid() {
	if p.p != nil && p.p["parenttemplateid"] != nil {
		delete(p.p, "parenttemplateid")
	}
}

func (p *ListVnfTemplatesParams) GetParenttemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parenttemplateid"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVnfTemplatesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVnfTemplatesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListVnfTemplatesParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListVnfTemplatesParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetShowremoved(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showremoved"] = v
}

func (p *ListVnfTemplatesParams) ResetShowremoved() {
	if p.p != nil && p.p["showremoved"] != nil {
		delete(p.p, "showremoved")
	}
}

func (p *ListVnfTemplatesParams) GetShowremoved() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showremoved"].(bool)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetShowunique(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showunique"] = v
}

func (p *ListVnfTemplatesParams) ResetShowunique() {
	if p.p != nil && p.p["showunique"] != nil {
		delete(p.p, "showunique")
	}
}

func (p *ListVnfTemplatesParams) GetShowunique() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showunique"].(bool)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVnfTemplatesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListVnfTemplatesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetTemplatefilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatefilter"] = v
}

func (p *ListVnfTemplatesParams) ResetTemplatefilter() {
	if p.p != nil && p.p["templatefilter"] != nil {
		delete(p.p, "templatefilter")
	}
}

func (p *ListVnfTemplatesParams) GetTemplatefilter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatefilter"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetTemplatetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetype"] = v
}

func (p *ListVnfTemplatesParams) ResetTemplatetype() {
	if p.p != nil && p.p["templatetype"] != nil {
		delete(p.p, "templatetype")
	}
}

func (p *ListVnfTemplatesParams) GetTemplatetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetype"].(string)
	return value, ok
}

func (p *ListVnfTemplatesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVnfTemplatesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVnfTemplatesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVnfTemplatesParams instance,
// as then you are sure you have configured all required params
func (s *VirtualNetworkFunctionsService) NewListVnfTemplatesParams(templatefilter string) *ListVnfTemplatesParams {
	p := &ListVnfTemplatesParams{}
	p.p = make(map[string]interface{})
	p.p["templatefilter"] = templatefilter
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualNetworkFunctionsService) GetVnfTemplateID(name string, templatefilter string, opts ...OptionFunc) (string, int, error) {
	p := &ListVnfTemplatesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name
	p.p["templatefilter"] = templatefilter

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVnfTemplates(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VnfTemplates[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VnfTemplates {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualNetworkFunctionsService) GetVnfTemplateByName(name string, templatefilter string, opts ...OptionFunc) (*VnfTemplate, int, error) {
	id, count, err := s.GetVnfTemplateID(name, templatefilter, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVnfTemplateByID(id, templatefilter, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualNetworkFunctionsService) GetVnfTemplateByID(id string, templatefilter string, opts ...OptionFunc) (*VnfTemplate, int, error) {
	p := &ListVnfTemplatesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id
	p.p["templatefilter"] = templatefilter

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVnfTemplates(p)
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
		return l.VnfTemplates[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VnfTemplate UUID: %s!", id)
}

// List all public, private, and privileged VNF templates.
func (s *VirtualNetworkFunctionsService) ListVnfTemplates(p *ListVnfTemplatesParams) (*ListVnfTemplatesResponse, error) {
	resp, err := s.cs.newRequest("listVnfTemplates", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVnfTemplatesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVnfTemplatesResponse struct {
	Count        int            `json:"count"`
	VnfTemplates []*VnfTemplate `json:"vnftemplate"`
}

type VnfTemplate struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Arch                  string              `json:"arch"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Domainpath            string              `json:"domainpath"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  interface{}         `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Userdataid            string              `json:"userdataid"`
	Userdataname          string              `json:"userdataname"`
	Userdataparams        string              `json:"userdataparams"`
	Userdatapolicy        string              `json:"userdatapolicy"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *VnfTemplate) UnmarshalJSON(b []byte) error {
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

	type alias VnfTemplate
	return json.Unmarshal(b, (*alias)(r))
}

type RegisterVnfTemplateParams struct {
	p map[string]interface{}
}

func (p *RegisterVnfTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["bits"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("bits", vv)
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := p.p["deployasis"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("deployasis", vv)
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["directdownload"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("directdownload", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["isextractable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isextractable", vv)
	}
	if v, found := p.p["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["isrouting"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrouting", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := p.p["sshkeyenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sshkeyenabled", vv)
	}
	if v, found := p.p["templatetag"]; found {
		u.Set("templatetag", v.(string))
	}
	if v, found := p.p["templatetype"]; found {
		u.Set("templatetype", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["vnfdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("vnfdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("vnfdetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["vnfnics"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("vnfnics[%d].key", i), k)
			u.Set(fmt.Sprintf("vnfnics[%d].value", i), m[k])
		}
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	if v, found := p.p["zoneids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneids", vv)
	}
	return u
}

func (p *RegisterVnfTemplateParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *RegisterVnfTemplateParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *RegisterVnfTemplateParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *RegisterVnfTemplateParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *RegisterVnfTemplateParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
}

func (p *RegisterVnfTemplateParams) ResetBits() {
	if p.p != nil && p.p["bits"] != nil {
		delete(p.p, "bits")
	}
}

func (p *RegisterVnfTemplateParams) GetBits() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bits"].(int)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
}

func (p *RegisterVnfTemplateParams) ResetChecksum() {
	if p.p != nil && p.p["checksum"] != nil {
		delete(p.p, "checksum")
	}
}

func (p *RegisterVnfTemplateParams) GetChecksum() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["checksum"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetDeployasis(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deployasis"] = v
}

func (p *RegisterVnfTemplateParams) ResetDeployasis() {
	if p.p != nil && p.p["deployasis"] != nil {
		delete(p.p, "deployasis")
	}
}

func (p *RegisterVnfTemplateParams) GetDeployasis() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deployasis"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *RegisterVnfTemplateParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *RegisterVnfTemplateParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetDirectdownload(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["directdownload"] = v
}

func (p *RegisterVnfTemplateParams) ResetDirectdownload() {
	if p.p != nil && p.p["directdownload"] != nil {
		delete(p.p, "directdownload")
	}
}

func (p *RegisterVnfTemplateParams) GetDirectdownload() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["directdownload"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *RegisterVnfTemplateParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *RegisterVnfTemplateParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *RegisterVnfTemplateParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *RegisterVnfTemplateParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *RegisterVnfTemplateParams) ResetFormat() {
	if p.p != nil && p.p["format"] != nil {
		delete(p.p, "format")
	}
}

func (p *RegisterVnfTemplateParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *RegisterVnfTemplateParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *RegisterVnfTemplateParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *RegisterVnfTemplateParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *RegisterVnfTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
}

func (p *RegisterVnfTemplateParams) ResetIsextractable() {
	if p.p != nil && p.p["isextractable"] != nil {
		delete(p.p, "isextractable")
	}
}

func (p *RegisterVnfTemplateParams) GetIsextractable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isextractable"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
}

func (p *RegisterVnfTemplateParams) ResetIsfeatured() {
	if p.p != nil && p.p["isfeatured"] != nil {
		delete(p.p, "isfeatured")
	}
}

func (p *RegisterVnfTemplateParams) GetIsfeatured() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isfeatured"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *RegisterVnfTemplateParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *RegisterVnfTemplateParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
}

func (p *RegisterVnfTemplateParams) ResetIsrouting() {
	if p.p != nil && p.p["isrouting"] != nil {
		delete(p.p, "isrouting")
	}
}

func (p *RegisterVnfTemplateParams) GetIsrouting() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrouting"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *RegisterVnfTemplateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *RegisterVnfTemplateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *RegisterVnfTemplateParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *RegisterVnfTemplateParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
}

func (p *RegisterVnfTemplateParams) ResetPasswordenabled() {
	if p.p != nil && p.p["passwordenabled"] != nil {
		delete(p.p, "passwordenabled")
	}
}

func (p *RegisterVnfTemplateParams) GetPasswordenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passwordenabled"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *RegisterVnfTemplateParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *RegisterVnfTemplateParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
}

func (p *RegisterVnfTemplateParams) ResetRequireshvm() {
	if p.p != nil && p.p["requireshvm"] != nil {
		delete(p.p, "requireshvm")
	}
}

func (p *RegisterVnfTemplateParams) GetRequireshvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["requireshvm"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetSshkeyenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sshkeyenabled"] = v
}

func (p *RegisterVnfTemplateParams) ResetSshkeyenabled() {
	if p.p != nil && p.p["sshkeyenabled"] != nil {
		delete(p.p, "sshkeyenabled")
	}
}

func (p *RegisterVnfTemplateParams) GetSshkeyenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sshkeyenabled"].(bool)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetTemplatetag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetag"] = v
}

func (p *RegisterVnfTemplateParams) ResetTemplatetag() {
	if p.p != nil && p.p["templatetag"] != nil {
		delete(p.p, "templatetag")
	}
}

func (p *RegisterVnfTemplateParams) GetTemplatetag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetag"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetTemplatetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetype"] = v
}

func (p *RegisterVnfTemplateParams) ResetTemplatetype() {
	if p.p != nil && p.p["templatetype"] != nil {
		delete(p.p, "templatetype")
	}
}

func (p *RegisterVnfTemplateParams) GetTemplatetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetype"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *RegisterVnfTemplateParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *RegisterVnfTemplateParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetVnfdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vnfdetails"] = v
}

func (p *RegisterVnfTemplateParams) ResetVnfdetails() {
	if p.p != nil && p.p["vnfdetails"] != nil {
		delete(p.p, "vnfdetails")
	}
}

func (p *RegisterVnfTemplateParams) GetVnfdetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vnfdetails"].(map[string]string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetVnfnics(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vnfnics"] = v
}

func (p *RegisterVnfTemplateParams) ResetVnfnics() {
	if p.p != nil && p.p["vnfnics"] != nil {
		delete(p.p, "vnfnics")
	}
}

func (p *RegisterVnfTemplateParams) GetVnfnics() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vnfnics"].(map[string]string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *RegisterVnfTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *RegisterVnfTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

func (p *RegisterVnfTemplateParams) SetZoneids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneids"] = v
}

func (p *RegisterVnfTemplateParams) ResetZoneids() {
	if p.p != nil && p.p["zoneids"] != nil {
		delete(p.p, "zoneids")
	}
}

func (p *RegisterVnfTemplateParams) GetZoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneids"].([]string)
	return value, ok
}

// You should always use this function to get a new RegisterVnfTemplateParams instance,
// as then you are sure you have configured all required params
func (s *VirtualNetworkFunctionsService) NewRegisterVnfTemplateParams(format string, hypervisor string, name string, url string) *RegisterVnfTemplateParams {
	p := &RegisterVnfTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["format"] = format
	p.p["hypervisor"] = hypervisor
	p.p["name"] = name
	p.p["url"] = url
	return p
}

// Registers an existing VNF template into the CloudStack cloud.
func (s *VirtualNetworkFunctionsService) RegisterVnfTemplate(p *RegisterVnfTemplateParams) (*RegisterVnfTemplateResponse, error) {
	resp, err := s.cs.newRequest("registerVnfTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterVnfTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterVnfTemplateResponse struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Arch                  string              `json:"arch"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Domainpath            string              `json:"domainpath"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  interface{}         `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Userdataid            string              `json:"userdataid"`
	Userdataname          string              `json:"userdataname"`
	Userdataparams        string              `json:"userdataparams"`
	Userdatapolicy        string              `json:"userdatapolicy"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *RegisterVnfTemplateResponse) UnmarshalJSON(b []byte) error {
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

	type alias RegisterVnfTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateVnfTemplateParams struct {
	p map[string]interface{}
}

func (p *UpdateVnfTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["arch"]; found {
		u.Set("arch", v.(string))
	}
	if v, found := p.p["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := p.p["cleanupdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupdetails", vv)
	}
	if v, found := p.p["cleanupvnfdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupvnfdetails", vv)
	}
	if v, found := p.p["cleanupvnfnics"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupvnfnics", vv)
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := p.p["isrouting"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrouting", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := p.p["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	if v, found := p.p["sshkeyenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sshkeyenabled", vv)
	}
	if v, found := p.p["templatetag"]; found {
		u.Set("templatetag", v.(string))
	}
	if v, found := p.p["templatetype"]; found {
		u.Set("templatetype", v.(string))
	}
	if v, found := p.p["vnfdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("vnfdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("vnfdetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["vnfnics"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("vnfnics[%d].key", i), k)
			u.Set(fmt.Sprintf("vnfnics[%d].value", i), m[k])
		}
	}
	return u
}

func (p *UpdateVnfTemplateParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *UpdateVnfTemplateParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *UpdateVnfTemplateParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
}

func (p *UpdateVnfTemplateParams) ResetBootable() {
	if p.p != nil && p.p["bootable"] != nil {
		delete(p.p, "bootable")
	}
}

func (p *UpdateVnfTemplateParams) GetBootable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootable"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetCleanupdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupdetails"] = v
}

func (p *UpdateVnfTemplateParams) ResetCleanupdetails() {
	if p.p != nil && p.p["cleanupdetails"] != nil {
		delete(p.p, "cleanupdetails")
	}
}

func (p *UpdateVnfTemplateParams) GetCleanupdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupdetails"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetCleanupvnfdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupvnfdetails"] = v
}

func (p *UpdateVnfTemplateParams) ResetCleanupvnfdetails() {
	if p.p != nil && p.p["cleanupvnfdetails"] != nil {
		delete(p.p, "cleanupvnfdetails")
	}
}

func (p *UpdateVnfTemplateParams) GetCleanupvnfdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupvnfdetails"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetCleanupvnfnics(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupvnfnics"] = v
}

func (p *UpdateVnfTemplateParams) ResetCleanupvnfnics() {
	if p.p != nil && p.p["cleanupvnfnics"] != nil {
		delete(p.p, "cleanupvnfnics")
	}
}

func (p *UpdateVnfTemplateParams) GetCleanupvnfnics() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupvnfnics"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateVnfTemplateParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateVnfTemplateParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateVnfTemplateParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *UpdateVnfTemplateParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *UpdateVnfTemplateParams) ResetFormat() {
	if p.p != nil && p.p["format"] != nil {
		delete(p.p, "format")
	}
}

func (p *UpdateVnfTemplateParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVnfTemplateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateVnfTemplateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *UpdateVnfTemplateParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *UpdateVnfTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
}

func (p *UpdateVnfTemplateParams) ResetIsrouting() {
	if p.p != nil && p.p["isrouting"] != nil {
		delete(p.p, "isrouting")
	}
}

func (p *UpdateVnfTemplateParams) GetIsrouting() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrouting"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVnfTemplateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateVnfTemplateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *UpdateVnfTemplateParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *UpdateVnfTemplateParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
}

func (p *UpdateVnfTemplateParams) ResetPasswordenabled() {
	if p.p != nil && p.p["passwordenabled"] != nil {
		delete(p.p, "passwordenabled")
	}
}

func (p *UpdateVnfTemplateParams) GetPasswordenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passwordenabled"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
}

func (p *UpdateVnfTemplateParams) ResetRequireshvm() {
	if p.p != nil && p.p["requireshvm"] != nil {
		delete(p.p, "requireshvm")
	}
}

func (p *UpdateVnfTemplateParams) GetRequireshvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["requireshvm"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

func (p *UpdateVnfTemplateParams) ResetSortkey() {
	if p.p != nil && p.p["sortkey"] != nil {
		delete(p.p, "sortkey")
	}
}

func (p *UpdateVnfTemplateParams) GetSortkey() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortkey"].(int)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetSshkeyenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sshkeyenabled"] = v
}

func (p *UpdateVnfTemplateParams) ResetSshkeyenabled() {
	if p.p != nil && p.p["sshkeyenabled"] != nil {
		delete(p.p, "sshkeyenabled")
	}
}

func (p *UpdateVnfTemplateParams) GetSshkeyenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sshkeyenabled"].(bool)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetTemplatetag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetag"] = v
}

func (p *UpdateVnfTemplateParams) ResetTemplatetag() {
	if p.p != nil && p.p["templatetag"] != nil {
		delete(p.p, "templatetag")
	}
}

func (p *UpdateVnfTemplateParams) GetTemplatetag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetag"].(string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetTemplatetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetype"] = v
}

func (p *UpdateVnfTemplateParams) ResetTemplatetype() {
	if p.p != nil && p.p["templatetype"] != nil {
		delete(p.p, "templatetype")
	}
}

func (p *UpdateVnfTemplateParams) GetTemplatetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetype"].(string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetVnfdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vnfdetails"] = v
}

func (p *UpdateVnfTemplateParams) ResetVnfdetails() {
	if p.p != nil && p.p["vnfdetails"] != nil {
		delete(p.p, "vnfdetails")
	}
}

func (p *UpdateVnfTemplateParams) GetVnfdetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vnfdetails"].(map[string]string)
	return value, ok
}

func (p *UpdateVnfTemplateParams) SetVnfnics(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vnfnics"] = v
}

func (p *UpdateVnfTemplateParams) ResetVnfnics() {
	if p.p != nil && p.p["vnfnics"] != nil {
		delete(p.p, "vnfnics")
	}
}

func (p *UpdateVnfTemplateParams) GetVnfnics() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vnfnics"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new UpdateVnfTemplateParams instance,
// as then you are sure you have configured all required params
func (s *VirtualNetworkFunctionsService) NewUpdateVnfTemplateParams(id string) *UpdateVnfTemplateParams {
	p := &UpdateVnfTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a template to VNF template or attributes of a VNF template.
func (s *VirtualNetworkFunctionsService) UpdateVnfTemplate(p *UpdateVnfTemplateParams) (*UpdateVnfTemplateResponse, error) {
	resp, err := s.cs.newRequest("updateVnfTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVnfTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateVnfTemplateResponse struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Arch                  string              `json:"arch"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Domainpath            string              `json:"domainpath"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  interface{}         `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Userdataid            string              `json:"userdataid"`
	Userdataname          string              `json:"userdataname"`
	Userdataparams        string              `json:"userdataparams"`
	Userdatapolicy        string              `json:"userdatapolicy"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *UpdateVnfTemplateResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateVnfTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}
