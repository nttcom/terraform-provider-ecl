package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/ecl/sss/v1/tenants"
)

func dataSourceSSSTenantV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSSSTenantV1Read,

		Schema: map[string]*schema.Schema{
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"contract_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// dataSourceSSSTenantV1Read performs the project lookup.
func dataSourceSSSTenantV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	listOpts := tenants.ListOpts{}

	log.Printf("[DEBUG] List Options: %#v", listOpts)

	var tenant tenants.Tenant
	allPages, err := tenants.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to query tenants: %s", err)
	}

	allTenants, err := tenants.ExtractTenants(allPages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve tenants: %s", err)
	}

	var refinedTenants []tenants.Tenant
	if len(allTenants) > 0 {
		for _, t := range allTenants {
			if t.TenantName == d.Get("tenant_name").(string) {
				refinedTenants = append(refinedTenants, t)
			}
		}
	}

	if len(refinedTenants) > 1 {
		log.Printf("[DEBUG] Multiple results found: %#v", allTenants)
		return fmt.Errorf("Your query returned more than one result")
	}

	if len(refinedTenants) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	tenant = refinedTenants[0]

	log.Printf("[DEBUG] Single tenant found: %s", tenant.TenantID)

	d.SetId(tenant.TenantID)
	d.Set("tenant_name", tenant.TenantName)
	d.Set("description", tenant.Description)
	d.Set("tenant_region", tenant.TenantRegion)
	d.Set("contract_id", tenant.ContractID)

	d.Set("tenant_id", tenant.TenantID)
	d.Set("start_time", tenant.StartTime.String())

	return nil
}
