package ecl

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/storage/v1/virtualstorages"
	"github.com/nttcom/eclcloud/ecl/storage/v1/volumes"
)

const IQN01 string = "iqn.2003-01.org.sample-iscsi.node1.x8664:sn.2613f8620d98"
const IQN02 string = "iqn.2003-01.org.sample-iscsi.node1.x8664:sn.2613f8620d99"

// TestAccStorageV1VolumeTimeout just check Timeout section is correctly worked or not
func TestAccStorageV1VolumeTimeout(t *testing.T) {
	var v volumes.Volume
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorageVolume(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeTimeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1"),
				),
			},
		},
	})
}

// TestAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolume is practical test case from user's stand point
// This function test followings
//		1. Create network and subnet which be connected from virtual storage
//		2. Create virtual storage by using environment values as Block Storage type
//		3. Create Volume under above virtual storage
//		4. Check if each parameters of volume are correctly set to created volume
//		5. Update it by modified HCL configurations
//		6. Check if new configuration is correctly applied
//		7. Check if new configuration is correctly applied
func TestAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolume(t *testing.T) {
	var v volumes.Volume
	var vs virtualstorages.VirtualStorage

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorageVolume(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolumeBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "100"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "initiator_iqns.0", IQN01),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolumeUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume updated"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "100"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "initiator_iqns.0", IQN01),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "initiator_iqns.1", IQN02),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolumeUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", ""),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "100"),
					testAccCheckStorageV1VolumeIQNLengthIsZERO(&v),
				),
			},
		},
	})
}

// TestAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolume is practical test case from user's stand point
// This test is almost same as TestAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolume
// Only difference between them is prior one is test for Block Storage type service,
// and later one is test for File Storage Premium type
//
// This function test followings
//		1. Create network and subnet which be connected from virtual storage
//		2. Create virtual storage by using environment values as FIle Storage Premium type
//		3. Create Volume under above virtual storage
//		4. Check if each parameters of volume are correctly set to created volume
//		5. Update it by modified HCL configurations
//		6. Check if new configuration is correctly applied
//		7. Repeat 5 and 6 by another configurations

func TestAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolume(t *testing.T) {
	var v volumes.Volume
	var vs virtualstorages.VirtualStorage

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckStorageVolume(t)
			testAccPreCheckFileStorageServiceType(t, true, false)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolumeBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "throughput", "50"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "256"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolumeUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume updated"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "throughput", "50"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "256"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolumeUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", ""),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "throughput", "50"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "256"),
				),
			},
		},
	})
}
func TestAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolume(t *testing.T) {
	var v volumes.Volume
	var vs virtualstorages.VirtualStorage

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckStorageVolume(t)
			testAccPreCheckFileStorageServiceType(t, false, true)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolumeBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "1024"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolumeUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume updated"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "1024"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolumeUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", ""),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "1024"),
				),
			},
		},
	})
}

func testCheckVolumeIDIsChanged(v1, v2 *volumes.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if v1.ID == v2.ID {
			log.Printf("[DEBUG] Volume (Comparison target) IDs: %s, %s", v1.ID, v2.ID)
			return fmt.Errorf("Resource ID is not changed. ForeceNew parameter does not work correctly")
		}
		return nil
	}
}

// TestAccStorageV1VolumeForceNewByIOPSPerGB checkes if ForceNew about volume works correctly
// This test changes iops_per_gb parameter by using 2 configs
// and check if resource is deleted and re-created
func TestAccStorageV1VolumeForceNewByIOPSPerGB(t *testing.T) {
	var v1, v2 volumes.Volume
	var vs virtualstorages.VirtualStorage

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorageVolume(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeForceNewByIOPSPerGBBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v1),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "100"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "initiator_iqns.0", IQN01),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeForceNewByIOPSPerGBUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v2),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					testCheckVolumeIDIsChanged(&v1, &v2),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-forcednew"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume forcednew"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "iops_per_gb", "4"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "100"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "initiator_iqns.0", IQN01),
				),
			},
		},
	})
}

// TestAccStorageV1VolumeForceNewBySize checkes if ForceNew about volume works correctly
// This test changes size parameter by using 2 configs
// and check if resource is deleted and re-created
func TestAccStorageV1VolumeForceNewBySize(t *testing.T) {
	var v1, v2 volumes.Volume
	var vs virtualstorages.VirtualStorage

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorageVolume(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeForceNewBySizeBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v1),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "100"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "initiator_iqns.0", IQN01),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeForceNewBySizeUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v2),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					testCheckVolumeIDIsChanged(&v1, &v2),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-forcednew"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume forcednew"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "iops_per_gb", "2"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "250"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "initiator_iqns.0", IQN01),
				),
			},
		},
	})
}

