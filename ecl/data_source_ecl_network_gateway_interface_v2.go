package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/network/v2/gateway_interfaces"
)

func dataSourceNetworkGatewayInterfaceV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkGatewayInterfaceV2Read,
		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"aws_gw_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"azure_gw_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gateway_interface_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gcp_gw_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gw_vipv4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gw_vipv6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"interdc_gw_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internet_gw_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"primary_ipv4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"primary_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secondary_ipv4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secondary_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vpn_gw_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vrid": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceNetworkGatewayInterfaceV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))

	listOpts := gateway_interfaces.ListOpts{}

	if v, ok := d.GetOk("aws_gw_id"); ok {
		listOpts.AwsGwID = v.(string)
	}

	if v, ok := d.GetOk("azure_gw_id"); ok {
		listOpts.AzureGwID = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("gcp_gw_id"); ok {
		listOpts.GcpGwID = v.(string)
	}

	if v, ok := d.GetOk("gw_vipv4"); ok {
		listOpts.GwVipv4 = v.(string)
	}

	if v, ok := d.GetOk("gw_vipv6"); ok {
		listOpts.GwVipv6 = v.(string)
	}

	if v, ok := d.GetOk("gateway_interface_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("interdc_gw_id"); ok {
		listOpts.InterdcGwID = v.(string)
	}

	if v, ok := d.GetOk("internet_gw_id"); ok {
		listOpts.InternetGwID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("netmask"); ok {
		listOpts.Netmask = v.(int)
	}

	if v, ok := d.GetOk("network_id"); ok {
		listOpts.NetworkID = v.(string)
	}

	if v, ok := d.GetOk("primary_ipv4"); ok {
		listOpts.PrimaryIpv4 = v.(string)
	}

	if v, ok := d.GetOk("primary_ipv6"); ok {
		listOpts.PrimaryIpv6 = v.(string)
	}

	if v, ok := d.GetOk("secondary_ipv4"); ok {
		listOpts.SecondaryIpv4 = v.(string)
	}

	if v, ok := d.GetOk("secondary_ipv6"); ok {
		listOpts.SecondaryIpv6 = v.(string)
	}

	if v, ok := d.GetOk("service_type"); ok {
		listOpts.ServiceType = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	if v, ok := d.GetOk("vpn_gw_id"); ok {
		listOpts.VpnGwID = v.(string)
	}

	if v, ok := d.GetOk("vrid"); ok {
		listOpts.VRID = v.(int)
	}

	pages, err := gateway_interfaces.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve gateway_interfaces: %s", err)
	}

	allGatewayInterfaces, err := gateway_interfaces.ExtractGatewayInterfaces(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract gateway_interfaces: %s", err)
	}

	if len(allGatewayInterfaces) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allGatewayInterfaces) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	gateway_interface := allGatewayInterfaces[0]

	log.Printf("[DEBUG] Retrieved GatewayInterface %s: %+v", gateway_interface.ID, gateway_interface)
	d.SetId(gateway_interface.ID)

	d.Set("aws_gw_id", gateway_interface.AwsGwID)
	d.Set("azure_gw_id", gateway_interface.AzureGwID)
	d.Set("description", gateway_interface.Description)
	d.Set("gcp_gw_id", gateway_interface.GcpGwID)
	d.Set("gw_vipv4", gateway_interface.GwVipv4)
	d.Set("gw_vipv6", gateway_interface.GwVipv6)
	d.Set("interdc_gw_id", gateway_interface.InterdcGwID)
	d.Set("internet_gw_id", gateway_interface.InternetGwID)
	d.Set("name", gateway_interface.Name)
	d.Set("netmask", gateway_interface.Netmask)
	d.Set("network_id", gateway_interface.NetworkID)
	d.Set("primary_ipv4", gateway_interface.PrimaryIpv4)
	d.Set("primary_ipv6", gateway_interface.PrimaryIpv6)
	d.Set("secondary_ipv4", gateway_interface.SecondaryIpv4)
	d.Set("secondary_ipv6", gateway_interface.SecondaryIpv6)
	d.Set("service_type", gateway_interface.ServiceType)
	d.Set("status", gateway_interface.Status)
	d.Set("tenant_id", gateway_interface.TenantID)
	d.Set("vpn_gw_id", gateway_interface.VpnGwID)
	d.Set("tenant_id", gateway_interface.TenantID)
	d.Set("vrid", gateway_interface.VRID)
	d.Set("region", GetRegion(d, config))

	return nil
}
