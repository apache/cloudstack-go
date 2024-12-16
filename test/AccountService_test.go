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

func TestAccountService(t *testing.T) {
	service := "AccountService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateAccount := func(t *testing.T) {
		if _, ok := response["createAccount"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewCreateAccountParams("email", "firstname", "lastname", "password", "username")
		r, err := client.Account.CreateAccount(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateAccount", testcreateAccount)

	testdeleteAccount := func(t *testing.T) {
		if _, ok := response["deleteAccount"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewDeleteAccountParams("id")
		_, err := client.Account.DeleteAccount(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteAccount", testdeleteAccount)

	testdisableAccount := func(t *testing.T) {
		if _, ok := response["disableAccount"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewDisableAccountParams(true)
		r, err := client.Account.DisableAccount(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DisableAccount", testdisableAccount)

	testenableAccount := func(t *testing.T) {
		if _, ok := response["enableAccount"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewEnableAccountParams()
		r, err := client.Account.EnableAccount(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("EnableAccount", testenableAccount)

	testisAccountAllowedToCreateOfferingsWithTags := func(t *testing.T) {
		if _, ok := response["isAccountAllowedToCreateOfferingsWithTags"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewIsAccountAllowedToCreateOfferingsWithTagsParams()
		_, err := client.Account.IsAccountAllowedToCreateOfferingsWithTags(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("IsAccountAllowedToCreateOfferingsWithTags", testisAccountAllowedToCreateOfferingsWithTags)

	testlinkAccountToLdap := func(t *testing.T) {
		if _, ok := response["linkAccountToLdap"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewLinkAccountToLdapParams("account", "domainid", "ldapdomain")
		_, err := client.Account.LinkAccountToLdap(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("LinkAccountToLdap", testlinkAccountToLdap)

	testlistAccounts := func(t *testing.T) {
		if _, ok := response["listAccounts"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewListAccountsParams()
		_, err := client.Account.ListAccounts(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListAccounts", testlistAccounts)

	testlistProjectAccounts := func(t *testing.T) {
		if _, ok := response["listProjectAccounts"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewListProjectAccountsParams("projectid")
		_, err := client.Account.ListProjectAccounts(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListProjectAccounts", testlistProjectAccounts)

	testlockAccount := func(t *testing.T) {
		if _, ok := response["lockAccount"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewLockAccountParams("account", "domainid")
		r, err := client.Account.LockAccount(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("LockAccount", testlockAccount)

	testmarkDefaultZoneForAccount := func(t *testing.T) {
		if _, ok := response["markDefaultZoneForAccount"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewMarkDefaultZoneForAccountParams("account", "domainid", "zoneid")
		r, err := client.Account.MarkDefaultZoneForAccount(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("MarkDefaultZoneForAccount", testmarkDefaultZoneForAccount)

	testupdateAccount := func(t *testing.T) {
		if _, ok := response["updateAccount"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Account.NewUpdateAccountParams()
		r, err := client.Account.UpdateAccount(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateAccount", testupdateAccount)

}
