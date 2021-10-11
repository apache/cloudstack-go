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

func TestImageStoreService_AddImageStore(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "addImageStore"
		response, err := ReadData(apiName, "ImageStoreService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.ImageStore.NewAddImageStoreParams("NFS")
	resp, err := client.ImageStore.AddImageStore(params)
	if err != nil {
		t.Errorf("Failed to add image store due to: %v", err)
		return
	}
	fmt.Println(resp)
	if resp == nil || resp.Providername != "NFS" {
		t.Errorf(" Failed to add image store")
	}
}

func TestImageStoreService_ListImageStores(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "listImageStores"
		response, err := ReadData(apiName, "ImageStoreService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.ImageStore.NewListImageStoresParams()
	resp, err := client.ImageStore.ListImageStores(params)
	if err != nil {
		t.Errorf("Failed to list image stores due to: %v", err)
		return
	}

	if resp == nil || resp.Count != 1 {
		t.Errorf(" Failed to list image stores")
	}
}

func TestImageStoreService_DeleteImageStore(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "deleteImageStore"
		response, err := ReadData(apiName, "ImageStoreService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.ImageStore.NewDeleteImageStoreParams("0ac85364-e31a-4840-97a4-a237b4291dfa")
	resp, err := client.ImageStore.DeleteImageStore(params)
	if err != nil {
		t.Errorf("Failed to delete image store due to: %v", err)
		return
	}

	if resp == nil || !resp.Success {
		t.Errorf(" Failed to delete image store")
	}
}
