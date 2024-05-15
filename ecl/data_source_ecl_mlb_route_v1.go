package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/routes"
)

func dataSourceMLBRouteV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBRouteV1Read,
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
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
			"configuration_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operation_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"load_balancer_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_hop_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}

	return result
}

func dataSourceMLBRouteV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := routes.ListOpts{}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("configuration_status"); ok {
		listOpts.ConfigurationStatus = v.(string)
	}

	if v, ok := d.GetOk("operation_status"); ok {
		listOpts.OperationStatus = v.(string)
	}

	if v, ok := d.GetOk("destination_cidr"); ok {
		listOpts.DestinationCidr = v.(string)
	}

	if v, ok := d.GetOk("next_hop_ip_address"); ok {
		listOpts.NextHopIPAddress = v.(string)
	}

	if v, ok := d.GetOk("load_balancer_id"); ok {
		listOpts.LoadBalancerID = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Retrieving ECL managed load balancer routes with options %+v", listOpts)

	pages, err := routes.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allRoutes, err := routes.ExtractRoutes(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer routes with options %+v: %s", listOpts, err)
	}

	if len(allRoutes) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allRoutes) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	route := allRoutes[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer route: %+v", route)

	d.SetId(route.ID)

	d.Set("name", route.Name)
	d.Set("description", route.Description)
	d.Set("tags", route.Tags)
	d.Set("configuration_status", route.ConfigurationStatus)
	d.Set("operation_status", route.OperationStatus)
	d.Set("destination_cidr", route.DestinationCidr)
	d.Set("load_balancer_id", route.LoadBalancerID)
	d.Set("tenant_id", route.TenantID)
	d.Set("next_hop_ip_address", route.NextHopIPAddress)

	return nil
}
