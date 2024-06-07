package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/routes"
)

func resourceMLBRouteV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceMLBRouteV1Create,
		Read:   resourceMLBRouteV1Read,
		Update: resourceMLBRouteV1Update,
		Delete: resourceMLBRouteV1Delete,
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
			"destination_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"next_hop_ip_address": &schema.Schema{
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

func resourceMLBRouteV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOpts := routes.CreateOpts{
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		Tags:             d.Get("tags").(map[string]interface{}),
		DestinationCidr:  d.Get("destination_cidr").(string),
		NextHopIPAddress: d.Get("next_hop_ip_address").(string),
		LoadBalancerID:   d.Get("load_balancer_id").(string),
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Creating ECL managed load balancer route with options %+v", createOpts)

	route, err := routes.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer route with options %+v: %s", createOpts, err)
	}

	d.SetId(route.ID)
	log.Printf("[INFO] ECL managed load balancer route ID: %s", route.ID)

	return resourceMLBRouteV1Read(d, meta)
}

func resourceMLBRouteV1Show(d *schema.ResourceData, client *eclcloud.ServiceClient, changes bool) (*routes.Route, error) {
	var route routes.Route

	showOpts := routes.ShowOpts{Changes: changes}
	err := routes.Show(client, d.Id(), showOpts).ExtractInto(&route)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer route (%s): %s", d.Id(), err)
	}

	return &route, nil
}

func resourceMLBRouteV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	route, err := resourceMLBRouteV1Show(d, managedLoadBalancerClient, true)
	if err != nil {
		return CheckDeleted(d, err, "route")
	}

	if route.ConfigurationStatus == "ACTIVE" {
		d.Set("next_hop_ip_address", route.NextHopIPAddress)
	} else if route.ConfigurationStatus == "CREATE_STAGED" {
		d.Set("next_hop_ip_address", route.Staged.NextHopIPAddress)
	} else if route.ConfigurationStatus == "UPDATE_STAGED" {
		d.Set("next_hop_ip_address", ternary(route.Staged.NextHopIPAddress == "", route.NextHopIPAddress, route.Staged.NextHopIPAddress))
	} else if route.ConfigurationStatus == "DELETE_STAGED" {
		d.SetId("")
		return nil
	}

	d.Set("name", route.Name)
	d.Set("description", route.Description)
	d.Set("tags", route.Tags)
	d.Set("destination_cidr", route.DestinationCidr)
	d.Set("load_balancer_id", route.LoadBalancerID)
	d.Set("tenant_id", route.TenantID)

	return nil
}

func resourceMLBRouteV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Start updating attributes of ECL managed load balancer route ...")

	err = resourceMLBRouteV1UpdateAttributes(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating attributes of ECL managed load balancer route: %s", err)
	}

	log.Printf("[DEBUG] Start updating configurations of ECL managed load balancer route ...")

	err = resourceMLBRouteV1UpdateConfigurations(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating configurations of ECL managed load balancer route: %s", err)
	}

	return resourceMLBRouteV1Read(d, meta)
}

func resourceMLBRouteV1UpdateAttributes(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isAttributesUpdated bool
	var updateOpts routes.UpdateOpts

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
		log.Printf("[DEBUG] Updating ECL managed load balancer route attributes (%s) with options %+v", d.Id(), updateOpts)

		_, err := routes.Update(client, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL managed load balancer route attributes (%s) with options %+v: %s", d.Id(), updateOpts, err)
		}
	}

	return nil
}

func resourceMLBRouteV1UpdateConfigurations(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isConfigurationsUpdated bool

	route, err := resourceMLBRouteV1Show(d, client, false)
	if err != nil {
		return err
	}

	if route.ConfigurationStatus == "ACTIVE" {
		var createStagedOpts routes.CreateStagedOpts

		if d.HasChange("next_hop_ip_address") {
			isConfigurationsUpdated = true
			createStagedOpts.NextHopIPAddress = d.Get("next_hop_ip_address").(string)
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer route configurations (%s) with options %+v", d.Id(), createStagedOpts)

			_, err := routes.CreateStaged(client, d.Id(), createStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer route configurations (%s) with options %+v: %s", d.Id(), createStagedOpts, err)
			}
		}
	} else {
		var updateStagedOpts routes.UpdateStagedOpts

		if d.HasChange("next_hop_ip_address") {
			isConfigurationsUpdated = true
			nextHopIPAddress := d.Get("next_hop_ip_address").(string)
			updateStagedOpts.NextHopIPAddress = &nextHopIPAddress
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer route configurations (%s) with options %+v", d.Id(), updateStagedOpts)

			_, err := routes.UpdateStaged(client, d.Id(), updateStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer route configurations (%s) with options %+v: %s", d.Id(), updateStagedOpts, err)
			}
		}
	}

	return nil
}

func resourceMLBRouteV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Deleting ECL managed load balancer route: %s", d.Id())

	err = routes.Delete(managedLoadBalancerClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer route (%s)", d.Id())
			return nil
		}

		route, err := resourceMLBRouteV1Show(d, managedLoadBalancerClient, false)
		if err != nil {
			return err
		}
		if route.ConfigurationStatus == "DELETE_STAGED" {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer route (%s)", d.Id())
			return nil
		}

		return fmt.Errorf("Error deleting ECL managed load balancer route: %s", err)
	}

	return nil
}
