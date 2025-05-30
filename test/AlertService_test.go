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

func TestAlertService(t *testing.T) {
	service := "AlertService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testarchiveAlerts := func(t *testing.T) {
		if _, ok := response["archiveAlerts"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Alert.NewArchiveAlertsParams()
		_, err := client.Alert.ArchiveAlerts(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ArchiveAlerts", testarchiveAlerts)

	testdeleteAlerts := func(t *testing.T) {
		if _, ok := response["deleteAlerts"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Alert.NewDeleteAlertsParams()
		_, err := client.Alert.DeleteAlerts(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteAlerts", testdeleteAlerts)

	testgenerateAlert := func(t *testing.T) {
		if _, ok := response["generateAlert"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Alert.NewGenerateAlertParams("description", "name", 0)
		_, err := client.Alert.GenerateAlert(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GenerateAlert", testgenerateAlert)

	testlistAlerts := func(t *testing.T) {
		if _, ok := response["listAlerts"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Alert.NewListAlertsParams()
		_, err := client.Alert.ListAlerts(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListAlerts", testlistAlerts)

	testlistAlertTypes := func(t *testing.T) {
		if _, ok := response["listAlertTypes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Alert.NewListAlertTypesParams()
		_, err := client.Alert.ListAlertTypes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListAlertTypes", testlistAlertTypes)

}
