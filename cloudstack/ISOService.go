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

type ISOServiceIface interface {
	AttachIso(p *AttachIsoParams) (*AttachIsoResponse, error)
	NewAttachIsoParams(id string, virtualmachineid string) *AttachIsoParams
	CopyIso(p *CopyIsoParams) (*CopyIsoResponse, error)
	NewCopyIsoParams(id string) *CopyIsoParams
	DeleteIso(p *DeleteIsoParams) (*DeleteIsoResponse, error)
	NewDeleteIsoParams(id string) *DeleteIsoParams
	DetachIso(p *DetachIsoParams) (*DetachIsoResponse, error)
	NewDetachIsoParams(virtualmachineid string) *DetachIsoParams
	ExtractIso(p *ExtractIsoParams) (*ExtractIsoResponse, error)
	NewExtractIsoParams(id string, mode string) *ExtractIsoParams
	GetUploadParamsForIso(p *GetUploadParamsForIsoParams) (*GetUploadParamsForIsoResponse, error)
	NewGetUploadParamsForIsoParams(format string, name string, zoneid string) *GetUploadParamsForIsoParams
	ListIsoPermissions(p *ListIsoPermissionsParams) (*ListIsoPermissionsResponse, error)
	NewListIsoPermissionsParams(id string) *ListIsoPermissionsParams
	GetIsoPermissionByID(id string, opts ...OptionFunc) (*IsoPermission, int, error)
	ListIsos(p *ListIsosParams) (*ListIsosResponse, error)
	NewListIsosParams() *ListIsosParams
	GetIsoID(name string, isofilter string, zoneid string, opts ...OptionFunc) (string, int, error)
	GetIsoByName(name string, isofilter string, zoneid string, opts ...OptionFunc) (*Iso, int, error)
	GetIsoByID(id string, opts ...OptionFunc) (*Iso, int, error)
	RegisterIso(p *RegisterIsoParams) (*RegisterIsoResponse, error)
	NewRegisterIsoParams(displaytext string, name string, url string, zoneid string) *RegisterIsoParams
	UpdateIso(p *UpdateIsoParams) (*UpdateIsoResponse, error)
	NewUpdateIsoParams(id string) *UpdateIsoParams
	UpdateIsoPermissions(p *UpdateIsoPermissionsParams) (*UpdateIsoPermissionsResponse, error)
	NewUpdateIsoPermissionsParams(id string) *UpdateIsoPermissionsParams
}

type AttachIsoParams struct {
	p map[string]interface{}
}

