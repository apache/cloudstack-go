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

type AutoScaleServiceIface interface {
	CreateAutoScalePolicy(p *CreateAutoScalePolicyParams) (*CreateAutoScalePolicyResponse, error)
	NewCreateAutoScalePolicyParams(action string, conditionids []string, duration int) *CreateAutoScalePolicyParams
	CreateAutoScaleVmGroup(p *CreateAutoScaleVmGroupParams) (*CreateAutoScaleVmGroupResponse, error)
	NewCreateAutoScaleVmGroupParams(lbruleid string, maxmembers int, minmembers int, scaledownpolicyids []string, scaleuppolicyids []string, vmprofileid string) *CreateAutoScaleVmGroupParams
	CreateAutoScaleVmProfile(p *CreateAutoScaleVmProfileParams) (*CreateAutoScaleVmProfileResponse, error)
	NewCreateAutoScaleVmProfileParams(serviceofferingid string, templateid string, zoneid string) *CreateAutoScaleVmProfileParams
	CreateCondition(p *CreateConditionParams) (*CreateConditionResponse, error)
	NewCreateConditionParams(counterid string, relationaloperator string, threshold int64) *CreateConditionParams
	CreateCounter(p *CreateCounterParams) (*CreateCounterResponse, error)
	NewCreateCounterParams(name string, provider string, source string, value string) *CreateCounterParams
	DeleteAutoScalePolicy(p *DeleteAutoScalePolicyParams) (*DeleteAutoScalePolicyResponse, error)
	NewDeleteAutoScalePolicyParams(id string) *DeleteAutoScalePolicyParams
	DeleteAutoScaleVmGroup(p *DeleteAutoScaleVmGroupParams) (*DeleteAutoScaleVmGroupResponse, error)
	NewDeleteAutoScaleVmGroupParams(id string) *DeleteAutoScaleVmGroupParams
	DeleteAutoScaleVmProfile(p *DeleteAutoScaleVmProfileParams) (*DeleteAutoScaleVmProfileResponse, error)
	NewDeleteAutoScaleVmProfileParams(id string) *DeleteAutoScaleVmProfileParams
	DeleteCondition(p *DeleteConditionParams) (*DeleteConditionResponse, error)
	NewDeleteConditionParams(id string) *DeleteConditionParams
	DeleteCounter(p *DeleteCounterParams) (*DeleteCounterResponse, error)
	NewDeleteCounterParams(id string) *DeleteCounterParams
	DisableAutoScaleVmGroup(p *DisableAutoScaleVmGroupParams) (*DisableAutoScaleVmGroupResponse, error)
	NewDisableAutoScaleVmGroupParams(id string) *DisableAutoScaleVmGroupParams
	EnableAutoScaleVmGroup(p *EnableAutoScaleVmGroupParams) (*EnableAutoScaleVmGroupResponse, error)
	NewEnableAutoScaleVmGroupParams(id string) *EnableAutoScaleVmGroupParams
	ListAutoScalePolicies(p *ListAutoScalePoliciesParams) (*ListAutoScalePoliciesResponse, error)
	NewListAutoScalePoliciesParams() *ListAutoScalePoliciesParams
	GetAutoScalePolicyID(name string, opts ...OptionFunc) (string, int, error)
	GetAutoScalePolicyByName(name string, opts ...OptionFunc) (*AutoScalePolicy, int, error)
	GetAutoScalePolicyByID(id string, opts ...OptionFunc) (*AutoScalePolicy, int, error)
	ListAutoScaleVmGroups(p *ListAutoScaleVmGroupsParams) (*ListAutoScaleVmGroupsResponse, error)
	NewListAutoScaleVmGroupsParams() *ListAutoScaleVmGroupsParams
	GetAutoScaleVmGroupID(name string, opts ...OptionFunc) (string, int, error)
	GetAutoScaleVmGroupByName(name string, opts ...OptionFunc) (*AutoScaleVmGroup, int, error)
	GetAutoScaleVmGroupByID(id string, opts ...OptionFunc) (*AutoScaleVmGroup, int, error)
	ListAutoScaleVmProfiles(p *ListAutoScaleVmProfilesParams) (*ListAutoScaleVmProfilesResponse, error)
	NewListAutoScaleVmProfilesParams() *ListAutoScaleVmProfilesParams
	GetAutoScaleVmProfileByID(id string, opts ...OptionFunc) (*AutoScaleVmProfile, int, error)
	ListConditions(p *ListConditionsParams) (*ListConditionsResponse, error)
	NewListConditionsParams() *ListConditionsParams
	GetConditionByID(id string, opts ...OptionFunc) (*Condition, int, error)
	ListCounters(p *ListCountersParams) (*ListCountersResponse, error)
	NewListCountersParams() *ListCountersParams
	GetCounterID(name string, opts ...OptionFunc) (string, int, error)
	GetCounterByName(name string, opts ...OptionFunc) (*Counter, int, error)
	GetCounterByID(id string, opts ...OptionFunc) (*Counter, int, error)
	UpdateAutoScalePolicy(p *UpdateAutoScalePolicyParams) (*UpdateAutoScalePolicyResponse, error)
	NewUpdateAutoScalePolicyParams(id string) *UpdateAutoScalePolicyParams
	UpdateAutoScaleVmGroup(p *UpdateAutoScaleVmGroupParams) (*UpdateAutoScaleVmGroupResponse, error)
	NewUpdateAutoScaleVmGroupParams(id string) *UpdateAutoScaleVmGroupParams
	UpdateAutoScaleVmProfile(p *UpdateAutoScaleVmProfileParams) (*UpdateAutoScaleVmProfileResponse, error)
	NewUpdateAutoScaleVmProfileParams(id string) *UpdateAutoScaleVmProfileParams
	UpdateCondition(p *UpdateConditionParams) (*UpdateConditionResponse, error)
	NewUpdateConditionParams(id string, relationaloperator string, threshold int64) *UpdateConditionParams
}

type CreateAutoScalePolicyParams struct {
	p map[string]interface{}
}

func (p *CreateAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["conditionids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("conditionids", vv)
	}
	if v, found := p.p["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["quiettime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("quiettime", vv)
	}
	return u
}

func (p *CreateAutoScalePolicyParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *CreateAutoScalePolicyParams) ResetAction() {
	if p.p != nil && p.p["action"] != nil {
		delete(p.p, "action")
	}
}

func (p *CreateAutoScalePolicyParams) GetAction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["action"].(string)
	return value, ok
}

func (p *CreateAutoScalePolicyParams) SetConditionids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["conditionids"] = v
}

func (p *CreateAutoScalePolicyParams) ResetConditionids() {
	if p.p != nil && p.p["conditionids"] != nil {
		delete(p.p, "conditionids")
	}
}

func (p *CreateAutoScalePolicyParams) GetConditionids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["conditionids"].([]string)
	return value, ok
}

func (p *CreateAutoScalePolicyParams) SetDuration(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["duration"] = v
}

func (p *CreateAutoScalePolicyParams) ResetDuration() {
	if p.p != nil && p.p["duration"] != nil {
		delete(p.p, "duration")
	}
}

