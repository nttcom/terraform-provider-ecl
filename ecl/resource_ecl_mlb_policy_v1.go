package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/policies"
)

func resourceMLBPolicyV1() *schema.Resource {
	var result *schema.Resource

	result = &schema.Resource{
		Create: resourceMLBPolicyV1Create,
		Read:   resourceMLBPolicyV1Read,
		Update: resourceMLBPolicyV1Update,
		Delete: resourceMLBPolicyV1Delete,
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
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sorry_page_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_monitor_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"listener_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"default_target_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tls_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"load_balancer_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}

	return result
}

func resourceMLBPolicyV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	createOpts := policies.CreateOpts{
		Name:                 d.Get("name").(string),
		Description:          d.Get("description").(string),
		Tags:                 d.Get("tags").(map[string]interface{}),
		Algorithm:            d.Get("algorithm").(string),
		Persistence:          d.Get("persistence").(string),
		IdleTimeout:          d.Get("idle_timeout").(int),
		SorryPageUrl:         d.Get("sorry_page_url").(string),
		SourceNat:            d.Get("source_nat").(string),
		CertificateID:        d.Get("certificate_id").(string),
		HealthMonitorID:      d.Get("health_monitor_id").(string),
		ListenerID:           d.Get("listener_id").(string),
		DefaultTargetGroupID: d.Get("default_target_group_id").(string),
		TLSPolicyID:          d.Get("tls_policy_id").(string),
		LoadBalancerID:       d.Get("load_balancer_id").(string),
	}

	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Creating ECL managed load balancer policy with options %+v", createOpts)

	policy, err := policies.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer policy with options %+v: %s", createOpts, err)
	}

	d.SetId(policy.ID)
	log.Printf("[INFO] ECL managed load balancer policy ID: %s", policy.ID)

	return resourceMLBPolicyV1Read(d, meta)
}

func resourceMLBPolicyV1Show(d *schema.ResourceData, client *eclcloud.ServiceClient, changes bool) (*policies.Policy, error) {
	var policy policies.Policy

	showOpts := policies.ShowOpts{Changes: changes}
	err := policies.Show(client, d.Id(), showOpts).ExtractInto(&policy)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve ECL managed load balancer policy (%s): %s", d.Id(), err)
	}

	return &policy, nil
}

func resourceMLBPolicyV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	policy, err := resourceMLBPolicyV1Show(d, managedLoadBalancerClient, true)
	if err != nil {
		return CheckDeleted(d, err, "policy")
	}

	if policy.ConfigurationStatus == "ACTIVE" {
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
	} else if policy.ConfigurationStatus == "CREATE_STAGED" {
		d.Set("algorithm", policy.Staged.Algorithm)
		d.Set("persistence", policy.Staged.Persistence)
		d.Set("idle_timeout", policy.Staged.IdleTimeout)
		d.Set("sorry_page_url", policy.Staged.SorryPageUrl)
		d.Set("source_nat", policy.Staged.SourceNat)
		d.Set("certificate_id", policy.Staged.CertificateID)
		d.Set("health_monitor_id", policy.Staged.HealthMonitorID)
		d.Set("listener_id", policy.Staged.ListenerID)
		d.Set("default_target_group_id", policy.Staged.DefaultTargetGroupID)
		d.Set("tls_policy_id", policy.Staged.TLSPolicyID)
	} else if policy.ConfigurationStatus == "UPDATE_STAGED" {
		d.Set("algorithm", ternary(policy.Staged.Algorithm == "", policy.Algorithm, policy.Staged.Algorithm))
		d.Set("persistence", ternary(policy.Staged.Persistence == "", policy.Persistence, policy.Staged.Persistence))
		d.Set("idle_timeout", ternary(policy.Staged.IdleTimeout == 0, policy.IdleTimeout, policy.Staged.IdleTimeout))
		d.Set("sorry_page_url", ternary(policy.Staged.SorryPageUrl == "", policy.SorryPageUrl, policy.Staged.SorryPageUrl))
		d.Set("source_nat", ternary(policy.Staged.SourceNat == "", policy.SourceNat, policy.Staged.SourceNat))
		d.Set("certificate_id", ternary(policy.Staged.CertificateID == "", policy.CertificateID, policy.Staged.CertificateID))
		d.Set("health_monitor_id", ternary(policy.Staged.HealthMonitorID == "", policy.HealthMonitorID, policy.Staged.HealthMonitorID))
		d.Set("listener_id", ternary(policy.Staged.ListenerID == "", policy.ListenerID, policy.Staged.ListenerID))
		d.Set("default_target_group_id", ternary(policy.Staged.DefaultTargetGroupID == "", policy.DefaultTargetGroupID, policy.Staged.DefaultTargetGroupID))
		d.Set("tls_policy_id", ternary(policy.Staged.TLSPolicyID == "", policy.TLSPolicyID, policy.Staged.TLSPolicyID))
	} else if policy.ConfigurationStatus == "DELETE_STAGED" {
		d.SetId("")
		return nil
	}

	d.Set("name", policy.Name)
	d.Set("description", policy.Description)
	d.Set("tags", policy.Tags)
	d.Set("load_balancer_id", policy.LoadBalancerID)
	d.Set("tenant_id", policy.TenantID)

	return nil
}

