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

func TestAssociateIpAddress(t *testing.T) {
	apiName := "associateIpAddress"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ParseAsyncResponse(apiName, "AddressService", *r)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(w, response)
	}))

	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	p := client.Address.NewAssociateIpAddressParams()
	p.SetZoneid("2e86e486-b472-4f12-a9b2-bb73701241e0")
	ip, err := client.Address.AssociateIpAddress(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if ip.Ipaddress != "10.70.3.100" {
		t.Errorf("ipaddress: actual %s, expected 10.70.3.100", ip.Ipaddress)
	}
}

func TestDisassociateIpAddress(t *testing.T) {
	apiName := "disassociateIpAddress"
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response, err := ParseAsyncResponse(apiName, "AddressService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))

	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	p := client.Address.NewDisassociateIpAddressParams("a767fbe1-ed7a-4d7c-8221-c7d736ca622d")
	address, err := client.Address.DisassociateIpAddress(p)
	if err != nil {
		t.Errorf("Failed to disassociate IP addres due to: %v", err.Error())
		return
	}
	if address.Success != true {
		t.Errorf("Failed to disassociate IP address")
	}
}
