package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v4/ecl/baremetal/v2/keypairs"
)

func dataSourceBaremetalKeypairV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBaremetalKeypairV2Read,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"public_key": &schema.Schema{
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

// dataSourceBaremetalKeypairV2Read performs the keypair lookup.
func dataSourceBaremetalKeypairV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	baremetalClient, err := config.baremetalV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	allPages, err := keypairs.List(baremetalClient).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to query keypair: %s", err)
	}

	allKeypairs, err := keypairs.ExtractKeyPairs(allPages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve keypair: %s", err)
	}

	// Loop through all keypair to find a more specific one.
	if len(allKeypairs) > 1 {
		var filteredKeypairs []keypairs.KeyPair
		for _, keypair := range allKeypairs {
			if v := d.Get("name").(string); v != "" {
				if keypair.Name != v {
					continue
				}
			}

			filteredKeypairs = append(filteredKeypairs, keypair)
		}

		allKeypairs = filteredKeypairs
	}

	if len(allKeypairs) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allKeypairs) > 1 {
		log.Printf("[DEBUG] Multiple results found: %#v", allKeypairs)
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria")
	}

	keypair := allKeypairs[0]
	log.Printf("[DEBUG] Single Keypair found: %s", keypair.Name)
	return dataSourceBaremetalKeypairV2Attributes(d, &keypair)
}

// dataSourceBaremetalKeypairV2Attributes populates the fields of an Keypair resource.
func dataSourceBaremetalKeypairV2Attributes(d *schema.ResourceData, keypair *keypairs.KeyPair) error {
	log.Printf("[DEBUG] ecl_baremetal_keypair_v2 details: %#v", keypair)

	d.SetId(keypair.Name)
	d.Set("name", keypair.Name)
	d.Set("public_key", keypair.PublicKey)
	d.Set("fingerprint", keypair.Fingerprint)

	return nil
}
