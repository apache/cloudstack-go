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

func TestDomainService_CreateDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "createDomain"
		response, err := ReadData(apiName, "DomainService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewCreateDomainParams("testDomain")
	params.SetParentdomainid("e4874e10-5fdf-11ea-9a56-1e006800018c")
	resp, err := client.Domain.CreateDomain(params)
	if err != nil {
		t.Errorf("Failed to create domain due to: %v", err)
	}

	if resp.Name != "testDomain" {
		t.Errorf("Failed to create domain")
	}
}

func TestDomainService_UpdateDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "updateDomain"
		response, err := ReadData(apiName, "DomainService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewUpdateDomainParams("ee05fd92-7365-4421-a15b-abfa11dfc4f6")
	params.SetName("testDomainUpdated")
	resp, err := client.Domain.UpdateDomain(params)
	if err != nil {
		t.Errorf("Failed to update domain name due to: %v", err)
	}

	if resp.Name != "testDomainUpdated" {
		t.Errorf("Failed to update domain name")
	}
}

func TestDomainService_ListDomains(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "listDomains"
		response, err := ReadData(apiName, "DomainService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewListDomainsParams()
	params.SetId("097d3992-7a67-42e1-afb5-b4d2d81e280f")
	resp, err := client.Domain.ListDomains(params)
	if err != nil {
		t.Errorf("Failed to list specific domain details due to: %v", err)
	}
	if resp.Count != 1 {
		t.Errorf("Failed to list specific domain details")
	}
	if resp.Domains[0].Name != "DummyDomain" {
		t.Errorf("Failed to fetch details of specific domain")
	}
}

func TestDomainService_ListDomainChildren(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "listDomainChildren"
		response, err := ReadData(apiName, "DomainService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewListDomainChildrenParams()
	params.SetId("99becf06-7f0f-4eb4-bdc3-44fecb8cb829")
	resp, err := client.Domain.ListDomainChildren(params)
	if err != nil {
		t.Errorf("Failed to list specific domain's children due to: %v", err)
	}
	if resp.Count != 1 {
		t.Errorf("Failed to list specific domain's children ")
	}
}

func TestDomainService_DeleteDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "deleteDomain"
		response, err := ParseAsyncResponse(apiName, "DomainService", *request)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewDeleteDomainParams("ee05fd92-7365-4421-a15b-abfa11dfc4f6")
	resp, err := client.Domain.DeleteDomain(params)
	if err != nil {
		t.Errorf("Failed to delete domain due to: %v", err)
	}
	if !resp.Success {
		t.Errorf("Failed to delete domain")
	}
}
