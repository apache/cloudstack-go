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
)

type ResourceServiceIface interface {
	PurgeExpungedResources(p *PurgeExpungedResourcesParams) (*PurgeExpungedResourcesResponse, error)
	NewPurgeExpungedResourcesParams() *PurgeExpungedResourcesParams
}

type PurgeExpungedResourcesParams struct {
	p map[string]interface{}
}

func (p *PurgeExpungedResourcesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["batchsize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("batchsize", vv)
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *PurgeExpungedResourcesParams) SetBatchsize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["batchsize"] = v
}

func (p *PurgeExpungedResourcesParams) ResetBatchsize() {
	if p.p != nil && p.p["batchsize"] != nil {
		delete(p.p, "batchsize")
	}
}

func (p *PurgeExpungedResourcesParams) GetBatchsize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["batchsize"].(int64)
	return value, ok
}

func (p *PurgeExpungedResourcesParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *PurgeExpungedResourcesParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *PurgeExpungedResourcesParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *PurgeExpungedResourcesParams) SetResourcetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["resourcetype"] = v
}

func (p *PurgeExpungedResourcesParams) ResetResourcetype() {
	if p.p != nil && p.p["resourcetype"] != nil {
		delete(p.p, "resourcetype")
	}
}

func (p *PurgeExpungedResourcesParams) GetResourcetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["resourcetype"].(string)
	return value, ok
}

func (p *PurgeExpungedResourcesParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *PurgeExpungedResourcesParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *PurgeExpungedResourcesParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new PurgeExpungedResourcesParams instance,
// as then you are sure you have configured all required params
func (s *ResourceService) NewPurgeExpungedResourcesParams() *PurgeExpungedResourcesParams {
	p := &PurgeExpungedResourcesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Purge expunged resources
func (s *ResourceService) PurgeExpungedResources(p *PurgeExpungedResourcesParams) (*PurgeExpungedResourcesResponse, error) {
	resp, err := s.cs.newRequest("purgeExpungedResources", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PurgeExpungedResourcesResponse
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

type PurgeExpungedResourcesResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
