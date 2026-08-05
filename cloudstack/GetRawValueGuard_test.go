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

import "testing"

// getRawValue must not panic when a count-wrapped response carries an empty
// data array ({"count":0,"<entity>":[]}); it should return a descriptive error.
func TestGetRawValueEmptyArrayReturnsErrorNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("getRawValue panicked on a count-wrapped empty array: %v", r)
		}
	}()
	_, err := getRawValue([]byte(`{"count":0,"entity":[]}`))
	if err == nil {
		t.Fatal("expected an error for a count-wrapped empty array, got nil")
	}
}
