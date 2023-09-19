package ecl

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud/v4"
	security "github.com/nttcom/eclcloud/v4/ecl/security_order/v3/network_based_device_ha"
	"github.com/nttcom/eclcloud/v4/ecl/security_order/v3/service_order_status"

	ports "github.com/nttcom/eclcloud/v4/ecl/security_portal/v3/ha_ports"
)

const securityDeviceHAPollIntervalSec = 30
const securityDeviceHACreatePollInterval = securityDeviceHAPollIntervalSec * time.Second
const securityDeviceHAUpdatePollInterval = securityDeviceHAPollIntervalSec * time.Second
const securityDeviceHADeletePollInterval = securityDeviceHAPollIntervalSec * time.Second

func haLinkSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		ForceNew: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"network_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"subnet_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"host_1_ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"host_2_ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func haDeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"tenant_id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"locale": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "en",
			ValidateFunc: validation.StringInSlice([]string{
				"ja", "en",
			}, false),
		},

		"operating_mode": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"FW_HA", "UTM_HA",
			}, false),
		},

		"license_kind": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"02", "08",
			}, false),
		},

		"host_1_az_group": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},

		"host_2_az_group": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},

		"ha_link_1": haLinkSchema(),
		"ha_link_2": haLinkSchema(),

		"port": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 7,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enable": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					"host_1_ip_address": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"host_1_ip_address_prefix": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					"host_2_ip_address": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"host_2_ip_address_prefix": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
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
						Type:     schema.TypeString,
						Optional: true,
					},
					"comment": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					"enable_ping": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						ValidateFunc: validation.StringInSlice([]string{
							"true", "false",
						}, false),
					},

					"vrrp_grp_id": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					"vrrp_id": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					"vrrp_ip_address": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					"preempt": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						ValidateFunc: validation.StringInSlice([]string{
							"true", "false",
						}, false),
					},
				},
			},
		},
	}
}

func gtHostForHADeviceUpdateAsOpts(d *schema.ResourceData) [2]security.GtHostInUpdate {
	result := [2]security.GtHostInUpdate{}

	gtHost1 := security.GtHostInUpdate{}
	gtHost2 := security.GtHostInUpdate{}

	licenseKind := d.Get("license_kind").(string)
	operatingMode := d.Get("operating_mode").(string)

	gtHost1.LicenseKind = licenseKind
	gtHost1.OperatingMode = operatingMode

	gtHost2.LicenseKind = licenseKind
	gtHost2.OperatingMode = operatingMode

	hostNames := strings.Split(d.Id(), "/")

	gtHost1.HostName = hostNames[0]
	gtHost2.HostName = hostNames[1]

	result[0] = gtHost1
	result[1] = gtHost2

	return result
}

func resourceSecurityNetworkBasedDeviceHAV3UpdateOrderAPIPart(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityOrderV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security order client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	updateOpts := security.UpdateOpts{
		SOKind:   "MH",
		TenantID: tenantID,
		Locale:   locale,
		GtHost:   gtHostForHADeviceUpdateAsOpts(d),
	}

	log.Printf("[DEBUG] Update Options: %#v", updateOpts)

	order, err := security.Update(client, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating ECL security single device: %s", err)
	}

	log.Printf("[DEBUG] Update request has successfully accepted with order: %#v", order)

	log.Printf("[DEBUG] Start waiting for HA device order becomes COMPLETE ...")

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PROCESSING"},
		Target:       []string{"COMPLETE"},
		Refresh:      waitForHADeviceOrderComplete(client, order.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceHAUpdatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	log.Printf("[DEBUG] Finish waiting for HA device update order becomes COMPLETE")

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for HA device order status (%s) to become ready: %s",
			order.ID, err)
	}

	return nil
}

