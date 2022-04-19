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

func TestLDAPService(t *testing.T) {
	service := "LDAPService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddLdapConfiguration := func(t *testing.T) {
		if _, ok := response["addLdapConfiguration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewAddLdapConfigurationParams("hostname", 0)
		_, err := client.LDAP.AddLdapConfiguration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddLdapConfiguration", testaddLdapConfiguration)

	testdeleteLdapConfiguration := func(t *testing.T) {
		if _, ok := response["deleteLdapConfiguration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewDeleteLdapConfigurationParams("hostname")
		_, err := client.LDAP.DeleteLdapConfiguration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteLdapConfiguration", testdeleteLdapConfiguration)

	testimportLdapUsers := func(t *testing.T) {
		if _, ok := response["importLdapUsers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewImportLdapUsersParams()
		_, err := client.LDAP.ImportLdapUsers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ImportLdapUsers", testimportLdapUsers)

	testldapConfig := func(t *testing.T) {
		if _, ok := response["ldapConfig"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewLdapConfigParams()
		_, err := client.LDAP.LdapConfig(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("LdapConfig", testldapConfig)

	testldapCreateAccount := func(t *testing.T) {
		if _, ok := response["ldapCreateAccount"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewLdapCreateAccountParams("username")
		_, err := client.LDAP.LdapCreateAccount(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("LdapCreateAccount", testldapCreateAccount)

	testldapRemove := func(t *testing.T) {
		if _, ok := response["ldapRemove"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewLdapRemoveParams()
		_, err := client.LDAP.LdapRemove(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("LdapRemove", testldapRemove)

	testlinkDomainToLdap := func(t *testing.T) {
		if _, ok := response["linkDomainToLdap"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewLinkDomainToLdapParams(0, "domainid", "type")
		_, err := client.LDAP.LinkDomainToLdap(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("LinkDomainToLdap", testlinkDomainToLdap)

	testlistLdapConfigurations := func(t *testing.T) {
		if _, ok := response["listLdapConfigurations"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewListLdapConfigurationsParams()
		_, err := client.LDAP.ListLdapConfigurations(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListLdapConfigurations", testlistLdapConfigurations)

	testlistLdapUsers := func(t *testing.T) {
		if _, ok := response["listLdapUsers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewListLdapUsersParams()
		_, err := client.LDAP.ListLdapUsers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListLdapUsers", testlistLdapUsers)

	testsearchLdap := func(t *testing.T) {
		if _, ok := response["searchLdap"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.LDAP.NewSearchLdapParams("query")
		_, err := client.LDAP.SearchLdap(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("SearchLdap", testsearchLdap)

}
