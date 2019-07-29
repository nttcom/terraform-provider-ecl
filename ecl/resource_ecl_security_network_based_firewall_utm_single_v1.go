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
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud"

	"github.com/nttcom/eclcloud/ecl/security_portal/v1/ports"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/processes"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_firewall_utm_single"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/device_interfaces"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/devices"

	"github.com/nttcom/eclcloud/ecl/security_order/v1/service_order_status"
)

const securityFirewallUTMSingleFirewallUTMPollIntervalSec = 1
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
				ForceNew: true,
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

			"port": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 7,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subnet_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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
	log.Printf("[DEBUG] Finish waiting for firewall/utm create order becomes COMPLETE")

	allPagesAfter, err := security.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list after creating single firewall/utm: %s", err)
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

func getUUIDFromServerHostName(client *eclcloud.ServiceClient, hostName string) (string, error) {

	listOpts := devices.ListOpts{
		TenantID:  os.Getenv("OS_TENANT_ID"),
		UserToken: client.TokenID,
	}

	allPages, err := devices.List(client, listOpts).AllPages()
	if err != nil {
		return "", fmt.Errorf("Unable to list firewall/utm to get device UUID: %s", err)
	}
	var allDevices []devices.Device

	err = devices.ExtractDevicesInto(allPages, &allDevices)
	if err != nil {
		return "", fmt.Errorf("Unable to extract result of list firewall/utm from portal api: %s", err)
	}

	for _, device := range allDevices {
		if device.MSADeviceID == hostName {
			log.Printf("[DEBUG] Host UUID looking result: Host %s has UUID %s", hostName, device.OSServerID)
			return device.OSServerID, nil
		}
	}

	return "", fmt.Errorf("Unable to find corresponding server of %s", hostName)
}

func resourceSecurityNetworkBasedFirewallUTMSingleV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	// Main Part
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

	allDevicePages, err := device_interfaces.List(client, hostUUID, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list interfaces: %s", err)
	}

	allDevices, err := device_interfaces.ExtractDeviceInterfaces(allDevicePages)
	if err != nil {
		return fmt.Errorf("Unable to extract device interfaces: %s", err)
	}

	deviceInterfaces := [7]map[string]string{}
	for _, d := range allDevices {
		thisDeviceInterface := map[string]string{}

		index, err := strconv.Atoi(strings.Replace(d.OSPortID, "port", "", 1))
		if err != nil {
			return fmt.Errorf("Error parsing device interface port number: %s", err)
		}
		index -= 4
		if index < 0 {
			return fmt.Errorf("Wrong index number is returned from device interface list API. %s", err)
		}

		thisDeviceInterface["enable"] = "true"
		thisDeviceInterface["ip_address"] = d.OSNetworkID
		thisDeviceInterface["network_id"] = d.OSNetworkID
		thisDeviceInterface["subnet_id"] = d.OSSubnetID

		deviceInterfaces[index] = thisDeviceInterface
	}

	d.Set("port", deviceInterfaces)
	return nil
}

func resourceSecurityNetworkBasedFirewallUTMSingleV1UpdateOrderAPIPart(d *schema.ResourceData, meta interface{}) error {
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

	log.Printf("[DEBUG] Finish waiting for firewall/utm update order becomes COMPLETE")

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single firewall/utm order status (%s) to become ready: %s",
			order.ID, err)
	}

	return nil
}

func resourceSecurityNetworkBasedFirewallUTMPortsForUpdate(d *schema.ResourceData) (ports.UpdateOpts, error) {
	resultPorts := [7]ports.SinglePort{}

	ifaces := d.Get("port").([]interface{})
	log.Printf("[DEBUG] Retrieved port information for update: %#v", ifaces)
	for index, iface := range ifaces {
		p := ports.SinglePort{}

		if _, ok := iface.(map[string]interface{}); ok {
			thisInterface := iface.(map[string]interface{})

			if thisInterface["enable"].(string) == "true" {
				p.EnablePort = "true"
				p.IPAddress = thisInterface["ip_address"].(string)
				p.NetworkID = thisInterface["network_id"].(string)
				p.SubnetID = thisInterface["subnet_id"].(string)
				p.Comment = thisInterface["comment"].(string)
			}
		} else {
			p.EnablePort = "false"
		}

		resultPorts[index] = p
	}

	log.Printf("[DEBUG] Port update parameters: %#v", resultPorts)
	result := ports.UpdateOpts{}
	result.Port = resultPorts
	return result, nil
}

