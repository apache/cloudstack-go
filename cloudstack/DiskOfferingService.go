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

type DiskOfferingServiceIface interface {
	CreateDiskOffering(p *CreateDiskOfferingParams) (*CreateDiskOfferingResponse, error)
	NewCreateDiskOfferingParams(displaytext string, name string) *CreateDiskOfferingParams
	DeleteDiskOffering(p *DeleteDiskOfferingParams) (*DeleteDiskOfferingResponse, error)
	NewDeleteDiskOfferingParams(id string) *DeleteDiskOfferingParams
	ListDiskOfferings(p *ListDiskOfferingsParams) (*ListDiskOfferingsResponse, error)
	NewListDiskOfferingsParams() *ListDiskOfferingsParams
	GetDiskOfferingID(name string, opts ...OptionFunc) (string, int, error)
	GetDiskOfferingByName(name string, opts ...OptionFunc) (*DiskOffering, int, error)
	GetDiskOfferingByID(id string, opts ...OptionFunc) (*DiskOffering, int, error)
	UpdateDiskOffering(p *UpdateDiskOfferingParams) (*UpdateDiskOfferingResponse, error)
	NewUpdateDiskOfferingParams(id string) *UpdateDiskOfferingParams
}

type CreateDiskOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateDiskOfferingParams) toURLValues() url.Values {
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
	if v, found := p.p["customized"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customized", vv)
	}
	if v, found := p.p["customizediops"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customizediops", vv)
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["disksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("disksize", vv)
	}
	if v, found := p.p["disksizestrictness"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("disksizestrictness", vv)
	}
	if v, found := p.p["displayoffering"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayoffering", vv)
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("domainid", vv)
	}
	if v, found := p.p["encrypt"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("encrypt", vv)
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
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["provisioningtype"]; found {
		u.Set("provisioningtype", v.(string))
	}
	if v, found := p.p["storagepolicy"]; found {
		u.Set("storagepolicy", v.(string))
	}
	if v, found := p.p["storagetype"]; found {
		u.Set("storagetype", v.(string))
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

func (p *CreateDiskOfferingParams) SetBytesreadrate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadrate"] = v
}

func (p *CreateDiskOfferingParams) ResetBytesreadrate() {
	if p.p != nil && p.p["bytesreadrate"] != nil {
		delete(p.p, "bytesreadrate")
	}
}

func (p *CreateDiskOfferingParams) GetBytesreadrate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadrate"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetBytesreadratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadratemax"] = v
}

func (p *CreateDiskOfferingParams) ResetBytesreadratemax() {
	if p.p != nil && p.p["bytesreadratemax"] != nil {
		delete(p.p, "bytesreadratemax")
	}
}

func (p *CreateDiskOfferingParams) GetBytesreadratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadratemax"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetBytesreadratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadratemaxlength"] = v
}

func (p *CreateDiskOfferingParams) ResetBytesreadratemaxlength() {
	if p.p != nil && p.p["bytesreadratemaxlength"] != nil {
		delete(p.p, "bytesreadratemaxlength")
	}
}

func (p *CreateDiskOfferingParams) GetBytesreadratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadratemaxlength"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetByteswriterate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriterate"] = v
}

func (p *CreateDiskOfferingParams) ResetByteswriterate() {
	if p.p != nil && p.p["byteswriterate"] != nil {
		delete(p.p, "byteswriterate")
	}
}

func (p *CreateDiskOfferingParams) GetByteswriterate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriterate"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetByteswriteratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriteratemax"] = v
}

func (p *CreateDiskOfferingParams) ResetByteswriteratemax() {
	if p.p != nil && p.p["byteswriteratemax"] != nil {
		delete(p.p, "byteswriteratemax")
	}
}

func (p *CreateDiskOfferingParams) GetByteswriteratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriteratemax"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetByteswriteratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriteratemaxlength"] = v
}

func (p *CreateDiskOfferingParams) ResetByteswriteratemaxlength() {
	if p.p != nil && p.p["byteswriteratemaxlength"] != nil {
		delete(p.p, "byteswriteratemaxlength")
	}
}

func (p *CreateDiskOfferingParams) GetByteswriteratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriteratemaxlength"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetCachemode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cachemode"] = v
}

