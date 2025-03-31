package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v4/ecl/network/v2/common_function_gateways"
)

func dataSourceNetworkCommonFunctionGatewayV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkCommonFunctionGatewayV2Read,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"common_function_pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"subnet_id": &schema.Schema{
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

func dataSourceNetworkCommonFunctionGatewayV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))

	listOpts := common_function_gateways.ListOpts{}

	if v, ok := d.GetOk("common_function_pool_id"); ok {
		listOpts.CommonFunctionPoolID = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("network_id"); ok {
		listOpts.NetworkID = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		listOpts.SubnetID = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	pages, err := common_function_gateways.List(networkClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allCommonFunctionGateways, err := common_function_gateways.ExtractCommonFunctionGateways(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve common function gateways: %s", err)
	}

	if len(allCommonFunctionGateways) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allCommonFunctionGateways) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	cfgw := allCommonFunctionGateways[0]

	log.Printf("[DEBUG] Retrieved Common Function Gateway %s: %+v", cfgw.ID, cfgw)
	d.SetId(cfgw.ID)

	d.Set("name", cfgw.Name)
	d.Set("description", cfgw.Description)
	d.Set("common_function_pool_id", cfgw.CommonFunctionPoolID)
	d.Set("network_id", cfgw.NetworkID)
	d.Set("subnet_id", cfgw.SubnetID)
	d.Set("tenant_id", cfgw.TenantID)

	return nil
}
