package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/listeners"
)

func dataSourceMLBListenerV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBListenerV1Read,
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
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}

	return result
}

func dataSourceMLBListenerV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := listeners.ListOpts{}

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

	if v, ok := d.GetOk("ip_address"); ok {
		listOpts.IPAddress = v.(string)
	}

	if v, ok := d.GetOk("port"); ok {
		listOpts.Port = v.(int)
	}

	if v, ok := d.GetOk("protocol"); ok {
		listOpts.Protocol = v.(string)
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

	log.Printf("[DEBUG] Retrieving ECL managed load balancer listeners with options %+v", listOpts)

	pages, err := listeners.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allListeners, err := listeners.ExtractListeners(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer listeners with options %+v: %s", listOpts, err)
	}

	if len(allListeners) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allListeners) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	listener := allListeners[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer listener: %+v", listener)

	d.SetId(listener.ID)

	d.Set("name", listener.Name)
	d.Set("description", listener.Description)
	d.Set("tags", listener.Tags)
	d.Set("configuration_status", listener.ConfigurationStatus)
	d.Set("operation_status", listener.OperationStatus)
	d.Set("load_balancer_id", listener.LoadBalancerID)
	d.Set("tenant_id", listener.TenantID)
	d.Set("ip_address", listener.IPAddress)
	d.Set("port", listener.Port)
	d.Set("protocol", listener.Protocol)

	return nil
}
