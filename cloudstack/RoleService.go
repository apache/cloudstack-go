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

type RoleServiceIface interface {
	CreateRole(p *CreateRoleParams) (*CreateRoleResponse, error)
	NewCreateRoleParams(name string) *CreateRoleParams
	CreateRolePermission(p *CreateRolePermissionParams) (*CreateRolePermissionResponse, error)
	NewCreateRolePermissionParams(permission string, roleid string, rule string) *CreateRolePermissionParams
	DeleteRole(p *DeleteRoleParams) (*DeleteRoleResponse, error)
	NewDeleteRoleParams(id string) *DeleteRoleParams
	DeleteRolePermission(p *DeleteRolePermissionParams) (*DeleteRolePermissionResponse, error)
	NewDeleteRolePermissionParams(id string) *DeleteRolePermissionParams
	ImportRole(p *ImportRoleParams) (*ImportRoleResponse, error)
	NewImportRoleParams(name string, rules map[string]string) *ImportRoleParams
	ListRolePermissions(p *ListRolePermissionsParams) (*ListRolePermissionsResponse, error)
	NewListRolePermissionsParams() *ListRolePermissionsParams
	ListRoles(p *ListRolesParams) (*ListRolesResponse, error)
	NewListRolesParams() *ListRolesParams
	GetRoleID(name string, opts ...OptionFunc) (string, int, error)
	GetRoleByName(name string, opts ...OptionFunc) (*Role, int, error)
	GetRoleByID(id string, opts ...OptionFunc) (*Role, int, error)
	UpdateRole(p *UpdateRoleParams) (*UpdateRoleResponse, error)
	NewUpdateRoleParams(id string) *UpdateRoleParams
	UpdateRolePermission(p *UpdateRolePermissionParams) (*UpdateRolePermissionResponse, error)
	NewUpdateRolePermissionParams(roleid string) *UpdateRolePermissionParams
}

type CreateRoleParams struct {
	p map[string]interface{}
}

func (p *CreateRoleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *CreateRoleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateRoleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateRoleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateRoleParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *CreateRoleParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *CreateRoleParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *CreateRoleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateRoleParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateRoleParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateRoleParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *CreateRoleParams) ResetRoleid() {
	if p.p != nil && p.p["roleid"] != nil {
		delete(p.p, "roleid")
	}
}

