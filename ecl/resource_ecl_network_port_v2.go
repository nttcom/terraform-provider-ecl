package ecl

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/ports"
)

func resourceNetworkPortV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkPortV2Create,
		Read:   resourceNetworkPortV2Read,
		Update: resourceNetworkPortV2Update,
		Delete: resourceNetworkPortV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"admin_state_up": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"all_fixed_ips": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"allowed_address_pairs": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"device_owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			/*
				In ECL2.0 removing existing fixed_ip is not allowed.
				Adding new element is only allowed.
			*/
			"fixed_ip": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"no_fixed_ip"},
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
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"managed_by_service": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"no_fixed_ip": &schema.Schema{
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"fixed_ip"},
			},
			"segmentation_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"segmentation_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"flat", "vlan"}, false),
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func resourceNetworkPortV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	createOpts := PortCreateOpts{
		ports.CreateOpts{
			AdminStateUp:        resourcePortAdminStateUpV2(d),
			AllowedAddressPairs: resourceAllowedAddressPairsV2(d),
			Description:         d.Get("description").(string),
			DeviceID:            d.Get("device_id").(string),
			DeviceOwner:         d.Get("device_owner").(string),
			FixedIPs:            resourcePortFixedIpsV2(d),
			MACAddress:          d.Get("mac_address").(string),
			Name:                d.Get("name").(string),
			NetworkID:           d.Get("network_id").(string),
			SegmentationID:      d.Get("segmentation_id").(int),
			SegmentationType:    d.Get("segmentation_type").(string),
			Tags:                resourceTags(d),
			TenantID:            d.Get("tenant_id").(string),
		},
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	p, err := ports.Create(networkingClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL Network network: %s", err)
	}

	d.SetId(p.ID)
	log.Printf("[INFO] Network ID: %s", p.ID)

	log.Printf("[DEBUG] Waiting for ECL Network Port (%s) to become available.", p.ID)

	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE"},
		Refresh:    waitForNetworkPortActive(networkingClient, p.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()

	return resourceNetworkPortV2Read(d, meta)
}

func resourceNetworkPortV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	p, err := ports.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "port")
	}

	log.Printf("[DEBUG] Retrieved Port %s: %+v", d.Id(), p)

	d.Set("admin_state_up", p.AdminStateUp)
	d.Set("allowed_address_pairs", flattenNetworkPortAllowedAddressPairsV2(p.MACAddress, p.AllowedAddressPairs))
	d.Set("description", p.Description)
	d.Set("device_id", p.DeviceID)
	d.Set("device_owner", p.DeviceOwner)
	d.Set("all_fixed_ips", expandNetworkPortFixedIPToStringSlice(p.FixedIPs))
	d.Set("fixed_ips", p.FixedIPs)
	d.Set("mac_address", p.MACAddress)
	d.Set("managed_by_service", p.ManagedByService)
	d.Set("name", p.Name)
	d.Set("network_id", p.NetworkID)
	d.Set("region", GetRegion(d, config))
	d.Set("segmentation_id", p.SegmentationID)
	d.Set("segmentation_type", p.SegmentationType)
	d.Set("status", p.Status)
	d.Set("tags", p.Tags)
	d.Set("tenant_id", p.TenantID)

	return nil
}

func resourceNetworkPortV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	var hasChange bool
	var updateOpts ports.UpdateOpts

	if d.HasChange("admin_state_up") {
		hasChange = true
		updateOpts.AdminStateUp = resourcePortAdminStateUpV2(d)
	}

	if d.HasChange("allowed_address_pairs") {
		hasChange = true
		aap := resourceAllowedAddressPairsV2(d)
		updateOpts.AllowedAddressPairs = &aap
	}

	if d.HasChange("description") {
		hasChange = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("device_id") {
		hasChange = true
		deviceID := d.Get("device_id").(string)
		updateOpts.DeviceID = &deviceID
	}

	if d.HasChange("device_owner") {
		hasChange = true
		deviceOwner := d.Get("device_owner").(string)
		updateOpts.DeviceOwner = &deviceOwner
	}

	if d.HasChange("fixed_ip") || d.HasChange("no_fixed_ip") {
		hasChange = true
		updateOpts.FixedIPs = resourcePortFixedIpsV2(d)
	}

	if d.HasChange("name") {
		hasChange = true
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("segmentation_id") {
		hasChange = true
		segmentationID := d.Get("segmentation_id").(int)
		updateOpts.SegmentationID = &segmentationID
	}

	if d.HasChange("segmentation_type") {
		hasChange = true
		segmentationType := d.Get("segmentation_type").(string)
		updateOpts.SegmentationType = &segmentationType
	}

	if d.HasChange("tags") {
		hasChange = true
		tags := resourceTags(d)
		updateOpts.Tags = &tags
	}

	if hasChange {
		log.Printf("[DEBUG] Updating Port %s with options: %+v", d.Id(), updateOpts)

		_, err = ports.Update(networkingClient, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating ECL Network Network: %s", err)
		}
	}

	return resourceNetworkPortV2Read(d, meta)
}

func resourceNetworkPortV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	p, err := ports.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		if _, ok := err.(eclcloud.ErrDefault404); ok {
			log.Printf("[DEBUG] ECL port already deleted %s", d.Id())
			d.SetId("")
			return nil
		}
	}

	if p.Status != "PENDING_DELETE" {
		err = ports.Delete(networkingClient, d.Id()).ExtractErr()
		if err != nil {
			return fmt.Errorf("Error deleting ECL network: %s", err)
		}
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForNetworkPortDelete(networkingClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Network Network: %s", err)
	}

	d.SetId("")
	return nil
}

