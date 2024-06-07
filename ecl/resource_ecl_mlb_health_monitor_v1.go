package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/health_monitors"
)

func resourceMLBHealthMonitorV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceMLBHealthMonitorV1Create,
		Read:   resourceMLBHealthMonitorV1Read,
		Update: resourceMLBHealthMonitorV1Update,
		Delete: resourceMLBHealthMonitorV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"load_balancer_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}

	return result
}

func resourceMLBHealthMonitorV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOpts := health_monitors.CreateOpts{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Tags:           d.Get("tags").(map[string]interface{}),
		Port:           d.Get("port").(int),
		Protocol:       d.Get("protocol").(string),
		Interval:       d.Get("interval").(int),
		Retry:          d.Get("retry").(int),
		Timeout:        d.Get("timeout").(int),
		Path:           d.Get("path").(string),
		HttpStatusCode: d.Get("http_status_code").(string),
		LoadBalancerID: d.Get("load_balancer_id").(string),
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Creating ECL managed load balancer health monitor with options %+v", createOpts)

	healthMonitor, err := health_monitors.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer health monitor with options %+v: %s", createOpts, err)
	}

	d.SetId(healthMonitor.ID)
	log.Printf("[INFO] ECL managed load balancer health monitor ID: %s", healthMonitor.ID)

	return resourceMLBHealthMonitorV1Read(d, meta)
}

func resourceMLBHealthMonitorV1Show(d *schema.ResourceData, client *eclcloud.ServiceClient, changes bool) (*health_monitors.HealthMonitor, error) {
	var healthMonitor health_monitors.HealthMonitor

	showOpts := health_monitors.ShowOpts{Changes: changes}
	err := health_monitors.Show(client, d.Id(), showOpts).ExtractInto(&healthMonitor)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer health monitor (%s): %s", d.Id(), err)
	}

	return &healthMonitor, nil
}

func resourceMLBHealthMonitorV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	healthMonitor, err := resourceMLBHealthMonitorV1Show(d, managedLoadBalancerClient, true)
	if err != nil {
		return CheckDeleted(d, err, "health_monitor")
	}

	log.Printf("[DEBUG] Retrieved ECL managed load balancer health monitor (%s): %+v", d.Id(), healthMonitor)

	if healthMonitor.ConfigurationStatus == "ACTIVE" {
		d.Set("port", healthMonitor.Port)
		d.Set("protocol", healthMonitor.Protocol)
		d.Set("interval", healthMonitor.Interval)
		d.Set("retry", healthMonitor.Retry)
		d.Set("timeout", healthMonitor.Timeout)
		d.Set("path", healthMonitor.Path)
		d.Set("http_status_code", healthMonitor.HttpStatusCode)
	} else if healthMonitor.ConfigurationStatus == "CREATE_STAGED" {
		d.Set("port", healthMonitor.Staged.Port)
		d.Set("protocol", healthMonitor.Staged.Protocol)
		d.Set("interval", healthMonitor.Staged.Interval)
		d.Set("retry", healthMonitor.Staged.Retry)
		d.Set("timeout", healthMonitor.Staged.Timeout)
		d.Set("path", healthMonitor.Staged.Path)
		d.Set("http_status_code", healthMonitor.Staged.HttpStatusCode)
	} else if healthMonitor.ConfigurationStatus == "UPDATE_STAGED" {
		d.Set("port", ternary(healthMonitor.Staged.Port == 0, healthMonitor.Port, healthMonitor.Staged.Port))
		d.Set("protocol", ternary(healthMonitor.Staged.Protocol == "", healthMonitor.Protocol, healthMonitor.Staged.Protocol))
		d.Set("interval", ternary(healthMonitor.Staged.Interval == 0, healthMonitor.Interval, healthMonitor.Staged.Interval))
		d.Set("retry", ternary(healthMonitor.Staged.Retry == 0, healthMonitor.Retry, healthMonitor.Staged.Retry))
		d.Set("timeout", ternary(healthMonitor.Staged.Timeout == 0, healthMonitor.Timeout, healthMonitor.Staged.Timeout))
		d.Set("path", ternary(healthMonitor.Staged.Path == "", healthMonitor.Path, healthMonitor.Staged.Path))
		d.Set("http_status_code", ternary(healthMonitor.Staged.HttpStatusCode == "", healthMonitor.HttpStatusCode, healthMonitor.Staged.HttpStatusCode))
	} else if healthMonitor.ConfigurationStatus == "DELETE_STAGED" {
		d.SetId("")
		return nil
	}

	d.Set("name", healthMonitor.Name)
	d.Set("description", healthMonitor.Description)
	d.Set("tags", healthMonitor.Tags)
	d.Set("load_balancer_id", healthMonitor.LoadBalancerID)
	d.Set("tenant_id", healthMonitor.TenantID)

	return nil
}

