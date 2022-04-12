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

func TestResourcetagsService_CreateTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createTags"
		response, err := ParseAsyncResponse(apiName, "ResourceTagsService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
		apiName := "deleteTags"
		response, err := ParseAsyncResponse(apiName, "ResourceTagsService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
