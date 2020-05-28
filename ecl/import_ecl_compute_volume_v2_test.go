package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeVolumeV2VolumeImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resourceName := "ecl_compute_volume_v2.volume_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
