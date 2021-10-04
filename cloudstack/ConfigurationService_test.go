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

func TestConfigurationService_ListCapabilities(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `{
			"listcapabilitiesresponse": {
			"capability": {
			"securitygroupsenabled": true,
			"dynamicrolesenabled": true,
			"cloudstackversion": "4.15.2.0",
			"userpublictemplateenabled": false,
			"supportELB": "false",
			"projectinviterequired": true,
			"allowusercreateprojects": true,
			"customdiskofferingminsize": 1,
			"customdiskofferingmaxsize": 1024,
			"regionsecondaryenabled": false,
			"kvmsnapshotenabled": true,
			"allowuserviewdestroyedvm": true,
			"allowuserexpungerecovervm": true,
			"allowuserexpungerecovervolume": true,
			"allowuserviewalldomainaccounts": false,
			"kubernetesserviceenabled": true,
			"kubernetesclusterexperimentalfeaturesenabled": true,
			"defaultuipagesize": 20
			}
		}
	}`
		fmt.Fprintf(w, resp)
	}))
	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	p := client.Configuration.NewListCapabilitiesParams()
	resp, err := client.Configuration.ListCapabilities(p)
	if err != nil {
		t.Errorf("Failed to list capabilities due to: %v", err)
	}

	if resp == nil {
		t.Errorf("Failed to list capabilities")
	}
}

func TestConfigurationService_ListConfigurations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `{
			"listconfigurationsresponse": {
			  "configuration": [
				{
				  "category": "Project Defaults",
				  "description": "If regular user can create a project; true by default",
				  "isdynamic": false,
				  "name": "allow.user.create.projects",
				  "value": "true"
				}
			  ],
			  "count": 1
			}
		}`
		fmt.Fprintf(w, resp)
	}))
	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	p := client.Configuration.NewListConfigurationsParams()
	p.SetName("allow.user.create.projects")
	resp, err := client.Configuration.ListConfigurations(p)
	if err != nil {
		t.Errorf("Failed to list configuration details due to: %v", err)
	}
	if resp.Count != 1 {
		t.Errorf("Failed to list configuration details")
	}
}

//func TestConfigurationService_UpdateConfigurations(t *testing.T) {
//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		resp := `{
//			"updateconfigurationresponse": {
//				"configuration": {
//					"category": "Project Defaults",
//					"name": "allow.user.create.projects",
//					"value": "false",
//					"description": "If regular user can create a project; true by default",
//					"isdynamic": false
//				}
//			}
//		}`
//		fmt.Fprintf(w, resp)
//	}))
//	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
//	p := client.Configuration.NewUpdateConfigurationParams("allow.user.create.projects")
//	p.SetValue("false")
//	resp, err := client.Configuration.UpdateConfiguration(p)
//	if err != nil {
//		t.Errorf("Failed to update configuration due to: %v", err)
//	}
//	fmt.Println(resp.Value)
//	if resp.Value != "false" {
//		t.Errorf("Failed to update configuration")
//	}
//}
