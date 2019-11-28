package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/network/v2/internet_gateways"
)

func dataSourceNetworkInternetGatewayV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkInternetGatewayV2Read,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internet_gateway_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"qos_option_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkInternetGatewayV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))

	listOpts := internet_gateways.ListOpts{}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("internet_gateway_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("internet_service_id"); ok {
		listOpts.InternetServiceID = v.(string)
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

	pages, err := internet_gateways.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve internet_gateways: %s", err)
	}

	allInternetGateways, err := internet_gateways.ExtractInternetGateways(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract internet_gateways: %s", err)
	}

	if len(allInternetGateways) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allInternetGateways) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	internetGateway := allInternetGateways[0]

	log.Printf("[DEBUG] Retrieved InternetGateway %s: %+v", internetGateway.ID, internetGateway)
	d.SetId(internetGateway.ID)

	d.Set("description", internetGateway.Description)
	d.Set("internet_service_id", internetGateway.InternetServiceID)
	d.Set("name", internetGateway.Name)
	d.Set("qos_option_id", internetGateway.QoSOptionID)
	d.Set("status", internetGateway.Status)
	d.Set("tenant_id", internetGateway.TenantID)
	d.Set("region", GetRegion(d, config))

	return nil
}
