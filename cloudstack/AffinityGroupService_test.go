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

func TestCreateAffinityGroup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"createAffinityGroup": `{
				"createaffinitygroupresponse": {
					"id": "5059d7e3-9213-448d-9314-6ae4e1d95a33",
					"jobid": "58e969e8-768b-44d0-b278-fd1b2f236c00"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.user.affinitygroup.CreateAffinityGroupCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"affinitygroup": {
							"id": "5059d7e3-9213-448d-9314-6ae4e1d95a33",
							"name": "testAffinityGroup",
							"description": "testAffinityGroup",
							"account": "admin",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"domain": "ROOT",
							"type": "host affinity"
						}
					},
					"jobinstancetype": "AffinityGroup",
					"created": "2021-10-01T16:59:04+0000",
					"completed": "2021-10-01T16:59:04+0000",
					"jobid": "58e969e8-768b-44d0-b278-fd1b2f236c00"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	p := client.AffinityGroup.NewCreateAffinityGroupParams("testAffinityGroup", "host affinity")
	ag, err := client.AffinityGroup.CreateAffinityGroup(p)
	if err != nil {
		t.Errorf("Failed to disassociate IP addres due to: %v", err.Error())
		return
	}
	if ag.Name != "testAffinityGroup" {
		t.Errorf("Failed to create affinity group of name: testAffinityGroup")
	}
}
