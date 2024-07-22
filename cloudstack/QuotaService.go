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

type QuotaServiceIface interface {
	QuotaBalance(p *QuotaBalanceParams) (*QuotaBalanceResponse, error)
	NewQuotaBalanceParams(account, domainid string) *QuotaBalanceParams
	QuotaCredits(p *QuotaCreditsParams) (*QuotaCreditsParams, error)
	NewQuotaCreditsParams(account, domainid, value string) *QuotaCreditsParams
	QuotaIsEnabled(p *QuotaIsEnabledParams) (*QuotaIsEnabledResponse, error)
	NewQuotaIsEnabledParams() *QuotaIsEnabledParams
	QuotaStatement(p *QuotaStatementParams) (*QuotaStatementResponse, error)
	NewQuotaStatementParams(account, domainid, enddate, startdate string) *QuotaStatementParams
	QuotaSummary(p *QuotaSummaryParams) (*QuotaSummaryResponse, error)
	NewQuotaSummaryParams() *QuotaSummaryParams
	QuotaTariffCreate(p *QuotaTariffCreateParams) (*QuotaTariffCreateResponse, error)
	NewQuotaTariffCreateParams(name, usagetype, value string) *QuotaTariffCreateParams
	QuotaTariffDelete(p *QuotaTariffDeleteParams) (*QuotaTariffDeleteResponse, error)
	NewQuotaTariffDeleteParams(id string) *QuotaTariffDeleteParams
	QuotaTariffList(p *QuotaTariffListParams) (*QuotaTariffListResponse, error)
	NewQuotaTariffListParams() *QuotaTariffListParams
	QuotaTariffUpdate(p *QuotaTariffUpdateParams) (*QuotaTariffUpdateResponse, error)
	NewQuotaTariffUpdateParams(name string) *QuotaTariffUpdateParams
	QuotaUpdate() (*QuotaUpdateResponse, error)
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

type QuotaSummaryParams struct {
	p map[string]interface{}
}

func (p *QuotaSummaryParams) toURLValues() url.Values {
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
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		u.Add("listall", v.(string))
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

func (p *QuotaSummaryParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *QuotaSummaryParams) ResetAccount() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "account")
}

func (p *QuotaSummaryParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *QuotaSummaryParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *QuotaSummaryParams) ResetDomainid() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "domainid")
}

func (p *QuotaSummaryParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *QuotaSummaryParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *QuotaSummaryParams) ResetKeyword() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "keyword")
}

func (p *QuotaSummaryParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *QuotaSummaryParams) SetListall(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *QuotaSummaryParams) ResetListall() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "listall")
}

func (p *QuotaSummaryParams) GetListall() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(string)
	return value, ok
}

func (p *QuotaSummaryParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *QuotaSummaryParams) ResetPage() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "page")
}

func (p *QuotaSummaryParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *QuotaSummaryParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *QuotaSummaryParams) ResetPagesize() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "pagesize")
}

