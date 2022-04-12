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

var (
	clusterName    = "TestCluster"
	clusterType    = "CloudManaged"
	hypervisorType = "KVM"
	podId          = "6137ef6f-753f-4e7b-a728-8d46a92358d2"
	zoneId         = "04ccc336-d730-42fe-8ff6-5ae36e141e81"
)

func TestClusterService_ListClusters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "listClusters"
		resp, err := ReadData(apiName, "ClusterService")
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, resp[apiName])
	}))
	defer server.Close()

	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
		apiName := "disableHAForCluster"
		responses, err := ParseAsyncResponse(apiName, "ClusterService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, responses)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
		apiName := "enableHAForCluster"
		responses, err := ParseAsyncResponse(apiName, "ClusterService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, responses)
	}))

	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
		apiName := "disableOutOfBandManagementForCluster"
		responses, err := ParseAsyncResponse(apiName, "ClusterService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, responses)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
		apiName := "enableOutOfBandManagementForCluster"
		responses, err := ParseAsyncResponse(apiName, "ClusterService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, responses)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
		apiName := "dedicateCluster"
		responses, err := ParseAsyncResponse(apiName, "ClusterService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, responses)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
		apiName := "releaseDedicatedCluster"
		responses, err := ParseAsyncResponse(apiName, "ClusterService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, responses)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Cluster.NewReleaseDedicatedClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	resp, err := client.Cluster.ReleaseDedicatedCluster(params)
	if err != nil {
		t.Errorf("Failed to release dedicated cluster to ROOT domain due to %v", err)
	}

	if !resp.Success {
		t.Errorf("Failed to release dedicated cluster to right domain")
	}
}

func TestClusterService_UpdateCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "updateCluster"
		response, err := ReadData(apiName, "ClusterService")
		if err != nil {
			t.Errorf("Failed to read response data, due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Cluster.NewUpdateClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	params.SetClustername("TestClusterUpdated")
	params.SetManagedstate("Unmanaged")
	resp, err := client.Cluster.UpdateCluster(params)
	if err != nil {
		t.Errorf("Failed to updated cluster details - name, due to: %v", err)
	}

	if resp.Name != "TestClusterUpdated" {
		t.Errorf("Failed to updated cluster name")
	}
}

func TestClusterService_DeleteCluster(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "deleteCluster"
		response, err := ReadData(apiName, "ClusterService")
		if err != nil {
			t.Errorf("Failed to read response data, due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Cluster.NewDeleteClusterParams("f1f5a259-1b55-431d-ad7e-6d39c28f1b15")
	resp, err := client.Cluster.DeleteCluster(params)
	if err != nil {
		t.Errorf("Failed to delete cluster due to: %v", err)
	}

	if !resp.Success {
		t.Errorf("Failed to delete cluster")
	}

}
