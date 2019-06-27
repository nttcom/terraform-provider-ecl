package ecl

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func fixedIPHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if m["ip_address"] != nil && m["subnet_id"] != nil {
		buf.WriteString(
			fmt.Sprintf(
				"%s-%s-",
				m["ip_address"].(string),
				m["subnet_id"].(string),
			))
	}
	return hashcode.String(buf.String())
}

func allowedAddressPairHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if m["ip_address"] != nil &&
		m["mac_address"] != nil &&
		m["type"] != nil &&
		m["vrid"] != nil {
		buf.WriteString(
			fmt.Sprintf(
				"%s-%s-%s-%s-",
				m["ip_address"].(string),
				m["subnet_id"].(string),
				m["type"].(string),
				m["vrid"].(string),
			))
	}
	return hashcode.String(buf.String())
}

func interfaceHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if m["slot_number"] != nil {
		buf.WriteString(fmt.Sprintf("%d-", m["slot_number"].(int)))
	}
	return hashcode.String(buf.String())
}

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

func convertApplianceInterfaceEachFromStructToMap(slotNumber int, structInterface appliances.InterfaceInResponse) map[string]interface{} {
	result := make(map[string]interface{}, 1)

	result["slot_number"] = slotNumber
	result["name"] = structInterface.Name
	result["description"] = structInterface.Description
	result["network_id"] = structInterface.NetworkID
	result["updatable"] = structInterface.Updatable

	resultTags := map[string]string{}
	for k, v := range structInterface.Tags {
		resultTags[k] = v
	}
	result["tags"] = resultTags
	log.Printf("[MYDEBUG] Tag complete")
	log.Printf("%#v", resultTags)

	resultFixedIPs := make([]interface{}, len(structInterface.FixedIPs))
	for i, fixedIP := range structInterface.FixedIPs {
		thisFixedIP := make(map[string]interface{}, 1)
		thisFixedIP["ip_address"] = fixedIP.IPAddress
		thisFixedIP["subnet_id"] = fixedIP.SubnetID

		resultFixedIPs[i] = thisFixedIP
	}
	result["fixed_ips"] = resultFixedIPs
	log.Printf("[MYDEBUG] FixedIPs complete")
	log.Printf("%#v", resultFixedIPs)

	resultAAPs := make([]interface{}, len(structInterface.AllowedAddressPairs))
	for i, aap := range structInterface.AllowedAddressPairs {
		thisAAP := make(map[string]interface{}, 1)
		thisAAP["ip_address"] = aap.IPAddress
		thisAAP["mac_address"] = aap.MACAddress
		thisAAP["type"] = aap.Type
		thisAAP["vrid"] = aap.VRID

		resultAAPs[i] = thisAAP
	}
	result["allowed_address_pairs"] = resultAAPs
	log.Printf("[MYDEBUG] AllowedAddressPairs complete")
	log.Printf("%#v", resultAAPs)

	return result
}

func convertApplianceInterfacesFromStructToMap(structInterfaces appliances.InterfacesInResponse) []map[string]interface{} {

	iface1 := convertApplianceInterfaceEachFromStructToMap(1, structInterfaces.Interface1)
	iface2 := convertApplianceInterfaceEachFromStructToMap(2, structInterfaces.Interface2)
	iface3 := convertApplianceInterfaceEachFromStructToMap(3, structInterfaces.Interface3)
	iface4 := convertApplianceInterfaceEachFromStructToMap(4, structInterfaces.Interface4)
	iface5 := convertApplianceInterfaceEachFromStructToMap(5, structInterfaces.Interface5)
	iface6 := convertApplianceInterfaceEachFromStructToMap(6, structInterfaces.Interface6)
	iface7 := convertApplianceInterfaceEachFromStructToMap(7, structInterfaces.Interface7)
	iface8 := convertApplianceInterfaceEachFromStructToMap(8, structInterfaces.Interface8)

	result := make([]map[string]interface{}, 8)

	result[0] = iface1
	result[1] = iface2
	result[2] = iface3
	result[3] = iface4
	result[4] = iface5
	result[5] = iface6
	result[6] = iface7
	result[7] = iface8

	return result
}

// func getCreateOptsForApplianceUpdate(d *schema.ResourceData) appliances.InterfacesInRequest {
// 	var resultInterfaces appliances.InterfacesInRequest

// 	return appliances.InterfacesInRequest{}
// }

func getFixedIpsForApplianceRequest(thisInterface map[string]interface{}) []appliances.FixedIPInRequest {
	var resultFixedIPs []appliances.FixedIPInRequest
	// thisRawFixedIPs := thisInterface["fixed_ips"].([]interface{})
	thisRawFixedIPs := thisInterface["fixed_ips"].(*schema.Set).List()

	for _, rawFip := range thisRawFixedIPs {
		var fixedIP appliances.FixedIPInRequest

		fip := rawFip.(map[string]interface{})
		ipAddress := fip["ip_address"].(string)
		fixedIP.IPAddress = ipAddress
		resultFixedIPs = append(resultFixedIPs, fixedIP)
	}

	return resultFixedIPs
}

