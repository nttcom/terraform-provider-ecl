package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/rules"
)

func conditionsSchemaForResource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"path_patterns": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func resourceMLBRuleV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceMLBRuleV1Create,
		Read:   resourceMLBRuleV1Read,
		Update: resourceMLBRuleV1Update,
		Delete: resourceMLBRuleV1Delete,
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
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"target_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"conditions": conditionsSchemaForResource(),
		},
	}

	return result
}

func resourceMLBRuleV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOpts := rules.CreateOpts{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Tags:          d.Get("tags").(map[string]interface{}),
		Priority:      d.Get("priority").(int),
		TargetGroupID: d.Get("target_group_id").(string),
		PolicyID:      d.Get("policy_id").(string),
	}

	pathPatterns := make([]string, len(d.Get("conditions").([]interface{})[0].(map[string]interface{})["path_patterns"].([]interface{})))
	for i, pathPattern := range d.Get("conditions").([]interface{})[0].(map[string]interface{})["path_patterns"].([]interface{}) {
		pathPatterns[i] = pathPattern.(string)
	}

	condition := rules.CreateOptsCondition{
		PathPatterns: pathPatterns,
	}
	createOpts.Conditions = &condition

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Creating ECL managed load balancer rule with options %+v", createOpts)

	rule, err := rules.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer rule with options %+v: %s", createOpts, err)
	}

	d.SetId(rule.ID)
	log.Printf("[INFO] ECL managed load balancer rule ID: %s", rule.ID)

	return resourceMLBRuleV1Read(d, meta)
}

func resourceMLBRuleV1Show(d *schema.ResourceData, client *eclcloud.ServiceClient, changes bool) (*rules.Rule, error) {
	var rule rules.Rule

	showOpts := rules.ShowOpts{Changes: changes}
	err := rules.Show(client, d.Id(), showOpts).ExtractInto(&rule)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer rule (%s): %s", d.Id(), err)
	}

	return &rule, nil
}

func resourceMLBRuleV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	rule, err := resourceMLBRuleV1Show(d, managedLoadBalancerClient, true)
	if err != nil {
		return CheckDeleted(d, err, "rule")
	}

	log.Printf("[DEBUG] Retrieved ECL managed load balancer rule (%s): %+v", d.Id(), rule)

	conditions := make(map[string]interface{})

	if rule.ConfigurationStatus == "ACTIVE" {
		d.Set("priority", rule.Priority)
		d.Set("target_group_id", rule.TargetGroupID)
		conditions["path_patterns"] = rule.Conditions.PathPatterns
	} else if rule.ConfigurationStatus == "CREATE_STAGED" {
		d.Set("priority", rule.Staged.Priority)
		d.Set("target_group_id", rule.Staged.TargetGroupID)
		conditions["path_patterns"] = rule.Staged.Conditions.PathPatterns
	} else if rule.ConfigurationStatus == "UPDATE_STAGED" {
		d.Set("priority", ternary(rule.Staged.Priority == 0, rule.Priority, rule.Staged.Priority))
		d.Set("target_group_id", ternary(rule.Staged.TargetGroupID == "", rule.TargetGroupID, rule.Staged.TargetGroupID))
		conditions["path_patterns"] = ternary(rule.Staged.Conditions.PathPatterns == nil, rule.Conditions.PathPatterns, rule.Staged.Conditions.PathPatterns)
	} else if rule.ConfigurationStatus == "DELETE_STAGED" {
		d.SetId("")
		return nil
	}

	d.Set("name", rule.Name)
	d.Set("description", rule.Description)
	d.Set("tags", rule.Tags)
	d.Set("policy_id", rule.PolicyID)
	d.Set("load_balancer_id", rule.LoadBalancerID)
	d.Set("tenant_id", rule.TenantID)
	d.Set("conditions", []interface{}{conditions})

	return nil
}

