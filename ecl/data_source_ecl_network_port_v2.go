package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/ecl/network/v2/ports"
)

func dataSourceNetworkPortV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkPortV2Read,

		Schema: map[string]*schema.Schema{
			"admin_state_up": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"all_fixed_ips": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"allowed_address_pairs": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Set:      allowedAddressPairsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"device_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"device_owner": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fixed_ip": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mac_address": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"managed_by_service": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"segmentation_id": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"segmentation_type": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceNetworkPortV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	listOpts := ports.ListOpts{}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("device_id"); ok {
		listOpts.DeviceID = v.(string)
	}

	if v, ok := d.GetOk("device_owner"); ok {
		listOpts.DeviceOwner = v.(string)
	}

	if v, ok := d.GetOk("port_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("mac_address"); ok {
		listOpts.MACAddress = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("network_id"); ok {
		listOpts.NetworkID = v.(string)
	}

	if v, ok := d.GetOk("segmentation_id"); ok {
		listOpts.SegmentationID = v.(int)
	}

	if v, ok := d.GetOk("segmentation_type"); ok {
		listOpts.SegmentationType = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	allPages, err := ports.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to list ecl_network_ports_v2: %s", err)
	}

	var allPorts []ports.Port

	err = ports.ExtractPortsInto(allPages, &allPorts)
	if err != nil {
		return fmt.Errorf("Unable to retrieve ecl_network_ports_v2: %s", err)
	}

	if len(allPorts) == 0 {
		return fmt.Errorf("No ecl_network_port_v2 found")
	}

	var portsList []ports.Port

	// Filter returned Fixed IPs by a "fixed_ip".
	if v, ok := d.GetOk("fixed_ip"); ok {
		for _, p := range allPorts {
			for _, ipObject := range p.FixedIPs {
				if v.(string) == ipObject.IPAddress {
					portsList = append(portsList, p)
				}
			}
		}
		if len(portsList) == 0 {
			log.Printf("No ecl_network_port_v2 found after the 'fixed_ip' filter")
			return fmt.Errorf("No ecl_network_port_v2 found")
		}
	} else {
		portsList = allPorts
	}

	if len(portsList) > 1 {
		return fmt.Errorf("More than one ecl_network_port_v2 found (%d)", len(portsList))
	}

	port := portsList[0]

	log.Printf("[DEBUG] Retrieved ecl_network_port_v2 %s: %+v", port.ID, port)
	d.SetId(port.ID)

	d.Set("admin_state_up", port.AdminStateUp)
	d.Set("allowed_address_pairs", flattenNetworkPortAllowedAddressPairsV2(port.MACAddress, port.AllowedAddressPairs))
	d.Set("description", port.Description)
	d.Set("device_id", port.DeviceID)
	d.Set("device_owner", port.DeviceOwner)
	d.Set("all_fixed_ips", expandNetworkPortFixedIPToStringSlice(port.FixedIPs))
	d.Set("fixed_ips", port.FixedIPs)
	d.Set("port_id", port.ID)
	d.Set("mac_address", port.MACAddress)
	d.Set("managed_by_service", port.ManagedByService)
	d.Set("name", port.Name)
	d.Set("network_id", port.NetworkID)
	d.Set("segmentation_id", port.SegmentationID)
	d.Set("segmentation_type", port.SegmentationType)
	d.Set("status", port.Status)
	d.Set("tags", port.Tags)
	d.Set("tenant_id", port.TenantID)

	return nil
}
