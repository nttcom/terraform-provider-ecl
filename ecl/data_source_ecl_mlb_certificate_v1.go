package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/certificates"
)

func certificateFileSchemaForDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"status": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
			},
		},
	}
}

func dataSourceMLBCertificateV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBCertificateV1Read,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert":  certificateFileSchemaForDataSource(),
			"ssl_cert": certificateFileSchemaForDataSource(),
			"ssl_key":  certificateFileSchemaForDataSource(),
		},
	}

	return result
}

func dataSourceMLBCertificateV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := certificates.ListOpts{}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Retrieving ECL managed load balancer certificates with options %+v", listOpts)

	pages, err := certificates.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allCertificates, err := certificates.ExtractCertificates(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer certificates with options %+v: %s", listOpts, err)
	}

	if len(allCertificates) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allCertificates) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	certificate := allCertificates[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer certificate: %+v", certificate)

	d.SetId(certificate.ID)

	sslKey := make(map[string]interface{})
	sslKey["status"] = certificate.SSLKey.Status

	sslCert := make(map[string]interface{})
	sslCert["status"] = certificate.SSLCert.Status

	caCert := make(map[string]interface{})
	caCert["status"] = certificate.CACert.Status

	d.Set("name", certificate.Name)
	d.Set("description", certificate.Description)
	d.Set("tags", certificate.Tags)
	d.Set("tenant_id", certificate.TenantID)
	d.Set("ca_cert", caCert)
	d.Set("ssl_cert", sslCert)
	d.Set("ssl_key", sslKey)

	return nil
}
