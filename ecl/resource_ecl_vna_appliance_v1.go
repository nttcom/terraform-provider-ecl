package ecl

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

const pollingSec = 30

const createPollInterval = pollingSec * time.Second
const updatePollInterval = pollingSec * time.Second
const deletePollInterval = pollingSec * time.Second

func noAllowedAddressPairsSchema(slotNumber int) *schema.Schema {
	conflictsWith := fmt.Sprintf("interface_%d_allowed_address_pairs", slotNumber)
	return &schema.Schema{
		Type:          schema.TypeBool,
		Optional:      true,
		ConflictsWith: []string{conflictsWith},
	}
}

func allowedAddessPairsSchema(slotNumber int) *schema.Schema {
	conflictsWith := fmt.Sprintf("interface_%d_no_allowed_address_pairs", slotNumber)
	return &schema.Schema{
		Type:          schema.TypeList,
		Optional:      true,
		Computed:      true,
		ConflictsWith: []string{conflictsWith},
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
					Required: true,
					ValidateFunc: validation.StringInSlice(
						[]string{"", "vrrp"}, false,
					),
				},

				"vrid": &schema.Schema{
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: ValidateVRID(),
				},
			},
		},
	}
}

func noFixedIPsSchema(slotNumber int) *schema.Schema {
	conflictsWith := fmt.Sprintf("interface_%d_fixed_ips", slotNumber)
	return &schema.Schema{
		Type:          schema.TypeBool,
		Optional:      true,
		ConflictsWith: []string{conflictsWith},
	}
}

