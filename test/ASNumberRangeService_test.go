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

func TestASNumberRangeService(t *testing.T) {
	service := "ASNumberRangeService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateASNRange := func(t *testing.T) {
		if _, ok := response["createASNRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ASNumberRange.NewCreateASNRangeParams(0, 0, "zoneid")
		r, err := client.ASNumberRange.CreateASNRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateASNRange", testcreateASNRange)

	testdeleteASNRange := func(t *testing.T) {
		if _, ok := response["deleteASNRange"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ASNumberRange.NewDeleteASNRangeParams("id")
		_, err := client.ASNumberRange.DeleteASNRange(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteASNRange", testdeleteASNRange)

	testlistASNRanges := func(t *testing.T) {
		if _, ok := response["listASNRanges"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ASNumberRange.NewListASNRangesParams()
		_, err := client.ASNumberRange.ListASNRanges(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListASNRanges", testlistASNRanges)

}
