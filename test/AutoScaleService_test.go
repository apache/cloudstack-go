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

func TestAutoScaleService(t *testing.T) {
	service := "AutoScaleService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateAutoScalePolicy := func(t *testing.T) {
		if _, ok := response["createAutoScalePolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewCreateAutoScalePolicyParams("action", []string{}, 0)
		r, err := client.AutoScale.CreateAutoScalePolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateAutoScalePolicy", testcreateAutoScalePolicy)

	testcreateAutoScaleVmGroup := func(t *testing.T) {
		if _, ok := response["createAutoScaleVmGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewCreateAutoScaleVmGroupParams("lbruleid", 0, 0, []string{}, []string{}, "vmprofileid")
		r, err := client.AutoScale.CreateAutoScaleVmGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateAutoScaleVmGroup", testcreateAutoScaleVmGroup)

	testcreateAutoScaleVmProfile := func(t *testing.T) {
		if _, ok := response["createAutoScaleVmProfile"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewCreateAutoScaleVmProfileParams("serviceofferingid", "templateid", "zoneid")
		r, err := client.AutoScale.CreateAutoScaleVmProfile(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateAutoScaleVmProfile", testcreateAutoScaleVmProfile)

	testcreateCondition := func(t *testing.T) {
		if _, ok := response["createCondition"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewCreateConditionParams("counterid", "relationaloperator", 0)
		r, err := client.AutoScale.CreateCondition(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateCondition", testcreateCondition)

	testcreateCounter := func(t *testing.T) {
		if _, ok := response["createCounter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewCreateCounterParams("name", "provider", "source", "value")
		r, err := client.AutoScale.CreateCounter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateCounter", testcreateCounter)

	testdeleteAutoScalePolicy := func(t *testing.T) {
		if _, ok := response["deleteAutoScalePolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewDeleteAutoScalePolicyParams("id")
		_, err := client.AutoScale.DeleteAutoScalePolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteAutoScalePolicy", testdeleteAutoScalePolicy)

	testdeleteAutoScaleVmGroup := func(t *testing.T) {
		if _, ok := response["deleteAutoScaleVmGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewDeleteAutoScaleVmGroupParams("id")
		_, err := client.AutoScale.DeleteAutoScaleVmGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteAutoScaleVmGroup", testdeleteAutoScaleVmGroup)

	testdeleteAutoScaleVmProfile := func(t *testing.T) {
		if _, ok := response["deleteAutoScaleVmProfile"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewDeleteAutoScaleVmProfileParams("id")
		_, err := client.AutoScale.DeleteAutoScaleVmProfile(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteAutoScaleVmProfile", testdeleteAutoScaleVmProfile)

	testdeleteCondition := func(t *testing.T) {
		if _, ok := response["deleteCondition"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewDeleteConditionParams("id")
		_, err := client.AutoScale.DeleteCondition(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteCondition", testdeleteCondition)

	testdeleteCounter := func(t *testing.T) {
		if _, ok := response["deleteCounter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewDeleteCounterParams("id")
		_, err := client.AutoScale.DeleteCounter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteCounter", testdeleteCounter)

	testdisableAutoScaleVmGroup := func(t *testing.T) {
		if _, ok := response["disableAutoScaleVmGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewDisableAutoScaleVmGroupParams("id")
		r, err := client.AutoScale.DisableAutoScaleVmGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DisableAutoScaleVmGroup", testdisableAutoScaleVmGroup)

	testenableAutoScaleVmGroup := func(t *testing.T) {
		if _, ok := response["enableAutoScaleVmGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewEnableAutoScaleVmGroupParams("id")
		r, err := client.AutoScale.EnableAutoScaleVmGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("EnableAutoScaleVmGroup", testenableAutoScaleVmGroup)

	testlistAutoScalePolicies := func(t *testing.T) {
		if _, ok := response["listAutoScalePolicies"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewListAutoScalePoliciesParams()
		_, err := client.AutoScale.ListAutoScalePolicies(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListAutoScalePolicies", testlistAutoScalePolicies)

	testlistAutoScaleVmGroups := func(t *testing.T) {
		if _, ok := response["listAutoScaleVmGroups"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewListAutoScaleVmGroupsParams()
		_, err := client.AutoScale.ListAutoScaleVmGroups(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListAutoScaleVmGroups", testlistAutoScaleVmGroups)

	testlistAutoScaleVmProfiles := func(t *testing.T) {
		if _, ok := response["listAutoScaleVmProfiles"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewListAutoScaleVmProfilesParams()
		_, err := client.AutoScale.ListAutoScaleVmProfiles(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListAutoScaleVmProfiles", testlistAutoScaleVmProfiles)

	testlistConditions := func(t *testing.T) {
		if _, ok := response["listConditions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewListConditionsParams()
		_, err := client.AutoScale.ListConditions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListConditions", testlistConditions)

	testlistCounters := func(t *testing.T) {
		if _, ok := response["listCounters"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewListCountersParams()
		_, err := client.AutoScale.ListCounters(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCounters", testlistCounters)

	testupdateAutoScalePolicy := func(t *testing.T) {
		if _, ok := response["updateAutoScalePolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewUpdateAutoScalePolicyParams("id")
		r, err := client.AutoScale.UpdateAutoScalePolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateAutoScalePolicy", testupdateAutoScalePolicy)

	testupdateAutoScaleVmGroup := func(t *testing.T) {
		if _, ok := response["updateAutoScaleVmGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewUpdateAutoScaleVmGroupParams("id")
		r, err := client.AutoScale.UpdateAutoScaleVmGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateAutoScaleVmGroup", testupdateAutoScaleVmGroup)

	testupdateAutoScaleVmProfile := func(t *testing.T) {
		if _, ok := response["updateAutoScaleVmProfile"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AutoScale.NewUpdateAutoScaleVmProfileParams("id")
		r, err := client.AutoScale.UpdateAutoScaleVmProfile(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateAutoScaleVmProfile", testupdateAutoScaleVmProfile)

}
