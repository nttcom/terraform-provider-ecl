package ecl

import (
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceImageStoragesMemberAccepterV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageStoragesMemberAccepterV2Create,
		Read:   resourceImageStoragesMemberV2Read,
		Update: resourceImageStoragesMemberV2Update,
		Delete: resourceImageStoragesMemberAccepterV2Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: "This region field is deprecated and will be removed from a future version.",
			},

			"image_member_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"accepted", "rejected",
				}, true),
			},
		},
	}
}

func resourceImageStoragesMemberAccepterV2Create(d *schema.ResourceData, meta interface{}) error {
	id := d.Get("image_member_id").(string)
	d.SetId(id)

	return resourceImageStoragesMemberV2Update(d, meta)
}

func resourceImageStoragesMemberAccepterV2Delete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Will not delete Image Member. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
