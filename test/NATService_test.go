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

	"github.com/sbrueseke/cloudstack-go/v2/cloudstack"
)

func TestNATService(t *testing.T) {
	service := "NATService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateIpForwardingRule := func(t *testing.T) {
		if _, ok := response["createIpForwardingRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NAT.NewCreateIpForwardingRuleParams("ipaddressid", "protocol", 0)
		r, err := client.NAT.CreateIpForwardingRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateIpForwardingRule", testcreateIpForwardingRule)

	testdeleteIpForwardingRule := func(t *testing.T) {
		if _, ok := response["deleteIpForwardingRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NAT.NewDeleteIpForwardingRuleParams("id")
		_, err := client.NAT.DeleteIpForwardingRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteIpForwardingRule", testdeleteIpForwardingRule)

	testdisableStaticNat := func(t *testing.T) {
		if _, ok := response["disableStaticNat"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NAT.NewDisableStaticNatParams("ipaddressid")
		_, err := client.NAT.DisableStaticNat(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DisableStaticNat", testdisableStaticNat)

	testenableStaticNat := func(t *testing.T) {
		if _, ok := response["enableStaticNat"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NAT.NewEnableStaticNatParams("ipaddressid", "virtualmachineid")
		_, err := client.NAT.EnableStaticNat(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("EnableStaticNat", testenableStaticNat)

	testlistIpForwardingRules := func(t *testing.T) {
		if _, ok := response["listIpForwardingRules"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.NAT.NewListIpForwardingRulesParams()
		_, err := client.NAT.ListIpForwardingRules(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListIpForwardingRules", testlistIpForwardingRules)

}
