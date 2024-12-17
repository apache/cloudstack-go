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

func TestManagementService(t *testing.T) {
	service := "ManagementService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcancelShutdown := func(t *testing.T) {
		if _, ok := response["cancelShutdown"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Management.NewCancelShutdownParams("managementserverid")
		_, err := client.Management.CancelShutdown(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CancelShutdown", testcancelShutdown)

	testlistManagementServers := func(t *testing.T) {
		if _, ok := response["listManagementServers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Management.NewListManagementServersParams()
		_, err := client.Management.ListManagementServers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListManagementServers", testlistManagementServers)

	testlistManagementServersMetrics := func(t *testing.T) {
		if _, ok := response["listManagementServersMetrics"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Management.NewListManagementServersMetricsParams()
		_, err := client.Management.ListManagementServersMetrics(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListManagementServersMetrics", testlistManagementServersMetrics)

	testprepareForShutdown := func(t *testing.T) {
		if _, ok := response["prepareForShutdown"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Management.NewPrepareForShutdownParams("managementserverid")
		_, err := client.Management.PrepareForShutdown(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("PrepareForShutdown", testprepareForShutdown)

	testreadyForShutdown := func(t *testing.T) {
		if _, ok := response["readyForShutdown"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Management.NewReadyForShutdownParams()
		_, err := client.Management.ReadyForShutdown(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ReadyForShutdown", testreadyForShutdown)

	testtriggerShutdown := func(t *testing.T) {
		if _, ok := response["triggerShutdown"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Management.NewTriggerShutdownParams("managementserverid")
		_, err := client.Management.TriggerShutdown(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("TriggerShutdown", testtriggerShutdown)

}
