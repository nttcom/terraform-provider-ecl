package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/fic_gateways"
)

func dataSourceNetworkFICGatewayV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkFICGatewayV2Read,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fic_service_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fic_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"qos_option_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkFICGatewayV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	var opts fic_gateways.ListOpts

	if v, ok := d.GetOk("description"); ok {
		opts.Description = v.(string)
	}

	if v, ok := d.GetOk("fic_service_id"); ok {
		opts.FICServiceID = v.(string)
	}

	if v, ok := d.GetOk("fic_gateway_id"); ok {
		opts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		opts.Name = v.(string)
	}

	if v, ok := d.GetOk("qos_option_id"); ok {
		opts.QoSOptionID = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		opts.Status = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		opts.TenantID = v.(string)
	}

	pages, err := fic_gateways.List(client, opts).AllPages()
	if err != nil {
		return fmt.Errorf("unable to retrieve fic_gateways: %w", err)
	}

	gws, err := fic_gateways.ExtractFICGateways(pages)
	if err != nil {
		return fmt.Errorf("unable to extract fic_gateways: %w", err)
	}

	if len(gws) < 1 {
		return fmt.Errorf("your query returned no results." +
			" please change your search criteria and try again")
	}

	if len(gws) > 1 {
		return fmt.Errorf("your query returned more than one result." +
			" please try a more specific search criteria")
	}

	gw := gws[0]

	log.Printf("[DEBUG] Retrieved FICGateway %s: %+v", gw.ID, gw)
	d.SetId(gw.ID)

	d.Set("description", gw.Description)
	d.Set("fic_service_id", gw.FICServiceID)
	d.Set("fic_gateway_id", gw.ID)
	d.Set("name", gw.Name)
	d.Set("qos_option_id", gw.QoSOptionID)
	d.Set("status", gw.Status)
	d.Set("tenant_id", gw.TenantID)

	return nil
}
