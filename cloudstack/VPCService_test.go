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

func TestVPCService_RestartVPC(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "restartVPC"
		response, err := ParseAsyncResponse(apiName, "VPCService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.VPC.NewRestartVPCParams("f9ec95f3-70be-448a-8ba2-cb6388dce55a")

	fmt.Printf(">>>>> 1 params: %v", params)
	resp, err := client.VPC.RestartVPC(params)
	if err != nil {
		t.Errorf("Failed to restart VPC network due to: %v", err)
	}
	if resp == nil {
		t.Errorf("Failed to restart VPC network")
	}
}

func TestVPCService_ListVPCs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "listVPCs"
		response, err := ParseAsyncResponse(apiName, "VPCService", *request)
		fmt.Printf(">>>>> VPC Resp: %v", response)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.VPC.NewListVPCsParams()
	// params.SetId("78d49100-e4af-4692-9a49-284fa2444eb9")
	fmt.Printf(">>>>> 1 params: %v", params)
	resp, err := client.VPC.ListVPCs(params)

	if err != nil {
		t.Errorf("Failed to list VPCs due to: %v", err)
	}
	if resp == nil {
		t.Errorf("Failed to list VPCs")
	}
}
