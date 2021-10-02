package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVirtualMachineService_DeployVirtualMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		responses := map[string]string{
			"deployVirtualMachine": `{
				"deployvirtualmachineresponse": {
					"id": "915653c4-298b-4d74-bdee-4ced282114f1",
					"jobid": "74bbb422-fb74-47ac-a22a-ad86a7831d3f"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.vm.DeployVMCmdByAdmin",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"virtualmachine": {
							"id": "915653c4-298b-4d74-bdee-4ced282114f1",
							"name": "testDummyVM",
							"displayname": "testDummyVM",
							"account": "admin",
							"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
							"username": "admin",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"created": "2021-10-04T09:24:38+0000",
							"state": "Running",
							"haenable": false,
							"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
							"zonename": "shouldwork",
							"hostid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
							"hostname": "SimulatedAgent.4234d24b-37fd-42bf-9184-874889001baf",
							"templateid": "50459b99-5fe0-11ea-9a56-1e006800018c",
							"templatename": "CentOS 5.6 (64-bit)",
							"templatedisplaytext": "CentOS 5.6 (64-bit) no GUI",
							"passwordenabled": false,
							"serviceofferingid": "498ce071-a077-4237-9492-e55c42553327",
							"serviceofferingname": "Very small instance",
							"cpunumber": 1,
							"cpuspeed": 100,
							"memory": 256,
							"guestosid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
							"rootdeviceid": 0,
							"rootdevicetype": "ROOT",
							"securitygroup": [],
							"nic": [
								{
									"id": "47d79da1-2fe1-4a44-a503-523055714a72",
									"networkid": "87503ba2-d59d-4a5e-8f7f-b3f486cedbd8",
									"networkname": "test",
									"netmask": "255.255.255.0",
									"gateway": "10.1.1.1",
									"ipaddress": "10.1.1.62",
									"isolationuri": "vlan://1850",
									"broadcasturi": "vlan://1850",
									"traffictype": "Guest",
									"type": "Isolated",
									"isdefault": true,
									"macaddress": "02:00:34:fb:00:01",
									"secondaryip": [],
									"extradhcpoption": []
								}
							],
							"hypervisor": "Simulator",
							"instancename": "i-2-683-QA",
							"details": {},
							"affinitygroup": [],
							"displayvm": true,
							"isdynamicallyscalable": false,
							"ostypeid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
							"osdisplayname": "CentOS 5.6 (64-bit)",
							"bootmode": "legacy",
							"boottype": "Bios",
							"tags": [],
							"jobid": "74bbb422-fb74-47ac-a22a-ad86a7831d3f",
							"jobstatus": 0
						}
					},
					"jobinstancetype": "VirtualMachine",
					"jobinstanceid": "915653c4-298b-4d74-bdee-4ced282114f1",
					"created": "2021-10-04T09:24:38+0000",
					"completed": "2021-10-04T09:24:42+0000",
					"jobid": "74bbb422-fb74-47ac-a22a-ad86a7831d3f"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[r.FormValue("command")])
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
		responses := map[string]string{
			"addNicToVirtualMachine": `{
				"addnictovirtualmachineresponse": {
					"jobid": "ec9d4711-b334-4750-868b-fb4d79d5801c"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.vm.AddNicToVMCmdByAdmin",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"virtualmachine": {
							"id": "915653c4-298b-4d74-bdee-4ced282114f1",
							"name": "testDummyVM",
							"displayname": "testDummyVM",
							"account": "admin",
							"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
							"username": "admin",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"created": "2021-10-04T09:24:38+0000",
							"state": "Running",
							"haenable": false,
							"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
							"zonename": "shouldwork",
							"hostid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
							"hostname": "SimulatedAgent.4234d24b-37fd-42bf-9184-874889001baf",
							"guestosid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
							"securitygroup": [],
							"nic": [
								{
									"id": "0921d201-f7ac-4f49-bf09-ffe9497fdab0",
									"networkid": "30358053-0f9d-4112-9948-976477896db6",
									"networkname": "test-network-2",
									"netmask": "255.255.255.0",
									"gateway": "10.1.2.1",
									"ipaddress": "10.1.2.102",
									"isolationuri": "vlan://1955",
									"broadcasturi": "vlan://1955",
									"traffictype": "Guest",
									"type": "Isolated",
									"isdefault": false,
									"macaddress": "02:00:13:a4:00:21",
									"secondaryip": [],
									"extradhcpoption": []
								},
								{
									"id": "47d79da1-2fe1-4a44-a503-523055714a72",
									"networkid": "87503ba2-d59d-4a5e-8f7f-b3f486cedbd8",
									"networkname": "test",
									"netmask": "255.255.255.0",
									"gateway": "10.1.1.1",
									"ipaddress": "10.1.1.62",
									"isolationuri": "vlan://1850",
									"broadcasturi": "vlan://1850",
									"traffictype": "Guest",
									"type": "Isolated",
									"isdefault": true,
									"macaddress": "02:00:34:fb:00:01",
									"secondaryip": [],
									"extradhcpoption": []
								}
							],
							"hypervisor": "Simulator",
							"instancename": "i-2-683-QA",
							"details": {},
							"affinitygroup": [],
							"displayvm": true,
							"isdynamicallyscalable": false,
							"ostypeid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
							"osdisplayname": "CentOS 5.6 (64-bit)",
							"bootmode": "legacy",
							"boottype": "Bios",
							"tags": []
						}
					},
					"created": "2021-10-04T09:29:51+0000",
					"completed": "2021-10-04T09:29:54+0000",
					"jobid": "ec9d4711-b334-4750-868b-fb4d79d5801c"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[r.FormValue("command")])
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
		responses := map[string]string{
			"stopVirtualMachine": `{
				"stopvirtualmachineresponse": {
					"jobid": "604bc5ed-1b7e-4d63-abc8-a55077a4c2c1"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.vm.StopVMCmdByAdmin",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"virtualmachine": {
							"id": "915653c4-298b-4d74-bdee-4ced282114f1",
							"name": "testDummyVM",
							"displayname": "testDummyVM",
							"account": "admin",
							"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
							"username": "admin",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"created": "2021-10-04T09:24:38+0000",
							"state": "Stopped",
							"haenable": false,
							"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
							"zonename": "shouldwork",
							"templateid": "50459b99-5fe0-11ea-9a56-1e006800018c",
							"templatename": "CentOS 5.6 (64-bit)",
							"templatedisplaytext": "CentOS 5.6 (64-bit) no GUI",
							"passwordenabled": false,
							"serviceofferingid": "498ce071-a077-4237-9492-e55c42553327",
							"serviceofferingname": "Very small instance",
							"cpunumber": 1,
							"cpuspeed": 100,
							"memory": 256,
							"cpuused": "10%",
							"networkkbsread": 917504,
							"networkkbswrite": 458752,
							"diskkbsread": 0,
							"diskkbswrite": 0,
							"memorykbs": 0,
							"memoryintfreekbs": 0,
							"memorytargetkbs": 0,
							"diskioread": 0,
							"diskiowrite": 0,
							"guestosid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
							"rootdeviceid": 0,
							"rootdevicetype": "ROOT",
							"securitygroup": [],
							"nic": [
								{
									"id": "47d79da1-2fe1-4a44-a503-523055714a72",
									"networkid": "87503ba2-d59d-4a5e-8f7f-b3f486cedbd8",
									"networkname": "test",
									"netmask": "255.255.255.0",
									"gateway": "10.1.1.1",
									"ipaddress": "10.1.1.62",
									"traffictype": "Guest",
									"type": "Isolated",
									"isdefault": true,
									"macaddress": "02:00:34:fb:00:01",
									"secondaryip": [],
									"extradhcpoption": []
								}
							],
							"hypervisor": "Simulator",
							"instancename": "i-2-683-QA",
							"details": {
								"Message.ReservedCapacityFreed.Flag": "false"
							},
							"affinitygroup": [],
							"displayvm": true,
							"isdynamicallyscalable": false,
							"ostypeid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
							"osdisplayname": "CentOS 5.6 (64-bit)",
							"tags": [],
							"jobid": "604bc5ed-1b7e-4d63-abc8-a55077a4c2c1",
							"jobstatus": 0
						}
					},
					"jobinstancetype": "VirtualMachine",
					"jobinstanceid": "915653c4-298b-4d74-bdee-4ced282114f1",
					"created": "2021-10-04T09:53:06+0000",
					"completed": "2021-10-04T09:53:08+0000",
					"jobid": "604bc5ed-1b7e-4d63-abc8-a55077a4c2c1"
				}
			}`,
		}
		fmt.Fprintln(writer, responses[r.FormValue("command")])
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
		responses := map[string]string{
			"startVirtualMachine": `{
				"startvirtualmachineresponse": {
					"jobid": "468d69d7-f327-4105-bec2-0e43107ae956"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.vm.StartVMCmdByAdmin",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"virtualmachine": {
							"id": "915653c4-298b-4d74-bdee-4ced282114f1",
							"name": "testDummyVM",
							"displayname": "testDummyVM",
							"account": "admin",
							"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
							"username": "admin",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"created": "2021-10-04T09:24:38+0000",
							"state": "Running",
							"haenable": false,
							"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
							"zonename": "shouldwork",
							"hostid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
							"hostname": "SimulatedAgent.4234d24b-37fd-42bf-9184-874889001baf",
							"templateid": "50459b99-5fe0-11ea-9a56-1e006800018c",
							"templatename": "CentOS 5.6 (64-bit)",
							"templatedisplaytext": "CentOS 5.6 (64-bit) no GUI",
							"passwordenabled": false,
							"serviceofferingid": "498ce071-a077-4237-9492-e55c42553327",
							"serviceofferingname": "Very small instance",
							"cpunumber": 1,
							"cpuspeed": 100,
							"memory": 256,
							"cpuused": "10%",
							"networkkbsread": 917504,
							"networkkbswrite": 458752,
							"diskkbsread": 0,
							"diskkbswrite": 0,
							"memorykbs": 0,
							"memoryintfreekbs": 0,
							"memorytargetkbs": 0,
							"diskioread": 0,
							"diskiowrite": 0,
							"guestosid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
							"rootdeviceid": 0,
							"rootdevicetype": "ROOT",
							"securitygroup": [],
							"nic": [
								{
									"id": "47d79da1-2fe1-4a44-a503-523055714a72",
									"networkid": "87503ba2-d59d-4a5e-8f7f-b3f486cedbd8",
									"networkname": "test",
									"netmask": "255.255.255.0",
									"gateway": "10.1.1.1",
									"ipaddress": "10.1.1.62",
									"isolationuri": "vlan://1994",
									"broadcasturi": "vlan://1994",
									"traffictype": "Guest",
									"type": "Isolated",
									"isdefault": true,
									"macaddress": "02:00:34:fb:00:01",
									"secondaryip": [],
									"extradhcpoption": []
								}
							],
							"hypervisor": "Simulator",
							"instancename": "i-2-683-QA",
							"details": {
								"Message.ReservedCapacityFreed.Flag": "false"
							},
							"affinitygroup": [],
							"displayvm": true,
							"isdynamicallyscalable": false,
							"ostypeid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
							"osdisplayname": "CentOS 5.6 (64-bit)",
							"tags": [],
							"jobid": "468d69d7-f327-4105-bec2-0e43107ae956",
							"jobstatus": 0
						}
					},
					"jobinstancetype": "VirtualMachine",
					"jobinstanceid": "915653c4-298b-4d74-bdee-4ced282114f1",
					"created": "2021-10-04T09:57:19+0000",
					"completed": "2021-10-04T09:57:22+0000",
					"jobid": "468d69d7-f327-4105-bec2-0e43107ae956"
				}
			}`,
		}
		fmt.Fprintln(writer, responses[r.FormValue("command")])
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
		response := `{
		"listVirtualMachines": {
		  "count": 1,
		  "virtualmachine": [
			{
			  "account": "admin",
			  "affinitygroup": [],
			  "cpunumber": 1,
			  "cpuspeed": 100,
			  "cpuused": "10%",
			  "created": "2021-10-04T09:24:38+0000",
			  "details": {
				"Message.ReservedCapacityFreed.Flag": "false"
			  },
			  "diskioread": 0,
			  "diskiowrite": 0,
			  "diskkbsread": 0,
			  "diskkbswrite": 0,
			  "displayname": "testDummyVM",
			  "displayvm": true,
			  "domain": "ROOT",
			  "domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
			  "guestosid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
			  "haenable": false,
			  "hostid": "8e8e68e7-19ea-4a78-bbdb-6e79d27729c9",
			  "hostname": "SimulatedAgent.4234d24b-37fd-42bf-9184-874889001baf",
			  "hypervisor": "Simulator",
			  "id": "915653c4-298b-4d74-bdee-4ced282114f1",
			  "instancename": "i-2-683-QA",
			  "isdynamicallyscalable": false,
			  "memory": 256,
			  "memoryintfreekbs": 0,
			  "memorykbs": 0,
			  "memorytargetkbs": 0,
			  "name": "testDummyVM",
			  "networkkbsread": 1015808,
			  "networkkbswrite": 507904,
			  "nic": [
				{
				  "broadcasturi": "vlan://1994",
				  "extradhcpoption": [],
				  "gateway": "10.1.1.1",
				  "id": "47d79da1-2fe1-4a44-a503-523055714a72",
				  "ipaddress": "10.1.1.62",
				  "isdefault": true,
				  "isolationuri": "vlan://1994",
				  "macaddress": "02:00:34:fb:00:01",
				  "netmask": "255.255.255.0",
				  "networkid": "87503ba2-d59d-4a5e-8f7f-b3f486cedbd8",
				  "networkname": "test",
				  "secondaryip": [],
				  "traffictype": "Guest",
				  "type": "Isolated"
				}
			  ],
			  "osdisplayname": "CentOS 5.6 (64-bit)",
			  "ostypeid": "e53f7606-5fdf-11ea-9a56-1e006800018c",
			  "passwordenabled": false,
			  "rootdeviceid": 0,
			  "rootdevicetype": "ROOT",
			  "securitygroup": [],
			  "serviceofferingid": "498ce071-a077-4237-9492-e55c42553327",
			  "serviceofferingname": "Very small instance",
			  "state": "Running",
			  "tags": [],
			  "templatedisplaytext": "CentOS 5.6 (64-bit) no GUI",
			  "templateid": "50459b99-5fe0-11ea-9a56-1e006800018c",
			  "templatename": "CentOS 5.6 (64-bit)",
			  "userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
			  "username": "admin",
			  "zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
			  "zonename": "shouldwork"
			}
		  ]
		}
	}`
		fmt.Fprintln(writer, response)
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
