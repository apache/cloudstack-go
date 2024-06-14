//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package test

import (
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestTemplateService(t *testing.T) {
	service := "TemplateService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcopyTemplate := func(t *testing.T) {
		if _, ok := response["copyTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewCopyTemplateParams("id")
		r, err := client.Template.CopyTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CopyTemplate", testcopyTemplate)

	testcreateTemplate := func(t *testing.T) {
		if _, ok := response["createTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewCreateTemplateParams("displaytext", "name", "ostypeid")
		r, err := client.Template.CreateTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateTemplate", testcreateTemplate)

	testdeleteTemplate := func(t *testing.T) {
		if _, ok := response["deleteTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewDeleteTemplateParams("id")
		_, err := client.Template.DeleteTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTemplate", testdeleteTemplate)

	testextractTemplate := func(t *testing.T) {
		if _, ok := response["extractTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewExtractTemplateParams("id", "mode")
		r, err := client.Template.ExtractTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ExtractTemplate", testextractTemplate)

	testgetUploadParamsForTemplate := func(t *testing.T) {
		if _, ok := response["getUploadParamsForTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewGetUploadParamsForTemplateParams("displaytext", "format", "hypervisor", "name", "zoneid")
		_, err := client.Template.GetUploadParamsForTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetUploadParamsForTemplate", testgetUploadParamsForTemplate)

	testlistTemplatePermissions := func(t *testing.T) {
		if _, ok := response["listTemplatePermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewListTemplatePermissionsParams("id")
		_, err := client.Template.ListTemplatePermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTemplatePermissions", testlistTemplatePermissions)

	testlistTemplates := func(t *testing.T) {
		if _, ok := response["listTemplates"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewListTemplatesParams("templatefilter")
		_, err := client.Template.ListTemplates(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTemplates", testlistTemplates)

	testprepareTemplate := func(t *testing.T) {
		if _, ok := response["prepareTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewPrepareTemplateParams("templateid", "zoneid")
		r, err := client.Template.PrepareTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("PrepareTemplate", testprepareTemplate)

	testregisterTemplate := func(t *testing.T) {
		if _, ok := response["registerTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewRegisterTemplateParams("displaytext", "format", "hypervisor", "name", "url")
		_, err := client.Template.RegisterTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RegisterTemplate", testregisterTemplate)

	testupdateTemplate := func(t *testing.T) {
		if _, ok := response["updateTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewUpdateTemplateParams("id")
		r, err := client.Template.UpdateTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateTemplate", testupdateTemplate)

	testupdateTemplatePermissions := func(t *testing.T) {
		if _, ok := response["updateTemplatePermissions"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewUpdateTemplatePermissionsParams("id")
		_, err := client.Template.UpdateTemplatePermissions(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateTemplatePermissions", testupdateTemplatePermissions)

	testupgradeRouterTemplate := func(t *testing.T) {
		if _, ok := response["upgradeRouterTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewUpgradeRouterTemplateParams()
		_, err := client.Template.UpgradeRouterTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpgradeRouterTemplate", testupgradeRouterTemplate)

	testlistTemplateDirectDownloadCertificates := func(t *testing.T) {
		if _, ok := response["listTemplateDirectDownloadCertificates"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewListTemplateDirectDownloadCertificatesParams()
		_, err := client.Template.ListTemplateDirectDownloadCertificates(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTemplateDirectDownloadCertificates", testlistTemplateDirectDownloadCertificates)

	testprovisionTemplateDirectDownloadCertificate := func(t *testing.T) {
		if _, ok := response["provisionTemplateDirectDownloadCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewProvisionTemplateDirectDownloadCertificateParams("hostid", "id")
		_, err := client.Template.ProvisionTemplateDirectDownloadCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ProvisionTemplateDirectDownloadCertificate", testprovisionTemplateDirectDownloadCertificate)

	testlinkUserDataToTemplate := func(t *testing.T) {
		if _, ok := response["linkUserDataToTemplate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Template.NewLinkUserDataToTemplateParams()
		r, err := client.Template.LinkUserDataToTemplate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("LinkUserDataToTemplate", testlinkUserDataToTemplate)

}
