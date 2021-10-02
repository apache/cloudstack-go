package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHostService_AddHost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"addhostresponse": {
			  "count": 1,
			  "host": [
				{
				  "capabilities": "hvm",
				  "clusterid": "2e4816cf-d8bf-4ef0-993a-05e42078c5a8",
				  "clustername": "C0",
				  "clustertype": "CloudManaged",
				  "cpuallocated": "0%",
				  "cpuallocatedpercentage": "0%",
				  "cpuallocatedvalue": 0,
				  "cpuallocatedwithoverprovisioning": "0%",
				  "cpunumber": 4,
				  "cpuspeed": 8000,
				  "cpuwithoverprovisioning": "32000",
				  "created": "2021-10-03T10:59:51+0530",
				  "events": "StartAgentRebalance; AgentConnected; ManagementServerDown; PingTimeout; ShutdownRequested; HostDown; AgentDisconnected; Ping; Remove",
				  "hahost": false,
				  "hasannotations": false,
				  "hostha": {
					"haenable": false,
					"hastate": "Disabled"
				  },
				  "hypervisor": "Simulator",
				  "hypervisorversion": "4.16.0.0-SNAPSHOT",
				  "id": "d9960e4d-5da9-47be-af88-5cd19473b844",
				  "ipaddress": "172.16.15.5",
				  "islocalstorageactive": false,
				  "lastpinged": "1970-01-19T16:32:39+0530",
				  "managementserverid": "028003eb-4d1a-4d17-a2e3-0ec901c9ee5c",
				  "memoryallocated": 0,
				  "memoryallocatedbytes": 0,
				  "memoryallocatedpercentage": "0%",
				  "memorytotal": 8589934592,
				  "memorywithoverprovisioning": "8589934592",
				  "name": "SimulatedAgent.5417728f-9fc2-424c-89a6-00a7e0d2852f",
				  "outofbandmanagement": {
					"enabled": false,
					"powerstate": "Disabled"
				  },
				  "podid": "5382edc2-e689-4074-bd67-0e1a236eb2bc",
				  "podname": "POD0",
				  "resourcestate": "Enabled",
				  "state": "Up",
				  "type": "Routing",
				  "version": "4.16.0.0-SNAPSHOT",
				  "zoneid": "d4a81f75-5d92-415e-ab59-e85cc2ce56d9",
				  "zonename": "Sandbox-simulator"
				}
			  ]
			}
		}`

		fmt.Fprintln(writer, response)
	}))

	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Host.NewAddHostParams("Simulator", "password",
		"5382edc2-e689-4074-bd67-0e1a236eb2bc", "http://sim/c0/h0",
		"root", "d4a81f75-5d92-415e-ab59-e85cc2ce56d9")
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
		response := `{
			"listhostsresponse": {
			  "count": 1,
			  "host": [
				{
				  "capabilities": "hvm",
				  "clusterid": "2e4816cf-d8bf-4ef0-993a-05e42078c5a8",
				  "clustername": "C0",
				  "clustertype": "CloudManaged",
				  "cpuallocated": "0%",
				  "cpuallocatedpercentage": "0%",
				  "cpuallocatedvalue": 0,
				  "cpuallocatedwithoverprovisioning": "0%",
				  "cpunumber": 4,
				  "cpuspeed": 8000,
				  "cpuwithoverprovisioning": "32000",
				  "created": "2021-10-03T10:59:51+0530",
				  "events": "StartAgentRebalance; AgentConnected; ManagementServerDown; PingTimeout; ShutdownRequested; HostDown; AgentDisconnected; Ping; Remove",
				  "hahost": false,
				  "hasannotations": false,
				  "hostha": {
					"haenable": false,
					"hastate": "Disabled"
				  },
				  "hypervisor": "Simulator",
				  "hypervisorversion": "4.16.0.0-SNAPSHOT",
				  "id": "d9960e4d-5da9-47be-af88-5cd19473b844",
				  "ipaddress": "172.16.15.5",
				  "islocalstorageactive": false,
				  "lastpinged": "1970-01-19T16:32:39+0530",
				  "managementserverid": "028003eb-4d1a-4d17-a2e3-0ec901c9ee5c",
				  "memoryallocated": 0,
				  "memoryallocatedbytes": 0,
				  "memoryallocatedpercentage": "0%",
				  "memorytotal": 8589934592,
				  "memorywithoverprovisioning": "8589934592",
				  "name": "SimulatedAgent.5417728f-9fc2-424c-89a6-00a7e0d2852f",
				  "outofbandmanagement": {
					"enabled": false,
					"powerstate": "Disabled"
				  },
				  "podid": "5382edc2-e689-4074-bd67-0e1a236eb2bc",
				  "podname": "POD0",
				  "resourcestate": "Enabled",
				  "state": "Up",
				  "type": "Routing",
				  "version": "4.16.0.0-SNAPSHOT",
				  "zoneid": "d4a81f75-5d92-415e-ab59-e85cc2ce56d9",
				  "zonename": "Sandbox-simulator"
				}
			  ]
			}
		}`
		fmt.Fprintln(writer, response)
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
		responses := map[string]string{
			"prepareHostForMaintenance": `{
				"preparehostformaintenanceresponse": {
					"jobid": "0931656b-d30d-474b-aa79-13077d275400"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.host.PrepareForMaintenanceCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"host": {
							"id": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
							"name": "SimulatedAgent.4234d24b-37fd-42bf-9184-874889001baf",
							"state": "Up",
							"disconnected": "2021-09-28T20:05:04+0000",
							"type": "Routing",
							"oscategoryid": "e4992bed-5fdf-11ea-9a56-1e006800018c",
							"oscategoryname": "CentOS",
							"ipaddress": "172.17.1.6",
							"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
							"zonename": "shouldwork",
							"podid": "647e68f7-5de7-459c-91fc-f36ccf6af9ad",
							"podname": "poiuj",
							"version": "4.14.0.0-SNAPSHOT",
							"hypervisor": "Simulator",
							"cpunumber": 4,
							"cpuspeed": 8000,
							"cpuallocated": "0%",
							"cpuallocatedvalue": 0,
							"cpuallocatedpercentage": "0%",
							"cpuallocatedwithoverprovisioning": "0%",
							"cpuused": "0%",
							"cpuwithoverprovisioning": "32000",
							"cpuloadaverage": 0,
							"networkkbsread": 32768,
							"networkkbswrite": 16384,
							"memorytotal": 8589934592,
							"memorywithoverprovisioning": "8589934592",
							"memoryallocated": 0,
							"memoryallocatedpercentage": "0%",
							"memoryallocatedbytes": 0,
							"memoryused": 0,
							"capabilities": "hvm",
							"lastpinged": "1970-01-19T10:58:33+0000",
							"managementserverid": "18c349c6-ed36-4890-b670-4ef76e48485e",
							"clusterid": "841bb416-2d65-4661-80cf-db708c8a1c4f",
							"clustername": "lkjlkj",
							"clustertype": "CloudManaged",
							"islocalstorageactive": false,
							"created": "2020-03-25T09:06:40+0000",
							"events": "Ping; ShutdownRequested; PingTimeout; ManagementServerDown; AgentConnected; AgentDisconnected; Remove; StartAgentRebalance; HostDown",
							"hosttags": "null",
							"hostha": {
								"haenable": true,
								"hastate": "Disabled",
								"haprovider": "simulatorhaprovider"
							},
							"outofbandmanagement": {
								"powerstate": "Disabled",
								"enabled": false
							},
							"resourcestate": "PrepareForMaintenance",
							"hypervisorversion": "4.14.0.0-SNAPSHOT",
							"hahost": false,
							"jobid": "0931656b-d30d-474b-aa79-13077d275400",
							"jobstatus": 0
						}
					},
					"jobinstancetype": "Host",
					"jobinstanceid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
					"created": "2021-10-03T06:52:21+0000",
					"completed": "2021-10-03T06:52:21+0000",
					"jobid": "0931656b-d30d-474b-aa79-13077d275400"
				}
			}`,
		}
		fmt.Fprintln(writer, responses[request.FormValue("command")])
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
		responses := map[string]string{
			"cancelHostMaintenance": `{
				"cancelhostmaintenanceresponse": {
					"jobid": "b40f561d-9d0d-4eae-a376-fbafdff966f1"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.host.CancelMaintenanceCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"host": {
							"id": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
							"name": "SimulatedAgent.4234d24b-37fd-42bf-9184-874889001baf",
							"state": "Up",
							"disconnected": "2021-09-28T20:05:04+0000",
							"type": "Routing",
							"oscategoryid": "e4992bed-5fdf-11ea-9a56-1e006800018c",
							"oscategoryname": "CentOS",
							"ipaddress": "172.17.1.6",
							"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
							"zonename": "shouldwork",
							"podid": "647e68f7-5de7-459c-91fc-f36ccf6af9ad",
							"podname": "poiuj",
							"version": "4.14.0.0-SNAPSHOT",
							"hypervisor": "Simulator",
							"cpunumber": 4,
							"cpuspeed": 8000,
							"cpuallocated": "10.18%",
							"cpuallocatedvalue": 3256,
							"cpuallocatedpercentage": "10.18%",
							"cpuallocatedwithoverprovisioning": "10.18%",
							"cpuused": "0%",
							"cpuwithoverprovisioning": "32000",
							"cpuloadaverage": 0,
							"networkkbsread": 32768,
							"networkkbswrite": 16384,
							"memorytotal": 8589934592,
							"memorywithoverprovisioning": "8589934592",
							"memoryallocated": 1610612736,
							"memoryallocatedpercentage": "18.75%",
							"memoryallocatedbytes": 1610612736,
							"memoryused": 0,
							"capabilities": "hvm",
							"lastpinged": "1970-01-19T10:58:33+0000",
							"managementserverid": "18c349c6-ed36-4890-b670-4ef76e48485e",
							"clusterid": "841bb416-2d65-4661-80cf-db708c8a1c4f",
							"clustername": "lkjlkj",
							"clustertype": "CloudManaged",
							"islocalstorageactive": false,
							"created": "2020-03-25T09:06:40+0000",
							"events": "Ping; ShutdownRequested; PingTimeout; ManagementServerDown; AgentConnected; AgentDisconnected; Remove; StartAgentRebalance; HostDown",
							"hosttags": "null",
							"hostha": {
								"haenable": true,
								"hastate": "Disabled",
								"haprovider": "simulatorhaprovider"
							},
							"outofbandmanagement": {
								"powerstate": "Disabled",
								"enabled": false
							},
							"resourcestate": "Enabled",
							"hypervisorversion": "4.14.0.0-SNAPSHOT",
							"hahost": false,
							"jobid": "b40f561d-9d0d-4eae-a376-fbafdff966f1",
							"jobstatus": 0
						}
					},
					"jobinstancetype": "Host",
					"jobinstanceid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
					"created": "2021-10-03T06:57:12+0000",
					"completed": "2021-10-03T06:57:12+0000",
					"jobid": "b40f561d-9d0d-4eae-a376-fbafdff966f1"
				}
			}`,
		}
		fmt.Fprintln(writer, responses[request.FormValue("command")])
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
		responses := map[string]string{
			"enableOutOfBandManagementForHost": `{
				"enableoutofbandmanagementforhostresponse": {
					"jobid": "049f5a42-343b-4156-9fb6-683f01433dca"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.outofbandmanagement.EnableOutOfBandManagementForHostCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"outofbandmanagement": {
							"hostid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
							"enabled": true,
							"status": true
						}
					},
					"created": "2021-10-03T07:05:08+0000",
					"completed": "2021-10-03T07:05:08+0000",
					"jobid": "1c3022d1-c882-430c-ada7-ebf9d06d3e84"
				}
			}`,
		}
		fmt.Fprintln(writer, responses[request.FormValue("command")])
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
		responses := map[string]string{
			"disableOutOfBandManagementForHost": `{
				"disableoutofbandmanagementforhostresponse": {
					"jobid": "a389b543-ce49-410f-8f9a-831be0480ee5"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.outofbandmanagement.DisableOutOfBandManagementForHostCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"outofbandmanagement": {
							"hostid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
							"enabled": false,
							"status": true
						}
					},
					"created": "2021-10-03T07:06:25+0000",
					"completed": "2021-10-03T07:06:25+0000",
					"jobid": "a389b543-ce49-410f-8f9a-831be0480ee5"
				}
			}`,
		}
		fmt.Fprintln(writer, responses[request.FormValue("command")])
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
		responses := map[string]string{
			"enableHAForHost": `{
				"enablehaforhostresponse": {
					"jobid": "79da9d4b-6368-4ce5-9054-952295fc2c8a"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.ha.EnableHAForHostCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"hostha": {
							"hostid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
							"haenable": true,
							"status": true
						}
					},
					"created": "2021-10-03T07:10:09+0000",
					"completed": "2021-10-03T07:10:09+0000",
					"jobid": "79da9d4b-6368-4ce5-9054-952295fc2c8a"
				}
			}`,
		}
		fmt.Fprintln(writer, responses[request.FormValue("command")])
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
