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

func TestNetworkOfferingService(t *testing.T) {
	service := "NetworkOfferingService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateNetworkOffering := func(t *testing.T) {
		if _, ok := response["createNetworkOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkOffering.NewCreateNetworkOfferingParams("displaytext", "guestiptype", "name", "traffictype")
		_, err := client.NetworkOffering.CreateNetworkOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateNetworkOffering", testcreateNetworkOffering)

	testdeleteNetworkOffering := func(t *testing.T) {
		if _, ok := response["deleteNetworkOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkOffering.NewDeleteNetworkOfferingParams("id")
		_, err := client.NetworkOffering.DeleteNetworkOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNetworkOffering", testdeleteNetworkOffering)

	testlistNetworkOfferings := func(t *testing.T) {
		if _, ok := response["listNetworkOfferings"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkOffering.NewListNetworkOfferingsParams()
		_, err := client.NetworkOffering.ListNetworkOfferings(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetworkOfferings", testlistNetworkOfferings)

	testupdateNetworkOffering := func(t *testing.T) {
		if _, ok := response["updateNetworkOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkOffering.NewUpdateNetworkOfferingParams()
		_, err := client.NetworkOffering.UpdateNetworkOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateNetworkOffering", testupdateNetworkOffering)

}
