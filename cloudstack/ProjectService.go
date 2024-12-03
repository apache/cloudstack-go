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

type ProjectServiceIface interface {
	ActivateProject(p *ActivateProjectParams) (*ActivateProjectResponse, error)
	NewActivateProjectParams(id string) *ActivateProjectParams
	AddAccountToProject(p *AddAccountToProjectParams) (*AddAccountToProjectResponse, error)
	NewAddAccountToProjectParams(projectid string) *AddAccountToProjectParams
	AddUserToProject(p *AddUserToProjectParams) (*AddUserToProjectResponse, error)
	NewAddUserToProjectParams(projectid string, username string) *AddUserToProjectParams
	CreateProject(p *CreateProjectParams) (*CreateProjectResponse, error)
	NewCreateProjectParams(displaytext string, name string) *CreateProjectParams
	DeleteAccountFromProject(p *DeleteAccountFromProjectParams) (*DeleteAccountFromProjectResponse, error)
	NewDeleteAccountFromProjectParams(account string, projectid string) *DeleteAccountFromProjectParams
	DeleteUserFromProject(p *DeleteUserFromProjectParams) (*DeleteUserFromProjectResponse, error)
	NewDeleteUserFromProjectParams(projectid string, userid string) *DeleteUserFromProjectParams
	DeleteProject(p *DeleteProjectParams) (*DeleteProjectResponse, error)
	NewDeleteProjectParams(id string) *DeleteProjectParams
	DeleteProjectInvitation(p *DeleteProjectInvitationParams) (*DeleteProjectInvitationResponse, error)
	NewDeleteProjectInvitationParams(id string) *DeleteProjectInvitationParams
	ListProjectInvitations(p *ListProjectInvitationsParams) (*ListProjectInvitationsResponse, error)
	NewListProjectInvitationsParams() *ListProjectInvitationsParams
	GetProjectInvitationByID(id string, opts ...OptionFunc) (*ProjectInvitation, int, error)
	ListProjects(p *ListProjectsParams) (*ListProjectsResponse, error)
	NewListProjectsParams() *ListProjectsParams
	GetProjectID(name string, opts ...OptionFunc) (string, int, error)
	GetProjectByName(name string, opts ...OptionFunc) (*Project, int, error)
	GetProjectByID(id string, opts ...OptionFunc) (*Project, int, error)
	SuspendProject(p *SuspendProjectParams) (*SuspendProjectResponse, error)
	NewSuspendProjectParams(id string) *SuspendProjectParams
	UpdateProject(p *UpdateProjectParams) (*UpdateProjectResponse, error)
	NewUpdateProjectParams(id string) *UpdateProjectParams
	UpdateProjectInvitation(p *UpdateProjectInvitationParams) (*UpdateProjectInvitationResponse, error)
	NewUpdateProjectInvitationParams(projectid string) *UpdateProjectInvitationParams
	ListProjectRolePermissions(p *ListProjectRolePermissionsParams) (*ListProjectRolePermissionsResponse, error)
	NewListProjectRolePermissionsParams(projectid string) *ListProjectRolePermissionsParams
	CreateProjectRolePermission(p *CreateProjectRolePermissionParams) (*CreateProjectRolePermissionResponse, error)
	NewCreateProjectRolePermissionParams(permission string, projectid string, projectroleid string, rule string) *CreateProjectRolePermissionParams
	UpdateProjectRolePermission(p *UpdateProjectRolePermissionParams) (*UpdateProjectRolePermissionResponse, error)
	NewUpdateProjectRolePermissionParams(projectid string, projectroleid string) *UpdateProjectRolePermissionParams
	DeleteProjectRolePermission(p *DeleteProjectRolePermissionParams) (*DeleteProjectRolePermissionResponse, error)
	NewDeleteProjectRolePermissionParams(id string, projectid string) *DeleteProjectRolePermissionParams
}

type ActivateProjectParams struct {
	p map[string]interface{}
}