func gtHostForHADeviceCreateAsOpts(d *schema.ResourceData) [2]security.GtHostInCreate {
	result := [2]security.GtHostInCreate{}

	gtHost1 := security.GtHostInCreate{}
	gtHost2 := security.GtHostInCreate{}

	// License Kind
	gtHost1.LicenseKind = d.Get("license_kind").(string)
	gtHost2.LicenseKind = d.Get("license_kind").(string)

	// Operating Mode
	gtHost1.OperatingMode = d.Get("operating_mode").(string)
	gtHost2.OperatingMode = d.Get("operating_mode").(string)

	// AZGroup
	gtHost1.AZGroup = d.Get("host_1_az_group").(string)
	gtHost2.AZGroup = d.Get("host_2_az_group").(string)

	// HALink NetworkID/SubnetID
	haLink1List := d.Get("ha_link_1").([]interface{})
	haLink2List := d.Get("ha_link_2").([]interface{})

	haLink1 := haLink1List[0].(map[string]interface{})
	haLink2 := haLink2List[0].(map[string]interface{})

	haLink1NetworkID := haLink1["network_id"].(string)
	haLink1SubnetID := haLink1["subnet_id"].(string)

	haLink2NetworkID := haLink2["network_id"].(string)
	haLink2SubnetID := haLink2["subnet_id"].(string)

	// GtHost-1 Network/Subnet
	gtHost1.HALink1NetworkID = haLink1NetworkID
	gtHost1.HALink1SubnetID = haLink1SubnetID
	gtHost1.HALink2NetworkID = haLink2NetworkID
	gtHost1.HALink2SubnetID = haLink2SubnetID

	// GtHost-1 Network/Subnet
	gtHost2.HALink1NetworkID = haLink1NetworkID
	gtHost2.HALink1SubnetID = haLink1SubnetID
	gtHost2.HALink2NetworkID = haLink2NetworkID
	gtHost2.HALink2SubnetID = haLink2SubnetID

	// HALink IP Address
	host1HALink1IPAddress := haLink1["host_1_ip_address"].(string)
	host1HALink2IPAddress := haLink2["host_1_ip_address"].(string)

	host2HALink1IPAddress := haLink1["host_2_ip_address"].(string)
	host2HALink2IPAddress := haLink2["host_2_ip_address"].(string)

	gtHost1.HALink1IPAddress = host1HALink1IPAddress
	gtHost1.HALink2IPAddress = host1HALink2IPAddress

	gtHost2.HALink1IPAddress = host2HALink1IPAddress
	gtHost2.HALink2IPAddress = host2HALink2IPAddress

	result[0] = gtHost1
	result[1] = gtHost2

	return result
}

func getHADeviceByHostName(client *eclcloud.ServiceClient, hostName string, tenantID string) (security.HADevice, error) {
	log.Printf("[DEBUG] Start getting HA Device by HostName %s ...", hostName)
	var hd = security.HADevice{}

	listOpts := security.ListOpts{
		TenantID: tenantID,
		Locale:   "en",
	}

	allPages, err := security.List(client, listOpts).AllPages()
	log.Printf("[DEBUG] Got HA Device pages as: %#v", allPages)
	if err != nil {
		return hd, fmt.Errorf("Unable to list HA device to get hostname from result: %s", err)
	}
	var allDevices []security.HADevice

	err = security.ExtractHADevicesInto(allPages, &allDevices)
	if err != nil {
		return hd, fmt.Errorf("Unable to extract result of HA device list api: %s", err)
	}
	log.Printf("[DEBUG] Extracted HA Devices as: %#v", allDevices)

	var thisDevice security.HADevice
	var found bool
	for _, device := range allDevices {
		if device.Cell[3] == hostName {
			thisDevice = device
			found = true
			break
		}
	}
	if !found {
		return hd, fmt.Errorf("[DEBUG] Specified HA device %s not found", hostName)
	}
	log.Printf("[DEBUG] Host has found as: %#v", thisDevice)
	return thisDevice, nil
}

func getIDFromHostNames(ids []string) string {

	rep := regexp.MustCompile(`^[A-Za-z]+`)
	id1s := rep.ReplaceAllString(ids[0], "")
	id2s := rep.ReplaceAllString(ids[1], "")

	id1i, _ := strconv.Atoi(id1s)
	id2i, _ := strconv.Atoi(id2s)

	if id1i < id2i {
		return fmt.Sprintf("%s/%s", ids[0], ids[1])
	}

	return fmt.Sprintf("%s/%s", ids[1], ids[0])
}

func getNewlyCreatedHADeviceID(before, after []security.HADevice) []string {
	result := []string{}

	for _, af := range after {
		hostNameAfter := af.Cell[3]
		match := false
		for _, bf := range before {
			hostNameBefore := bf.Cell[3]
			if hostNameAfter == hostNameBefore {
				match = true
			}
		}
		if !match {
			result = append(result, hostNameAfter)
		}
	}

	return result
}

