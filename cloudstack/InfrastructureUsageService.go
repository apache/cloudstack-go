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

type InfrastructureUsageServiceIface interface {
	ListManagementServersMetrics(p *ListManagementServersMetricsParams) (*ListManagementServersMetricsResponse, error)
	NewListManagementServersMetricsParams() *ListManagementServersMetricsParams
	GetManagementServersMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetManagementServersMetricByName(name string, opts ...OptionFunc) (*ManagementServersMetric, int, error)
	GetManagementServersMetricByID(id string, opts ...OptionFunc) (*ManagementServersMetric, int, error)
	ListDbMetrics(p *ListDbMetricsParams) (*ListDbMetricsResponse, error)
	NewListDbMetricsParams() *ListDbMetricsParams
}

type ListManagementServersMetricsParams struct {
	p map[string]interface{}
}

func (p *ListManagementServersMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
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
	if v, found := p.p["system"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("system", vv)
	}
	return u
}

func (p *ListManagementServersMetricsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListManagementServersMetricsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListManagementServersMetricsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListManagementServersMetricsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListManagementServersMetricsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListManagementServersMetricsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListManagementServersMetricsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListManagementServersMetricsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListManagementServersMetricsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListManagementServersMetricsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListManagementServersMetricsParams) SetSystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["system"] = v
}

func (p *ListManagementServersMetricsParams) GetSystem() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["system"].(bool)
	return value, ok
}

// You should always use this function to get a new ListManagementServersMetricsParams instance,
// as then you are sure you have configured all required params
func (s *InfrastructureUsageService) NewListManagementServersMetricsParams() *ListManagementServersMetricsParams {
	p := &ListManagementServersMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InfrastructureUsageService) GetManagementServersMetricID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListManagementServersMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListManagementServersMetrics(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ManagementServersMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ManagementServersMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InfrastructureUsageService) GetManagementServersMetricByName(name string, opts ...OptionFunc) (*ManagementServersMetric, int, error) {
	id, count, err := s.GetManagementServersMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetManagementServersMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InfrastructureUsageService) GetManagementServersMetricByID(id string, opts ...OptionFunc) (*ManagementServersMetric, int, error) {
	p := &ListManagementServersMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListManagementServersMetrics(p)
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
		return l.ManagementServersMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ManagementServersMetric UUID: %s!", id)
}

// Lists Management Server metrics
func (s *InfrastructureUsageService) ListManagementServersMetrics(p *ListManagementServersMetricsParams) (*ListManagementServersMetricsResponse, error) {
	resp, err := s.cs.newRequest("listManagementServersMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListManagementServersMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListManagementServersMetricsResponse struct {
	Count                    int                        `json:"count"`
	ManagementServersMetrics []*ManagementServersMetric `json:"managementserversmetric"`
}

type ManagementServersMetric struct {
	Agentcount              int     `json:"agentcount"`
	Availableprocessors     int     `json:"availableprocessors"`
	Collectiontime          string  `json:"collectiontime"`
	Cpuload                 string  `json:"cpuload"`
	Dbislocal               bool    `json:"dbislocal"`
	Heapmemorytotal         int64   `json:"heapmemorytotal"`
	Heapmemoryused          int64   `json:"heapmemoryused"`
	Id                      string  `json:"id"`
	Javadistribution        string  `json:"javadistribution"`
	Javaversion             string  `json:"javaversion"`
	JobID                   string  `json:"jobid"`
	Jobstatus               int     `json:"jobstatus"`
	Kernelversion           string  `json:"kernelversion"`
	Lastboottime            string  `json:"lastboottime"`
	Lastserverstart         string  `json:"lastserverstart"`
	Lastserverstop          string  `json:"lastserverstop"`
	Loginfo                 string  `json:"loginfo"`
	Name                    string  `json:"name"`
	Osdistribution          string  `json:"osdistribution"`
	Sessions                int64   `json:"sessions"`
	State                   string  `json:"state"`
	Systemcycleusage        string  `json:"systemcycleusage"`
	Systemloadaverages      string  `json:"systemloadaverages"`
	Systemmemoryfree        string  `json:"systemmemoryfree"`
	Systemmemorytotal       string  `json:"systemmemorytotal"`
	Systemmemoryused        string  `json:"systemmemoryused"`
	Systemmemoryvirtualsize string  `json:"systemmemoryvirtualsize"`
	Systemtotalcpucycles    float64 `json:"systemtotalcpucycles"`
	Threadsblockedcount     int     `json:"threadsblockedcount"`
	Threadsdaemoncount      int     `json:"threadsdaemoncount"`
	Threadsrunnablecount    int     `json:"threadsrunnablecount"`
	Threadsteminatedcount   int     `json:"threadsteminatedcount"`
	Threadstotalcount       int     `json:"threadstotalcount"`
	Threadswaitingcount     int     `json:"threadswaitingcount"`
	Usageislocal            bool    `json:"usageislocal"`
	Version                 string  `json:"version"`
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
	Count     int         `json:"count"`
	DbMetrics []*DbMetric `json:"dbmetric"`
}

type DbMetric struct {
	Collectiontime string `json:"collectiontime"`
	Connections    int    `json:"connections"`
	Dbloadaverages string `json:"dbloadaverages"`
	Hostname       string `json:"hostname"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Queries        int64  `json:"queries"`
	Replicas       string `json:"replicas"`
	Tlsversions    string `json:"tlsversions"`
	Uptime         int64  `json:"uptime"`
	Version        string `json:"version"`
	Versioncomment string `json:"versioncomment"`
}
