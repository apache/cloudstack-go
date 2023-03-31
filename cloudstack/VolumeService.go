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

type VolumeServiceIface interface {
	AttachVolume(p *AttachVolumeParams) (*AttachVolumeResponse, error)
	NewAttachVolumeParams(id string, virtualmachineid string) *AttachVolumeParams
	CreateVolume(p *CreateVolumeParams) (*CreateVolumeResponse, error)
	NewCreateVolumeParams() *CreateVolumeParams
	DeleteVolume(p *DeleteVolumeParams) (*DeleteVolumeResponse, error)
	NewDeleteVolumeParams(id string) *DeleteVolumeParams
	DestroyVolume(p *DestroyVolumeParams) (*DestroyVolumeResponse, error)
	NewDestroyVolumeParams(id string) *DestroyVolumeParams
	DetachVolume(p *DetachVolumeParams) (*DetachVolumeResponse, error)
	NewDetachVolumeParams() *DetachVolumeParams
	ExtractVolume(p *ExtractVolumeParams) (*ExtractVolumeResponse, error)
	NewExtractVolumeParams(id string, mode string, zoneid string) *ExtractVolumeParams
	GetPathForVolume(p *GetPathForVolumeParams) (*GetPathForVolumeResponse, error)
	NewGetPathForVolumeParams(volumeid string) *GetPathForVolumeParams
	GetSolidFireVolumeSize(p *GetSolidFireVolumeSizeParams) (*GetSolidFireVolumeSizeResponse, error)
	NewGetSolidFireVolumeSizeParams(volumeid string) *GetSolidFireVolumeSizeParams
	GetUploadParamsForVolume(p *GetUploadParamsForVolumeParams) (*GetUploadParamsForVolumeResponse, error)
	NewGetUploadParamsForVolumeParams(format string, name string, zoneid string) *GetUploadParamsForVolumeParams
	GetVolumeiScsiName(p *GetVolumeiScsiNameParams) (*GetVolumeiScsiNameResponse, error)
	NewGetVolumeiScsiNameParams(volumeid string) *GetVolumeiScsiNameParams
	ListVolumes(p *ListVolumesParams) (*ListVolumesResponse, error)
	NewListVolumesParams() *ListVolumesParams
	GetVolumeID(name string, opts ...OptionFunc) (string, int, error)
	GetVolumeByName(name string, opts ...OptionFunc) (*Volume, int, error)
	GetVolumeByID(id string, opts ...OptionFunc) (*Volume, int, error)
	ListVolumesMetrics(p *ListVolumesMetricsParams) (*ListVolumesMetricsResponse, error)
	NewListVolumesMetricsParams() *ListVolumesMetricsParams
	GetVolumesMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetVolumesMetricByName(name string, opts ...OptionFunc) (*VolumesMetric, int, error)
	GetVolumesMetricByID(id string, opts ...OptionFunc) (*VolumesMetric, int, error)
	MigrateVolume(p *MigrateVolumeParams) (*MigrateVolumeResponse, error)
	NewMigrateVolumeParams(storageid string, volumeid string) *MigrateVolumeParams
	RecoverVolume(p *RecoverVolumeParams) (*RecoverVolumeResponse, error)
	NewRecoverVolumeParams(id string) *RecoverVolumeParams
	ResizeVolume(p *ResizeVolumeParams) (*ResizeVolumeResponse, error)
	NewResizeVolumeParams(id string) *ResizeVolumeParams
	UpdateVolume(p *UpdateVolumeParams) (*UpdateVolumeResponse, error)
	NewUpdateVolumeParams() *UpdateVolumeParams
	UploadVolume(p *UploadVolumeParams) (*UploadVolumeResponse, error)
	NewUploadVolumeParams(format string, name string, url string, zoneid string) *UploadVolumeParams
	ChangeOfferingForVolume(p *ChangeOfferingForVolumeParams) (*ChangeOfferingForVolumeResponse, error)
	NewChangeOfferingForVolumeParams(diskofferingid string, id string) *ChangeOfferingForVolumeParams
}

type AttachVolumeParams struct {
	p map[string]interface{}
}

