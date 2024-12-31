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

func TestQuotaService(t *testing.T) {
	service := "QuotaService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testquotaBalance := func(t *testing.T) {
		if _, ok := response["quotaBalance"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaBalanceParams("account", "domainid")
		_, err := client.Quota.QuotaBalance(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("QuotaBalance", testquotaBalance)

	testquotaCredits := func(t *testing.T) {
		if _, ok := response["quotaCredits"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaCreditsParams("account", "domainid", 0)
		_, err := client.Quota.QuotaCredits(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("QuotaCredits", testquotaCredits)

	testquotaIsEnabled := func(t *testing.T) {
		if _, ok := response["quotaIsEnabled"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaIsEnabledParams()
		_, err := client.Quota.QuotaIsEnabled(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("QuotaIsEnabled", testquotaIsEnabled)

	testquotaStatement := func(t *testing.T) {
		if _, ok := response["quotaStatement"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaStatementParams("account", "domainid", "enddate", "startdate")
		_, err := client.Quota.QuotaStatement(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("QuotaStatement", testquotaStatement)

	testquotaSummary := func(t *testing.T) {
		if _, ok := response["quotaSummary"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaSummaryParams()
		_, err := client.Quota.QuotaSummary(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("QuotaSummary", testquotaSummary)

	testquotaTariffCreate := func(t *testing.T) {
		if _, ok := response["quotaTariffCreate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaTariffCreateParams("name", 0, 0)
		r, err := client.Quota.QuotaTariffCreate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("QuotaTariffCreate", testquotaTariffCreate)

	testquotaTariffDelete := func(t *testing.T) {
		if _, ok := response["quotaTariffDelete"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaTariffDeleteParams("id")
		_, err := client.Quota.QuotaTariffDelete(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("QuotaTariffDelete", testquotaTariffDelete)

	testquotaTariffList := func(t *testing.T) {
		if _, ok := response["quotaTariffList"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaTariffListParams()
		_, err := client.Quota.QuotaTariffList(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("QuotaTariffList", testquotaTariffList)

	testquotaTariffUpdate := func(t *testing.T) {
		if _, ok := response["quotaTariffUpdate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaTariffUpdateParams("name")
		r, err := client.Quota.QuotaTariffUpdate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("QuotaTariffUpdate", testquotaTariffUpdate)

	testquotaUpdate := func(t *testing.T) {
		if _, ok := response["quotaUpdate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Quota.NewQuotaUpdateParams()
		_, err := client.Quota.QuotaUpdate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("QuotaUpdate", testquotaUpdate)

}
