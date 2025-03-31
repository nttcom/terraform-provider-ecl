package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/target_groups"
)

func membersSchemaForResource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"port": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"weight": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
			},
		},
	}
}

func resourceMLBTargetGroupV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read:   resourceMLBTargetGroupV1Read,
		Create: resourceMLBTargetGroupV1Create,
		Update: resourceMLBTargetGroupV1Update,
		Delete: resourceMLBTargetGroupV1Delete,
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
			"members": membersSchemaForResource(),
		},
	}

	return result
}

func resourceMLBTargetGroupV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOpts := target_groups.CreateOpts{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Tags:           d.Get("tags").(map[string]interface{}),
		LoadBalancerID: d.Get("load_balancer_id").(string),
	}

	members := make([]target_groups.CreateOptsMember, len(d.Get("members").([]interface{})))
	for i, member := range d.Get("members").([]interface{}) {
		members[i] = target_groups.CreateOptsMember{
			IPAddress: member.(map[string]interface{})["ip_address"].(string),
			Port:      member.(map[string]interface{})["port"].(int),
			Weight:    member.(map[string]interface{})["weight"].(int),
		}
	}
	createOpts.Members = &members

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Creating ECL managed load balancer target group with options %+v", createOpts)

	rule, err := target_groups.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer target group with options %+v: %s", createOpts, err)
	}

	d.SetId(rule.ID)
	log.Printf("[INFO] ECL managed load balancer target group ID: %s", rule.ID)

	return resourceMLBTargetGroupV1Read(d, meta)
}

func resourceMLBTargetGroupV1Show(d *schema.ResourceData, client *eclcloud.ServiceClient, changes bool) (*target_groups.TargetGroup, error) {
	var targetGroup target_groups.TargetGroup

	showOpts := target_groups.ShowOpts{Changes: changes}
	err := target_groups.Show(client, d.Id(), showOpts).ExtractInto(&targetGroup)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer target group (%s): %s", d.Id(), err)
	}

	return &targetGroup, nil
}

func resourceMLBTargetGroupV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	targetGroup, err := resourceMLBTargetGroupV1Show(d, managedLoadBalancerClient, true)
	if err != nil {
		return CheckDeleted(d, err, "target_group")
	}

	log.Printf("[DEBUG] Retrieved ECL managed load balancer target group (%s): %+v", d.Id(), targetGroup)

	if targetGroup.ConfigurationStatus == "ACTIVE" || (targetGroup.ConfigurationStatus == "UPDATE_STAGED" && targetGroup.Members == nil) {
		members := make([]interface{}, len(targetGroup.Members))
		for i, member := range targetGroup.Members {
			members[i] = map[string]interface{}{
				"ip_address": member.IPAddress,
				"port":       member.Port,
				"weight":     member.Weight,
			}
		}

		d.Set("members", members)
	} else if targetGroup.ConfigurationStatus == "CREATE_STAGED" || (targetGroup.ConfigurationStatus == "UPDATE_STAGED" && targetGroup.Members != nil) {
		members := make([]interface{}, len(targetGroup.Staged.Members))
		for i, member := range targetGroup.Staged.Members {
			members[i] = map[string]interface{}{
				"ip_address": member.IPAddress,
				"port":       member.Port,
				"weight":     member.Weight,
			}
		}

		d.Set("members", members)
	} else if targetGroup.ConfigurationStatus == "DELETE_STAGED" {
		d.SetId("")
		return nil
	}

	d.Set("name", targetGroup.Name)
	d.Set("description", targetGroup.Description)
	d.Set("tags", targetGroup.Tags)
	d.Set("load_balancer_id", targetGroup.LoadBalancerID)
	d.Set("tenant_id", targetGroup.TenantID)

	return nil
}

func resourceMLBTargetGroupV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Start updating attributes of ECL managed load balancer target group ...")

	err = resourceMLBTargetGroupV1UpdateAttributes(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating attributes of ECL managed load balancer target group: %s", err)
	}

	log.Printf("[DEBUG] Start updating configurations of ECL managed load balancer target group ...")

	err = resourceMLBTargetGroupV1UpdateConfigurations(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating configurations of ECL managed load balancer target group: %s", err)
	}

	return resourceMLBTargetGroupV1Read(d, meta)
}

func resourceMLBTargetGroupV1UpdateAttributes(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isAttributesUpdated bool
	var updateOpts target_groups.UpdateOpts

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
		log.Printf("[DEBUG] Updating ECL managed load balancer target group attributes (%s) with options %+v", d.Id(), updateOpts)

		_, err := target_groups.Update(client, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL managed load balancer target group attributes (%s) with options %+v: %s", d.Id(), updateOpts, err)
		}
	}

	return nil
}

func resourceMLBTargetGroupV1UpdateConfigurations(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isConfigurationsUpdated bool

	targetGroup, err := resourceMLBTargetGroupV1Show(d, client, false)
	if err != nil {
		return err
	}

	if targetGroup.ConfigurationStatus == "ACTIVE" {
		var createStagedOpts target_groups.CreateStagedOpts

		if d.HasChange("members") {
			isConfigurationsUpdated = true

			members := make([]target_groups.CreateStagedOptsMember, len(d.Get("members").([]interface{})))
			for i, member := range d.Get("members").([]interface{}) {
				members[i] = target_groups.CreateStagedOptsMember{
					IPAddress: member.(map[string]interface{})["ip_address"].(string),
					Port:      member.(map[string]interface{})["port"].(int),
					Weight:    member.(map[string]interface{})["weight"].(int),
				}
			}
			createStagedOpts.Members = &members
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer target group configurations (%s) with options %+v", d.Id(), createStagedOpts)

			_, err := target_groups.CreateStaged(client, d.Id(), createStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer target group configurations (%s) with options %+v: %s", d.Id(), createStagedOpts, err)
			}
		}
	} else {
		var updateStagedOpts target_groups.UpdateStagedOpts

		if d.HasChange("members") {
			isConfigurationsUpdated = true

			members := make([]target_groups.UpdateStagedOptsMember, len(d.Get("members").([]interface{})))
			for i, member := range d.Get("members").([]interface{}) {
				ipAddress := member.(map[string]interface{})["ip_address"].(string)
				port := member.(map[string]interface{})["port"].(int)
				weight := member.(map[string]interface{})["weight"].(int)
				members[i] = target_groups.UpdateStagedOptsMember{
					IPAddress: &ipAddress,
					Port:      &port,
					Weight:    &weight,
				}
			}
			updateStagedOpts.Members = &members
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer target group configurations (%s) with options %+v", d.Id(), updateStagedOpts)

			_, err := target_groups.UpdateStaged(client, d.Id(), updateStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer target group configurations (%s) with options %+v: %s", d.Id(), updateStagedOpts, err)
			}
		}
	}

	return nil
}

func resourceMLBTargetGroupV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Deleting ECL managed load balancer target group: %s", d.Id())

	err = target_groups.Delete(managedLoadBalancerClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer target group (%s)", d.Id())
			return nil
		}

		targetGroup, err := resourceMLBTargetGroupV1Show(d, managedLoadBalancerClient, false)
		if err != nil {
			return err
		}
		if targetGroup.ConfigurationStatus == "DELETE_STAGED" {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer target group (%s)", d.Id())
			return nil
		}

		return fmt.Errorf("Error deleting ECL managed load balancer target group: %s", err)
	}

	return nil
}
