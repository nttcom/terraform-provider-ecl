package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkV2CommonFunctionGatewayImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resourceName := "ecl_network_common_function_gateway_v2.common_function_gateway_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCommonFunctionGateway(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2CommonFunctionGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2CommonFunctionGatewayBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
