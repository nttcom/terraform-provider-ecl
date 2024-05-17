package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/v3/ecl/vna/v1/appliances"
)

func allowedAddessPairsSchemaForDatasource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},

				"mac_address": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},

				"type": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},

				"vrid": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func fixedIPsSchemaForDatasource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"ip_address": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"subnet_id": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func interfaceInfoSchemaForDatasource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"description": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"network_id": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"updatable": &schema.Schema{
					Type:     schema.TypeBool,
					Computed: true,
				},
				"tags": &schema.Schema{
					Type:     schema.TypeMap,
					Computed: true,
				},
			},
		},
	}
}
func dataSourceVNAApplianceV1() *schema.Resource {
	var result *schema.Resource
	result = &schema.Resource{
		Read: dataSourceVNAApplianceV1Read,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"virtual_network_appliance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"appliance_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"os_monitoring_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"os_login_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"operation_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"virtual_network_appliance_plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},

			"interface_1_info":                  interfaceInfoSchemaForDatasource(),
			"interface_1_fixed_ips":             fixedIPsSchemaForDatasource(),
			"interface_1_allowed_address_pairs": allowedAddessPairsSchemaForDatasource(),

			"interface_2_info":                  interfaceInfoSchemaForDatasource(),
			"interface_2_fixed_ips":             fixedIPsSchemaForDatasource(),
			"interface_2_allowed_address_pairs": allowedAddessPairsSchemaForDatasource(),

			"interface_3_info":                  interfaceInfoSchemaForDatasource(),
			"interface_3_fixed_ips":             fixedIPsSchemaForDatasource(),
			"interface_3_allowed_address_pairs": allowedAddessPairsSchemaForDatasource(),

			"interface_4_info":                  interfaceInfoSchemaForDatasource(),
			"interface_4_fixed_ips":             fixedIPsSchemaForDatasource(),
			"interface_4_allowed_address_pairs": allowedAddessPairsSchemaForDatasource(),

			"interface_5_info":                  interfaceInfoSchemaForDatasource(),
			"interface_5_fixed_ips":             fixedIPsSchemaForDatasource(),
			"interface_5_allowed_address_pairs": allowedAddessPairsSchemaForDatasource(),

			"interface_6_info":                  interfaceInfoSchemaForDatasource(),
			"interface_6_fixed_ips":             fixedIPsSchemaForDatasource(),
			"interface_6_allowed_address_pairs": allowedAddessPairsSchemaForDatasource(),

			"interface_7_info":                  interfaceInfoSchemaForDatasource(),
			"interface_7_fixed_ips":             fixedIPsSchemaForDatasource(),
			"interface_7_allowed_address_pairs": allowedAddessPairsSchemaForDatasource(),

			"interface_8_info":                  interfaceInfoSchemaForDatasource(),
			"interface_8_fixed_ips":             fixedIPsSchemaForDatasource(),
			"interface_8_allowed_address_pairs": allowedAddessPairsSchemaForDatasource(),
		},
	}

	for i := 1; i <= maxNumberOfInterfaces; i++ {
		result.Schema[fmt.Sprintf("interface_%d_info", i)] = interfaceInfoSchemaForDatasource()
		result.Schema[fmt.Sprintf("interface_%d_fixed_ips", i)] = fixedIPsSchemaForDatasource()
		result.Schema[fmt.Sprintf("interface_%d_allowed_address_pairs", i)] = allowedAddessPairsSchemaForDatasource()
	}

	return result
}

func dataSourceVNAApplianceV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	listOpts := appliances.ListOpts{}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	if v, ok := d.GetOk("virtual_network_appliance_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("appliance_type"); ok {
		listOpts.ApplianceType = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		listOpts.Description = v.(string)
	}

	if v, ok := d.GetOk("availability_zone"); ok {
		listOpts.AvailabilityZone = v.(string)
	}

	if v, ok := d.GetOk("os_monitoring_status"); ok {
		listOpts.OSMonitoringStatus = v.(string)
	}

	if v, ok := d.GetOk("os_login_status"); ok {
		listOpts.OSLoginStatus = v.(string)
	}

	if v, ok := d.GetOk("vm_status"); ok {
		listOpts.VMStatus = v.(string)
	}

	if v, ok := d.GetOk("operation_status"); ok {
		listOpts.OperationStatus = v.(string)
	}

	if v, ok := d.GetOk("virtual_network_appliance_plan_id"); ok {
		listOpts.VirtualNetworkAppliancePlanID = v.(string)
	}
	if v, ok := d.GetOk("tenant_id"); ok {
		listOpts.TenantID = v.(string)
	}

	vnaClient, err := config.vnaV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
	}

	pages, err := appliances.List(vnaClient, listOpts).AllPages()
	if err != nil {
		return err
	}

	allAppliances, err := appliances.ExtractAppliances(pages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve virtual network appliances: %s", err)
	}

	if len(allAppliances) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(allAppliances) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	vna := allAppliances[0]

	log.Printf("[DEBUG] Retrieved Virtual Network Appliance %s: %+v", d.Id(), vna)

	d.SetId(vna.ID)

	d.Set("name", vna.Name)
	d.Set("description", vna.Description)
	d.Set("default_gateway", vna.DefaultGateway)
	d.Set("availability_zone", vna.AvailabilityZone)
	d.Set("virtual_network_appliance_plan_id", vna.AppliancePlanID)
	d.Set("tenant_id", vna.TenantID)
	d.Set("tags", vna.Tags)

	for i := 1; i <= maxNumberOfInterfaces; i++ {
		targetMeta := getInterfaceBySlotNumber(&vna, i)
		targetFIPs := getFixedIPsBySlotNumber(&vna, i)
		targetAAPs := getAllowedAddressPairsBySlotNumber(&vna, i)

		d.Set(
			fmt.Sprintf("interface_%d_info", i),
			getInterfaceInfoAsState(targetMeta))

		d.Set(
			fmt.Sprintf("interface_%d_fixed_ips", i),
			getInterfaceFixedIPsAsState(targetFIPs))

		d.Set(
			fmt.Sprintf("interface_%d_allowed_address_pairs", i),
			getInterfaceAllowedAddressPairsAsState(targetAAPs))
	}

	return nil
}
