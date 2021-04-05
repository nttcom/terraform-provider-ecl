package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v2/ecl/network/v2/internet_services"
)

func dataSourceNetworkInternetServiceV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkInternetServiceV2Read,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "This attribute is not used to set up the resource.",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"minimal_submask_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceNetworkInternetServiceV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkClient, err := config.networkV2Client(GetRegion(d, config))

	listOpts := internet_services.ListOpts{}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("internet_service_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("minimal_submask_length"); ok {
		listOpts.MinimalSubmaskLength = v.(int)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("zone"); ok {
		listOpts.Zone = v.(string)
	}

	pages, err := internet_services.List(networkClient, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to retrieve internet_services: %s", err)
	}

	allInternetServices, err := internet_services.ExtractInternetServices(pages)
	if err != nil {
		return fmt.Errorf("Unable to extract internet_services: %s", err)
	}

	if len(allInternetServices) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allInternetServices) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	internetService := allInternetServices[0]

	log.Printf("[DEBUG] Retrieved InternetService %s: %+v", internetService.ID, internetService)
	d.SetId(internetService.ID)

	d.Set("description", internetService.Description)
	d.Set("minimal_submask_length", internetService.MinimalSubmaskLength)
	d.Set("name", internetService.Name)
	d.Set("zone", internetService.Zone)
	d.Set("region", GetRegion(d, config))

	return nil
}
