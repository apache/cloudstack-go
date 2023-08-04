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

type ServiceOfferingServiceIface interface {
	CreateServiceOffering(p *CreateServiceOfferingParams) (*CreateServiceOfferingResponse, error)
	NewCreateServiceOfferingParams(name string) *CreateServiceOfferingParams
	DeleteServiceOffering(p *DeleteServiceOfferingParams) (*DeleteServiceOfferingResponse, error)
	NewDeleteServiceOfferingParams(id string) *DeleteServiceOfferingParams
	ListServiceOfferings(p *ListServiceOfferingsParams) (*ListServiceOfferingsResponse, error)
	NewListServiceOfferingsParams() *ListServiceOfferingsParams
	GetServiceOfferingID(name string, opts ...OptionFunc) (string, int, error)
	GetServiceOfferingByName(name string, opts ...OptionFunc) (*ServiceOffering, int, error)
	GetServiceOfferingByID(id string, opts ...OptionFunc) (*ServiceOffering, int, error)
	UpdateServiceOffering(p *UpdateServiceOfferingParams) (*UpdateServiceOfferingResponse, error)
	NewUpdateServiceOfferingParams(id string) *UpdateServiceOfferingParams
}

type CreateServiceOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bytesreadrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadrate", vv)
	}
	if v, found := p.p["bytesreadratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadratemax", vv)
	}
	if v, found := p.p["bytesreadratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadratemaxlength", vv)
	}
	if v, found := p.p["byteswriterate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriterate", vv)
	}
	if v, found := p.p["byteswriteratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriteratemax", vv)
	}
	if v, found := p.p["byteswriteratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriteratemaxlength", vv)
	}
	if v, found := p.p["cachemode"]; found {
		u.Set("cachemode", v.(string))
	}
	if v, found := p.p["cpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpunumber", vv)
	}
	if v, found := p.p["cpuspeed"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpuspeed", vv)
	}
	if v, found := p.p["customized"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customized", vv)
	}
	if v, found := p.p["customizediops"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customizediops", vv)
	}
	if v, found := p.p["deploymentplanner"]; found {
		u.Set("deploymentplanner", v.(string))
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["diskofferingstrictness"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("diskofferingstrictness", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("domainid", vv)
	}
	if v, found := p.p["dynamicscalingenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dynamicscalingenabled", vv)
	}
	if v, found := p.p["encryptroot"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("encryptroot", vv)
	}
	if v, found := p.p["hosttags"]; found {
		u.Set("hosttags", v.(string))
	}
	if v, found := p.p["hypervisorsnapshotreserve"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("hypervisorsnapshotreserve", vv)
	}
	if v, found := p.p["iopsreadrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadrate", vv)
	}
	if v, found := p.p["iopsreadratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadratemax", vv)
	}
	if v, found := p.p["iopsreadratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadratemaxlength", vv)
	}
	if v, found := p.p["iopswriterate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriterate", vv)
	}
	if v, found := p.p["iopswriteratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriteratemax", vv)
	}
	if v, found := p.p["iopswriteratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriteratemaxlength", vv)
	}
	if v, found := p.p["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := p.p["isvolatile"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isvolatile", vv)
	}
	if v, found := p.p["limitcpuuse"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("limitcpuuse", vv)
	}
	if v, found := p.p["maxcpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxcpunumber", vv)
	}
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["maxmemory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxmemory", vv)
	}
	if v, found := p.p["memory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("memory", vv)
	}
	if v, found := p.p["mincpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("mincpunumber", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["minmemory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmemory", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkrate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("networkrate", vv)
	}
	if v, found := p.p["offerha"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("offerha", vv)
	}
	if v, found := p.p["provisioningtype"]; found {
		u.Set("provisioningtype", v.(string))
	}
	if v, found := p.p["rootdisksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("rootdisksize", vv)
	}
	if v, found := p.p["serviceofferingdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("serviceofferingdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("serviceofferingdetails[%d].value", i), m[k])
		}
	}
	if v, found := p.p["storagepolicy"]; found {
		u.Set("storagepolicy", v.(string))
	}
	if v, found := p.p["storagetype"]; found {
		u.Set("storagetype", v.(string))
	}
	if v, found := p.p["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneid", vv)
	}
	return u
}