func resourceSecurityNetworkBasedFirewallUTMSingleV1UpdatePortalAPIPart(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityPortalV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security portal client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts, err := resourceSecurityNetworkBasedFirewallUTMPortsForUpdate(d)
	if err != nil {
		return fmt.Errorf("Error getting port option in update: %s", err)
	}
	updateQueryOpts := ports.UpdateQueryOpts{
		TenantID:  tenantID,
		UserToken: client.TokenID,
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)
	log.Printf("[DEBUG] Update Query Options: %#v", updateQueryOpts)
	process, err := ports.Update(
		client,
		"utm",
		d.Id(),
		updateOpts,
		updateQueryOpts).Extract()
	log.Printf("[MYDEBUG] process: %#v", process)
	log.Printf("[MYDEBUG] error: %#v", err)

	if err != nil {
		return fmt.Errorf("Error updating ECL security single firewall/utm port: %s", err)
	}

	log.Printf("[DEBUG] Update request has successfully accepted with process: %#v", process)

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"RUNNING"},
		Target:       []string{"ENDED"},
		Refresh:      waitForSingleFirewallUTMProcessComplete(client, process.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityFirewallUTMSingleUpdatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for single firewall/utm port management status (%s) to become ready: %s",
			process.ID, err)
	}

	log.Printf("[DEBUG] Finish waiting for firewall/utm portal api order becomes COMPLETE")

	return nil
}

func resourceSecurityNetworkBasedFirewallUTMSingleV1Update(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("locale") || d.HasChange("operating_mode") || d.HasChange("license_kind") {
		log.Printf("[DEBUG] Start changing firwall/utm by order api.")
		resourceSecurityNetworkBasedFirewallUTMSingleV1UpdateOrderAPIPart(d, meta)
	}

	if d.HasChange("port") {
		log.Printf("[DEBUG] Start changing firwall/utm by portal api.")
		resourceSecurityNetworkBasedFirewallUTMSingleV1UpdatePortalAPIPart(d, meta)
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

		log.Printf("[DEBUG] Start waiting for single firewall/utm order becomes COMPLETE ...")

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
	log.Printf("[DEBUG] Start getting hostname ...")
	var sd = security.SingleFirewallUTM{}

	listOpts := security.ListOpts{
		TenantID: os.Getenv("OS_TENANT_ID"),
		Locale:   "en",
	}

	allPages, err := security.List(client, listOpts).AllPages()
	if err != nil {
		return sd, fmt.Errorf("Unable to list firewall/utm to get hostname from result: %s", err)
	}
	var allDevices []security.SingleFirewallUTM

	err = security.ExtractSingleFirewallUTMsInto(allPages, &allDevices)
	if err != nil {
		return sd, fmt.Errorf("Unable to extract result of single firewall/utm list api: %s", err)
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

func waitForSingleFirewallUTMProcessComplete(client *eclcloud.ServiceClient, processID, tenantID, locale string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		log.Printf("[DEBUG] Start waiting for single firewall/utm process becomes ENDED ...")
		opts := processes.GetOpts{
			TenantID:  tenantID,
			UserToken: client.TokenID,
		}
		process, err := processes.Get(client, processID, opts).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Service Process Status: %+v", process)

		// if process.Status.Status == "ENDED" {
		// 	return process, "COMPLETE", nil
		// }

		return process, process.Status.Status, nil
	}
}
