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
