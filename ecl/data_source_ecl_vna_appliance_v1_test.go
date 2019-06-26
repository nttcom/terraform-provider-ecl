package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccVNAV1ApplianceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckVNA(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceDataSourceID("data.ecl_vna_appliance_v1.appliance_1"),
					// resource.TestCheckResourceAttr(
					// 	"data.ecl_network_common_function_gateway_v2.common_function_gateway_1", "name", "tf_common_function_gatewy"),
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

		return nil
	}
}

var testAccVNAV1ApplianceDataSourceBasic = fmt.Sprintf(`
data "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
}
`)
