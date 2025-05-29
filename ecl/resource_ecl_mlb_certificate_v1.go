package ecl

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/certificates"
)

func certificateCertFileSchemaForResource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"content": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
			},
		},
	}
}

func certificateKeyFileSchemaForResource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"content": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
				"passphrase": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					ForceNew: true,
				},
			},
		},
	}
}

func resourceMLBCertificateV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceMLBCertificateV1Create,
		Read:   resourceMLBCertificateV1Read,
		Update: resourceMLBCertificateV1Update,
		Delete: resourceMLBCertificateV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ca_cert":  certificateCertFileSchemaForResource(),
			"ssl_cert": certificateCertFileSchemaForResource(),
			"ssl_key":  certificateKeyFileSchemaForResource(),
		},
	}

	return result
}

func resourceMLBCertificateV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOpts := certificates.CreateOpts{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        d.Get("tags").(map[string]interface{}),
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Creating ECL managed load balancer certificate with options %+v", createOpts)

	certificate, err := certificates.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer certificate with options %+v: %s", createOpts, err)
	}

	d.SetId(certificate.ID)
	log.Printf("[INFO] ECL managed load balancer certificate ID: %s", certificate.ID)

	for _, fileType := range []string{"ca_cert", "ssl_cert", "ssl_key"} {
		file := d.Get(fileType).(map[string]interface{})
		uploadFileOpts := certificates.UploadFileOpts{
			Type:    strings.Replace(fileType, "_", "-", -1),
			Content: file["content"].(string),
		}

		if fileType == "ssl_key" {
			if passphrase, ok := file["passphrase"].(string); ok {
				uploadFileOpts.Passphrase = passphrase
			}
		}

		log.Printf("[DEBUG] Uploading ECL managed load balancer certificate file (%s) with options %+v", d.Id(), uploadFileOpts)

		err = certificates.UploadFile(managedLoadBalancerClient, certificate.ID, uploadFileOpts).ExtractErr()
		if err != nil {
			return fmt.Errorf("Error uploading ECL managed load balancer certificate file (%s) with options %+v: %s", d.Id(), uploadFileOpts, err)
		}
	}

	return resourceMLBCertificateV1Read(d, meta)
}

func resourceMLBCertificateV1Read(d *schema.ResourceData, meta interface{}) error {
	var certificate certificates.Certificate

	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	err = certificates.Show(managedLoadBalancerClient, d.Id()).ExtractInto(&certificate)
	if err != nil {
		return CheckDeleted(d, err, "certificate")
	}

	log.Printf("[DEBUG] Retrieved ECL managed load balancer certificate (%s): %+v", d.Id(), certificate)

	d.Set("name", certificate.Name)
	d.Set("description", certificate.Description)
	d.Set("tags", certificate.Tags)
	d.Set("tenant_id", certificate.TenantID)

	return nil
}

func resourceMLBCertificateV1Update(d *schema.ResourceData, meta interface{}) error {
	var isUpdated bool
	var updateOpts certificates.UpdateOpts

	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	if d.HasChange("name") {
		isUpdated = true
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		isUpdated = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("tags") {
		isUpdated = true
		tags := d.Get("tags").(map[string]interface{})
		updateOpts.Tags = &tags
	}

	if isUpdated {
		log.Printf("[DEBUG] Updating ECL managed load balancer certificate (%s) with options %+v", d.Id(), updateOpts)

		_, err := certificates.Update(managedLoadBalancerClient, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL managed load balancer certificate (%s) with options %+v: %s", d.Id(), updateOpts, err)
		}
	}

	return resourceMLBCertificateV1Read(d, meta)
}

func resourceMLBCertificateV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Deleting ECL managed load balancer certificate (%s)", d.Id())

	err = certificates.Delete(managedLoadBalancerClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer certificate (%s)", d.Id())
			return nil
		}
		return fmt.Errorf("Error deleting ECL managed load balancer certificate: %s", err)
	}

	return nil
}
