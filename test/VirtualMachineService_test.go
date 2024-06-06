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

func TestVirtualMachineService(t *testing.T) {
	service := "VirtualMachineService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddNicToVirtualMachine := func(t *testing.T) {
		if _, ok := response["addNicToVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewAddNicToVirtualMachineParams("networkid", "virtualmachineid")
		r, err := client.VirtualMachine.AddNicToVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddNicToVirtualMachine", testaddNicToVirtualMachine)

	testassignVirtualMachine := func(t *testing.T) {
		if _, ok := response["assignVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewAssignVirtualMachineParams("virtualmachineid")
		r, err := client.VirtualMachine.AssignVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AssignVirtualMachine", testassignVirtualMachine)

	testchangeServiceForVirtualMachine := func(t *testing.T) {
		if _, ok := response["changeServiceForVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewChangeServiceForVirtualMachineParams("id", "serviceofferingid")
		r, err := client.VirtualMachine.ChangeServiceForVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ChangeServiceForVirtualMachine", testchangeServiceForVirtualMachine)

	testcleanVMReservations := func(t *testing.T) {
		if _, ok := response["cleanVMReservations"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewCleanVMReservationsParams()
		_, err := client.VirtualMachine.CleanVMReservations(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CleanVMReservations", testcleanVMReservations)

	testdeployVirtualMachine := func(t *testing.T) {
		if _, ok := response["deployVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewDeployVirtualMachineParams("serviceofferingid", "templateid", "zoneid")
		r, err := client.VirtualMachine.DeployVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DeployVirtualMachine", testdeployVirtualMachine)

	testdestroyVirtualMachine := func(t *testing.T) {
		if _, ok := response["destroyVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewDestroyVirtualMachineParams("id")
		r, err := client.VirtualMachine.DestroyVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DestroyVirtualMachine", testdestroyVirtualMachine)

	testexpungeVirtualMachine := func(t *testing.T) {
		if _, ok := response["expungeVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewExpungeVirtualMachineParams("id")
		_, err := client.VirtualMachine.ExpungeVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ExpungeVirtualMachine", testexpungeVirtualMachine)

	testgetVMPassword := func(t *testing.T) {
		if _, ok := response["getVMPassword"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewGetVMPasswordParams("id")
		_, err := client.VirtualMachine.GetVMPassword(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetVMPassword", testgetVMPassword)

	testlistVirtualMachines := func(t *testing.T) {
		if _, ok := response["listVirtualMachines"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewListVirtualMachinesParams()
		_, err := client.VirtualMachine.ListVirtualMachines(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVirtualMachines", testlistVirtualMachines)

	testlistVirtualMachinesMetrics := func(t *testing.T) {
		if _, ok := response["listVirtualMachinesMetrics"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewListVirtualMachinesMetricsParams()
		_, err := client.VirtualMachine.ListVirtualMachinesMetrics(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVirtualMachinesMetrics", testlistVirtualMachinesMetrics)

	testmigrateVirtualMachine := func(t *testing.T) {
		if _, ok := response["migrateVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewMigrateVirtualMachineParams("virtualmachineid")
		r, err := client.VirtualMachine.MigrateVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("MigrateVirtualMachine", testmigrateVirtualMachine)

	testmigrateVirtualMachineWithVolume := func(t *testing.T) {
		if _, ok := response["migrateVirtualMachineWithVolume"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewMigrateVirtualMachineWithVolumeParams("virtualmachineid")
		r, err := client.VirtualMachine.MigrateVirtualMachineWithVolume(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("MigrateVirtualMachineWithVolume", testmigrateVirtualMachineWithVolume)

	testrebootVirtualMachine := func(t *testing.T) {
		if _, ok := response["rebootVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewRebootVirtualMachineParams("id")
		r, err := client.VirtualMachine.RebootVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RebootVirtualMachine", testrebootVirtualMachine)

	testrecoverVirtualMachine := func(t *testing.T) {
		if _, ok := response["recoverVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewRecoverVirtualMachineParams("id")
		r, err := client.VirtualMachine.RecoverVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RecoverVirtualMachine", testrecoverVirtualMachine)

	testremoveNicFromVirtualMachine := func(t *testing.T) {
		if _, ok := response["removeNicFromVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewRemoveNicFromVirtualMachineParams("nicid", "virtualmachineid")
		r, err := client.VirtualMachine.RemoveNicFromVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RemoveNicFromVirtualMachine", testremoveNicFromVirtualMachine)

	testresetPasswordForVirtualMachine := func(t *testing.T) {
		if _, ok := response["resetPasswordForVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewResetPasswordForVirtualMachineParams("id")
		r, err := client.VirtualMachine.ResetPasswordForVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ResetPasswordForVirtualMachine", testresetPasswordForVirtualMachine)

	testrestoreVirtualMachine := func(t *testing.T) {
		if _, ok := response["restoreVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewRestoreVirtualMachineParams("virtualmachineid")
		r, err := client.VirtualMachine.RestoreVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RestoreVirtualMachine", testrestoreVirtualMachine)

	testscaleVirtualMachine := func(t *testing.T) {
		if _, ok := response["scaleVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewScaleVirtualMachineParams("id", "serviceofferingid")
		_, err := client.VirtualMachine.ScaleVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ScaleVirtualMachine", testscaleVirtualMachine)

	teststartVirtualMachine := func(t *testing.T) {
		if _, ok := response["startVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewStartVirtualMachineParams("id")
		r, err := client.VirtualMachine.StartVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("StartVirtualMachine", teststartVirtualMachine)

	teststopVirtualMachine := func(t *testing.T) {
		if _, ok := response["stopVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewStopVirtualMachineParams("id")
		r, err := client.VirtualMachine.StopVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("StopVirtualMachine", teststopVirtualMachine)

	testupdateDefaultNicForVirtualMachine := func(t *testing.T) {
		if _, ok := response["updateDefaultNicForVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewUpdateDefaultNicForVirtualMachineParams("nicid", "virtualmachineid")
		r, err := client.VirtualMachine.UpdateDefaultNicForVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateDefaultNicForVirtualMachine", testupdateDefaultNicForVirtualMachine)

	testupdateVirtualMachine := func(t *testing.T) {
		if _, ok := response["updateVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewUpdateVirtualMachineParams("id")
		r, err := client.VirtualMachine.UpdateVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateVirtualMachine", testupdateVirtualMachine)

	testlistVirtualMachinesUsageHistory := func(t *testing.T) {
		if _, ok := response["listVirtualMachinesUsageHistory"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewListVirtualMachinesUsageHistoryParams()
		_, err := client.VirtualMachine.ListVirtualMachinesUsageHistory(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVirtualMachinesUsageHistory", testlistVirtualMachinesUsageHistory)

	testimportVm := func(t *testing.T) {
		if _, ok := response["importVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewImportVmParams("clusterid", "hypervisor", "importsource", "name", "serviceofferingid", "zoneid")
		r, err := client.VirtualMachine.ImportVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ImportVm", testimportVm)

	testunmanageVirtualMachine := func(t *testing.T) {
		if _, ok := response["unmanageVirtualMachine"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewUnmanageVirtualMachineParams("id")
		_, err := client.VirtualMachine.UnmanageVirtualMachine(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UnmanageVirtualMachine", testunmanageVirtualMachine)

	testlistUnmanagedInstances := func(t *testing.T) {
		if _, ok := response["listUnmanagedInstances"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewListUnmanagedInstancesParams("clusterid")
		_, err := client.VirtualMachine.ListUnmanagedInstances(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListUnmanagedInstances", testlistUnmanagedInstances)

	testimportUnmanagedInstance := func(t *testing.T) {
		if _, ok := response["importUnmanagedInstance"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.VirtualMachine.NewImportUnmanagedInstanceParams("clusterid", "name", "serviceofferingid")
		r, err := client.VirtualMachine.ImportUnmanagedInstance(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ImportUnmanagedInstance", testimportUnmanagedInstance)

}