func (p *CreateServiceOfferingParams) SetBytesreadrate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadrate"] = v
}

func (p *CreateServiceOfferingParams) GetBytesreadrate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadrate"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetBytesreadratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadratemax"] = v
}

func (p *CreateServiceOfferingParams) GetBytesreadratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadratemax"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetBytesreadratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadratemaxlength"] = v
}

func (p *CreateServiceOfferingParams) GetBytesreadratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadratemaxlength"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetByteswriterate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriterate"] = v
}

func (p *CreateServiceOfferingParams) GetByteswriterate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriterate"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetByteswriteratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriteratemax"] = v
}

func (p *CreateServiceOfferingParams) GetByteswriteratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriteratemax"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetByteswriteratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriteratemaxlength"] = v
}

func (p *CreateServiceOfferingParams) GetByteswriteratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriteratemaxlength"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetCachemode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cachemode"] = v
}

func (p *CreateServiceOfferingParams) GetCachemode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cachemode"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetCpunumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cpunumber"] = v
}

func (p *CreateServiceOfferingParams) GetCpunumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cpunumber"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetCpuspeed(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cpuspeed"] = v
}

func (p *CreateServiceOfferingParams) GetCpuspeed() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cpuspeed"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetCustomized(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customized"] = v
}

func (p *CreateServiceOfferingParams) GetCustomized() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customized"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetCustomizediops(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customizediops"] = v
}

func (p *CreateServiceOfferingParams) GetCustomizediops() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customizediops"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetDeploymentplanner(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deploymentplanner"] = v
}

func (p *CreateServiceOfferingParams) GetDeploymentplanner() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deploymentplanner"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
}

func (p *CreateServiceOfferingParams) GetDiskofferingid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingid"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetDiskofferingstrictness(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingstrictness"] = v
}

func (p *CreateServiceOfferingParams) GetDiskofferingstrictness() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["diskofferingstrictness"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateServiceOfferingParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetDomainid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateServiceOfferingParams) GetDomainid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].([]string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetDynamicscalingenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dynamicscalingenabled"] = v
}

func (p *CreateServiceOfferingParams) GetDynamicscalingenabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["dynamicscalingenabled"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetEncryptroot(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["encryptroot"] = v
}

func (p *CreateServiceOfferingParams) GetEncryptroot() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["encryptroot"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetHosttags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
}

func (p *CreateServiceOfferingParams) GetHosttags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hosttags"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetHypervisorsnapshotreserve(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisorsnapshotreserve"] = v
}

func (p *CreateServiceOfferingParams) GetHypervisorsnapshotreserve() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisorsnapshotreserve"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetIopsreadrate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadrate"] = v
}

func (p *CreateServiceOfferingParams) GetIopsreadrate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadrate"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetIopsreadratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadratemax"] = v
}

func (p *CreateServiceOfferingParams) GetIopsreadratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadratemax"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetIopsreadratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadratemaxlength"] = v
}

func (p *CreateServiceOfferingParams) GetIopsreadratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadratemaxlength"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetIopswriterate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriterate"] = v
}

func (p *CreateServiceOfferingParams) GetIopswriterate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriterate"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetIopswriteratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriteratemax"] = v
}

func (p *CreateServiceOfferingParams) GetIopswriteratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriteratemax"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetIopswriteratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriteratemaxlength"] = v
}

func (p *CreateServiceOfferingParams) GetIopswriteratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriteratemaxlength"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
}

func (p *CreateServiceOfferingParams) GetIssystem() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["issystem"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetIsvolatile(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isvolatile"] = v
}

func (p *CreateServiceOfferingParams) GetIsvolatile() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isvolatile"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetLimitcpuuse(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["limitcpuuse"] = v
}

