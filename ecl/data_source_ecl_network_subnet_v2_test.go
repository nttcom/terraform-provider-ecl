package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2SubnetDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceSubnet,
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
					testAccCheckNetworkV2SubnetDataSourceGoodNetwork("data.ecl_network_subnet_v2.subnet_1", "ecl_network_network_v2.network_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_subnet_v2.subnet_1", "name", "subnet_1"),
				),
			},
		},
	})
}

func TestAccNetworkV2SubnetDataSourceTestQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceSubnet,
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceCIDR,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceGatewayIP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceIPVersion,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceNetworkID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceStatus,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceTenantID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
		},
	})
}

func TestAccNetworkV2SubnetDataSourceNetworkIdAttribute(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceNetworkIDAttribute,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
					testAccCheckNetworkV2SubnetDataSourceGoodNetwork("data.ecl_network_subnet_v2.subnet_1", "ecl_network_network_v2.network_1"),
					testAccCheckNetworkV2PortID("ecl_network_port_v2.port_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2SubnetDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find subnet data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Subnet data source ID not set")
		}

		return nil
	}
}

func testAccCheckNetworkV2PortID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find port resource: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Port resource ID not set")
		}

		return nil
	}
}

func testAccCheckNetworkV2SubnetDataSourceGoodNetwork(n1, n2 string) resource.TestCheckFunc {
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

const testAccNetworkV2SubnetDataSourceSubnet = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"
  tags = {
	  key1 = "value1"
  }
}
`

var testAccNetworkV2SubnetDataSourceCIDR = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	cidr = "${ecl_network_subnet_v2.subnet_1.cidr}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceDescription = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	description = "${ecl_network_subnet_v2.subnet_1.description}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceGatewayIP = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	gateway_ip = "${ecl_network_subnet_v2.subnet_1.gateway_ip}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceIPVersion = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	ip_version = "${ecl_network_subnet_v2.subnet_1.ip_version}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceName = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	name = "${ecl_network_subnet_v2.subnet_1.name}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceNetworkID = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	network_id = "${ecl_network_subnet_v2.subnet_1.network_id}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceStatus = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	status = "${ecl_network_subnet_v2.subnet_1.status}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceID = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceTenantID = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
	tenant_id = "${ecl_network_subnet_v2.subnet_1.tenant_id}"
}
`, testAccNetworkV2SubnetDataSourceSubnet)

var testAccNetworkV2SubnetDataSourceNetworkIDAttribute = fmt.Sprintf(`
%s

data "ecl_network_subnet_v2" "subnet_1" {
  subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
}

resource "ecl_network_port_v2" "port_1" {
  name               = "test_port"
  network_id         = "${data.ecl_network_subnet_v2.subnet_1.network_id}"
  admin_state_up  = "true"
}

`, testAccNetworkV2SubnetDataSourceSubnet)
