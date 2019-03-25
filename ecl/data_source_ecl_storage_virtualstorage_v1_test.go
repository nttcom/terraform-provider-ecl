package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccStorageV1VirtualStorageDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckStorage(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageBasic,
			},
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageDataSourceByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageDataSourceID("data.ecl_storage_virtualstorage_v1.virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_id", OS_STORAGE_VOLUME_TYPE_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.10"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.20"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
				),
			},
		},
	})
}
func TestAccStorageV1VirtualStorageDataSourceID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckStorage(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageBasic,
			},
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageDataSourceByID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageDataSourceID("data.ecl_storage_virtualstorage_v1.virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_id", OS_STORAGE_VOLUME_TYPE_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.10"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.20"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
				),
			},
		},
	})
}

func testAccCheckStorageV1VirtualStorageDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find virtual storage data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VirtualStorage data source ID not set")
		}

		return nil
	}
}

var testAccStorageV1VirtualStorageDataSourceByName = fmt.Sprintf(`
%s

data "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "${ecl_storage_virtualstorage_v1.virtualstorage_1.name}"
}
`, testAccStorageV1VirtualStorageBasic)

var testAccStorageV1VirtualStorageDataSourceByID = fmt.Sprintf(`
%s

data "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
}
`, testAccStorageV1VirtualStorageBasic)
