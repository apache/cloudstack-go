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

func TestLimitService_ResetApiLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		apiName := "resetApiLimit"
		response, err := ReadData(apiName, "LimitService")
		if err != nil {
			t.Errorf("Failed to read response data due to: %v", err)
		}
		fmt.Fprintln(writer, response[apiName])
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Limit.NewResetApiLimitParams()
	resp, err := client.Limit.ResetApiLimit(params)
	if err != nil {
		t.Errorf("Failed to reset API Params limit due to: %v", err)
		return
	}

	if resp == nil {
		t.Errorf("Failed to register template")
	}
}
