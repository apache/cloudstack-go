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

func TestNsxService(t *testing.T) {
	service := "NsxService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddNsxController := func(t *testing.T) {
		if _, ok := response["addNsxController"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Nsx.NewAddNsxControllerParams("edgecluster", "name", "nsxproviderhostname", "password", "tier0gateway", "transportzone", "username", "zoneid")
		_, err := client.Nsx.AddNsxController(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddNsxController", testaddNsxController)

	testdeleteNsxController := func(t *testing.T) {
		if _, ok := response["deleteNsxController"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Nsx.NewDeleteNsxControllerParams("nsxcontrollerid")
		_, err := client.Nsx.DeleteNsxController(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNsxController", testdeleteNsxController)

	testlistNsxControllers := func(t *testing.T) {
		if _, ok := response["listNsxControllers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Nsx.NewListNsxControllersParams()
		_, err := client.Nsx.ListNsxControllers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNsxControllers", testlistNsxControllers)

}
