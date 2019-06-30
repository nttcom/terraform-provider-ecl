package ecl

import (
	"fmt"
	// "log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	// "github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func TestAccVNAV1ApplianceDataSourceBasic(t *testing.T) {
	// var vna appliances.Appliance
	// var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckVNA(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceDataSourceID("data.ecl_vna_appliance_v1.appliance_1"),
					// testAccCheckVNAV1ApplianceExists("data.ecl_vna_appliance_v1.appliance_1", &vna),
					// testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					// Check about meta
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_v1.appliance_1", "interface_1_meta.0.name", "interface_1"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_v1.appliance_1", "interface_1_meta.0.description", "interface_1_description"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
					// Check about interface
					// testAccCheckVNAV1FixedIP(
					// testAccCheckVNAV1FixedIP(vna.Interfaces.Interface1.FixedIPs, 0, "192.168.1.50"),
					// testAccCheckVNAV1FixedIP(&vna.Interfaces.Interface1.FixedIPs[0], "192.168.1.50"),
				),
			},
		},
	})
}

func testAccCheckVNAV1ApplianceDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find virtual network appliance data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Virtual Network Appliance data source ID not set")
		}

		// fmt.Printf("[MYDEBUG] datasource existence VNA: %#v", rs)
		return nil
	}
}

var testAccVNAV1ApplianceDataSourceBasic = fmt.Sprintf(`
data "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
}
`)
