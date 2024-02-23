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

	"github.com/sbrueseke/cloudstack-go/v2/cloudstack"
)

func TestEventService(t *testing.T) {
	service := "EventService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testarchiveEvents := func(t *testing.T) {
		if _, ok := response["archiveEvents"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Event.NewArchiveEventsParams()
		_, err := client.Event.ArchiveEvents(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ArchiveEvents", testarchiveEvents)

	testdeleteEvents := func(t *testing.T) {
		if _, ok := response["deleteEvents"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Event.NewDeleteEventsParams()
		_, err := client.Event.DeleteEvents(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteEvents", testdeleteEvents)

	testlistEventTypes := func(t *testing.T) {
		if _, ok := response["listEventTypes"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Event.NewListEventTypesParams()
		_, err := client.Event.ListEventTypes(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListEventTypes", testlistEventTypes)

	testlistEvents := func(t *testing.T) {
		if _, ok := response["listEvents"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Event.NewListEventsParams()
		_, err := client.Event.ListEvents(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListEvents", testlistEvents)

}
