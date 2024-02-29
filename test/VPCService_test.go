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

func TestVPCService(t *testing.T) {
	service := "VPCService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreatePrivateGateway := func(t *testing.T) {
		if _, ok := response["createPrivateGateway"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewCreatePrivateGatewayParams("gateway", "ipaddress", "netmask", "vpcid")
		r, err := client.VPC.CreatePrivateGateway(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreatePrivateGateway", testcreatePrivateGateway)

	testcreateStaticRoute := func(t *testing.T) {
		if _, ok := response["createStaticRoute"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewCreateStaticRouteParams("cidr", "gatewayid")
		r, err := client.VPC.CreateStaticRoute(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateStaticRoute", testcreateStaticRoute)

	testcreateVPC := func(t *testing.T) {
		if _, ok := response["createVPC"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewCreateVPCParams("cidr", "name", "vpcofferingid", "zoneid")
		r, err := client.VPC.CreateVPC(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVPC", testcreateVPC)

	testcreateVPCOffering := func(t *testing.T) {
		if _, ok := response["createVPCOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewCreateVPCOfferingParams("name", []string{})
		r, err := client.VPC.CreateVPCOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVPCOffering", testcreateVPCOffering)

	testdeletePrivateGateway := func(t *testing.T) {
		if _, ok := response["deletePrivateGateway"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewDeletePrivateGatewayParams("id")
		_, err := client.VPC.DeletePrivateGateway(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeletePrivateGateway", testdeletePrivateGateway)

	testdeleteStaticRoute := func(t *testing.T) {
		if _, ok := response["deleteStaticRoute"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewDeleteStaticRouteParams("id")
		_, err := client.VPC.DeleteStaticRoute(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteStaticRoute", testdeleteStaticRoute)

	testdeleteVPC := func(t *testing.T) {
		if _, ok := response["deleteVPC"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewDeleteVPCParams("id")
		_, err := client.VPC.DeleteVPC(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVPC", testdeleteVPC)

	testdeleteVPCOffering := func(t *testing.T) {
		if _, ok := response["deleteVPCOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewDeleteVPCOfferingParams("id")
		_, err := client.VPC.DeleteVPCOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVPCOffering", testdeleteVPCOffering)

	testlistPrivateGateways := func(t *testing.T) {
		if _, ok := response["listPrivateGateways"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewListPrivateGatewaysParams()
		_, err := client.VPC.ListPrivateGateways(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListPrivateGateways", testlistPrivateGateways)

	testlistStaticRoutes := func(t *testing.T) {
		if _, ok := response["listStaticRoutes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewListStaticRoutesParams()
		_, err := client.VPC.ListStaticRoutes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListStaticRoutes", testlistStaticRoutes)

	testlistVPCOfferings := func(t *testing.T) {
		if _, ok := response["listVPCOfferings"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewListVPCOfferingsParams()
		_, err := client.VPC.ListVPCOfferings(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVPCOfferings", testlistVPCOfferings)

	testlistVPCs := func(t *testing.T) {
		if _, ok := response["listVPCs"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewListVPCsParams()
		_, err := client.VPC.ListVPCs(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVPCs", testlistVPCs)

	testrestartVPC := func(t *testing.T) {
		if _, ok := response["restartVPC"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewRestartVPCParams("id")
		_, err := client.VPC.RestartVPC(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RestartVPC", testrestartVPC)

	testupdateVPC := func(t *testing.T) {
		if _, ok := response["updateVPC"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewUpdateVPCParams("id")
		r, err := client.VPC.UpdateVPC(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVPC", testupdateVPC)

	testupdateVPCOffering := func(t *testing.T) {
		if _, ok := response["updateVPCOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPC.NewUpdateVPCOfferingParams("id")
		r, err := client.VPC.UpdateVPCOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVPCOffering", testupdateVPCOffering)

}
