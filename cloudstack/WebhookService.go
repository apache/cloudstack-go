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

type WebhookServiceIface interface {
	CreateWebhook(p *CreateWebhookParams) (*CreateWebhookResponse, error)
	NewCreateWebhookParams(name string, payloadurl string) *CreateWebhookParams
	DeleteWebhook(p *DeleteWebhookParams) (*DeleteWebhookResponse, error)
	NewDeleteWebhookParams(id string) *DeleteWebhookParams
	DeleteWebhookDelivery(p *DeleteWebhookDeliveryParams) (*DeleteWebhookDeliveryResponse, error)
	NewDeleteWebhookDeliveryParams() *DeleteWebhookDeliveryParams
	ExecuteWebhookDelivery(p *ExecuteWebhookDeliveryParams) (*ExecuteWebhookDeliveryResponse, error)
	NewExecuteWebhookDeliveryParams() *ExecuteWebhookDeliveryParams
	ListWebhookDeliveries(p *ListWebhookDeliveriesParams) (*ListWebhookDeliveriesResponse, error)
	NewListWebhookDeliveriesParams() *ListWebhookDeliveriesParams
	GetWebhookDeliveryID(keyword string, opts ...OptionFunc) (string, int, error)
	GetWebhookDeliveryByName(name string, opts ...OptionFunc) (*WebhookDelivery, int, error)
	GetWebhookDeliveryByID(id string, opts ...OptionFunc) (*WebhookDelivery, int, error)
	ListWebhooks(p *ListWebhooksParams) (*ListWebhooksResponse, error)
	NewListWebhooksParams() *ListWebhooksParams
	GetWebhookID(name string, opts ...OptionFunc) (string, int, error)
	GetWebhookByName(name string, opts ...OptionFunc) (*Webhook, int, error)
	GetWebhookByID(id string, opts ...OptionFunc) (*Webhook, int, error)
	UpdateWebhook(p *UpdateWebhookParams) (*UpdateWebhookResponse, error)
	NewUpdateWebhookParams(id string) *UpdateWebhookParams
}

type CreateWebhookParams struct {
	p map[string]interface{}
}

func (p *CreateWebhookParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["payloadurl"]; found {
		u.Set("payloadurl", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["secretkey"]; found {
		u.Set("secretkey", v.(string))
	}
	if v, found := p.p["sslverification"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sslverification", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *CreateWebhookParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateWebhookParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateWebhookParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateWebhookParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *CreateWebhookParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *CreateWebhookParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *CreateWebhookParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateWebhookParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateWebhookParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateWebhookParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateWebhookParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateWebhookParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateWebhookParams) SetPayloadurl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["payloadurl"] = v
}

func (p *CreateWebhookParams) ResetPayloadurl() {
	if p.p != nil && p.p["payloadurl"] != nil {
		delete(p.p, "payloadurl")
	}
}

func (p *CreateWebhookParams) GetPayloadurl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["payloadurl"].(string)
	return value, ok
}

func (p *CreateWebhookParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateWebhookParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateWebhookParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateWebhookParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *CreateWebhookParams) ResetScope() {
	if p.p != nil && p.p["scope"] != nil {
		delete(p.p, "scope")
	}
}

func (p *CreateWebhookParams) GetScope() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scope"].(string)
	return value, ok
}

func (p *CreateWebhookParams) SetSecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secretkey"] = v
}

func (p *CreateWebhookParams) ResetSecretkey() {
	if p.p != nil && p.p["secretkey"] != nil {
		delete(p.p, "secretkey")
	}
}

func (p *CreateWebhookParams) GetSecretkey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["secretkey"].(string)
	return value, ok
}

func (p *CreateWebhookParams) SetSslverification(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sslverification"] = v
}

func (p *CreateWebhookParams) ResetSslverification() {
	if p.p != nil && p.p["sslverification"] != nil {
		delete(p.p, "sslverification")
	}
}

func (p *CreateWebhookParams) GetSslverification() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sslverification"].(bool)
	return value, ok
}

