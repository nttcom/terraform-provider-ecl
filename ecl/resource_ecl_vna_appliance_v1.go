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

// const createPollInterval = 30 * time.Second
// const updatePollInterval = 30 * time.Second
// const deletePollInterval = 30 * time.Second

const createPollInterval = 2 * time.Second
const updatePollInterval = 2 * time.Second
const deletePollInterval = 2 * time.Second

func allowedAddessPairsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"mac_address": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
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

func interfaceInfoSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
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
					Optional: true,
				},
				"updatable": &schema.Schema{
					Type:     schema.TypeBool,
					Computed: true,
				},
				"tags": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
					Computed: true,
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
				Computed: true,
			},

			"interface_1_info":                  interfaceInfoSchema(),
			"interface_1_fixed_ips":             fixedIPsSchema(),
			"interface_1_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_2_info":                  interfaceInfoSchema(),
			"interface_2_fixed_ips":             fixedIPsSchema(),
			"interface_2_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_3_info":                  interfaceInfoSchema(),
			"interface_3_fixed_ips":             fixedIPsSchema(),
			"interface_3_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_4_info":                  interfaceInfoSchema(),
			"interface_4_fixed_ips":             fixedIPsSchema(),
			"interface_4_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_5_info":                  interfaceInfoSchema(),
			"interface_5_fixed_ips":             fixedIPsSchema(),
			"interface_5_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_6_info":                  interfaceInfoSchema(),
			"interface_6_fixed_ips":             fixedIPsSchema(),
			"interface_6_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_7_info":                  interfaceInfoSchema(),
			"interface_7_fixed_ips":             fixedIPsSchema(),
			"interface_7_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_8_info":                  interfaceInfoSchema(),
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
			fmt.Sprintf("interface_%d_info", i),
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

	err = updateMetadata(d, meta, vnaClient)
	if err != nil {
		return fmt.Errorf("Error in updating virtual network appliance metadata: %s", err)
	}

	err = updateFixedIP(d, meta, vnaClient)
	if err != nil {
		return fmt.Errorf("Error in updating virtual network appliance interface: %s", err)
	}

	// err = updateAllowedAddressPair(d, meta, vnaClient)
	// if err != nil {
	// 	return fmt.Errorf("Error in updating virtual network appliance allowed address pair: %s", err)
	// }

	return resourceVNAApplianceV1Read(d, meta)
}

