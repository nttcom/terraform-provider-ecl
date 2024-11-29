package ecl

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/vna/v1/appliances"
)

const maxNumberOfInterfaces = 8

const vnaPollingSec = 30
const vnaCreatePollInterval = vnaPollingSec * time.Second
const vnaUpdatePollInterval = vnaPollingSec * time.Second
const vnaDeletePollInterval = vnaPollingSec * time.Second

func getApplianceTags(d *schema.ResourceData) map[string]string {
	rawTags := d.Get("tags").(map[string]interface{})
	tags := map[string]string{}
	for key, value := range rawTags {
		if v, ok := value.(string); ok {
			tags[key] = v
		}
	}
	return tags
}

func getInterfaceAllowedAddressPairsAsState(allowedAddressPairs []appliances.AllowedAddressPairInResponse) []interface{} {
	result := make([]interface{}, len(allowedAddressPairs))
	for i, aap := range allowedAddressPairs {
		thisAAP := map[string]string{}
		thisAAP["ip_address"] = aap.IPAddress
		thisAAP["mac_address"] = aap.MACAddress
		thisAAP["type"] = aap.Type

		var vrid string

		if aap.VRID == interface{}(nil) {
			log.Printf("[DEBUG] VRID has converted into null")
			vrid = "null"
		} else {
			v, ok := aap.VRID.(float64)
			if !ok {
				log.Printf("[DEBUG] VRID float assertion failed v : ok =  %#v : %#v", v, ok)
			}
			iv := int(v)
			sv := strconv.Itoa(iv)
			vrid = sv
		}
		thisAAP["vrid"] = vrid

		result[i] = thisAAP
	}
	log.Printf("[DEBUG] Result Allowed Address Pairs: %#v", result)
	return result
}

func getInterfaceFixedIPsAsState(fixedIPs []appliances.FixedIPInResponse) []interface{} {
	result := make([]interface{}, len(fixedIPs))

	for i, fixedIP := range fixedIPs {
		thisFixedIP := map[string]string{}
		thisFixedIP["ip_address"] = fixedIP.IPAddress
		thisFixedIP["subnet_id"] = fixedIP.SubnetID

		result[i] = thisFixedIP
	}
	log.Printf("[DEBUG] Result FixedIPs: %#v", result)
	return result
}

func getInterfaceInfoAsState(singleInterface appliances.InterfaceInResponse) []interface{} {
	var result []interface{}
	result = []interface{}{}

	var meta map[string]interface{}
	meta = map[string]interface{}{}

	meta["name"] = singleInterface.Name
	meta["description"] = singleInterface.Description
	meta["network_id"] = singleInterface.NetworkID
	meta["updatable"] = singleInterface.Updatable
	meta["tags"] = singleInterface.Tags

	result = append(result, meta)
	log.Printf("[DEBUG] Result Interface data: %#v", result)
	return result
}

func getTagsAsOpts(rawTags map[string]interface{}) map[string]string {
	var tags map[string]string
	tags = map[string]string{}
	for k, v := range rawTags {
		tags[k] = v.(string)
	}
	return tags
}

func getInitialConfigCreateOpts(d *schema.ResourceData) *appliances.CreateOptsInitialConfig {
	if _, ok := d.GetOk("initial_config"); ok {
		return &appliances.CreateOptsInitialConfig{
			Format: d.Get("initial_config.format").(string),
			Data:   d.Get("initial_config.data").(string),
		}
	}
	return nil
}