func (p *CreateWebhookParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *CreateWebhookParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *CreateWebhookParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

// You should always use this function to get a new CreateWebhookParams instance,
// as then you are sure you have configured all required params
func (s *WebhookService) NewCreateWebhookParams(name string, payloadurl string) *CreateWebhookParams {
	p := &CreateWebhookParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["payloadurl"] = payloadurl
	return p
}

// Creates a Webhook
func (s *WebhookService) CreateWebhook(p *CreateWebhookParams) (*CreateWebhookResponse, error) {
	resp, err := s.cs.newRequest("createWebhook", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateWebhookResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateWebhookResponse struct {
	Account         string `json:"account"`
	Created         string `json:"created"`
	Description     string `json:"description"`
	Domain          string `json:"domain"`
	Domainid        string `json:"domainid"`
	Domainpath      string `json:"domainpath"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Name            string `json:"name"`
	Payloadurl      string `json:"payloadurl"`
	Project         string `json:"project"`
	Projectid       string `json:"projectid"`
	Scope           string `json:"scope"`
	Secretkey       string `json:"secretkey"`
	Sslverification bool   `json:"sslverification"`
	State           string `json:"state"`
}

type DeleteWebhookParams struct {
	p map[string]interface{}
}

func (p *DeleteWebhookParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteWebhookParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteWebhookParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteWebhookParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteWebhookParams instance,
// as then you are sure you have configured all required params
func (s *WebhookService) NewDeleteWebhookParams(id string) *DeleteWebhookParams {
	p := &DeleteWebhookParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a Webhook
func (s *WebhookService) DeleteWebhook(p *DeleteWebhookParams) (*DeleteWebhookResponse, error) {
	resp, err := s.cs.newRequest("deleteWebhook", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteWebhookResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteWebhookResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteWebhookResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteWebhookResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteWebhookDeliveryParams struct {
	p map[string]interface{}
}

func (p *DeleteWebhookDeliveryParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	if v, found := p.p["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := p.p["webhookid"]; found {
		u.Set("webhookid", v.(string))
	}
	return u
}

func (p *DeleteWebhookDeliveryParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *DeleteWebhookDeliveryParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *DeleteWebhookDeliveryParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *DeleteWebhookDeliveryParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteWebhookDeliveryParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteWebhookDeliveryParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *DeleteWebhookDeliveryParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *DeleteWebhookDeliveryParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *DeleteWebhookDeliveryParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

func (p *DeleteWebhookDeliveryParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *DeleteWebhookDeliveryParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *DeleteWebhookDeliveryParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *DeleteWebhookDeliveryParams) SetWebhookid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["webhookid"] = v
}

func (p *DeleteWebhookDeliveryParams) ResetWebhookid() {
	if p.p != nil && p.p["webhookid"] != nil {
		delete(p.p, "webhookid")
	}
}

func (p *DeleteWebhookDeliveryParams) GetWebhookid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["webhookid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteWebhookDeliveryParams instance,
// as then you are sure you have configured all required params
func (s *WebhookService) NewDeleteWebhookDeliveryParams() *DeleteWebhookDeliveryParams {
	p := &DeleteWebhookDeliveryParams{}
	p.p = make(map[string]interface{})
	return p
}

// Deletes Webhook delivery
func (s *WebhookService) DeleteWebhookDelivery(p *DeleteWebhookDeliveryParams) (*DeleteWebhookDeliveryResponse, error) {
	resp, err := s.cs.newRequest("deleteWebhookDelivery", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteWebhookDeliveryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteWebhookDeliveryResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteWebhookDeliveryResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteWebhookDeliveryResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ExecuteWebhookDeliveryParams struct {
	p map[string]interface{}
}

func (p *ExecuteWebhookDeliveryParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["payload"]; found {
		u.Set("payload", v.(string))
	}
	if v, found := p.p["payloadurl"]; found {
		u.Set("payloadurl", v.(string))
	}
	if v, found := p.p["secretkey"]; found {
		u.Set("secretkey", v.(string))
	}
	if v, found := p.p["sslverification"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sslverification", vv)
	}
	if v, found := p.p["webhookid"]; found {
		u.Set("webhookid", v.(string))
	}
	return u
}

func (p *ExecuteWebhookDeliveryParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ExecuteWebhookDeliveryParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ExecuteWebhookDeliveryParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ExecuteWebhookDeliveryParams) SetPayload(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["payload"] = v
}

func (p *ExecuteWebhookDeliveryParams) ResetPayload() {
	if p.p != nil && p.p["payload"] != nil {
		delete(p.p, "payload")
	}
}

func (p *ExecuteWebhookDeliveryParams) GetPayload() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["payload"].(string)
	return value, ok
}

func (p *ExecuteWebhookDeliveryParams) SetPayloadurl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["payloadurl"] = v
}

func (p *ExecuteWebhookDeliveryParams) ResetPayloadurl() {
	if p.p != nil && p.p["payloadurl"] != nil {
		delete(p.p, "payloadurl")
	}
}

func (p *ExecuteWebhookDeliveryParams) GetPayloadurl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["payloadurl"].(string)
	return value, ok
}

func (p *ExecuteWebhookDeliveryParams) SetSecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secretkey"] = v
}

func (p *ExecuteWebhookDeliveryParams) ResetSecretkey() {
	if p.p != nil && p.p["secretkey"] != nil {
		delete(p.p, "secretkey")
	}
}

func (p *ExecuteWebhookDeliveryParams) GetSecretkey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["secretkey"].(string)
	return value, ok
}

func (p *ExecuteWebhookDeliveryParams) SetSslverification(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sslverification"] = v
}

func (p *ExecuteWebhookDeliveryParams) ResetSslverification() {
	if p.p != nil && p.p["sslverification"] != nil {
		delete(p.p, "sslverification")
	}
}

func (p *ExecuteWebhookDeliveryParams) GetSslverification() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sslverification"].(bool)
	return value, ok
}

func (p *ExecuteWebhookDeliveryParams) SetWebhookid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["webhookid"] = v
}

func (p *ExecuteWebhookDeliveryParams) ResetWebhookid() {
	if p.p != nil && p.p["webhookid"] != nil {
		delete(p.p, "webhookid")
	}
}

func (p *ExecuteWebhookDeliveryParams) GetWebhookid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["webhookid"].(string)
	return value, ok
}

// You should always use this function to get a new ExecuteWebhookDeliveryParams instance,
// as then you are sure you have configured all required params
func (s *WebhookService) NewExecuteWebhookDeliveryParams() *ExecuteWebhookDeliveryParams {
	p := &ExecuteWebhookDeliveryParams{}
	p.p = make(map[string]interface{})
	return p
}

// Executes a Webhook delivery
func (s *WebhookService) ExecuteWebhookDelivery(p *ExecuteWebhookDeliveryParams) (*ExecuteWebhookDeliveryResponse, error) {
	resp, err := s.cs.newRequest("executeWebhookDelivery", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExecuteWebhookDeliveryResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ExecuteWebhookDeliveryResponse struct {
	Enddate              string `json:"enddate"`
	Eventid              string `json:"eventid"`
	Eventtype            string `json:"eventtype"`
	Headers              string `json:"headers"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Managementserverid   UUID   `json:"managementserverid"`
	Managementservername string `json:"managementservername"`
	Payload              string `json:"payload"`
	Response             string `json:"response"`
	Startdate            string `json:"startdate"`
	Success              bool   `json:"success"`
	Webhookid            string `json:"webhookid"`
	Webhookname          string `json:"webhookname"`
}

func (r *ExecuteWebhookDeliveryResponse) UnmarshalJSON(b []byte) error {
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

	type alias ExecuteWebhookDeliveryResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListWebhookDeliveriesParams struct {
	p map[string]interface{}
}

func (p *ListWebhookDeliveriesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := p.p["eventtype"]; found {
		u.Set("eventtype", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
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
	if v, found := p.p["webhookid"]; found {
		u.Set("webhookid", v.(string))
	}
	return u
}

func (p *ListWebhookDeliveriesParams) SetEnddate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enddate"] = v
}

func (p *ListWebhookDeliveriesParams) ResetEnddate() {
	if p.p != nil && p.p["enddate"] != nil {
		delete(p.p, "enddate")
	}
}

func (p *ListWebhookDeliveriesParams) GetEnddate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enddate"].(string)
	return value, ok
}

func (p *ListWebhookDeliveriesParams) SetEventtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["eventtype"] = v
}

func (p *ListWebhookDeliveriesParams) ResetEventtype() {
	if p.p != nil && p.p["eventtype"] != nil {
		delete(p.p, "eventtype")
	}
}

func (p *ListWebhookDeliveriesParams) GetEventtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["eventtype"].(string)
	return value, ok
}

func (p *ListWebhookDeliveriesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListWebhookDeliveriesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListWebhookDeliveriesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListWebhookDeliveriesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListWebhookDeliveriesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListWebhookDeliveriesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListWebhookDeliveriesParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *ListWebhookDeliveriesParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *ListWebhookDeliveriesParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

func (p *ListWebhookDeliveriesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListWebhookDeliveriesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListWebhookDeliveriesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListWebhookDeliveriesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListWebhookDeliveriesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListWebhookDeliveriesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListWebhookDeliveriesParams) SetStartdate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startdate"] = v
}

func (p *ListWebhookDeliveriesParams) ResetStartdate() {
	if p.p != nil && p.p["startdate"] != nil {
		delete(p.p, "startdate")
	}
}

func (p *ListWebhookDeliveriesParams) GetStartdate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startdate"].(string)
	return value, ok
}

func (p *ListWebhookDeliveriesParams) SetWebhookid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["webhookid"] = v
}