func (p *CreateRoleParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

func (p *CreateRoleParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *CreateRoleParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *CreateRoleParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new CreateRoleParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewCreateRoleParams(name string) *CreateRoleParams {
	p := &CreateRoleParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Creates a role
func (s *RoleService) CreateRole(p *CreateRoleParams) (*CreateRoleResponse, error) {
	resp, err := s.cs.newRequest("createRole", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response CreateRoleResponse `json:"role"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type CreateRoleResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Isdefault   bool   `json:"isdefault"`
	Ispublic    bool   `json:"ispublic"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type CreateRolePermissionParams struct {
	p map[string]interface{}
}

func (p *CreateRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["permission"]; found {
		u.Set("permission", v.(string))
	}
	if v, found := p.p["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := p.p["rule"]; found {
		u.Set("rule", v.(string))
	}
	return u
}

func (p *CreateRolePermissionParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateRolePermissionParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateRolePermissionParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateRolePermissionParams) SetPermission(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["permission"] = v
}

func (p *CreateRolePermissionParams) ResetPermission() {
	if p.p != nil && p.p["permission"] != nil {
		delete(p.p, "permission")
	}
}

func (p *CreateRolePermissionParams) GetPermission() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["permission"].(string)
	return value, ok
}

func (p *CreateRolePermissionParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *CreateRolePermissionParams) ResetRoleid() {
	if p.p != nil && p.p["roleid"] != nil {
		delete(p.p, "roleid")
	}
}

func (p *CreateRolePermissionParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

func (p *CreateRolePermissionParams) SetRule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rule"] = v
}

func (p *CreateRolePermissionParams) ResetRule() {
	if p.p != nil && p.p["rule"] != nil {
		delete(p.p, "rule")
	}
}

func (p *CreateRolePermissionParams) GetRule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rule"].(string)
	return value, ok
}

// You should always use this function to get a new CreateRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewCreateRolePermissionParams(permission string, roleid string, rule string) *CreateRolePermissionParams {
	p := &CreateRolePermissionParams{}
	p.p = make(map[string]interface{})
	p.p["permission"] = permission
	p.p["roleid"] = roleid
	p.p["rule"] = rule
	return p
}

// Adds an API permission to a role
func (s *RoleService) CreateRolePermission(p *CreateRolePermissionParams) (*CreateRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("createRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response CreateRolePermissionResponse `json:"rolepermission"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type CreateRolePermissionResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Permission  string `json:"permission"`
	Roleid      string `json:"roleid"`
	Rolename    string `json:"rolename"`
	Rule        string `json:"rule"`
}

type DeleteRoleParams struct {
	p map[string]interface{}
}

func (p *DeleteRoleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteRoleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteRoleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteRoleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteRoleParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewDeleteRoleParams(id string) *DeleteRoleParams {
	p := &DeleteRoleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a role
func (s *RoleService) DeleteRole(p *DeleteRoleParams) (*DeleteRoleResponse, error) {
	resp, err := s.cs.newRequest("deleteRole", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteRoleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteRoleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteRoleResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteRoleResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteRolePermissionParams struct {
	p map[string]interface{}
}

func (p *DeleteRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteRolePermissionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteRolePermissionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteRolePermissionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewDeleteRolePermissionParams(id string) *DeleteRolePermissionParams {
	p := &DeleteRolePermissionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a role permission
func (s *RoleService) DeleteRolePermission(p *DeleteRolePermissionParams) (*DeleteRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("deleteRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteRolePermissionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteRolePermissionResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteRolePermissionResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ImportRoleParams struct {
	p map[string]interface{}
}

func (p *ImportRoleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["rules"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("rules[%d].key", i), k)
			u.Set(fmt.Sprintf("rules[%d].value", i), m[k])
		}
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *ImportRoleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *ImportRoleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *ImportRoleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *ImportRoleParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *ImportRoleParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *ImportRoleParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *ImportRoleParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *ImportRoleParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *ImportRoleParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *ImportRoleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ImportRoleParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ImportRoleParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ImportRoleParams) SetRules(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rules"] = v
}

func (p *ImportRoleParams) ResetRules() {
	if p.p != nil && p.p["rules"] != nil {
		delete(p.p, "rules")
	}
}

func (p *ImportRoleParams) GetRules() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rules"].(map[string]string)
	return value, ok
}

func (p *ImportRoleParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ImportRoleParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ImportRoleParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new ImportRoleParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewImportRoleParams(name string, rules map[string]string) *ImportRoleParams {
	p := &ImportRoleParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["rules"] = rules
	return p
}

// Imports a role based on provided map of rule permissions
func (s *RoleService) ImportRole(p *ImportRoleParams) (*ImportRoleResponse, error) {
	resp, err := s.cs.newRequest("importRole", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportRoleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ImportRoleResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Isdefault   bool   `json:"isdefault"`
	Ispublic    bool   `json:"ispublic"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type ListRolePermissionsParams struct {
	p map[string]interface{}
}

func (p *ListRolePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	return u
}

func (p *ListRolePermissionsParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *ListRolePermissionsParams) ResetRoleid() {
	if p.p != nil && p.p["roleid"] != nil {
		delete(p.p, "roleid")
	}
}

func (p *ListRolePermissionsParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

// You should always use this function to get a new ListRolePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewListRolePermissionsParams() *ListRolePermissionsParams {
	p := &ListRolePermissionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists role permissions
func (s *RoleService) ListRolePermissions(p *ListRolePermissionsParams) (*ListRolePermissionsResponse, error) {
	resp, err := s.cs.newRequest("listRolePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRolePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRolePermissionsResponse struct {
	Count           int               `json:"count"`
	RolePermissions []*RolePermission `json:"rolepermission"`
}

type RolePermission struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Permission  string `json:"permission"`
	Roleid      string `json:"roleid"`
	Rolename    string `json:"rolename"`
	Rule        string `json:"rule"`
}

type ListRolesParams struct {
	p map[string]interface{}
}

func (p *ListRolesParams) toURLValues() url.Values {
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
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *ListRolesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListRolesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListRolesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListRolesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListRolesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListRolesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListRolesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListRolesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListRolesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListRolesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListRolesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListRolesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListRolesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListRolesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListRolesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListRolesParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListRolesParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ListRolesParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new ListRolesParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewListRolesParams() *ListRolesParams {
	p := &ListRolesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RoleService) GetRoleID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListRolesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListRoles(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Roles[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Roles {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RoleService) GetRoleByName(name string, opts ...OptionFunc) (*Role, int, error) {
	id, count, err := s.GetRoleID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetRoleByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RoleService) GetRoleByID(id string, opts ...OptionFunc) (*Role, int, error) {
	p := &ListRolesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListRoles(p)
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
		return l.Roles[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Role UUID: %s!", id)
}

// Lists dynamic roles in CloudStack
func (s *RoleService) ListRoles(p *ListRolesParams) (*ListRolesResponse, error) {
	resp, err := s.cs.newRequest("listRoles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRolesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRolesResponse struct {
	Count int     `json:"count"`
	Roles []*Role `json:"role"`
}

type Role struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Isdefault   bool   `json:"isdefault"`
	Ispublic    bool   `json:"ispublic"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type UpdateRoleParams struct {
	p map[string]interface{}
}

func (p *UpdateRoleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *UpdateRoleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateRoleParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateRoleParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateRoleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateRoleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateRoleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateRoleParams) SetIspublic(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispublic"] = v
}

func (p *UpdateRoleParams) ResetIspublic() {
	if p.p != nil && p.p["ispublic"] != nil {
		delete(p.p, "ispublic")
	}
}

func (p *UpdateRoleParams) GetIspublic() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispublic"].(bool)
	return value, ok
}

func (p *UpdateRoleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateRoleParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateRoleParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateRoleParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *UpdateRoleParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *UpdateRoleParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateRoleParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewUpdateRoleParams(id string) *UpdateRoleParams {
	p := &UpdateRoleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a role
func (s *RoleService) UpdateRole(p *UpdateRoleParams) (*UpdateRoleResponse, error) {
	resp, err := s.cs.newRequest("updateRole", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateRoleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateRoleResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Isdefault   bool   `json:"isdefault"`
	Ispublic    bool   `json:"ispublic"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type UpdateRolePermissionParams struct {
	p map[string]interface{}
}

func (p *UpdateRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["permission"]; found {
		u.Set("permission", v.(string))
	}
	if v, found := p.p["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := p.p["ruleid"]; found {
		u.Set("ruleid", v.(string))
	}
	if v, found := p.p["ruleorder"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ruleorder", vv)
	}
	return u
}

func (p *UpdateRolePermissionParams) SetPermission(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["permission"] = v
}

func (p *UpdateRolePermissionParams) ResetPermission() {
	if p.p != nil && p.p["permission"] != nil {
		delete(p.p, "permission")
	}
}

func (p *UpdateRolePermissionParams) GetPermission() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["permission"].(string)
	return value, ok
}

func (p *UpdateRolePermissionParams) SetRoleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roleid"] = v
}

func (p *UpdateRolePermissionParams) ResetRoleid() {
	if p.p != nil && p.p["roleid"] != nil {
		delete(p.p, "roleid")
	}
}

func (p *UpdateRolePermissionParams) GetRoleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roleid"].(string)
	return value, ok
}

func (p *UpdateRolePermissionParams) SetRuleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ruleid"] = v
}

func (p *UpdateRolePermissionParams) ResetRuleid() {
	if p.p != nil && p.p["ruleid"] != nil {
		delete(p.p, "ruleid")
	}
}

func (p *UpdateRolePermissionParams) GetRuleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ruleid"].(string)
	return value, ok
}

func (p *UpdateRolePermissionParams) SetRuleorder(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ruleorder"] = v
}

func (p *UpdateRolePermissionParams) ResetRuleorder() {
	if p.p != nil && p.p["ruleorder"] != nil {
		delete(p.p, "ruleorder")
	}
}

func (p *UpdateRolePermissionParams) GetRuleorder() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ruleorder"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewUpdateRolePermissionParams(roleid string) *UpdateRolePermissionParams {
	p := &UpdateRolePermissionParams{}
	p.p = make(map[string]interface{})
	p.p["roleid"] = roleid
	return p
}

// Updates a role permission order
func (s *RoleService) UpdateRolePermission(p *UpdateRolePermissionParams) (*UpdateRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("updateRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateRolePermissionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateRolePermissionResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateRolePermissionResponse
	return json.Unmarshal(b, (*alias)(r))
}