func resourcePortFixedIpsV2(d *schema.ResourceData) interface{} {
	// if no_fixed_ip was specified, then just return
	// an empty array. Since no_fixed_ip is mutually
	// exclusive to fixed_ip, we can safely do this.
	//
	// Since we're only concerned about no_fixed_ip
	// being set to "true", GetOk is used.
	if _, ok := d.GetOk("no_fixed_ip"); ok {
		return []interface{}{}
	}

	rawIP := d.Get("fixed_ip").([]interface{})

	if len(rawIP) == 0 {
		return nil
	}

	ip := make([]ports.IP, len(rawIP))
	for i, raw := range rawIP {
		rawMap := raw.(map[string]interface{})
		ip[i] = ports.IP{
			SubnetID:  rawMap["subnet_id"].(string),
			IPAddress: rawMap["ip_address"].(string),
		}
	}
	return ip
}

func resourceAllowedAddressPairsV2(d *schema.ResourceData) []ports.AddressPair {
	// ports.AddressPair
	rawPairs := d.Get("allowed_address_pairs").(*schema.Set).List()

	pairs := make([]ports.AddressPair, len(rawPairs))
	for i, raw := range rawPairs {
		rawMap := raw.(map[string]interface{})
		pairs[i] = ports.AddressPair{
			IPAddress:  rawMap["ip_address"].(string),
			MACAddress: rawMap["mac_address"].(string),
		}
	}
	return pairs
}

func resourcePortAdminStateUpV2(d *schema.ResourceData) *bool {
	value := false

	if raw, ok := d.GetOk("admin_state_up"); ok && raw == true {
		value = true
	}

	return &value
}

func allowedAddressPairsHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-%s", m["ip_address"].(string), m["mac_address"].(string)))

	return hashcode.String(buf.String())
}

func flattenNetworkPortAllowedAddressPairsV2(mac string, allowedAddressPairs []ports.AddressPair) []map[string]interface{} {
	// Convert AllowedAddressPairs to list of map
	var pairs []map[string]interface{}
	for _, pairObject := range allowedAddressPairs {
		pair := make(map[string]interface{})
		pair["ip_address"] = pairObject.IPAddress

		// Only set the MAC address if it is different than the
		// port's MAC. This means that a specific MAC was set.
		if mac != pairObject.MACAddress {
			pair["mac_address"] = pairObject.MACAddress
		}

		pairs = append(pairs, pair)
	}
	return pairs
}

func expandNetworkPortFixedIPToStringSlice(fixedIPs []ports.IP) []string {
	s := make([]string, len(fixedIPs))
	for i, fixedIP := range fixedIPs {
		s[i] = fixedIP.IPAddress
	}

	return s
}

func waitForNetworkPortActive(networkingClient *eclcloud.ServiceClient, portID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		p, err := ports.Get(networkingClient, portID).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Network Port: %+v", p)
		if p.Status == "DOWN" || p.Status == "ACTIVE" {
			return p, "ACTIVE", nil
		}

		return p, p.Status, nil
	}
}

func waitForNetworkPortDelete(networkingClient *eclcloud.ServiceClient, portID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Network Port %s", portID)

		p, err := ports.Get(networkingClient, portID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Port %s", portID)
				return p, "DELETED", nil
			}
			return p, "ACTIVE", err
		}

		log.Printf("[DEBUG] ECL Port %s still active.\n", portID)
		return p, "ACTIVE", nil
	}
}
