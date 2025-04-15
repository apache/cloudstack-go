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
	"strings"
)

type RollingMaintenanceServiceIface interface {
	StartRollingMaintenance(p *StartRollingMaintenanceParams) (*StartRollingMaintenanceResponse, error)
	NewStartRollingMaintenanceParams() *StartRollingMaintenanceParams
}

type StartRollingMaintenanceParams struct {
	p map[string]interface{}
}

func (p *StartRollingMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("clusterids", vv)
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["hostids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("hostids", vv)
	}
	if v, found := p.p["payload"]; found {
		u.Set("payload", v.(string))
	}
	if v, found := p.p["podids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("podids", vv)
	}
	if v, found := p.p["timeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("timeout", vv)
	}
	if v, found := p.p["zoneids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneids", vv)
	}
	return u
}

func (p *StartRollingMaintenanceParams) SetClusterids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterids"] = v
}

func (p *StartRollingMaintenanceParams) ResetClusterids() {
	if p.p != nil && p.p["clusterids"] != nil {
		delete(p.p, "clusterids")
	}
}

func (p *StartRollingMaintenanceParams) GetClusterids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterids"].([]string)
	return value, ok
}

func (p *StartRollingMaintenanceParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *StartRollingMaintenanceParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *StartRollingMaintenanceParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *StartRollingMaintenanceParams) SetHostids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostids"] = v
}

func (p *StartRollingMaintenanceParams) ResetHostids() {
	if p.p != nil && p.p["hostids"] != nil {
		delete(p.p, "hostids")
	}
}

func (p *StartRollingMaintenanceParams) GetHostids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostids"].([]string)
	return value, ok
}

func (p *StartRollingMaintenanceParams) SetPayload(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["payload"] = v
}

func (p *StartRollingMaintenanceParams) ResetPayload() {
	if p.p != nil && p.p["payload"] != nil {
		delete(p.p, "payload")
	}
}

func (p *StartRollingMaintenanceParams) GetPayload() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["payload"].(string)
	return value, ok
}

func (p *StartRollingMaintenanceParams) SetPodids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podids"] = v
}

func (p *StartRollingMaintenanceParams) ResetPodids() {
	if p.p != nil && p.p["podids"] != nil {
		delete(p.p, "podids")
	}
}

func (p *StartRollingMaintenanceParams) GetPodids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podids"].([]string)
	return value, ok
}

func (p *StartRollingMaintenanceParams) SetTimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timeout"] = v
}

func (p *StartRollingMaintenanceParams) ResetTimeout() {
	if p.p != nil && p.p["timeout"] != nil {
		delete(p.p, "timeout")
	}
}

func (p *StartRollingMaintenanceParams) GetTimeout() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timeout"].(int)
	return value, ok
}

func (p *StartRollingMaintenanceParams) SetZoneids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneids"] = v
}

func (p *StartRollingMaintenanceParams) ResetZoneids() {
	if p.p != nil && p.p["zoneids"] != nil {
		delete(p.p, "zoneids")
	}
}

func (p *StartRollingMaintenanceParams) GetZoneids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneids"].([]string)
	return value, ok
}

// You should always use this function to get a new StartRollingMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *RollingMaintenanceService) NewStartRollingMaintenanceParams() *StartRollingMaintenanceParams {
	p := &StartRollingMaintenanceParams{}
	p.p = make(map[string]interface{})
	return p
}

// Start rolling maintenance
func (s *RollingMaintenanceService) StartRollingMaintenance(p *StartRollingMaintenanceParams) (*StartRollingMaintenanceResponse, error) {
	resp, err := s.cs.newPostRequest("startRollingMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartRollingMaintenanceResponse
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

type StartRollingMaintenanceResponse struct {
	Details      string                                        `json:"details"`
	Hostsskipped []StartRollingMaintenanceResponseHostsskipped `json:"hostsskipped"`
	Hostsupdated []StartRollingMaintenanceResponseHostsupdated `json:"hostsupdated"`
	JobID        string                                        `json:"jobid"`
	Jobstatus    int                                           `json:"jobstatus"`
	Success      bool                                          `json:"success"`
}

type StartRollingMaintenanceResponseHostsupdated struct {
	Enddate   string `json:"enddate"`
	Hostid    string `json:"hostid"`
	Hostname  string `json:"hostname"`
	Output    string `json:"output"`
	Startdate string `json:"startdate"`
}

type StartRollingMaintenanceResponseHostsskipped struct {
	Hostid   string `json:"hostid"`
	Hostname string `json:"hostname"`
	Reason   string `json:"reason"`
}
