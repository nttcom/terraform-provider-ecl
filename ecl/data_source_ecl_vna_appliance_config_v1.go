package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/nttcom/eclcloud/v3/ecl/vna/v1/appliances"
)

func dataSourceVNAApplianceConfigV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNAApplianceConfigV1Read,

		Schema: map[string]*schema.Schema{
			"virtual_network_appliance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"format": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{"set", "text", "a10-json"}, false,
				),
			},

			"data": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceVNAApplianceConfigV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	id := d.Get("virtual_network_appliance_id").(string)
	opts := appliances.ExportConfigOpts{
		Format: d.Get("format").(string),
	}

	log.Printf("[DEBUG] Exporting configuration of Virtual Network Appliance %s with options: %#v", id, opts)

	configuration, err := appliances.ExportConfig(vnaClient, id, opts).Extract()
	if err != nil {
		return fmt.Errorf("Error exporting configuration of Virtual Network Appliance %s: %s", id, err)
	}

	log.Printf("[DEBUG] Exported configuration of Virtual Network Appliance %s (format: %s)", id, configuration.Format)

	d.SetId(fmt.Sprintf("%s/%s", id, configuration.Format))
	d.Set("format", configuration.Format)
	d.Set("data", configuration.Data)

	return nil
}
