package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func dataSourceVNAApplianceV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNAApplianceV1Read,

		Schema: map[string]*schema.Schema{

			"virtual_network_appliance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"virtual_network_appliance_plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"interfaces": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      interfaceHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"slot_number": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"network_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"tags": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
						},

						"fixed_ips": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"subnet_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"allowed_address_pairs": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},

									"vrid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVNAApplianceV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	listOpts := appliances.ListOpts{}

	if v, ok := d.GetOk("virtual_network_appliance_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	// TODO maintenance followings later
	// if v, ok := d.GetOk("appliance_type"); ok {
	// 	listOpts.Name = v.(string)
	// }

	// if v, ok := d.GetOk("description"); ok {
	// 	listOpts.Description = v.(string)
	// }

	// if v, ok := d.GetOk("id"); ok {
	// 	listOpts.ID = v.(string)
	// }

	// if v, ok := d.GetOk("network_id"); ok {
	// 	listOpts.NetworkID = v.(string)
	// }

	// if v, ok := d.GetOk("status"); ok {
	// 	listOpts.Status = v.(string)
	// }

	// if v, ok := d.GetOk("subnet_id"); ok {
	// 	listOpts.SubnetID = v.(string)
	// }

	// if v, ok := d.GetOk("tenant_id"); ok {
	// 	listOpts.TenantID = v.(string)
	// }

	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	pages, err := appliances.List(vnaClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allAppliances, err := appliances.ExtractAppliances(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve virtual network appliances: %s", err)
	}

	if len(allAppliances) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allAppliances) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	vna := allAppliances[0]

	log.Printf("[DEBUG] Retrieved Virtual Network Appliance %s: %+v", d.Id(), vna)

	d.SetId(vna.ID)

	d.Set("name", vna.Name)
	d.Set("description", vna.Description)
	d.Set("default_gateway", vna.DefaultGateway)
	d.Set("availability_zone", vna.AvailabilityZone)
	d.Set("virtual_network_appliance_plan_id", vna.AppliancePlanID)
	d.Set("tenant_id", vna.TenantID)
	d.Set("tags", vna.Tags)
	// d.Set("interfaces", convertApplianceInterfacesFromStructToMap(vna.Interfaces))

	return nil
}
