package ecl

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/network/v2/subnets"
)

func resourceNetworkSubnetV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSubnetV2Create,
		Read:   resourceNetworkSubnetV2Read,
		Update: resourceNetworkSubnetV2Update,
		Delete: resourceNetworkSubnetV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"allocation_pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"end": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_nameservers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enable_dhcp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"gateway_ip": &schema.Schema{
				Type:          schema.TypeString,
				ConflictsWith: []string{"no_gateway"},
				Optional:      true,
				Computed:      true,
			},
			"host_routes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_cidr": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"next_hop": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  4,
			},
			"ipv6_address_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"ipv6_ra_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
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
			"no_gateway": &schema.Schema{
				Type:          schema.TypeBool,
				ConflictsWith: []string{"gateway_ip"},
				Optional:      true,
			},
			"ntp_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: "This attribute is not used to set up the resource.",
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

func resourceNetworkSubnetV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	if err = resourceSubnetDNSNameserversV2CheckIsSet(d); err != nil {
		return err
	}

	createOpts := SubnetCreateOpts{
		subnets.CreateOpts{
			CIDR:      d.Get("cidr").(string),
			NetworkID: d.Get("network_id").(string),
		},
	}

	if v, ok := resourceSubnetAllocationPoolsV2(d); ok {
		createOpts.AllocationPools = v
	}

	if v, ok := d.GetOk("description"); ok {
		createOpts.Description = v.(string)
	}

	if v, ok := resourceSubnetDNSNameserversV2(d); ok {
		createOpts.DNSNameservers = v
	}

	if v, ok := d.GetOkExists("enable_dhcp"); ok {
		enableDHCP := v.(bool)
		createOpts.EnableDHCP = &enableDHCP
	}

	if v, ok := d.GetOk("gateway_ip"); ok {
		gatewayIP := v.(string)
		createOpts.GatewayIP = &gatewayIP
	}

	if v, ok := resourceSubnetHostRoutesV2(d); ok {
		createOpts.HostRoutes = v
	}

	if v, ok := d.GetOk("ip_version"); ok {
		createOpts.IPVersion = resourceNetworkSubnetV2DetermineIPVersion(v.(int))
	}

	if v, ok := d.GetOk("name"); ok {
		createOpts.Name = v.(string)
	}

	noGateway := d.Get("no_gateway").(bool)
	if noGateway {
		gatewayIP := ""
		createOpts.GatewayIP = &gatewayIP
	}

	if v, ok := resourceSubnetNTPServersV2(d); ok {
		createOpts.NTPServers = v
	}

	if _, ok := d.GetOk("tags"); ok {
		createOpts.Tags = resourceTags(d)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		createOpts.TenantID = v.(string)
	}

	log.Printf("[DEBUG] Creating subnet: %v", createOpts)

	s, err := subnets.Create(networkClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL Network subnet: %s", err)
	}

	d.SetId(s.ID)

	log.Printf("[DEBUG] Waiting for Subnet (%s) to become available", s.ID)
	stateConf := &resource.StateChangeConf{
		Target:     []string{"ACTIVE"},
		Refresh:    waitForSubnetActive(networkClient, s.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()

	log.Printf("[DEBUG] Created Subnet %s: %#v", s.ID, s)
	return resourceNetworkSubnetV2Read(d, meta)
}

func resourceNetworkSubnetV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	s, err := subnets.Get(networkClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "subnet")
	}

	log.Printf("[DEBUG] Retrieved Subnet %s: %#v", d.Id(), s)

	// allocation_pools is set below
	d.Set("cidr", s.CIDR)
	d.Set("description", s.Description)
	// dns_nameservers is set below
	d.Set("enable_dhcp", s.EnableDHCP)
	d.Set("gateway_ip", s.GatewayIP)
	// host_routes is set below
	d.SetId(s.ID)
	d.Set("ip_version", s.IPVersion)
	d.Set("ipv6_address_mode", s.IPv6AddressMode)
	d.Set("ipv6_ra_mode", s.IPv6RAMode)
	d.Set("name", s.Name)
	d.Set("network_id", s.NetworkID)
	d.Set("region", GetRegion(d, config))
	// ntp_servers is set below
	d.Set("status", s.Status)
	// tags is set below
	d.Set("tenant_id", s.TenantID)

	err = d.Set("allocation_pools", flattenAllocationPools(s.AllocationPools))
	if err != nil {
		log.Printf("[DEBUG] Unable to set allocation_pools: %s", err)
	}

	err = d.Set("dns_nameservers", s.DNSNameservers)
	if err != nil {
		log.Printf("[DEBUG] Unable to set dns_nameservers: %s", err)
	}

	err = d.Set("host_routes", flattenHostRoutes(s.HostRoutes))
	if err != nil {
		log.Printf("[DEBUG] Unable to set host_routes: %s", err)
	}

	err = d.Set("ntp_servers", s.NTPServers)
	if err != nil {
		log.Printf("[DEBUG] Unable to set ntp_servers: %s", err)
	}

	err = d.Set("tags", s.Tags)
	if err != nil {
		log.Printf("[DEBUG] Unable to set ntp_servers: %s", err)
	}

	// Based on the subnet's Gateway IP, set `no_gateway` accordingly.
	if s.GatewayIP == "" {
		d.Set("no_gateway", true)
	} else {
		d.Set("no_gateway", false)
	}

	return nil
}

