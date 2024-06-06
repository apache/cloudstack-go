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

func TestUserService(t *testing.T) {
	service := "UserService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateUser := func(t *testing.T) {
		if _, ok := response["createUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewCreateUserParams("account", "email", "firstname", "lastname", "password", "username")
		r, err := client.User.CreateUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateUser", testcreateUser)

	testdeleteUser := func(t *testing.T) {
		if _, ok := response["deleteUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewDeleteUserParams("id")
		_, err := client.User.DeleteUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteUser", testdeleteUser)

	testdisableUser := func(t *testing.T) {
		if _, ok := response["disableUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewDisableUserParams("id")
		r, err := client.User.DisableUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DisableUser", testdisableUser)

	testenableUser := func(t *testing.T) {
		if _, ok := response["enableUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewEnableUserParams("id")
		r, err := client.User.EnableUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("EnableUser", testenableUser)

	testgetUser := func(t *testing.T) {
		if _, ok := response["getUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewGetUserParams("userapikey")
		r, err := client.User.GetUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("GetUser", testgetUser)

	testgetUserKeys := func(t *testing.T) {
		if _, ok := response["getUserKeys"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewGetUserKeysParams("id")
		_, err := client.User.GetUserKeys(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetUserKeys", testgetUserKeys)

	testgetVirtualMachineUserData := func(t *testing.T) {
		if _, ok := response["getVirtualMachineUserData"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewGetVirtualMachineUserDataParams("virtualmachineid")
		_, err := client.User.GetVirtualMachineUserData(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetVirtualMachineUserData", testgetVirtualMachineUserData)

	testlistUsers := func(t *testing.T) {
		if _, ok := response["listUsers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewListUsersParams()
		_, err := client.User.ListUsers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListUsers", testlistUsers)

	testlockUser := func(t *testing.T) {
		if _, ok := response["lockUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewLockUserParams("id")
		r, err := client.User.LockUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("LockUser", testlockUser)

	testregisterUserKeys := func(t *testing.T) {
		if _, ok := response["registerUserKeys"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewRegisterUserKeysParams("id")
		_, err := client.User.RegisterUserKeys(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RegisterUserKeys", testregisterUserKeys)

	testupdateUser := func(t *testing.T) {
		if _, ok := response["updateUser"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewUpdateUserParams("id")
		r, err := client.User.UpdateUser(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateUser", testupdateUser)

	testlistUserData := func(t *testing.T) {
		if _, ok := response["listUserData"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewListUserDataParams()
		_, err := client.User.ListUserData(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListUserData", testlistUserData)

	testdeleteUserData := func(t *testing.T) {
		if _, ok := response["deleteUserData"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewDeleteUserDataParams("id")
		_, err := client.User.DeleteUserData(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteUserData", testdeleteUserData)

	testregisterUserData := func(t *testing.T) {
		if _, ok := response["registerUserData"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.User.NewRegisterUserDataParams("name", "userdata")
		_, err := client.User.RegisterUserData(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RegisterUserData", testregisterUserData)

}
