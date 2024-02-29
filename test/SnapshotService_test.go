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

func TestSnapshotService(t *testing.T) {
	service := "SnapshotService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateSnapshot := func(t *testing.T) {
		if _, ok := response["createSnapshot"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewCreateSnapshotParams("volumeid")
		r, err := client.Snapshot.CreateSnapshot(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateSnapshot", testcreateSnapshot)

	testcreateSnapshotPolicy := func(t *testing.T) {
		if _, ok := response["createSnapshotPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewCreateSnapshotPolicyParams("intervaltype", 0, "schedule", "timezone", "volumeid")
		r, err := client.Snapshot.CreateSnapshotPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateSnapshotPolicy", testcreateSnapshotPolicy)

	testcreateVMSnapshot := func(t *testing.T) {
		if _, ok := response["createVMSnapshot"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewCreateVMSnapshotParams("virtualmachineid")
		r, err := client.Snapshot.CreateVMSnapshot(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVMSnapshot", testcreateVMSnapshot)

	testdeleteSnapshot := func(t *testing.T) {
		if _, ok := response["deleteSnapshot"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewDeleteSnapshotParams("id")
		_, err := client.Snapshot.DeleteSnapshot(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteSnapshot", testdeleteSnapshot)

	testdeleteSnapshotPolicies := func(t *testing.T) {
		if _, ok := response["deleteSnapshotPolicies"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewDeleteSnapshotPoliciesParams()
		_, err := client.Snapshot.DeleteSnapshotPolicies(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteSnapshotPolicies", testdeleteSnapshotPolicies)

	testdeleteVMSnapshot := func(t *testing.T) {
		if _, ok := response["deleteVMSnapshot"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewDeleteVMSnapshotParams("vmsnapshotid")
		_, err := client.Snapshot.DeleteVMSnapshot(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVMSnapshot", testdeleteVMSnapshot)

	testlistSnapshotPolicies := func(t *testing.T) {
		if _, ok := response["listSnapshotPolicies"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewListSnapshotPoliciesParams()
		_, err := client.Snapshot.ListSnapshotPolicies(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSnapshotPolicies", testlistSnapshotPolicies)

	testlistSnapshots := func(t *testing.T) {
		if _, ok := response["listSnapshots"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewListSnapshotsParams()
		_, err := client.Snapshot.ListSnapshots(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSnapshots", testlistSnapshots)

	testlistVMSnapshot := func(t *testing.T) {
		if _, ok := response["listVMSnapshot"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewListVMSnapshotParams()
		_, err := client.Snapshot.ListVMSnapshot(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVMSnapshot", testlistVMSnapshot)

	testrevertSnapshot := func(t *testing.T) {
		if _, ok := response["revertSnapshot"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewRevertSnapshotParams("id")
		r, err := client.Snapshot.RevertSnapshot(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RevertSnapshot", testrevertSnapshot)

	testrevertToVMSnapshot := func(t *testing.T) {
		if _, ok := response["revertToVMSnapshot"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewRevertToVMSnapshotParams("vmsnapshotid")
		r, err := client.Snapshot.RevertToVMSnapshot(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RevertToVMSnapshot", testrevertToVMSnapshot)

	testupdateSnapshotPolicy := func(t *testing.T) {
		if _, ok := response["updateSnapshotPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Snapshot.NewUpdateSnapshotPolicyParams()
		r, err := client.Snapshot.UpdateSnapshotPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateSnapshotPolicy", testupdateSnapshotPolicy)

}
