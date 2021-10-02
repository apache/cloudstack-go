package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNetworkOfferingService_CreateNetworkOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"createnetworkofferingresponse": {
				"networkoffering": {
					"id": "c2000683-68fd-437e-a4cf-f3676e1d18c1",
					"name": "testNetOffering",
					"displaytext": "testNetOffering",
					"traffictype": "Guest",
					"isdefault": false,
					"specifyvlan": true,
					"conservemode": true,
					"specifyipranges": false,
					"availability": "Optional",
					"networkrate": 200,
					"state": "Disabled",
					"guestiptype": "L2",
					"serviceofferingid": "c4b50796-22a5-4347-ae81-79911e44ac8d",
					"service": [],
					"forvpc": false,
					"ispersistent": false,
					"egressdefaultpolicy": true,
					"supportsstrechedl2subnet": false,
					"supportspublicaccess": false
				}
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.NetworkOffering.NewCreateNetworkOfferingParams("testNetOffering", "L2", "testNetOffering", []string{}, "Guest")
	resp, err := client.NetworkOffering.CreateNetworkOffering(params)
	if err != nil {
		t.Errorf("Failed to create network offering due to: %v", err)
	}
	if resp.Name != "testNetOffering" {
		t.Errorf("Failed to create network offering")
	}
}

func TestNetworkOfferingService_UpdateNetworkOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"updatenetworkofferingresponse": {
				"networkoffering": {
					"id": "c2000683-68fd-437e-a4cf-f3676e1d18c1",
					"name": "testNetOffering",
					"displaytext": "testNetOffering",
					"traffictype": "Guest",
					"isdefault": false,
					"specifyvlan": true,
					"conservemode": true,
					"specifyipranges": false,
					"availability": "Optional",
					"networkrate": 200,
					"state": "Enabled",
					"guestiptype": "L2",
					"serviceofferingid": "c4b50796-22a5-4347-ae81-79911e44ac8d",
					"service": [],
					"forvpc": false,
					"ispersistent": false,
					"egressdefaultpolicy": true,
					"supportsstrechedl2subnet": false,
					"supportspublicaccess": false
				}
			}
		}`
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.NetworkOffering.NewUpdateNetworkOfferingParams()
	params.SetState("Enabled")
	resp, err := client.NetworkOffering.UpdateNetworkOffering(params)
	if err != nil {
		t.Errorf("Failed to update network offering state due to: %v", err)
	}
	fmt.Println(resp)
	if resp.State != "Enabled" {
		t.Errorf("Failed to enable network offering")
	}
}

func TestNetworkOfferingService_ListNetworkOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listnetworkofferingsresponse": {
				"count": 1,
				"networkoffering": [
					{
						"id": "c2000683-68fd-437e-a4cf-f3676e1d18c1",
						"name": "testNetOffering",
						"displaytext": "testNetOffering",
						"traffictype": "Guest",
						"isdefault": false,
						"specifyvlan": true,
						"conservemode": true,
						"specifyipranges": false,
						"availability": "Optional",
						"networkrate": 200,
						"state": "Enabled",
						"guestiptype": "L2",
						"serviceofferingid": "c4b50796-22a5-4347-ae81-79911e44ac8d",
						"service": [],
						"forvpc": false,
						"ispersistent": false,
						"egressdefaultpolicy": true,
						"supportsstrechedl2subnet": false,
						"supportspublicaccess": false
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.NetworkOffering.NewListNetworkOfferingsParams()
	params.SetId("c2000683-68fd-437e-a4cf-f3676e1d18c1")
	resp, err := client.NetworkOffering.ListNetworkOfferings(params)
	if err != nil {
		t.Errorf("Failed to list network offering due to: %v", err)
	}

	if resp.Count != 1 {
		t.Errorf("Failed to list network offering")
	}
}

func TestNetworkOfferingService_DeleteNetworkOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"deletenetworkofferingresponse": {
				"success": true
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.NetworkOffering.NewDeleteNetworkOfferingParams("c2000683-68fd-437e-a4cf-f3676e1d18c1")
	resp, err := client.NetworkOffering.DeleteNetworkOffering(params)
	if err != nil {
		t.Errorf("Failed to delete network offering due to: %v", err)
	}

	if !resp.Success {
		t.Errorf("Failed to list network offering")
	}
}
