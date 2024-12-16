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

func TestDomainService(t *testing.T) {
	service := "DomainService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateDomain := func(t *testing.T) {
		if _, ok := response["createDomain"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Domain.NewCreateDomainParams("name")
		r, err := client.Domain.CreateDomain(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateDomain", testcreateDomain)

	testdeleteDomain := func(t *testing.T) {
		if _, ok := response["deleteDomain"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Domain.NewDeleteDomainParams("id")
		_, err := client.Domain.DeleteDomain(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteDomain", testdeleteDomain)

	testlistDomainChildren := func(t *testing.T) {
		if _, ok := response["listDomainChildren"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Domain.NewListDomainChildrenParams()
		_, err := client.Domain.ListDomainChildren(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListDomainChildren", testlistDomainChildren)

	testlistDomains := func(t *testing.T) {
		if _, ok := response["listDomains"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Domain.NewListDomainsParams()
		_, err := client.Domain.ListDomains(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListDomains", testlistDomains)

	testmoveDomain := func(t *testing.T) {
		if _, ok := response["moveDomain"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Domain.NewMoveDomainParams("domainid", "parentdomainid")
		r, err := client.Domain.MoveDomain(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("MoveDomain", testmoveDomain)

	testupdateDomain := func(t *testing.T) {
		if _, ok := response["updateDomain"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Domain.NewUpdateDomainParams("id")
		r, err := client.Domain.UpdateDomain(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateDomain", testupdateDomain)

}
