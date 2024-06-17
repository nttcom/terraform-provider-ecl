package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/listeners"
)

func resourceMLBListenerV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceMLBListenerV1Create,
		Read:   resourceMLBListenerV1Read,
		Update: resourceMLBListenerV1Update,
		Delete: resourceMLBListenerV1Delete,
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
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func resourceMLBListenerV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOpts := listeners.CreateOpts{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Tags:           d.Get("tags").(map[string]interface{}),
		IPAddress:      d.Get("ip_address").(string),
		Port:           d.Get("port").(int),
		Protocol:       d.Get("protocol").(string),
		LoadBalancerID: d.Get("load_balancer_id").(string),
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Creating ECL managed load balancer listener with options %+v", createOpts)

	listener, err := listeners.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer listener with options %+v: %s", createOpts, err)
	}

	d.SetId(listener.ID)
	log.Printf("[INFO] ECL managed load balancer listener ID: %s", listener.ID)

	return resourceMLBListenerV1Read(d, meta)
}

func resourceMLBListenerV1Show(d *schema.ResourceData, client *eclcloud.ServiceClient, changes bool) (*listeners.Listener, error) {
	var listener listeners.Listener

	showOpts := listeners.ShowOpts{Changes: changes}
	err := listeners.Show(client, d.Id(), showOpts).ExtractInto(&listener)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer listener (%s): %s", d.Id(), err)
	}

	return &listener, nil
}

func resourceMLBListenerV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	listener, err := resourceMLBListenerV1Show(d, managedLoadBalancerClient, true)
	if err != nil {
		return CheckDeleted(d, err, "listener")
	}

	log.Printf("[DEBUG] Retrieved ECL managed load balancer listener (%s): %+v", d.Id(), listener)

	if listener.ConfigurationStatus == "ACTIVE" {
		d.Set("ip_address", listener.IPAddress)
		d.Set("port", listener.Port)
		d.Set("protocol", listener.Protocol)
	} else if listener.ConfigurationStatus == "CREATE_STAGED" {
		d.Set("ip_address", listener.Staged.IPAddress)
		d.Set("port", listener.Staged.Port)
		d.Set("protocol", listener.Staged.Protocol)
	} else if listener.ConfigurationStatus == "UPDATE_STAGED" {
		d.Set("ip_address", ternary(listener.Staged.IPAddress == "", listener.IPAddress, listener.Staged.IPAddress))
		d.Set("port", ternary(listener.Staged.Port == 0, listener.Port, listener.Staged.Port))
		d.Set("protocol", ternary(listener.Staged.Protocol == "", listener.Protocol, listener.Staged.Protocol))
	} else if listener.ConfigurationStatus == "DELETE_STAGED" {
		d.SetId("")
		return nil
	}

	d.Set("name", listener.Name)
	d.Set("description", listener.Description)
	d.Set("tags", listener.Tags)
	d.Set("load_balancer_id", listener.LoadBalancerID)
	d.Set("tenant_id", listener.TenantID)

	return nil
}

func resourceMLBListenerV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Start updating attributes of ECL managed load balancer listener ...")

	err = resourceMLBListenerV1UpdateAttributes(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating attributes of ECL managed load balancer listener: %s", err)
	}

	log.Printf("[DEBUG] Start updating configurations of ECL managed load balancer listener ...")

	err = resourceMLBListenerV1UpdateConfigurations(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating configurations of ECL managed load balancer listener: %s", err)
	}

	return resourceMLBListenerV1Read(d, meta)
}

func resourceMLBListenerV1UpdateAttributes(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isAttributesUpdated bool
	var updateOpts listeners.UpdateOpts

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
		log.Printf("[DEBUG] Updating ECL managed load balancer listener attributes (%s) with options %+v", d.Id(), updateOpts)

		_, err := listeners.Update(client, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL managed load balancer listener attributes (%s) with options %+v: %s", d.Id(), updateOpts, err)
		}
	}

	return nil
}

func resourceMLBListenerV1UpdateConfigurations(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isConfigurationsUpdated bool

	listener, err := resourceMLBListenerV1Show(d, client, false)
	if err != nil {
		return err
	}

	if listener.ConfigurationStatus == "ACTIVE" {
		var createStagedOpts listeners.CreateStagedOpts

		if d.HasChange("ip_address") {
			isConfigurationsUpdated = true
			createStagedOpts.IPAddress = d.Get("ip_address").(string)
		}

		if d.HasChange("port") {
			isConfigurationsUpdated = true
			createStagedOpts.Port = d.Get("port").(int)
		}

		if d.HasChange("protocol") {
			isConfigurationsUpdated = true
			createStagedOpts.Protocol = d.Get("protocol").(string)
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer listener configurations (%s) with options %+v", d.Id(), createStagedOpts)

			_, err := listeners.CreateStaged(client, d.Id(), createStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer listener configurations (%s) with options %+v: %s", d.Id(), createStagedOpts, err)
			}
		}
	} else {
		var updateStagedOpts listeners.UpdateStagedOpts

		if d.HasChange("ip_address") {
			isConfigurationsUpdated = true
			ipAddress := d.Get("ip_address").(string)
			updateStagedOpts.IPAddress = &ipAddress
		}

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

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer listener configurations (%s) with options %+v", d.Id(), updateStagedOpts)

			_, err := listeners.UpdateStaged(client, d.Id(), updateStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer listener configurations (%s) with options %+v: %s", d.Id(), updateStagedOpts, err)
			}
		}
	}

	return nil
}

func resourceMLBListenerV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Deleting ECL managed load balancer listener: %s", d.Id())

	err = listeners.Delete(managedLoadBalancerClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer listener (%s)", d.Id())
			return nil
		}

		listener, err := resourceMLBListenerV1Show(d, managedLoadBalancerClient, false)
		if err != nil {
			return err
		}
		if listener.ConfigurationStatus == "DELETE_STAGED" {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer listener (%s)", d.Id())
			return nil
		}

		return fmt.Errorf("Error deleting ECL managed load balancer listener: %s", err)
	}

	return nil
}
