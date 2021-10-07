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

func TestStoragePoolService_CancelStorageMaintenance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		responses := map[string]string{
			"cancelStorageMaintenance": `{
				"cancelprimarystoragemaintenanceresponse": {
					"jobid": "7cc3faa4-73c5-45cb-a982-683a13023827"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.storage.CancelPrimaryStorageMaintenanceCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"storagepool": {
							"id": "44ad900d-785b-3eff-addc-a5e6bf4927ef",
							"zoneid": "753941b6-5239-4ab3-b2b3-7522d2a6d1ba",
							"zonename": "DummySG",
							"podid": "b52f7877-227e-4e3f-974b-2397a37d6e1c",
							"podname": "SGAdvPod",
							"name": "SGAdvPS",
							"ipaddress": "10.1.1.3",
							"path": "/export/testing/primary ",
							"created": "2020-07-27T06:01:56+0000",
							"type": "NetworkFilesystem",
							"clusterid": "a7b4c0e1-8ba4-41ea-a2f4-490798ddd166",
							"clustername": "SGAdvCluster",
							"disksizetotal": 1099511627776,
							"disksizeallocated": 100,
							"state": "Up",
							"scope": "CLUSTER",
							"overprovisionfactor": "2.0",
							"provider": "DefaultPrimary",
							"jobid": "7cc3faa4-73c5-45cb-a982-683a13023827",
							"jobstatus": 0
						}
					},
					"jobinstancetype": "StoragePool",
					"jobinstanceid": "44ad900d-785b-3eff-addc-a5e6bf4927ef",
					"created": "2021-10-04T08:06:55+0000",
					"completed": "2021-10-04T08:06:55+0000",
					"jobid": "7cc3faa4-73c5-45cb-a982-683a13023827"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[r.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.StoragePool.NewCancelStorageMaintenanceParams("44ad900d-785b-3eff-addc-a5e6bf4927ef")
	resp, err := client.StoragePool.CancelStorageMaintenance(params)
	if err != nil {
		t.Errorf("Failed to cancel storage maintenance due to %v", err)
		return
	}
	if resp == nil || resp.State != "Up" {
		t.Errorf("Failed to cancel Storage pool maintenance")
	}
}
func TestStoragePoolService_EnableStorageMaintenance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		responses := map[string]string{
			"enableStorageMaintenance": `{
				"prepareprimarystorageformaintenanceresponse": {
					"jobid": "3036df18-12ca-44a0-9c77-19fcf81dad75"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.storage.PreparePrimaryStorageForMaintenanceCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"storagepool": {
							"id": "44ad900d-785b-3eff-addc-a5e6bf4927ef",
							"zoneid": "753941b6-5239-4ab3-b2b3-7522d2a6d1ba",
							"zonename": "DummySG",
							"podid": "b52f7877-227e-4e3f-974b-2397a37d6e1c",
							"podname": "SGAdvPod",
							"name": "SGAdvPS",
							"ipaddress": "10.1.1.3",
							"path": "/export/testing/primary ",
							"created": "2020-07-27T06:01:56+0000",
							"type": "NetworkFilesystem",
							"clusterid": "a7b4c0e1-8ba4-41ea-a2f4-490798ddd166",
							"clustername": "SGAdvCluster",
							"disksizetotal": 1099511627776,
							"disksizeallocated": 100,
							"state": "Maintenance",
							"scope": "CLUSTER",
							"overprovisionfactor": "2.0",
							"provider": "DefaultPrimary",
							"jobid": "3036df18-12ca-44a0-9c77-19fcf81dad75",
							"jobstatus": 0
						}
					},
					"jobinstancetype": "StoragePool",
					"jobinstanceid": "44ad900d-785b-3eff-addc-a5e6bf4927ef",
					"created": "2021-10-04T08:13:01+0000",
					"completed": "2021-10-04T08:13:02+0000",
					"jobid": "3036df18-12ca-44a0-9c77-19fcf81dad75"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[r.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.StoragePool.NewEnableStorageMaintenanceParams("44ad900d-785b-3eff-addc-a5e6bf4927ef")
	resp, err := client.StoragePool.EnableStorageMaintenance(params)
	if err != nil {
		t.Errorf("Failed to enable storage maintenance due to %v", err)
		return
	}
	if resp == nil || resp.State != "Maintenance" {
		t.Errorf("Failed to enable Storage pool maintenance")
	}
}