func getInterfaceCreateOpts(d *schema.ResourceData) *appliances.CreateOptsInterfaces {
	var interfaces appliances.CreateOptsInterfaces
	isInterfacesCreated := false

	for slotNumber := 1; slotNumber <= maxNumberOfInterfaces; slotNumber++ {
		isInterfaceCreated := false
		var interfaceEach appliances.CreateOptsInterface
		rawMeta := d.Get(fmt.Sprintf("interface_%d_info", slotNumber)).([]interface{})
		rawFips := d.Get(fmt.Sprintf("interface_%d_fixed_ips", slotNumber)).([]interface{})
		noFixedIPsKey := fmt.Sprintf("interface_%d_no_fixed_ips", slotNumber)

		for index, rm := range rawMeta {
			thisRawMeta := rm.(map[string]interface{})
			if index == 0 {
				isInterfaceCreated = true
				interfaceEach.Name = thisRawMeta["name"].(string)
				interfaceEach.Description = thisRawMeta["description"].(string)
				interfaceEach.NetworkID = thisRawMeta["network_id"].(string)
				tags := getTagsAsOpts(thisRawMeta["tags"].(map[string]interface{}))
				interfaceEach.Tags = tags
			}
		}

		interfaceEach.FixedIPs = nil
		for index, rawFip := range rawFips {
			if index == 0 {
				fip := rawFip.(map[string]interface{})

				ipAddress := fip["ip_address"].(string)
				var fixedIP appliances.CreateOptsFixedIP
				fixedIP.IPAddress = ipAddress
				resultFixedIPs := []appliances.CreateOptsFixedIP{}
				resultFixedIPs = append(resultFixedIPs, fixedIP)

				interfaceEach.FixedIPs = &resultFixedIPs
			}
		}

		if _, ok := d.GetOk(noFixedIPsKey); ok {
			resultFixedIPs := []appliances.CreateOptsFixedIP{}
			interfaceEach.FixedIPs = &resultFixedIPs
		}

		interfaceOrNil := &interfaceEach
		if !isInterfaceCreated {
			interfaceOrNil = nil
		} else {
			isInterfacesCreated = true
		}
		switch slotNumber {
		case 1:
			interfaces.Interface1 = interfaceOrNil
			break
		case 2:
			interfaces.Interface2 = interfaceOrNil
			break
		case 3:
			interfaces.Interface3 = interfaceOrNil
			break
		case 4:
			interfaces.Interface4 = interfaceOrNil
			break
		case 5:
			interfaces.Interface5 = interfaceOrNil
			break
		case 6:
			interfaces.Interface6 = interfaceOrNil
			break
		case 7:
			interfaces.Interface7 = interfaceOrNil
			break
		case 8:
			interfaces.Interface8 = interfaceOrNil
			break
		}
	}
	if !isInterfacesCreated {
		return nil
	}
	return &interfaces
}

func getInterfaceBySlotNumber(vna *appliances.Appliance, slotNumber int) appliances.InterfaceInResponse {
	var result appliances.InterfaceInResponse
	switch slotNumber {
	case 1:
		result = vna.Interfaces.Interface1
		break
	case 2:
		result = vna.Interfaces.Interface2
		break
	case 3:
		result = vna.Interfaces.Interface3
		break
	case 4:
		result = vna.Interfaces.Interface4
		break
	case 5:
		result = vna.Interfaces.Interface5
		break
	case 6:
		result = vna.Interfaces.Interface6
		break
	case 7:
		result = vna.Interfaces.Interface7
		break
	case 8:
		result = vna.Interfaces.Interface8
		break
	default:
		break
	}
	log.Printf("[DEBUG] Retrieved Interface by slotNumber %d : %#v", slotNumber, result)
	return result
}

func getFixedIPsBySlotNumber(vna *appliances.Appliance, slotNumber int) []appliances.FixedIPInResponse {
	var result []appliances.FixedIPInResponse
	switch slotNumber {
	case 1:
		result = vna.Interfaces.Interface1.FixedIPs
		break
	case 2:
		result = vna.Interfaces.Interface2.FixedIPs
		break
	case 3:
		result = vna.Interfaces.Interface3.FixedIPs
		break
	case 4:
		result = vna.Interfaces.Interface4.FixedIPs
		break
	case 5:
		result = vna.Interfaces.Interface5.FixedIPs
		break
	case 6:
		result = vna.Interfaces.Interface6.FixedIPs
		break
	case 7:
		result = vna.Interfaces.Interface7.FixedIPs
		break
	case 8:
		result = vna.Interfaces.Interface8.FixedIPs
		break
	default:
		break
	}
	log.Printf("[DEBUG] Retrieved FixedIP by slotNumber %d : %#v", slotNumber, result)
	return result
}

