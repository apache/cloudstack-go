package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNetworkService_DedicateGuestVLANRange(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"dedicateguestvlanrangeresponse": {
				"dedicatedguestvlanrange": {
					"id": "b1afc5fe-b17c-4db3-8411-c323e4997ead",
					"account": "admin",
					"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
					"domain": "ROOT",
					"guestvlanrange": "100-110",
					"physicalnetworkid": 200,
					"zoneid": 1
				}
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.VLAN.NewDedicateGuestVlanRangeParams("8e27a637-7525-49ed-81ce-52bd5e5d9ea2", "100-110")
	params.SetAccount("admin")
	params.SetDomainid("e4874e10-5fdf-11ea-9a56-1e006800018c")
	resp, err := client.VLAN.DedicateGuestVlanRange(params)
	if err != nil {
		t.Errorf("Failed to dedicate guest VLAN range for physical network due to: %v", err)
		return
	}
	fmt.Println(resp)
	if resp == nil || resp.Guestvlanrange != "100-110" {
		t.Errorf("Failed to dedicate guest VLAN range for physical network")
	}
}
