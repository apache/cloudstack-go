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

type CertificateServiceIface interface {
	IssueCertificate(p *IssueCertificateParams) (*IssueCertificateResponse, error)
	NewIssueCertificateParams() *IssueCertificateParams
	ListCAProviders(p *ListCAProvidersParams) (*ListCAProvidersResponse, error)
	NewListCAProvidersParams() *ListCAProvidersParams
	ListCaCertificate(p *ListCaCertificateParams) (*ListCaCertificateResponse, error)
	NewListCaCertificateParams() *ListCaCertificateParams
	ListTemplateDirectDownloadCertificates(p *ListTemplateDirectDownloadCertificatesParams) (*ListTemplateDirectDownloadCertificatesResponse, error)
	NewListTemplateDirectDownloadCertificatesParams() *ListTemplateDirectDownloadCertificatesParams
	GetTemplateDirectDownloadCertificateByID(id string, opts ...OptionFunc) (*TemplateDirectDownloadCertificate, int, error)
	ProvisionCertificate(p *ProvisionCertificateParams) (*ProvisionCertificateResponse, error)
	NewProvisionCertificateParams(hostid string) *ProvisionCertificateParams
	ProvisionTemplateDirectDownloadCertificate(p *ProvisionTemplateDirectDownloadCertificateParams) (*ProvisionTemplateDirectDownloadCertificateResponse, error)
	NewProvisionTemplateDirectDownloadCertificateParams(hostid string, id string) *ProvisionTemplateDirectDownloadCertificateParams
	RevokeCertificate(p *RevokeCertificateParams) (*RevokeCertificateResponse, error)
	NewRevokeCertificateParams(serial string) *RevokeCertificateParams
	RevokeTemplateDirectDownloadCertificate(p *RevokeTemplateDirectDownloadCertificateParams) (*RevokeTemplateDirectDownloadCertificateResponse, error)
	NewRevokeTemplateDirectDownloadCertificateParams(zoneid string) *RevokeTemplateDirectDownloadCertificateParams
	UploadCustomCertificate(p *UploadCustomCertificateParams) (*UploadCustomCertificateResponse, error)
	NewUploadCustomCertificateParams(certificate string, domainsuffix string) *UploadCustomCertificateParams
	UploadTemplateDirectDownloadCertificate(p *UploadTemplateDirectDownloadCertificateParams) (*UploadTemplateDirectDownloadCertificateResponse, error)
	NewUploadTemplateDirectDownloadCertificateParams(certificate string, hypervisor string, name string, zoneid string) *UploadTemplateDirectDownloadCertificateParams
}

type IssueCertificateParams struct {
	p map[string]interface{}
}

func (p *IssueCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["csr"]; found {
		u.Set("csr", v.(string))
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (p *IssueCertificateParams) SetCsr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["csr"] = v
}

func (p *IssueCertificateParams) ResetCsr() {
	if p.p != nil && p.p["csr"] != nil {
		delete(p.p, "csr")
	}
}

func (p *IssueCertificateParams) GetCsr() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["csr"].(string)
	return value, ok
}

func (p *IssueCertificateParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *IssueCertificateParams) ResetDomain() {
	if p.p != nil && p.p["domain"] != nil {
		delete(p.p, "domain")
	}
}

func (p *IssueCertificateParams) GetDomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domain"].(string)
	return value, ok
}

func (p *IssueCertificateParams) SetDuration(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["duration"] = v
}

func (p *IssueCertificateParams) ResetDuration() {
	if p.p != nil && p.p["duration"] != nil {
		delete(p.p, "duration")
	}
}

func (p *IssueCertificateParams) GetDuration() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["duration"].(int)
	return value, ok
}

func (p *IssueCertificateParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *IssueCertificateParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *IssueCertificateParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *IssueCertificateParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *IssueCertificateParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *IssueCertificateParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

// You should always use this function to get a new IssueCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewIssueCertificateParams() *IssueCertificateParams {
	p := &IssueCertificateParams{}
	p.p = make(map[string]interface{})
	return p
}

