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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

// Regression test for issue #103: the generated Get<Entity>ID helpers must not
// panic with "index out of range [0] with length 0" when the API reports
// count >= 1 but returns an empty result slice (a count/slice mismatch that
// occurs e.g. on metrics APIs). They should return a descriptive error instead.
func TestGetVirtualMachinesMetricIDEmptyListWithNonZeroCount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// count = 1, but the "virtualmachine" list is empty
		fmt.Fprint(w, `{"listvirtualmachinesmetricsresponse":{"count":1,"virtualmachine":[]}}`)
	}))
	defer server.Close()

	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("GetVirtualMachinesMetricID panicked instead of returning an error: %v", r)
		}
	}()

	id, count, err := client.VirtualMachine.GetVirtualMachinesMetricID("anyvm")
	if err == nil {
		t.Fatalf("expected an error when count=1 but the result list is empty; got id=%q count=%d", id, count)
	}
	if id != "" {
		t.Fatalf("expected an empty id on the empty-list path, got %q", id)
	}
}
