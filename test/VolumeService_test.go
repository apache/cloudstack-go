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

func TestVolumeService(t *testing.T) {
	service := "VolumeService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testattachVolume := func(t *testing.T) {
		if _, ok := response["attachVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewAttachVolumeParams("id", "virtualmachineid")
		r, err := client.Volume.AttachVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AttachVolume", testattachVolume)

	testchangeOfferingForVolume := func(t *testing.T) {
		if _, ok := response["changeOfferingForVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewChangeOfferingForVolumeParams("diskofferingid", "id")
		r, err := client.Volume.ChangeOfferingForVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ChangeOfferingForVolume", testchangeOfferingForVolume)

	testcheckVolume := func(t *testing.T) {
		if _, ok := response["checkVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewCheckVolumeParams("id")
		r, err := client.Volume.CheckVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CheckVolume", testcheckVolume)

	testcreateVolume := func(t *testing.T) {
		if _, ok := response["createVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewCreateVolumeParams()
		r, err := client.Volume.CreateVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVolume", testcreateVolume)

	testdeleteVolume := func(t *testing.T) {
		if _, ok := response["deleteVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewDeleteVolumeParams("id")
		_, err := client.Volume.DeleteVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteVolume", testdeleteVolume)

	testdestroyVolume := func(t *testing.T) {
		if _, ok := response["destroyVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewDestroyVolumeParams("id")
		r, err := client.Volume.DestroyVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DestroyVolume", testdestroyVolume)

	testdetachVolume := func(t *testing.T) {
		if _, ok := response["detachVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewDetachVolumeParams()
		r, err := client.Volume.DetachVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DetachVolume", testdetachVolume)

	testextractVolume := func(t *testing.T) {
		if _, ok := response["extractVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewExtractVolumeParams("id", "mode", "zoneid")
		r, err := client.Volume.ExtractVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ExtractVolume", testextractVolume)

	testgetPathForVolume := func(t *testing.T) {
		if _, ok := response["getPathForVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewGetPathForVolumeParams("volumeid")
		_, err := client.Volume.GetPathForVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetPathForVolume", testgetPathForVolume)

	testgetUploadParamsForVolume := func(t *testing.T) {
		if _, ok := response["getUploadParamsForVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewGetUploadParamsForVolumeParams("format", "name", "zoneid")
		_, err := client.Volume.GetUploadParamsForVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetUploadParamsForVolume", testgetUploadParamsForVolume)

	testgetVolumeiScsiName := func(t *testing.T) {
		if _, ok := response["getVolumeiScsiName"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewGetVolumeiScsiNameParams("volumeid")
		_, err := client.Volume.GetVolumeiScsiName(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetVolumeiScsiName", testgetVolumeiScsiName)

	testimportVolume := func(t *testing.T) {
		if _, ok := response["importVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewImportVolumeParams("path", "storageid")
		r, err := client.Volume.ImportVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ImportVolume", testimportVolume)

	testlistElastistorVolume := func(t *testing.T) {
		if _, ok := response["listElastistorVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewListElastistorVolumeParams("id")
		_, err := client.Volume.ListElastistorVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListElastistorVolume", testlistElastistorVolume)

	testlistVolumes := func(t *testing.T) {
		if _, ok := response["listVolumes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewListVolumesParams()
		_, err := client.Volume.ListVolumes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVolumes", testlistVolumes)

	testlistVolumesForImport := func(t *testing.T) {
		if _, ok := response["listVolumesForImport"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewListVolumesForImportParams("storageid")
		_, err := client.Volume.ListVolumesForImport(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVolumesForImport", testlistVolumesForImport)

	testlistVolumesMetrics := func(t *testing.T) {
		if _, ok := response["listVolumesMetrics"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewListVolumesMetricsParams()
		_, err := client.Volume.ListVolumesMetrics(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVolumesMetrics", testlistVolumesMetrics)

	testmigrateVolume := func(t *testing.T) {
		if _, ok := response["migrateVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewMigrateVolumeParams("storageid", "volumeid")
		r, err := client.Volume.MigrateVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("MigrateVolume", testmigrateVolume)

	testrecoverVolume := func(t *testing.T) {
		if _, ok := response["recoverVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewRecoverVolumeParams("id")
		r, err := client.Volume.RecoverVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RecoverVolume", testrecoverVolume)

	testresizeVolume := func(t *testing.T) {
		if _, ok := response["resizeVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewResizeVolumeParams("id")
		r, err := client.Volume.ResizeVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ResizeVolume", testresizeVolume)

	testunmanageVolume := func(t *testing.T) {
		if _, ok := response["unmanageVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewUnmanageVolumeParams("id")
		_, err := client.Volume.UnmanageVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UnmanageVolume", testunmanageVolume)

	testupdateVolume := func(t *testing.T) {
		if _, ok := response["updateVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewUpdateVolumeParams()
		r, err := client.Volume.UpdateVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVolume", testupdateVolume)

	testuploadVolume := func(t *testing.T) {
		if _, ok := response["uploadVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewUploadVolumeParams("format", "name", "url", "zoneid")
		r, err := client.Volume.UploadVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UploadVolume", testuploadVolume)

	testlistVolumesUsageHistory := func(t *testing.T) {
		if _, ok := response["listVolumesUsageHistory"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewListVolumesUsageHistoryParams()
		_, err := client.Volume.ListVolumesUsageHistory(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVolumesUsageHistory", testlistVolumesUsageHistory)

	testassignVolume := func(t *testing.T) {
		if _, ok := response["assignVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Volume.NewAssignVolumeParams("volumeid")
		r, err := client.Volume.AssignVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AssignVolume", testassignVolume)

}