// TestAccStorageV1VolumeForceNewByThroughput checkes if ForceNew about volume works correctly
// This test changes throughput parameter by using 2 configs
// and check if resource is deleted and re-created
// [Note] This test CRUD all of resources, different from above two ForceNew test
func TestAccStorageV1VolumeForceNewByThroughput(t *testing.T) {
	var v1, v2 volumes.Volume
	var vs virtualstorages.VirtualStorage

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckStorageVolume(t)
			testAccPreCheckFileStorageServiceType(t, true, false)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VolumeForceNewByThroughputBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v1),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "throughput", "50"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "256"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VolumeForceNewByThroughputUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VolumeExists("ecl_storage_volume_v1.volume_1", &v2),
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					testCheckVolumeIDIsChanged(&v1, &v2),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "name", "volume_1-forcednew"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "description", "first test volume forcednew"),
					resource.TestCheckResourceAttrPtr(
						"ecl_storage_volume_v1.volume_1", "virtual_storage_id", &vs.ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "throughput", "100"),
					resource.TestCheckResourceAttr(
						"ecl_storage_volume_v1.volume_1", "size", "256"),
				),
			},
		},
	})
}

// TestValidateStorageV1VolumeCheckConflictsIOPSPerGBAndThroughput checks "ConflictsWith" parameter
// correctly works or not between iops_per_gb and throuthput
func TestValidateStorageV1VolumeCheckConflictsIOPSPerGBAndThroughput(t *testing.T) {
	pattern1 := "\"iops_per_gb\": conflicts with throughput"
	pattern2 := "\"throughput\": conflicts with iops_per_gb"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testValidateStorageV1VolumeConflictsIOPSPerGBAndThroughput,
				ExpectError: regexp.MustCompile(pattern1),
			},
			resource.TestStep{
				Config:      testValidateStorageV1VolumeConflictsIOPSPerGBAndThroughput,
				ExpectError: regexp.MustCompile(pattern2),
			},
		},
	})
}

// TestValidateStorageV1VolumeCheckConflictsInitiatorIQNsAndThroughput checks "ConflictsWith" parameter
// correctly works or not between initiator_iqns and throuthput
func TestValidateStorageV1VolumeCheckConflictsInitiatorIQNsAndThroughput(t *testing.T) {
	pattern1 := "\"initiator_iqns\": conflicts with throughput"
	pattern2 := "\"throughput\": conflicts with initiator_iqns"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testValidateStorageV1VolumeConflictsInitiatorIQNsAndThroughput,
				ExpectError: regexp.MustCompile(pattern1),
			},
			resource.TestStep{
				Config:      testValidateStorageV1VolumeConflictsInitiatorIQNsAndThroughput,
				ExpectError: regexp.MustCompile(pattern2),
			},
		},
	})
}

func testAccCheckStorageV1VolumeDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.storageV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_storage_volume_v1" {
			continue
		}

		_, err := volumes.Get(client, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Volume still exists")
		}
	}

	return nil
}

func testAccCheckStorageV1VolumeExists(n string, v *volumes.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.storageV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL storage client: %s", err)
		}

		found, err := volumes.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Volume not found")
		}

		*v = *found

		return nil
	}
}

func testAccCheckStorageV1VolumeIQNLengthIsZERO(v *volumes.Volume) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(v.InitiatorIQNs) != 0 {
			return fmt.Errorf("IQN length is not ZERO in volume: %#v", v)
		}

		return nil
	}
}

var testAccStorageV1VolumeBasic = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
  name = "volume_1"
  description = "first test volume"
  virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
  iops_per_gb = "2"
  size = 100
  initiator_iqns = ["%s"]
}
`, testAccStorageV1VirtualStorageBasic, IQN01)

var testAccStorageV1VolumeTimeout = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
  name = "volume_1"
  description = "first test volume"
  virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
  iops_per_gb = "2"
  size = 100
  initiator_iqns = ["%s"]

  timeouts {
    create = "30m"
	delete = "30m"
  }
}
`, testAccStorageV1VirtualStorageBasic, IQN01)

var testAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolumeBasic = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1"
	description = "first test volume"
	iops_per_gb = "2"
	size = 100
	initiator_iqns = ["%s"]
}`,
	testAccStorageV1VirtualStorageVolumeTypeNameBlock,
	IQN01,
)

var testAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolumeUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1-updated"
	description = "first test volume updated"
	iops_per_gb = "2"
	size = 100
	initiator_iqns = ["%s", "%s"]
  }
  
`, testAccStorageV1VirtualStorageVolumeTypeNameBlock,
	IQN01,
	IQN02,
)

var testAccStorageV1VolumeCreateNetworkAndBlockVirtualStorageAndVolumeUpdate2 = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1-updated"
	description = ""
	iops_per_gb = "2"
	size = 100
  }
`, testAccStorageV1VirtualStorageVolumeTypeNameBlock,
)

var testAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolumeBasic = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1"
	description = "first test volume"
	size = 256
	throughput = "50"
}`,
	testAccStorageV1VirtualStorageVolumeTypeNameFilePremium,
)

var testAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolumeUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1-updated"
	description = "first test volume updated"
	size = 256
	throughput = "50"
  }
  
`, testAccStorageV1VirtualStorageVolumeTypeNameFilePremium,
)
var testAccStorageV1VolumeCreateNetworkAndFilePremiumVirtualStorageAndVolumeUpdate2 = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1-updated"
	description = ""
	size = 256
	throughput = "50"
  }
  
`, testAccStorageV1VirtualStorageVolumeTypeNameFilePremium,
)

var testAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolumeBasic = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1"
	description = "first test volume"
	size = 1024
}`,
	testAccStorageV1VirtualStorageVolumeTypeNameFileStandard,
)

var testAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolumeUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1-updated"
	description = "first test volume updated"
	size = 1024
  }
  
`, testAccStorageV1VirtualStorageVolumeTypeNameFileStandard,
)
var testAccStorageV1VolumeCreateNetworkAndFileStandardVirtualStorageAndVolumeUpdate2 = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1-updated"
	description = ""
	size = 1024
  }
  
`, testAccStorageV1VirtualStorageVolumeTypeNameFileStandard,
)

var testAccStorageV1VolumeForceNewByIOPSPerGBBasic = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
  name = "volume_1"
  description = "first test volume"
  virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
  iops_per_gb = "2"
  size = 100
  initiator_iqns = ["%s"]
}
`, testAccStorageV1VirtualStorageBasic, IQN01)

var testAccStorageV1VolumeForceNewByIOPSPerGBUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
  name = "volume_1-forcednew"
  description = "first test volume forcednew"
  virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
  iops_per_gb = "4"
  size = 100
  initiator_iqns = ["%s"]
}
`, testAccStorageV1VirtualStorageBasic, IQN01)

var testAccStorageV1VolumeForceNewBySizeBasic = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
  name = "volume_1"
  description = "first test volume"
  virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
  iops_per_gb = "2"
  size = 100
  initiator_iqns = ["%s"]
}
`, testAccStorageV1VirtualStorageBasic, IQN01)

var testAccStorageV1VolumeForceNewBySizeUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
  name = "volume_1-forcednew"
  description = "first test volume forcednew"
  virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
  iops_per_gb = "2"
  size = 250
  initiator_iqns = ["%s"]
}
`, testAccStorageV1VirtualStorageBasic, IQN01)

var testAccStorageV1VolumeForceNewByThroughputBasic = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1"
	description = "first test volume"
	size = 256
	throughput = "50"
}`,
	testAccStorageV1VirtualStorageVolumeTypeNameFilePremium,
)

var testAccStorageV1VolumeForceNewByThroughputUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_volume_v1" "volume_1" {
	virtual_storage_id = "${ecl_storage_virtualstorage_v1.virtualstorage_1.id}"
	name = "volume_1-forcednew"
	description = "first test volume forcednew"
	size = 256
	throughput = "100"
  }
  
`, testAccStorageV1VirtualStorageVolumeTypeNameFilePremium,
)

const testValidateStorageV1VolumeConflictsIOPSPerGBAndThroughput = `
resource "ecl_storage_volume_v1" "volume_1" {
	size = "100"
	iops_per_gb = "2"
	throughput = "50"
	virtual_storage_id = "dummyid"
}
`

var testValidateStorageV1VolumeConflictsInitiatorIQNsAndThroughput = fmt.Sprintf(`
resource "ecl_storage_volume_v1" "volume_1" {
	size = "100"
	initiator_iqns = ["%s"]
	throughput = "50"
	virtual_storage_id = "dummyid"
}
`, IQN01,
)
