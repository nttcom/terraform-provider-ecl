package ecl

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

const maxNumberOfInterfaces = 8

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

func getInterfaceAllowedAddressPairsAsState(allowedAddressPairs []appliances.AllowedAddressPairInResponse) []interface{} {
	result := make([]interface{}, len(allowedAddressPairs))
	for i, aap := range allowedAddressPairs {
		thisAAP := make(map[string]interface{}, 1)
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
		thisFixedIP := make(map[string]interface{}, 1)
		thisFixedIP["ip_address"] = fixedIP.IPAddress
		thisFixedIP["subnet_id"] = fixedIP.SubnetID

		result[i] = thisFixedIP
	}
	log.Printf("[MYDEBUG] Result FixedIPs: %#v", result)
	return result
}

func getInterfaceMetaAsState(singleInterface appliances.InterfaceInResponse) []map[string]interface{} {
	result := make([]map[string]interface{}, 1)
	meta := make(map[string]interface{}, 1)

	meta["name"] = singleInterface.Name
	meta["description"] = singleInterface.Description
	meta["network_id"] = singleInterface.NetworkID
	meta["updatable"] = singleInterface.Updatable

	resultTags := map[string]string{}
	for k, v := range singleInterface.Tags {
		resultTags[k] = v
	}
	meta["tags"] = resultTags

	log.Printf("[MYDEBUG] Result Meta: %#v", result)

	result[0] = meta
	return result
}

func getTagsAsOpts(rawTags map[string]interface{}) map[string]string {
	var tags map[string]string
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
	rawMeta := d.Get("interface_1_meta").(*schema.Set).List()
	rawFips := d.Get("interface_1_fixed_ips").(*schema.Set).List()

	log.Printf("[MYDEBUG] rawMeta: %#v", rawMeta)
	log.Printf("[MYDEBUG] rawFips: %#v", rawFips)

	thisRawMeta := rawMeta[0].(map[string]interface{})
	interface1.Name = thisRawMeta["name"].(string)
	interface1.Description = thisRawMeta["description"].(string)
	interface1.NetworkID = thisRawMeta["network_id"].(string)

	tags := getTagsAsOpts(thisRawMeta["tags"].(map[string]interface{}))
	interface1.Tags = tags

	// FixedIPs part
	var resultFixedIPs [1]appliances.CreateOptsFixedIP
	var fixedIP appliances.CreateOptsFixedIP

	rawFip := rawFips[0]
	log.Printf("[MYDEBUG] rawFip: %#v", rawFip)

	fip := rawFip.(map[string]interface{})
	log.Printf("[MYDEBUG] fip: %#v", fip)

	ipAddress := fip["ip_address"].(string)
	fixedIP.IPAddress = ipAddress
	resultFixedIPs[0] = fixedIP

	interface1.FixedIPs = resultFixedIPs

	interfaces.Interface1 = interface1

	return interfaces
}
