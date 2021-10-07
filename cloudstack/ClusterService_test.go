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

var (
	clusterName    = "TestCluster"
	clusterType    = "CloudManaged"
	hypervisorType = "KVM"
	podId          = "6137ef6f-753f-4e7b-a728-8d46a92358d2"
	zoneId         = "04ccc336-d730-42fe-8ff6-5ae36e141e81"
)

//func TestCreateCluster(t *testing.T) {
//	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		response := `{
//			"addclusterresponse": {
//				"count": 1,
//				"cluster": [
//					{
//						"id": "f1f5a259-1b55-431d-ad7e-6d39c28f1b15",
//						"name": "TestCluster",
//						"podid": "6137ef6f-753f-4e7b-a728-8d46a92358d2",
//						"podname": "POD0",
//						"zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
//						"zonename": "SimZone1",
//						"hypervisortype": "KVM",
//						"clustertype": "CloudManaged",
//						"allocationstate": "Enabled",
//						"managedstate": "Managed",
//						"cpuovercommitratio": "1.0",
//						"memoryovercommitratio": "1.0",
//						"resourcedetails": {
//							"memoryOvercommitRatio": "1.0",
//							"cpuOvercommitRatio": "1.0"
//						}
//					}
//				]
//			}
//		}`
//		fmt.Fprintf(writer, response)
//	}))
//	defer server.Close()
//
//	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
//	params := client.Cluster.NewAddClusterParams(clusterName, clusterType, hypervisorType, podId, zoneId)
//	resp, err := client.Cluster.AddCluster(params)
//
//	if err != nil {
//		t.Errorf("Failed to add cluster %s to zone %s, due to: %v", clusterName, zoneId, err)
//	}
//
//	fmt.Printf("%v", resp)
//	if resp.Name != clusterName {
//		t.Errorf("Failed to add cluster")
//	}
//}

func TestClusterService_ListClusters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		resp := `{
"listclustersresponse":
{
  "cluster": [
    {
      "allocationstate": "Enabled",
      "clustertype": "CloudManaged",
      "cpuovercommitratio": "1.0",
      "hypervisortype": "Simulator",
      "id": "38f93102-4f87-49ab-bafa-854bdf31cfe6",
      "managedstate": "Managed",
      "memoryovercommitratio": "1.0",
      "name": "SimCluster0",
      "podid": "6137ef6f-753f-4e7b-a728-8d46a92358d2",
      "podname": "POD0",
      "resourcedetails": {
        "cluster.cpu.allocated.capacity.disablethreshold": "0.85",
        "cluster.cpu.allocated.capacity.notificationthreshold": "0.01",
        "cluster.memory.allocated.capacity.disablethreshold": "0.02",
        "cluster.memory.allocated.capacity.notificationthreshold": "0.01",
        "cluster.storage.allocated.capacity.notificationthreshold": "0.01",
        "cluster.storage.capacity.notificationthreshold": "0.01",
        "cpuOvercommitRatio": "1.0",
        "memoryOvercommitRatio": "1.0",
        "outOfBandManagementEnabled": "true",
        "resourceHAEnabled": "true"
      },
      "zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
      "zonename": "SimZone1"
    },
    {
      "allocationstate": "Enabled",
      "clustertype": "CloudManaged",
      "cpuovercommitratio": "1.0",
      "hypervisortype": "Simulator",
      "id": "388a9023-f0a9-43f1-a962-b8689e0fca09",
      "managedstate": "Unmanaged",
      "memoryovercommitratio": "1.0",
      "name": "SimCluster1",
      "podid": "6137ef6f-753f-4e7b-a728-8d46a92358d2",
      "podname": "POD0",
      "resourcedetails": {
        "cpuOvercommitRatio": "1.0",
        "memoryOvercommitRatio": "1.0"
      },
      "zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
      "zonename": "SimZone1"
    },
    {
      "allocationstate": "Enabled",
      "clustertype": "CloudManaged",
      "cpuovercommitratio": "1.0",
      "hypervisortype": "KVM",
      "id": "f1f5a259-1b55-431d-ad7e-6d39c28f1b15",
      "managedstate": "Managed",
      "memoryovercommitratio": "1.0",
      "name": "TestCluster",
      "podid": "6137ef6f-753f-4e7b-a728-8d46a92358d2",
      "podname": "POD0",
      "resourcedetails": {
        "cpuOvercommitRatio": "1.0",
        "memoryOvercommitRatio": "1.0"
      },
      "zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
      "zonename": "SimZone1"
    }
  ],
  "count": 3
}
}`
		fmt.Fprintln(writer, resp)
	}))
	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	params := client.Cluster.NewListClustersParams()
	params.SetZoneid(zoneId)
	clusterResp, err := client.Cluster.ListClusters(params)
	if err != nil {
		t.Errorf("Failed to list clusters in zone %s due to: %v", zoneId, err)
	}

	if clusterResp.Count != 3 {
		t.Errorf("Failed to list all clusters in the zone")
	}
}

