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

type ManagementServiceIface interface {
	CancelShutdown(p *CancelShutdownParams) (*CancelShutdownResponse, error)
	NewCancelShutdownParams(managementserverid UUID) *CancelShutdownParams
	ListManagementServers(p *ListManagementServersParams) (*ListManagementServersResponse, error)
	NewListManagementServersParams() *ListManagementServersParams
	GetManagementServerID(name string, opts ...OptionFunc) (string, int, error)
	GetManagementServerByName(name string, opts ...OptionFunc) (*ManagementServer, int, error)
	GetManagementServerByID(id string, opts ...OptionFunc) (*ManagementServer, int, error)
	ListManagementServersMetrics(p *ListManagementServersMetricsParams) (*ListManagementServersMetricsResponse, error)
	NewListManagementServersMetricsParams() *ListManagementServersMetricsParams
	GetManagementServersMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetManagementServersMetricByName(name string, opts ...OptionFunc) (*ManagementServersMetric, int, error)
	GetManagementServersMetricByID(id string, opts ...OptionFunc) (*ManagementServersMetric, int, error)
	PrepareForShutdown(p *PrepareForShutdownParams) (*PrepareForShutdownResponse, error)
	NewPrepareForShutdownParams(managementserverid UUID) *PrepareForShutdownParams
	ReadyForShutdown(p *ReadyForShutdownParams) (*ReadyForShutdownResponse, error)
	NewReadyForShutdownParams() *ReadyForShutdownParams
	TriggerShutdown(p *TriggerShutdownParams) (*TriggerShutdownResponse, error)
	NewTriggerShutdownParams(managementserverid UUID) *TriggerShutdownParams
}

type CancelShutdownParams struct {
	p map[string]interface{}
}

func (p *CancelShutdownParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	return u
}

