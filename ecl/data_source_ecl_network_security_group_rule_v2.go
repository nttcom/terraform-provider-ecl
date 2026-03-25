package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_group_rules"
)

func dataSourceNetworkSecurityGroupRuleV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkSecurityGroupRuleV2Read,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},
			"security_group_rule_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ethertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_range_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port_range_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"remote_ip_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: descriptions["tenant_id"],
			},
		},
	}
}

func dataSourceNetworkSecurityGroupRuleV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	listOpts := security_group_rules.ListOpts{
		Description:     d.Get("description").(string),
		Direction:       d.Get("direction").(string),
		Ethertype:       d.Get("ethertype").(string),
		ID:              d.Get("security_group_rule_id").(string),
		Protocol:        d.Get("protocol").(string),
		RemoteGroupID:   d.Get("remote_group_id").(string),
		RemoteIPPrefix:  d.Get("remote_ip_prefix").(string),
		SecurityGroupID: d.Get("security_group_id").(string),
		TenantID:        d.Get("tenant_id").(string),
	}

	if v, ok := d.GetOk("port_range_min"); ok {
		listOpts.PortRangeMin = v.(int)
	}
	if v, ok := d.GetOk("port_range_max"); ok {
		listOpts.PortRangeMax = v.(int)
	}

	pages, err := security_group_rules.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve security group rules: %s", err)
	}

	allRules, err := security_group_rules.ExtractSecurityGroupRules(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract security group rules: %s", err)
	}

	if len(allRules) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allRules) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	rule := allRules[0]

	log.Printf("[DEBUG] Retrieved Security Group Rule %s: %+v", rule.ID, rule)
	d.SetId(rule.ID)

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
