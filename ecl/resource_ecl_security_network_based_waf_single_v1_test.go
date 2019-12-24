package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/ecl/network/v2/subnets"
	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_single"
)

func TestAccSecurityV1NetworkBasedWAFSingle_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var sd security.SingleDevice

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedWAFSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedWAFSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v1.waf_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "operating_mode", "WAF"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "az_group", "zone1-groupb"),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedWAFSingleUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v1.waf_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "operating_mode", "WAF"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "az_group", "zone1-groupb"),
				),
			},
		},
	})
}

func TestAccSecurityV1NetworkBasedWAFSingle_updateInterface(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var sd security.SingleDevice
	var n1, n2 networks.Network
	var sn1, sn2 subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedWAFSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedWAFSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v1.waf_1", &sd),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedWAFSingleUpdateInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v1.waf_1", &sd),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn1),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2", &sn2),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "port.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_waf_single_v1.waf_1", "port.0.network_id", &n1.ID),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_waf_single_v1.waf_1", "port.0.subnet_id", &sn1.ID),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "port.0.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "port.0.comment", "port 0 comment"),
				),
			},

			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedWAFSingleUpdateInterface2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v1.waf_1", &sd),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn1),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2", &sn2),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v1.waf_1", "port.0.enable", "false"),
				),
			},
		},
	})
}

func testAccCheckSecurityV1NetworkBasedWAFSingleExists(n string, sd *security.SingleDevice) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.securityOrderV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL security client: %s", err)
		}

		found, err := getSingleDeviceByHostName(client, "WAF", rs.Primary.ID)
		if err != nil {
			return err
		}

		if found.Cell[2] != rs.Primary.ID {
			return fmt.Errorf("Security single WAF not found")
		}

		*sd = found

		return nil
	}
}

func testAccCheckSecurityV1NetworkBasedWAFSingleDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_network_based_waf_single_v1" {
			continue
		}

		_, err := getSingleDeviceByHostName(client, "WAF", rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Security single WAF still exists")
		}

	}

	return nil
}

var testAccSecurityV1NetworkBasedWAFSingleBasic = fmt.Sprintf(`
resource "ecl_security_network_based_waf_single_v1" "waf_1" {
	tenant_id = "%s"
	locale = "ja"
	license_kind = "02"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)

var testAccSecurityV1NetworkBasedWAFSingleUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_waf_single_v1" "waf_1" {
	tenant_id = "%s"
	locale = "en"
	license_kind = "08"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)

const testAccSecurityV1NetworkBasedWAFSingleUpdateInterfaceNetworkSubnet1 = `
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
const testAccSecurityV1NetworkBasedWAFSingleUpdateInterfaceNetworkSubnet2 = `
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

var testAccSecurityV1NetworkBasedWAFSingleUpdateInterface = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_waf_single_v1" "waf_1" {
	tenant_id = "%s"
	locale = "ja"
	license_kind = "02"
	az_group = "zone1-groupb"

    port {
        enable = "true"
        ip_address = "192.168.1.50"
        ip_address_prefix = 24
        network_id = "${ecl_network_network_v2.network_1.id}"
        subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
        mtu = "1500"
        comment = "port 0 comment"
    }
}
`,
	testAccSecurityV1NetworkBasedWAFSingleUpdateInterfaceNetworkSubnet1,
	testAccSecurityV1NetworkBasedWAFSingleUpdateInterfaceNetworkSubnet2,
	OS_TENANT_ID,
)

var testAccSecurityV1NetworkBasedWAFSingleUpdateInterface2 = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_waf_single_v1" "waf_1" {
    tenant_id = "%s"
    locale = "ja"
    license_kind = "02"
    az_group = "zone1-groupb"

    port {
        enable = "false"
    }
}
`,
	testAccSecurityV1NetworkBasedWAFSingleUpdateInterfaceNetworkSubnet1,
	testAccSecurityV1NetworkBasedWAFSingleUpdateInterfaceNetworkSubnet2,
	OS_TENANT_ID,
)
