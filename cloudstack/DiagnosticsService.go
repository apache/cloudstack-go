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
	"strings"
)

type DiagnosticsServiceIface interface {
	GetDiagnosticsData(p *GetDiagnosticsDataParams) (*GetDiagnosticsDataResponse, error)
	NewGetDiagnosticsDataParams(targetid string) *GetDiagnosticsDataParams
}

type GetDiagnosticsDataParams struct {
	p map[string]interface{}
}

func (p *GetDiagnosticsDataParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["files"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("files", vv)
	}
	if v, found := p.p["targetid"]; found {
		u.Set("targetid", v.(string))
	}
	return u
}

func (p *GetDiagnosticsDataParams) SetFiles(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["files"] = v
}

func (p *GetDiagnosticsDataParams) ResetFiles() {
	if p.p != nil && p.p["files"] != nil {
		delete(p.p, "files")
	}
}

func (p *GetDiagnosticsDataParams) GetFiles() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["files"].([]string)
	return value, ok
}

func (p *GetDiagnosticsDataParams) SetTargetid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["targetid"] = v
}

func (p *GetDiagnosticsDataParams) ResetTargetid() {
	if p.p != nil && p.p["targetid"] != nil {
		delete(p.p, "targetid")
	}
}

func (p *GetDiagnosticsDataParams) GetTargetid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["targetid"].(string)
	return value, ok
}

// You should always use this function to get a new GetDiagnosticsDataParams instance,
// as then you are sure you have configured all required params
func (s *DiagnosticsService) NewGetDiagnosticsDataParams(targetid string) *GetDiagnosticsDataParams {
	p := &GetDiagnosticsDataParams{}
	p.p = make(map[string]interface{})
	p.p["targetid"] = targetid
	return p
}

// Get diagnostics and files from system VMs
func (s *DiagnosticsService) GetDiagnosticsData(p *GetDiagnosticsDataParams) (*GetDiagnosticsDataResponse, error) {
	resp, err := s.cs.newRequest("getDiagnosticsData", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetDiagnosticsDataResponse
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

type GetDiagnosticsDataResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Url       string `json:"url"`
}
