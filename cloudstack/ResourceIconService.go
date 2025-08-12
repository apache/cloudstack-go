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
	"net/url"
	"strconv"
	"strings"
)

type ResourceIconServiceIface interface {
	DeleteResourceIcon(p *DeleteResourceIconParams) (*DeleteResourceIconResponse, error)
	NewDeleteResourceIconParams(resourceids []string, resourcetype string) *DeleteResourceIconParams
	ListResourceIcon(p *ListResourceIconParams) (*ListResourceIconResponse, error)
	NewListResourceIconParams(resourceids []string, resourcetype string) *ListResourceIconParams
	UploadResourceIcon(p *UploadResourceIconParams) (*UploadResourceIconResponse, error)
	NewUploadResourceIconParams(base64image string, resourceids []string, resourcetype string) *UploadResourceIconParams
}

type DeleteResourceIconParams struct {
	p map[string]interface{}
}

func (p *DeleteResourceIconParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["resourceids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("resourceids", vv)
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *DeleteResourceIconParams) SetResourceids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceids"] = v
}

func (p *DeleteResourceIconParams) ResetResourceids() {
	if p.p != nil && p.p["resourceids"] != nil {
		delete(p.p, "resourceids")
	}
}

func (p *DeleteResourceIconParams) GetResourceids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceids"].([]string)
	return value, ok
}

func (p *DeleteResourceIconParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *DeleteResourceIconParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *DeleteResourceIconParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteResourceIconParams instance,
// as then you are sure you have configured all required params
func (s *ResourceIconService) NewDeleteResourceIconParams(resourceids []string, resourcetype string) *DeleteResourceIconParams {
	p := &DeleteResourceIconParams{}
	p.p = make(map[string]interface{})
	p.p["resourceids"] = resourceids
	p.p["resourcetype"] = resourcetype
	return p
}

// deletes the resource icon from the specified resource(s)
func (s *ResourceIconService) DeleteResourceIcon(p *DeleteResourceIconParams) (*DeleteResourceIconResponse, error) {
	resp, err := s.cs.newPostRequest("deleteResourceIcon", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteResourceIconResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteResourceIconResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteResourceIconResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteResourceIconResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListResourceIconParams struct {
	p map[string]interface{}
}

func (p *ListResourceIconParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["resourceids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("resourceids", vv)
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *ListResourceIconParams) SetResourceids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceids"] = v
}

func (p *ListResourceIconParams) ResetResourceids() {
	if p.p != nil && p.p["resourceids"] != nil {
		delete(p.p, "resourceids")
	}
}

func (p *ListResourceIconParams) GetResourceids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceids"].([]string)
	return value, ok
}

func (p *ListResourceIconParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *ListResourceIconParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *ListResourceIconParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new ListResourceIconParams instance,
// as then you are sure you have configured all required params
func (s *ResourceIconService) NewListResourceIconParams(resourceids []string, resourcetype string) *ListResourceIconParams {
	p := &ListResourceIconParams{}
	p.p = make(map[string]interface{})
	p.p["resourceids"] = resourceids
	p.p["resourcetype"] = resourcetype
	return p
}

// Lists the resource icon for the specified resource(s)
func (s *ResourceIconService) ListResourceIcon(p *ListResourceIconParams) (*ListResourceIconResponse, error) {
	resp, err := s.cs.newRequest("listResourceIcon", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListResourceIconResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListResourceIconResponse struct {
	Count        int             `json:"count"`
	ResourceIcon []*ResourceIcon `json:"resourceicon"`
}

type ResourceIcon struct {
	Base64image  string `json:"base64image"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Resourceid   string `json:"resourceid"`
	Resourcetype string `json:"resourcetype"`
}

type UploadResourceIconParams struct {
	p map[string]interface{}
}

func (p *UploadResourceIconParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["base64image"]; found {
		u.Set("base64image", v.(string))
	}
	if v, found := p.p["resourceids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("resourceids", vv)
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *UploadResourceIconParams) SetBase64image(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["base64image"] = v
}

func (p *UploadResourceIconParams) ResetBase64image() {
	if p.p != nil && p.p["base64image"] != nil {
		delete(p.p, "base64image")
	}
}

func (p *UploadResourceIconParams) GetBase64image() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["base64image"].(string)
	return value, ok
}

func (p *UploadResourceIconParams) SetResourceids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceids"] = v
}

func (p *UploadResourceIconParams) ResetResourceids() {
	if p.p != nil && p.p["resourceids"] != nil {
		delete(p.p, "resourceids")
	}
}

func (p *UploadResourceIconParams) GetResourceids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceids"].([]string)
	return value, ok
}

func (p *UploadResourceIconParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *UploadResourceIconParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *UploadResourceIconParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new UploadResourceIconParams instance,
// as then you are sure you have configured all required params
func (s *ResourceIconService) NewUploadResourceIconParams(base64image string, resourceids []string, resourcetype string) *UploadResourceIconParams {
	p := &UploadResourceIconParams{}
	p.p = make(map[string]interface{})
	p.p["base64image"] = base64image
	p.p["resourceids"] = resourceids
	p.p["resourcetype"] = resourcetype
	return p
}

// Uploads an icon for the specified resource(s)
func (s *ResourceIconService) UploadResourceIcon(p *UploadResourceIconParams) (*UploadResourceIconResponse, error) {
	resp, err := s.cs.newPostRequest("uploadResourceIcon", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadResourceIconResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UploadResourceIconResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UploadResourceIconResponse) UnmarshalJSON(b []byte) error {
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

	type alias UploadResourceIconResponse
	return json.Unmarshal(b, (*alias)(r))
}