// updateFixedIP handles following updates
// - interface_N.network_id
// - interface_N.fixedips list
//  Above updates are correspond to "interface update" in VNA API
func updateFixedIP(d *schema.ResourceData, meta interface{}, client *eclcloud.ServiceClient) error {
	var updated bool
	var updateFixedIPOpts appliances.UpdateFixedIPOpts
	// var updateFixedIPOpts = appliances.UpdateFixedIPOpts{}

	for _, slotNumber := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		updateFixedIPInterface := appliances.UpdateFixedIPInterface{}

		// var networkID string
		networkIDKey := fmt.Sprintf("interface_%d_info.0.network_id", slotNumber)
		if d.HasChange(networkIDKey) {
			updated = true
			networkID := d.Get(networkIDKey).(string)
			log.Printf("[MYDEBUG] slotNumber %d, %s has some change to: %s", slotNumber, networkIDKey, networkID)
			updateFixedIPInterface.NetworkID = &networkID
		}
		log.Printf("[MYDEBUG] updateFixedIPInterface in slot %d: %#v", slotNumber, updateFixedIPInterface)

		addressInfo := []appliances.UpdateFixedIPAddressInfo{}

		fixedIPsKey := fmt.Sprintf("interface_%d_fixed_ips", slotNumber)
		if d.HasChange(fixedIPsKey) {
			updated = true
			rawFixedIPs := d.Get(fixedIPsKey).([]interface{})
			for _, rawFixedIP := range rawFixedIPs {
				tmpFixedIP := rawFixedIP.(map[string]interface{})
				fixedIP := appliances.UpdateFixedIPAddressInfo{}
				fixedIP.IPAddress = tmpFixedIP["ip_address"].(string)
				addressInfo = append(addressInfo, fixedIP)
			}
			log.Printf("[MYDEBUG] slotNumber %d, %s has some change to: %#v", slotNumber, fixedIPsKey, addressInfo)
		}
		updateFixedIPInterface.FixedIPs = &addressInfo
		log.Printf("[MYDEBUG] updateFixedIPInterface in slot %d: %#v", slotNumber, updateFixedIPInterface)

		switch slotNumber {
		case 1:
			updateFixedIPOpts.Interfaces.Interface1 = &updateFixedIPInterface
			break
		case 2:
			updateFixedIPOpts.Interfaces.Interface2 = &updateFixedIPInterface
			break
		case 3:
			updateFixedIPOpts.Interfaces.Interface3 = &updateFixedIPInterface
			break
		case 4:
			updateFixedIPOpts.Interfaces.Interface4 = &updateFixedIPInterface
			break
		case 5:
			updateFixedIPOpts.Interfaces.Interface5 = &updateFixedIPInterface
			break
		case 6:
			updateFixedIPOpts.Interfaces.Interface6 = &updateFixedIPInterface
			break
		case 7:
			updateFixedIPOpts.Interfaces.Interface7 = &updateFixedIPInterface
			break
		case 8:
			updateFixedIPOpts.Interfaces.Interface8 = &updateFixedIPInterface
			break
		}
	}

	if updated {
		log.Printf("[DEBUG] Updating VNA Interface %s with options: %+v", d.Id(), updateFixedIPOpts)
		_, err := appliances.Update(client, d.Id(), updateFixedIPOpts).Extract()
		if err != nil {
			return fmt.Errorf(
				"Error updating for virtual network appliance (%s) with option %#v: %s,",
				d.Id(), updateFixedIPOpts, err)
		}

		stateConf := &resource.StateChangeConf{
			Pending:      []string{"PROCESSING"},
			Target:       []string{"COMPLETE"},
			Refresh:      waitForVirtualNetworkApplianceComplete(client, d.Id()),
			Timeout:      d.Timeout(schema.TimeoutDelete),
			Delay:        5 * time.Second,
			PollInterval: updatePollInterval,
			MinTimeout:   3 * time.Second,
		}

		_, err = stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf(
				"Error waiting for virtual network appliance (%s) to become COMPLETE(after interface update): %s",
				d.Id(), err)
		}
	}

	return nil
}
func updateMetadata(d *schema.ResourceData, meta interface{}, client *eclcloud.ServiceClient) error {
	var updateMeta, updateMetaInInterface bool

	var updateMetadataOpts appliances.UpdateMetadataOpts
	if d.HasChange("name") {
		updateMeta = true
		name := d.Get("name").(string)
		updateMetadataOpts.Name = &name
	}

	if d.HasChange("description") {
		updateMeta = true
		description := d.Get("description").(string)
		updateMetadataOpts.Description = &description
	}

	if d.HasChange("tags") {
		updateMeta = true
		tags := resourceTags(d)
		updateMetadataOpts.Tags = &tags
	}

	updateMetadataInterfaces := appliances.UpdateMetadataInterfaces{}
	for _, slotNumber := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		updateMetadataInterface := appliances.UpdateMetadataInterface{}

		nameKey := fmt.Sprintf("interface_%d_info.0.name", slotNumber)
		if d.HasChange(nameKey) {
			updateMeta = true
			updateMetaInInterface = true
			name := d.Get(nameKey).(string)
			updateMetadataInterface.Name = &name
		}

		descriptionKey := fmt.Sprintf("interface_%d_info.0.description", slotNumber)
		if d.HasChange(descriptionKey) {
			updateMeta = true
			updateMetaInInterface = true
			description := d.Get(descriptionKey).(string)
			updateMetadataInterface.Description = &description
		}

		tagsKey := fmt.Sprintf("interface_%d_info.0.tags", slotNumber)
		if d.HasChange(tagsKey) {
			updateMeta = true
			updateMetaInInterface = true

			schemaTags := d.Get(tagsKey)
			newTags := map[string]string{}
			for k, v := range schemaTags.(map[string]interface{}) {
				newTags[k] = v.(string)
			}
			updateMetadataInterface.Tags = &newTags
		}
		switch slotNumber {
		case 1:
			updateMetadataInterfaces.Interface1 = &updateMetadataInterface
			break
		case 2:
			updateMetadataInterfaces.Interface2 = &updateMetadataInterface
			break
		case 3:
			updateMetadataInterfaces.Interface3 = &updateMetadataInterface
			break
		case 4:
			updateMetadataInterfaces.Interface4 = &updateMetadataInterface
			break
		case 5:
			updateMetadataInterfaces.Interface5 = &updateMetadataInterface
			break
		case 6:
			updateMetadataInterfaces.Interface6 = &updateMetadataInterface
			break
		case 7:
			updateMetadataInterfaces.Interface7 = &updateMetadataInterface
			break
		case 8:
			updateMetadataInterfaces.Interface8 = &updateMetadataInterface
			break
		}
	}

	if updateMetaInInterface {
		updateMetadataOpts.Interfaces = updateMetadataInterfaces
	}

	if updateMeta {
		log.Printf("[DEBUG] Updating VNA Metadata %s with options: %+v", d.Id(), updateMetadataOpts)
		_, err := appliances.Update(client, d.Id(), updateMetadataOpts).Extract()
		if err != nil {
			return fmt.Errorf(
				"Error updating for virtual network appliance (%s) with option %#v: %s,",
				d.Id(), updateMetadataOpts, err)
		}

		stateConf := &resource.StateChangeConf{
			Pending:      []string{"PROCESSING"},
			Target:       []string{"COMPLETE"},
			Refresh:      waitForVirtualNetworkApplianceComplete(client, d.Id()),
			Timeout:      d.Timeout(schema.TimeoutDelete),
			Delay:        5 * time.Second,
			PollInterval: updatePollInterval,
			MinTimeout:   3 * time.Second,
		}

		_, err = stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf(
				"Error waiting for virtual network appliance (%s) to become COMPLETE(after metadata update): %s",
				d.Id(), err)
		}
	}

	return nil
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
