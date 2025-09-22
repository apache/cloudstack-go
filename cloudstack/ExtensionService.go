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

type ExtensionServiceIface interface {
	AddCustomAction(p *AddCustomActionParams) (*AddCustomActionResponse, error)
	NewAddCustomActionParams(extensionid string, name string) *AddCustomActionParams
	CreateExtension(p *CreateExtensionParams) (*CreateExtensionResponse, error)
	NewCreateExtensionParams(name string, extensionType string) *CreateExtensionParams
	DeleteCustomAction(p *DeleteCustomActionParams) (*DeleteCustomActionResponse, error)
	NewDeleteCustomActionParams() *DeleteCustomActionParams
	DeleteExtension(p *DeleteExtensionParams) (*DeleteExtensionResponse, error)
	NewDeleteExtensionParams() *DeleteExtensionParams
	ListCustomActions(p *ListCustomActionsParams) (*ListCustomActionsResponse, error)
	NewListCustomActionsParams() *ListCustomActionsParams
	GetCustomActionID(name string, opts ...OptionFunc) (string, int, error)
	GetCustomActionByName(name string, opts ...OptionFunc) (*CustomAction, int, error)
	GetCustomActionByID(id string, opts ...OptionFunc) (*CustomAction, int, error)
	ListExtensions(p *ListExtensionsParams) (*ListExtensionsResponse, error)
	NewListExtensionsParams() *ListExtensionsParams
	GetExtensionID(name string, opts ...OptionFunc) (string, int, error)
	GetExtensionByName(name string, opts ...OptionFunc) (*Extension, int, error)
	GetExtensionByID(id string, opts ...OptionFunc) (*Extension, int, error)
	RegisterExtension(p *RegisterExtensionParams) (*RegisterExtensionResponse, error)
	NewRegisterExtensionParams(extensionid string, resourceid string, resourcetype string) *RegisterExtensionParams
	RunCustomAction(p *RunCustomActionParams) (*RunCustomActionResponse, error)
	NewRunCustomActionParams(customactionid string, resourceid string) *RunCustomActionParams
	UnregisterExtension(p *UnregisterExtensionParams) (*UnregisterExtensionResponse, error)
	NewUnregisterExtensionParams(extensionid string, resourceid string, resourcetype string) *UnregisterExtensionParams
	UpdateCustomAction(p *UpdateCustomActionParams) (*UpdateCustomActionResponse, error)
	NewUpdateCustomActionParams(id string) *UpdateCustomActionParams
	UpdateExtension(p *UpdateExtensionParams) (*UpdateExtensionResponse, error)
	NewUpdateExtensionParams(id string) *UpdateExtensionParams
}

type AddCustomActionParams struct {
	p map[string]interface{}
}

func (p *AddCustomActionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allowedroletypes"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("allowedroletypes", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["errormessage"]; found {
		u.Set("errormessage", v.(string))
	}
	if v, found := p.p["extensionid"]; found {
		u.Set("extensionid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["parameters"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("parameters[%d].key", i), k)
			u.Set(fmt.Sprintf("parameters[%d].value", i), m[k])
		}
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := p.p["successmessage"]; found {
		u.Set("successmessage", v.(string))
	}
	if v, found := p.p["timeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("timeout", vv)
	}
	return u
}

func (p *AddCustomActionParams) SetAllowedroletypes(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allowedroletypes"] = v
}

func (p *AddCustomActionParams) ResetAllowedroletypes() {
	if p.p != nil && p.p["allowedroletypes"] != nil {
		delete(p.p, "allowedroletypes")
	}
}

func (p *AddCustomActionParams) GetAllowedroletypes() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allowedroletypes"].([]string)
	return value, ok
}

func (p *AddCustomActionParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *AddCustomActionParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *AddCustomActionParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *AddCustomActionParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *AddCustomActionParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *AddCustomActionParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *AddCustomActionParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *AddCustomActionParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *AddCustomActionParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *AddCustomActionParams) SetErrormessage(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["errormessage"] = v
}

