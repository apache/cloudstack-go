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

func TestSSHService_CreateSSHKeyPair(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "createSSHKeyPair"
		response, err := ReadData(apiName, "SSHService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.SSH.NewCreateSSHKeyPairParams("testSSHKey")
	resp, err := client.SSH.CreateSSHKeyPair(params)
	if err != nil {
		t.Errorf("Failed to create SSH key pair due to %v", err)
		return
	}
	if resp == nil || resp.Name != "testSSHKey" {
		t.Errorf("Failed to create SSH key pair")
	}
}

func TestSSHService_ListSSHKeyPairs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "listSSHKeyPairs"
		response, err := ReadData(apiName, "SSHService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.SSH.NewListSSHKeyPairsParams()
	params.SetName("testSSHKey")
	resp, err := client.SSH.ListSSHKeyPairs(params)
	if err != nil {
		t.Errorf("Failed to list SSH key pair due to %v", err)
		return
	}
	if resp == nil || resp.Count != 1 {
		t.Errorf("Failed to list SSH key pair")
	}
}

func TestSSHService_ResetSSHKeyForVirtualMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "resetSSHKeyForVirtualMachine"
		response, err := ParseAsyncResponse(apiName, "SSHService", *r)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.SSH.NewResetSSHKeyForVirtualMachineParams("8aa29529-b238-45f3-8992-5befadcd8bb0", "testSSHKey")
	resp, err := client.SSH.ResetSSHKeyForVirtualMachine(params)
	if err != nil {
		t.Errorf("Failed to reset SSH key pair for VM due to %v", err)
		return
	}
	if resp == nil || resp.Keypair != "testSSHKey" {
		t.Errorf("Failed to reset SSH key pair for VM")
	}
}

func TestSSHService_DeleteSSHKeyPair(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "deleteSSHKeyPair"
		response, err := ReadData(apiName, "SSHService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.SSH.NewDeleteSSHKeyPairParams("testSSHKey")
	resp, err := client.SSH.DeleteSSHKeyPair(params)
	if err != nil {
		t.Errorf("Failed to delete SSH key pair due to %v", err)
		return
	}
	if resp == nil || !resp.Success {
		t.Errorf("Failed to delete SSH key pair")
	}
}
