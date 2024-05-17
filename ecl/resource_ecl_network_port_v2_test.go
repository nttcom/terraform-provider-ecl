package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v3/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/ports"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/subnets"
)

func TestAccNetworkV2Port_basic(t *testing.T) {
	var network networks.Network
	var port ports.Port
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "name", "port_1"),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "description", "port_1_description"),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "admin_state_up", "true"),
					testAccCheckNetworkV2PortTag(&port, "k1", "v1"),
					testAccCheckNetworkV2PortNoTagKey(&port, "k2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortUpdate1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "name", "port_1-update"),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "description", "port_1_description-update"),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "admin_state_up", "false"),
					testAccCheckNetworkV2PortTag(&port, "k1", "v1"),
					testAccCheckNetworkV2PortTag(&port, "k2", "v2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "name", ""),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "description", "port_1_description-update"),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "admin_state_up", "false"),
					testAccCheckNetworkV2PortNoTagKey(&port, "k1"),
					testAccCheckNetworkV2PortTag(&port, "k2", "v2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortUpdate3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "name", "port_1_name"),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "description", ""),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "admin_state_up", "true"),
					testAccCheckNetworkV2PortTagLengthIsZERO(&port),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_conflictsNoFixedIPAndFixedIPs(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	pattern1 := "\"no_fixed_ip\": conflicts with fixed_ip"
	pattern2 := "\"fixed_ip\": conflicts with no_fixed_ip"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testValidateNetworkV2PortConflictsNoFixedIPAndFixedIPs,
				ExpectError: regexp.MustCompile(pattern1),
			},
			resource.TestStep{
				Config:      testValidateNetworkV2PortConflictsNoFixedIPAndFixedIPs,
				ExpectError: regexp.MustCompile(pattern2),
			},
		},
	})
}

func TestAccNetworkV2Port_FixedIPBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network
	var port ports.Port
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortFixedIPBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "fixed_ip.0.ip_address", "192.168.199.21"),
					testAccCheckNetworkV2PortCountFixedIPs(&port, 1),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PortFixedIPUpdate1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "fixed_ip.0.ip_address", "192.168.199.21"),
					resource.TestCheckResourceAttr("ecl_network_port_v2.port_1", "fixed_ip.1.ip_address", "192.168.199.22"),
					testAccCheckNetworkV2PortCountFixedIPs(&port, 2),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_noIP(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network
	var port ports.Port
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortNoIP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					testAccCheckNetworkV2PortCountFixedIPs(&port, 1),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_multipleNoIP(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network
	var port ports.Port
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortMultipleNoIP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					testAccCheckNetworkV2PortCountFixedIPs(&port, 3),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_allowedAddressPairs(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network
	var subnet subnets.Subnet
	var vrrpPort1, vrrpPort2, instancePort ports.Port

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortAllowedAddressPairs,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.vrrp_subnet", &subnet),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.vrrp_network", &network),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.vrrp_port1", &vrrpPort1),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.vrrp_port2", &vrrpPort2),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.instance_port", &instancePort),
					testAccCheckNetworkV2PortCountAllowedAddressPairs(&instancePort, 2),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_allowedAddressPairsNoMAC(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network
	var subnet subnets.Subnet
	var vrrpPort1, vrrpPort2, instancePort ports.Port

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortAllowedAddressPairsNoMAC,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.vrrp_subnet", &subnet),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.vrrp_network", &network),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.vrrp_port1", &vrrpPort1),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.vrrp_port2", &vrrpPort2),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.instance_port", &instancePort),
					testAccCheckNetworkV2PortCountAllowedAddressPairs(&instancePort, 2),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_multipleFixedIPs(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network
	var port ports.Port
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortMultipleFixedIPs,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					testAccCheckNetworkV2PortCountFixedIPs(&port, 3),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_timeout(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network
	var port ports.Port
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortTimeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_fixedIPs(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortFixedIPs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_port_v2.port_1", "all_fixed_ips.0", "192.168.199.23"),
					resource.TestCheckResourceAttr(
						"ecl_network_port_v2.port_1", "all_fixed_ips.1", "192.168.199.24"),
				),
			},
		},
	})
}