func resourceNetworkSubnetV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	var updateOpts subnets.UpdateOpts

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if d.HasChange("dns_nameservers") {
		if err = resourceSubnetDNSNameserversV2CheckIsSet(d); err != nil {
			return err
		}
		dns, _ := resourceSubnetDNSNameserversV2(d)
		log.Printf("[DEBUG] dns is set as : %#v", dns)
		updateOpts.DNSNameservers = dns
	}

	if d.HasChange("enable_dhcp") {
		v := d.Get("enable_dhcp").(bool)
		updateOpts.EnableDHCP = &v
	}

	if d.HasChange("gateway_ip") {
		updateOpts.GatewayIP = nil
		if v, ok := d.GetOk("gateway_ip"); ok {
			gatewayIP := v.(string)
			updateOpts.GatewayIP = &gatewayIP
		}
	}

	if d.HasChange("host_routes") {
		newHostRoutes, _ := resourceSubnetHostRoutesV2(d)
		updateOpts.HostRoutes = &newHostRoutes
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateOpts.Name = &name
	}

	if d.HasChange("ntp_servers") {
		ntpServers, _ := resourceSubnetNTPServersV2(d)
		updateOpts.NTPServers = &ntpServers
	}

	if d.HasChange("tags") {
		tags := resourceTags(d)
		updateOpts.Tags = &tags
	}

	log.Printf("[DEBUG] Updating Subnet %s with options: %+v", d.Id(), updateOpts)

	_, err = subnets.Update(networkClient, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating ECL Network Subnet: %s", err)
	}

	return resourceNetworkSubnetV2Read(d, meta)
}

func resourceNetworkSubnetV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	s, err := subnets.Get(networkClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "subnet")
	}

	if s.Status != "PENDING_DELETE" {
		err := avoidConflictForSubnetDelete(networkClient, d.Id())
		if err != nil {
			return CheckDeleted(d, err, "subnet")
		}
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE", "PENDING_DELETE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForSubnetDelete(networkClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting ECL Network Subnet: %s", err)
	}

	d.SetId("")
	return nil
}

