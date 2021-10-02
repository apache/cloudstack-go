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
