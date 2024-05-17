package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v3/ecl/baremetal/v2/keypairs"
)

func resourceBaremetalKeypairV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceBaremetalKeypairV2Create,
		Read:   resourceBaremetalKeypairV2Read,
		Delete: resourceBaremetalKeypairV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceBaremetalKeypairV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	createOpts := keypairs.CreateOpts{
		Name:      d.Get("name").(string),
		PublicKey: d.Get("public_key").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	kp, err := keypairs.Create(baremetalClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL keypair: %s", err)
	}

	d.SetId(kp.Name)

	// Private Key is only available in the response to a create.
	d.Set("private_key", kp.PrivateKey)

	return resourceBaremetalKeypairV2Read(d, meta)
}

func resourceBaremetalKeypairV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	kp, err := keypairs.Get(baremetalClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "keypair")
	}

	d.Set("name", kp.Name)
	d.Set("public_key", kp.PublicKey)
	d.Set("fingerprint", kp.Fingerprint)

	return nil
}

func resourceBaremetalKeypairV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	err = keypairs.Delete(baremetalClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL keypair: %s", err)
	}

	d.SetId("")

	return nil
}
