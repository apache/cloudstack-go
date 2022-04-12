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

func TestNetworkService_CreateNetwork(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createNetwork"
		response, err := ReadData(apiName, "NetworkService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Network.NewCreateNetworkParams("testIsolatedNet", "testIsolatedNet", "69b7f746-208a-47c3-a940-4d3ebb720372", "04ccc336-d730-42fe-8ff6-5ae36e141e81")
	resp, err := client.Network.CreateNetwork(params)
	if err != nil {
		t.Errorf("Failed to create network due to: %v", err)
	}
	if resp == nil || resp.Name != "testIsolatedNet" {
		t.Errorf("Failed to create network")
	}
}

func TestNetworkService_ListNetworks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "listNetworks"
		response, err := ReadData(apiName, "NetworkService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Network.NewListNetworksParams()
	params.SetId("eb9c270d-dd66-443b-9524-ada1eff4442a")
	resp, err := client.Network.ListNetworks(params)
	if err != nil {
		t.Errorf("Failed to list specific network details due to : %v", err)
	}
	if resp.Count != 1 {
		t.Errorf("Failed to list specific network details")
	}
	if resp.Networks[0].Id != "eb9c270d-dd66-443b-9524-ada1eff4442a" {
		t.Errorf("Failed to list specific network details")
	}
}

func TestNetworkService_DeleteNetwork(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "deleteNetwork"
		response, err := ParseAsyncResponse(apiName, "NetworkService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Network.NewDeleteNetworkParams("eb9c270d-dd66-443b-9524-ada1eff4442a")
	resp, err := client.Network.DeleteNetwork(params)
	if err != nil {
		t.Errorf("Failed to delete network due to: %v", err)
		return
	}
	if resp == nil || !resp.Success {
		t.Errorf("Failed to delete network")
	}
}

//func TestNetworkService_RestartNetworkWithCleanUp(t *testing.T) {
//	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		responses := map[string]string{
//			"restartNetwork": `{
//				"restartnetworkresponse": {
//					"jobid": "09b1ab5d-02d0-410f-b373-6520962b98eb"
//				}
//			}`,
//			"queryAsyncJobResult": `{
//				"queryasyncjobresultresponse": {
//					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
//					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
//					"cmd": "org.apache.cloudstack.api.command.user.network.RestartNetworkCmd",
//					"jobstatus": 1,
//					"jobprocstatus": 0,
//					"jobresultcode": 0,
//					"jobresulttype": "object",
//					"jobresult": {
//						"success": true
//					},
//					"created": "2021-10-04T05:37:23+0000",
//					"completed": "2021-10-04T05:37:41+0000",
//					"jobid": "09b1ab5d-02d0-410f-b373-6520962b98eb"
//				}
//			}`,
//		}
//		fmt.Fprintf(writer, responses[request.FormValue("command")])
//	}))
//	defer server.Close()
//	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
//	params := client.Network.NewRestartNetworkParams("30358053-0f9d-4112-9948-976477896db6")
//	params.SetCleanup(true)
//	_, err := client.Network.RestartNetwork(params)
//	if err != nil {
//		t.Errorf("Failed to restart network with cleanup due to: %v", err)
//		return
//	}
//}

func TestNetworkService_CreatePhysicalNetwork(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createPhysicalNetwork"
		response, err := ParseAsyncResponse(apiName, "NetworkService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Network.NewCreatePhysicalNetworkParams("testPhyNet", "d4a81f75-5d92-415e-ab59-e85cc2ce56d9")
	resp, err := client.Network.CreatePhysicalNetwork(params)
	if err != nil {
		t.Errorf("Failed to create physical network due to: %v", err)
		return
	}
	if resp == nil || resp.Name != "testPhyNet" {
		t.Errorf("Failed to create physical network")
	}
}
