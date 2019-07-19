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
	"github.com/nttcom/eclcloud/ecl/sss/v1/users"
	// "github.com/nttcom/eclcloud/ecl/sss/v1/users"
)

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

func gtHostForCreateAsOpts(d *schema.ResourceData) []single_devices.GtHost {
	result := []single_devices.GtHost{}

	gtHost := single_devices.GtHost{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.AZGroup = d.Get("az_group").(string)

	result = append(result, gtHost)

	return result
}

func resourceSecurityNetworkBasedSingleDeviceV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
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
		GtHost:   gtHostForCreateAsOpts(d),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	order, err := single_devices.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] User has successfully created.")

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleDeviceOrderComplete(client, order.SoID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: 1 * time.Minute,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single device order status (%s) to become ready: %s",
			order.SoID, err)
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
		}
		order, err := service_order_status.Get(client, soID, opts).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Service Order Status: %+v", order)

		if order.ProgressRate == "100" {
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
	// config := meta.(*Config)
	// client, err := config.sssV1Client(GetRegion(d, config))
	// if err != nil {
	// 	return fmt.Errorf("Error creating ECL sss client: %s", err)
	// }

	// var hasChange bool
	// var updateOpts users.UpdateOpts

	// if d.HasChange("login_id") {
	// 	hasChange = true
	// 	loginID := d.Get("login_id").(string)
	// 	updateOpts.LoginID = &loginID
	// }

	// if d.HasChange("mail_address") {
	// 	hasChange = true
	// 	mailAddress := d.Get("mail_address").(string)
	// 	updateOpts.MailAddress = &mailAddress
	// }

	// if d.HasChange("password") {
	// 	hasChange = true
	// 	newPassword := d.Get("password").(string)
	// 	updateOpts.NewPassword = &newPassword
	// }

	// if hasChange {
	// 	r := users.Update(client, d.Id(), updateOpts)
	// 	if r.Err != nil {
	// 		return fmt.Errorf("Error updating ECL user: %s", r.Err)
	// 	}
	// 	log.Printf("[DEBUG] User has successfully updated.")
	// }

	return resourceSecurityNetworkBasedSingleDeviceV1Read(d, meta)
}

func resourceSecurityNetworkBasedSingleDeviceV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	err = users.Delete(client, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL user: %s", err)
	}

	log.Printf("[DEBUG] User has successfully deleted.")
	return nil
}
