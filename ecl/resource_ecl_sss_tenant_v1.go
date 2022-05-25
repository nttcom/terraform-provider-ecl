package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v2/ecl/sss/v1/tenants"
)

func resourceSSSTenantV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSSTenantV1Create,
		Read:   resourceSSSTenantV1Read,
		Update: resourceSSSTenantV1Update,
		Delete: resourceSSSTenantV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"tenant_region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"contract_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

		DeprecationMessage: "The ecl_sss_tenant resource has been deprecated and will be removed in a future version.",
	}
}

func resourceSSSTenantV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	createOpts := tenants.CreateOpts{
		TenantName:   d.Get("tenant_name").(string),
		Description:  d.Get("description").(string),
		TenantRegion: d.Get("tenant_region").(string),
		ContractID:   d.Get("contract_id").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	tenant, err := tenants.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL tenant: %s", err)
	}

	log.Printf("[DEBUG] Tenant has successfully created.")
	d.SetId(tenant.TenantID)

	return resourceSSSTenantV1Read(d, meta)
}

func resourceSSSTenantV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	tenant, err := tenants.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "tenant")
	}
	log.Printf("[DEBUG] Retrieved ECL tenant: %#v", tenant)

	d.Set("tenant_name", tenant.TenantName)
	d.Set("description", tenant.Description)
	d.Set("tenant_region", tenant.TenantRegion)
	d.Set("contract_id", tenant.ContractID)

	d.Set("tenant_id", tenant.TenantID)
	d.Set("start_time", tenant.StartTime.String())

	log.Printf("[DEBUG] resourceSSSTenantV1Read Succeeded")

	return nil
}

func resourceSSSTenantV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	var hasChange bool
	var updateOpts tenants.UpdateOpts

	if d.HasChange("description") {
		hasChange = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if hasChange {
		r := tenants.Update(client, d.Id(), updateOpts)
		if r.Err != nil {
			return fmt.Errorf("Error updating ECL tenant: %s", r.Err)
		}
		log.Printf("[DEBUG] Tenant has successfully updated.")
	}

	return resourceSSSTenantV1Read(d, meta)
}

func resourceSSSTenantV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	err = tenants.Delete(client, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL tenant: %s", err)
	}

	log.Printf("[DEBUG] Tenant has successfully deleted.")
	return nil
}
