package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDomainService_CreateDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"createdomainresponse": {
				"domain": {
					"id": "1530cc8c-efe7-4dff-84c4-96a1e8a1cd7a",
					"name": "testDomain",
					"level": 1,
					"parentdomainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
					"parentdomainname": "ROOT",
					"haschild": false,
					"path": "ROOT/testDomain",
					"secondarystoragetotal": 0
				}
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewCreateDomainParams("testDomain")
	params.SetParentdomainid("e4874e10-5fdf-11ea-9a56-1e006800018c")
	resp, err := client.Domain.CreateDomain(params)
	if err != nil {
		t.Errorf("Failed to create domain due to: %v", err)
	}
	fmt.Println(resp)
	if resp.Name != "testDomain" {
		t.Errorf("Failed to create domain")
	}
}

func TestDomainService_UpdateDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"updatedomainresponse": {
				"domain": {
					"id": "ee05fd92-7365-4421-a15b-abfa11dfc4f6",
					"name": "testDomainUpdated",
					"level": 1,
					"parentdomainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
					"parentdomainname": "ROOT",
					"haschild": false,
					"path": "ROOT/testDomainUpdated",
					"secondarystoragetotal": 0
				}
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewUpdateDomainParams("ee05fd92-7365-4421-a15b-abfa11dfc4f6")
	params.SetName("testDomainUpdated")
	resp, err := client.Domain.UpdateDomain(params)
	if err != nil {
		t.Errorf("Failed to update domain name due to: %v", err)
	}

	fmt.Println(resp)
	if resp.Name != "testDomainUpdated" {
		t.Errorf("Failed to update domain name")
	}
}

func TestDomainService_ListDomains(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listdomainsresponse": {
				"count": 1,
				"domain": [
					{
						"id": "097d3992-7a67-42e1-afb5-b4d2d81e280f",
						"name": "DummyDomain",
						"level": 1,
						"parentdomainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
						"parentdomainname": "ROOT",
						"haschild": false,
						"path": "ROOT/DummyDomain",
						"state": "Active",
						"vmlimit": "Unlimited",
						"vmtotal": 0,
						"vmavailable": "Unlimited",
						"iplimit": "Unlimited",
						"iptotal": 0,
						"ipavailable": "Unlimited",
						"volumelimit": "Unlimited",
						"volumetotal": 0,
						"volumeavailable": "Unlimited",
						"snapshotlimit": "Unlimited",
						"snapshottotal": 0,
						"snapshotavailable": "Unlimited",
						"templatelimit": "Unlimited",
						"templatetotal": 0,
						"templateavailable": "Unlimited",
						"projectlimit": "Unlimited",
						"projecttotal": 0,
						"projectavailable": "Unlimited",
						"networklimit": "Unlimited",
						"networktotal": 0,
						"networkavailable": "Unlimited",
						"vpclimit": "Unlimited",
						"vpctotal": 0,
						"vpcavailable": "Unlimited",
						"cpulimit": "Unlimited",
						"cputotal": 0,
						"cpuavailable": "Unlimited",
						"memorylimit": "Unlimited",
						"memorytotal": 0,
						"memoryavailable": "Unlimited",
						"primarystoragelimit": "-1",
						"primarystoragetotal": 0,
						"primarystorageavailable": "-1",
						"secondarystoragelimit": "-1",
						"secondarystoragetotal": 0,
						"secondarystorageavailable": "-1.0"
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewListDomainsParams()
	params.SetId("097d3992-7a67-42e1-afb5-b4d2d81e280f")
	resp, err := client.Domain.ListDomains(params)
	if err != nil {
		t.Errorf("Failed to list specific domain details due to: %v", err)
	}
	if resp.Count != 1 {
		t.Errorf("Failed to list specific domain details")
	}
	if resp.Domains[0].Name != "DummyDomain" {
		t.Errorf("Failed to fetch details of specific domain")
	}
}

func TestDomainService_ListDomainChildren(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listdomainchildrenresponse": {
				"count": 1,
				"domain": [
					{
						"id": "bfe3c14c-3a4a-4f85-b03c-308852967296",
						"name": "subdomain",
						"level": 2,
						"parentdomainid": "99becf06-7f0f-4eb4-bdc3-44fecb8cb829",
						"parentdomainname": "domain",
						"haschild": false,
						"path": "ROOT/domain/subdomain",
						"secondarystoragetotal": 0
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewListDomainChildrenParams()
	params.SetId("99becf06-7f0f-4eb4-bdc3-44fecb8cb829")
	resp, err := client.Domain.ListDomainChildren(params)
	if err != nil {
		t.Errorf("Failed to list specific domain's children due to: %v", err)
	}
	if resp.Count != 1 {
		t.Errorf("Failed to list specific domain's children ")
	}
}

func TestDomainService_DeleteDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := map[string]string{
			"deleteDomain": `{
				"deletedomainresponse": {
					"jobid": "318c77c0-8c42-42fb-8c23-363819c85dd3"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.domain.DeleteDomainCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"success": true
					},
					"created": "2021-10-02T17:44:42+0000",
					"completed": "2021-10-02T17:44:42+0000",
					"jobid": "318c77c0-8c42-42fb-8c23-363819c85dd3"
				}
			}`,
		}
		fmt.Fprintf(writer, response[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", false)
	params := client.Domain.NewDeleteDomainParams("ee05fd92-7365-4421-a15b-abfa11dfc4f6")
	resp, err := client.Domain.DeleteDomain(params)
	if err != nil {
		t.Errorf("Failed to delete domain due to: %v", err)
	}
	if !resp.Success {
		t.Errorf("Failed to delete domain")
	}
}
