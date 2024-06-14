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

func TestPodService(t *testing.T) {
	service := "PodService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreatePod := func(t *testing.T) {
		if _, ok := response["createPod"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pod.NewCreatePodParams("name", "zoneid")
		r, err := client.Pod.CreatePod(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreatePod", testcreatePod)

	testdedicatePod := func(t *testing.T) {
		if _, ok := response["dedicatePod"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pod.NewDedicatePodParams("domainid", "podid")
		r, err := client.Pod.DedicatePod(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DedicatePod", testdedicatePod)

	testdeletePod := func(t *testing.T) {
		if _, ok := response["deletePod"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pod.NewDeletePodParams("id")
		_, err := client.Pod.DeletePod(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeletePod", testdeletePod)

	testlistDedicatedPods := func(t *testing.T) {
		if _, ok := response["listDedicatedPods"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pod.NewListDedicatedPodsParams()
		_, err := client.Pod.ListDedicatedPods(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListDedicatedPods", testlistDedicatedPods)

	testlistPods := func(t *testing.T) {
		if _, ok := response["listPods"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pod.NewListPodsParams()
		_, err := client.Pod.ListPods(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListPods", testlistPods)

	testreleaseDedicatedPod := func(t *testing.T) {
		if _, ok := response["releaseDedicatedPod"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pod.NewReleaseDedicatedPodParams("podid")
		_, err := client.Pod.ReleaseDedicatedPod(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReleaseDedicatedPod", testreleaseDedicatedPod)

	testupdatePod := func(t *testing.T) {
		if _, ok := response["updatePod"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pod.NewUpdatePodParams("id")
		r, err := client.Pod.UpdatePod(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdatePod", testupdatePod)

	testupdatePodManagementNetworkIpRange := func(t *testing.T) {
		if _, ok := response["updatePodManagementNetworkIpRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Pod.NewUpdatePodManagementNetworkIpRangeParams("currentendip", "currentstartip", "podid")
		_, err := client.Pod.UpdatePodManagementNetworkIpRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdatePodManagementNetworkIpRange", testupdatePodManagementNetworkIpRange)

}
