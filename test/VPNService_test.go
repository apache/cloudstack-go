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

func TestVPNService(t *testing.T) {
	service := "VPNService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddVpnUser := func(t *testing.T) {
		if _, ok := response["addVpnUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewAddVpnUserParams("password", "username")
		r, err := client.VPN.AddVpnUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddVpnUser", testaddVpnUser)

	testcreateRemoteAccessVpn := func(t *testing.T) {
		if _, ok := response["createRemoteAccessVpn"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewCreateRemoteAccessVpnParams("publicipid")
		r, err := client.VPN.CreateRemoteAccessVpn(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateRemoteAccessVpn", testcreateRemoteAccessVpn)

	testcreateVpnConnection := func(t *testing.T) {
		if _, ok := response["createVpnConnection"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewCreateVpnConnectionParams("s2scustomergatewayid", "s2svpngatewayid")
		r, err := client.VPN.CreateVpnConnection(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVpnConnection", testcreateVpnConnection)

	testcreateVpnCustomerGateway := func(t *testing.T) {
		if _, ok := response["createVpnCustomerGateway"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewCreateVpnCustomerGatewayParams("cidrlist", "esppolicy", "gateway", "ikepolicy", "ipsecpsk")
		r, err := client.VPN.CreateVpnCustomerGateway(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVpnCustomerGateway", testcreateVpnCustomerGateway)

	testcreateVpnGateway := func(t *testing.T) {
		if _, ok := response["createVpnGateway"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewCreateVpnGatewayParams("vpcid")
		r, err := client.VPN.CreateVpnGateway(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVpnGateway", testcreateVpnGateway)

	testdeleteRemoteAccessVpn := func(t *testing.T) {
		if _, ok := response["deleteRemoteAccessVpn"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewDeleteRemoteAccessVpnParams("publicipid")
		_, err := client.VPN.DeleteRemoteAccessVpn(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteRemoteAccessVpn", testdeleteRemoteAccessVpn)

	testdeleteVpnConnection := func(t *testing.T) {
		if _, ok := response["deleteVpnConnection"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewDeleteVpnConnectionParams("id")
		_, err := client.VPN.DeleteVpnConnection(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVpnConnection", testdeleteVpnConnection)

	testdeleteVpnCustomerGateway := func(t *testing.T) {
		if _, ok := response["deleteVpnCustomerGateway"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewDeleteVpnCustomerGatewayParams("id")
		_, err := client.VPN.DeleteVpnCustomerGateway(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVpnCustomerGateway", testdeleteVpnCustomerGateway)

	testdeleteVpnGateway := func(t *testing.T) {
		if _, ok := response["deleteVpnGateway"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewDeleteVpnGatewayParams("id")
		_, err := client.VPN.DeleteVpnGateway(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVpnGateway", testdeleteVpnGateway)

	testlistRemoteAccessVpns := func(t *testing.T) {
		if _, ok := response["listRemoteAccessVpns"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewListRemoteAccessVpnsParams()
		_, err := client.VPN.ListRemoteAccessVpns(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListRemoteAccessVpns", testlistRemoteAccessVpns)

	testlistVpnConnections := func(t *testing.T) {
		if _, ok := response["listVpnConnections"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewListVpnConnectionsParams()
		_, err := client.VPN.ListVpnConnections(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVpnConnections", testlistVpnConnections)

	testlistVpnCustomerGateways := func(t *testing.T) {
		if _, ok := response["listVpnCustomerGateways"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewListVpnCustomerGatewaysParams()
		_, err := client.VPN.ListVpnCustomerGateways(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVpnCustomerGateways", testlistVpnCustomerGateways)

	testlistVpnGateways := func(t *testing.T) {
		if _, ok := response["listVpnGateways"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewListVpnGatewaysParams()
		_, err := client.VPN.ListVpnGateways(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVpnGateways", testlistVpnGateways)

	testlistVpnUsers := func(t *testing.T) {
		if _, ok := response["listVpnUsers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewListVpnUsersParams()
		_, err := client.VPN.ListVpnUsers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVpnUsers", testlistVpnUsers)

	testremoveVpnUser := func(t *testing.T) {
		if _, ok := response["removeVpnUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewRemoveVpnUserParams("username")
		_, err := client.VPN.RemoveVpnUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveVpnUser", testremoveVpnUser)

	testresetVpnConnection := func(t *testing.T) {
		if _, ok := response["resetVpnConnection"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewResetVpnConnectionParams("id")
		r, err := client.VPN.ResetVpnConnection(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ResetVpnConnection", testresetVpnConnection)

	testupdateRemoteAccessVpn := func(t *testing.T) {
		if _, ok := response["updateRemoteAccessVpn"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewUpdateRemoteAccessVpnParams("id")
		r, err := client.VPN.UpdateRemoteAccessVpn(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateRemoteAccessVpn", testupdateRemoteAccessVpn)

	testupdateVpnConnection := func(t *testing.T) {
		if _, ok := response["updateVpnConnection"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewUpdateVpnConnectionParams("id")
		r, err := client.VPN.UpdateVpnConnection(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVpnConnection", testupdateVpnConnection)

	testupdateVpnCustomerGateway := func(t *testing.T) {
		if _, ok := response["updateVpnCustomerGateway"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewUpdateVpnCustomerGatewayParams("cidrlist", "esppolicy", "gateway", "id", "ikepolicy", "ipsecpsk")
		r, err := client.VPN.UpdateVpnCustomerGateway(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVpnCustomerGateway", testupdateVpnCustomerGateway)

	testupdateVpnGateway := func(t *testing.T) {
		if _, ok := response["updateVpnGateway"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VPN.NewUpdateVpnGatewayParams("id")
		r, err := client.VPN.UpdateVpnGateway(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVpnGateway", testupdateVpnGateway)

}