func (p *CreateServiceOfferingParams) GetLimitcpuuse() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["limitcpuuse"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetMaxcpunumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxcpunumber"] = v
}

func (p *CreateServiceOfferingParams) GetMaxcpunumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxcpunumber"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *CreateServiceOfferingParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetMaxmemory(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxmemory"] = v
}

func (p *CreateServiceOfferingParams) GetMaxmemory() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxmemory"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetMemory(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["memory"] = v
}

func (p *CreateServiceOfferingParams) GetMemory() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["memory"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetMincpunumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["mincpunumber"] = v
}

func (p *CreateServiceOfferingParams) GetMincpunumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["mincpunumber"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *CreateServiceOfferingParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetMinmemory(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minmemory"] = v
}

func (p *CreateServiceOfferingParams) GetMinmemory() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["minmemory"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateServiceOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetNetworkrate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkrate"] = v
}

func (p *CreateServiceOfferingParams) GetNetworkrate() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkrate"].(int)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetOfferha(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["offerha"] = v
}

func (p *CreateServiceOfferingParams) GetOfferha() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["offerha"].(bool)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetProvisioningtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provisioningtype"] = v
}

func (p *CreateServiceOfferingParams) GetProvisioningtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provisioningtype"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetRootdisksize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["rootdisksize"] = v
}

func (p *CreateServiceOfferingParams) GetRootdisksize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["rootdisksize"].(int64)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetServiceofferingdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingdetails"] = v
}

func (p *CreateServiceOfferingParams) GetServiceofferingdetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serviceofferingdetails"].(map[string]string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetStoragepolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagepolicy"] = v
}

func (p *CreateServiceOfferingParams) GetStoragepolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagepolicy"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetStoragetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagetype"] = v
}

func (p *CreateServiceOfferingParams) GetStoragetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagetype"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetSystemvmtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["systemvmtype"] = v
}

func (p *CreateServiceOfferingParams) GetSystemvmtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["systemvmtype"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreateServiceOfferingParams) GetTags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(string)
	return value, ok
}

func (p *CreateServiceOfferingParams) SetZoneid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateServiceOfferingParams) GetZoneid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewCreateServiceOfferingParams(name string) *CreateServiceOfferingParams {
	p := &CreateServiceOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Creates a service offering.
func (s *ServiceOfferingService) CreateServiceOffering(p *CreateServiceOfferingParams) (*CreateServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("createServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateServiceOfferingResponse struct {
	CacheMode                   string            `json:"cacheMode"`
	Cpunumber                   int               `json:"cpunumber"`
	Cpuspeed                    int               `json:"cpuspeed"`
	Created                     string            `json:"created"`
	Defaultuse                  bool              `json:"defaultuse"`
	Deploymentplanner           string            `json:"deploymentplanner"`
	DiskBytesReadRate           int64             `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64             `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64             `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64             `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64             `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64             `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64             `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64             `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64             `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64             `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64             `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64             `json:"diskIopsWriteRateMaxLength"`
	Diskofferingdisplaytext     string            `json:"diskofferingdisplaytext"`
	Diskofferingid              string            `json:"diskofferingid"`
	Diskofferingname            string            `json:"diskofferingname"`
	Diskofferingstrictness      bool              `json:"diskofferingstrictness"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Dynamicscalingenabled       bool              `json:"dynamicscalingenabled"`
	Encryptroot                 bool              `json:"encryptroot"`
	Hasannotations              bool              `json:"hasannotations"`
	Hosttags                    string            `json:"hosttags"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	Issystem                    bool              `json:"issystem"`
	Isvolatile                  bool              `json:"isvolatile"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Limitcpuuse                 bool              `json:"limitcpuuse"`
	Maxiops                     int64             `json:"maxiops"`
	Memory                      int               `json:"memory"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Networkrate                 int               `json:"networkrate"`
	Offerha                     bool              `json:"offerha"`
	Provisioningtype            string            `json:"provisioningtype"`
	Rootdisksize                int64             `json:"rootdisksize"`
	Serviceofferingdetails      map[string]string `json:"serviceofferingdetails"`
	Storagetags                 string            `json:"storagetags"`
	Storagetype                 string            `json:"storagetype"`
	Systemvmtype                string            `json:"systemvmtype"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}

type DeleteServiceOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteServiceOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteServiceOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewDeleteServiceOfferingParams(id string) *DeleteServiceOfferingParams {
	p := &DeleteServiceOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a service offering.
func (s *ServiceOfferingService) DeleteServiceOffering(p *DeleteServiceOfferingParams) (*DeleteServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteServiceOfferingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteServiceOfferingResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteServiceOfferingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListServiceOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListServiceOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpunumber", vv)
	}
	if v, found := p.p["cpuspeed"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpuspeed", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["encryptroot"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("encryptroot", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["memory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("memory", vv)
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
	if v, found := p.p["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListServiceOfferingsParams) SetCpunumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cpunumber"] = v
}

func (p *ListServiceOfferingsParams) GetCpunumber() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cpunumber"].(int)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetCpuspeed(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cpuspeed"] = v
}

func (p *ListServiceOfferingsParams) GetCpuspeed() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cpuspeed"].(int)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListServiceOfferingsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetEncryptroot(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["encryptroot"] = v
}

