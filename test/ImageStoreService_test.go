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
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestImageStoreService(t *testing.T) {
	service := "ImageStoreService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddImageStore := func(t *testing.T) {
		if _, ok := response["addImageStore"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewAddImageStoreParams("provider")
		r, err := client.ImageStore.AddImageStore(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddImageStore", testaddImageStore)

	testaddImageStoreS3 := func(t *testing.T) {
		if _, ok := response["addImageStoreS3"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewAddImageStoreS3Params("accesskey", "bucket", "endpoint", "secretkey")
		r, err := client.ImageStore.AddImageStoreS3(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddImageStoreS3", testaddImageStoreS3)

	testcreateSecondaryStagingStore := func(t *testing.T) {
		if _, ok := response["createSecondaryStagingStore"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewCreateSecondaryStagingStoreParams("url")
		r, err := client.ImageStore.CreateSecondaryStagingStore(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateSecondaryStagingStore", testcreateSecondaryStagingStore)

	testdeleteImageStore := func(t *testing.T) {
		if _, ok := response["deleteImageStore"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewDeleteImageStoreParams("id")
		_, err := client.ImageStore.DeleteImageStore(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteImageStore", testdeleteImageStore)

	testdeleteSecondaryStagingStore := func(t *testing.T) {
		if _, ok := response["deleteSecondaryStagingStore"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewDeleteSecondaryStagingStoreParams("id")
		_, err := client.ImageStore.DeleteSecondaryStagingStore(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteSecondaryStagingStore", testdeleteSecondaryStagingStore)

	testlistImageStores := func(t *testing.T) {
		if _, ok := response["listImageStores"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewListImageStoresParams()
		_, err := client.ImageStore.ListImageStores(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListImageStores", testlistImageStores)

	testlistSecondaryStagingStores := func(t *testing.T) {
		if _, ok := response["listSecondaryStagingStores"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewListSecondaryStagingStoresParams()
		_, err := client.ImageStore.ListSecondaryStagingStores(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSecondaryStagingStores", testlistSecondaryStagingStores)

	testmigrateResourceToAnotherSecondaryStorage := func(t *testing.T) {
		if _, ok := response["migrateResourceToAnotherSecondaryStorage"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewMigrateResourceToAnotherSecondaryStorageParams("destpool", "srcpool")
		_, err := client.ImageStore.MigrateResourceToAnotherSecondaryStorage(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("MigrateResourceToAnotherSecondaryStorage", testmigrateResourceToAnotherSecondaryStorage)

	testupdateCloudToUseObjectStore := func(t *testing.T) {
		if _, ok := response["updateCloudToUseObjectStore"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewUpdateCloudToUseObjectStoreParams("provider")
		r, err := client.ImageStore.UpdateCloudToUseObjectStore(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateCloudToUseObjectStore", testupdateCloudToUseObjectStore)

	testlistImageStoreObjects := func(t *testing.T) {
		if _, ok := response["listImageStoreObjects"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewListImageStoreObjectsParams("id")
		_, err := client.ImageStore.ListImageStoreObjects(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListImageStoreObjects", testlistImageStoreObjects)

	testupdateImageStore := func(t *testing.T) {
		if _, ok := response["updateImageStore"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewUpdateImageStoreParams("id")
		r, err := client.ImageStore.UpdateImageStore(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateImageStore", testupdateImageStore)

	testdownloadImageStoreObject := func(t *testing.T) {
		if _, ok := response["downloadImageStoreObject"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ImageStore.NewDownloadImageStoreObjectParams("id")
		r, err := client.ImageStore.DownloadImageStoreObject(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DownloadImageStoreObject", testdownloadImageStoreObject)

}
