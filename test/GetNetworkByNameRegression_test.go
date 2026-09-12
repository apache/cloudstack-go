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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestNetworkLookupCompatibility(t *testing.T) {
	const networkID = "11111111-2222-4333-8444-555555555555"
	const zoneID = "22222222-3333-4444-8555-666666666666"
	existing := &cloudstack.Network{Id: networkID, Name: "existing"}
	multiple := []*cloudstack.Network{
		{Id: "33333333-4444-4555-8666-777777777777", Name: "existing-extra"},
		{Id: "44444444-5555-4666-8777-888888888888", Name: "other"},
		existing,
	}
	cases := []struct {
		name      string
		lookup    string
		networks  []*cloudstack.Network
		wantID    string
		wantCount int
	}{
		{"single match", "existing", []*cloudstack.Network{existing}, networkID, 1},
		{"missing with one network", "missing", []*cloudstack.Network{existing}, "", 0},
		{"exact match among partial matches", "existing", multiple, networkID, 2},
		{"missing with multiple networks", "missing", multiple, "", 0},
	}

	for _, supportsName := range []bool{false, true} {
		version := "legacy"
		if supportsName {
			version = "current"
		}
		t.Run(version, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					wantCount := tc.wantCount
					if supportsName && wantCount > 1 {
						wantCount = 1
					}
					server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						q := r.URL.Query()
						if q.Get("command") != "listNetworks" || q.Get("zoneid") != zoneID {
							t.Errorf("unexpected request parameters: %v", q)
							http.Error(w, "unexpected request parameters", http.StatusBadRequest)
							return
						}
						var matches []*cloudstack.Network
						for _, network := range tc.networks {
							if id := q.Get("id"); id != "" && network.Id != id {
								continue
							}
							if keyword := q.Get("keyword"); !strings.Contains(network.Name, keyword) {
								continue
							}
							if name := q.Get("name"); supportsName && name != "" && network.Name != name {
								continue
							}
							matches = append(matches, network)
						}
						w.Header().Set("Content-Type", "application/json")
						response := map[string]*cloudstack.ListNetworksResponse{
							"listnetworksresponse": {Count: len(matches), Networks: matches},
						}
						if err := json.NewEncoder(w).Encode(response); err != nil {
							t.Errorf("encode response: %v", err)
						}
					}))
					defer server.Close()
					client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)

					id, count, err := client.Network.GetNetworkID(tc.lookup, cloudstack.WithZone(zoneID))
					if id != tc.wantID || count != wantCount || (err != nil) != (tc.wantID == "") {
						t.Errorf("GetNetworkID(%q) = %q, %d, %v; want %q, %d",
							tc.lookup, id, count, err, tc.wantID, wantCount)
					}

					network, count, err := client.Network.GetNetworkByName(tc.lookup, cloudstack.WithZone(zoneID))
					if tc.wantID == "" {
						if network != nil || count != 0 || err == nil {
							t.Errorf("GetNetworkByName(%q) = %+v, %d, %v; want no match",
								tc.lookup, network, count, err)
						}
					} else if err != nil || network == nil || network.Id != tc.wantID || count != 1 {
						t.Errorf("GetNetworkByName(%q) = %+v, %d, %v; want network %s",
							tc.lookup, network, count, err, tc.wantID)
					}
				})
			}
		})
	}
}
