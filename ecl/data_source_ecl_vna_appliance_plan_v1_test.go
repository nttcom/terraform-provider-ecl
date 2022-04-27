package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccVNAV1AppliancePlanDataSourcePlanName = "vSRX_20.4R2_2CPU_4GB_8IF_STD"

func TestAccVNAV1AppliancePlanDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVNAV1AppliancePlanDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AppliancePlanDataSourceID("data.ecl_vna_appliance_plan_v1.appliance_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "name", testAccVNAV1AppliancePlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "description", testAccVNAV1AppliancePlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "appliance_type", "ECL::VirtualNetworkAppliance::VSRX"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "flavor", "2CPU-4GB"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "licenses.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "licenses.0.license_type", "STD"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "availability_zones.#", "3"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "availability_zones.0.availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "availability_zones.0.available", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "availability_zones.0.rank", "1"),
				),
			},
		},
	})
}

func TestAccVNAV1AppliancePlanDataSource_queries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVNAV1AppliancePlanDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AppliancePlanDataSourceID("data.ecl_vna_appliance_plan_v1.appliance_plan_1"),
				),
			},
			{
				Config: testAccVNAV1AppliancePlanDataSourceQueryID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AppliancePlanDataSourceID("data.ecl_vna_appliance_plan_v1.appliance_plan_2"),
				),
			},
		},
	})
}

func testAccCheckVNAV1AppliancePlanDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("can't find vna_appliance_plan data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("vna plan data source ID not set")
		}

		return nil
	}
}

var testAccVNAV1AppliancePlanDataSourceBasic = fmt.Sprintf(`
data "ecl_vna_appliance_plan_v1" "appliance_plan_1" {
  name = %q
  description = %q
  appliance_type = "ECL::VirtualNetworkAppliance::VSRX"
  flavor = "2CPU-8GB"
  enabled = true
  licenses {
    license_type = "STD"
  }
  availability_zones {
    availability_zone = "zone1_groupa"
    available = true
    rank = 1
  }
}
`, testAccVNAV1AppliancePlanDataSourcePlanName, testAccVNAV1AppliancePlanDataSourcePlanName,
)

var testAccVNAV1AppliancePlanDataSourceQueryID = fmt.Sprintf(`
data "ecl_vna_appliance_plan_v1" "appliance_plan_1" {
  name = %q
}
data "ecl_vna_appliance_plan_v1" "appliance_plan_2" {
  id = "${data.ecl_network_appliance_plan_v1.appliance_plan_1.id}"
}
`, testAccVNAV1AppliancePlanDataSourcePlanName,
)

var testAccVNAV1AppliancePlanDataSourceQueryName = fmt.Sprintf(`
data "ecl_vna_appliance_plan_v1" "appliance_plan_1" {
  name = %q
}
`, testAccVNAV1AppliancePlanDataSourcePlanName,
)

var testAccVNAV1AppliancePlanDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_vna_appliance_plan_v1" "appliance_plan_1" {
  description = %q
}
`, testAccVNAV1AppliancePlanDataSourcePlanName,
)
