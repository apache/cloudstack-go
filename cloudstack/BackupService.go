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

type BackupServiceIface interface {
	AddBackupRepository(p *AddBackupRepositoryParams) (*AddBackupRepositoryResponse, error)
	NewAddBackupRepositoryParams(address string, name string, backupType string, zoneid string) *AddBackupRepositoryParams
	CreateBackup(p *CreateBackupParams) (*CreateBackupResponse, error)
	NewCreateBackupParams(virtualmachineid string) *CreateBackupParams
	CreateBackupSchedule(p *CreateBackupScheduleParams) (*CreateBackupScheduleResponse, error)
	NewCreateBackupScheduleParams(intervaltype string, schedule string, timezone string, virtualmachineid string) *CreateBackupScheduleParams
	CreateVMFromBackup(p *CreateVMFromBackupParams) (*CreateVMFromBackupResponse, error)
	NewCreateVMFromBackupParams(backupid string, zoneid string) *CreateVMFromBackupParams
	DeleteBackup(p *DeleteBackupParams) (*DeleteBackupResponse, error)
	NewDeleteBackupParams(id string) *DeleteBackupParams
	DeleteBackupOffering(p *DeleteBackupOfferingParams) (*DeleteBackupOfferingResponse, error)
	NewDeleteBackupOfferingParams(id string) *DeleteBackupOfferingParams
	DeleteBackupRepository(p *DeleteBackupRepositoryParams) (*DeleteBackupRepositoryResponse, error)
	NewDeleteBackupRepositoryParams(id string) *DeleteBackupRepositoryParams
	DeleteBackupSchedule(p *DeleteBackupScheduleParams) (*DeleteBackupScheduleResponse, error)
	NewDeleteBackupScheduleParams() *DeleteBackupScheduleParams
	ImportBackupOffering(p *ImportBackupOfferingParams) (*ImportBackupOfferingResponse, error)
	NewImportBackupOfferingParams(allowuserdrivenbackups bool, description string, externalid string, name string, zoneid string) *ImportBackupOfferingParams
	ListBackupOfferings(p *ListBackupOfferingsParams) (*ListBackupOfferingsResponse, error)
	NewListBackupOfferingsParams() *ListBackupOfferingsParams
	GetBackupOfferingID(keyword string, opts ...OptionFunc) (string, int, error)
	GetBackupOfferingByName(name string, opts ...OptionFunc) (*BackupOffering, int, error)
	GetBackupOfferingByID(id string, opts ...OptionFunc) (*BackupOffering, int, error)
	ListBackupProviderOfferings(p *ListBackupProviderOfferingsParams) (*ListBackupProviderOfferingsResponse, error)
	NewListBackupProviderOfferingsParams(zoneid string) *ListBackupProviderOfferingsParams
	GetBackupProviderOfferingID(keyword string, zoneid string, opts ...OptionFunc) (string, int, error)
	ListBackupProviders(p *ListBackupProvidersParams) (*ListBackupProvidersResponse, error)
	NewListBackupProvidersParams() *ListBackupProvidersParams
	ListBackupRepositories(p *ListBackupRepositoriesParams) (*ListBackupRepositoriesResponse, error)
	NewListBackupRepositoriesParams() *ListBackupRepositoriesParams
	GetBackupRepositoryID(name string, opts ...OptionFunc) (string, int, error)
	GetBackupRepositoryByName(name string, opts ...OptionFunc) (*BackupRepository, int, error)
	GetBackupRepositoryByID(id string, opts ...OptionFunc) (*BackupRepository, int, error)
	ListBackupSchedule(p *ListBackupScheduleParams) (*ListBackupScheduleResponse, error)
	NewListBackupScheduleParams() *ListBackupScheduleParams
	GetBackupScheduleByID(id string, opts ...OptionFunc) (*BackupSchedule, int, error)
	ListBackups(p *ListBackupsParams) (*ListBackupsResponse, error)
	NewListBackupsParams() *ListBackupsParams
	GetBackupID(name string, opts ...OptionFunc) (string, int, error)
	GetBackupByName(name string, opts ...OptionFunc) (*Backup, int, error)
	GetBackupByID(id string, opts ...OptionFunc) (*Backup, int, error)
	RestoreBackup(p *RestoreBackupParams) (*RestoreBackupResponse, error)
	NewRestoreBackupParams(id string) *RestoreBackupParams
	UpdateBackupRepository(p *UpdateBackupRepositoryParams) (*UpdateBackupRepositoryResponse, error)
	NewUpdateBackupRepositoryParams(id string) *UpdateBackupRepositoryParams
	UpdateBackupOffering(p *UpdateBackupOfferingParams) (*UpdateBackupOfferingResponse, error)
	NewUpdateBackupOfferingParams(id string) *UpdateBackupOfferingParams
	UpdateBackupSchedule(p *UpdateBackupScheduleParams) (*UpdateBackupScheduleResponse, error)
	NewUpdateBackupScheduleParams(intervaltype string, schedule string, timezone string, virtualmachineid string) *UpdateBackupScheduleParams
}

type AddBackupRepositoryParams struct {
	p map[string]interface{}
}

func (p *AddBackupRepositoryParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["address"]; found {
		u.Set("address", v.(string))
	}
	if v, found := p.p["capacitybytes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacitybytes", vv)
	}
	if v, found := p.p["crosszoneinstancecreation"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("crosszoneinstancecreation", vv)
	}
	if v, found := p.p["mountopts"]; found {
		u.Set("mountopts", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddBackupRepositoryParams) SetAddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["address"] = v
}

func (p *AddBackupRepositoryParams) ResetAddress() {
	if p.p != nil && p.p["address"] != nil {
		delete(p.p, "address")
	}
}

func (p *AddBackupRepositoryParams) GetAddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["address"].(string)
	return value, ok
}

func (p *AddBackupRepositoryParams) SetCapacitybytes(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacitybytes"] = v
}

func (p *AddBackupRepositoryParams) ResetCapacitybytes() {
	if p.p != nil && p.p["capacitybytes"] != nil {
		delete(p.p, "capacitybytes")
	}
}

func (p *AddBackupRepositoryParams) GetCapacitybytes() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["capacitybytes"].(int64)
	return value, ok
}

func (p *AddBackupRepositoryParams) SetCrosszoneinstancecreation(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["crosszoneinstancecreation"] = v
}

func (p *AddBackupRepositoryParams) ResetCrosszoneinstancecreation() {
	if p.p != nil && p.p["crosszoneinstancecreation"] != nil {
		delete(p.p, "crosszoneinstancecreation")
	}
}

func (p *AddBackupRepositoryParams) GetCrosszoneinstancecreation() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["crosszoneinstancecreation"].(bool)
	return value, ok
}

func (p *AddBackupRepositoryParams) SetMountopts(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mountopts"] = v
}

func (p *AddBackupRepositoryParams) ResetMountopts() {
	if p.p != nil && p.p["mountopts"] != nil {
		delete(p.p, "mountopts")
	}
}

func (p *AddBackupRepositoryParams) GetMountopts() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mountopts"].(string)
	return value, ok
}

