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

func TestNetscalerService(t *testing.T) {
	service := "NetscalerService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddNetscalerLoadBalancer := func(t *testing.T) {
		if _, ok := response["addNetscalerLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewAddNetscalerLoadBalancerParams("networkdevicetype", "password", "physicalnetworkid", "url", "username")
		_, err := client.Netscaler.AddNetscalerLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddNetscalerLoadBalancer", testaddNetscalerLoadBalancer)

	testconfigureNetscalerLoadBalancer := func(t *testing.T) {
		if _, ok := response["configureNetscalerLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewConfigureNetscalerLoadBalancerParams("lbdeviceid")
		_, err := client.Netscaler.ConfigureNetscalerLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ConfigureNetscalerLoadBalancer", testconfigureNetscalerLoadBalancer)

	testdeleteNetscalerControlCenter := func(t *testing.T) {
		if _, ok := response["deleteNetscalerControlCenter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewDeleteNetscalerControlCenterParams("id")
		_, err := client.Netscaler.DeleteNetscalerControlCenter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNetscalerControlCenter", testdeleteNetscalerControlCenter)

	testdeleteNetscalerLoadBalancer := func(t *testing.T) {
		if _, ok := response["deleteNetscalerLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewDeleteNetscalerLoadBalancerParams("lbdeviceid")
		_, err := client.Netscaler.DeleteNetscalerLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNetscalerLoadBalancer", testdeleteNetscalerLoadBalancer)

	testlistNetscalerControlCenter := func(t *testing.T) {
		if _, ok := response["listNetscalerControlCenter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewListNetscalerControlCenterParams()
		_, err := client.Netscaler.ListNetscalerControlCenter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetscalerControlCenter", testlistNetscalerControlCenter)

	testlistNetscalerLoadBalancerNetworks := func(t *testing.T) {
		if _, ok := response["listNetscalerLoadBalancerNetworks"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewListNetscalerLoadBalancerNetworksParams("lbdeviceid")
		_, err := client.Netscaler.ListNetscalerLoadBalancerNetworks(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetscalerLoadBalancerNetworks", testlistNetscalerLoadBalancerNetworks)

	testlistNetscalerLoadBalancers := func(t *testing.T) {
		if _, ok := response["listNetscalerLoadBalancers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewListNetscalerLoadBalancersParams()
		_, err := client.Netscaler.ListNetscalerLoadBalancers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetscalerLoadBalancers", testlistNetscalerLoadBalancers)

	testregisterNetscalerControlCenter := func(t *testing.T) {
		if _, ok := response["registerNetscalerControlCenter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewRegisterNetscalerControlCenterParams("ipaddress", 0, "password", "username")
		_, err := client.Netscaler.RegisterNetscalerControlCenter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RegisterNetscalerControlCenter", testregisterNetscalerControlCenter)

	testregisterNetscalerServicePackage := func(t *testing.T) {
		if _, ok := response["registerNetscalerServicePackage"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Netscaler.NewRegisterNetscalerServicePackageParams("description", "name")
		r, err := client.Netscaler.RegisterNetscalerServicePackage(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RegisterNetscalerServicePackage", testregisterNetscalerServicePackage)

}