// Issues a client certificate using configured or provided CA plugin
func (s *CertificateService) IssueCertificate(p *IssueCertificateParams) (*IssueCertificateResponse, error) {
	resp, err := s.cs.newPostRequest("issueCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r IssueCertificateResponse
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

type IssueCertificateResponse struct {
	Cacertificates string `json:"cacertificates"`
	Certificate    string `json:"certificate"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Privatekey     string `json:"privatekey"`
}

type ListCAProvidersParams struct {
	p map[string]interface{}
}

func (p *ListCAProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *ListCAProvidersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListCAProvidersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListCAProvidersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new ListCAProvidersParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewListCAProvidersParams() *ListCAProvidersParams {
	p := &ListCAProvidersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists available certificate authority providers in CloudStack
func (s *CertificateService) ListCAProviders(p *ListCAProvidersParams) (*ListCAProvidersResponse, error) {
	resp, err := s.cs.newRequest("listCAProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCAProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCAProvidersResponse struct {
	Count       int           `json:"count"`
	CAProviders []*CAProvider `json:"caprovider"`
}

type CAProvider struct {
	Description string `json:"description"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
}

type ListCaCertificateParams struct {
	p map[string]interface{}
}

func (p *ListCaCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (p *ListCaCertificateParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListCaCertificateParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListCaCertificateParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

// You should always use this function to get a new ListCaCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewListCaCertificateParams() *ListCaCertificateParams {
	p := &ListCaCertificateParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists the CA public certificate(s) as support by the configured/provided CA plugin
func (s *CertificateService) ListCaCertificate(p *ListCaCertificateParams) (*ListCaCertificateResponse, error) {
	resp, err := s.cs.newRequest("listCaCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCaCertificateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCaCertificateResponse struct {
	Count         int              `json:"count"`
	CaCertificate []*CaCertificate `json:"cacertificate"`
}

type CaCertificate struct {
	Cacertificates string `json:"cacertificates"`
	Certificate    string `json:"certificate"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Privatekey     string `json:"privatekey"`
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

func (p *ListTemplateDirectDownloadCertificatesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *ListTemplateDirectDownloadCertificatesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListTemplateDirectDownloadCertificatesParams) ResetListhosts() {
	if p.p != nil && p.p["listhosts"] != nil {
		delete(p.p, "listhosts")
	}
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

func (p *ListTemplateDirectDownloadCertificatesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListTemplateDirectDownloadCertificatesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

func (p *ListTemplateDirectDownloadCertificatesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
func (s *CertificateService) NewListTemplateDirectDownloadCertificatesParams() *ListTemplateDirectDownloadCertificatesParams {
	p := &ListTemplateDirectDownloadCertificatesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *CertificateService) GetTemplateDirectDownloadCertificateByID(id string, opts ...OptionFunc) (*TemplateDirectDownloadCertificate, int, error) {
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
func (s *CertificateService) ListTemplateDirectDownloadCertificates(p *ListTemplateDirectDownloadCertificatesParams) (*ListTemplateDirectDownloadCertificatesResponse, error) {
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

type ProvisionCertificateParams struct {
	p map[string]interface{}
}

func (p *ProvisionCertificateParams) toURLValues() url.Values {
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
	if v, found := p.p["reconnect"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("reconnect", vv)
	}
	return u
}

func (p *ProvisionCertificateParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ProvisionCertificateParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ProvisionCertificateParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ProvisionCertificateParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ProvisionCertificateParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ProvisionCertificateParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *ProvisionCertificateParams) SetReconnect(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["reconnect"] = v
}

func (p *ProvisionCertificateParams) ResetReconnect() {
	if p.p != nil && p.p["reconnect"] != nil {
		delete(p.p, "reconnect")
	}
}

func (p *ProvisionCertificateParams) GetReconnect() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["reconnect"].(bool)
	return value, ok
}

// You should always use this function to get a new ProvisionCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewProvisionCertificateParams(hostid string) *ProvisionCertificateParams {
	p := &ProvisionCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	return p
}

// Issues and propagates client certificate on a connected host/agent using configured CA plugin
func (s *CertificateService) ProvisionCertificate(p *ProvisionCertificateParams) (*ProvisionCertificateResponse, error) {
	resp, err := s.cs.newPostRequest("provisionCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ProvisionCertificateResponse
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

type ProvisionCertificateResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
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

func (p *ProvisionTemplateDirectDownloadCertificateParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
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

func (p *ProvisionTemplateDirectDownloadCertificateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
func (s *CertificateService) NewProvisionTemplateDirectDownloadCertificateParams(hostid string, id string) *ProvisionTemplateDirectDownloadCertificateParams {
	p := &ProvisionTemplateDirectDownloadCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	p.p["id"] = id
	return p
}

// Provisions a host with a direct download certificate
func (s *CertificateService) ProvisionTemplateDirectDownloadCertificate(p *ProvisionTemplateDirectDownloadCertificateParams) (*ProvisionTemplateDirectDownloadCertificateResponse, error) {
	resp, err := s.cs.newPostRequest("provisionTemplateDirectDownloadCertificate", p.toURLValues())
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

type RevokeCertificateParams struct {
	p map[string]interface{}
}

func (p *RevokeCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cn"]; found {
		u.Set("cn", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["serial"]; found {
		u.Set("serial", v.(string))
	}
	return u
}

func (p *RevokeCertificateParams) SetCn(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cn"] = v
}

func (p *RevokeCertificateParams) ResetCn() {
	if p.p != nil && p.p["cn"] != nil {
		delete(p.p, "cn")
	}
}

func (p *RevokeCertificateParams) GetCn() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cn"].(string)
	return value, ok
}

func (p *RevokeCertificateParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *RevokeCertificateParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *RevokeCertificateParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *RevokeCertificateParams) SetSerial(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serial"] = v
}

func (p *RevokeCertificateParams) ResetSerial() {
	if p.p != nil && p.p["serial"] != nil {
		delete(p.p, "serial")
	}
}

func (p *RevokeCertificateParams) GetSerial() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serial"].(string)
	return value, ok
}

// You should always use this function to get a new RevokeCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewRevokeCertificateParams(serial string) *RevokeCertificateParams {
	p := &RevokeCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["serial"] = serial
	return p
}

// Revokes certificate using configured CA plugin
func (s *CertificateService) RevokeCertificate(p *RevokeCertificateParams) (*RevokeCertificateResponse, error) {
	resp, err := s.cs.newPostRequest("revokeCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeCertificateResponse
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

type RevokeCertificateResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RevokeTemplateDirectDownloadCertificateParams struct {
	p map[string]interface{}
}

func (p *RevokeTemplateDirectDownloadCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *RevokeTemplateDirectDownloadCertificateParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *RevokeTemplateDirectDownloadCertificateParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *RevokeTemplateDirectDownloadCertificateParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *RevokeTemplateDirectDownloadCertificateParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *RevokeTemplateDirectDownloadCertificateParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *RevokeTemplateDirectDownloadCertificateParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *RevokeTemplateDirectDownloadCertificateParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RevokeTemplateDirectDownloadCertificateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RevokeTemplateDirectDownloadCertificateParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *RevokeTemplateDirectDownloadCertificateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *RevokeTemplateDirectDownloadCertificateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *RevokeTemplateDirectDownloadCertificateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *RevokeTemplateDirectDownloadCertificateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *RevokeTemplateDirectDownloadCertificateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *RevokeTemplateDirectDownloadCertificateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new RevokeTemplateDirectDownloadCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewRevokeTemplateDirectDownloadCertificateParams(zoneid string) *RevokeTemplateDirectDownloadCertificateParams {
	p := &RevokeTemplateDirectDownloadCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// Revoke a direct download certificate from hosts in a zone
func (s *CertificateService) RevokeTemplateDirectDownloadCertificate(p *RevokeTemplateDirectDownloadCertificateParams) (*RevokeTemplateDirectDownloadCertificateResponse, error) {
	resp, err := s.cs.newPostRequest("revokeTemplateDirectDownloadCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeTemplateDirectDownloadCertificateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RevokeTemplateDirectDownloadCertificateResponse struct {
	Details   string `json:"details"`
	Hostid    string `json:"hostid"`
	Hostname  string `json:"hostname"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Status    string `json:"status"`
}

type UploadCustomCertificateParams struct {
	p map[string]interface{}
}

func (p *UploadCustomCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["certificate"]; found {
		u.Set("certificate", v.(string))
	}
	if v, found := p.p["domainsuffix"]; found {
		u.Set("domainsuffix", v.(string))
	}
	if v, found := p.p["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["privatekey"]; found {
		u.Set("privatekey", v.(string))
	}
	return u
}

func (p *UploadCustomCertificateParams) SetCertificate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certificate"] = v
}

func (p *UploadCustomCertificateParams) ResetCertificate() {
	if p.p != nil && p.p["certificate"] != nil {
		delete(p.p, "certificate")
	}
}

func (p *UploadCustomCertificateParams) GetCertificate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["certificate"].(string)
	return value, ok
}

func (p *UploadCustomCertificateParams) SetDomainsuffix(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainsuffix"] = v
}

func (p *UploadCustomCertificateParams) ResetDomainsuffix() {
	if p.p != nil && p.p["domainsuffix"] != nil {
		delete(p.p, "domainsuffix")
	}
}

func (p *UploadCustomCertificateParams) GetDomainsuffix() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainsuffix"].(string)
	return value, ok
}

func (p *UploadCustomCertificateParams) SetId(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UploadCustomCertificateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UploadCustomCertificateParams) GetId() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int)
	return value, ok
}

