package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/plans"
)

func dataSourceMLBPlanV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBPlanV1Read,
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
			"bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redundancy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_number_of_interfaces": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_number_of_health_monitors": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_number_of_listeners": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_number_of_policies": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_number_of_routes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_number_of_target_groups": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_number_of_members": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_number_of_rules": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_number_of_conditions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}

	return result
}

func dataSourceMLBPlanV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := plans.ListOpts{}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("bandwidth"); ok {
		listOpts.Bandwidth = v.(string)
	}

	if v, ok := d.GetOk("redundancy"); ok {
		listOpts.Redundancy = v.(string)
	}

	if v, ok := d.GetOk("max_number_of_interfaces"); ok {
		listOpts.MaxNumberOfInterfaces = v.(int)
	}

	if v, ok := d.GetOk("max_number_of_health_monitors"); ok {
		listOpts.MaxNumberOfHealthMonitors = v.(int)
	}

	if v, ok := d.GetOk("max_number_of_listeners"); ok {
		listOpts.MaxNumberOfListeners = v.(int)
	}

	if v, ok := d.GetOk("max_number_of_policies"); ok {
		listOpts.MaxNumberOfPolicies = v.(int)
	}

	if v, ok := d.GetOk("max_number_of_routes"); ok {
		listOpts.MaxNumberOfRoutes = v.(int)
	}

	if v, ok := d.GetOk("max_number_of_target_groups"); ok {
		listOpts.MaxNumberOfTargetGroups = v.(int)
	}

	if v, ok := d.GetOk("max_number_of_members"); ok {
		listOpts.MaxNumberOfMembers = v.(int)
	}

	if v, ok := d.GetOk("max_number_of_rules"); ok {
		listOpts.MaxNumberOfRules = v.(int)
	}

	if v, ok := d.GetOk("max_number_of_conditions"); ok {
		listOpts.MaxNumberOfConditions = v.(int)
	}

	if v, ok := d.GetOk("enabled"); ok {
		listOpts.Enabled = v.(bool)
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Retrieving ECL managed load balancer plans with options %+v", listOpts)

	pages, err := plans.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allPlans, err := plans.ExtractPlans(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer plans with options %+v: %s", listOpts, err)
	}

	if len(allPlans) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allPlans) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	plan := allPlans[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer plan: %+v", plan)

	d.SetId(plan.ID)

	d.Set("name", plan.Name)
	d.Set("description", plan.Description)
	d.Set("bandwidth", plan.Bandwidth)
	d.Set("redundancy", plan.Redundancy)
	d.Set("max_number_of_interfaces", plan.MaxNumberOfInterfaces)
	d.Set("max_number_of_health_monitors", plan.MaxNumberOfHealthMonitors)
	d.Set("max_number_of_listeners", plan.MaxNumberOfListeners)
	d.Set("max_number_of_policies", plan.MaxNumberOfPolicies)
	d.Set("max_number_of_routes", plan.MaxNumberOfRoutes)
	d.Set("max_number_of_target_groups", plan.MaxNumberOfTargetGroups)
	d.Set("max_number_of_members", plan.MaxNumberOfMembers)
	d.Set("max_number_of_rules", plan.MaxNumberOfRules)
	d.Set("max_number_of_conditions", plan.MaxNumberOfConditions)
	d.Set("enabled", plan.Enabled)

	return nil
}
