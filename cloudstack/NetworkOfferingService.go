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

type NetworkOfferingServiceIface interface {
	CreateNetworkOffering(p *CreateNetworkOfferingParams) (*CreateNetworkOfferingResponse, error)
	NewCreateNetworkOfferingParams(displaytext string, guestiptype string, name string, traffictype string) *CreateNetworkOfferingParams
	DeleteNetworkOffering(p *DeleteNetworkOfferingParams) (*DeleteNetworkOfferingResponse, error)
	NewDeleteNetworkOfferingParams(id string) *DeleteNetworkOfferingParams
	ListNetworkOfferings(p *ListNetworkOfferingsParams) (*ListNetworkOfferingsResponse, error)
	NewListNetworkOfferingsParams() *ListNetworkOfferingsParams
	GetNetworkOfferingID(name string, opts ...OptionFunc) (string, int, error)
	GetNetworkOfferingByName(name string, opts ...OptionFunc) (*NetworkOffering, int, error)
	GetNetworkOfferingByID(id string, opts ...OptionFunc) (*NetworkOffering, int, error)
	UpdateNetworkOffering(p *UpdateNetworkOfferingParams) (*UpdateNetworkOfferingResponse, error)
	NewUpdateNetworkOfferingParams() *UpdateNetworkOfferingParams
}

type CreateNetworkOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := p.p["conservemode"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("conservemode", vv)
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("domainid", vv)
	}
	if v, found := p.p["egressdefaultpolicy"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("egressdefaultpolicy", vv)
	}
	if v, found := p.p["enable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enable", vv)
	}
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["guestiptype"]; found {
		u.Set("guestiptype", v.(string))
	}
	if v, found := p.p["ispersistent"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispersistent", vv)
	}
	if v, found := p.p["keepaliveenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("keepaliveenabled", vv)
	}
	if v, found := p.p["maxconnections"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxconnections", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkrate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("networkrate", vv)
	}
	if v, found := p.p["servicecapabilitylist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].key", i), k)
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].value", i), m[k])
		}
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["serviceproviderlist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("serviceproviderlist[%d].service", i), k)
			u.Set(fmt.Sprintf("serviceproviderlist[%d].provider", i), m[k])
		}
	}
	if v, found := p.p["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
	}
	if v, found := p.p["specifyvlan"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyvlan", vv)
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneid", vv)
	}
	return u
}

func (p *CreateNetworkOfferingParams) SetAvailability(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["availability"] = v
}

func (p *CreateNetworkOfferingParams) GetAvailability() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["availability"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetConservemode(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["conservemode"] = v
}

func (p *CreateNetworkOfferingParams) GetConservemode() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["conservemode"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateNetworkOfferingParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateNetworkOfferingParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetDomainid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateNetworkOfferingParams) GetDomainid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].([]string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetEgressdefaultpolicy(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["egressdefaultpolicy"] = v
}

func (p *CreateNetworkOfferingParams) GetEgressdefaultpolicy() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["egressdefaultpolicy"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetEnable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enable"] = v
}

func (p *CreateNetworkOfferingParams) GetEnable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enable"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
}

func (p *CreateNetworkOfferingParams) GetForvpc() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forvpc"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetGuestiptype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestiptype"] = v
}

func (p *CreateNetworkOfferingParams) GetGuestiptype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["guestiptype"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetIspersistent(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispersistent"] = v
}

func (p *CreateNetworkOfferingParams) GetIspersistent() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ispersistent"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetKeepaliveenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keepaliveenabled"] = v
}

func (p *CreateNetworkOfferingParams) GetKeepaliveenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keepaliveenabled"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetMaxconnections(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxconnections"] = v
}

func (p *CreateNetworkOfferingParams) GetMaxconnections() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxconnections"].(int)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateNetworkOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetNetworkrate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkrate"] = v
}

func (p *CreateNetworkOfferingParams) GetNetworkrate() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkrate"].(int)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetServicecapabilitylist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicecapabilitylist"] = v
}

func (p *CreateNetworkOfferingParams) GetServicecapabilitylist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["servicecapabilitylist"].(map[string]string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
}