func (p *AddCustomActionParams) ResetErrormessage() {
	if p.p != nil && p.p["errormessage"] != nil {
		delete(p.p, "errormessage")
	}
}

func (p *AddCustomActionParams) GetErrormessage() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["errormessage"].(string)
	return value, ok
}

func (p *AddCustomActionParams) SetExtensionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extensionid"] = v
}

func (p *AddCustomActionParams) ResetExtensionid() {
	if p.p != nil && p.p["extensionid"] != nil {
		delete(p.p, "extensionid")
	}
}

func (p *AddCustomActionParams) GetExtensionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extensionid"].(string)
	return value, ok
}

func (p *AddCustomActionParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddCustomActionParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddCustomActionParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddCustomActionParams) SetParameters(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parameters"] = v
}

func (p *AddCustomActionParams) ResetParameters() {
	if p.p != nil && p.p["parameters"] != nil {
		delete(p.p, "parameters")
	}
}

func (p *AddCustomActionParams) GetParameters() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parameters"].(map[string]string)
	return value, ok
}

func (p *AddCustomActionParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *AddCustomActionParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *AddCustomActionParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

func (p *AddCustomActionParams) SetSuccessmessage(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["successmessage"] = v
}

func (p *AddCustomActionParams) ResetSuccessmessage() {
	if p.p != nil && p.p["successmessage"] != nil {
		delete(p.p, "successmessage")
	}
}

func (p *AddCustomActionParams) GetSuccessmessage() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["successmessage"].(string)
	return value, ok
}

func (p *AddCustomActionParams) SetTimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timeout"] = v
}

func (p *AddCustomActionParams) ResetTimeout() {
	if p.p != nil && p.p["timeout"] != nil {
		delete(p.p, "timeout")
	}
}

func (p *AddCustomActionParams) GetTimeout() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timeout"].(int)
	return value, ok
}

// You should always use this function to get a new AddCustomActionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewAddCustomActionParams(extensionid string, name string) *AddCustomActionParams {
	p := &AddCustomActionParams{}
	p.p = make(map[string]interface{})
	p.p["extensionid"] = extensionid
	p.p["name"] = name
	return p
}

// Add a custom action for an extension
func (s *ExtensionService) AddCustomAction(p *AddCustomActionParams) (*AddCustomActionResponse, error) {
	resp, err := s.cs.newPostRequest("addCustomAction", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddCustomActionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddCustomActionResponse struct {
	Allowedroletypes []string                            `json:"allowedroletypes"`
	Created          string                              `json:"created"`
	Description      string                              `json:"description"`
	Details          map[string]string                   `json:"details"`
	Enabled          bool                                `json:"enabled"`
	Errormessage     string                              `json:"errormessage"`
	Extensionid      string                              `json:"extensionid"`
	Extensionname    string                              `json:"extensionname"`
	Id               string                              `json:"id"`
	JobID            string                              `json:"jobid"`
	Jobstatus        int                                 `json:"jobstatus"`
	Name             string                              `json:"name"`
	Parameters       []AddCustomActionResponseParameters `json:"parameters"`
	Resourcetype     string                              `json:"resourcetype"`
	Successmessage   string                              `json:"successmessage"`
	Timeout          int                                 `json:"timeout"`
}

type AddCustomActionResponseParameters struct {
	Name             string   `json:"name"`
	Required         bool     `json:"required"`
	Type             string   `json:"type"`
	Validationformat string   `json:"validationformat"`
	Valueoptions     []string `json:"valueoptions"`
}

type CreateExtensionParams struct {
	p map[string]interface{}
}

func (p *CreateExtensionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["orchestratorrequirespreparevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("orchestratorrequirespreparevm", vv)
	}
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *CreateExtensionParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateExtensionParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateExtensionParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateExtensionParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateExtensionParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *CreateExtensionParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *CreateExtensionParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateExtensionParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateExtensionParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateExtensionParams) SetOrchestratorrequirespreparevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["orchestratorrequirespreparevm"] = v
}

func (p *CreateExtensionParams) ResetOrchestratorrequirespreparevm() {
	if p.p != nil && p.p["orchestratorrequirespreparevm"] != nil {
		delete(p.p, "orchestratorrequirespreparevm")
	}
}

func (p *CreateExtensionParams) GetOrchestratorrequirespreparevm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["orchestratorrequirespreparevm"].(bool)
	return value, ok
}

