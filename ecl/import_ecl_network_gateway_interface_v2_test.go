package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkV2GatewayInterfaceImportBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resourceName := "ecl_network_gateway_interface_v2.gateway_interface_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckGatewayInterface(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2GatewayInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
