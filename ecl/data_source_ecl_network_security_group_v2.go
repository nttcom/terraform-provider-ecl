package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_groups"
)

func dataSourceNetworkSecurityGroupV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkSecurityGroupV2Read,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
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
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
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

func dataSourceNetworkSecurityGroupV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	listOpts := security_groups.ListOpts{
		Description: d.Get("description").(string),
		ID:          d.Get("security_group_id").(string),
		Name:        d.Get("name").(string),
		Status:      d.Get("status").(string),
		TenantID:    d.Get("tenant_id").(string),
	}

	pages, err := security_groups.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve security groups: %s", err)
	}

	allSecurityGroups, err := security_groups.ExtractSecurityGroups(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract security groups: %s", err)
	}

	if len(allSecurityGroups) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allSecurityGroups) > 1 {
		return fmt.Errorf("Your query returned more than one result. " +
			"Please try a more specific search criteria.")
	}

	sg := allSecurityGroups[0]

	log.Printf("[DEBUG] Retrieved Security Group %s: %+v", sg.ID, sg)
	d.SetId(sg.ID)

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
