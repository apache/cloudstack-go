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

func TestResourcetagsService(t *testing.T) {
	service := "ResourcetagsService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateTags := func(t *testing.T) {
		if _, ok := response["createTags"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Resourcetags.NewCreateTagsParams([]string{}, "resourcetype", map[string]string{})
		_, err := client.Resourcetags.CreateTags(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTags", testcreateTags)

	testdeleteTags := func(t *testing.T) {
		if _, ok := response["deleteTags"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Resourcetags.NewDeleteTagsParams([]string{}, "resourcetype")
		_, err := client.Resourcetags.DeleteTags(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTags", testdeleteTags)

	testlistStorageTags := func(t *testing.T) {
		if _, ok := response["listStorageTags"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Resourcetags.NewListStorageTagsParams()
		_, err := client.Resourcetags.ListStorageTags(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListStorageTags", testlistStorageTags)

	testlistTags := func(t *testing.T) {
		if _, ok := response["listTags"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Resourcetags.NewListTagsParams()
		_, err := client.Resourcetags.ListTags(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTags", testlistTags)

}
