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

func TestHostService(t *testing.T) {
	service := "HostService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddBaremetalHost := func(t *testing.T) {
		if _, ok := response["addBaremetalHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewAddBaremetalHostParams("hypervisor", "podid", "url", "zoneid")
		r, err := client.Host.AddBaremetalHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddBaremetalHost", testaddBaremetalHost)

	testaddGloboDnsHost := func(t *testing.T) {
		if _, ok := response["addGloboDnsHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewAddGloboDnsHostParams("password", "physicalnetworkid", "url", "username")
		_, err := client.Host.AddGloboDnsHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddGloboDnsHost", testaddGloboDnsHost)

	testaddHost := func(t *testing.T) {
		if _, ok := response["addHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewAddHostParams("hypervisor", "podid", "url", "zoneid")
		r, err := client.Host.AddHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddHost", testaddHost)

	testaddSecondaryStorage := func(t *testing.T) {
		if _, ok := response["addSecondaryStorage"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewAddSecondaryStorageParams("url")
		r, err := client.Host.AddSecondaryStorage(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddSecondaryStorage", testaddSecondaryStorage)

	testcancelHostMaintenance := func(t *testing.T) {
		if _, ok := response["cancelHostMaintenance"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewCancelHostMaintenanceParams("id")
		r, err := client.Host.CancelHostMaintenance(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CancelHostMaintenance", testcancelHostMaintenance)

	testconfigureHAForHost := func(t *testing.T) {
		if _, ok := response["configureHAForHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewConfigureHAForHostParams("hostid", "provider")
		_, err := client.Host.ConfigureHAForHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ConfigureHAForHost", testconfigureHAForHost)

	testenableHAForHost := func(t *testing.T) {
		if _, ok := response["enableHAForHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewEnableHAForHostParams("hostid")
		_, err := client.Host.EnableHAForHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("EnableHAForHost", testenableHAForHost)

	testdedicateHost := func(t *testing.T) {
		if _, ok := response["dedicateHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewDedicateHostParams("domainid", "hostid")
		r, err := client.Host.DedicateHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DedicateHost", testdedicateHost)

	testdeleteHost := func(t *testing.T) {
		if _, ok := response["deleteHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewDeleteHostParams("id")
		_, err := client.Host.DeleteHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteHost", testdeleteHost)

	testdisableOutOfBandManagementForHost := func(t *testing.T) {
		if _, ok := response["disableOutOfBandManagementForHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewDisableOutOfBandManagementForHostParams("hostid")
		_, err := client.Host.DisableOutOfBandManagementForHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DisableOutOfBandManagementForHost", testdisableOutOfBandManagementForHost)

	testenableOutOfBandManagementForHost := func(t *testing.T) {
		if _, ok := response["enableOutOfBandManagementForHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewEnableOutOfBandManagementForHostParams("hostid")
		_, err := client.Host.EnableOutOfBandManagementForHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("EnableOutOfBandManagementForHost", testenableOutOfBandManagementForHost)

	testfindHostsForMigration := func(t *testing.T) {
		if _, ok := response["findHostsForMigration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewFindHostsForMigrationParams("virtualmachineid")
		r, err := client.Host.FindHostsForMigration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if len(r.Host) != 1 {
			t.Errorf("Failed to parse response. Host is missing")
		}
		if r.Host[0].Id != "6a5b8975-225f-4630-9093-3d55cec439d8" {
			t.Errorf("Failed to parse response. ID incorrect")
		}
	}
	t.Run("FindHostsForMigration", testfindHostsForMigration)

	testlistDedicatedHosts := func(t *testing.T) {
		if _, ok := response["listDedicatedHosts"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewListDedicatedHostsParams()
		_, err := client.Host.ListDedicatedHosts(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListDedicatedHosts", testlistDedicatedHosts)

	testlistHostTags := func(t *testing.T) {
		if _, ok := response["listHostTags"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewListHostTagsParams()
		_, err := client.Host.ListHostTags(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListHostTags", testlistHostTags)

	testlistHosts := func(t *testing.T) {
		if _, ok := response["listHosts"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewListHostsParams()
		_, err := client.Host.ListHosts(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListHosts", testlistHosts)

	testlistHostsMetrics := func(t *testing.T) {
		if _, ok := response["listHostsMetrics"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewListHostsMetricsParams()
		_, err := client.Host.ListHostsMetrics(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListHostsMetrics", testlistHostsMetrics)

	testprepareHostForMaintenance := func(t *testing.T) {
		if _, ok := response["prepareHostForMaintenance"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewPrepareHostForMaintenanceParams("id")
		r, err := client.Host.PrepareHostForMaintenance(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("PrepareHostForMaintenance", testprepareHostForMaintenance)

	testreconnectHost := func(t *testing.T) {
		if _, ok := response["reconnectHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewReconnectHostParams("id")
		r, err := client.Host.ReconnectHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ReconnectHost", testreconnectHost)

	testreleaseDedicatedHost := func(t *testing.T) {
		if _, ok := response["releaseDedicatedHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewReleaseDedicatedHostParams("hostid")
		_, err := client.Host.ReleaseDedicatedHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReleaseDedicatedHost", testreleaseDedicatedHost)

	testreleaseHostReservation := func(t *testing.T) {
		if _, ok := response["releaseHostReservation"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewReleaseHostReservationParams("id")
		_, err := client.Host.ReleaseHostReservation(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReleaseHostReservation", testreleaseHostReservation)

	testupdateHost := func(t *testing.T) {
		if _, ok := response["updateHost"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewUpdateHostParams("id")
		r, err := client.Host.UpdateHost(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateHost", testupdateHost)

	testupdateHostPassword := func(t *testing.T) {
		if _, ok := response["updateHostPassword"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Host.NewUpdateHostPasswordParams("password", "username")
		_, err := client.Host.UpdateHostPassword(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateHostPassword", testupdateHostPassword)

}
