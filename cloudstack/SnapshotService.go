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

type SnapshotServiceIface interface {
	ArchiveSnapshot(p *ArchiveSnapshotParams) (*ArchiveSnapshotResponse, error)
	NewArchiveSnapshotParams(id string) *ArchiveSnapshotParams
	CopySnapshot(p *CopySnapshotParams) (*CopySnapshotResponse, error)
	NewCopySnapshotParams(id string) *CopySnapshotParams
	CreateSnapshot(p *CreateSnapshotParams) (*CreateSnapshotResponse, error)
	NewCreateSnapshotParams(volumeid string) *CreateSnapshotParams
	CreateSnapshotFromVMSnapshot(p *CreateSnapshotFromVMSnapshotParams) (*CreateSnapshotFromVMSnapshotResponse, error)
	NewCreateSnapshotFromVMSnapshotParams(vmsnapshotid string, volumeid string) *CreateSnapshotFromVMSnapshotParams
	CreateSnapshotPolicy(p *CreateSnapshotPolicyParams) (*CreateSnapshotPolicyResponse, error)
	NewCreateSnapshotPolicyParams(intervaltype string, maxsnaps int, schedule string, timezone string, volumeid string) *CreateSnapshotPolicyParams
	CreateVMSnapshot(p *CreateVMSnapshotParams) (*CreateVMSnapshotResponse, error)
	NewCreateVMSnapshotParams(virtualmachineid string) *CreateVMSnapshotParams
	DeleteSnapshot(p *DeleteSnapshotParams) (*DeleteSnapshotResponse, error)
	NewDeleteSnapshotParams(id string) *DeleteSnapshotParams
	DeleteSnapshotPolicies(p *DeleteSnapshotPoliciesParams) (*DeleteSnapshotPoliciesResponse, error)
	NewDeleteSnapshotPoliciesParams() *DeleteSnapshotPoliciesParams
	DeleteVMSnapshot(p *DeleteVMSnapshotParams) (*DeleteVMSnapshotResponse, error)
	NewDeleteVMSnapshotParams(vmsnapshotid string) *DeleteVMSnapshotParams
	ExtractSnapshot(p *ExtractSnapshotParams) (*ExtractSnapshotResponse, error)
	NewExtractSnapshotParams(id string, zoneid string) *ExtractSnapshotParams
	ListSnapshotPolicies(p *ListSnapshotPoliciesParams) (*ListSnapshotPoliciesResponse, error)
	NewListSnapshotPoliciesParams() *ListSnapshotPoliciesParams
	GetSnapshotPolicyByID(id string, opts ...OptionFunc) (*SnapshotPolicy, int, error)
	ListSnapshots(p *ListSnapshotsParams) (*ListSnapshotsResponse, error)
	NewListSnapshotsParams() *ListSnapshotsParams
	GetSnapshotID(name string, opts ...OptionFunc) (string, int, error)
	GetSnapshotByName(name string, opts ...OptionFunc) (*Snapshot, int, error)
	GetSnapshotByID(id string, opts ...OptionFunc) (*Snapshot, int, error)
	ListVMSnapshot(p *ListVMSnapshotParams) (*ListVMSnapshotResponse, error)
	NewListVMSnapshotParams() *ListVMSnapshotParams
	GetVMSnapshotID(name string, opts ...OptionFunc) (string, int, error)
	RevertSnapshot(p *RevertSnapshotParams) (*RevertSnapshotResponse, error)
	NewRevertSnapshotParams(id string) *RevertSnapshotParams
	RevertToVMSnapshot(p *RevertToVMSnapshotParams) (*RevertToVMSnapshotResponse, error)
	NewRevertToVMSnapshotParams(vmsnapshotid string) *RevertToVMSnapshotParams
	UpdateSnapshotPolicy(p *UpdateSnapshotPolicyParams) (*UpdateSnapshotPolicyResponse, error)
	NewUpdateSnapshotPolicyParams() *UpdateSnapshotPolicyParams
}

type ArchiveSnapshotParams struct {
	p map[string]interface{}
}

