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
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	simpleNetOffering = "DefaultIsolatedNetworkOfferingWithSourceNatService"
	myZone            = "my-zone"
)

func TestUtils_singleID(t *testing.T) {

	utils := MakeUtils()

	id, err := utils.GetID(myZone, "zone", myZone)
	assert.NoError(t, err)
	assert.NotEqual(t, id, "")

	fmt.Printf("zone ID: %s", id)
}

func TestUtils_parID(t *testing.T) {
	testLen := 2

	utils := MakeUtils()

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
