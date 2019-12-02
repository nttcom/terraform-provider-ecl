package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkV2PublicIPImportBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resourceName := "ecl_network_public_ip_v2.public_ip_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckPublicIP(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PublicIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PublicIPBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
