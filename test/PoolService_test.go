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

func TestPoolService(t *testing.T) {
	service := "PoolService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateStoragePool := func(t *testing.T) {
		if _, ok := response["createStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pool.NewCreateStoragePoolParams("name", "url", "zoneid")
		r, err := client.Pool.CreateStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateStoragePool", testcreateStoragePool)

	testdeleteStoragePool := func(t *testing.T) {
		if _, ok := response["deleteStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pool.NewDeleteStoragePoolParams("id")
		_, err := client.Pool.DeleteStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteStoragePool", testdeleteStoragePool)

	testfindStoragePoolsForMigration := func(t *testing.T) {
		if _, ok := response["findStoragePoolsForMigration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pool.NewFindStoragePoolsForMigrationParams("id")
		r, err := client.Pool.FindStoragePoolsForMigration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("FindStoragePoolsForMigration", testfindStoragePoolsForMigration)

	testlistStoragePools := func(t *testing.T) {
		if _, ok := response["listStoragePools"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pool.NewListStoragePoolsParams()
		_, err := client.Pool.ListStoragePools(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListStoragePools", testlistStoragePools)

	testsyncStoragePool := func(t *testing.T) {
		if _, ok := response["syncStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pool.NewSyncStoragePoolParams("id")
		r, err := client.Pool.SyncStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("SyncStoragePool", testsyncStoragePool)

	testupdateStoragePool := func(t *testing.T) {
		if _, ok := response["updateStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pool.NewUpdateStoragePoolParams("id")
		r, err := client.Pool.UpdateStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateStoragePool", testupdateStoragePool)

}
