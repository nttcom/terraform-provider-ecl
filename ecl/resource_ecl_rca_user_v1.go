package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud/v2/ecl/rca/v1/users"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceRCAUserV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceRCAUserV1Create,
		Read:   resourceRCAUserV1Read,
		Update: resourceRCAUserV1Update,
		Delete: resourceRCAUserV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(8, 127),
			},
			"vpn_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceRCAUserV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.rcaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL RCA client: %s", err)
	}

	opts := users.CreateOpts{
		Password: d.Get("password").(string),
	}
	log.Printf("[DEBUG] Create Options: %#v", opts)

	user, err := users.Create(client, opts).Extract()
	if err != nil {
		return fmt.Errorf("error creating ECL RCA user: %s", err)
	}

	d.SetId(user.Name)
	d.Set("password", user.Password)

	log.Printf("[DEBUG] Created ECL RCA user %s: %#v", user.Name, user)

	return resourceRCAUserV1Read(d, meta)
}

func resourceRCAUserV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.rcaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL RCA client: %s", err)
	}

	user, err := users.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "rca user")
	}
	log.Printf("[DEBUG] Retrieved RCA user %s: %+v", user.Name, user)

	d.Set("name", user.Name)

	var endpoints []map[string]string
	for _, v := range user.VPNEndpoints {
		endpoint := map[string]string{
			"endpoint": v.Endpoint,
			"type":     v.Type,
		}
		endpoints = append(endpoints, endpoint)
	}
	d.Set("vpn_endpoints", endpoints)

	return nil
}

func resourceRCAUserV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.rcaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL RCA client: %s", err)
	}

	if err := users.Delete(client, d.Id()).ExtractErr(); err != nil {
		return fmt.Errorf("error deleting ECL RCA user: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceRCAUserV1Update(d *schema.ResourceData, meta interface{}) error {
	if !d.HasChange("password") {
		return nil
	}

	config := meta.(*Config)
	client, err := config.rcaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL RCA client: %s", err)
	}

	password := d.Get("password").(string)
	opts := users.UpdateOpts{
		Password: password,
	}
	log.Printf("[DEBUG] Update Options: %#v", opts)

	user, err := users.Update(client, d.Id(), opts).Extract()
	if err != nil {
		return fmt.Errorf("error updating ECL RCA user: %s", err)
	}

	d.Set("password", password)

	log.Printf("[DEBUG] Updated ECL RCA user %s: %#v", user.Name, user)

	return resourceRCAUserV1Read(d, meta)
}
