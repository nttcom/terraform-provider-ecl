package ecl

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2PortDataSource_basic(t *testing.T) {
	name, description, segmentationID := generateNetworkPortQueryParams()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2PortDataSourceBasic(name, description, segmentationID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.ecl_network_port_v2.port_1", "id",
						"ecl_network_port_v2.port_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_port_v2.port_1", "name", name),
				),
			},
		},
	})
}

func TestAccNetworkV2PortDataSource_queries(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	name, description, segmentationID := generateNetworkPortQueryParams()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourcePort(name, description, segmentationID),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceID(name, description, segmentationID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortDataSourceID("data.ecl_network_port_v2.port_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceDescription(name, description, segmentationID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortDataSourceID("data.ecl_network_port_v2.port_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceDeviceOwnerAndNetworkID(name, description, segmentationID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortDataSourceID("data.ecl_network_port_v2.port_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceMACAddress(name, description, segmentationID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortDataSourceID("data.ecl_network_port_v2.port_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceSegmentationTypeAndID(name, description, segmentationID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortDataSourceID("data.ecl_network_port_v2.port_1"),
				),
			},
		},
	})
}

func TestAccNetworkV2PortDataSource_queriesDeviceID(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	name, description, segmentationID := generateNetworkPortQueryParams()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourcePortForDeviceID(name, description, segmentationID),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceDeviceID(name, description, segmentationID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortDataSourceID("data.ecl_network_port_v2.port_1"),
				),
			},
		},
	})
}
func TestAccNetworkV2PortDataSource_networkIDAttribute(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	name, description, segmentationID := generateNetworkPortQueryParams()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceNetworkIDAttribute(name, description, segmentationID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortDataSourceID("data.ecl_network_port_v2.port_1"),
					testAccCheckNetworkV2PortDataSourceGoodNetwork("data.ecl_network_port_v2.port_1", "ecl_network_network_v2.network_1"),
					testAccCheckNetworkV2PortID("ecl_network_port_v2.port_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2PortDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find port data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Port data source ID not set")
		}

		return nil
	}
}

func testAccCheckNetworkV2SubnetID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find subnet resource: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Subnet resource ID not set")
		}

		return nil
	}
}

func testAccCheckNetworkV2PortDataSourceGoodNetwork(n1, n2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds1, ok := s.RootModule().Resources[n1]
		if !ok {
			return fmt.Errorf("Can't find subnet data source: %s", n1)
		}

		if ds1.Primary.ID == "" {
			return fmt.Errorf("Subnet data source ID not set")
		}

		rs2, ok := s.RootModule().Resources[n2]
		if !ok {
			return fmt.Errorf("Can't find network resource: %s", n2)
		}

		if rs2.Primary.ID == "" {
			return fmt.Errorf("Network resource ID not set")
		}

		if rs2.Primary.ID != ds1.Primary.Attributes["network_id"] {
			return fmt.Errorf("Network id and subnet network_id don't match")
		}

		return nil
	}
}

func testAccNetworkV2PortDataSourcePort(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
	resource "ecl_network_network_v2" "network_1" {
	  name = "network_1"
	  admin_state_up = "true"
	}
	
	resource "ecl_network_port_v2" "port_1" {
	  name = "%s"
	  description = "%s"
	  segmentation_type = "vlan"
	  segmentation_id = %d
	  network_id = "${ecl_network_network_v2.network_1.id}"
	}`, name, description, segmentationID)
}

func testAccNetworkingV2PortDataSourceBasic(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		resource "ecl_network_network_v2" "network_1" {
		  name           = "network_1"
		  admin_state_up = "true"
		}

		resource "ecl_network_subnet_v2" "subnet_1" {
		  name       = "subnet_1"
		  network_id = "${ecl_network_network_v2.network_1.id}"
		  cidr       = "10.0.0.0/24"
		  ip_version = 4
		}

		resource "ecl_network_port_v2" "port_1" {
		  name           = "%s"
		  description    = "%s"
		  network_id     = "${ecl_network_network_v2.network_1.id}"
		  admin_state_up = "true"
		  segmentation_type = "vlan"
		  segmentation_id = %d
		}

		data "ecl_network_port_v2" "port_1" {
		  name           = "${ecl_network_port_v2.port_1.name}"
		}`, name, description, segmentationID)
}

