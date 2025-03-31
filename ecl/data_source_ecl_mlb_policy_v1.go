package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/policies"
)

func dataSourceMLBPolicyV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Read: dataSourceMLBPolicyV1Read,
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
			"configuration_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"operation_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"load_balancer_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sorry_page_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health_monitor_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listener_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_target_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}

	return result
}

func dataSourceMLBPolicyV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	listOpts := policies.ListOpts{}

	if v, ok := d.GetOk("id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("configuration_status"); ok {
		listOpts.ConfigurationStatus = v.(string)
	}

	if v, ok := d.GetOk("operation_status"); ok {
		listOpts.OperationStatus = v.(string)
	}

	if v, ok := d.GetOk("algorithm"); ok {
		listOpts.Algorithm = v.(string)
	}

	if v, ok := d.GetOk("persistence"); ok {
		listOpts.Persistence = v.(string)
	}

	if v, ok := d.GetOk("idle_timeout"); ok {
		listOpts.IdleTimeout = v.(int)
	}

	if v, ok := d.GetOk("sorry_page_url"); ok {
		listOpts.SorryPageUrl = v.(string)
	}

	if v, ok := d.GetOk("source_nat"); ok {
		listOpts.SourceNat = v.(string)
	}

	if v, ok := d.GetOk("certificate_id"); ok {
		listOpts.CertificateID = v.(string)
	}

	if v, ok := d.GetOk("health_monitor_id"); ok {
		listOpts.HealthMonitorID = v.(string)
	}

	if v, ok := d.GetOk("listener_id"); ok {
		listOpts.ListenerID = v.(string)
	}

	if v, ok := d.GetOk("default_target_group_id"); ok {
		listOpts.DefaultTargetGroupID = v.(string)
	}

	if v, ok := d.GetOk("tls_policy_id"); ok {
		listOpts.TLSPolicyID = v.(string)
	}

	if v, ok := d.GetOk("load_balancer_id"); ok {
		listOpts.LoadBalancerID = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Retrieving ECL managed load balancer policies with options %+v", listOpts)

	pages, err := policies.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allPolicies, err := policies.ExtractPolicies(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ECL managed load balancer policies with options %+v: %s", listOpts, err)
	}

	if len(allPolicies) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allPolicies) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	policy := allPolicies[0]

	log.Printf("[DEBUG] Retrieved ECL managed load balancer policy: %+v", policy)

	d.SetId(policy.ID)

	d.Set("name", policy.Name)
	d.Set("description", policy.Description)
	d.Set("tags", policy.Tags)
	d.Set("configuration_status", policy.ConfigurationStatus)
	d.Set("operation_status", policy.OperationStatus)
	d.Set("load_balancer_id", policy.LoadBalancerID)
	d.Set("tenant_id", policy.TenantID)
	d.Set("algorithm", policy.Algorithm)
	d.Set("persistence", policy.Persistence)
	d.Set("idle_timeout", policy.IdleTimeout)
	d.Set("sorry_page_url", policy.SorryPageUrl)
	d.Set("source_nat", policy.SourceNat)
	d.Set("certificate_id", policy.CertificateID)
	d.Set("health_monitor_id", policy.HealthMonitorID)
	d.Set("listener_id", policy.ListenerID)
	d.Set("default_target_group_id", policy.DefaultTargetGroupID)
	d.Set("tls_policy_id", policy.TLSPolicyID)

	return nil
}
