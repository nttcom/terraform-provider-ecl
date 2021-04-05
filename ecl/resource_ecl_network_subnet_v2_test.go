package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v2/ecl/network/v2/subnets"
)

func TestAccNetworkV2Subnet_basic(t *testing.T) {
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetDNSConsistency("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "name", "subnet_1"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "description", "subnet_1_description"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "cidr", "192.168.199.0/24"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "dns_nameservers.0", "1.1.1.1"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "dns_nameservers.1", "2.2.2.2"),
					testAccCheckNetworkV2SubnetDNSLength(&subnet, 2),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "enable_dhcp", "true"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "allocation_pools.0.start", "192.168.199.100"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "allocation_pools.0.end", "192.168.199.200"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "host_routes.0.destination_cidr", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "host_routes.0.next_hop", "192.168.199.1"),
					testAccCheckNetworkV2SubnetHostRoutesLength(&subnet, 1),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetUpdate1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetDNSConsistency("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "name", "subnet_1-update"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "description", "subnet_1_description-update"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "cidr", "192.168.199.0/24"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "dns_nameservers.0", "1.1.1.1"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "dns_nameservers.1", "3.3.3.3"),
					testAccCheckNetworkV2SubnetDNSLength(&subnet, 2),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "enable_dhcp", "false"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "allocation_pools.0.start", "192.168.199.100"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "allocation_pools.0.end", "192.168.199.200"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "host_routes.0.destination_cidr", "1.1.1.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "host_routes.0.next_hop", "192.168.199.1"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "host_routes.1.destination_cidr", "2.2.2.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "host_routes.1.next_hop", "192.168.199.1"),
					testAccCheckNetworkV2SubnetHostRoutesLength(&subnet, 2),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetDNSConsistency("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "name", ""),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "description", "subnet_1_description-update"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "cidr", "192.168.199.0/24"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "dns_nameservers.0", "3.3.3.3"),
					testAccCheckNetworkV2SubnetDNSLength(&subnet, 1),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "enable_dhcp", "false"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "allocation_pools.0.start", "192.168.199.100"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "allocation_pools.0.end", "192.168.199.200"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "host_routes.0.destination_cidr", "2.2.2.0/24"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "host_routes.0.next_hop", "192.168.199.1"),
					testAccCheckNetworkV2SubnetHostRoutesLength(&subnet, 1),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetUpdate3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetDNSConsistency("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "name", "name_1"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "description", ""),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "cidr", "192.168.199.0/24"),
					// TODO : length of dns_nameservers can not be set 0
					// testAccCheckNetworkV2SubnetDNSLength(&subnet, 0),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "enable_dhcp", "false"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "allocation_pools.0.start", "192.168.199.100"),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "allocation_pools.0.end", "192.168.199.200"),
					testAccCheckNetworkV2SubnetHostRoutesLength(&subnet, 0),
				),
			},
		},
	})
}

func TestAccNetworkV2Subnet_forceNewByCIDR(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var sn1 subnets.Subnet
	var sn2 subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetForceNew1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn1),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetForceNew2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn2),
					testAccCheckNetworkV2SubnetIDsDoNotMatch(&sn1, &sn2),
				),
			},
		},
	})
}

func TestAccNetworkV2Subnet_tag(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetDNSConsistency("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetTag(&subnet, "k1", "v1"),
					testAccCheckNetworkV2SubnetNoTagKey(&subnet, "k2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetTagUpdate1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetDNSConsistency("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetTag(&subnet, "k1", "v1"),
					testAccCheckNetworkV2SubnetTag(&subnet, "k2", "v2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetTagUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetDNSConsistency("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetTagLengthIsZERO(&subnet),
				),
			},
		},
	})
}
func TestAccNetworkV2Subnet_enableDHCP(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetEnableDHCP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "enable_dhcp", "true"),
				),
			},
		},
	})
}

func TestAccNetworkV2Subnet_disableDHCP(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetDisableDHCP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "enable_dhcp", "false"),
				),
			},
		},
	})
}

func TestAccNetworkV2Subnet_noGateway(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetNoGateway,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "gateway_ip", ""),
				),
			},
		},
	})
}

func TestAccNetworkV2Subnet_impliedGateway(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetImpliedGateway,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr(
						"ecl_network_subnet_v2.subnet_1", "gateway_ip", "192.168.199.1"),
				),
			},
		},
	})
}

