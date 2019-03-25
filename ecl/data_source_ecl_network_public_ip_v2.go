package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/network/v2/public_ips"
)

func dataSourceNetworkPublicIPV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkPublicIPV2Read,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"internet_gw_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_ip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"submask_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkPublicIPV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))

	listOpts := public_ips.ListOpts{}

	if v, ok := d.GetOk("cidr"); ok {
		listOpts.Cidr = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("public_ip_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		listOpts.Status = v.(string)
	}

	if v, ok := d.GetOk("submask_length"); ok {
		listOpts.SubmaskLength = v.(int)
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	pages, err := public_ips.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve public_ips: %s", err)
	}

	allPublicIPs, err := public_ips.ExtractPublicIPs(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract public_ips: %s", err)
	}

	if len(allPublicIPs) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allPublicIPs) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	public_ip := allPublicIPs[0]

	log.Printf("[DEBUG] Retrieved PublicIP %s: %+v", public_ip.ID, public_ip)
	d.SetId(public_ip.ID)

	d.Set("cidr", public_ip.Cidr)
	d.Set("description", public_ip.Description)
	d.Set("internet_gw_id", public_ip.InternetGwID)
	d.Set("name", public_ip.Name)
	d.Set("status", public_ip.Status)
	d.Set("submask_length", public_ip.SubmaskLength)
	d.Set("tenant_id", public_ip.TenantID)
	d.Set("region", GetRegion(d, config))

	return nil
}