func resourceMLBPolicyV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Start updating attributes of ECL managed load balancer policy ...")

	err = resourceMLBPolicyV1UpdateAttributes(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating attributes of ECL managed load balancer policy: %s", err)
	}

	log.Printf("[DEBUG] Start updating configurations of ECL managed load balancer policy ...")

	err = resourceMLBPolicyV1UpdateConfigurations(d, managedLoadBalancerClient)
	if err != nil {
		return fmt.Errorf("Error in updating configurations of ECL managed load balancer policy: %s", err)
	}

	return resourceMLBPolicyV1Read(d, meta)
}

func resourceMLBPolicyV1UpdateAttributes(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isAttributesUpdated bool
	var updateOpts policies.UpdateOpts

	if d.HasChange("name") {
		isAttributesUpdated = true
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		isAttributesUpdated = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("tags") {
		isAttributesUpdated = true
		tags := d.Get("tags").(map[string]interface{})
		updateOpts.Tags = &tags
	}

	if isAttributesUpdated {
		log.Printf("[DEBUG] Updating ECL managed load balancer policy attributes (%s) with options %+v", d.Id(), updateOpts)

		_, err := policies.Update(client, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL managed load balancer policy attributes (%s) with options %+v: %s", d.Id(), updateOpts, err)
		}
	}

	return nil
}

func resourceMLBPolicyV1UpdateConfigurations(d *schema.ResourceData, client *eclcloud.ServiceClient) error {
	var isConfigurationsUpdated bool

	policy, err := resourceMLBPolicyV1Show(d, client, false)
	if err != nil {
		return err
	}

	if policy.ConfigurationStatus == "ACTIVE" {
		var createStagedOpts policies.CreateStagedOpts

		if d.HasChange("algorithm") {
			isConfigurationsUpdated = true
			createStagedOpts.Algorithm = d.Get("algorithm").(string)
		}

		if d.HasChange("persistence") {
			isConfigurationsUpdated = true
			createStagedOpts.Persistence = d.Get("persistence").(string)
		}

		if d.HasChange("idle_timeout") {
			isConfigurationsUpdated = true
			createStagedOpts.IdleTimeout = d.Get("idle_timeout").(int)
		}

		if d.HasChange("sorry_page_url") {
			isConfigurationsUpdated = true
			createStagedOpts.SorryPageUrl = d.Get("sorry_page_url").(string)
		}

		if d.HasChange("source_nat") {
			isConfigurationsUpdated = true
			createStagedOpts.SourceNat = d.Get("source_nat").(string)
		}

		if d.HasChange("certificate_id") {
			isConfigurationsUpdated = true
			createStagedOpts.CertificateID = d.Get("certificate_id").(string)
		}

		if d.HasChange("health_monitor_id") {
			isConfigurationsUpdated = true
			createStagedOpts.HealthMonitorID = d.Get("health_monitor_id").(string)
		}

		if d.HasChange("listener_id") {
			isConfigurationsUpdated = true
			createStagedOpts.ListenerID = d.Get("listener_id").(string)
		}

		if d.HasChange("default_target_group_id") {
			isConfigurationsUpdated = true
			createStagedOpts.DefaultTargetGroupID = d.Get("default_target_group_id").(string)
		}

		if d.HasChange("tls_policy_id") {
			isConfigurationsUpdated = true
			createStagedOpts.TLSPolicyID = d.Get("tls_policy_id").(string)
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer policy configurations (%s) with options %+v", d.Id(), createStagedOpts)

			_, err := policies.CreateStaged(client, d.Id(), createStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer policy configurations (%s) with options %+v: %s", d.Id(), createStagedOpts, err)
			}
		}
	} else {
		var updateStagedOpts policies.UpdateStagedOpts

		if d.HasChange("algorithm") {
			isConfigurationsUpdated = true
			algorithm := d.Get("algorithm").(string)
			updateStagedOpts.Algorithm = &algorithm
		}

		if d.HasChange("persistence") {
			isConfigurationsUpdated = true
			persistence := d.Get("persistence").(string)
			updateStagedOpts.Persistence = &persistence
		}

		if d.HasChange("idle_timeout") {
			isConfigurationsUpdated = true
			idleTimeout := d.Get("idle_timeout").(int)
			updateStagedOpts.IdleTimeout = &idleTimeout
		}

		if d.HasChange("sorry_page_url") {
			isConfigurationsUpdated = true
			sorryPageUrl := d.Get("sorry_page_url").(string)
			updateStagedOpts.SorryPageUrl = &sorryPageUrl
		}

		if d.HasChange("source_nat") {
			isConfigurationsUpdated = true
			sourceNat := d.Get("source_nat").(string)
			updateStagedOpts.SourceNat = &sourceNat
		}

		if d.HasChange("certificate_id") {
			isConfigurationsUpdated = true
			certificateID := d.Get("certificate_id").(string)
			updateStagedOpts.CertificateID = &certificateID
		}

		if d.HasChange("health_monitor_id") {
			isConfigurationsUpdated = true
			healthMonitorID := d.Get("health_monitor_id").(string)
			updateStagedOpts.HealthMonitorID = &healthMonitorID
		}

		if d.HasChange("listener_id") {
			isConfigurationsUpdated = true
			listenerID := d.Get("listener_id").(string)
			updateStagedOpts.ListenerID = &listenerID
		}

		if d.HasChange("default_target_group_id") {
			isConfigurationsUpdated = true
			defaultTargetGroupID := d.Get("default_target_group_id").(string)
			updateStagedOpts.DefaultTargetGroupID = &defaultTargetGroupID
		}

		if d.HasChange("tls_policy_id") {
			isConfigurationsUpdated = true
			tlsPolicyID := d.Get("tls_policy_id").(string)
			updateStagedOpts.TLSPolicyID = &tlsPolicyID
		}

		if isConfigurationsUpdated {
			log.Printf("[DEBUG] Updating ECL managed load balancer policy configurations (%s) with options %+v", d.Id(), updateStagedOpts)

			_, err := policies.UpdateStaged(client, d.Id(), updateStagedOpts).Extract()
			if err != nil {
				return fmt.Errorf("Error updating ECL managed load balancer policy configurations (%s) with options %+v: %s", d.Id(), updateStagedOpts, err)
			}
		}
	}

	return nil
}

func resourceMLBPolicyV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	managedLoadBalancerClient, err := config.managedLoadBalancerV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL managed load balancer client: %s", err)
	}

	log.Printf("[DEBUG] Deleting ECL managed load balancer policy: %s", d.Id())

	err = policies.Delete(managedLoadBalancerClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer policy (%s)", d.Id())
			return nil
		}

		policy, err := resourceMLBPolicyV1Show(d, managedLoadBalancerClient, false)
		if err != nil {
			return err
		}
		if policy.ConfigurationStatus == "DELETE_STAGED" {
			log.Printf("[DEBUG] Already deleted ECL managed load balancer policy (%s)", d.Id())
			return nil
		}

		return fmt.Errorf("Error deleting ECL managed load balancer policy: %s", err)
	}

	return nil
}
