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

func TestSSHService(t *testing.T) {
	service := "SSHService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateSSHKeyPair := func(t *testing.T) {
		if _, ok := response["createSSHKeyPair"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SSH.NewCreateSSHKeyPairParams("name")
		r, err := client.SSH.CreateSSHKeyPair(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateSSHKeyPair", testcreateSSHKeyPair)

	testdeleteSSHKeyPair := func(t *testing.T) {
		if _, ok := response["deleteSSHKeyPair"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SSH.NewDeleteSSHKeyPairParams("name")
		_, err := client.SSH.DeleteSSHKeyPair(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteSSHKeyPair", testdeleteSSHKeyPair)

	testlistSSHKeyPairs := func(t *testing.T) {
		if _, ok := response["listSSHKeyPairs"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SSH.NewListSSHKeyPairsParams()
		_, err := client.SSH.ListSSHKeyPairs(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSSHKeyPairs", testlistSSHKeyPairs)

	testregisterSSHKeyPair := func(t *testing.T) {
		if _, ok := response["registerSSHKeyPair"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SSH.NewRegisterSSHKeyPairParams("name", "publickey")
		r, err := client.SSH.RegisterSSHKeyPair(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RegisterSSHKeyPair", testregisterSSHKeyPair)

	testresetSSHKeyForVirtualMachine := func(t *testing.T) {
		if _, ok := response["resetSSHKeyForVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SSH.NewResetSSHKeyForVirtualMachineParams("id")
		r, err := client.SSH.ResetSSHKeyForVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ResetSSHKeyForVirtualMachine", testresetSSHKeyForVirtualMachine)

}
