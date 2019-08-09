package ecl

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud"

	"github.com/nttcom/eclcloud/ecl/security_portal/v1/ports"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_ha"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/device_interfaces"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/devices"
)

func resourceSecurityNetworkBasedDeviceHAV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedDeviceHAV1Create,
		Read:   resourceSecurityNetworkBasedDeviceHAV1Read,
		Update: resourceSecurityNetworkBasedDeviceHAV1Update,
		Delete: resourceSecurityNetworkBasedDeviceHAV1Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: haDeviceSchema(),
	}
}

func resourceSecurityNetworkBasedDeviceHAV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	listOpts := security.ListOpts{
		TenantID: tenantID,
		Locale:   locale,
	}

	allPagesBefore, err := security.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to get page of devices before creation: %s", err)
	}
	var allDevicesBefore []security.HADevice

	err = security.ExtractHADevicesInto(allPagesBefore, &allDevicesBefore)

	if err != nil {
		return fmt.Errorf("Unable to retrieve device list before create: %s", err)
	}
	log.Printf("[DEBUG] allSingleDevices before creation: %#v", allDevicesBefore)
	createOpts := security.CreateOpts{
		SOKind:   "A",
		TenantID: tenantID,
		Locale:   locale,
		GtHost:   gtHostForHADeviceCreateAsOpts(d),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	order, err := security.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL security HA device: %s", err)
	}

	log.Printf("[DEBUG] SingleDevice has successfully created with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.ID, tenantID, locale, deviceType),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceHACreatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for HA device order status (%s) to become ready: %s",
			order.ID, err)
	}
	log.Printf("[DEBUG] Finish waiting for HA device create order becomes COMPLETE")

	allPagesAfter, err := security.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to get page of devices after creation: %s", err)
	}
	var allDevicesAfter []security.HADevice

	err = security.ExtractHADevicesInto(allPagesAfter, &allDevicesAfter)
	if err != nil {
		return fmt.Errorf("Unable to retrieve list of HA device after create: %s", err)
	}
	log.Printf("[DEBUG] allSingleDevices after creation: %#v", allDevicesAfter)

	if len(allDevicesBefore) == len(allDevicesAfter) {
		return fmt.Errorf("Unable to find newly created HA device")
	}

	id := getNewlyCreatedDeviceID(allDevicesBefore, allDevicesAfter, d.Get("operating_mode").(string))
	if id == "" {
		return fmt.Errorf("Unable to find newly created HA device after hostname matching")
	}

	log.Printf("[DEBUG] Newly created HA device is found as ID: %s", id)

	d.SetId(id)

	return resourceSecurityNetworkBasedDeviceHAV1Read(d, meta)
}

func getUUIDFromServerHostName(client *eclcloud.ServiceClient, hostName string) (string, error) {

	listOpts := devices.ListOpts{
		TenantID:  os.Getenv("OS_TENANT_ID"),
		UserToken: client.TokenID,
	}

	allPages, err := devices.List(client, listOpts).AllPages()
	if err != nil {
		return "", fmt.Errorf("Unable to list HA device to get device UUID: %s", err)
	}
	var allDevices []devices.Device

	err = devices.ExtractDevicesInto(allPages, &allDevices)
	if err != nil {
		return "", fmt.Errorf("Unable to extract list of HA device by portal api: %s", err)
	}

	for _, device := range allDevices {
		if device.MSADeviceID == hostName {
			log.Printf("[DEBUG] Host UUID looking result: Host %s has UUID %s", hostName, device.OSServerID)
			return device.OSServerID, nil
		}
	}

	return "", fmt.Errorf("Unable to find corresponding server of %s", hostName)
}

