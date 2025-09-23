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

func TestConfigurationService(t *testing.T) {
	service := "ConfigurationService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testlistCapabilities := func(t *testing.T) {
		if _, ok := response["listCapabilities"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewListCapabilitiesParams()
		_, err := client.Configuration.ListCapabilities(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCapabilities", testlistCapabilities)

	testlistConfigurationGroups := func(t *testing.T) {
		if _, ok := response["listConfigurationGroups"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewListConfigurationGroupsParams()
		_, err := client.Configuration.ListConfigurationGroups(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListConfigurationGroups", testlistConfigurationGroups)

	testlistConfigurations := func(t *testing.T) {
		if _, ok := response["listConfigurations"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewListConfigurationsParams()
		_, err := client.Configuration.ListConfigurations(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListConfigurations", testlistConfigurations)

	testlistDeploymentPlanners := func(t *testing.T) {
		if _, ok := response["listDeploymentPlanners"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewListDeploymentPlannersParams()
		_, err := client.Configuration.ListDeploymentPlanners(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListDeploymentPlanners", testlistDeploymentPlanners)

	testupdateConfiguration := func(t *testing.T) {
		if _, ok := response["updateConfiguration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewUpdateConfigurationParams("name")
		_, err := client.Configuration.UpdateConfiguration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateConfiguration", testupdateConfiguration)

	testresetConfiguration := func(t *testing.T) {
		if _, ok := response["resetConfiguration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewResetConfigurationParams("name")
		_, err := client.Configuration.ResetConfiguration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ResetConfiguration", testresetConfiguration)

	testupdateStorageCapabilities := func(t *testing.T) {
		if _, ok := response["updateStorageCapabilities"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewUpdateStorageCapabilitiesParams("id")
		r, err := client.Configuration.UpdateStorageCapabilities(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateStorageCapabilities", testupdateStorageCapabilities)

	testregisterCniConfiguration := func(t *testing.T) {
		if _, ok := response["registerCniConfiguration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewRegisterCniConfigurationParams("name")
		_, err := client.Configuration.RegisterCniConfiguration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RegisterCniConfiguration", testregisterCniConfiguration)

	testlistCniConfiguration := func(t *testing.T) {
		if _, ok := response["listCniConfiguration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewListCniConfigurationParams()
		_, err := client.Configuration.ListCniConfiguration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCniConfiguration", testlistCniConfiguration)

	testdeleteCniConfiguration := func(t *testing.T) {
		if _, ok := response["deleteCniConfiguration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Configuration.NewDeleteCniConfigurationParams("id")
		_, err := client.Configuration.DeleteCniConfiguration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteCniConfiguration", testdeleteCniConfiguration)

}