func (p *ArchiveSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ArchiveSnapshotParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ArchiveSnapshotParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ArchiveSnapshotParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new ArchiveSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewArchiveSnapshotParams(id string) *ArchiveSnapshotParams {
	p := &ArchiveSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Archives (moves) a snapshot on primary storage to secondary storage
func (s *SnapshotService) ArchiveSnapshot(p *ArchiveSnapshotParams) (*ArchiveSnapshotResponse, error) {
	resp, err := s.cs.newRequest("archiveSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ArchiveSnapshotResponse
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

type ArchiveSnapshotResponse struct {
	Account         string            `json:"account"`
	Created         string            `json:"created"`
	Datastoreid     string            `json:"datastoreid"`
	Datastorename   string            `json:"datastorename"`
	Datastorestate  string            `json:"datastorestate"`
	Datastoretype   string            `json:"datastoretype"`
	Domain          string            `json:"domain"`
	Domainid        string            `json:"domainid"`
	Domainpath      string            `json:"domainpath"`
	Downloaddetails map[string]string `json:"downloaddetails"`
	Hasannotations  bool              `json:"hasannotations"`
	Id              string            `json:"id"`
	Intervaltype    string            `json:"intervaltype"`
	JobID           string            `json:"jobid"`
	Jobstatus       int               `json:"jobstatus"`
	Locationtype    string            `json:"locationtype"`
	Name            string            `json:"name"`
	Osdisplayname   string            `json:"osdisplayname"`
	Ostypeid        string            `json:"ostypeid"`
	Physicalsize    int64             `json:"physicalsize"`
	Project         string            `json:"project"`
	Projectid       string            `json:"projectid"`
	Revertable      bool              `json:"revertable"`
	Snapshottype    string            `json:"snapshottype"`
	State           string            `json:"state"`
	Status          string            `json:"status"`
	Tags            []Tags            `json:"tags"`
	Virtualsize     int64             `json:"virtualsize"`
	Volumeid        string            `json:"volumeid"`
	Volumename      string            `json:"volumename"`
	Volumestate     string            `json:"volumestate"`
	Volumetype      string            `json:"volumetype"`
	Zoneid          string            `json:"zoneid"`
	Zonename        string            `json:"zonename"`
}

func (r *ArchiveSnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias ArchiveSnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CopySnapshotParams struct {
	p map[string]interface{}
}

func (p *CopySnapshotParams) toURLValues() url.Values {
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

func (p *CopySnapshotParams) SetDestzoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destzoneid"] = v
}

func (p *CopySnapshotParams) ResetDestzoneid() {
	if p.p != nil && p.p["destzoneid"] != nil {
		delete(p.p, "destzoneid")
	}
}

func (p *CopySnapshotParams) GetDestzoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destzoneid"].(string)
	return value, ok
}

func (p *CopySnapshotParams) SetDestzoneids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destzoneids"] = v
}

func (p *CopySnapshotParams) ResetDestzoneids() {
	if p.p != nil && p.p["destzoneids"] != nil {
		delete(p.p, "destzoneids")
	}
}

func (p *CopySnapshotParams) GetDestzoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destzoneids"].([]string)
	return value, ok
}

func (p *CopySnapshotParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *CopySnapshotParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *CopySnapshotParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *CopySnapshotParams) SetSourcezoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcezoneid"] = v
}

func (p *CopySnapshotParams) ResetSourcezoneid() {
	if p.p != nil && p.p["sourcezoneid"] != nil {
		delete(p.p, "sourcezoneid")
	}
}

func (p *CopySnapshotParams) GetSourcezoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcezoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CopySnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCopySnapshotParams(id string) *CopySnapshotParams {
	p := &CopySnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Copies a snapshot from one zone to another.
func (s *SnapshotService) CopySnapshot(p *CopySnapshotParams) (*CopySnapshotResponse, error) {
	resp, err := s.cs.newRequest("copySnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CopySnapshotResponse
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

type CopySnapshotResponse struct {
	Account         string            `json:"account"`
	Created         string            `json:"created"`
	Datastoreid     string            `json:"datastoreid"`
	Datastorename   string            `json:"datastorename"`
	Datastorestate  string            `json:"datastorestate"`
	Datastoretype   string            `json:"datastoretype"`
	Domain          string            `json:"domain"`
	Domainid        string            `json:"domainid"`
	Domainpath      string            `json:"domainpath"`
	Downloaddetails map[string]string `json:"downloaddetails"`
	Hasannotations  bool              `json:"hasannotations"`
	Id              string            `json:"id"`
	Intervaltype    string            `json:"intervaltype"`
	JobID           string            `json:"jobid"`
	Jobstatus       int               `json:"jobstatus"`
	Locationtype    string            `json:"locationtype"`
	Name            string            `json:"name"`
	Osdisplayname   string            `json:"osdisplayname"`
	Ostypeid        string            `json:"ostypeid"`
	Physicalsize    int64             `json:"physicalsize"`
	Project         string            `json:"project"`
	Projectid       string            `json:"projectid"`
	Revertable      bool              `json:"revertable"`
	Snapshottype    string            `json:"snapshottype"`
	State           string            `json:"state"`
	Status          string            `json:"status"`
	Tags            []Tags            `json:"tags"`
	Virtualsize     int64             `json:"virtualsize"`
	Volumeid        string            `json:"volumeid"`
	Volumename      string            `json:"volumename"`
	Volumestate     string            `json:"volumestate"`
	Volumetype      string            `json:"volumetype"`
	Zoneid          string            `json:"zoneid"`
	Zonename        string            `json:"zonename"`
}

func (r *CopySnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias CopySnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CreateSnapshotParams struct {
	p map[string]interface{}
}

func (p *CreateSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["asyncbackup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("asyncbackup", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["locationtype"]; found {
		u.Set("locationtype", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	if v, found := p.p["quiescevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("quiescevm", vv)
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	if v, found := p.p["zoneids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneids", vv)
	}
	return u
}

func (p *CreateSnapshotParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateSnapshotParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateSnapshotParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateSnapshotParams) SetAsyncbackup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["asyncbackup"] = v
}

func (p *CreateSnapshotParams) ResetAsyncbackup() {
	if p.p != nil && p.p["asyncbackup"] != nil {
		delete(p.p, "asyncbackup")
	}
}

func (p *CreateSnapshotParams) GetAsyncbackup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["asyncbackup"].(bool)
	return value, ok
}

func (p *CreateSnapshotParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateSnapshotParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateSnapshotParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateSnapshotParams) SetLocationtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["locationtype"] = v
}

func (p *CreateSnapshotParams) ResetLocationtype() {
	if p.p != nil && p.p["locationtype"] != nil {
		delete(p.p, "locationtype")
	}
}

func (p *CreateSnapshotParams) GetLocationtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["locationtype"].(string)
	return value, ok
}

func (p *CreateSnapshotParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateSnapshotParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateSnapshotParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateSnapshotParams) SetPolicyid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyid"] = v
}

func (p *CreateSnapshotParams) ResetPolicyid() {
	if p.p != nil && p.p["policyid"] != nil {
		delete(p.p, "policyid")
	}
}

func (p *CreateSnapshotParams) GetPolicyid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyid"].(string)
	return value, ok
}

func (p *CreateSnapshotParams) SetQuiescevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiescevm"] = v
}

