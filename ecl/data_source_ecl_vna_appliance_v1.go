package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func dataSourceVNAApplianceV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNAApplianceV1Read,

		Schema: map[string]*schema.Schema{

			"virtual_network_appliance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"default_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"virtual_network_appliance_plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"interface_1_meta":                  interfaceMetaSchema(),
			"interface_1_fixed_ips":             fixedIPsSchema(),
			"interface_1_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_2_meta":                  interfaceMetaSchema(),
			"interface_2_fixed_ips":             fixedIPsSchema(),
			"interface_2_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_3_meta":                  interfaceMetaSchema(),
			"interface_3_fixed_ips":             fixedIPsSchema(),
			"interface_3_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_4_meta":                  interfaceMetaSchema(),
			"interface_4_fixed_ips":             fixedIPsSchema(),
			"interface_4_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_5_meta":                  interfaceMetaSchema(),
			"interface_5_fixed_ips":             fixedIPsSchema(),
			"interface_5_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_6_meta":                  interfaceMetaSchema(),
			"interface_6_fixed_ips":             fixedIPsSchema(),
			"interface_6_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_7_meta":                  interfaceMetaSchema(),
			"interface_7_fixed_ips":             fixedIPsSchema(),
			"interface_7_allowed_address_pairs": allowedAddessPairsSchema(),

			"interface_8_meta":                  interfaceMetaSchema(),
			"interface_8_fixed_ips":             fixedIPsSchema(),
			"interface_8_allowed_address_pairs": allowedAddessPairsSchema(),
		},
	}
}

func dataSourceVNAApplianceV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	listOpts := appliances.ListOpts{}

	if v, ok := d.GetOk("virtual_network_appliance_id"); ok {
		listOpts.ID = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		listOpts.Name = v.(string)
	}

	// TODO maintenance followings later
	// if v, ok := d.GetOk("appliance_type"); ok {
	// 	listOpts.Name = v.(string)
	// }

	// if v, ok := d.GetOk("description"); ok {
	// 	listOpts.Description = v.(string)
	// }

	// if v, ok := d.GetOk("id"); ok {
	// 	listOpts.ID = v.(string)
	// }

	// if v, ok := d.GetOk("network_id"); ok {
	// 	listOpts.NetworkID = v.(string)
	// }

	// if v, ok := d.GetOk("status"); ok {
	// 	listOpts.Status = v.(string)
	// }

	// if v, ok := d.GetOk("subnet_id"); ok {
	// 	listOpts.SubnetID = v.(string)
	// }

	// if v, ok := d.GetOk("tenant_id"); ok {
	// 	listOpts.TenantID = v.(string)
	// }

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

	log.Printf("[MYDEBUG] VNA: %#v", vna)
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
		var targetMeta appliances.InterfaceInResponse
		var targetFIPs []appliances.FixedIPInResponse
		var targetAAPs []appliances.AllowedAddressPairInResponse

		switch i {
		case 1:
			targetMeta = vna.Interfaces.Interface1
			targetFIPs = vna.Interfaces.Interface1.FixedIPs
			targetAAPs = vna.Interfaces.Interface1.AllowedAddressPairs
			break
		case 2:
			targetMeta = vna.Interfaces.Interface2
			targetFIPs = vna.Interfaces.Interface2.FixedIPs
			targetAAPs = vna.Interfaces.Interface2.AllowedAddressPairs
			break
		case 3:
			targetMeta = vna.Interfaces.Interface3
			targetFIPs = vna.Interfaces.Interface3.FixedIPs
			targetAAPs = vna.Interfaces.Interface3.AllowedAddressPairs
			break
		case 4:
			targetMeta = vna.Interfaces.Interface4
			targetFIPs = vna.Interfaces.Interface4.FixedIPs
			targetAAPs = vna.Interfaces.Interface4.AllowedAddressPairs
			break
		case 5:
			targetMeta = vna.Interfaces.Interface5
			targetFIPs = vna.Interfaces.Interface5.FixedIPs
			targetAAPs = vna.Interfaces.Interface5.AllowedAddressPairs
			break
		case 6:
			targetMeta = vna.Interfaces.Interface6
			targetFIPs = vna.Interfaces.Interface6.FixedIPs
			targetAAPs = vna.Interfaces.Interface6.AllowedAddressPairs
			break
		case 7:
			targetMeta = vna.Interfaces.Interface7
			targetFIPs = vna.Interfaces.Interface7.FixedIPs
			targetAAPs = vna.Interfaces.Interface7.AllowedAddressPairs
			break
		case 8:
			targetMeta = vna.Interfaces.Interface8
			targetFIPs = vna.Interfaces.Interface8.FixedIPs
			targetAAPs = vna.Interfaces.Interface8.AllowedAddressPairs
			break
		default:
			break
		}

		d.Set(
			fmt.Sprintf("interface_%d_meta", i),
			getInterfaceMetaAsState(targetMeta))

		d.Set(
			fmt.Sprintf("interface_%d_fixed_ips", i),
			getInterfaceFixedIPsAsState(targetFIPs))

		d.Set(
			fmt.Sprintf("interface_%d_allowed_address_pairs", i),
			getInterfaceAllowedAddressPairsAsState(targetAAPs))
	}

	return nil
}
