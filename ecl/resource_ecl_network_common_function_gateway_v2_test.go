package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/common_function_gateways"
)

func TestAccNetworkV2CommonFunctionGatewayBasic(t *testing.T) {
	var cfGw common_function_gateways.CommonFunctionGateway

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCommonFunctionGateway(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2CommonFunctionGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2CommonFunctionGatewayBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2CommonFunctionGatewayExists(
						"ecl_network_common_function_gateway_v2.common_function_gateway_1",
						&cfGw),
					resource.TestCheckResourceAttr(
						"ecl_network_common_function_gateway_v2.common_function_gateway_1",
						"name", "terraform-common-function-gateway1"),
					resource.TestCheckResourceAttr(
						"ecl_network_common_function_gateway_v2.common_function_gateway_1",
						"description", "terraform-common-function-gateway1-description"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2CommonFunctionGatewayUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_common_function_gateway_v2.common_function_gateway_1",
						"name", repeatedString("a", 255)),
					resource.TestCheckResourceAttr(
						"ecl_network_common_function_gateway_v2.common_function_gateway_1",
						"description", ""),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2CommonFunctionGatewayUpdate2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_common_function_gateway_v2.common_function_gateway_1",
						"name", ""),
					resource.TestCheckResourceAttr(
						"ecl_network_common_function_gateway_v2.common_function_gateway_1",
						"description", repeatedString("a", 255)),
				),
			},
		},
	})
}

func testAccCheckNetworkV2CommonFunctionGatewayExists(n string, cfGw *common_function_gateways.CommonFunctionGateway) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL network client: %s", err)
		}

		found, err := common_function_gateways.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Common Function Gateway not found")
		}

		*cfGw = *found

		return nil
	}
}

func testAccCheckNetworkV2CommonFunctionGatewayDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_common_function_gateway_v2" {
			continue
		}
		_, err := common_function_gateways.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Common Function Gateway still exists")
		}

	}

	return nil
}

var testAccNetworkV2CommonFunctionGatewayBasic = fmt.Sprintf(`
resource "ecl_network_common_function_gateway_v2" "common_function_gateway_1" {
  name = "terraform-common-function-gateway1"
  description = "terraform-common-function-gateway1-description"
  common_function_pool_id = "%s"
}
`, OS_COMMON_FUNCTION_POOL_ID)

var testAccNetworkV2CommonFunctionGatewayUpdate = fmt.Sprintf(`
resource "ecl_network_common_function_gateway_v2" "common_function_gateway_1" {
  name = "%s"
  description = ""
  common_function_pool_id = "%s"
}
`, repeatedString("a", 255),
	OS_COMMON_FUNCTION_POOL_ID)

var testAccNetworkV2CommonFunctionGatewayUpdate2 = fmt.Sprintf(`
resource "ecl_network_common_function_gateway_v2" "common_function_gateway_1" {
  name = ""
  description = "%s"
  common_function_pool_id = "%s"
}
`, repeatedString("a", 255),
	OS_COMMON_FUNCTION_POOL_ID)