func (p *CreateSnapshotParams) ResetQuiescevm() {
	if p.p != nil && p.p["quiescevm"] != nil {
		delete(p.p, "quiescevm")
	}
}

func (p *CreateSnapshotParams) GetQuiescevm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quiescevm"].(bool)
	return value, ok
}

func (p *CreateSnapshotParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreateSnapshotParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *CreateSnapshotParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *CreateSnapshotParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *CreateSnapshotParams) ResetVolumeid() {
	if p.p != nil && p.p["volumeid"] != nil {
		delete(p.p, "volumeid")
	}
}

func (p *CreateSnapshotParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

func (p *CreateSnapshotParams) SetZoneids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneids"] = v
}

func (p *CreateSnapshotParams) ResetZoneids() {
	if p.p != nil && p.p["zoneids"] != nil {
		delete(p.p, "zoneids")
	}
}

func (p *CreateSnapshotParams) GetZoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneids"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateSnapshotParams(volumeid string) *CreateSnapshotParams {
	p := &CreateSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["volumeid"] = volumeid
	return p
}

// Creates an instant snapshot of a volume.
func (s *SnapshotService) CreateSnapshot(p *CreateSnapshotParams) (*CreateSnapshotResponse, error) {
	resp, err := s.cs.newRequest("createSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSnapshotResponse
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

type CreateSnapshotResponse struct {
	Account         string            `json:"account"`
	Created         string            `json:"created"`
	Datastoreid     string            `json:"datastoreid"`
	Datastorename   string            `json:"datastorename"`
	Datastorestate  string            `json:"datastorestate"`
	Datastoretype   string            `json:"datastoretype"`
	Domain          string            `json:"domain"`
	Domainid        string            `json:"domainid"`
	Domainpath      string            `json:"domainpath"`
	Downloaddetails map[string]string `json:"downloaddetails"`
	Hasannotations  bool              `json:"hasannotations"`
	Id              string            `json:"id"`
	Intervaltype    string            `json:"intervaltype"`
	JobID           string            `json:"jobid"`
	Jobstatus       int               `json:"jobstatus"`
	Locationtype    string            `json:"locationtype"`
	Name            string            `json:"name"`
	Osdisplayname   string            `json:"osdisplayname"`
	Ostypeid        string            `json:"ostypeid"`
	Physicalsize    int64             `json:"physicalsize"`
	Project         string            `json:"project"`
	Projectid       string            `json:"projectid"`
	Revertable      bool              `json:"revertable"`
	Snapshottype    string            `json:"snapshottype"`
	State           string            `json:"state"`
	Status          string            `json:"status"`
	Tags            []Tags            `json:"tags"`
	Virtualsize     int64             `json:"virtualsize"`
	Volumeid        string            `json:"volumeid"`
	Volumename      string            `json:"volumename"`
	Volumestate     string            `json:"volumestate"`
	Volumetype      string            `json:"volumetype"`
	Zoneid          string            `json:"zoneid"`
	Zonename        string            `json:"zonename"`
}

func (r *CreateSnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias CreateSnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CreateSnapshotFromVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *CreateSnapshotFromVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *CreateSnapshotFromVMSnapshotParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateSnapshotFromVMSnapshotParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateSnapshotFromVMSnapshotParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateSnapshotFromVMSnapshotParams) SetVmsnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotid"] = v
}

func (p *CreateSnapshotFromVMSnapshotParams) ResetVmsnapshotid() {
	if p.p != nil && p.p["vmsnapshotid"] != nil {
		delete(p.p, "vmsnapshotid")
	}
}

func (p *CreateSnapshotFromVMSnapshotParams) GetVmsnapshotid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmsnapshotid"].(string)
	return value, ok
}

func (p *CreateSnapshotFromVMSnapshotParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *CreateSnapshotFromVMSnapshotParams) ResetVolumeid() {
	if p.p != nil && p.p["volumeid"] != nil {
		delete(p.p, "volumeid")
	}
}