func (p *ActivateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ActivateProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ActivateProjectParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ActivateProjectParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ActivateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewActivateProjectParams(id string) *ActivateProjectParams {
	p := &ActivateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Activates a project
func (s *ProjectService) ActivateProject(p *ActivateProjectParams) (*ActivateProjectResponse, error) {
	resp, err := s.cs.newRequest("activateProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ActivateProjectResponse
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

type ActivateProjectResponse struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      interface{}         `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Taggedresources           []string            `json:"taggedresources"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type AddAccountToProjectParams struct {
	p map[string]interface{}
}

func (p *AddAccountToProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := p.p["roletype"]; found {
		u.Set("roletype", v.(string))
	}
	return u
}

func (p *AddAccountToProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *AddAccountToProjectParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *AddAccountToProjectParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *AddAccountToProjectParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
}

func (p *AddAccountToProjectParams) ResetEmail() {
	if p.p != nil && p.p["email"] != nil {
		delete(p.p, "email")
	}
}

func (p *AddAccountToProjectParams) GetEmail() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["email"].(string)
	return value, ok
}

func (p *AddAccountToProjectParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AddAccountToProjectParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *AddAccountToProjectParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *AddAccountToProjectParams) SetProjectroleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectroleid"] = v
}

func (p *AddAccountToProjectParams) ResetProjectroleid() {
	if p.p != nil && p.p["projectroleid"] != nil {
		delete(p.p, "projectroleid")
	}
}

func (p *AddAccountToProjectParams) GetProjectroleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectroleid"].(string)
	return value, ok
}

func (p *AddAccountToProjectParams) SetRoletype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roletype"] = v
}

func (p *AddAccountToProjectParams) ResetRoletype() {
	if p.p != nil && p.p["roletype"] != nil {
		delete(p.p, "roletype")
	}
}

func (p *AddAccountToProjectParams) GetRoletype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roletype"].(string)
	return value, ok
}

// You should always use this function to get a new AddAccountToProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewAddAccountToProjectParams(projectid string) *AddAccountToProjectParams {
	p := &AddAccountToProjectParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	return p
}

// Adds account to a project
func (s *ProjectService) AddAccountToProject(p *AddAccountToProjectParams) (*AddAccountToProjectResponse, error) {
	resp, err := s.cs.newRequest("addAccountToProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddAccountToProjectResponse
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

type AddAccountToProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AddUserToProjectParams struct {
	p map[string]interface{}
}

func (p *AddUserToProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := p.p["roletype"]; found {
		u.Set("roletype", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddUserToProjectParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
}

func (p *AddUserToProjectParams) ResetEmail() {
	if p.p != nil && p.p["email"] != nil {
		delete(p.p, "email")
	}
}

func (p *AddUserToProjectParams) GetEmail() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["email"].(string)
	return value, ok
}

func (p *AddUserToProjectParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *AddUserToProjectParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *AddUserToProjectParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *AddUserToProjectParams) SetProjectroleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectroleid"] = v
}

func (p *AddUserToProjectParams) ResetProjectroleid() {
	if p.p != nil && p.p["projectroleid"] != nil {
		delete(p.p, "projectroleid")
	}
}

func (p *AddUserToProjectParams) GetProjectroleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectroleid"].(string)
	return value, ok
}

func (p *AddUserToProjectParams) SetRoletype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roletype"] = v
}

func (p *AddUserToProjectParams) ResetRoletype() {
	if p.p != nil && p.p["roletype"] != nil {
		delete(p.p, "roletype")
	}
}

func (p *AddUserToProjectParams) GetRoletype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roletype"].(string)
	return value, ok
}