func getAllowedAddressPairsBySlotNumber(vna *appliances.Appliance, slotNumber int) []appliances.AllowedAddressPairInResponse {
	var result []appliances.AllowedAddressPairInResponse
	switch slotNumber {
	case 1:
		result = vna.Interfaces.Interface1.AllowedAddressPairs
		break
	case 2:
		result = vna.Interfaces.Interface2.AllowedAddressPairs
		break
	case 3:
		result = vna.Interfaces.Interface3.AllowedAddressPairs
		break
	case 4:
		result = vna.Interfaces.Interface4.AllowedAddressPairs
		break
	case 5:
		result = vna.Interfaces.Interface5.AllowedAddressPairs
		break
	case 6:
		result = vna.Interfaces.Interface6.AllowedAddressPairs
		break
	case 7:
		result = vna.Interfaces.Interface7.AllowedAddressPairs
		break
	case 8:
		result = vna.Interfaces.Interface8.AllowedAddressPairs
		break
	default:
		break
	}
	log.Printf("[DEBUG] Retrieved Allowed Address Pair by slotNumber %d : %#v", slotNumber, result)
	return result
}

func updateAllowedAddressPairs(d *schema.ResourceData, meta interface{}, client *eclcloud.ServiceClient) error {
	var isAtLeastOneInterfaceUpdated bool

	var updateAllowedAddressPairOpts appliances.UpdateAllowedAddressPairOpts

	allInterfaces := appliances.UpdateAllowedAddressPairInterfaces{}

	for slotNumber := 1; slotNumber <= maxNumberOfInterfaces; slotNumber++ {
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

			updateAllowedAddressPairInterface.AllowedAddressPairs = &addressInfo

			if _, ok := d.GetOk(noAllowedAddressPairsKey); ok {
				isInterfaceUpdated = true
				addressInfo := []appliances.UpdateAllowedAddressPairAddressInfo{}
				updateAllowedAddressPairInterface.AllowedAddressPairs = &addressInfo
			}
		}

		if isInterfaceUpdated {
			isAtLeastOneInterfaceUpdated = true
			switch slotNumber {
			case 1:
				allInterfaces.Interface1 = &updateAllowedAddressPairInterface
				break
			case 2:
				allInterfaces.Interface2 = &updateAllowedAddressPairInterface
				break
			case 3:
				allInterfaces.Interface3 = &updateAllowedAddressPairInterface
				break
			case 4:
				allInterfaces.Interface4 = &updateAllowedAddressPairInterface
				break
			case 5:
				allInterfaces.Interface5 = &updateAllowedAddressPairInterface
				break
			case 6:
				allInterfaces.Interface6 = &updateAllowedAddressPairInterface
				break
			case 7:
				allInterfaces.Interface7 = &updateAllowedAddressPairInterface
				break
			case 8:
				allInterfaces.Interface8 = &updateAllowedAddressPairInterface
				break
			}
		}
	}

	if isAtLeastOneInterfaceUpdated {
		updateAllowedAddressPairOpts.Interfaces = &allInterfaces
	} else {
		updateAllowedAddressPairOpts.Interfaces = nil
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
			PollInterval: vnaUpdatePollInterval,
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

func updateFixedIPs(d *schema.ResourceData, meta interface{}, client *eclcloud.ServiceClient) error {
	var isAtLeastOneInterfaceUpdated bool

	var updateFixedIPOpts appliances.UpdateFixedIPOpts

	allInterfaces := appliances.UpdateFixedIPInterfaces{}

	for slotNumber := 1; slotNumber <= maxNumberOfInterfaces; slotNumber++ {
		isInterfaceUpdated := false
		updateFixedIPInterface := appliances.UpdateFixedIPInterface{}

		networkIDKey := fmt.Sprintf("interface_%d_info.0.network_id", slotNumber)
		if d.HasChange(networkIDKey) {
			isInterfaceUpdated = true
			networkID := d.Get(networkIDKey).(string)
			updateFixedIPInterface.NetworkID = &networkID
		}

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

			updateFixedIPInterface.FixedIPs = &addressInfo
		}

		if _, ok := d.GetOk(noFixedIPsKey); ok {
			isInterfaceUpdated = true
			addressInfo := []appliances.UpdateFixedIPAddressInfo{}
			updateFixedIPInterface.FixedIPs = &addressInfo
		}

		if isInterfaceUpdated {
			isAtLeastOneInterfaceUpdated = true
			switch slotNumber {
			case 1:
				allInterfaces.Interface1 = &updateFixedIPInterface
				break
			case 2:
				allInterfaces.Interface2 = &updateFixedIPInterface
				break
			case 3:
				allInterfaces.Interface3 = &updateFixedIPInterface
				break
			case 4:
				allInterfaces.Interface4 = &updateFixedIPInterface
				break
			case 5:
				allInterfaces.Interface5 = &updateFixedIPInterface
				break
			case 6:
				allInterfaces.Interface6 = &updateFixedIPInterface
				break
			case 7:
				allInterfaces.Interface7 = &updateFixedIPInterface
				break
			case 8:
				allInterfaces.Interface8 = &updateFixedIPInterface
				break
			}
		}
	}

	if isAtLeastOneInterfaceUpdated {
		updateFixedIPOpts.Interfaces = &allInterfaces
	} else {
		updateFixedIPOpts.Interfaces = nil
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
			PollInterval: vnaUpdatePollInterval,
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
	var isMetaUpdated bool
	var isAtLeastOneInterfaceUpdated bool

	var updateMetadataOpts appliances.UpdateMetadataOpts

	allInterfaces := appliances.UpdateMetadataInterfaces{}

	if d.HasChange("name") {
		isMetaUpdated = true
		name := d.Get("name").(string)
		updateMetadataOpts.Name = &name
	}

	if d.HasChange("description") {
		isMetaUpdated = true
		description := d.Get("description").(string)
		updateMetadataOpts.Description = &description
	}

	if d.HasChange("tags") {
		isMetaUpdated = true
		tags := resourceTags(d)
		updateMetadataOpts.Tags = &tags
	}

	for slotNumber := 1; slotNumber <= maxNumberOfInterfaces; slotNumber++ {
		isInterfaceMetaUpdated := false
		updateMetadataInterface := appliances.UpdateMetadataInterface{}

		nameKey := fmt.Sprintf("interface_%d_info.0.name", slotNumber)
		if d.HasChange(nameKey) {
			isInterfaceMetaUpdated = true
			name := d.Get(nameKey).(string)
			updateMetadataInterface.Name = &name
		}

		descriptionKey := fmt.Sprintf("interface_%d_info.0.description", slotNumber)
		if d.HasChange(descriptionKey) {
			isInterfaceMetaUpdated = true
			description := d.Get(descriptionKey).(string)
			updateMetadataInterface.Description = &description
		}

		tagsKey := fmt.Sprintf("interface_%d_info.0.tags", slotNumber)

		if d.HasChange(tagsKey) {
			isInterfaceMetaUpdated = true

			schemaTags := d.Get(tagsKey)
			newTags := map[string]string{}
			for k, v := range schemaTags.(map[string]interface{}) {
				newTags[k] = v.(string)
			}
			updateMetadataInterface.Tags = &newTags
		}

		if isInterfaceMetaUpdated {
			isAtLeastOneInterfaceUpdated = true

			switch slotNumber {
			case 1:
				allInterfaces.Interface1 = &updateMetadataInterface
				break
			case 2:
				allInterfaces.Interface2 = &updateMetadataInterface
				break
			case 3:
				allInterfaces.Interface3 = &updateMetadataInterface
				break
			case 4:
				allInterfaces.Interface4 = &updateMetadataInterface
				break
			case 5:
				allInterfaces.Interface5 = &updateMetadataInterface
				break
			case 6:
				allInterfaces.Interface6 = &updateMetadataInterface
				break
			case 7:
				allInterfaces.Interface7 = &updateMetadataInterface
				break
			case 8:
				allInterfaces.Interface8 = &updateMetadataInterface
				break
			}
		}
	}

	if isAtLeastOneInterfaceUpdated {
		updateMetadataOpts.Interfaces = &allInterfaces
	} else {
		updateMetadataOpts.Interfaces = nil
	}

	if isMetaUpdated || isAtLeastOneInterfaceUpdated {
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
			PollInterval: vnaUpdatePollInterval,
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
