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
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

// Regression test for list-response JSON-tag mismatches: the response struct
// tag must match the object key CloudStack actually returns, otherwise Count
// parses but the slice stays nil (silent data loss). Keys below are the
// authoritative server object names (setObjectName in the CloudStack source).
func TestListResponseJSONTagsPopulateSlices(t *testing.T) {
	t.Run("HypervisorCapabilities", func(t *testing.T) {
		var r cloudstack.ListHypervisorCapabilitiesResponse
		if err := json.Unmarshal([]byte(`{"count":1,"hypervisorCapabilities":[{"id":"h1"}]}`), &r); err != nil {
			t.Fatal(err)
		}
		if len(r.HypervisorCapabilities) != 1 {
			t.Fatalf("expected 1 item under key 'hypervisorCapabilities', got %d (nil slice = wrong json tag)", len(r.HypervisorCapabilities))
		}
	})

	t.Run("GuestNetworkIpv6Prefixes", func(t *testing.T) {
		var r cloudstack.ListGuestNetworkIpv6PrefixesResponse
		if err := json.Unmarshal([]byte(`{"count":1,"guestnetworkipv6prefix":[{"id":"p1"}]}`), &r); err != nil {
			t.Fatal(err)
		}
		if len(r.GuestNetworkIpv6Prefixes) != 1 {
			t.Fatalf("expected 1 item under key 'guestnetworkipv6prefix', got %d", len(r.GuestNetworkIpv6Prefixes))
		}
	})

	t.Run("LBHealthCheckPolicies", func(t *testing.T) {
		var r cloudstack.ListLBHealthCheckPoliciesResponse
		if err := json.Unmarshal([]byte(`{"count":1,"healthcheckpolicies":[{"lbruleid":"r1"}]}`), &r); err != nil {
			t.Fatal(err)
		}
		if len(r.LBHealthCheckPolicies) != 1 {
			t.Fatalf("expected 1 item under key 'healthcheckpolicies', got %d", len(r.LBHealthCheckPolicies))
		}
	})

	t.Run("LBStickinessPolicies", func(t *testing.T) {
		var r cloudstack.ListLBStickinessPoliciesResponse
		if err := json.Unmarshal([]byte(`{"count":1,"stickinesspolicies":[{"lbruleid":"r1"}]}`), &r); err != nil {
			t.Fatal(err)
		}
		if len(r.LBStickinessPolicies) != 1 {
			t.Fatalf("expected 1 item under key 'stickinesspolicies', got %d", len(r.LBStickinessPolicies))
		}
	})
}
