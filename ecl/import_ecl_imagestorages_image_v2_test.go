package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccImageStoragesV2ImageImport_basic(t *testing.T) {
	resourceName := "ecl_imagestorages_image_v2.image_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2ImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2ImageBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"region",
					"local_file_path",
					// "image_cache_path",
					// "image_source_url",
					"verify_checksum",
				},
			},
		},
	})
}
