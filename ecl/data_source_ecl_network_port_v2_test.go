package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2PortDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2PortDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.ecl_network_port_v2.port_1", "id",
						"ecl_network_port_v2.port_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_port_v2.port_1", "name", "port"),
				),
			},
		},
	})
}

func TestAccNetworkV2PortDataSourceTestQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourcePort,
			},
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortDataSourceID("data.ecl_network_port_v2.port_1"),
				),
			},
		},
	})
}

func TestAccNetworkV2PortDataSourceNetworkIdAttribute(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortDataSourceNetworkIDAttribute,
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

const testAccNetworkV2PortDataSourcePort = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  network_id = "${ecl_network_network_v2.network_1.id}"
  tags = {
	  key1 = "value1"
  }
}
`

const testAccNetworkingV2PortDataSourceBasic = `
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
  name           = "port"
  description    = "test port"
  network_id     = "${ecl_network_network_v2.network_1.id}"
  admin_state_up = "true"
}

data "ecl_network_port_v2" "port_1" {
  name           = "${ecl_network_port_v2.port_1.name}"
}
`

var testAccNetworkV2PortDataSourceID = fmt.Sprintf(`
%s

data "ecl_network_port_v2" "port_1" {
	port_id = "${ecl_network_port_v2.port_1.id}"
}
`, testAccNetworkV2PortDataSourcePort)

var testAccNetworkV2PortDataSourceNetworkIDAttribute = fmt.Sprintf(`
%s

data "ecl_network_port_v2" "port_1" {
  port_id = "${ecl_network_port_v2.port_1.id}"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name               = "test_subnet"
  network_id         = "${data.ecl_network_port_v2.port_1.network_id}"
  cidr               = "192.168.1.0/24"
}

`, testAccNetworkV2PortDataSourcePort)
