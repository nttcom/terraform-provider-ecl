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
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/processes"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_single"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/device_interfaces"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/devices"

	"github.com/nttcom/eclcloud/ecl/security_order/v1/service_order_status"
)

const securityDeviceSinglePollIntervalSec = 30
const securityDeviceSingleCreatePollInterval = securityDeviceSinglePollIntervalSec * time.Second
const securityDeviceSingleUpdatePollInterval = securityDeviceSinglePollIntervalSec * time.Second
const securityDeviceSingleDeletePollInterval = securityDeviceSinglePollIntervalSec * time.Second

func resourceSecurityNetworkBasedDeviceSingleV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedDeviceSingleV1Create,
		Read:   resourceSecurityNetworkBasedDeviceSingleV1Read,
		Update: resourceSecurityNetworkBasedDeviceSingleV1Update,
		Delete: resourceSecurityNetworkBasedDeviceSingleV1Delete,

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

func getTypeOfSingleDevice(d *schema.ResourceData) string {
	operatingMode := d.Get("operating_mode").(string)
	switch operatingMode {
	case "WAF":
		return "WAF"
	default:
		return "UTM"
	}
}

func resourceSecurityNetworkBasedDeviceSingleV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
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

	return resourceSecurityNetworkBasedDeviceSingleV1Read(d, meta)
}

func getNewlyCreatedDeviceID(beforeTemp, afterTemp []security.SingleDevice, deviceType string) string {
	var before, after []security.SingleDevice
	before = []security.SingleDevice{}
	after = []security.SingleDevice{}

	for _, a := range afterTemp {
		if a.Cell[3] == deviceType {
			after = append(after, a)
		}
	}

	for _, b := range beforeTemp {
		if b.Cell[3] == deviceType {
			before = append(before, b)
		}
	}

	for _, af := range after {
		hostNameAfter := af.Cell[2]
		match := false
		for _, bf := range before {
			if bf.Cell[3] != deviceType {
				continue
			}
			hostNameBefore := bf.Cell[2]
			if hostNameAfter == hostNameBefore {
				match = true
			}
		}
		if !match {
			return hostNameAfter
		}
	}
	return ""
}

func getUUIDFromServerHostName(client *eclcloud.ServiceClient, hostName string) (string, error) {

	listOpts := devices.ListOpts{
		TenantID:  os.Getenv("OS_TENANT_ID"),
		UserToken: client.TokenID,
	}

	allPages, err := devices.List(client, listOpts).AllPages()
	if err != nil {
		return "", fmt.Errorf("Unable to list single device to get device UUID: %s", err)
	}
	var allDevices []devices.Device

	err = devices.ExtractDevicesInto(allPages, &allDevices)
	if err != nil {
		return "", fmt.Errorf("Unable to extract list of single device by portal api: %s", err)
	}

	for _, device := range allDevices {
		if device.MSADeviceID == hostName {
			log.Printf("[DEBUG] Host UUID looking result: Host %s has UUID %s", hostName, device.OSServerID)
			return device.OSServerID, nil
		}
	}

	return "", fmt.Errorf("Unable to find corresponding server of %s", hostName)
}

func resourceSecurityNetworkBasedDeviceSingleV1Read(d *schema.ResourceData, meta interface{}) error {
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

	deviceInterfaces := [7]map[string]interface{}{}
	// initialize
	for index := range []int{0, 1, 2, 3, 4, 5, 6} {
		thisDeviceInterface := map[string]interface{}{}
		thisDeviceInterface["enable"] = "false"
		deviceInterfaces[index] = thisDeviceInterface
	}

	for _, dev := range allDevices {
		thisDeviceInterface := map[string]interface{}{}

		index, err := strconv.Atoi(strings.Replace(dev.MSAPortID, "port", "", 1))
		if err != nil {
			return fmt.Errorf("Error parsing device interface port number: %s", err)
		}
		index -= 4
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

func resourceSecurityNetworkBasedDeviceSingleV1UpdateOrderAPIPart(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts := security.UpdateOpts{
		SOKind:   "M",
		TenantID: tenantID,
		Locale:   locale,
		GtHost:   gtHostForSingleDeviceUpdateAsOpts(d),
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)

	deviceType := getTypeOfSingleDevice(d)
	order, err := security.Update(client, deviceType, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] Update request has successfully accepted with order: %#v", order)

	log.Printf("[DEBUG] Start waiting for single device order becomes COMPLETE ...")

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.ID, tenantID, locale, deviceType),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceSingleUpdatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	log.Printf("[DEBUG] Finish waiting for single device update order becomes COMPLETE")

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}

	return nil
}

func resourceSecurityNetworkBasedSingleDevicePortsForUpdate(d *schema.ResourceData) (ports.UpdateOpts, error) {
	resultPorts := [7]ports.SinglePort{}

	ifaces := d.Get("port").([]interface{})
	log.Printf("[DEBUG] Retrieved port information for update: %#v", ifaces)
	for index, iface := range ifaces {
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
		resultPorts[index] = p
	}

	log.Printf("[DEBUG] Port update parameters: %#v", resultPorts)
	result := ports.UpdateOpts{}
	result.Port = resultPorts
	return result, nil
}

