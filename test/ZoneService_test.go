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

func TestZoneService(t *testing.T) {
	service := "ZoneService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateZone := func(t *testing.T) {
		if _, ok := response["createZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewCreateZoneParams("dns1", "internaldns1", "name", "networktype")
		r, err := client.Zone.CreateZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateZone", testcreateZone)

	testdedicateZone := func(t *testing.T) {
		if _, ok := response["dedicateZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewDedicateZoneParams("domainid", "zoneid")
		r, err := client.Zone.DedicateZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DedicateZone", testdedicateZone)

	testdeleteZone := func(t *testing.T) {
		if _, ok := response["deleteZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewDeleteZoneParams("id")
		_, err := client.Zone.DeleteZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteZone", testdeleteZone)

	testdisableOutOfBandManagementForZone := func(t *testing.T) {
		if _, ok := response["disableOutOfBandManagementForZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewDisableOutOfBandManagementForZoneParams("zoneid")
		_, err := client.Zone.DisableOutOfBandManagementForZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DisableOutOfBandManagementForZone", testdisableOutOfBandManagementForZone)

	testenableOutOfBandManagementForZone := func(t *testing.T) {
		if _, ok := response["enableOutOfBandManagementForZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewEnableOutOfBandManagementForZoneParams("zoneid")
		_, err := client.Zone.EnableOutOfBandManagementForZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("EnableOutOfBandManagementForZone", testenableOutOfBandManagementForZone)

	testdisableHAForZone := func(t *testing.T) {
		if _, ok := response["disableHAForZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewDisableHAForZoneParams("zoneid")
		_, err := client.Zone.DisableHAForZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DisableHAForZone", testdisableHAForZone)

	testenableHAForZone := func(t *testing.T) {
		if _, ok := response["enableHAForZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewEnableHAForZoneParams("zoneid")
		_, err := client.Zone.EnableHAForZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("EnableHAForZone", testenableHAForZone)

	testlistDedicatedZones := func(t *testing.T) {
		if _, ok := response["listDedicatedZones"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewListDedicatedZonesParams()
		_, err := client.Zone.ListDedicatedZones(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListDedicatedZones", testlistDedicatedZones)

	testlistZones := func(t *testing.T) {
		if _, ok := response["listZones"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewListZonesParams()
		_, err := client.Zone.ListZones(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListZones", testlistZones)

	testlistZonesMetrics := func(t *testing.T) {
		if _, ok := response["listZonesMetrics"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewListZonesMetricsParams()
		_, err := client.Zone.ListZonesMetrics(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListZonesMetrics", testlistZonesMetrics)

	testreleaseDedicatedZone := func(t *testing.T) {
		if _, ok := response["releaseDedicatedZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewReleaseDedicatedZoneParams("zoneid")
		_, err := client.Zone.ReleaseDedicatedZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReleaseDedicatedZone", testreleaseDedicatedZone)

	testupdateZone := func(t *testing.T) {
		if _, ok := response["updateZone"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Zone.NewUpdateZoneParams("id")
		r, err := client.Zone.UpdateZone(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateZone", testupdateZone)

}
