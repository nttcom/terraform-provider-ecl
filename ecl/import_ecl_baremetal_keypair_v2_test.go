package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccBaremetalV2KeypairImport_basic(t *testing.T) {
	resourceName := "ecl_baremetal_keypair_v2.kp_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBaremetalV2KeypairDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBaremetalV2KeypairBasic,
			},

			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"private_key"},
			},
		},
	})
}
