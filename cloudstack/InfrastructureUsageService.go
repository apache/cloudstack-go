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

type InfrastructureUsageServiceIface interface {
	ListDbMetrics(p *ListDbMetricsParams) (*ListDbMetricsResponse, error)
	NewListDbMetricsParams() *ListDbMetricsParams
}

type ListDbMetricsParams struct {
	p map[string]interface{}
}

func (p *ListDbMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListDbMetricsParams instance,
// as then you are sure you have configured all required params
func (s *InfrastructureUsageService) NewListDbMetricsParams() *ListDbMetricsParams {
	p := &ListDbMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// list the db hosts and statistics
func (s *InfrastructureUsageService) ListDbMetrics(p *ListDbMetricsParams) (*ListDbMetricsResponse, error) {
	resp, err := s.cs.newRequest("listDbMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDbMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDbMetricsResponse struct {
	DbMetrics DbMetric `json:"dbMetrics"`
}

type DbMetric struct {
	Collectiontime string    `json:"collectiontime"`
	Connections    int       `json:"connections"`
	Dbloadaverages []float64 `json:"dbloadaverages"`
	Hostname       string    `json:"hostname"`
	JobID          string    `json:"jobid"`
	Jobstatus      int       `json:"jobstatus"`
	Queries        int64     `json:"queries"`
	Replicas       []string  `json:"replicas"`
	Tlsversions    string    `json:"tlsversions"`
	Uptime         int64     `json:"uptime"`
	Version        string    `json:"version"`
	Versioncomment string    `json:"versioncomment"`
}
