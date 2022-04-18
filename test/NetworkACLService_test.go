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

func TestNetworkACLService(t *testing.T) {
	service := "NetworkACLService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateNetworkACL := func(t *testing.T) {
		if _, ok := response["createNetworkACL"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewCreateNetworkACLParams("protocol")
		_, err := client.NetworkACL.CreateNetworkACL(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateNetworkACL", testcreateNetworkACL)

	testcreateNetworkACLList := func(t *testing.T) {
		if _, ok := response["createNetworkACLList"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewCreateNetworkACLListParams("name", "vpcid")
		_, err := client.NetworkACL.CreateNetworkACLList(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateNetworkACLList", testcreateNetworkACLList)

	testdeleteNetworkACL := func(t *testing.T) {
		if _, ok := response["deleteNetworkACL"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewDeleteNetworkACLParams("id")
		_, err := client.NetworkACL.DeleteNetworkACL(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNetworkACL", testdeleteNetworkACL)

	testdeleteNetworkACLList := func(t *testing.T) {
		if _, ok := response["deleteNetworkACLList"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewDeleteNetworkACLListParams("id")
		_, err := client.NetworkACL.DeleteNetworkACLList(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNetworkACLList", testdeleteNetworkACLList)

	testlistNetworkACLLists := func(t *testing.T) {
		if _, ok := response["listNetworkACLLists"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewListNetworkACLListsParams()
		_, err := client.NetworkACL.ListNetworkACLLists(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetworkACLLists", testlistNetworkACLLists)

	testlistNetworkACLs := func(t *testing.T) {
		if _, ok := response["listNetworkACLs"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewListNetworkACLsParams()
		_, err := client.NetworkACL.ListNetworkACLs(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetworkACLs", testlistNetworkACLs)

	testreplaceNetworkACLList := func(t *testing.T) {
		if _, ok := response["replaceNetworkACLList"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewReplaceNetworkACLListParams("aclid")
		_, err := client.NetworkACL.ReplaceNetworkACLList(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReplaceNetworkACLList", testreplaceNetworkACLList)

	testupdateNetworkACLItem := func(t *testing.T) {
		if _, ok := response["updateNetworkACLItem"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewUpdateNetworkACLItemParams("id")
		_, err := client.NetworkACL.UpdateNetworkACLItem(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateNetworkACLItem", testupdateNetworkACLItem)

	testupdateNetworkACLList := func(t *testing.T) {
		if _, ok := response["updateNetworkACLList"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NetworkACL.NewUpdateNetworkACLListParams("id")
		_, err := client.NetworkACL.UpdateNetworkACLList(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateNetworkACLList", testupdateNetworkACLList)

}
