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

type TemplateServiceIface interface {
	CopyTemplate(p *CopyTemplateParams) (*CopyTemplateResponse, error)
	NewCopyTemplateParams(id string) *CopyTemplateParams
	CreateTemplate(p *CreateTemplateParams) (*CreateTemplateResponse, error)
	NewCreateTemplateParams(displaytext string, name string, ostypeid string) *CreateTemplateParams
	DeleteTemplate(p *DeleteTemplateParams) (*DeleteTemplateResponse, error)
	NewDeleteTemplateParams(id string) *DeleteTemplateParams
	ExtractTemplate(p *ExtractTemplateParams) (*ExtractTemplateResponse, error)
	NewExtractTemplateParams(id string, mode string) *ExtractTemplateParams
	GetUploadParamsForTemplate(p *GetUploadParamsForTemplateParams) (*GetUploadParamsForTemplateResponse, error)
	NewGetUploadParamsForTemplateParams(displaytext string, format string, hypervisor string, name string, zoneid string) *GetUploadParamsForTemplateParams
	ListTemplatePermissions(p *ListTemplatePermissionsParams) (*ListTemplatePermissionsResponse, error)
	NewListTemplatePermissionsParams(id string) *ListTemplatePermissionsParams
	GetTemplatePermissionByID(id string, opts ...OptionFunc) (*TemplatePermission, int, error)
	ListTemplates(p *ListTemplatesParams) (*ListTemplatesResponse, error)
	NewListTemplatesParams(templatefilter string) *ListTemplatesParams
	GetTemplateID(name string, templatefilter string, zoneid string, opts ...OptionFunc) (string, int, error)
	GetTemplateByName(name string, templatefilter string, zoneid string, opts ...OptionFunc) (*Template, int, error)
	GetTemplateByID(id string, templatefilter string, opts ...OptionFunc) (*Template, int, error)
	PrepareTemplate(p *PrepareTemplateParams) (*PrepareTemplateResponse, error)
	NewPrepareTemplateParams(templateid string, zoneid string) *PrepareTemplateParams
	RegisterTemplate(p *RegisterTemplateParams) (*RegisterTemplateResponse, error)
	NewRegisterTemplateParams(displaytext string, format string, hypervisor string, name string, url string) *RegisterTemplateParams
	UpdateTemplate(p *UpdateTemplateParams) (*UpdateTemplateResponse, error)
	NewUpdateTemplateParams(id string) *UpdateTemplateParams
	UpdateTemplatePermissions(p *UpdateTemplatePermissionsParams) (*UpdateTemplatePermissionsResponse, error)
	NewUpdateTemplatePermissionsParams(id string) *UpdateTemplatePermissionsParams
	UpgradeRouterTemplate(p *UpgradeRouterTemplateParams) (*UpgradeRouterTemplateResponse, error)
	NewUpgradeRouterTemplateParams() *UpgradeRouterTemplateParams
	LinkUserDataToTemplate(p *LinkUserDataToTemplateParams) (*LinkUserDataToTemplateResponse, error)
	NewLinkUserDataToTemplateParams() *LinkUserDataToTemplateParams
}

type CopyTemplateParams struct {
	p map[string]interface{}
}

func (p *CopyTemplateParams) toURLValues() url.Values {
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

func (p *CopyTemplateParams) SetDestzoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destzoneid"] = v
}

func (p *CopyTemplateParams) ResetDestzoneid() {
	if p.p != nil && p.p["destzoneid"] != nil {
		delete(p.p, "destzoneid")
	}
}

func (p *CopyTemplateParams) GetDestzoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destzoneid"].(string)
	return value, ok
}

func (p *CopyTemplateParams) SetDestzoneids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destzoneids"] = v
}

func (p *CopyTemplateParams) ResetDestzoneids() {
	if p.p != nil && p.p["destzoneids"] != nil {
		delete(p.p, "destzoneids")
	}
}

func (p *CopyTemplateParams) GetDestzoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destzoneids"].([]string)
	return value, ok
}

func (p *CopyTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *CopyTemplateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *CopyTemplateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *CopyTemplateParams) SetSourcezoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcezoneid"] = v
}

func (p *CopyTemplateParams) ResetSourcezoneid() {
	if p.p != nil && p.p["sourcezoneid"] != nil {
		delete(p.p, "sourcezoneid")
	}
}

