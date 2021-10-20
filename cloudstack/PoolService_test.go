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

package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPoolService_CreateStoragePool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "createStoragePool"
		response, err := ReadData(apiName, "PoolService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Pool.NewCreateStoragePoolParams("testPrimary1", "nfs://10.1.1.2/export/primary1", "6ebc7a1c-ff98-425c-ac0d-b6c2e3ae6e33")
	resp, err := client.Pool.CreateStoragePool(params)
	if err != nil {
		t.Errorf("Failed to create storage pool due to %v", err)
		return
	}

	if resp == nil || resp.Name != "testPrimary1" {
		t.Errorf("Failed to create storage pool")
	}
}

func TestPoolService_ListStoragePools(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "listStoragePools"
		response, err := ReadData(apiName, "PoolService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Pool.NewListStoragePoolsParams()
	params.SetId("3fd44942-4881-3fb0-95f1-5d7b7ae64cfb")
	resp, err := client.Pool.ListStoragePools(params)
	if err != nil {
		t.Errorf("Failed to list storage pool details due to %v", err)
		return
	}

	if resp == nil || resp.Count != 1 {
		t.Errorf("Failed to list storage pool")
	}
}

func TestPoolService_DeleteStoragePool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "deleteStoragePool"
		response, err := ReadData(apiName, "PoolService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Pool.NewDeleteStoragePoolParams("3fd44942-4881-3fb0-95f1-5d7b7ae64cfb")
	resp, err := client.Pool.DeleteStoragePool(params)
	if err != nil {
		t.Errorf("Failed to delete storage pool details due to %v", err)
		return
	}

	if resp == nil || !resp.Success {
		t.Errorf("Failed to delete storage pool")
	}
}
