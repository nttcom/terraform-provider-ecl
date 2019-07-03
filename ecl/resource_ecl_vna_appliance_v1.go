package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func allowedAddessPairsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Computed: true,
		Set:      allowedAddressPairHash,
		// Default:  &schema.Set{},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
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
	}
}

func fixedIPsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Computed: true,
		Set:      fixedIPHash,
		// Default:  &schema.Set{},
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"subnet_id": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func interfaceMetaSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		Set:      interfaceHash,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					// Default:  "",
				},
				"description": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					// Default:  "",
				},
				"network_id": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					// Default:  "",
				},
				"updatable": &schema.Schema{
					Type:     schema.TypeBool,
					Computed: true,
				},
				"tags": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
					// Default:  map[string]string{},
				},
			},
		},
	}
}
func resourceVNAApplianceV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceVNAApplianceV1Create,
		Read:   resourceVNAApplianceV1Read,
		Update: resourceVNAApplianceV1Update,
		Delete: resourceVNAApplianceV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
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

			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"virtual_network_appliance_plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

			"interface_1_meta":                  interfaceMetaSchema(),
			"interface_1_fixed_ips":             fixedIPsSchema(),
			"interface_1_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_2_meta":                  interfaceMetaSchema(),
			"interface_2_fixed_ips":             fixedIPsSchema(),
			"interface_2_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_3_meta":                  interfaceMetaSchema(),
			"interface_3_fixed_ips":             fixedIPsSchema(),
			"interface_3_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_4_meta":                  interfaceMetaSchema(),
			"interface_4_fixed_ips":             fixedIPsSchema(),
			"interface_4_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_5_meta":                  interfaceMetaSchema(),
			"interface_5_fixed_ips":             fixedIPsSchema(),
			"interface_5_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_6_meta":                  interfaceMetaSchema(),
			"interface_6_fixed_ips":             fixedIPsSchema(),
			"interface_6_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_7_meta":                  interfaceMetaSchema(),
			"interface_7_fixed_ips":             fixedIPsSchema(),
			"interface_7_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_8_meta":                  interfaceMetaSchema(),
			"interface_8_fixed_ips":             fixedIPsSchema(),
			"interface_8_allowed_address_pairs": allowedAddessPairsSchema(),
		},
	}
}

func resourceVNAApplianceV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	createOpts := VirtualNetworkApplianceCreateOpts{
		appliances.CreateOpts{
			Name:                          d.Get("name").(string),
			Description:                   d.Get("description").(string),
			DefaultGateway:                d.Get("default_gateway").(string),
			AvailabilityZone:              d.Get("availability_zone").(string),
			VirtualNetworkAppliancePlanID: d.Get("virtual_network_appliance_plan_id").(string),
			TenantID:                      d.Get("tenant_id").(string),
			Tags:                          resourceTags(d),
			Interfaces:                    getInterfaceCreateOpts(d),
		},
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	vna, err := appliances.Create(vnaClient, createOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance: %s", err)
	}

	d.SetId(vna.ID)
	log.Printf("[INFO] Virtual Network Appliance ID: %s", vna.ID)

	log.Printf("[DEBUG] Waiting for Virtual Network Appliance (%s) to become available", vna.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForVirtualNetworkApplianceComplete(vnaClient, vna.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: 30 * time.Second,
		MinTimeout:   10 * time.Second,
	}

	_, err = stateConf.WaitForState()

	return resourceVNAApplianceV1Read(d, meta)
}

func resourceVNAApplianceV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	var vna appliances.Appliance
	err = appliances.Get(vnaClient, d.Id()).ExtractInto(&vna)
	if err != nil {
		return CheckDeleted(d, err, "virtual-network-appliance")
	}

	log.Printf("[DEBUG] Retrieved Virtual Network Appliance %s: %+v", d.Id(), vna)
	log.Printf("[MYDEBUG] VNA: %#v", vna)

	d.Set("name", vna.Name)
	d.Set("description", vna.Description)
	d.Set("default_gateway", vna.DefaultGateway)
	d.Set("availability_zone", vna.AvailabilityZone)
	d.Set("virtual_network_appliance_plan_id", vna.AppliancePlanID)
	d.Set("tenant_id", vna.TenantID)
	d.Set("tags", vna.Tags)

	for i := 1; i <= maxNumberOfInterfaces; i++ {
		targetMeta := getInterfaceBySlotNumber(&vna, i)
		targetFIPs := getFixedIPsBySlotNumber(&vna, i)
		targetAAPs := getAllowedAddressPairsBySlotNumber(&vna, i)

		d.Set(
			fmt.Sprintf("interface_%d_meta", i),
			getInterfaceMetaAsState(targetMeta))

		d.Set(
			fmt.Sprintf("interface_%d_fixed_ips", i),
			getInterfaceFixedIPsAsState(targetFIPs))

		d.Set(
			fmt.Sprintf("interface_%d_allowed_address_pairs", i),
			getInterfaceAllowedAddressPairsAsState(targetAAPs))
	}

	return nil
}

func resourceVNAApplianceV1Update(d *schema.ResourceData, meta interface{}) error {
	// config := meta.(*Config)
	// vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	// if err != nil {
	// 	return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	// }

	// var updateOpts appliances.UpdateOpts
	// if d.HasChange("name") {
	// 	name := d.Get("name").(string)
	// 	updateOpts.Name = &name
	// }
	// if v, ok := d.GetOkExists("admin_state_up"); ok {
	// 	asu := v.(bool)
	// 	updateOpts.AdminStateUp = &asu
	// }
	// if d.HasChange("description") {
	// 	description := d.Get("description").(string)
	// 	updateOpts.Description = &description
	// }

	// if d.HasChange("tags") {
	// 	tags := resourceTags(d)
	// 	updateOpts.Tags = &tags
	// }

	// log.Printf("[DEBUG] Updating Network %s with options: %+v", d.Id(), updateOpts)
	// _, err = appliances.Update(vnaClient, d.Id(), updateOpts).Extract()

	// if err != nil {
	// 	return fmt.Errorf("Error updating ECL Neutron Network: %s", err)
	// }

	return resourceVNAApplianceV1Read(d, meta)
}

func resourceVNAApplianceV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	err = appliances.Delete(vnaClient, d.Id()).ExtractErr()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"PROCESSING"},
		Target:     []string{"DELETED"},
		Refresh:    waitForVirtualNetworkApplianceDelete(vnaClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Network: %s", err)
	}

	d.SetId("")
	return nil
}

func waitForVirtualNetworkApplianceComplete(vnaClient *eclcloud.ServiceClient, virtualNetworkApplianceID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := appliances.Get(vnaClient, virtualNetworkApplianceID).Extract()
		if err != nil {
			return nil, "", err
		}

		return n, n.OperationStatus, nil
	}
}

func waitForVirtualNetworkApplianceDelete(vnaClient *eclcloud.ServiceClient, virtualNetworkApplianceID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Virtual Network Appliance %s.\n", virtualNetworkApplianceID)

		n, err := appliances.Get(vnaClient, virtualNetworkApplianceID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Virtual Network Appliance %s",
					virtualNetworkApplianceID)
				return n, "DELETED", nil
			}
			return n, "PROCESSING", err
		}

		log.Printf("[DEBUG] ECL Virtual Network Appliance %s still active.\n", virtualNetworkApplianceID)
		return n, "PROCESSING", nil
	}
}
