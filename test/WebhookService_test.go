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

func TestWebhookService(t *testing.T) {
	service := "WebhookService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testcreateWebhook := func(t *testing.T) {
		if _, ok := response["createWebhook"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Webhook.NewCreateWebhookParams("name", "payloadurl")
		r, err := client.Webhook.CreateWebhook(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateWebhook", testcreateWebhook)

	testdeleteWebhook := func(t *testing.T) {
		if _, ok := response["deleteWebhook"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Webhook.NewDeleteWebhookParams("id")
		_, err := client.Webhook.DeleteWebhook(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteWebhook", testdeleteWebhook)

	testdeleteWebhookDelivery := func(t *testing.T) {
		if _, ok := response["deleteWebhookDelivery"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Webhook.NewDeleteWebhookDeliveryParams()
		_, err := client.Webhook.DeleteWebhookDelivery(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteWebhookDelivery", testdeleteWebhookDelivery)

	testexecuteWebhookDelivery := func(t *testing.T) {
		if _, ok := response["executeWebhookDelivery"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Webhook.NewExecuteWebhookDeliveryParams()
		r, err := client.Webhook.ExecuteWebhookDelivery(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ExecuteWebhookDelivery", testexecuteWebhookDelivery)

	testlistWebhookDeliveries := func(t *testing.T) {
		if _, ok := response["listWebhookDeliveries"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Webhook.NewListWebhookDeliveriesParams()
		_, err := client.Webhook.ListWebhookDeliveries(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListWebhookDeliveries", testlistWebhookDeliveries)

	testlistWebhooks := func(t *testing.T) {
		if _, ok := response["listWebhooks"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Webhook.NewListWebhooksParams()
		_, err := client.Webhook.ListWebhooks(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListWebhooks", testlistWebhooks)

	testupdateWebhook := func(t *testing.T) {
		if _, ok := response["updateWebhook"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Webhook.NewUpdateWebhookParams("id")
		_, err := client.Webhook.UpdateWebhook(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateWebhook", testupdateWebhook)

}
