package ecl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/storage/v1/virtualstorages"
)

const VOLUME_TYPE_NAME_BLOCK = "piops_iscsi_na"
const VOLUME_TYPE_NAME_FILE_PREMIUM = "pre_nfs_na"
const VOLUME_TYPE_NAME_FILE_STANDARD = "standard_nfs_na"

// TestAccStorageV1VirtualStorageBasic is basic test for storage/virtual storage
// This function test followings
//		1. Create network and subnet which be connected from virtual storage
//		2. Create virtual storage by using environment values
//		3. Check if each parameters are correctly set to created virtual storage
//		4. Update it by modified HCL configurations
//			All parameters are updated except ip addr pool start
//		5. Check if new configuration is correctly applied
func TestAccStorageV1VirtualStorageBasic(t *testing.T) {
	var vs virtualstorages.VirtualStorage
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorage(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VirtualStorageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_id", OS_STORAGE_VOLUME_TYPE_ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.10"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.20"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_id", OS_STORAGE_VOLUME_TYPE_ID),

					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.9"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.21"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.1.destination", "1.1.2.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.1.nexthop", "192.168.1.1"),
				),
			},
		},
	})
}

// TestAccStorageV1VirtualStorageTimeout just check Timeout section is correctly worked or not
func TestAccStorageV1VirtualStorageTimeout(t *testing.T) {
	var vs virtualstorages.VirtualStorage
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorage(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VirtualStorageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageTimeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1"),
				),
			},
		},
	})
}

// TestAccStorageV1VirtualStorageVolumeTypeNameBlock is practical test case from user's stand point
// In this test case, volume_type_name is used. Even this is not API original parameter
// but is meant to enhance user's convenience
// This test is the case of Block Storage service
// Test items are almost same as TestAccStorageV1VirtualStorageBasic
// Create -> Parameter Confirmation -> Update -> Parameter Confirmation(All parameters are correctly updated or not)
func TestAccStorageV1VirtualStorageVolumeTypeNameBlock(t *testing.T) {
	var vs virtualstorages.VirtualStorage
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorage(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VirtualStorageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageVolumeTypeNameBlock,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_name", VOLUME_TYPE_NAME_BLOCK),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.10"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.20"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageVolumeTypeNameBlockUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_name", VOLUME_TYPE_NAME_BLOCK),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.9"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.21"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageVolumeTypeNameBlockUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", ""),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_name", VOLUME_TYPE_NAME_BLOCK),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.9"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.21"),
					testAccCheckStorageV1VirtualStorageHostRouteLengthIsZERO(&vs),
				),
			},
		},
	})
}

// TestAccStorageV1VirtualStorageVolumeTypeNameFilePremium is practical test case from user's stand point
// In this test case, volume_type_name is used. Even this is not API original parameter
// but is meant to enhance user's convenience
// This test is the case of File Storage(Premium) service
// Test items are almost same as TestAccStorageV1VirtualStorageBasic
// Create -> Parameter Confirmation -> Update -> Parameter Confirmation(All parameters are correctly updated or not)
func TestAccStorageV1VirtualStorageVolumeTypeNameFilePremium(t *testing.T) {
	var vs virtualstorages.VirtualStorage
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckStorage(t)
			testAccPreCheckFileStorageServiceType(t, true, false)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VirtualStorageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageVolumeTypeNameFilePremium,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_name", VOLUME_TYPE_NAME_FILE_PREMIUM),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.10"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.20"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageVolumeTypeNameFilePremiumUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_name", VOLUME_TYPE_NAME_FILE_PREMIUM),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.9"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.21"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.1.destination", "1.1.2.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.1.nexthop", "192.168.1.1"),
				),
			},
		},
	})
}

// TestAccStorageV1VirtualStorageVolumeTypeNameFileStandard is practical test case from user's stand point
// In this test case, volume_type_name is used. Even this is not API original parameter
// but is meant to enhance user's convenience
// This test is the case of File Storage(Standard) service
// Test items are almost same as TestAccStorageV1VirtualStorageBasic
// Create -> Parameter Confirmation -> Update -> Parameter Confirmation(All parameters are correctly updated or not)
func TestAccStorageV1VirtualStorageVolumeTypeNameFileStandard(t *testing.T) {
	var vs virtualstorages.VirtualStorage
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckStorage(t)
			testAccPreCheckFileStorageServiceType(t, false, true)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VirtualStorageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageVolumeTypeNameFileStandard,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_name", VOLUME_TYPE_NAME_FILE_STANDARD),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.10"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.20"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageVolumeTypeNameFileStandardUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage-updated"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_name", VOLUME_TYPE_NAME_FILE_STANDARD),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.9"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.21"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.1.destination", "1.1.2.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.1.nexthop", "192.168.1.1"),
				),
			},
		},
	})
}

