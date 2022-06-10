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

package test

import (
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestGuestOSService(t *testing.T) {
	service := "GuestOSService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddGuestOs := func(t *testing.T) {
		if _, ok := response["addGuestOs"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewAddGuestOsParams(map[string]string{}, "oscategoryid", "osdisplayname")
		r, err := client.GuestOS.AddGuestOs(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddGuestOs", testaddGuestOs)

	testaddGuestOsMapping := func(t *testing.T) {
		if _, ok := response["addGuestOsMapping"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewAddGuestOsMappingParams("hypervisor", "hypervisorversion", "osnameforhypervisor")
		r, err := client.GuestOS.AddGuestOsMapping(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddGuestOsMapping", testaddGuestOsMapping)

	testlistGuestOsMapping := func(t *testing.T) {
		if _, ok := response["listGuestOsMapping"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewListGuestOsMappingParams()
		_, err := client.GuestOS.ListGuestOsMapping(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListGuestOsMapping", testlistGuestOsMapping)

	testlistOsCategories := func(t *testing.T) {
		if _, ok := response["listOsCategories"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewListOsCategoriesParams()
		_, err := client.GuestOS.ListOsCategories(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListOsCategories", testlistOsCategories)

	testlistOsTypes := func(t *testing.T) {
		if _, ok := response["listOsTypes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewListOsTypesParams()
		_, err := client.GuestOS.ListOsTypes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListOsTypes", testlistOsTypes)

	testremoveGuestOs := func(t *testing.T) {
		if _, ok := response["removeGuestOs"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewRemoveGuestOsParams("id")
		_, err := client.GuestOS.RemoveGuestOs(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveGuestOs", testremoveGuestOs)

	testremoveGuestOsMapping := func(t *testing.T) {
		if _, ok := response["removeGuestOsMapping"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewRemoveGuestOsMappingParams("id")
		_, err := client.GuestOS.RemoveGuestOsMapping(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveGuestOsMapping", testremoveGuestOsMapping)

	testupdateGuestOs := func(t *testing.T) {
		if _, ok := response["updateGuestOs"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewUpdateGuestOsParams(map[string]string{}, "id", "osdisplayname")
		r, err := client.GuestOS.UpdateGuestOs(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateGuestOs", testupdateGuestOs)

	testupdateGuestOsMapping := func(t *testing.T) {
		if _, ok := response["updateGuestOsMapping"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GuestOS.NewUpdateGuestOsMappingParams("id", "osnameforhypervisor")
		r, err := client.GuestOS.UpdateGuestOsMapping(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateGuestOsMapping", testupdateGuestOsMapping)

}