func resourceMLBHealthMonitorV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Start updating attributes of ECL managed load balancer health monitor ...")

	err = resourceMLBHealthMonitorV1UpdateAttributes(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating attributes of ECL managed load balancer health monitor: %s", err)
	}

	log.Printf("[DEBUG] Start updating configurations of ECL managed load balancer health monitor ...")

	err = resourceMLBHealthMonitorV1UpdateConfigurations(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating configurations of ECL managed load balancer health monitor: %s", err)
	}

	return resourceMLBHealthMonitorV1Read(d, meta)
}

func resourceMLBHealthMonitorV1UpdateAttributes(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isAttributesUpdated bool
	var updateOpts health_monitors.UpdateOpts

	if d.HasChange("name") {
		isAttributesUpdated = true
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		isAttributesUpdated = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("tags") {
		isAttributesUpdated = true
		tags := d.Get("tags").(map[string]interface{})
		updateOpts.Tags = &tags
	}

	if isAttributesUpdated {
		log.Printf("[DEBUG] Updating ECL managed load balancer health monitor attributes (%s) with options %+v", d.Id(), updateOpts)

		_, err := health_monitors.Update(client, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL managed load balancer health monitor attributes (%s) with options %+v: %s", d.Id(), updateOpts, err)
		}
	}

	return nil
}

func resourceMLBHealthMonitorV1UpdateConfigurations(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isConfigurationsUpdated bool

	healthMonitor, err := resourceMLBHealthMonitorV1Show(d, client, false)
	if err != nil {
		return err
	}

	if healthMonitor.ConfigurationStatus == "ACTIVE" {
		var createStagedOpts health_monitors.CreateStagedOpts

		if d.HasChange("port") {
			isConfigurationsUpdated = true
			createStagedOpts.Port = d.Get("port").(int)
		}

		if d.HasChange("protocol") {
			isConfigurationsUpdated = true
			createStagedOpts.Protocol = d.Get("protocol").(string)
		}

		if d.HasChange("interval") {
			isConfigurationsUpdated = true
			createStagedOpts.Interval = d.Get("interval").(int)
		}

		if d.HasChange("retry") {
			isConfigurationsUpdated = true
			createStagedOpts.Retry = d.Get("retry").(int)
		}

		if d.HasChange("timeout") {
			isConfigurationsUpdated = true
			createStagedOpts.Timeout = d.Get("timeout").(int)
		}

		if d.HasChange("path") {
			isConfigurationsUpdated = true
			createStagedOpts.Path = d.Get("path").(string)
		}

		if d.HasChange("http_status_code") {
			isConfigurationsUpdated = true
			createStagedOpts.HttpStatusCode = d.Get("http_status_code").(string)
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer health monitor configurations (%s) with options %+v", d.Id(), createStagedOpts)

			_, err := health_monitors.CreateStaged(client, d.Id(), createStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer health monitor configurations (%s) with options %+v: %s", d.Id(), createStagedOpts, err)
			}
		}
	} else {
		var updateStagedOpts health_monitors.UpdateStagedOpts

		if d.HasChange("port") {
			isConfigurationsUpdated = true
			port := d.Get("port").(int)
			updateStagedOpts.Port = &port
		}

		if d.HasChange("protocol") {
			isConfigurationsUpdated = true
			protocol := d.Get("protocol").(string)
			updateStagedOpts.Protocol = &protocol
		}

		if d.HasChange("interval") {
			isConfigurationsUpdated = true
			interval := d.Get("interval").(int)
			updateStagedOpts.Interval = &interval
		}

		if d.HasChange("retry") {
			isConfigurationsUpdated = true
			retry := d.Get("retry").(int)
			updateStagedOpts.Retry = &retry
		}

		if d.HasChange("timeout") {
			isConfigurationsUpdated = true
			timeout := d.Get("timeout").(int)
			updateStagedOpts.Timeout = &timeout
		}

		if d.HasChange("path") {
			isConfigurationsUpdated = true
			path := d.Get("path").(string)
			updateStagedOpts.Path = &path
		}

		if d.HasChange("http_status_code") {
			isConfigurationsUpdated = true
			httpStatusCode := d.Get("http_status_code").(string)
			updateStagedOpts.HttpStatusCode = &httpStatusCode
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer health monitor configurations (%s) with options %+v", d.Id(), updateStagedOpts)

			_, err := health_monitors.UpdateStaged(client, d.Id(), updateStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer health monitor configurations (%s) with options %+v: %s", d.Id(), updateStagedOpts, err)
			}
		}
	}

	return nil
}

func resourceMLBHealthMonitorV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Deleting ECL managed load balancer health monitor: %s", d.Id())

	err = health_monitors.Delete(managedLoadBalancerClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer health monitor (%s)", d.Id())
			return nil
		}

		healthMonitor, err := resourceMLBHealthMonitorV1Show(d, managedLoadBalancerClient, false)
		if err != nil {
			return err
		}
		if healthMonitor.ConfigurationStatus == "DELETE_STAGED" {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer health monitor (%s)", d.Id())
			return nil
		}

		return fmt.Errorf("Error deleting ECL managed load balancer health monitor: %s", err)
	}

	return nil
}
