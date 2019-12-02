package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2CommonFunctionGatewayDataSourceBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommonFunctionGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2CommonFunctionGatewayDataSourceCommonFunctionGateway,
			},
			resource.TestStep{
				Config: testAccNetworkV2CommonFunctionGatewayDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2CommonFunctionGatewayDataSourceID("data.ecl_network_common_function_gateway_v2.common_function_gateway_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_gateway_v2.common_function_gateway_1", "name", "tf_common_function_gatewy"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2CommonFunctionGatewayDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find common function gateway data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Common Function Gateway data source ID not set")
		}

		return nil
	}
}

var testAccNetworkV2CommonFunctionGatewayDataSourceCommonFunctionGateway = fmt.Sprintf(`
resource "ecl_network_common_function_gateway_v2" "common_function_gateway_1" {
        name = "tf_common_function_gatewy"
		description = "test description"
		common_function_pool_id = "%s"
}`, OS_COMMON_FUNCTION_POOL_ID)

var testAccNetworkV2CommonFunctionGatewayDataSourceBasic = fmt.Sprintf(`
%s

data "ecl_network_common_function_gateway_v2" "common_function_gateway_1" {
	name = "${ecl_network_common_function_gateway_v2.common_function_gateway_1.name}"
}
`, testAccNetworkV2CommonFunctionGatewayDataSourceCommonFunctionGateway)
