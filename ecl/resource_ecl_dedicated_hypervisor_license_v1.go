package ecl

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/dedicated_hypervisor/v1/licenses"
)

var licenseNotFoundError = errors.New("license not found")

func resourceDedicatedHypervisorLicenseV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceDedicatedHypervisorLicenseV1Create,
		Read:   resourceDedicatedHypervisorLicenseV1Read,
		Delete: resourceDedicatedHypervisorLicenseV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assigned_from": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDedicatedHypervisorLicenseV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.dedicatedHypervisorV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor client: %s", err)
	}

	opts := licenses.CreateOpts{
		LicenseType: d.Get("license_type").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", opts)
	license, err := licenses.Create(client, opts).ExtractLicenseInfo()
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor license: %s", err)
	}
	d.SetId(license.ID)

	log.Printf("[DEBUG] Created ECL Dedicated Hypervisor license %s: %#v", license.ID, license)
	return resourceDedicatedHypervisorLicenseV1Read(d, meta)
}

func resourceDedicatedHypervisorLicenseV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.dedicatedHypervisorV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor client: %s", err)
	}

	license, err := getLicense(client, d.Id())
	if err != nil {
		if err == licenseNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("license_type", license.LicenseType)
	d.Set("key", license.Key)
	d.Set("assigned_from", license.AssignedFrom.Format(time.RFC3339))
	if license.ExpiresAt != nil {
		d.Set("expires_at", license.ExpiresAt.Format(time.RFC3339))
	}

	log.Printf("[DEBUG] Retrieved Dedicated Hypervisor license %s: %#v", d.Id(), license)
	return nil
}

func resourceDedicatedHypervisorLicenseV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.dedicatedHypervisorV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor client: %s", err)
	}

	if err := licenses.Delete(client, d.Id()).ExtractErr(); err != nil {
		return fmt.Errorf("error deleting ECL Dedicated Hypervisor license: %s", err)
	}

	d.SetId("")
	return nil
}

func getLicense(client *eclcloud.ServiceClient, id string) (*licenses.License, error) {
	pages, err := licenses.List(client, nil).AllPages()
	if err != nil {
		return nil, fmt.Errorf("error getting ECL Dedicated Hypervisor licenses: %s", err)
	}

	ls, err := licenses.ExtractLicenses(pages)
	if err != nil {
		return nil, fmt.Errorf("error getting ECL Dedicated Hypervisor licenses: %s", err)
	}

	for _, l := range ls {
		if l.ID == id {
			return &l, nil
		}
	}

	return nil, licenseNotFoundError
}
