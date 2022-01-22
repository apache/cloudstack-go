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
package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/stretchr/testify/assert"
)

const (
	simpleNetOffering = "DefaultIsolatedNetworkOfferingWithSourceNatService"
	myZone            = "my-zone"
)

func testClient() *cloudstack.CloudStackClient {
	apiURL := os.Getenv("ACS_API_URL")
	apiKey := os.Getenv("ACS_API_KEY")
	secretKey := os.Getenv("ACS_SECRET_KEY")
	cs := cloudstack.NewAsyncClient(apiURL, apiKey, secretKey, false)

	return cs
}

func TestUtils_singleID(t *testing.T) {

	utils := AcsUtils{Client: testClient()}

	id, err := utils.GetID(myZone, "zone", myZone)
	assert.NoError(t, err)
	assert.NotEqual(t, id, "")

	fmt.Printf("zone ID: %s", id)
}

func TestUtils_parID(t *testing.T) {
	testLen := 2

	utils := AcsUtils{Client: testClient()}

	params := make([][]string, testLen)
	params[0] = []string{myZone, "zone", myZone}
	params[1] = []string{myZone, "network_offering", simpleNetOffering}

	ids, err := utils.GetIDPar(&params)
	assert.NoError(t, err)
	assert.Len(t, ids, testLen)

	for _, id := range ids {
		fmt.Printf("resource ID: %s \n", id)
	}
}