func (p *CreateNetworkOfferingParams) GetServiceofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingid"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetServiceproviderlist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceproviderlist"] = v
}

func (p *CreateNetworkOfferingParams) GetServiceproviderlist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceproviderlist"].(map[string]string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetSpecifyipranges(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyipranges"] = v
}

func (p *CreateNetworkOfferingParams) GetSpecifyipranges() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["specifyipranges"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetSpecifyvlan(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyvlan"] = v
}

func (p *CreateNetworkOfferingParams) GetSpecifyvlan() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["specifyvlan"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

func (p *CreateNetworkOfferingParams) GetSupportedservices() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["supportedservices"].([]string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreateNetworkOfferingParams) GetTags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *CreateNetworkOfferingParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetZoneid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateNetworkOfferingParams) GetZoneid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewCreateNetworkOfferingParams(displaytext string, guestiptype string, name string, traffictype string) *CreateNetworkOfferingParams {
	p := &CreateNetworkOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["guestiptype"] = guestiptype
	p.p["name"] = name
	p.p["traffictype"] = traffictype
	return p
}

// Creates a network offering.
func (s *NetworkOfferingService) CreateNetworkOffering(p *CreateNetworkOfferingParams) (*CreateNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("createNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateNetworkOfferingResponse struct {
	Availability             string                                 `json:"availability"`
	Conservemode             bool                                   `json:"conservemode"`
	Created                  string                                 `json:"created"`
	Details                  map[string]string                      `json:"details"`
	Displaytext              string                                 `json:"displaytext"`
	Domain                   string                                 `json:"domain"`
	Domainid                 string                                 `json:"domainid"`
	Egressdefaultpolicy      bool                                   `json:"egressdefaultpolicy"`
	Forvpc                   bool                                   `json:"forvpc"`
	Guestiptype              string                                 `json:"guestiptype"`
	Hasannotations           bool                                   `json:"hasannotations"`
	Id                       string                                 `json:"id"`
	Isdefault                bool                                   `json:"isdefault"`
	Ispersistent             bool                                   `json:"ispersistent"`
	JobID                    string                                 `json:"jobid"`
	Jobstatus                int                                    `json:"jobstatus"`
	Maxconnections           int                                    `json:"maxconnections"`
	Name                     string                                 `json:"name"`
	Networkrate              int                                    `json:"networkrate"`
	Service                  []CreateNetworkOfferingResponseService `json:"service"`
	Serviceofferingid        string                                 `json:"serviceofferingid"`
	Specifyipranges          bool                                   `json:"specifyipranges"`
	Specifyvlan              bool                                   `json:"specifyvlan"`
	State                    string                                 `json:"state"`
	Supportspublicaccess     bool                                   `json:"supportspublicaccess"`
	Supportsstrechedl2subnet bool                                   `json:"supportsstrechedl2subnet"`
	Tags                     string                                 `json:"tags"`
	Traffictype              string                                 `json:"traffictype"`
	Zone                     string                                 `json:"zone"`
	Zoneid                   string                                 `json:"zoneid"`
}

type CreateNetworkOfferingResponseService struct {
	Capability []CreateNetworkOfferingResponseServiceCapability `json:"capability"`
	Name       string                                           `json:"name"`
	Provider   []CreateNetworkOfferingResponseServiceProvider   `json:"provider"`
}

type CreateNetworkOfferingResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type CreateNetworkOfferingResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type DeleteNetworkOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteNetworkOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewDeleteNetworkOfferingParams(id string) *DeleteNetworkOfferingParams {
	p := &DeleteNetworkOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a network offering.
func (s *NetworkOfferingService) DeleteNetworkOffering(p *DeleteNetworkOfferingParams) (*DeleteNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteNetworkOfferingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteNetworkOfferingResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteNetworkOfferingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListNetworkOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["guestiptype"]; found {
		u.Set("guestiptype", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdefault"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdefault", vv)
	}
	if v, found := p.p["istagged"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("istagged", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["sourcenatsupported"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sourcenatsupported", vv)
	}
	if v, found := p.p["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
	}
	if v, found := p.p["specifyvlan"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyvlan", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListNetworkOfferingsParams) SetAvailability(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["availability"] = v
}

func (p *ListNetworkOfferingsParams) GetAvailability() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["availability"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *ListNetworkOfferingsParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListNetworkOfferingsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
}

func (p *ListNetworkOfferingsParams) GetForvpc() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forvpc"].(bool)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetGuestiptype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestiptype"] = v
}

func (p *ListNetworkOfferingsParams) GetGuestiptype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["guestiptype"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListNetworkOfferingsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetIsdefault(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdefault"] = v
}

func (p *ListNetworkOfferingsParams) GetIsdefault() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isdefault"].(bool)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetIstagged(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["istagged"] = v
}

func (p *ListNetworkOfferingsParams) GetIstagged() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["istagged"].(bool)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListNetworkOfferingsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListNetworkOfferingsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListNetworkOfferingsParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListNetworkOfferingsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListNetworkOfferingsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetSourcenatsupported(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatsupported"] = v
}

func (p *ListNetworkOfferingsParams) GetSourcenatsupported() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sourcenatsupported"].(bool)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetSpecifyipranges(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyipranges"] = v
}

func (p *ListNetworkOfferingsParams) GetSpecifyipranges() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["specifyipranges"].(bool)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetSpecifyvlan(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyvlan"] = v
}

func (p *ListNetworkOfferingsParams) GetSpecifyvlan() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["specifyvlan"].(bool)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListNetworkOfferingsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
}

func (p *ListNetworkOfferingsParams) GetSupportedservices() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["supportedservices"].([]string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *ListNetworkOfferingsParams) GetTags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
}

