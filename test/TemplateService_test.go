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

func TestTemplateService_RegisterTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "registerTemplate"
		response, err := ReadData(apiName, "TemplateService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintf(writer, response[apiName])
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Template.NewRegisterTemplateParams("testTemplate", "VHD",
		"Simulator", "testTemplate", "http://dl.openvm.eu/cloudstack/macchinina/x86_64/macchinina-xen.vhd.bz2")
	resp, err := client.Template.RegisterTemplate(params)
	if err != nil {
		t.Errorf("Failed to register template due to %v", err)
		return
	}

	if resp == nil || resp.Count != 1 {
		t.Errorf("Failed to register template")
	}
}

func TestTemplateService_CreateTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "createTemplate"
		response, err := ParseAsyncResponse(apiName, "TemplateService", *r)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Template.NewCreateTemplateParams("createTempFromVol", "createTempFromVol", "e510f742-5fdf-11ea-9a56-1e006800018c")
	resp, err := client.Template.CreateTemplate(params)
	if err != nil {
		t.Errorf("Failed to create template from volume due to %v", err)
		return
	}
	if resp == nil || resp.Name != "createTempFromVol" {
		t.Errorf("Failed to create template from volume")
	}
}

func TestTemplateService_ExtractTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		apiName := "extractTemplate"
		response, err := ParseAsyncResponse(apiName, "TemplateService", *r)
		if err != nil {
			t.Errorf("Failed to parse response, due to: %v", err)
		}
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := cloudstack.NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Template.NewExtractTemplateParams("5ce8f0b1-a910-4631-a8de-1e332bf3a6b7", "HTTP_DOWNLOAD")
	resp, err := client.Template.ExtractTemplate(params)
	if err != nil {
		t.Errorf("Failed to download template due to %v", err)
		return
	}
	if resp == nil || resp.Url == "" {
		t.Errorf("Failed to download template")
	}
}
