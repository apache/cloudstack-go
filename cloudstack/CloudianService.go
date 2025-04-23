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

type CloudianServiceIface interface {
	CloudianIsEnabled(p *CloudianIsEnabledParams) (*CloudianIsEnabledResponse, error)
	NewCloudianIsEnabledParams() *CloudianIsEnabledParams
}

type CloudianIsEnabledParams struct {
	p map[string]interface{}
}

func (p *CloudianIsEnabledParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new CloudianIsEnabledParams instance,
// as then you are sure you have configured all required params
func (s *CloudianService) NewCloudianIsEnabledParams() *CloudianIsEnabledParams {
	p := &CloudianIsEnabledParams{}
	p.p = make(map[string]interface{})
	return p
}

// Checks if the Cloudian Connector is enabled
func (s *CloudianService) CloudianIsEnabled(p *CloudianIsEnabledParams) (*CloudianIsEnabledResponse, error) {
	resp, err := s.cs.newRequest("cloudianIsEnabled", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CloudianIsEnabledResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CloudianIsEnabledResponse struct {
	Enabled   bool   `json:"enabled"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Url       string `json:"url"`
}
