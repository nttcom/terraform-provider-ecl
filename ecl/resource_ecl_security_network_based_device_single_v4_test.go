package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v4/network_based_device_single"
)

func TestAccSecurityV4NetworkBasedDeviceSingle_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var sd security.SingleDevice

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV4NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV4NetworkBasedDeviceSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV4NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v4.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1",
						"operating_mode", "FW"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1",
						"az_group", OS_DEFAULT_ZONE),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV4NetworkBasedDeviceSingleUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV4NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v4.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1",
						"locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1",
						"operating_mode", "UTM"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1",
						"license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1",
						"az_group", OS_DEFAULT_ZONE),
				),
			},
		},
	})
}

func TestAccSecurityV4NetworkBasedDeviceSingle_updateInterface(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var sd security.SingleDevice

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV4NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV4NetworkBasedDeviceSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV4NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v4.device_1", &sd),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV4NetworkBasedDeviceSingleUpdateInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV4NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v4.device_1", &sd),

					// port 0
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.0.ip_address", "192.168.7.3"),
					resource.TestCheckResourceAttrPair(
						"ecl_security_network_based_device_single_v4.device_1", "port.0.network_id", "data.ecl_network_subnet_v2.subnet_1", "network_id"),
					resource.TestCheckResourceAttrPair(
						"ecl_security_network_based_device_single_v4.device_1", "port.0.subnet_id", "data.ecl_network_subnet_v2.subnet_1", "id"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.0.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.0.comment", "port 0 comment"),

					// port 1
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.1.enable", "false"),
					// port 2
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.2.enable", "false"),

					// port 3
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.3.ip_address", "192.168.8.3"),
					resource.TestCheckResourceAttrPair(
						"ecl_security_network_based_device_single_v4.device_1", "port.3.network_id", "data.ecl_network_subnet_v2.subnet_2", "network_id"),
					resource.TestCheckResourceAttrPair(
						"ecl_security_network_based_device_single_v4.device_1", "port.3.subnet_id", "data.ecl_network_subnet_v2.subnet_2", "id"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.3.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.3.comment", "port 3 comment"),

					// port 4
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.4.enable", "false"),
					// port 5
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.5.enable", "false"),
					// port 6
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.6.enable", "false"),
				),
			},

			resource.TestStep{
				Config: testAccSecurityV4NetworkBasedDeviceSingleUpdateInterface2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV4NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v4.device_1", &sd),

					// port 0
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.0.enable", "false"),

					// port 1
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.1.enable", "false"),
					// port 2
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.2.enable", "false"),

					// port 3
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.3.enable", "false"),

					// port 4
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.4.enable", "false"),
					// port 5
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.5.enable", "false"),
					// port 6
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v4.device_1", "port.6.enable", "false"),
				),
			},
		},
	})
}

func testAccCheckSecurityV4NetworkBasedDeviceSingleExists(n string, sd *security.SingleDevice) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.securityOrderV4Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL security client: %s", err)
		}

		found, err := getSingleDeviceByHostNameV4(client, "UTM", rs.Primary.ID, OS_TENANT_ID)
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

func testAccCheckSecurityV4NetworkBasedDeviceSingleDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV4Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_network_based_device_single_v4" {
			continue
		}

		_, err := getSingleDeviceByHostNameV4(client, "UTM", rs.Primary.ID, OS_TENANT_ID)

		if err == nil {
			return fmt.Errorf("Security single device still exists")
		}

	}

	return nil
}

var testAccSecurityV4NetworkBasedDeviceSingleBasic = fmt.Sprintf(`
resource "ecl_security_network_based_device_single_v4" "device_1" {
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

var testAccSecurityV4NetworkBasedDeviceSingleUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_device_single_v4" "device_1" {
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

const testAccSecurityV4NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet1 = `
data "ecl_network_subnet_v2" "subnet_1" {
	name = "IF-17729_inoue_subnet_v4-1"
	cidr = "192.168.7.0/29"
	network_id = "3ffe509c-cd11-4010-938f-0b85d8be800f"
	subnet_id = "84996ed7-d36f-4acc-abae-b22b4fd2349f"
}
`
const testAccSecurityV4NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet2 = `
data "ecl_network_subnet_v2" "subnet_2" {
	name = "IF-17729_inoue_subnet_v4-2"
	cidr = "192.168.8.0/29"
	network_id = "43fa9d4f-380b-4f4f-96f2-c23d92b23718"
	subnet_id = "72bb5418-ba9f-45cb-a0aa-dd040f322df2"
}
`

var testAccSecurityV4NetworkBasedDeviceSingleUpdateInterface = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_device_single_v4" "device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "%s"

  port {
      enable = "true"
      ip_address = "192.168.7.3"
      ip_address_prefix = 29
      network_id = "${data.ecl_network_subnet_v2.subnet_1.network_id}"
      subnet_id = "${data.ecl_network_subnet_v2.subnet_1.id}"
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
      ip_address = "192.168.8.3"
      ip_address_prefix = 29
      network_id = "${data.ecl_network_subnet_v2.subnet_2.network_id}"
      subnet_id = "${data.ecl_network_subnet_v2.subnet_2.id}"
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
	testAccSecurityV4NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet1,
	testAccSecurityV4NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet2,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testAccSecurityV4NetworkBasedDeviceSingleUpdateInterface2 = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_device_single_v4" "device_1" {
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
	testAccSecurityV4NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet1,
	testAccSecurityV4NetworkBasedDeviceSingleUpdateInterfaceNetworkSubnet2,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)
