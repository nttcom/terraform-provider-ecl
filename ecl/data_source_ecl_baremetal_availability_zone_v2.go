package ecl

import (
	"fmt"
	"log"

	"github.com/nttcom/eclcloud/v3/ecl/baremetal/v2/availabilityzones"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceBaremetalAvailabilityZoneV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBaremetalAvailabilityZoneV2Read,

		Schema: map[string]*schema.Schema{
			"zone_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

// dataSourceBaremetalAvailabilityZoneV2Read performs the availability zone lookup.
func dataSourceBaremetalAvailabilityZoneV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	allPages, err := availabilityzones.List(baremetalClient).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to query availability zones: %s", err)
	}

	allAvailabilityZones, err := availabilityzones.ExtractAvailabilityZones(allPages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve availability zones: %s", err)
	}

	// Loop through all availability zones to find a more specific one.
	if len(allAvailabilityZones) > 1 {
		var filteredAvailabilityZones []availabilityzones.AvailabilityZone
		for _, availabilityzone := range allAvailabilityZones {
			if v := d.Get("zone_name").(string); v != "" {
				if availabilityzone.ZoneName != v {
					continue
				}
			}

			filteredAvailabilityZones = append(filteredAvailabilityZones, availabilityzone)
		}

		allAvailabilityZones = filteredAvailabilityZones
	}

	if len(allAvailabilityZones) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allAvailabilityZones) > 1 {
		log.Printf("[DEBUG] Multiple results found: %#v", allAvailabilityZones)
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria")
	}

	availabilityzone := allAvailabilityZones[0]
	log.Printf("[DEBUG] Single AvailabilityZone found: %s", availabilityzone.ZoneName)
	return dataSourceBaremetalAvailabilityZoneV2Attributes(d, &availabilityzone)
}

// dataSourceBaremetalAvailabilityZoneV2Attributes populates the fields of an AvailabilityZone resource.
func dataSourceBaremetalAvailabilityZoneV2Attributes(d *schema.ResourceData, availabilityzone *availabilityzones.AvailabilityZone) error {
	log.Printf("[DEBUG] ecl_baremetal_availabilityzone_v2 details: %#v", availabilityzone)

	d.SetId(availabilityzone.ZoneName)
	d.Set("zone_name", availabilityzone.ZoneName)

	return nil
}
