package ecl

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/service_order_status"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/single_devices"
	// "github.com/nttcom/eclcloud/ecl/sss/v1/users"
)

const securitySingleDevicePollIntervalSec = 30
const securitySingleDeviceCreatePollInterval = securitySingleDevicePollIntervalSec * time.Second
const securitySingleDeviceUpdatePollInterval = securitySingleDevicePollIntervalSec * time.Second
const securitySingleDeviceDeletePollInterval = securitySingleDevicePollIntervalSec * time.Second

// func SecurityNetworkBasedSingleDeviceV1() *schema.Resource {
func resourceSecurityNetworkBasedSingleDeviceV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedSingleDeviceV1Create,
		Read:   resourceSecurityNetworkBasedSingleDeviceV1Read,
		Update: resourceSecurityNetworkBasedSingleDeviceV1Update,
		Delete: resourceSecurityNetworkBasedSingleDeviceV1Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// "so_kind": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Required: true,
			// 	ValidateFunc: validation.StringInSlice([]string{
			// 		"A", "M",
			// 	}, false),
			// },

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"locale": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ja", "en",
				}, false),
			},

			"operating_mode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"FW", "UTM", "WAF",
				}, false),
			},

			"license_kind": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"02", "04", "08",
				}, false),
			},

			"az_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func gtHostForSingleDeviceCreateAsOpts(d *schema.ResourceData) [1]single_devices.GtHostInCreate {
	result := [1]single_devices.GtHostInCreate{}

	gtHost := single_devices.GtHostInCreate{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.AZGroup = d.Get("az_group").(string)

	result[0] = gtHost

	return result
}

func gtHostForSingleDeviceUpdateAsOpts(d *schema.ResourceData) [1]single_devices.GtHostInUpdate {
	result := [1]single_devices.GtHostInUpdate{}

	gtHost := single_devices.GtHostInUpdate{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.HostName = d.Id()

	result[0] = gtHost

	return result
}

func gtHostForSingleDeviceDeleteAsOpts(d *schema.ResourceData) [1]single_devices.GtHostInDelete {
	result := [1]single_devices.GtHostInDelete{}

	gtHost := single_devices.GtHostInDelete{}

	gtHost.HostName = d.Id()

	result[0] = gtHost

	return result
}

func resourceSecurityNetworkBasedSingleDeviceV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	listOpts := single_devices.ListOpts{
		TenantID: tenantID,
		Locale:   locale,
	}

	allPagesBefore, err := single_devices.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list of devices, before creating new single device: %s", err)
	}
	var allDevicesBefore []single_devices.SingleDevice

	err = single_devices.ExtractSingleDevicesInto(allPagesBefore, &allDevicesBefore)

	if err != nil {
		return fmt.Errorf("Unable to retrieve device list before create: %s", err)
	}
	log.Printf("[DEBUG] allSingleDevices before creation: %#v", allDevicesBefore)
	createOpts := single_devices.CreateOpts{
		SOKind:   "A",
		TenantID: tenantID,
		Locale:   locale,
		GtHost:   gtHostForSingleDeviceCreateAsOpts(d),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	order, err := single_devices.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] User has successfully created with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securitySingleDeviceCreatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}

	allPagesAfter, err := single_devices.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list after to create single device: %s", err)
	}
	var allDevicesAfter []single_devices.SingleDevice

	err = single_devices.ExtractSingleDevicesInto(allPagesAfter, &allDevicesAfter)
	if err != nil {
		return fmt.Errorf("Unable to retrieve list of devices after create: %s", err)
	}
	log.Printf("[DEBUG] allSingleDevices after creation: %#v", allDevicesAfter)

	if len(allDevicesBefore) == len(allDevicesAfter) {
		return fmt.Errorf("Unable to find newly created device")
	}

	id := getNewlyCreatedDeviceID(allDevicesBefore, allDevicesAfter)

	if id == "" {
		return fmt.Errorf("Unable to find newly created device after hostname matching")

	}

	log.Printf("[DEBUG] Newly created device is found as ID: %s", id)

	d.SetId(id)

	return resourceSecurityNetworkBasedSingleDeviceV1Read(d, meta)
}

