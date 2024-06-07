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

func TestSecurityGroupService(t *testing.T) {
	service := "SecurityGroupService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testauthorizeSecurityGroupEgress := func(t *testing.T) {
		if _, ok := response["authorizeSecurityGroupEgress"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SecurityGroup.NewAuthorizeSecurityGroupEgressParams()
		_, err := client.SecurityGroup.AuthorizeSecurityGroupEgress(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AuthorizeSecurityGroupEgress", testauthorizeSecurityGroupEgress)

	testauthorizeSecurityGroupIngress := func(t *testing.T) {
		if _, ok := response["authorizeSecurityGroupIngress"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SecurityGroup.NewAuthorizeSecurityGroupIngressParams()
		_, err := client.SecurityGroup.AuthorizeSecurityGroupIngress(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AuthorizeSecurityGroupIngress", testauthorizeSecurityGroupIngress)

	testcreateSecurityGroup := func(t *testing.T) {
		if _, ok := response["createSecurityGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SecurityGroup.NewCreateSecurityGroupParams("name")
		r, err := client.SecurityGroup.CreateSecurityGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateSecurityGroup", testcreateSecurityGroup)

	testdeleteSecurityGroup := func(t *testing.T) {
		if _, ok := response["deleteSecurityGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SecurityGroup.NewDeleteSecurityGroupParams()
		_, err := client.SecurityGroup.DeleteSecurityGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteSecurityGroup", testdeleteSecurityGroup)

	testlistSecurityGroups := func(t *testing.T) {
		if _, ok := response["listSecurityGroups"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SecurityGroup.NewListSecurityGroupsParams()
		_, err := client.SecurityGroup.ListSecurityGroups(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSecurityGroups", testlistSecurityGroups)

	testrevokeSecurityGroupEgress := func(t *testing.T) {
		if _, ok := response["revokeSecurityGroupEgress"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SecurityGroup.NewRevokeSecurityGroupEgressParams("id")
		_, err := client.SecurityGroup.RevokeSecurityGroupEgress(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RevokeSecurityGroupEgress", testrevokeSecurityGroupEgress)

	testrevokeSecurityGroupIngress := func(t *testing.T) {
		if _, ok := response["revokeSecurityGroupIngress"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SecurityGroup.NewRevokeSecurityGroupIngressParams("id")
		_, err := client.SecurityGroup.RevokeSecurityGroupIngress(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RevokeSecurityGroupIngress", testrevokeSecurityGroupIngress)

	testupdateSecurityGroup := func(t *testing.T) {
		if _, ok := response["updateSecurityGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SecurityGroup.NewUpdateSecurityGroupParams("id")
		r, err := client.SecurityGroup.UpdateSecurityGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateSecurityGroup", testupdateSecurityGroup)

}
