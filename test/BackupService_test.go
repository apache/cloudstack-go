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

func TestBackupService(t *testing.T) {
	service := "BackupService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddBackupRepository := func(t *testing.T) {
		if _, ok := response["addBackupRepository"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewAddBackupRepositoryParams("address", "name", "type", "zoneid")
		r, err := client.Backup.AddBackupRepository(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddBackupRepository", testaddBackupRepository)

	testcreateBackup := func(t *testing.T) {
		if _, ok := response["createBackup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewCreateBackupParams("virtualmachineid")
		_, err := client.Backup.CreateBackup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateBackup", testcreateBackup)

	testcreateBackupSchedule := func(t *testing.T) {
		if _, ok := response["createBackupSchedule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewCreateBackupScheduleParams("intervaltype", "schedule", "timezone", "virtualmachineid")
		r, err := client.Backup.CreateBackupSchedule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateBackupSchedule", testcreateBackupSchedule)

	testcreateVMFromBackup := func(t *testing.T) {
		if _, ok := response["createVMFromBackup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewCreateVMFromBackupParams("backupid", "zoneid")
		r, err := client.Backup.CreateVMFromBackup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVMFromBackup", testcreateVMFromBackup)

	testdeleteBackup := func(t *testing.T) {
		if _, ok := response["deleteBackup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewDeleteBackupParams("id")
		_, err := client.Backup.DeleteBackup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteBackup", testdeleteBackup)

	testdeleteBackupOffering := func(t *testing.T) {
		if _, ok := response["deleteBackupOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewDeleteBackupOfferingParams("id")
		_, err := client.Backup.DeleteBackupOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteBackupOffering", testdeleteBackupOffering)

	testdeleteBackupRepository := func(t *testing.T) {
		if _, ok := response["deleteBackupRepository"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewDeleteBackupRepositoryParams("id")
		_, err := client.Backup.DeleteBackupRepository(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteBackupRepository", testdeleteBackupRepository)

	testdeleteBackupSchedule := func(t *testing.T) {
		if _, ok := response["deleteBackupSchedule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewDeleteBackupScheduleParams()
		_, err := client.Backup.DeleteBackupSchedule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteBackupSchedule", testdeleteBackupSchedule)

	testimportBackupOffering := func(t *testing.T) {
		if _, ok := response["importBackupOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewImportBackupOfferingParams(true, "description", "externalid", "name", "zoneid")
		r, err := client.Backup.ImportBackupOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ImportBackupOffering", testimportBackupOffering)

	testlistBackupOfferings := func(t *testing.T) {
		if _, ok := response["listBackupOfferings"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewListBackupOfferingsParams()
		_, err := client.Backup.ListBackupOfferings(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBackupOfferings", testlistBackupOfferings)

	testlistBackupProviderOfferings := func(t *testing.T) {
		if _, ok := response["listBackupProviderOfferings"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewListBackupProviderOfferingsParams("zoneid")
		_, err := client.Backup.ListBackupProviderOfferings(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBackupProviderOfferings", testlistBackupProviderOfferings)

	testlistBackupProviders := func(t *testing.T) {
		if _, ok := response["listBackupProviders"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewListBackupProvidersParams()
		_, err := client.Backup.ListBackupProviders(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBackupProviders", testlistBackupProviders)

	testlistBackupRepositories := func(t *testing.T) {
		if _, ok := response["listBackupRepositories"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewListBackupRepositoriesParams()
		_, err := client.Backup.ListBackupRepositories(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBackupRepositories", testlistBackupRepositories)

	testlistBackupSchedule := func(t *testing.T) {
		if _, ok := response["listBackupSchedule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewListBackupScheduleParams()
		_, err := client.Backup.ListBackupSchedule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBackupSchedule", testlistBackupSchedule)

	testlistBackups := func(t *testing.T) {
		if _, ok := response["listBackups"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewListBackupsParams()
		_, err := client.Backup.ListBackups(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBackups", testlistBackups)

	testrestoreBackup := func(t *testing.T) {
		if _, ok := response["restoreBackup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewRestoreBackupParams("id")
		_, err := client.Backup.RestoreBackup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RestoreBackup", testrestoreBackup)

	testupdateBackupOffering := func(t *testing.T) {
		if _, ok := response["updateBackupOffering"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewUpdateBackupOfferingParams("id")
		r, err := client.Backup.UpdateBackupOffering(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateBackupOffering", testupdateBackupOffering)

	testupdateBackupSchedule := func(t *testing.T) {
		if _, ok := response["updateBackupSchedule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Backup.NewUpdateBackupScheduleParams("intervaltype", "schedule", "timezone", "virtualmachineid")
		r, err := client.Backup.UpdateBackupSchedule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateBackupSchedule", testupdateBackupSchedule)

}
