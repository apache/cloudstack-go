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

func TestListAsyncJobs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
						"listasyncjobsresponse": {
  							"asyncjobs": [
								{
								  "accountid": "bc1b465f-1d18-11ec-9173-50eb7122da94",
								  "created": "2021-09-29T17:50:11+0530",
								  "jobid": "6679e6b9-4bf2-4d83-b9f0-235fc1609227",
								  "jobprocstatus": 0,
								  "jobresultcode": 0,
								  "userid": "bc1b60db-1d18-11ec-9173-50eb7122da94"
   								}
							],
							"count": 1
							}
						}`
		fmt.Fprintf(writer, response)
	}))

	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	p := client.Asyncjob.NewListAsyncJobsParams()
	p.SetListall(true)
	resp, err := client.Asyncjob.ListAsyncJobs(p)
	if err != nil {
		t.Errorf("Failed to fetch listAsyncJobs response, due to: %v", err)
	}

	if resp.Count != 1 {
		t.Errorf("Failed to fetch listAsyncJobs response")
	}
}
