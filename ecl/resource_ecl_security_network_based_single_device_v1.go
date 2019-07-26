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
	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_firewall_utm_single"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/service_order_status"
	// "github.com/nttcom/eclcloud/ecl/sss/v1/users"
)

const securityFirewallUTMSingleDevicePollIntervalSec = 1
const securityFirewallUTMSingleCreatePollInterval = securityFirewallUTMSingleDevicePollIntervalSec * time.Second
const securityFirewallUTMSingleUpdatePollInterval = securityFirewallUTMSingleDevicePollIntervalSec * time.Second
const securityFirewallUTMSingleDeletePollInterval = securityFirewallUTMSingleDevicePollIntervalSec * time.Second

// func SecurityNetworkBasedFirewallUTMSingleV1() *schema.Resource {
func resourceSecurityNetworkBasedFirewallUTMSingleV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityNetworkBasedFirewallUTMSingleV1Create,
		Read:   resourceSecurityNetworkBasedFirewallUTMSingleV1Read,
		Update: resourceSecurityNetworkBasedFirewallUTMSingleV1Update,
		Delete: resourceSecurityNetworkBasedFirewallUTMSingleV1Delete,

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

func gtHostForFirewallUTMSingleCreateAsOpts(d *schema.ResourceData) [1]security.GtHostInCreate {
	result := [1]security.GtHostInCreate{}

	gtHost := security.GtHostInCreate{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.AZGroup = d.Get("az_group").(string)

	result[0] = gtHost

	return result
}

func gtHostForFirewallUTMSingleUpdateAsOpts(d *schema.ResourceData) [1]security.GtHostInUpdate {
	result := [1]security.GtHostInUpdate{}

	gtHost := security.GtHostInUpdate{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.HostName = d.Id()

	result[0] = gtHost

	return result
}

func gtHostForFirewallUTMSingleDeleteAsOpts(d *schema.ResourceData) [1]security.GtHostInDelete {
	result := [1]security.GtHostInDelete{}

	gtHost := security.GtHostInDelete{}

	gtHost.HostName = d.Id()

	result[0] = gtHost

	return result
}

func resourceSecurityNetworkBasedFirewallUTMSingleV1Create(d *schema.ResourceData, meta interface{}) error {
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
		return fmt.Errorf("Unable to list of devices, before creating new single device: %s", err)
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
		GtHost:   gtHostForFirewallUTMSingleCreateAsOpts(d),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	order, err := security.Create(client, createOpts).Extract()
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
		PollInterval: securityFirewallUTMSingleCreatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}

	allPagesAfter, err := security.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list after to create single device: %s", err)
	}
	var allDevicesAfter []security.SingleDevice

	err = security.ExtractSingleDevicesInto(allPagesAfter, &allDevicesAfter)
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

	return resourceSecurityNetworkBasedFirewallUTMSingleV1Read(d, meta)
}

func getNewlyCreatedDeviceID(before, after []security.SingleDevice) string {
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

func getSingleDeviceByHostName(client *eclcloud.ServiceClient, hostName string) (security.SingleDevice, error) {
	var sd = security.SingleDevice{}

	listOpts := security.ListOpts{
		TenantID: os.Getenv("OS_TENANT_ID"),
		Locale:   "en",
	}

	allPages, err := security.List(client, listOpts).AllPages()
	if err != nil {
		return sd, fmt.Errorf("Unable to list after to create single device: %s", err)
	}
	var allDevices []security.SingleDevice

	err = security.ExtractSingleDevicesInto(allPages, &allDevices)
	if err != nil {
		return sd, fmt.Errorf("Unable to retrieve list of devices after create: %s", err)
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
	return thisDevice, nil
}

func resourceSecurityNetworkBasedFirewallUTMSingleV1Read(d *schema.ResourceData, meta interface{}) error {
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
	// device := dev.(security.SingleDevice)
	// allPages, err := security.List(client, listOpts).AllPages()
	// if err != nil {
	// 	return fmt.Errorf("Unable to list after to create single device: %s", err)
	// }
	// var allDevices []security.SingleDevice

	// err = security.ExtractSingleDevicesInto(allPages, &allDevices)
	// if err != nil {
	// 	return fmt.Errorf("Unable to retrieve list of devices after create: %s", err)
	// }

	// var thisDevice security.SingleDevice
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

	log.Printf("[DEBUG] SecurityNetworkBasedFirewallUTMSingleV1Read Succeeded")

	return nil
}

func resourceSecurityNetworkBasedFirewallUTMSingleV1Update(d *schema.ResourceData, meta interface{}) error {
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
		GtHost:   gtHostForFirewallUTMSingleUpdateAsOpts(d),
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)
	order, err := security.Update(client, updateOpts).Extract()
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
		PollInterval: securityFirewallUTMSingleUpdatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.ID, err)
	}
	return resourceSecurityNetworkBasedFirewallUTMSingleV1Read(d, meta)
}

func resourceSecurityNetworkBasedFirewallUTMSingleV1Delete(d *schema.ResourceData, meta interface{}) error {
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
		GtHost:   gtHostForFirewallUTMSingleDeleteAsOpts(d),
	}

	log.Printf("[DEBUG] Delete Options: %#v", deleteOpts)
	order, err := security.Delete(client, deleteOpts).Extract()
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
		PollInterval: securityFirewallUTMSingleDeletePollInterval,
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
