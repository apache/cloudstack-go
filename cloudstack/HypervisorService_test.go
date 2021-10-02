package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHypervisorService_ListSpecificHypervisorCapabilities(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listhypervisorcapabilitiesresponse": {
				"count": 1,
				"hypervisorCapabilities": [
					{
						"id": "1",
						"hypervisorversion": "default",
						"hypervisor": "XenServer",
						"maxguestslimit": 50,
						"securitygroupenabled": true,
						"maxdatavolumeslimit": 6,
						"storagemotionenabled": false
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Hypervisor.NewListHypervisorCapabilitiesParams()
	params.SetId("1")
	resp, err := client.Hypervisor.ListHypervisorCapabilities(params)
	if err != nil {
		t.Errorf("Failed to list hypervisor capabilities due to: %v", err)
		return
	}

	if resp.Count != 1 {
		t.Errorf("Failed list specific hypervisor capability")
	}
}
