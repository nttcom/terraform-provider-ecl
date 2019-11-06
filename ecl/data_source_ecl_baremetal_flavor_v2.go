package ecl

import (
	"fmt"
	"log"

	"github.com/nttcom/eclcloud/ecl/baremetal/v2/flavors"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceBaremetalFlavorV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBaremetalFlavorV2Read,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"ram": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"vcpus": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"disk": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

// dataSourceBaremetalFlavorV2Read performs the flavor lookup.
func dataSourceBaremetalFlavorV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	allPages, err := flavors.List(baremetalClient, nil).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to query flavors: %s", err)
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve flavors: %s", err)
	}

	// Loop through all flavors to find a more specific one.
	if len(allFlavors) > 1 {
		var filteredFlavors []flavors.Flavor
		for _, flavor := range allFlavors {
			if v := d.Get("name").(string); v != "" {
				if flavor.Name != v {
					continue
				}
			}

			if v, ok := d.GetOk("ram"); ok {
				if flavor.RAM != v.(int) {
					continue
				}
			}

			if v, ok := d.GetOk("vcpus"); ok {
				if flavor.VCPUs != v.(int) {
					continue
				}
			}

			if v, ok := d.GetOk("disk"); ok {
				if flavor.Disk != v.(int) {
					continue
				}
			}

			filteredFlavors = append(filteredFlavors, flavor)
		}

		allFlavors = filteredFlavors
	}

	if len(allFlavors) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allFlavors) > 1 {
		log.Printf("[DEBUG] Multiple results found: %#v", allFlavors)
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria")
	}

	flavor := allFlavors[0]
	log.Printf("[DEBUG] Single Flavor found: %s", flavor.ID)
	return dataSourceBaremetalFlavorV2Attributes(d, &flavor)
}

// dataSourceBaremetalFlavorV2Attributes populates the fields of an Flavor resource.
func dataSourceBaremetalFlavorV2Attributes(d *schema.ResourceData, flavor *flavors.Flavor) error {
	log.Printf("[DEBUG] ecl_baremetal_flavor_v2 details: %#v", flavor)

	d.SetId(flavor.ID)
	d.Set("name", flavor.Name)
	d.Set("disk", flavor.Disk)
	d.Set("ram", flavor.RAM)
	d.Set("vcpus", flavor.VCPUs)

	return nil
}