func (p *CreateSnapshotFromVMSnapshotParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSnapshotFromVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateSnapshotFromVMSnapshotParams(vmsnapshotid string, volumeid string) *CreateSnapshotFromVMSnapshotParams {
	p := &CreateSnapshotFromVMSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["vmsnapshotid"] = vmsnapshotid
	p.p["volumeid"] = volumeid
	return p
}

// Creates an instant snapshot of a volume from existing vm snapshot.
func (s *SnapshotService) CreateSnapshotFromVMSnapshot(p *CreateSnapshotFromVMSnapshotParams) (*CreateSnapshotFromVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("createSnapshotFromVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSnapshotFromVMSnapshotResponse
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

type CreateSnapshotFromVMSnapshotResponse struct {
	Account         string            `json:"account"`
	Created         string            `json:"created"`
	Datastoreid     string            `json:"datastoreid"`
	Datastorename   string            `json:"datastorename"`
	Datastorestate  string            `json:"datastorestate"`
	Datastoretype   string            `json:"datastoretype"`
	Domain          string            `json:"domain"`
	Domainid        string            `json:"domainid"`
	Domainpath      string            `json:"domainpath"`
	Downloaddetails map[string]string `json:"downloaddetails"`
	Hasannotations  bool              `json:"hasannotations"`
	Id              string            `json:"id"`
	Intervaltype    string            `json:"intervaltype"`
	JobID           string            `json:"jobid"`
	Jobstatus       int               `json:"jobstatus"`
	Locationtype    string            `json:"locationtype"`
	Name            string            `json:"name"`
	Osdisplayname   string            `json:"osdisplayname"`
	Ostypeid        string            `json:"ostypeid"`
	Physicalsize    int64             `json:"physicalsize"`
	Project         string            `json:"project"`
	Projectid       string            `json:"projectid"`
	Revertable      bool              `json:"revertable"`
	Snapshottype    string            `json:"snapshottype"`
	State           string            `json:"state"`
	Status          string            `json:"status"`
	Tags            []Tags            `json:"tags"`
	Virtualsize     int64             `json:"virtualsize"`
	Volumeid        string            `json:"volumeid"`
	Volumename      string            `json:"volumename"`
	Volumestate     string            `json:"volumestate"`
	Volumetype      string            `json:"volumetype"`
	Zoneid          string            `json:"zoneid"`
	Zonename        string            `json:"zonename"`
}

func (r *CreateSnapshotFromVMSnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias CreateSnapshotFromVMSnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CreateSnapshotPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateSnapshotPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["intervaltype"]; found {
		u.Set("intervaltype", v.(string))
	}
	if v, found := p.p["maxsnaps"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxsnaps", vv)
	}
	if v, found := p.p["schedule"]; found {
		u.Set("schedule", v.(string))
	}
	if v, found := p.p["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	if v, found := p.p["zoneids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneids", vv)
	}
	return u
}

func (p *CreateSnapshotPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateSnapshotPolicyParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateSnapshotPolicyParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateSnapshotPolicyParams) SetIntervaltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltype"] = v
}

func (p *CreateSnapshotPolicyParams) ResetIntervaltype() {
	if p.p != nil && p.p["intervaltype"] != nil {
		delete(p.p, "intervaltype")
	}
}

func (p *CreateSnapshotPolicyParams) GetIntervaltype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["intervaltype"].(string)
	return value, ok
}

func (p *CreateSnapshotPolicyParams) SetMaxsnaps(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxsnaps"] = v
}

func (p *CreateSnapshotPolicyParams) ResetMaxsnaps() {
	if p.p != nil && p.p["maxsnaps"] != nil {
		delete(p.p, "maxsnaps")
	}
}

func (p *CreateSnapshotPolicyParams) GetMaxsnaps() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxsnaps"].(int)
	return value, ok
}

func (p *CreateSnapshotPolicyParams) SetSchedule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["schedule"] = v
}

func (p *CreateSnapshotPolicyParams) ResetSchedule() {
	if p.p != nil && p.p["schedule"] != nil {
		delete(p.p, "schedule")
	}
}

func (p *CreateSnapshotPolicyParams) GetSchedule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["schedule"].(string)
	return value, ok
}

func (p *CreateSnapshotPolicyParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreateSnapshotPolicyParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *CreateSnapshotPolicyParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *CreateSnapshotPolicyParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *CreateSnapshotPolicyParams) ResetTimezone() {
	if p.p != nil && p.p["timezone"] != nil {
		delete(p.p, "timezone")
	}
}

func (p *CreateSnapshotPolicyParams) GetTimezone() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timezone"].(string)
	return value, ok
}

func (p *CreateSnapshotPolicyParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *CreateSnapshotPolicyParams) ResetVolumeid() {
	if p.p != nil && p.p["volumeid"] != nil {
		delete(p.p, "volumeid")
	}
}

func (p *CreateSnapshotPolicyParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

func (p *CreateSnapshotPolicyParams) SetZoneids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneids"] = v
}

func (p *CreateSnapshotPolicyParams) ResetZoneids() {
	if p.p != nil && p.p["zoneids"] != nil {
		delete(p.p, "zoneids")
	}
}

func (p *CreateSnapshotPolicyParams) GetZoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneids"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateSnapshotPolicyParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateSnapshotPolicyParams(intervaltype string, maxsnaps int, schedule string, timezone string, volumeid string) *CreateSnapshotPolicyParams {
	p := &CreateSnapshotPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["intervaltype"] = intervaltype
	p.p["maxsnaps"] = maxsnaps
	p.p["schedule"] = schedule
	p.p["timezone"] = timezone
	p.p["volumeid"] = volumeid
	return p
}

// Creates a snapshot policy for the account.
func (s *SnapshotService) CreateSnapshotPolicy(p *CreateSnapshotPolicyParams) (*CreateSnapshotPolicyResponse, error) {
	resp, err := s.cs.newRequest("createSnapshotPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSnapshotPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateSnapshotPolicyResponse struct {
	Fordisplay     bool          `json:"fordisplay"`
	Hasannotations bool          `json:"hasannotations"`
	Id             string        `json:"id"`
	Intervaltype   int           `json:"intervaltype"`
	JobID          string        `json:"jobid"`
	Jobstatus      int           `json:"jobstatus"`
	Maxsnaps       int           `json:"maxsnaps"`
	Schedule       string        `json:"schedule"`
	Tags           []Tags        `json:"tags"`
	Timezone       string        `json:"timezone"`
	Volumeid       string        `json:"volumeid"`
	Zone           []interface{} `json:"zone"`
}

type CreateVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *CreateVMSnapshotParams) toURLValues() url.Values {
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
	if v, found := p.p["snapshotmemory"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("snapshotmemory", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *CreateVMSnapshotParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateVMSnapshotParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateVMSnapshotParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateVMSnapshotParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateVMSnapshotParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateVMSnapshotParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateVMSnapshotParams) SetQuiescevm(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiescevm"] = v
}

func (p *CreateVMSnapshotParams) ResetQuiescevm() {
	if p.p != nil && p.p["quiescevm"] != nil {
		delete(p.p, "quiescevm")
	}
}

func (p *CreateVMSnapshotParams) GetQuiescevm() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quiescevm"].(bool)
	return value, ok
}

func (p *CreateVMSnapshotParams) SetSnapshotmemory(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshotmemory"] = v
}

func (p *CreateVMSnapshotParams) ResetSnapshotmemory() {
	if p.p != nil && p.p["snapshotmemory"] != nil {
		delete(p.p, "snapshotmemory")
	}
}

func (p *CreateVMSnapshotParams) GetSnapshotmemory() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["snapshotmemory"].(bool)
	return value, ok
}

