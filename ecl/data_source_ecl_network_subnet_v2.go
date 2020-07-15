package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/network/v2/subnets"
)

func dataSourceNetworkSubnetV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkSubnetV2Read,

		Schema: map[string]*schema.Schema{
			"allocation_pools": &schema.Schema{
				Type:     schema.TypeList,
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
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_nameservers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enable_dhcp": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_routes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
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
				Computed: true,
			},
			"ipv6_address_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_ra_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_servers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "This region field is deprecated and will be removed from a future version.",
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: descriptions["tenant_id"],
			},
		},
	}
}

func dataSourceNetworkSubnetV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))

	listOpts := subnets.ListOpts{}

	if v, ok := d.GetOk("cidr"); ok {
		listOpts.CIDR = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("gateway_ip"); ok {
		listOpts.GatewayIP = v.(string)
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("network_id"); ok {
		listOpts.NetworkID = v.(string)
	}

	pages, err := subnets.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve subnets: %s", err)
	}

	allSubnets, err := subnets.ExtractSubnets(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract subnets: %s", err)
	}

	if len(allSubnets) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allSubnets) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	subnet := allSubnets[0]

	log.Printf("[DEBUG] Retrieved Subnet %s: %+v", subnet.ID, subnet)

	// allocation_pools is set below
	d.Set("cidr", subnet.CIDR)
	d.Set("description", subnet.Description)
	// dns_nameservers is set below
	d.Set("enable_dhcp", subnet.EnableDHCP)
	d.Set("gateway_ip", subnet.GatewayIP)
	// host_routes is set below
	d.SetId(subnet.ID)
	d.Set("ip_version", subnet.IPVersion)
	d.Set("ipv6_address_mode", subnet.IPv6AddressMode)
	d.Set("ipv6_ra_mode", subnet.IPv6RAMode)
	d.Set("name", subnet.Name)
	d.Set("network_id", subnet.NetworkID)
	// ntp_servers is set below
	d.Set("status", subnet.Status)
	// tags is set below
	d.Set("tenant_id", subnet.TenantID)

	err = d.Set("allocation_pools", flattenAllocationPools(subnet.AllocationPools))
	if err != nil {
		log.Printf("[DEBUG] Unable to set allocation_pools: %s", err)
	}

	err = d.Set("dns_nameservers", subnet.DNSNameservers)
	if err != nil {
		log.Printf("[DEBUG] Unable to set dns_nameservers: %s", err)
	}

	err = d.Set("host_routes", flattenHostRoutes(subnet.HostRoutes))
	if err != nil {
		log.Printf("[DEBUG] Unable to set host_routes: %s", err)
	}

	err = d.Set("ntp_servers", subnet.NTPServers)
	if err != nil {
		log.Printf("[DEBUG] Unable to set ntp_servers: %s", err)
	}

	err = d.Set("tags", subnet.Tags)
	if err != nil {
		log.Printf("[DEBUG] Unable to set ntp_servers: %s", err)
	}

	return nil
}

func flattenAllocationPools(in []subnets.AllocationPool) []map[string]interface{} {
	var out []map[string]interface{}
	for _, v := range in {
		pool := make(map[string]interface{})
		pool["start"] = v.Start
		pool["end"] = v.End

		out = append(out, pool)
	}
	return out
}

func flattenHostRoutes(in []subnets.HostRoute) []map[string]interface{} {
	var out []map[string]interface{}
	for _, v := range in {
		route := make(map[string]interface{})
		route["destination_cidr"] = v.DestinationCIDR
		route["next_hop"] = v.NextHop

		out = append(out, route)
	}
	return out
}
