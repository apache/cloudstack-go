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

type AnnotationServiceIface interface {
	AddAnnotation(p *AddAnnotationParams) (*AddAnnotationResponse, error)
	NewAddAnnotationParams() *AddAnnotationParams
	ListAnnotations(p *ListAnnotationsParams) (*ListAnnotationsResponse, error)
	NewListAnnotationsParams() *ListAnnotationsParams
	GetAnnotationByID(id string, opts ...OptionFunc) (*Annotation, int, error)
	RemoveAnnotation(p *RemoveAnnotationParams) (*RemoveAnnotationResponse, error)
	NewRemoveAnnotationParams(id string) *RemoveAnnotationParams
	UpdateAnnotationVisibility(p *UpdateAnnotationVisibilityParams) (*UpdateAnnotationVisibilityResponse, error)
	NewUpdateAnnotationVisibilityParams(adminsonly bool, id string) *UpdateAnnotationVisibilityParams
}

type AddAnnotationParams struct {
	p map[string]interface{}
}

func (p *AddAnnotationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["adminsonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("adminsonly", vv)
	}
	if v, found := p.p["annotation"]; found {
		u.Set("annotation", v.(string))
	}
	if v, found := p.p["entityid"]; found {
		u.Set("entityid", v.(string))
	}
	if v, found := p.p["entitytype"]; found {
		u.Set("entitytype", v.(string))
	}
	return u
}

func (p *AddAnnotationParams) SetAdminsonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["adminsonly"] = v
}

func (p *AddAnnotationParams) ResetAdminsonly() {
	if p.p != nil && p.p["adminsonly"] != nil {
		delete(p.p, "adminsonly")
	}
}

func (p *AddAnnotationParams) GetAdminsonly() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["adminsonly"].(bool)
	return value, ok
}

func (p *AddAnnotationParams) SetAnnotation(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["annotation"] = v
}

func (p *AddAnnotationParams) ResetAnnotation() {
	if p.p != nil && p.p["annotation"] != nil {
		delete(p.p, "annotation")
	}
}

func (p *AddAnnotationParams) GetAnnotation() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["annotation"].(string)
	return value, ok
}

func (p *AddAnnotationParams) SetEntityid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["entityid"] = v
}

func (p *AddAnnotationParams) ResetEntityid() {
	if p.p != nil && p.p["entityid"] != nil {
		delete(p.p, "entityid")
	}
}

func (p *AddAnnotationParams) GetEntityid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["entityid"].(string)
	return value, ok
}

func (p *AddAnnotationParams) SetEntitytype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["entitytype"] = v
}

func (p *AddAnnotationParams) ResetEntitytype() {
	if p.p != nil && p.p["entitytype"] != nil {
		delete(p.p, "entitytype")
	}
}

func (p *AddAnnotationParams) GetEntitytype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["entitytype"].(string)
	return value, ok
}

// You should always use this function to get a new AddAnnotationParams instance,
// as then you are sure you have configured all required params
func (s *AnnotationService) NewAddAnnotationParams() *AddAnnotationParams {
	p := &AddAnnotationParams{}
	p.p = make(map[string]interface{})
	return p
}

// add an annotation.
func (s *AnnotationService) AddAnnotation(p *AddAnnotationParams) (*AddAnnotationResponse, error) {
	resp, err := s.cs.newRequest("addAnnotation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r AddAnnotationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddAnnotationResponse struct {
	Adminsonly bool   `json:"adminsonly"`
	Annotation string `json:"annotation"`
	Created    string `json:"created"`
	Entityid   string `json:"entityid"`
	Entityname string `json:"entityname"`
	Entitytype string `json:"entitytype"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Removed    string `json:"removed"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
}

type ListAnnotationsParams struct {
	p map[string]interface{}
}

func (p *ListAnnotationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["annotationfilter"]; found {
		u.Set("annotationfilter", v.(string))
	}
	if v, found := p.p["entityid"]; found {
		u.Set("entityid", v.(string))
	}
	if v, found := p.p["entitytype"]; found {
		u.Set("entitytype", v.(string))
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
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (p *ListAnnotationsParams) SetAnnotationfilter(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["annotationfilter"] = v
}

func (p *ListAnnotationsParams) ResetAnnotationfilter() {
	if p.p != nil && p.p["annotationfilter"] != nil {
		delete(p.p, "annotationfilter")
	}
}

func (p *ListAnnotationsParams) GetAnnotationfilter() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["annotationfilter"].(string)
	return value, ok
}

func (p *ListAnnotationsParams) SetEntityid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["entityid"] = v
}

func (p *ListAnnotationsParams) ResetEntityid() {
	if p.p != nil && p.p["entityid"] != nil {
		delete(p.p, "entityid")
	}
}

func (p *ListAnnotationsParams) GetEntityid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["entityid"].(string)
	return value, ok
}

