package ecl

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	security "github.com/nttcom/eclcloud/v4/ecl/security_order/v3/network_based_device_single"
	"github.com/nttcom/eclcloud/v4/ecl/security_portal/v3/device_interfaces"
)

func resourceSecurityNetworkBasedDeviceSingleV3() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedDeviceSingleV3Create,
		Read:   resourceSecurityNetworkBasedDeviceSingleV3Read,
		Update: resourceSecurityNetworkBasedDeviceSingleV3Update,
		Delete: resourceSecurityNetworkBasedDeviceSingleV3Delete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: singleDeviceSchema(),
	}
}

func resourceSecurityNetworkBasedDeviceSingleV3Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	deviceType := getTypeOfSingleDevice(d)

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	listOpts := security.ListOpts{
		TenantID: tenantID,
		Locale:   locale,
	}

	allPagesBefore, err := security.List(client, deviceType, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to get page of devices before creation: %s", err)
	}
	var allDevicesBefore []security.SingleDevice

	err = security.ExtractSingleDevicesInto(allPagesBefore, &allDevicesBefore)

	if err != nil {
		return fmt.Errorf("Unable to retrieve device list before create: %s", err)
	}
	log.Printf("[DEBUG] allSingleDevices before creation: %#v", allDevicesBefore)
	createOpts := security.CreateOpts{
		SOKind:   "A",
		TenantID: tenantID,
		Locale:   locale,
		GtHost:   gtHostForSingleDeviceCreateAsOpts(d),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	order, err := security.Create(client, deviceType, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] SingleDevice has successfully created with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.ID, tenantID, locale, deviceType),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceSingleCreatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}
	log.Printf("[DEBUG] Finish waiting for single device create order becomes COMPLETE")

	allPagesAfter, err := security.List(client, deviceType, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to get page of devices after creation: %s", err)
	}
	var allDevicesAfter []security.SingleDevice

	err = security.ExtractSingleDevicesInto(allPagesAfter, &allDevicesAfter)
	if err != nil {
		return fmt.Errorf("Unable to retrieve list of single device after create: %s", err)
	}
	log.Printf("[DEBUG] allSingleDevices after creation: %#v", allDevicesAfter)

	if len(allDevicesBefore) == len(allDevicesAfter) {
		return fmt.Errorf("Unable to find newly created single device")
	}

	id := getNewlyCreatedDeviceID(allDevicesBefore, allDevicesAfter, d.Get("operating_mode").(string))
	if id == "" {
		return fmt.Errorf("Unable to find newly created single device after hostname matching")
	}

	log.Printf("[DEBUG] Newly created single device is found as ID: %s", id)

	d.SetId(id)

	return resourceSecurityNetworkBasedDeviceSingleV3Read(d, meta)
}

func resourceSecurityNetworkBasedDeviceSingleV3Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	client, err := config.securityOrderV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	if tenantID == "" {
		tenantID = config.TenantID
	}

	deviceType := getTypeOfSingleDevice(d)
	device, err := getSingleDeviceByHostName(client, deviceType, d.Id(), tenantID)
	if err != nil {
		return err
	}

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
	pClient, err := config.securityPortalV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security portal client: %s", err)
	}

	hostUUID, err := getUUIDFromServerHostName(pClient, d.Id(), tenantID)
	if err != nil {
		return fmt.Errorf("Unable to get host UUID: %s", err)
	}

	listOpts := device_interfaces.ListOpts{
		TenantID:  tenantID,
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

func resourceSecurityNetworkBasedDeviceSingleV3Update(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("locale") || d.HasChange("operating_mode") || d.HasChange("license_kind") {
		log.Printf("[DEBUG] Start changing device by order api.")
		resourceSecurityNetworkBasedDeviceSingleV3UpdateOrderAPIPart(d, meta)
	}

	if d.HasChange("port") {
		log.Printf("[DEBUG] Start changing device by portal api.")
		resourceSecurityNetworkBasedDeviceSingleV3UpdatePortalAPIPart(d, meta)
	}

	return resourceSecurityNetworkBasedDeviceSingleV3Read(d, meta)
}

func resourceSecurityNetworkBasedDeviceSingleV3Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV3Client(GetRegion(d, config))
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
		return fmt.Errorf("Error deleting ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] Delete request has successfully accepted with order: %#v", order)

	log.Printf("[DEBUG] Start waiting for single device order becomes COMPLETE ...")

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.ID, tenantID, locale, deviceType),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceSingleDeletePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}

	d.SetId("")

	return nil
}