func (p *AttachVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["deviceid"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("deviceid", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AttachVolumeParams) SetDeviceid(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deviceid"] = v
}

func (p *AttachVolumeParams) GetDeviceid() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deviceid"].(int64)
	return value, ok
}

func (p *AttachVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *AttachVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *AttachVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *AttachVolumeParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new AttachVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewAttachVolumeParams(id string, virtualmachineid string) *AttachVolumeParams {
	p := &AttachVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attaches a disk volume to a virtual machine.
func (s *VolumeService) AttachVolume(p *AttachVolumeParams) (*AttachVolumeResponse, error) {
	resp, err := s.cs.newRequest("attachVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AttachVolumeResponse
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

type AttachVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type CreateVolumeParams struct {
	p map[string]interface{}
}

func (p *CreateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := p.p["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateVolumeParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateVolumeParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateVolumeParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *CreateVolumeParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *CreateVolumeParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *CreateVolumeParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *CreateVolumeParams) SetDisplayvolume(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvolume"] = v
}

func (p *CreateVolumeParams) GetDisplayvolume() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvolume"].(bool)
	return value, ok
}

func (p *CreateVolumeParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateVolumeParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateVolumeParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *CreateVolumeParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *CreateVolumeParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *CreateVolumeParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *CreateVolumeParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVolumeParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateVolumeParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateVolumeParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateVolumeParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *CreateVolumeParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

func (p *CreateVolumeParams) SetSnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshotid"] = v
}

func (p *CreateVolumeParams) GetSnapshotid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["snapshotid"].(string)
	return value, ok
}

func (p *CreateVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *CreateVolumeParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *CreateVolumeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateVolumeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewCreateVolumeParams() *CreateVolumeParams {
	p := &CreateVolumeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Creates a disk volume from a disk offering. This disk volume must still be attached to a virtual machine to make use of it.
func (s *VolumeService) CreateVolume(p *CreateVolumeParams) (*CreateVolumeResponse, error) {
	resp, err := s.cs.newRequest("createVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVolumeResponse
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

type CreateVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type DeleteVolumeParams struct {
	p map[string]interface{}
}

func (p *DeleteVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDeleteVolumeParams(id string) *DeleteVolumeParams {
	p := &DeleteVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a detached disk volume.
func (s *VolumeService) DeleteVolume(p *DeleteVolumeParams) (*DeleteVolumeResponse, error) {
	resp, err := s.cs.newRequest("deleteVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVolumeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteVolumeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteVolumeResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteVolumeResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DestroyVolumeParams struct {
	p map[string]interface{}
}

func (p *DestroyVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["expunge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("expunge", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DestroyVolumeParams) SetExpunge(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expunge"] = v
}

func (p *DestroyVolumeParams) GetExpunge() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["expunge"].(bool)
	return value, ok
}

func (p *DestroyVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DestroyVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DestroyVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDestroyVolumeParams(id string) *DestroyVolumeParams {
	p := &DestroyVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroys a Volume.
func (s *VolumeService) DestroyVolume(p *DestroyVolumeParams) (*DestroyVolumeResponse, error) {
	resp, err := s.cs.newRequest("destroyVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroyVolumeResponse
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

type DestroyVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type DetachVolumeParams struct {
	p map[string]interface{}
}

func (p *DetachVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["deviceid"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("deviceid", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *DetachVolumeParams) SetDeviceid(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deviceid"] = v
}

func (p *DetachVolumeParams) GetDeviceid() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deviceid"].(int64)
	return value, ok
}

func (p *DetachVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DetachVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DetachVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *DetachVolumeParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new DetachVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDetachVolumeParams() *DetachVolumeParams {
	p := &DetachVolumeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Detaches a disk volume from a virtual machine.
func (s *VolumeService) DetachVolume(p *DetachVolumeParams) (*DetachVolumeResponse, error) {
	resp, err := s.cs.newRequest("detachVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DetachVolumeResponse
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

type DetachVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type ExtractVolumeParams struct {
	p map[string]interface{}
}

func (p *ExtractVolumeParams) toURLValues() url.Values {
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

func (p *ExtractVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ExtractVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ExtractVolumeParams) SetMode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mode"] = v
}

func (p *ExtractVolumeParams) GetMode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mode"].(string)
	return value, ok
}

func (p *ExtractVolumeParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *ExtractVolumeParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *ExtractVolumeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ExtractVolumeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ExtractVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewExtractVolumeParams(id string, mode string, zoneid string) *ExtractVolumeParams {
	p := &ExtractVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["mode"] = mode
	p.p["zoneid"] = zoneid
	return p
}

// Extracts volume
func (s *VolumeService) ExtractVolume(p *ExtractVolumeParams) (*ExtractVolumeResponse, error) {
	resp, err := s.cs.newRequest("extractVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractVolumeResponse
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

type ExtractVolumeResponse struct {
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

type GetPathForVolumeParams struct {
	p map[string]interface{}
}

func (p *GetPathForVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *GetPathForVolumeParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *GetPathForVolumeParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new GetPathForVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewGetPathForVolumeParams(volumeid string) *GetPathForVolumeParams {
	p := &GetPathForVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["volumeid"] = volumeid
	return p
}

// Get the path associated with the provided volume UUID
func (s *VolumeService) GetPathForVolume(p *GetPathForVolumeParams) (*GetPathForVolumeResponse, error) {
	resp, err := s.cs.newRequest("getPathForVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetPathForVolumeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetPathForVolumeResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Path      string `json:"path"`
}

type GetSolidFireVolumeSizeParams struct {
	p map[string]interface{}
}

func (p *GetSolidFireVolumeSizeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *GetSolidFireVolumeSizeParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *GetSolidFireVolumeSizeParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new GetSolidFireVolumeSizeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewGetSolidFireVolumeSizeParams(volumeid string) *GetSolidFireVolumeSizeParams {
	p := &GetSolidFireVolumeSizeParams{}
	p.p = make(map[string]interface{})
	p.p["volumeid"] = volumeid
	return p
}

// Get the SF volume size including Hypervisor Snapshot Reserve
func (s *VolumeService) GetSolidFireVolumeSize(p *GetSolidFireVolumeSizeParams) (*GetSolidFireVolumeSizeResponse, error) {
	resp, err := s.cs.newRequest("getSolidFireVolumeSize", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetSolidFireVolumeSizeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetSolidFireVolumeSizeResponse struct {
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	SolidFireVolumeSize int64  `json:"solidFireVolumeSize"`
}

type GetUploadParamsForVolumeParams struct {
	p map[string]interface{}
}

func (p *GetUploadParamsForVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := p.p["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *GetUploadParamsForVolumeParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *GetUploadParamsForVolumeParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *GetUploadParamsForVolumeParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
}

func (p *GetUploadParamsForVolumeParams) GetChecksum() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["checksum"].(string)
	return value, ok
}

func (p *GetUploadParamsForVolumeParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *GetUploadParamsForVolumeParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *GetUploadParamsForVolumeParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *GetUploadParamsForVolumeParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *GetUploadParamsForVolumeParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *GetUploadParamsForVolumeParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *GetUploadParamsForVolumeParams) SetImagestoreuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["imagestoreuuid"] = v
}

func (p *GetUploadParamsForVolumeParams) GetImagestoreuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["imagestoreuuid"].(string)
	return value, ok
}

func (p *GetUploadParamsForVolumeParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *GetUploadParamsForVolumeParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *GetUploadParamsForVolumeParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *GetUploadParamsForVolumeParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *GetUploadParamsForVolumeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *GetUploadParamsForVolumeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new GetUploadParamsForVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewGetUploadParamsForVolumeParams(format string, name string, zoneid string) *GetUploadParamsForVolumeParams {
	p := &GetUploadParamsForVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["format"] = format
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// Upload a data disk to the cloudstack cloud.
func (s *VolumeService) GetUploadParamsForVolume(p *GetUploadParamsForVolumeParams) (*GetUploadParamsForVolumeResponse, error) {
	resp, err := s.cs.newRequest("getUploadParamsForVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response GetUploadParamsForVolumeResponse `json:"getuploadparams"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type GetUploadParamsForVolumeResponse struct {
	Expires   string `json:"expires"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Metadata  string `json:"metadata"`
	PostURL   string `json:"postURL"`
	Signature string `json:"signature"`
}

type GetVolumeiScsiNameParams struct {
	p map[string]interface{}
}

func (p *GetVolumeiScsiNameParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *GetVolumeiScsiNameParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *GetVolumeiScsiNameParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new GetVolumeiScsiNameParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewGetVolumeiScsiNameParams(volumeid string) *GetVolumeiScsiNameParams {
	p := &GetVolumeiScsiNameParams{}
	p.p = make(map[string]interface{})
	p.p["volumeid"] = volumeid
	return p
}

// Get Volume's iSCSI Name
func (s *VolumeService) GetVolumeiScsiName(p *GetVolumeiScsiNameParams) (*GetVolumeiScsiNameResponse, error) {
	resp, err := s.cs.newRequest("getVolumeiScsiName", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetVolumeiScsiNameResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetVolumeiScsiNameResponse struct {
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	VolumeiScsiName string `json:"volumeiScsiName"`
}

type ListVolumesParams struct {
	p map[string]interface{}
}

func (p *ListVolumesParams) toURLValues() url.Values {
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
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
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
	if v, found := p.p["listsystemvms"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listsystemvms", vv)
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
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
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVolumesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVolumesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListVolumesParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *ListVolumesParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetDisplayvolume(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvolume"] = v
}

func (p *ListVolumesParams) GetDisplayvolume() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvolume"].(bool)
	return value, ok
}

func (p *ListVolumesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVolumesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListVolumesParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVolumesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListVolumesParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListVolumesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVolumesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVolumesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVolumesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVolumesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVolumesParams) SetListsystemvms(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listsystemvms"] = v
}

func (p *ListVolumesParams) GetListsystemvms() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listsystemvms"].(bool)
	return value, ok
}

func (p *ListVolumesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVolumesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVolumesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVolumesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVolumesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVolumesParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListVolumesParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVolumesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVolumesParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListVolumesParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVolumesParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListVolumesParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListVolumesParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListVolumesParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListVolumesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVolumesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVolumesParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewListVolumesParams() *ListVolumesParams {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVolumes(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Volumes[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Volumes {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeByName(name string, opts ...OptionFunc) (*Volume, int, error) {
	id, count, err := s.GetVolumeID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVolumeByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeByID(id string, opts ...OptionFunc) (*Volume, int, error) {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVolumes(p)
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
		return l.Volumes[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Volume UUID: %s!", id)
}

// Lists all volumes.
func (s *VolumeService) ListVolumes(p *ListVolumesParams) (*ListVolumesResponse, error) {
	resp, err := s.cs.newRequest("listVolumes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVolumesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVolumesResponse struct {
	Count   int       `json:"count"`
	Volumes []*Volume `json:"volume"`
}

type Volume struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type ListVolumesMetricsParams struct {
	p map[string]interface{}
}

func (p *ListVolumesMetricsParams) toURLValues() url.Values {
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
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
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
	if v, found := p.p["listsystemvms"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listsystemvms", vv)
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
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
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
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVolumesMetricsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVolumesMetricsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListVolumesMetricsParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *ListVolumesMetricsParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetDisplayvolume(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvolume"] = v
}

func (p *ListVolumesMetricsParams) GetDisplayvolume() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvolume"].(bool)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVolumesMetricsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListVolumesMetricsParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListVolumesMetricsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListVolumesMetricsParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVolumesMetricsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVolumesMetricsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVolumesMetricsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetListsystemvms(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listsystemvms"] = v
}

func (p *ListVolumesMetricsParams) GetListsystemvms() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listsystemvms"].(bool)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVolumesMetricsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVolumesMetricsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVolumesMetricsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListVolumesMetricsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVolumesMetricsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVolumesMetricsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListVolumesMetricsParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVolumesMetricsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListVolumesMetricsParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListVolumesMetricsParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListVolumesMetricsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVolumesMetricsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVolumesMetricsParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewListVolumesMetricsParams() *ListVolumesMetricsParams {
	p := &ListVolumesMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumesMetricID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVolumesMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVolumesMetrics(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VolumesMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VolumesMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumesMetricByName(name string, opts ...OptionFunc) (*VolumesMetric, int, error) {
	id, count, err := s.GetVolumesMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVolumesMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumesMetricByID(id string, opts ...OptionFunc) (*VolumesMetric, int, error) {
	p := &ListVolumesMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVolumesMetrics(p)
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
		return l.VolumesMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VolumesMetric UUID: %s!", id)
}

// Lists volume metrics
func (s *VolumeService) ListVolumesMetrics(p *ListVolumesMetricsParams) (*ListVolumesMetricsResponse, error) {
	resp, err := s.cs.newRequest("listVolumesMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVolumesMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVolumesMetricsResponse struct {
	Count          int              `json:"count"`
	VolumesMetrics []*VolumesMetric `json:"volumesmetric"`
}

type VolumesMetric struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskiopstotal              int64  `json:"diskiopstotal"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Sizegb                     string `json:"sizegb"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type MigrateVolumeParams struct {
	p map[string]interface{}
}

func (p *MigrateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["livemigrate"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("livemigrate", vv)
	}
	if v, found := p.p["newdiskofferingid"]; found {
		u.Set("newdiskofferingid", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *MigrateVolumeParams) SetLivemigrate(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["livemigrate"] = v
}

func (p *MigrateVolumeParams) GetLivemigrate() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["livemigrate"].(bool)
	return value, ok
}

func (p *MigrateVolumeParams) SetNewdiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["newdiskofferingid"] = v
}

func (p *MigrateVolumeParams) GetNewdiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["newdiskofferingid"].(string)
	return value, ok
}

func (p *MigrateVolumeParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *MigrateVolumeParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *MigrateVolumeParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *MigrateVolumeParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewMigrateVolumeParams(storageid string, volumeid string) *MigrateVolumeParams {
	p := &MigrateVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["storageid"] = storageid
	p.p["volumeid"] = volumeid
	return p
}

// Migrate volume
func (s *VolumeService) MigrateVolume(p *MigrateVolumeParams) (*MigrateVolumeResponse, error) {
	resp, err := s.cs.newRequest("migrateVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVolumeResponse
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

type MigrateVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type RecoverVolumeParams struct {
	p map[string]interface{}
}

func (p *RecoverVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RecoverVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RecoverVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RecoverVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewRecoverVolumeParams(id string) *RecoverVolumeParams {
	p := &RecoverVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Recovers a Destroy volume.
func (s *VolumeService) RecoverVolume(p *RecoverVolumeParams) (*RecoverVolumeResponse, error) {
	resp, err := s.cs.newRequest("recoverVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RecoverVolumeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RecoverVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type ResizeVolumeParams struct {
	p map[string]interface{}
}

func (p *ResizeVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["shrinkok"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("shrinkok", vv)
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	return u
}

func (p *ResizeVolumeParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *ResizeVolumeParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *ResizeVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ResizeVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ResizeVolumeParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *ResizeVolumeParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *ResizeVolumeParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *ResizeVolumeParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *ResizeVolumeParams) SetShrinkok(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["shrinkok"] = v
}

func (p *ResizeVolumeParams) GetShrinkok() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["shrinkok"].(bool)
	return value, ok
}

func (p *ResizeVolumeParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *ResizeVolumeParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

// You should always use this function to get a new ResizeVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewResizeVolumeParams(id string) *ResizeVolumeParams {
	p := &ResizeVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Resizes a volume
func (s *VolumeService) ResizeVolume(p *ResizeVolumeParams) (*ResizeVolumeResponse, error) {
	resp, err := s.cs.newRequest("resizeVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResizeVolumeResponse
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

type ResizeVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type UpdateVolumeParams struct {
	p map[string]interface{}
}

func (p *UpdateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["chaininfo"]; found {
		u.Set("chaininfo", v.(string))
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	return u
}

func (p *UpdateVolumeParams) SetChaininfo(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["chaininfo"] = v
}

func (p *UpdateVolumeParams) GetChaininfo() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["chaininfo"].(string)
	return value, ok
}

func (p *UpdateVolumeParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateVolumeParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateVolumeParams) SetDisplayvolume(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayvolume"] = v
}

func (p *UpdateVolumeParams) GetDisplayvolume() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayvolume"].(bool)
	return value, ok
}

func (p *UpdateVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateVolumeParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateVolumeParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateVolumeParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *UpdateVolumeParams) GetPath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["path"].(string)
	return value, ok
}

func (p *UpdateVolumeParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdateVolumeParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *UpdateVolumeParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *UpdateVolumeParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewUpdateVolumeParams() *UpdateVolumeParams {
	p := &UpdateVolumeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Updates the volume.
func (s *VolumeService) UpdateVolume(p *UpdateVolumeParams) (*UpdateVolumeResponse, error) {
	resp, err := s.cs.newRequest("updateVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVolumeResponse
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

type UpdateVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type UploadVolumeParams struct {
	p map[string]interface{}
}

func (p *UploadVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := p.p["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
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

func (p *UploadVolumeParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *UploadVolumeParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetChecksum(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["checksum"] = v
}

func (p *UploadVolumeParams) GetChecksum() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["checksum"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *UploadVolumeParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UploadVolumeParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetFormat(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["format"] = v
}

func (p *UploadVolumeParams) GetFormat() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["format"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetImagestoreuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["imagestoreuuid"] = v
}

func (p *UploadVolumeParams) GetImagestoreuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["imagestoreuuid"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UploadVolumeParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *UploadVolumeParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *UploadVolumeParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *UploadVolumeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *UploadVolumeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UploadVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewUploadVolumeParams(format string, name string, url string, zoneid string) *UploadVolumeParams {
	p := &UploadVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["format"] = format
	p.p["name"] = name
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Uploads a data disk.
func (s *VolumeService) UploadVolume(p *UploadVolumeParams) (*UploadVolumeResponse, error) {
	resp, err := s.cs.newRequest("uploadVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadVolumeResponse
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

type UploadVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type ChangeOfferingForVolumeParams struct {
	p map[string]interface{}
}

func (p *ChangeOfferingForVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["automigrate"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("automigrate", vv)
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["shrinkok"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("shrinkok", vv)
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	return u
}

func (p *ChangeOfferingForVolumeParams) SetAutomigrate(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["automigrate"] = v
}

func (p *ChangeOfferingForVolumeParams) GetAutomigrate() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["automigrate"].(bool)
	return value, ok
}

func (p *ChangeOfferingForVolumeParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *ChangeOfferingForVolumeParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *ChangeOfferingForVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ChangeOfferingForVolumeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ChangeOfferingForVolumeParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *ChangeOfferingForVolumeParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *ChangeOfferingForVolumeParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *ChangeOfferingForVolumeParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *ChangeOfferingForVolumeParams) SetShrinkok(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["shrinkok"] = v
}

func (p *ChangeOfferingForVolumeParams) GetShrinkok() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["shrinkok"].(bool)
	return value, ok
}

func (p *ChangeOfferingForVolumeParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
}

func (p *ChangeOfferingForVolumeParams) GetSize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["size"].(int64)
	return value, ok
}

// You should always use this function to get a new ChangeOfferingForVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewChangeOfferingForVolumeParams(diskofferingid string, id string) *ChangeOfferingForVolumeParams {
	p := &ChangeOfferingForVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["diskofferingid"] = diskofferingid
	p.p["id"] = id
	return p
}

// Change disk offering of the volume and also an option to auto migrate if required to apply the new disk offering
func (s *VolumeService) ChangeOfferingForVolume(p *ChangeOfferingForVolumeParams) (*ChangeOfferingForVolumeResponse, error) {
	resp, err := s.cs.newRequest("changeOfferingForVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeOfferingForVolumeResponse
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

type ChangeOfferingForVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Externaluuid               string `json:"externaluuid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Vmtype                     string `json:"vmtype"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}
