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

func TestVLANService(t *testing.T) {
	service := "VLANService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateVlanIpRange := func(t *testing.T) {
		if _, ok := response["createVlanIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VLAN.NewCreateVlanIpRangeParams()
		r, err := client.VLAN.CreateVlanIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVlanIpRange", testcreateVlanIpRange)

	testdedicateGuestVlanRange := func(t *testing.T) {
		if _, ok := response["dedicateGuestVlanRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VLAN.NewDedicateGuestVlanRangeParams("physicalnetworkid", "vlanrange")
		r, err := client.VLAN.DedicateGuestVlanRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DedicateGuestVlanRange", testdedicateGuestVlanRange)

	testdeleteVlanIpRange := func(t *testing.T) {
		if _, ok := response["deleteVlanIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VLAN.NewDeleteVlanIpRangeParams("id")
		_, err := client.VLAN.DeleteVlanIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVlanIpRange", testdeleteVlanIpRange)

	testlistDedicatedGuestVlanRanges := func(t *testing.T) {
		if _, ok := response["listDedicatedGuestVlanRanges"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VLAN.NewListDedicatedGuestVlanRangesParams()
		_, err := client.VLAN.ListDedicatedGuestVlanRanges(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListDedicatedGuestVlanRanges", testlistDedicatedGuestVlanRanges)

	testlistVlanIpRanges := func(t *testing.T) {
		if _, ok := response["listVlanIpRanges"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VLAN.NewListVlanIpRangesParams()
		_, err := client.VLAN.ListVlanIpRanges(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVlanIpRanges", testlistVlanIpRanges)

	testreleaseDedicatedGuestVlanRange := func(t *testing.T) {
		if _, ok := response["releaseDedicatedGuestVlanRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VLAN.NewReleaseDedicatedGuestVlanRangeParams("id")
		_, err := client.VLAN.ReleaseDedicatedGuestVlanRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReleaseDedicatedGuestVlanRange", testreleaseDedicatedGuestVlanRange)

	testlistGuestVlans := func(t *testing.T) {
		if _, ok := response["listGuestVlans"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VLAN.NewListGuestVlansParams()
		_, err := client.VLAN.ListGuestVlans(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListGuestVlans", testlistGuestVlans)

}