func waitForHADeviceOrderComplete(client *eclcloud.ServiceClient, soID, tenantID, locale string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		opts := service_order_status.GetOpts{
			Locale:   locale,
			TenantID: tenantID,
			SoID:     soID,
		}
		order, err := service_order_status.Get(client, "UTM", opts).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Security Service Order Status: %#v", order)

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

func gtHostForHADeviceDeleteAsOpts(d *schema.ResourceData) [2]security.GtHostInDelete {
	result := [2]security.GtHostInDelete{}

	gtHost1 := security.GtHostInDelete{}
	gtHost2 := security.GtHostInDelete{}

	hostNames := strings.Split(d.Id(), "/")

	gtHost1.HostName = hostNames[0]
	gtHost2.HostName = hostNames[1]

	result[0] = gtHost1
	result[1] = gtHost2

	return result
}

func resourceSecurityNetworkBasedHADevicePortsForUpdate(d *schema.ResourceData) (ports.UpdateOpts, error) {
	resultPorts := []ports.SinglePort{}

	ifaces := d.Get("port").([]interface{})
	log.Printf("[DEBUG] Retrieved port information for update: %#v", ifaces)
	for _, iface := range ifaces {
		p := ports.SinglePort{}

		if _, ok := iface.(map[string]interface{}); ok {
			thisInterface := iface.(map[string]interface{})

			if thisInterface["enable"].(string) == "true" {
				p.EnablePort = "true"

				// host_1_ip_address, prefix and
				// host_2_ip_address, prefix are the real IP address
				// for both interface on host1 and host2.
				// You need to add prefix for this configuration.
				ipAddress1 := thisInterface["host_1_ip_address"].(string)
				prefix1 := thisInterface["host_1_ip_address_prefix"].(int)

				ipAddress2 := thisInterface["host_2_ip_address"].(string)
				prefix2 := thisInterface["host_2_ip_address_prefix"].(int)

				// In HA device case, you need to specify port.IPAddress
				// as array with length = 2.
				p.IPAddress = []string{
					fmt.Sprintf("%s/%d", ipAddress1, prefix1),
					fmt.Sprintf("%s/%d", ipAddress2, prefix2),
				}

				p.NetworkID = thisInterface["network_id"].(string)
				p.SubnetID = thisInterface["subnet_id"].(string)
				p.MTU = thisInterface["mtu"].(string)
				p.Comment = thisInterface["comment"].(string)

				p.EnablePing = thisInterface["enable_ping"].(string)
				p.VRRPGroupID = thisInterface["vrrp_grp_id"].(string)

				p.VRRPID = thisInterface["vrrp_id"].(string)

				// VRRPIP is "Virtual" IP Address in VRRP configuration
				// You do not need to add prefix for this VRRPIP
				p.VRRPIPAddress = thisInterface["vrrp_ip_address"].(string)
				p.Preempt = thisInterface["preempt"].(string)

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

func resourceSecurityNetworkBasedDeviceHAV3UpdatePortalAPIPart(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.securityPortalV3Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL security portal client: %s", err)
	}

	tenantID := d.Get("tenant_id").(string)
	locale := d.Get("locale").(string)

	hostNames := strings.Split(d.Id(), "/")
	host1 := hostNames[0]

	updateOpts, err := resourceSecurityNetworkBasedHADevicePortsForUpdate(d)
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
		host1,
		updateOpts,
		updateQueryOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error updating ECL security single device port: %s", err)
	}

	log.Printf("[DEBUG] Update request for %s has successfully accepted with process: %#v", d.Id(), process)

	log.Printf("[DEBUG] Start waiting for HA device process for %s becomes ENDED ...", d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"RUNNING"},
		Target:       []string{"ENDED"},
		Refresh:      waitForSingleDeviceProcessComplete(client, process.ID, tenantID, locale),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: securityDeviceHAUpdatePollInterval,
		MinTimeout:   30 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf(
			"Error waiting for HA device port management status (%s) to become ready: %s",
			process.ID, err)
	}

	log.Printf("[DEBUG] Finish waiting for HA device portal api order becomes COMPLETE")

	return nil
}