func (p *CreateDiskOfferingParams) ResetCachemode() {
	if p.p != nil && p.p["cachemode"] != nil {
		delete(p.p, "cachemode")
	}
}

func (p *CreateDiskOfferingParams) GetCachemode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cachemode"].(string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetCustomized(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customized"] = v
}

func (p *CreateDiskOfferingParams) ResetCustomized() {
	if p.p != nil && p.p["customized"] != nil {
		delete(p.p, "customized")
	}
}

func (p *CreateDiskOfferingParams) GetCustomized() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customized"].(bool)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetCustomizediops(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customizediops"] = v
}

func (p *CreateDiskOfferingParams) ResetCustomizediops() {
	if p.p != nil && p.p["customizediops"] != nil {
		delete(p.p, "customizediops")
	}
}

func (p *CreateDiskOfferingParams) GetCustomizediops() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["customizediops"].(bool)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateDiskOfferingParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *CreateDiskOfferingParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetDisksize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["disksize"] = v
}

func (p *CreateDiskOfferingParams) ResetDisksize() {
	if p.p != nil && p.p["disksize"] != nil {
		delete(p.p, "disksize")
	}
}

func (p *CreateDiskOfferingParams) GetDisksize() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["disksize"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetDisksizestrictness(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["disksizestrictness"] = v
}

func (p *CreateDiskOfferingParams) ResetDisksizestrictness() {
	if p.p != nil && p.p["disksizestrictness"] != nil {
		delete(p.p, "disksizestrictness")
	}
}

func (p *CreateDiskOfferingParams) GetDisksizestrictness() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["disksizestrictness"].(bool)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetDisplayoffering(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayoffering"] = v
}

func (p *CreateDiskOfferingParams) ResetDisplayoffering() {
	if p.p != nil && p.p["displayoffering"] != nil {
		delete(p.p, "displayoffering")
	}
}

func (p *CreateDiskOfferingParams) GetDisplayoffering() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayoffering"].(bool)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateDiskOfferingParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *CreateDiskOfferingParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetDomainid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateDiskOfferingParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *CreateDiskOfferingParams) GetDomainid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].([]string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetEncrypt(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["encrypt"] = v
}

func (p *CreateDiskOfferingParams) ResetEncrypt() {
	if p.p != nil && p.p["encrypt"] != nil {
		delete(p.p, "encrypt")
	}
}

func (p *CreateDiskOfferingParams) GetEncrypt() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["encrypt"].(bool)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetHypervisorsnapshotreserve(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisorsnapshotreserve"] = v
}

func (p *CreateDiskOfferingParams) ResetHypervisorsnapshotreserve() {
	if p.p != nil && p.p["hypervisorsnapshotreserve"] != nil {
		delete(p.p, "hypervisorsnapshotreserve")
	}
}

func (p *CreateDiskOfferingParams) GetHypervisorsnapshotreserve() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisorsnapshotreserve"].(int)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetIopsreadrate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadrate"] = v
}

func (p *CreateDiskOfferingParams) ResetIopsreadrate() {
	if p.p != nil && p.p["iopsreadrate"] != nil {
		delete(p.p, "iopsreadrate")
	}
}

func (p *CreateDiskOfferingParams) GetIopsreadrate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadrate"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetIopsreadratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadratemax"] = v
}

func (p *CreateDiskOfferingParams) ResetIopsreadratemax() {
	if p.p != nil && p.p["iopsreadratemax"] != nil {
		delete(p.p, "iopsreadratemax")
	}
}

func (p *CreateDiskOfferingParams) GetIopsreadratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadratemax"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetIopsreadratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadratemaxlength"] = v
}

func (p *CreateDiskOfferingParams) ResetIopsreadratemaxlength() {
	if p.p != nil && p.p["iopsreadratemaxlength"] != nil {
		delete(p.p, "iopsreadratemaxlength")
	}
}

func (p *CreateDiskOfferingParams) GetIopsreadratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadratemaxlength"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetIopswriterate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriterate"] = v
}

func (p *CreateDiskOfferingParams) ResetIopswriterate() {
	if p.p != nil && p.p["iopswriterate"] != nil {
		delete(p.p, "iopswriterate")
	}
}

