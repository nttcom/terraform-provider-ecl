package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/operations"
)

func dataSourceMLBOperationV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBOperationV1Read,
		Schema: map[string]*schema.Schema{
			"resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_types": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reception_datetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"commit_datetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"warning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func dataSourceMLBOperationV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := operations.ListOpts{}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("resource_id"); ok {
		listOpts.ResourceID = v.(string)
	}

	if v, ok := d.GetOk("resource_type"); ok {
		listOpts.ResourceType = v.(string)
	}

	if v, ok := d.GetOk("request_id"); ok {
		listOpts.RequestID = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Retrieving ECL managed load balancer operations with options %+v", listOpts)

	pages, err := operations.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allOperations, err := operations.ExtractOperations(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer operations with options %+v: %s", listOpts, err)
	}

	if len(allOperations) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allOperations) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	operation := allOperations[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer operation: %+v", operation)

	d.SetId(operation.ID)

	d.Set("resource_id", operation.ResourceID)
	d.Set("resource_type", operation.ResourceType)
	d.Set("request_id", operation.RequestID)
	d.Set("request_types", operation.RequestTypes)
	d.Set("status", operation.Status)
	d.Set("reception_datetime", operation.ReceptionDatetime)
	d.Set("commit_datetime", operation.CommitDatetime)
	d.Set("warning", operation.Warning)
	d.Set("error", operation.Error)
	d.Set("tenant_id", operation.TenantID)

	return nil
}
