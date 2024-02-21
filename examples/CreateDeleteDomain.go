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
	"log"

	"github.com/sbrueseke/cloudstack-go/v2/cloudstack"
)

func CreateDomain() {
	domainName := "DummyDomain"
	cs := cloudstack.NewAsyncClient(ApiUrl, ApiKey, SecretKey, false)
	domain, _, err := cs.Domain.GetDomainByName(domainName)
	if err != nil {
		log.Printf("Failed to get domain with name %s due to %v", domainName, err)
	}
	if domain != nil {
		deleteParams := cs.Domain.NewDeleteDomainParams(domain.Id)
		delResp, err := cs.Domain.DeleteDomain(deleteParams)
		if err != nil {
			log.Fatalf("Failed to delete domain %s due to %v", domainName, err)
		}
		r, err := parseResponse(delResp)
		if err != nil {
			log.Fatalf("Failed to parse delete domain response due to %v", err)
			return
		}
		fmt.Printf("Delete response: %v\n\n", string(r))
	}

	p := &cloudstack.CreateDomainParams{}
	p.SetParentdomainid("e4874e10-5fdf-11ea-9a56-1e006800018c")
	p.SetName("DummyDomain")
	resp, err := cs.Domain.CreateDomain(p)
	if err != nil {
		log.Fatalf("Failed to create domain due to: %v", err)
		return
	}
	b, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Printf("%v", err)
		return
	}
	fmt.Printf("Create Domain response : %v\n\n", string(b))
}

func parseResponse(resp interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	return b, nil
}
