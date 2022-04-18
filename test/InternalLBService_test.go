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

func TestInternalLBService(t *testing.T) {
	service := "InternalLBService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testconfigureInternalLoadBalancerElement := func(t *testing.T) {
		if _, ok := response["configureInternalLoadBalancerElement"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.InternalLB.NewConfigureInternalLoadBalancerElementParams(true, "id")
		_, err := client.InternalLB.ConfigureInternalLoadBalancerElement(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ConfigureInternalLoadBalancerElement", testconfigureInternalLoadBalancerElement)

	testcreateInternalLoadBalancerElement := func(t *testing.T) {
		if _, ok := response["createInternalLoadBalancerElement"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.InternalLB.NewCreateInternalLoadBalancerElementParams("nspid")
		_, err := client.InternalLB.CreateInternalLoadBalancerElement(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateInternalLoadBalancerElement", testcreateInternalLoadBalancerElement)

	testlistInternalLoadBalancerElements := func(t *testing.T) {
		if _, ok := response["listInternalLoadBalancerElements"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.InternalLB.NewListInternalLoadBalancerElementsParams()
		_, err := client.InternalLB.ListInternalLoadBalancerElements(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListInternalLoadBalancerElements", testlistInternalLoadBalancerElements)

	testlistInternalLoadBalancerVMs := func(t *testing.T) {
		if _, ok := response["listInternalLoadBalancerVMs"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.InternalLB.NewListInternalLoadBalancerVMsParams()
		_, err := client.InternalLB.ListInternalLoadBalancerVMs(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListInternalLoadBalancerVMs", testlistInternalLoadBalancerVMs)

	teststartInternalLoadBalancerVM := func(t *testing.T) {
		if _, ok := response["startInternalLoadBalancerVM"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.InternalLB.NewStartInternalLoadBalancerVMParams("id")
		_, err := client.InternalLB.StartInternalLoadBalancerVM(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("StartInternalLoadBalancerVM", teststartInternalLoadBalancerVM)

	teststopInternalLoadBalancerVM := func(t *testing.T) {
		if _, ok := response["stopInternalLoadBalancerVM"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.InternalLB.NewStopInternalLoadBalancerVMParams("id")
		_, err := client.InternalLB.StopInternalLoadBalancerVM(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("StopInternalLoadBalancerVM", teststopInternalLoadBalancerVM)

}
