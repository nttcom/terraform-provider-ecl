package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/ecl/network/v2/fic_gateways"
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
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	listOpts := fic_gateways.ListOpts{}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("fic_service_id"); ok {
		listOpts.FICServiceID = v.(string)
	}

	if v, ok := d.GetOk("fic_gateway_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("qos_option_id"); ok {
		listOpts.QoSOptionID = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	pages, err := fic_gateways.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("unable to retrieve fic_gateways: %w", err)
	}

	allFICGateways, err := fic_gateways.ExtractFICGateways(pages)
	if err != nil {
		return fmt.Errorf("unable to extract fic_gateways: %w", err)
	}

	if len(allFICGateways) < 1 {
		return fmt.Errorf("Your query returned no results." +
			" Please change your search criteria and try again.")
	}

	if len(allFICGateways) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	FICGateway := allFICGateways[0]

	log.Printf("[DEBUG] Retrieved FICGateway %s: %+v", FICGateway.ID, FICGateway)
	d.SetId(FICGateway.ID)

	d.Set("description", FICGateway.Description)
	d.Set("fic_service_id", FICGateway.FICServiceID)
	d.Set("fic_gateway_id", FICGateway.ID)
	d.Set("name", FICGateway.Name)
	d.Set("qos_option_id", FICGateway.QoSOptionID)
	d.Set("status", FICGateway.Status)
	d.Set("tenant_id", FICGateway.TenantID)

	return nil
}
