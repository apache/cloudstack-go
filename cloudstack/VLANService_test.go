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

func TestNetworkService_DedicateGuestVLANRange(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "dedicateGuestVlanRange"
		response, err := ReadData(apiName, "VLANService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.VLAN.NewDedicateGuestVlanRangeParams("8e27a637-7525-49ed-81ce-52bd5e5d9ea2", "100-110")
	params.SetAccount("admin")
	params.SetDomainid("e4874e10-5fdf-11ea-9a56-1e006800018c")
	resp, err := client.VLAN.DedicateGuestVlanRange(params)
	if err != nil {
		t.Errorf("Failed to dedicate guest VLAN range for physical network due to: %v", err)
		return
	}

	if resp == nil || resp.Guestvlanrange != "100-110" {
		t.Errorf("Failed to dedicate guest VLAN range for physical network")
	}
}
