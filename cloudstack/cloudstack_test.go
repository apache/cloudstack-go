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

package cloudstack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	CS_API_URL    = "http://localhost:8080/client/api"
	CS_API_KEY    = "valid-api-key"
	CS_SECRET_KEY = "valid-secret-key"
)

func ReadData(apiName string, testDataFile string) (map[string]string, error) {
	var data interface{}
	apis, err := ioutil.ReadFile("testdata/" + testDataFile + "Data.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(apis), &data); err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(data.(map[string]interface{})[apiName])

	if err != nil {
		return nil, err
	}
	response := make(map[string]string)
	response[apiName] = string(jsonBytes)

	return response, nil
}

func ParseAsyncResponse(apiName string, filename string, r http.Request) (string, error) {
	response, err := ReadData(apiName, filename)
	if err != nil {
		fmt.Printf("Failed to read response data, due to: %v", err)
		return "", err
	}
	jsonStr := response[apiName]
	responses := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonStr), &responses)
	if err != nil {
		fmt.Printf("Failed to parse the response, due to: %v", err)
		return "", err
	}
	s, _ := json.Marshal(responses[r.FormValue("command")])
	return string(s), nil
}

func TestCreateAsyncClient(t *testing.T) {
	client := NewAsyncClient(CS_API_URL, CS_API_KEY, CS_SECRET_KEY, true)

	if client == nil {
		t.Errorf("Failed to create Cloudstack Async Client")
	}
}

func TestCreateSyncClient(t *testing.T) {
	client := NewClient(CS_API_URL, CS_API_KEY, CS_SECRET_KEY, true)

	if client == nil {
		t.Errorf("Failed to create Cloudstack Client")
	}
}

type CSLongStruct struct {
	Value CSLong `json:"value"`
}

func TestCSLong(t *testing.T) {
	valLong := `{"value": 4801878}`
	valString := `{"value": "994801878"}`
	valBool := `{"value": false}`
	res := CSLongStruct{}

	res.Value = ""
	if err := json.Unmarshal([]byte(valLong), &res); err != nil {
		t.Errorf("could not unserialize long into CSLong: %s", err)
	}
	if res.Value != "4801878" {
		t.Errorf("unepected value '%s', expecting 4801878", res.Value)
	}

	res.Value = ""
	if err := json.Unmarshal([]byte(valString), &res); err != nil {
		t.Errorf("could not unserialize string into CSLong: %s", err)
	}
	if res.Value != "994801878" {
		t.Errorf("unepected value '%s', expecting 994801878", res.Value)
	}

	res.Value = ""
	if err := json.Unmarshal([]byte(valBool), &res); err == nil {
		t.Errorf("missing expected error when serializing bool into CSLong")
	}
}
