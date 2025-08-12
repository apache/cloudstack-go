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
	if v, found := p.p["fornsx"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fornsx", vv)
	}
	if v, found := p.p["fortungsten"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fortungsten", vv)
	}
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["guestiptype"]; found {
		u.Set("guestiptype", v.(string))
	}
	if v, found := p.p["internetprotocol"]; found {
		u.Set("internetprotocol", v.(string))
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
	if v, found := p.p["networkmode"]; found {
		u.Set("networkmode", v.(string))
	}
	if v, found := p.p["networkrate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("networkrate", vv)
	}
	if v, found := p.p["nsxsupportlb"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("nsxsupportlb", vv)
	}
	if v, found := p.p["nsxsupportsinternallb"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("nsxsupportsinternallb", vv)
	}
	if v, found := p.p["routingmode"]; found {
		u.Set("routingmode", v.(string))
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
	if v, found := p.p["specifyasnumber"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyasnumber", vv)
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

func (p *CreateNetworkOfferingParams) ResetAvailability() {
	if p.p != nil && p.p["availability"] != nil {
		delete(p.p, "availability")
	}
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

func (p *CreateNetworkOfferingParams) ResetConservemode() {
	if p.p != nil && p.p["conservemode"] != nil {
		delete(p.p, "conservemode")
	}
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

func (p *CreateNetworkOfferingParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
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

func (p *CreateNetworkOfferingParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
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

func (p *CreateNetworkOfferingParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *CreateNetworkOfferingParams) ResetEgressdefaultpolicy() {
	if p.p != nil && p.p["egressdefaultpolicy"] != nil {
		delete(p.p, "egressdefaultpolicy")
	}
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

func (p *CreateNetworkOfferingParams) ResetEnable() {
	if p.p != nil && p.p["enable"] != nil {
		delete(p.p, "enable")
	}
}

func (p *CreateNetworkOfferingParams) GetEnable() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enable"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetFornsx(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fornsx"] = v
}

func (p *CreateNetworkOfferingParams) ResetFornsx() {
	if p.p != nil && p.p["fornsx"] != nil {
		delete(p.p, "fornsx")
	}
}

func (p *CreateNetworkOfferingParams) GetFornsx() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fornsx"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetFortungsten(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fortungsten"] = v
}

func (p *CreateNetworkOfferingParams) ResetFortungsten() {
	if p.p != nil && p.p["fortungsten"] != nil {
		delete(p.p, "fortungsten")
	}
}

func (p *CreateNetworkOfferingParams) GetFortungsten() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["fortungsten"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
}

func (p *CreateNetworkOfferingParams) ResetForvpc() {
	if p.p != nil && p.p["forvpc"] != nil {
		delete(p.p, "forvpc")
	}
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

func (p *CreateNetworkOfferingParams) ResetGuestiptype() {
	if p.p != nil && p.p["guestiptype"] != nil {
		delete(p.p, "guestiptype")
	}
}

func (p *CreateNetworkOfferingParams) GetGuestiptype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["guestiptype"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetInternetprotocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["internetprotocol"] = v
}

func (p *CreateNetworkOfferingParams) ResetInternetprotocol() {
	if p.p != nil && p.p["internetprotocol"] != nil {
		delete(p.p, "internetprotocol")
	}
}

func (p *CreateNetworkOfferingParams) GetInternetprotocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["internetprotocol"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetIspersistent(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispersistent"] = v
}

func (p *CreateNetworkOfferingParams) ResetIspersistent() {
	if p.p != nil && p.p["ispersistent"] != nil {
		delete(p.p, "ispersistent")
	}
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

func (p *CreateNetworkOfferingParams) ResetKeepaliveenabled() {
	if p.p != nil && p.p["keepaliveenabled"] != nil {
		delete(p.p, "keepaliveenabled")
	}
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

func (p *CreateNetworkOfferingParams) ResetMaxconnections() {
	if p.p != nil && p.p["maxconnections"] != nil {
		delete(p.p, "maxconnections")
	}
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

func (p *CreateNetworkOfferingParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateNetworkOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetNetworkmode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkmode"] = v
}

func (p *CreateNetworkOfferingParams) ResetNetworkmode() {
	if p.p != nil && p.p["networkmode"] != nil {
		delete(p.p, "networkmode")
	}
}

func (p *CreateNetworkOfferingParams) GetNetworkmode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkmode"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetNetworkrate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkrate"] = v
}

func (p *CreateNetworkOfferingParams) ResetNetworkrate() {
	if p.p != nil && p.p["networkrate"] != nil {
		delete(p.p, "networkrate")
	}
}

func (p *CreateNetworkOfferingParams) GetNetworkrate() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkrate"].(int)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetNsxsupportlb(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nsxsupportlb"] = v
}

func (p *CreateNetworkOfferingParams) ResetNsxsupportlb() {
	if p.p != nil && p.p["nsxsupportlb"] != nil {
		delete(p.p, "nsxsupportlb")
	}
}

func (p *CreateNetworkOfferingParams) GetNsxsupportlb() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nsxsupportlb"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetNsxsupportsinternallb(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nsxsupportsinternallb"] = v
}

func (p *CreateNetworkOfferingParams) ResetNsxsupportsinternallb() {
	if p.p != nil && p.p["nsxsupportsinternallb"] != nil {
		delete(p.p, "nsxsupportsinternallb")
	}
}

func (p *CreateNetworkOfferingParams) GetNsxsupportsinternallb() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nsxsupportsinternallb"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetRoutingmode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["routingmode"] = v
}

func (p *CreateNetworkOfferingParams) ResetRoutingmode() {
	if p.p != nil && p.p["routingmode"] != nil {
		delete(p.p, "routingmode")
	}
}

func (p *CreateNetworkOfferingParams) GetRoutingmode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["routingmode"].(string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetServicecapabilitylist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicecapabilitylist"] = v
}

func (p *CreateNetworkOfferingParams) ResetServicecapabilitylist() {
	if p.p != nil && p.p["servicecapabilitylist"] != nil {
		delete(p.p, "servicecapabilitylist")
	}
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

func (p *CreateNetworkOfferingParams) ResetServiceofferingid() {
	if p.p != nil && p.p["serviceofferingid"] != nil {
		delete(p.p, "serviceofferingid")
	}
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

func (p *CreateNetworkOfferingParams) ResetServiceproviderlist() {
	if p.p != nil && p.p["serviceproviderlist"] != nil {
		delete(p.p, "serviceproviderlist")
	}
}

func (p *CreateNetworkOfferingParams) GetServiceproviderlist() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceproviderlist"].(map[string]string)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetSpecifyasnumber(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyasnumber"] = v
}

func (p *CreateNetworkOfferingParams) ResetSpecifyasnumber() {
	if p.p != nil && p.p["specifyasnumber"] != nil {
		delete(p.p, "specifyasnumber")
	}
}

func (p *CreateNetworkOfferingParams) GetSpecifyasnumber() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["specifyasnumber"].(bool)
	return value, ok
}

func (p *CreateNetworkOfferingParams) SetSpecifyipranges(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyipranges"] = v
}

func (p *CreateNetworkOfferingParams) ResetSpecifyipranges() {
	if p.p != nil && p.p["specifyipranges"] != nil {
		delete(p.p, "specifyipranges")
	}
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

func (p *CreateNetworkOfferingParams) ResetSpecifyvlan() {
	if p.p != nil && p.p["specifyvlan"] != nil {
		delete(p.p, "specifyvlan")
	}
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

func (p *CreateNetworkOfferingParams) ResetSupportedservices() {
	if p.p != nil && p.p["supportedservices"] != nil {
		delete(p.p, "supportedservices")
	}
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

func (p *CreateNetworkOfferingParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
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

func (p *CreateNetworkOfferingParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
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

func (p *CreateNetworkOfferingParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
	resp, err := s.cs.newPostRequest("createNetworkOffering", p.toURLValues())
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
	Fornsx                   bool                                   `json:"fornsx"`
	Fortungsten              bool                                   `json:"fortungsten"`
	Forvpc                   bool                                   `json:"forvpc"`
	Guestiptype              string                                 `json:"guestiptype"`
	Hasannotations           bool                                   `json:"hasannotations"`
	Id                       string                                 `json:"id"`
	Internetprotocol         string                                 `json:"internetprotocol"`
	Isdefault                bool                                   `json:"isdefault"`
	Ispersistent             bool                                   `json:"ispersistent"`
	JobID                    string                                 `json:"jobid"`
	Jobstatus                int                                    `json:"jobstatus"`
	Maxconnections           int                                    `json:"maxconnections"`
	Name                     string                                 `json:"name"`
	Networkmode              string                                 `json:"networkmode"`
	Networkrate              int                                    `json:"networkrate"`
	Routingmode              string                                 `json:"routingmode"`
	Service                  []CreateNetworkOfferingResponseService `json:"service"`
	Serviceofferingid        string                                 `json:"serviceofferingid"`
	Specifyasnumber          bool                                   `json:"specifyasnumber"`
	Specifyipranges          bool                                   `json:"specifyipranges"`
	Specifyvlan              bool                                   `json:"specifyvlan"`
	State                    string                                 `json:"state"`
	Supportsinternallb       bool                                   `json:"supportsinternallb"`
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

func (p *DeleteNetworkOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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
	resp, err := s.cs.newPostRequest("deleteNetworkOffering", p.toURLValues())
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
	if v, found := p.p["routingmode"]; found {
		u.Set("routingmode", v.(string))
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

func (p *ListNetworkOfferingsParams) ResetAvailability() {
	if p.p != nil && p.p["availability"] != nil {
		delete(p.p, "availability")
	}
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

func (p *ListNetworkOfferingsParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
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

func (p *ListNetworkOfferingsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *ListNetworkOfferingsParams) ResetForvpc() {
	if p.p != nil && p.p["forvpc"] != nil {
		delete(p.p, "forvpc")
	}
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

func (p *ListNetworkOfferingsParams) ResetGuestiptype() {
	if p.p != nil && p.p["guestiptype"] != nil {
		delete(p.p, "guestiptype")
	}
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

func (p *ListNetworkOfferingsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *ListNetworkOfferingsParams) ResetIsdefault() {
	if p.p != nil && p.p["isdefault"] != nil {
		delete(p.p, "isdefault")
	}
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

func (p *ListNetworkOfferingsParams) ResetIstagged() {
	if p.p != nil && p.p["istagged"] != nil {
		delete(p.p, "istagged")
	}
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

func (p *ListNetworkOfferingsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
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

func (p *ListNetworkOfferingsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *ListNetworkOfferingsParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
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

func (p *ListNetworkOfferingsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
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

func (p *ListNetworkOfferingsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListNetworkOfferingsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetRoutingmode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["routingmode"] = v
}

func (p *ListNetworkOfferingsParams) ResetRoutingmode() {
	if p.p != nil && p.p["routingmode"] != nil {
		delete(p.p, "routingmode")
	}
}

func (p *ListNetworkOfferingsParams) GetRoutingmode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["routingmode"].(string)
	return value, ok
}

func (p *ListNetworkOfferingsParams) SetSourcenatsupported(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatsupported"] = v
}

func (p *ListNetworkOfferingsParams) ResetSourcenatsupported() {
	if p.p != nil && p.p["sourcenatsupported"] != nil {
		delete(p.p, "sourcenatsupported")
	}
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

func (p *ListNetworkOfferingsParams) ResetSpecifyipranges() {
	if p.p != nil && p.p["specifyipranges"] != nil {
		delete(p.p, "specifyipranges")
	}
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

func (p *ListNetworkOfferingsParams) ResetSpecifyvlan() {
	if p.p != nil && p.p["specifyvlan"] != nil {
		delete(p.p, "specifyvlan")
	}
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

func (p *ListNetworkOfferingsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
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

func (p *ListNetworkOfferingsParams) ResetSupportedservices() {
	if p.p != nil && p.p["supportedservices"] != nil {
		delete(p.p, "supportedservices")
	}
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

func (p *ListNetworkOfferingsParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
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

func (p *ListNetworkOfferingsParams) ResetTraffictype() {
	if p.p != nil && p.p["traffictype"] != nil {
		delete(p.p, "traffictype")
	}
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

func (p *ListNetworkOfferingsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
	Fornsx                   bool                             `json:"fornsx"`
	Fortungsten              bool                             `json:"fortungsten"`
	Forvpc                   bool                             `json:"forvpc"`
	Guestiptype              string                           `json:"guestiptype"`
	Hasannotations           bool                             `json:"hasannotations"`
	Id                       string                           `json:"id"`
	Internetprotocol         string                           `json:"internetprotocol"`
	Isdefault                bool                             `json:"isdefault"`
	Ispersistent             bool                             `json:"ispersistent"`
	JobID                    string                           `json:"jobid"`
	Jobstatus                int                              `json:"jobstatus"`
	Maxconnections           int                              `json:"maxconnections"`
	Name                     string                           `json:"name"`
	Networkmode              string                           `json:"networkmode"`
	Networkrate              int                              `json:"networkrate"`
	Routingmode              string                           `json:"routingmode"`
	Service                  []NetworkOfferingServiceInternal `json:"service"`
	Serviceofferingid        string                           `json:"serviceofferingid"`
	Specifyasnumber          bool                             `json:"specifyasnumber"`
	Specifyipranges          bool                             `json:"specifyipranges"`
	Specifyvlan              bool                             `json:"specifyvlan"`
	State                    string                           `json:"state"`
	Supportsinternallb       bool                             `json:"supportsinternallb"`
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

func (p *UpdateNetworkOfferingParams) ResetAvailability() {
	if p.p != nil && p.p["availability"] != nil {
		delete(p.p, "availability")
	}
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

func (p *UpdateNetworkOfferingParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
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

func (p *UpdateNetworkOfferingParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
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

func (p *UpdateNetworkOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
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

func (p *UpdateNetworkOfferingParams) ResetKeepaliveenabled() {
	if p.p != nil && p.p["keepaliveenabled"] != nil {
		delete(p.p, "keepaliveenabled")
	}
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

func (p *UpdateNetworkOfferingParams) ResetMaxconnections() {
	if p.p != nil && p.p["maxconnections"] != nil {
		delete(p.p, "maxconnections")
	}
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

func (p *UpdateNetworkOfferingParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
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

func (p *UpdateNetworkOfferingParams) ResetSortkey() {
	if p.p != nil && p.p["sortkey"] != nil {
		delete(p.p, "sortkey")
	}
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

func (p *UpdateNetworkOfferingParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
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

func (p *UpdateNetworkOfferingParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
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

func (p *UpdateNetworkOfferingParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
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
	resp, err := s.cs.newPostRequest("updateNetworkOffering", p.toURLValues())
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
	Fornsx                   bool                                   `json:"fornsx"`
	Fortungsten              bool                                   `json:"fortungsten"`
	Forvpc                   bool                                   `json:"forvpc"`
	Guestiptype              string                                 `json:"guestiptype"`
	Hasannotations           bool                                   `json:"hasannotations"`
	Id                       string                                 `json:"id"`
	Internetprotocol         string                                 `json:"internetprotocol"`
	Isdefault                bool                                   `json:"isdefault"`
	Ispersistent             bool                                   `json:"ispersistent"`
	JobID                    string                                 `json:"jobid"`
	Jobstatus                int                                    `json:"jobstatus"`
	Maxconnections           int                                    `json:"maxconnections"`
	Name                     string                                 `json:"name"`
	Networkmode              string                                 `json:"networkmode"`
	Networkrate              int                                    `json:"networkrate"`
	Routingmode              string                                 `json:"routingmode"`
	Service                  []UpdateNetworkOfferingResponseService `json:"service"`
	Serviceofferingid        string                                 `json:"serviceofferingid"`
	Specifyasnumber          bool                                   `json:"specifyasnumber"`
	Specifyipranges          bool                                   `json:"specifyipranges"`
	Specifyvlan              bool                                   `json:"specifyvlan"`
	State                    string                                 `json:"state"`
	Supportsinternallb       bool                                   `json:"supportsinternallb"`
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
