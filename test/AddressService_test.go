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

	"github.com/sbrueseke/cloudstack-go/v2/cloudstack"
)

func TestAddressService(t *testing.T) {
	service := "AddressService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testassociateIpAddress := func(t *testing.T) {
		if _, ok := response["associateIpAddress"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Address.NewAssociateIpAddressParams()
		r, err := client.Address.AssociateIpAddress(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AssociateIpAddress", testassociateIpAddress)

	testdisassociateIpAddress := func(t *testing.T) {
		if _, ok := response["disassociateIpAddress"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Address.NewDisassociateIpAddressParams()
		_, err := client.Address.DisassociateIpAddress(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DisassociateIpAddress", testdisassociateIpAddress)

	testlistPublicIpAddresses := func(t *testing.T) {
		if _, ok := response["listPublicIpAddresses"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Address.NewListPublicIpAddressesParams()
		_, err := client.Address.ListPublicIpAddresses(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListPublicIpAddresses", testlistPublicIpAddresses)

	testupdateIpAddress := func(t *testing.T) {
		if _, ok := response["updateIpAddress"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Address.NewUpdateIpAddressParams("id")
		r, err := client.Address.UpdateIpAddress(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateIpAddress", testupdateIpAddress)

	testreleaseIpAddress := func(t *testing.T) {
		if _, ok := response["releaseIpAddress"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Address.NewReleaseIpAddressParams("id")
		_, err := client.Address.ReleaseIpAddress(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReleaseIpAddress", testreleaseIpAddress)

}