func resourceSecurityNetworkBasedDeviceHAV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	// Main Part
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := os.Getenv("OS_TENANT_ID")
	locale := d.Get("locale")

	deviceType := getTypeOfSingleDevice(d)
	device, err := getSingleDeviceByHostName(client, deviceType, d.Id())
	if err != nil {
		return err
	}

	d.Set("tenant_id", tenantID)
	d.Set("locale", locale)

	operatingMode := device.Cell[3]
	licenseKind := device.Cell[4]

	var azGroup string
	if operatingMode == "WAF" {
		azGroup = device.Cell[5]
	} else {
		azGroup = device.Cell[6]
	}

	d.Set("operating_mode", operatingMode)

	d.Set("license_kind", licenseKind)
	d.Set("az_group", azGroup)

	// Device Interface Part
	pClient, err := config.securityPortalV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security portal client: %s", err)
	}

	hostUUID, err := getUUIDFromServerHostName(pClient, d.Id())
	if err != nil {
		return fmt.Errorf("Unable to get host UUID: %s", err)
	}

	listOpts := device_interfaces.ListOpts{
		TenantID:  os.Getenv("OS_TENANT_ID"),
		UserToken: pClient.TokenID,
	}

	allDevicePages, err := device_interfaces.List(pClient, hostUUID, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list interfaces: %s", err)
	}

	allDevices, err := device_interfaces.ExtractDeviceInterfaces(allDevicePages)
	if err != nil {
		return fmt.Errorf("Unable to extract device interfaces: %s", err)
	}

	// initialize
	deviceInterfaces := []map[string]interface{}{}
	var loopCounter []int

	if deviceType == "WAF" {
		loopCounter = []int{0}
	} else {
		loopCounter = []int{0, 1, 2, 3, 4, 5, 6}
	}

	for range loopCounter {
		thisDeviceInterface := map[string]interface{}{}
		thisDeviceInterface["enable"] = "false"
		deviceInterfaces = append(deviceInterfaces, thisDeviceInterface)
	}

	for _, dev := range allDevices {
		thisDeviceInterface := map[string]interface{}{}

		index, err := strconv.Atoi(strings.Replace(dev.MSAPortID, "port", "", 1))
		if err != nil {
			return fmt.Errorf("Error parsing device interface port number: %s", err)
		}

		if deviceType == "WAF" {
			// map port 2(actual) as 0 to handle by list index in WAF
			index -= 2
		} else {
			// map port 4 to 10 (actual) as 0 to 6 to handle by list index in FW/UTM
			index -= 4
		}

		if index < 0 {
			return fmt.Errorf("Wrong index number is returned from device interface list API. %s", err)
		}

		thisDeviceInterface["enable"] = "true"
		thisDeviceInterface["ip_address"] = dev.OSIPAddress

		prefix := d.Get(fmt.Sprintf("port.%d.ip_address_prefix", index)).(int)
		thisDeviceInterface["ip_addess_prefix"] = prefix

		thisDeviceInterface["network_id"] = dev.OSNetworkID
		thisDeviceInterface["subnet_id"] = dev.OSSubnetID

		mtu := d.Get(fmt.Sprintf("port.%d.mtu", index)).(string)
		comment := d.Get(fmt.Sprintf("port.%d.comment", index)).(string)
		thisDeviceInterface["mtu"] = mtu
		thisDeviceInterface["comment"] = comment

		deviceInterfaces[index] = thisDeviceInterface
	}

	d.Set("port", deviceInterfaces)
	return nil
}

func resourceSecurityNetworkBasedSingleDevicePortsForUpdate(d *schema.ResourceData) (ports.UpdateOpts, error) {
	resultPorts := []ports.SinglePort{}

	ifaces := d.Get("port").([]interface{})
	log.Printf("[DEBUG] Retrieved port information for update: %#v", ifaces)
	for _, iface := range ifaces {
		p := ports.SinglePort{}

		if _, ok := iface.(map[string]interface{}); ok {
			thisInterface := iface.(map[string]interface{})

			if thisInterface["enable"].(string) == "true" {
				p.EnablePort = "true"

				ipAddress := thisInterface["ip_address"].(string)
				prefix := thisInterface["ip_address_prefix"].(int)

				p.IPAddress = fmt.Sprintf("%s/%d", ipAddress, prefix)

				p.NetworkID = thisInterface["network_id"].(string)
				p.SubnetID = thisInterface["subnet_id"].(string)
				p.MTU = thisInterface["mtu"].(string)
				p.Comment = thisInterface["comment"].(string)
			} else {
				p.EnablePort = "false"
			}
		}
		resultPorts = append(resultPorts, p)
	}

	log.Printf("[DEBUG] Port update parameters: %#v", resultPorts)
	result := ports.UpdateOpts{}
	result.Port = resultPorts
	return result, nil
}

func resourceSecurityNetworkBasedDeviceHAV1Update(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("locale") || d.HasChange("operating_mode") || d.HasChange("license_kind") {
		log.Printf("[DEBUG] Start changing device by order api.")
		resourceSecurityNetworkBasedDeviceHAV1UpdateOrderAPIPart(d, meta)
	}

	if d.HasChange("port") {
		log.Printf("[DEBUG] Start changing device by portal api.")
		resourceSecurityNetworkBasedDeviceHAV1UpdatePortalAPIPart(d, meta)
	}

	return resourceSecurityNetworkBasedDeviceHAV1Read(d, meta)
}

func resourceSecurityNetworkBasedDeviceHAV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	deleteOpts := security.DeleteOpts{
		SOKind:   "D",
		TenantID: tenantID,
		GtHost:   gtHostForSingleDeviceDeleteAsOpts(d),
	}

	log.Printf("[DEBUG] Delete Options: %#v", deleteOpts)

	deviceType := getTypeOfSingleDevice(d)
	order, err := security.Delete(client, deviceType, deleteOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error deleting ECL security HA device: %s", err)
	}

	log.Printf("[DEBUG] Delete request has successfully accepted with order: %#v", order)

	log.Printf("[DEBUG] Start waiting for HA device order becomes COMPLETE ...")

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.ID, tenantID, locale, deviceType),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceHADeletePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for HA device order status (%s) to become ready: %s",
			order.ID, err)
	}

	d.SetId("")

	return nil
}
