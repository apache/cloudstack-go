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

func TestCreateAffinityGroup(t *testing.T) {
	apiName := "createAffinityGroup"
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses, err := ParseAsyncResponse(apiName, "AffinityGroupService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, responses)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
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
