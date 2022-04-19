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

func TestSystemVMService(t *testing.T) {
	service := "SystemVMService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testchangeServiceForSystemVm := func(t *testing.T) {
		if _, ok := response["changeServiceForSystemVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SystemVM.NewChangeServiceForSystemVmParams("id", "serviceofferingid")
		_, err := client.SystemVM.ChangeServiceForSystemVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ChangeServiceForSystemVm", testchangeServiceForSystemVm)

	testdestroySystemVm := func(t *testing.T) {
		if _, ok := response["destroySystemVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SystemVM.NewDestroySystemVmParams("id")
		_, err := client.SystemVM.DestroySystemVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DestroySystemVm", testdestroySystemVm)

	testlistSystemVms := func(t *testing.T) {
		if _, ok := response["listSystemVms"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SystemVM.NewListSystemVmsParams()
		_, err := client.SystemVM.ListSystemVms(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSystemVms", testlistSystemVms)

	testmigrateSystemVm := func(t *testing.T) {
		if _, ok := response["migrateSystemVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SystemVM.NewMigrateSystemVmParams("virtualmachineid")
		_, err := client.SystemVM.MigrateSystemVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("MigrateSystemVm", testmigrateSystemVm)

	testrebootSystemVm := func(t *testing.T) {
		if _, ok := response["rebootSystemVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SystemVM.NewRebootSystemVmParams("id")
		_, err := client.SystemVM.RebootSystemVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RebootSystemVm", testrebootSystemVm)

	testscaleSystemVm := func(t *testing.T) {
		if _, ok := response["scaleSystemVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SystemVM.NewScaleSystemVmParams("id", "serviceofferingid")
		_, err := client.SystemVM.ScaleSystemVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ScaleSystemVm", testscaleSystemVm)

	teststartSystemVm := func(t *testing.T) {
		if _, ok := response["startSystemVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SystemVM.NewStartSystemVmParams("id")
		_, err := client.SystemVM.StartSystemVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("StartSystemVm", teststartSystemVm)

	teststopSystemVm := func(t *testing.T) {
		if _, ok := response["stopSystemVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.SystemVM.NewStopSystemVmParams("id")
		_, err := client.SystemVM.StopSystemVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("StopSystemVm", teststopSystemVm)

}
