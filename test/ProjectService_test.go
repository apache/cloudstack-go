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

func TestProjectService_CreateProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createProject"
		response, err := ParseAsyncResponse(apiName, "ProjectService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Project.NewCreateProjectParams("testProject", "testProject")
	resp, err := client.Project.CreateProject(params)
	if err != nil {
		t.Errorf("Failed to create Project due to: %v", err)
		return
	}
	if resp == nil || resp.Name != "testProject" {
		t.Errorf("Failed to create project")
	}
}

func TestProjectService_ActivateProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "activateProject"
		response, err := ParseAsyncResponse(apiName, "ProjectService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Project.NewActivateProjectParams("99a842a4-e50f-4265-8ca7-249959506c13")
	resp, err := client.Project.ActivateProject(params)
	if err != nil {
		t.Errorf("Failed to activate Project due to: %v", err)
		return
	}
	if resp == nil || resp.Id != "99a842a4-e50f-4265-8ca7-249959506c13" || resp.State != "Active" {
		t.Errorf("Failed to activate project")
	}
}

func TestProjectService_SuspendProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "suspendProject"
		response, err := ParseAsyncResponse(apiName, "ProjectService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Project.NewSuspendProjectParams("99a842a4-e50f-4265-8ca7-249959506c13")
	resp, err := client.Project.SuspendProject(params)
	if err != nil {
		t.Errorf("Failed to suspend Project due to: %v", err)
		return
	}
	if resp == nil || resp.Id != "99a842a4-e50f-4265-8ca7-249959506c13" || resp.State != "Suspended" {
		t.Errorf("Failed to suspend project")
	}
}

func TestProjectService_UpdateProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "updateProject"
		response, err := ParseAsyncResponse(apiName, "ProjectService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Project.NewUpdateProjectParams("99a842a4-e50f-4265-8ca7-249959506c13")
	params.SetDisplaytext("testProjectUpdated")
	resp, err := client.Project.UpdateProject(params)
	if err != nil {
		t.Errorf("Failed to update Project details due to: %v", err)
		return
	}

	if resp == nil || resp.Id != "69646881-8d7f-4800-987d-106698a42608" || resp.Displaytext != "testProjectUpdate" {
		t.Errorf("Failed to update project name")
	}
}

func TestProjectService_listProjectRolePermissions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "listProjectRolePermissions"
		response, err := ReadData(apiName, "ProjectService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Project.NewListProjectRolePermissionsParams("69646881-8d7f-4800-987d-106698a42608")
	params.SetProjectroleid("fa089002-d055-46b5-aa0a-18f2eae2b4fc")
	resp, err := client.Project.ListProjectRolePermissions(params)
	if err != nil {
		t.Errorf("Failed to list project role permissions due to %v", err)
		return
	}

	fmt.Println(resp)
	if resp == nil || resp.Count != 1 {
		t.Errorf("Failed to list VM")
	}
}
