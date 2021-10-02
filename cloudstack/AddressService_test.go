package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAssociateIpAddress(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responses := map[string]string{
			"associateIpAddress": `
		{
			"associateipaddressresponse": {
			"id": "dc51835f-b0e2-4a2e-91a4-6cccc44bdae3",
				"jobid": "8e805516-1729-46cf-a7d0-7289d523871e"
			}
		}`,
			"queryAsyncJobResult": `
		{
			"queryasyncjobresultresponse": {
			"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
				"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
				"cmd": "org.apache.cloudstack.api.command.admin.address.AssociateIPAddrCmdByAdmin",
				"jobstatus": 1,
				"jobprocstatus": 0,
				"jobresultcode": 0,
				"jobresulttype": "object",
				"jobresult": {
					"ipaddress": {
						"id": "dc51835f-b0e2-4a2e-91a4-6cccc44bdae3",
						"ipaddress": "10.70.3.100",
						"allocated": "2021-10-01T16:25:11+0000",
						"zoneid": "3fc049b0-87ae-4d77-90c1-cce70da17db6",
						"zonename": "testAdvZone2",
						"issourcenat": true,
						"account": "admin",
						"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
						"domain": "ROOT",
						"forvirtualnetwork": true,
						"vlanid": "2e86e486-b472-4f12-a9b2-bb73701241e0",
						"vlanname": "vlan://untagged",
						"isstaticnat": false,
						"issystem": false,
						"associatednetworkid": "cf8056db-25e1-49d1-b023-f13a717e5ecc",
						"associatednetworkname": "test-adv-network",
						"networkid": "f17d38db-4810-437e-be28-d38cd30e3034",
						"state": "Allocating",
						"physicalnetworkid": "0d6c1e76-2d83-459a-bbce-c133fbd732b7",
						"tags":[],
						"isportable": false,
						"fordisplay": true
					}
				},
				"jobinstancetype": "IpAddress",
				"jobinstanceid": "dc51835f-b0e2-4a2e-91a4-6cccc44bdae3",
				"created": "2021-10-01T16:25:12+0000",
				"completed": "2021-10-01T16:25:12+0000",
				"jobid": "8e805516-1729-46cf-a7d0-7289d523871e"
			}
		}`,
		}
		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))

	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	p := client.Address.NewAssociateIpAddressParams()
	p.SetZoneid("2e86e486-b472-4f12-a9b2-bb73701241e0")
	ip, err := client.Address.AssociateIpAddress(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if ip.Ipaddress != "10.70.3.100" {
		t.Errorf("ipaddress: actual %s, expected 10.70.3.100", ip.Ipaddress)
	}
}

func TestDisassociateIpAddress(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"disassociateIpAddress": `{
				"disassociateipaddressresponse": {
					"jobid": "44fc961e-9d57-4313-9f11-7556508b319c"
				}
			}`,
			"queryAsyncJobResult": `{
			"queryasyncjobresultresponse": {
			"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
				"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
				"cmd": "org.apache.cloudstack.api.command.user.address.DisassociateIPAddrCmd",
				"jobstatus": 1,
				"jobprocstatus": 0,
				"jobresultcode": 0,
				"jobresulttype": "object",
				"jobresult": {
					"success": true
				},
				"jobinstancetype": "IpAddress",
				"jobinstanceid": "a767fbe1-ed7a-4d7c-8221-c7d736ca622d",
				"created": "2021-10-01T16:42:32+0000",
				"completed": "2021-10-01T16:42:33+0000",
				"jobid": "44fc961e-9d57-4313-9f11-7556508b319c"
			}
		}`,
		}
		fmt.Fprintln(writer, responses[request.FormValue("command")])
	}))

	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	p := client.Address.NewDisassociateIpAddressParams("a767fbe1-ed7a-4d7c-8221-c7d736ca622d")
	address, err := client.Address.DisassociateIpAddress(p)
	if err != nil {
		t.Errorf("Failed to disassociate IP addres due to: %v", err.Error())
		return
	}
	if address.Success != true {
		t.Errorf("Failed to disassociate IP address")
	}
}