func (p *AddBackupRepositoryParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddBackupRepositoryParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddBackupRepositoryParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddBackupRepositoryParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *AddBackupRepositoryParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *AddBackupRepositoryParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *AddBackupRepositoryParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *AddBackupRepositoryParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *AddBackupRepositoryParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *AddBackupRepositoryParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddBackupRepositoryParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddBackupRepositoryParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddBackupRepositoryParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewAddBackupRepositoryParams(address string, name string, backupType string, zoneid string) *AddBackupRepositoryParams {
	p := &AddBackupRepositoryParams{}
	p.p = make(map[string]interface{})
	p.p["address"] = address
	p.p["name"] = name
	p.p["type"] = backupType
	p.p["zoneid"] = zoneid
	return p
}

// Adds a backup repository to store NAS backups
func (s *BackupService) AddBackupRepository(p *AddBackupRepositoryParams) (*AddBackupRepositoryResponse, error) {
	resp, err := s.cs.newPostRequest("addBackupRepository", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBackupRepositoryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddBackupRepositoryResponse struct {
	Address                   string `json:"address"`
	Capacitybytes             int64  `json:"capacitybytes"`
	Created                   string `json:"created"`
	Crosszoneinstancecreation bool   `json:"crosszoneinstancecreation"`
	Id                        string `json:"id"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Name                      string `json:"name"`
	Provider                  string `json:"provider"`
	Type                      string `json:"type"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type CreateBackupParams struct {
	p map[string]interface{}
}

func (p *CreateBackupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["quiescevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("quiescevm", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *CreateBackupParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateBackupParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateBackupParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateBackupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateBackupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateBackupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateBackupParams) SetQuiescevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiescevm"] = v
}

func (p *CreateBackupParams) ResetQuiescevm() {
	if p.p != nil && p.p["quiescevm"] != nil {
		delete(p.p, "quiescevm")
	}
}

func (p *CreateBackupParams) GetQuiescevm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quiescevm"].(bool)
	return value, ok
}

func (p *CreateBackupParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *CreateBackupParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *CreateBackupParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateBackupParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewCreateBackupParams(virtualmachineid string) *CreateBackupParams {
	p := &CreateBackupParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Create VM backup
func (s *BackupService) CreateBackup(p *CreateBackupParams) (*CreateBackupResponse, error) {
	resp, err := s.cs.newPostRequest("createBackup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateBackupResponse
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

type CreateBackupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type CreateBackupScheduleParams struct {
	p map[string]interface{}
}

func (p *CreateBackupScheduleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["intervaltype"]; found {
		u.Set("intervaltype", v.(string))
	}
	if v, found := p.p["maxbackups"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxbackups", vv)
	}
	if v, found := p.p["quiescevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("quiescevm", vv)
	}
	if v, found := p.p["schedule"]; found {
		u.Set("schedule", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *CreateBackupScheduleParams) SetIntervaltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltype"] = v
}

func (p *CreateBackupScheduleParams) ResetIntervaltype() {
	if p.p != nil && p.p["intervaltype"] != nil {
		delete(p.p, "intervaltype")
	}
}

func (p *CreateBackupScheduleParams) GetIntervaltype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["intervaltype"].(string)
	return value, ok
}

func (p *CreateBackupScheduleParams) SetMaxbackups(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxbackups"] = v
}

func (p *CreateBackupScheduleParams) ResetMaxbackups() {
	if p.p != nil && p.p["maxbackups"] != nil {
		delete(p.p, "maxbackups")
	}
}

func (p *CreateBackupScheduleParams) GetMaxbackups() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxbackups"].(int)
	return value, ok
}

func (p *CreateBackupScheduleParams) SetQuiescevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiescevm"] = v
}

func (p *CreateBackupScheduleParams) ResetQuiescevm() {
	if p.p != nil && p.p["quiescevm"] != nil {
		delete(p.p, "quiescevm")
	}
}

func (p *CreateBackupScheduleParams) GetQuiescevm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quiescevm"].(bool)
	return value, ok
}

func (p *CreateBackupScheduleParams) SetSchedule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["schedule"] = v
}

func (p *CreateBackupScheduleParams) ResetSchedule() {
	if p.p != nil && p.p["schedule"] != nil {
		delete(p.p, "schedule")
	}
}

func (p *CreateBackupScheduleParams) GetSchedule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["schedule"].(string)
	return value, ok
}

func (p *CreateBackupScheduleParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *CreateBackupScheduleParams) ResetTimezone() {
	if p.p != nil && p.p["timezone"] != nil {
		delete(p.p, "timezone")
	}
}

func (p *CreateBackupScheduleParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

func (p *CreateBackupScheduleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *CreateBackupScheduleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *CreateBackupScheduleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateBackupScheduleParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewCreateBackupScheduleParams(intervaltype string, schedule string, timezone string, virtualmachineid string) *CreateBackupScheduleParams {
	p := &CreateBackupScheduleParams{}
	p.p = make(map[string]interface{})
	p.p["intervaltype"] = intervaltype
	p.p["schedule"] = schedule
	p.p["timezone"] = timezone
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Creates a user-defined VM backup schedule
func (s *BackupService) CreateBackupSchedule(p *CreateBackupScheduleParams) (*CreateBackupScheduleResponse, error) {
	resp, err := s.cs.newPostRequest("createBackupSchedule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateBackupScheduleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateBackupScheduleResponse struct {
	Account                 string            `json:"account"`
	Accountid               string            `json:"accountid"`
	Backupofferingid        string            `json:"backupofferingid"`
	Backupofferingname      string            `json:"backupofferingname"`
	Created                 string            `json:"created"`
	Description             string            `json:"description"`
	Domain                  string            `json:"domain"`
	Domainid                string            `json:"domainid"`
	Externalid              string            `json:"externalid"`
	Id                      string            `json:"id"`
	Intervaltype            string            `json:"intervaltype"`
	Isbackupvmexpunged      bool              `json:"isbackupvmexpunged"`
	JobID                   string            `json:"jobid"`
	Jobstatus               int               `json:"jobstatus"`
	Name                    string            `json:"name"`
	Size                    int64             `json:"size"`
	Status                  string            `json:"status"`
	Type                    string            `json:"type"`
	Virtualmachineid        string            `json:"virtualmachineid"`
	Virtualmachinename      string            `json:"virtualmachinename"`
	Virtualsize             int64             `json:"virtualsize"`
	Vmbackupofferingremoved bool              `json:"vmbackupofferingremoved"`
	Vmdetails               map[string]string `json:"vmdetails"`
	Volumes                 string            `json:"volumes"`
	Zone                    string            `json:"zone"`
	Zoneid                  string            `json:"zoneid"`
}

type CreateVMFromBackupParams struct {
	p map[string]interface{}
}

func (p *CreateVMFromBackupParams) toURLValues() url.Values {
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
	if v, found := p.p["backupid"]; found {
		u.Set("backupid", v.(string))
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
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
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
	if v, found := p.p["datadisksdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("datadisksdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("datadisksdetails[%d].value", i), m[k])
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
	if v, found := p.p["externaldetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("externaldetails[%d].key", i), k)
			u.Set(fmt.Sprintf("externaldetails[%d].value", i), m[k])
		}
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
	if v, found := p.p["leaseduration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("leaseduration", vv)
	}
	if v, found := p.p["leaseexpiryaction"]; found {
		u.Set("leaseexpiryaction", v.(string))
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["preserveip"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("preserveip", vv)
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateVMFromBackupParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateVMFromBackupParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateVMFromBackupParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetAffinitygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupids"] = v
}

func (p *CreateVMFromBackupParams) ResetAffinitygroupids() {
	if p.p != nil && p.p["affinitygroupids"] != nil {
		delete(p.p, "affinitygroupids")
	}
}

func (p *CreateVMFromBackupParams) GetAffinitygroupids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupids"].([]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetAffinitygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["affinitygroupnames"] = v
}

func (p *CreateVMFromBackupParams) ResetAffinitygroupnames() {
	if p.p != nil && p.p["affinitygroupnames"] != nil {
		delete(p.p, "affinitygroupnames")
	}
}

func (p *CreateVMFromBackupParams) GetAffinitygroupnames() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["affinitygroupnames"].([]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetBackupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["backupid"] = v
}

func (p *CreateVMFromBackupParams) ResetBackupid() {
	if p.p != nil && p.p["backupid"] != nil {
		delete(p.p, "backupid")
	}
}

func (p *CreateVMFromBackupParams) GetBackupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["backupid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetBootintosetup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootintosetup"] = v
}

func (p *CreateVMFromBackupParams) ResetBootintosetup() {
	if p.p != nil && p.p["bootintosetup"] != nil {
		delete(p.p, "bootintosetup")
	}
}

func (p *CreateVMFromBackupParams) GetBootintosetup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootintosetup"].(bool)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetBootmode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bootmode"] = v
}

func (p *CreateVMFromBackupParams) ResetBootmode() {
	if p.p != nil && p.p["bootmode"] != nil {
		delete(p.p, "bootmode")
	}
}

func (p *CreateVMFromBackupParams) GetBootmode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bootmode"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetBoottype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["boottype"] = v
}

func (p *CreateVMFromBackupParams) ResetBoottype() {
	if p.p != nil && p.p["boottype"] != nil {
		delete(p.p, "boottype")
	}
}

func (p *CreateVMFromBackupParams) GetBoottype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["boottype"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *CreateVMFromBackupParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *CreateVMFromBackupParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetCopyimagetags(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["copyimagetags"] = v
}

func (p *CreateVMFromBackupParams) ResetCopyimagetags() {
	if p.p != nil && p.p["copyimagetags"] != nil {
		delete(p.p, "copyimagetags")
	}
}

func (p *CreateVMFromBackupParams) GetCopyimagetags() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["copyimagetags"].(bool)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *CreateVMFromBackupParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *CreateVMFromBackupParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDatadiskofferinglist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["datadiskofferinglist"] = v
}

func (p *CreateVMFromBackupParams) ResetDatadiskofferinglist() {
	if p.p != nil && p.p["datadiskofferinglist"] != nil {
		delete(p.p, "datadiskofferinglist")
	}
}

func (p *CreateVMFromBackupParams) GetDatadiskofferinglist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["datadiskofferinglist"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDatadisksdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["datadisksdetails"] = v
}

func (p *CreateVMFromBackupParams) ResetDatadisksdetails() {
	if p.p != nil && p.p["datadisksdetails"] != nil {
		delete(p.p, "datadisksdetails")
	}
}

func (p *CreateVMFromBackupParams) GetDatadisksdetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["datadisksdetails"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDeploymentplanner(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deploymentplanner"] = v
}

func (p *CreateVMFromBackupParams) ResetDeploymentplanner() {
	if p.p != nil && p.p["deploymentplanner"] != nil {
		delete(p.p, "deploymentplanner")
	}
}

func (p *CreateVMFromBackupParams) GetDeploymentplanner() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deploymentplanner"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateVMFromBackupParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *CreateVMFromBackupParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDhcpoptionsnetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dhcpoptionsnetworklist"] = v
}

func (p *CreateVMFromBackupParams) ResetDhcpoptionsnetworklist() {
	if p.p != nil && p.p["dhcpoptionsnetworklist"] != nil {
		delete(p.p, "dhcpoptionsnetworklist")
	}
}

func (p *CreateVMFromBackupParams) GetDhcpoptionsnetworklist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dhcpoptionsnetworklist"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *CreateVMFromBackupParams) ResetDiskofferingid() {
	if p.p != nil && p.p["diskofferingid"] != nil {
		delete(p.p, "diskofferingid")
	}
}

func (p *CreateVMFromBackupParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDisplayname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayname"] = v
}

func (p *CreateVMFromBackupParams) ResetDisplayname() {
	if p.p != nil && p.p["displayname"] != nil {
		delete(p.p, "displayname")
	}
}

func (p *CreateVMFromBackupParams) GetDisplayname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayname"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDisplayvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvm"] = v
}

