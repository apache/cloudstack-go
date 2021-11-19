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

func TestHostService_AddHost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "addHost"
		response, err := ReadData(apiName, "HostService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))

	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Host.NewAddHostParams("Simulator", "5382edc2-e689-4074-bd67-0e1a236eb2bc", "http://sim/c0/h0",
		"d4a81f75-5d92-415e-ab59-e85cc2ce56d9")
	resp, err := client.Host.AddHost(params)
	if err != nil {
		t.Errorf("Failed to add host due to: %v", err)
		return
	}

	if resp == nil {
		t.Errorf("Failed to add host")
	}
}

func TestHostService_ListHosts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "listHosts"
		response, err := ReadData(apiName, "HostService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Host.NewListHostsParams()
	resp, err := client.Host.ListHosts(params)
	if err != nil {
		t.Errorf("Failed to add host due to: %v", err)
		return
	}
	if resp.Count != 1 {
		t.Errorf("Failed to add host")
	}
}

func TestHostService_PrepareHostForMaintenance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "prepareHostForMaintenance"
		response, err := ParseAsyncResponse(apiName, "HostService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Host.NewPrepareHostForMaintenanceParams("8e8e68e7-19ea-4a78-bbdb-6e79d27729c9")
	resp, err := client.Host.PrepareHostForMaintenance(params)
	if err != nil {
		t.Errorf("Failed to prepare host for maintenance due to: %v", err)
	}
	if resp.Resourcestate != "PrepareForMaintenance" {
		t.Errorf("Failed to prepare host for maintenance")
	}
}

func TestHostService_CancelHostForMaintenance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "cancelHostForMaintenance"
		response, err := ParseAsyncResponse(apiName, "HostService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Host.NewCancelHostMaintenanceParams("8e8e68e7-19ea-4a78-bbdb-6e79d27729c9")
	resp, err := client.Host.CancelHostMaintenance(params)
	if err != nil {
		t.Errorf("Failed to cancel host for maintenance due to: %v", err)
	}
	if resp.Resourcestate != "Enabled" {
		t.Errorf("Failed to cancel host for maintenance")
	}
}

func TestHostService_EnableOutOfBandManagementForHost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "enableOutOfBandManagementForHost"
		response, err := ParseAsyncResponse(apiName, "HostService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Host.NewEnableOutOfBandManagementForHostParams("8e8e68e7-19ea-4a78-bbdb-6e79d27729c9")
	resp, err := client.Host.EnableOutOfBandManagementForHost(params)
	if err != nil {
		t.Errorf("Failed to enable out of band management for due to: %v", err)
	}
	if !resp.Enabled {
		t.Errorf("Failed to enable out of band management")
	}
}

func TestHostService_DisableOutOfBandManagementForHost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "disableOutOfBandManagementForHost"
		response, err := ParseAsyncResponse(apiName, "HostService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Host.NewDisableOutOfBandManagementForHostParams("8e8e68e7-19ea-4a78-bbdb-6e79d27729c9")
	resp, err := client.Host.DisableOutOfBandManagementForHost(params)
	if err != nil {
		t.Errorf("Failed to disable out of band management for due to: %v", err)
	}
	if resp.Enabled {
		t.Errorf("Failed to disable out of band management")
	}
}

func TestHostService_EnableHAForHost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "enableHAForHost"
		response, err := ParseAsyncResponse(apiName, "HostService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Host.NewEnableHAForHostParams("8e8e68e7-19ea-4a78-bbdb-6e79d27729c9")
	resp, err := client.Host.EnableHAForHost(params)
	if err != nil {
		t.Errorf("Failed to enable HA for due to: %v", err)
	}
	if !resp.Haenable {
		t.Errorf("Failed to enable HA for host")
	}
}
