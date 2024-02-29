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

func TestOutofbandManagementService(t *testing.T) {
	service := "OutofbandManagementService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testchangeOutOfBandManagementPassword := func(t *testing.T) {
		if _, ok := response["changeOutOfBandManagementPassword"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.OutofbandManagement.NewChangeOutOfBandManagementPasswordParams("hostid")
		_, err := client.OutofbandManagement.ChangeOutOfBandManagementPassword(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ChangeOutOfBandManagementPassword", testchangeOutOfBandManagementPassword)

	testconfigureOutOfBandManagement := func(t *testing.T) {
		if _, ok := response["configureOutOfBandManagement"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.OutofbandManagement.NewConfigureOutOfBandManagementParams("address", "driver", "hostid", "password", "port", "username")
		_, err := client.OutofbandManagement.ConfigureOutOfBandManagement(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ConfigureOutOfBandManagement", testconfigureOutOfBandManagement)

	testissueOutOfBandManagementPowerAction := func(t *testing.T) {
		if _, ok := response["issueOutOfBandManagementPowerAction"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.OutofbandManagement.NewIssueOutOfBandManagementPowerActionParams("action", "hostid")
		_, err := client.OutofbandManagement.IssueOutOfBandManagementPowerAction(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("IssueOutOfBandManagementPowerAction", testissueOutOfBandManagementPowerAction)

}
