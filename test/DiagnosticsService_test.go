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

func TestDiagnosticsService(t *testing.T) {
	service := "DiagnosticsService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testgetDiagnosticsData := func(t *testing.T) {
		if _, ok := response["getDiagnosticsData"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Diagnostics.NewGetDiagnosticsDataParams("targetid")
		_, err := client.Diagnostics.GetDiagnosticsData(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetDiagnosticsData", testgetDiagnosticsData)

	testrunDiagnostics := func(t *testing.T) {
		if _, ok := response["runDiagnostics"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Diagnostics.NewRunDiagnosticsParams("ipaddress", "targetid", "type")
		_, err := client.Diagnostics.RunDiagnostics(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RunDiagnostics", testrunDiagnostics)

}
