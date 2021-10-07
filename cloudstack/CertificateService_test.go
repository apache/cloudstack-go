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

func TestUploadCustomCertificate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"uploadCustomCertificate": `{
			"uploadcustomcertificateresponse": {
				"jobid": "bb14e24d-af95-4d9d-8ead-e4f42ccd83f7"
			}
		}`,
			"queryAsyncJobResult": `{
			"queryasyncjobresultresponse": {
				"accountid": "3bcbd8be-19eb-11ec-83a0-1e00bd000159",
				"userid": "3bcd30ee-19eb-11ec-83a0-1e00bd000159",
				"cmd": "org.apache.cloudstack.api.command.admin.resource.UploadCustomCertificateCmd",
				"jobstatus": 1,
				"jobprocstatus": 0,
				"jobresultcode": 0,
				"jobresulttype": "object",
				"jobresult": {
					"customcertificate": {
						"message": "Certificate has been successfully updated, if its the server certificate we would reboot all running console proxy VMs and secondary storage VMs to propagate the new certificate, please give a few minutes for console access and storage services service to be up and working again"
					}
				},
				"created": "2021-10-02T07:43:59+0000",
				"completed": "2021-10-02T07:43:59+0000",
				"jobid": "bb14e24d-af95-4d9d-8ead-e4f42ccd83f7"
			}
		}`,
		}
		fmt.Fprintln(writer, responses[request.FormValue("command")])
	}))

	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Certificate.NewUploadCustomCertificateParams("test", "xyz.com")
	resp, err := client.Certificate.UploadCustomCertificate(params)

	if err != nil {
		t.Errorf("Failed to upload custom certificate, due to: %v", err)
	}

	if resp.Jobstatus == 2 {
		t.Errorf("Failed to upload custom certificate")
	}

	if resp.Jobstatus == 1 {
		fmt.Println("Successfully uploaded certificate")
	}

}
