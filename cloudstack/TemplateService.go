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
	NewCreateTemplateParams(name string, ostypeid string) *CreateTemplateParams
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
	NewRegisterTemplateParams(format string, hypervisor string, name string, url string) *RegisterTemplateParams
	UpdateTemplate(p *UpdateTemplateParams) (*UpdateTemplateResponse, error)
	NewUpdateTemplateParams(id string) *UpdateTemplateParams
	UpdateTemplatePermissions(p *UpdateTemplatePermissionsParams) (*UpdateTemplatePermissionsResponse, error)
	NewUpdateTemplatePermissionsParams(id string) *UpdateTemplatePermissionsParams
	UpgradeRouterTemplate(p *UpgradeRouterTemplateParams) (*UpgradeRouterTemplateResponse, error)
	NewUpgradeRouterTemplateParams() *UpgradeRouterTemplateParams
	ListTemplateDirectDownloadCertificates(p *ListTemplateDirectDownloadCertificatesParams) (*ListTemplateDirectDownloadCertificatesResponse, error)
	NewListTemplateDirectDownloadCertificatesParams() *ListTemplateDirectDownloadCertificatesParams
	GetTemplateDirectDownloadCertificateByID(id string, opts ...OptionFunc) (*TemplateDirectDownloadCertificate, int, error)
	ProvisionTemplateDirectDownloadCertificate(p *ProvisionTemplateDirectDownloadCertificateParams) (*ProvisionTemplateDirectDownloadCertificateResponse, error)
	NewProvisionTemplateDirectDownloadCertificateParams(hostid string, id string) *ProvisionTemplateDirectDownloadCertificateParams
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
	return u
}

func (p *CreateTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
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

func (p *CreateTemplateParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateTemplateParams) SetIsdynamicallyscalable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdynamicallyscalable"] = v
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

func (p *CreateTemplateParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewCreateTemplateParams(name string, ostypeid string) *CreateTemplateParams {
	p := &CreateTemplateParams{}
	p.p = make(map[string]interface{})
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

func (p *DeleteTemplateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteTemplateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
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

func (p *GetUploadParamsForTemplateParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *GetUploadParamsForTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
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

func (p *ListTemplatesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
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

func (p *ListTemplatesParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListTemplatesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListTemplatesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListTemplatesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
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

func (p *ListTemplatesParams) GetShowunique() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showunique"].(bool)
	return value, ok
}

func (p *ListTemplatesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
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

func (p *ListTemplatesParams) GetTemplatefilter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatefilter"].(string)
	return value, ok
}

func (p *ListTemplatesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
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

func (p *RegisterTemplateParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetBits(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bits"] = v
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

func (p *RegisterTemplateParams) GetTemplatetag() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templatetag"].(string)
	return value, ok
}

func (p *RegisterTemplateParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
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

func (p *RegisterTemplateParams) GetZoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneids"].([]string)
	return value, ok
}

// You should always use this function to get a new RegisterTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewRegisterTemplateParams(format string, hypervisor string, name string, url string) *RegisterTemplateParams {
	p := &RegisterTemplateParams{}
	p.p = make(map[string]interface{})
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
	if v, found := p.p["templatetype"]; found {
		u.Set("templatetype", v.(string))
	}
	return u
}

func (p *UpdateTemplateParams) SetBootable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootable"] = v
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

func (p *UpdateTemplateParams) GetSshkeyenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sshkeyenabled"].(bool)
	return value, ok
}

func (p *UpdateTemplateParams) SetTemplatetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templatetype"] = v
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

type ListTemplateDirectDownloadCertificatesParams struct {
	p map[string]interface{}
}

func (p *ListTemplateDirectDownloadCertificatesParams) toURLValues() url.Values {
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
	if v, found := p.p["listhosts"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listhosts", vv)
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

func (p *ListTemplateDirectDownloadCertificatesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListTemplateDirectDownloadCertificatesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListTemplateDirectDownloadCertificatesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTemplateDirectDownloadCertificatesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTemplateDirectDownloadCertificatesParams) SetListhosts(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listhosts"] = v
}

func (p *ListTemplateDirectDownloadCertificatesParams) GetListhosts() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listhosts"].(bool)
	return value, ok
}

func (p *ListTemplateDirectDownloadCertificatesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTemplateDirectDownloadCertificatesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTemplateDirectDownloadCertificatesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTemplateDirectDownloadCertificatesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTemplateDirectDownloadCertificatesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTemplateDirectDownloadCertificatesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTemplateDirectDownloadCertificatesParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewListTemplateDirectDownloadCertificatesParams() *ListTemplateDirectDownloadCertificatesParams {
	p := &ListTemplateDirectDownloadCertificatesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateDirectDownloadCertificateByID(id string, opts ...OptionFunc) (*TemplateDirectDownloadCertificate, int, error) {
	p := &ListTemplateDirectDownloadCertificatesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListTemplateDirectDownloadCertificates(p)
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
		return l.TemplateDirectDownloadCertificates[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for TemplateDirectDownloadCertificate UUID: %s!", id)
}

// List the uploaded certificates for direct download templates
func (s *TemplateService) ListTemplateDirectDownloadCertificates(p *ListTemplateDirectDownloadCertificatesParams) (*ListTemplateDirectDownloadCertificatesResponse, error) {
	resp, err := s.cs.newRequest("listTemplateDirectDownloadCertificates", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTemplateDirectDownloadCertificatesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTemplateDirectDownloadCertificatesResponse struct {
	Count                              int                                  `json:"count"`
	TemplateDirectDownloadCertificates []*TemplateDirectDownloadCertificate `json:"templatedirectdownloadcertificate"`
}

type TemplateDirectDownloadCertificate struct {
	Alias      string   `json:"alias"`
	Hostsmap   []string `json:"hostsmap"`
	Hypervisor string   `json:"hypervisor"`
	Id         string   `json:"id"`
	Issuer     string   `json:"issuer"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Serialnum  string   `json:"serialnum"`
	Subject    string   `json:"subject"`
	Validity   string   `json:"validity"`
	Version    string   `json:"version"`
	Zoneid     string   `json:"zoneid"`
	Zonename   string   `json:"zonename"`
}

type ProvisionTemplateDirectDownloadCertificateParams struct {
	p map[string]interface{}
}

func (p *ProvisionTemplateDirectDownloadCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ProvisionTemplateDirectDownloadCertificateParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ProvisionTemplateDirectDownloadCertificateParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ProvisionTemplateDirectDownloadCertificateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ProvisionTemplateDirectDownloadCertificateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ProvisionTemplateDirectDownloadCertificateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewProvisionTemplateDirectDownloadCertificateParams(hostid string, id string) *ProvisionTemplateDirectDownloadCertificateParams {
	p := &ProvisionTemplateDirectDownloadCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	p.p["id"] = id
	return p
}

// Provisions a host with a direct download certificate
func (s *TemplateService) ProvisionTemplateDirectDownloadCertificate(p *ProvisionTemplateDirectDownloadCertificateParams) (*ProvisionTemplateDirectDownloadCertificateResponse, error) {
	resp, err := s.cs.newRequest("provisionTemplateDirectDownloadCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ProvisionTemplateDirectDownloadCertificateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ProvisionTemplateDirectDownloadCertificateResponse struct {
	Details   string `json:"details"`
	Hostid    string `json:"hostid"`
	Hostname  string `json:"hostname"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Status    string `json:"status"`
}
