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

func TestLoadBalancerService_CreateLoadBalancerRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createLoadBalancerRule"
		response, err := ParseAsyncResponse(apiName, "LoadBalancerService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.LoadBalancer.NewCreateLoadBalancerRuleParams("roundrobin", "testLBRule", 9090, 9090)
	params.SetProtocol("tcp")
	params.SetNetworkid("cc39d938-5ea0-4d9c-b89d-421da3274e54")
	params.SetPublicipid("bab02a09-1244-4235-a938-150e75e04ce0")
	resp, err := client.LoadBalancer.CreateLoadBalancerRule(params)
	if err != nil {
		t.Errorf("Failed to create load balancer rule due to: %v", err)
	}
	if resp.Name != "testLBRule" {
		t.Errorf("Failed to create load balancer rule")
	}
}

func TestLoadBalancerService_AssignLoadBalancerRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "assignLoadBalancerRule"
		response, err := ParseAsyncResponse(apiName, "LoadBalancerService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.LoadBalancer.NewAssignToLoadBalancerRuleParams("be6294cc-8f99-42c8-970d-f54c511972a3")
	vmIpMap := make(map[string]string)
	vmIpMap["vmid"] = "c7749abf-fadb-4658-9346-bbef95a2456f"
	vmIpMap["vmip"] = "10.1.1.178"
	params.SetVmidipmap(vmIpMap)
	resp, err := client.LoadBalancer.AssignToLoadBalancerRule(params)
	if err != nil {
		t.Errorf("Failed to assign vm to loadbalancer rule due to: %v", err)
		return
	}

	if resp == nil {
		t.Errorf("Failed to assign vm to loadbalancer rule")
	}
}