func avoidConflictForSubnetDelete(client *eclcloud.ServiceClient, id string) error {
	// Keep waiting 5 minutes in deletion at most
	retryMax := 10
	sleepSecondInEachRetry := 30

	for i := 0; i < retryMax; i++ {
		err := subnets.Delete(client, id).ExtractErr()
		log.Printf("%d th deleting try: result is %#v", i, err)

		if err != nil {
			_, ok := err.(eclcloud.ErrDefault409)
			if ok {
				log.Printf("[DEBUG] Sleeping for retry deletion")
				time.Sleep(time.Second * time.Duration(sleepSecondInEachRetry))
				continue
			} else {
				return fmt.Errorf("Failed in subnet deletion")
			}
		}
		return nil
	}
	log.Printf("[DEBUG] Reached maximun retry count of deletion")
	return nil
}

func resourceSubnetAllocationPoolsV2(d *schema.ResourceData) ([]subnets.AllocationPool, bool) {
	resources, ok := d.GetOk("allocation_pools")
	rawAPs := resources.([]interface{})
	aps := make([]subnets.AllocationPool, len(rawAPs))
	for i, raw := range rawAPs {
		rawMap := raw.(map[string]interface{})
		aps[i] = subnets.AllocationPool{
			Start: rawMap["start"].(string),
			End:   rawMap["end"].(string),
		}
	}
	return aps, ok
}

func resourceSubnetDNSNameserversV2(d *schema.ResourceData) ([]string, bool) {
	resources, ok := d.GetOk("dns_nameservers")
	rawDNSN := resources.([]interface{})
	dnsn := make([]string, len(rawDNSN))
	for i, raw := range rawDNSN {
		dnsn[i] = raw.(string)
	}
	return dnsn, ok
}

func resourceSubnetNTPServersV2(d *schema.ResourceData) ([]string, bool) {
	resources, ok := d.GetOk("ntp_servers")
	rawNTP := resources.([]interface{})
	ntps := make([]string, len(rawNTP))
	for i, raw := range rawNTP {
		ntps[i] = raw.(string)
	}
	return ntps, ok
}

func resourceSubnetDNSNameserversV2CheckIsSet(d *schema.ResourceData) error {
	rawDNSN := d.Get("dns_nameservers").([]interface{})
	set := make(map[string]*string)
	for _, raw := range rawDNSN {
		dns := raw.(string)
		if set[dns] != nil {
			return fmt.Errorf("DNS nameservers must appear exactly once: %q", dns)
		} else {
			set[dns] = &dns
		}
	}
	return nil
}

func resourceSubnetHostRoutesV2(d *schema.ResourceData) ([]subnets.HostRoute, bool) {
	resources, ok := d.GetOk("host_routes")
	rawHR := resources.([]interface{})
	hr := make([]subnets.HostRoute, len(rawHR))
	for i, raw := range rawHR {
		rawMap := raw.(map[string]interface{})
		hr[i] = subnets.HostRoute{
			DestinationCIDR: rawMap["destination_cidr"].(string),
			NextHop:         rawMap["next_hop"].(string),
		}
	}
	return hr, ok
}

func resourceNetworkSubnetV2DetermineIPVersion(v int) eclcloud.IPVersion {
	var ipVersion eclcloud.IPVersion
	switch v {
	case 4:
		ipVersion = eclcloud.IPv4
	case 6:
		ipVersion = eclcloud.IPv6
	}

	return ipVersion
}

func waitForSubnetActive(networkClient *eclcloud.ServiceClient, subnetId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		s, err := subnets.Get(networkClient, subnetId).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] ECL Network Subnet: %+v", s)
		return s, "ACTIVE", nil
	}
}

func waitForSubnetDelete(networkClient *eclcloud.ServiceClient, subnetId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete ECL Subnet %s.\n", subnetId)

		s, err := subnets.Get(networkClient, subnetId).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted ECL Subnet %s", subnetId)
				return s, "DELETED", nil
			}
			return s, "ACTIVE", err
		}

		log.Printf("[DEBUG] ECL Subnet %s still active.\n", subnetId)
		return s, "ACTIVE", nil
	}
}
