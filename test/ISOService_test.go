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

func TestISOService(t *testing.T) {
	service := "ISOService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testattachIso := func(t *testing.T) {
		if _, ok := response["attachIso"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewAttachIsoParams("id", "virtualmachineid")
		r, err := client.ISO.AttachIso(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AttachIso", testattachIso)

	testcopyIso := func(t *testing.T) {
		if _, ok := response["copyIso"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewCopyIsoParams("id")
		r, err := client.ISO.CopyIso(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CopyIso", testcopyIso)

	testdeleteIso := func(t *testing.T) {
		if _, ok := response["deleteIso"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewDeleteIsoParams("id")
		_, err := client.ISO.DeleteIso(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteIso", testdeleteIso)

	testdetachIso := func(t *testing.T) {
		if _, ok := response["detachIso"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewDetachIsoParams("virtualmachineid")
		r, err := client.ISO.DetachIso(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DetachIso", testdetachIso)

	testextractIso := func(t *testing.T) {
		if _, ok := response["extractIso"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewExtractIsoParams("id", "mode")
		r, err := client.ISO.ExtractIso(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ExtractIso", testextractIso)

	testlistIsoPermissions := func(t *testing.T) {
		if _, ok := response["listIsoPermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewListIsoPermissionsParams("id")
		_, err := client.ISO.ListIsoPermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListIsoPermissions", testlistIsoPermissions)

	testlistIsos := func(t *testing.T) {
		if _, ok := response["listIsos"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewListIsosParams()
		_, err := client.ISO.ListIsos(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListIsos", testlistIsos)

	testregisterIso := func(t *testing.T) {
		if _, ok := response["registerIso"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewRegisterIsoParams("name", "url", "zoneid")
		r, err := client.ISO.RegisterIso(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RegisterIso", testregisterIso)

	testupdateIso := func(t *testing.T) {
		if _, ok := response["updateIso"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewUpdateIsoParams("id")
		r, err := client.ISO.UpdateIso(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateIso", testupdateIso)

	testupdateIsoPermissions := func(t *testing.T) {
		if _, ok := response["updateIsoPermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ISO.NewUpdateIsoPermissionsParams("id")
		_, err := client.ISO.UpdateIsoPermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateIsoPermissions", testupdateIsoPermissions)

}
