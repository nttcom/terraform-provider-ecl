package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/health_monitors"
)

func dataSourceMLBHealthMonitorV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBHealthMonitorV1Read,
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
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}

	return result
}

func dataSourceMLBHealthMonitorV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := health_monitors.ListOpts{}

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

	if v, ok := d.GetOk("port"); ok {
		listOpts.Port = v.(int)
	}

	if v, ok := d.GetOk("protocol"); ok {
		listOpts.Protocol = v.(string)
	}

	if v, ok := d.GetOk("interval"); ok {
		listOpts.Interval = v.(int)
	}

	if v, ok := d.GetOk("retry"); ok {
		listOpts.Retry = v.(int)
	}

	if v, ok := d.GetOk("timeout"); ok {
		listOpts.Timeout = v.(int)
	}

	if v, ok := d.GetOk("path"); ok {
		listOpts.Path = v.(string)
	}

	if v, ok := d.GetOk("http_status_code"); ok {
		listOpts.HttpStatusCode = v.(string)
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

	log.Printf("[DEBUG] Retrieving ECL managed load balancer health monitors with options %+v", listOpts)

	pages, err := health_monitors.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allHealthMonitors, err := health_monitors.ExtractHealthMonitors(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer health monitors with options %+v: %s", listOpts, err)
	}

	if len(allHealthMonitors) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allHealthMonitors) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	healthMonitor := allHealthMonitors[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer health monitor: %+v", healthMonitor)

	d.SetId(healthMonitor.ID)

	d.Set("name", healthMonitor.Name)
	d.Set("description", healthMonitor.Description)
	d.Set("tags", healthMonitor.Tags)
	d.Set("configuration_status", healthMonitor.ConfigurationStatus)
	d.Set("operation_status", healthMonitor.OperationStatus)
	d.Set("load_balancer_id", healthMonitor.LoadBalancerID)
	d.Set("tenant_id", healthMonitor.TenantID)
	d.Set("port", healthMonitor.Port)
	d.Set("protocol", healthMonitor.Protocol)
	d.Set("interval", healthMonitor.Interval)
	d.Set("retry", healthMonitor.Retry)
	d.Set("timeout", healthMonitor.Timeout)
	d.Set("path", healthMonitor.Path)
	d.Set("http_status_code", healthMonitor.HttpStatusCode)

	return nil
}