func resourceMLBRuleV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Start updating attributes of ECL managed load balancer rule ...")

	err = resourceMLBRuleV1UpdateAttributes(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating attributes of ECL managed load balancer rule: %s", err)
	}

	log.Printf("[DEBUG] Start updating configurations of ECL managed load balancer rule ...")

	err = resourceMLBRuleV1UpdateConfigurations(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating configurations of ECL managed load balancer rule: %s", err)
	}

	return resourceMLBRuleV1Read(d, meta)
}

func resourceMLBRuleV1UpdateAttributes(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isAttributesUpdated bool
	var updateOpts rules.UpdateOpts

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
		log.Printf("[DEBUG] Updating ECL managed load balancer rule attributes (%s) with options %+v", d.Id(), updateOpts)

		_, err := rules.Update(client, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL managed load balancer rule attributes (%s) with options %+v: %s", d.Id(), updateOpts, err)
		}
	}

	return nil
}

func resourceMLBRuleV1UpdateConfigurations(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isConfigurationsUpdated bool

	pathPatterns := make([]string, len(d.Get("conditions").([]interface{})[0].(map[string]interface{})["path_patterns"].([]interface{})))
	for i, pathPattern := range d.Get("conditions").([]interface{})[0].(map[string]interface{})["path_patterns"].([]interface{}) {
		pathPatterns[i] = pathPattern.(string)
	}

	rule, err := resourceMLBRuleV1Show(d, client, false)
	if err != nil {
		return err
	}

	if rule.ConfigurationStatus == "ACTIVE" {
		var createStagedOpts rules.CreateStagedOpts

		if d.HasChange("priority") {
			isConfigurationsUpdated = true
			createStagedOpts.Priority = d.Get("priority").(int)
		}

		if d.HasChange("target_group_id") {
			isConfigurationsUpdated = true
			createStagedOpts.TargetGroupID = d.Get("target_group_id").(string)
		}

		if d.HasChange("conditions") {
			isConfigurationsUpdated = true
			condition := rules.CreateStagedOptsCondition{
				PathPatterns: pathPatterns,
			}
			createStagedOpts.Conditions = &condition
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer rule configurations (%s) with options %+v", d.Id(), createStagedOpts)

			_, err := rules.CreateStaged(client, d.Id(), createStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer rule configurations (%s) with options %+v: %s", d.Id(), createStagedOpts, err)
			}
		}
	} else {
		var updateStagedOpts rules.UpdateStagedOpts

		if d.HasChange("priority") {
			isConfigurationsUpdated = true
			priority := d.Get("priority").(int)
			updateStagedOpts.Priority = &priority
		}

		if d.HasChange("target_group_id") {
			isConfigurationsUpdated = true
			targetGroupID := d.Get("target_group_id").(string)
			updateStagedOpts.TargetGroupID = &targetGroupID
		}

		if d.HasChange("conditions") {
			isConfigurationsUpdated = true
			condition := rules.UpdateStagedOptsCondition{
				PathPatterns: &pathPatterns,
			}
			updateStagedOpts.Conditions = &condition
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer rule configurations (%s) with options %+v", d.Id(), updateStagedOpts)

			_, err := rules.UpdateStaged(client, d.Id(), updateStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer rule configurations (%s) with options %+v: %s", d.Id(), updateStagedOpts, err)
			}
		}
	}

	return nil
}

func resourceMLBRuleV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Deleting ECL managed load balancer rule: %s", d.Id())

	err = rules.Delete(managedLoadBalancerClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer rule (%s)", d.Id())
			return nil
		}

		rule, err := resourceMLBRuleV1Show(d, managedLoadBalancerClient, false)
		if err != nil {
			return err
		}
		if rule.ConfigurationStatus == "DELETE_STAGED" {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer rule (%s)", d.Id())
			return nil
		}

		return fmt.Errorf("Error deleting ECL managed load balancer rule: %s", err)
	}

	return nil
}
