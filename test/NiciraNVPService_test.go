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

func TestNiciraNVPService(t *testing.T) {
	service := "NiciraNVPService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddNiciraNvpDevice := func(t *testing.T) {
		if _, ok := response["addNiciraNvpDevice"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NiciraNVP.NewAddNiciraNvpDeviceParams("hostname", "password", "physicalnetworkid", "transportzoneuuid", "username")
		_, err := client.NiciraNVP.AddNiciraNvpDevice(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddNiciraNvpDevice", testaddNiciraNvpDevice)

	testdeleteNiciraNvpDevice := func(t *testing.T) {
		if _, ok := response["deleteNiciraNvpDevice"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NiciraNVP.NewDeleteNiciraNvpDeviceParams("nvpdeviceid")
		_, err := client.NiciraNVP.DeleteNiciraNvpDevice(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNiciraNvpDevice", testdeleteNiciraNvpDevice)

	testlistNiciraNvpDevices := func(t *testing.T) {
		if _, ok := response["listNiciraNvpDevices"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NiciraNVP.NewListNiciraNvpDevicesParams()
		_, err := client.NiciraNVP.ListNiciraNvpDevices(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNiciraNvpDevices", testlistNiciraNvpDevices)

}
