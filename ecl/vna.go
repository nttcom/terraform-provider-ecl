package ecl

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

const maxNumberOfInterfaces = 8

// func fixedIPHash(v interface{}) int {
// 	var buf bytes.Buffer
// 	m := v.(map[string]interface{})
// 	if m["ip_address"] != nil && m["subnet_id"] != nil {
// 		buf.WriteString(
// 			fmt.Sprintf(
// 				"%s-%s-",
// 				m["ip_address"].(string),
// 				m["subnet_id"].(string),
// 			))
// 	}
// 	return hashcode.String(buf.String())
// }

// func allowedAddressPairHash(v interface{}) int {
// 	var buf bytes.Buffer
// 	m := v.(map[string]interface{})
// 	if m["ip_address"] != nil &&
// 		m["mac_address"] != nil &&
// 		m["type"] != nil &&
// 		m["vrid"] != nil {
// 		buf.WriteString(
// 			fmt.Sprintf(
// 				"%s-%s-%s-%s-",
// 				m["ip_address"].(string),
// 				m["subnet_id"].(string),
// 				m["type"].(string),
// 				m["vrid"].(string),
// 			))
// 	}
// 	return hashcode.String(buf.String())
// }

// func interfaceHash(v interface{}) int {
// 	var buf bytes.Buffer
// 	m := v.(map[string]interface{})
// 	if m["slot_number"] != nil {
// 		buf.WriteString(fmt.Sprintf("%d-", m["slot_number"].(int)))
// 	}
// 	return hashcode.String(buf.String())
// }

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
		thisAAP["vrid"] = aap.VRID

		result[i] = thisAAP
	}
	log.Printf("[MYDEBUG] Result Allowed Address Pairs: %#v", result)
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
	log.Printf("[MYDEBUG] Result FixedIPs: %#v", result)
	return result
}

func getInterfaceMetaAsState(singleInterface appliances.InterfaceInResponse) []interface{} {
	var result []interface{}
	result = []interface{}{}

	var meta map[string]interface{}
	meta = map[string]interface{}{}

	meta["name"] = singleInterface.Name
	meta["description"] = singleInterface.Description
	meta["network_id"] = singleInterface.NetworkID
	meta["updatable"] = singleInterface.Updatable

	resultTags := map[string]string{}
	for k, v := range singleInterface.Tags {
		resultTags[k] = v
	}
	meta["tags"] = resultTags

	result = append(result, meta)
	log.Printf("[MYDEBUG] Result Meta: %#v", result)
	return result
}

func getTagsAsOpts(rawTags map[string]interface{}) map[string]string {
	var tags map[string]string
	tags = map[string]string{}
	// rawTags := thisInterface["tags"].(map[string]interface{})
	for k, v := range rawTags {
		tags[k] = v.(string)
	}
	return tags
}

func getInterfaceCreateOpts(d *schema.ResourceData) appliances.CreateOptsInterfaces {
	var interface1 appliances.CreateOptsInterface
	var interfaces appliances.CreateOptsInterfaces

	// Meta part
	rawMeta := d.Get("interface_1_meta").([]interface{})
	rawFips := d.Get("interface_1_fixed_ips").([]interface{})

	// log.Printf("[MYDEBUG] rawMeta: %#v", rawMeta)
	// log.Printf("[MYDEBUG] rawFips: %#v", rawFips)

	for index, rm := range rawMeta {
		thisRawMeta := rm.(map[string]interface{})
		if index == 0 {
			interface1.Name = thisRawMeta["name"].(string)
			interface1.Description = thisRawMeta["description"].(string)
			interface1.NetworkID = thisRawMeta["network_id"].(string)
			tags := getTagsAsOpts(thisRawMeta["tags"].(map[string]interface{}))
			interface1.Tags = tags
		}
	}
	// thisRawMeta := rawMeta[0].(map[string]interface{})

	// FixedIPs part
	var resultFixedIPs [1]appliances.CreateOptsFixedIP
	var fixedIP appliances.CreateOptsFixedIP

	for index, rawFip := range rawFips {
		if index == 0 {
			fip := rawFip.(map[string]interface{})
			// log.Printf("[MYDEBUG] fip: %#v", fip)

			ipAddress := fip["ip_address"].(string)
			fixedIP.IPAddress = ipAddress
			resultFixedIPs[0] = fixedIP

			interface1.FixedIPs = resultFixedIPs
		}
	}
	// rawFip := rawFips[0]
	// log.Printf("[MYDEBUG] rawFip: %#v", rawFip)

	interfaces.Interface1 = interface1

	return interfaces
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
	return result
}