func TestAccNetworkV2Port_noFixedIP(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var port ports.Port

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PortDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PortNoFixedIP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					resource.TestCheckResourceAttr(
						"ecl_network_port_v2.port_1", "all_fixed_ips.#", "0"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2PortDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkingClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_port_v2" {
			continue
		}

		_, err := ports.Get(networkingClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Port still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2PortExists(n string, port *ports.Port) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkingClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL networking client: %s", err)
		}

		found, err := ports.Get(networkingClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Port not found")
		}

		*port = *found

		return nil
	}
}

func testAccCheckNetworkV2PortCountFixedIPs(port *ports.Port, expected int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(port.FixedIPs) != expected {
			return fmt.Errorf("Expected %d Fixed IPs, got %d", expected, len(port.FixedIPs))
		}

		return nil
	}
}

func testAccCheckNetworkV2PortCountAllowedAddressPairs(
	port *ports.Port, expected int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(port.AllowedAddressPairs) != expected {
			return fmt.Errorf("Expected %d Allowed Address Pairs, got %d", expected, len(port.AllowedAddressPairs))
		}

		return nil
	}
}

func testAccCheckNetworkV2PortTagLengthIsZERO(port *ports.Port) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(port.Tags) != 0 {
			return fmt.Errorf("Tag length is not ZERO")
		}
		return nil
	}
}

func testAccCheckNetworkV2PortTag(
	port *ports.Port, k string, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if port.Tags == nil {
			return fmt.Errorf("No tag")
		}

		for key, value := range port.Tags {
			if k != key {
				continue
			}

			if v == value {
				return nil
			}

			return fmt.Errorf("Bad value for %s: %s", k, value)
		}

		return fmt.Errorf("Tag not found: %s", k)
	}
}

func testAccCheckNetworkV2PortNoTagKey(
	port *ports.Port, k string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if port.Tags == nil {
			return nil
		}

		for key := range port.Tags {
			if k == key {
				return fmt.Errorf("Tag found: %s", k)
			}
		}

		return nil
	}
}

const testBaseNetworkAndSubnet = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
}
`

var testAccNetworkV2PortBasic = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  description = "port_1_description"
  admin_state_up = "true"
	network_id = "${ecl_network_network_v2.network_1.id}"
	segmentation_id = 100
	segmentation_type = "vlan"

	tags = {
		k1 = "v1"
	}

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.23"
  }
}`, testBaseNetworkAndSubnet)

var testAccNetworkV2PortUpdate1 = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  name = "port_1-update"
  description = "port_1_description-update"
  admin_state_up = "false"
  network_id = "${ecl_network_network_v2.network_1.id}"
	segmentation_type = "flat"

	tags = {
		k1 = "v1"
		k2 = "v2"
	}

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.23"
  }
}`, testBaseNetworkAndSubnet)

var testAccNetworkV2PortUpdate2 = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  name = ""
  description = "port_1_description-update"
  admin_state_up = "false"
  network_id = "${ecl_network_network_v2.network_1.id}"
	segmentation_type = "flat"

	tags = {
		k2 = "v2"
	}

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.23"
  }
}`, testBaseNetworkAndSubnet)

var testAccNetworkV2PortUpdate3 = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  name = "port_1_name"
  description = ""
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.23"
  }
}`, testBaseNetworkAndSubnet)

var testAccNetworkV2PortFixedIPBasic = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.21"
  }
}`, testBaseNetworkAndSubnet)

var testAccNetworkV2PortFixedIPUpdate1 = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.21"
  }

	fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.22"
  }
}`, testBaseNetworkAndSubnet)

var testAccNetworkV2PortFixedIPUpdate2 = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  network_id = "${ecl_network_network_v2.network_1.id}"
}`, testBaseNetworkAndSubnet)