func (p *CreateVMSnapshotParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *CreateVMSnapshotParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *CreateVMSnapshotParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateVMSnapshotParams(virtualmachineid string) *CreateVMSnapshotParams {
	p := &CreateVMSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Creates snapshot for a vm.
func (s *SnapshotService) CreateVMSnapshot(p *CreateVMSnapshotParams) (*CreateVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("createVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVMSnapshotResponse
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

type CreateVMSnapshotResponse struct {
	Account            string `json:"account"`
	Created            string `json:"created"`
	Current            bool   `json:"current"`
	Description        string `json:"description"`
	Displayname        string `json:"displayname"`
	Domain             string `json:"domain"`
	Domainid           string `json:"domainid"`
	Domainpath         string `json:"domainpath"`
	Hasannotations     bool   `json:"hasannotations"`
	Hypervisor         string `json:"hypervisor"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Name               string `json:"name"`
	Parent             string `json:"parent"`
	ParentName         string `json:"parentName"`
	Project            string `json:"project"`
	Projectid          string `json:"projectid"`
	State              string `json:"state"`
	Tags               []Tags `json:"tags"`
	Type               string `json:"type"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
	Zoneid             string `json:"zoneid"`
	Zonename           string `json:"zonename"`
}

type DeleteSnapshotParams struct {
	p map[string]interface{}
}

func (p *DeleteSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteSnapshotParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteSnapshotParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteSnapshotParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteSnapshotParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteSnapshotParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteSnapshotParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteSnapshotParams(id string) *DeleteSnapshotParams {
	p := &DeleteSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a snapshot of a disk volume.
func (s *SnapshotService) DeleteSnapshot(p *DeleteSnapshotParams) (*DeleteSnapshotResponse, error) {
	resp, err := s.cs.newRequest("deleteSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSnapshotResponse
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

type DeleteSnapshotResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteSnapshotPoliciesParams struct {
	p map[string]interface{}
}

func (p *DeleteSnapshotPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	return u
}

func (p *DeleteSnapshotPoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteSnapshotPoliciesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteSnapshotPoliciesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteSnapshotPoliciesParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *DeleteSnapshotPoliciesParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *DeleteSnapshotPoliciesParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

// You should always use this function to get a new DeleteSnapshotPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteSnapshotPoliciesParams() *DeleteSnapshotPoliciesParams {
	p := &DeleteSnapshotPoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Deletes snapshot policies for the account.
func (s *SnapshotService) DeleteSnapshotPolicies(p *DeleteSnapshotPoliciesParams) (*DeleteSnapshotPoliciesResponse, error) {
	resp, err := s.cs.newRequest("deleteSnapshotPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSnapshotPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSnapshotPoliciesResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSnapshotPoliciesResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSnapshotPoliciesResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *DeleteVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	return u
}

func (p *DeleteVMSnapshotParams) SetVmsnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotid"] = v
}

func (p *DeleteVMSnapshotParams) ResetVmsnapshotid() {
	if p.p != nil && p.p["vmsnapshotid"] != nil {
		delete(p.p, "vmsnapshotid")
	}
}

func (p *DeleteVMSnapshotParams) GetVmsnapshotid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmsnapshotid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteVMSnapshotParams(vmsnapshotid string) *DeleteVMSnapshotParams {
	p := &DeleteVMSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["vmsnapshotid"] = vmsnapshotid
	return p
}

// Deletes a vmsnapshot.
func (s *SnapshotService) DeleteVMSnapshot(p *DeleteVMSnapshotParams) (*DeleteVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("deleteVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVMSnapshotResponse
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

type DeleteVMSnapshotResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ExtractSnapshotParams struct {
	p map[string]interface{}
}

func (p *ExtractSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ExtractSnapshotParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ExtractSnapshotParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ExtractSnapshotParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ExtractSnapshotParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ExtractSnapshotParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ExtractSnapshotParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ExtractSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewExtractSnapshotParams(id string, zoneid string) *ExtractSnapshotParams {
	p := &ExtractSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["zoneid"] = zoneid
	return p
}

// Returns a download URL for extracting a snapshot. It must be in the Backed Up state.
func (s *SnapshotService) ExtractSnapshot(p *ExtractSnapshotParams) (*ExtractSnapshotResponse, error) {
	resp, err := s.cs.newRequest("extractSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractSnapshotResponse
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

type ExtractSnapshotResponse struct {
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

type ListSnapshotPoliciesParams struct {
	p map[string]interface{}
}

func (p *ListSnapshotPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
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
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (p *ListSnapshotPoliciesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListSnapshotPoliciesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListSnapshotPoliciesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListSnapshotPoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSnapshotPoliciesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListSnapshotPoliciesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListSnapshotPoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSnapshotPoliciesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSnapshotPoliciesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSnapshotPoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSnapshotPoliciesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSnapshotPoliciesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSnapshotPoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSnapshotPoliciesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSnapshotPoliciesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListSnapshotPoliciesParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *ListSnapshotPoliciesParams) ResetVolumeid() {
	if p.p != nil && p.p["volumeid"] != nil {
		delete(p.p, "volumeid")
	}
}

func (p *ListSnapshotPoliciesParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSnapshotPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListSnapshotPoliciesParams() *ListSnapshotPoliciesParams {
	p := &ListSnapshotPoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotPolicyByID(id string, opts ...OptionFunc) (*SnapshotPolicy, int, error) {
	p := &ListSnapshotPoliciesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSnapshotPolicies(p)
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
		return l.SnapshotPolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SnapshotPolicy UUID: %s!", id)
}

// Lists snapshot policies.
func (s *SnapshotService) ListSnapshotPolicies(p *ListSnapshotPoliciesParams) (*ListSnapshotPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listSnapshotPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSnapshotPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSnapshotPoliciesResponse struct {
	Count            int               `json:"count"`
	SnapshotPolicies []*SnapshotPolicy `json:"snapshotpolicy"`
}

type SnapshotPolicy struct {
	Fordisplay     bool          `json:"fordisplay"`
	Hasannotations bool          `json:"hasannotations"`
	Id             string        `json:"id"`
	Intervaltype   int           `json:"intervaltype"`
	JobID          string        `json:"jobid"`
	Jobstatus      int           `json:"jobstatus"`
	Maxsnaps       int           `json:"maxsnaps"`
	Schedule       string        `json:"schedule"`
	Tags           []Tags        `json:"tags"`
	Timezone       string        `json:"timezone"`
	Volumeid       string        `json:"volumeid"`
	Zone           []interface{} `json:"zone"`
}

type ListSnapshotsParams struct {
	p map[string]interface{}
}

func (p *ListSnapshotsParams) toURLValues() url.Values {
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
	if v, found := p.p["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := p.p["imagestoreid"]; found {
		u.Set("imagestoreid", v.(string))
	}
	if v, found := p.p["intervaltype"]; found {
		u.Set("intervaltype", v.(string))
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
	if v, found := p.p["locationtype"]; found {
		u.Set("locationtype", v.(string))
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
	if v, found := p.p["showunique"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showunique", vv)
	}
	if v, found := p.p["snapshottype"]; found {
		u.Set("snapshottype", v.(string))
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
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListSnapshotsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListSnapshotsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListSnapshotsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListSnapshotsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListSnapshotsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSnapshotsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListSnapshotsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetIds(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ids"] = v
}

func (p *ListSnapshotsParams) ResetIds() {
	if p.p != nil && p.p["ids"] != nil {
		delete(p.p, "ids")
	}
}

func (p *ListSnapshotsParams) GetIds() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ids"].([]string)
	return value, ok
}

func (p *ListSnapshotsParams) SetImagestoreid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["imagestoreid"] = v
}

func (p *ListSnapshotsParams) ResetImagestoreid() {
	if p.p != nil && p.p["imagestoreid"] != nil {
		delete(p.p, "imagestoreid")
	}
}

func (p *ListSnapshotsParams) GetImagestoreid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["imagestoreid"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetIntervaltype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltype"] = v
}

func (p *ListSnapshotsParams) ResetIntervaltype() {
	if p.p != nil && p.p["intervaltype"] != nil {
		delete(p.p, "intervaltype")
	}
}

func (p *ListSnapshotsParams) GetIntervaltype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["intervaltype"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListSnapshotsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListSnapshotsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListSnapshotsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSnapshotsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListSnapshotsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListSnapshotsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListSnapshotsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListSnapshotsParams) SetLocationtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["locationtype"] = v
}

func (p *ListSnapshotsParams) ResetLocationtype() {
	if p.p != nil && p.p["locationtype"] != nil {
		delete(p.p, "locationtype")
	}
}

func (p *ListSnapshotsParams) GetLocationtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["locationtype"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListSnapshotsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListSnapshotsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSnapshotsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListSnapshotsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListSnapshotsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSnapshotsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListSnapshotsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListSnapshotsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListSnapshotsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListSnapshotsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetShowunique(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showunique"] = v
}

func (p *ListSnapshotsParams) ResetShowunique() {
	if p.p != nil && p.p["showunique"] != nil {
		delete(p.p, "showunique")
	}
}

func (p *ListSnapshotsParams) GetShowunique() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["showunique"].(bool)
	return value, ok
}

func (p *ListSnapshotsParams) SetSnapshottype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshottype"] = v
}

func (p *ListSnapshotsParams) ResetSnapshottype() {
	if p.p != nil && p.p["snapshottype"] != nil {
		delete(p.p, "snapshottype")
	}
}

func (p *ListSnapshotsParams) GetSnapshottype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["snapshottype"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListSnapshotsParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *ListSnapshotsParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListSnapshotsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListSnapshotsParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListSnapshotsParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *ListSnapshotsParams) ResetVolumeid() {
	if p.p != nil && p.p["volumeid"] != nil {
		delete(p.p, "volumeid")
	}
}

func (p *ListSnapshotsParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

func (p *ListSnapshotsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListSnapshotsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListSnapshotsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSnapshotsParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListSnapshotsParams() *ListSnapshotsParams {
	p := &ListSnapshotsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListSnapshotsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSnapshots(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Snapshots[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Snapshots {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotByName(name string, opts ...OptionFunc) (*Snapshot, int, error) {
	id, count, err := s.GetSnapshotID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSnapshotByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotByID(id string, opts ...OptionFunc) (*Snapshot, int, error) {
	p := &ListSnapshotsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSnapshots(p)
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
		return l.Snapshots[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Snapshot UUID: %s!", id)
}

// Lists all available snapshots for the account.
func (s *SnapshotService) ListSnapshots(p *ListSnapshotsParams) (*ListSnapshotsResponse, error) {
	resp, err := s.cs.newRequest("listSnapshots", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSnapshotsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSnapshotsResponse struct {
	Count     int         `json:"count"`
	Snapshots []*Snapshot `json:"snapshot"`
}

type Snapshot struct {
	Account         string            `json:"account"`
	Created         string            `json:"created"`
	Datastoreid     string            `json:"datastoreid"`
	Datastorename   string            `json:"datastorename"`
	Datastorestate  string            `json:"datastorestate"`
	Datastoretype   string            `json:"datastoretype"`
	Domain          string            `json:"domain"`
	Domainid        string            `json:"domainid"`
	Domainpath      string            `json:"domainpath"`
	Downloaddetails map[string]string `json:"downloaddetails"`
	Hasannotations  bool              `json:"hasannotations"`
	Id              string            `json:"id"`
	Intervaltype    string            `json:"intervaltype"`
	JobID           string            `json:"jobid"`
	Jobstatus       int               `json:"jobstatus"`
	Locationtype    string            `json:"locationtype"`
	Name            string            `json:"name"`
	Osdisplayname   string            `json:"osdisplayname"`
	Ostypeid        string            `json:"ostypeid"`
	Physicalsize    int64             `json:"physicalsize"`
	Project         string            `json:"project"`
	Projectid       string            `json:"projectid"`
	Revertable      bool              `json:"revertable"`
	Snapshottype    string            `json:"snapshottype"`
	State           string            `json:"state"`
	Status          string            `json:"status"`
	Tags            []Tags            `json:"tags"`
	Virtualsize     int64             `json:"virtualsize"`
	Volumeid        string            `json:"volumeid"`
	Volumename      string            `json:"volumename"`
	Volumestate     string            `json:"volumestate"`
	Volumetype      string            `json:"volumetype"`
	Zoneid          string            `json:"zoneid"`
	Zonename        string            `json:"zonename"`
}

func (r *Snapshot) UnmarshalJSON(b []byte) error {
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

	type alias Snapshot
	return json.Unmarshal(b, (*alias)(r))
}

type ListVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *ListVMSnapshotParams) toURLValues() url.Values {
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
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
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	if v, found := p.p["vmsnapshotids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("vmsnapshotids", vv)
	}
	return u
}

func (p *ListVMSnapshotParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListVMSnapshotParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListVMSnapshotParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListVMSnapshotParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListVMSnapshotParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListVMSnapshotParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListVMSnapshotParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListVMSnapshotParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVMSnapshotParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVMSnapshotParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListVMSnapshotParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListVMSnapshotParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListVMSnapshotParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListVMSnapshotParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListVMSnapshotParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVMSnapshotParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVMSnapshotParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVMSnapshotParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVMSnapshotParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVMSnapshotParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVMSnapshotParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListVMSnapshotParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListVMSnapshotParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListVMSnapshotParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListVMSnapshotParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListVMSnapshotParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *ListVMSnapshotParams) GetTags() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(map[string]string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListVMSnapshotParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *ListVMSnapshotParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetVmsnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotid"] = v
}

func (p *ListVMSnapshotParams) ResetVmsnapshotid() {
	if p.p != nil && p.p["vmsnapshotid"] != nil {
		delete(p.p, "vmsnapshotid")
	}
}

func (p *ListVMSnapshotParams) GetVmsnapshotid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmsnapshotid"].(string)
	return value, ok
}

func (p *ListVMSnapshotParams) SetVmsnapshotids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotids"] = v
}

func (p *ListVMSnapshotParams) ResetVmsnapshotids() {
	if p.p != nil && p.p["vmsnapshotids"] != nil {
		delete(p.p, "vmsnapshotids")
	}
}

func (p *ListVMSnapshotParams) GetVmsnapshotids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmsnapshotids"].([]string)
	return value, ok
}

// You should always use this function to get a new ListVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListVMSnapshotParams() *ListVMSnapshotParams {
	p := &ListVMSnapshotParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetVMSnapshotID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVMSnapshotParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVMSnapshot(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VMSnapshot[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VMSnapshot {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// List virtual machine snapshot by conditions
func (s *SnapshotService) ListVMSnapshot(p *ListVMSnapshotParams) (*ListVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("listVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVMSnapshotResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVMSnapshotResponse struct {
	Count      int           `json:"count"`
	VMSnapshot []*VMSnapshot `json:"vmsnapshot"`
}

type VMSnapshot struct {
	Account            string `json:"account"`
	Created            string `json:"created"`
	Current            bool   `json:"current"`
	Description        string `json:"description"`
	Displayname        string `json:"displayname"`
	Domain             string `json:"domain"`
	Domainid           string `json:"domainid"`
	Domainpath         string `json:"domainpath"`
	Hasannotations     bool   `json:"hasannotations"`
	Hypervisor         string `json:"hypervisor"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Name               string `json:"name"`
	Parent             string `json:"parent"`
	ParentName         string `json:"parentName"`
	Project            string `json:"project"`
	Projectid          string `json:"projectid"`
	State              string `json:"state"`
	Tags               []Tags `json:"tags"`
	Type               string `json:"type"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
	Zoneid             string `json:"zoneid"`
	Zonename           string `json:"zonename"`
}

type RevertSnapshotParams struct {
	p map[string]interface{}
}

func (p *RevertSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RevertSnapshotParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *RevertSnapshotParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *RevertSnapshotParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new RevertSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewRevertSnapshotParams(id string) *RevertSnapshotParams {
	p := &RevertSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is supposed to revert a volume snapshot. This command is only supported with KVM so far
func (s *SnapshotService) RevertSnapshot(p *RevertSnapshotParams) (*RevertSnapshotResponse, error) {
	resp, err := s.cs.newRequest("revertSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevertSnapshotResponse
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

type RevertSnapshotResponse struct {
	Account         string            `json:"account"`
	Created         string            `json:"created"`
	Datastoreid     string            `json:"datastoreid"`
	Datastorename   string            `json:"datastorename"`
	Datastorestate  string            `json:"datastorestate"`
	Datastoretype   string            `json:"datastoretype"`
	Domain          string            `json:"domain"`
	Domainid        string            `json:"domainid"`
	Domainpath      string            `json:"domainpath"`
	Downloaddetails map[string]string `json:"downloaddetails"`
	Hasannotations  bool              `json:"hasannotations"`
	Id              string            `json:"id"`
	Intervaltype    string            `json:"intervaltype"`
	JobID           string            `json:"jobid"`
	Jobstatus       int               `json:"jobstatus"`
	Locationtype    string            `json:"locationtype"`
	Name            string            `json:"name"`
	Osdisplayname   string            `json:"osdisplayname"`
	Ostypeid        string            `json:"ostypeid"`
	Physicalsize    int64             `json:"physicalsize"`
	Project         string            `json:"project"`
	Projectid       string            `json:"projectid"`
	Revertable      bool              `json:"revertable"`
	Snapshottype    string            `json:"snapshottype"`
	State           string            `json:"state"`
	Status          string            `json:"status"`
	Tags            []Tags            `json:"tags"`
	Virtualsize     int64             `json:"virtualsize"`
	Volumeid        string            `json:"volumeid"`
	Volumename      string            `json:"volumename"`
	Volumestate     string            `json:"volumestate"`
	Volumetype      string            `json:"volumetype"`
	Zoneid          string            `json:"zoneid"`
	Zonename        string            `json:"zonename"`
}

func (r *RevertSnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias RevertSnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RevertToVMSnapshotParams struct {
	p map[string]interface{}
}

func (p *RevertToVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	return u
}

func (p *RevertToVMSnapshotParams) SetVmsnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmsnapshotid"] = v
}

func (p *RevertToVMSnapshotParams) ResetVmsnapshotid() {
	if p.p != nil && p.p["vmsnapshotid"] != nil {
		delete(p.p, "vmsnapshotid")
	}
}

func (p *RevertToVMSnapshotParams) GetVmsnapshotid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmsnapshotid"].(string)
	return value, ok
}

// You should always use this function to get a new RevertToVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewRevertToVMSnapshotParams(vmsnapshotid string) *RevertToVMSnapshotParams {
	p := &RevertToVMSnapshotParams{}
	p.p = make(map[string]interface{})
	p.p["vmsnapshotid"] = vmsnapshotid
	return p
}

// Revert VM from a vmsnapshot.
func (s *SnapshotService) RevertToVMSnapshot(p *RevertToVMSnapshotParams) (*RevertToVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("revertToVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevertToVMSnapshotResponse
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

type RevertToVMSnapshotResponse struct {
	Account               string                                    `json:"account"`
	Affinitygroup         []RevertToVMSnapshotResponseAffinitygroup `json:"affinitygroup"`
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
	Securitygroup         []RevertToVMSnapshotResponseSecuritygroup `json:"securitygroup"`
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
	Vmtype                string                                    `json:"vmtype"`
	Vnfdetails            map[string]string                         `json:"vnfdetails"`
	Vnfnics               []string                                  `json:"vnfnics"`
	Zoneid                string                                    `json:"zoneid"`
	Zonename              string                                    `json:"zonename"`
}

type RevertToVMSnapshotResponseSecuritygroup struct {
	Account             string                                        `json:"account"`
	Description         string                                        `json:"description"`
	Domain              string                                        `json:"domain"`
	Domainid            string                                        `json:"domainid"`
	Domainpath          string                                        `json:"domainpath"`
	Egressrule          []RevertToVMSnapshotResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                        `json:"id"`
	Ingressrule         []RevertToVMSnapshotResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                        `json:"name"`
	Project             string                                        `json:"project"`
	Projectid           string                                        `json:"projectid"`
	Tags                []Tags                                        `json:"tags"`
	Virtualmachinecount int                                           `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                 `json:"virtualmachineids"`
}

type RevertToVMSnapshotResponseSecuritygroupRule struct {
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

type RevertToVMSnapshotResponseAffinitygroup struct {
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

func (r *RevertToVMSnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias RevertToVMSnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateSnapshotPolicyParams struct {
	p map[string]interface{}
}

func (p *UpdateSnapshotPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateSnapshotPolicyParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateSnapshotPolicyParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateSnapshotPolicyParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateSnapshotPolicyParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateSnapshotPolicyParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateSnapshotPolicyParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateSnapshotPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateSnapshotPolicyParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateSnapshotPolicyParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateSnapshotPolicyParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewUpdateSnapshotPolicyParams() *UpdateSnapshotPolicyParams {
	p := &UpdateSnapshotPolicyParams{}
	p.p = make(map[string]interface{})
	return p
}

// Updates the snapshot policy.
func (s *SnapshotService) UpdateSnapshotPolicy(p *UpdateSnapshotPolicyParams) (*UpdateSnapshotPolicyResponse, error) {
	resp, err := s.cs.newRequest("updateSnapshotPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateSnapshotPolicyResponse
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

type UpdateSnapshotPolicyResponse struct {
	Fordisplay     bool          `json:"fordisplay"`
	Hasannotations bool          `json:"hasannotations"`
	Id             string        `json:"id"`
	Intervaltype   int           `json:"intervaltype"`
	JobID          string        `json:"jobid"`
	Jobstatus      int           `json:"jobstatus"`
	Maxsnaps       int           `json:"maxsnaps"`
	Schedule       string        `json:"schedule"`
	Tags           []Tags        `json:"tags"`
	Timezone       string        `json:"timezone"`
	Volumeid       string        `json:"volumeid"`
	Zone           []interface{} `json:"zone"`
}
