package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v4/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/subnets"
	security "github.com/nttcom/eclcloud/v4/ecl/security_order/v3/network_based_device_single"
)

func TestAccSecurityV3NetworkBasedDeviceSingle_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var sd security.SingleDevice

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV3NetworkBasedDeviceSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"operating_mode", "FW"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"az_group", OS_DEFAULT_ZONE),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV3NetworkBasedDeviceSingleUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"operating_mode", "UTM"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"az_group", OS_DEFAULT_ZONE),
				),
			},
		},
	})
}

func TestAccSecurityV3NetworkBasedDeviceSingle_updateInterface(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var sd security.SingleDevice
	var n1, n2 networks.Network
	var sn1, sn2 subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV3NetworkBasedDeviceSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV3NetworkBasedDeviceSingleUpdateInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn1),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2", &sn2),

					// port 0
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.network_id", &n1.ID),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.subnet_id", &sn1.ID),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.comment", "port 0 comment"),

					// port 1
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.1.enable", "false"),
					// port 2
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.2.enable", "false"),

					// port 3
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.ip_address", "192.168.2.50"),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.network_id", &n2.ID),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.subnet_id", &sn2.ID),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.comment", "port 3 comment"),

					// port 4
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.4.enable", "false"),
					// port 5
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.5.enable", "false"),
					// port 6
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.6.enable", "false"),
				),
			},

			resource.TestStep{
				Config: testAccSecurityV3NetworkBasedDeviceSingleUpdateInterface2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn1),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2", &sn2),

					// port 0
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.enable", "false"),

					// port 1
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.1.enable", "false"),
					// port 2
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.2.enable", "false"),

					// port 3
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.enable", "false"),

					// port 4
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.4.enable", "false"),
					// port 5
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.5.enable", "false"),
					// port 6
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.6.enable", "false"),
				),
			},
		},
	})
}

func testAccCheckSecurityV3NetworkBasedDeviceSingleExists(n string, sd *security.SingleDevice) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.securityOrderV3Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL security client: %s", err)
		}

		found, err := getSingleDeviceByHostName(client, "UTM", rs.Primary.ID, OS_TENANT_ID)
		if err != nil {
			return err
		}

		if found.Cell[2] != rs.Primary.ID {
			return fmt.Errorf("Security single device not found")
		}

		*sd = found

		return nil
	}
}

func testAccCheckSecurityV3NetworkBasedDeviceSingleDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV3Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_network_based_device_single_v3" {
			continue
		}

		_, err := getSingleDeviceByHostName(client, "UTM", rs.Primary.ID, OS_TENANT_ID)

		if err == nil {
			return fmt.Errorf("Security single device still exists")
		}

	}

	return nil
}

var testAccSecurityV3NetworkBasedDeviceSingleBasic = fmt.Sprintf(`
resource "ecl_security_network_based_device_single_v3" "device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "%s"
}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testAccSecurityV3NetworkBasedDeviceSingleUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_device_single_v3" "device_1" {
	tenant_id = "%s"
	locale = "en"
	operating_mode = "UTM"
	license_kind = "08"
	az_group = "%s"
}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

const testAccSecurityV3NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet1 = `
resource "ecl_network_network_v2" "network_1" {
	name = "network_1_for_utm_single"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "subnet_1_for_utm_single"
	cidr = "192.168.1.0/24"
	network_id = "${ecl_network_network_v2.network_1.id}"
	gateway_ip = "192.168.1.1"
	allocation_pools {
		start = "192.168.1.100"
		end = "192.168.1.200"
	}
}
`
const testAccSecurityV3NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet2 = `
resource "ecl_network_network_v2" "network_2" {
	name = "network_2_for_utm_single"
}

resource "ecl_network_subnet_v2" "subnet_2" {
	name = "subnet_2_for_utm_single"
	cidr = "192.168.2.0/24"
	network_id = "${ecl_network_network_v2.network_2.id}"
	gateway_ip = "192.168.2.1"
	allocation_pools {
		start = "192.168.2.100"
		end = "192.168.2.200"
	}
}
`

var testAccSecurityV3NetworkBasedDeviceSingleUpdateInterface = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_device_single_v3" "device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "%s"

  port {
      enable = "true"
      ip_address = "192.168.1.50"
      ip_address_prefix = 24
      network_id = "${ecl_network_network_v2.network_1.id}"
      subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
      mtu = "1500"
      comment = "port 0 comment"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }
  
  port {
      enable = "true"
      ip_address = "192.168.2.50"
      ip_address_prefix = 24
      network_id = "${ecl_network_network_v2.network_2.id}"
      subnet_id = "${ecl_network_subnet_v2.subnet_2.id}"
      mtu = "1500"
      comment = "port 3 comment"
  }

  port {
    enable = "false"
  }
  port {
    enable = "false"
  }
  port {
    enable = "false"
  }

}
`,
	testAccSecurityV3NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet1,
	testAccSecurityV3NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet2,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testAccSecurityV3NetworkBasedDeviceSingleUpdateInterface2 = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_device_single_v3" "device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "%s"

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }
  
  port {
    enable = "false"
  }

  port {
    enable = "false"
  }
  port {
    enable = "false"
  }
  port {
    enable = "false"
  }

}
`,
	testAccSecurityV3NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet1,
	testAccSecurityV3NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet2,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)
