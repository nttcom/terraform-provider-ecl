package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_plans"
)

func dataSourceNetworkLoadBalancerPlanV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkLoadBalancerPlanV2Read,

		Schema: map[string]*schema.Schema{

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"maximum_syslog_servers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"model": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"edition": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vendor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkLoadBalancerPlanV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	listOpts := load_balancer_plans.ListOpts{
		Description:          d.Get("description").(string),
		ID:                   d.Get("id").(string),
		MaximumSyslogServers: d.Get("maximum_syslog_servers").(int),
		Name:                 d.Get("name").(string),
		Vendor:               d.Get("vendor").(string),
		Version:              d.Get("version").(string),
	}

	allPlans, err := getLoadBalancerPlans(networkClient, listOpts)
	if err != nil {
		return fmt.Errorf("error getting Load Balancer Plans: %w", err)
	}

	var filteredPlans []load_balancer_plans.LoadBalancerPlan

	// Loop through all Plans to find a more specific one.
	for _, plan := range allPlans {
		if v, ok := d.GetOk("enabled"); ok {
			if plan.Enabled != v.(bool) {
				continue
			}
		}
		if d.Get("model.#") == 1 {
			if v, ok := d.GetOk("model.0.edition"); ok {
				if plan.Model.Edition != v.(string) {
					continue
				}
			}

			if v, ok := d.GetOk("model.0.size"); ok {
				if plan.Model.Size != v.(string) {
					continue
				}
			}
		}
		filteredPlans = append(filteredPlans, plan)
	}

	if len(filteredPlans) > 1 {
		return fmt.Errorf("specified Load Balancer Plan query returned more than one result")
	}

	if len(filteredPlans) == 0 {
		return fmt.Errorf("specified Load Balancer Plan query returned no results")
	}

	plan := filteredPlans[0]

	log.Printf("[DEBUG] Retrieved Load Balancer Plan %s: %+v", plan.ID, plan)

	d.SetId(plan.ID)
	d.Set("description", plan.Description)
	d.Set("enabled", plan.Enabled)
	d.Set("id", plan.ID)
	d.Set("maximum_syslog_servers", plan.MaximumSyslogServers)

	model := make(map[string]interface{})
	model["size"] = plan.Model.Size
	model["edition"] = plan.Model.Edition
	d.Set("model", []interface{}{model})

	d.Set("name", plan.Name)
	d.Set("vendor", plan.Vendor)
	d.Set("version", plan.Version)

	return nil
}

func getLoadBalancerPlans(networkClient *eclcloud.ServiceClient, listOpts load_balancer_plans.ListOpts) ([]load_balancer_plans.LoadBalancerPlan, error) {
	pages, err := load_balancer_plans.List(networkClient, listOpts).AllPages()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Load Balancer Plans: %w", err)
	}

	allPlans, err := load_balancer_plans.ExtractLoadBalancerPlans(pages)
	if err != nil {
		return nil, fmt.Errorf("unable to extract retrieved Load Balancer Plans: %w", err)
	}

	if len(allPlans) < 1 {
		return nil, fmt.Errorf("specified Load Balancer Plan query returned no results")
	}
	return allPlans, nil
}
