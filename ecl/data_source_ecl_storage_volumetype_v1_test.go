package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// Skip test for "File Storage Standard" service type.
// Because in some region , that service type is not available.

func TestAccStorageV1VolumeTypeDataSourceBlockStorageBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeTypeOfBlockStorageByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeTypeDataSourceID("data.ecl_storage_volumetype_v1.vt"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volumetype_v1.vt", "name", VOLUME_TYPE_NAME_BLOCK),
				),
			},
		},
	})
}
func TestAccStorageV1VolumeTypeDataSourceBlockStorageByID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckStorage(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeTypeOfBlockStorageByID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeTypeDataSourceID("data.ecl_storage_volumetype_v1.vt"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volumetype_v1.vt", "name", VOLUME_TYPE_NAME_BLOCK),
				),
			},
		},
	})
}
func TestAccStorageV1VolumeTypeDataSourceFileStoragePremiumBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFileStorageServiceType(t, true, false) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeTypeOfFileStoragePremiumByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeTypeDataSourceID("data.ecl_storage_volumetype_v1.vt"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volumetype_v1.vt", "name", VOLUME_TYPE_NAME_FILE_PREMIUM),
				),
			},
		},
	})
}

func testAccCheckStorageV1VolumeTypeDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find volume type data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VolumeType data source ID not set")
		}

		return nil
	}
}

var testAccStorageV1VolumeTypeOfBlockStorageByName = fmt.Sprintf(`
data "ecl_storage_volumetype_v1" "vt" {
  name = "%s"
}
`, VOLUME_TYPE_NAME_BLOCK)

var testAccStorageV1VolumeTypeOfBlockStorageByID = fmt.Sprintf(`
data "ecl_storage_volumetype_v1" "vt" {
  volume_type_id = "%s"
}
`, OS_STORAGE_VOLUME_TYPE_ID)

var testAccStorageV1VolumeTypeOfFileStoragePremiumByName = fmt.Sprintf(`
data "ecl_storage_volumetype_v1" "vt" {
  name = "%s"
}
`, VOLUME_TYPE_NAME_FILE_PREMIUM)
