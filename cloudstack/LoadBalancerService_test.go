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
		responses := map[string]string{
			"createLoadBalancerRule": `{
				"createloadbalancerruleresponse": {
					"id": "d9186a3e-b39a-473f-8d5e-70cf1a7aed44",
					"jobid": "83ab0ded-225b-4a91-9aed-c17b1da07989"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse":{
					  "accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					  "cmd": "org.apache.cloudstack.api.command.user.loadbalancer.CreateLoadBalancerRuleCmd",
					  "completed": "2021-10-04T04:37:08+0000",
					  "created": "2021-10-04T04:37:08+0000",
					  "jobid": "83ab0ded-225b-4a91-9aed-c17b1da07989",
					  "jobinstanceid": "d9186a3e-b39a-473f-8d5e-70cf1a7aed44",
					  "jobinstancetype": "FirewallRule",
					  "jobprocstatus": 0,
					  "jobresult": {
						"loadbalancer": {
						  "account": "admin",
						  "algorithm": "roundrobin",
						  "cidrlist": "",
						  "domain": "ROOT",
						  "domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
						  "fordisplay": true,
						  "id": "d9186a3e-b39a-473f-8d5e-70cf1a7aed44",
						  "name": "testLBRule",
						  "networkid": "cc39d938-5ea0-4d9c-b89d-421da3274e54",
						  "privateport": "9090",
						  "protocol": "tcp",
						  "publicip": "192.168.2.107",
						  "publicipid": "bab02a09-1244-4235-a938-150e75e04ce0",
						  "publicport": "9090",
						  "state": "Add",
						  "tags": [],
						  "zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
						  "zonename": "SimZone1"
						}
					  },
					  "jobresultcode": 0,
					  "jobresulttype": "object",
					  "jobstatus": 1,
					  "userid": "27f2484f-5fe0-11ea-9a56-1e006800018c"
					}
				}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
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
		responses := map[string]string{
			"assignToLoadBalancerRule": `{
				"assigntoloadbalancerruleresponse": {
					"jobid": "7ab605d2-cddd-4f90-bd6b-884a8f62c16b"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.user.loadbalancer.AssignToLoadBalancerRuleCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"success": true
					},
					"created": "2021-10-03T07:52:26+0000",
					"completed": "2021-10-03T07:52:26+0000",
					"jobid": "7ab605d2-cddd-4f90-bd6b-884a8f62c16b"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
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