func (p *CancelShutdownParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *CancelShutdownParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *CancelShutdownParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

// You should always use this function to get a new CancelShutdownParams instance,
// as then you are sure you have configured all required params
func (s *ManagementService) NewCancelShutdownParams(managementserverid UUID) *CancelShutdownParams {
	p := &CancelShutdownParams{}
	p.p = make(map[string]interface{})
	p.p["managementserverid"] = managementserverid
	return p
}

// Cancels a triggered shutdown
func (s *ManagementService) CancelShutdown(p *CancelShutdownParams) (*CancelShutdownResponse, error) {
	resp, err := s.cs.newRequest("cancelShutdown", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelShutdownResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CancelShutdownResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementserverid UUID   `json:"managementserverid"`
	Pendingjobscount   int64  `json:"pendingjobscount"`
	Readyforshutdown   bool   `json:"readyforshutdown"`
	Shutdowntriggered  bool   `json:"shutdowntriggered"`
}

type ListManagementServersParams struct {
	p map[string]interface{}
}

func (p *ListManagementServersParams) toURLValues() url.Values {
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
	return u
}

func (p *ListManagementServersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListManagementServersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListManagementServersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListManagementServersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListManagementServersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListManagementServersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListManagementServersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListManagementServersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListManagementServersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListManagementServersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListManagementServersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListManagementServersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListManagementServersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListManagementServersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListManagementServersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListManagementServersParams instance,
// as then you are sure you have configured all required params
func (s *ManagementService) NewListManagementServersParams() *ListManagementServersParams {
	p := &ListManagementServersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ManagementService) GetManagementServerID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListManagementServersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListManagementServers(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ManagementServers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ManagementServers {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ManagementService) GetManagementServerByName(name string, opts ...OptionFunc) (*ManagementServer, int, error) {
	id, count, err := s.GetManagementServerID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetManagementServerByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ManagementService) GetManagementServerByID(id string, opts ...OptionFunc) (*ManagementServer, int, error) {
	p := &ListManagementServersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListManagementServers(p)
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
		return l.ManagementServers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ManagementServer UUID: %s!", id)
}

// Lists management servers.
func (s *ManagementService) ListManagementServers(p *ListManagementServersParams) (*ListManagementServersResponse, error) {
	resp, err := s.cs.newRequest("listManagementServers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListManagementServersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListManagementServersResponse struct {
	Count             int                 `json:"count"`
	ManagementServers []*ManagementServer `json:"managementserver"`
}

type ManagementServer struct {
	Id               string `json:"id"`
	Javadistribution string `json:"javadistribution"`
	Javaversion      string `json:"javaversion"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Kernelversion    string `json:"kernelversion"`
	Lastboottime     string `json:"lastboottime"`
	Lastserverstart  string `json:"lastserverstart"`
	Lastserverstop   string `json:"lastserverstop"`
	Name             string `json:"name"`
	Osdistribution   string `json:"osdistribution"`
	Serviceip        string `json:"serviceip"`
	State            string `json:"state"`
	Version          string `json:"version"`
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

func (p *ListManagementServersMetricsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *ListManagementServersMetricsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListManagementServersMetricsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *ListManagementServersMetricsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListManagementServersMetricsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
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

func (p *ListManagementServersMetricsParams) ResetSystem() {
	if p.p != nil && p.p["system"] != nil {
		delete(p.p, "system")
	}
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
func (s *ManagementService) NewListManagementServersMetricsParams() *ListManagementServersMetricsParams {
	p := &ListManagementServersMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ManagementService) GetManagementServersMetricID(name string, opts ...OptionFunc) (string, int, error) {
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
func (s *ManagementService) GetManagementServersMetricByName(name string, opts ...OptionFunc) (*ManagementServersMetric, int, error) {
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
func (s *ManagementService) GetManagementServersMetricByID(id string, opts ...OptionFunc) (*ManagementServersMetric, int, error) {
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
func (s *ManagementService) ListManagementServersMetrics(p *ListManagementServersMetricsParams) (*ListManagementServersMetricsResponse, error) {
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
	ManagementServersMetrics []*ManagementServersMetric `json:"managementserver"`
}

type ManagementServersMetric struct {
	Agentcount              int       `json:"agentcount"`
	Availableprocessors     int       `json:"availableprocessors"`
	Collectiontime          string    `json:"collectiontime"`
	Cpuload                 string    `json:"cpuload"`
	Dbislocal               bool      `json:"dbislocal"`
	Heapmemorytotal         int64     `json:"heapmemorytotal"`
	Heapmemoryused          int64     `json:"heapmemoryused"`
	Id                      string    `json:"id"`
	Javadistribution        string    `json:"javadistribution"`
	Javaversion             string    `json:"javaversion"`
	JobID                   string    `json:"jobid"`
	Jobstatus               int       `json:"jobstatus"`
	Kernelversion           string    `json:"kernelversion"`
	Lastboottime            string    `json:"lastboottime"`
	Lastserverstart         string    `json:"lastserverstart"`
	Lastserverstop          string    `json:"lastserverstop"`
	Loginfo                 string    `json:"loginfo"`
	Name                    string    `json:"name"`
	Osdistribution          string    `json:"osdistribution"`
	Serviceip               string    `json:"serviceip"`
	Sessions                int64     `json:"sessions"`
	State                   string    `json:"state"`
	Systemcycleusage        []int64   `json:"systemcycleusage"`
	Systemloadaverages      []float64 `json:"systemloadaverages"`
	Systemmemoryfree        string    `json:"systemmemoryfree"`
	Systemmemorytotal       string    `json:"systemmemorytotal"`
	Systemmemoryused        string    `json:"systemmemoryused"`
	Systemmemoryvirtualsize string    `json:"systemmemoryvirtualsize"`
	Systemtotalcpucycles    float64   `json:"systemtotalcpucycles"`
	Threadsblockedcount     int       `json:"threadsblockedcount"`
	Threadsdaemoncount      int       `json:"threadsdaemoncount"`
	Threadsrunnablecount    int       `json:"threadsrunnablecount"`
	Threadsteminatedcount   int       `json:"threadsteminatedcount"`
	Threadstotalcount       int       `json:"threadstotalcount"`
	Threadswaitingcount     int       `json:"threadswaitingcount"`
	Usageislocal            bool      `json:"usageislocal"`
	Version                 string    `json:"version"`
}

type PrepareForShutdownParams struct {
	p map[string]interface{}
}

func (p *PrepareForShutdownParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	return u
}

func (p *PrepareForShutdownParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *PrepareForShutdownParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *PrepareForShutdownParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

// You should always use this function to get a new PrepareForShutdownParams instance,
// as then you are sure you have configured all required params
func (s *ManagementService) NewPrepareForShutdownParams(managementserverid UUID) *PrepareForShutdownParams {
	p := &PrepareForShutdownParams{}
	p.p = make(map[string]interface{})
	p.p["managementserverid"] = managementserverid
	return p
}

// Prepares CloudStack for a safe manual shutdown by preventing new jobs from being accepted
func (s *ManagementService) PrepareForShutdown(p *PrepareForShutdownParams) (*PrepareForShutdownResponse, error) {
	resp, err := s.cs.newRequest("prepareForShutdown", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PrepareForShutdownResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type PrepareForShutdownResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementserverid UUID   `json:"managementserverid"`
	Pendingjobscount   int64  `json:"pendingjobscount"`
	Readyforshutdown   bool   `json:"readyforshutdown"`
	Shutdowntriggered  bool   `json:"shutdowntriggered"`
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
func (s *ManagementService) NewReadyForShutdownParams() *ReadyForShutdownParams {
	p := &ReadyForShutdownParams{}
	p.p = make(map[string]interface{})
	return p
}

// Returns the status of CloudStack, whether a shutdown has been triggered and if ready to shutdown
func (s *ManagementService) ReadyForShutdown(p *ReadyForShutdownParams) (*ReadyForShutdownResponse, error) {
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

type TriggerShutdownParams struct {
	p map[string]interface{}
}

func (p *TriggerShutdownParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	return u
}

func (p *TriggerShutdownParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *TriggerShutdownParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *TriggerShutdownParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

// You should always use this function to get a new TriggerShutdownParams instance,
// as then you are sure you have configured all required params
func (s *ManagementService) NewTriggerShutdownParams(managementserverid UUID) *TriggerShutdownParams {
	p := &TriggerShutdownParams{}
	p.p = make(map[string]interface{})
	p.p["managementserverid"] = managementserverid
	return p
}

// Triggers an automatic safe shutdown of CloudStack by not accepting new jobs and shutting down when all pending jobbs have been completed. Triggers an immediate shutdown if forced
func (s *ManagementService) TriggerShutdown(p *TriggerShutdownParams) (*TriggerShutdownResponse, error) {
	resp, err := s.cs.newRequest("triggerShutdown", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r TriggerShutdownResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type TriggerShutdownResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementserverid UUID   `json:"managementserverid"`
	Pendingjobscount   int64  `json:"pendingjobscount"`
	Readyforshutdown   bool   `json:"readyforshutdown"`
	Shutdowntriggered  bool   `json:"shutdowntriggered"`
}
