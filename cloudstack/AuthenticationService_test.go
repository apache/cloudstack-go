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

func TestAuthenticationService_Login(t *testing.T) {
	apiName := "login"
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		loginResponse, err := ReadData(apiName, "AuthenticationService")
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintf(writer, loginResponse[apiName])
	}))
	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	loginParams := client.Authentication.NewLoginParams("admin", "password")
	resp, err := client.Authentication.Login(loginParams)

	if err != nil {
		t.Errorf("Failed to login user 'admin', due to: %v", err)
	}

	if resp == nil {
		t.Errorf("Failed to login user 'admin'")
	}

	if resp.Username != "admin" {
		t.Errorf("Failed to login user 'admin'")
	}
}

func TestAuthenticationService_Logout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "logout"
		logoutResponse, err := ReadData(apiName, "AuthenticationService")
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, logoutResponse[apiName])
	}))

	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	logoutParams := client.Authentication.NewLogoutParams()
	logoutResp, err := client.Authentication.Logout(logoutParams)

	if err != nil {
		t.Errorf("Failed to logout user 'admin', due to: %v", err)
	}

	if logoutResp.Description != "success" {
		t.Errorf("Failed to logout user 'admin'")
	}
}
