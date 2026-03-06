package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_group_rules"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_groups"
)

func resourceNetworkSecurityGroupRuleV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityGroupRuleV2Create,
		Read:   resourceNetworkSecurityGroupRuleV2Read,
		Delete: resourceNetworkSecurityGroupRuleV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},
			"security_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ingress", "egress",
				}, false),
			},
			"ethertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "IPv4",
				ValidateFunc: validation.StringInSlice([]string{
					"IPv4", "IPv6",
				}, false),
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"port_range_min": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(0, 65535),
			},
			"port_range_max": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(0, 65535),
			},
			"remote_ip_prefix": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"remote_group_id"},
			},
			"remote_group_id": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"remote_ip_prefix"},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func resourceNetworkSecurityGroupRuleV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	sgID := d.Get("security_group_id").(string)

	// Lock the security group to serialize operations on it
	osMutexKV.Lock(sgID)
	defer osMutexKV.Unlock(sgID)

	// Wait for the parent security group to become active before creating rule
	log.Printf("[DEBUG] Waiting for Security Group (%s) to become available before rule creation", sgID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE", "PENDING_UPDATE", "PENDING_DELETE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForSecurityGroupActiveAfterRuleChange(networkClient, sgID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for Security Group (%s) to become ready before rule creation: %s", sgID, err)
	}

	portRangeMin := d.Get("port_range_min").(int)
	portRangeMax := d.Get("port_range_max").(int)

	var createOpts security_group_rules.CreateOpts
	createOpts.Description = d.Get("description").(string)
	createOpts.Direction = d.Get("direction").(string)
	createOpts.Ethertype = d.Get("ethertype").(string)
	createOpts.SecurityGroupID = sgID
	createOpts.TenantID = d.Get("tenant_id").(string)

	if v, ok := d.GetOk("protocol"); ok {
		createOpts.Protocol = v.(string)
	}

	if portRangeMin > 0 {
		createOpts.PortRangeMin = &portRangeMin
	}

	if portRangeMax > 0 {
		createOpts.PortRangeMax = &portRangeMax
	}

	if v, ok := d.GetOk("remote_ip_prefix"); ok {
		remoteIPPrefix := v.(string)
		createOpts.RemoteIPPrefix = &remoteIPPrefix
	}

	if v, ok := d.GetOk("remote_group_id"); ok {
		remoteGroupID := v.(string)
		createOpts.RemoteGroupID = &remoteGroupID
	}

	log.Printf("[DEBUG] Creating Security Group Rule: %#v", createOpts)

	rule, err := security_group_rules.Create(networkClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL Security Group Rule: %s", err)
	}

	d.SetId(rule.ID)

	// Wait for the parent security group to become active after rule creation
	log.Printf("[DEBUG] Waiting for Security Group (%s) to become available after rule creation", sgID)

	stateConf = &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE", "PENDING_UPDATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForSecurityGroupActiveAfterRuleChange(networkClient, sgID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for Security Group (%s) to become ready after rule creation: %s", sgID, err)
	}

	return resourceNetworkSecurityGroupRuleV2Read(d, meta)
}

func resourceNetworkSecurityGroupRuleV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	rule, err := security_group_rules.Get(networkClient, d.Id()).Extract()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error retrieving ECL Security Group Rule: %s", err)
	}

	log.Printf("[DEBUG] Retrieved Security Group Rule %s: %+v", d.Id(), rule)

	d.Set("description", rule.Description)
	d.Set("direction", rule.Direction)
	d.Set("ethertype", rule.Ethertype)
	d.Set("protocol", rule.Protocol)
	d.Set("security_group_id", rule.SecurityGroupID)
	d.Set("tenant_id", rule.TenantID)
	d.Set("region", GetRegion(d, config))

	if rule.PortRangeMin != nil {
		d.Set("port_range_min", *rule.PortRangeMin)
	}
	if rule.PortRangeMax != nil {
		d.Set("port_range_max", *rule.PortRangeMax)
	}
	if rule.RemoteIPPrefix != nil {
		d.Set("remote_ip_prefix", *rule.RemoteIPPrefix)
	}
	if rule.RemoteGroupID != nil {
		d.Set("remote_group_id", *rule.RemoteGroupID)
	}

	return nil
}

func resourceNetworkSecurityGroupRuleV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	sgID := d.Get("security_group_id").(string)

	// Lock the security group to serialize operations on it
	osMutexKV.Lock(sgID)
	defer osMutexKV.Unlock(sgID)

	// Wait for the parent security group to become active before deleting rule
	log.Printf("[DEBUG] Waiting for Security Group (%s) to become available before rule deletion", sgID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE", "PENDING_UPDATE", "PENDING_DELETE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForSecurityGroupActiveAfterRuleChange(networkClient, sgID),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		// If the security group itself was deleted, that's okay
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error waiting for Security Group (%s) to become ready before rule deletion: %s", sgID, err)
	}

	// Delete the rule
	err = security_group_rules.Delete(networkClient, d.Id()).ExtractErr()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			// Rule already deleted
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error deleting ECL Security Group Rule: %s", err)
	}

	// Wait for the parent security group to become active after rule deletion
	log.Printf("[DEBUG] Waiting for Security Group (%s) to become available after rule deletion", sgID)

	stateConf = &resource.StateChangeConf{
		Pending:    []string{"PENDING_UPDATE", "PENDING_DELETE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForSecurityGroupActiveAfterRuleChange(networkClient, sgID),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		// If the security group itself was deleted, that's okay
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error waiting for Security Group (%s) to become ready after rule deletion: %s", sgID, err)
	}

	d.SetId("")
	return nil
}

func waitForSecurityGroupActiveAfterRuleChange(networkClient *eclcloud.ServiceClient, sgID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		sg, err := security_groups.Get(networkClient, sgID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return sg, "DELETED", nil
			}
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Group: %+v", sg)
		return sg, sg.Status, nil
	}
}