func (p *CreateVMFromBackupParams) ResetDisplayvm() {
	if p.p != nil && p.p["displayvm"] != nil {
		delete(p.p, "displayvm")
	}
}

func (p *CreateVMFromBackupParams) GetDisplayvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvm"].(bool)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateVMFromBackupParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateVMFromBackupParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetDynamicscalingenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dynamicscalingenabled"] = v
}

func (p *CreateVMFromBackupParams) ResetDynamicscalingenabled() {
	if p.p != nil && p.p["dynamicscalingenabled"] != nil {
		delete(p.p, "dynamicscalingenabled")
	}
}

func (p *CreateVMFromBackupParams) GetDynamicscalingenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dynamicscalingenabled"].(bool)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetExternaldetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["externaldetails"] = v
}

func (p *CreateVMFromBackupParams) ResetExternaldetails() {
	if p.p != nil && p.p["externaldetails"] != nil {
		delete(p.p, "externaldetails")
	}
}

func (p *CreateVMFromBackupParams) GetExternaldetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["externaldetails"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetExtraconfig(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["extraconfig"] = v
}

func (p *CreateVMFromBackupParams) ResetExtraconfig() {
	if p.p != nil && p.p["extraconfig"] != nil {
		delete(p.p, "extraconfig")
	}
}

func (p *CreateVMFromBackupParams) GetExtraconfig() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["extraconfig"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetGroup(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["group"] = v
}

func (p *CreateVMFromBackupParams) ResetGroup() {
	if p.p != nil && p.p["group"] != nil {
		delete(p.p, "group")
	}
}

func (p *CreateVMFromBackupParams) GetGroup() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["group"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *CreateVMFromBackupParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *CreateVMFromBackupParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *CreateVMFromBackupParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *CreateVMFromBackupParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetIodriverpolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iodriverpolicy"] = v
}

func (p *CreateVMFromBackupParams) ResetIodriverpolicy() {
	if p.p != nil && p.p["iodriverpolicy"] != nil {
		delete(p.p, "iodriverpolicy")
	}
}

func (p *CreateVMFromBackupParams) GetIodriverpolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iodriverpolicy"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetIothreadsenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iothreadsenabled"] = v
}

func (p *CreateVMFromBackupParams) ResetIothreadsenabled() {
	if p.p != nil && p.p["iothreadsenabled"] != nil {
		delete(p.p, "iothreadsenabled")
	}
}

func (p *CreateVMFromBackupParams) GetIothreadsenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iothreadsenabled"].(bool)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetIp6address(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ip6address"] = v
}

func (p *CreateVMFromBackupParams) ResetIp6address() {
	if p.p != nil && p.p["ip6address"] != nil {
		delete(p.p, "ip6address")
	}
}

func (p *CreateVMFromBackupParams) GetIp6address() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ip6address"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *CreateVMFromBackupParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *CreateVMFromBackupParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetIptonetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iptonetworklist"] = v
}