func (p *ListAnnotationsParams) SetEntitytype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["entitytype"] = v
}

func (p *ListAnnotationsParams) ResetEntitytype() {
	if p.p != nil && p.p["entitytype"] != nil {
		delete(p.p, "entitytype")
	}
}

func (p *ListAnnotationsParams) GetEntitytype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["entitytype"].(string)
	return value, ok
}

func (p *ListAnnotationsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListAnnotationsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListAnnotationsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListAnnotationsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAnnotationsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListAnnotationsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListAnnotationsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAnnotationsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListAnnotationsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListAnnotationsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAnnotationsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListAnnotationsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListAnnotationsParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *ListAnnotationsParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *ListAnnotationsParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

// You should always use this function to get a new ListAnnotationsParams instance,
// as then you are sure you have configured all required params
func (s *AnnotationService) NewListAnnotationsParams() *ListAnnotationsParams {
	p := &ListAnnotationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AnnotationService) GetAnnotationByID(id string, opts ...OptionFunc) (*Annotation, int, error) {
	p := &ListAnnotationsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAnnotations(p)
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
		return l.Annotations[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Annotation UUID: %s!", id)
}

// Lists annotations.
func (s *AnnotationService) ListAnnotations(p *ListAnnotationsParams) (*ListAnnotationsResponse, error) {
	resp, err := s.cs.newRequest("listAnnotations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAnnotationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAnnotationsResponse struct {
	Count       int           `json:"count"`
	Annotations []*Annotation `json:"annotation"`
}

type Annotation struct {
	Adminsonly bool   `json:"adminsonly"`
	Annotation string `json:"annotation"`
	Created    string `json:"created"`
	Entityid   string `json:"entityid"`
	Entityname string `json:"entityname"`
	Entitytype string `json:"entitytype"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Removed    string `json:"removed"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
}

type RemoveAnnotationParams struct {
	p map[string]interface{}
}

func (p *RemoveAnnotationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RemoveAnnotationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RemoveAnnotationParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RemoveAnnotationParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveAnnotationParams instance,
// as then you are sure you have configured all required params
func (s *AnnotationService) NewRemoveAnnotationParams(id string) *RemoveAnnotationParams {
	p := &RemoveAnnotationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// remove an annotation.
func (s *AnnotationService) RemoveAnnotation(p *RemoveAnnotationParams) (*RemoveAnnotationResponse, error) {
	resp, err := s.cs.newRequest("removeAnnotation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r RemoveAnnotationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveAnnotationResponse struct {
	Adminsonly bool   `json:"adminsonly"`
	Annotation string `json:"annotation"`
	Created    string `json:"created"`
	Entityid   string `json:"entityid"`
	Entityname string `json:"entityname"`
	Entitytype string `json:"entitytype"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Removed    string `json:"removed"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
}

type UpdateAnnotationVisibilityParams struct {
	p map[string]interface{}
}

func (p *UpdateAnnotationVisibilityParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["adminsonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("adminsonly", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateAnnotationVisibilityParams) SetAdminsonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["adminsonly"] = v
}

func (p *UpdateAnnotationVisibilityParams) ResetAdminsonly() {
	if p.p != nil && p.p["adminsonly"] != nil {
		delete(p.p, "adminsonly")
	}
}

func (p *UpdateAnnotationVisibilityParams) GetAdminsonly() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["adminsonly"].(bool)
	return value, ok
}

func (p *UpdateAnnotationVisibilityParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateAnnotationVisibilityParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateAnnotationVisibilityParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateAnnotationVisibilityParams instance,
// as then you are sure you have configured all required params
func (s *AnnotationService) NewUpdateAnnotationVisibilityParams(adminsonly bool, id string) *UpdateAnnotationVisibilityParams {
	p := &UpdateAnnotationVisibilityParams{}
	p.p = make(map[string]interface{})
	p.p["adminsonly"] = adminsonly
	p.p["id"] = id
	return p
}

// update an annotation visibility.
func (s *AnnotationService) UpdateAnnotationVisibility(p *UpdateAnnotationVisibilityParams) (*UpdateAnnotationVisibilityResponse, error) {
	resp, err := s.cs.newRequest("updateAnnotationVisibility", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAnnotationVisibilityResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateAnnotationVisibilityResponse struct {
	Adminsonly bool   `json:"adminsonly"`
	Annotation string `json:"annotation"`
	Created    string `json:"created"`
	Entityid   string `json:"entityid"`
	Entityname string `json:"entityname"`
	Entitytype string `json:"entitytype"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Removed    string `json:"removed"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
}
