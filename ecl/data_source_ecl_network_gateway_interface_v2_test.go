package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2GatewayInterfaceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckGatewayInterface(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface,
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_gateway_interface_v2.gateway_interface_1", "name", "Terraform_Test_Gateway_Interface_01"),
				),
			},
		},
	})
}

func TestAccNetworkV2GatewayInterfaceDataSourceTestQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckGatewayInterface(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface,
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceGwVipv4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceNetmask,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourcePrimaryIpv4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceSecondaryIpv4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceServiceType,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceVRID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceDataSourceID("data.ecl_network_gateway_interface_v2.gateway_interface_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2GatewayInterfaceDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find gateway interface data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Gateway interface data source ID not set")
		}

		return nil
	}
}

var testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface = fmt.Sprintf(`
resource "ecl_network_network_v2" "network_1" {
    name = "Terraform_Test_Network_01"
}

resource "ecl_network_subnet_v2" "subnet_1" {
    name = "Terraform_Test_Subnet_01"
    cidr = "192.168.200.0/29"
    enable_dhcp = false
    no_gateway = true
    network_id = "${ecl_network_network_v2.network_1.id}"
}

data "ecl_network_internet_service_v2" "internet_service_1" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
    name = "Terraform_Test_Internet_Gateway_01"
	description = "test_internet_gateway"
	internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
    qos_option_id = "%s"
}

resource "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    description = "test_gateway_interface"
    gw_vipv4 = "192.168.200.1"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    name = "Terraform_Test_Gateway_Interface_01"
    netmask = 29
    network_id = "${ecl_network_network_v2.network_1.id}"
    primary_ipv4 = "192.168.200.2"
    secondary_ipv4 = "192.168.200.3"
    service_type = "internet"
    vrid=1
    depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`,
	OS_QOS_OPTION_ID_10M)

var testAccNetworkV2GatewayInterfaceDataSourceBasic = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    name = "${ecl_network_gateway_interface_v2.gateway_interface_1.name}"
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourceName = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
  name = "Terraform_Test_Gateway_Interface_01"
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourceDescription = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    description = "test_gateway_interface"
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourceGwVipv4 = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    gw_vipv4 = "192.168.200.1"
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourceNetmask = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    netmask = 29
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourcePrimaryIpv4 = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    primary_ipv4 = "192.168.200.2"
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourceSecondaryIpv4 = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    secondary_ipv4 = "192.168.200.3"
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourceServiceType = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    service_type = "internet"
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourceVRID = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
    vrid = 1
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)

var testAccNetworkV2GatewayInterfaceDataSourceID = fmt.Sprintf(`
%s

data "ecl_network_gateway_interface_v2" "gateway_interface_1" {
  gateway_interface_id = "${ecl_network_gateway_interface_v2.gateway_interface_1.id}"
}
`, testAccNetworkV2GatewayInterfaceDataSourceGatewayInterface)
