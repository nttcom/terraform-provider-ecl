package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/ecl/network/v2/subnets"
)

func dataSourceNetworkNetworkV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkNetworkV2Read,

		Schema: map[string]*schema.Schema{
			"admin_state_up": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"matching_subnet_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"plane": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnets": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func dataSourceNetworkNetworkV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))

	var listOpts networks.ListOptsBuilder

	var status string
	if v, ok := d.GetOk("status"); ok {
		status = v.(string)
	}

	listOpts = networks.ListOpts{
		Description: d.Get("description").(string),
		ID:          d.Get("network_id").(string),
		Name:        d.Get("name").(string),
		Plane:       d.Get("plane").(string),
		Status:      status,
		TenantID:    d.Get("tenant_id").(string),
	}

	pages, err := networks.List(networkClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	// First extract into a normal networks.Network in order to see if
	// there were any results at all.
	tmpAllNetworks, err := networks.ExtractNetworks(pages)
	if err != nil {
		return err
	}

	if len(tmpAllNetworks) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	var allNetworks []networks.Network
	err = networks.ExtractNetworksInto(pages, &allNetworks)
	if err != nil {
		return fmt.Errorf("Unable to retrieve networks: %s", err)
	}

	var refinedNetworks []networks.Network
	if cidr := d.Get("matching_subnet_cidr").(string); cidr != "" {
		for _, n := range allNetworks {
			for _, s := range n.Subnets {
				subnet, err := subnets.Get(networkClient, s).Extract()
				if err != nil {
					if _, ok := err.(eclcloud.ErrDefault404); ok {
						continue
					}
					return fmt.Errorf("Unable to retrieve network subnet: %s", err)
				}
				if cidr == subnet.CIDR {
					refinedNetworks = append(refinedNetworks, n)
				}
			}
		}
	} else {
		refinedNetworks = allNetworks
	}

	if len(refinedNetworks) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(refinedNetworks) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	network := refinedNetworks[0]

	log.Printf("[DEBUG] Retrieved Network %s: %+v", network.ID, network)
	d.SetId(network.ID)

	d.Set("admin_state_up", network.AdminStateUp)
	d.Set("description", network.Description)
	d.Set("name", network.Name)
	d.Set("plane", network.Plane)
	d.Set("status", network.Status)
	d.Set("subnets", network.Subnets)
	d.Set("tags", network.Tags)
	d.Set("tenant_id", network.TenantID)

	return nil
}
