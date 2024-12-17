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
)

type SolidFireServiceIface interface {
	GetSolidFireAccountId(p *GetSolidFireAccountIdParams) (*GetSolidFireAccountIdResponse, error)
	NewGetSolidFireAccountIdParams(accountid string, storageid string) *GetSolidFireAccountIdParams
	GetSolidFireVolumeAccessGroupIds(p *GetSolidFireVolumeAccessGroupIdsParams) (*GetSolidFireVolumeAccessGroupIdsResponse, error)
	NewGetSolidFireVolumeAccessGroupIdsParams(clusterid string, storageid string) *GetSolidFireVolumeAccessGroupIdsParams
	GetSolidFireVolumeSize(p *GetSolidFireVolumeSizeParams) (*GetSolidFireVolumeSizeResponse, error)
	NewGetSolidFireVolumeSizeParams(volumeid string) *GetSolidFireVolumeSizeParams
}

type GetSolidFireAccountIdParams struct {
	p map[string]interface{}
}

func (p *GetSolidFireAccountIdParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	return u
}

func (p *GetSolidFireAccountIdParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *GetSolidFireAccountIdParams) ResetAccountid() {
	if p.p != nil && p.p["accountid"] != nil {
		delete(p.p, "accountid")
	}
}

func (p *GetSolidFireAccountIdParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *GetSolidFireAccountIdParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *GetSolidFireAccountIdParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *GetSolidFireAccountIdParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

// You should always use this function to get a new GetSolidFireAccountIdParams instance,
// as then you are sure you have configured all required params
func (s *SolidFireService) NewGetSolidFireAccountIdParams(accountid string, storageid string) *GetSolidFireAccountIdParams {
	p := &GetSolidFireAccountIdParams{}
	p.p = make(map[string]interface{})
	p.p["accountid"] = accountid
	p.p["storageid"] = storageid
	return p
}

// Get SolidFire Account ID
func (s *SolidFireService) GetSolidFireAccountId(p *GetSolidFireAccountIdParams) (*GetSolidFireAccountIdResponse, error) {
	resp, err := s.cs.newRequest("getSolidFireAccountId", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetSolidFireAccountIdResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetSolidFireAccountIdResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	SolidFireAccountId int64  `json:"solidFireAccountId"`
}

type GetSolidFireVolumeAccessGroupIdsParams struct {
	p map[string]interface{}
}

func (p *GetSolidFireVolumeAccessGroupIdsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	return u
}

func (p *GetSolidFireVolumeAccessGroupIdsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *GetSolidFireVolumeAccessGroupIdsParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *GetSolidFireVolumeAccessGroupIdsParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *GetSolidFireVolumeAccessGroupIdsParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *GetSolidFireVolumeAccessGroupIdsParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *GetSolidFireVolumeAccessGroupIdsParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

// You should always use this function to get a new GetSolidFireVolumeAccessGroupIdsParams instance,
// as then you are sure you have configured all required params
func (s *SolidFireService) NewGetSolidFireVolumeAccessGroupIdsParams(clusterid string, storageid string) *GetSolidFireVolumeAccessGroupIdsParams {
	p := &GetSolidFireVolumeAccessGroupIdsParams{}
	p.p = make(map[string]interface{})
	p.p["clusterid"] = clusterid
	p.p["storageid"] = storageid
	return p
}

// Get the SF Volume Access Group IDs
func (s *SolidFireService) GetSolidFireVolumeAccessGroupIds(p *GetSolidFireVolumeAccessGroupIdsParams) (*GetSolidFireVolumeAccessGroupIdsResponse, error) {
	resp, err := s.cs.newRequest("getSolidFireVolumeAccessGroupIds", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetSolidFireVolumeAccessGroupIdsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetSolidFireVolumeAccessGroupIdsResponse struct {
	JobID                         string  `json:"jobid"`
	Jobstatus                     int     `json:"jobstatus"`
	SolidFireVolumeAccessGroupIds []int64 `json:"solidFireVolumeAccessGroupIds"`
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

func (p *GetSolidFireVolumeSizeParams) ResetVolumeid() {
	if p.p != nil && p.p["volumeid"] != nil {
		delete(p.p, "volumeid")
	}
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
func (s *SolidFireService) NewGetSolidFireVolumeSizeParams(volumeid string) *GetSolidFireVolumeSizeParams {
	p := &GetSolidFireVolumeSizeParams{}
	p.p = make(map[string]interface{})
	p.p["volumeid"] = volumeid
	return p
}

// Get the SF volume size including Hypervisor Snapshot Reserve
func (s *SolidFireService) GetSolidFireVolumeSize(p *GetSolidFireVolumeSizeParams) (*GetSolidFireVolumeSizeResponse, error) {
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