const testAccNetworkV2PortNoIP = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
}

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
	admin_state_up = "true"
	description = "port_description"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
  }
}
`

var testAccNetworkV2PortFixdIPBasic = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  description = "port_1_description"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.network_1.id}"

	tags = {
		k1 = "v1"
	}

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.23"
  }
}`, testBaseNetworkAndSubnet)

const testAccNetworkV2PortMultipleNoIP = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
}

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
  }

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
  }

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
  }
}
`

const testAccNetworkV2PortAllowedAddressPairs = `
resource "ecl_network_network_v2" "vrrp_network" {
  name = "vrrp_network"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "vrrp_subnet" {
  name = "vrrp_subnet"
  cidr = "10.0.0.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.vrrp_network.id}"

  allocation_pools {
    start = "10.0.0.2"
    end = "10.0.0.200"
  }
}

resource "ecl_network_port_v2" "vrrp_port1" {
  name = "vrrp_port1"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.vrrp_network.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.vrrp_subnet.id}"
    ip_address = "10.0.0.202"
  }
}

resource "ecl_network_port_v2" "vrrp_port2" {
  name = "vrrp_port2"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.vrrp_network.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.vrrp_subnet.id}"
    ip_address = "10.0.0.201"
  }
}

resource "ecl_network_port_v2" "instance_port" {
  name = "instance_port"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.vrrp_network.id}"

  allowed_address_pairs {
    ip_address = "${ecl_network_port_v2.vrrp_port1.fixed_ip.0.ip_address}"
    mac_address = "${ecl_network_port_v2.vrrp_port1.mac_address}"
  }

  allowed_address_pairs {
    ip_address = "${ecl_network_port_v2.vrrp_port2.fixed_ip.0.ip_address}"
    mac_address = "${ecl_network_port_v2.vrrp_port2.mac_address}"
  }
}
`

const testAccNetworkV2PortMultipleFixedIPs = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
}

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.23"
  }

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.20"
  }

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.40"
  }
}
`

const testAccNetworkV2PortTimeout = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
}

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.23"
  }

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`

const testAccNetworkV2PortFixedIPs = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
}

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.23"
  }

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.199.24"
  }
}
`

const testAccNetworkV2PortAllowedAddressPairsNoMAC = `
resource "ecl_network_network_v2" "vrrp_network" {
  name = "vrrp_network"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "vrrp_subnet" {
  name = "vrrp_subnet"
  cidr = "10.0.0.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.vrrp_network.id}"

  allocation_pools {
    start = "10.0.0.2"
    end = "10.0.0.200"
  }
}

resource "ecl_network_port_v2" "vrrp_port1" {
  name = "vrrp_port1"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.vrrp_network.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.vrrp_subnet.id}"
    ip_address = "10.0.0.202"
  }
}

resource "ecl_network_port_v2" "vrrp_port2" {
  name = "vrrp_port2"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.vrrp_network.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.vrrp_subnet.id}"
    ip_address = "10.0.0.201"
  }
}

resource "ecl_network_port_v2" "instance_port" {
  name = "instance_port"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.vrrp_network.id}"

  allowed_address_pairs {
    ip_address = "${ecl_network_port_v2.vrrp_port1.fixed_ip.0.ip_address}"
  }

  allowed_address_pairs {
    ip_address = "${ecl_network_port_v2.vrrp_port2.fixed_ip.0.ip_address}"
  }
}
`

const testAccNetworkV2PortNoFixedIP = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
}

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.network_1.id}"
	no_fixed_ip = true
}
`

const testValidateNetworkV2PortConflictsNoFixedIPAndFixedIPs = `
resource "ecl_network_port_v2" "port_1" {
	network_id = "dummy_network"
	no_fixed_ip = "true"
	fixed_ip {
		subnet_id =  "dummy_subnet"
    ip_address = "192.168.199.1"
	}
}`
