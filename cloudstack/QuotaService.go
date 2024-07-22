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

type QuotaServiceIface interface {
	QuotaBalance(p *QuotaBalanceParams) (*QuotaBalanceResponse, error)
	NewQuotaBalanceParams(account string, domainid string) *QuotaBalanceParams
	//CreateCredit(p *CreateCreditParams) (*CreateCreditResponse, error)
	QuotaIsEnabled(p *QuotaIsEnabledParams) (*QuotaIsEnabledResponse, error)
	NewQuotaIsEnabledParams() *QuotaIsEnabledParams
	//CreateStatement(p *CreateStatementParams) (*CreateStatementResponse, error)
	//ListSummary(p *GetSummaryParams) (*GetSummaryResponse, error)
	//CreateTariff(p *CreateTariffParams) (*CreateTariffResponse, error)
	//DeleteTariff(p *DeleteTariffParams) (*DeleteTariffResponse, error)
	//ListTariffs(p *ListTariffsParams) (*ListTariffsResponse, error)
	//UpdateTariff(p *UpdateTariffParams) (*UpdateTariffResponse, error)
	//UpdateQuota(p *UpdateQuotaParams) (*UpdateQuotaResponse, error)
}

type QuotaBalanceParams struct {
	p map[string]interface{}
}

func (p *QuotaBalanceParams) toURLValues() url.Values {
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
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *QuotaBalanceParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *QuotaBalanceParams) ResetAccount() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "account")
}

func (p *QuotaBalanceParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *QuotaBalanceParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *QuotaBalanceParams) ResetDomainid() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "domainid")
}

func (p *QuotaBalanceParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *QuotaBalanceParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *QuotaBalanceParams) ResetAccountid() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "accountid")
}

func (p *QuotaBalanceParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *QuotaBalanceParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *QuotaBalanceParams) ResetEnddate() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "enddate")
}

func (p *QuotaBalanceParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *QuotaBalanceParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *QuotaBalanceParams) ResetStartdate() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "startdate")
}

func (p *QuotaBalanceParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new QuotaBalanceParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaBalanceParams(account string, domainid string) *QuotaBalanceParams {
	p := &QuotaBalanceParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	return p
}

func (s *QuotaService) QuotaBalance(p *QuotaBalanceParams) (*QuotaBalanceResponse, error) {
	resp, err := s.cs.newRequest("quotaBalance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaBalanceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaBalanceResponse struct {
	Account   string `json:"account"`
	Accountid string `json:"accountid"`
	Domain    string `json:"domain"`
	Name      string `json:"name"`
	Quota     string `json:"quota"`
	Type      string `json:"type"`
	Unit      string `json:"unit"`
}

type QuotaIsEnabledParams struct {
	p map[string]interface{}
}

func (p *QuotaIsEnabledParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new QuotaIsEnabledParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaIsEnabledParams() *QuotaIsEnabledParams {
	p := &QuotaIsEnabledParams{}
	p.p = make(map[string]interface{})
	return p
}

// Return true if the plugin is enabled
func (s *QuotaService) QuotaIsEnabled(p *QuotaIsEnabledParams) (*QuotaIsEnabledResponse, error) {
	resp, err := s.cs.newRequest("quotaIsEnabled", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaIsEnabledResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaIsEnabledResponse struct {
	Isenabled bool   `json:"isenabled"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
}
