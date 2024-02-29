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

func TestFirewallService(t *testing.T) {
	service := "FirewallService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddPaloAltoFirewall := func(t *testing.T) {
		if _, ok := response["addPaloAltoFirewall"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewAddPaloAltoFirewallParams("networkdevicetype", "password", "physicalnetworkid", "url", "username")
		_, err := client.Firewall.AddPaloAltoFirewall(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddPaloAltoFirewall", testaddPaloAltoFirewall)

	testconfigurePaloAltoFirewall := func(t *testing.T) {
		if _, ok := response["configurePaloAltoFirewall"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewConfigurePaloAltoFirewallParams("fwdeviceid")
		_, err := client.Firewall.ConfigurePaloAltoFirewall(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ConfigurePaloAltoFirewall", testconfigurePaloAltoFirewall)

	testcreateEgressFirewallRule := func(t *testing.T) {
		if _, ok := response["createEgressFirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewCreateEgressFirewallRuleParams("networkid", "protocol")
		r, err := client.Firewall.CreateEgressFirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateEgressFirewallRule", testcreateEgressFirewallRule)

	testcreateFirewallRule := func(t *testing.T) {
		if _, ok := response["createFirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewCreateFirewallRuleParams("ipaddressid", "protocol")
		r, err := client.Firewall.CreateFirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateFirewallRule", testcreateFirewallRule)

	testcreatePortForwardingRule := func(t *testing.T) {
		if _, ok := response["createPortForwardingRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewCreatePortForwardingRuleParams("ipaddressid", 0, "protocol", 0, "virtualmachineid")
		r, err := client.Firewall.CreatePortForwardingRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreatePortForwardingRule", testcreatePortForwardingRule)

	testdeleteEgressFirewallRule := func(t *testing.T) {
		if _, ok := response["deleteEgressFirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewDeleteEgressFirewallRuleParams("id")
		_, err := client.Firewall.DeleteEgressFirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteEgressFirewallRule", testdeleteEgressFirewallRule)

	testdeleteFirewallRule := func(t *testing.T) {
		if _, ok := response["deleteFirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewDeleteFirewallRuleParams("id")
		_, err := client.Firewall.DeleteFirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteFirewallRule", testdeleteFirewallRule)

	testdeletePaloAltoFirewall := func(t *testing.T) {
		if _, ok := response["deletePaloAltoFirewall"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewDeletePaloAltoFirewallParams("fwdeviceid")
		_, err := client.Firewall.DeletePaloAltoFirewall(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeletePaloAltoFirewall", testdeletePaloAltoFirewall)

	testdeletePortForwardingRule := func(t *testing.T) {
		if _, ok := response["deletePortForwardingRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewDeletePortForwardingRuleParams("id")
		_, err := client.Firewall.DeletePortForwardingRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeletePortForwardingRule", testdeletePortForwardingRule)

	testlistEgressFirewallRules := func(t *testing.T) {
		if _, ok := response["listEgressFirewallRules"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewListEgressFirewallRulesParams()
		_, err := client.Firewall.ListEgressFirewallRules(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListEgressFirewallRules", testlistEgressFirewallRules)

	testlistFirewallRules := func(t *testing.T) {
		if _, ok := response["listFirewallRules"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewListFirewallRulesParams()
		_, err := client.Firewall.ListFirewallRules(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListFirewallRules", testlistFirewallRules)

	testlistPaloAltoFirewalls := func(t *testing.T) {
		if _, ok := response["listPaloAltoFirewalls"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewListPaloAltoFirewallsParams()
		_, err := client.Firewall.ListPaloAltoFirewalls(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListPaloAltoFirewalls", testlistPaloAltoFirewalls)

	testlistPortForwardingRules := func(t *testing.T) {
		if _, ok := response["listPortForwardingRules"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewListPortForwardingRulesParams()
		_, err := client.Firewall.ListPortForwardingRules(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListPortForwardingRules", testlistPortForwardingRules)

	testupdateEgressFirewallRule := func(t *testing.T) {
		if _, ok := response["updateEgressFirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewUpdateEgressFirewallRuleParams("id")
		r, err := client.Firewall.UpdateEgressFirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateEgressFirewallRule", testupdateEgressFirewallRule)

	testupdateFirewallRule := func(t *testing.T) {
		if _, ok := response["updateFirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewUpdateFirewallRuleParams("id")
		r, err := client.Firewall.UpdateFirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateFirewallRule", testupdateFirewallRule)

	testupdatePortForwardingRule := func(t *testing.T) {
		if _, ok := response["updatePortForwardingRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewUpdatePortForwardingRuleParams("id")
		r, err := client.Firewall.UpdatePortForwardingRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdatePortForwardingRule", testupdatePortForwardingRule)

	testlistIpv6FirewallRules := func(t *testing.T) {
		if _, ok := response["listIpv6FirewallRules"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewListIpv6FirewallRulesParams()
		_, err := client.Firewall.ListIpv6FirewallRules(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListIpv6FirewallRules", testlistIpv6FirewallRules)

	testcreateIpv6FirewallRule := func(t *testing.T) {
		if _, ok := response["createIpv6FirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewCreateIpv6FirewallRuleParams("networkid", "protocol")
		r, err := client.Firewall.CreateIpv6FirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateIpv6FirewallRule", testcreateIpv6FirewallRule)

	testupdateIpv6FirewallRule := func(t *testing.T) {
		if _, ok := response["updateIpv6FirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewUpdateIpv6FirewallRuleParams("id")
		r, err := client.Firewall.UpdateIpv6FirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateIpv6FirewallRule", testupdateIpv6FirewallRule)

	testdeleteIpv6FirewallRule := func(t *testing.T) {
		if _, ok := response["deleteIpv6FirewallRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Firewall.NewDeleteIpv6FirewallRuleParams("id")
		_, err := client.Firewall.DeleteIpv6FirewallRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteIpv6FirewallRule", testdeleteIpv6FirewallRule)

}
