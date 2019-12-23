package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDedicatedHypervisorV1LicenseImport_basic(t *testing.T) {
	resourceName := "ecl_dedicated_hypervisor_license_v1.license_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDedicatedHypervisorV1LicenseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDedicatedHypervisorV1LicenseBasic,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}
