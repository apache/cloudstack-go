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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestUserService_CreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createUser"
		response, err := ReadData(apiName, "UserService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewCreateUserParams("admin", "user.xyz.com",
		"firstname", "lastname", "password", "dummyUser")
	resp, err := client.User.CreateUser(params)
	if err != nil {
		t.Errorf("Failed to create user due to %v", err)
		return
	}
	if resp == nil || resp.Username != "dummyUser" {
		t.Errorf("Failed to create user")
	}
}

func TestUserService_EnableUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "enableUser"
		response, err := ReadData(apiName, "UserService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewEnableUserParams("cd2b5afa-db4d-4532-88ec-1356a273e534")
	resp, err := client.User.EnableUser(params)
	if err != nil {
		t.Errorf("Failed to enable user due to %v", err)
		return
	}

	if resp == nil || resp.State != "enabled" {
		t.Errorf("Failed to enable user")
	}
}

func TestUserService_DisableUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "disableUser"
		response, err := ParseAsyncResponse(apiName, "UserService", *request)
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewDisableUserParams("cd2b5afa-db4d-4532-88ec-1356a273e534")
	resp, err := client.User.DisableUser(params)
	if err != nil {
		t.Errorf("Failed to disable user due to %v", err)
		return
	}

	if resp == nil || resp.State != "disabled" {
		t.Errorf("Failed to disable user")
	}
}

func TestUserService_ListUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "listUsers"
		response, err := ReadData(apiName, "UserService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewListUsersParams()
	params.SetId("cd2b5afa-db4d-4532-88ec-1356a273e534")
	resp, err := client.User.ListUsers(params)
	if err != nil {
		t.Errorf("Failed to list user details due to %v", err)
		return
	}

	if resp == nil || resp.Count != 1 {
		t.Errorf("Failed to list user details")
	}
}

func TestUserService_LockUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "lockUser"
		response, err := ReadData(apiName, "UserService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewLockUserParams("3b7325b3-849c-4314-ad19-dd3c483b6d1a")
	resp, err := client.User.LockUser(params)
	if err != nil {
		t.Errorf("Failed to lock user due to %v", err)
		return
	}

	if resp == nil || resp.State != "locked" {
		t.Errorf("Failed to lock user")
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "deleteUser"
		response, err := ReadData(apiName, "UserService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewDeleteUserParams("cd2b5afa-db4d-4532-88ec-1356a273e534")
	resp, err := client.User.DeleteUser(params)
	if err != nil {
		t.Errorf("Failed to delete user due to %v", err)
		return
	}

	if resp == nil || !resp.Success {
		t.Errorf("Failed to delete user")
	}
}

func TestGetUserKeys(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "getUserKeys"
		response, err := ReadData(apiName, "UserService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewGetUserKeysParams("random-id")
	resp, err := client.User.GetUserKeys(params)
	if err != nil {
		t.Errorf("Failed to get user keys due to %v", err)
		return
	}
	if resp.Apikey == "" || resp.Secretkey == "" {
		t.Errorf("Parsing failure")
	}
}
