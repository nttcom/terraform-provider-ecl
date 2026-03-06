package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_groups"
)

func resourceNetworkSecurityGroupV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityGroupV2Create,
		Read:   resourceNetworkSecurityGroupV2Read,
		Update: resourceNetworkSecurityGroupV2Update,
		Delete: resourceNetworkSecurityGroupV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Computed:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_group_rules": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ethertype": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_range_max": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"port_range_min": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_group_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_ip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_group_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkSecurityGroupV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	createOpts := SecurityGroupCreateOpts{
		security_groups.CreateOpts{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			TenantID:    d.Get("tenant_id").(string),
			Tags:        resourceTags(d),
		},
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	sg, err := security_groups.Create(networkClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL Security Group: %s", err)
	}

	d.SetId(sg.ID)
	log.Printf("[INFO] Security Group ID: %s", sg.ID)

	log.Printf("[DEBUG] Waiting for Security Group (%s) to become available", sg.ID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_CREATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForSecurityGroupActive(networkClient, sg.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for Security Group (%s) to become ready: %s", sg.ID, err)
	}

	return resourceNetworkSecurityGroupV2Read(d, meta)
}

func resourceNetworkSecurityGroupV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	sg, err := security_groups.Get(networkClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "security_group")
	}

	log.Printf("[DEBUG] Retrieved Security Group %s: %+v", d.Id(), sg)

	d.Set("name", sg.Name)
	d.Set("description", sg.Description)
	d.Set("tenant_id", sg.TenantID)
	d.Set("status", sg.Status)
	d.Set("tags", sg.Tags)
	d.Set("region", GetRegion(d, config))

	// Set security_group_rules
	rules := make([]map[string]interface{}, len(sg.SecurityGroupRules))
	for i, rule := range sg.SecurityGroupRules {
		ruleMap := make(map[string]interface{})
		ruleMap["id"] = rule.ID
		ruleMap["description"] = rule.Description
		ruleMap["direction"] = rule.Direction
		ruleMap["ethertype"] = rule.Ethertype
		ruleMap["protocol"] = rule.Protocol
		ruleMap["security_group_id"] = rule.SecurityGroupID
		ruleMap["tenant_id"] = rule.TenantID

		if rule.PortRangeMax != nil {
			ruleMap["port_range_max"] = *rule.PortRangeMax
		}
		if rule.PortRangeMin != nil {
			ruleMap["port_range_min"] = *rule.PortRangeMin
		}
		if rule.RemoteGroupID != nil {
			ruleMap["remote_group_id"] = *rule.RemoteGroupID
		}
		if rule.RemoteIPPrefix != nil {
			ruleMap["remote_ip_prefix"] = *rule.RemoteIPPrefix
		}

		rules[i] = ruleMap
	}
	d.Set("security_group_rules", rules)

	return nil
}

func resourceNetworkSecurityGroupV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	var updateOpts security_groups.UpdateOpts

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("tags") {
		tags := resourceTags(d)
		updateOpts.Tags = &tags
	}

	log.Printf("[DEBUG] Updating Security Group %s with options: %+v", d.Id(), updateOpts)
	_, err = security_groups.Update(networkClient, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating ECL Security Group: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PENDING_UPDATE"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForSecurityGroupActive(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for Security Group (%s) to become ready: %s", d.Id(), err)
	}

	return resourceNetworkSecurityGroupV2Read(d, meta)
}

func resourceNetworkSecurityGroupV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	err = security_groups.Delete(networkClient, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Security Group: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForSecurityGroupDelete(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Security Group: %s", err)
	}

	d.SetId("")
	return nil
}

func waitForSecurityGroupActive(networkClient *eclcloud.ServiceClient, sgID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		sg, err := security_groups.Get(networkClient, sgID).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Group: %+v", sg)
		return sg, sg.Status, nil
	}
}

func waitForSecurityGroupDelete(networkClient *eclcloud.ServiceClient, sgID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Security Group %s.\n", sgID)

		sg, err := security_groups.Get(networkClient, sgID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Security Group %s", sgID)
				return sg, "DELETED", nil
			}
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Group %s still active.\n", sgID)
		return sg, sg.Status, nil
	}
}
