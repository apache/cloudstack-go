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

func TestNetworkService(t *testing.T) {
	service := "NetworkService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddNetworkServiceProvider := func(t *testing.T) {
		if _, ok := response["addNetworkServiceProvider"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewAddNetworkServiceProviderParams("name", "physicalnetworkid")
		r, err := client.Network.AddNetworkServiceProvider(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddNetworkServiceProvider", testaddNetworkServiceProvider)

	testaddOpenDaylightController := func(t *testing.T) {
		if _, ok := response["addOpenDaylightController"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewAddOpenDaylightControllerParams("password", "physicalnetworkid", "url", "username")
		r, err := client.Network.AddOpenDaylightController(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddOpenDaylightController", testaddOpenDaylightController)

	testcreateNetwork := func(t *testing.T) {
		if _, ok := response["createNetwork"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewCreateNetworkParams("displaytext", "name", "networkofferingid", "zoneid")
		r, err := client.Network.CreateNetwork(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateNetwork", testcreateNetwork)

	testcreatePhysicalNetwork := func(t *testing.T) {
		if _, ok := response["createPhysicalNetwork"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewCreatePhysicalNetworkParams("name", "zoneid")
		r, err := client.Network.CreatePhysicalNetwork(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreatePhysicalNetwork", testcreatePhysicalNetwork)

	testcreateServiceInstance := func(t *testing.T) {
		if _, ok := response["createServiceInstance"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewCreateServiceInstanceParams("leftnetworkid", "name", "rightnetworkid", "serviceofferingid", "templateid", "zoneid")
		r, err := client.Network.CreateServiceInstance(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateServiceInstance", testcreateServiceInstance)

	testcreateStorageNetworkIpRange := func(t *testing.T) {
		if _, ok := response["createStorageNetworkIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewCreateStorageNetworkIpRangeParams("gateway", "netmask", "podid", "startip")
		r, err := client.Network.CreateStorageNetworkIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateStorageNetworkIpRange", testcreateStorageNetworkIpRange)

	testdedicatePublicIpRange := func(t *testing.T) {
		if _, ok := response["dedicatePublicIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewDedicatePublicIpRangeParams("domainid", "id")
		r, err := client.Network.DedicatePublicIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DedicatePublicIpRange", testdedicatePublicIpRange)

	testdeleteNetwork := func(t *testing.T) {
		if _, ok := response["deleteNetwork"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewDeleteNetworkParams("id")
		_, err := client.Network.DeleteNetwork(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNetwork", testdeleteNetwork)

	testdeleteNetworkServiceProvider := func(t *testing.T) {
		if _, ok := response["deleteNetworkServiceProvider"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewDeleteNetworkServiceProviderParams("id")
		_, err := client.Network.DeleteNetworkServiceProvider(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNetworkServiceProvider", testdeleteNetworkServiceProvider)

	testdeleteOpenDaylightController := func(t *testing.T) {
		if _, ok := response["deleteOpenDaylightController"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewDeleteOpenDaylightControllerParams("id")
		r, err := client.Network.DeleteOpenDaylightController(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DeleteOpenDaylightController", testdeleteOpenDaylightController)

	testdeletePhysicalNetwork := func(t *testing.T) {
		if _, ok := response["deletePhysicalNetwork"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewDeletePhysicalNetworkParams("id")
		_, err := client.Network.DeletePhysicalNetwork(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeletePhysicalNetwork", testdeletePhysicalNetwork)

	testdeleteStorageNetworkIpRange := func(t *testing.T) {
		if _, ok := response["deleteStorageNetworkIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewDeleteStorageNetworkIpRangeParams("id")
		_, err := client.Network.DeleteStorageNetworkIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteStorageNetworkIpRange", testdeleteStorageNetworkIpRange)

	testlistNetscalerLoadBalancerNetworks := func(t *testing.T) {
		if _, ok := response["listNetscalerLoadBalancerNetworks"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListNetscalerLoadBalancerNetworksParams("lbdeviceid")
		_, err := client.Network.ListNetscalerLoadBalancerNetworks(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetscalerLoadBalancerNetworks", testlistNetscalerLoadBalancerNetworks)

	testlistNetworkIsolationMethods := func(t *testing.T) {
		if _, ok := response["listNetworkIsolationMethods"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListNetworkIsolationMethodsParams()
		_, err := client.Network.ListNetworkIsolationMethods(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetworkIsolationMethods", testlistNetworkIsolationMethods)

	testlistNetworkServiceProviders := func(t *testing.T) {
		if _, ok := response["listNetworkServiceProviders"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListNetworkServiceProvidersParams()
		_, err := client.Network.ListNetworkServiceProviders(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetworkServiceProviders", testlistNetworkServiceProviders)

	testlistNetworks := func(t *testing.T) {
		if _, ok := response["listNetworks"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListNetworksParams()
		_, err := client.Network.ListNetworks(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetworks", testlistNetworks)

	testlistNiciraNvpDeviceNetworks := func(t *testing.T) {
		if _, ok := response["listNiciraNvpDeviceNetworks"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListNiciraNvpDeviceNetworksParams("nvpdeviceid")
		_, err := client.Network.ListNiciraNvpDeviceNetworks(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNiciraNvpDeviceNetworks", testlistNiciraNvpDeviceNetworks)

	testlistOpenDaylightControllers := func(t *testing.T) {
		if _, ok := response["listOpenDaylightControllers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListOpenDaylightControllersParams()
		_, err := client.Network.ListOpenDaylightControllers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListOpenDaylightControllers", testlistOpenDaylightControllers)

	testlistPaloAltoFirewallNetworks := func(t *testing.T) {
		if _, ok := response["listPaloAltoFirewallNetworks"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListPaloAltoFirewallNetworksParams("lbdeviceid")
		_, err := client.Network.ListPaloAltoFirewallNetworks(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListPaloAltoFirewallNetworks", testlistPaloAltoFirewallNetworks)

	testlistPhysicalNetworks := func(t *testing.T) {
		if _, ok := response["listPhysicalNetworks"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListPhysicalNetworksParams()
		_, err := client.Network.ListPhysicalNetworks(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListPhysicalNetworks", testlistPhysicalNetworks)

	testlistStorageNetworkIpRange := func(t *testing.T) {
		if _, ok := response["listStorageNetworkIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListStorageNetworkIpRangeParams()
		_, err := client.Network.ListStorageNetworkIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListStorageNetworkIpRange", testlistStorageNetworkIpRange)

	testlistSupportedNetworkServices := func(t *testing.T) {
		if _, ok := response["listSupportedNetworkServices"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListSupportedNetworkServicesParams()
		_, err := client.Network.ListSupportedNetworkServices(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSupportedNetworkServices", testlistSupportedNetworkServices)

	testreleasePublicIpRange := func(t *testing.T) {
		if _, ok := response["releasePublicIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewReleasePublicIpRangeParams("id")
		_, err := client.Network.ReleasePublicIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReleasePublicIpRange", testreleasePublicIpRange)

	testrestartNetwork := func(t *testing.T) {
		if _, ok := response["restartNetwork"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewRestartNetworkParams("id")
		_, err := client.Network.RestartNetwork(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RestartNetwork", testrestartNetwork)

	testupdateNetwork := func(t *testing.T) {
		if _, ok := response["updateNetwork"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewUpdateNetworkParams("id")
		r, err := client.Network.UpdateNetwork(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateNetwork", testupdateNetwork)

	testupdateNetworkServiceProvider := func(t *testing.T) {
		if _, ok := response["updateNetworkServiceProvider"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewUpdateNetworkServiceProviderParams("id")
		r, err := client.Network.UpdateNetworkServiceProvider(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateNetworkServiceProvider", testupdateNetworkServiceProvider)

	testupdatePhysicalNetwork := func(t *testing.T) {
		if _, ok := response["updatePhysicalNetwork"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewUpdatePhysicalNetworkParams("id")
		r, err := client.Network.UpdatePhysicalNetwork(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdatePhysicalNetwork", testupdatePhysicalNetwork)

	testupdateStorageNetworkIpRange := func(t *testing.T) {
		if _, ok := response["updateStorageNetworkIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewUpdateStorageNetworkIpRangeParams("id")
		r, err := client.Network.UpdateStorageNetworkIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateStorageNetworkIpRange", testupdateStorageNetworkIpRange)

	testdeleteGuestNetworkIpv6Prefix := func(t *testing.T) {
		if _, ok := response["deleteGuestNetworkIpv6Prefix"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewDeleteGuestNetworkIpv6PrefixParams("id")
		_, err := client.Network.DeleteGuestNetworkIpv6Prefix(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteGuestNetworkIpv6Prefix", testdeleteGuestNetworkIpv6Prefix)

	testcreateGuestNetworkIpv6Prefix := func(t *testing.T) {
		if _, ok := response["createGuestNetworkIpv6Prefix"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewCreateGuestNetworkIpv6PrefixParams("prefix", "zoneid")
		r, err := client.Network.CreateGuestNetworkIpv6Prefix(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateGuestNetworkIpv6Prefix", testcreateGuestNetworkIpv6Prefix)

	testlistGuestNetworkIpv6Prefixes := func(t *testing.T) {
		if _, ok := response["listGuestNetworkIpv6Prefixes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListGuestNetworkIpv6PrefixesParams()
		_, err := client.Network.ListGuestNetworkIpv6Prefixes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListGuestNetworkIpv6Prefixes", testlistGuestNetworkIpv6Prefixes)

	testcreateNetworkPermissions := func(t *testing.T) {
		if _, ok := response["createNetworkPermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewCreateNetworkPermissionsParams("networkid")
		_, err := client.Network.CreateNetworkPermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateNetworkPermissions", testcreateNetworkPermissions)

	testresetNetworkPermissions := func(t *testing.T) {
		if _, ok := response["resetNetworkPermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewResetNetworkPermissionsParams("networkid")
		_, err := client.Network.ResetNetworkPermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ResetNetworkPermissions", testresetNetworkPermissions)

	testlistNetworkPermissions := func(t *testing.T) {
		if _, ok := response["listNetworkPermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewListNetworkPermissionsParams("networkid")
		_, err := client.Network.ListNetworkPermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetworkPermissions", testlistNetworkPermissions)

	testremoveNetworkPermissions := func(t *testing.T) {
		if _, ok := response["removeNetworkPermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Network.NewRemoveNetworkPermissionsParams("networkid")
		_, err := client.Network.RemoveNetworkPermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveNetworkPermissions", testremoveNetworkPermissions)

}
