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

const createPollInterval = 2 * time.Second
const updatePollInterval = 2 * time.Second
const deletePollInterval = 2 * time.Second

func allowedAddessPairsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		// Set:      allowedAddressPairHash,
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
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		// Set:      fixedIPHash,
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
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		// Set:      interfaceHash,
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
	log.Printf("[DEBUG] Waiting for Virtual Network Appliance (%s) to become COMPLETE", vna.ID)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForVirtualNetworkApplianceComplete(vnaClient, vna.ID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: createPollInterval,
		MinTimeout:   10 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for virtual network appliance (%s) to become COMPLETE: %s",
			vna.ID, err)
	}

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

// TODO remove following comments later.
/*
Update samples

Update Metadata:
	Request

		PATCH /v1.0/virtual_network_appliances/[vna.ID]
		{
			"virtual_network_appliance": {
				"description": "appliance_1_description-update",
				"name": "appliance_1-update",
				"tags": {
					"k2": "v2", <-- it is okay to send all tags even difference is only this key-value pair.
					"k1": "v1"
				}
			}
		}

Update Interface:
	Request

	interface2: Changing from no-connection to connect with auto assigned IP addresses
	interface3: Changing from no-connection to 2 fixed IPs as 192.168.3.50 and .60

		PATCH /v1.0/virtual_network_appliances/[vna.ID]
		{
			"virtual_network_appliance": {
				"interfaces": {
					"interface_3": {
						"network_id": "989c8daf-9769-4c3a-8aec-5d1744ce5787",
						"fixed_ips": [{
							"ip_address": "192.168.3.50"
						}, {
							"ip_address": "192.168.3.60"
						}]
					},
					"interface_2": {
						"network_id": "e9e3c929-331b-4e4c-b182-53dd26472411"
					}
				}
			}
		}
*/
func resourceVNAApplianceV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	var updateMeta, updateInterface, updateAAP bool

	var updateOptsMeta appliances.UpdateOptsMeta
	if d.HasChange("name") {
		updateMeta = true

		name := d.Get("name").(string)
		a, b := d.GetChange("name")
		log.Printf("[MYDEBUG:name] %#v %#v", a, b)
		updateOptsMeta.Name = &name
	}

	if d.HasChange("description") {
		updateMeta = true

		description := d.Get("description").(string)
		a, b := d.GetChange("description")
		log.Printf("[MYDEBUG:description] %#v %#v", a, b)
		updateOptsMeta.Description = &description
	}

	if d.HasChange("tags") {
		updateMeta = true

		tags := resourceTags(d)
		a, b := d.GetChange("tags")
		log.Printf("[MYDEBUG:tags] %#v %#v", a, b)
		updateOptsMeta.Tags = &tags
	}

	if updateMeta {
		log.Printf("[DEBUG] Updating VNA Metadata %s with options: %+v", d.Id(), updateOptsMeta)
		_, err = appliances.Update(vnaClient, d.Id(), updateOptsMeta).Extract()

		stateConf := &resource.StateChangeConf{
			Pending:      []string{"PROCESSING"},
			Target:       []string{"COMPLETE"},
			Refresh:      waitForVirtualNetworkApplianceComplete(vnaClient, d.Id()),
			Timeout:      d.Timeout(schema.TimeoutDelete),
			Delay:        5 * time.Second,
			PollInterval: updatePollInterval,
			MinTimeout:   3 * time.Second,
		}

		_, err = stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf(
				"Error waiting for virtual network appliance (%s) to become COMPLETE: %s",
				d.Id(), err)
		}
	}
	// [MYDEBUG:tags] map[string]interface {}{"k1":"v1"} map[string]interface {}{"k1":"v1", "k2":"v2"}

	// updateOptsInterface := appliances.UpdateOptsMeta{}
	var slotNumbers = []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, slotNumber := range slotNumbers {
		log.Printf("[MYDEBUG] slotNumber: %d", slotNumber)
		ifMetaKey := fmt.Sprintf("interface_%d_meta", slotNumber)
		// ifFixedIPsKey := fmt.Sprintf("interface_%d_fixed_ips", slotNumber)

		log.Printf("[MYDEBUG] d.HasChange(ifMetaKey): %#v", d.HasChange(ifMetaKey))
		if d.HasChange(ifMetaKey) {
			updateInterface = true

			b, a := d.GetChange(ifMetaKey)
			before := b.(*schema.Set).List()
			after := a.(*schema.Set).List()
			log.Printf("[MYDEBUG] before: %#v", before)
			log.Printf("[MYDEBUG] after: %#v", after)
		}

	}

	if updateInterface {
	}
	if updateAAP {
	}

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
		Delay:      deletePollInterval,
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
