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

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestKubernetesService(t *testing.T) {
	service := "KubernetesService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddKubernetesSupportedVersion := func(t *testing.T) {
		if _, ok := response["addKubernetesSupportedVersion"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewAddKubernetesSupportedVersionParams(0, 0, "semanticversion")
		r, err := client.Kubernetes.AddKubernetesSupportedVersion(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddKubernetesSupportedVersion", testaddKubernetesSupportedVersion)

	testcreateKubernetesCluster := func(t *testing.T) {
		if _, ok := response["createKubernetesCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewCreateKubernetesClusterParams("description", "kubernetesversionid", "name", "serviceofferingid", 0, "zoneid")
		r, err := client.Kubernetes.CreateKubernetesCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateKubernetesCluster", testcreateKubernetesCluster)

	testdeleteKubernetesCluster := func(t *testing.T) {
		if _, ok := response["deleteKubernetesCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewDeleteKubernetesClusterParams("id")
		_, err := client.Kubernetes.DeleteKubernetesCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteKubernetesCluster", testdeleteKubernetesCluster)

	testdeleteKubernetesSupportedVersion := func(t *testing.T) {
		if _, ok := response["deleteKubernetesSupportedVersion"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewDeleteKubernetesSupportedVersionParams("id")
		_, err := client.Kubernetes.DeleteKubernetesSupportedVersion(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteKubernetesSupportedVersion", testdeleteKubernetesSupportedVersion)

	testgetKubernetesClusterConfig := func(t *testing.T) {
		if _, ok := response["getKubernetesClusterConfig"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewGetKubernetesClusterConfigParams()
		r, err := client.Kubernetes.GetKubernetesClusterConfig(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("GetKubernetesClusterConfig", testgetKubernetesClusterConfig)

	testlistKubernetesClusters := func(t *testing.T) {
		if _, ok := response["listKubernetesClusters"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewListKubernetesClustersParams()
		_, err := client.Kubernetes.ListKubernetesClusters(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListKubernetesClusters", testlistKubernetesClusters)

	testlistKubernetesSupportedVersions := func(t *testing.T) {
		if _, ok := response["listKubernetesSupportedVersions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewListKubernetesSupportedVersionsParams()
		_, err := client.Kubernetes.ListKubernetesSupportedVersions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListKubernetesSupportedVersions", testlistKubernetesSupportedVersions)

	testscaleKubernetesCluster := func(t *testing.T) {
		if _, ok := response["scaleKubernetesCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewScaleKubernetesClusterParams("id")
		r, err := client.Kubernetes.ScaleKubernetesCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ScaleKubernetesCluster", testscaleKubernetesCluster)

	teststartKubernetesCluster := func(t *testing.T) {
		if _, ok := response["startKubernetesCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewStartKubernetesClusterParams("id")
		r, err := client.Kubernetes.StartKubernetesCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("StartKubernetesCluster", teststartKubernetesCluster)

	teststopKubernetesCluster := func(t *testing.T) {
		if _, ok := response["stopKubernetesCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewStopKubernetesClusterParams("id")
		_, err := client.Kubernetes.StopKubernetesCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("StopKubernetesCluster", teststopKubernetesCluster)

	testupdateKubernetesSupportedVersion := func(t *testing.T) {
		if _, ok := response["updateKubernetesSupportedVersion"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewUpdateKubernetesSupportedVersionParams("id", "state")
		r, err := client.Kubernetes.UpdateKubernetesSupportedVersion(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateKubernetesSupportedVersion", testupdateKubernetesSupportedVersion)

	testupgradeKubernetesCluster := func(t *testing.T) {
		if _, ok := response["upgradeKubernetesCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewUpgradeKubernetesClusterParams("id", "kubernetesversionid")
		r, err := client.Kubernetes.UpgradeKubernetesCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpgradeKubernetesCluster", testupgradeKubernetesCluster)

	testaddVirtualMachinesToKubernetesCluster := func(t *testing.T) {
		if _, ok := response["addVirtualMachinesToKubernetesCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewAddVirtualMachinesToKubernetesClusterParams("id", []string{})
		_, err := client.Kubernetes.AddVirtualMachinesToKubernetesCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddVirtualMachinesToKubernetesCluster", testaddVirtualMachinesToKubernetesCluster)

	testremoveVirtualMachinesFromKubernetesCluster := func(t *testing.T) {
		if _, ok := response["removeVirtualMachinesFromKubernetesCluster"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Kubernetes.NewRemoveVirtualMachinesFromKubernetesClusterParams("id", []string{})
		r, err := client.Kubernetes.RemoveVirtualMachinesFromKubernetesCluster(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RemoveVirtualMachinesFromKubernetesCluster", testremoveVirtualMachinesFromKubernetesCluster)

}
