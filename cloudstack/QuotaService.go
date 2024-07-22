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
	NewQuotaBalanceParams(account, domainid string) *QuotaBalanceParams
	QuotaCredits(p *QuotaCreditsParams) (*QuotaCreditsParams, error)
	NewQuotaCreditsParams(account, domainid, value string) *QuotaCreditsParams
	QuotaIsEnabled(p *QuotaIsEnabledParams) (*QuotaIsEnabledResponse, error)
	NewQuotaIsEnabledParams() *QuotaIsEnabledParams
	QuotaStatement(p *QuotaStatementParams) (*QuotaStatementResponse, error)
	NewQuotaStatementParams(account, domainid, enddate, startdate string) *QuotaStatementParams
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

type QuotaCreditsParams struct {
	p map[string]interface{}
}

func (p *QuotaCreditsParams) toURLValues() url.Values {
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
	if v, found := p.p["value"]; found {
		u.Set("value", v.(string))
	}
	if v, found := p.p["min_balance"]; found {
		u.Set("min_balance", v.(string))

	}
	if v, found := p.p["quota_enforce"]; found {
		u.Set("quota_enforce", v.(string))
	}
	return u
}

func (p *QuotaCreditsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *QuotaCreditsParams) ResetAccount() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "account")
}

func (p *QuotaCreditsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *QuotaCreditsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *QuotaCreditsParams) ResetDomainid() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "domainid")
}

func (p *QuotaCreditsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *QuotaCreditsParams) SetValue(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["value"] = v
}

func (p *QuotaCreditsParams) ResetValue() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "value")
}

func (p *QuotaCreditsParams) GetValue() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["value"].(string)
	return value, ok
}

func (p *QuotaCreditsParams) SetMinBalance(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["min_balance"] = v
}

func (p *QuotaCreditsParams) ResetMinBalance() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "min_balance")
}

func (p *QuotaCreditsParams) GetMinBalance() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["min_balance"].(string)
	return value, ok
}

func (p *QuotaCreditsParams) SetQuotaEnforce(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quota_enforce"] = v
}

func (p *QuotaCreditsParams) ResetQuotaEnforce() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "quota_enforce")
}

func (p *QuotaCreditsParams) GetQuotaEnforce() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quota_enforce"].(string)
	return value, ok
}

// You should always use this function to get a new QuotaCreditsParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaCreditsParams(account, domainid, value string) *QuotaCreditsParams {
	p := &QuotaCreditsParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	p.p["value"] = value
	return p
}

func (s *QuotaService) QuotaCredits(p *QuotaCreditsParams) (*QuotaCreditsResponse, error) {
	resp, err := s.cs.newRequest("quotaCredits", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaCreditsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaCreditsResponse struct {
	Credits    string `json:"credits"`
	Currency   string `json:"currency"`
	Updated_by string `json:"updated_by"`
	Updated_on string `json:"updated_on"`
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

type QuotaStatementParams struct {
	p map[string]interface{}
}

func (p *QuotaStatementParams) toURLValues() url.Values {
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
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *QuotaStatementParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *QuotaStatementParams) ResetAccount() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "account")
}

func (p *QuotaStatementParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *QuotaStatementParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *QuotaStatementParams) ResetDomainid() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "domainid")
}

func (p *QuotaStatementParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *QuotaStatementParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *QuotaStatementParams) ResetEnddate() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "enddate")
}

func (p *QuotaStatementParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *QuotaStatementParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *QuotaStatementParams) ResetStartdate() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "startdate")
}

func (p *QuotaStatementParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *QuotaStatementParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
}

func (p *QuotaStatementParams) ResetAccountid() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "accountid")
}

func (p *QuotaStatementParams) GetAccountid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["accountid"].(string)
	return value, ok
}

func (p *QuotaStatementParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *QuotaStatementParams) ResetType() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "type")
}

func (p *QuotaStatementParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new QuotaStatementParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaStatementParams(account, domainid, enddate, startdate string) *QuotaStatementParams {
	p := &QuotaStatementParams{}
	p.p = make(map[string]interface{})
	p.p["account"] = account
	p.p["domainid"] = domainid
	p.p["enddate"] = enddate
	p.p["startdate"] = startdate
	return p
}

func (s *QuotaService) QuotaStatement(p *QuotaStatementParams) (*QuotaStatementResponse, error) {
	resp, err := s.cs.newRequest("quotaStatement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaStatementResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaStatementResponse struct {
	Account   string `json:"account"`
	Accountid string `json:"accountid"`
	Domain    string `json:"domain"`
	Name      string `json:"name"`
	Quota     string `json:"quota"`
	Type      string `json:"type"`
	Unit      string `json:"unit"`
}