func (p *ListNetworkOfferingsParams) GetTraffictype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["traffictype"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListNetworkOfferingsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewListNetworkOfferingsParams() *ListNetworkOfferingsParams {
	p := &ListNetworkOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkOfferingService) GetNetworkOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListNetworkOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetworkOfferings(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.NetworkOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetworkOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkOfferingService) GetNetworkOfferingByName(name string, opts ...OptionFunc) (*NetworkOffering, int, error) {
	id, count, err := s.GetNetworkOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetNetworkOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkOfferingService) GetNetworkOfferingByID(id string, opts ...OptionFunc) (*NetworkOffering, int, error) {
	p := &ListNetworkOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListNetworkOfferings(p)
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
		return l.NetworkOfferings[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for NetworkOffering UUID: %s!", id)
}

// Lists all available network offerings.
func (s *NetworkOfferingService) ListNetworkOfferings(p *ListNetworkOfferingsParams) (*ListNetworkOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkOfferingsResponse struct {
	Count            int                `json:"count"`
	NetworkOfferings []*NetworkOffering `json:"networkoffering"`
}

type NetworkOffering struct {
	Availability             string                           `json:"availability"`
	Conservemode             bool                             `json:"conservemode"`
	Created                  string                           `json:"created"`
	Details                  map[string]string                `json:"details"`
	Displaytext              string                           `json:"displaytext"`
	Domain                   string                           `json:"domain"`
	Domainid                 string                           `json:"domainid"`
	Egressdefaultpolicy      bool                             `json:"egressdefaultpolicy"`
	Forvpc                   bool                             `json:"forvpc"`
	Guestiptype              string                           `json:"guestiptype"`
	Hasannotations           bool                             `json:"hasannotations"`
	Id                       string                           `json:"id"`
	Isdefault                bool                             `json:"isdefault"`
	Ispersistent             bool                             `json:"ispersistent"`
	JobID                    string                           `json:"jobid"`
	Jobstatus                int                              `json:"jobstatus"`
	Maxconnections           int                              `json:"maxconnections"`
	Name                     string                           `json:"name"`
	Networkrate              int                              `json:"networkrate"`
	Service                  []NetworkOfferingServiceInternal `json:"service"`
	Serviceofferingid        string                           `json:"serviceofferingid"`
	Specifyipranges          bool                             `json:"specifyipranges"`
	Specifyvlan              bool                             `json:"specifyvlan"`
	State                    string                           `json:"state"`
	Supportspublicaccess     bool                             `json:"supportspublicaccess"`
	Supportsstrechedl2subnet bool                             `json:"supportsstrechedl2subnet"`
	Tags                     string                           `json:"tags"`
	Traffictype              string                           `json:"traffictype"`
	Zone                     string                           `json:"zone"`
	Zoneid                   string                           `json:"zoneid"`
}

type NetworkOfferingServiceInternal struct {
	Capability []NetworkOfferingServiceInternalCapability `json:"capability"`
	Name       string                                     `json:"name"`
	Provider   []NetworkOfferingServiceInternalProvider   `json:"provider"`
}

type NetworkOfferingServiceInternalProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NetworkOfferingServiceInternalCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type UpdateNetworkOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keepaliveenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("keepaliveenabled", vv)
	}
	if v, found := p.p["maxconnections"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxconnections", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *UpdateNetworkOfferingParams) SetAvailability(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["availability"] = v
}

func (p *UpdateNetworkOfferingParams) GetAvailability() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["availability"].(string)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateNetworkOfferingParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpdateNetworkOfferingParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateNetworkOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetKeepaliveenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keepaliveenabled"] = v
}

func (p *UpdateNetworkOfferingParams) GetKeepaliveenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keepaliveenabled"].(bool)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetMaxconnections(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxconnections"] = v
}

