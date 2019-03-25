package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeVolumeV2AttachImportBasic(t *testing.T) {
	resourceName := "ecl_compute_volume_attach_v2.va_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2AttachDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2AttachBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
