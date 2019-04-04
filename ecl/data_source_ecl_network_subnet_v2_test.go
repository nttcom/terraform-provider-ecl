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

func TestAccNetworkV2SubnetDataSourceBasic(t *testing.T) {
	name, description, cidr, gatewayIP := generateNetworkSubnetQueryParams()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceName(name, description, cidr, gatewayIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
					testAccCheckNetworkV2SubnetDataSourceGoodNetwork("data.ecl_network_subnet_v2.subnet_1", "ecl_network_network_v2.network_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_subnet_v2.subnet_1", "name", name),
				),
			},
		},
	})
}

func TestAccNetworkV2SubnetDataSourceTestQueries(t *testing.T) {
	name, description, cidr, gatewayIP := generateNetworkSubnetQueryParams()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceCIDR(name, description, cidr, gatewayIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceDescription(name, description, cidr, gatewayIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceGatewayIP(name, description, cidr, gatewayIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceName(name, description, cidr, gatewayIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceNetworkID(name, description, cidr, gatewayIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceID(name, description, cidr, gatewayIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetDataSourceID("data.ecl_network_subnet_v2.subnet_1"),
				),
			},
		},
	})
}

func TestAccNetworkV2SubnetDataSourceNetworkIdAttribute(t *testing.T) {
	name, description, cidr, gatewayIP := generateNetworkSubnetQueryParams()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetDataSourceNetworkIDAttribute(name, description, cidr, gatewayIP),
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

func testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP string) string {
	return fmt.Sprintf(`
	resource "ecl_network_network_v2" "network_1" {
	  name = "network_1"
	  admin_state_up = "true"
	}
	
	resource "ecl_network_subnet_v2" "subnet_1" {
	  name = "%s"
	  description = "%s"
	  cidr = "%s"
	  gateway_ip = "%s"
	  network_id = "${ecl_network_network_v2.network_1.id}"
	  tags = {
		  key1 = "value1"
	  }
	}`, name, description, cidr, gatewayIP)

}

func testAccNetworkV2SubnetDataSourceCIDR(name, description, cidr, gatewayIP string) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_subnet_v2" "subnet_1" {
			cidr = "${ecl_network_subnet_v2.subnet_1.cidr}"
		}`, testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP))

}

func testAccNetworkV2SubnetDataSourceDescription(name, description, cidr, gatewayIP string) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_subnet_v2" "subnet_1" {
			description = "${ecl_network_subnet_v2.subnet_1.description}"
		}`, testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP))
}

func testAccNetworkV2SubnetDataSourceGatewayIP(name, description, cidr, gatewayIP string) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_subnet_v2" "subnet_1" {
			gateway_ip = "${ecl_network_subnet_v2.subnet_1.gateway_ip}"
		}`, testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP))
}

func testAccNetworkV2SubnetDataSourceName(name, description, cidr, gatewayIP string) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_subnet_v2" "subnet_1" {
			name = "${ecl_network_subnet_v2.subnet_1.name}"
		}`, testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP))
}

func testAccNetworkV2SubnetDataSourceNetworkID(name, description, cidr, gatewayIP string) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_subnet_v2" "subnet_1" {
			network_id = "${ecl_network_subnet_v2.subnet_1.network_id}"
		}`, testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP))
}

func testAccNetworkV2SubnetDataSourceID(name, description, cidr, gatewayIP string) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_subnet_v2" "subnet_1" {
			subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
		}`, testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP))
}

func testAccNetworkV2SubnetDataSourceNetworkIDAttribute(name, description, cidr, gatewayIP string) string {
	return fmt.Sprintf(`
			%s

		data "ecl_network_subnet_v2" "subnet_1" {
			subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
		}

		resource "ecl_network_port_v2" "port_1" {
			name               = "test_port"
			network_id         = "${data.ecl_network_subnet_v2.subnet_1.network_id}"
			admin_state_up  = "true"
		}`, testAccNetworkV2SubnetDataSourceSubnet(name, description, cidr, gatewayIP))
}

func generateNetworkSubnetQueryParams() (string, string, string, string) {
	name := fmt.Sprintf("ACPTTEST%s-network", acctest.RandString(5))
	description := fmt.Sprintf("ACPTTEST%s-network-description", acctest.RandString(5))

	rand.Seed(time.Now().UnixNano())
	thirdOctet := rand.Intn(255)
	cidr := fmt.Sprintf("192.168.%d.0/24", thirdOctet)
	gatewayIP := fmt.Sprintf("192.168.%d.1", thirdOctet)

	return name, description, cidr, gatewayIP
}
