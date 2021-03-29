package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v2/ecl/network/v2/common_function_pool"
)

func dataSourceNetworkCommonFunctionPoolV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNetworkCommonFunctionPoolV2Read,

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceNetworkCommonFunctionPoolV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.networkV2Client(GetRegion(d, config))

	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	var opts common_function_pool.ListOpts

	if v, ok := d.GetOk("description"); ok {
		opts.Description = v.(string)
	}

	if v, ok := d.GetOk("id"); ok {
		opts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		opts.Name = v.(string)
	}

	pages, err := common_function_pool.List(client, opts).AllPages()
	if err != nil {
		return fmt.Errorf("unable to retrieve common_function_pool: %w", err)
	}

	cfps, err := common_function_pool.ExtractCommonFunctionPools(pages)
	if err != nil {
		return fmt.Errorf("unable to extract common_function_pool: %w", err)
	}

	if len(cfps) < 1 {
		return fmt.Errorf("your query returned no results." +
			" please change your search criteria and try again")
	}

	if len(cfps) > 1 {
		return fmt.Errorf("your query returned more than one result." +
			" please try a more specific search criteria")
	}

	cfp := cfps[0]

	log.Printf("[DEBUG] Retrieved CommonFunctionPool %s: %+v", cfp.ID, cfp)
	d.SetId(cfp.ID)

	d.Set("description", cfp.Description)
	d.Set("name", cfp.Name)

	return nil
}
