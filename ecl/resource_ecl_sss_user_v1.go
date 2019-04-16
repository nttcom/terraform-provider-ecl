package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/nttcom/eclcloud/ecl/sss/v1/users"
)

func resourceSSSUserV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSSUserV1Create,
		Read:   resourceSSSUserV1Read,
		Update: resourceSSSUserV1Update,
		Delete: resourceSSSUserV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"login_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"mail_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"notify_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"true", "false",
				}, false),
			},

			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"contract_owner": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},

			"keystone_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"keystone_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"keystone_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"sss_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"contract_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"login_integration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"external_reference_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSSSUserV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	createOpts := users.CreateOpts{
		LoginID:        d.Get("login_id").(string),
		MailAddress:    d.Get("mail_address").(string),
		Password:       d.Get("password").(string),
		NotifyPassword: d.Get("notify_password").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	user, err := users.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL user: %s", err)
	}

	log.Printf("[DEBUG] User has successfully created.")
	d.SetId(user.UserID)

	return resourceSSSUserV1Read(d, meta)
}

func resourceSSSUserV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	user, err := users.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "user")
	}
	log.Printf("[DEBUG] Retrieved ECL user: %#v", user)

	d.Set("login_id", user.LoginID)
	d.Set("mail_address", user.MailAddress)
	d.Set("user_id", user.UserID)
	d.Set("contract_owner", user.ContractOwner)
	d.Set("keystone_endpoint", user.KeystoneEndpoint)
	d.Set("keystone_name", user.KeystoneName)
	d.Set("sss_endpoint", user.SSSEndpoint)
	d.Set("contract_id", user.ContractID)
	d.Set("login_integration", user.LoginIntegration)
	d.Set("external_reference_id", user.ExternalReferenceID)
	d.Set("brand_id", user.BrandID)
	d.Set("start_time", user.StartTime)

	log.Printf("[DEBUG] resourceSSSUserV1Read Succeeded")

	return nil
}

func resourceSSSUserV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	var hasChange bool
	var updateOpts users.UpdateOpts

	if d.HasChange("login_id") {
		hasChange = true
		loginID := d.Get("login_id").(string)
		updateOpts.LoginID = &loginID
	}

	if d.HasChange("mail_address") {
		hasChange = true
		mailAddress := d.Get("mail_address").(string)
		updateOpts.MailAddress = &mailAddress
	}

	if d.HasChange("password") {
		hasChange = true
		newPassword := d.Get("password").(string)
		updateOpts.NewPassword = &newPassword
	}

	if hasChange {
		r := users.Update(client, d.Id(), updateOpts)
		if r.Err != nil {
			return fmt.Errorf("Error updating ECL user: %s", r.Err)
		}
		log.Printf("[DEBUG] User has successfully updated.")
	}

	return resourceSSSUserV1Read(d, meta)
}

func resourceSSSUserV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	err = users.Delete(client, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL user: %s", err)
	}

	log.Printf("[DEBUG] User has successfully deleted.")
	return nil
}