func (p *CopyTemplateParams) GetSourcezoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcezoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CopyTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewCopyTemplateParams(id string) *CopyTemplateParams {
	p := &CopyTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Copies a template from one zone to another.
func (s *TemplateService) CopyTemplate(p *CopyTemplateParams) (*CopyTemplateResponse, error) {
	resp, err := s.cs.newRequest("copyTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CopyTemplateResponse
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

type CopyTemplateResponse struct {
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

func (r *CopyTemplateResponse) UnmarshalJSON(b []byte) error {
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

	type alias CopyTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CreateTemplateParams struct {
	p map[string]interface{}
}

func (p *CreateTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["bits"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("bits", vv)
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
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
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
	if v, found := p.p["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := p.p["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	if v, found := p.p["sshkeyenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sshkeyenabled", vv)
	}
	if v, found := p.p["templatetag"]; found {
		u.Set("templatetag", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateTemplateParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateTemplateParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateTemplateParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
}

func (p *CreateTemplateParams) ResetBits() {
	if p.p != nil && p.p["bits"] != nil {
		delete(p.p, "bits")
	}
}

func (p *CreateTemplateParams) GetBits() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bits"].(int)
	return value, ok
}

func (p *CreateTemplateParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateTemplateParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *CreateTemplateParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *CreateTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateTemplateParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *CreateTemplateParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateTemplateParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateTemplateParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *CreateTemplateParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *CreateTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *CreateTemplateParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
}

func (p *CreateTemplateParams) ResetIsfeatured() {
	if p.p != nil && p.p["isfeatured"] != nil {
		delete(p.p, "isfeatured")
	}
}

func (p *CreateTemplateParams) GetIsfeatured() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isfeatured"].(bool)
	return value, ok
}

func (p *CreateTemplateParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *CreateTemplateParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *CreateTemplateParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *CreateTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateTemplateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateTemplateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *CreateTemplateParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *CreateTemplateParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
}

func (p *CreateTemplateParams) ResetPasswordenabled() {
	if p.p != nil && p.p["passwordenabled"] != nil {
		delete(p.p, "passwordenabled")
	}
}

func (p *CreateTemplateParams) GetPasswordenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passwordenabled"].(bool)
	return value, ok
}

func (p *CreateTemplateParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateTemplateParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateTemplateParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
}

func (p *CreateTemplateParams) ResetRequireshvm() {
	if p.p != nil && p.p["requireshvm"] != nil {
		delete(p.p, "requireshvm")
	}
}

func (p *CreateTemplateParams) GetRequireshvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["requireshvm"].(bool)
	return value, ok
}

func (p *CreateTemplateParams) SetSnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshotid"] = v
}

func (p *CreateTemplateParams) ResetSnapshotid() {
	if p.p != nil && p.p["snapshotid"] != nil {
		delete(p.p, "snapshotid")
	}
}

func (p *CreateTemplateParams) GetSnapshotid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["snapshotid"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetSshkeyenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sshkeyenabled"] = v
}

func (p *CreateTemplateParams) ResetSshkeyenabled() {
	if p.p != nil && p.p["sshkeyenabled"] != nil {
		delete(p.p, "sshkeyenabled")
	}
}

func (p *CreateTemplateParams) GetSshkeyenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sshkeyenabled"].(bool)
	return value, ok
}

func (p *CreateTemplateParams) SetTemplatetag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetag"] = v
}

func (p *CreateTemplateParams) ResetTemplatetag() {
	if p.p != nil && p.p["templatetag"] != nil {
		delete(p.p, "templatetag")
	}
}

func (p *CreateTemplateParams) GetTemplatetag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetag"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *CreateTemplateParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *CreateTemplateParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *CreateTemplateParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *CreateTemplateParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *CreateTemplateParams) ResetVolumeid() {
	if p.p != nil && p.p["volumeid"] != nil {
		delete(p.p, "volumeid")
	}
}