func (p *CreateDiskOfferingParams) GetIopswriterate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriterate"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetIopswriteratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriteratemax"] = v
}

func (p *CreateDiskOfferingParams) ResetIopswriteratemax() {
	if p.p != nil && p.p["iopswriteratemax"] != nil {
		delete(p.p, "iopswriteratemax")
	}
}

func (p *CreateDiskOfferingParams) GetIopswriteratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriteratemax"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetIopswriteratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriteratemaxlength"] = v
}

func (p *CreateDiskOfferingParams) ResetIopswriteratemaxlength() {
	if p.p != nil && p.p["iopswriteratemaxlength"] != nil {
		delete(p.p, "iopswriteratemaxlength")
	}
}

func (p *CreateDiskOfferingParams) GetIopswriteratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriteratemaxlength"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *CreateDiskOfferingParams) ResetMaxiops() {
	if p.p != nil && p.p["maxiops"] != nil {
		delete(p.p, "maxiops")
	}
}

func (p *CreateDiskOfferingParams) GetMaxiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["maxiops"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *CreateDiskOfferingParams) ResetMiniops() {
	if p.p != nil && p.p["miniops"] != nil {
		delete(p.p, "miniops")
	}
}

func (p *CreateDiskOfferingParams) GetMiniops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["miniops"].(int64)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateDiskOfferingParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateDiskOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetProvisioningtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provisioningtype"] = v
}

func (p *CreateDiskOfferingParams) ResetProvisioningtype() {
	if p.p != nil && p.p["provisioningtype"] != nil {
		delete(p.p, "provisioningtype")
	}
}

func (p *CreateDiskOfferingParams) GetProvisioningtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provisioningtype"].(string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetStoragepolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagepolicy"] = v
}

func (p *CreateDiskOfferingParams) ResetStoragepolicy() {
	if p.p != nil && p.p["storagepolicy"] != nil {
		delete(p.p, "storagepolicy")
	}
}

func (p *CreateDiskOfferingParams) GetStoragepolicy() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagepolicy"].(string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetStoragetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagetype"] = v
}

func (p *CreateDiskOfferingParams) ResetStoragetype() {
	if p.p != nil && p.p["storagetype"] != nil {
		delete(p.p, "storagetype")
	}
}

func (p *CreateDiskOfferingParams) GetStoragetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagetype"].(string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreateDiskOfferingParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *CreateDiskOfferingParams) GetTags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(string)
	return value, ok
}

func (p *CreateDiskOfferingParams) SetZoneid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateDiskOfferingParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateDiskOfferingParams) GetZoneid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewCreateDiskOfferingParams(displaytext string, name string) *CreateDiskOfferingParams {
	p := &CreateDiskOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	return p
}