func (p *AttachIsoParams) toURLValues() url.Values {
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
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AttachIsoParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *AttachIsoParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *AttachIsoParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *AttachIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *AttachIsoParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *AttachIsoParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *AttachIsoParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *AttachIsoParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *AttachIsoParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new AttachIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewAttachIsoParams(id string, virtualmachineid string) *AttachIsoParams {
	p := &AttachIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attaches an ISO to a virtual machine.
func (s *ISOService) AttachIso(p *AttachIsoParams) (*AttachIsoResponse, error) {
	resp, err := s.cs.newPostRequest("attachIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AttachIsoResponse
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

type AttachIsoResponse struct {
	Account               string                           `json:"account"`
	Affinitygroup         []AttachIsoResponseAffinitygroup `json:"affinitygroup"`
	Arch                  string                           `json:"arch"`
	Autoscalevmgroupid    string                           `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                           `json:"autoscalevmgroupname"`
	Backupofferingid      string                           `json:"backupofferingid"`
	Backupofferingname    string                           `json:"backupofferingname"`
	Bootmode              string                           `json:"bootmode"`
	Boottype              string                           `json:"boottype"`
	Cpunumber             int                              `json:"cpunumber"`
	Cpuspeed              int                              `json:"cpuspeed"`
	Cpuused               string                           `json:"cpuused"`
	Created               string                           `json:"created"`
	Deleteprotection      bool                             `json:"deleteprotection"`
	Details               map[string]string                `json:"details"`
	Diskioread            int64                            `json:"diskioread"`
	Diskiowrite           int64                            `json:"diskiowrite"`
	Diskkbsread           int64                            `json:"diskkbsread"`
	Diskkbswrite          int64                            `json:"diskkbswrite"`
	Diskofferingid        string                           `json:"diskofferingid"`
	Diskofferingname      string                           `json:"diskofferingname"`
	Displayname           string                           `json:"displayname"`
	Displayvm             bool                             `json:"displayvm"`
	Domain                string                           `json:"domain"`
	Domainid              string                           `json:"domainid"`
	Domainpath            string                           `json:"domainpath"`
	Forvirtualnetwork     bool                             `json:"forvirtualnetwork"`
	Gpucardid             string                           `json:"gpucardid"`
	Gpucardname           string                           `json:"gpucardname"`
	Gpucount              int                              `json:"gpucount"`
	Group                 string                           `json:"group"`
	Groupid               string                           `json:"groupid"`
	Guestosid             string                           `json:"guestosid"`
	Haenable              bool                             `json:"haenable"`
	Hasannotations        bool                             `json:"hasannotations"`
	Hostcontrolstate      string                           `json:"hostcontrolstate"`
	Hostid                string                           `json:"hostid"`
	Hostname              string                           `json:"hostname"`
	Hypervisor            string                           `json:"hypervisor"`
	Icon                  interface{}                      `json:"icon"`
	Id                    string                           `json:"id"`
	Instancename          string                           `json:"instancename"`
	Ipaddress             string                           `json:"ipaddress"`
	Isdynamicallyscalable bool                             `json:"isdynamicallyscalable"`
	Isodisplaytext        string                           `json:"isodisplaytext"`
	Isoid                 string                           `json:"isoid"`
	Isoname               string                           `json:"isoname"`
	JobID                 string                           `json:"jobid"`
	Jobstatus             int                              `json:"jobstatus"`
	Keypairs              string                           `json:"keypairs"`
	Lastupdated           string                           `json:"lastupdated"`
	Leaseduration         int                              `json:"leaseduration"`
	Leaseexpiryaction     string                           `json:"leaseexpiryaction"`
	Leaseexpirydate       string                           `json:"leaseexpirydate"`
	Maxheads              int64                            `json:"maxheads"`
	Maxresolutionx        int64                            `json:"maxresolutionx"`
	Maxresolutiony        int64                            `json:"maxresolutiony"`
	Memory                int                              `json:"memory"`
	Memoryintfreekbs      int64                            `json:"memoryintfreekbs"`
	Memorykbs             int64                            `json:"memorykbs"`
	Memorytargetkbs       int64                            `json:"memorytargetkbs"`
	Name                  string                           `json:"name"`
	Networkkbsread        int64                            `json:"networkkbsread"`
	Networkkbswrite       int64                            `json:"networkkbswrite"`
	Nic                   []Nic                            `json:"nic"`
	Osdisplayname         string                           `json:"osdisplayname"`
	Ostypeid              string                           `json:"ostypeid"`
	Password              string                           `json:"password"`
	Passwordenabled       bool                             `json:"passwordenabled"`
	Pooltype              string                           `json:"pooltype"`
	Project               string                           `json:"project"`
	Projectid             string                           `json:"projectid"`
	Publicip              string                           `json:"publicip"`
	Publicipid            string                           `json:"publicipid"`
	Readonlydetails       string                           `json:"readonlydetails"`
	Receivedbytes         int64                            `json:"receivedbytes"`
	Rootdeviceid          int64                            `json:"rootdeviceid"`
	Rootdevicetype        string                           `json:"rootdevicetype"`
	Securitygroup         []AttachIsoResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                            `json:"sentbytes"`
	Serviceofferingid     string                           `json:"serviceofferingid"`
	Serviceofferingname   string                           `json:"serviceofferingname"`
	Servicestate          string                           `json:"servicestate"`
	State                 string                           `json:"state"`
	Tags                  []Tags                           `json:"tags"`
	Templatedisplaytext   string                           `json:"templatedisplaytext"`
	Templateformat        string                           `json:"templateformat"`
	Templateid            string                           `json:"templateid"`
	Templatename          string                           `json:"templatename"`
	Templatetype          string                           `json:"templatetype"`
	Userdata              string                           `json:"userdata"`
	Userdatadetails       string                           `json:"userdatadetails"`
	Userdataid            string                           `json:"userdataid"`
	Userdataname          string                           `json:"userdataname"`
	Userdatapolicy        string                           `json:"userdatapolicy"`
	Userid                string                           `json:"userid"`
	Username              string                           `json:"username"`
	Vgpu                  string                           `json:"vgpu"`
	Vgpuprofileid         string                           `json:"vgpuprofileid"`
	Vgpuprofilename       string                           `json:"vgpuprofilename"`
	Videoram              int64                            `json:"videoram"`
	Vmtype                string                           `json:"vmtype"`
	Vnfdetails            map[string]string                `json:"vnfdetails"`
	Vnfnics               []string                         `json:"vnfnics"`
	Zoneid                string                           `json:"zoneid"`
	Zonename              string                           `json:"zonename"`
}

type AttachIsoResponseSecuritygroup struct {
	Account             string                               `json:"account"`
	Description         string                               `json:"description"`
	Domain              string                               `json:"domain"`
	Domainid            string                               `json:"domainid"`
	Domainpath          string                               `json:"domainpath"`
	Egressrule          []AttachIsoResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                               `json:"id"`
	Ingressrule         []AttachIsoResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                               `json:"name"`
	Project             string                               `json:"project"`
	Projectid           string                               `json:"projectid"`
	Tags                []Tags                               `json:"tags"`
	Virtualmachinecount int                                  `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                        `json:"virtualmachineids"`
}

type AttachIsoResponseSecuritygroupRule struct {
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

type AttachIsoResponseAffinitygroup struct {
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

func (r *AttachIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias AttachIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CopyIsoParams struct {
	p map[string]interface{}
}

func (p *CopyIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["destzoneid"]; found {
		u.Set("destzoneid", v.(string))
	}
	if v, found := p.p["destzoneids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("destzoneids", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["sourcezoneid"]; found {
		u.Set("sourcezoneid", v.(string))
	}
	return u
}

func (p *CopyIsoParams) SetDestzoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destzoneid"] = v
}

func (p *CopyIsoParams) ResetDestzoneid() {
	if p.p != nil && p.p["destzoneid"] != nil {
		delete(p.p, "destzoneid")
	}
}

func (p *CopyIsoParams) GetDestzoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destzoneid"].(string)
	return value, ok
}

func (p *CopyIsoParams) SetDestzoneids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destzoneids"] = v
}

func (p *CopyIsoParams) ResetDestzoneids() {
	if p.p != nil && p.p["destzoneids"] != nil {
		delete(p.p, "destzoneids")
	}
}

func (p *CopyIsoParams) GetDestzoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destzoneids"].([]string)
	return value, ok
}

func (p *CopyIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *CopyIsoParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *CopyIsoParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *CopyIsoParams) SetSourcezoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcezoneid"] = v
}

func (p *CopyIsoParams) ResetSourcezoneid() {
	if p.p != nil && p.p["sourcezoneid"] != nil {
		delete(p.p, "sourcezoneid")
	}
}

func (p *CopyIsoParams) GetSourcezoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcezoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CopyIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewCopyIsoParams(id string) *CopyIsoParams {
	p := &CopyIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Copies an iso from one zone to another.
func (s *ISOService) CopyIso(p *CopyIsoParams) (*CopyIsoResponse, error) {
	resp, err := s.cs.newPostRequest("copyIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CopyIsoResponse
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

type CopyIsoResponse struct {
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
	Extensionid           string              `json:"extensionid"`
	Extensionname         string              `json:"extensionname"`
	Forcks                bool                `json:"forcks"`
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

func (r *CopyIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias CopyIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteIsoParams struct {
	p map[string]interface{}
}

func (p *DeleteIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteIsoParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteIsoParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteIsoParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteIsoParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteIsoParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewDeleteIsoParams(id string) *DeleteIsoParams {
	p := &DeleteIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an ISO file.
func (s *ISOService) DeleteIso(p *DeleteIsoParams) (*DeleteIsoResponse, error) {
	resp, err := s.cs.newPostRequest("deleteIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteIsoResponse
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

type DeleteIsoResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DetachIsoParams struct {
	p map[string]interface{}
}

func (p *DetachIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *DetachIsoParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DetachIsoParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *DetachIsoParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *DetachIsoParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *DetachIsoParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *DetachIsoParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new DetachIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewDetachIsoParams(virtualmachineid string) *DetachIsoParams {
	p := &DetachIsoParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Detaches any ISO file (if any) currently attached to a virtual machine.
func (s *ISOService) DetachIso(p *DetachIsoParams) (*DetachIsoResponse, error) {
	resp, err := s.cs.newPostRequest("detachIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DetachIsoResponse
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

type DetachIsoResponse struct {
	Account               string                           `json:"account"`
	Affinitygroup         []DetachIsoResponseAffinitygroup `json:"affinitygroup"`
	Arch                  string                           `json:"arch"`
	Autoscalevmgroupid    string                           `json:"autoscalevmgroupid"`
	Autoscalevmgroupname  string                           `json:"autoscalevmgroupname"`
	Backupofferingid      string                           `json:"backupofferingid"`
	Backupofferingname    string                           `json:"backupofferingname"`
	Bootmode              string                           `json:"bootmode"`
	Boottype              string                           `json:"boottype"`
	Cpunumber             int                              `json:"cpunumber"`
	Cpuspeed              int                              `json:"cpuspeed"`
	Cpuused               string                           `json:"cpuused"`
	Created               string                           `json:"created"`
	Deleteprotection      bool                             `json:"deleteprotection"`
	Details               map[string]string                `json:"details"`
	Diskioread            int64                            `json:"diskioread"`
	Diskiowrite           int64                            `json:"diskiowrite"`
	Diskkbsread           int64                            `json:"diskkbsread"`
	Diskkbswrite          int64                            `json:"diskkbswrite"`
	Diskofferingid        string                           `json:"diskofferingid"`
	Diskofferingname      string                           `json:"diskofferingname"`
	Displayname           string                           `json:"displayname"`
	Displayvm             bool                             `json:"displayvm"`
	Domain                string                           `json:"domain"`
	Domainid              string                           `json:"domainid"`
	Domainpath            string                           `json:"domainpath"`
	Forvirtualnetwork     bool                             `json:"forvirtualnetwork"`
	Gpucardid             string                           `json:"gpucardid"`
	Gpucardname           string                           `json:"gpucardname"`
	Gpucount              int                              `json:"gpucount"`
	Group                 string                           `json:"group"`
	Groupid               string                           `json:"groupid"`
	Guestosid             string                           `json:"guestosid"`
	Haenable              bool                             `json:"haenable"`
	Hasannotations        bool                             `json:"hasannotations"`
	Hostcontrolstate      string                           `json:"hostcontrolstate"`
	Hostid                string                           `json:"hostid"`
	Hostname              string                           `json:"hostname"`
	Hypervisor            string                           `json:"hypervisor"`
	Icon                  interface{}                      `json:"icon"`
	Id                    string                           `json:"id"`
	Instancename          string                           `json:"instancename"`
	Ipaddress             string                           `json:"ipaddress"`
	Isdynamicallyscalable bool                             `json:"isdynamicallyscalable"`
	Isodisplaytext        string                           `json:"isodisplaytext"`
	Isoid                 string                           `json:"isoid"`
	Isoname               string                           `json:"isoname"`
	JobID                 string                           `json:"jobid"`
	Jobstatus             int                              `json:"jobstatus"`
	Keypairs              string                           `json:"keypairs"`
	Lastupdated           string                           `json:"lastupdated"`
	Leaseduration         int                              `json:"leaseduration"`
	Leaseexpiryaction     string                           `json:"leaseexpiryaction"`
	Leaseexpirydate       string                           `json:"leaseexpirydate"`
	Maxheads              int64                            `json:"maxheads"`
	Maxresolutionx        int64                            `json:"maxresolutionx"`
	Maxresolutiony        int64                            `json:"maxresolutiony"`
	Memory                int                              `json:"memory"`
	Memoryintfreekbs      int64                            `json:"memoryintfreekbs"`
	Memorykbs             int64                            `json:"memorykbs"`
	Memorytargetkbs       int64                            `json:"memorytargetkbs"`
	Name                  string                           `json:"name"`
	Networkkbsread        int64                            `json:"networkkbsread"`
	Networkkbswrite       int64                            `json:"networkkbswrite"`
	Nic                   []Nic                            `json:"nic"`
	Osdisplayname         string                           `json:"osdisplayname"`
	Ostypeid              string                           `json:"ostypeid"`
	Password              string                           `json:"password"`
	Passwordenabled       bool                             `json:"passwordenabled"`
	Pooltype              string                           `json:"pooltype"`
	Project               string                           `json:"project"`
	Projectid             string                           `json:"projectid"`
	Publicip              string                           `json:"publicip"`
	Publicipid            string                           `json:"publicipid"`
	Readonlydetails       string                           `json:"readonlydetails"`
	Receivedbytes         int64                            `json:"receivedbytes"`
	Rootdeviceid          int64                            `json:"rootdeviceid"`
	Rootdevicetype        string                           `json:"rootdevicetype"`
	Securitygroup         []DetachIsoResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                            `json:"sentbytes"`
	Serviceofferingid     string                           `json:"serviceofferingid"`
	Serviceofferingname   string                           `json:"serviceofferingname"`
	Servicestate          string                           `json:"servicestate"`
	State                 string                           `json:"state"`
	Tags                  []Tags                           `json:"tags"`
	Templatedisplaytext   string                           `json:"templatedisplaytext"`
	Templateformat        string                           `json:"templateformat"`
	Templateid            string                           `json:"templateid"`
	Templatename          string                           `json:"templatename"`
	Templatetype          string                           `json:"templatetype"`
	Userdata              string                           `json:"userdata"`
	Userdatadetails       string                           `json:"userdatadetails"`
	Userdataid            string                           `json:"userdataid"`
	Userdataname          string                           `json:"userdataname"`
	Userdatapolicy        string                           `json:"userdatapolicy"`
	Userid                string                           `json:"userid"`
	Username              string                           `json:"username"`
	Vgpu                  string                           `json:"vgpu"`
	Vgpuprofileid         string                           `json:"vgpuprofileid"`
	Vgpuprofilename       string                           `json:"vgpuprofilename"`
	Videoram              int64                            `json:"videoram"`
	Vmtype                string                           `json:"vmtype"`
	Vnfdetails            map[string]string                `json:"vnfdetails"`
	Vnfnics               []string                         `json:"vnfnics"`
	Zoneid                string                           `json:"zoneid"`
	Zonename              string                           `json:"zonename"`
}

type DetachIsoResponseSecuritygroup struct {
	Account             string                               `json:"account"`
	Description         string                               `json:"description"`
	Domain              string                               `json:"domain"`
	Domainid            string                               `json:"domainid"`
	Domainpath          string                               `json:"domainpath"`
	Egressrule          []DetachIsoResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                               `json:"id"`
	Ingressrule         []DetachIsoResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                               `json:"name"`
	Project             string                               `json:"project"`
	Projectid           string                               `json:"projectid"`
	Tags                []Tags                               `json:"tags"`
	Virtualmachinecount int                                  `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                        `json:"virtualmachineids"`
}

type DetachIsoResponseSecuritygroupRule struct {
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

type DetachIsoResponseAffinitygroup struct {
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

func (r *DetachIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias DetachIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ExtractIsoParams struct {
	p map[string]interface{}
}

func (p *ExtractIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["mode"]; found {
		u.Set("mode", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ExtractIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ExtractIsoParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ExtractIsoParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ExtractIsoParams) SetMode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mode"] = v
}

func (p *ExtractIsoParams) ResetMode() {
	if p.p != nil && p.p["mode"] != nil {
		delete(p.p, "mode")
	}
}

func (p *ExtractIsoParams) GetMode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mode"].(string)
	return value, ok
}

func (p *ExtractIsoParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *ExtractIsoParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *ExtractIsoParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *ExtractIsoParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ExtractIsoParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ExtractIsoParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ExtractIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewExtractIsoParams(id string, mode string) *ExtractIsoParams {
	p := &ExtractIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["mode"] = mode
	return p
}

// Extracts an ISO
func (s *ISOService) ExtractIso(p *ExtractIsoParams) (*ExtractIsoResponse, error) {
	resp, err := s.cs.newPostRequest("extractIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractIsoResponse
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

type ExtractIsoResponse struct {
	Accountid        string `json:"accountid"`
	Created          string `json:"created"`
	ExtractId        string `json:"extractId"`
	ExtractMode      string `json:"extractMode"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Resultstring     string `json:"resultstring"`
	State            string `json:"state"`
	Status           string `json:"status"`
	Storagetype      string `json:"storagetype"`
	Uploadpercentage int    `json:"uploadpercentage"`
	Url              string `json:"url"`
	Zoneid           string `json:"zoneid"`
	Zonename         string `json:"zonename"`
}

type GetUploadParamsForIsoParams struct {
	p map[string]interface{}
}

func (p *GetUploadParamsForIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *GetUploadParamsForIsoParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *GetUploadParamsForIsoParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *GetUploadParamsForIsoParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
}

func (p *GetUploadParamsForIsoParams) ResetBootable() {
	if p.p != nil && p.p["bootable"] != nil {
		delete(p.p, "bootable")
	}
}

func (p *GetUploadParamsForIsoParams) GetBootable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootable"].(bool)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
}

func (p *GetUploadParamsForIsoParams) ResetChecksum() {
	if p.p != nil && p.p["checksum"] != nil {
		delete(p.p, "checksum")
	}
}

func (p *GetUploadParamsForIsoParams) GetChecksum() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["checksum"].(string)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *GetUploadParamsForIsoParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *GetUploadParamsForIsoParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *GetUploadParamsForIsoParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *GetUploadParamsForIsoParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *GetUploadParamsForIsoParams) ResetFormat() {
	if p.p != nil && p.p["format"] != nil {
		delete(p.p, "format")
	}
}

func (p *GetUploadParamsForIsoParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
}

func (p *GetUploadParamsForIsoParams) ResetIsextractable() {
	if p.p != nil && p.p["isextractable"] != nil {
		delete(p.p, "isextractable")
	}
}

func (p *GetUploadParamsForIsoParams) GetIsextractable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isextractable"].(bool)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
}

func (p *GetUploadParamsForIsoParams) ResetIsfeatured() {
	if p.p != nil && p.p["isfeatured"] != nil {
		delete(p.p, "isfeatured")
	}
}

func (p *GetUploadParamsForIsoParams) GetIsfeatured() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isfeatured"].(bool)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *GetUploadParamsForIsoParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *GetUploadParamsForIsoParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *GetUploadParamsForIsoParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *GetUploadParamsForIsoParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *GetUploadParamsForIsoParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *GetUploadParamsForIsoParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *GetUploadParamsForIsoParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *GetUploadParamsForIsoParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *GetUploadParamsForIsoParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *GetUploadParamsForIsoParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *GetUploadParamsForIsoParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new GetUploadParamsForIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewGetUploadParamsForIsoParams(format string, name string, zoneid string) *GetUploadParamsForIsoParams {
	p := &GetUploadParamsForIsoParams{}
	p.p = make(map[string]interface{})
	p.p["format"] = format
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// upload an existing ISO into the CloudStack cloud.
func (s *ISOService) GetUploadParamsForIso(p *GetUploadParamsForIsoParams) (*GetUploadParamsForIsoResponse, error) {
	resp, err := s.cs.newRequest("getUploadParamsForIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetUploadParamsForIsoResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetUploadParamsForIsoResponse struct {
	Expires   string `json:"expires"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Metadata  string `json:"metadata"`
	PostURL   string `json:"postURL"`
	Signature string `json:"signature"`
}

type ListIsoPermissionsParams struct {
	p map[string]interface{}
}

func (p *ListIsoPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ListIsoPermissionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListIsoPermissionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListIsoPermissionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ListIsoPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewListIsoPermissionsParams(id string) *ListIsoPermissionsParams {
	p := &ListIsoPermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoPermissionByID(id string, opts ...OptionFunc) (*IsoPermission, int, error) {
	p := &ListIsoPermissionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListIsoPermissions(p)
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
		return l.IsoPermissions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for IsoPermission UUID: %s!", id)
}

// List iso visibility and all accounts that have permissions to view this iso.
func (s *ISOService) ListIsoPermissions(p *ListIsoPermissionsParams) (*ListIsoPermissionsResponse, error) {
	resp, err := s.cs.newRequest("listIsoPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIsoPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListIsoPermissionsResponse struct {
	Count          int              `json:"count"`
	IsoPermissions []*IsoPermission `json:"isopermission"`
}

type IsoPermission struct {
	Account    []string `json:"account"`
	Domainid   string   `json:"domainid"`
	Id         string   `json:"id"`
	Ispublic   bool     `json:"ispublic"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Projectids []string `json:"projectids"`
}

type ListIsosParams struct {
	p map[string]interface{}
}

func (p *ListIsosParams) toURLValues() url.Values {
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
	if v, found := p.p["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
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
	if v, found := p.p["imagestoreid"]; found {
		u.Set("imagestoreid", v.(string))
	}
	if v, found := p.p["isofilter"]; found {
		u.Set("isofilter", v.(string))
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["isready"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isready", vv)
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
	if v, found := p.p["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListIsosParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListIsosParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListIsosParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListIsosParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *ListIsosParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *ListIsosParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *ListIsosParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
}

func (p *ListIsosParams) ResetBootable() {
	if p.p != nil && p.p["bootable"] != nil {
		delete(p.p, "bootable")
	}
}

func (p *ListIsosParams) GetBootable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootable"].(bool)
	return value, ok
}

func (p *ListIsosParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListIsosParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListIsosParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListIsosParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListIsosParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListIsosParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListIsosParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListIsosParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListIsosParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListIsosParams) SetImagestoreid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["imagestoreid"] = v
}

func (p *ListIsosParams) ResetImagestoreid() {
	if p.p != nil && p.p["imagestoreid"] != nil {
		delete(p.p, "imagestoreid")
	}
}

func (p *ListIsosParams) GetImagestoreid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["imagestoreid"].(string)
	return value, ok
}

func (p *ListIsosParams) SetIsofilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isofilter"] = v
}

func (p *ListIsosParams) ResetIsofilter() {
	if p.p != nil && p.p["isofilter"] != nil {
		delete(p.p, "isofilter")
	}
}

func (p *ListIsosParams) GetIsofilter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isofilter"].(string)
	return value, ok
}

func (p *ListIsosParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *ListIsosParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *ListIsosParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *ListIsosParams) SetIsready(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isready"] = v
}

func (p *ListIsosParams) ResetIsready() {
	if p.p != nil && p.p["isready"] != nil {
		delete(p.p, "isready")
	}
}

func (p *ListIsosParams) GetIsready() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isready"].(bool)
	return value, ok
}

func (p *ListIsosParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListIsosParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListIsosParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListIsosParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListIsosParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListIsosParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListIsosParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListIsosParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListIsosParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListIsosParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListIsosParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListIsosParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListIsosParams) SetOscategoryid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["oscategoryid"] = v
}

func (p *ListIsosParams) ResetOscategoryid() {
	if p.p != nil && p.p["oscategoryid"] != nil {
		delete(p.p, "oscategoryid")
	}
}

func (p *ListIsosParams) GetOscategoryid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["oscategoryid"].(string)
	return value, ok
}

func (p *ListIsosParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListIsosParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListIsosParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListIsosParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListIsosParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListIsosParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListIsosParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListIsosParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListIsosParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListIsosParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListIsosParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListIsosParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListIsosParams) SetShowremoved(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showremoved"] = v
}

func (p *ListIsosParams) ResetShowremoved() {
	if p.p != nil && p.p["showremoved"] != nil {
		delete(p.p, "showremoved")
	}
}

func (p *ListIsosParams) GetShowremoved() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showremoved"].(bool)
	return value, ok
}

func (p *ListIsosParams) SetShowunique(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showunique"] = v
}

func (p *ListIsosParams) ResetShowunique() {
	if p.p != nil && p.p["showunique"] != nil {
		delete(p.p, "showunique")
	}
}

func (p *ListIsosParams) GetShowunique() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showunique"].(bool)
	return value, ok
}

func (p *ListIsosParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListIsosParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *ListIsosParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListIsosParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListIsosParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListIsosParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListIsosParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListIsosParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListIsosParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListIsosParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewListIsosParams() *ListIsosParams {
	p := &ListIsosParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoID(name string, isofilter string, zoneid string, opts ...OptionFunc) (string, int, error) {
	p := &ListIsosParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name
	p.p["isofilter"] = isofilter
	p.p["zoneid"] = zoneid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListIsos(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Isos[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Isos {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoByName(name string, isofilter string, zoneid string, opts ...OptionFunc) (*Iso, int, error) {
	id, count, err := s.GetIsoID(name, isofilter, zoneid, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetIsoByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoByID(id string, opts ...OptionFunc) (*Iso, int, error) {
	p := &ListIsosParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListIsos(p)
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
		return l.Isos[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Iso UUID: %s!", id)
}

// Lists all available ISO files.
func (s *ISOService) ListIsos(p *ListIsosParams) (*ListIsosResponse, error) {
	resp, err := s.cs.newRequest("listIsos", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIsosResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListIsosResponse struct {
	Count int    `json:"count"`
	Isos  []*Iso `json:"iso"`
}

type Iso struct {
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
	Extensionid           string              `json:"extensionid"`
	Extensionname         string              `json:"extensionname"`
	Forcks                bool                `json:"forcks"`
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

func (r *Iso) UnmarshalJSON(b []byte) error {
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

	type alias Iso
	return json.Unmarshal(b, (*alias)(r))
}

type RegisterIsoParams struct {
	p map[string]interface{}
}

func (p *RegisterIsoParams) toURLValues() url.Values {
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
	if v, found := p.p["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
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
	if v, found := p.p["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
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
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *RegisterIsoParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *RegisterIsoParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *RegisterIsoParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *RegisterIsoParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *RegisterIsoParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
}

func (p *RegisterIsoParams) ResetBootable() {
	if p.p != nil && p.p["bootable"] != nil {
		delete(p.p, "bootable")
	}
}

func (p *RegisterIsoParams) GetBootable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootable"].(bool)
	return value, ok
}

func (p *RegisterIsoParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
}

func (p *RegisterIsoParams) ResetChecksum() {
	if p.p != nil && p.p["checksum"] != nil {
		delete(p.p, "checksum")
	}
}

func (p *RegisterIsoParams) GetChecksum() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["checksum"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetDirectdownload(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["directdownload"] = v
}

func (p *RegisterIsoParams) ResetDirectdownload() {
	if p.p != nil && p.p["directdownload"] != nil {
		delete(p.p, "directdownload")
	}
}

func (p *RegisterIsoParams) GetDirectdownload() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["directdownload"].(bool)
	return value, ok
}

func (p *RegisterIsoParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *RegisterIsoParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *RegisterIsoParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *RegisterIsoParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *RegisterIsoParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetImagestoreuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["imagestoreuuid"] = v
}

func (p *RegisterIsoParams) ResetImagestoreuuid() {
	if p.p != nil && p.p["imagestoreuuid"] != nil {
		delete(p.p, "imagestoreuuid")
	}
}

func (p *RegisterIsoParams) GetImagestoreuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["imagestoreuuid"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *RegisterIsoParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *RegisterIsoParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *RegisterIsoParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
}

func (p *RegisterIsoParams) ResetIsextractable() {
	if p.p != nil && p.p["isextractable"] != nil {
		delete(p.p, "isextractable")
	}
}

func (p *RegisterIsoParams) GetIsextractable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isextractable"].(bool)
	return value, ok
}

func (p *RegisterIsoParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
}

func (p *RegisterIsoParams) ResetIsfeatured() {
	if p.p != nil && p.p["isfeatured"] != nil {
		delete(p.p, "isfeatured")
	}
}

func (p *RegisterIsoParams) GetIsfeatured() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isfeatured"].(bool)
	return value, ok
}

func (p *RegisterIsoParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *RegisterIsoParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *RegisterIsoParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *RegisterIsoParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *RegisterIsoParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *RegisterIsoParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *RegisterIsoParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *RegisterIsoParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
}

func (p *RegisterIsoParams) ResetPasswordenabled() {
	if p.p != nil && p.p["passwordenabled"] != nil {
		delete(p.p, "passwordenabled")
	}
}

func (p *RegisterIsoParams) GetPasswordenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passwordenabled"].(bool)
	return value, ok
}

func (p *RegisterIsoParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *RegisterIsoParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *RegisterIsoParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *RegisterIsoParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *RegisterIsoParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *RegisterIsoParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *RegisterIsoParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *RegisterIsoParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewRegisterIsoParams(displaytext string, name string, url string, zoneid string) *RegisterIsoParams {
	p := &RegisterIsoParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Registers an existing ISO into the CloudStack Cloud.
func (s *ISOService) RegisterIso(p *RegisterIsoParams) (*RegisterIsoResponse, error) {
	resp, err := s.cs.newPostRequest("registerIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r RegisterIsoResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterIsoResponse struct {
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
	Extensionid           string              `json:"extensionid"`
	Extensionname         string              `json:"extensionname"`
	Forcks                bool                `json:"forcks"`
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

func (r *RegisterIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias RegisterIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateIsoParams struct {
	p map[string]interface{}
}

func (p *UpdateIsoParams) toURLValues() url.Values {
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
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["forceupdateostype"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forceupdateostype", vv)
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
	return u
}

func (p *UpdateIsoParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *UpdateIsoParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *UpdateIsoParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *UpdateIsoParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
}

func (p *UpdateIsoParams) ResetBootable() {
	if p.p != nil && p.p["bootable"] != nil {
		delete(p.p, "bootable")
	}
}

func (p *UpdateIsoParams) GetBootable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootable"].(bool)
	return value, ok
}

func (p *UpdateIsoParams) SetCleanupdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupdetails"] = v
}

func (p *UpdateIsoParams) ResetCleanupdetails() {
	if p.p != nil && p.p["cleanupdetails"] != nil {
		delete(p.p, "cleanupdetails")
	}
}

func (p *UpdateIsoParams) GetCleanupdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupdetails"].(bool)
	return value, ok
}

func (p *UpdateIsoParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateIsoParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateIsoParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateIsoParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateIsoParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *UpdateIsoParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateIsoParams) SetForceupdateostype(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forceupdateostype"] = v
}

func (p *UpdateIsoParams) ResetForceupdateostype() {
	if p.p != nil && p.p["forceupdateostype"] != nil {
		delete(p.p, "forceupdateostype")
	}
}

func (p *UpdateIsoParams) GetForceupdateostype() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forceupdateostype"].(bool)
	return value, ok
}

func (p *UpdateIsoParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *UpdateIsoParams) ResetFormat() {
	if p.p != nil && p.p["format"] != nil {
		delete(p.p, "format")
	}
}

func (p *UpdateIsoParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *UpdateIsoParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateIsoParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateIsoParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateIsoParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *UpdateIsoParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *UpdateIsoParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *UpdateIsoParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
}

func (p *UpdateIsoParams) ResetIsrouting() {
	if p.p != nil && p.p["isrouting"] != nil {
		delete(p.p, "isrouting")
	}
}

func (p *UpdateIsoParams) GetIsrouting() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrouting"].(bool)
	return value, ok
}

func (p *UpdateIsoParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateIsoParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateIsoParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateIsoParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *UpdateIsoParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *UpdateIsoParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *UpdateIsoParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
}

func (p *UpdateIsoParams) ResetPasswordenabled() {
	if p.p != nil && p.p["passwordenabled"] != nil {
		delete(p.p, "passwordenabled")
	}
}

func (p *UpdateIsoParams) GetPasswordenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passwordenabled"].(bool)
	return value, ok
}

func (p *UpdateIsoParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
}

func (p *UpdateIsoParams) ResetRequireshvm() {
	if p.p != nil && p.p["requireshvm"] != nil {
		delete(p.p, "requireshvm")
	}
}

func (p *UpdateIsoParams) GetRequireshvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["requireshvm"].(bool)
	return value, ok
}

func (p *UpdateIsoParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

func (p *UpdateIsoParams) ResetSortkey() {
	if p.p != nil && p.p["sortkey"] != nil {
		delete(p.p, "sortkey")
	}
}

func (p *UpdateIsoParams) GetSortkey() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortkey"].(int)
	return value, ok
}

func (p *UpdateIsoParams) SetSshkeyenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sshkeyenabled"] = v
}

func (p *UpdateIsoParams) ResetSshkeyenabled() {
	if p.p != nil && p.p["sshkeyenabled"] != nil {
		delete(p.p, "sshkeyenabled")
	}
}

func (p *UpdateIsoParams) GetSshkeyenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sshkeyenabled"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewUpdateIsoParams(id string) *UpdateIsoParams {
	p := &UpdateIsoParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an ISO file.
func (s *ISOService) UpdateIso(p *UpdateIsoParams) (*UpdateIsoResponse, error) {
	resp, err := s.cs.newPostRequest("updateIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIsoResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateIsoResponse struct {
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
	Extensionid           string              `json:"extensionid"`
	Extensionname         string              `json:"extensionname"`
	Forcks                bool                `json:"forcks"`
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

func (r *UpdateIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateIsoPermissionsParams struct {
	p map[string]interface{}
}

func (p *UpdateIsoPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accounts"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("accounts", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["op"]; found {
		u.Set("op", v.(string))
	}
	if v, found := p.p["projectids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("projectids", vv)
	}
	return u
}

func (p *UpdateIsoPermissionsParams) SetAccounts(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounts"] = v
}

func (p *UpdateIsoPermissionsParams) ResetAccounts() {
	if p.p != nil && p.p["accounts"] != nil {
		delete(p.p, "accounts")
	}
}

func (p *UpdateIsoPermissionsParams) GetAccounts() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounts"].([]string)
	return value, ok
}

func (p *UpdateIsoPermissionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateIsoPermissionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateIsoPermissionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateIsoPermissionsParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
}

func (p *UpdateIsoPermissionsParams) ResetIsextractable() {
	if p.p != nil && p.p["isextractable"] != nil {
		delete(p.p, "isextractable")
	}
}

func (p *UpdateIsoPermissionsParams) GetIsextractable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isextractable"].(bool)
	return value, ok
}

func (p *UpdateIsoPermissionsParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
}

func (p *UpdateIsoPermissionsParams) ResetIsfeatured() {
	if p.p != nil && p.p["isfeatured"] != nil {
		delete(p.p, "isfeatured")
	}
}

func (p *UpdateIsoPermissionsParams) GetIsfeatured() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isfeatured"].(bool)
	return value, ok
}

func (p *UpdateIsoPermissionsParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *UpdateIsoPermissionsParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *UpdateIsoPermissionsParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *UpdateIsoPermissionsParams) SetOp(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["op"] = v
}

func (p *UpdateIsoPermissionsParams) ResetOp() {
	if p.p != nil && p.p["op"] != nil {
		delete(p.p, "op")
	}
}

func (p *UpdateIsoPermissionsParams) GetOp() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["op"].(string)
	return value, ok
}

func (p *UpdateIsoPermissionsParams) SetProjectids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectids"] = v
}

func (p *UpdateIsoPermissionsParams) ResetProjectids() {
	if p.p != nil && p.p["projectids"] != nil {
		delete(p.p, "projectids")
	}
}

func (p *UpdateIsoPermissionsParams) GetProjectids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectids"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateIsoPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewUpdateIsoPermissionsParams(id string) *UpdateIsoPermissionsParams {
	p := &UpdateIsoPermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates ISO permissions
func (s *ISOService) UpdateIsoPermissions(p *UpdateIsoPermissionsParams) (*UpdateIsoPermissionsResponse, error) {
	resp, err := s.cs.newPostRequest("updateIsoPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIsoPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateIsoPermissionsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateIsoPermissionsResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateIsoPermissionsResponse
	return json.Unmarshal(b, (*alias)(r))
}
