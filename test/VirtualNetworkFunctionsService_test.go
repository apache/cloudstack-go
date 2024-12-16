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

func TestVirtualNetworkFunctionsService(t *testing.T) {
	service := "VirtualNetworkFunctionsService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testdeleteVnfTemplate := func(t *testing.T) {
		if _, ok := response["deleteVnfTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualNetworkFunctions.NewDeleteVnfTemplateParams("id")
		_, err := client.VirtualNetworkFunctions.DeleteVnfTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVnfTemplate", testdeleteVnfTemplate)

	testdeployVnfAppliance := func(t *testing.T) {
		if _, ok := response["deployVnfAppliance"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualNetworkFunctions.NewDeployVnfApplianceParams("serviceofferingid", "templateid", "zoneid")
		r, err := client.VirtualNetworkFunctions.DeployVnfAppliance(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DeployVnfAppliance", testdeployVnfAppliance)

	testlistVnfAppliances := func(t *testing.T) {
		if _, ok := response["listVnfAppliances"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualNetworkFunctions.NewListVnfAppliancesParams()
		_, err := client.VirtualNetworkFunctions.ListVnfAppliances(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVnfAppliances", testlistVnfAppliances)

	testlistVnfTemplates := func(t *testing.T) {
		if _, ok := response["listVnfTemplates"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualNetworkFunctions.NewListVnfTemplatesParams("templatefilter")
		_, err := client.VirtualNetworkFunctions.ListVnfTemplates(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVnfTemplates", testlistVnfTemplates)

	testregisterVnfTemplate := func(t *testing.T) {
		if _, ok := response["registerVnfTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualNetworkFunctions.NewRegisterVnfTemplateParams("format", "hypervisor", "name", "url")
		r, err := client.VirtualNetworkFunctions.RegisterVnfTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RegisterVnfTemplate", testregisterVnfTemplate)

	testupdateVnfTemplate := func(t *testing.T) {
		if _, ok := response["updateVnfTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualNetworkFunctions.NewUpdateVnfTemplateParams("id")
		r, err := client.VirtualNetworkFunctions.UpdateVnfTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVnfTemplate", testupdateVnfTemplate)

}