func (p *CreateVMFromBackupParams) ResetIptonetworklist() {
	if p.p != nil && p.p["iptonetworklist"] != nil {
		delete(p.p, "iptonetworklist")
	}
}

func (p *CreateVMFromBackupParams) GetIptonetworklist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iptonetworklist"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetKeyboard(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyboard"] = v
}

func (p *CreateVMFromBackupParams) ResetKeyboard() {
	if p.p != nil && p.p["keyboard"] != nil {
		delete(p.p, "keyboard")
	}
}

func (p *CreateVMFromBackupParams) GetKeyboard() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyboard"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetKeypair(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypair"] = v
}

func (p *CreateVMFromBackupParams) ResetKeypair() {
	if p.p != nil && p.p["keypair"] != nil {
		delete(p.p, "keypair")
	}
}

func (p *CreateVMFromBackupParams) GetKeypair() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypair"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetKeypairs(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keypairs"] = v
}

func (p *CreateVMFromBackupParams) ResetKeypairs() {
	if p.p != nil && p.p["keypairs"] != nil {
		delete(p.p, "keypairs")
	}
}

func (p *CreateVMFromBackupParams) GetKeypairs() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keypairs"].([]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetLeaseduration(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["leaseduration"] = v
}

func (p *CreateVMFromBackupParams) ResetLeaseduration() {
	if p.p != nil && p.p["leaseduration"] != nil {
		delete(p.p, "leaseduration")
	}
}

func (p *CreateVMFromBackupParams) GetLeaseduration() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["leaseduration"].(int)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetLeaseexpiryaction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["leaseexpiryaction"] = v
}

func (p *CreateVMFromBackupParams) ResetLeaseexpiryaction() {
	if p.p != nil && p.p["leaseexpiryaction"] != nil {
		delete(p.p, "leaseexpiryaction")
	}
}

func (p *CreateVMFromBackupParams) GetLeaseexpiryaction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["leaseexpiryaction"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetMacaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["macaddress"] = v
}

func (p *CreateVMFromBackupParams) ResetMacaddress() {
	if p.p != nil && p.p["macaddress"] != nil {
		delete(p.p, "macaddress")
	}
}

func (p *CreateVMFromBackupParams) GetMacaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["macaddress"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVMFromBackupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateVMFromBackupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetNetworkids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
}

func (p *CreateVMFromBackupParams) ResetNetworkids() {
	if p.p != nil && p.p["networkids"] != nil {
		delete(p.p, "networkids")
	}
}

func (p *CreateVMFromBackupParams) GetNetworkids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkids"].([]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetNicmultiqueuenumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicmultiqueuenumber"] = v
}

func (p *CreateVMFromBackupParams) ResetNicmultiqueuenumber() {
	if p.p != nil && p.p["nicmultiqueuenumber"] != nil {
		delete(p.p, "nicmultiqueuenumber")
	}
}

func (p *CreateVMFromBackupParams) GetNicmultiqueuenumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicmultiqueuenumber"].(int)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetNicnetworklist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicnetworklist"] = v
}

func (p *CreateVMFromBackupParams) ResetNicnetworklist() {
	if p.p != nil && p.p["nicnetworklist"] != nil {
		delete(p.p, "nicnetworklist")
	}
}

func (p *CreateVMFromBackupParams) GetNicnetworklist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicnetworklist"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetNicpackedvirtqueuesenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicpackedvirtqueuesenabled"] = v
}

func (p *CreateVMFromBackupParams) ResetNicpackedvirtqueuesenabled() {
	if p.p != nil && p.p["nicpackedvirtqueuesenabled"] != nil {
		delete(p.p, "nicpackedvirtqueuesenabled")
	}
}

func (p *CreateVMFromBackupParams) GetNicpackedvirtqueuesenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicpackedvirtqueuesenabled"].(bool)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetOverridediskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["overridediskofferingid"] = v
}

func (p *CreateVMFromBackupParams) ResetOverridediskofferingid() {
	if p.p != nil && p.p["overridediskofferingid"] != nil {
		delete(p.p, "overridediskofferingid")
	}
}

func (p *CreateVMFromBackupParams) GetOverridediskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["overridediskofferingid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *CreateVMFromBackupParams) ResetPassword() {
	if p.p != nil && p.p["password"] != nil {
		delete(p.p, "password")
	}
}

func (p *CreateVMFromBackupParams) GetPassword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["password"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *CreateVMFromBackupParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *CreateVMFromBackupParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetPreserveip(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["preserveip"] = v
}

func (p *CreateVMFromBackupParams) ResetPreserveip() {
	if p.p != nil && p.p["preserveip"] != nil {
		delete(p.p, "preserveip")
	}
}

func (p *CreateVMFromBackupParams) GetPreserveip() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["preserveip"].(bool)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateVMFromBackupParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateVMFromBackupParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetProperties(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["properties"] = v
}

func (p *CreateVMFromBackupParams) ResetProperties() {
	if p.p != nil && p.p["properties"] != nil {
		delete(p.p, "properties")
	}
}

func (p *CreateVMFromBackupParams) GetProperties() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["properties"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetRootdisksize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rootdisksize"] = v
}

func (p *CreateVMFromBackupParams) ResetRootdisksize() {
	if p.p != nil && p.p["rootdisksize"] != nil {
		delete(p.p, "rootdisksize")
	}
}

func (p *CreateVMFromBackupParams) GetRootdisksize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rootdisksize"].(int64)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetSecuritygroupids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupids"] = v
}

func (p *CreateVMFromBackupParams) ResetSecuritygroupids() {
	if p.p != nil && p.p["securitygroupids"] != nil {
		delete(p.p, "securitygroupids")
	}
}

func (p *CreateVMFromBackupParams) GetSecuritygroupids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupids"].([]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetSecuritygroupnames(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["securitygroupnames"] = v
}

func (p *CreateVMFromBackupParams) ResetSecuritygroupnames() {
	if p.p != nil && p.p["securitygroupnames"] != nil {
		delete(p.p, "securitygroupnames")
	}
}

func (p *CreateVMFromBackupParams) GetSecuritygroupnames() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["securitygroupnames"].([]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateVMFromBackupParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *CreateVMFromBackupParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *CreateVMFromBackupParams) ResetSize() {
	if p.p != nil && p.p["size"] != nil {
		delete(p.p, "size")
	}
}

func (p *CreateVMFromBackupParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetStartvm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startvm"] = v
}

func (p *CreateVMFromBackupParams) ResetStartvm() {
	if p.p != nil && p.p["startvm"] != nil {
		delete(p.p, "startvm")
	}
}

func (p *CreateVMFromBackupParams) GetStartvm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startvm"].(bool)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *CreateVMFromBackupParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *CreateVMFromBackupParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *CreateVMFromBackupParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *CreateVMFromBackupParams) GetUserdata() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetUserdatadetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdatadetails"] = v
}

func (p *CreateVMFromBackupParams) ResetUserdatadetails() {
	if p.p != nil && p.p["userdatadetails"] != nil {
		delete(p.p, "userdatadetails")
	}
}

func (p *CreateVMFromBackupParams) GetUserdatadetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdatadetails"].(map[string]string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetUserdataid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdataid"] = v
}

