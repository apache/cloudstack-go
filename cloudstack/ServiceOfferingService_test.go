package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceOfferingService_CreateServiceOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"createserviceofferingresponse": {
				"serviceoffering": {
					"id": "efaeeab0-4b09-4729-8b6f-62645db41b37",
					"name": "testServiceOffering",
					"displaytext": "testServiceOffering",
					"cpunumber": 2,
					"cpuspeed": 1000,
					"memory": 2048,
					"created": "2021-10-04T07:16:42+0000",
					"storagetype": "shared",
					"provisioningtype": "thin",
					"offerha": false,
					"limitcpuuse": false,
					"isvolatile": false,
					"issystem": false,
					"defaultuse": false,
					"iscustomized": false,
					"cacheMode": "none",
					"rootdisksize": 0
				}
			}
		}`
		fmt.Fprintf(w, response)
	}))
	defer server.Close()

	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	p := client.ServiceOffering.NewCreateServiceOfferingParams("testServiceOffering", "testServiceOffering")
	resp, err := client.ServiceOffering.CreateServiceOffering(p)
	if err != nil {
		t.Errorf("Failed to create service offering due to %v", err)
		return
	}
	if resp == nil || resp.Name != "testServiceOffering" {
		t.Errorf("Failed to create service offering")
	}
}

func TestServiceOfferingService_UpdateServiceOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"updateserviceofferingresponse": {
				"serviceoffering": {
					"id": "efaeeab0-4b09-4729-8b6f-62645db41b37",
					"name": "testServiceOfferingUpdated",
					"displaytext": "testServiceOffering",
					"cpunumber": 2,
					"cpuspeed": 1000,
					"memory": 2048,
					"created": "2021-10-04T07:16:42+0000",
					"storagetype": "shared",
					"provisioningtype": "thin",
					"offerha": false,
					"limitcpuuse": false,
					"isvolatile": false,
					"issystem": false,
					"defaultuse": false,
					"iscustomized": false,
					"cacheMode": "none",
					"rootdisksize": 0
				}
			}
		}`
		fmt.Fprintln(w, response)
	}))
	defer server.Close()

	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	p := client.ServiceOffering.NewUpdateServiceOfferingParams("efaeeab0-4b09-4729-8b6f-62645db41b37")
	p.SetName("testServiceOfferingUpdated")
	resp, err := client.ServiceOffering.UpdateServiceOffering(p)
	if err != nil {
		t.Errorf("Failed to update service offering due to %v", err)
		return
	}
	fmt.Println(resp)
	if resp == nil || resp.Name != "testServiceOfferingUpdated" {
		t.Errorf("Failed to create service offering name")
	}
}

func TestServiceOfferingService_DeleteServiceOffering(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"deleteserviceofferingresponse": {
				"success": true
			}
		}`
		fmt.Fprintf(w, response)
	}))
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	p := client.ServiceOffering.NewDeleteServiceOfferingParams("efaeeab0-4b09-4729-8b6f-62645db41b37")
	resp, err := client.ServiceOffering.DeleteServiceOffering(p)
	if err != nil {
		t.Errorf("Failed to delete service offering due to %v", err)
		return
	}
	if resp == nil || !resp.Success {
		t.Errorf("Failed to delete service offering")
	}
}
