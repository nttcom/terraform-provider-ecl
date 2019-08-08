package ecl

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/nttcom/eclcloud"
	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_single"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/service_order_status"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/ports"
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/processes"
)

const securityDeviceSinglePollIntervalSec = 30
const securityDeviceSingleCreatePollInterval = securityDeviceSinglePollIntervalSec * time.Second
const securityDeviceSingleUpdatePollInterval = securityDeviceSinglePollIntervalSec * time.Second
const securityDeviceSingleDeletePollInterval = securityDeviceSinglePollIntervalSec * time.Second

func singleDeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

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
			ForceNew: true,
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
						Required: true,
					},
					"ip_address": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"ip_address_prefix": &schema.Schema{
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
				},
			},
		},
	}
}

func singleWAFSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{

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
			Optional: true,
			Default:  "WAF",
			ValidateFunc: validation.StringInSlice([]string{
				"WAF",
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
			ForceNew: true,
		},

		"port": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enable": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"ip_address": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"ip_address_prefix": &schema.Schema{
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
				},
			},
		},
	}
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

func getTypeOfSingleDevice(d *schema.ResourceData) string {
	operatingMode := d.Get("operating_mode").(string)
	switch operatingMode {
	case "WAF":
		return "WAF"
	default:
		return "UTM"
	}
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