func (p *CreateExtensionParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *CreateExtensionParams) ResetPath() {
	if p.p != nil && p.p["path"] != nil {
		delete(p.p, "path")
	}
}

func (p *CreateExtensionParams) GetPath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["path"].(string)
	return value, ok
}

func (p *CreateExtensionParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *CreateExtensionParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *CreateExtensionParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *CreateExtensionParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *CreateExtensionParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *CreateExtensionParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new CreateExtensionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewCreateExtensionParams(name string, extensionType string) *CreateExtensionParams {
	p := &CreateExtensionParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["type"] = extensionType
	return p
}

// Create an extension
func (s *ExtensionService) CreateExtension(p *CreateExtensionParams) (*CreateExtensionResponse, error) {
	resp, err := s.cs.newPostRequest("createExtension", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateExtensionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateExtensionResponse struct {
	Created       string                             `json:"created"`
	Description   string                             `json:"description"`
	Details       map[string]string                  `json:"details"`
	Id            string                             `json:"id"`
	Isuserdefined bool                               `json:"isuserdefined"`
	JobID         string                             `json:"jobid"`
	Jobstatus     int                                `json:"jobstatus"`
	Name          string                             `json:"name"`
	Path          string                             `json:"path"`
	Pathready     bool                               `json:"pathready"`
	Removed       string                             `json:"removed"`
	Resources     []CreateExtensionResponseResources `json:"resources"`
	State         string                             `json:"state"`
	Type          string                             `json:"type"`
}

type CreateExtensionResponseResources struct {
	Created string            `json:"created"`
	Details map[string]string `json:"details"`
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Type    string            `json:"type"`
}

type DeleteCustomActionParams struct {
	p map[string]interface{}
}

func (p *DeleteCustomActionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteCustomActionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteCustomActionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteCustomActionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteCustomActionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewDeleteCustomActionParams() *DeleteCustomActionParams {
	p := &DeleteCustomActionParams{}
	p.p = make(map[string]interface{})
	return p
}

// Delete the custom action
func (s *ExtensionService) DeleteCustomAction(p *DeleteCustomActionParams) (*DeleteCustomActionResponse, error) {
	resp, err := s.cs.newPostRequest("deleteCustomAction", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteCustomActionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteCustomActionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteCustomActionResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteCustomActionResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteExtensionParams struct {
	p map[string]interface{}
}

func (p *DeleteExtensionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteExtensionParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *DeleteExtensionParams) ResetCleanup() {
	if p.p != nil && p.p["cleanup"] != nil {
		delete(p.p, "cleanup")
	}
}

func (p *DeleteExtensionParams) GetCleanup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanup"].(bool)
	return value, ok
}

func (p *DeleteExtensionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteExtensionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteExtensionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteExtensionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewDeleteExtensionParams() *DeleteExtensionParams {
	p := &DeleteExtensionParams{}
	p.p = make(map[string]interface{})
	return p
}

// Delete the extensions
func (s *ExtensionService) DeleteExtension(p *DeleteExtensionParams) (*DeleteExtensionResponse, error) {
	resp, err := s.cs.newPostRequest("deleteExtension", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteExtensionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteExtensionResponse struct {
	Created       string                             `json:"created"`
	Description   string                             `json:"description"`
	Details       map[string]string                  `json:"details"`
	Id            string                             `json:"id"`
	Isuserdefined bool                               `json:"isuserdefined"`
	JobID         string                             `json:"jobid"`
	Jobstatus     int                                `json:"jobstatus"`
	Name          string                             `json:"name"`
	Path          string                             `json:"path"`
	Pathready     bool                               `json:"pathready"`
	Removed       string                             `json:"removed"`
	Resources     []DeleteExtensionResponseResources `json:"resources"`
	State         string                             `json:"state"`
	Type          string                             `json:"type"`
}

type DeleteExtensionResponseResources struct {
	Created string            `json:"created"`
	Details map[string]string `json:"details"`
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Type    string            `json:"type"`
}

type ListCustomActionsParams struct {
	p map[string]interface{}
}

func (p *ListCustomActionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["extensionid"]; found {
		u.Set("extensionid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *ListCustomActionsParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *ListCustomActionsParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *ListCustomActionsParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *ListCustomActionsParams) SetExtensionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extensionid"] = v
}

func (p *ListCustomActionsParams) ResetExtensionid() {
	if p.p != nil && p.p["extensionid"] != nil {
		delete(p.p, "extensionid")
	}
}

func (p *ListCustomActionsParams) GetExtensionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extensionid"].(string)
	return value, ok
}

func (p *ListCustomActionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListCustomActionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListCustomActionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListCustomActionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListCustomActionsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListCustomActionsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListCustomActionsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListCustomActionsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListCustomActionsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListCustomActionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListCustomActionsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListCustomActionsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListCustomActionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListCustomActionsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListCustomActionsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListCustomActionsParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *ListCustomActionsParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *ListCustomActionsParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *ListCustomActionsParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *ListCustomActionsParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *ListCustomActionsParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new ListCustomActionsParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewListCustomActionsParams() *ListCustomActionsParams {
	p := &ListCustomActionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ExtensionService) GetCustomActionID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListCustomActionsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListCustomActions(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.CustomActions[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.CustomActions {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ExtensionService) GetCustomActionByName(name string, opts ...OptionFunc) (*CustomAction, int, error) {
	id, count, err := s.GetCustomActionID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetCustomActionByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ExtensionService) GetCustomActionByID(id string, opts ...OptionFunc) (*CustomAction, int, error) {
	p := &ListCustomActionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListCustomActions(p)
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
		return l.CustomActions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for CustomAction UUID: %s!", id)
}

// Lists the custom actions
func (s *ExtensionService) ListCustomActions(p *ListCustomActionsParams) (*ListCustomActionsResponse, error) {
	resp, err := s.cs.newRequest("listCustomActions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCustomActionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCustomActionsResponse struct {
	Count         int             `json:"count"`
	CustomActions []*CustomAction `json:"customaction"`
}

type CustomAction struct {
	Allowedroletypes []string                 `json:"allowedroletypes"`
	Created          string                   `json:"created"`
	Description      string                   `json:"description"`
	Details          map[string]string        `json:"details"`
	Enabled          bool                     `json:"enabled"`
	Errormessage     string                   `json:"errormessage"`
	Extensionid      string                   `json:"extensionid"`
	Extensionname    string                   `json:"extensionname"`
	Id               string                   `json:"id"`
	JobID            string                   `json:"jobid"`
	Jobstatus        int                      `json:"jobstatus"`
	Name             string                   `json:"name"`
	Parameters       []CustomActionParameters `json:"parameters"`
	Resourcetype     string                   `json:"resourcetype"`
	Successmessage   string                   `json:"successmessage"`
	Timeout          int                      `json:"timeout"`
}

type CustomActionParameters struct {
	Name             string   `json:"name"`
	Required         bool     `json:"required"`
	Type             string   `json:"type"`
	Validationformat string   `json:"validationformat"`
	Valueoptions     []string `json:"valueoptions"`
}

type ListExtensionsParams struct {
	p map[string]interface{}
}

func (p *ListExtensionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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

func (p *ListExtensionsParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListExtensionsParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListExtensionsParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListExtensionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListExtensionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListExtensionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListExtensionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListExtensionsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListExtensionsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListExtensionsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListExtensionsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListExtensionsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListExtensionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListExtensionsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListExtensionsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListExtensionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListExtensionsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListExtensionsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListExtensionsParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewListExtensionsParams() *ListExtensionsParams {
	p := &ListExtensionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ExtensionService) GetExtensionID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListExtensionsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListExtensions(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Extensions[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Extensions {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ExtensionService) GetExtensionByName(name string, opts ...OptionFunc) (*Extension, int, error) {
	id, count, err := s.GetExtensionID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetExtensionByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ExtensionService) GetExtensionByID(id string, opts ...OptionFunc) (*Extension, int, error) {
	p := &ListExtensionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListExtensions(p)
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
		return l.Extensions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Extension UUID: %s!", id)
}

// Lists extensions
func (s *ExtensionService) ListExtensions(p *ListExtensionsParams) (*ListExtensionsResponse, error) {
	resp, err := s.cs.newRequest("listExtensions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListExtensionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListExtensionsResponse struct {
	Count      int          `json:"count"`
	Extensions []*Extension `json:"extension"`
}

type Extension struct {
	Created       string               `json:"created"`
	Description   string               `json:"description"`
	Details       map[string]string    `json:"details"`
	Id            string               `json:"id"`
	Isuserdefined bool                 `json:"isuserdefined"`
	JobID         string               `json:"jobid"`
	Jobstatus     int                  `json:"jobstatus"`
	Name          string               `json:"name"`
	Path          string               `json:"path"`
	Pathready     bool                 `json:"pathready"`
	Removed       string               `json:"removed"`
	Resources     []ExtensionResources `json:"resources"`
	State         string               `json:"state"`
	Type          string               `json:"type"`
}

type ExtensionResources struct {
	Created string            `json:"created"`
	Details map[string]string `json:"details"`
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Type    string            `json:"type"`
}

type RegisterExtensionParams struct {
	p map[string]interface{}
}

func (p *RegisterExtensionParams) toURLValues() url.Values {
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
	if v, found := p.p["extensionid"]; found {
		u.Set("extensionid", v.(string))
	}
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *RegisterExtensionParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *RegisterExtensionParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *RegisterExtensionParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *RegisterExtensionParams) SetExtensionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extensionid"] = v
}

func (p *RegisterExtensionParams) ResetExtensionid() {
	if p.p != nil && p.p["extensionid"] != nil {
		delete(p.p, "extensionid")
	}
}

func (p *RegisterExtensionParams) GetExtensionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extensionid"].(string)
	return value, ok
}

func (p *RegisterExtensionParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *RegisterExtensionParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *RegisterExtensionParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *RegisterExtensionParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *RegisterExtensionParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *RegisterExtensionParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterExtensionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewRegisterExtensionParams(extensionid string, resourceid string, resourcetype string) *RegisterExtensionParams {
	p := &RegisterExtensionParams{}
	p.p = make(map[string]interface{})
	p.p["extensionid"] = extensionid
	p.p["resourceid"] = resourceid
	p.p["resourcetype"] = resourcetype
	return p
}

// Register an extension with a resource
func (s *ExtensionService) RegisterExtension(p *RegisterExtensionParams) (*RegisterExtensionResponse, error) {
	resp, err := s.cs.newPostRequest("registerExtension", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterExtensionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterExtensionResponse struct {
	Created       string                               `json:"created"`
	Description   string                               `json:"description"`
	Details       map[string]string                    `json:"details"`
	Id            string                               `json:"id"`
	Isuserdefined bool                                 `json:"isuserdefined"`
	JobID         string                               `json:"jobid"`
	Jobstatus     int                                  `json:"jobstatus"`
	Name          string                               `json:"name"`
	Path          string                               `json:"path"`
	Pathready     bool                                 `json:"pathready"`
	Removed       string                               `json:"removed"`
	Resources     []RegisterExtensionResponseResources `json:"resources"`
	State         string                               `json:"state"`
	Type          string                               `json:"type"`
}

type RegisterExtensionResponseResources struct {
	Created string            `json:"created"`
	Details map[string]string `json:"details"`
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Type    string            `json:"type"`
}

type RunCustomActionParams struct {
	p map[string]interface{}
}

func (p *RunCustomActionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customactionid"]; found {
		u.Set("customactionid", v.(string))
	}
	if v, found := p.p["parameters"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("parameters[%d].key", i), k)
			u.Set(fmt.Sprintf("parameters[%d].value", i), m[k])
		}
	}
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *RunCustomActionParams) SetCustomactionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customactionid"] = v
}

func (p *RunCustomActionParams) ResetCustomactionid() {
	if p.p != nil && p.p["customactionid"] != nil {
		delete(p.p, "customactionid")
	}
}

func (p *RunCustomActionParams) GetCustomactionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customactionid"].(string)
	return value, ok
}

func (p *RunCustomActionParams) SetParameters(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parameters"] = v
}

func (p *RunCustomActionParams) ResetParameters() {
	if p.p != nil && p.p["parameters"] != nil {
		delete(p.p, "parameters")
	}
}

func (p *RunCustomActionParams) GetParameters() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parameters"].(map[string]string)
	return value, ok
}

func (p *RunCustomActionParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *RunCustomActionParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *RunCustomActionParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *RunCustomActionParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *RunCustomActionParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *RunCustomActionParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new RunCustomActionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewRunCustomActionParams(customactionid string, resourceid string) *RunCustomActionParams {
	p := &RunCustomActionParams{}
	p.p = make(map[string]interface{})
	p.p["customactionid"] = customactionid
	p.p["resourceid"] = resourceid
	return p
}

// Run the custom action
func (s *ExtensionService) RunCustomAction(p *RunCustomActionParams) (*RunCustomActionResponse, error) {
	resp, err := s.cs.newPostRequest("runCustomAction", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RunCustomActionResponse
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

type RunCustomActionResponse struct {
	Id        string            `json:"id"`
	JobID     string            `json:"jobid"`
	Jobstatus int               `json:"jobstatus"`
	Name      string            `json:"name"`
	Result    map[string]string `json:"result"`
	Success   bool              `json:"success"`
}

type UnregisterExtensionParams struct {
	p map[string]interface{}
}

func (p *UnregisterExtensionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["extensionid"]; found {
		u.Set("extensionid", v.(string))
	}
	if v, found := p.p["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (p *UnregisterExtensionParams) SetExtensionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extensionid"] = v
}

func (p *UnregisterExtensionParams) ResetExtensionid() {
	if p.p != nil && p.p["extensionid"] != nil {
		delete(p.p, "extensionid")
	}
}

func (p *UnregisterExtensionParams) GetExtensionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extensionid"].(string)
	return value, ok
}

func (p *UnregisterExtensionParams) SetResourceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourceid"] = v
}

func (p *UnregisterExtensionParams) ResetResourceid() {
	if p.p != nil && p.p["resourceid"] != nil {
		delete(p.p, "resourceid")
	}
}

func (p *UnregisterExtensionParams) GetResourceid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourceid"].(string)
	return value, ok
}

func (p *UnregisterExtensionParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *UnregisterExtensionParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *UnregisterExtensionParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new UnregisterExtensionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewUnregisterExtensionParams(extensionid string, resourceid string, resourcetype string) *UnregisterExtensionParams {
	p := &UnregisterExtensionParams{}
	p.p = make(map[string]interface{})
	p.p["extensionid"] = extensionid
	p.p["resourceid"] = resourceid
	p.p["resourcetype"] = resourcetype
	return p
}

// Unregister an extension with a resource
func (s *ExtensionService) UnregisterExtension(p *UnregisterExtensionParams) (*UnregisterExtensionResponse, error) {
	resp, err := s.cs.newPostRequest("unregisterExtension", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UnregisterExtensionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UnregisterExtensionResponse struct {
	Created       string                                 `json:"created"`
	Description   string                                 `json:"description"`
	Details       map[string]string                      `json:"details"`
	Id            string                                 `json:"id"`
	Isuserdefined bool                                   `json:"isuserdefined"`
	JobID         string                                 `json:"jobid"`
	Jobstatus     int                                    `json:"jobstatus"`
	Name          string                                 `json:"name"`
	Path          string                                 `json:"path"`
	Pathready     bool                                   `json:"pathready"`
	Removed       string                                 `json:"removed"`
	Resources     []UnregisterExtensionResponseResources `json:"resources"`
	State         string                                 `json:"state"`
	Type          string                                 `json:"type"`
}

type UnregisterExtensionResponseResources struct {
	Created string            `json:"created"`
	Details map[string]string `json:"details"`
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Type    string            `json:"type"`
}

type UpdateCustomActionParams struct {
	p map[string]interface{}
}

func (p *UpdateCustomActionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allowedroletypes"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("allowedroletypes", vv)
	}
	if v, found := p.p["cleanupdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupdetails", vv)
	}
	if v, found := p.p["cleanupparameters"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupparameters", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["errormessage"]; found {
		u.Set("errormessage", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["parameters"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("parameters[%d].key", i), k)
			u.Set(fmt.Sprintf("parameters[%d].value", i), m[k])
		}
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := p.p["successmessage"]; found {
		u.Set("successmessage", v.(string))
	}
	if v, found := p.p["timeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("timeout", vv)
	}
	return u
}

func (p *UpdateCustomActionParams) SetAllowedroletypes(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allowedroletypes"] = v
}

func (p *UpdateCustomActionParams) ResetAllowedroletypes() {
	if p.p != nil && p.p["allowedroletypes"] != nil {
		delete(p.p, "allowedroletypes")
	}
}

func (p *UpdateCustomActionParams) GetAllowedroletypes() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allowedroletypes"].([]string)
	return value, ok
}

func (p *UpdateCustomActionParams) SetCleanupdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupdetails"] = v
}

func (p *UpdateCustomActionParams) ResetCleanupdetails() {
	if p.p != nil && p.p["cleanupdetails"] != nil {
		delete(p.p, "cleanupdetails")
	}
}

func (p *UpdateCustomActionParams) GetCleanupdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupdetails"].(bool)
	return value, ok
}

func (p *UpdateCustomActionParams) SetCleanupparameters(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupparameters"] = v
}

func (p *UpdateCustomActionParams) ResetCleanupparameters() {
	if p.p != nil && p.p["cleanupparameters"] != nil {
		delete(p.p, "cleanupparameters")
	}
}

func (p *UpdateCustomActionParams) GetCleanupparameters() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupparameters"].(bool)
	return value, ok
}

func (p *UpdateCustomActionParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateCustomActionParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateCustomActionParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateCustomActionParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateCustomActionParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateCustomActionParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateCustomActionParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *UpdateCustomActionParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *UpdateCustomActionParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *UpdateCustomActionParams) SetErrormessage(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["errormessage"] = v
}

func (p *UpdateCustomActionParams) ResetErrormessage() {
	if p.p != nil && p.p["errormessage"] != nil {
		delete(p.p, "errormessage")
	}
}

func (p *UpdateCustomActionParams) GetErrormessage() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["errormessage"].(string)
	return value, ok
}

func (p *UpdateCustomActionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateCustomActionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateCustomActionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateCustomActionParams) SetParameters(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["parameters"] = v
}

func (p *UpdateCustomActionParams) ResetParameters() {
	if p.p != nil && p.p["parameters"] != nil {
		delete(p.p, "parameters")
	}
}

func (p *UpdateCustomActionParams) GetParameters() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["parameters"].(map[string]string)
	return value, ok
}

func (p *UpdateCustomActionParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *UpdateCustomActionParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *UpdateCustomActionParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

func (p *UpdateCustomActionParams) SetSuccessmessage(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["successmessage"] = v
}

func (p *UpdateCustomActionParams) ResetSuccessmessage() {
	if p.p != nil && p.p["successmessage"] != nil {
		delete(p.p, "successmessage")
	}
}

func (p *UpdateCustomActionParams) GetSuccessmessage() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["successmessage"].(string)
	return value, ok
}

func (p *UpdateCustomActionParams) SetTimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timeout"] = v
}

func (p *UpdateCustomActionParams) ResetTimeout() {
	if p.p != nil && p.p["timeout"] != nil {
		delete(p.p, "timeout")
	}
}

func (p *UpdateCustomActionParams) GetTimeout() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timeout"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateCustomActionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewUpdateCustomActionParams(id string) *UpdateCustomActionParams {
	p := &UpdateCustomActionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Update the custom action
func (s *ExtensionService) UpdateCustomAction(p *UpdateCustomActionParams) (*UpdateCustomActionResponse, error) {
	resp, err := s.cs.newPostRequest("updateCustomAction", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateCustomActionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateCustomActionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateCustomActionResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateCustomActionResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateExtensionParams struct {
	p map[string]interface{}
}

func (p *UpdateExtensionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanupdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupdetails", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
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
	if v, found := p.p["orchestratorrequirespreparevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("orchestratorrequirespreparevm", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *UpdateExtensionParams) SetCleanupdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanupdetails"] = v
}

func (p *UpdateExtensionParams) ResetCleanupdetails() {
	if p.p != nil && p.p["cleanupdetails"] != nil {
		delete(p.p, "cleanupdetails")
	}
}

func (p *UpdateExtensionParams) GetCleanupdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanupdetails"].(bool)
	return value, ok
}

func (p *UpdateExtensionParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateExtensionParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateExtensionParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateExtensionParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateExtensionParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateExtensionParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateExtensionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateExtensionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateExtensionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateExtensionParams) SetOrchestratorrequirespreparevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["orchestratorrequirespreparevm"] = v
}

func (p *UpdateExtensionParams) ResetOrchestratorrequirespreparevm() {
	if p.p != nil && p.p["orchestratorrequirespreparevm"] != nil {
		delete(p.p, "orchestratorrequirespreparevm")
	}
}

func (p *UpdateExtensionParams) GetOrchestratorrequirespreparevm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["orchestratorrequirespreparevm"].(bool)
	return value, ok
}

func (p *UpdateExtensionParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdateExtensionParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *UpdateExtensionParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateExtensionParams instance,
// as then you are sure you have configured all required params
func (s *ExtensionService) NewUpdateExtensionParams(id string) *UpdateExtensionParams {
	p := &UpdateExtensionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Update the extension
func (s *ExtensionService) UpdateExtension(p *UpdateExtensionParams) (*UpdateExtensionResponse, error) {
	resp, err := s.cs.newPostRequest("updateExtension", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateExtensionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateExtensionResponse struct {
	Created       string                             `json:"created"`
	Description   string                             `json:"description"`
	Details       map[string]string                  `json:"details"`
	Id            string                             `json:"id"`
	Isuserdefined bool                               `json:"isuserdefined"`
	JobID         string                             `json:"jobid"`
	Jobstatus     int                                `json:"jobstatus"`
	Name          string                             `json:"name"`
	Path          string                             `json:"path"`
	Pathready     bool                               `json:"pathready"`
	Removed       string                             `json:"removed"`
	Resources     []UpdateExtensionResponseResources `json:"resources"`
	State         string                             `json:"state"`
	Type          string                             `json:"type"`
}

type UpdateExtensionResponseResources struct {
	Created string            `json:"created"`
	Details map[string]string `json:"details"`
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Type    string            `json:"type"`
}
