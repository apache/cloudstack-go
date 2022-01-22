package utils

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

import (
	"fmt"
	"log"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

type AcsError struct {
	zone  string
	name  string
	value string
	err   error
}

func (e AcsError) Error() error {
	return fmt.Errorf("Error retrieving ID of %s %s %s: %s", e.zone, e.name, e.value, e.err)
}

type CloudstackUtils struct {
	Client *cloudstack.CloudStackClient
	Err    AcsError
}

func (c CloudstackUtils) RetrieveID(keys ...string) (string, AcsError) {

	len := len(keys)

	if len < 3 {
		c.Err.err = fmt.Errorf("invalid number of keys passed to RetrieveID. Should be 3 at least")
		return "", c.Err
	}

	var name, value, zone string

	// There may be more input parameters in future, so reserve more space with switch
	switch len {
	default:
		name, value, zone = keys[0], keys[1], keys[2]
	}

	// If the supplied value isn't a ID, try to retrieve the ID ourselves
	if cloudstack.IsID(value) {
		return value, c.Err
	}

	c.Err.zone = zone
	c.Err.name = name
	c.Err.value = value

	var err error
	var id string

	log.Printf("[DEBUG] Retrieving ID of %s %s: %s", zone, name, value)

	// Ignore counts, since an error is returned if there is no exact match
	switch name {

	case "zone":
		id, _, err = c.Client.Zone.GetZoneID(value)

	case "network_offering":
		id, _, err = c.Client.NetworkOffering.GetNetworkOfferingID(value)

	case "vpc_offering":
		id, _, err = c.Client.VPC.GetVPCOfferingID(value)

	case "service_offering":
		id, _, err = c.Client.ServiceOffering.GetServiceOfferingID(value)

	case "disk_offering":
		id, _, err = c.Client.DiskOffering.GetDiskOfferingID(value)

	case "template_id":
		id, _, err = c.Client.Template.GetTemplateID(value, "executable", zone)

	case "project":
		id, _, err = c.Client.Project.GetProjectID(value)

	case "os_type":
		p := c.Client.GuestOS.NewListOsTypesParams()
		p.SetDescription(value)
		l, e := c.Client.GuestOS.ListOsTypes(p)
		if e != nil {
			err = e
			break
		}
		if l.Count == 1 {
			id = l.OsTypes[0].Id
			break
		}
		err = fmt.Errorf("could not find ID of OS Type: %s", value)

	default:
		c.Err.err = fmt.Errorf("unknown request: %s", name)
		return id, c.Err
	}

	if err != nil {
		c.Err.err = err
		return id, c.Err
	}

	return id, c.Err
}
