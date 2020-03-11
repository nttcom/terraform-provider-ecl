package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkV2LoadBalancerImport_basic(t *testing.T) {
	resourceName := "ecl_network_load_balancer_v2.lb_test1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2LoadBalancerBasic,
			},

			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"admin_password", "user_password"},
			},
		},
	})
}
