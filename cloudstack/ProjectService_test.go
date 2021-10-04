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

package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProjectService_CreateProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"createProject": `{
				"createprojectresponse": {
					"id": "69646881-8d7f-4800-987d-106698a42608",
					"jobid": "0666daab-af9b-4001-ae70-78ac6bc697e8"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.user.project.CreateProjectCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"project": {
							"id": "69646881-8d7f-4800-987d-106698a42608",
							"name": "testProject",
							"displaytext": "testProject",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"owner": [
								{
									"account": "admin"
								}
							],
							"projectaccountname": "PrjAcct-testProject-1",
							"state": "Active",
							"tags": [],
							"networklimit": "20",
							"networktotal": 0,
							"networkavailable": "20",
							"vpclimit": "20",
							"vpctotal": 0,
							"vpcavailable": "20",
							"cpulimit": "40",
							"cputotal": 0,
							"cpuavailable": "40",
							"memorylimit": "40960",
							"memorytotal": 0,
							"memoryavailable": "40960",
							"primarystoragelimit": "200",
							"primarystoragetotal": 0,
							"primarystorageavailable": "200",
							"secondarystoragelimit": "400",
							"secondarystoragetotal": 0,
							"secondarystorageavailable": "400.0",
							"vmlimit": "20",
							"vmtotal": 0,
							"vmavailable": "20",
							"iplimit": "20",
							"iptotal": 0,
							"ipavailable": "20",
							"volumelimit": "20",
							"volumetotal": 0,
							"volumeavailable": "20",
							"snapshotlimit": "20",
							"snapshottotal": 0,
							"snapshotavailable": "20",
							"templatelimit": "20",
							"templatetotal": 0,
							"templateavailable": "20",
							"vmstopped": 0,
							"vmrunning": 0
						}
					},
					"jobinstancetype": "None",
					"created": "2021-10-04T06:21:54+0000",
					"completed": "2021-10-04T06:21:54+0000",
					"jobid": "0666daab-af9b-4001-ae70-78ac6bc697e8"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
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
		responses := map[string]string{
			"activateProject": `{
				"activaterojectresponse": {
					"jobid": "fa12a90a-29fc-43c0-8cbd-5be49b72ed22"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.user.project.ActivateProjectCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"project": {
							"id": "99a842a4-e50f-4265-8ca7-249959506c13",
							"name": "Admin Project",
							"displaytext": "Admin Project",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"owner": [
								{
									"account": "admin"
								}
							],
							"projectaccountname": "PrjAcct-Admin Project-1",
							"state": "Active",
							"tags": [],
							"networklimit": "20",
							"networktotal": 0,
							"networkavailable": "20",
							"vpclimit": "20",
							"vpctotal": 0,
							"vpcavailable": "20",
							"cpulimit": "40",
							"cputotal": 0,
							"cpuavailable": "40",
							"memorylimit": "40960",
							"memorytotal": 0,
							"memoryavailable": "40960",
							"primarystoragelimit": "200",
							"primarystoragetotal": 0,
							"primarystorageavailable": "200",
							"secondarystoragelimit": "400",
							"secondarystoragetotal": 0,
							"secondarystorageavailable": "400.0",
							"vmlimit": "20",
							"vmtotal": 0,
							"vmavailable": "20",
							"iplimit": "20",
							"iptotal": 0,
							"ipavailable": "20",
							"volumelimit": "20",
							"volumetotal": 0,
							"volumeavailable": "20",
							"snapshotlimit": "20",
							"snapshottotal": 0,
							"snapshotavailable": "20",
							"templatelimit": "20",
							"templatetotal": 0,
							"templateavailable": "20",
							"vmstopped": 0,
							"vmrunning": 0
						}
					},
					"created": "2021-10-04T06:30:39+0000",
					"completed": "2021-10-04T06:30:39+0000",
					"jobid": "fa12a90a-29fc-43c0-8cbd-5be49b72ed22"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
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
		responses := map[string]string{
			"suspendProject": `{
				"suspendprojectresponse": {
					"jobid": "608f4d53-ceae-4747-9c0d-e8c15fc52135"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.user.project.SuspendProjectCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"project": {
							"id": "99a842a4-e50f-4265-8ca7-249959506c13",
							"name": "Admin Project",
							"displaytext": "Admin Project",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"owner": [
								{
									"account": "admin"
								}
							],
							"projectaccountname": "PrjAcct-Admin Project-1",
							"state": "Suspended",
							"tags": [],
							"networklimit": "20",
							"networktotal": 0,
							"networkavailable": "20",
							"vpclimit": "20",
							"vpctotal": 0,
							"vpcavailable": "20",
							"cpulimit": "40",
							"cputotal": 0,
							"cpuavailable": "40",
							"memorylimit": "40960",
							"memorytotal": 0,
							"memoryavailable": "40960",
							"primarystoragelimit": "200",
							"primarystoragetotal": 0,
							"primarystorageavailable": "200",
							"secondarystoragelimit": "400",
							"secondarystoragetotal": 0,
							"secondarystorageavailable": "400.0",
							"vmlimit": "20",
							"vmtotal": 0,
							"vmavailable": "20",
							"iplimit": "20",
							"iptotal": 0,
							"ipavailable": "20",
							"volumelimit": "20",
							"volumetotal": 0,
							"volumeavailable": "20",
							"snapshotlimit": "20",
							"snapshottotal": 0,
							"snapshotavailable": "20",
							"templatelimit": "20",
							"templatetotal": 0,
							"templateavailable": "20",
							"vmstopped": 0,
							"vmrunning": 0
						}
					},
					"created": "2021-10-04T06:34:01+0000",
					"completed": "2021-10-04T06:34:01+0000",
					"jobid": "608f4d53-ceae-4747-9c0d-e8c15fc52135"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
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
		responses := map[string]string{
			"updateProject": `{
				"updateprojectresponse": {
					"jobid": "1330d540-c099-4943-ac22-3fd5846c0e5b"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.user.project.UpdateProjectCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"project": {
							"id": "69646881-8d7f-4800-987d-106698a42608",
							"name": "testProject",
							"displaytext": "testProjectUpdate",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"owner": [
								{
									"account": "admin"
								}
							],
							"projectaccountname": "PrjAcct-testProject-1",
							"state": "Active",
							"tags": [],
							"networklimit": "20",
							"networktotal": 0,
							"networkavailable": "20",
							"vpclimit": "20",
							"vpctotal": 0,
							"vpcavailable": "20",
							"cpulimit": "40",
							"cputotal": 0,
							"cpuavailable": "40",
							"memorylimit": "40960",
							"memorytotal": 0,
							"memoryavailable": "40960",
							"primarystoragelimit": "200",
							"primarystoragetotal": 0,
							"primarystorageavailable": "200",
							"secondarystoragelimit": "400",
							"secondarystoragetotal": 0,
							"secondarystorageavailable": "400.0",
							"vmlimit": "20",
							"vmtotal": 0,
							"vmavailable": "20",
							"iplimit": "20",
							"iptotal": 0,
							"ipavailable": "20",
							"volumelimit": "20",
							"volumetotal": 0,
							"volumeavailable": "20",
							"snapshotlimit": "20",
							"snapshottotal": 0,
							"snapshotavailable": "20",
							"templatelimit": "20",
							"templatetotal": 0,
							"templateavailable": "20",
							"vmstopped": 0,
							"vmrunning": 0
						}
					},
					"created": "2021-10-04T06:35:29+0000",
					"completed": "2021-10-04T06:35:29+0000",
					"jobid": "1330d540-c099-4943-ac22-3fd5846c0e5b"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
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
