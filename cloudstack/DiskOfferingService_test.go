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

func TestDiskOfferingService_CreateDiskOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"creatediskofferingresponse": {
				"diskoffering": {
					"id": "7662b6ae-f00b-4268-973f-f3f87eaf82c5",
						"name": "test",
						"displaytext": "test",
						"disksize": 0,
						"created": "2021-10-02T15:19:25+0000",
						"iscustomized": true,
						"storagetype": "shared",
						"provisioningtype": "thin",
						"cacheMode": "none",
						"displayoffering": true
				}
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.DiskOffering.NewCreateDiskOfferingParams("test", "test")
	resp, err := client.DiskOffering.CreateDiskOffering(params)
	if err != nil {
		t.Errorf("Failed to create disk offering due to: %v", err)
	}

	if resp == nil {
		t.Errorf("Failed to create disk offering")
	}
}

func TestDiskOfferingService_ListDiskOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listdiskofferingsresponse": {
				"count": 1,
				"diskoffering": [
					{
						"id": "7662b6ae-f00b-4268-973f-f3f87eaf82c5",
						"name": "test",
						"displaytext": "test",
						"disksize": 0,
						"created": "2021-10-02T15:19:25+0000",
						"iscustomized": true,
						"storagetype": "shared",
						"provisioningtype": "thin",
						"cacheMode": "none",
						"displayoffering": true
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.DiskOffering.NewListDiskOfferingsParams()
	params.SetId("7662b6ae-f00b-4268-973f-f3f87eaf82c5")
	resp, err := client.DiskOffering.ListDiskOfferings(params)
	if err != nil {
		t.Errorf("Failed to list disk offering due to: %v", err)
	}

	if resp.Count != 1 {
		t.Errorf("Failed to list disk offering")
	}
}

func TestDiskOfferingService_DeleteDiskOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"deletediskofferingresponse": {
				"success": true
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.DiskOffering.NewDeleteDiskOfferingParams("7662b6ae-f00b-4268-973f-f3f87eaf82c5")
	resp, err := client.DiskOffering.DeleteDiskOffering(params)
	if err != nil {
		t.Errorf("Failed to delete disk offering due to : %v", err)
	}
	if !resp.Success {
		t.Errorf("Failed to delete disk offering")
	}
}