func (p *ListWebhookDeliveriesParams) ResetWebhookid() {
	if p.p != nil && p.p["webhookid"] != nil {
		delete(p.p, "webhookid")
	}
}

func (p *ListWebhookDeliveriesParams) GetWebhookid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["webhookid"].(string)
	return value, ok
}

// You should always use this function to get a new ListWebhookDeliveriesParams instance,
// as then you are sure you have configured all required params
func (s *WebhookService) NewListWebhookDeliveriesParams() *ListWebhookDeliveriesParams {
	p := &ListWebhookDeliveriesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *WebhookService) GetWebhookDeliveryID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListWebhookDeliveriesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListWebhookDeliveries(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.WebhookDeliveries[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.WebhookDeliveries {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *WebhookService) GetWebhookDeliveryByName(name string, opts ...OptionFunc) (*WebhookDelivery, int, error) {
	id, count, err := s.GetWebhookDeliveryID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetWebhookDeliveryByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *WebhookService) GetWebhookDeliveryByID(id string, opts ...OptionFunc) (*WebhookDelivery, int, error) {
	p := &ListWebhookDeliveriesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListWebhookDeliveries(p)
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
		return l.WebhookDeliveries[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for WebhookDelivery UUID: %s!", id)
}

// Lists Webhook deliveries
func (s *WebhookService) ListWebhookDeliveries(p *ListWebhookDeliveriesParams) (*ListWebhookDeliveriesResponse, error) {
	resp, err := s.cs.newRequest("listWebhookDeliveries", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListWebhookDeliveriesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListWebhookDeliveriesResponse struct {
	Count             int                `json:"count"`
	WebhookDeliveries []*WebhookDelivery `json:"webhookdelivery"`
}

type WebhookDelivery struct {
	Account         string `json:"account"`
	Created         string `json:"created"`
	Description     string `json:"description"`
	Domain          string `json:"domain"`
	Domainid        string `json:"domainid"`
	Domainpath      string `json:"domainpath"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Name            string `json:"name"`
	Payloadurl      string `json:"payloadurl"`
	Project         string `json:"project"`
	Projectid       string `json:"projectid"`
	Scope           string `json:"scope"`
	Secretkey       string `json:"secretkey"`
	Sslverification bool   `json:"sslverification"`
	State           string `json:"state"`
}

type ListWebhooksParams struct {
	p map[string]interface{}
}

func (p *ListWebhooksParams) toURLValues() url.Values {
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
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *ListWebhooksParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListWebhooksParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListWebhooksParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListWebhooksParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListWebhooksParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListWebhooksParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListWebhooksParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListWebhooksParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListWebhooksParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListWebhooksParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListWebhooksParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListWebhooksParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListWebhooksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListWebhooksParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListWebhooksParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListWebhooksParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListWebhooksParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListWebhooksParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListWebhooksParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListWebhooksParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListWebhooksParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListWebhooksParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListWebhooksParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListWebhooksParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListWebhooksParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListWebhooksParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListWebhooksParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListWebhooksParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListWebhooksParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListWebhooksParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListWebhooksParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *ListWebhooksParams) ResetScope() {
	if p.p != nil && p.p["scope"] != nil {
		delete(p.p, "scope")
	}
}

func (p *ListWebhooksParams) GetScope() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scope"].(string)
	return value, ok
}

func (p *ListWebhooksParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListWebhooksParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListWebhooksParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

// You should always use this function to get a new ListWebhooksParams instance,
// as then you are sure you have configured all required params
func (s *WebhookService) NewListWebhooksParams() *ListWebhooksParams {
	p := &ListWebhooksParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *WebhookService) GetWebhookID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListWebhooksParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListWebhooks(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Webhooks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Webhooks {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *WebhookService) GetWebhookByName(name string, opts ...OptionFunc) (*Webhook, int, error) {
	id, count, err := s.GetWebhookID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetWebhookByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *WebhookService) GetWebhookByID(id string, opts ...OptionFunc) (*Webhook, int, error) {
	p := &ListWebhooksParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListWebhooks(p)
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
		return l.Webhooks[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Webhook UUID: %s!", id)
}

// Lists Webhooks
func (s *WebhookService) ListWebhooks(p *ListWebhooksParams) (*ListWebhooksResponse, error) {
	resp, err := s.cs.newRequest("listWebhooks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListWebhooksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListWebhooksResponse struct {
	Count    int        `json:"count"`
	Webhooks []*Webhook `json:"webhook"`
}

type Webhook struct {
	Account         string `json:"account"`
	Created         string `json:"created"`
	Description     string `json:"description"`
	Domain          string `json:"domain"`
	Domainid        string `json:"domainid"`
	Domainpath      string `json:"domainpath"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Name            string `json:"name"`
	Payloadurl      string `json:"payloadurl"`
	Project         string `json:"project"`
	Projectid       string `json:"projectid"`
	Scope           string `json:"scope"`
	Secretkey       string `json:"secretkey"`
	Sslverification bool   `json:"sslverification"`
	State           string `json:"state"`
}

type UpdateWebhookParams struct {
	p map[string]interface{}
}

func (p *UpdateWebhookParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["payloadurl"]; found {
		u.Set("payloadurl", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["secretkey"]; found {
		u.Set("secretkey", v.(string))
	}
	if v, found := p.p["sslverification"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sslverification", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *UpdateWebhookParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *UpdateWebhookParams) ResetDescription() {
	if p.p != nil && p.p["description"] != nil {
		delete(p.p, "description")
	}
}

func (p *UpdateWebhookParams) GetDescription() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["description"].(string)
	return value, ok
}

func (p *UpdateWebhookParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateWebhookParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateWebhookParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateWebhookParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateWebhookParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateWebhookParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateWebhookParams) SetPayloadurl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["payloadurl"] = v
}

func (p *UpdateWebhookParams) ResetPayloadurl() {
	if p.p != nil && p.p["payloadurl"] != nil {
		delete(p.p, "payloadurl")
	}
}

func (p *UpdateWebhookParams) GetPayloadurl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["payloadurl"].(string)
	return value, ok
}

func (p *UpdateWebhookParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *UpdateWebhookParams) ResetScope() {
	if p.p != nil && p.p["scope"] != nil {
		delete(p.p, "scope")
	}
}

func (p *UpdateWebhookParams) GetScope() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scope"].(string)
	return value, ok
}

func (p *UpdateWebhookParams) SetSecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["secretkey"] = v
}

func (p *UpdateWebhookParams) ResetSecretkey() {
	if p.p != nil && p.p["secretkey"] != nil {
		delete(p.p, "secretkey")
	}
}

func (p *UpdateWebhookParams) GetSecretkey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["secretkey"].(string)
	return value, ok
}

func (p *UpdateWebhookParams) SetSslverification(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sslverification"] = v
}

func (p *UpdateWebhookParams) ResetSslverification() {
	if p.p != nil && p.p["sslverification"] != nil {
		delete(p.p, "sslverification")
	}
}

func (p *UpdateWebhookParams) GetSslverification() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sslverification"].(bool)
	return value, ok
}

func (p *UpdateWebhookParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdateWebhookParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *UpdateWebhookParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateWebhookParams instance,
// as then you are sure you have configured all required params
func (s *WebhookService) NewUpdateWebhookParams(id string) *UpdateWebhookParams {
	p := &UpdateWebhookParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a Webhook
func (s *WebhookService) UpdateWebhook(p *UpdateWebhookParams) (*UpdateWebhookResponse, error) {
	resp, err := s.cs.newRequest("updateWebhook", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateWebhookResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateWebhookResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateWebhookResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateWebhookResponse
	return json.Unmarshal(b, (*alias)(r))
}
