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
	"time"
)

type AsyncjobServiceIface interface {
	ListAsyncJobs(p *ListAsyncJobsParams) (*ListAsyncJobsResponse, error)
	NewListAsyncJobsParams() *ListAsyncJobsParams
	QueryAsyncJobResult(p *QueryAsyncJobResultParams) (*QueryAsyncJobResultResponse, error)
	NewQueryAsyncJobResultParams(jobid string) *QueryAsyncJobResultParams
}

type ListAsyncJobsParams struct {
	p map[string]interface{}
}

func (p *ListAsyncJobsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *ListAsyncJobsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListAsyncJobsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListAsyncJobsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListAsyncJobsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListAsyncJobsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListAsyncJobsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListAsyncJobsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAsyncJobsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListAsyncJobsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListAsyncJobsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListAsyncJobsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAsyncJobsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListAsyncJobsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAsyncJobsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListAsyncJobsParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *ListAsyncJobsParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new ListAsyncJobsParams instance,
// as then you are sure you have configured all required params
func (s *AsyncjobService) NewListAsyncJobsParams() *ListAsyncJobsParams {
	p := &ListAsyncJobsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all pending asynchronous jobs for the account.
func (s *AsyncjobService) ListAsyncJobs(p *ListAsyncJobsParams) (*ListAsyncJobsResponse, error) {
	resp, err := s.cs.newRequest("listAsyncJobs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAsyncJobsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAsyncJobsResponse struct {
	Count     int         `json:"count"`
	AsyncJobs []*AsyncJob `json:"asyncjobs"`
}

type AsyncJob struct {
	Accountid       string          `json:"accountid"`
	Cmd             string          `json:"cmd"`
	Completed       string          `json:"completed"`
	Created         string          `json:"created"`
	JobID           string          `json:"jobid"`
	Jobinstanceid   string          `json:"jobinstanceid"`
	Jobinstancetype string          `json:"jobinstancetype"`
	Jobprocstatus   int             `json:"jobprocstatus"`
	Jobresult       json.RawMessage `json:"jobresult"`
	Jobresultcode   int             `json:"jobresultcode"`
	Jobresulttype   string          `json:"jobresulttype"`
	Jobstatus       int             `json:"jobstatus"`
	Userid          string          `json:"userid"`
}

type QueryAsyncJobResultParams struct {
	p map[string]interface{}
}

func (p *QueryAsyncJobResultParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["jobid"]; found {
		u.Set("jobid", v.(string))
	}
	return u
}

func (p *QueryAsyncJobResultParams) SetJobID(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["jobid"] = v
}

func (p *QueryAsyncJobResultParams) GetJobID() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["jobid"].(string)
	return value, ok
}

// You should always use this function to get a new QueryAsyncJobResultParams instance,
// as then you are sure you have configured all required params
func (s *AsyncjobService) NewQueryAsyncJobResultParams(jobid string) *QueryAsyncJobResultParams {
	p := &QueryAsyncJobResultParams{}
	p.p = make(map[string]interface{})
	p.p["jobid"] = jobid
	return p
}

// Retrieves the current status of asynchronous job.
func (s *AsyncjobService) QueryAsyncJobResult(p *QueryAsyncJobResultParams) (*QueryAsyncJobResultResponse, error) {
	var resp json.RawMessage
	var err error

	// We should be able to retry on failure as this call is idempotent
	for i := 0; i < 3; i++ {
		resp, err = s.cs.newRequest("queryAsyncJobResult", p.toURLValues())
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		return nil, err
	}

	var r QueryAsyncJobResultResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QueryAsyncJobResultResponse struct {
	Accountid       string          `json:"accountid"`
	Cmd             string          `json:"cmd"`
	Completed       string          `json:"completed"`
	Created         string          `json:"created"`
	JobID           string          `json:"jobid"`
	Jobinstanceid   string          `json:"jobinstanceid"`
	Jobinstancetype string          `json:"jobinstancetype"`
	Jobprocstatus   int             `json:"jobprocstatus"`
	Jobresult       json.RawMessage `json:"jobresult"`
	Jobresultcode   int             `json:"jobresultcode"`
	Jobresulttype   string          `json:"jobresulttype"`
	Jobstatus       int             `json:"jobstatus"`
	Userid          string          `json:"userid"`
}
