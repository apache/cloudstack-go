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

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

type AsyncCall func(string, string, string) (string, error)

type TaskResponse struct {
	Err  error
	Resp string
}

// CloudstackUtils - utils package for Cloudstack
type AcsUtils struct {
	Client *cloudstack.CloudStackClient
}

// GetID - gets an ID of CloudStack resource
func (c AcsUtils) GetID(zone, resourceType, resourceName string) (string, error) {

	// If the supplied resourceName isn't a ID, try to retrieve the ID ourselves
	if cloudstack.IsID(resourceName) {
		return resourceName, nil
	}

	var err error
	var id string

	// Ignore counts, since an error is returned if there is no exact match
	switch resourceType {

	case "zone":
		id, _, err = c.Client.Zone.GetZoneID(resourceName)

	case "network_offering":
		id, _, err = c.Client.NetworkOffering.GetNetworkOfferingID(resourceName)

	case "vpc_offering":
		id, _, err = c.Client.VPC.GetVPCOfferingID(resourceName)

	case "service_offering":
		id, _, err = c.Client.ServiceOffering.GetServiceOfferingID(resourceName)

	case "disk_offering":
		id, _, err = c.Client.DiskOffering.GetDiskOfferingID(resourceName)

	case "template_id":
		id, _, err = c.Client.Template.GetTemplateID(resourceName, "executable", zone)

	case "project":
		id, _, err = c.Client.Project.GetProjectID(resourceName)

	case "os_type":
		p := c.Client.GuestOS.NewListOsTypesParams()
		p.SetDescription(resourceName)
		l, e := c.Client.GuestOS.ListOsTypes(p)
		if e != nil {
			err = e
			break
		}
		if l.Count == 1 {
			id = l.OsTypes[0].Id
			break
		}
		err = fmt.Errorf("could not find ID of OS Type: %s", resourceName)

	default:
		return id, fmt.Errorf("unknown resource type for zone %s: %s", zone, resourceType)
	}

	if err != nil {
		return "", err
	}

	return id, nil
}

// GetIDPar - runs multiple GetID requests in parallel
// Results are merged into a single output slice
func (c AcsUtils) GetIDPar(params *[][]string) ([]string, error) {
	len := len(*params)

	in := make(chan TaskResponse, len)
	ids := make([]string, len)

	// Run async requests
	for _, p := range *params {
		go makeIDRequest(in, c.GetID, p)
	}

	// Collect responses
	for i := 0; i < len; i++ {
		resp := <-in
		if resp.Err != nil {
			return nil, resp.Err
		}
		ids[i] = resp.Resp
	}

	return ids, nil
}

func makeIDRequest(out chan<- TaskResponse, f AsyncCall, keys []string) {

	if len(keys) < 3 {
		return
	}

	var res TaskResponse

	resp, err := f(keys[0], keys[1], keys[2])
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = resp
	out <- res
}
