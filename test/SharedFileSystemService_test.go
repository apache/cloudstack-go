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

func TestSharedFileSystemService(t *testing.T) {
	service := "SharedFileSystemService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testchangeSharedFileSystemDiskOffering := func(t *testing.T) {
		if _, ok := response["changeSharedFileSystemDiskOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewChangeSharedFileSystemDiskOfferingParams("id")
		r, err := client.SharedFileSystem.ChangeSharedFileSystemDiskOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ChangeSharedFileSystemDiskOffering", testchangeSharedFileSystemDiskOffering)

	testchangeSharedFileSystemServiceOffering := func(t *testing.T) {
		if _, ok := response["changeSharedFileSystemServiceOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewChangeSharedFileSystemServiceOfferingParams("id", "serviceofferingid")
		r, err := client.SharedFileSystem.ChangeSharedFileSystemServiceOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ChangeSharedFileSystemServiceOffering", testchangeSharedFileSystemServiceOffering)

	testcreateSharedFileSystem := func(t *testing.T) {
		if _, ok := response["createSharedFileSystem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewCreateSharedFileSystemParams("diskofferingid", "filesystem", "name", "networkid", "serviceofferingid", "zoneid")
		r, err := client.SharedFileSystem.CreateSharedFileSystem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateSharedFileSystem", testcreateSharedFileSystem)

	testdestroySharedFileSystem := func(t *testing.T) {
		if _, ok := response["destroySharedFileSystem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewDestroySharedFileSystemParams()
		_, err := client.SharedFileSystem.DestroySharedFileSystem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DestroySharedFileSystem", testdestroySharedFileSystem)

	testexpungeSharedFileSystem := func(t *testing.T) {
		if _, ok := response["expungeSharedFileSystem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewExpungeSharedFileSystemParams()
		_, err := client.SharedFileSystem.ExpungeSharedFileSystem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ExpungeSharedFileSystem", testexpungeSharedFileSystem)

	testlistSharedFileSystemProviders := func(t *testing.T) {
		if _, ok := response["listSharedFileSystemProviders"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewListSharedFileSystemProvidersParams()
		_, err := client.SharedFileSystem.ListSharedFileSystemProviders(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSharedFileSystemProviders", testlistSharedFileSystemProviders)

	testlistSharedFileSystems := func(t *testing.T) {
		if _, ok := response["listSharedFileSystems"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewListSharedFileSystemsParams()
		_, err := client.SharedFileSystem.ListSharedFileSystems(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSharedFileSystems", testlistSharedFileSystems)

	testrecoverSharedFileSystem := func(t *testing.T) {
		if _, ok := response["recoverSharedFileSystem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewRecoverSharedFileSystemParams()
		_, err := client.SharedFileSystem.RecoverSharedFileSystem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RecoverSharedFileSystem", testrecoverSharedFileSystem)

	testrestartSharedFileSystem := func(t *testing.T) {
		if _, ok := response["restartSharedFileSystem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewRestartSharedFileSystemParams("id")
		_, err := client.SharedFileSystem.RestartSharedFileSystem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RestartSharedFileSystem", testrestartSharedFileSystem)

	teststartSharedFileSystem := func(t *testing.T) {
		if _, ok := response["startSharedFileSystem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewStartSharedFileSystemParams("id")
		r, err := client.SharedFileSystem.StartSharedFileSystem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("StartSharedFileSystem", teststartSharedFileSystem)

	teststopSharedFileSystem := func(t *testing.T) {
		if _, ok := response["stopSharedFileSystem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewStopSharedFileSystemParams("id")
		r, err := client.SharedFileSystem.StopSharedFileSystem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("StopSharedFileSystem", teststopSharedFileSystem)

	testupdateSharedFileSystem := func(t *testing.T) {
		if _, ok := response["updateSharedFileSystem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SharedFileSystem.NewUpdateSharedFileSystemParams("id")
		r, err := client.SharedFileSystem.UpdateSharedFileSystem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateSharedFileSystem", testupdateSharedFileSystem)

}
