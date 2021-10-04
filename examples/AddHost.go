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

package main

import (
	"encoding/json"
	"fmt"
	"github.com/apache/cloudstack-go/v2/cloudstack"
	"log"
)

func main()  {
	cs := cloudstack.NewAsyncClient(ApiUrl, ApiKey, SecretKey, false)
	p := cs.Host.NewAddHostParams("Simulator", "password", PodId,
		"http://sim/c0/h0", "root", ZoneId)
	resp, err := cs.Host.AddHost(p)
	if err != nil {
		fmt.Errorf("Failed to add host due to: %v", err)
	}

	b, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Printf("Host response : %v", string(b))
}