func (p *CreateTemplateParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewCreateTemplateParams(displaytext string, name string, ostypeid string) *CreateTemplateParams {
	p := &CreateTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	p.p["ostypeid"] = ostypeid
	return p
}

// Creates a template of a virtual machine. The virtual machine must be in a STOPPED state. A template created from this command is automatically designated as a private template visible to the account that created it.
func (s *TemplateService) CreateTemplate(p *CreateTemplateParams) (*CreateTemplateResponse, error) {
	resp, err := s.cs.newRequest("createTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTemplateResponse
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

type CreateTemplateResponse struct {
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

func (r *CreateTemplateResponse) UnmarshalJSON(b []byte) error {
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

	type alias CreateTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteTemplateParams struct {
	p map[string]interface{}
}

func (p *DeleteTemplateParams) toURLValues() url.Values {
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

func (p *DeleteTemplateParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DeleteTemplateParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *DeleteTemplateParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *DeleteTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteTemplateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteTemplateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteTemplateParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
}

func (p *DeleteTemplateParams) ResetIssystem() {
	if p.p != nil && p.p["issystem"] != nil {
		delete(p.p, "issystem")
	}
}

func (p *DeleteTemplateParams) GetIssystem() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["issystem"].(bool)
	return value, ok
}

func (p *DeleteTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewDeleteTemplateParams(id string) *DeleteTemplateParams {
	p := &DeleteTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a template from the system. All virtual machines using the deleted template will not be affected.
func (s *TemplateService) DeleteTemplate(p *DeleteTemplateParams) (*DeleteTemplateResponse, error) {
	resp, err := s.cs.newRequest("deleteTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTemplateResponse
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

type DeleteTemplateResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ExtractTemplateParams struct {
	p map[string]interface{}
}

func (p *ExtractTemplateParams) toURLValues() url.Values {
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

func (p *ExtractTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ExtractTemplateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ExtractTemplateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ExtractTemplateParams) SetMode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mode"] = v
}

func (p *ExtractTemplateParams) ResetMode() {
	if p.p != nil && p.p["mode"] != nil {
		delete(p.p, "mode")
	}
}

func (p *ExtractTemplateParams) GetMode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mode"].(string)
	return value, ok
}

func (p *ExtractTemplateParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *ExtractTemplateParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *ExtractTemplateParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *ExtractTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ExtractTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ExtractTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ExtractTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewExtractTemplateParams(id string, mode string) *ExtractTemplateParams {
	p := &ExtractTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["mode"] = mode
	return p
}

// Extracts a template
func (s *TemplateService) ExtractTemplate(p *ExtractTemplateParams) (*ExtractTemplateResponse, error) {
	resp, err := s.cs.newRequest("extractTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractTemplateResponse
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

type ExtractTemplateResponse struct {
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

type GetUploadParamsForTemplateParams struct {
	p map[string]interface{}
}

func (p *GetUploadParamsForTemplateParams) toURLValues() url.Values {
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *GetUploadParamsForTemplateParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *GetUploadParamsForTemplateParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *GetUploadParamsForTemplateParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetBits() {
	if p.p != nil && p.p["bits"] != nil {
		delete(p.p, "bits")
	}
}

func (p *GetUploadParamsForTemplateParams) GetBits() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bits"].(int)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetChecksum() {
	if p.p != nil && p.p["checksum"] != nil {
		delete(p.p, "checksum")
	}
}

func (p *GetUploadParamsForTemplateParams) GetChecksum() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["checksum"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetDeployasis(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deployasis"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetDeployasis() {
	if p.p != nil && p.p["deployasis"] != nil {
		delete(p.p, "deployasis")
	}
}

func (p *GetUploadParamsForTemplateParams) GetDeployasis() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deployasis"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *GetUploadParamsForTemplateParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *GetUploadParamsForTemplateParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *GetUploadParamsForTemplateParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetFormat() {
	if p.p != nil && p.p["format"] != nil {
		delete(p.p, "format")
	}
}

func (p *GetUploadParamsForTemplateParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *GetUploadParamsForTemplateParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *GetUploadParamsForTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetIsextractable() {
	if p.p != nil && p.p["isextractable"] != nil {
		delete(p.p, "isextractable")
	}
}

func (p *GetUploadParamsForTemplateParams) GetIsextractable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isextractable"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetIsfeatured() {
	if p.p != nil && p.p["isfeatured"] != nil {
		delete(p.p, "isfeatured")
	}
}

func (p *GetUploadParamsForTemplateParams) GetIsfeatured() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isfeatured"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *GetUploadParamsForTemplateParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetIsrouting() {
	if p.p != nil && p.p["isrouting"] != nil {
		delete(p.p, "isrouting")
	}
}

func (p *GetUploadParamsForTemplateParams) GetIsrouting() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrouting"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *GetUploadParamsForTemplateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *GetUploadParamsForTemplateParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetPasswordenabled() {
	if p.p != nil && p.p["passwordenabled"] != nil {
		delete(p.p, "passwordenabled")
	}
}

func (p *GetUploadParamsForTemplateParams) GetPasswordenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passwordenabled"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *GetUploadParamsForTemplateParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetRequireshvm() {
	if p.p != nil && p.p["requireshvm"] != nil {
		delete(p.p, "requireshvm")
	}
}

func (p *GetUploadParamsForTemplateParams) GetRequireshvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["requireshvm"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetSshkeyenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sshkeyenabled"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetSshkeyenabled() {
	if p.p != nil && p.p["sshkeyenabled"] != nil {
		delete(p.p, "sshkeyenabled")
	}
}

func (p *GetUploadParamsForTemplateParams) GetSshkeyenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sshkeyenabled"].(bool)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetTemplatetag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetag"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetTemplatetag() {
	if p.p != nil && p.p["templatetag"] != nil {
		delete(p.p, "templatetag")
	}
}

func (p *GetUploadParamsForTemplateParams) GetTemplatetag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetag"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *GetUploadParamsForTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *GetUploadParamsForTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new GetUploadParamsForTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewGetUploadParamsForTemplateParams(displaytext string, format string, hypervisor string, name string, zoneid string) *GetUploadParamsForTemplateParams {
	p := &GetUploadParamsForTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["format"] = format
	p.p["hypervisor"] = hypervisor
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// upload an existing template into the CloudStack cloud.
func (s *TemplateService) GetUploadParamsForTemplate(p *GetUploadParamsForTemplateParams) (*GetUploadParamsForTemplateResponse, error) {
	resp, err := s.cs.newRequest("getUploadParamsForTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response GetUploadParamsForTemplateResponse `json:"getuploadparams"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type GetUploadParamsForTemplateResponse struct {
	Expires   string `json:"expires"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Metadata  string `json:"metadata"`
	PostURL   string `json:"postURL"`
	Signature string `json:"signature"`
}

type ListTemplatePermissionsParams struct {
	p map[string]interface{}
}

func (p *ListTemplatePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ListTemplatePermissionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListTemplatePermissionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListTemplatePermissionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ListTemplatePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewListTemplatePermissionsParams(id string) *ListTemplatePermissionsParams {
	p := &ListTemplatePermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplatePermissionByID(id string, opts ...OptionFunc) (*TemplatePermission, int, error) {
	p := &ListTemplatePermissionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListTemplatePermissions(p)
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
		return l.TemplatePermissions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for TemplatePermission UUID: %s!", id)
}

// List template visibility and all accounts that have permissions to view this template.
func (s *TemplateService) ListTemplatePermissions(p *ListTemplatePermissionsParams) (*ListTemplatePermissionsResponse, error) {
	resp, err := s.cs.newRequest("listTemplatePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTemplatePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTemplatePermissionsResponse struct {
	Count               int                   `json:"count"`
	TemplatePermissions []*TemplatePermission `json:"templatepermission"`
}

type TemplatePermission struct {
	Account    []string `json:"account"`
	Domainid   string   `json:"domainid"`
	Id         string   `json:"id"`
	Ispublic   bool     `json:"ispublic"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Projectids []string `json:"projectids"`
}

type ListTemplatesParams struct {
	p map[string]interface{}
}

func (p *ListTemplatesParams) toURLValues() url.Values {
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
	if v, found := p.p["imagestoreid"]; found {
		u.Set("imagestoreid", v.(string))
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

func (p *ListTemplatesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListTemplatesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListTemplatesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *ListTemplatesParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *ListTemplatesParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListTemplatesParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListTemplatesParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListTemplatesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListTemplatesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListTemplatesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *ListTemplatesParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *ListTemplatesParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListTemplatesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListTemplatesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListTemplatesParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListTemplatesParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListTemplatesParams) SetImagestoreid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["imagestoreid"] = v
}

func (p *ListTemplatesParams) ResetImagestoreid() {
	if p.p != nil && p.p["imagestoreid"] != nil {
		delete(p.p, "imagestoreid")
	}
}

func (p *ListTemplatesParams) GetImagestoreid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["imagestoreid"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListTemplatesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListTemplatesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListTemplatesParams) SetIsvnf(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isvnf"] = v
}

func (p *ListTemplatesParams) ResetIsvnf() {
	if p.p != nil && p.p["isvnf"] != nil {
		delete(p.p, "isvnf")
	}
}

func (p *ListTemplatesParams) GetIsvnf() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isvnf"].(bool)
	return value, ok
}

func (p *ListTemplatesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTemplatesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTemplatesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListTemplatesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListTemplatesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListTemplatesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListTemplatesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListTemplatesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTemplatesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTemplatesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTemplatesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTemplatesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTemplatesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTemplatesParams) SetParenttemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parenttemplateid"] = v
}

func (p *ListTemplatesParams) ResetParenttemplateid() {
	if p.p != nil && p.p["parenttemplateid"] != nil {
		delete(p.p, "parenttemplateid")
	}
}

func (p *ListTemplatesParams) GetParenttemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parenttemplateid"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListTemplatesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListTemplatesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListTemplatesParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListTemplatesParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListTemplatesParams) SetShowremoved(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showremoved"] = v
}

func (p *ListTemplatesParams) ResetShowremoved() {
	if p.p != nil && p.p["showremoved"] != nil {
		delete(p.p, "showremoved")
	}
}

func (p *ListTemplatesParams) GetShowremoved() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showremoved"].(bool)
	return value, ok
}

func (p *ListTemplatesParams) SetShowunique(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showunique"] = v
}

func (p *ListTemplatesParams) ResetShowunique() {
	if p.p != nil && p.p["showunique"] != nil {
		delete(p.p, "showunique")
	}
}

func (p *ListTemplatesParams) GetShowunique() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showunique"].(bool)
	return value, ok
}

func (p *ListTemplatesParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListTemplatesParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *ListTemplatesParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListTemplatesParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListTemplatesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListTemplatesParams) SetTemplatefilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatefilter"] = v
}

func (p *ListTemplatesParams) ResetTemplatefilter() {
	if p.p != nil && p.p["templatefilter"] != nil {
		delete(p.p, "templatefilter")
	}
}

func (p *ListTemplatesParams) GetTemplatefilter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatefilter"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetTemplatetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetype"] = v
}

func (p *ListTemplatesParams) ResetTemplatetype() {
	if p.p != nil && p.p["templatetype"] != nil {
		delete(p.p, "templatetype")
	}
}

func (p *ListTemplatesParams) GetTemplatetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetype"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTemplatesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTemplatesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTemplatesParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewListTemplatesParams(templatefilter string) *ListTemplatesParams {
	p := &ListTemplatesParams{}
	p.p = make(map[string]interface{})
	p.p["templatefilter"] = templatefilter
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateID(name string, templatefilter string, zoneid string, opts ...OptionFunc) (string, int, error) {
	p := &ListTemplatesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name
	p.p["templatefilter"] = templatefilter
	p.p["zoneid"] = zoneid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListTemplates(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Templates[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Templates {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateByName(name string, templatefilter string, zoneid string, opts ...OptionFunc) (*Template, int, error) {
	id, count, err := s.GetTemplateID(name, templatefilter, zoneid, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetTemplateByID(id, templatefilter, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateByID(id string, templatefilter string, opts ...OptionFunc) (*Template, int, error) {
	p := &ListTemplatesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id
	p.p["templatefilter"] = templatefilter

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListTemplates(p)
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
		return l.Templates[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Template UUID: %s!", id)
}

// List all public, private, and privileged templates.
func (s *TemplateService) ListTemplates(p *ListTemplatesParams) (*ListTemplatesResponse, error) {
	resp, err := s.cs.newRequest("listTemplates", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTemplatesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTemplatesResponse struct {
	Count     int         `json:"count"`
	Templates []*Template `json:"template"`
}

type Template struct {
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

func (r *Template) UnmarshalJSON(b []byte) error {
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

	type alias Template
	return json.Unmarshal(b, (*alias)(r))
}

type PrepareTemplateParams struct {
	p map[string]interface{}
}

func (p *PrepareTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *PrepareTemplateParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *PrepareTemplateParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *PrepareTemplateParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *PrepareTemplateParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *PrepareTemplateParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *PrepareTemplateParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *PrepareTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *PrepareTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *PrepareTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new PrepareTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewPrepareTemplateParams(templateid string, zoneid string) *PrepareTemplateParams {
	p := &PrepareTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	return p
}

// load template into primary storage
func (s *TemplateService) PrepareTemplate(p *PrepareTemplateParams) (*PrepareTemplateResponse, error) {
	resp, err := s.cs.newRequest("prepareTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PrepareTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type PrepareTemplateResponse struct {
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

func (r *PrepareTemplateResponse) UnmarshalJSON(b []byte) error {
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

	type alias PrepareTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RegisterTemplateParams struct {
	p map[string]interface{}
}

func (p *RegisterTemplateParams) toURLValues() url.Values {
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
		for _, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[0].%s", k), m[k])
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	if v, found := p.p["zoneids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneids", vv)
	}
	return u
}

func (p *RegisterTemplateParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *RegisterTemplateParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *RegisterTemplateParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *RegisterTemplateParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *RegisterTemplateParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
}

func (p *RegisterTemplateParams) ResetBits() {
	if p.p != nil && p.p["bits"] != nil {
		delete(p.p, "bits")
	}
}

func (p *RegisterTemplateParams) GetBits() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bits"].(int)
	return value, ok
}

func (p *RegisterTemplateParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
}

func (p *RegisterTemplateParams) ResetChecksum() {
	if p.p != nil && p.p["checksum"] != nil {
		delete(p.p, "checksum")
	}
}

func (p *RegisterTemplateParams) GetChecksum() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["checksum"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetDeployasis(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deployasis"] = v
}

func (p *RegisterTemplateParams) ResetDeployasis() {
	if p.p != nil && p.p["deployasis"] != nil {
		delete(p.p, "deployasis")
	}
}

func (p *RegisterTemplateParams) GetDeployasis() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deployasis"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *RegisterTemplateParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *RegisterTemplateParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *RegisterTemplateParams) SetDirectdownload(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["directdownload"] = v
}

func (p *RegisterTemplateParams) ResetDirectdownload() {
	if p.p != nil && p.p["directdownload"] != nil {
		delete(p.p, "directdownload")
	}
}

func (p *RegisterTemplateParams) GetDirectdownload() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["directdownload"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *RegisterTemplateParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *RegisterTemplateParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *RegisterTemplateParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *RegisterTemplateParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *RegisterTemplateParams) ResetFormat() {
	if p.p != nil && p.p["format"] != nil {
		delete(p.p, "format")
	}
}

func (p *RegisterTemplateParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *RegisterTemplateParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *RegisterTemplateParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *RegisterTemplateParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *RegisterTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
}

func (p *RegisterTemplateParams) ResetIsextractable() {
	if p.p != nil && p.p["isextractable"] != nil {
		delete(p.p, "isextractable")
	}
}

func (p *RegisterTemplateParams) GetIsextractable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isextractable"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
}

func (p *RegisterTemplateParams) ResetIsfeatured() {
	if p.p != nil && p.p["isfeatured"] != nil {
		delete(p.p, "isfeatured")
	}
}

func (p *RegisterTemplateParams) GetIsfeatured() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isfeatured"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *RegisterTemplateParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *RegisterTemplateParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
}

func (p *RegisterTemplateParams) ResetIsrouting() {
	if p.p != nil && p.p["isrouting"] != nil {
		delete(p.p, "isrouting")
	}
}

func (p *RegisterTemplateParams) GetIsrouting() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrouting"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *RegisterTemplateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *RegisterTemplateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *RegisterTemplateParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *RegisterTemplateParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
}

func (p *RegisterTemplateParams) ResetPasswordenabled() {
	if p.p != nil && p.p["passwordenabled"] != nil {
		delete(p.p, "passwordenabled")
	}
}

func (p *RegisterTemplateParams) GetPasswordenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passwordenabled"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *RegisterTemplateParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *RegisterTemplateParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
}

func (p *RegisterTemplateParams) ResetRequireshvm() {
	if p.p != nil && p.p["requireshvm"] != nil {
		delete(p.p, "requireshvm")
	}
}

func (p *RegisterTemplateParams) GetRequireshvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["requireshvm"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetSshkeyenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sshkeyenabled"] = v
}

func (p *RegisterTemplateParams) ResetSshkeyenabled() {
	if p.p != nil && p.p["sshkeyenabled"] != nil {
		delete(p.p, "sshkeyenabled")
	}
}

func (p *RegisterTemplateParams) GetSshkeyenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sshkeyenabled"].(bool)
	return value, ok
}

func (p *RegisterTemplateParams) SetTemplatetag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetag"] = v
}

func (p *RegisterTemplateParams) ResetTemplatetag() {
	if p.p != nil && p.p["templatetag"] != nil {
		delete(p.p, "templatetag")
	}
}

func (p *RegisterTemplateParams) GetTemplatetag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetag"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetTemplatetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetype"] = v
}

func (p *RegisterTemplateParams) ResetTemplatetype() {
	if p.p != nil && p.p["templatetype"] != nil {
		delete(p.p, "templatetype")
	}
}

func (p *RegisterTemplateParams) GetTemplatetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetype"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *RegisterTemplateParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *RegisterTemplateParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *RegisterTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *RegisterTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetZoneids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneids"] = v
}

func (p *RegisterTemplateParams) ResetZoneids() {
	if p.p != nil && p.p["zoneids"] != nil {
		delete(p.p, "zoneids")
	}
}

func (p *RegisterTemplateParams) GetZoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneids"].([]string)
	return value, ok
}

// You should always use this function to get a new RegisterTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewRegisterTemplateParams(displaytext string, format string, hypervisor string, name string, url string) *RegisterTemplateParams {
	p := &RegisterTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["format"] = format
	p.p["hypervisor"] = hypervisor
	p.p["name"] = name
	p.p["url"] = url
	return p
}

// Registers an existing template into the CloudStack cloud.
func (s *TemplateService) RegisterTemplate(p *RegisterTemplateParams) (*RegisterTemplateResponse, error) {
	resp, err := s.cs.newRequest("registerTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterTemplateResponse struct {
	Count            int                 `json:"count"`
	RegisterTemplate []*RegisterTemplate `json:"template"`
}

type RegisterTemplate struct {
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

func (r *RegisterTemplate) UnmarshalJSON(b []byte) error {
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

	type alias RegisterTemplate
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateTemplateParams struct {
	p map[string]interface{}
}

func (p *UpdateTemplateParams) toURLValues() url.Values {
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
		for _, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[0].%s", k), m[k])
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
	return u
}

func (p *UpdateTemplateParams) SetArch(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["arch"] = v
}

func (p *UpdateTemplateParams) ResetArch() {
	if p.p != nil && p.p["arch"] != nil {
		delete(p.p, "arch")
	}
}

func (p *UpdateTemplateParams) GetArch() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["arch"].(string)
	return value, ok
}

func (p *UpdateTemplateParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
}

func (p *UpdateTemplateParams) ResetBootable() {
	if p.p != nil && p.p["bootable"] != nil {
		delete(p.p, "bootable")
	}
}

func (p *UpdateTemplateParams) GetBootable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootable"].(bool)
	return value, ok
}

func (p *UpdateTemplateParams) SetCleanupdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupdetails"] = v
}

func (p *UpdateTemplateParams) ResetCleanupdetails() {
	if p.p != nil && p.p["cleanupdetails"] != nil {
		delete(p.p, "cleanupdetails")
	}
}

func (p *UpdateTemplateParams) GetCleanupdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupdetails"].(bool)
	return value, ok
}

func (p *UpdateTemplateParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateTemplateParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateTemplateParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateTemplateParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateTemplateParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *UpdateTemplateParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateTemplateParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *UpdateTemplateParams) ResetFormat() {
	if p.p != nil && p.p["format"] != nil {
		delete(p.p, "format")
	}
}

func (p *UpdateTemplateParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *UpdateTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateTemplateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateTemplateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
}

func (p *UpdateTemplateParams) ResetIsdynamicallyscalable() {
	if p.p != nil && p.p["isdynamicallyscalable"] != nil {
		delete(p.p, "isdynamicallyscalable")
	}
}

func (p *UpdateTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdynamicallyscalable"].(bool)
	return value, ok
}

func (p *UpdateTemplateParams) SetIsrouting(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrouting"] = v
}

func (p *UpdateTemplateParams) ResetIsrouting() {
	if p.p != nil && p.p["isrouting"] != nil {
		delete(p.p, "isrouting")
	}
}

func (p *UpdateTemplateParams) GetIsrouting() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrouting"].(bool)
	return value, ok
}

func (p *UpdateTemplateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateTemplateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateTemplateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateTemplateParams) SetOstypeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ostypeid"] = v
}

