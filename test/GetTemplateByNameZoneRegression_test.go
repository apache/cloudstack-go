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

// Regression test for issue #87: GetTemplateByName must carry the zoneid
// constraint into the by-ID lookup. Otherwise a template registered in multiple
// zones lists multiple rows for one UUID and the helper fails with
// "There is more then one result for Template UUID: ...!".
func TestGetTemplateByNameScopesByIDLookupToZone(t *testing.T) {
	const tmplID = "06145677-058a-456a-89a0-af4afd6fffcf"
	const zoneID = "11111111-2222-3333-4444-555555555555"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		if q.Get("id") != "" {
			// by-ID lookup: a multi-zone template returns one row PER zone unless
			// the zoneid is passed through (the bug). With the fix, zoneid is set.
			if q.Get("zoneid") != "" {
				fmt.Fprintf(w, `{"listtemplatesresponse":{"count":1,"template":[{"id":%q,"name":"mytmpl","zoneid":"zoneA"}]}}`, tmplID)
			} else {
				fmt.Fprintf(w, `{"listtemplatesresponse":{"count":2,"template":[{"id":%q,"name":"mytmpl","zoneid":"zoneA"},{"id":%q,"name":"mytmpl","zoneid":"zoneB"}]}}`, tmplID, tmplID)
			}
			return
		}
		// by-name lookup (GetTemplateID) resolves to a single template id
		fmt.Fprintf(w, `{"listtemplatesresponse":{"count":1,"template":[{"id":%q,"name":"mytmpl"}]}}`, tmplID)
	}))
	defer server.Close()

	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)

	tmpl, _, err := client.Template.GetTemplateByName("mytmpl", "all", zoneID)
	if err != nil {
		t.Fatalf("expected zone-scoped GetTemplateByName to succeed, got error: %v", err)
	}
	if tmpl == nil || tmpl.Id != tmplID {
		t.Fatalf("expected template %s, got %+v", tmplID, tmpl)
	}
}
