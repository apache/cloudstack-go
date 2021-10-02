package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateService_RegisterTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		response := `{
			"registertemplateresponse": {
				"count": 1,
				"template": [
					{
						"id": "5ce8f0b1-a910-4631-a8de-1e332bf3a6b7",
						"name": "testTemplate",
						"displaytext": "testTemplate",
						"ispublic": false,
						"created": "2021-10-04T08:41:23+0000",
						"isready": true,
						"passwordenabled": false,
						"format": "VHD",
						"isfeatured": false,
						"crossZones": false,
						"ostypeid": "f3404cc6-c38c-11eb-848b-1e006800018c",
						"ostypename": "Amazon Linux 3 (64 bit)",
						"account": "admin",
						"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
						"zonename": "shouldwork",
						"status": "Download Complete",
						"size": 5242880,
						"physicalsize": 5242880,
						"templatetype": "USER",
						"hypervisor": "Simulator",
						"domain": "ROOT",
						"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
						"isextractable": false,
						"details": {},
						"downloaddetails": [
							{
								"downloadState": "DOWNLOADED",
								"datastore": "secon",
								"downloadPercent": "100"
							}
						],
						"bits": 0,
						"sshkeyenabled": false,
						"isdynamicallyscalable": false,
						"directdownload": false,
						"deployasis": false,
						"requireshvm": true,
						"url": "http://dl.openvm.eu/cloudstack/macchinina/x86_64/macchinina-xen.vhd.bz2",
						"tags": []
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Template.NewRegisterTemplateParams("testTemplate", "VHD",
		"Simulator", "testTemplate", "http://dl.openvm.eu/cloudstack/macchinina/x86_64/macchinina-xen.vhd.bz2")
	resp, err := client.Template.RegisterTemplate(params)
	if err != nil {
		t.Errorf("Failed to register template due to %v", err)
		return
	}

	if resp == nil || resp.Count != 1 {
		t.Errorf("Failed to register template")
	}
}

func TestTemplateService_CreateTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		responses := map[string]string{
			"createTemplate": `{
				"createtemplateresponse": {
					"id": "b07dd9e4-1720-4122-b840-1e57a79e0dd9",
					"jobid": "13cc57e2-f388-49d5-95f3-de5ba43835fb"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.admin.template.CreateTemplateCmdByAdmin",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"template": {
							"id": "b07dd9e4-1720-4122-b840-1e57a79e0dd9",
							"name": "createTempFromVol",
							"displaytext": "createTempFromVol",
							"ispublic": false,
							"created": "2021-10-04T08:45:37+0000",
							"isready": true,
							"passwordenabled": false,
							"format": "RAW",
							"isfeatured": false,
							"crossZones": false,
							"ostypeid": "e510f742-5fdf-11ea-9a56-1e006800018c",
							"ostypename": "Other (64-bit)",
							"account": "admin",
							"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
							"zonename": "shouldwork",
							"status": "Download Complete",
							"size": 21474836480,
							"templatetype": "USER",
							"hypervisor": "Simulator",
							"domain": "ROOT",
							"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
							"isextractable": false,
							"details": {},
							"downloaddetails": [
								{
									"downloadState": "DOWNLOADED",
									"datastore": "secon",
									"downloadPercent": "100"
								}
							],
							"bits": 0,
							"sshkeyenabled": false,
							"isdynamicallyscalable": false,
							"directdownload": false,
							"deployasis": false,
							"requireshvm": false,
							"tags": []
						}
					},
					"jobinstancetype": "Template",
					"jobinstanceid": "b07dd9e4-1720-4122-b840-1e57a79e0dd9",
					"created": "2021-10-04T08:45:37+0000",
					"completed": "2021-10-04T08:45:37+0000",
					"jobid": "13cc57e2-f388-49d5-95f3-de5ba43835fb"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[r.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Template.NewCreateTemplateParams("createTempFromVol", "createTempFromVol", "e510f742-5fdf-11ea-9a56-1e006800018c")
	resp, err := client.Template.CreateTemplate(params)
	if err != nil {
		t.Errorf("Failed to create template from volume due to %v", err)
		return
	}
	if resp == nil || resp.Name != "createTempFromVol" {
		t.Errorf("Failed to create template from volume")
	}
}

func TestTemplateService_ExtractTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		responses := map[string]string{
			"extractTemplate": `{
				"extracttemplateresponse": {
					"jobid": "d2c86fca-26ca-436e-b5cf-60754c598da3"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
					"cmd": "org.apache.cloudstack.api.command.user.template.ExtractTemplateCmd",
					"jobstatus": 1,
					"jobprocstatus": 0,
					"jobresultcode": 0,
					"jobresulttype": "object",
					"jobresult": {
						"template": {
							"id": "5ce8f0b1-a910-4631-a8de-1e332bf3a6b7",
							"name": "testTemplate",
							"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
							"state": "DOWNLOAD_URL_CREATED",
							"zoneid": "1d8d87d4-1425-459c-8d81-c6f57dca2bd2",
							"zonename": "shouldwork",
							"extractMode": "HTTP_DOWNLOAD",
							"url": "http://172.17.1.11/userdata/a7a2ca0e-65c7-4b9e-bc61-3a4ab6adf6c0.vhd"
						}
					},
					"jobinstancetype": "Template",
					"jobinstanceid": "5ce8f0b1-a910-4631-a8de-1e332bf3a6b7",
					"created": "2021-10-04T08:50:45+0000",
					"completed": "2021-10-04T08:50:46+0000",
					"jobid": "d2c86fca-26ca-436e-b5cf-60754c598da3"
				}
			}`,
		}
		fmt.Fprintf(writer, responses[r.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.Template.NewExtractTemplateParams("5ce8f0b1-a910-4631-a8de-1e332bf3a6b7", "HTTP_DOWNLOAD")
	resp, err := client.Template.ExtractTemplate(params)
	if err != nil {
		t.Errorf("Failed to download template due to %v", err)
		return
	}
	if resp == nil || resp.Url == "" {
		t.Errorf("Failed to download template")
	}
}
