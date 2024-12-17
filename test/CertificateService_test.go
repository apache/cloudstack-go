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

func TestCertificateService(t *testing.T) {
	service := "CertificateService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testissueCertificate := func(t *testing.T) {
		if _, ok := response["issueCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewIssueCertificateParams()
		_, err := client.Certificate.IssueCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("IssueCertificate", testissueCertificate)

	testlistCAProviders := func(t *testing.T) {
		if _, ok := response["listCAProviders"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewListCAProvidersParams()
		_, err := client.Certificate.ListCAProviders(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCAProviders", testlistCAProviders)

	testlistCaCertificate := func(t *testing.T) {
		if _, ok := response["listCaCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewListCaCertificateParams()
		_, err := client.Certificate.ListCaCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCaCertificate", testlistCaCertificate)

	testlistTemplateDirectDownloadCertificates := func(t *testing.T) {
		if _, ok := response["listTemplateDirectDownloadCertificates"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewListTemplateDirectDownloadCertificatesParams()
		_, err := client.Certificate.ListTemplateDirectDownloadCertificates(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTemplateDirectDownloadCertificates", testlistTemplateDirectDownloadCertificates)

	testprovisionCertificate := func(t *testing.T) {
		if _, ok := response["provisionCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewProvisionCertificateParams("hostid")
		_, err := client.Certificate.ProvisionCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ProvisionCertificate", testprovisionCertificate)

	testprovisionTemplateDirectDownloadCertificate := func(t *testing.T) {
		if _, ok := response["provisionTemplateDirectDownloadCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewProvisionTemplateDirectDownloadCertificateParams("hostid", "id")
		_, err := client.Certificate.ProvisionTemplateDirectDownloadCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ProvisionTemplateDirectDownloadCertificate", testprovisionTemplateDirectDownloadCertificate)

	testrevokeCertificate := func(t *testing.T) {
		if _, ok := response["revokeCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewRevokeCertificateParams("serial")
		_, err := client.Certificate.RevokeCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RevokeCertificate", testrevokeCertificate)

	testrevokeTemplateDirectDownloadCertificate := func(t *testing.T) {
		if _, ok := response["revokeTemplateDirectDownloadCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewRevokeTemplateDirectDownloadCertificateParams("zoneid")
		_, err := client.Certificate.RevokeTemplateDirectDownloadCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RevokeTemplateDirectDownloadCertificate", testrevokeTemplateDirectDownloadCertificate)

	testuploadCustomCertificate := func(t *testing.T) {
		if _, ok := response["uploadCustomCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewUploadCustomCertificateParams("certificate", "domainsuffix")
		_, err := client.Certificate.UploadCustomCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UploadCustomCertificate", testuploadCustomCertificate)

	testuploadTemplateDirectDownloadCertificate := func(t *testing.T) {
		if _, ok := response["uploadTemplateDirectDownloadCertificate"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Certificate.NewUploadTemplateDirectDownloadCertificateParams("certificate", "hypervisor", "name", "zoneid")
		r, err := client.Certificate.UploadTemplateDirectDownloadCertificate(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UploadTemplateDirectDownloadCertificate", testuploadTemplateDirectDownloadCertificate)

}
