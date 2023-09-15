package ecl

import (
	"fmt"
	"log"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/vna/v1/appliance_plans"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceVNAAppliancePlanV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNAAppliancePlanV1Read,

		Schema: map[string]*schema.Schema{

			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"appliance_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"flavor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"number_of_interfaces": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"max_number_of_aap": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"licenses": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"license_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			"availability_zones": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"availability_zone": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"available": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"rank": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVNAAppliancePlanV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	var opts appliance_plans.ListOpts

	if v, ok := d.GetOk("id"); ok {
		opts.ID = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		opts.Name = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		opts.Description = v.(string)
	}
	if v, ok := d.GetOk("appliance_type"); ok {
		opts.ApplianceType = v.(string)
	}
	if v, ok := d.GetOk("version"); ok {
		opts.Version = v.(string)
	}
	if v, ok := d.GetOk("flavor"); ok {
		opts.Flavor = v.(string)
	}
	if v, ok := d.GetOk("number_of_interfaces"); ok {
		opts.NumberOfInterfaces = v.(int)
	}
	if v, ok := d.GetOk("enabled"); ok {
		opts.Enabled = v.(bool)
	}
	if v, ok := d.GetOk("max_number_of_aap"); ok {
		opts.MaxNumberOfAap = v.(int)
	}
	if v, ok := d.GetOk("availability_zone"); ok {
		opts.AvailabilityZone = v.(string)
	}
	if v, ok := d.GetOk("availability_zone.available"); ok {
		opts.AvailabilityZoneAvailable = v.(bool)
	}

	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	allPlans, err := getVirtualNetworkAppliancePlans(vnaClient, opts)
	if err != nil {
		return fmt.Errorf("error getting Load Balancer Plans: %w", err)
	}

	if len(allPlans) > 1 {
		return fmt.Errorf("specified Virtual Network Appliance Plan query returned more than one result")
	}

	if len(allPlans) == 0 {
		return fmt.Errorf("specified Virtual Network Appliance Plan query returned no results")
	}

	plan := allPlans[0]

	log.Printf("[DEBUG] Retrieved Virtual Network Appliance Plan %s: %+v", plan.ID, plan)

	d.SetId(plan.ID)
	d.Set("id", plan.ID)
	d.Set("name", plan.Name)
	d.Set("description", plan.Description)
	d.Set("appliance_type", plan.ApplianceType)
	d.Set("version", plan.Version)
	d.Set("flavor", plan.Flavor)
	d.Set("number_of_interfaces", plan.NumberOfInterfaces)
	d.Set("enabled", plan.Enabled)
	d.Set("max_number_of_aap", plan.MaxNumberOfAap)

	licenses := make([]map[string]string, len(plan.Licenses))
	for i, l := range plan.Licenses {
		license := map[string]string{"license_type": l.LicenseType}
		licenses[i] = license
	}
	d.Set("licenses", licenses)

	availability_zones := make([]map[string]interface{}, len(plan.AvailabilityZones))
	for i, az := range plan.AvailabilityZones {
		availability_zone := make(map[string]interface{})
		availability_zone["availability_zone"] = az.AvailabilityZone
		availability_zone["available"] = az.Available
		availability_zone["rank"] = az.Rank
		availability_zones[i] = availability_zone
	}
	d.Set("availability_zones", availability_zones)

	return nil
}

func getVirtualNetworkAppliancePlans(vnaClient *eclcloud.ServiceClient, listOpts appliance_plans.ListOpts) ([]appliance_plans.VirtualNetworkAppliancePlan, error) {
	pages, err := appliance_plans.List(vnaClient, listOpts).AllPages()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Load Balancer Plans: %w", err)
	}

	allPlans, err := appliance_plans.ExtractVirtualNetworkAppliancePlans(pages)
	if err != nil {
		return nil, fmt.Errorf("unable to extract retrieved Load Balancer Plans: %w", err)
	}

	if len(allPlans) < 1 {
		return nil, fmt.Errorf("specified Load Balancer Plan query returned no results")
	}
	return allPlans, nil
}