func (p *UploadCustomCertificateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UploadCustomCertificateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UploadCustomCertificateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UploadCustomCertificateParams) SetPrivatekey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privatekey"] = v
}

func (p *UploadCustomCertificateParams) ResetPrivatekey() {
	if p.p != nil && p.p["privatekey"] != nil {
		delete(p.p, "privatekey")
	}
}

func (p *UploadCustomCertificateParams) GetPrivatekey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privatekey"].(string)
	return value, ok
}

// You should always use this function to get a new UploadCustomCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewUploadCustomCertificateParams(certificate string, domainsuffix string) *UploadCustomCertificateParams {
	p := &UploadCustomCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["certificate"] = certificate
	p.p["domainsuffix"] = domainsuffix
	return p
}

// Uploads a custom certificate for the console proxy VMs to use for SSL. Can be used to upload a single certificate signed by a known CA. Can also be used, through multiple calls, to upload a chain of certificates from CA to the custom certificate itself.
func (s *CertificateService) UploadCustomCertificate(p *UploadCustomCertificateParams) (*UploadCustomCertificateResponse, error) {
	resp, err := s.cs.newPostRequest("uploadCustomCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadCustomCertificateResponse
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

type UploadCustomCertificateResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Message   string `json:"message"`
}

type UploadTemplateDirectDownloadCertificateParams struct {
	p map[string]interface{}
}

func (p *UploadTemplateDirectDownloadCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["certificate"]; found {
		u.Set("certificate", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *UploadTemplateDirectDownloadCertificateParams) SetCertificate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certificate"] = v
}

func (p *UploadTemplateDirectDownloadCertificateParams) ResetCertificate() {
	if p.p != nil && p.p["certificate"] != nil {
		delete(p.p, "certificate")
	}
}

func (p *UploadTemplateDirectDownloadCertificateParams) GetCertificate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["certificate"].(string)
	return value, ok
}

func (p *UploadTemplateDirectDownloadCertificateParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *UploadTemplateDirectDownloadCertificateParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *UploadTemplateDirectDownloadCertificateParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *UploadTemplateDirectDownloadCertificateParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *UploadTemplateDirectDownloadCertificateParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *UploadTemplateDirectDownloadCertificateParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *UploadTemplateDirectDownloadCertificateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UploadTemplateDirectDownloadCertificateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UploadTemplateDirectDownloadCertificateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UploadTemplateDirectDownloadCertificateParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *UploadTemplateDirectDownloadCertificateParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *UploadTemplateDirectDownloadCertificateParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UploadTemplateDirectDownloadCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewUploadTemplateDirectDownloadCertificateParams(certificate string, hypervisor string, name string, zoneid string) *UploadTemplateDirectDownloadCertificateParams {
	p := &UploadTemplateDirectDownloadCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["certificate"] = certificate
	p.p["hypervisor"] = hypervisor
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// Upload a certificate for HTTPS direct template download on KVM hosts
func (s *CertificateService) UploadTemplateDirectDownloadCertificate(p *UploadTemplateDirectDownloadCertificateParams) (*UploadTemplateDirectDownloadCertificateResponse, error) {
	resp, err := s.cs.newPostRequest("uploadTemplateDirectDownloadCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadTemplateDirectDownloadCertificateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UploadTemplateDirectDownloadCertificateResponse struct {
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
