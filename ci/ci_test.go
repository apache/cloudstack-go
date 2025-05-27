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

/*
Package citest contains comprehensive tests for the CloudStack Go SDK.

This file is used in the CI pipeline to test the CloudStack API client
and ensure that parameter handling works correctly across all supported
data types.

The tests cover all CloudStack API parameter types:
- BOOLEAN: Boolean values (true/false)
- DATE: Date strings in CloudStack format
- INTEGER: Regular integers (int)
- SHORT: Short integers (handled as int)
- LIST: Array/slice parameters ([]string)
- LONG: Long integers (int64)
- MAP: Key-value pairs (map[string]string)
- STRING: Text strings and UUIDs
- UUID: Universally unique identifiers (string format)

These tests validate that:
1. Parameters can be set and retrieved correctly
2. Type conversions work as expected
3. Parameter validation functions properly
4. The SDK handles all CloudStack API data types
*/

package citest

import (
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

// testResources holds IDs of created test resources for cleanup
type testResources struct {
	zoneID             string
	vmID1              string
	vmID2              string
	serviceOfferingID1 string
	serviceOfferingID2 string
	networkID          string
	templateID         string
	networkOfferingID  string
}

func TestCloudstackAPI(t *testing.T) {
	// Check if CloudStack credentials are provided, otherwise use localhost defaults
	apiURL := os.Getenv("CLOUDSTACK_API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080/client/api"
	}

	apiKey := os.Getenv("CLOUDSTACK_API_KEY")
	if apiKey == "" {
		t.Fatalf("CLOUDSTACK_API_KEY is not set")
	}

	secretKey := os.Getenv("CLOUDSTACK_SECRET_KEY")
	if secretKey == "" {
		t.Fatalf("CLOUDSTACK_SECRET_KEY is not set")
	}

	client := cloudstack.NewClient(apiURL, apiKey, secretKey, false) // Don't verify SSL for testing

	// Initialize test resources
	resources := &testResources{}

	// Set up test environment
	setupTestEnvironment(t, client, resources)

	// Run the actual tests
	t.Run("BasicTypes", func(t *testing.T) {
		testBasicTypes(t, client, resources)
	})

	// Clean up test resources we created
	t.Cleanup(func() {
		cleanupTestResources(t, client, resources)
	})
}

func setupTestEnvironment(t *testing.T, client *cloudstack.CloudStackClient, resources *testResources) {
	t.Helper()

	// Get zone ID
	resources.zoneID = getTestZone(t, client)

	// Get service offering IDs
	resources.serviceOfferingID1, resources.serviceOfferingID2 = getTestServiceOfferings(t, client)

	// Get network offerings to create a network
	resources.networkOfferingID = getTestNetworkOffering(t, client)

	// Create test network
	resources.networkID = createTestNetwork(t, client, "ci-test-network", resources.networkOfferingID, resources.zoneID)

	// Get template ID
	resources.templateID = getTestTemplate(t, client, resources.zoneID)

	// Create VMs for testing
	resources.vmID1 = createTestVM(t, client, "ci-test-vm-1", resources.serviceOfferingID1, resources.zoneID, resources.networkID, resources.templateID)
	resources.vmID2 = createTestVM(t, client, "ci-test-vm-2", resources.serviceOfferingID1, resources.zoneID, resources.networkID, resources.templateID)
}

func testBasicTypes(t *testing.T, client *cloudstack.CloudStackClient, resources *testResources) {
	// Test STRING parameters with actual API call
	t.Run("StringParameters", func(t *testing.T) {
		// Test with ListZones API which uses string parameters
		params := client.Zone.NewListZonesParams()
		params.SetKeyword("zone") // string parameter

		// Verify parameter was set correctly
		if keyword, ok := params.GetKeyword(); !ok || keyword != "zone" {
			t.Errorf("Expected keyword to be 'zone', got %s", keyword)
		}

		// Make actual API call with the parameters we set (not new params)
		resp, err := client.Zone.ListZones(params)
		if err != nil {
			t.Fatalf("ListZones API call with keyword='zone' failed: %v", err)
		}
		t.Logf("ListZones with keyword='zone' successful, found %d zones", resp.Count)

		// Test parameter reset functionality
		params.ResetKeyword()
		if _, ok := params.GetKeyword(); ok {
			t.Error("Expected keyword to be reset")
		}
	})

	// Test BOOLEAN parameters with actual API call
	t.Run("BooleanParameters", func(t *testing.T) {
		// Test with ListVirtualMachines API which has boolean parameters
		params := client.VirtualMachine.NewListVirtualMachinesParams()
		params.SetListall(true) // boolean parameter

		// Verify parameter was set correctly
		if listall, ok := params.GetListall(); !ok || !listall {
			t.Errorf("Expected listall to be true, got %v", listall)
		}

		// Make actual API call with the boolean parameter we set
		resp, err := client.VirtualMachine.ListVirtualMachines(params)
		if err != nil {
			t.Fatalf("ListVirtualMachines API call with listall=true failed: %v", err)
		}
		t.Logf("ListVirtualMachines with listall=true successful, found %d VMs", resp.Count)

		// Test setting to false and making another call
		params.SetListall(false)
		if listall, ok := params.GetListall(); !ok || listall {
			t.Errorf("Expected listall to be false, got %v", listall)
		}

		// Make API call with listall=false
		respFalse, err := client.VirtualMachine.ListVirtualMachines(params)
		if err != nil {
			t.Fatalf("ListVirtualMachines API call with listall=false failed: %v", err)
		}
		t.Logf("ListVirtualMachines with listall=false successful, found %d VMs", respFalse.Count)
	})

	// Test INTEGER parameters with actual API call
	t.Run("IntegerParameters", func(t *testing.T) {
		// Test with ListVirtualMachines API which has integer parameters
		params := client.VirtualMachine.NewListVirtualMachinesParams()
		params.SetPagesize(5) // integer parameter
		params.SetPage(1)     // required when pagesize is specified

		// Verify parameters were set correctly
		if pagesize, ok := params.GetPagesize(); !ok || pagesize != 5 {
			t.Errorf("Expected pagesize to be 5, got %d", pagesize)
		}
		if page, ok := params.GetPage(); !ok || page != 1 {
			t.Errorf("Expected page to be 1, got %d", page)
		}

		// Make actual API call with the integer parameters we set
		resp, err := client.VirtualMachine.ListVirtualMachines(params)
		if err != nil {
			t.Fatalf("ListVirtualMachines with pagesize=5 API call failed: %v", err)
		}
		t.Logf("ListVirtualMachines with pagesize=5, page=1 successful, found %d VMs", resp.Count)
		if resp.Count > 5 {
			t.Logf("Note: Response has more than 5 items, pagesize parameter might not be enforced by test server")
		}
	})

	// Test LONG (int64) parameters with actual API call
	t.Run("LongParameters", func(t *testing.T) {
		// Test with ChangeServiceForVirtualMachine API which has int64 parameters
		params := client.VirtualMachine.NewChangeServiceForVirtualMachineParams(resources.vmID1, resources.serviceOfferingID2)
		params.SetMaxiops(1000) // int64 parameter

		// Verify parameter was set correctly
		if maxiops, ok := params.GetMaxiops(); !ok || maxiops != 1000 {
			t.Errorf("Expected maxiops to be 1000, got %d", maxiops)
		}

		// Make actual API call with the int64 parameter we set
		// Note: This might fail if the VM is not in the right state, but that's okay for parameter testing
		resp, err := client.VirtualMachine.ChangeServiceForVirtualMachine(params)
		if err != nil {
			t.Fatalf("ChangeServiceForVirtualMachine with maxiops=1000 API call failed: %v", err)
		} else {
			t.Logf("ChangeServiceForVirtualMachine with maxiops=1000 successful, job ID: %s", resp.JobID)
		}
	})

	// Test DATE parameters with actual API call
	t.Run("DateParameters", func(t *testing.T) {
		// Test with ListUsageRecords API which has date parameters
		params := client.Usage.NewListUsageRecordsParams("2024-01-31", "2024-01-01")

		// Verify date parameters were set correctly during construction
		if startdate, ok := params.GetStartdate(); !ok || startdate != "2024-01-01" {
			t.Errorf("Expected startdate to be '2024-01-01', got %s", startdate)
		}
		if enddate, ok := params.GetEnddate(); !ok || enddate != "2024-01-31" {
			t.Errorf("Expected enddate to be '2024-01-31', got %s", enddate)
		}

		// Set additional date parameters
		params.SetStartdate("2024-01-01")
		params.SetEnddate("2024-01-31")

		// Verify parameters can be modified
		if startdate, ok := params.GetStartdate(); !ok || startdate != "2024-01-01" {
			t.Errorf("Expected modified startdate to be '2024-01-01', got %s", startdate)
		}

		// Make actual API call with the date parameters we set
		resp, err := client.Usage.ListUsageRecords(params)
		if err != nil {
			t.Fatalf("ListUsageRecords API call with date parameters failed: %v", err)
		}
		t.Logf("ListUsageRecords with startdate='2024-01-01' and enddate='2024-01-31' successful, found %d usage records", resp.Count)
	})

	// Test LIST parameters with actual API call
	t.Run("ListParameters", func(t *testing.T) {
		// Test with ListVirtualMachines API which has list parameters
		params := client.VirtualMachine.NewListVirtualMachinesParams()
		vmIDList := []string{resources.vmID1, resources.vmID2}
		params.SetIds(vmIDList) // []string parameter

		// Verify parameter was set correctly
		if ids, ok := params.GetIds(); !ok || len(ids) != 2 || ids[0] != resources.vmID1 || ids[1] != resources.vmID2 {
			t.Errorf("Expected ids to be %v, got %v", vmIDList, ids)
		}

		// Make actual API call with the list parameter we set
		resp, err := client.VirtualMachine.ListVirtualMachines(params)
		if err != nil {
			t.Fatalf("ListVirtualMachines API call with ids=%v failed: %v", vmIDList, err)
		}
		t.Logf("ListVirtualMachines with ids=%v successful, found %d VMs", vmIDList, resp.Count)
		if resp.Count != 2 {
			t.Errorf("Expected 2 VMs, got %d", resp.Count)
		}
		// Verify returned VMs are in our expected list
		if resp.Count >= 2 {
			foundIDs := make([]string, resp.Count)
			for i, vm := range resp.VirtualMachines {
				foundIDs[i] = vm.Id
			}
			for _, expectedID := range vmIDList {
				if !slices.Contains(foundIDs, expectedID) {
					t.Errorf("Expected VM ID %s not found in response", expectedID)
				}
			}
		}
	})

	// Test MAP parameters with actual API call
	t.Run("MapParameters", func(t *testing.T) {
		// Test with AddNicToVirtualMachine API which has map parameters
		params := client.VirtualMachine.NewAddNicToVirtualMachineParams(resources.networkID, resources.vmID1)
		dhcpOptions := map[string]string{
			"option1": "value1",
			"option2": "value2",
		}
		params.SetDhcpoptions(dhcpOptions) // map[string]string parameter

		// Verify parameter was set correctly
		if options, ok := params.GetDhcpoptions(); !ok || len(options) != 2 || options["option1"] != "value1" {
			t.Errorf("Expected dhcpoptions to be %v, got %v", dhcpOptions, options)
		}

		// Make actual API call with the map parameter we set
		resp, err := client.VirtualMachine.AddNicToVirtualMachine(params)
		if err != nil {
			t.Fatalf("AddNicToVirtualMachine with dhcpoptions=%v API call failed: %v", dhcpOptions, err)
		} else {
			t.Logf("AddNicToVirtualMachine with dhcpoptions=%v successful, job ID: %s", dhcpOptions, resp.JobID)
		}
	})

	// Test UUID parameters with actual API call
	t.Run("UUIDParameters", func(t *testing.T) {
		// Test with ListServiceOfferings API using zoneid UUID parameter
		params := client.ServiceOffering.NewListServiceOfferingsParams()
		params.SetZoneid(resources.zoneID) // UUID parameter

		// Verify UUID parameter was set correctly
		if zoneid, ok := params.GetZoneid(); !ok || zoneid != resources.zoneID {
			t.Errorf("Expected zoneid to be %s, got %s", resources.zoneID, zoneid)
		}

		// Make actual API call with the UUID parameter we set
		resp, err := client.ServiceOffering.ListServiceOfferings(params)
		if err != nil {
			t.Fatalf("ListServiceOfferings with zoneid=%s API call failed: %v", resources.zoneID, err)
		}
		t.Logf("ListServiceOfferings with zoneid=%s successful, found %d service offerings", resources.zoneID, resp.Count)
		if resp.Count > 0 {
			// Verify that returned IDs are valid UUIDs
			firstOffering := resp.ServiceOfferings[0]
			if !cloudstack.IsID(firstOffering.Id) {
				t.Errorf("Service offering ID %s is not a valid UUID format", firstOffering.Id)
			} else {
				t.Logf("Service offering ID %s is valid UUID format", firstOffering.Id)
			}
		}
	})

	// Test SHORT parameters with actual API call
	t.Run("ShortParameters", func(t *testing.T) {
		// Test with RemoveRawUsageRecords API which has interval parameter
		params := client.Usage.NewRemoveRawUsageRecordsParams(30)

		// Verify parameter was set correctly during construction
		if interval, ok := params.GetInterval(); !ok || interval != 30 {
			t.Errorf("Expected interval to be 30, got %d", interval)
		}

		// Modify the parameter to a safer value
		params.SetInterval(365) // Remove usage records older than 1 year

		// Verify parameter was modified correctly
		if interval, ok := params.GetInterval(); !ok || interval != 365 {
			t.Errorf("Expected modified interval to be 365, got %d", interval)
		}

		// Make actual API call with the short parameter we set
		resp, err := client.Usage.RemoveRawUsageRecords(params)
		if err != nil {
			// Log the error but don't fail the test - might not have permission
			t.Logf("RemoveRawUsageRecords with interval=365 API call failed (might not have permission): %v", err)
			// Verify it's a CloudStack API error, not a parameter handling error
			if !strings.Contains(err.Error(), "CloudStack API error") {
				t.Fatalf("Unexpected error type - might be parameter handling issue: %v", err)
			}
		} else {
			t.Logf("RemoveRawUsageRecords with interval=365 successful: %s", resp.Displaytext)
		}
	})
}

/*
	Helper functions to get test resources
*/

func getTestZone(t *testing.T, client *cloudstack.CloudStackClient) string {
	t.Helper()

	zoneParams := client.Zone.NewListZonesParams()
	zones, err := client.Zone.ListZones(zoneParams)
	if err != nil || zones.Count == 0 {
		t.Fatalf("Failed to get zones for testing: %v", err)
	}
	zoneID := zones.Zones[0].Id
	t.Logf("Using zone ID: %s", zoneID)
	return zoneID
}

func getTestServiceOfferings(t *testing.T, client *cloudstack.CloudStackClient) (string, string) {
	t.Helper()

	serviceParams := client.ServiceOffering.NewListServiceOfferingsParams()
	services, err := client.ServiceOffering.ListServiceOfferings(serviceParams)
	if err != nil || services.Count < 2 {
		t.Fatalf("Failed to get at least 2 service offerings for testing: %v (found %d)", err, services.Count)
	}
	serviceOfferingID1 := services.ServiceOfferings[0].Id
	serviceOfferingID2 := services.ServiceOfferings[1].Id
	t.Logf("Using service offering IDs: %s, %s", serviceOfferingID1, serviceOfferingID2)
	return serviceOfferingID1, serviceOfferingID2
}

func getTestNetworkOffering(t *testing.T, client *cloudstack.CloudStackClient) string {
	t.Helper()

	netOfferingParams := client.NetworkOffering.NewListNetworkOfferingsParams()
	netOfferingParams.SetState("Enabled")
	netOfferings, err := client.NetworkOffering.ListNetworkOfferings(netOfferingParams)
	if err != nil || netOfferings.Count == 0 {
		t.Fatalf("Failed to get network offerings for network creation: %v", err)
	}

	// Find a suitable network offering (preferably isolated)
	var networkOfferingID string
	for _, offering := range netOfferings.NetworkOfferings {
		if offering.Guestiptype == "Isolated" && !offering.Specifyvlan {
			networkOfferingID = offering.Id
			break
		}
	}
	if networkOfferingID == "" {
		// Use first available offering if no isolated found
		networkOfferingID = netOfferings.NetworkOfferings[0].Id
	}
	t.Logf("Using network offering ID: %s", networkOfferingID)
	return networkOfferingID
}

func getTestTemplate(t *testing.T, client *cloudstack.CloudStackClient, zoneID string) string {
	t.Helper()

	templateParams := client.Template.NewListTemplatesParams("executable")
	templateParams.SetZoneid(zoneID)
	templates, err := client.Template.ListTemplates(templateParams)
	if err != nil || templates.Count == 0 {
		t.Fatalf("Failed to get templates for VM creation: %v", err)
	}
	templateID := templates.Templates[0].Id
	t.Logf("Using template ID: %s", templateID)
	return templateID
}

/*
	Helper functions to create test resources
*/

func createTestNetwork(t *testing.T, client *cloudstack.CloudStackClient, name, offeringID, zoneID string) string {
	t.Helper()

	createNetParams := client.Network.NewCreateNetworkParams(name, offeringID, zoneID)
	createNetParams.SetDisplaytext(name)
	createNetResp, err := client.Network.CreateNetwork(createNetParams)
	if err != nil {
		t.Fatalf("Failed to create test network: %v", err)
	}
	t.Logf("Created test network with ID: %s", createNetResp.Id)
	return createNetResp.Id
}

func createTestVM(t *testing.T, client *cloudstack.CloudStackClient, name, offeringID, zoneID, networkID, templateID string) string {
	t.Helper()

	deployParams := client.VirtualMachine.NewDeployVirtualMachineParams(offeringID, templateID, zoneID)
	deployParams.SetName(name)
	deployParams.SetDisplayname(name)
	deployParams.SetNetworkids([]string{networkID})
	deployParams.SetStartvm(false) // Don't start VM to avoid resource consumption
	deployResp, err := client.VirtualMachine.DeployVirtualMachine(deployParams)
	if err != nil {
		t.Fatalf("Failed to create test VM %s: %v", name, err)
	}
	t.Logf("Created test VM %s with ID: %s", name, deployResp.Id)
	return deployResp.Id
}

/*
	Helper functions to delete test resources
*/

func cleanupTestResources(t *testing.T, client *cloudstack.CloudStackClient, resources *testResources) {
	t.Helper()

	// Clean up VMs first (order matters for dependencies)
	if resources.vmID1 != "" {
		expungeVM(t, client, resources.vmID1)
	}
	if resources.vmID2 != "" {
		expungeVM(t, client, resources.vmID2)
	}

	// Clean up network after VMs are destroyed
	if resources.networkID != "" {
		deleteNetwork(t, client, resources.networkID)
	}
}

func deleteNetwork(t *testing.T, client *cloudstack.CloudStackClient, networkID string) {
	t.Helper()

	params := client.Network.NewDeleteNetworkParams(networkID)
	params.SetForced(true)
	_, err := client.Network.DeleteNetwork(params)
	if err != nil {
		t.Logf("Failed to delete test network %s: %v", networkID, err)
	} else {
		t.Logf("Successfully deleted test network %s", networkID)
	}
}

func expungeVM(t *testing.T, client *cloudstack.CloudStackClient, vmID string) {
	t.Helper()

	params := client.VirtualMachine.NewDestroyVirtualMachineParams(vmID)
	params.SetExpunge(true)
	_, err := client.VirtualMachine.DestroyVirtualMachine(params)
	if err != nil {
		t.Logf("Failed to destroy test VM %s: %v", vmID, err)
	} else {
		t.Logf("Successfully destroyed test VM %s", vmID)
	}
}