func testCheckVirtualStorageIDIsChanged(vs1, vs2 *virtualstorages.VirtualStorage) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if vs1.ID == vs2.ID {
			log.Printf("[DEBUG] Virtual Storage (Comparison target) IDs: %s, %s", vs1.ID, vs2.ID)
			return fmt.Errorf("Resource ID is not changed. ForeceNew parameter does not work correctly")
		}
		return nil
	}
}

// TestAccStorageV1VirtualStorageForceNewByNetwork check if ForceNew works correctly
// This function has 2 TestSteps
// Firstly, create virtual storage with new network
// Secondly, create another network and change network_id of virtual storage as new network
// This cases ForceNew operation
// So this test case check whether 1st and 2nd virtual storage ID points different UUID correctly or not
// Each ID is inserted into virtualStorageIDs slice by beggining of each TestStep by using "testCheckIDIsChanged"
// and also be compared in same function
func TestAccStorageV1VirtualStorageForceNewByNetwork(t *testing.T) {
	var vs1, vs2 virtualstorages.VirtualStorage

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStorage(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStorageV1VirtualStorageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageCheckForceNewByNetworkBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs1),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_id", OS_STORAGE_VOLUME_TYPE_ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.1.10"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.1.20"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.1.1"),
				),
			},
			resource.TestStep{
				Config: testAccStorageV1VirtualStorageCheckForceNewByNetworkUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStorageV1VirtualStorageExists("ecl_storage_virtualstorage_v1.virtualstorage_1", &vs2),
					testCheckVirtualStorageIDIsChanged(&vs1, &vs2),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "name", "virtualstorage_1-forcednew"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "description", "first test virtual storage-forcednew"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "volume_type_id", OS_STORAGE_VOLUME_TYPE_ID),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.start", "192.168.2.10"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "ip_addr_pool.end", "192.168.2.20"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.destination", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.0.nexthop", "192.168.2.1"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.1.destination", "1.1.2.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_storage_virtualstorage_v1.virtualstorage_1", "host_routes.1.nexthop", "192.168.2.1"),
				),
			},
		},
	})
}

func testAccCheckStorageV1VirtualStorageDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.storageV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL storage client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_storage_virtualstorage_v1" {
			continue
		}

		_, err := virtualstorages.Get(client, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("VirtualStorage still exists")
		}
	}

	return nil
}

func testAccCheckStorageV1VirtualStorageExists(n string, vs *virtualstorages.VirtualStorage) resource.TestCheckFunc {
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

		found, err := virtualstorages.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("VirtualStorage not found")
		}

		*vs = *found

		return nil
	}
}

func testAccCheckStorageV1VirtualStorageHostRouteLengthIsZERO(vs *virtualstorages.VirtualStorage) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(vs.HostRoutes) != 0 {
			return fmt.Errorf("Route length is not ZERO in virtual storage: %#v", vs)
		}

		return nil
	}
}

const testAccStorageV1VirtualStorageNetworkAndSubnetBasic = `
resource "ecl_network_network_v2" "network_1" {
	name = "terraform-temp-network-data"
	plane = "%s"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	cidr = "192.168.1.0/24"
	network_id = "${ecl_network_network_v2.network_1.id}"
	gateway_ip = "192.168.1.1"
    
	allocation_pools {
	  start = "192.168.1.100"
	  end = "192.168.1.200"
	}
}   
`

var testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane = fmt.Sprintf(
	testAccStorageV1VirtualStorageNetworkAndSubnetBasic,
	"data",
)

var testAccStorageV1VirtualStorageNetworkAndSubnetStoragePlane = fmt.Sprintf(
	testAccStorageV1VirtualStorageNetworkAndSubnetBasic,
	"storage",
)

var testAccStorageV1VirtualStorageBasic = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1"
  description = "first test virtual storage"
  volume_type_id = "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.10"
    end = "192.168.1.20"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane,
	OS_STORAGE_VOLUME_TYPE_ID,
)

