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

func TestBaremetalService(t *testing.T) {
	service := "BaremetalService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddBaremetalDhcp := func(t *testing.T) {
		if _, ok := response["addBaremetalDhcp"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewAddBaremetalDhcpParams("dhcpservertype", "password", "physicalnetworkid", "url", "username")
		r, err := client.Baremetal.AddBaremetalDhcp(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddBaremetalDhcp", testaddBaremetalDhcp)

	testaddBaremetalPxeKickStartServer := func(t *testing.T) {
		if _, ok := response["addBaremetalPxeKickStartServer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewAddBaremetalPxeKickStartServerParams("password", "physicalnetworkid", "pxeservertype", "tftpdir", "url", "username")
		r, err := client.Baremetal.AddBaremetalPxeKickStartServer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddBaremetalPxeKickStartServer", testaddBaremetalPxeKickStartServer)

	testaddBaremetalPxePingServer := func(t *testing.T) {
		if _, ok := response["addBaremetalPxePingServer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewAddBaremetalPxePingServerParams("password", "physicalnetworkid", "pingdir", "pingstorageserverip", "pxeservertype", "tftpdir", "url", "username")
		r, err := client.Baremetal.AddBaremetalPxePingServer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddBaremetalPxePingServer", testaddBaremetalPxePingServer)

	testaddBaremetalRct := func(t *testing.T) {
		if _, ok := response["addBaremetalRct"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewAddBaremetalRctParams("baremetalrcturl")
		r, err := client.Baremetal.AddBaremetalRct(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddBaremetalRct", testaddBaremetalRct)

	testdeleteBaremetalRct := func(t *testing.T) {
		if _, ok := response["deleteBaremetalRct"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewDeleteBaremetalRctParams("id")
		_, err := client.Baremetal.DeleteBaremetalRct(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteBaremetalRct", testdeleteBaremetalRct)

	testlistBaremetalDhcp := func(t *testing.T) {
		if _, ok := response["listBaremetalDhcp"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewListBaremetalDhcpParams("physicalnetworkid")
		_, err := client.Baremetal.ListBaremetalDhcp(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBaremetalDhcp", testlistBaremetalDhcp)

	testlistBaremetalPxeServers := func(t *testing.T) {
		if _, ok := response["listBaremetalPxeServers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewListBaremetalPxeServersParams("physicalnetworkid")
		_, err := client.Baremetal.ListBaremetalPxeServers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBaremetalPxeServers", testlistBaremetalPxeServers)

	testlistBaremetalRct := func(t *testing.T) {
		if _, ok := response["listBaremetalRct"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewListBaremetalRctParams()
		_, err := client.Baremetal.ListBaremetalRct(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBaremetalRct", testlistBaremetalRct)

	testnotifyBaremetalProvisionDone := func(t *testing.T) {
		if _, ok := response["notifyBaremetalProvisionDone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Baremetal.NewNotifyBaremetalProvisionDoneParams("mac")
		_, err := client.Baremetal.NotifyBaremetalProvisionDone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("NotifyBaremetalProvisionDone", testnotifyBaremetalProvisionDone)

}