func TestAccNetworkV2Subnet_conflictsGateway(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	pattern1 := "\"no_gateway\": conflicts with gateway_ip"
	pattern2 := "\"gateway_ip\": conflicts with no_gateway"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testValidateNetworkV2SubnetConflictsGateway,
				ExpectError: regexp.MustCompile(pattern1),
			},
			resource.TestStep{
				Config:      testValidateNetworkV2SubnetConflictsGateway,
				ExpectError: regexp.MustCompile(pattern2),
			},
		},
	})
}

func TestAccNetworkV2Subnet_timeout(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetTimeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
				),
			},
		},
	})
}

func TestAccNetworkV2Subnet_ntp(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SubnetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2SubnetNTPBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "ntp_servers.0", "1.1.1.1"),
					testAccCheckNetworkV2SubnetNTPServersLength(&subnet, 1),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetNTPUpdate1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "ntp_servers.0", "1.1.1.1"),
					resource.TestCheckResourceAttr("ecl_network_subnet_v2.subnet_1", "ntp_servers.1", "2.2.2.2"),
					testAccCheckNetworkV2SubnetNTPServersLength(&subnet, 2),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2SubnetNTPUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2SubnetNTPServersLength(&subnet, 0),
				),
			},
		},
	})
}

func testAccCheckNetworkV2SubnetHostRoutesLength(sn *subnets.Subnet, length int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(sn.HostRoutes) != length {
			return fmt.Errorf(
				"Tag host_routes length does not match. Actual is %d . Expected is %d",
				len(sn.HostRoutes), length)
		}
		return nil
	}
}

func testAccCheckNetworkV2SubnetDNSLength(sn *subnets.Subnet, length int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(sn.DNSNameservers) != length {
			return fmt.Errorf(
				"Tag dns_nameservers length does not match. Actual is %d . Expected is %d",
				len(sn.DNSNameservers), length)
		}
		return nil
	}
}

func testAccCheckNetworkV2SubnetNTPServersLength(sn *subnets.Subnet, length int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(sn.NTPServers) != length {
			return fmt.Errorf(
				"Tag ntp_servers length does not match. Actual is %d . Expected is %d",
				len(sn.NTPServers), length)
		}
		return nil
	}
}

func testAccCheckNetworkV2SubnetIDsDoNotMatch(sn1, sn2 *subnets.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if sn1.ID == sn2.ID {
			return fmt.Errorf("Network was not recreated")
		}

		return nil
	}
}

func testAccCheckNetworkV2SubnetTagLengthIsZERO(sn *subnets.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(sn.Tags) != 0 {
			return fmt.Errorf("Tag length is not ZERO")
		}
		return nil
	}
}

func testAccCheckNetworkV2SubnetTag(
	sn *subnets.Subnet, k string, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if sn.Tags == nil {
			return fmt.Errorf("No tag")
		}

		for key, value := range sn.Tags {
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

func testAccCheckNetworkV2SubnetNoTagKey(
	sn *subnets.Subnet, k string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if sn.Tags == nil {
			return nil
		}

		for key := range sn.Tags {
			if k == key {
				return fmt.Errorf("Tag found: %s", k)
			}
		}

		return nil
	}
}

func testAccCheckNetworkV2SubnetDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_subnet_v2" {
			continue
		}

		_, err := subnets.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Subnet still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2SubnetExists(n string, subnet *subnets.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL network client: %s", err)
		}

		found, err := subnets.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Subnet not found")
		}

		*subnet = *found

		return nil
	}
}

func testAccCheckNetworkV2SubnetDNSConsistency(n string, subnet *subnets.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		for i, dns := range subnet.DNSNameservers {
			if dns != rs.Primary.Attributes[fmt.Sprintf("dns_nameservers.%d", i)] {
				return fmt.Errorf("Dns Nameservers list elements or order is not consistent")
			}
		}

		return nil
	}
}

const testAccNetworkV2SubnetTag = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  tags = {
    k1 = "v1"
  }
}`

const testAccNetworkV2SubnetTagUpdate1 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  tags = {
    k1 = "v1",
    k2 = "v2"
  }
}`

