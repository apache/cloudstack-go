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

func TestLoadBalancerService(t *testing.T) {
	service := "LoadBalancerService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddNetscalerLoadBalancer := func(t *testing.T) {
		if _, ok := response["addNetscalerLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewAddNetscalerLoadBalancerParams("networkdevicetype", "password", "physicalnetworkid", "url", "username")
		_, err := client.LoadBalancer.AddNetscalerLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddNetscalerLoadBalancer", testaddNetscalerLoadBalancer)

	testassignCertToLoadBalancer := func(t *testing.T) {
		if _, ok := response["assignCertToLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewAssignCertToLoadBalancerParams("certid", "lbruleid")
		_, err := client.LoadBalancer.AssignCertToLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AssignCertToLoadBalancer", testassignCertToLoadBalancer)

	testassignToGlobalLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["assignToGlobalLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewAssignToGlobalLoadBalancerRuleParams("id", []string{})
		_, err := client.LoadBalancer.AssignToGlobalLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AssignToGlobalLoadBalancerRule", testassignToGlobalLoadBalancerRule)

	testassignToLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["assignToLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewAssignToLoadBalancerRuleParams("id")
		_, err := client.LoadBalancer.AssignToLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AssignToLoadBalancerRule", testassignToLoadBalancerRule)

	testconfigureNetscalerLoadBalancer := func(t *testing.T) {
		if _, ok := response["configureNetscalerLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewConfigureNetscalerLoadBalancerParams("lbdeviceid")
		_, err := client.LoadBalancer.ConfigureNetscalerLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ConfigureNetscalerLoadBalancer", testconfigureNetscalerLoadBalancer)

	testcreateGlobalLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["createGlobalLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewCreateGlobalLoadBalancerRuleParams("gslbdomainname", "gslbservicetype", "name", 0)
		_, err := client.LoadBalancer.CreateGlobalLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateGlobalLoadBalancerRule", testcreateGlobalLoadBalancerRule)

	testcreateLBHealthCheckPolicy := func(t *testing.T) {
		if _, ok := response["createLBHealthCheckPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewCreateLBHealthCheckPolicyParams("lbruleid")
		_, err := client.LoadBalancer.CreateLBHealthCheckPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateLBHealthCheckPolicy", testcreateLBHealthCheckPolicy)

	testcreateLBStickinessPolicy := func(t *testing.T) {
		if _, ok := response["createLBStickinessPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewCreateLBStickinessPolicyParams("lbruleid", "methodname", "name")
		_, err := client.LoadBalancer.CreateLBStickinessPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateLBStickinessPolicy", testcreateLBStickinessPolicy)

	testcreateLoadBalancer := func(t *testing.T) {
		if _, ok := response["createLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewCreateLoadBalancerParams("algorithm", 0, "name", "networkid", "scheme", "sourceipaddressnetworkid", 0)
		_, err := client.LoadBalancer.CreateLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateLoadBalancer", testcreateLoadBalancer)

	testcreateLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["createLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewCreateLoadBalancerRuleParams("algorithm", "name", 0, 0)
		_, err := client.LoadBalancer.CreateLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateLoadBalancerRule", testcreateLoadBalancerRule)

	testdeleteGlobalLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["deleteGlobalLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewDeleteGlobalLoadBalancerRuleParams("id")
		_, err := client.LoadBalancer.DeleteGlobalLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteGlobalLoadBalancerRule", testdeleteGlobalLoadBalancerRule)

	testdeleteLBHealthCheckPolicy := func(t *testing.T) {
		if _, ok := response["deleteLBHealthCheckPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewDeleteLBHealthCheckPolicyParams("id")
		_, err := client.LoadBalancer.DeleteLBHealthCheckPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteLBHealthCheckPolicy", testdeleteLBHealthCheckPolicy)

	testdeleteLBStickinessPolicy := func(t *testing.T) {
		if _, ok := response["deleteLBStickinessPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewDeleteLBStickinessPolicyParams("id")
		_, err := client.LoadBalancer.DeleteLBStickinessPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteLBStickinessPolicy", testdeleteLBStickinessPolicy)

	testdeleteLoadBalancer := func(t *testing.T) {
		if _, ok := response["deleteLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewDeleteLoadBalancerParams("id")
		_, err := client.LoadBalancer.DeleteLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteLoadBalancer", testdeleteLoadBalancer)

	testdeleteLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["deleteLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewDeleteLoadBalancerRuleParams("id")
		_, err := client.LoadBalancer.DeleteLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteLoadBalancerRule", testdeleteLoadBalancerRule)

	testdeleteNetscalerLoadBalancer := func(t *testing.T) {
		if _, ok := response["deleteNetscalerLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewDeleteNetscalerLoadBalancerParams("lbdeviceid")
		_, err := client.LoadBalancer.DeleteNetscalerLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteNetscalerLoadBalancer", testdeleteNetscalerLoadBalancer)

	testdeleteSslCert := func(t *testing.T) {
		if _, ok := response["deleteSslCert"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewDeleteSslCertParams("id")
		_, err := client.LoadBalancer.DeleteSslCert(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteSslCert", testdeleteSslCert)

	testlistGlobalLoadBalancerRules := func(t *testing.T) {
		if _, ok := response["listGlobalLoadBalancerRules"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewListGlobalLoadBalancerRulesParams()
		_, err := client.LoadBalancer.ListGlobalLoadBalancerRules(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListGlobalLoadBalancerRules", testlistGlobalLoadBalancerRules)

	testlistLBHealthCheckPolicies := func(t *testing.T) {
		if _, ok := response["listLBHealthCheckPolicies"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewListLBHealthCheckPoliciesParams()
		_, err := client.LoadBalancer.ListLBHealthCheckPolicies(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListLBHealthCheckPolicies", testlistLBHealthCheckPolicies)

	testlistLBStickinessPolicies := func(t *testing.T) {
		if _, ok := response["listLBStickinessPolicies"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewListLBStickinessPoliciesParams()
		_, err := client.LoadBalancer.ListLBStickinessPolicies(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListLBStickinessPolicies", testlistLBStickinessPolicies)

	testlistLoadBalancerRuleInstances := func(t *testing.T) {
		if _, ok := response["listLoadBalancerRuleInstances"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewListLoadBalancerRuleInstancesParams("id")
		_, err := client.LoadBalancer.ListLoadBalancerRuleInstances(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListLoadBalancerRuleInstances", testlistLoadBalancerRuleInstances)

	testlistLoadBalancerRules := func(t *testing.T) {
		if _, ok := response["listLoadBalancerRules"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewListLoadBalancerRulesParams()
		_, err := client.LoadBalancer.ListLoadBalancerRules(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListLoadBalancerRules", testlistLoadBalancerRules)

	testlistLoadBalancers := func(t *testing.T) {
		if _, ok := response["listLoadBalancers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewListLoadBalancersParams()
		_, err := client.LoadBalancer.ListLoadBalancers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListLoadBalancers", testlistLoadBalancers)

	testlistNetscalerLoadBalancers := func(t *testing.T) {
		if _, ok := response["listNetscalerLoadBalancers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewListNetscalerLoadBalancersParams()
		_, err := client.LoadBalancer.ListNetscalerLoadBalancers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListNetscalerLoadBalancers", testlistNetscalerLoadBalancers)

	testlistSslCerts := func(t *testing.T) {
		if _, ok := response["listSslCerts"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewListSslCertsParams()
		_, err := client.LoadBalancer.ListSslCerts(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListSslCerts", testlistSslCerts)

	testremoveCertFromLoadBalancer := func(t *testing.T) {
		if _, ok := response["removeCertFromLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewRemoveCertFromLoadBalancerParams("lbruleid")
		_, err := client.LoadBalancer.RemoveCertFromLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveCertFromLoadBalancer", testremoveCertFromLoadBalancer)

	testremoveFromGlobalLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["removeFromGlobalLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewRemoveFromGlobalLoadBalancerRuleParams("id", []string{})
		_, err := client.LoadBalancer.RemoveFromGlobalLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveFromGlobalLoadBalancerRule", testremoveFromGlobalLoadBalancerRule)

	testremoveFromLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["removeFromLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewRemoveFromLoadBalancerRuleParams("id")
		_, err := client.LoadBalancer.RemoveFromLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveFromLoadBalancerRule", testremoveFromLoadBalancerRule)

	testupdateGlobalLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["updateGlobalLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewUpdateGlobalLoadBalancerRuleParams("id")
		_, err := client.LoadBalancer.UpdateGlobalLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateGlobalLoadBalancerRule", testupdateGlobalLoadBalancerRule)

	testupdateLBHealthCheckPolicy := func(t *testing.T) {
		if _, ok := response["updateLBHealthCheckPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewUpdateLBHealthCheckPolicyParams("id")
		_, err := client.LoadBalancer.UpdateLBHealthCheckPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateLBHealthCheckPolicy", testupdateLBHealthCheckPolicy)

	testupdateLBStickinessPolicy := func(t *testing.T) {
		if _, ok := response["updateLBStickinessPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewUpdateLBStickinessPolicyParams("id")
		_, err := client.LoadBalancer.UpdateLBStickinessPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateLBStickinessPolicy", testupdateLBStickinessPolicy)

	testupdateLoadBalancer := func(t *testing.T) {
		if _, ok := response["updateLoadBalancer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewUpdateLoadBalancerParams("id")
		_, err := client.LoadBalancer.UpdateLoadBalancer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateLoadBalancer", testupdateLoadBalancer)

	testupdateLoadBalancerRule := func(t *testing.T) {
		if _, ok := response["updateLoadBalancerRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewUpdateLoadBalancerRuleParams("id")
		_, err := client.LoadBalancer.UpdateLoadBalancerRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateLoadBalancerRule", testupdateLoadBalancerRule)

	testuploadSslCert := func(t *testing.T) {
		if _, ok := response["uploadSslCert"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LoadBalancer.NewUploadSslCertParams("certificate", "name", "privatekey")
		_, err := client.LoadBalancer.UploadSslCert(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UploadSslCert", testuploadSslCert)

}
