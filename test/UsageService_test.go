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

func TestUsageService(t *testing.T) {
	service := "UsageService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddTrafficMonitor := func(t *testing.T) {
		if _, ok := response["addTrafficMonitor"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewAddTrafficMonitorParams("url", "zoneid")
		r, err := client.Usage.AddTrafficMonitor(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddTrafficMonitor", testaddTrafficMonitor)

	testaddTrafficType := func(t *testing.T) {
		if _, ok := response["addTrafficType"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewAddTrafficTypeParams("physicalnetworkid", "traffictype")
		r, err := client.Usage.AddTrafficType(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddTrafficType", testaddTrafficType)

	testdeleteTrafficMonitor := func(t *testing.T) {
		if _, ok := response["deleteTrafficMonitor"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewDeleteTrafficMonitorParams("id")
		_, err := client.Usage.DeleteTrafficMonitor(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTrafficMonitor", testdeleteTrafficMonitor)

	testdeleteTrafficType := func(t *testing.T) {
		if _, ok := response["deleteTrafficType"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewDeleteTrafficTypeParams("id")
		_, err := client.Usage.DeleteTrafficType(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTrafficType", testdeleteTrafficType)

	testgenerateUsageRecords := func(t *testing.T) {
		if _, ok := response["generateUsageRecords"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewGenerateUsageRecordsParams("enddate", "startdate")
		_, err := client.Usage.GenerateUsageRecords(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GenerateUsageRecords", testgenerateUsageRecords)

	testlistTrafficMonitors := func(t *testing.T) {
		if _, ok := response["listTrafficMonitors"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewListTrafficMonitorsParams("zoneid")
		_, err := client.Usage.ListTrafficMonitors(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTrafficMonitors", testlistTrafficMonitors)

	testlistTrafficTypeImplementors := func(t *testing.T) {
		if _, ok := response["listTrafficTypeImplementors"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewListTrafficTypeImplementorsParams()
		_, err := client.Usage.ListTrafficTypeImplementors(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTrafficTypeImplementors", testlistTrafficTypeImplementors)

	testlistTrafficTypes := func(t *testing.T) {
		if _, ok := response["listTrafficTypes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewListTrafficTypesParams("physicalnetworkid")
		_, err := client.Usage.ListTrafficTypes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTrafficTypes", testlistTrafficTypes)

	testlistUsageRecords := func(t *testing.T) {
		if _, ok := response["listUsageRecords"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewListUsageRecordsParams("enddate", "startdate")
		_, err := client.Usage.ListUsageRecords(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListUsageRecords", testlistUsageRecords)

	testlistUsageTypes := func(t *testing.T) {
		if _, ok := response["listUsageTypes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewListUsageTypesParams()
		_, err := client.Usage.ListUsageTypes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListUsageTypes", testlistUsageTypes)

	testremoveRawUsageRecords := func(t *testing.T) {
		if _, ok := response["removeRawUsageRecords"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewRemoveRawUsageRecordsParams(0)
		_, err := client.Usage.RemoveRawUsageRecords(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveRawUsageRecords", testremoveRawUsageRecords)

	testupdateTrafficType := func(t *testing.T) {
		if _, ok := response["updateTrafficType"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Usage.NewUpdateTrafficTypeParams("id")
		r, err := client.Usage.UpdateTrafficType(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateTrafficType", testupdateTrafficType)

}
