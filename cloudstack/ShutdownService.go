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

type ShutdownServiceIface interface {
	ReadyForShutdown(p *ReadyForShutdownParams) (*ReadyForShutdownResponse, error)
	NewReadyForShutdownParams() *ReadyForShutdownParams
}

type ReadyForShutdownParams struct {
	p map[string]interface{}
}

func (p *ReadyForShutdownParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	return u
}

func (p *ReadyForShutdownParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *ReadyForShutdownParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *ReadyForShutdownParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

// You should always use this function to get a new ReadyForShutdownParams instance,
// as then you are sure you have configured all required params
func (s *ShutdownService) NewReadyForShutdownParams() *ReadyForShutdownParams {
	p := &ReadyForShutdownParams{}
	p.p = make(map[string]interface{})
	return p
}

// Returns the status of CloudStack, whether a shutdown has been triggered and if ready to shutdown
func (s *ShutdownService) ReadyForShutdown(p *ReadyForShutdownParams) (*ReadyForShutdownResponse, error) {
	resp, err := s.cs.newRequest("readyForShutdown", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response ReadyForShutdownResponse `json:"readyforshutdown"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type ReadyForShutdownResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementserverid UUID   `json:"managementserverid"`
	Pendingjobscount   int64  `json:"pendingjobscount"`
	Readyforshutdown   bool   `json:"readyforshutdown"`
	Shutdowntriggered  bool   `json:"shutdowntriggered"`
}
