package main

import (
	"encoding/json"
	"github.com/apache/cloudstack-go/v2/cloudstack"
	"log"
)

func main() {
	cs := cloudstack.NewAsyncClient(ApiUrl, ApiKey, SecretKey, false)
	zone, _, err := cs.Zone.GetZoneByID(ZoneId)
	if err != nil {
		log.Fatalf("Failed to find zone: %v", err)
		return
	}

	if zone.Resourcedetails["resourceHAEnabled"] == "true" {
		log.Printf("High Availability enabled for zone: %s. Attempting to disable it", ZoneId)
		zoneHAParams := &cloudstack.DisableHAForZoneParams{}
		zoneHAParams.SetZoneid(zone.Id)
		resp, err := cs.Zone.DisableHAForZone(zoneHAParams)
		if err != nil {
			log.Fatalf("Failed to disable HA for zone due to: %v", err)
			return
		}
		log.Printf("HA for zone - disabled? : %v", resp.Success)
	}

	if zone.Resourcedetails["resourceHAEnabled"] == "false" {
		log.Printf("High Availability disabled for zone: %s. Attempting to enable it", ZoneId)
		zoneHAParams := &cloudstack.EnableHAForZoneParams{}
		zoneHAParams.SetZoneid(zone.Id)
		resp, err := cs.Zone.EnableHAForZone(zoneHAParams)
		if err != nil {
			log.Fatalf("Failed to enable HA for zone due to: %v", err)
			return
		}
		log.Printf("HA for zone - enabled? : %v", resp.Success)
	}

	p := &cloudstack.ListHostsParams{}
	p.SetZoneid(ZoneId)
	hosts, err := cs.Host.ListHosts(p)
	if err != nil {
		log.Fatalf("Failed to list all hosts in zone: %s due to: %v", zone.Name, err)
		return
	}
	for _, host := range hosts.Hosts {
		log.Printf("Host - Name: %s \t Host - HA Response: %v", host.Name, host.Hostha)
	}

	host := hosts.Hosts[0]
	log.Println(host.Resourcestate)
	if host.Resourcestate == "Enabled" {
		log.Printf("Prepare host %s for maintenance", host.Name)
		prepHostMnt := &cloudstack.PrepareHostForMaintenanceParams{}
		prepHostMnt.SetId(host.Id)
		resp, err := cs.Host.PrepareHostForMaintenance(prepHostMnt)
		if err != nil {
			log.Fatalf("Failed to prepare host %s for maintenance due to: %v", host.Name, err)
			return
		}
		b, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			log.Printf("%v", err)
			return
		}
		log.Printf("Host response : %v", string(b))
	}

	if host.Resourcestate == "Maintenance" {
		log.Printf("Cancelinng host for maintenance on host: %s", host.Name)
		cancelHostMnt := &cloudstack.CancelHostMaintenanceParams{}
		cancelHostMnt.SetId(host.Id)
		resp, err := cs.Host.CancelHostMaintenance(cancelHostMnt)
		if err != nil {
			log.Fatalf("Failed to cancel host for maintenance on %s due to: %v", host.Name, err)
			return
		}
		b, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			log.Printf("%v", err)
			return
		}
		log.Printf("Host response : %v", string(b))
	}
}
