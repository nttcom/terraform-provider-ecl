package ecl

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v3"
	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v4/network_based_device_single"
	"github.com/nttcom/eclcloud/v3/ecl/security_order/v4/service_order_status"
	"github.com/nttcom/eclcloud/v3/ecl/security_portal/v4/devices"
	"github.com/nttcom/eclcloud/v3/ecl/security_portal/v4/ports"
	"github.com/nttcom/eclcloud/v3/ecl/security_portal/v4/processes"
)

const securityDeviceSinglePollIntervalSecV4 = 30
const securityDeviceSingleCreatePollIntervalV4 = securityDeviceSinglePollIntervalSecV4 * time.Second
const securityDeviceSingleUpdatePollIntervalV4 = securityDeviceSinglePollIntervalSecV4 * time.Second
const securityDeviceSingleDeletePollIntervalV4 = securityDeviceSinglePollIntervalSecV4 * time.Second

func getUUIDFromServerHostNameV4(client *eclcloud.ServiceClient, hostName string, tenantID string) (string, error) {

	listOpts := devices.ListOpts{
		TenantID:  tenantID,
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

func gtHostForSingleDeviceCreateAsOptsV4(d *schema.ResourceData) [1]security.GtHostInCreate {
	result := [1]security.GtHostInCreate{}

	gtHost := security.GtHostInCreate{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.AZGroup = d.Get("az_group").(string)

	result[0] = gtHost

	return result
}

func gtHostForSingleDeviceUpdateAsOptsV4(d *schema.ResourceData) [1]security.GtHostInUpdate {
	result := [1]security.GtHostInUpdate{}

	gtHost := security.GtHostInUpdate{}

	gtHost.LicenseKind = d.Get("license_kind").(string)
	gtHost.OperatingMode = d.Get("operating_mode").(string)
	gtHost.HostName = d.Id()

	result[0] = gtHost

	return result
}

func gtHostForSingleDeviceDeleteAsOptsV4(d *schema.ResourceData) [1]security.GtHostInDelete {
	result := [1]security.GtHostInDelete{}

	gtHost := security.GtHostInDelete{}

	gtHost.HostName = d.Id()

	result[0] = gtHost

	return result
}

func getSingleDeviceByHostNameV4(client *eclcloud.ServiceClient, deviceType, hostName string, tenantID string) (security.SingleDevice, error) {
	log.Printf("[DEBUG] Start getting %s by HostName %s ...", deviceType, hostName)
	var sd = security.SingleDevice{}

	listOpts := security.ListOpts{
		TenantID: tenantID,
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

func resourceSecurityNetworkBasedSingleDevicePortsForUpdateV4(d *schema.ResourceData) (ports.UpdateOpts, error) {
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

func resourceSecurityNetworkBasedDeviceSingleV4UpdatePortalAPIPart(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityPortalV4Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security portal client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts, err := resourceSecurityNetworkBasedSingleDevicePortsForUpdateV4(d)
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
		Refresh:      waitForSingleDeviceProcessCompleteV4(client, process.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceSingleUpdatePollIntervalV4,
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

func resourceSecurityNetworkBasedDeviceSingleV4UpdateOrderAPIPart(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV4Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts := security.UpdateOpts{
		SOKind:   "M",
		TenantID: tenantID,
		Locale:   locale,
		GtHost:   gtHostForSingleDeviceUpdateAsOptsV4(d),
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
		Refresh:      waitForSingleDeviceOrderCompleteV4(client, order.ID, tenantID, locale, deviceType),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceSingleUpdatePollIntervalV4,
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

func getNewlyCreatedDeviceIDV4(beforeTemp, afterTemp []security.SingleDevice, deviceType string) string {
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

func waitForSingleDeviceOrderCompleteV4(client *eclcloud.ServiceClient, soID, tenantID, locale, deviceType string) resource.StateRefreshFunc {
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

		r := regexp.MustCompile(`^FOV-E`)
		if r.MatchString(order.Code) {
			return order, "ERROR", fmt.Errorf("Status becomes error %s: %s", order.Code, order.Message)
		}

		if order.ProgressRate == 100 {
			return order, "COMPLETE", nil
		}

		return order, "PROCESSING", nil
	}
}

func waitForSingleDeviceProcessCompleteV4(client *eclcloud.ServiceClient, processID, tenantID, locale string) resource.StateRefreshFunc {
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
