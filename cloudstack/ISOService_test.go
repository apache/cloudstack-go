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

func TestISOService_RegisterIso(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"registerisoresponse": {
				"count": 1,
				"iso": [
					{
						"id": "e6d86706-aa8c-444c-bf5c-ba677f8e02c2",
						"name": "testIso",
						"displaytext": "testIso",
						"ispublic": false,
						"created": "2021-10-03T07:28:41+0000",
						"isready": true,
						"passwordenabled": false,
						"bootable": true,
						"isfeatured": false,
						"crossZones": false,
						"ostypeid": "e510f742-5fdf-11ea-9a56-1e006800018c",
						"ostypename": "Other (64-bit)",
						"account": "admin",
						"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
						"zonename": "shouldwork",
						"status": "Successfully Installed",
						"size": 5242880,
						"domain": "ROOT",
						"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
						"isextractable": false,
						"bits": 64,
						"isdynamicallyscalable": false,
						"directdownload": false,
						"url": "http://dl.openvm.eu/cloudstack/iso/TinyCore-8.0.iso",
						"tags": []
					}
				]
			}
		}`
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.ISO.NewRegisterIsoParams("testIso", "testIso",
		"http://dl.openvm.eu/cloudstack/iso/TinyCore-8.0.iso", "1d8d87d4-1425-459c-8d81-c6f57dca2bd2")
	resp, err := client.ISO.RegisterIso(params)
	if err != nil {
		t.Errorf("Failed to register ISO due to: %v", err)
		return
	}

	if resp == nil {
		t.Errorf("Failed to register template")
	}
}
