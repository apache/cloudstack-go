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

func TestNetworkService_CreateNetwork(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"createnetworkresponse": {
				"network": {
					"id": "eb9c270d-dd66-443b-9524-ada1eff4442a",
					"name": "testIsolatedNet",
					"displaytext": "testIsolatedNet",
					"broadcastdomaintype": "Vlan",
					"traffictype": "Guest",
					"gateway": "10.1.1.1",
					"netmask": "255.255.255.0",
					"cidr": "10.1.1.0/24",
					"zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
					"zonename": "SimZone1",
					"networkofferingid": "69b7f746-208a-47c3-a940-4d3ebb720372",
					"networkofferingname": "DefaultIsolatedNetworkOfferingWithSourceNatService",
					"networkofferingdisplaytext": "Offering for Isolated networks with Source Nat service enabled",
					"networkofferingconservemode": true,
					"networkofferingavailability": "Required",
					"issystem": false,
					"state": "Allocated",
					"related": "eb9c270d-dd66-443b-9524-ada1eff4442a",
					"dns1": "10.147.28.6",
					"type": "Isolated",
					"acltype": "Account",
					"account": "admin",
					"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
					"domain": "ROOT",
					"service": [
						{
							"name": "Lb",
							"capability": [
								{
									"name": "SupportedLBIsolation",
									"value": "dedicated",
									"canchooseservicecapability": false
								},
								{
									"name": "SupportedLbAlgorithms",
									"value": "roundrobin,leastconn,source",
									"canchooseservicecapability": false
								},
								{
									"name": "SupportedProtocols",
									"value": "tcp, udp, tcp-proxy",
									"canchooseservicecapability": false
								},
								{
									"name": "LbSchemes",
									"value": "Public",
									"canchooseservicecapability": false
								},
								{
									"name": "SupportedStickinessMethods",
									"value": "[{\"methodname\":\"LbCookie\",\"paramlist\":[{\"paramname\":\"cookie-name\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"mode\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"nocache\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"indirect\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"postonly\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"domain\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is loadbalancer cookie based stickiness method.\"},{\"methodname\":\"AppCookie\",\"paramlist\":[{\"paramname\":\"cookie-name\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"length\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"holdtime\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"request-learn\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"prefix\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"mode\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is App session based sticky method. Define session stickiness on an existing application cookie. It can be used only for a specific http traffic\"},{\"methodname\":\"SourceBased\",\"paramlist\":[{\"paramname\":\"tablesize\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"expire\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is source based Stickiness method, it can be used for any type of protocol.\"}]",
									"canchooseservicecapability": false
								},
								{
									"name": "AutoScaleCounters",
									"value": "[{\"methodname\":\"cpu\",\"paramlist\":[]},{\"methodname\":\"memory\",\"paramlist\":[]}]",
									"canchooseservicecapability": false
								}
							]
						},
						{
							"name": "PortForwarding",
							"capability": [
								{
									"name": "SupportedProtocols",
									"value": "tcp,udp",
									"canchooseservicecapability": false
								}
							]
						},
						{
							"name": "SourceNat",
							"capability": [
								{
									"name": "SupportedSourceNatTypes",
									"value": "peraccount",
									"canchooseservicecapability": false
								},
								{
									"name": "RedundantRouter",
									"value": "true",
									"canchooseservicecapability": false
								}
							]
						},
						{
							"name": "Vpn",
							"capability": [
								{
									"name": "VpnTypes",
									"value": "removeaccessvpn",
									"canchooseservicecapability": false
								},
								{
									"name": "SupportedVpnTypes",
									"value": "pptp,l2tp,ipsec",
									"canchooseservicecapability": false
								}
							]
						},
						{
							"name": "Dns",
							"capability": [
								{
									"name": "AllowDnsSuffixModification",
									"value": "true",
									"canchooseservicecapability": false
								}
							]
						},
						{
							"name": "Dhcp",
							"capability": [
								{
									"name": "DhcpAccrossMultipleSubnets",
									"value": "true",
									"canchooseservicecapability": false
								}
							]
						},
						{
							"name": "StaticNat"
						},
						{
							"name": "UserData"
						},
						{
							"name": "Firewall",
							"capability": [
								{
									"name": "TrafficStatistics",
									"value": "per public ip",
									"canchooseservicecapability": false
								},
								{
									"name": "SupportedProtocols",
									"value": "tcp,udp,icmp",
									"canchooseservicecapability": false
								},
								{
									"name": "MultipleIps",
									"value": "true",
									"canchooseservicecapability": false
								},
								{
									"name": "SupportedTrafficDirection",
									"value": "ingress, egress",
									"canchooseservicecapability": false
								},
								{
									"name": "SupportedEgressProtocols",
									"value": "tcp,udp,icmp, all",
									"canchooseservicecapability": false
								}
							]
						}
					],
					"networkdomain": "cs2sandbox.simulator",
					"physicalnetworkid": "8e27a637-7525-49ed-81ce-52bd5e5d9ea2",
					"restartrequired": false,
					"specifyipranges": false,
					"canusefordeploy": true,
					"ispersistent": false,
					"tags": [],
					"details": {},
					"displaynetwork": true,
					"strechedl2subnet": false,
					"redundantrouter": false
				}
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
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
		response := `{
			"listnetworksresponse": {
				"count": 1,
				"network": [
					{
						"id": "eb9c270d-dd66-443b-9524-ada1eff4442a",
						"name": "testIsolatedNet",
						"displaytext": "testIsolatedNet",
						"broadcastdomaintype": "Vlan",
						"traffictype": "Guest",
						"gateway": "10.1.1.1",
						"netmask": "255.255.255.0",
						"cidr": "10.1.1.0/24",
						"zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
						"zonename": "SimZone1",
						"networkofferingid": "69b7f746-208a-47c3-a940-4d3ebb720372",
						"networkofferingname": "DefaultIsolatedNetworkOfferingWithSourceNatService",
						"networkofferingdisplaytext": "Offering for Isolated networks with Source Nat service enabled",
						"networkofferingconservemode": true,
						"networkofferingavailability": "Required",
						"issystem": false,
						"state": "Allocated",
						"related": "eb9c270d-dd66-443b-9524-ada1eff4442a",
						"dns1": "10.147.28.6",
						"type": "Isolated",
						"acltype": "Account",
						"account": "admin",
						"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
						"domain": "ROOT",
						"service": [
							{
								"name": "Lb",
								"capability": [
									{
										"name": "SupportedLBIsolation",
										"value": "dedicated",
										"canchooseservicecapability": false
									},
									{
										"name": "SupportedLbAlgorithms",
										"value": "roundrobin,leastconn,source",
										"canchooseservicecapability": false
									},
									{
										"name": "SupportedProtocols",
										"value": "tcp, udp, tcp-proxy",
										"canchooseservicecapability": false
									},
									{
										"name": "LbSchemes",
										"value": "Public",
										"canchooseservicecapability": false
									},
									{
										"name": "SupportedStickinessMethods",
										"value": "[{\"methodname\":\"LbCookie\",\"paramlist\":[{\"paramname\":\"cookie-name\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"mode\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"nocache\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"indirect\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"postonly\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"domain\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is loadbalancer cookie based stickiness method.\"},{\"methodname\":\"AppCookie\",\"paramlist\":[{\"paramname\":\"cookie-name\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"length\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"holdtime\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"request-learn\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"prefix\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"mode\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is App session based sticky method. Define session stickiness on an existing application cookie. It can be used only for a specific http traffic\"},{\"methodname\":\"SourceBased\",\"paramlist\":[{\"paramname\":\"tablesize\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"expire\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is source based Stickiness method, it can be used for any type of protocol.\"}]",
										"canchooseservicecapability": false
									},
									{
										"name": "AutoScaleCounters",
										"value": "[{\"methodname\":\"cpu\",\"paramlist\":[]},{\"methodname\":\"memory\",\"paramlist\":[]}]",
										"canchooseservicecapability": false
									}
								]
							},
							{
								"name": "PortForwarding",
								"capability": [
									{
										"name": "SupportedProtocols",
										"value": "tcp,udp",
										"canchooseservicecapability": false
									}
								]
							},
							{
								"name": "SourceNat",
								"capability": [
									{
										"name": "SupportedSourceNatTypes",
										"value": "peraccount",
										"canchooseservicecapability": false
									},
									{
										"name": "RedundantRouter",
										"value": "true",
										"canchooseservicecapability": false
									}
								]
							},
							{
								"name": "Vpn",
								"capability": [
									{
										"name": "VpnTypes",
										"value": "removeaccessvpn",
										"canchooseservicecapability": false
									},
									{
										"name": "SupportedVpnTypes",
										"value": "pptp,l2tp,ipsec",
										"canchooseservicecapability": false
									}
								]
							},
							{
								"name": "Dns",
								"capability": [
									{
										"name": "AllowDnsSuffixModification",
										"value": "true",
										"canchooseservicecapability": false
									}
								]
							},
							{
								"name": "Dhcp",
								"capability": [
									{
										"name": "DhcpAccrossMultipleSubnets",
										"value": "true",
										"canchooseservicecapability": false
									}
								]
							},
							{
								"name": "StaticNat"
							},
							{
								"name": "UserData"
							},
							{
								"name": "Firewall",
								"capability": [
									{
										"name": "TrafficStatistics",
										"value": "per public ip",
										"canchooseservicecapability": false
									},
									{
										"name": "SupportedProtocols",
										"value": "tcp,udp,icmp",
										"canchooseservicecapability": false
									},
									{
										"name": "MultipleIps",
										"value": "true",
										"canchooseservicecapability": false
									},
									{
										"name": "SupportedTrafficDirection",
										"value": "ingress, egress",
										"canchooseservicecapability": false
									},
									{
										"name": "SupportedEgressProtocols",
										"value": "tcp,udp,icmp, all",
										"canchooseservicecapability": false
									}
								]
							}
						],
						"networkdomain": "cs2sandbox.simulator",
						"physicalnetworkid": "8e27a637-7525-49ed-81ce-52bd5e5d9ea2",
						"restartrequired": false,
						"specifyipranges": false,
						"canusefordeploy": true,
						"ispersistent": false,
						"tags": [],
						"details": {},
						"displaynetwork": true,
						"strechedl2subnet": false,
						"redundantrouter": false
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
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
		responses := map[string]string{
			"deleteNetwork": `{
				"deletenetworkresponse": {
					"jobid": "8fe59207-90f5-4382-af17-e8b413bda517"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.user.network.DeleteNetworkCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"success": true
					},
					"jobinstancetype": "Network",
					"created": "2021-10-04T05:11:01+0000",
					"completed": "2021-10-04T05:11:02+0000",
					"jobid": "8fe59207-90f5-4382-af17-e8b413bda517"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
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
//	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
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
		responses := map[string]string{
			"createPhysicalNetwork": `{
				"createphysicalnetworkresponse": {
					"id": "256b3b3f-4a31-48f3-a939-79b01357fb87",
					"jobid": "a52138ee-752e-400a-8289-c534eebca026"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "bc1b501c-1d18-11ec-9173-50eb7122da94",
					"userid": "bc1b6b47-1d18-11ec-9173-50eb7122da94",
					"cmd": "org.apache.cloudstack.api.command.admin.network.CreatePhysicalNetworkCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"physicalnetwork": {
							"id": "256b3b3f-4a31-48f3-a939-79b01357fb87",
							"name": "testPhyNet",
							"broadcastdomainrange": "ZONE",
							"zoneid": "d4a81f75-5d92-415e-ab59-e85cc2ce56d9",
							"zonename": "Sandbox-simulator",
							"state": "Disabled",
							"isolationmethods": "VLAN"
						}
					},
					"jobinstancetype": "PhysicalNetwork",
					"jobinstanceid": "256b3b3f-4a31-48f3-a939-79b01357fb87",
					"created": "2021-10-04T11:26:38+0530",
					"completed": "2021-10-04T11:26:38+0530",
					"jobid": "a52138ee-752e-400a-8289-c534eebca026"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
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
