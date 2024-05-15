package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/rules"
)

func conditionsSchemaForDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"path_patterns": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Computed: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func dataSourceMLBRuleV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBRuleV1Read,
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
			"policy_id": &schema.Schema{
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
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"target_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conditions": conditionsSchemaForDataSource(),
		},
	}

	return result
}

func dataSourceMLBRuleV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := rules.ListOpts{}

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

	if v, ok := d.GetOk("priority"); ok {
		listOpts.Priority = v.(int)
	}

	if v, ok := d.GetOk("target_group_id"); ok {
		listOpts.TargetGroupID = v.(string)
	}

	if v, ok := d.GetOk("policy_id"); ok {
		listOpts.PolicyID = v.(string)
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

	log.Printf("[DEBUG] Retrieving ECL managed load balancer rules with options %+v", listOpts)

	pages, err := rules.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allRules, err := rules.ExtractRules(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer rules with options %+v: %s", listOpts, err)
	}

	if len(allRules) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allRules) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	rule := allRules[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer rule: %+v", rule)

	d.SetId(rule.ID)

	d.Set("name", rule.Name)
	d.Set("description", rule.Description)
	d.Set("tags", rule.Tags)
	d.Set("configuration_status", rule.ConfigurationStatus)
	d.Set("operation_status", rule.OperationStatus)
	d.Set("policy_id", rule.PolicyID)
	d.Set("load_balancer_id", rule.LoadBalancerID)
	d.Set("tenant_id", rule.TenantID)
	d.Set("priority", rule.Priority)
	d.Set("target_group_id", rule.TargetGroupID)
	d.Set("conditions", []interface{}{map[string]interface{}{
		"path_patterns": rule.Conditions.PathPatterns,
	}})

	return nil
}
