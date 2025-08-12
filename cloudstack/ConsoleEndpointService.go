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

type ConsoleEndpointServiceIface interface {
	CreateConsoleEndpoint(p *CreateConsoleEndpointParams) (*CreateConsoleEndpointResponse, error)
	NewCreateConsoleEndpointParams(virtualmachineid string) *CreateConsoleEndpointParams
}

type CreateConsoleEndpointParams struct {
	p map[string]interface{}
}

func (p *CreateConsoleEndpointParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["token"]; found {
		u.Set("token", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *CreateConsoleEndpointParams) SetToken(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["token"] = v
}

func (p *CreateConsoleEndpointParams) ResetToken() {
	if p.p != nil && p.p["token"] != nil {
		delete(p.p, "token")
	}
}

func (p *CreateConsoleEndpointParams) GetToken() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["token"].(string)
	return value, ok
}

func (p *CreateConsoleEndpointParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *CreateConsoleEndpointParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *CreateConsoleEndpointParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateConsoleEndpointParams instance,
// as then you are sure you have configured all required params
func (s *ConsoleEndpointService) NewCreateConsoleEndpointParams(virtualmachineid string) *CreateConsoleEndpointParams {
	p := &CreateConsoleEndpointParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Create a console endpoint to connect to a VM console
func (s *ConsoleEndpointService) CreateConsoleEndpoint(p *CreateConsoleEndpointParams) (*CreateConsoleEndpointResponse, error) {
	resp, err := s.cs.newPostRequest("createConsoleEndpoint", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response CreateConsoleEndpointResponse `json:"consoleendpoint"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type CreateConsoleEndpointResponse struct {
	Details   string                 `json:"details"`
	JobID     string                 `json:"jobid"`
	Jobstatus int                    `json:"jobstatus"`
	Success   bool                   `json:"success"`
	Url       string                 `json:"url"`
	Websocket map[string]interface{} `json:"websocket"`
}

func (r *CreateConsoleEndpointResponse) UnmarshalJSON(b []byte) error {
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

	type alias CreateConsoleEndpointResponse
	return json.Unmarshal(b, (*alias)(r))
}
