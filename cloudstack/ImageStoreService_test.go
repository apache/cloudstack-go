package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestImageStoreService_AddImageStore(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"addimagestoreresponse": {
				"imagestore": {
					"id": "0ac85364-e31a-4840-97a4-a237b4291dfa",
					"zoneid": "04ccc336-d730-42fe-8ff6-5ae36e141e81",
					"zonename": "SimZone1",
					"name": "nfs://192.168.0.20/export/testing/secondary",
					"url": "nfs://192.168.0.20/export/testing/secondary",
					"protocol": "nfs",
					"providername": "NFS",
					"scope": "ZONE",
					"readonly": false
				}
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.ImageStore.NewAddImageStoreParams("NFS")
	resp, err := client.ImageStore.AddImageStore(params)
	if err != nil {
		t.Errorf("Failed to add image store due to: %v", err)
		return
	}
	fmt.Println(resp)
	if resp == nil || resp.Providername != "NFS" {
		t.Errorf(" Failed to add image store")
	}
}

func TestImageStoreService_ListImageStores(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listimagestoresresponse": {
				"count": 1,
				"imagestore": [
					{
						"id": "06f9e780-fac0-42fa-ac0e-e6eb6a038178",
						"zoneid": "3fc049b0-87ae-4d77-90c1-cce70da17db6",
						"zonename": "testAdvZone2",
						"name": "testAdvSecondaryStorage",
						"url": "nfs://10.70.4.150/export/testing/secondary2",
						"protocol": "nfs",
						"providername": "NFS",
						"scope": "ZONE",
						"readonly": false
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.ImageStore.NewListImageStoresParams()
	resp, err := client.ImageStore.ListImageStores(params)
	if err != nil {
		t.Errorf("Failed to list image stores due to: %v", err)
		return
	}

	if resp == nil || resp.Count != 1 {
		t.Errorf(" Failed to list image stores")
	}
}

func TestImageStoreService_DeleteImageStore(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"deleteimagestoreresponse": {
				"success": true
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.ImageStore.NewDeleteImageStoreParams("0ac85364-e31a-4840-97a4-a237b4291dfa")
	resp, err := client.ImageStore.DeleteImageStore(params)
	if err != nil {
		t.Errorf("Failed to delete image store due to: %v", err)
		return
	}

	if resp == nil || !resp.Success {
		t.Errorf(" Failed to delete image store")
	}
}
