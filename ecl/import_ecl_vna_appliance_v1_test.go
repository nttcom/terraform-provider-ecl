package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccVNAV1ApplianceImportBasic(t *testing.T) {
	resourceName := "ecl_vna_appliance_v1.appliance_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
