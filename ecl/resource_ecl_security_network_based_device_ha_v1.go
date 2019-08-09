package ecl

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_ha"
	// "github.com/nttcom/eclcloud/ecl/security_portal/v1/device_interfaces"
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

	return resourceSecurityNetworkBasedDeviceHAV1Read(d, meta)
}

func resourceSecurityNetworkBasedDeviceHAV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := os.Getenv("OS_TENANT_ID")
	locale := d.Get("locale")
	d.Set("tenant_id", tenantID)
	d.Set("locale", locale)

	ids := d.Id()
	idArr := strings.Split(ids, "/")

	id1 := idArr[0]
	id2 := idArr[1]

	device1, err := getHADeviceByHostName(client, id1)
	if err != nil {
		return err
	}

	device2, err := getHADeviceByHostName(client, id2)
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

	d.Set("ha_link_1", haLink1Info)
	d.Set("ha_link_2", haLink2Info)

	// Device Interface is later.

	return nil
}

func resourceSecurityNetworkBasedDeviceHAV1Update(d *schema.ResourceData, meta interface{}) error {

	// if d.HasChange("locale") || d.HasChange("operating_mode") || d.HasChange("license_kind") {
	// 	log.Printf("[DEBUG] Start changing device by order api.")
	// 	resourceSecurityNetworkBasedDeviceHAV1UpdateOrderAPIPart(d, meta)
	// }

	// if d.HasChange("port") {
	// 	log.Printf("[DEBUG] Start changing device by portal api.")
	// 	resourceSecurityNetworkBasedDeviceHAV1UpdatePortalAPIPart(d, meta)
	// }

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
