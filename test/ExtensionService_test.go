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

func TestExtensionService(t *testing.T) {
	service := "ExtensionService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddCustomAction := func(t *testing.T) {
		if _, ok := response["addCustomAction"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewAddCustomActionParams("extensionid", "name")
		r, err := client.Extension.AddCustomAction(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddCustomAction", testaddCustomAction)

	testcreateExtension := func(t *testing.T) {
		if _, ok := response["createExtension"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewCreateExtensionParams("name", "type")
		r, err := client.Extension.CreateExtension(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateExtension", testcreateExtension)

	testdeleteCustomAction := func(t *testing.T) {
		if _, ok := response["deleteCustomAction"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewDeleteCustomActionParams()
		_, err := client.Extension.DeleteCustomAction(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteCustomAction", testdeleteCustomAction)

	testdeleteExtension := func(t *testing.T) {
		if _, ok := response["deleteExtension"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewDeleteExtensionParams()
		r, err := client.Extension.DeleteExtension(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DeleteExtension", testdeleteExtension)

	testlistCustomActions := func(t *testing.T) {
		if _, ok := response["listCustomActions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewListCustomActionsParams()
		_, err := client.Extension.ListCustomActions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCustomActions", testlistCustomActions)

	testlistExtensions := func(t *testing.T) {
		if _, ok := response["listExtensions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewListExtensionsParams()
		_, err := client.Extension.ListExtensions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListExtensions", testlistExtensions)

	testregisterExtension := func(t *testing.T) {
		if _, ok := response["registerExtension"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewRegisterExtensionParams("extensionid", "resourceid", "resourcetype")
		r, err := client.Extension.RegisterExtension(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RegisterExtension", testregisterExtension)

	testrunCustomAction := func(t *testing.T) {
		if _, ok := response["runCustomAction"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewRunCustomActionParams("customactionid", "resourceid")
		r, err := client.Extension.RunCustomAction(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RunCustomAction", testrunCustomAction)

	testunregisterExtension := func(t *testing.T) {
		if _, ok := response["unregisterExtension"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewUnregisterExtensionParams("extensionid", "resourceid", "resourcetype")
		r, err := client.Extension.UnregisterExtension(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UnregisterExtension", testunregisterExtension)

	testupdateCustomAction := func(t *testing.T) {
		if _, ok := response["updateCustomAction"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewUpdateCustomActionParams("id")
		_, err := client.Extension.UpdateCustomAction(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateCustomAction", testupdateCustomAction)

	testupdateExtension := func(t *testing.T) {
		if _, ok := response["updateExtension"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Extension.NewUpdateExtensionParams("id")
		r, err := client.Extension.UpdateExtension(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateExtension", testupdateExtension)

}
