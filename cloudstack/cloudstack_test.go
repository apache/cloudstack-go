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
	"testing"
)

var (
	CS_API_URL    = "http://localhost:8080/client/api"
	CS_API_KEY    = "valid-api-key"
	CS_SECRET_KEY = "valid-secret-key"
)

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