func (p *QuotaSummaryParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new QuotaSummaryParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaSummaryParams() *QuotaSummaryParams {
	p := &QuotaSummaryParams{}
	p.p = make(map[string]interface{})
	return p
}

func (s *QuotaService) QuotaSummary(p *QuotaSummaryParams) (*QuotaSummaryResponse, error) {
	resp, err := s.cs.newRequest("quotaSummary", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaSummaryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaSummaryResponse struct {
	Account      string `json:"account"`
	Accountid    string `json:"accountid"`
	Balance      string `json:"balance"`
	Currency     string `json:"currency"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Enddate      string `json:"enddate"`
	Quota        string `json:"quota"`
	Quotaenabled string `json:"quotaenabled"`
	Startdate    string `json:"startdate"`
	State        string `json:"state"`
}

type QuotaTariffCreateParams struct {
	p map[string]interface{}
}

func (p *QuotaTariffCreateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["usagetype"]; found {
		u.Set("usagetype", v.(string))
	}
	if v, found := p.p["value"]; found {
		u.Set("value", v.(string))
	}
	if v, found := p.p["activationrule"]; found {
		u.Set("activationrule", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *QuotaTariffCreateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *QuotaTariffCreateParams) ResetName() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "name")
}

func (p *QuotaTariffCreateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *QuotaTariffCreateParams) SetUsagetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usagetype"] = v
}

func (p *QuotaTariffCreateParams) ResetUsagetype() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "usagetype")
}

func (p *QuotaTariffCreateParams) GetUsagetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["usagetype"].(string)
	return value, ok
}

func (p *QuotaTariffCreateParams) SetValue(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["value"] = v
}

func (p *QuotaTariffCreateParams) ResetValue() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "value")
}

func (p *QuotaTariffCreateParams) GetValue() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["value"].(string)
	return value, ok
}

func (p *QuotaTariffCreateParams) SetActivationrule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["activationrule"] = v
}

func (p *QuotaTariffCreateParams) ResetActivationrule() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "activationrule")
}

func (p *QuotaTariffCreateParams) GetActivationrule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["activationrule"].(string)
	return value, ok
}

func (p *QuotaTariffCreateParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *QuotaTariffCreateParams) ResetDescription() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "description")
}

func (p *QuotaTariffCreateParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *QuotaTariffCreateParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *QuotaTariffCreateParams) ResetEnddate() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "enddate")
}

func (p *QuotaTariffCreateParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *QuotaTariffCreateParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *QuotaTariffCreateParams) ResetStartdate() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "startdate")
}

func (p *QuotaTariffCreateParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new QuotaTariffCreateParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaTariffCreateParams(name, usagetype, value string) *QuotaTariffCreateParams {
	p := &QuotaTariffCreateParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["usagetype"] = usagetype
	p.p["value"] = value
	return p
}

func (s *QuotaService) QuotaTariffCreate(p *QuotaTariffCreateParams) (*QuotaTariffCreateResponse, error) {
	resp, err := s.cs.newRequest("quotaTariffCreate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaTariffCreateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaTariffCreateResponse struct {
	Activationrule       string `json:"activationRule"`
	Currency             string `json:"currency"`
	Description          string `json:"description"`
	EffectiveDate        string `json:"effectiveDate"`
	EndDate              string `json:"endDate"`
	Name                 string `json:"name"`
	Removed              string `json:"removed"`
	TariffValue          string `json:"tariffValue"`
	UsageDiscriminator   string `json:"usageDiscriminator"`
	UsageName            string `json:"usageName"`
	UsageType            string `json:"usageType"`
	UsageTypeDescription string `json:"usageTypeDescription"`
	UsageUnit            string `json:"usageUnit"`
	Uuid                 string `json:"uuid"`
}

type QuotaTariffDeleteParams struct {
	p map[string]interface{}
}

func (p *QuotaTariffDeleteParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *QuotaTariffDeleteParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *QuotaTariffDeleteParams) ResetId() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "id")
}

func (p *QuotaTariffDeleteParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new QuotaTariffDeleteParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaTariffDeleteParams(id string) *QuotaTariffDeleteParams {
	p := &QuotaTariffDeleteParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

func (s *QuotaService) QuotaTariffDelete(p *QuotaTariffDeleteParams) (*QuotaTariffDeleteResponse, error) {
	resp, err := s.cs.newRequest("quotaTariffDelete", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaTariffDeleteResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaTariffDeleteResponse struct {
	DisplayText string `json:"displaytext"`
	Success     bool   `json:"success"`
}

type QuotaTariffListParams struct {
	p map[string]interface{}
}

func (p *QuotaTariffListParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new QuotaTariffListParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaTariffListParams() *QuotaTariffListParams {
	p := &QuotaTariffListParams{}
	p.p = make(map[string]interface{})
	return p
}

func (s *QuotaService) QuotaTariffList(p *QuotaTariffListParams) (*QuotaTariffListResponse, error) {
	resp, err := s.cs.newRequest("quotaTariffList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaTariffListResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaTariffListResponse struct {
	Activationrule       string `json:"activationRule"`
	Currency             string `json:"currency"`
	Description          string `json:"description"`
	EffectiveDate        string `json:"effectiveDate"`
	EndDate              string `json:"endDate"`
	Name                 string `json:"name"`
	Removed              string `json:"removed"`
	TariffValue          string `json:"tariffValue"`
	UsageDiscriminator   string `json:"usageDiscriminator"`
	UsageName            string `json:"usageName"`
	UsageType            string `json:"usageType"`
	UsageTypeDescription string `json:"usageTypeDescription"`
	UsageUnit            string `json:"usageUnit"`
	Uuid                 string `json:"uuid"`
}

type QuotaTariffUpdateParams struct {
	p map[string]interface{}
}

func (p *QuotaTariffUpdateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["activationrule"]; found {
		u.Set("activationrule", v.(string))
	}
	if v, found := p.p["usagetype"]; found {
		u.Set("usagetype", v.(string))
	}
	if v, found := p.p["value"]; found {
		u.Set("value", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (p *QuotaTariffUpdateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *QuotaTariffUpdateParams) ResetName() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "name")
}

func (p *QuotaTariffUpdateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *QuotaTariffUpdateParams) SetActivationrule(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["activationrule"] = v
}

func (p *QuotaTariffUpdateParams) ResetActivationrule() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "activationrule")
}

func (p *QuotaTariffUpdateParams) GetActivationrule() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["activationrule"].(string)
	return value, ok
}

func (p *QuotaTariffUpdateParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *QuotaTariffUpdateParams) ResetDescription() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "description")
}

func (p *QuotaTariffUpdateParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *QuotaTariffUpdateParams) SetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *QuotaTariffUpdateParams) ResetEnddate() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "enddate")
}

func (p *QuotaTariffUpdateParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *QuotaTariffUpdateParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *QuotaTariffUpdateParams) ResetStartdate() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "startdate")
}

func (p *QuotaTariffUpdateParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *QuotaTariffUpdateParams) SetUsagetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usagetype"] = v
}

func (p *QuotaTariffUpdateParams) ResetUsagetype() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "usagetype")
}

func (p *QuotaTariffUpdateParams) GetUsagetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["usagetype"].(string)
	return value, ok
}

func (p *QuotaTariffUpdateParams) SetValue(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["value"] = v
}

func (p *QuotaTariffUpdateParams) ResetValue() {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	delete(p.p, "value")
}

func (p *QuotaTariffUpdateParams) GetValue() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["value"].(string)
	return value, ok
}

// You should always use this function to get a new QuotaTariffUpdateParams instance,
// as then you are sure you have configured all required params
func (s *QuotaService) NewQuotaTariffUpdateParams(name string) *QuotaTariffUpdateParams {
	p := &QuotaTariffUpdateParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

func (s *QuotaService) QuotaTariffUpdate(p *QuotaTariffUpdateParams) (*QuotaTariffUpdateResponse, error) {
	resp, err := s.cs.newRequest("quotaTariffUpdate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r QuotaTariffUpdateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaTariffUpdateResponse struct {
	Activationrule       string `json:"activationRule"`
	Currency             string `json:"currency"`
	Description          string `json:"description"`
	EffectiveDate        string `json:"effectiveDate"`
	EndDate              string `json:"endDate"`
	Name                 string `json:"name"`
	Removed              string `json:"removed"`
	TariffValue          string `json:"tariffValue"`
	UsageDiscriminator   string `json:"usageDiscriminator"`
	UsageName            string `json:"usageName"`
	UsageType            string `json:"usageType"`
	UsageTypeDescription string `json:"usageTypeDescription"`
	UsageUnit            string `json:"usageUnit"`
	Uuid                 string `json:"uuid"`
}

func (s *QuotaService) QuotaUpdate() (*QuotaUpdateResponse, error) {
	resp, err := s.cs.newRequest("quotaUpdate")
	if err != nil {
		return nil, err
	}

	var r QuotaUpdateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QuotaUpdateResponse struct {
	UpdatedOn string `json:"updated_on"`
}
