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

func TestAffinityGroupService(t *testing.T) {
	service := "AffinityGroupService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateAffinityGroup := func(t *testing.T) {
		if _, ok := response["createAffinityGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AffinityGroup.NewCreateAffinityGroupParams("name", "type")
		r, err := client.AffinityGroup.CreateAffinityGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateAffinityGroup", testcreateAffinityGroup)

	testdeleteAffinityGroup := func(t *testing.T) {
		if _, ok := response["deleteAffinityGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AffinityGroup.NewDeleteAffinityGroupParams()
		_, err := client.AffinityGroup.DeleteAffinityGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteAffinityGroup", testdeleteAffinityGroup)

	testlistAffinityGroupTypes := func(t *testing.T) {
		if _, ok := response["listAffinityGroupTypes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AffinityGroup.NewListAffinityGroupTypesParams()
		_, err := client.AffinityGroup.ListAffinityGroupTypes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListAffinityGroupTypes", testlistAffinityGroupTypes)

	testlistAffinityGroups := func(t *testing.T) {
		if _, ok := response["listAffinityGroups"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AffinityGroup.NewListAffinityGroupsParams()
		_, err := client.AffinityGroup.ListAffinityGroups(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListAffinityGroups", testlistAffinityGroups)

	testupdateVMAffinityGroup := func(t *testing.T) {
		if _, ok := response["updateVMAffinityGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.AffinityGroup.NewUpdateVMAffinityGroupParams("id")
		r, err := client.AffinityGroup.UpdateVMAffinityGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVMAffinityGroup", testupdateVMAffinityGroup)

}
