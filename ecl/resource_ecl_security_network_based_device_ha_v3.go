package ecl

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	security "github.com/nttcom/eclcloud/v4/ecl/security_order/v3/network_based_device_ha"
	"github.com/nttcom/eclcloud/v4/ecl/security_portal/v3/device_interfaces"
)

func resourceSecurityNetworkBasedDeviceHAV3() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedDeviceHAV3Create,
		Read:   resourceSecurityNetworkBasedDeviceHAV3Read,
		Update: resourceSecurityNetworkBasedDeviceHAV3Update,
		Delete: resourceSecurityNetworkBasedDeviceHAV3Delete,

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

func resourceSecurityNetworkBasedDeviceHAV3Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV3Client(GetRegion(d, config))
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
		SOKind:   "AH",
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
		Refresh:      waitForHADeviceOrderComplete(client, order.ID, tenantID, locale),
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

	ids := getNewlyCreatedHADeviceID(allDevicesBefore, allDevicesAfter)
	if len(ids) != 2 {
		return fmt.Errorf("Unable to find newly created HA device after hostname matching. IDs are: %#v", ids)
	}

	log.Printf("[DEBUG] Newly created HA devices are found as ID %s and %s", ids[0], ids[1])

	id := getIDFromHostNames(ids)
	d.SetId(id)

	return resourceSecurityNetworkBasedDeviceHAV3Read(d, meta)
}