func TestClusterService_DisableHAForCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"disableHAForCluster": `{
					"disablehaforclusterresponse": {
						"jobid": "c0591985-d022-4a2d-b5b5-1ac209849f66"
					}
				}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.ha.DisableHAForClusterCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"success": true
					},
					"created": "2021-10-02T08:59:13+0000",
					"completed": "2021-10-02T08:59:13+0000",
					"jobid": "c0591985-d022-4a2d-b5b5-1ac209849f66"
				}
			}
			`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Cluster.NewDisableHAForClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	resp, err := client.Cluster.DisableHAForCluster(params)
	if err != nil {
		t.Errorf("Failed to disable HA for cluster due to %v", err)
	}

	if !resp.Success {
		t.Errorf("Failed to disable HA for cluster")
	}
}

func TestClusterService_EnableHAForCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"enableHAForCluster": `{
				"enablehaforclusterresponse": {
					"jobid": "875ead2d-7b78-4bb0-83f5-c40a2a944dde"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.ha.EnableHAForClusterCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"success": true
					},
					"created": "2021-10-02T09:05:15+0000",
					"completed": "2021-10-02T09:05:15+0000",
					"jobid": "875ead2d-7b78-4bb0-83f5-c40a2a944dde"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))

	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Cluster.NewEnableHAForClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	resp, err := client.Cluster.EnableHAForCluster(params)
	if err != nil {
		t.Errorf("Failed to enable HA for cluster due to %v", err)
	}

	if !resp.Success {
		t.Errorf("Failed to enable HA for cluster")
	}
}

func TestClusterService_DisableOutOfBandManagementForCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"disableOutOfBandManagementForCluster": `{
				"disableoutofbandmanagementforclusterresponse": {
					"jobid": "067b171f-ca63-4d41-bb87-461589e20223"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.outofbandmanagement.DisableOutOfBandManagementForClusterCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"outofbandmanagement": {
							"enabled": false,
							"status": true
						}
					},
					"created": "2021-10-02T09:09:04+0000",
					"completed": "2021-10-02T09:09:04+0000",
					"jobid": "067b171f-ca63-4d41-bb87-461589e20223"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Cluster.NewDisableOutOfBandManagementForClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	resp, err := client.Cluster.DisableOutOfBandManagementForCluster(params)
	if err != nil {
		t.Errorf("Failed to disable Out of band management for cluster due to %v", err)
	}

	if resp.Enabled {
		t.Errorf("Failed to disable out of band management for cluster")
	}
}

func TestClusterService_EnableOutOfBandManagementForCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"enableOutOfBandManagementForCluster": `{
				"enableoutofbandmanagementforclusterresponse": {
					"jobid": "16f517a5-247b-4f96-aba4-0e876e1e9720"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.outofbandmanagement.EnableOutOfBandManagementForClusterCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"outofbandmanagement": {
							"enabled": true,
							"status": true
						}
					},
					"created": "2021-10-02T09:14:30+0000",
					"completed": "2021-10-02T09:14:30+0000",
					"jobid": "16f517a5-247b-4f96-aba4-0e876e1e9720"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Cluster.NewEnableOutOfBandManagementForClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	resp, err := client.Cluster.EnableOutOfBandManagementForCluster(params)
	if err != nil {
		t.Errorf("Failed to enable Out of band management for cluster due to %v", err)
	}

	if !resp.Enabled {
		t.Errorf("Failed to enable out of band management for cluster")
	}
}

func TestClusterService_DedicateCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"dedicateCluster": `{
				"dedicateclusterresponse": {
					"jobid": "64e777d8-9c72-4639-aa75-057971ecd289"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.commands.DedicateClusterCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"dedicatedcluster": {
							"id": "60a97ad1-702b-40dd-9ca5-a67954a28fbd",
							"clusterid": "f1f5a259-1b55-431d-ad7e-6d39c28f1b15",
							"clustername": "TestCluster",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
							"affinitygroupid": "a8725454-1680-4517-85ce-6f8ec08300fb"
						}
					},
					"created": "2021-10-02T09:16:25+0000",
					"completed": "2021-10-02T09:16:25+0000",
					"jobid": "64e777d8-9c72-4639-aa75-057971ecd289"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Cluster.NewDedicateClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15", "e4874e10-5fdf-11ea-9a56-1e006800018c")
	resp, err := client.Cluster.DedicateCluster(params)
	if err != nil {
		t.Errorf("Failed to dedicate cluster to ROOT domain due to %v", err)
	}

	if resp.Domainid != "e4874e10-5fdf-11ea-9a56-1e006800018c" {
		t.Errorf("Failed to dedicate cluster to right domain")
	}
}

func TestClusterService_ReleaseDedicatedCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"releaseDedicatedCluster": `{
				"releasededicatedclusterresponse": {
					"jobid": "e7cb7eb7-b77d-4568-9bcf-11bb5fd78938"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.commands.ReleaseDedicatedClusterCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"success": true
					},
					"created": "2021-10-02T14:40:30+0000",
					"completed": "2021-10-02T14:40:30+0000",
					"jobid": "e7cb7eb7-b77d-4568-9bcf-11bb5fd78938"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Cluster.NewReleaseDedicatedClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	resp, err := client.Cluster.ReleaseDedicatedCluster(params)
	if err != nil {
		t.Errorf("Failed to release dedicated cluster to ROOT domain due to %v", err)
	}

	if !resp.Success {
		t.Errorf("Failed to release dedicated cluster to right domain")
	}
}

//func TestClusterService_UpdateCluster(t *testing.T) {
//	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		response := `{
//			"updateclusterresponse": {
//				"cluster": {
//					"id": "f1f5a259-1b55-431d-ad7e-6d39c28f1b15",
//					"name": "TestClusterUpdated",
//					"podid": "6137ef6f-753f-4e7b-a728-8d46a92358d2",
//					"podname": "POD0",
//					"zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
//					"zonename": "SimZone1",
//					"hypervisortype": "KVM",
//					"clustertype": "CloudManaged",
//					"allocationstate": "Enabled",
//					"managedstate": "Managed",
//					"cpuovercommitratio": "1.0",
//					"memoryovercommitratio": "1.0",
//					"resourcedetails": {
//						"cpuOvercommitRatio": "1.0",
//						"resourceHAEnabled": "true",
//						"memoryOvercommitRatio": "1.0",
//						"outOfBandManagementEnabled": "true"
//					}
//				}
//			}
//		}`
//		fmt.Fprintf(writer, response)
//	}))
//	defer server.Close()
//	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
//	params := client.Cluster.NewUpdateClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
//	params.SetClustername("TestClusterUpdated")
//	params.SetManagedstate("Unmanaged")
//	resp, err := client.Cluster.UpdateCluster(params)
//	if err != nil {
//		t.Errorf("Failed to updated cluster details - name, due to: %v", err)
//	}
//
//	fmt.Println(resp)
//	if resp.Name != "TestClusterUpdated" {
//		t.Errorf("Failed to updated cluster name")
//	}
//}

func TestClusterService_DeleteCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := `{
			"deleteclusterresponse": {
				"success": true
			}
		}`
		fmt.Fprintf(writer, responses)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Cluster.NewDeleteClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	resp, err := client.Cluster.DeleteCluster(params)
	if err != nil {
		t.Errorf("Failed to delete cluster due to: %v", err)
	}

	if !resp.Success {
		t.Errorf("Failed to delete cluster")
	}

}