func (p *UpdateTemplateParams) ResetOstypeid() {
	if p.p != nil && p.p["ostypeid"] != nil {
		delete(p.p, "ostypeid")
	}
}

func (p *UpdateTemplateParams) GetOstypeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ostypeid"].(string)
	return value, ok
}

func (p *UpdateTemplateParams) SetPasswordenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passwordenabled"] = v
}

func (p *UpdateTemplateParams) ResetPasswordenabled() {
	if p.p != nil && p.p["passwordenabled"] != nil {
		delete(p.p, "passwordenabled")
	}
}

func (p *UpdateTemplateParams) GetPasswordenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["passwordenabled"].(bool)
	return value, ok
}

func (p *UpdateTemplateParams) SetRequireshvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["requireshvm"] = v
}

func (p *UpdateTemplateParams) ResetRequireshvm() {
	if p.p != nil && p.p["requireshvm"] != nil {
		delete(p.p, "requireshvm")
	}
}

func (p *UpdateTemplateParams) GetRequireshvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["requireshvm"].(bool)
	return value, ok
}

func (p *UpdateTemplateParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

func (p *UpdateTemplateParams) ResetSortkey() {
	if p.p != nil && p.p["sortkey"] != nil {
		delete(p.p, "sortkey")
	}
}

func (p *UpdateTemplateParams) GetSortkey() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortkey"].(int)
	return value, ok
}