func (p *AddUserToProjectParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *AddUserToProjectParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *AddUserToProjectParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddUserToProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewAddUserToProjectParams(projectid string, username string) *AddUserToProjectParams {
	p := &AddUserToProjectParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	p.p["username"] = username
	return p
}

// Adds user to a project
func (s *ProjectService) AddUserToProject(p *AddUserToProjectParams) (*AddUserToProjectResponse, error) {
	resp, err := s.cs.newRequest("addUserToProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddUserToProjectResponse
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

type AddUserToProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type CreateProjectParams struct {
	p map[string]interface{}
}

func (p *CreateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (p *CreateProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateProjectParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateProjectParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateProjectParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *CreateProjectParams) ResetAccountid() {
	if p.p != nil && p.p["accountid"] != nil {
		delete(p.p, "accountid")
	}
}

func (p *CreateProjectParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *CreateProjectParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateProjectParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *CreateProjectParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateProjectParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateProjectParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateProjectParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateProjectParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateProjectParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateProjectParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateProjectParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *CreateProjectParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *CreateProjectParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewCreateProjectParams(displaytext string, name string) *CreateProjectParams {
	p := &CreateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	return p
}

// Creates a project
func (s *ProjectService) CreateProject(p *CreateProjectParams) (*CreateProjectResponse, error) {
	resp, err := s.cs.newRequest("createProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateProjectResponse
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

type CreateProjectResponse struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      interface{}         `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Taggedresources           []string            `json:"taggedresources"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type DeleteAccountFromProjectParams struct {
	p map[string]interface{}
}

func (p *DeleteAccountFromProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DeleteAccountFromProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *DeleteAccountFromProjectParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *DeleteAccountFromProjectParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *DeleteAccountFromProjectParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DeleteAccountFromProjectParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *DeleteAccountFromProjectParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAccountFromProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteAccountFromProjectParams(account string, projectid string) *DeleteAccountFromProjectParams {
	p := &DeleteAccountFromProjectParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["projectid"] = projectid
	return p
}

// Deletes account from the project
func (s *ProjectService) DeleteAccountFromProject(p *DeleteAccountFromProjectParams) (*DeleteAccountFromProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteAccountFromProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAccountFromProjectResponse
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

type DeleteAccountFromProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteUserFromProjectParams struct {
	p map[string]interface{}
}

func (p *DeleteUserFromProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (p *DeleteUserFromProjectParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DeleteUserFromProjectParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *DeleteUserFromProjectParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *DeleteUserFromProjectParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *DeleteUserFromProjectParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *DeleteUserFromProjectParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteUserFromProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteUserFromProjectParams(projectid string, userid string) *DeleteUserFromProjectParams {
	p := &DeleteUserFromProjectParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	p.p["userid"] = userid
	return p
}

// Deletes user from the project
func (s *ProjectService) DeleteUserFromProject(p *DeleteUserFromProjectParams) (*DeleteUserFromProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteUserFromProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUserFromProjectResponse
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

type DeleteUserFromProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteProjectParams struct {
	p map[string]interface{}
}

func (p *DeleteProjectParams) toURLValues() url.Values {
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

func (p *DeleteProjectParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *DeleteProjectParams) ResetCleanup() {
	if p.p != nil && p.p["cleanup"] != nil {
		delete(p.p, "cleanup")
	}
}

func (p *DeleteProjectParams) GetCleanup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanup"].(bool)
	return value, ok
}

func (p *DeleteProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteProjectParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteProjectParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectParams(id string) *DeleteProjectParams {
	p := &DeleteProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a project
func (s *ProjectService) DeleteProject(p *DeleteProjectParams) (*DeleteProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectResponse
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

type DeleteProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteProjectInvitationParams struct {
	p map[string]interface{}
}

func (p *DeleteProjectInvitationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteProjectInvitationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteProjectInvitationParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteProjectInvitationParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteProjectInvitationParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectInvitationParams(id string) *DeleteProjectInvitationParams {
	p := &DeleteProjectInvitationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes project invitation
func (s *ProjectService) DeleteProjectInvitation(p *DeleteProjectInvitationParams) (*DeleteProjectInvitationResponse, error) {
	resp, err := s.cs.newRequest("deleteProjectInvitation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectInvitationResponse
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

type DeleteProjectInvitationResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListProjectInvitationsParams struct {
	p map[string]interface{}
}

func (p *ListProjectInvitationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["activeonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("activeonly", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (p *ListProjectInvitationsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListProjectInvitationsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListProjectInvitationsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetActiveonly(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["activeonly"] = v
}

func (p *ListProjectInvitationsParams) ResetActiveonly() {
	if p.p != nil && p.p["activeonly"] != nil {
		delete(p.p, "activeonly")
	}
}

func (p *ListProjectInvitationsParams) GetActiveonly() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["activeonly"].(bool)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListProjectInvitationsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListProjectInvitationsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListProjectInvitationsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListProjectInvitationsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListProjectInvitationsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListProjectInvitationsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListProjectInvitationsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListProjectInvitationsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListProjectInvitationsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListProjectInvitationsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListProjectInvitationsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListProjectInvitationsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListProjectInvitationsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListProjectInvitationsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListProjectInvitationsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListProjectInvitationsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListProjectInvitationsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListProjectInvitationsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListProjectInvitationsParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *ListProjectInvitationsParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *ListProjectInvitationsParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

// You should always use this function to get a new ListProjectInvitationsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectInvitationsParams() *ListProjectInvitationsParams {
	p := &ListProjectInvitationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectInvitationByID(id string, opts ...OptionFunc) (*ProjectInvitation, int, error) {
	p := &ListProjectInvitationsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListProjectInvitations(p)
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
		return l.ProjectInvitations[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ProjectInvitation UUID: %s!", id)
}

// Lists project invitations and provides detailed information for listed invitations
func (s *ProjectService) ListProjectInvitations(p *ListProjectInvitationsParams) (*ListProjectInvitationsResponse, error) {
	resp, err := s.cs.newRequest("listProjectInvitations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectInvitationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectInvitationsResponse struct {
	Count              int                  `json:"count"`
	ProjectInvitations []*ProjectInvitation `json:"projectinvitation"`
}

type ProjectInvitation struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Domainpath string `json:"domainpath"`
	Email      string `json:"email"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	State      string `json:"state"`
	Userid     string `json:"userid"`
}

type ListProjectsParams struct {
	p map[string]interface{}
}

func (p *ListProjectsParams) toURLValues() url.Values {
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
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *ListProjectsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListProjectsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListProjectsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListProjectsParams) SetDetails(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *ListProjectsParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *ListProjectsParams) GetDetails() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].([]string)
	return value, ok
}

func (p *ListProjectsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *ListProjectsParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *ListProjectsParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *ListProjectsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListProjectsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListProjectsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListProjectsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListProjectsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListProjectsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListProjectsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListProjectsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListProjectsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListProjectsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListProjectsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListProjectsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListProjectsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListProjectsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListProjectsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListProjectsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListProjectsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListProjectsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListProjectsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListProjectsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListProjectsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListProjectsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListProjectsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListProjectsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListProjectsParams) SetShowicon(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showicon"] = v
}

func (p *ListProjectsParams) ResetShowicon() {
	if p.p != nil && p.p["showicon"] != nil {
		delete(p.p, "showicon")
	}
}

func (p *ListProjectsParams) GetShowicon() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showicon"].(bool)
	return value, ok
}

func (p *ListProjectsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListProjectsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListProjectsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListProjectsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListProjectsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListProjectsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListProjectsParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *ListProjectsParams) ResetUsername() {
	if p.p != nil && p.p["username"] != nil {
		delete(p.p, "username")
	}
}

func (p *ListProjectsParams) GetUsername() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["username"].(string)
	return value, ok
}

// You should always use this function to get a new ListProjectsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectsParams() *ListProjectsParams {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListProjects(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Projects[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Projects {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectByName(name string, opts ...OptionFunc) (*Project, int, error) {
	id, count, err := s.GetProjectID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetProjectByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectByID(id string, opts ...OptionFunc) (*Project, int, error) {
	p := &ListProjectsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListProjects(p)
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
		return l.Projects[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Project UUID: %s!", id)
}

// Lists projects and provides detailed information for listed projects
func (s *ProjectService) ListProjects(p *ListProjectsParams) (*ListProjectsResponse, error) {
	resp, err := s.cs.newRequest("listProjects", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectsResponse struct {
	Count    int        `json:"count"`
	Projects []*Project `json:"project"`
}

type Project struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      interface{}         `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Taggedresources           []string            `json:"taggedresources"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type SuspendProjectParams struct {
	p map[string]interface{}
}

func (p *SuspendProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *SuspendProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *SuspendProjectParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *SuspendProjectParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new SuspendProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewSuspendProjectParams(id string) *SuspendProjectParams {
	p := &SuspendProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Suspends a project
func (s *ProjectService) SuspendProject(p *SuspendProjectParams) (*SuspendProjectResponse, error) {
	resp, err := s.cs.newRequest("suspendProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r SuspendProjectResponse
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

type SuspendProjectResponse struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      interface{}         `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Taggedresources           []string            `json:"taggedresources"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type UpdateProjectParams struct {
	p map[string]interface{}
}

func (p *UpdateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["roletype"]; found {
		u.Set("roletype", v.(string))
	}
	if v, found := p.p["swapowner"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("swapowner", vv)
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (p *UpdateProjectParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpdateProjectParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *UpdateProjectParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UpdateProjectParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateProjectParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *UpdateProjectParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateProjectParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateProjectParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateProjectParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateProjectParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateProjectParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateProjectParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateProjectParams) SetRoletype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["roletype"] = v
}

func (p *UpdateProjectParams) ResetRoletype() {
	if p.p != nil && p.p["roletype"] != nil {
		delete(p.p, "roletype")
	}
}

func (p *UpdateProjectParams) GetRoletype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["roletype"].(string)
	return value, ok
}

func (p *UpdateProjectParams) SetSwapowner(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["swapowner"] = v
}

func (p *UpdateProjectParams) ResetSwapowner() {
	if p.p != nil && p.p["swapowner"] != nil {
		delete(p.p, "swapowner")
	}
}

func (p *UpdateProjectParams) GetSwapowner() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["swapowner"].(bool)
	return value, ok
}

func (p *UpdateProjectParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *UpdateProjectParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *UpdateProjectParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectParams(id string) *UpdateProjectParams {
	p := &UpdateProjectParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a project
func (s *ProjectService) UpdateProject(p *UpdateProjectParams) (*UpdateProjectResponse, error) {
	resp, err := s.cs.newRequest("updateProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectResponse
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

type UpdateProjectResponse struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      interface{}         `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Taggedresources           []string            `json:"taggedresources"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type UpdateProjectInvitationParams struct {
	p map[string]interface{}
}

func (p *UpdateProjectInvitationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accept"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("accept", vv)
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["token"]; found {
		u.Set("token", v.(string))
	}
	if v, found := p.p["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (p *UpdateProjectInvitationParams) SetAccept(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accept"] = v
}

func (p *UpdateProjectInvitationParams) ResetAccept() {
	if p.p != nil && p.p["accept"] != nil {
		delete(p.p, "accept")
	}
}

func (p *UpdateProjectInvitationParams) GetAccept() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accept"].(bool)
	return value, ok
}

func (p *UpdateProjectInvitationParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UpdateProjectInvitationParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *UpdateProjectInvitationParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UpdateProjectInvitationParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *UpdateProjectInvitationParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *UpdateProjectInvitationParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *UpdateProjectInvitationParams) SetToken(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["token"] = v
}

func (p *UpdateProjectInvitationParams) ResetToken() {
	if p.p != nil && p.p["token"] != nil {
		delete(p.p, "token")
	}
}

func (p *UpdateProjectInvitationParams) GetToken() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["token"].(string)
	return value, ok
}

func (p *UpdateProjectInvitationParams) SetUserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userid"] = v
}

func (p *UpdateProjectInvitationParams) ResetUserid() {
	if p.p != nil && p.p["userid"] != nil {
		delete(p.p, "userid")
	}
}

func (p *UpdateProjectInvitationParams) GetUserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateProjectInvitationParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectInvitationParams(projectid string) *UpdateProjectInvitationParams {
	p := &UpdateProjectInvitationParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	return p
}

// Accepts or declines project invitation
func (s *ProjectService) UpdateProjectInvitation(p *UpdateProjectInvitationParams) (*UpdateProjectInvitationResponse, error) {
	resp, err := s.cs.newRequest("updateProjectInvitation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectInvitationResponse
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

type UpdateProjectInvitationResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListProjectRolePermissionsParams struct {
	p map[string]interface{}
}

func (p *ListProjectRolePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	return u
}

func (p *ListProjectRolePermissionsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListProjectRolePermissionsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListProjectRolePermissionsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListProjectRolePermissionsParams) SetProjectroleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectroleid"] = v
}

func (p *ListProjectRolePermissionsParams) ResetProjectroleid() {
	if p.p != nil && p.p["projectroleid"] != nil {
		delete(p.p, "projectroleid")
	}
}

func (p *ListProjectRolePermissionsParams) GetProjectroleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectroleid"].(string)
	return value, ok
}

// You should always use this function to get a new ListProjectRolePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectRolePermissionsParams(projectid string) *ListProjectRolePermissionsParams {
	p := &ListProjectRolePermissionsParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	return p
}

// Lists a project's project role permissions
func (s *ProjectService) ListProjectRolePermissions(p *ListProjectRolePermissionsParams) (*ListProjectRolePermissionsResponse, error) {
	resp, err := s.cs.newRequest("listProjectRolePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectRolePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectRolePermissionsResponse struct {
	Count                  int                      `json:"count"`
	ProjectRolePermissions []*ProjectRolePermission `json:"projectrolepermission"`
}

type ProjectRolePermission struct {
	Description     string `json:"description"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Permission      string `json:"permission"`
	Projectid       string `json:"projectid"`
	Projectroleid   string `json:"projectroleid"`
	Projectrolename string `json:"projectrolename"`
	Rule            string `json:"rule"`
}

type CreateProjectRolePermissionParams struct {
	p map[string]interface{}
}

func (p *CreateProjectRolePermissionParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := p.p["rule"]; found {
		u.Set("rule", v.(string))
	}
	return u
}

func (p *CreateProjectRolePermissionParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateProjectRolePermissionParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateProjectRolePermissionParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateProjectRolePermissionParams) SetPermission(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["permission"] = v
}

func (p *CreateProjectRolePermissionParams) ResetPermission() {
	if p.p != nil && p.p["permission"] != nil {
		delete(p.p, "permission")
	}
}

func (p *CreateProjectRolePermissionParams) GetPermission() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["permission"].(string)
	return value, ok
}

func (p *CreateProjectRolePermissionParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateProjectRolePermissionParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateProjectRolePermissionParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateProjectRolePermissionParams) SetProjectroleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectroleid"] = v
}

func (p *CreateProjectRolePermissionParams) ResetProjectroleid() {
	if p.p != nil && p.p["projectroleid"] != nil {
		delete(p.p, "projectroleid")
	}
}

func (p *CreateProjectRolePermissionParams) GetProjectroleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectroleid"].(string)
	return value, ok
}

func (p *CreateProjectRolePermissionParams) SetRule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rule"] = v
}

func (p *CreateProjectRolePermissionParams) ResetRule() {
	if p.p != nil && p.p["rule"] != nil {
		delete(p.p, "rule")
	}
}

func (p *CreateProjectRolePermissionParams) GetRule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rule"].(string)
	return value, ok
}

// You should always use this function to get a new CreateProjectRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewCreateProjectRolePermissionParams(permission string, projectid string, projectroleid string, rule string) *CreateProjectRolePermissionParams {
	p := &CreateProjectRolePermissionParams{}
	p.p = make(map[string]interface{})
	p.p["permission"] = permission
	p.p["projectid"] = projectid
	p.p["projectroleid"] = projectroleid
	p.p["rule"] = rule
	return p
}

// Adds API permissions to a project role
func (s *ProjectService) CreateProjectRolePermission(p *CreateProjectRolePermissionParams) (*CreateProjectRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("createProjectRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateProjectRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateProjectRolePermissionResponse struct {
	Description     string `json:"description"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Permission      string `json:"permission"`
	Projectid       string `json:"projectid"`
	Projectroleid   string `json:"projectroleid"`
	Projectrolename string `json:"projectrolename"`
	Rule            string `json:"rule"`
}

type UpdateProjectRolePermissionParams struct {
	p map[string]interface{}
}

func (p *UpdateProjectRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["permission"]; found {
		u.Set("permission", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := p.p["projectrolepermissionid"]; found {
		u.Set("projectrolepermissionid", v.(string))
	}
	if v, found := p.p["ruleorder"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ruleorder", vv)
	}
	return u
}

func (p *UpdateProjectRolePermissionParams) SetPermission(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["permission"] = v
}

func (p *UpdateProjectRolePermissionParams) ResetPermission() {
	if p.p != nil && p.p["permission"] != nil {
		delete(p.p, "permission")
	}
}

func (p *UpdateProjectRolePermissionParams) GetPermission() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["permission"].(string)
	return value, ok
}

func (p *UpdateProjectRolePermissionParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *UpdateProjectRolePermissionParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *UpdateProjectRolePermissionParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *UpdateProjectRolePermissionParams) SetProjectroleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectroleid"] = v
}

func (p *UpdateProjectRolePermissionParams) ResetProjectroleid() {
	if p.p != nil && p.p["projectroleid"] != nil {
		delete(p.p, "projectroleid")
	}
}

func (p *UpdateProjectRolePermissionParams) GetProjectroleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectroleid"].(string)
	return value, ok
}

func (p *UpdateProjectRolePermissionParams) SetProjectrolepermissionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectrolepermissionid"] = v
}

func (p *UpdateProjectRolePermissionParams) ResetProjectrolepermissionid() {
	if p.p != nil && p.p["projectrolepermissionid"] != nil {
		delete(p.p, "projectrolepermissionid")
	}
}

func (p *UpdateProjectRolePermissionParams) GetProjectrolepermissionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectrolepermissionid"].(string)
	return value, ok
}

func (p *UpdateProjectRolePermissionParams) SetRuleorder(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ruleorder"] = v
}

func (p *UpdateProjectRolePermissionParams) ResetRuleorder() {
	if p.p != nil && p.p["ruleorder"] != nil {
		delete(p.p, "ruleorder")
	}
}

func (p *UpdateProjectRolePermissionParams) GetRuleorder() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ruleorder"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateProjectRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectRolePermissionParams(projectid string, projectroleid string) *UpdateProjectRolePermissionParams {
	p := &UpdateProjectRolePermissionParams{}
	p.p = make(map[string]interface{})
	p.p["projectid"] = projectid
	p.p["projectroleid"] = projectroleid
	return p
}

// Updates a project role permission and/or order
func (s *ProjectService) UpdateProjectRolePermission(p *UpdateProjectRolePermissionParams) (*UpdateProjectRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("updateProjectRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateProjectRolePermissionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateProjectRolePermissionResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateProjectRolePermissionResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteProjectRolePermissionParams struct {
	p map[string]interface{}
}

func (p *DeleteProjectRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *DeleteProjectRolePermissionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteProjectRolePermissionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteProjectRolePermissionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteProjectRolePermissionParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *DeleteProjectRolePermissionParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *DeleteProjectRolePermissionParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteProjectRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectRolePermissionParams(id string, projectid string) *DeleteProjectRolePermissionParams {
	p := &DeleteProjectRolePermissionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["projectid"] = projectid
	return p
}

// Deletes a project role permission in the project
func (s *ProjectService) DeleteProjectRolePermission(p *DeleteProjectRolePermissionParams) (*DeleteProjectRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("deleteProjectRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteProjectRolePermissionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteProjectRolePermissionResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteProjectRolePermissionResponse
	return json.Unmarshal(b, (*alias)(r))
}
