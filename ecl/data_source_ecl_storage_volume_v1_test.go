package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// TestAccStorageV1VolumeDataSource_basic is basic test of volume
// 	(Note) you need to prepare Virtual Storage and set the ID of that as Env value.
//	This test function does followings
// 		1. Create volume under pre created virtual storage
// 		2. Refer that volume as datasource by using volume name to find it
func TestAccStorageV1VolumeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeBaseResourceForDataSource,
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeDataSourceByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeDataSourceID("data.ecl_storage_volume_v1.volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "size", "100"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
				),
			},
		},
	})
}

// TestAccStorageV1VolumeDataSource_id is basic test of volume
// 	(Note) you need to prepare Virtual Storage and set the ID of that as Env value.
//	This test function does followings
// 		1. Create volume under pre created virtual storage
// 		2. Refer that volume as datasource by using volume ID to find it
func TestAccStorageV1VolumeDataSource_id(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeBaseResourceForDataSource,
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeDataSourceByID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeDataSourceID("data.ecl_storage_volume_v1.volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "size", "100"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
				),
			},
		},
	})
}

// TestAccStorageV1VolumeDataSource_createAllResourcesAsBlock create all of relevant resources
// about data source test as Block Storage service
// 	This test function does followings
// 		1. Create network virtual storage, volume prior to data source
// 		2. Refer that volume as datasource by using volume ID to find it
func TestAccStorageV1VolumeDataSource_createAllResourcesAsBlock(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolumeBasic,
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeDataSourceCreateAllResourcesAsBlock,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeDataSourceID("data.ecl_storage_volume_v1.volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "size", "100"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
				),
			},
		},
	})
}

// TestAccStorageV1VolumeDataSource_createAllResourcesAsFilePremium create all of relevant resources
// about data source test as File Storage Premium service
// 	This test function does followings
// 		1. Create network virtual storage, volume prior to data source
// 		2. Refer that volume as datasource by using volume ID to find it
func TestAccStorageV1VolumeDataSource_createAllResourcesAsFilePremium(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckFileStorageServiceType(t, true, false)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolumeBasic,
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeDataSourceCreateAllResourcesAsFilePremium,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeDataSourceID("data.ecl_storage_volume_v1.volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "size", "256"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "throughput", "50"),
				),
			},
		},
	})
}
func TestAccStorageV1VolumeDataSource_createAllResourcesAsFileStandard(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckFileStorageServiceType(t, false, true)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolumeBasic,
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeDataSourceCreateAllResourcesAsFileStandard,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeDataSourceID("data.ecl_storage_volume_v1.volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_storage_volume_v1.volume_1", "size", "1024"),
				),
			},
		},
	})
}

func testAccCheckStorageV1VolumeDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find volume data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Volume data source ID not set")
		}

		return nil
	}
}

var testAccStorageV1VolumeBaseResourceForDataSource = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
  name = "volume_1"
  description = "first test volume"
  virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
  iops_per_gb = "2"
  size = 100
  initiator_iqns = ["%s", "%s"]
}  
`, testAccStorageV1VirtualStorageBasic, IQN01, IQN02)

var testAccStorageV1VolumeDataSourceByName = fmt.Sprintf(`
%s

data "ecl_storage_volume_v1" "volume_1" {
  name = "${ecl_storage_volume_v1.volume_1.name}"
}
`, testAccStorageV1VolumeBaseResourceForDataSource)

var testAccStorageV1VolumeDataSourceByID = fmt.Sprintf(`
%s

data "ecl_storage_volume_v1" "volume_1" {
  volume_id = "${ecl_storage_volume_v1.volume_1.id}"
}
`, testAccStorageV1VolumeBaseResourceForDataSource)

var testAccStorageV1VolumeDataSourceCreateAllResourcesAsBlock = fmt.Sprintf(`
%s

data "ecl_storage_volume_v1" "volume_1" {
	name = "${ecl_storage_volume_v1.volume_1.name}"
}`,
	testAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolumeBasic,
)

var testAccStorageV1VolumeDataSourceCreateAllResourcesAsFilePremium = fmt.Sprintf(`
%s

data "ecl_storage_volume_v1" "volume_1" {
	name = "${ecl_storage_volume_v1.volume_1.name}"
}`,
	testAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolumeBasic,
)

var testAccStorageV1VolumeDataSourceCreateAllResourcesAsFileStandard = fmt.Sprintf(`
%s

data "ecl_storage_volume_v1" "volume_1" {
	name = "${ecl_storage_volume_v1.volume_1.name}"
}`,
	testAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolumeBasic,
)