func (p *UpdateTemplateParams) SetSshkeyenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sshkeyenabled"] = v
}

func (p *UpdateTemplateParams) ResetSshkeyenabled() {
	if p.p != nil && p.p["sshkeyenabled"] != nil {
		delete(p.p, "sshkeyenabled")
	}
}

func (p *UpdateTemplateParams) GetSshkeyenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sshkeyenabled"].(bool)
	return value, ok
}

func (p *UpdateTemplateParams) SetTemplatetag(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetag"] = v
}

func (p *UpdateTemplateParams) ResetTemplatetag() {
	if p.p != nil && p.p["templatetag"] != nil {
		delete(p.p, "templatetag")
	}
}

func (p *UpdateTemplateParams) GetTemplatetag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetag"].(string)
	return value, ok
}

func (p *UpdateTemplateParams) SetTemplatetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetype"] = v
}

func (p *UpdateTemplateParams) ResetTemplatetype() {
	if p.p != nil && p.p["templatetype"] != nil {
		delete(p.p, "templatetype")
	}
}

func (p *UpdateTemplateParams) GetTemplatetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetype"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpdateTemplateParams(id string) *UpdateTemplateParams {
	p := &UpdateTemplateParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates attributes of a template.
func (s *TemplateService) UpdateTemplate(p *UpdateTemplateParams) (*UpdateTemplateResponse, error) {
	resp, err := s.cs.newRequest("updateTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateTemplateResponse struct {
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

func (r *UpdateTemplateResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateTemplatePermissionsParams struct {
	p map[string]interface{}
}

func (p *UpdateTemplatePermissionsParams) toURLValues() url.Values {
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

func (p *UpdateTemplatePermissionsParams) SetAccounts(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounts"] = v
}

func (p *UpdateTemplatePermissionsParams) ResetAccounts() {
	if p.p != nil && p.p["accounts"] != nil {
		delete(p.p, "accounts")
	}
}

func (p *UpdateTemplatePermissionsParams) GetAccounts() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accounts"].([]string)
	return value, ok
}

func (p *UpdateTemplatePermissionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateTemplatePermissionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateTemplatePermissionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateTemplatePermissionsParams) SetIsextractable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isextractable"] = v
}

func (p *UpdateTemplatePermissionsParams) ResetIsextractable() {
	if p.p != nil && p.p["isextractable"] != nil {
		delete(p.p, "isextractable")
	}
}

func (p *UpdateTemplatePermissionsParams) GetIsextractable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isextractable"].(bool)
	return value, ok
}

func (p *UpdateTemplatePermissionsParams) SetIsfeatured(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isfeatured"] = v
}

func (p *UpdateTemplatePermissionsParams) ResetIsfeatured() {
	if p.p != nil && p.p["isfeatured"] != nil {
		delete(p.p, "isfeatured")
	}
}

func (p *UpdateTemplatePermissionsParams) GetIsfeatured() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isfeatured"].(bool)
	return value, ok
}

func (p *UpdateTemplatePermissionsParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *UpdateTemplatePermissionsParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *UpdateTemplatePermissionsParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *UpdateTemplatePermissionsParams) SetOp(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["op"] = v
}

func (p *UpdateTemplatePermissionsParams) ResetOp() {
	if p.p != nil && p.p["op"] != nil {
		delete(p.p, "op")
	}
}

func (p *UpdateTemplatePermissionsParams) GetOp() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["op"].(string)
	return value, ok
}

func (p *UpdateTemplatePermissionsParams) SetProjectids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectids"] = v
}

func (p *UpdateTemplatePermissionsParams) ResetProjectids() {
	if p.p != nil && p.p["projectids"] != nil {
		delete(p.p, "projectids")
	}
}

func (p *UpdateTemplatePermissionsParams) GetProjectids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectids"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateTemplatePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpdateTemplatePermissionsParams(id string) *UpdateTemplatePermissionsParams {
	p := &UpdateTemplatePermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a template visibility permissions. A public template is visible to all accounts within the same domain. A private template is visible only to the owner of the template. A privileged template is a private template with account permissions added. Only accounts specified under the template permissions are visible to them.
func (s *TemplateService) UpdateTemplatePermissions(p *UpdateTemplatePermissionsParams) (*UpdateTemplatePermissionsResponse, error) {
	resp, err := s.cs.newRequest("updateTemplatePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTemplatePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateTemplatePermissionsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateTemplatePermissionsResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateTemplatePermissionsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpgradeRouterTemplateParams struct {
	p map[string]interface{}
}

func (p *UpgradeRouterTemplateParams) toURLValues() url.Values {
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
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *UpgradeRouterTemplateParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpgradeRouterTemplateParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *UpgradeRouterTemplateParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UpgradeRouterTemplateParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *UpgradeRouterTemplateParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *UpgradeRouterTemplateParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *UpgradeRouterTemplateParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpgradeRouterTemplateParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *UpgradeRouterTemplateParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpgradeRouterTemplateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpgradeRouterTemplateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpgradeRouterTemplateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpgradeRouterTemplateParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *UpgradeRouterTemplateParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *UpgradeRouterTemplateParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *UpgradeRouterTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *UpgradeRouterTemplateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *UpgradeRouterTemplateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpgradeRouterTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpgradeRouterTemplateParams() *UpgradeRouterTemplateParams {
	p := &UpgradeRouterTemplateParams{}
	p.p = make(map[string]interface{})
	return p
}

// Upgrades router to use newer template
func (s *TemplateService) UpgradeRouterTemplate(p *UpgradeRouterTemplateParams) (*UpgradeRouterTemplateResponse, error) {
	resp, err := s.cs.newRequest("upgradeRouterTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpgradeRouterTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpgradeRouterTemplateResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
}

type LinkUserDataToTemplateParams struct {
	p map[string]interface{}
}

func (p *LinkUserDataToTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["isoid"]; found {
		u.Set("isoid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["userdataid"]; found {
		u.Set("userdataid", v.(string))
	}
	if v, found := p.p["userdatapolicy"]; found {
		u.Set("userdatapolicy", v.(string))
	}
	return u
}

func (p *LinkUserDataToTemplateParams) SetIsoid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isoid"] = v
}

func (p *LinkUserDataToTemplateParams) ResetIsoid() {
	if p.p != nil && p.p["isoid"] != nil {
		delete(p.p, "isoid")
	}
}

func (p *LinkUserDataToTemplateParams) GetIsoid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isoid"].(string)
	return value, ok
}

func (p *LinkUserDataToTemplateParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *LinkUserDataToTemplateParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *LinkUserDataToTemplateParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *LinkUserDataToTemplateParams) SetUserdataid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdataid"] = v
}

func (p *LinkUserDataToTemplateParams) ResetUserdataid() {
	if p.p != nil && p.p["userdataid"] != nil {
		delete(p.p, "userdataid")
	}
}

func (p *LinkUserDataToTemplateParams) GetUserdataid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdataid"].(string)
	return value, ok
}

func (p *LinkUserDataToTemplateParams) SetUserdatapolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdatapolicy"] = v
}

func (p *LinkUserDataToTemplateParams) ResetUserdatapolicy() {
	if p.p != nil && p.p["userdatapolicy"] != nil {
		delete(p.p, "userdatapolicy")
	}
}

func (p *LinkUserDataToTemplateParams) GetUserdatapolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdatapolicy"].(string)
	return value, ok
}

// You should always use this function to get a new LinkUserDataToTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewLinkUserDataToTemplateParams() *LinkUserDataToTemplateParams {
	p := &LinkUserDataToTemplateParams{}
	p.p = make(map[string]interface{})
	return p
}

// Link or unlink a userdata to a template.
func (s *TemplateService) LinkUserDataToTemplate(p *LinkUserDataToTemplateParams) (*LinkUserDataToTemplateResponse, error) {
	resp, err := s.cs.newRequest("linkUserDataToTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response LinkUserDataToTemplateResponse `json:"template"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type LinkUserDataToTemplateResponse struct {
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

func (r *LinkUserDataToTemplateResponse) UnmarshalJSON(b []byte) error {
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

	type alias LinkUserDataToTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}
