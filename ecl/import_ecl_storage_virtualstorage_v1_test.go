package ecl

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccStorageV1VirtualStorageImport_basic(t *testing.T) {
	resourceName := "ecl_storage_virtualstorage_v1.virtualstorage_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorage(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VirtualStorageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
