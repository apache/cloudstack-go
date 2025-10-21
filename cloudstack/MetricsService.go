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

type MetricsServiceIface interface {
	ListInfrastructure(p *ListInfrastructureParams) (*ListInfrastructureResponse, error)
	NewListInfrastructureParams() *ListInfrastructureParams
}

type ListInfrastructureParams struct {
	p map[string]interface{}
}

func (p *ListInfrastructureParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListInfrastructureParams instance,
// as then you are sure you have configured all required params
func (s *MetricsService) NewListInfrastructureParams() *ListInfrastructureParams {
	p := &ListInfrastructureParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists infrastructure
func (s *MetricsService) ListInfrastructure(p *ListInfrastructureParams) (*ListInfrastructureResponse, error) {
	resp, err := s.cs.newRequest("listInfrastructure", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListInfrastructureResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListInfrastructureResponse struct {
	Count          int             `json:"count"`
	Infrastructure *Infrastructure `json:"infrastructure"`
}

type Infrastructure struct {
	Alerts             int    `json:"alerts"`
	Backuprepositories int    `json:"backuprepositories"`
	Clusters           int    `json:"clusters"`
	Cpusockets         int    `json:"cpusockets"`
	Hosts              int    `json:"hosts"`
	Ilbvms             int    `json:"ilbvms"`
	Imagestores        int    `json:"imagestores"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementservers  int    `json:"managementservers"`
	Objectstores       int    `json:"objectstores"`
	Pods               int    `json:"pods"`
	Routers            int    `json:"routers"`
	Storagepools       int    `json:"storagepools"`
	Systemvms          int    `json:"systemvms"`
	Zones              int    `json:"zones"`
}