func resourceSecurityNetworkBasedDeviceHAV3Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	client, err := config.SecurityOrderV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	if tenantID == "" {
		tenantID = config.TenantID
	}

	log.Printf("[DEBUG] Setting Basic information into state.")

	ids := d.Id()
	idArr := strings.Split(ids, "/")

	id1 := idArr[0]
	id2 := idArr[1]

	device1, err := getHADeviceByHostName(client, id1, tenantID)
	if err != nil {
		return err
	}

	device2, err := getHADeviceByHostName(client, id2, tenantID)
	if err != nil {
		return err
	}

	operatingMode := device1.Cell[4]
	licenseKind := device1.Cell[5]

	d.Set("operating_mode", operatingMode)
	d.Set("license_kind", licenseKind)

	az1 := device1.Cell[7]
	az2 := device2.Cell[7]
	d.Set("host_1_az_group", az1)
	d.Set("host_2_az_group", az2)

	haLink1NetworkID := device1.Cell[9]
	haLink1SubnetID := device1.Cell[10]

	haLink2NetworkID := device1.Cell[12]
	haLink2SubnetID := device1.Cell[13]

	haLink1Host1IPAddress := device1.Cell[11]
	haLink1Host2IPAddress := device2.Cell[11]

	haLink2Host1IPAddress := device1.Cell[14]
	haLink2Host2IPAddress := device2.Cell[14]

	haLink1Info := map[string]string{}
	haLink1Info["network_id"] = haLink1NetworkID
	haLink1Info["subnet_id"] = haLink1SubnetID
	haLink1Info["host_1_ip_address"] = haLink1Host1IPAddress
	haLink1Info["host_2_ip_address"] = haLink1Host2IPAddress

	haLink2Info := map[string]string{}
	haLink1Info["network_id"] = haLink2NetworkID
	haLink1Info["subnet_id"] = haLink2SubnetID
	haLink1Info["host_1_ip_address"] = haLink2Host1IPAddress
	haLink1Info["host_2_ip_address"] = haLink2Host2IPAddress

	log.Printf("[DEBUG] Setting HA Link information into state.")
	d.Set("ha_link_1", haLink1Info)
	d.Set("ha_link_2", haLink2Info)

	log.Printf("[DEBUG] Setting Port information into state.")
	pClient, err := config.SecurityPortalV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security portal client: %s", err)
	}

	hostUUID1, err := getUUIDFromServerHostName(pClient, id1, tenantID)
	if err != nil {
		return fmt.Errorf("Unable to get host UUID of %s: %s", id1, err)
	}

	hostUUID2, err := getUUIDFromServerHostName(pClient, id2, tenantID)
	if err != nil {
		return fmt.Errorf("Unable to get host UUID of %s: %s", id2, err)
	}

	listOpts := device_interfaces.ListOpts{
		TenantID:  tenantID,
		UserToken: pClient.TokenID,
	}

	host1AllDevicePages, err := device_interfaces.List(pClient, hostUUID1, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list interfaces of host-1: %s", err)
	}

	host1AllDevices, err := device_interfaces.ExtractDeviceInterfaces(host1AllDevicePages)
	if err != nil {
		return fmt.Errorf("Unable to extract device interfaces of host-1: %s", err)
	}

	host2AllDevicePages, err := device_interfaces.List(pClient, hostUUID2, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list interfaces of host-2: %s", err)
	}

	host2AllDevices, err := device_interfaces.ExtractDeviceInterfaces(host2AllDevicePages)
	if err != nil {
		return fmt.Errorf("Unable to extract device interfaces of host-2: %s", err)
	}

	deviceInterfaces := []map[string]interface{}{}
	loopCounter := []int{0, 1, 2, 3, 4, 5, 6}

	for range loopCounter {
		thisDeviceInterface := map[string]interface{}{}
		thisDeviceInterface["enable"] = "false"
		deviceInterfaces = append(deviceInterfaces, thisDeviceInterface)
	}

	for _, dev1 := range host1AllDevices {
		log.Printf("[DEBUG] Setting Port information into state from device interface information: %#v", dev1)
		thisDeviceInterface := map[string]interface{}{}

		index1, err := strconv.Atoi(strings.Replace(dev1.MSAPortID, "port", "", 1))
		if err != nil {
			return fmt.Errorf("Error parsing host-1 device interface port number: %s", err)
		}

		index1 -= 4

		if index1 < 0 {
			log.Printf("[DEBUG] Index number %d has found. Skip this interface to store state.", index1)
			continue
		}
		log.Printf("[DEBUG] Processing port %d", index1)

		thisDeviceInterface["enable"] = "true"

		vrrpIP := d.Get(fmt.Sprintf("port.%d.vrrp_ip_address", index1)).(string)
		thisDeviceInterface["vrrp_ip_addess"] = vrrpIP

		thisDeviceInterface["host_1_ip_address"] = dev1.OSIPAddress
		prefixDev1 := d.Get(fmt.Sprintf("port.%d.host_1_ip_address_prefix", index1)).(int)
		thisDeviceInterface["host_1_ip_address_prefix"] = prefixDev1

		for _, dev2 := range host2AllDevices {
			index2, err := strconv.Atoi(strings.Replace(dev2.MSAPortID, "port", "", 1))
			if err != nil {
				return fmt.Errorf("Error parsing host-2 device interface port number: %s", err)
			}

			if index2 != index1 {
				continue
			}
			thisDeviceInterface["host_2_ip_address"] = dev2.OSIPAddress
			prefixDev2 := d.Get(fmt.Sprintf("port.%d.host_2_ip_address_prefix", index1)).(int)
			thisDeviceInterface["host_2_ip_address_prefix"] = prefixDev2
		}

		thisDeviceInterface["network_id"] = dev1.OSNetworkID
		thisDeviceInterface["subnet_id"] = dev1.OSSubnetID

		mtu := d.Get(fmt.Sprintf("port.%d.mtu", index1)).(string)
		comment := d.Get(fmt.Sprintf("port.%d.comment", index1)).(string)

		thisDeviceInterface["mtu"] = mtu

		thisDeviceInterface["comment"] = comment

		thisDeviceInterface["enable_ping"] = d.Get(fmt.Sprintf("port.%d.enable_ping", index1)).(string)

		thisDeviceInterface["vrrp_id"] = d.Get(fmt.Sprintf("port.%d.vrrp_id", index1)).(string)

		deviceInterfaces[index1] = thisDeviceInterface
	}

	d.Set("port", deviceInterfaces)
	log.Printf("[DEBUG] Finished setting state.")

	return nil
}

func resourceSecurityNetworkBasedDeviceHAV3Update(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("locale") || d.HasChange("operating_mode") || d.HasChange("license_kind") {
		log.Printf("[DEBUG] Start changing device by order api.")
		resourceSecurityNetworkBasedDeviceHAV3UpdateOrderAPIPart(d, meta)
	}

	if d.HasChange("port") {
		log.Printf("[DEBUG] Start changing device by portal api.")
		resourceSecurityNetworkBasedDeviceHAV3UpdatePortalAPIPart(d, meta)
	}

	return resourceSecurityNetworkBasedDeviceHAV3Read(d, meta)
}

func resourceSecurityNetworkBasedDeviceHAV3Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.SecurityOrderV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	deleteOpts := security.DeleteOpts{
		SOKind:   "DH",
		TenantID: tenantID,
		GtHost:   gtHostForHADeviceDeleteAsOpts(d),
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
		Refresh:      waitForHADeviceOrderComplete(client, order.ID, tenantID, locale),
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