func testAccNetworkV2PortDataSourceID(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_port_v2" "port_1" {
			port_id = "${ecl_network_port_v2.port_1.id}"
		}`, testAccNetworkV2PortDataSourcePort(name, description, segmentationID))
}

func testAccNetworkV2PortDataSourceDescription(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_port_v2" "port_1" {
			description = "${ecl_network_port_v2.port_1.description}"
		}`, testAccNetworkV2PortDataSourcePort(name, description, segmentationID))
}

func testAccNetworkV2PortDataSourceDeviceOwnerAndNetworkID(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_port_v2" "port_1" {
			device_owner = "${ecl_network_port_v2.port_1.device_owner}"
			network_id = "${ecl_network_port_v2.port_1.network_id}"
		}`, testAccNetworkV2PortDataSourcePort(name, description, segmentationID))
}

func testAccNetworkV2PortDataSourceMACAddress(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_port_v2" "port_1" {
			mac_address = "${ecl_network_port_v2.port_1.mac_address}"
		}`, testAccNetworkV2PortDataSourcePort(name, description, segmentationID))
}

func testAccNetworkV2PortDataSourceName(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_port_v2" "port_1" {
			name = "${ecl_network_port_v2.port_1.name}"
		}`, testAccNetworkV2PortDataSourcePort(name, description, segmentationID))
}

func testAccNetworkV2PortDataSourceSegmentationTypeAndID(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_port_v2" "port_1" {
			segmentation_type = "${ecl_network_port_v2.port_1.segmentation_type}"
			segmentation_id = "${ecl_network_port_v2.port_1.segmentation_id}"
		}`, testAccNetworkV2PortDataSourcePort(name, description, segmentationID))
}

func testAccNetworkV2PortDataSourceNetworkIDAttribute(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		%s
	
		data "ecl_network_port_v2" "port_1" {
			port_id = "${ecl_network_port_v2.port_1.id}"
		}
	
		resource "ecl_network_subnet_v2" "subnet_1" {
			name               = "test_subnet"
			network_id         = "${data.ecl_network_port_v2.port_1.network_id}"
			cidr               = "192.168.1.0/24"
		}`, testAccNetworkV2PortDataSourcePort(name, description, segmentationID))

}

func generateNetworkPortQueryParams() (string, string, int) {
	name := fmt.Sprintf("ACPTTEST%s-port", acctest.RandString(5))
	description := fmt.Sprintf("ACPTTEST%s-port-description", acctest.RandString(5))

	rand.Seed(time.Now().UnixNano())
	segmentationID := rand.Intn(250) + 1

	return name, description, segmentationID
}

func testAccNetworkV2PortDataSourcePortForDeviceID(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
	resource "ecl_network_network_v2" "network_1" {
	  name = "network_1"
	  admin_state_up = "true"
	}

	resource "ecl_network_subnet_v2" "subnet_1" {
		name = "subnet_1"
		network_id = "${ecl_network_network_v2.network_1.id}"
		cidr = "192.168.1.0/24"
		gateway_ip = "192.168.1.1"
		allocation_pools {
			start = "192.168.1.100"
			end = "192.168.1.200"
		}
	}

	resource "ecl_network_port_v2" "port_1" {
		name = "%s"
		description = "%s"
		segmentation_type = "vlan"
		segmentation_id = %d
		network_id = "${ecl_network_network_v2.network_1.id}"
		fixed_ip {
			subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
			ip_address = "192.168.1.50"
		}
	}
	
	resource "ecl_compute_instance_v2" "instance_1" {
		depends_on = ["ecl_network_subnet_v2.subnet_1"]
		name = "instance_1"
		image_name = "Ubuntu-18.04.1_64_virtual-server_02"
		flavor_id = "1CPU-2GB"
		network {
			port = "${ecl_network_port_v2.port_1.id}"
		}
	}`, name, description, segmentationID)
}

func testAccNetworkV2PortDataSourceDeviceID(name, description string, segmentationID int) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_port_v2" "port_1" {
			device_id = "${ecl_network_port_v2.port_1.device_id}"
		}`, testAccNetworkV2PortDataSourcePortForDeviceID(name, description, segmentationID))
}