func fixedIPsSchema(slotNumber int) *schema.Schema {
	conflictsWith := fmt.Sprintf("interface_%d_no_fixed_ips", slotNumber)
	return &schema.Schema{
		Type:          schema.TypeList,
		Optional:      true,
		Computed:      true,
		ConflictsWith: []string{conflictsWith},
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

			"interface_1_info":                     interfaceInfoSchema(),
			"interface_1_fixed_ips":                fixedIPsSchema(1),
			"interface_1_no_fixed_ips":             noFixedIPsSchema(1),
			"interface_1_allowed_address_pairs":    allowedAddessPairsSchema(1),
			"interface_1_no_allowed_address_pairs": noAllowedAddressPairsSchema(1),

			"interface_2_info":                     interfaceInfoSchema(),
			"interface_2_fixed_ips":                fixedIPsSchema(2),
			"interface_2_no_fixed_ips":             noFixedIPsSchema(2),
			"interface_2_allowed_address_pairs":    allowedAddessPairsSchema(2),
			"interface_2_no_allowed_address_pairs": noAllowedAddressPairsSchema(2),

			"interface_3_info":                     interfaceInfoSchema(),
			"interface_3_fixed_ips":                fixedIPsSchema(3),
			"interface_3_no_fixed_ips":             noFixedIPsSchema(3),
			"interface_3_allowed_address_pairs":    allowedAddessPairsSchema(3),
			"interface_3_no_allowed_address_pairs": noAllowedAddressPairsSchema(3),

			"interface_4_info":                     interfaceInfoSchema(),
			"interface_4_fixed_ips":                fixedIPsSchema(4),
			"interface_4_no_fixed_ips":             noFixedIPsSchema(4),
			"interface_4_allowed_address_pairs":    allowedAddessPairsSchema(4),
			"interface_4_no_allowed_address_pairs": noAllowedAddressPairsSchema(4),

			"interface_5_info":                     interfaceInfoSchema(),
			"interface_5_fixed_ips":                fixedIPsSchema(5),
			"interface_5_no_fixed_ips":             noFixedIPsSchema(5),
			"interface_5_allowed_address_pairs":    allowedAddessPairsSchema(5),
			"interface_5_no_allowed_address_pairs": noAllowedAddressPairsSchema(5),

			"interface_6_info":                     interfaceInfoSchema(),
			"interface_6_fixed_ips":                fixedIPsSchema(6),
			"interface_6_no_fixed_ips":             noFixedIPsSchema(6),
			"interface_6_allowed_address_pairs":    allowedAddessPairsSchema(6),
			"interface_6_no_allowed_address_pairs": noAllowedAddressPairsSchema(6),

			"interface_7_info":                     interfaceInfoSchema(),
			"interface_7_fixed_ips":                fixedIPsSchema(7),
			"interface_7_no_fixed_ips":             noFixedIPsSchema(7),
			"interface_7_allowed_address_pairs":    allowedAddessPairsSchema(7),
			"interface_7_no_allowed_address_pairs": noAllowedAddressPairsSchema(7),

			"interface_8_info":                     interfaceInfoSchema(),
			"interface_8_fixed_ips":                fixedIPsSchema(8),
			"interface_8_no_fixed_ips":             noFixedIPsSchema(8),
			"interface_8_allowed_address_pairs":    allowedAddessPairsSchema(8),
			"interface_8_no_allowed_address_pairs": noAllowedAddressPairsSchema(8),
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

func resourceVNAApplianceV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	log.Printf("[DEBUG] Start updating metadata of virtual network appliance ...")
	err = updateMetadata(d, meta, vnaClient)
	if err != nil {
		return fmt.Errorf("Error in updating virtual network appliance metadata: %s", err)
	}

	log.Printf("[DEBUG] Start updating attached network and fixed_ips of virtual network appliance ...")
	err = updateFixedIPs(d, meta, vnaClient)
	if err != nil {
		return fmt.Errorf("Error in updating virtual network appliance interface/fixed_ips or interface/network_id: %s", err)
	}

	log.Printf("[DEBUG] Start updating allowed address pairs of virtual network appliance ...")
	err = updateAllowedAddressPairs(d, meta, vnaClient)
	if err != nil {
		return fmt.Errorf("Error in updating virtual network appliance interface/allowed_address_pairs: %s", err)
	}

	return resourceVNAApplianceV1Read(d, meta)
}

func updateAllowedAddressPairs(d *schema.ResourceData, meta interface{}, client *eclcloud.ServiceClient) error {
	var isAtLeastOneInterfaceUpdated bool

	var updateAllowedAddressPairOpts appliances.UpdateAllowedAddressPairOpts

	allInterfaces := appliances.UpdateAllowedAddressPairInterfaces{}

	for _, slotNumber := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		isInterfaceUpdated := false
		updateAllowedAddressPairInterface := appliances.UpdateAllowedAddressPairInterface{}

		allowedAddressPairsKey := fmt.Sprintf("interface_%d_allowed_address_pairs", slotNumber)
		noAllowedAddressPairsKey := fmt.Sprintf("interface_%d_no_allowed_address_pairs", slotNumber)

		if d.HasChange(allowedAddressPairsKey) {
			isInterfaceUpdated = true

			var addressInfo []appliances.UpdateAllowedAddressPairAddressInfo

			addressInfo = []appliances.UpdateAllowedAddressPairAddressInfo{}
			rawAllowedAddressPairs := d.Get(allowedAddressPairsKey).([]interface{})
			for _, rawAllowedAddressPair := range rawAllowedAddressPairs {
				tmpAllowedAddressPair := rawAllowedAddressPair.(map[string]interface{})
				allowedAddressPair := appliances.UpdateAllowedAddressPairAddressInfo{}

				allowedAddressPair.IPAddress = tmpAllowedAddressPair["ip_address"].(string)

				macAddress := tmpAllowedAddressPair["mac_address"].(string)
				allowedAddressPair.MACAddress = &macAddress

				aapType := tmpAllowedAddressPair["type"].(string)
				allowedAddressPair.Type = &aapType

				vridString := tmpAllowedAddressPair["vrid"].(string)
				switch vridString {
				case "null":
					var v interface{}
					v = interface{}(nil)
					allowedAddressPair.VRID = &v
					break
				default:
					var v interface{}
					v, _ = strconv.Atoi(vridString)
					allowedAddressPair.VRID = &v
				}

				addressInfo = append(addressInfo, allowedAddressPair)
			}
			log.Printf("[MYDEBUG] slotNumber %d, %s has some change to: %#v", slotNumber, allowedAddressPairsKey, addressInfo)
			updateAllowedAddressPairInterface.AllowedAddressPairs = &addressInfo

			if _, ok := d.GetOk(noAllowedAddressPairsKey); ok {
				isInterfaceUpdated = true
				addressInfo := []appliances.UpdateAllowedAddressPairAddressInfo{}
				updateAllowedAddressPairInterface.AllowedAddressPairs = &addressInfo
			}
		}

		log.Printf("[MYDEBUG] updateAllowedAddressPairInterface in slot %d: %#v", slotNumber, updateAllowedAddressPairInterface)

		if isInterfaceUpdated {
			isAtLeastOneInterfaceUpdated = true
			// thisInterface := interface {}
			switch slotNumber {
			case 1:
				allInterfaces.Interface1 = updateAllowedAddressPairInterface
				break
			case 2:
				allInterfaces.Interface2 = updateAllowedAddressPairInterface
				break
			case 3:
				allInterfaces.Interface3 = updateAllowedAddressPairInterface
				break
			case 4:
				allInterfaces.Interface4 = updateAllowedAddressPairInterface
				break
			case 5:
				allInterfaces.Interface5 = updateAllowedAddressPairInterface
				break
			case 6:
				allInterfaces.Interface6 = updateAllowedAddressPairInterface
				break
			case 7:
				allInterfaces.Interface7 = updateAllowedAddressPairInterface
				break
			case 8:
				allInterfaces.Interface8 = updateAllowedAddressPairInterface
				break
			}
		} else {
			switch slotNumber {
			case 1:
				allInterfaces.Interface1 = interface{}(nil)
				break
			case 2:
				allInterfaces.Interface2 = interface{}(nil)
				break
			case 3:
				allInterfaces.Interface3 = interface{}(nil)
				break
			case 4:
				allInterfaces.Interface4 = interface{}(nil)
				break
			case 5:
				allInterfaces.Interface5 = interface{}(nil)
				break
			case 6:
				allInterfaces.Interface6 = interface{}(nil)
				break
			case 7:
				allInterfaces.Interface7 = interface{}(nil)
				break
			case 8:
				allInterfaces.Interface8 = interface{}(nil)
				break
			}
		}
	}

	if isAtLeastOneInterfaceUpdated {
		updateAllowedAddressPairOpts.Interfaces = allInterfaces
	} else {
		updateAllowedAddressPairOpts.Interfaces = interface{}(nil)
	}

	if isAtLeastOneInterfaceUpdated {
		log.Printf("[DEBUG] Updating VNA Allowed Address Pair %s with options: %+v", d.Id(), updateAllowedAddressPairOpts)
		_, err := appliances.Update(client, d.Id(), updateAllowedAddressPairOpts).Extract()
		if err != nil {
			return fmt.Errorf(
				"Error updating for virtual network appliance (%s) with option %#v: %s,",
				d.Id(), updateAllowedAddressPairOpts, err)
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
				"Error waiting for virtual network appliance (%s) to become COMPLETE(after allowed address pairs update): %s",
				d.Id(), err)
		}
	}

	return nil
}

