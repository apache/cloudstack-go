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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func readData() (map[string]interface{} , error) {
	var resp interface{}
	apis, err := ioutil.ReadFile("AccountServiceTestData.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(apis, &resp); err != nil {
		return nil, err
	}

	data := resp.(map[string]interface{})
	return data, nil
}

func TestListAccounts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := readData()
		if err != nil {
			return
		}
		resp := data["listAccountsResponse"]
		fmt.Fprintln(w, resp)
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
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "createaccountresponse": {
        "account": {
            "accountdetails": {
                "key0": "value0",
                "key1": "value1"
            },
            "accounttype": 2,
            "cpuavailable": "40",
            "cpulimit": "40",
            "cputotal": 0,
            "domain": "ROOT",
            "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
            "groups": [],
            "id": "c9cb9df8-dcd5-44c5-a40d-2e7d266669a6",
            "ipavailable": "20",
            "iplimit": "20",
            "iptotal": 0,
            "isdefault": false,
            "memoryavailable": "40960",
            "memorylimit": "40960",
            "memorytotal": 0,
            "name": "user01",
            "networkavailable": "20",
            "networklimit": "20",
            "networktotal": 0,
            "primarystorageavailable": "200",
            "primarystoragelimit": "200",
            "primarystoragetotal": 0,
            "projectavailable": "Unlimited",
            "projectlimit": "Unlimited",
            "projecttotal": 0,
            "secondarystorageavailable": "400",
            "secondarystoragelimit": "400",
            "secondarystoragetotal": 0,
            "snapshotavailable": "20",
            "snapshotlimit": "20",
            "snapshottotal": 0,
            "state": "enabled",
            "templateavailable": "20",
            "templatelimit": "20",
            "templatetotal": 0,
            "user": [
                {
                    "account": "user01",
                    "accountid": "c9cb9df8-dcd5-44c5-a40d-2e7d266669a6",
                    "accounttype": 2,
                    "created": "2016-07-26T02:17:17+0000",
                    "domain": "ROOT",
                    "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                    "email": "user01@example.com",
                    "firstname": "user01",
                    "id": "1d69ac45-a8f7-469c-ab24-d2b0382fa8be",
                    "iscallerchilddomain": false,
                    "isdefault": false,
                    "lastname": "user01",
                    "state": "enabled",
                    "username": "user01"
                }
            ],
            "vmavailable": "20",
            "vmlimit": "20",
            "vmtotal": 0,
            "volumeavailable": "20",
            "volumelimit": "20",
            "volumetotal": 0,
            "vpcavailable": "20",
            "vpclimit": "20",
            "vpctotal": 0
        }
    }
}
		`
		fmt.Fprintln(w, resp)
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
