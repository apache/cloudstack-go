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

func TestUCSService(t *testing.T) {
	service := "UCSService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddUcsManager := func(t *testing.T) {
		if _, ok := response["addUcsManager"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.UCS.NewAddUcsManagerParams("password", "url", "username", "zoneid")
		r, err := client.UCS.AddUcsManager(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddUcsManager", testaddUcsManager)

	testassociateUcsProfileToBlade := func(t *testing.T) {
		if _, ok := response["associateUcsProfileToBlade"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.UCS.NewAssociateUcsProfileToBladeParams("bladeid", "profiledn", "ucsmanagerid")
		r, err := client.UCS.AssociateUcsProfileToBlade(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AssociateUcsProfileToBlade", testassociateUcsProfileToBlade)

	testdeleteUcsManager := func(t *testing.T) {
		if _, ok := response["deleteUcsManager"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.UCS.NewDeleteUcsManagerParams("ucsmanagerid")
		_, err := client.UCS.DeleteUcsManager(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteUcsManager", testdeleteUcsManager)

	testlistUcsBlades := func(t *testing.T) {
		if _, ok := response["listUcsBlades"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.UCS.NewListUcsBladesParams("ucsmanagerid")
		_, err := client.UCS.ListUcsBlades(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListUcsBlades", testlistUcsBlades)

	testlistUcsManagers := func(t *testing.T) {
		if _, ok := response["listUcsManagers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.UCS.NewListUcsManagersParams()
		_, err := client.UCS.ListUcsManagers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListUcsManagers", testlistUcsManagers)

	testlistUcsProfiles := func(t *testing.T) {
		if _, ok := response["listUcsProfiles"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.UCS.NewListUcsProfilesParams("ucsmanagerid")
		_, err := client.UCS.ListUcsProfiles(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListUcsProfiles", testlistUcsProfiles)

}
