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
	"testing"

	"github.com/sbrueseke/cloudstack-go/v2/cloudstack"
)

func TestClusterService(t *testing.T) {
	service := "ClusterService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddCluster := func(t *testing.T) {
		if _, ok := response["addCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewAddClusterParams("clustername", "clustertype", "hypervisor", "podid", "zoneid")
		r, err := client.Cluster.AddCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddCluster", testaddCluster)

	testdedicateCluster := func(t *testing.T) {
		if _, ok := response["dedicateCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewDedicateClusterParams("clusterid", "domainid")
		r, err := client.Cluster.DedicateCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DedicateCluster", testdedicateCluster)

	testdeleteCluster := func(t *testing.T) {
		if _, ok := response["deleteCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewDeleteClusterParams("id")
		_, err := client.Cluster.DeleteCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteCluster", testdeleteCluster)

	testdisableOutOfBandManagementForCluster := func(t *testing.T) {
		if _, ok := response["disableOutOfBandManagementForCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewDisableOutOfBandManagementForClusterParams("clusterid")
		_, err := client.Cluster.DisableOutOfBandManagementForCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DisableOutOfBandManagementForCluster", testdisableOutOfBandManagementForCluster)

	testenableOutOfBandManagementForCluster := func(t *testing.T) {
		if _, ok := response["enableOutOfBandManagementForCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewEnableOutOfBandManagementForClusterParams("clusterid")
		_, err := client.Cluster.EnableOutOfBandManagementForCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("EnableOutOfBandManagementForCluster", testenableOutOfBandManagementForCluster)

	testenableHAForCluster := func(t *testing.T) {
		if _, ok := response["enableHAForCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewEnableHAForClusterParams("clusterid")
		_, err := client.Cluster.EnableHAForCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("EnableHAForCluster", testenableHAForCluster)

	testdisableHAForCluster := func(t *testing.T) {
		if _, ok := response["disableHAForCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewDisableHAForClusterParams("clusterid")
		_, err := client.Cluster.DisableHAForCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DisableHAForCluster", testdisableHAForCluster)

	testlistClusters := func(t *testing.T) {
		if _, ok := response["listClusters"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewListClustersParams()
		_, err := client.Cluster.ListClusters(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListClusters", testlistClusters)

	testlistClustersMetrics := func(t *testing.T) {
		if _, ok := response["listClustersMetrics"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewListClustersMetricsParams()
		_, err := client.Cluster.ListClustersMetrics(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListClustersMetrics", testlistClustersMetrics)

	testlistDedicatedClusters := func(t *testing.T) {
		if _, ok := response["listDedicatedClusters"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewListDedicatedClustersParams()
		_, err := client.Cluster.ListDedicatedClusters(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListDedicatedClusters", testlistDedicatedClusters)

	testreleaseDedicatedCluster := func(t *testing.T) {
		if _, ok := response["releaseDedicatedCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewReleaseDedicatedClusterParams("clusterid")
		_, err := client.Cluster.ReleaseDedicatedCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReleaseDedicatedCluster", testreleaseDedicatedCluster)

	testupdateCluster := func(t *testing.T) {
		if _, ok := response["updateCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Cluster.NewUpdateClusterParams("id")
		r, err := client.Cluster.UpdateCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateCluster", testupdateCluster)

}
