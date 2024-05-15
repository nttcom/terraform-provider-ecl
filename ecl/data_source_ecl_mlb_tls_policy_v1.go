package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/tls_policies"
)

func dataSourceMLBTLSPolicyV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBTLSPolicyV1Read,
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
			"default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"tls_protocols": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tls_ciphers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}

	return result
}

func dataSourceMLBTLSPolicyV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := tls_policies.ListOpts{}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("default"); ok {
		listOpts.Default = v.(bool)
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Retrieving ECL managed load balancer TLS policies with options %+v", listOpts)

	pages, err := tls_policies.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allTLSPolicies, err := tls_policies.ExtractTLSPolicies(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer TLS policies with options %+v: %s", listOpts, err)
	}

	if len(allTLSPolicies) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allTLSPolicies) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	tLSPolicy := allTLSPolicies[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer TLS policy: %+v", tLSPolicy)

	d.SetId(tLSPolicy.ID)

	d.Set("name", tLSPolicy.Name)
	d.Set("description", tLSPolicy.Description)
	d.Set("default", tLSPolicy.Default)
	d.Set("tls_protocols", tLSPolicy.TLSProtocols)
	d.Set("tls_ciphers", tLSPolicy.TLSCiphers)

	return nil
}