// Creates a disk offering.
func (s *DiskOfferingService) CreateDiskOffering(p *CreateDiskOfferingParams) (*CreateDiskOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("createDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateDiskOfferingResponse struct {
	CacheMode                   string            `json:"cacheMode"`
	Created                     string            `json:"created"`
	Details                     map[string]string `json:"details"`
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
	Disksize                    int64             `json:"disksize"`
	Disksizestrictness          bool              `json:"disksizestrictness"`
	Displayoffering             bool              `json:"displayoffering"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Encrypt                     bool              `json:"encrypt"`
	Hasannotations              bool              `json:"hasannotations"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Maxiops                     int64             `json:"maxiops"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Provisioningtype            string            `json:"provisioningtype"`
	State                       string            `json:"state"`
	Storagetype                 string            `json:"storagetype"`
	Suitableforvirtualmachine   bool              `json:"suitableforvirtualmachine"`
	Tags                        string            `json:"tags"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}

type DeleteDiskOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteDiskOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteDiskOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteDiskOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteDiskOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewDeleteDiskOfferingParams(id string) *DeleteDiskOfferingParams {
	p := &DeleteDiskOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a disk offering.
func (s *DiskOfferingService) DeleteDiskOffering(p *DeleteDiskOfferingParams) (*DeleteDiskOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("deleteDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteDiskOfferingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteDiskOfferingResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteDiskOfferingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListDiskOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListDiskOfferingsParams) toURLValues() url.Values {
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
	if v, found := p.p["encrypt"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("encrypt", vv)
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
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["storagetype"]; found {
		u.Set("storagetype", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListDiskOfferingsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListDiskOfferingsParams) ResetAccount() {
	if p.p != nil && p.p["account"] != nil {
		delete(p.p, "account")
	}
}

func (p *ListDiskOfferingsParams) GetAccount() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["account"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListDiskOfferingsParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *ListDiskOfferingsParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetEncrypt(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["encrypt"] = v
}

func (p *ListDiskOfferingsParams) ResetEncrypt() {
	if p.p != nil && p.p["encrypt"] != nil {
		delete(p.p, "encrypt")
	}
}

func (p *ListDiskOfferingsParams) GetEncrypt() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["encrypt"].(bool)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListDiskOfferingsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListDiskOfferingsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListDiskOfferingsParams) ResetIsrecursive() {
	if p.p != nil && p.p["isrecursive"] != nil {
		delete(p.p, "isrecursive")
	}
}

func (p *ListDiskOfferingsParams) GetIsrecursive() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["isrecursive"].(bool)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListDiskOfferingsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListDiskOfferingsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListDiskOfferingsParams) ResetListall() {
	if p.p != nil && p.p["listall"] != nil {
		delete(p.p, "listall")
	}
}

func (p *ListDiskOfferingsParams) GetListall() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["listall"].(bool)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListDiskOfferingsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListDiskOfferingsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListDiskOfferingsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListDiskOfferingsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListDiskOfferingsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListDiskOfferingsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
}

func (p *ListDiskOfferingsParams) ResetProjectid() {
	if p.p != nil && p.p["projectid"] != nil {
		delete(p.p, "projectid")
	}
}

func (p *ListDiskOfferingsParams) GetProjectid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["projectid"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListDiskOfferingsParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *ListDiskOfferingsParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *ListDiskOfferingsParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *ListDiskOfferingsParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetStoragetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagetype"] = v
}

func (p *ListDiskOfferingsParams) ResetStoragetype() {
	if p.p != nil && p.p["storagetype"] != nil {
		delete(p.p, "storagetype")
	}
}

func (p *ListDiskOfferingsParams) GetStoragetype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagetype"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

func (p *ListDiskOfferingsParams) ResetVirtualmachineid() {
	if p.p != nil && p.p["virtualmachineid"] != nil {
		delete(p.p, "virtualmachineid")
	}
}

func (p *ListDiskOfferingsParams) GetVirtualmachineid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["virtualmachineid"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetVolumeid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["volumeid"] = v
}

func (p *ListDiskOfferingsParams) ResetVolumeid() {
	if p.p != nil && p.p["volumeid"] != nil {
		delete(p.p, "volumeid")
	}
}

func (p *ListDiskOfferingsParams) GetVolumeid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["volumeid"].(string)
	return value, ok
}

func (p *ListDiskOfferingsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListDiskOfferingsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListDiskOfferingsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListDiskOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewListDiskOfferingsParams() *ListDiskOfferingsParams {
	p := &ListDiskOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DiskOfferingService) GetDiskOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListDiskOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListDiskOfferings(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.DiskOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.DiskOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DiskOfferingService) GetDiskOfferingByName(name string, opts ...OptionFunc) (*DiskOffering, int, error) {
	id, count, err := s.GetDiskOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetDiskOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DiskOfferingService) GetDiskOfferingByID(id string, opts ...OptionFunc) (*DiskOffering, int, error) {
	p := &ListDiskOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListDiskOfferings(p)
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
		return l.DiskOfferings[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for DiskOffering UUID: %s!", id)
}

// Lists all available disk offerings.
func (s *DiskOfferingService) ListDiskOfferings(p *ListDiskOfferingsParams) (*ListDiskOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listDiskOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDiskOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDiskOfferingsResponse struct {
	Count         int             `json:"count"`
	DiskOfferings []*DiskOffering `json:"diskoffering"`
}

type DiskOffering struct {
	CacheMode                   string            `json:"cacheMode"`
	Created                     string            `json:"created"`
	Details                     map[string]string `json:"details"`
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
	Disksize                    int64             `json:"disksize"`
	Disksizestrictness          bool              `json:"disksizestrictness"`
	Displayoffering             bool              `json:"displayoffering"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Encrypt                     bool              `json:"encrypt"`
	Hasannotations              bool              `json:"hasannotations"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Maxiops                     int64             `json:"maxiops"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Provisioningtype            string            `json:"provisioningtype"`
	State                       string            `json:"state"`
	Storagetype                 string            `json:"storagetype"`
	Suitableforvirtualmachine   bool              `json:"suitableforvirtualmachine"`
	Tags                        string            `json:"tags"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}

type UpdateDiskOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateDiskOfferingParams) toURLValues() url.Values {
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
	if v, found := p.p["displayoffering"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayoffering", vv)
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

func (p *UpdateDiskOfferingParams) SetBytesreadrate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadrate"] = v
}

func (p *UpdateDiskOfferingParams) ResetBytesreadrate() {
	if p.p != nil && p.p["bytesreadrate"] != nil {
		delete(p.p, "bytesreadrate")
	}
}

func (p *UpdateDiskOfferingParams) GetBytesreadrate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadrate"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetBytesreadratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadratemax"] = v
}

func (p *UpdateDiskOfferingParams) ResetBytesreadratemax() {
	if p.p != nil && p.p["bytesreadratemax"] != nil {
		delete(p.p, "bytesreadratemax")
	}
}

func (p *UpdateDiskOfferingParams) GetBytesreadratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadratemax"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetBytesreadratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadratemaxlength"] = v
}

func (p *UpdateDiskOfferingParams) ResetBytesreadratemaxlength() {
	if p.p != nil && p.p["bytesreadratemaxlength"] != nil {
		delete(p.p, "bytesreadratemaxlength")
	}
}

func (p *UpdateDiskOfferingParams) GetBytesreadratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["bytesreadratemaxlength"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetByteswriterate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriterate"] = v
}

func (p *UpdateDiskOfferingParams) ResetByteswriterate() {
	if p.p != nil && p.p["byteswriterate"] != nil {
		delete(p.p, "byteswriterate")
	}
}

func (p *UpdateDiskOfferingParams) GetByteswriterate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriterate"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetByteswriteratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriteratemax"] = v
}

