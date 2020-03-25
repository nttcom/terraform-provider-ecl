package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/computevolume/v2/volumes"
)

func TestAccComputeVolumeV2Volume_basic(t *testing.T) {
	var volume volumes.Volume

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageV2VolumeExists("ecl_compute_volume_v2.volume_1", &volume),
					testAccCheckComputeVolumeV2VolumeMetadata(&volume, "foo", "bar"),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "description", "volume description"),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "size", "15"),
				),
			},
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageV2VolumeExists("ecl_compute_volume_v2.volume_1", &volume),
					testAccCheckComputeVolumeV2VolumeMetadata(&volume, "foo", "bar"),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "name", ""),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "description", "volume description-updated"),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "size", "15"),
				),
			},
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageV2VolumeExists("ecl_compute_volume_v2.volume_1", &volume),
					testAccCheckComputeVolumeV2VolumeMetadata(&volume, "foo", "bar"),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "name", "volume_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "description", ""),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "size", "15"),
				),
			},
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeUpdate3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageV2VolumeExists("ecl_compute_volume_v2.volume_1", &volume),
					testAccCheckComputeVolumeV2VolumeMetadataIsBlankMap(&volume),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "name", "volume_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "description", ""),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "size", "40"),
				),
			},
		},
	})
}

func TestAccComputeVolumeV2Volume_fromImageByName(t *testing.T) {
	var volume volumes.Volume

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeFromImageByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageV2VolumeExists("ecl_compute_volume_v2.volume_1", &volume),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "name", "volume_1"),
				),
			},
		},
	})
}

func TestAccComputeVolumeV2Volume_fromImageByID(t *testing.T) {
	var volume volumes.Volume

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeFromImageByID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageV2VolumeExists("ecl_compute_volume_v2.volume_1", &volume),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "name", "volume_1"),
				),
			},
		},
	})
}
func TestAccComputeVolumeV2Volume_fromImageByNameAndID(t *testing.T) {
	var volume volumes.Volume

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeFromImageByNameAndID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageV2VolumeExists("ecl_compute_volume_v2.volume_1", &volume),
					resource.TestCheckResourceAttr(
						"ecl_compute_volume_v2.volume_1", "name", "volume_1"),
				),
			},
		},
	})
}

func TestAccComputeVolumeV2Volume_timeout(t *testing.T) {
	var volume volumes.Volume

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2VolumeTimeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBlockStorageV2VolumeExists("ecl_compute_volume_v2.volume_1", &volume),
				),
			},
		},
	})
}

func testAccCheckComputeVolumeV2VolumeDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	computeVolumeClient, err := config.computeVolumeV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_compute_volume_v2" {
			continue
		}

		_, err := volumes.Get(computeVolumeClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Volume still exists")
		}
	}

	return nil
}

func testAccCheckBlockStorageV2VolumeExists(n string, volume *volumes.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		computeVolumeClient, err := config.computeVolumeV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL compute volume client: %s", err)
		}

		found, err := volumes.Get(computeVolumeClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Volume not found")
		}

		*volume = *found

		return nil
	}
}

func testAccCheckComputeVolumeV2VolumeDoesNotExist(t *testing.T, n string, volume *volumes.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		computeVolumeClient, err := config.computeVolumeV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL compute volume client: %s", err)
		}

		_, err = volumes.Get(computeVolumeClient, volume.ID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return nil
			}
			return err
		}

		return fmt.Errorf("Volume still exists")
	}
}

func testAccCheckComputeVolumeV2VolumeMetadataIsBlankMap(volume *volumes.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(volume.Metadata) != 0 {
			return fmt.Errorf("Not blank metadata")
		}
		return nil
	}
}

func testAccCheckComputeVolumeV2VolumeMetadata(
	volume *volumes.Volume, k string, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if volume.Metadata == nil {
			return fmt.Errorf("No metadata")
		}

		for key, value := range volume.Metadata {
			if k != key {
				continue
			}

			if v == value {
				return nil
			}

			return fmt.Errorf("Bad value for %s: %s", k, value)
		}

		return fmt.Errorf("Metadata not found: %s", k)
	}
}

const testAccComputeVolumeV2VolumeBasic = `
resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1"
  description = "volume description"
  metadata = {
    foo = "bar"
  }
  size = 15
}
`

const testAccComputeVolumeV2VolumeUpdate = `
resource "ecl_compute_volume_v2" "volume_1" {
  name = ""
  description = "volume description-updated"
  metadata = {
    foo = "bar"
  }
  size = 15
}
`
const testAccComputeVolumeV2VolumeUpdate2 = `
resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1-updated"
  description = ""
  metadata = {
    foo = "bar"
  }
  size = 15
}
`
const testAccComputeVolumeV2VolumeUpdate3 = `
resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1-updated"
  description = ""
  size = 40
}
`

const testAccComputeVolumeV2VolumeFromImageByName = `
resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1"
  size = 15
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
}
`

const testAccComputeVolumeV2VolumeFromImageByID = `
resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1"
  size = 15
  image_id = "e74134b7-f20a-40ee-918b-66503b2f13be"
}
`

const testAccComputeVolumeV2VolumeFromImageByNameAndID = `
resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1"
  size = 15
  image_id = "e74134b7-f20a-40ee-918b-66503b2f13be"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
}
`

const testAccComputeVolumeV2VolumeTimeout = `
resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1"
  description = "first test volume"
  size = 15

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`
