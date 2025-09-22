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

func TestGPUService(t *testing.T) {
	service := "GPUService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateGpuCard := func(t *testing.T) {
		if _, ok := response["createGpuCard"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewCreateGpuCardParams("deviceid", "devicename", "name", "vendorid", "vendorname")
		r, err := client.GPU.CreateGpuCard(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateGpuCard", testcreateGpuCard)

	testcreateGpuDevice := func(t *testing.T) {
		if _, ok := response["createGpuDevice"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewCreateGpuDeviceParams("busaddress", "gpucardid", "hostid", "vgpuprofileid")
		r, err := client.GPU.CreateGpuDevice(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateGpuDevice", testcreateGpuDevice)

	testcreateVgpuProfile := func(t *testing.T) {
		if _, ok := response["createVgpuProfile"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewCreateVgpuProfileParams("gpucardid", "name")
		r, err := client.GPU.CreateVgpuProfile(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVgpuProfile", testcreateVgpuProfile)

	testdeleteGpuCard := func(t *testing.T) {
		if _, ok := response["deleteGpuCard"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewDeleteGpuCardParams("id")
		_, err := client.GPU.DeleteGpuCard(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteGpuCard", testdeleteGpuCard)

	testdeleteGpuDevice := func(t *testing.T) {
		if _, ok := response["deleteGpuDevice"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewDeleteGpuDeviceParams([]string{})
		_, err := client.GPU.DeleteGpuDevice(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteGpuDevice", testdeleteGpuDevice)

	testdeleteVgpuProfile := func(t *testing.T) {
		if _, ok := response["deleteVgpuProfile"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewDeleteVgpuProfileParams("id")
		_, err := client.GPU.DeleteVgpuProfile(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVgpuProfile", testdeleteVgpuProfile)

	testdiscoverGpuDevices := func(t *testing.T) {
		if _, ok := response["discoverGpuDevices"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewDiscoverGpuDevicesParams("id")
		r, err := client.GPU.DiscoverGpuDevices(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DiscoverGpuDevices", testdiscoverGpuDevices)

	testlistGpuCards := func(t *testing.T) {
		if _, ok := response["listGpuCards"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewListGpuCardsParams()
		_, err := client.GPU.ListGpuCards(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListGpuCards", testlistGpuCards)

	testlistGpuDevices := func(t *testing.T) {
		if _, ok := response["listGpuDevices"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewListGpuDevicesParams()
		_, err := client.GPU.ListGpuDevices(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListGpuDevices", testlistGpuDevices)

	testlistVgpuProfiles := func(t *testing.T) {
		if _, ok := response["listVgpuProfiles"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewListVgpuProfilesParams()
		_, err := client.GPU.ListVgpuProfiles(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVgpuProfiles", testlistVgpuProfiles)

	testmanageGpuDevice := func(t *testing.T) {
		if _, ok := response["manageGpuDevice"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewManageGpuDeviceParams([]string{})
		_, err := client.GPU.ManageGpuDevice(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ManageGpuDevice", testmanageGpuDevice)

	testunmanageGpuDevice := func(t *testing.T) {
		if _, ok := response["unmanageGpuDevice"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewUnmanageGpuDeviceParams([]string{})
		_, err := client.GPU.UnmanageGpuDevice(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UnmanageGpuDevice", testunmanageGpuDevice)

	testupdateGpuCard := func(t *testing.T) {
		if _, ok := response["updateGpuCard"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewUpdateGpuCardParams("id")
		r, err := client.GPU.UpdateGpuCard(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateGpuCard", testupdateGpuCard)

	testupdateGpuDevice := func(t *testing.T) {
		if _, ok := response["updateGpuDevice"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewUpdateGpuDeviceParams("id")
		r, err := client.GPU.UpdateGpuDevice(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateGpuDevice", testupdateGpuDevice)

	testupdateVgpuProfile := func(t *testing.T) {
		if _, ok := response["updateVgpuProfile"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.GPU.NewUpdateVgpuProfileParams("id")
		r, err := client.GPU.UpdateVgpuProfile(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVgpuProfile", testupdateVgpuProfile)

}