func (p *UpdateDiskOfferingParams) ResetByteswriteratemax() {
	if p.p != nil && p.p["byteswriteratemax"] != nil {
		delete(p.p, "byteswriteratemax")
	}
}

func (p *UpdateDiskOfferingParams) GetByteswriteratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriteratemax"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetByteswriteratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriteratemaxlength"] = v
}

func (p *UpdateDiskOfferingParams) ResetByteswriteratemaxlength() {
	if p.p != nil && p.p["byteswriteratemaxlength"] != nil {
		delete(p.p, "byteswriteratemaxlength")
	}
}

func (p *UpdateDiskOfferingParams) GetByteswriteratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["byteswriteratemaxlength"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetCachemode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cachemode"] = v
}

func (p *UpdateDiskOfferingParams) ResetCachemode() {
	if p.p != nil && p.p["cachemode"] != nil {
		delete(p.p, "cachemode")
	}
}

func (p *UpdateDiskOfferingParams) GetCachemode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cachemode"].(string)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetDisplayoffering(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displayoffering"] = v
}

func (p *UpdateDiskOfferingParams) ResetDisplayoffering() {
	if p.p != nil && p.p["displayoffering"] != nil {
		delete(p.p, "displayoffering")
	}
}

func (p *UpdateDiskOfferingParams) GetDisplayoffering() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displayoffering"].(bool)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateDiskOfferingParams) ResetDisplaytext() {
	if p.p != nil && p.p["displaytext"] != nil {
		delete(p.p, "displaytext")
	}
}

func (p *UpdateDiskOfferingParams) GetDisplaytext() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["displaytext"].(string)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *UpdateDiskOfferingParams) ResetDomainid() {
	if p.p != nil && p.p["domainid"] != nil {
		delete(p.p, "domainid")
	}
}

