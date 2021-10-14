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

func TestVirtualMachineService_DeployVirtualMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "deployVirtualMachine"
		response, err := ParseAsyncResponse(apiName, "VirtualMachineService", *r)
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.VirtualMachine.NewDeployVirtualMachineParams("498ce071-a077-4237-9492-e55c42553327",
		"50459b99-5fe0-11ea-9a56-1e006800018c", "1d8d87d4-1425-459c-8d81-c6f57dca2bd2")
	params.SetName("testDummyVM")
	params.SetNetworkids([]string{"87503ba2-d59d-4a5e-8f7f-b3f486cedbd8"})
	resp, err := client.VirtualMachine.DeployVirtualMachine(params)
	if err != nil {
		t.Errorf("Failed to deploy VM with given specs due to %v", err)
		return
	}

	if resp == nil || resp.Name != "testDummyVM" {
		t.Errorf("Failed to deploy VM with given specs")
	}
}

func TestVirtualMachineService_AddNicToVirtualMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "addNicToVirtualMachine"
		response, err := ParseAsyncResponse(apiName, "VirtualMachineService", *r)
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.VirtualMachine.NewAddNicToVirtualMachineParams("30358053-0f9d-4112-9948-976477896db6",
		"915653c4-298b-4d74-bdee-4ced282114f1")
	resp, err := client.VirtualMachine.AddNicToVirtualMachine(params)
	if err != nil {
		t.Errorf("Failed to add nic to VM due to %v", err)
		return
	}

	if resp == nil || len(resp.Nic) != 2 {
		t.Errorf("Failed to add nic to VM")
	}
}

func TestVirtualMachineService_StopVirtualMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "stopVirtualMachine"
		response, err := ParseAsyncResponse(apiName, "VirtualMachineService", *r)
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.VirtualMachine.NewStopVirtualMachineParams("915653c4-298b-4d74-bdee-4ced282114f1")
	resp, err := client.VirtualMachine.StopVirtualMachine(params)
	if err != nil {
		t.Errorf("Failed to stop VM due to %v", err)
		return
	}

	if resp == nil || resp.State != "Stopped" {
		t.Errorf("Failed to stop VM")
	}
}

func TestVirtualMachineService_StartVirtualMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "startVirtualMachine"
		response, err := ParseAsyncResponse(apiName, "VirtualMachineService", *r)
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.VirtualMachine.NewStartVirtualMachineParams("915653c4-298b-4d74-bdee-4ced282114f1")
	resp, err := client.VirtualMachine.StartVirtualMachine(params)
	if err != nil {
		t.Errorf("Failed to start VM due to %v", err)
		return
	}

	if resp == nil || resp.State != "Running" {
		t.Errorf("Failed to start VM")
	}
}

func TestVirtualMachineService_ListVirtualMachines(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "listVirtualMachines"
		response, err := ReadData(apiName, "VirtualMachineService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.VirtualMachine.NewListVirtualMachinesParams()
	params.SetId("915653c4-298b-4d74-bdee-4ced282114f1")
	resp, err := client.VirtualMachine.ListVirtualMachines(params)
	if err != nil {
		t.Errorf("Failed to list VM due to %v", err)
		return
	}

	if resp == nil || resp.Count != 1 {
		t.Errorf("Failed to list VM")
	}
}
