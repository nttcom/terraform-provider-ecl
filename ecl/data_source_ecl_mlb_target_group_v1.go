package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/target_groups"
)

func membersSchemaForDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
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
				"weight": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
			},
		},
	}
}

func dataSourceMLBTargetGroupV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBTargetGroupV1Read,
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
			"members": membersSchemaForDataSource(),
		},
	}

	return result
}

func dataSourceMLBTargetGroupV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := target_groups.ListOpts{}

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

	log.Printf("[DEBUG] Retrieving ECL managed load balancer target groups with options %+v", listOpts)

	pages, err := target_groups.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allTargetGroups, err := target_groups.ExtractTargetGroups(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer target groups with options %+v: %s", listOpts, err)
	}

	if len(allTargetGroups) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allTargetGroups) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	targetGroup := allTargetGroups[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer target group: %+v", targetGroup)

	d.SetId(targetGroup.ID)

	members := make([]interface{}, len(targetGroup.Members))
	for i, member := range targetGroup.Members {
		result := make(map[string]interface{})
		result["ip_address"] = member.IPAddress
		result["port"] = member.Port
		result["weight"] = member.Weight
		members[i] = result
	}

	d.Set("name", targetGroup.Name)
	d.Set("description", targetGroup.Description)
	d.Set("tags", targetGroup.Tags)
	d.Set("configuration_status", targetGroup.ConfigurationStatus)
	d.Set("operation_status", targetGroup.OperationStatus)
	d.Set("load_balancer_id", targetGroup.LoadBalancerID)
	d.Set("tenant_id", targetGroup.TenantID)
	d.Set("members", members)

	return nil
}