func (p *UpdateDiskOfferingParams) GetDomainid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainid"].(string)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateDiskOfferingParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateDiskOfferingParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetIopsreadrate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadrate"] = v
}

func (p *UpdateDiskOfferingParams) ResetIopsreadrate() {
	if p.p != nil && p.p["iopsreadrate"] != nil {
		delete(p.p, "iopsreadrate")
	}
}

func (p *UpdateDiskOfferingParams) GetIopsreadrate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadrate"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetIopsreadratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadratemax"] = v
}

func (p *UpdateDiskOfferingParams) ResetIopsreadratemax() {
	if p.p != nil && p.p["iopsreadratemax"] != nil {
		delete(p.p, "iopsreadratemax")
	}
}

func (p *UpdateDiskOfferingParams) GetIopsreadratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadratemax"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetIopsreadratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadratemaxlength"] = v
}

func (p *UpdateDiskOfferingParams) ResetIopsreadratemaxlength() {
	if p.p != nil && p.p["iopsreadratemaxlength"] != nil {
		delete(p.p, "iopsreadratemaxlength")
	}
}

func (p *UpdateDiskOfferingParams) GetIopsreadratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsreadratemaxlength"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetIopswriterate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriterate"] = v
}

func (p *UpdateDiskOfferingParams) ResetIopswriterate() {
	if p.p != nil && p.p["iopswriterate"] != nil {
		delete(p.p, "iopswriterate")
	}
}

func (p *UpdateDiskOfferingParams) GetIopswriterate() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriterate"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetIopswriteratemax(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriteratemax"] = v
}

func (p *UpdateDiskOfferingParams) ResetIopswriteratemax() {
	if p.p != nil && p.p["iopswriteratemax"] != nil {
		delete(p.p, "iopswriteratemax")
	}
}

func (p *UpdateDiskOfferingParams) GetIopswriteratemax() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriteratemax"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetIopswriteratemaxlength(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriteratemaxlength"] = v
}

func (p *UpdateDiskOfferingParams) ResetIopswriteratemaxlength() {
	if p.p != nil && p.p["iopswriteratemaxlength"] != nil {
		delete(p.p, "iopswriteratemaxlength")
	}
}

func (p *UpdateDiskOfferingParams) GetIopswriteratemaxlength() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopswriteratemaxlength"].(int64)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateDiskOfferingParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateDiskOfferingParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

func (p *UpdateDiskOfferingParams) ResetSortkey() {
	if p.p != nil && p.p["sortkey"] != nil {
		delete(p.p, "sortkey")
	}
}

func (p *UpdateDiskOfferingParams) GetSortkey() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sortkey"].(int)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *UpdateDiskOfferingParams) ResetState() {
	if p.p != nil && p.p["state"] != nil {
		delete(p.p, "state")
	}
}

func (p *UpdateDiskOfferingParams) GetState() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["state"].(string)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *UpdateDiskOfferingParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *UpdateDiskOfferingParams) GetTags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(string)
	return value, ok
}

func (p *UpdateDiskOfferingParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *UpdateDiskOfferingParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *UpdateDiskOfferingParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewUpdateDiskOfferingParams(id string) *UpdateDiskOfferingParams {
	p := &UpdateDiskOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a disk offering.
func (s *DiskOfferingService) UpdateDiskOffering(p *UpdateDiskOfferingParams) (*UpdateDiskOfferingResponse, error) {
	resp, err := s.cs.newPostRequest("updateDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateDiskOfferingResponse struct {
	CacheMode                   string            `json:"cacheMode"`
	Created                     string            `json:"created"`
	Details                     map[string]string `json:"details"`
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
	Disksize                    int64             `json:"disksize"`
	Disksizestrictness          bool              `json:"disksizestrictness"`
	Displayoffering             bool              `json:"displayoffering"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Encrypt                     bool              `json:"encrypt"`
	Hasannotations              bool              `json:"hasannotations"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Maxiops                     int64             `json:"maxiops"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Provisioningtype            string            `json:"provisioningtype"`
	State                       string            `json:"state"`
	Storagetype                 string            `json:"storagetype"`
	Suitableforvirtualmachine   bool              `json:"suitableforvirtualmachine"`
	Tags                        string            `json:"tags"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}