const testAccNetworkV2SubnetTagUpdate2 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"
}`

const testAccNetworkV2SubnetBasic = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "subnet_1"
	description = "subnet_1_description"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  dns_nameservers = [
		"1.1.1.1", 
		"2.2.2.2"
	]

	enable_dhcp = "true"

	allocation_pools {
    start = "192.168.199.100"
    end = "192.168.199.200"
	}

	host_routes {
		destination_cidr = "1.1.1.0/24"
		next_hop = "192.168.199.1"
	}
}`

const testAccNetworkV2SubnetUpdate1 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "subnet_1-update"
	description = "subnet_1_description-update"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  dns_nameservers = [
		"1.1.1.1", 
		"3.3.3.3"
	]

	enable_dhcp = "false"

	allocation_pools {
    start = "192.168.199.100"
    end = "192.168.199.200"
	}

	host_routes {
		destination_cidr = "1.1.1.0/24"
		next_hop = "192.168.199.1"
	}
	host_routes {
		destination_cidr = "2.2.2.0/24"
		next_hop = "192.168.199.1"
	}
}
`

const testAccNetworkV2SubnetUpdate2 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = ""
	description = "subnet_1_description-update"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  dns_nameservers = [
		"3.3.3.3"
	]

	enable_dhcp = "false"

	allocation_pools {
    start = "192.168.199.100"
    end = "192.168.199.200"
	}

	host_routes {
		destination_cidr = "2.2.2.0/24"
		next_hop = "192.168.199.1"
	}
}
`
const testAccNetworkV2SubnetUpdate3 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "name_1"
	description = ""
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

	enable_dhcp = "false"

	allocation_pools {
    start = "192.168.199.100"
    end = "192.168.199.200"
	}
}
`

const testAccNetworkV2SubnetEnableDHCP = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  gateway_ip = "192.168.199.1"
  enable_dhcp = true
  network_id = "${ecl_network_network_v2.network_1.id}"
}
`

const testAccNetworkV2SubnetDisableDHCP = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  enable_dhcp = false
  network_id = "${ecl_network_network_v2.network_1.id}"
}
`

const testAccNetworkV2SubnetNoGateway = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  no_gateway = true
  network_id = "${ecl_network_network_v2.network_1.id}"
}
`

const testAccNetworkV2SubnetImpliedGateway = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}
resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"
}
`

const testAccNetworkV2SubnetTimeout = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  allocation_pools {
    start = "192.168.199.100"
    end = "192.168.199.200"
  }

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`

const testValidateNetworkV2SubnetConflictsGateway = `
resource "ecl_network_subnet_v2" "subnet_1" {
  cidr = "192.168.199.0/24"
	network_id = "dummy_network_id"
	gateway_ip = "192.168.1.1"
	no_gateway = "true"

}
`

const testAccNetworkV2SubnetForceNew1 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  no_gateway = true
  network_id = "${ecl_network_network_v2.network_1.id}"
}
`
const testAccNetworkV2SubnetForceNew2 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.198.0/24"
  no_gateway = true
  network_id = "${ecl_network_network_v2.network_1.id}"
}
`

const testAccNetworkV2SubnetNTPBasic = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "subnet_1"
	description = "subnet_1_description"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  dns_nameservers = [
		"1.1.1.1", 
		"2.2.2.2"
	]

  ntp_servers = [
		"1.1.1.1", 
	]

	enable_dhcp = "true"

	allocation_pools {
    start = "192.168.199.100"
    end = "192.168.199.200"
	}
}`

const testAccNetworkV2SubnetNTPUpdate1 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "subnet_1"
	description = "subnet_1_description"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  dns_nameservers = [
		"1.1.1.1", 
		"2.2.2.2"
	]

  ntp_servers = [
		"1.1.1.1", 
		"2.2.2.2",
	]

	enable_dhcp = "true"

	allocation_pools {
    start = "192.168.199.100"
    end = "192.168.199.200"
	}
}`

const testAccNetworkV2SubnetNTPUpdate2 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "subnet_1"
	description = "subnet_1_description"
  cidr = "192.168.199.0/24"
  network_id = "${ecl_network_network_v2.network_1.id}"

  dns_nameservers = [
		"1.1.1.1", 
		"2.2.2.2"
	]

	enable_dhcp = "true"

	allocation_pools {
    start = "192.168.199.100"
    end = "192.168.199.200"
	}
}`
