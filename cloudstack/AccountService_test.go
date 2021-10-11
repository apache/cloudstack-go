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

package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAccounts(t *testing.T) {
	apiName := "listAccountsResponse"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := ReadData(apiName, "AccountService")
		if err != nil {
			t.Errorf("Failed to read response data, due to: %v", err)
			return
		}
		fmt.Fprintln(w, resp[apiName])
	}))
	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	p := client.Account.NewListAccountsParams()
	acc, _ := client.Account.ListAccounts(p)
	accounts := acc.Accounts
	if len(accounts) != 1 {
		t.Errorf("length: actual %d, expected 1", len(accounts))
	}

	if accounts[0].Name != "admin" {
		t.Errorf("name: actual %s, expected admin", accounts[0].Name)
	}

	if len(accounts[0].User) != 1 {
		t.Errorf("length of user: actual %d, expected 1", len(accounts[0].User))
	}

	if accounts[0].User[0].Username != "admin" {
		t.Errorf("username: actual %s, expected admin", accounts[0].User[0].Username)
	}

	if accounts[0].Accountdetails["key0"] != "value0" {
		t.Errorf("accountdetails[\"key0\"]: actual %s, expected value0", accounts[0].Accountdetails["key0"])
	}
}

func TestCreateAccounts(t *testing.T) {
	apiName := "createAccountsResponse"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := ReadData(apiName, "AccountService")
		if err != nil {
			t.Errorf("Failed to read response data, due to: %v", err)
			return
		}
		fmt.Fprintln(w, resp[apiName])
	}))
	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	email := "user01@example.com"
	firstname := "user01"
	lastname := "user01"
	password := "password"
	username := "user01"

	p := client.Account.NewCreateAccountParams(email, firstname, lastname, password, username)
	p.SetAccountdetails(map[string]string{"key0": "value0", "key1": "value1"})
	account, _ := client.Account.CreateAccount(p)

	if account.Name != "user01" {
		t.Errorf("name: actual %s, expected user01", account.Name)
	}

	if len(account.User) != 1 {
		t.Errorf("length of user: actual %d, expected 1", len(account.User))
	}

	if account.User[0].Username != "user01" {
		t.Errorf("username: actual %s, expected admin", account.User[0].Username)
	}

	if account.Accountdetails["key0"] != "value0" {
		t.Errorf("accountdetails[\"key0\"]: actual %s, expected value0", account.Accountdetails["key0"])
	}

	if account.Accountdetails["key1"] != "value1" {
		t.Errorf("accountdetails[\"key1\"]: actual %s, expected value1", account.Accountdetails["key1"])
	}
}
