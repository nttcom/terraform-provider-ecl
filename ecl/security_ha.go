package ecl

import (
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud"
	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_ha"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/service_order_status"
)

const securityDeviceHAPollIntervalSec = 30
const securityDeviceHACreatePollInterval = securityDeviceHAPollIntervalSec * time.Second
const securityDeviceHAUpdatePollInterval = securityDeviceHAPollIntervalSec * time.Second
const securityDeviceHADeletePollInterval = securityDeviceHAPollIntervalSec * time.Second

func haLinkSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
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
					Required: schema.TypeInt,
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
			Required: true,
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
	haLink1 := d.Get("ha_link_1").(map[string]interface{})
	haLink2 := d.Get("ha_link_2").(map[string]interface{})

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

func waitForHADeviceOrderComplete(client *eclcloud.ServiceClient, soID, tenantID, locale, deviceType string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		opts := service_order_status.GetOpts{
			Locale:   locale,
			TenantID: tenantID,
			SoID:     soID,
		}
		order, err := service_order_status.Get(client, "HA", opts).Extract()
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
