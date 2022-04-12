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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestStoragePoolService_CancelStorageMaintenance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "cancelStorageMaintenance"
		response, err := ParseAsyncResponse(apiName, "StoragePoolService", *r)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.StoragePool.NewCancelStorageMaintenanceParams("44ad900d-785b-3eff-addc-a5e6bf4927ef")
	resp, err := client.StoragePool.CancelStorageMaintenance(params)
	if err != nil {
		t.Errorf("Failed to cancel storage maintenance due to %v", err)
		return
	}
	if resp == nil || resp.State != "Up" {
		t.Errorf("Failed to cancel Storage pool maintenance")
	}
}
func TestStoragePoolService_EnableStorageMaintenance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "enableStorageMaintenance"
		response, err := ParseAsyncResponse(apiName, "StoragePoolService", *r)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.StoragePool.NewEnableStorageMaintenanceParams("44ad900d-785b-3eff-addc-a5e6bf4927ef")
	resp, err := client.StoragePool.EnableStorageMaintenance(params)
	if err != nil {
		t.Errorf("Failed to enable storage maintenance due to %v", err)
		return
	}
	if resp == nil || resp.State != "Maintenance" {
		t.Errorf("Failed to enable Storage pool maintenance")
	}
}