func (p *CreateVMFromBackupParams) ResetUserdataid() {
	if p.p != nil && p.p["userdataid"] != nil {
		delete(p.p, "userdataid")
	}
}

func (p *CreateVMFromBackupParams) GetUserdataid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdataid"].(string)
	return value, ok
}

func (p *CreateVMFromBackupParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateVMFromBackupParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateVMFromBackupParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVMFromBackupParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewCreateVMFromBackupParams(backupid string, zoneid string) *CreateVMFromBackupParams {
	p := &CreateVMFromBackupParams{}
	p.p = make(map[string]interface{})
	p.p["backupid"] = backupid
	p.p["zoneid"] = zoneid
	return p
}

// Creates and automatically starts a VM from a backup.
func (s *BackupService) CreateVMFromBackup(p *CreateVMFromBackupParams) (*CreateVMFromBackupResponse, error) {
	resp, err := s.cs.newPostRequest("createVMFromBackup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVMFromBackupResponse
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

type CreateVMFromBackupResponse struct {
	Account               string                                    `json:"account"`
	Affinitygroup         []CreateVMFromBackupResponseAffinitygroup `json:"affinitygroup"`
	Arch                  string                                    `json:"arch"`
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
	Gpucardid             string                                    `json:"gpucardid"`
	Gpucardname           string                                    `json:"gpucardname"`
	Gpucount              int                                       `json:"gpucount"`
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
	Leaseduration         int                                       `json:"leaseduration"`
	Leaseexpiryaction     string                                    `json:"leaseexpiryaction"`
	Leaseexpirydate       string                                    `json:"leaseexpirydate"`
	Maxheads              int64                                     `json:"maxheads"`
	Maxresolutionx        int64                                     `json:"maxresolutionx"`
	Maxresolutiony        int64                                     `json:"maxresolutiony"`
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
	Securitygroup         []CreateVMFromBackupResponseSecuritygroup `json:"securitygroup"`
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
	Vgpuprofileid         string                                    `json:"vgpuprofileid"`
	Vgpuprofilename       string                                    `json:"vgpuprofilename"`
	Videoram              int64                                     `json:"videoram"`
	Vmtype                string                                    `json:"vmtype"`
	Vnfdetails            map[string]string                         `json:"vnfdetails"`
	Vnfnics               []string                                  `json:"vnfnics"`
	Zoneid                string                                    `json:"zoneid"`
	Zonename              string                                    `json:"zonename"`
}

type CreateVMFromBackupResponseSecuritygroup struct {
	Account             string                                        `json:"account"`
	Description         string                                        `json:"description"`
	Domain              string                                        `json:"domain"`
	Domainid            string                                        `json:"domainid"`
	Domainpath          string                                        `json:"domainpath"`
	Egressrule          []CreateVMFromBackupResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                        `json:"id"`
	Ingressrule         []CreateVMFromBackupResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                        `json:"name"`
	Project             string                                        `json:"project"`
	Projectid           string                                        `json:"projectid"`
	Tags                []Tags                                        `json:"tags"`
	Virtualmachinecount int                                           `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                 `json:"virtualmachineids"`
}

type CreateVMFromBackupResponseSecuritygroupRule struct {
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

type CreateVMFromBackupResponseAffinitygroup struct {
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

func (r *CreateVMFromBackupResponse) UnmarshalJSON(b []byte) error {
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

	type alias CreateVMFromBackupResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteBackupParams struct {
	p map[string]interface{}
}

func (p *DeleteBackupParams) toURLValues() url.Values {
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
	return u
}

func (p *DeleteBackupParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DeleteBackupParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *DeleteBackupParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *DeleteBackupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteBackupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteBackupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBackupParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewDeleteBackupParams(id string) *DeleteBackupParams {
	p := &DeleteBackupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete VM backup
func (s *BackupService) DeleteBackup(p *DeleteBackupParams) (*DeleteBackupResponse, error) {
	resp, err := s.cs.newPostRequest("deleteBackup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBackupResponse
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

type DeleteBackupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteBackupOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteBackupOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteBackupOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteBackupOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteBackupOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBackupOfferingParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewDeleteBackupOfferingParams(id string) *DeleteBackupOfferingParams {
	p := &DeleteBackupOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a backup offering
func (s *BackupService) DeleteBackupOffering(p *DeleteBackupOfferingParams) (*DeleteBackupOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("deleteBackupOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBackupOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteBackupOfferingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteBackupOfferingResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteBackupOfferingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteBackupRepositoryParams struct {
	p map[string]interface{}
}

func (p *DeleteBackupRepositoryParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteBackupRepositoryParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteBackupRepositoryParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteBackupRepositoryParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBackupRepositoryParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewDeleteBackupRepositoryParams(id string) *DeleteBackupRepositoryParams {
	p := &DeleteBackupRepositoryParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// delete a backup repository
func (s *BackupService) DeleteBackupRepository(p *DeleteBackupRepositoryParams) (*DeleteBackupRepositoryResponse, error) {
	resp, err := s.cs.newPostRequest("deleteBackupRepository", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBackupRepositoryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteBackupRepositoryResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteBackupRepositoryResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteBackupRepositoryResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteBackupScheduleParams struct {
	p map[string]interface{}
}

func (p *DeleteBackupScheduleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *DeleteBackupScheduleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteBackupScheduleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteBackupScheduleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteBackupScheduleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *DeleteBackupScheduleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *DeleteBackupScheduleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBackupScheduleParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewDeleteBackupScheduleParams() *DeleteBackupScheduleParams {
	p := &DeleteBackupScheduleParams{}
	p.p = make(map[string]interface{})
	return p
}

// Deletes the backup schedule of a VM
func (s *BackupService) DeleteBackupSchedule(p *DeleteBackupScheduleParams) (*DeleteBackupScheduleResponse, error) {
	resp, err := s.cs.newPostRequest("deleteBackupSchedule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBackupScheduleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteBackupScheduleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteBackupScheduleResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteBackupScheduleResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ImportBackupOfferingParams struct {
	p map[string]interface{}
}

func (p *ImportBackupOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allowuserdrivenbackups"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("allowuserdrivenbackups", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["externalid"]; found {
		u.Set("externalid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ImportBackupOfferingParams) SetAllowuserdrivenbackups(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allowuserdrivenbackups"] = v
}

func (p *ImportBackupOfferingParams) ResetAllowuserdrivenbackups() {
	if p.p != nil && p.p["allowuserdrivenbackups"] != nil {
		delete(p.p, "allowuserdrivenbackups")
	}
}

func (p *ImportBackupOfferingParams) GetAllowuserdrivenbackups() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allowuserdrivenbackups"].(bool)
	return value, ok
}

func (p *ImportBackupOfferingParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *ImportBackupOfferingParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *ImportBackupOfferingParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *ImportBackupOfferingParams) SetExternalid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["externalid"] = v
}

func (p *ImportBackupOfferingParams) ResetExternalid() {
	if p.p != nil && p.p["externalid"] != nil {
		delete(p.p, "externalid")
	}
}

func (p *ImportBackupOfferingParams) GetExternalid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["externalid"].(string)
	return value, ok
}

func (p *ImportBackupOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ImportBackupOfferingParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ImportBackupOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ImportBackupOfferingParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ImportBackupOfferingParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ImportBackupOfferingParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ImportBackupOfferingParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewImportBackupOfferingParams(allowuserdrivenbackups bool, description string, externalid string, name string, zoneid string) *ImportBackupOfferingParams {
	p := &ImportBackupOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["allowuserdrivenbackups"] = allowuserdrivenbackups
	p.p["description"] = description
	p.p["externalid"] = externalid
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// Imports a backup offering using a backup provider
func (s *BackupService) ImportBackupOffering(p *ImportBackupOfferingParams) (*ImportBackupOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("importBackupOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportBackupOfferingResponse
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

type ImportBackupOfferingResponse struct {
	Allowuserdrivenbackups    bool   `json:"allowuserdrivenbackups"`
	Created                   string `json:"created"`
	Crosszoneinstancecreation bool   `json:"crosszoneinstancecreation"`
	Description               string `json:"description"`
	Externalid                string `json:"externalid"`
	Id                        string `json:"id"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Name                      string `json:"name"`
	Provider                  string `json:"provider"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type ListBackupOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListBackupOfferingsParams) toURLValues() url.Values {
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

func (p *ListBackupOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListBackupOfferingsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListBackupOfferingsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListBackupOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBackupOfferingsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBackupOfferingsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBackupOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBackupOfferingsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBackupOfferingsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBackupOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBackupOfferingsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBackupOfferingsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBackupOfferingsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListBackupOfferingsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListBackupOfferingsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBackupOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewListBackupOfferingsParams() *ListBackupOfferingsParams {
	p := &ListBackupOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupOfferingID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListBackupOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListBackupOfferings(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.BackupOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.BackupOfferings {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupOfferingByName(name string, opts ...OptionFunc) (*BackupOffering, int, error) {
	id, count, err := s.GetBackupOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetBackupOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupOfferingByID(id string, opts ...OptionFunc) (*BackupOffering, int, error) {
	p := &ListBackupOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListBackupOfferings(p)
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
		return l.BackupOfferings[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for BackupOffering UUID: %s!", id)
}

// Lists backup offerings
func (s *BackupService) ListBackupOfferings(p *ListBackupOfferingsParams) (*ListBackupOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listBackupOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBackupOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBackupOfferingsResponse struct {
	Count           int               `json:"count"`
	BackupOfferings []*BackupOffering `json:"backupoffering"`
}

type BackupOffering struct {
	Allowuserdrivenbackups    bool   `json:"allowuserdrivenbackups"`
	Created                   string `json:"created"`
	Crosszoneinstancecreation bool   `json:"crosszoneinstancecreation"`
	Description               string `json:"description"`
	Externalid                string `json:"externalid"`
	Id                        string `json:"id"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Name                      string `json:"name"`
	Provider                  string `json:"provider"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type ListBackupProviderOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListBackupProviderOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListBackupProviderOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBackupProviderOfferingsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBackupProviderOfferingsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBackupProviderOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBackupProviderOfferingsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBackupProviderOfferingsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBackupProviderOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBackupProviderOfferingsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBackupProviderOfferingsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBackupProviderOfferingsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListBackupProviderOfferingsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListBackupProviderOfferingsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBackupProviderOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewListBackupProviderOfferingsParams(zoneid string) *ListBackupProviderOfferingsParams {
	p := &ListBackupProviderOfferingsParams{}
	p.p = make(map[string]interface{})
	p.p["zoneid"] = zoneid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupProviderOfferingID(keyword string, zoneid string, opts ...OptionFunc) (string, int, error) {
	p := &ListBackupProviderOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["zoneid"] = zoneid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListBackupProviderOfferings(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.BackupProviderOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.BackupProviderOfferings {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists external backup offerings of the provider
func (s *BackupService) ListBackupProviderOfferings(p *ListBackupProviderOfferingsParams) (*ListBackupProviderOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listBackupProviderOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBackupProviderOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBackupProviderOfferingsResponse struct {
	Count                   int                       `json:"count"`
	BackupProviderOfferings []*BackupProviderOffering `json:"backupprovideroffering"`
}

type BackupProviderOffering struct {
	Allowuserdrivenbackups    bool   `json:"allowuserdrivenbackups"`
	Created                   string `json:"created"`
	Crosszoneinstancecreation bool   `json:"crosszoneinstancecreation"`
	Description               string `json:"description"`
	Externalid                string `json:"externalid"`
	Id                        string `json:"id"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Name                      string `json:"name"`
	Provider                  string `json:"provider"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type ListBackupProvidersParams struct {
	p map[string]interface{}
}

func (p *ListBackupProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *ListBackupProvidersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListBackupProvidersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListBackupProvidersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new ListBackupProvidersParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewListBackupProvidersParams() *ListBackupProvidersParams {
	p := &ListBackupProvidersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Backup and Recovery providers
func (s *BackupService) ListBackupProviders(p *ListBackupProvidersParams) (*ListBackupProvidersResponse, error) {
	resp, err := s.cs.newRequest("listBackupProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBackupProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBackupProvidersResponse struct {
	Count           int               `json:"count"`
	BackupProviders []*BackupProvider `json:"backupprovider"`
}

type BackupProvider struct {
	Description string `json:"description"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
}

type ListBackupRepositoriesParams struct {
	p map[string]interface{}
}

func (p *ListBackupRepositoriesParams) toURLValues() url.Values {
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
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListBackupRepositoriesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListBackupRepositoriesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListBackupRepositoriesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListBackupRepositoriesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBackupRepositoriesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBackupRepositoriesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBackupRepositoriesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListBackupRepositoriesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListBackupRepositoriesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListBackupRepositoriesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBackupRepositoriesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBackupRepositoriesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBackupRepositoriesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBackupRepositoriesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBackupRepositoriesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBackupRepositoriesParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListBackupRepositoriesParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListBackupRepositoriesParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *ListBackupRepositoriesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListBackupRepositoriesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListBackupRepositoriesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBackupRepositoriesParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewListBackupRepositoriesParams() *ListBackupRepositoriesParams {
	p := &ListBackupRepositoriesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupRepositoryID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListBackupRepositoriesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListBackupRepositories(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.BackupRepositories[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.BackupRepositories {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupRepositoryByName(name string, opts ...OptionFunc) (*BackupRepository, int, error) {
	id, count, err := s.GetBackupRepositoryID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetBackupRepositoryByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupRepositoryByID(id string, opts ...OptionFunc) (*BackupRepository, int, error) {
	p := &ListBackupRepositoriesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListBackupRepositories(p)
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
		return l.BackupRepositories[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for BackupRepository UUID: %s!", id)
}

// Lists all backup repositories
func (s *BackupService) ListBackupRepositories(p *ListBackupRepositoriesParams) (*ListBackupRepositoriesResponse, error) {
	resp, err := s.cs.newRequest("listBackupRepositories", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBackupRepositoriesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBackupRepositoriesResponse struct {
	Count              int                 `json:"count"`
	BackupRepositories []*BackupRepository `json:"backuprepository"`
}

type BackupRepository struct {
	Address                   string `json:"address"`
	Capacitybytes             int64  `json:"capacitybytes"`
	Created                   string `json:"created"`
	Crosszoneinstancecreation bool   `json:"crosszoneinstancecreation"`
	Id                        string `json:"id"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Name                      string `json:"name"`
	Provider                  string `json:"provider"`
	Type                      string `json:"type"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type ListBackupScheduleParams struct {
	p map[string]interface{}
}

func (p *ListBackupScheduleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
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
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListBackupScheduleParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListBackupScheduleParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListBackupScheduleParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListBackupScheduleParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListBackupScheduleParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListBackupScheduleParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListBackupScheduleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListBackupScheduleParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListBackupScheduleParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListBackupScheduleParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListBackupScheduleParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListBackupScheduleParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListBackupScheduleParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBackupScheduleParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBackupScheduleParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBackupScheduleParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListBackupScheduleParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListBackupScheduleParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListBackupScheduleParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBackupScheduleParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBackupScheduleParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBackupScheduleParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBackupScheduleParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBackupScheduleParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBackupScheduleParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListBackupScheduleParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListBackupScheduleParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListBackupScheduleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListBackupScheduleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *ListBackupScheduleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBackupScheduleParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewListBackupScheduleParams() *ListBackupScheduleParams {
	p := &ListBackupScheduleParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupScheduleByID(id string, opts ...OptionFunc) (*BackupSchedule, int, error) {
	p := &ListBackupScheduleParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListBackupSchedule(p)
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
		return l.BackupSchedule[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for BackupSchedule UUID: %s!", id)
}

// List backup schedule of a VM
func (s *BackupService) ListBackupSchedule(p *ListBackupScheduleParams) (*ListBackupScheduleResponse, error) {
	resp, err := s.cs.newRequest("listBackupSchedule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBackupScheduleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBackupScheduleResponse struct {
	Count          int               `json:"count"`
	BackupSchedule []*BackupSchedule `json:"backupschedule"`
}

type BackupSchedule struct {
	Id                 string `json:"id"`
	Intervaltype       string `json:"intervaltype"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Maxbackups         int    `json:"maxbackups"`
	Quiescevm          bool   `json:"quiescevm"`
	Schedule           string `json:"schedule"`
	Timezone           string `json:"timezone"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
}

type ListBackupsParams struct {
	p map[string]interface{}
}

func (p *ListBackupsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["backupofferingid"]; found {
		u.Set("backupofferingid", v.(string))
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
	if v, found := p.p["listvmdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listvmdetails", vv)
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListBackupsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListBackupsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListBackupsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListBackupsParams) SetBackupofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["backupofferingid"] = v
}

func (p *ListBackupsParams) ResetBackupofferingid() {
	if p.p != nil && p.p["backupofferingid"] != nil {
		delete(p.p, "backupofferingid")
	}
}

func (p *ListBackupsParams) GetBackupofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["backupofferingid"].(string)
	return value, ok
}

func (p *ListBackupsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListBackupsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListBackupsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListBackupsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListBackupsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListBackupsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListBackupsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListBackupsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListBackupsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListBackupsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListBackupsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListBackupsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListBackupsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListBackupsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListBackupsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListBackupsParams) SetListvmdetails(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listvmdetails"] = v
}

func (p *ListBackupsParams) ResetListvmdetails() {
	if p.p != nil && p.p["listvmdetails"] != nil {
		delete(p.p, "listvmdetails")
	}
}

func (p *ListBackupsParams) GetListvmdetails() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listvmdetails"].(bool)
	return value, ok
}

func (p *ListBackupsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListBackupsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListBackupsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListBackupsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListBackupsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListBackupsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListBackupsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListBackupsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListBackupsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListBackupsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListBackupsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListBackupsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListBackupsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListBackupsParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *ListBackupsParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListBackupsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListBackupsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListBackupsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBackupsParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewListBackupsParams() *ListBackupsParams {
	p := &ListBackupsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListBackupsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListBackups(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Backups[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Backups {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupByName(name string, opts ...OptionFunc) (*Backup, int, error) {
	id, count, err := s.GetBackupID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetBackupByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BackupService) GetBackupByID(id string, opts ...OptionFunc) (*Backup, int, error) {
	p := &ListBackupsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListBackups(p)
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
		return l.Backups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Backup UUID: %s!", id)
}

// Lists VM backups
func (s *BackupService) ListBackups(p *ListBackupsParams) (*ListBackupsResponse, error) {
	resp, err := s.cs.newRequest("listBackups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBackupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBackupsResponse struct {
	Count   int       `json:"count"`
	Backups []*Backup `json:"backup"`
}

type Backup struct {
	Account                 string            `json:"account"`
	Accountid               string            `json:"accountid"`
	Backupofferingid        string            `json:"backupofferingid"`
	Backupofferingname      string            `json:"backupofferingname"`
	Created                 string            `json:"created"`
	Description             string            `json:"description"`
	Domain                  string            `json:"domain"`
	Domainid                string            `json:"domainid"`
	Externalid              string            `json:"externalid"`
	Id                      string            `json:"id"`
	Intervaltype            string            `json:"intervaltype"`
	Isbackupvmexpunged      bool              `json:"isbackupvmexpunged"`
	JobID                   string            `json:"jobid"`
	Jobstatus               int               `json:"jobstatus"`
	Name                    string            `json:"name"`
	Size                    int64             `json:"size"`
	Status                  string            `json:"status"`
	Type                    string            `json:"type"`
	Virtualmachineid        string            `json:"virtualmachineid"`
	Virtualmachinename      string            `json:"virtualmachinename"`
	Virtualsize             int64             `json:"virtualsize"`
	Vmbackupofferingremoved bool              `json:"vmbackupofferingremoved"`
	Vmdetails               map[string]string `json:"vmdetails"`
	Volumes                 string            `json:"volumes"`
	Zone                    string            `json:"zone"`
	Zoneid                  string            `json:"zoneid"`
}

type RestoreBackupParams struct {
	p map[string]interface{}
}

func (p *RestoreBackupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RestoreBackupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RestoreBackupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RestoreBackupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RestoreBackupParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewRestoreBackupParams(id string) *RestoreBackupParams {
	p := &RestoreBackupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Restores an existing stopped or deleted VM using a VM backup
func (s *BackupService) RestoreBackup(p *RestoreBackupParams) (*RestoreBackupResponse, error) {
	resp, err := s.cs.newPostRequest("restoreBackup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestoreBackupResponse
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

type RestoreBackupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateBackupRepositoryParams struct {
	p map[string]interface{}
}

func (p *UpdateBackupRepositoryParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["address"]; found {
		u.Set("address", v.(string))
	}
	if v, found := p.p["crosszoneinstancecreation"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("crosszoneinstancecreation", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["mountopts"]; found {
		u.Set("mountopts", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateBackupRepositoryParams) SetAddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["address"] = v
}

func (p *UpdateBackupRepositoryParams) ResetAddress() {
	if p.p != nil && p.p["address"] != nil {
		delete(p.p, "address")
	}
}

func (p *UpdateBackupRepositoryParams) GetAddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["address"].(string)
	return value, ok
}

func (p *UpdateBackupRepositoryParams) SetCrosszoneinstancecreation(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["crosszoneinstancecreation"] = v
}

func (p *UpdateBackupRepositoryParams) ResetCrosszoneinstancecreation() {
	if p.p != nil && p.p["crosszoneinstancecreation"] != nil {
		delete(p.p, "crosszoneinstancecreation")
	}
}

func (p *UpdateBackupRepositoryParams) GetCrosszoneinstancecreation() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["crosszoneinstancecreation"].(bool)
	return value, ok
}

func (p *UpdateBackupRepositoryParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateBackupRepositoryParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateBackupRepositoryParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateBackupRepositoryParams) SetMountopts(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mountopts"] = v
}

func (p *UpdateBackupRepositoryParams) ResetMountopts() {
	if p.p != nil && p.p["mountopts"] != nil {
		delete(p.p, "mountopts")
	}
}

func (p *UpdateBackupRepositoryParams) GetMountopts() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mountopts"].(string)
	return value, ok
}

func (p *UpdateBackupRepositoryParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateBackupRepositoryParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateBackupRepositoryParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateBackupRepositoryParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewUpdateBackupRepositoryParams(id string) *UpdateBackupRepositoryParams {
	p := &UpdateBackupRepositoryParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Update a backup repository
func (s *BackupService) UpdateBackupRepository(p *UpdateBackupRepositoryParams) (*UpdateBackupRepositoryResponse, error) {
	resp, err := s.cs.newPostRequest("updateBackupRepository", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateBackupRepositoryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateBackupRepositoryResponse struct {
	Address                   string `json:"address"`
	Capacitybytes             int64  `json:"capacitybytes"`
	Created                   string `json:"created"`
	Crosszoneinstancecreation bool   `json:"crosszoneinstancecreation"`
	Id                        string `json:"id"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Name                      string `json:"name"`
	Provider                  string `json:"provider"`
	Type                      string `json:"type"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type UpdateBackupOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateBackupOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["allowuserdrivenbackups"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("allowuserdrivenbackups", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateBackupOfferingParams) SetAllowuserdrivenbackups(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["allowuserdrivenbackups"] = v
}

func (p *UpdateBackupOfferingParams) ResetAllowuserdrivenbackups() {
	if p.p != nil && p.p["allowuserdrivenbackups"] != nil {
		delete(p.p, "allowuserdrivenbackups")
	}
}

func (p *UpdateBackupOfferingParams) GetAllowuserdrivenbackups() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["allowuserdrivenbackups"].(bool)
	return value, ok
}

func (p *UpdateBackupOfferingParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateBackupOfferingParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateBackupOfferingParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateBackupOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateBackupOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateBackupOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateBackupOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateBackupOfferingParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateBackupOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateBackupOfferingParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewUpdateBackupOfferingParams(id string) *UpdateBackupOfferingParams {
	p := &UpdateBackupOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a backup offering.
func (s *BackupService) UpdateBackupOffering(p *UpdateBackupOfferingParams) (*UpdateBackupOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("updateBackupOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateBackupOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateBackupOfferingResponse struct {
	Allowuserdrivenbackups    bool   `json:"allowuserdrivenbackups"`
	Created                   string `json:"created"`
	Crosszoneinstancecreation bool   `json:"crosszoneinstancecreation"`
	Description               string `json:"description"`
	Externalid                string `json:"externalid"`
	Id                        string `json:"id"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Name                      string `json:"name"`
	Provider                  string `json:"provider"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type UpdateBackupScheduleParams struct {
	p map[string]interface{}
}

func (p *UpdateBackupScheduleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["intervaltype"]; found {
		u.Set("intervaltype", v.(string))
	}
	if v, found := p.p["maxbackups"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxbackups", vv)
	}
	if v, found := p.p["quiescevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("quiescevm", vv)
	}
	if v, found := p.p["schedule"]; found {
		u.Set("schedule", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *UpdateBackupScheduleParams) SetIntervaltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltype"] = v
}

func (p *UpdateBackupScheduleParams) ResetIntervaltype() {
	if p.p != nil && p.p["intervaltype"] != nil {
		delete(p.p, "intervaltype")
	}
}

func (p *UpdateBackupScheduleParams) GetIntervaltype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["intervaltype"].(string)
	return value, ok
}

func (p *UpdateBackupScheduleParams) SetMaxbackups(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxbackups"] = v
}

func (p *UpdateBackupScheduleParams) ResetMaxbackups() {
	if p.p != nil && p.p["maxbackups"] != nil {
		delete(p.p, "maxbackups")
	}
}

func (p *UpdateBackupScheduleParams) GetMaxbackups() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxbackups"].(int)
	return value, ok
}

func (p *UpdateBackupScheduleParams) SetQuiescevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiescevm"] = v
}

func (p *UpdateBackupScheduleParams) ResetQuiescevm() {
	if p.p != nil && p.p["quiescevm"] != nil {
		delete(p.p, "quiescevm")
	}
}

func (p *UpdateBackupScheduleParams) GetQuiescevm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quiescevm"].(bool)
	return value, ok
}

func (p *UpdateBackupScheduleParams) SetSchedule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["schedule"] = v
}

func (p *UpdateBackupScheduleParams) ResetSchedule() {
	if p.p != nil && p.p["schedule"] != nil {
		delete(p.p, "schedule")
	}
}

func (p *UpdateBackupScheduleParams) GetSchedule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["schedule"].(string)
	return value, ok
}

func (p *UpdateBackupScheduleParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *UpdateBackupScheduleParams) ResetTimezone() {
	if p.p != nil && p.p["timezone"] != nil {
		delete(p.p, "timezone")
	}
}

func (p *UpdateBackupScheduleParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

func (p *UpdateBackupScheduleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *UpdateBackupScheduleParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *UpdateBackupScheduleParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateBackupScheduleParams instance,
// as then you are sure you have configured all required params
func (s *BackupService) NewUpdateBackupScheduleParams(intervaltype string, schedule string, timezone string, virtualmachineid string) *UpdateBackupScheduleParams {
	p := &UpdateBackupScheduleParams{}
	p.p = make(map[string]interface{})
	p.p["intervaltype"] = intervaltype
	p.p["schedule"] = schedule
	p.p["timezone"] = timezone
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Updates a user-defined VM backup schedule
func (s *BackupService) UpdateBackupSchedule(p *UpdateBackupScheduleParams) (*UpdateBackupScheduleResponse, error) {
	resp, err := s.cs.newPostRequest("updateBackupSchedule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateBackupScheduleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateBackupScheduleResponse struct {
	Account                 string            `json:"account"`
	Accountid               string            `json:"accountid"`
	Backupofferingid        string            `json:"backupofferingid"`
	Backupofferingname      string            `json:"backupofferingname"`
	Created                 string            `json:"created"`
	Description             string            `json:"description"`
	Domain                  string            `json:"domain"`
	Domainid                string            `json:"domainid"`
	Externalid              string            `json:"externalid"`
	Id                      string            `json:"id"`
	Intervaltype            string            `json:"intervaltype"`
	Isbackupvmexpunged      bool              `json:"isbackupvmexpunged"`
	JobID                   string            `json:"jobid"`
	Jobstatus               int               `json:"jobstatus"`
	Name                    string            `json:"name"`
	Size                    int64             `json:"size"`
	Status                  string            `json:"status"`
	Type                    string            `json:"type"`
	Virtualmachineid        string            `json:"virtualmachineid"`
	Virtualmachinename      string            `json:"virtualmachinename"`
	Virtualsize             int64             `json:"virtualsize"`
	Vmbackupofferingremoved bool              `json:"vmbackupofferingremoved"`
	Vmdetails               map[string]string `json:"vmdetails"`
	Volumes                 string            `json:"volumes"`
	Zone                    string            `json:"zone"`
	Zoneid                  string            `json:"zoneid"`
}