func getAllowedAddressPairsForApplianceRequest(thisInterface map[string]interface{}) []appliances.AllowedAddressPairInRequest {
	var resultAllowedAddressPairs []appliances.AllowedAddressPairInRequest
	// thisRawAllowedAddressPairs := thisInterface["allowed_address_pairs"].([]interface{})
	thisRawAllowedAddressPairs := thisInterface["allowed_address_pairs"].(*schema.Set).List()

	for _, rawAAP := range thisRawAllowedAddressPairs {
		var allowedAddressPair appliances.AllowedAddressPairInRequest
		aap := rawAAP.(map[string]interface{})
		ipAddress := aap["ip_address"].(string)
		tp := aap["type"].(string)
		macAddress := aap["mac_address"].(string)
		vrid := aap["vrid"].(string)

		allowedAddressPair.Type = tp
		allowedAddressPair.IPAddress = ipAddress
		allowedAddressPair.MACAddress = macAddress
		allowedAddressPair.VRID = vrid
		resultAllowedAddressPairs = append(resultAllowedAddressPairs, allowedAddressPair)
	}

	return resultAllowedAddressPairs
}

func getTagsForApplianceRequest(thisInterface map[string]interface{}) map[string]string {
	var tags map[string]string
	rawTags := thisInterface["tags"].(map[string]interface{})
	for k, v := range rawTags {
		tags[k] = v.(string)
	}
	return tags
}

func getCreateOptsForApplianceCreate(d *schema.ResourceData) appliances.InterfacesInCreate {
	rawInterfaces := d.Get("interfaces").(*schema.Set).List()

	var resultInterfaces appliances.InterfacesInCreate

	for _, tmpIface := range rawInterfaces {

		thisInterface := tmpIface.(map[string]interface{})
		slotNumber := thisInterface["slot_number"].(int)

		if slotNumber != 1 {
			continue
		}

		var iface appliances.InterfaceInCreate

		// top level data
		iface.Name = thisInterface["name"].(string)
		iface.Description = thisInterface["description"].(string)
		iface.NetworkID = thisInterface["network_id"].(string)

		// tags
		tags := getTagsForApplianceRequest(thisInterface)
		iface.Tags = tags

		resultFixedIPs := getFixedIpsForApplianceRequest(thisInterface)
		iface.FixedIPs = resultFixedIPs

		// resultAllowedAddressPairs := getAllowedAddressPairsForApplianceRequest(thisInterface)
		// iface.AllowedAddressPairs = &resultAllowedAddressPairs

		resultInterfaces.Interface1 = iface
	}

	return resultInterfaces
}

func getCreateOptsForApplianceUpdate(d *schema.ResourceData) appliances.InterfacesInRequest {
	// func convertApplianceInterfacesFromSchemaToStruct(d *schema.ResourceData) appliances.InterfacesInRequest {
	rawInterfaces := d.Get("interfaces").(*schema.Set).List()

	var resultInterfaces appliances.InterfacesInRequest

	for _, tmpIface := range rawInterfaces {

		thisInterface := tmpIface.(map[string]interface{})
		slotNumber := thisInterface["slot_number"].(int)

		var iface appliances.InterfaceInRequest

		// top level data
		iface.Name = thisInterface["name"].(string)
		iface.Description = thisInterface["description"].(string)
		iface.NetworkID = thisInterface["network_id"].(string)

		// tags
		tags := getTagsForApplianceRequest(thisInterface)
		iface.Tags = tags

		resultFixedIPs := getFixedIpsForApplianceRequest(thisInterface)
		iface.FixedIPs = resultFixedIPs

		resultAllowedAddressPairs := getAllowedAddressPairsForApplianceRequest(thisInterface)
		iface.AllowedAddressPairs = resultAllowedAddressPairs

		switch slotNumber {
		case 1:
			resultInterfaces.Interface1 = iface
			break
		case 2:
			resultInterfaces.Interface2 = iface
			break
		case 3:
			resultInterfaces.Interface3 = iface
			break
		case 4:
			resultInterfaces.Interface4 = iface
			break
		case 5:
			resultInterfaces.Interface5 = iface
			break
		case 6:
			resultInterfaces.Interface6 = iface
			break
		case 7:
			resultInterfaces.Interface7 = iface
			break
		case 8:
			resultInterfaces.Interface8 = iface
			break
		default:
		}
	}

	return resultInterfaces
}