func (p *CreateAutoScalePolicyParams) GetDuration() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["duration"].(int)
	return value, ok
}

func (p *CreateAutoScalePolicyParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateAutoScalePolicyParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateAutoScalePolicyParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateAutoScalePolicyParams) SetQuiettime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiettime"] = v
}

func (p *CreateAutoScalePolicyParams) ResetQuiettime() {
	if p.p != nil && p.p["quiettime"] != nil {
		delete(p.p, "quiettime")
	}
}

func (p *CreateAutoScalePolicyParams) GetQuiettime() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quiettime"].(int)
	return value, ok
}

// You should always use this function to get a new CreateAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScalePolicyParams(action string, conditionids []string, duration int) *CreateAutoScalePolicyParams {
	p := &CreateAutoScalePolicyParams{}
	p.p = make(map[string]interface{})
	p.p["action"] = action
	p.p["conditionids"] = conditionids
	p.p["duration"] = duration
	return p
}

// Creates an autoscale policy for a provision or deprovision action, the action is taken when the all the conditions evaluates to true for the specified duration. The policy is in effect once it is attached to a autscale vm group.
func (s *AutoScaleService) CreateAutoScalePolicy(p *CreateAutoScalePolicyParams) (*CreateAutoScalePolicyResponse, error) {
	resp, err := s.cs.newPostRequest("createAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScalePolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateAutoScalePolicyResponse struct {
	Account    string       `json:"account"`
	Action     string       `json:"action"`
	Conditions []*Condition `json:"conditions"`
	Domain     string       `json:"domain"`
	Domainid   string       `json:"domainid"`
	Domainpath string       `json:"domainpath"`
	Duration   int          `json:"duration"`
	Id         string       `json:"id"`
	JobID      string       `json:"jobid"`
	Jobstatus  int          `json:"jobstatus"`
	Name       string       `json:"name"`
	Project    string       `json:"project"`
	Projectid  string       `json:"projectid"`
	Quiettime  int          `json:"quiettime"`
}

type CreateAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *CreateAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["maxmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxmembers", vv)
	}
	if v, found := p.p["minmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmembers", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["scaledownpolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaledownpolicyids", vv)
	}
	if v, found := p.p["scaleuppolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaleuppolicyids", vv)
	}
	if v, found := p.p["vmprofileid"]; found {
		u.Set("vmprofileid", v.(string))
	}
	return u
}

func (p *CreateAutoScaleVmGroupParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateAutoScaleVmGroupParams) SetInterval(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["interval"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetInterval() {
	if p.p != nil && p.p["interval"] != nil {
		delete(p.p, "interval")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetInterval() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["interval"].(int)
	return value, ok
}

func (p *CreateAutoScaleVmGroupParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmGroupParams) SetMaxmembers(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxmembers"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetMaxmembers() {
	if p.p != nil && p.p["maxmembers"] != nil {
		delete(p.p, "maxmembers")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetMaxmembers() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxmembers"].(int)
	return value, ok
}

func (p *CreateAutoScaleVmGroupParams) SetMinmembers(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minmembers"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetMinmembers() {
	if p.p != nil && p.p["minmembers"] != nil {
		delete(p.p, "minmembers")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetMinmembers() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["minmembers"].(int)
	return value, ok
}

func (p *CreateAutoScaleVmGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmGroupParams) SetScaledownpolicyids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scaledownpolicyids"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetScaledownpolicyids() {
	if p.p != nil && p.p["scaledownpolicyids"] != nil {
		delete(p.p, "scaledownpolicyids")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetScaledownpolicyids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scaledownpolicyids"].([]string)
	return value, ok
}

func (p *CreateAutoScaleVmGroupParams) SetScaleuppolicyids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scaleuppolicyids"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetScaleuppolicyids() {
	if p.p != nil && p.p["scaleuppolicyids"] != nil {
		delete(p.p, "scaleuppolicyids")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetScaleuppolicyids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scaleuppolicyids"].([]string)
	return value, ok
}

func (p *CreateAutoScaleVmGroupParams) SetVmprofileid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmprofileid"] = v
}

func (p *CreateAutoScaleVmGroupParams) ResetVmprofileid() {
	if p.p != nil && p.p["vmprofileid"] != nil {
		delete(p.p, "vmprofileid")
	}
}

func (p *CreateAutoScaleVmGroupParams) GetVmprofileid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmprofileid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScaleVmGroupParams(lbruleid string, maxmembers int, minmembers int, scaledownpolicyids []string, scaleuppolicyids []string, vmprofileid string) *CreateAutoScaleVmGroupParams {
	p := &CreateAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	p.p["maxmembers"] = maxmembers
	p.p["minmembers"] = minmembers
	p.p["scaledownpolicyids"] = scaledownpolicyids
	p.p["scaleuppolicyids"] = scaleuppolicyids
	p.p["vmprofileid"] = vmprofileid
	return p
}

// Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
func (s *AutoScaleService) CreateAutoScaleVmGroup(p *CreateAutoScaleVmGroupParams) (*CreateAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newPostRequest("createAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScaleVmGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateAutoScaleVmGroupResponse struct {
	Account                      string             `json:"account"`
	Associatednetworkid          string             `json:"associatednetworkid"`
	Associatednetworkname        string             `json:"associatednetworkname"`
	Availablevirtualmachinecount int                `json:"availablevirtualmachinecount"`
	Created                      string             `json:"created"`
	Domain                       string             `json:"domain"`
	Domainid                     string             `json:"domainid"`
	Domainpath                   string             `json:"domainpath"`
	Fordisplay                   bool               `json:"fordisplay"`
	Hasannotations               bool               `json:"hasannotations"`
	Id                           string             `json:"id"`
	Interval                     int                `json:"interval"`
	JobID                        string             `json:"jobid"`
	Jobstatus                    int                `json:"jobstatus"`
	Lbprovider                   string             `json:"lbprovider"`
	Lbruleid                     string             `json:"lbruleid"`
	Maxmembers                   int                `json:"maxmembers"`
	Minmembers                   int                `json:"minmembers"`
	Name                         string             `json:"name"`
	Privateport                  string             `json:"privateport"`
	Project                      string             `json:"project"`
	Projectid                    string             `json:"projectid"`
	Publicip                     string             `json:"publicip"`
	Publicipid                   string             `json:"publicipid"`
	Publicport                   string             `json:"publicport"`
	Scaledownpolicies            []*AutoScalePolicy `json:"scaledownpolicies"`
	Scaleuppolicies              []*AutoScalePolicy `json:"scaleuppolicies"`
	State                        string             `json:"state"`
	Vmprofileid                  string             `json:"vmprofileid"`
}

type CreateAutoScaleVmProfileParams struct {
	p map[string]interface{}
}

func (p *CreateAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["autoscaleuserid"]; found {
		u.Set("autoscaleuserid", v.(string))
	}
	if v, found := p.p["counterparam"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("counterparam[%d].key", i), k)
			u.Set(fmt.Sprintf("counterparam[%d].value", i), m[k])
		}
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["expungevmgraceperiod"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("expungevmgraceperiod", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["otherdeployparams"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("otherdeployparams[%d].name", i), k)
			u.Set(fmt.Sprintf("otherdeployparams[%d].value", i), m[k])
		}
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	if v, found := p.p["userdatadetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("userdatadetails[%d].key", i), k)
			u.Set(fmt.Sprintf("userdatadetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["userdataid"]; found {
		u.Set("userdataid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateAutoScaleVmProfileParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetAutoscaleuserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoscaleuserid"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetAutoscaleuserid() {
	if p.p != nil && p.p["autoscaleuserid"] != nil {
		delete(p.p, "autoscaleuserid")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetAutoscaleuserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoscaleuserid"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetCounterparam(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["counterparam"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetCounterparam() {
	if p.p != nil && p.p["counterparam"] != nil {
		delete(p.p, "counterparam")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetCounterparam() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["counterparam"].(map[string]string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetExpungevmgraceperiod(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expungevmgraceperiod"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetExpungevmgraceperiod() {
	if p.p != nil && p.p["expungevmgraceperiod"] != nil {
		delete(p.p, "expungevmgraceperiod")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetExpungevmgraceperiod() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["expungevmgraceperiod"].(int)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetOtherdeployparams(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["otherdeployparams"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetOtherdeployparams() {
	if p.p != nil && p.p["otherdeployparams"] != nil {
		delete(p.p, "otherdeployparams")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetOtherdeployparams() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["otherdeployparams"].(map[string]string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetUserdata() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetUserdatadetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdatadetails"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetUserdatadetails() {
	if p.p != nil && p.p["userdatadetails"] != nil {
		delete(p.p, "userdatadetails")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetUserdatadetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdatadetails"].(map[string]string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetUserdataid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdataid"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetUserdataid() {
	if p.p != nil && p.p["userdataid"] != nil {
		delete(p.p, "userdataid")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetUserdataid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdataid"].(string)
	return value, ok
}

func (p *CreateAutoScaleVmProfileParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateAutoScaleVmProfileParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateAutoScaleVmProfileParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScaleVmProfileParams(serviceofferingid string, templateid string, zoneid string) *CreateAutoScaleVmProfileParams {
	p := &CreateAutoScaleVmProfileParams{}
	p.p = make(map[string]interface{})
	p.p["serviceofferingid"] = serviceofferingid
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	return p
}

// Creates a profile that contains information about the virtual machine which will be provisioned automatically by autoscale feature.
func (s *AutoScaleService) CreateAutoScaleVmProfile(p *CreateAutoScaleVmProfileParams) (*CreateAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newPostRequest("createAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScaleVmProfileResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateAutoScaleVmProfileResponse struct {
	Account              string            `json:"account"`
	Autoscaleuserid      string            `json:"autoscaleuserid"`
	Domain               string            `json:"domain"`
	Domainid             string            `json:"domainid"`
	Domainpath           string            `json:"domainpath"`
	Expungevmgraceperiod int               `json:"expungevmgraceperiod"`
	Fordisplay           bool              `json:"fordisplay"`
	Id                   string            `json:"id"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Otherdeployparams    map[string]string `json:"otherdeployparams"`
	Project              string            `json:"project"`
	Projectid            string            `json:"projectid"`
	Serviceofferingid    string            `json:"serviceofferingid"`
	Templateid           string            `json:"templateid"`
	Userdata             string            `json:"userdata"`
	Userdatadetails      string            `json:"userdatadetails"`
	Userdataid           string            `json:"userdataid"`
	Userdataname         string            `json:"userdataname"`
	Userdatapolicy       string            `json:"userdatapolicy"`
	Zoneid               string            `json:"zoneid"`
}

type CreateConditionParams struct {
	p map[string]interface{}
}

func (p *CreateConditionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["counterid"]; found {
		u.Set("counterid", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["relationaloperator"]; found {
		u.Set("relationaloperator", v.(string))
	}
	if v, found := p.p["threshold"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("threshold", vv)
	}
	return u
}

func (p *CreateConditionParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *CreateConditionParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *CreateConditionParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *CreateConditionParams) SetCounterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["counterid"] = v
}

func (p *CreateConditionParams) ResetCounterid() {
	if p.p != nil && p.p["counterid"] != nil {
		delete(p.p, "counterid")
	}
}

func (p *CreateConditionParams) GetCounterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["counterid"].(string)
	return value, ok
}

func (p *CreateConditionParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateConditionParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateConditionParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *CreateConditionParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *CreateConditionParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *CreateConditionParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *CreateConditionParams) SetRelationaloperator(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["relationaloperator"] = v
}

func (p *CreateConditionParams) ResetRelationaloperator() {
	if p.p != nil && p.p["relationaloperator"] != nil {
		delete(p.p, "relationaloperator")
	}
}

func (p *CreateConditionParams) GetRelationaloperator() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["relationaloperator"].(string)
	return value, ok
}

func (p *CreateConditionParams) SetThreshold(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["threshold"] = v
}

func (p *CreateConditionParams) ResetThreshold() {
	if p.p != nil && p.p["threshold"] != nil {
		delete(p.p, "threshold")
	}
}

func (p *CreateConditionParams) GetThreshold() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["threshold"].(int64)
	return value, ok
}

// You should always use this function to get a new CreateConditionParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateConditionParams(counterid string, relationaloperator string, threshold int64) *CreateConditionParams {
	p := &CreateConditionParams{}
	p.p = make(map[string]interface{})
	p.p["counterid"] = counterid
	p.p["relationaloperator"] = relationaloperator
	p.p["threshold"] = threshold
	return p
}

// Creates a condition for VM auto scaling
func (s *AutoScaleService) CreateCondition(p *CreateConditionParams) (*CreateConditionResponse, error) {
	resp, err := s.cs.newPostRequest("createCondition", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateConditionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateConditionResponse struct {
	Account            string   `json:"account"`
	Counter            *Counter `json:"counter"`
	Counterid          string   `json:"counterid"`
	Countername        string   `json:"countername"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	JobID              string   `json:"jobid"`
	Jobstatus          int      `json:"jobstatus"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Relationaloperator string   `json:"relationaloperator"`
	Threshold          int64    `json:"threshold"`
	Zoneid             string   `json:"zoneid"`
}

type CreateCounterParams struct {
	p map[string]interface{}
}

func (p *CreateCounterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["source"]; found {
		u.Set("source", v.(string))
	}
	if v, found := p.p["value"]; found {
		u.Set("value", v.(string))
	}
	return u
}

func (p *CreateCounterParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateCounterParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateCounterParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateCounterParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *CreateCounterParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *CreateCounterParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *CreateCounterParams) SetSource(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["source"] = v
}

func (p *CreateCounterParams) ResetSource() {
	if p.p != nil && p.p["source"] != nil {
		delete(p.p, "source")
	}
}

func (p *CreateCounterParams) GetSource() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["source"].(string)
	return value, ok
}

func (p *CreateCounterParams) SetValue(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["value"] = v
}

func (p *CreateCounterParams) ResetValue() {
	if p.p != nil && p.p["value"] != nil {
		delete(p.p, "value")
	}
}

func (p *CreateCounterParams) GetValue() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["value"].(string)
	return value, ok
}

// You should always use this function to get a new CreateCounterParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateCounterParams(name string, provider string, source string, value string) *CreateCounterParams {
	p := &CreateCounterParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["provider"] = provider
	p.p["source"] = source
	p.p["value"] = value
	return p
}

// Adds metric counter for VM auto scaling
func (s *AutoScaleService) CreateCounter(p *CreateCounterParams) (*CreateCounterResponse, error) {
	resp, err := s.cs.newPostRequest("createCounter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateCounterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateCounterResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Provider  string `json:"provider"`
	Source    string `json:"source"`
	Value     string `json:"value"`
	Zoneid    string `json:"zoneid"`
}

type DeleteAutoScalePolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteAutoScalePolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteAutoScalePolicyParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteAutoScalePolicyParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScalePolicyParams(id string) *DeleteAutoScalePolicyParams {
	p := &DeleteAutoScalePolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a autoscale policy.
func (s *AutoScaleService) DeleteAutoScalePolicy(p *DeleteAutoScalePolicyParams) (*DeleteAutoScalePolicyResponse, error) {
	resp, err := s.cs.newPostRequest("deleteAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScalePolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteAutoScalePolicyResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *DeleteAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteAutoScaleVmGroupParams) SetCleanup(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cleanup"] = v
}

func (p *DeleteAutoScaleVmGroupParams) ResetCleanup() {
	if p.p != nil && p.p["cleanup"] != nil {
		delete(p.p, "cleanup")
	}
}

func (p *DeleteAutoScaleVmGroupParams) GetCleanup() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cleanup"].(bool)
	return value, ok
}

func (p *DeleteAutoScaleVmGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteAutoScaleVmGroupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteAutoScaleVmGroupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScaleVmGroupParams(id string) *DeleteAutoScaleVmGroupParams {
	p := &DeleteAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a autoscale vm group.
func (s *AutoScaleService) DeleteAutoScaleVmGroup(p *DeleteAutoScaleVmGroupParams) (*DeleteAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newPostRequest("deleteAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScaleVmGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteAutoScaleVmGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteAutoScaleVmProfileParams struct {
	p map[string]interface{}
}

func (p *DeleteAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteAutoScaleVmProfileParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteAutoScaleVmProfileParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteAutoScaleVmProfileParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScaleVmProfileParams(id string) *DeleteAutoScaleVmProfileParams {
	p := &DeleteAutoScaleVmProfileParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a autoscale vm profile.
func (s *AutoScaleService) DeleteAutoScaleVmProfile(p *DeleteAutoScaleVmProfileParams) (*DeleteAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newPostRequest("deleteAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScaleVmProfileResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteAutoScaleVmProfileResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteConditionParams struct {
	p map[string]interface{}
}

func (p *DeleteConditionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteConditionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteConditionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteConditionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteConditionParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteConditionParams(id string) *DeleteConditionParams {
	p := &DeleteConditionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes a condition for VM auto scaling
func (s *AutoScaleService) DeleteCondition(p *DeleteConditionParams) (*DeleteConditionResponse, error) {
	resp, err := s.cs.newPostRequest("deleteCondition", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteConditionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteConditionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteCounterParams struct {
	p map[string]interface{}
}

func (p *DeleteCounterParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteCounterParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteCounterParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteCounterParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteCounterParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteCounterParams(id string) *DeleteCounterParams {
	p := &DeleteCounterParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a counter for VM auto scaling
func (s *AutoScaleService) DeleteCounter(p *DeleteCounterParams) (*DeleteCounterResponse, error) {
	resp, err := s.cs.newPostRequest("deleteCounter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteCounterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteCounterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DisableAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *DisableAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DisableAutoScaleVmGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DisableAutoScaleVmGroupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DisableAutoScaleVmGroupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DisableAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDisableAutoScaleVmGroupParams(id string) *DisableAutoScaleVmGroupParams {
	p := &DisableAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Disables an AutoScale Vm Group
func (s *AutoScaleService) DisableAutoScaleVmGroup(p *DisableAutoScaleVmGroupParams) (*DisableAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newPostRequest("disableAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableAutoScaleVmGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DisableAutoScaleVmGroupResponse struct {
	Account                      string             `json:"account"`
	Associatednetworkid          string             `json:"associatednetworkid"`
	Associatednetworkname        string             `json:"associatednetworkname"`
	Availablevirtualmachinecount int                `json:"availablevirtualmachinecount"`
	Created                      string             `json:"created"`
	Domain                       string             `json:"domain"`
	Domainid                     string             `json:"domainid"`
	Domainpath                   string             `json:"domainpath"`
	Fordisplay                   bool               `json:"fordisplay"`
	Hasannotations               bool               `json:"hasannotations"`
	Id                           string             `json:"id"`
	Interval                     int                `json:"interval"`
	JobID                        string             `json:"jobid"`
	Jobstatus                    int                `json:"jobstatus"`
	Lbprovider                   string             `json:"lbprovider"`
	Lbruleid                     string             `json:"lbruleid"`
	Maxmembers                   int                `json:"maxmembers"`
	Minmembers                   int                `json:"minmembers"`
	Name                         string             `json:"name"`
	Privateport                  string             `json:"privateport"`
	Project                      string             `json:"project"`
	Projectid                    string             `json:"projectid"`
	Publicip                     string             `json:"publicip"`
	Publicipid                   string             `json:"publicipid"`
	Publicport                   string             `json:"publicport"`
	Scaledownpolicies            []*AutoScalePolicy `json:"scaledownpolicies"`
	Scaleuppolicies              []*AutoScalePolicy `json:"scaleuppolicies"`
	State                        string             `json:"state"`
	Vmprofileid                  string             `json:"vmprofileid"`
}

type EnableAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *EnableAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableAutoScaleVmGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *EnableAutoScaleVmGroupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *EnableAutoScaleVmGroupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new EnableAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewEnableAutoScaleVmGroupParams(id string) *EnableAutoScaleVmGroupParams {
	p := &EnableAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Enables an AutoScale Vm Group
func (s *AutoScaleService) EnableAutoScaleVmGroup(p *EnableAutoScaleVmGroupParams) (*EnableAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newPostRequest("enableAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableAutoScaleVmGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type EnableAutoScaleVmGroupResponse struct {
	Account                      string             `json:"account"`
	Associatednetworkid          string             `json:"associatednetworkid"`
	Associatednetworkname        string             `json:"associatednetworkname"`
	Availablevirtualmachinecount int                `json:"availablevirtualmachinecount"`
	Created                      string             `json:"created"`
	Domain                       string             `json:"domain"`
	Domainid                     string             `json:"domainid"`
	Domainpath                   string             `json:"domainpath"`
	Fordisplay                   bool               `json:"fordisplay"`
	Hasannotations               bool               `json:"hasannotations"`
	Id                           string             `json:"id"`
	Interval                     int                `json:"interval"`
	JobID                        string             `json:"jobid"`
	Jobstatus                    int                `json:"jobstatus"`
	Lbprovider                   string             `json:"lbprovider"`
	Lbruleid                     string             `json:"lbruleid"`
	Maxmembers                   int                `json:"maxmembers"`
	Minmembers                   int                `json:"minmembers"`
	Name                         string             `json:"name"`
	Privateport                  string             `json:"privateport"`
	Project                      string             `json:"project"`
	Projectid                    string             `json:"projectid"`
	Publicip                     string             `json:"publicip"`
	Publicipid                   string             `json:"publicipid"`
	Publicport                   string             `json:"publicport"`
	Scaledownpolicies            []*AutoScalePolicy `json:"scaledownpolicies"`
	Scaleuppolicies              []*AutoScalePolicy `json:"scaleuppolicies"`
	State                        string             `json:"state"`
	Vmprofileid                  string             `json:"vmprofileid"`
}

type ListAutoScalePoliciesParams struct {
	p map[string]interface{}
}

func (p *ListAutoScalePoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["conditionid"]; found {
		u.Set("conditionid", v.(string))
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
	if v, found := p.p["vmgroupid"]; found {
		u.Set("vmgroupid", v.(string))
	}
	return u
}

func (p *ListAutoScalePoliciesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListAutoScalePoliciesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListAutoScalePoliciesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *ListAutoScalePoliciesParams) ResetAction() {
	if p.p != nil && p.p["action"] != nil {
		delete(p.p, "action")
	}
}

func (p *ListAutoScalePoliciesParams) GetAction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["action"].(string)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetConditionid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["conditionid"] = v
}

func (p *ListAutoScalePoliciesParams) ResetConditionid() {
	if p.p != nil && p.p["conditionid"] != nil {
		delete(p.p, "conditionid")
	}
}

func (p *ListAutoScalePoliciesParams) GetConditionid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["conditionid"].(string)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListAutoScalePoliciesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListAutoScalePoliciesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListAutoScalePoliciesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListAutoScalePoliciesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListAutoScalePoliciesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListAutoScalePoliciesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAutoScalePoliciesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListAutoScalePoliciesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListAutoScalePoliciesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListAutoScalePoliciesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListAutoScalePoliciesParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListAutoScalePoliciesParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAutoScalePoliciesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListAutoScalePoliciesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAutoScalePoliciesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListAutoScalePoliciesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListAutoScalePoliciesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListAutoScalePoliciesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListAutoScalePoliciesParams) SetVmgroupid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmgroupid"] = v
}

func (p *ListAutoScalePoliciesParams) ResetVmgroupid() {
	if p.p != nil && p.p["vmgroupid"] != nil {
		delete(p.p, "vmgroupid")
	}
}

func (p *ListAutoScalePoliciesParams) GetVmgroupid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmgroupid"].(string)
	return value, ok
}

// You should always use this function to get a new ListAutoScalePoliciesParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScalePoliciesParams() *ListAutoScalePoliciesParams {
	p := &ListAutoScalePoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScalePolicyID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListAutoScalePoliciesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListAutoScalePolicies(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.AutoScalePolicies[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.AutoScalePolicies {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScalePolicyByName(name string, opts ...OptionFunc) (*AutoScalePolicy, int, error) {
	id, count, err := s.GetAutoScalePolicyID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetAutoScalePolicyByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScalePolicyByID(id string, opts ...OptionFunc) (*AutoScalePolicy, int, error) {
	p := &ListAutoScalePoliciesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScalePolicies(p)
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
		return l.AutoScalePolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScalePolicy UUID: %s!", id)
}

// Lists autoscale policies.
func (s *AutoScaleService) ListAutoScalePolicies(p *ListAutoScalePoliciesParams) (*ListAutoScalePoliciesResponse, error) {
	resp, err := s.cs.newRequest("listAutoScalePolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScalePoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScalePoliciesResponse struct {
	Count             int                `json:"count"`
	AutoScalePolicies []*AutoScalePolicy `json:"autoscalepolicy"`
}

type AutoScalePolicy struct {
	Account    string       `json:"account"`
	Action     string       `json:"action"`
	Conditions []*Condition `json:"conditions"`
	Domain     string       `json:"domain"`
	Domainid   string       `json:"domainid"`
	Domainpath string       `json:"domainpath"`
	Duration   int          `json:"duration"`
	Id         string       `json:"id"`
	JobID      string       `json:"jobid"`
	Jobstatus  int          `json:"jobstatus"`
	Name       string       `json:"name"`
	Project    string       `json:"project"`
	Projectid  string       `json:"projectid"`
	Quiettime  int          `json:"quiettime"`
}

type ListAutoScaleVmGroupsParams struct {
	p map[string]interface{}
}

func (p *ListAutoScaleVmGroupsParams) toURLValues() url.Values {
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
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
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
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
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
	if v, found := p.p["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["vmprofileid"]; found {
		u.Set("vmprofileid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListAutoScaleVmGroupsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetPolicyid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyid"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetPolicyid() {
	if p.p != nil && p.p["policyid"] != nil {
		delete(p.p, "policyid")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetPolicyid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetVmprofileid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmprofileid"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetVmprofileid() {
	if p.p != nil && p.p["vmprofileid"] != nil {
		delete(p.p, "vmprofileid")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetVmprofileid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmprofileid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmGroupsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListAutoScaleVmGroupsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListAutoScaleVmGroupsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListAutoScaleVmGroupsParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScaleVmGroupsParams() *ListAutoScaleVmGroupsParams {
	p := &ListAutoScaleVmGroupsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScaleVmGroupID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListAutoScaleVmGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListAutoScaleVmGroups(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.AutoScaleVmGroups[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.AutoScaleVmGroups {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScaleVmGroupByName(name string, opts ...OptionFunc) (*AutoScaleVmGroup, int, error) {
	id, count, err := s.GetAutoScaleVmGroupID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetAutoScaleVmGroupByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScaleVmGroupByID(id string, opts ...OptionFunc) (*AutoScaleVmGroup, int, error) {
	p := &ListAutoScaleVmGroupsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScaleVmGroups(p)
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
		return l.AutoScaleVmGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScaleVmGroup UUID: %s!", id)
}

// Lists autoscale vm groups.
func (s *AutoScaleService) ListAutoScaleVmGroups(p *ListAutoScaleVmGroupsParams) (*ListAutoScaleVmGroupsResponse, error) {
	resp, err := s.cs.newRequest("listAutoScaleVmGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScaleVmGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScaleVmGroupsResponse struct {
	Count             int                 `json:"count"`
	AutoScaleVmGroups []*AutoScaleVmGroup `json:"autoscalevmgroup"`
}

type AutoScaleVmGroup struct {
	Account                      string             `json:"account"`
	Associatednetworkid          string             `json:"associatednetworkid"`
	Associatednetworkname        string             `json:"associatednetworkname"`
	Availablevirtualmachinecount int                `json:"availablevirtualmachinecount"`
	Created                      string             `json:"created"`
	Domain                       string             `json:"domain"`
	Domainid                     string             `json:"domainid"`
	Domainpath                   string             `json:"domainpath"`
	Fordisplay                   bool               `json:"fordisplay"`
	Hasannotations               bool               `json:"hasannotations"`
	Id                           string             `json:"id"`
	Interval                     int                `json:"interval"`
	JobID                        string             `json:"jobid"`
	Jobstatus                    int                `json:"jobstatus"`
	Lbprovider                   string             `json:"lbprovider"`
	Lbruleid                     string             `json:"lbruleid"`
	Maxmembers                   int                `json:"maxmembers"`
	Minmembers                   int                `json:"minmembers"`
	Name                         string             `json:"name"`
	Privateport                  string             `json:"privateport"`
	Project                      string             `json:"project"`
	Projectid                    string             `json:"projectid"`
	Publicip                     string             `json:"publicip"`
	Publicipid                   string             `json:"publicipid"`
	Publicport                   string             `json:"publicport"`
	Scaledownpolicies            []*AutoScalePolicy `json:"scaledownpolicies"`
	Scaleuppolicies              []*AutoScalePolicy `json:"scaleuppolicies"`
	State                        string             `json:"state"`
	Vmprofileid                  string             `json:"vmprofileid"`
}

type ListAutoScaleVmProfilesParams struct {
	p map[string]interface{}
}

func (p *ListAutoScaleVmProfilesParams) toURLValues() url.Values {
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
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
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
	if v, found := p.p["otherdeployparams"]; found {
		u.Set("otherdeployparams", v.(string))
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
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListAutoScaleVmProfilesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetOtherdeployparams(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["otherdeployparams"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetOtherdeployparams() {
	if p.p != nil && p.p["otherdeployparams"] != nil {
		delete(p.p, "otherdeployparams")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetOtherdeployparams() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["otherdeployparams"].(string)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *ListAutoScaleVmProfilesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListAutoScaleVmProfilesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListAutoScaleVmProfilesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListAutoScaleVmProfilesParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScaleVmProfilesParams() *ListAutoScaleVmProfilesParams {
	p := &ListAutoScaleVmProfilesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScaleVmProfileByID(id string, opts ...OptionFunc) (*AutoScaleVmProfile, int, error) {
	p := &ListAutoScaleVmProfilesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScaleVmProfiles(p)
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
		return l.AutoScaleVmProfiles[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScaleVmProfile UUID: %s!", id)
}

// Lists autoscale vm profiles.
func (s *AutoScaleService) ListAutoScaleVmProfiles(p *ListAutoScaleVmProfilesParams) (*ListAutoScaleVmProfilesResponse, error) {
	resp, err := s.cs.newRequest("listAutoScaleVmProfiles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScaleVmProfilesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScaleVmProfilesResponse struct {
	Count               int                   `json:"count"`
	AutoScaleVmProfiles []*AutoScaleVmProfile `json:"autoscalevmprofile"`
}

type AutoScaleVmProfile struct {
	Account              string            `json:"account"`
	Autoscaleuserid      string            `json:"autoscaleuserid"`
	Domain               string            `json:"domain"`
	Domainid             string            `json:"domainid"`
	Domainpath           string            `json:"domainpath"`
	Expungevmgraceperiod int               `json:"expungevmgraceperiod"`
	Fordisplay           bool              `json:"fordisplay"`
	Id                   string            `json:"id"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Otherdeployparams    map[string]string `json:"otherdeployparams"`
	Project              string            `json:"project"`
	Projectid            string            `json:"projectid"`
	Serviceofferingid    string            `json:"serviceofferingid"`
	Templateid           string            `json:"templateid"`
	Userdata             string            `json:"userdata"`
	Userdatadetails      string            `json:"userdatadetails"`
	Userdataid           string            `json:"userdataid"`
	Userdataname         string            `json:"userdataname"`
	Userdatapolicy       string            `json:"userdatapolicy"`
	Zoneid               string            `json:"zoneid"`
}

type ListConditionsParams struct {
	p map[string]interface{}
}

func (p *ListConditionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["counterid"]; found {
		u.Set("counterid", v.(string))
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
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *ListConditionsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListConditionsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListConditionsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListConditionsParams) SetCounterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["counterid"] = v
}

func (p *ListConditionsParams) ResetCounterid() {
	if p.p != nil && p.p["counterid"] != nil {
		delete(p.p, "counterid")
	}
}

func (p *ListConditionsParams) GetCounterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["counterid"].(string)
	return value, ok
}

func (p *ListConditionsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListConditionsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListConditionsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListConditionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListConditionsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListConditionsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListConditionsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListConditionsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListConditionsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListConditionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListConditionsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListConditionsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListConditionsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListConditionsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListConditionsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListConditionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListConditionsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListConditionsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListConditionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListConditionsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListConditionsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListConditionsParams) SetPolicyid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyid"] = v
}

func (p *ListConditionsParams) ResetPolicyid() {
	if p.p != nil && p.p["policyid"] != nil {
		delete(p.p, "policyid")
	}
}

func (p *ListConditionsParams) GetPolicyid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyid"].(string)
	return value, ok
}

func (p *ListConditionsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListConditionsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListConditionsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListConditionsParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListConditionsParams() *ListConditionsParams {
	p := &ListConditionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetConditionByID(id string, opts ...OptionFunc) (*Condition, int, error) {
	p := &ListConditionsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListConditions(p)
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
		return l.Conditions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Condition UUID: %s!", id)
}

// List Conditions for VM auto scaling
func (s *AutoScaleService) ListConditions(p *ListConditionsParams) (*ListConditionsResponse, error) {
	resp, err := s.cs.newRequest("listConditions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListConditionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListConditionsResponse struct {
	Count      int          `json:"count"`
	Conditions []*Condition `json:"condition"`
}

type Condition struct {
	Account            string   `json:"account"`
	Counter            *Counter `json:"counter"`
	Counterid          string   `json:"counterid"`
	Countername        string   `json:"countername"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Domainpath         string   `json:"domainpath"`
	Id                 string   `json:"id"`
	JobID              string   `json:"jobid"`
	Jobstatus          int      `json:"jobstatus"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Relationaloperator string   `json:"relationaloperator"`
	Threshold          int64    `json:"threshold"`
	Zoneid             string   `json:"zoneid"`
}

type ListCountersParams struct {
	p map[string]interface{}
}

func (p *ListCountersParams) toURLValues() url.Values {
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
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["source"]; found {
		u.Set("source", v.(string))
	}
	return u
}

func (p *ListCountersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListCountersParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListCountersParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListCountersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListCountersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListCountersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListCountersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListCountersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListCountersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListCountersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListCountersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListCountersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListCountersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListCountersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListCountersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListCountersParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListCountersParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListCountersParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *ListCountersParams) SetSource(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["source"] = v
}

func (p *ListCountersParams) ResetSource() {
	if p.p != nil && p.p["source"] != nil {
		delete(p.p, "source")
	}
}

func (p *ListCountersParams) GetSource() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["source"].(string)
	return value, ok
}

// You should always use this function to get a new ListCountersParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListCountersParams() *ListCountersParams {
	p := &ListCountersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListCountersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListCounters(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Counters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Counters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterByName(name string, opts ...OptionFunc) (*Counter, int, error) {
	id, count, err := s.GetCounterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetCounterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterByID(id string, opts ...OptionFunc) (*Counter, int, error) {
	p := &ListCountersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListCounters(p)
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
		return l.Counters[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Counter UUID: %s!", id)
}

// List the counters for VM auto scaling
func (s *AutoScaleService) ListCounters(p *ListCountersParams) (*ListCountersResponse, error) {
	resp, err := s.cs.newRequest("listCounters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCountersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCountersResponse struct {
	Count    int        `json:"count"`
	Counters []*Counter `json:"counter"`
}

type Counter struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Provider  string `json:"provider"`
	Source    string `json:"source"`
	Value     string `json:"value"`
	Zoneid    string `json:"zoneid"`
}

type UpdateAutoScalePolicyParams struct {
	p map[string]interface{}
}

func (p *UpdateAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["conditionids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("conditionids", vv)
	}
	if v, found := p.p["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["quiettime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("quiettime", vv)
	}
	return u
}

func (p *UpdateAutoScalePolicyParams) SetConditionids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["conditionids"] = v
}

func (p *UpdateAutoScalePolicyParams) ResetConditionids() {
	if p.p != nil && p.p["conditionids"] != nil {
		delete(p.p, "conditionids")
	}
}

func (p *UpdateAutoScalePolicyParams) GetConditionids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["conditionids"].([]string)
	return value, ok
}

func (p *UpdateAutoScalePolicyParams) SetDuration(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["duration"] = v
}

func (p *UpdateAutoScalePolicyParams) ResetDuration() {
	if p.p != nil && p.p["duration"] != nil {
		delete(p.p, "duration")
	}
}

func (p *UpdateAutoScalePolicyParams) GetDuration() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["duration"].(int)
	return value, ok
}

func (p *UpdateAutoScalePolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateAutoScalePolicyParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateAutoScalePolicyParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateAutoScalePolicyParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateAutoScalePolicyParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateAutoScalePolicyParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateAutoScalePolicyParams) SetQuiettime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["quiettime"] = v
}

func (p *UpdateAutoScalePolicyParams) ResetQuiettime() {
	if p.p != nil && p.p["quiettime"] != nil {
		delete(p.p, "quiettime")
	}
}

func (p *UpdateAutoScalePolicyParams) GetQuiettime() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["quiettime"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScalePolicyParams(id string) *UpdateAutoScalePolicyParams {
	p := &UpdateAutoScalePolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing autoscale policy.
func (s *AutoScaleService) UpdateAutoScalePolicy(p *UpdateAutoScalePolicyParams) (*UpdateAutoScalePolicyResponse, error) {
	resp, err := s.cs.newPostRequest("updateAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScalePolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateAutoScalePolicyResponse struct {
	Account    string       `json:"account"`
	Action     string       `json:"action"`
	Conditions []*Condition `json:"conditions"`
	Domain     string       `json:"domain"`
	Domainid   string       `json:"domainid"`
	Domainpath string       `json:"domainpath"`
	Duration   int          `json:"duration"`
	Id         string       `json:"id"`
	JobID      string       `json:"jobid"`
	Jobstatus  int          `json:"jobstatus"`
	Name       string       `json:"name"`
	Project    string       `json:"project"`
	Projectid  string       `json:"projectid"`
	Quiettime  int          `json:"quiettime"`
}

type UpdateAutoScaleVmGroupParams struct {
	p map[string]interface{}
}

func (p *UpdateAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	if v, found := p.p["maxmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxmembers", vv)
	}
	if v, found := p.p["minmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmembers", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["scaledownpolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaledownpolicyids", vv)
	}
	if v, found := p.p["scaleuppolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaleuppolicyids", vv)
	}
	return u
}

func (p *UpdateAutoScaleVmGroupParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmGroupParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateAutoScaleVmGroupParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmGroupParams) SetInterval(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["interval"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetInterval() {
	if p.p != nil && p.p["interval"] != nil {
		delete(p.p, "interval")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetInterval() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["interval"].(int)
	return value, ok
}

func (p *UpdateAutoScaleVmGroupParams) SetMaxmembers(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxmembers"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetMaxmembers() {
	if p.p != nil && p.p["maxmembers"] != nil {
		delete(p.p, "maxmembers")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetMaxmembers() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxmembers"].(int)
	return value, ok
}

func (p *UpdateAutoScaleVmGroupParams) SetMinmembers(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minmembers"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetMinmembers() {
	if p.p != nil && p.p["minmembers"] != nil {
		delete(p.p, "minmembers")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetMinmembers() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["minmembers"].(int)
	return value, ok
}

func (p *UpdateAutoScaleVmGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmGroupParams) SetScaledownpolicyids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scaledownpolicyids"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetScaledownpolicyids() {
	if p.p != nil && p.p["scaledownpolicyids"] != nil {
		delete(p.p, "scaledownpolicyids")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetScaledownpolicyids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scaledownpolicyids"].([]string)
	return value, ok
}

func (p *UpdateAutoScaleVmGroupParams) SetScaleuppolicyids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scaleuppolicyids"] = v
}

func (p *UpdateAutoScaleVmGroupParams) ResetScaleuppolicyids() {
	if p.p != nil && p.p["scaleuppolicyids"] != nil {
		delete(p.p, "scaleuppolicyids")
	}
}

func (p *UpdateAutoScaleVmGroupParams) GetScaleuppolicyids() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scaleuppolicyids"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScaleVmGroupParams(id string) *UpdateAutoScaleVmGroupParams {
	p := &UpdateAutoScaleVmGroupParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing autoscale vm group.
func (s *AutoScaleService) UpdateAutoScaleVmGroup(p *UpdateAutoScaleVmGroupParams) (*UpdateAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newPostRequest("updateAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScaleVmGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateAutoScaleVmGroupResponse struct {
	Account                      string             `json:"account"`
	Associatednetworkid          string             `json:"associatednetworkid"`
	Associatednetworkname        string             `json:"associatednetworkname"`
	Availablevirtualmachinecount int                `json:"availablevirtualmachinecount"`
	Created                      string             `json:"created"`
	Domain                       string             `json:"domain"`
	Domainid                     string             `json:"domainid"`
	Domainpath                   string             `json:"domainpath"`
	Fordisplay                   bool               `json:"fordisplay"`
	Hasannotations               bool               `json:"hasannotations"`
	Id                           string             `json:"id"`
	Interval                     int                `json:"interval"`
	JobID                        string             `json:"jobid"`
	Jobstatus                    int                `json:"jobstatus"`
	Lbprovider                   string             `json:"lbprovider"`
	Lbruleid                     string             `json:"lbruleid"`
	Maxmembers                   int                `json:"maxmembers"`
	Minmembers                   int                `json:"minmembers"`
	Name                         string             `json:"name"`
	Privateport                  string             `json:"privateport"`
	Project                      string             `json:"project"`
	Projectid                    string             `json:"projectid"`
	Publicip                     string             `json:"publicip"`
	Publicipid                   string             `json:"publicipid"`
	Publicport                   string             `json:"publicport"`
	Scaledownpolicies            []*AutoScalePolicy `json:"scaledownpolicies"`
	Scaleuppolicies              []*AutoScalePolicy `json:"scaleuppolicies"`
	State                        string             `json:"state"`
	Vmprofileid                  string             `json:"vmprofileid"`
}

type UpdateAutoScaleVmProfileParams struct {
	p map[string]interface{}
}

func (p *UpdateAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["autoscaleuserid"]; found {
		u.Set("autoscaleuserid", v.(string))
	}
	if v, found := p.p["counterparam"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("counterparam[%d].key", i), k)
			u.Set(fmt.Sprintf("counterparam[%d].value", i), m[k])
		}
	}
	if v, found := p.p["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := p.p["expungevmgraceperiod"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("expungevmgraceperiod", vv)
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["otherdeployparams"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("otherdeployparams[%d].name", i), k)
			u.Set(fmt.Sprintf("otherdeployparams[%d].value", i), m[k])
		}
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	if v, found := p.p["userdatadetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("userdatadetails[%d].key", i), k)
			u.Set(fmt.Sprintf("userdatadetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["userdataid"]; found {
		u.Set("userdataid", v.(string))
	}
	return u
}

func (p *UpdateAutoScaleVmProfileParams) SetAutoscaleuserid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["autoscaleuserid"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetAutoscaleuserid() {
	if p.p != nil && p.p["autoscaleuserid"] != nil {
		delete(p.p, "autoscaleuserid")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetAutoscaleuserid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["autoscaleuserid"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetCounterparam(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["counterparam"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetCounterparam() {
	if p.p != nil && p.p["counterparam"] != nil {
		delete(p.p, "counterparam")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetCounterparam() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["counterparam"].(map[string]string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetCustomid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customid"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetCustomid() {
	if p.p != nil && p.p["customid"] != nil {
		delete(p.p, "customid")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetCustomid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customid"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetExpungevmgraceperiod(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expungevmgraceperiod"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetExpungevmgraceperiod() {
	if p.p != nil && p.p["expungevmgraceperiod"] != nil {
		delete(p.p, "expungevmgraceperiod")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetExpungevmgraceperiod() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["expungevmgraceperiod"].(int)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetFordisplay() {
	if p.p != nil && p.p["fordisplay"] != nil {
		delete(p.p, "fordisplay")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetFordisplay() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fordisplay"].(bool)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetOtherdeployparams(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["otherdeployparams"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetOtherdeployparams() {
	if p.p != nil && p.p["otherdeployparams"] != nil {
		delete(p.p, "otherdeployparams")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetOtherdeployparams() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["otherdeployparams"].(map[string]string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetTemplateid() {
	if p.p != nil && p.p["templateid"] != nil {
		delete(p.p, "templateid")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetTemplateid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["templateid"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetUserdata(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdata"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetUserdata() {
	if p.p != nil && p.p["userdata"] != nil {
		delete(p.p, "userdata")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetUserdata() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdata"].(string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetUserdatadetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdatadetails"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetUserdatadetails() {
	if p.p != nil && p.p["userdatadetails"] != nil {
		delete(p.p, "userdatadetails")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetUserdatadetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdatadetails"].(map[string]string)
	return value, ok
}

func (p *UpdateAutoScaleVmProfileParams) SetUserdataid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userdataid"] = v
}

func (p *UpdateAutoScaleVmProfileParams) ResetUserdataid() {
	if p.p != nil && p.p["userdataid"] != nil {
		delete(p.p, "userdataid")
	}
}

func (p *UpdateAutoScaleVmProfileParams) GetUserdataid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["userdataid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScaleVmProfileParams(id string) *UpdateAutoScaleVmProfileParams {
	p := &UpdateAutoScaleVmProfileParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates an existing autoscale vm profile.
func (s *AutoScaleService) UpdateAutoScaleVmProfile(p *UpdateAutoScaleVmProfileParams) (*UpdateAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newPostRequest("updateAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScaleVmProfileResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateAutoScaleVmProfileResponse struct {
	Account              string            `json:"account"`
	Autoscaleuserid      string            `json:"autoscaleuserid"`
	Domain               string            `json:"domain"`
	Domainid             string            `json:"domainid"`
	Domainpath           string            `json:"domainpath"`
	Expungevmgraceperiod int               `json:"expungevmgraceperiod"`
	Fordisplay           bool              `json:"fordisplay"`
	Id                   string            `json:"id"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Otherdeployparams    map[string]string `json:"otherdeployparams"`
	Project              string            `json:"project"`
	Projectid            string            `json:"projectid"`
	Serviceofferingid    string            `json:"serviceofferingid"`
	Templateid           string            `json:"templateid"`
	Userdata             string            `json:"userdata"`
	Userdatadetails      string            `json:"userdatadetails"`
	Userdataid           string            `json:"userdataid"`
	Userdataname         string            `json:"userdataname"`
	Userdatapolicy       string            `json:"userdatapolicy"`
	Zoneid               string            `json:"zoneid"`
}

type UpdateConditionParams struct {
	p map[string]interface{}
}

func (p *UpdateConditionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["relationaloperator"]; found {
		u.Set("relationaloperator", v.(string))
	}
	if v, found := p.p["threshold"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("threshold", vv)
	}
	return u
}

func (p *UpdateConditionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateConditionParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateConditionParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateConditionParams) SetRelationaloperator(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["relationaloperator"] = v
}

func (p *UpdateConditionParams) ResetRelationaloperator() {
	if p.p != nil && p.p["relationaloperator"] != nil {
		delete(p.p, "relationaloperator")
	}
}

func (p *UpdateConditionParams) GetRelationaloperator() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["relationaloperator"].(string)
	return value, ok
}

func (p *UpdateConditionParams) SetThreshold(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["threshold"] = v
}

func (p *UpdateConditionParams) ResetThreshold() {
	if p.p != nil && p.p["threshold"] != nil {
		delete(p.p, "threshold")
	}
}

func (p *UpdateConditionParams) GetThreshold() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["threshold"].(int64)
	return value, ok
}

// You should always use this function to get a new UpdateConditionParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateConditionParams(id string, relationaloperator string, threshold int64) *UpdateConditionParams {
	p := &UpdateConditionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["relationaloperator"] = relationaloperator
	p.p["threshold"] = threshold
	return p
}

// Updates a condition for VM auto scaling
func (s *AutoScaleService) UpdateCondition(p *UpdateConditionParams) (*UpdateConditionResponse, error) {
	resp, err := s.cs.newPostRequest("updateCondition", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateConditionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateConditionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
