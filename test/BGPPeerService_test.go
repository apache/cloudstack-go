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

func TestBGPPeerService(t *testing.T) {
	service := "BGPPeerService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testchangeBgpPeersForVpc := func(t *testing.T) {
		if _, ok := response["changeBgpPeersForVpc"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.BGPPeer.NewChangeBgpPeersForVpcParams("vpcid")
		r, err := client.BGPPeer.ChangeBgpPeersForVpc(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ChangeBgpPeersForVpc", testchangeBgpPeersForVpc)

	testcreateBgpPeer := func(t *testing.T) {
		if _, ok := response["createBgpPeer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.BGPPeer.NewCreateBgpPeerParams(0, "zoneid")
		r, err := client.BGPPeer.CreateBgpPeer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateBgpPeer", testcreateBgpPeer)

	testdedicateBgpPeer := func(t *testing.T) {
		if _, ok := response["dedicateBgpPeer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.BGPPeer.NewDedicateBgpPeerParams("id")
		r, err := client.BGPPeer.DedicateBgpPeer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DedicateBgpPeer", testdedicateBgpPeer)

	testdeleteBgpPeer := func(t *testing.T) {
		if _, ok := response["deleteBgpPeer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.BGPPeer.NewDeleteBgpPeerParams("id")
		_, err := client.BGPPeer.DeleteBgpPeer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteBgpPeer", testdeleteBgpPeer)

	testlistBgpPeers := func(t *testing.T) {
		if _, ok := response["listBgpPeers"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.BGPPeer.NewListBgpPeersParams()
		_, err := client.BGPPeer.ListBgpPeers(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListBgpPeers", testlistBgpPeers)

	testreleaseBgpPeer := func(t *testing.T) {
		if _, ok := response["releaseBgpPeer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.BGPPeer.NewReleaseBgpPeerParams("id")
		r, err := client.BGPPeer.ReleaseBgpPeer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ReleaseBgpPeer", testreleaseBgpPeer)

	testupdateBgpPeer := func(t *testing.T) {
		if _, ok := response["updateBgpPeer"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.BGPPeer.NewUpdateBgpPeerParams("id")
		r, err := client.BGPPeer.UpdateBgpPeer(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateBgpPeer", testupdateBgpPeer)

}