// updateFixedIP handles following updates
// - interface_N.network_id
// - interface_N.fixedips list
//  Above updates are correspond to "interface update" in VNA API
func updateFixedIPs(d *schema.ResourceData, meta interface{}, client *eclcloud.ServiceClient) error {
	var isAtLeastOneInterfaceUpdated bool

	var updateFixedIPOpts appliances.UpdateFixedIPOpts
	// var updateFixedIPOpts = appliances.UpdateFixedIPOpts{}

	allInterfaces := appliances.UpdateFixedIPInterfaces{}

	for _, slotNumber := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		isInterfaceUpdated := false
		updateFixedIPInterface := appliances.UpdateFixedIPInterface{}

		// var networkID string
		networkIDKey := fmt.Sprintf("interface_%d_info.0.network_id", slotNumber)
		if d.HasChange(networkIDKey) {
			isInterfaceUpdated = true
			networkID := d.Get(networkIDKey).(string)
			log.Printf("[MYDEBUG] slotNumber %d, %s has some change to: %s", slotNumber, networkIDKey, networkID)
			updateFixedIPInterface.NetworkID = &networkID
		}
		log.Printf("[MYDEBUG] updateFixedIPInterface in slot %d: %#v", slotNumber, updateFixedIPInterface)

		fixedIPsKey := fmt.Sprintf("interface_%d_fixed_ips", slotNumber)
		noFixedIPsKey := fmt.Sprintf("interface_%d_no_fixed_ips", slotNumber)

		if d.HasChange(fixedIPsKey) {
			isInterfaceUpdated = true

			var addressInfo []appliances.UpdateFixedIPAddressInfo
			addressInfo = []appliances.UpdateFixedIPAddressInfo{}
			rawFixedIPs := d.Get(fixedIPsKey).([]interface{})

			for _, rawFixedIP := range rawFixedIPs {
				tmpFixedIP := rawFixedIP.(map[string]interface{})
				fixedIP := appliances.UpdateFixedIPAddressInfo{}
				fixedIP.IPAddress = tmpFixedIP["ip_address"].(string)
				addressInfo = append(addressInfo, fixedIP)
			}

			log.Printf("[MYDEBUG] slotNumber %d, %s has some change to: %#v", slotNumber, fixedIPsKey, addressInfo)
			updateFixedIPInterface.FixedIPs = &addressInfo
		}

		if _, ok := d.GetOk(noFixedIPsKey); ok {
			isInterfaceUpdated = true
			addressInfo := []appliances.UpdateFixedIPAddressInfo{}
			updateFixedIPInterface.FixedIPs = &addressInfo
		}

		log.Printf("[MYDEBUG] updateFixedIPInterface in slot %d: %#v", slotNumber, updateFixedIPInterface)

		if isInterfaceUpdated {
			isAtLeastOneInterfaceUpdated = true
			// thisInterface := interface {}
			switch slotNumber {
			case 1:
				allInterfaces.Interface1 = updateFixedIPInterface
				break
			case 2:
				allInterfaces.Interface2 = updateFixedIPInterface
				break
			case 3:
				allInterfaces.Interface3 = updateFixedIPInterface
				break
			case 4:
				allInterfaces.Interface4 = updateFixedIPInterface
				break
			case 5:
				allInterfaces.Interface5 = updateFixedIPInterface
				break
			case 6:
				allInterfaces.Interface6 = updateFixedIPInterface
				break
			case 7:
				allInterfaces.Interface7 = updateFixedIPInterface
				break
			case 8:
				allInterfaces.Interface8 = updateFixedIPInterface
				break
			}
		} else {
			switch slotNumber {
			case 1:
				allInterfaces.Interface1 = interface{}(nil)
				break
			case 2:
				allInterfaces.Interface2 = interface{}(nil)
				break
			case 3:
				allInterfaces.Interface3 = interface{}(nil)
				break
			case 4:
				allInterfaces.Interface4 = interface{}(nil)
				break
			case 5:
				allInterfaces.Interface5 = interface{}(nil)
				break
			case 6:
				allInterfaces.Interface6 = interface{}(nil)
				break
			case 7:
				allInterfaces.Interface7 = interface{}(nil)
				break
			case 8:
				allInterfaces.Interface8 = interface{}(nil)
				break
			}
		}
	}

	if isAtLeastOneInterfaceUpdated {
		updateFixedIPOpts.Interfaces = allInterfaces
	} else {
		updateFixedIPOpts.Interfaces = interface{}(nil)
	}

	if isAtLeastOneInterfaceUpdated {
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