func resourceSecurityNetworkBasedDeviceSingleV1UpdatePortalAPIPart(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityPortalV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security portal client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts, err := resourceSecurityNetworkBasedSingleDevicePortsForUpdate(d)
	if err != nil {
		return fmt.Errorf("Error getting port option in update: %s", err)
	}
	updateQueryOpts := ports.UpdateQueryOpts{
		TenantID:  tenantID,
		UserToken: client.TokenID,
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)
	log.Printf("[DEBUG] Update Query Options: %#v", updateQueryOpts)

	deviceType := getTypeOfSingleDevice(d)
	process, err := ports.Update(
		client,
		strings.ToLower(deviceType),
		d.Id(),
		updateOpts,
		updateQueryOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error updating ECL security single device port: %s", err)
	}

	log.Printf("[DEBUG] Update request has successfully accepted with process: %#v", process)

	log.Printf("[DEBUG] Start waiting for single device process becomes ENDED ...")

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"RUNNING"},
		Target:       []string{"ENDED"},
		Refresh:      waitForSingleDeviceProcessComplete(client, process.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceSingleUpdatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device port management status (%s) to become ready: %s",
			process.ID, err)
	}

	log.Printf("[DEBUG] Finish waiting for single device portal api order becomes COMPLETE")

	return nil
}

func resourceSecurityNetworkBasedDeviceSingleV1Update(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("locale") || d.HasChange("operating_mode") || d.HasChange("license_kind") {
		log.Printf("[DEBUG] Start changing device by order api.")
		resourceSecurityNetworkBasedDeviceSingleV1UpdateOrderAPIPart(d, meta)
	}

	if d.HasChange("port") {
		log.Printf("[DEBUG] Start changing device by portal api.")
		resourceSecurityNetworkBasedDeviceSingleV1UpdatePortalAPIPart(d, meta)
	}

	return resourceSecurityNetworkBasedDeviceSingleV1Read(d, meta)
}

func resourceSecurityNetworkBasedDeviceSingleV1Delete(d *schema.ResourceData, meta interface{}) error {
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

func waitForSingleDeviceOrderComplete(client *eclcloud.ServiceClient, soID, tenantID, locale, deviceType string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		opts := service_order_status.GetOpts{
			Locale:   locale,
			TenantID: tenantID,
			SoID:     soID,
		}
		order, err := service_order_status.Get(client, deviceType, opts).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Service Order Status: %+v", order)

		if order.ProgressRate == 100 {
			return order, "COMPLETE", nil
		}

		return order, "PROCESSING", nil
	}
}

func getSingleDeviceByHostName(client *eclcloud.ServiceClient, deviceType, hostName string) (security.SingleDevice, error) {
	log.Printf("[DEBUG] Start getting %s by HostName %s ...", deviceType, hostName)
	var sd = security.SingleDevice{}

	listOpts := security.ListOpts{
		TenantID: os.Getenv("OS_TENANT_ID"),
		Locale:   "en",
	}

	allPages, err := security.List(client, deviceType, listOpts).AllPages()
	if err != nil {
		return sd, fmt.Errorf("Unable to list single device to get hostname from result: %s", err)
	}
	var allDevices []security.SingleDevice

	err = security.ExtractSingleDevicesInto(allPages, &allDevices)
	if err != nil {
		return sd, fmt.Errorf("Unable to extract result of single device list api: %s", err)
	}

	var thisDevice security.SingleDevice
	var found bool
	for _, device := range allDevices {
		if device.Cell[2] == hostName {
			thisDevice = device
			found = true
			break
		}
	}
	if !found {
		return sd, fmt.Errorf("[DEBUG] Specified single device %s not found", hostName)
	}
	log.Printf("[DEBUG] Host has found as: %#v", thisDevice)
	return thisDevice, nil
}

func gtHostForSingleDeviceCreateAsOpts(d *schema.ResourceData) [1]security.GtHostInCreate {
	result := [1]security.GtHostInCreate{}

	gtHost := security.GtHostInCreate{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.AZGroup = d.Get("az_group").(string)

	result[0] = gtHost

	return result
}

func gtHostForSingleDeviceUpdateAsOpts(d *schema.ResourceData) [1]security.GtHostInUpdate {
	result := [1]security.GtHostInUpdate{}

	gtHost := security.GtHostInUpdate{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.HostName = d.Id()

	result[0] = gtHost

	return result
}

func gtHostForSingleDeviceDeleteAsOpts(d *schema.ResourceData) [1]security.GtHostInDelete {
	result := [1]security.GtHostInDelete{}

	gtHost := security.GtHostInDelete{}

	gtHost.HostName = d.Id()

	result[0] = gtHost

	return result
}

func waitForSingleDeviceProcessComplete(client *eclcloud.ServiceClient, processID, tenantID, locale string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		opts := processes.GetOpts{
			TenantID:  tenantID,
			UserToken: client.TokenID,
		}
		process, err := processes.Get(client, processID, opts).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Service Process Status: %#v", process)

		return process, process.Status.Status, nil
	}
}