func getNewlyCreatedDeviceID(before, after []single_devices.SingleDevice) string {
	// var beforeHostNames, afterHostNames []string
	// var HostNameAfter string

	// var result string
	for _, af := range after {
		hostNameAfter := af.Cell[2]
		match := false
		for _, bf := range before {
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

func waitForSingleDeviceOrderComplete(client *eclcloud.ServiceClient, soID, tenantID, locale string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		opts := service_order_status.GetOpts{
			Locale:   locale,
			TenantID: tenantID,
			SoID:     soID,
		}
		order, err := service_order_status.Get(client, opts).Extract()
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

func getSingleDeviceByHostName(client *eclcloud.ServiceClient, hostName string) (single_devices.SingleDevice, error) {
	var sd = single_devices.SingleDevice{}

	listOpts := single_devices.ListOpts{
		TenantID: os.Getenv("OS_TENANT_ID"),
		Locale:   "en",
	}

	allPages, err := single_devices.List(client, listOpts).AllPages()
	if err != nil {
		return sd, fmt.Errorf("Unable to list after to create single device: %s", err)
	}
	var allDevices []single_devices.SingleDevice

	err = single_devices.ExtractSingleDevicesInto(allPages, &allDevices)
	if err != nil {
		return sd, fmt.Errorf("Unable to retrieve list of devices after create: %s", err)
	}

	var thisDevice single_devices.SingleDevice
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
	return thisDevice, nil
}

func resourceSecurityNetworkBasedSingleDeviceV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	tenantID := os.Getenv("OS_TENANT_ID")
	locale := d.Get("locale")

	device, err := getSingleDeviceByHostName(client, d.Id())
	if err != nil {
		return err
	}
	// device := dev.(single_devices.SingleDevice)
	// allPages, err := single_devices.List(client, listOpts).AllPages()
	// if err != nil {
	// 	return fmt.Errorf("Unable to list after to create single device: %s", err)
	// }
	// var allDevices []single_devices.SingleDevice

	// err = single_devices.ExtractSingleDevicesInto(allPages, &allDevices)
	// if err != nil {
	// 	return fmt.Errorf("Unable to retrieve list of devices after create: %s", err)
	// }

	// var thisDevice single_devices.SingleDevice
	// for _, device := range allDevices {
	// 	if device.Cell[2] == d.Id()
	// 	thisDevice = device
	// 	break
	// }

	// d.Set("so_kind", soKind)
	d.Set("tenant_id", tenantID)
	d.Set("locale", locale)

	// resultGtHost := []interface{}{}

	operatingMode := device.Cell[3]
	licenseKind := device.Cell[4]
	azGroup := device.Cell[6]

	d.Set("operating_mode", operatingMode)
	d.Set("license_kind", licenseKind)
	d.Set("az_group", azGroup)

	log.Printf("[DEBUG] SecurityNetworkBasedSingleDeviceV1Read Succeeded")

	return nil
}

func resourceSecurityNetworkBasedSingleDeviceV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts := single_devices.UpdateOpts{
		SOKind: "M",
		Locale: locale,
		GtHost: gtHostForSingleDeviceUpdateAsOpts(d),
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)
	order, err := single_devices.Update(client, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] Update request has successfully accepted with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securitySingleDeviceUpdatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}
	return resourceSecurityNetworkBasedSingleDeviceV1Read(d, meta)
}

func resourceSecurityNetworkBasedSingleDeviceV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	deleteOpts := single_devices.DeleteOpts{
		SOKind:   "D",
		TenantID: tenantID,
		GtHost:   gtHostForSingleDeviceDeleteAsOpts(d),
	}

	log.Printf("[DEBUG] Delete Options: %#v", deleteOpts)
	order, err := single_devices.Delete(client, deleteOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error deleting ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] Delete request has successfully accepted with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securitySingleDeviceDeletePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}

	// log.Printf("[DEBUG] Delete device is found as ID: %s", id)

	d.SetId("")

	return nil
}
