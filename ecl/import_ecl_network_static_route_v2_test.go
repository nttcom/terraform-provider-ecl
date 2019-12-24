package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkV2StaticRouteImport_basic(t *testing.T) {
	resourceName := "ecl_network_static_route_v2.static_route_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStaticRoute(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2StaticRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
