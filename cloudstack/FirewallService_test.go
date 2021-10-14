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

func TestFirewallService_CreateFirewallRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createFirewallRule"
		response, err := ParseAsyncResponse(apiName, "FirewallService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Firewall.NewCreateFirewallRuleParams("192.168.2.4", "tcp")
	resp, err := client.Firewall.CreateFirewallRule(params)
	if err != nil {
		t.Errorf("Failed to create firewall rule due to: %v", err)
	}
	if resp.Ipaddress != "192.168.2.4" {
		t.Errorf("Failed to create firewall rule")
	}

	if resp.State != "Active" {
		t.Errorf("Failed to create firewall rule")
	}
}

func TestFirewallService_DeleteFirewallRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "deleteFirewallRule"
		response, err := ParseAsyncResponse(apiName, "FirewallService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Firewall.NewDeleteFirewallRuleParams("fb4ad2ee-02c8-433e-a769-6f18afddc750")
	resp, err := client.Firewall.DeleteFirewallRule(params)
	if err != nil {
		t.Errorf("Failed to delete firewall rule due to: %v", err)
	}
	if !resp.Success {
		t.Errorf("Failed to delete firewall rule")
	}
}

func TestFirewallService_CreateEgressFirewallRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createEgressFirewallRule"
		response, err := ParseAsyncResponse(apiName, "FirewallService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Firewall.NewCreateEgressFirewallRuleParams("c4a3303c-376d-4d56-b336-1bd91cb130b6", "tcp")
	resp, err := client.Firewall.CreateEgressFirewallRule(params)
	if err != nil {
		t.Errorf("Failed to create egress firewall rule due to: %v", err)
	}

	if resp.State != "Active" {
		t.Errorf("Failed to create egress firewall rule")
	}
}

func TestFirewallService_DeleteEgressFirewallRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "deleteEgressFirewallRule"
		response, err := ParseAsyncResponse(apiName, "FirewallService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Firewall.NewDeleteEgressFirewallRuleParams("fb4ad2ee-02c8-433e-a769-6f18afddc750")
	resp, err := client.Firewall.DeleteEgressFirewallRule(params)
	if err != nil {
		t.Errorf("Failed to delete egress firewall rule due to: %v", err)
	}
	if !resp.Success {
		t.Errorf("Failed to delete egress firewall rule")
	}
}
