package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/compute/v2/servers"
	"github.com/nttcom/eclcloud/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/ecl/network/v2/ports"
	"github.com/nttcom/eclcloud/ecl/network/v2/subnets"
)

func TestAccNetworkV2Network_basic(t *testing.T) {
	var network networks.Network

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2NetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "name", "network_1"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "description", "network_1_description"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "plane", "data"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "admin_state_up", "true"),
					testAccCheckNetworkV2NetworkTag(&network, "k1", "v1"),
					testAccCheckNetworkV2NetworkNoTagKey(&network, "k2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkUpdate1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "name", "network_1-update"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "description", "network_1_description-update"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "plane", "data"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "admin_state_up", "false"),
					testAccCheckNetworkV2NetworkTag(&network, "k1", "v1"),
					testAccCheckNetworkV2NetworkTag(&network, "k2", "v2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "name", ""),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "description", "network_1_description-update"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "plane", "data"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "admin_state_up", "false"),
					testAccCheckNetworkV2NetworkNoTagKey(&network, "k1"),
					testAccCheckNetworkV2NetworkTag(&network, "k2", "v2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkUpdate3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "name", "name_1"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "description", ""),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "plane", "data"),
					resource.TestCheckResourceAttr("ecl_network_network_v2.network_1", "admin_state_up", "false"),
					testAccCheckNetworkV2NetworkTagLengthIsZERO(&network),
				),
			},
		},
	})
}

func TestAccNetworkV2Network_planeForceNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var n1 networks.Network
	var n2 networks.Network

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2NetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkForceNew1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkForceNew2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n2),
					testAccCheckNetworkV2NetworkIDsDoNotMatch(&n1, &n2),
				),
			},
		},
	})
}
func TestAccNetworkV2Network_netstack(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2NetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkNetstack,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2NetworkHasSubnets("ecl_network_network_v2.network_1", 1),
				),
			},
		},
	})
}

func TestAccNetworkV2Network_fullstack(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server
	var network networks.Network
	var port ports.Port
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2NetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkFullstack,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &subnet),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
				),
			},
		},
	})
}

func TestAccNetworkV2Network_withTag(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2NetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkWithTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					testAccCheckNetworkV2NewtorkHasTag("ecl_network_network_v2.network_1", "sample_key", &network),
				),
			},
		},
	})
}

func TestAccNetworkV2Network_timeout(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var network networks.Network

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2NetworkDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkTimeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
				),
			},
		},
	})
}

func testAccCheckNetworkV2NetworkIDsDoNotMatch(n1, n2 *networks.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if n1.ID == n2.ID {
			return fmt.Errorf("Network was not recreated")
		}

		return nil
	}
}

func testAccCheckNetworkV2NetworkDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_network_v2" {
			continue
		}

		_, err := networks.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Network still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2NetworkExists(n string, network *networks.Network) resource.TestCheckFunc {
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

		found, err := networks.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Network not found")
		}

		*network = *found

		return nil
	}
}

func testAccCheckNetworkV2NewtorkHasTag(n string, key string, network *networks.Network) resource.TestCheckFunc {
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

		found, err := networks.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if _, ok := found.Tags[key]; !ok {
			return fmt.Errorf("Certain tag not found: %s", key)
		}

		return nil
	}
}

func testAccCheckNetworkV2NetworkHasSubnets(n string, numOfSubnets int) resource.TestCheckFunc {
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

		found, err := networks.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if len(found.Subnets) != numOfSubnets {
			return fmt.Errorf("Network doesn't have certain number of subnets: %+v", found)
		}

		return nil
	}
}

func testAccCheckNetworkV2NetworkTagLengthIsZERO(n *networks.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(n.Tags) != 0 {
			return fmt.Errorf("Tag length is not ZERO")
		}
		return nil
	}
}

func testAccCheckNetworkV2NetworkTag(
	n *networks.Network, k string, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if n.Tags == nil {
			return fmt.Errorf("No tag")
		}

		for key, value := range n.Tags {
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

func testAccCheckNetworkV2NetworkNoTagKey(
	n *networks.Network, k string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if n.Tags == nil {
			return nil
		}

		for key := range n.Tags {
			if k == key {
				return fmt.Errorf("Tag found: %s", k)
			}
		}

		return nil
	}
}

const testAccNetworkV2NetworkBasic = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  description = "network_1_description"
  plane = "data"
	admin_state_up = "true"
	tags = {
		k1 = "v1"
	}
}
`
const testAccNetworkV2NetworkUpdate1 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1-update"
  description = "network_1_description-update"
  plane = "data"
	admin_state_up = "false"
	tags = {
		k1 = "v1"
		k2 = "v2"
	}
}
`
const testAccNetworkV2NetworkUpdate2 = `
resource "ecl_network_network_v2" "network_1" {
  name = ""
  description = "network_1_description-update"
  plane = "data"
	admin_state_up = "false"
	tags = {
		k2 = "v2"
	}
}
`
const testAccNetworkV2NetworkUpdate3 = `
resource "ecl_network_network_v2" "network_1" {
  name = "name_1"
  description = ""
  plane = "data"
	admin_state_up = "false"
}
`

const testAccNetworkV2NetworkForceNew1 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  description = "network_1_description"
  plane = "data"
	admin_state_up = "true"
}
`
const testAccNetworkV2NetworkForceNew2 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  description = "network_1_description"
  plane = "storage"
  admin_state_up = "true"
}
`

const testAccNetworkV2NetworkNetstack = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.10.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
}
`

const testAccNetworkV2NetworkFullstack = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  cidr = "192.168.199.0/24"
  ip_version = 4
  network_id = "${ecl_network_network_v2.network_1.id}"
  depends_on = ["ecl_network_network_v2.network_1"]
}

resource "ecl_network_port_v2" "port_1" {
  name = "port_1"
  admin_state_up = "true"
  network_id = "${ecl_network_network_v2.network_1.id}"

  fixed_ip {
    subnet_id =  "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address =  "192.168.199.23"
  }
	
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
}

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"

  network {
    port = "${ecl_network_port_v2.port_1.id}"
  }

  depends_on = ["ecl_network_port_v2.port_1"]
}
`

const testAccNetworkV2NetworkWithTag = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"
  tags = {
    "sample_key" = "sample_value"
  }
}
`

const testAccNetworkV2NetworkTimeout = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  admin_state_up = "true"

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`