func (p *ListServiceOfferingsParams) GetEncryptroot() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["encryptroot"].(bool)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListServiceOfferingsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListServiceOfferingsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
}

func (p *ListServiceOfferingsParams) GetIssystem() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["issystem"].(bool)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListServiceOfferingsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListServiceOfferingsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetMemory(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["memory"] = v
}

func (p *ListServiceOfferingsParams) GetMemory() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["memory"].(int)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListServiceOfferingsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListServiceOfferingsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListServiceOfferingsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetSystemvmtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["systemvmtype"] = v
}

func (p *ListServiceOfferingsParams) GetSystemvmtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["systemvmtype"].(string)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListServiceOfferingsParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListServiceOfferingsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListServiceOfferingsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListServiceOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewListServiceOfferingsParams() *ListServiceOfferingsParams {
	p := &ListServiceOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListServiceOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListServiceOfferings(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ServiceOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ServiceOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingByName(name string, opts ...OptionFunc) (*ServiceOffering, int, error) {
	id, count, err := s.GetServiceOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetServiceOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingByID(id string, opts ...OptionFunc) (*ServiceOffering, int, error) {
	p := &ListServiceOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListServiceOfferings(p)
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
		return l.ServiceOfferings[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ServiceOffering UUID: %s!", id)
}

// Lists all available service offerings.
func (s *ServiceOfferingService) ListServiceOfferings(p *ListServiceOfferingsParams) (*ListServiceOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listServiceOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListServiceOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListServiceOfferingsResponse struct {
	Count            int                `json:"count"`
	ServiceOfferings []*ServiceOffering `json:"serviceoffering"`
}

type ServiceOffering struct {
	CacheMode                   string            `json:"cacheMode"`
	Cpunumber                   int               `json:"cpunumber"`
	Cpuspeed                    int               `json:"cpuspeed"`
	Created                     string            `json:"created"`
	Defaultuse                  bool              `json:"defaultuse"`
	Deploymentplanner           string            `json:"deploymentplanner"`
	DiskBytesReadRate           int64             `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64             `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64             `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64             `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64             `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64             `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64             `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64             `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64             `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64             `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64             `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64             `json:"diskIopsWriteRateMaxLength"`
	Diskofferingdisplaytext     string            `json:"diskofferingdisplaytext"`
	Diskofferingid              string            `json:"diskofferingid"`
	Diskofferingname            string            `json:"diskofferingname"`
	Diskofferingstrictness      bool              `json:"diskofferingstrictness"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Dynamicscalingenabled       bool              `json:"dynamicscalingenabled"`
	Encryptroot                 bool              `json:"encryptroot"`
	Hasannotations              bool              `json:"hasannotations"`
	Hosttags                    string            `json:"hosttags"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	Issystem                    bool              `json:"issystem"`
	Isvolatile                  bool              `json:"isvolatile"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Limitcpuuse                 bool              `json:"limitcpuuse"`
	Maxiops                     int64             `json:"maxiops"`
	Memory                      int               `json:"memory"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Networkrate                 int               `json:"networkrate"`
	Offerha                     bool              `json:"offerha"`
	Provisioningtype            string            `json:"provisioningtype"`
	Rootdisksize                int64             `json:"rootdisksize"`
	Serviceofferingdetails      map[string]string `json:"serviceofferingdetails"`
	Storagetags                 string            `json:"storagetags"`
	Storagetype                 string            `json:"storagetype"`
	Systemvmtype                string            `json:"systemvmtype"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}

type UpdateServiceOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hosttags"]; found {
		u.Set("hosttags", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	if v, found := p.p["storagetags"]; found {
		u.Set("storagetags", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *UpdateServiceOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateServiceOfferingParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateServiceOfferingParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpdateServiceOfferingParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpdateServiceOfferingParams) SetHosttags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
}

func (p *UpdateServiceOfferingParams) GetHosttags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hosttags"].(string)
	return value, ok
}

func (p *UpdateServiceOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateServiceOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateServiceOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateServiceOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateServiceOfferingParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

func (p *UpdateServiceOfferingParams) GetSortkey() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortkey"].(int)
	return value, ok
}

func (p *UpdateServiceOfferingParams) SetStoragetags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagetags"] = v
}

func (p *UpdateServiceOfferingParams) GetStoragetags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagetags"].(string)
	return value, ok
}

func (p *UpdateServiceOfferingParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *UpdateServiceOfferingParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewUpdateServiceOfferingParams(id string) *UpdateServiceOfferingParams {
	p := &UpdateServiceOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a service offering.
func (s *ServiceOfferingService) UpdateServiceOffering(p *UpdateServiceOfferingParams) (*UpdateServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateServiceOfferingResponse struct {
	CacheMode                   string            `json:"cacheMode"`
	Cpunumber                   int               `json:"cpunumber"`
	Cpuspeed                    int               `json:"cpuspeed"`
	Created                     string            `json:"created"`
	Defaultuse                  bool              `json:"defaultuse"`
	Deploymentplanner           string            `json:"deploymentplanner"`
	DiskBytesReadRate           int64             `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64             `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64             `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64             `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64             `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64             `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64             `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64             `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64             `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64             `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64             `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64             `json:"diskIopsWriteRateMaxLength"`
	Diskofferingdisplaytext     string            `json:"diskofferingdisplaytext"`
	Diskofferingid              string            `json:"diskofferingid"`
	Diskofferingname            string            `json:"diskofferingname"`
	Diskofferingstrictness      bool              `json:"diskofferingstrictness"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Dynamicscalingenabled       bool              `json:"dynamicscalingenabled"`
	Encryptroot                 bool              `json:"encryptroot"`
	Hasannotations              bool              `json:"hasannotations"`
	Hosttags                    string            `json:"hosttags"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	Issystem                    bool              `json:"issystem"`
	Isvolatile                  bool              `json:"isvolatile"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Limitcpuuse                 bool              `json:"limitcpuuse"`
	Maxiops                     int64             `json:"maxiops"`
	Memory                      int               `json:"memory"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Networkrate                 int               `json:"networkrate"`
	Offerha                     bool              `json:"offerha"`
	Provisioningtype            string            `json:"provisioningtype"`
	Rootdisksize                int64             `json:"rootdisksize"`
	Serviceofferingdetails      map[string]string `json:"serviceofferingdetails"`
	Storagetags                 string            `json:"storagetags"`
	Storagetype                 string            `json:"storagetype"`
	Systemvmtype                string            `json:"systemvmtype"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}
