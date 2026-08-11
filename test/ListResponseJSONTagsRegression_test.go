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

// These assertions exist because the key cannot be derived from the API name, so
// a regeneration can otherwise quietly reintroduce any of them.
func TestListResponseKeysObservedFromServer(t *testing.T) {
	sliceCases := []struct {
		name   string
		key    string
		body   string
		decode func([]byte) (int, error)
	}{
		{
			name: "ASNRanges", key: "asnumberrange",
			body: `{"count":1,"asnumberrange":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListASNRangesResponse
				err := json.Unmarshal(b, &r)
				return len(r.ASNRanges), err
			},
		},
		{
			name: "Ipv4SubnetsForZone", key: "zoneipv4subnet",
			body: `{"count":1,"zoneipv4subnet":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListIpv4SubnetsForZoneResponse
				err := json.Unmarshal(b, &r)
				return len(r.Ipv4SubnetsForZone), err
			},
		},
		{
			name: "BackupProviders", key: "providers",
			body: `{"count":1,"providers":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListBackupProvidersResponse
				err := json.Unmarshal(b, &r)
				return len(r.BackupProviders), err
			},
		},
		{
			name: "ClustersMetrics", key: "cluster",
			body: `{"count":1,"cluster":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListClustersMetricsResponse
				err := json.Unmarshal(b, &r)
				return len(r.ClustersMetrics), err
			},
		},
		{
			name: "CustomActions", key: "extensioncustomaction",
			body: `{"count":1,"extensioncustomaction":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListCustomActionsResponse
				err := json.Unmarshal(b, &r)
				return len(r.CustomActions), err
			},
		},
		{
			name: "HostsMetrics", key: "host",
			body: `{"count":1,"host":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListHostsMetricsResponse
				err := json.Unmarshal(b, &r)
				return len(r.HostsMetrics), err
			},
		},
		{
			name: "NetworkIsolationMethods", key: "isolationmethod",
			body: `{"count":1,"isolationmethod":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListNetworkIsolationMethodsResponse
				err := json.Unmarshal(b, &r)
				return len(r.NetworkIsolationMethods), err
			},
		},
		{
			name: "RoutingFirewallRules", key: "firewallrule",
			body: `{"count":1,"firewallrule":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListRoutingFirewallRulesResponse
				err := json.Unmarshal(b, &r)
				return len(r.RoutingFirewallRules), err
			},
		},
		{
			name: "SupportedNetworkServices", key: "networkservice",
			body: `{"count":1,"networkservice":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListSupportedNetworkServicesResponse
				err := json.Unmarshal(b, &r)
				return len(r.SupportedNetworkServices), err
			},
		},
		{
			name: "SystemVmsUsageHistory", key: "virtualmachine",
			body: `{"count":1,"virtualmachine":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListSystemVmsUsageHistoryResponse
				err := json.Unmarshal(b, &r)
				return len(r.SystemVmsUsageHistory), err
			},
		},
		{
			name: "TrafficTypeImplementors", key: "traffictypeimplementorresponse",
			body: `{"count":1,"traffictypeimplementorresponse":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListTrafficTypeImplementorsResponse
				err := json.Unmarshal(b, &r)
				return len(r.TrafficTypeImplementors), err
			},
		},
		{
			name: "UserTwoFactorAuthenticatorProviders", key: "providers",
			body: `{"count":1,"providers":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListUserTwoFactorAuthenticatorProvidersResponse
				err := json.Unmarshal(b, &r)
				return len(r.UserTwoFactorAuthenticatorProviders), err
			},
		},
		{
			name: "VolumesMetrics", key: "volume",
			body: `{"count":1,"volume":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListVolumesMetricsResponse
				err := json.Unmarshal(b, &r)
				return len(r.VolumesMetrics), err
			},
		},
		{
			name: "ZonesMetrics", key: "zone",
			body: `{"count":1,"zone":[{}]}`,
			decode: func(b []byte) (int, error) {
				var r cloudstack.ListZonesMetricsResponse
				err := json.Unmarshal(b, &r)
				return len(r.ZonesMetrics), err
			},
		},
	}

	for _, tc := range sliceCases {
		t.Run(tc.name, func(t *testing.T) {
			n, err := tc.decode([]byte(tc.body))
			if err != nil {
				t.Fatalf("decoding key %q: %v", tc.key, err)
			}
			if n != 1 {
				t.Fatalf("expected 1 item under key %q, got %d (nil slice means the json tag does not match the key the server sends)", tc.key, n)
			}
		})
	}

	// These two return a single object and no count, so the field is a pointer
	// rather than a slice.
	objectCases := []struct {
		name   string
		key    string
		body   string
		decode func([]byte) (bool, error)
	}{
		{
			name: "CaCertificate", key: "cacertificates",
			body: `{"cacertificates":{}}`,
			decode: func(b []byte) (bool, error) {
				var r cloudstack.ListCaCertificateResponse
				err := json.Unmarshal(b, &r)
				return r.CaCertificate != nil, err
			},
		},
		{
			name: "UsageServerMetrics", key: "usageMetrics",
			body: `{"usageMetrics":{}}`,
			decode: func(b []byte) (bool, error) {
				var r cloudstack.ListUsageServerMetricsResponse
				err := json.Unmarshal(b, &r)
				return r.UsageServerMetrics != nil, err
			},
		},
	}

	for _, tc := range objectCases {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := tc.decode([]byte(tc.body))
			if err != nil {
				t.Fatalf("decoding key %q: %v", tc.key, err)
			}
			if !ok {
				t.Fatalf("expected an object under key %q, got nil", tc.key)
			}
		})
	}
}
