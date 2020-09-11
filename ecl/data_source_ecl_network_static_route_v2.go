package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/network/v2/static_routes"
)

func dataSourceNetworkStaticRouteV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkStaticRouteV2Read,
		Schema: map[string]*schema.Schema{
			"region": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: "This attribute is not used to set up the resource",
			},
			"aws_gw_id": {
				Type:          schema.TypeString,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"azure_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"azure_gw_id": {
				Type:          schema.TypeString,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"aws_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"destination": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gcp_gw_id": {
				Type:          schema.TypeString,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "interdc_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"interdc_gw_id": {
				Type:          schema.TypeString,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "gcp_gw_id", "internet_gw_id", "vpn_gw_id"},
			},
			"internet_gw_id": {
				Type:          schema.TypeString,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "vpn_gw_id"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"nexthop": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"service_type": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"static_route_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vpn_gw_id": {
				Type:          schema.TypeString,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"aws_gw_id", "azure_gw_id", "gcp_gw_id", "interdc_gw_id", "internet_gw_id"},
			},
		},
	}
}

func dataSourceNetworkStaticRouteV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))

	listOpts := static_routes.ListOpts{}

	if v, ok := d.GetOk("aws_gw_id"); ok {
		listOpts.AwsGwID = v.(string)
	}

	if v, ok := d.GetOk("azure_gw_id"); ok {
		listOpts.AzureGwID = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("destination"); ok {
		listOpts.Destination = v.(string)
	}

	if v, ok := d.GetOk("gcp_gw_id"); ok {
		listOpts.GcpGwID = v.(string)
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

	if v, ok := d.GetOk("nexthop"); ok {
		listOpts.Nexthop = v.(string)
	}

	if v, ok := d.GetOk("service_type"); ok {
		listOpts.ServiceType = v.(string)
	}

	if v, ok := d.GetOk("static_route_id"); ok {
		listOpts.ID = v.(string)
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

	pages, err := static_routes.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("unable to retrieve static_routes: %w", err)
	}

	allStaticRoutes, err := static_routes.ExtractStaticRoutes(pages)
	if err != nil {
		return fmt.Errorf("unable to extract static_routes: %w", err)
	}

	if len(allStaticRoutes) < 1 {
		return fmt.Errorf("your query returned no results. " +
			"please change your search criteria and try again")
	}

	if len(allStaticRoutes) > 1 {
		return fmt.Errorf("your query returned more than one result. " +
			"please try a more specific search criteria")
	}

	staticRoute := allStaticRoutes[0]

	log.Printf("[DEBUG] Retrieved StaticRoute %s: %+v", staticRoute.ID, staticRoute)
	d.SetId(staticRoute.ID)

	d.Set("aws_gw_id", staticRoute.AwsGwID)
	d.Set("azure_gw_id", staticRoute.AzureGwID)
	d.Set("description", staticRoute.Description)
	d.Set("destination", staticRoute.Destination)
	d.Set("gcp_gw_id", staticRoute.GcpGwID)
	d.Set("interdc_gw_id", staticRoute.InterdcGwID)
	d.Set("internet_gw_id", staticRoute.InternetGwID)
	d.Set("name", staticRoute.Name)
	d.Set("nexthop", staticRoute.Nexthop)
	d.Set("service_type", staticRoute.ServiceType)
	d.Set("status", staticRoute.Status)
	d.Set("tenant_id", staticRoute.TenantID)
	d.Set("vpn_gw_id", staticRoute.VpnGwID)
	d.Set("region", GetRegion(d, config))

	return nil
}