func (p *UpdateNetworkOfferingParams) GetMaxconnections() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxconnections"].(int)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateNetworkOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

func (p *UpdateNetworkOfferingParams) GetSortkey() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortkey"].(int)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdateNetworkOfferingParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *UpdateNetworkOfferingParams) GetTags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(string)
	return value, ok
}

func (p *UpdateNetworkOfferingParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *UpdateNetworkOfferingParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewUpdateNetworkOfferingParams() *UpdateNetworkOfferingParams {
	p := &UpdateNetworkOfferingParams{}
	p.p = make(map[string]interface{})
	return p
}

// Updates a network offering.
func (s *NetworkOfferingService) UpdateNetworkOffering(p *UpdateNetworkOfferingParams) (*UpdateNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateNetworkOfferingResponse struct {
	Availability             string                                 `json:"availability"`
	Conservemode             bool                                   `json:"conservemode"`
	Created                  string                                 `json:"created"`
	Details                  map[string]string                      `json:"details"`
	Displaytext              string                                 `json:"displaytext"`
	Domain                   string                                 `json:"domain"`
	Domainid                 string                                 `json:"domainid"`
	Egressdefaultpolicy      bool                                   `json:"egressdefaultpolicy"`
	Forvpc                   bool                                   `json:"forvpc"`
	Guestiptype              string                                 `json:"guestiptype"`
	Hasannotations           bool                                   `json:"hasannotations"`
	Id                       string                                 `json:"id"`
	Isdefault                bool                                   `json:"isdefault"`
	Ispersistent             bool                                   `json:"ispersistent"`
	JobID                    string                                 `json:"jobid"`
	Jobstatus                int                                    `json:"jobstatus"`
	Maxconnections           int                                    `json:"maxconnections"`
	Name                     string                                 `json:"name"`
	Networkrate              int                                    `json:"networkrate"`
	Service                  []UpdateNetworkOfferingResponseService `json:"service"`
	Serviceofferingid        string                                 `json:"serviceofferingid"`
	Specifyipranges          bool                                   `json:"specifyipranges"`
	Specifyvlan              bool                                   `json:"specifyvlan"`
	State                    string                                 `json:"state"`
	Supportspublicaccess     bool                                   `json:"supportspublicaccess"`
	Supportsstrechedl2subnet bool                                   `json:"supportsstrechedl2subnet"`
	Tags                     string                                 `json:"tags"`
	Traffictype              string                                 `json:"traffictype"`
	Zone                     string                                 `json:"zone"`
	Zoneid                   string                                 `json:"zoneid"`
}

type UpdateNetworkOfferingResponseService struct {
	Capability []UpdateNetworkOfferingResponseServiceCapability `json:"capability"`
	Name       string                                           `json:"name"`
	Provider   []UpdateNetworkOfferingResponseServiceProvider   `json:"provider"`
}

type UpdateNetworkOfferingResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type UpdateNetworkOfferingResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}