var testAccStorageV1VirtualStorageUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1-updated"
  description = "first test virtual storage-updated"
  volume_type_id = "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.9"
    end = "192.168.1.21"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
  host_routes {
    destination = "1.1.2.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane,
	OS_STORAGE_VOLUME_TYPE_ID,
)

var testAccStorageV1VirtualStorageTimeout = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1"
  description = "first test virtual storage"
  volume_type_id = "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.10"
    end = "192.168.1.20"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
  timeouts {
    create = "30m"
    delete = "30m"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane,
	OS_STORAGE_VOLUME_TYPE_ID,
)

var testAccStorageV1VirtualStorageVolumeTypeNameBlock = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1"
  description = "first test virtual storage"
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.10"
    end = "192.168.1.20"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane,
	VOLUME_TYPE_NAME_BLOCK,
)

var testAccStorageV1VirtualStorageVolumeTypeNameBlockUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1-updated"
  description = "first test virtual storage-updated"
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.9"
    end = "192.168.1.21"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
  host_routes {
    destination = "1.1.2.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane,
	VOLUME_TYPE_NAME_BLOCK,
)
var testAccStorageV1VirtualStorageVolumeTypeNameBlockUpdate2 = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1-updated"
  description = ""
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.9"
    end = "192.168.1.21"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane,
	VOLUME_TYPE_NAME_BLOCK,
)

var testAccStorageV1VirtualStorageVolumeTypeNameFilePremium = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1"
  description = "first test virtual storage"
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.10"
    end = "192.168.1.20"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetStoragePlane,
	VOLUME_TYPE_NAME_FILE_PREMIUM,
)

var testAccStorageV1VirtualStorageVolumeTypeNameFilePremiumUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1-updated"
  description = "first test virtual storage-updated"
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.9"
    end = "192.168.1.21"
  }

  host_routes{
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
  host_routes {
    destination = "1.1.2.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetStoragePlane,
	VOLUME_TYPE_NAME_FILE_PREMIUM,
)

var testAccStorageV1VirtualStorageVolumeTypeNameFileStandard = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1"
  description = "first test virtual storage"
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.10"
    end = "192.168.1.20"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetStoragePlane,
	VOLUME_TYPE_NAME_FILE_STANDARD,
)

var testAccStorageV1VirtualStorageVolumeTypeNameFileStandardUpdate = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1-updated"
  description = "first test virtual storage-updated"
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.9"
    end = "192.168.1.21"
  }

  host_routes {
		destination = "1.1.1.0/24"
		nexthop = "192.168.1.1"
  }

  host_routes {
    destination = "1.1.2.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetStoragePlane,
	VOLUME_TYPE_NAME_FILE_STANDARD,
)

var testAccStorageV1VirtualStorageCheckForceNewByNetworkBasic = fmt.Sprintf(`
%s

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1"
  description = "first test virtual storage"
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_1.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
  ip_addr_pool = {
    start = "192.168.1.10"
    end = "192.168.1.20"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.1.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane,
	VOLUME_TYPE_NAME_BLOCK,
)

var testAccStorageV1VirtualStorageCheckForceNewByNetworkUpdate = fmt.Sprintf(`
%s

resource "ecl_network_network_v2" "network_2" {
	name = "terraform-temp-network-data2"
	plane = "data"
}

resource "ecl_network_subnet_v2" "subnet_2" {
	cidr = "192.168.2.0/24"
	network_id = "${ecl_network_network_v2.network_2.id}"
	gateway_ip = "192.168.2.1"
    
	allocation_pools {
	  start = "192.168.2.100"
	  end = "192.168.2.200"
	}
}   

resource "ecl_storage_virtualstorage_v1" "virtualstorage_1" {
  name = "virtualstorage_1-forcednew"
  description = "first test virtual storage-forcednew"
  volume_type_name= "%s"
  network_id = "${ecl_network_network_v2.network_2.id}"
  subnet_id = "${ecl_network_subnet_v2.subnet_2.id}"
  ip_addr_pool = {
    start = "192.168.2.10"
    end = "192.168.2.20"
  }
  host_routes {
    destination = "1.1.1.0/24"
    nexthop = "192.168.2.1"
  }
  host_routes {
    destination = "1.1.2.0/24"
    nexthop = "192.168.2.1"
  }
}
`, testAccStorageV1VirtualStorageNetworkAndSubnetDataPlane,
	VOLUME_TYPE_NAME_BLOCK,
)
