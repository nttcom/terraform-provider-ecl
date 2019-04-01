package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/gateway_interfaces"
)

func TestAccNetworkV2GatewayInterfaceBasic(t *testing.T) {
	var gatewayInterface gateway_interfaces.GatewayInterface

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckGatewayInterface(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2GatewayInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceExists("ecl_network_gateway_interface_v2.gateway_interface_1", &gatewayInterface),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_gateway_interface_v2.gateway_interface_1", "name", stringMaxLength),
					resource.TestCheckResourceAttr(
						"ecl_network_gateway_interface_v2.gateway_interface_1", "description", stringMaxLength),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2GatewayInterfaceUpdate2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_gateway_interface_v2.gateway_interface_1", "name", ""),
					resource.TestCheckResourceAttr(
						"ecl_network_gateway_interface_v2.gateway_interface_1", "description", ""),
				),
			},
		},
	})
}

func TestAccNetworkV2GatewayInterfaceMultiGateway(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckGatewayInterface(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2GatewayInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccNetworkV2GatewayInterfaceMultiGateway,
				ExpectError: regexp.MustCompile("\"internet_gw_id\": conflicts with vpn_gw_id"),
			},
		},
	})
}

func testAccCheckNetworkV2GatewayInterfaceDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_gateway_interface_v2" {
			continue
		}

		_, err := gateway_interfaces.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Gateway interface still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2GatewayInterfaceExists(n string, gatewayInterface *gateway_interfaces.GatewayInterface) resource.TestCheckFunc {
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

		found, err := gateway_interfaces.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Internet gateway not found")
		}

		*gatewayInterface = *found

		return nil
	}
}

var testAccNetworkV2GatewayInterfaceBasic = fmt.Sprintf(`
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

var testAccNetworkV2GatewayInterfaceUpdate = fmt.Sprintf(`
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
    description = "%s"
    gw_vipv4 = "192.168.200.1"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    name = "%s"
    netmask = 29
    network_id = "${ecl_network_network_v2.network_1.id}"
    primary_ipv4 = "192.168.200.2"
    secondary_ipv4 = "192.168.200.3"
    service_type = "internet"
    vrid=1
    depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`,
	OS_QOS_OPTION_ID_10M,
	stringMaxLength,
	stringMaxLength)

var testAccNetworkV2GatewayInterfaceUpdate2 = fmt.Sprintf(`
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
	description = ""
	gw_vipv4 = "192.168.200.1"
	internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
	name = ""
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

var testAccNetworkV2GatewayInterfaceMultiGateway = fmt.Sprintf(`
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
	vpn_gw_id = "dummy_id"
	vrid=1
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`,
	OS_QOS_OPTION_ID_10M)
