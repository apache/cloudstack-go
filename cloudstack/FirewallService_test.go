package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFirewallService_CreateFirewallRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"createFirewallRule": `{
				"createfirewallruleresponse": {
					"id": "fb4ad2ee-02c8-433e-a769-6f18afddc750",
					"jobid": "104d5139-91dc-40a1-af2b-a94730e42e89"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "bc1b501c-1d18-11ec-9173-50eb7122da94",
					"userid": "bc1b6b47-1d18-11ec-9173-50eb7122da94",
					"cmd": "org.apache.cloudstack.api.command.user.firewall.CreateFirewallRuleCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"firewallrule": {
							"id": "fb4ad2ee-02c8-433e-a769-6f18afddc750",
							"protocol": "tcp",
							"ipaddressid": "a4aa35ef-34ad-4332-9247-378c83b8f927",
							"networkid": "c4a3303c-376d-4d56-b336-1bd91cb130b6",
							"ipaddress": "192.168.2.4",
							"state": "Active",
							"cidrlist": "0.0.0.0/0",
							"tags": [],
							"fordisplay": true
						}
					},
					"jobinstancetype": "FirewallRule",
					"jobinstanceid": "fb4ad2ee-02c8-433e-a769-6f18afddc750",
					"created": "2021-10-02T23:20:07+0530",
					"completed": "2021-10-02T23:20:07+0530",
					"jobid": "104d5139-91dc-40a1-af2b-a94730e42e89"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Firewall.NewCreateFirewallRuleParams("192.168.2.4", "tcp")
	resp, err := client.Firewall.CreateFirewallRule(params)
	if err != nil {
		t.Errorf("Failed to create firewall rule due to: %v", err)
	}
	if resp.Ipaddress != "192.168.2.4" {
		t.Errorf("Failed to create firewall rule")
	}

	if resp.State != "Active" {
		t.Errorf("Failed to create firewall rule")
	}
}

func TestFirewallService_DeleteFirewallRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"deleteFirewallRule": `{
				"deletefirewallruleresponse": {
					"jobid": "6ae96f78-1c28-45f9-a30b-761ac5f3a87b"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "bc1b501c-1d18-11ec-9173-50eb7122da94",
					"userid": "bc1b6b47-1d18-11ec-9173-50eb7122da94",
					"cmd": "org.apache.cloudstack.api.command.user.firewall.DeleteFirewallRuleCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"success": true
					},
					"created": "2021-10-02T23:25:22+0530",
					"completed": "2021-10-02T23:25:22+0530",
					"jobid": "6ae96f78-1c28-45f9-a30b-761ac5f3a87b"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Firewall.NewDeleteFirewallRuleParams("fb4ad2ee-02c8-433e-a769-6f18afddc750")
	resp, err := client.Firewall.DeleteFirewallRule(params)
	if err != nil {
		t.Errorf("Failed to delete firewall rule due to: %v", err)
	}
	if !resp.Success {
		t.Errorf("Failed to delete firewall rule")
	}
}

func TestFirewallService_CreateEgressFirewallRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"createEgressFirewallRule": `{
				"createegressfirewallruleresponse": {
					"id": "b7d8b539-8c72-4d25-8539-625f665681ad",
					"jobid": "8841a004-eb50-4d4e-aed5-7bc2494dc856"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "bc1b501c-1d18-11ec-9173-50eb7122da94",
					"userid": "bc1b6b47-1d18-11ec-9173-50eb7122da94",
					"cmd": "org.apache.cloudstack.api.command.user.firewall.CreateEgressFirewallRuleCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"firewallrule": {
							"id": "b7d8b539-8c72-4d25-8539-625f665681ad",
							"protocol": "tcp",
							"networkid": "c4a3303c-376d-4d56-b336-1bd91cb130b6",
							"state": "Active",
							"cidrlist": "10.1.1.0/24",
							"tags": [],
							"fordisplay": true,
							"destcidrlist": ""
						}
					},
					"jobinstancetype": "FirewallRule",
					"jobinstanceid": "b7d8b539-8c72-4d25-8539-625f665681ad",
					"created": "2021-10-02T23:29:41+0530",
					"completed": "2021-10-02T23:29:41+0530",
					"jobid": "8841a004-eb50-4d4e-aed5-7bc2494dc856"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Firewall.NewCreateEgressFirewallRuleParams("c4a3303c-376d-4d56-b336-1bd91cb130b6", "tcp")
	resp, err := client.Firewall.CreateEgressFirewallRule(params)
	if err != nil {
		t.Errorf("Failed to create egress firewall rule due to: %v", err)
	}

	if resp.State != "Active" {
		t.Errorf("Failed to create egress firewall rule")
	}
}

func TestFirewallService_DeleteEgressFirewallRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"deleteEgressFirewallRule": `{
				"deleteegressfirewallruleresponse": {
					"jobid": "91b10008-2e9f-412d-afa6-3c54a0bb19f5"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "bc1b501c-1d18-11ec-9173-50eb7122da94",
					"userid": "bc1b6b47-1d18-11ec-9173-50eb7122da94",
					"cmd": "org.apache.cloudstack.api.command.user.firewall.DeleteEgressFirewallRuleCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"success": true
					},
					"created": "2021-10-02T23:32:08+0530",
					"completed": "2021-10-02T23:32:08+0530",
					"jobid": "91b10008-2e9f-412d-afa6-3c54a0bb19f5"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Firewall.NewDeleteEgressFirewallRuleParams("fb4ad2ee-02c8-433e-a769-6f18afddc750")
	resp, err := client.Firewall.DeleteEgressFirewallRule(params)
	if err != nil {
		t.Errorf("Failed to delete egress firewall rule due to: %v", err)
	}
	if !resp.Success {
		t.Errorf("Failed to delete egress firewall rule")
	}
}
