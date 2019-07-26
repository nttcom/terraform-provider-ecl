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
)

const securityFirewallUTMSingleFirewallUTMPollIntervalSec = 30
const securityFirewallUTMSingleCreatePollInterval = securityFirewallUTMSingleFirewallUTMPollIntervalSec * time.Second
const securityFirewallUTMSingleUpdatePollInterval = securityFirewallUTMSingleFirewallUTMPollIntervalSec * time.Second
const securityFirewallUTMSingleDeletePollInterval = securityFirewallUTMSingleFirewallUTMPollIntervalSec * time.Second

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
					"FW", "UTM",
				}, false),
			},

			"license_kind": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"02", "08",
				}, false),
			},

			"az_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
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
		return fmt.Errorf("Unable to list of devices, before creating new single firewall/utm: %s", err)
	}
	var allDevicesBefore []security.SingleFirewallUTM

	err = security.ExtractSingleFirewallUTMsInto(allPagesBefore, &allDevicesBefore)

	if err != nil {
		return fmt.Errorf("Unable to retrieve device list before create: %s", err)
	}
	log.Printf("[DEBUG] allSingleFirewallUTMs before creation: %#v", allDevicesBefore)
	createOpts := security.CreateOpts{
		SOKind:   "A",
		TenantID: tenantID,
		Locale:   locale,
		GtHost:   gtHostForFirewallUTMSingleCreateAsOpts(d),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	order, err := security.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL security single firewall/utm: %s", err)
	}

	log.Printf("[DEBUG] Firewall/UTM has successfully created with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleFirewallUTMOrderComplete(client, order.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityFirewallUTMSingleCreatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single firewall/utm order status (%s) to become ready: %s",
			order.ID, err)
	}

	allPagesAfter, err := security.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list after to create single firewall/utm: %s", err)
	}
	var allDevicesAfter []security.SingleFirewallUTM

	err = security.ExtractSingleFirewallUTMsInto(allPagesAfter, &allDevicesAfter)
	if err != nil {
		return fmt.Errorf("Unable to retrieve list of firewall/utm after create: %s", err)
	}
	log.Printf("[DEBUG] allSingleFirewallUTMs after creation: %#v", allDevicesAfter)

	if len(allDevicesBefore) == len(allDevicesAfter) {
		return fmt.Errorf("Unable to find newly created firewall/utm")
	}

	id := getNewlyCreatedDeviceID(allDevicesBefore, allDevicesAfter)
	if id == "" {
		return fmt.Errorf("Unable to find newly created firewall/utm after hostname matching")
	}

	log.Printf("[DEBUG] Newly created firewall/utm is found as ID: %s", id)

	d.SetId(id)

	return resourceSecurityNetworkBasedFirewallUTMSingleV1Read(d, meta)
}

func getNewlyCreatedDeviceID(before, after []security.SingleFirewallUTM) string {
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

func resourceSecurityNetworkBasedFirewallUTMSingleV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := os.Getenv("OS_TENANT_ID")
	locale := d.Get("locale")

	device, err := getSingleFirewallUTMByHostName(client, d.Id())
	if err != nil {
		return err
	}

	d.Set("tenant_id", tenantID)
	d.Set("locale", locale)

	operatingMode := device.Cell[3]
	licenseKind := device.Cell[4]
	azGroup := device.Cell[6]

	d.Set("operating_mode", operatingMode)
	d.Set("license_kind", licenseKind)
	d.Set("az_group", azGroup)

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
		return fmt.Errorf("Error updating ECL security single firewall/utm: %s", err)
	}

	log.Printf("[DEBUG] Update request has successfully accepted with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleFirewallUTMOrderComplete(client, order.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityFirewallUTMSingleUpdatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single firewall/utm order status (%s) to become ready: %s",
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
		return fmt.Errorf("Error deleting ECL security single firewall/utm: %s", err)
	}

	log.Printf("[DEBUG] Delete request has successfully accepted with order: %#v", order)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForSingleFirewallUTMOrderComplete(client, order.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityFirewallUTMSingleDeletePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single firewall/utm order status (%s) to become ready: %s",
			order.ID, err)
	}

	d.SetId("")

	return nil
}

func waitForSingleFirewallUTMOrderComplete(client *eclcloud.ServiceClient, soID, tenantID, locale string) resource.StateRefreshFunc {
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

func getSingleFirewallUTMByHostName(client *eclcloud.ServiceClient, hostName string) (security.SingleFirewallUTM, error) {
	var sd = security.SingleFirewallUTM{}

	listOpts := security.ListOpts{
		TenantID: os.Getenv("OS_TENANT_ID"),
		Locale:   "en",
	}

	allPages, err := security.List(client, listOpts).AllPages()
	if err != nil {
		return sd, fmt.Errorf("Unable to list after to create single device: %s", err)
	}
	var allDevices []security.SingleFirewallUTM

	err = security.ExtractSingleFirewallUTMsInto(allPages, &allDevices)
	if err != nil {
		return sd, fmt.Errorf("Unable to retrieve list of devices after create: %s", err)
	}

	var thisDevice security.SingleFirewallUTM
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
