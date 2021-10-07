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

func TestResourcetagsService_CreateTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"createTags": `{
				"createtagsresponse": {
					"jobid": "fb96bb6b-c192-492d-9671-42cea59a3709"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
				  "accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
				  "cmd": "org.apache.cloudstack.api.command.user.tag.CreateTagsCmd",
				  "completed": "2021-10-04T06:58:35+0000",
				  "created": "2021-10-04T06:58:35+0000",
				  "jobid": "4901a56e-60aa-4385-8ab6-8f9591ef0469",
				  "jobprocstatus": 0,
				  "jobresult": {
					"success": true
				  },
				  "jobresultcode": 0,
				  "jobresulttype": "object",
				  "jobstatus": 1,
				  "userid": "27f2484f-5fe0-11ea-9a56-1e006800018c"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	tags := make(map[string]string)
	tags["key"] = "testKey"
	tags["value"] = "testValue"
	params := client.Resourcetags.NewCreateTagsParams([]string{"99a842a4-e50f-4265-8ca7-249959506c13"}, "Project", tags)
	resp, err := client.Resourcetags.CreateTags(params)
	if err != nil {
		t.Errorf("Failed to create tags for project due to: %v", err)
		return
	}

	if resp == nil || !resp.Success {
		t.Errorf("Failed to create project tags")
	}
}

func TestResourcetagsService_DeleteTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"deleteTags": `{
				"deletetagsresponse": {
					"jobid": "494d0f35-057a-4cc7-bf5d-b52d3b803a36"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
				  "accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
				  "cmd": "org.apache.cloudstack.api.command.user.tag.DeleteTagsCmd",
				  "completed": "2021-10-04T07:07:22+0000",
				  "created": "2021-10-04T07:07:22+0000",
				  "jobid": "494d0f35-057a-4cc7-bf5d-b52d3b803a36",
				  "jobprocstatus": 0,
				  "jobresult": {
					"success": true
				  },
				  "jobresultcode": 0,
				  "jobresulttype": "object",
				  "jobstatus": 1,
				  "userid": "27f2484f-5fe0-11ea-9a56-1e006800018c"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Resourcetags.NewDeleteTagsParams([]string{"99a842a4-e50f-4265-8ca7-249959506c13"}, "Project")
	resp, err := client.Resourcetags.DeleteTags(params)
	if err != nil {
		t.Errorf("Failed to delete tags for project due to: %v", err)
		return
	}

	if resp == nil || !resp.Success {
		t.Errorf("Failed to delete project tags")
	}
}
