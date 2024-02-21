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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/sbrueseke/cloudstack-go/v2/cloudstack"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	CS_API_URL    = "http://localhost:8080/client/api"
	CS_API_KEY    = "valid-api-key"
	CS_SECRET_KEY = "valid-secret-key"
)

func CreateTestServer(t *testing.T, responses map[string]json.RawMessage) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := responses[r.FormValue("command")]
		rawResult, _ := getRawValue(response)

		var result map[string]json.RawMessage
		err := json.Unmarshal(rawResult, &result)
		if err != nil {
			fmt.Fprintln(w, string(response))
			return
		}

		// Since we're using a sync client, pass the job result as the response
		val, ok := result["jobresult"]
		if !ok {
			fmt.Fprintln(w, string(response))
			return
		}

		err = json.Unmarshal(val, &result)
		if err != nil {
			fmt.Fprintln(w, string(response))
			return
		}

		// Handle success response separately
		if _, ok := result["success"]; ok {
			fmt.Fprintf(w, `{"jobresult":{"success":true}}`)
			return
		}

		fmt.Fprintln(w, string(val))

	}))
}

func getRawValue(b json.RawMessage) (json.RawMessage, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for _, v := range m {
		return v, nil
	}
	return nil, fmt.Errorf("Unable to extract the raw value from:\n\n%s\n\n", string(b))
}

func readData(file string) (map[string]json.RawMessage, error) {
	var data map[string]json.RawMessage
	apis, err := ioutil.ReadFile("testdata/" + file + ".json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(apis, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func TestCreateAsyncClient(t *testing.T) {
	client := cloudstack.NewAsyncClient(CS_API_URL, CS_API_KEY, CS_SECRET_KEY, true)

	if client == nil {
		t.Errorf("Failed to create Cloudstack Async Client")
	}
}

func TestCreateSyncClient(t *testing.T) {
	client := cloudstack.NewClient(CS_API_URL, CS_API_KEY, CS_SECRET_KEY, true)

	if client == nil {
		t.Errorf("Failed to create Cloudstack Client")
	}
}

func TestEncodeValues(t *testing.T) {
	RegisterFailHandler(Fail)

	DescribeTable("Encoding Values", func(value, expected string) {
		key := "testvalue"
		values := url.Values{}
		values.Set(key, value)
		Expect(cloudstack.EncodeValues(values)).To(Equal(fmt.Sprintf("%s=%s", key, expected)))
	},
		Entry("When asterisk is in value it shouldn't encode asterisk", "*.example.com", "*.example.com"),
		Entry("When question mark is in value it should encode question mark", "foo?", "foo%3F"),
		Entry("When question mark and asterisk is in value it should encode only question mark", "*.foo?", "*.foo%3F"),
		Entry("When a space is involved, should encode to %20 and not a +", "this that", "this%20that"),
	)

	RunSpecs(t, "Encoding Suite")
}

type UUIDStruct struct {
	Value cloudstack.UUID `json:"value"`
}

func TestUUID(t *testing.T) {
	valLong := `{"value": 4801878}`
	valString := `{"value": "994801878"}`
	valBool := `{"value": false}`
	res := UUIDStruct{}

	res.Value = ""
	if err := json.Unmarshal([]byte(valLong), &res); err != nil {
		t.Errorf("could not unserialize long into UUID: %s", err)
	}
	if res.Value != "4801878" {
		t.Errorf("unexpected value '%s', expecting 4801878", res.Value)
	}

	res.Value = ""
	if err := json.Unmarshal([]byte(valString), &res); err != nil {
		t.Errorf("could not unserialize string into UUID: %s", err)
	}
	if res.Value != "994801878" {
		t.Errorf("unexpected value '%s', expecting 994801878", res.Value)
	}

	res.Value = ""
	if err := json.Unmarshal([]byte(valBool), &res); err == nil {
		t.Errorf("missing expected error when serializing bool into UUID")
	}
}
