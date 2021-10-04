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

func TestHypervisorService_ListSpecificHypervisorCapabilities(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listhypervisorcapabilitiesresponse": {
				"count": 1,
				"hypervisorCapabilities": [
					{
						"id": "1",
						"hypervisorversion": "default",
						"hypervisor": "XenServer",
						"maxguestslimit": 50,
						"securitygroupenabled": true,
						"maxdatavolumeslimit": 6,
						"storagemotionenabled": false
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Hypervisor.NewListHypervisorCapabilitiesParams()
	params.SetId("1")
	resp, err := client.Hypervisor.ListHypervisorCapabilities(params)
	if err != nil {
		t.Errorf("Failed to list hypervisor capabilities due to: %v", err)
		return
	}

	if resp.Count != 1 {
		t.Errorf("Failed list specific hypervisor capability")
	}
}
