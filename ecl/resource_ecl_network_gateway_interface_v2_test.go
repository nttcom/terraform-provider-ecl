package ecl

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"

	"github.com/nttcom/eclcloud"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/gateway_interfaces"
)

func TestAccNetworkV2GatewayInterface_basic(t *testing.T) {
	var gatewayInterface gateway_interfaces.GatewayInterface
	rName := acctest.RandomWithPrefix("tf-acc-test")
	resourceName := "ecl_network_gateway_interface_v2.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckGatewayInterface(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2GatewayInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2GatewayInterfaceConfig(rName, "create", "created"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceExists(resourceName, &gatewayInterface),
					resource.TestCheckResourceAttr(resourceName, "name", rName+"-create"),
					resource.TestCheckResourceAttr(resourceName, "description", "created"),
					resource.TestCheckResourceAttrSet(resourceName, "internet_gw_id"),
					resource.TestCheckResourceAttr(resourceName, "gw_vipv4", "192.168.200.1"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "29"),
					resource.TestCheckResourceAttrSet(resourceName, "network_id"),
					resource.TestCheckResourceAttr(resourceName, "primary_ipv4", "192.168.200.2"),
					resource.TestCheckResourceAttr(resourceName, "secondary_ipv4", "192.168.200.3"),
					resource.TestCheckResourceAttr(resourceName, "service_type", "internet"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "vrid", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "azure_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "fic_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "gcp_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "interdc_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "vpn_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "gw_vipv6", ""),
					resource.TestCheckResourceAttr(resourceName, "primary_ipv6", ""),
					resource.TestCheckResourceAttr(resourceName, "secondary_ipv6", ""),
				),
			},
			{
				Config: testAccNetworkV2GatewayInterfaceConfig(rName, "update", "name and description are updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2GatewayInterfaceExists(resourceName, &gatewayInterface),
					resource.TestCheckResourceAttr(resourceName, "name", rName+"-update"),
					resource.TestCheckResourceAttr(resourceName, "description", "name and description are updated"),
					resource.TestCheckResourceAttrSet(resourceName, "internet_gw_id"),
					resource.TestCheckResourceAttr(resourceName, "gw_vipv4", "192.168.200.1"),
					resource.TestCheckResourceAttr(resourceName, "netmask", "29"),
					resource.TestCheckResourceAttrSet(resourceName, "network_id"),
					resource.TestCheckResourceAttr(resourceName, "primary_ipv4", "192.168.200.2"),
					resource.TestCheckResourceAttr(resourceName, "secondary_ipv4", "192.168.200.3"),
					resource.TestCheckResourceAttr(resourceName, "service_type", "internet"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "vrid", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "azure_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "fic_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "gcp_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "interdc_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "vpn_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "gw_vipv6", ""),
					resource.TestCheckResourceAttr(resourceName, "primary_ipv6", ""),
					resource.TestCheckResourceAttr(resourceName, "secondary_ipv6", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckNetworkV2GatewayInterfaceDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_gateway_interface_v2" {
			continue
		}

		if result := gateway_interfaces.Get(networkClient, rs.Primary.ID); result.Err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(result.Err, &e) {
				return nil
			}
			return fmt.Errorf("error getting ECL Gateway interface: %w", result.Err)
		}

		return fmt.Errorf("gateway interface (%s) still exists", rs.Primary.ID)
	}

	return nil
}

func testAccCheckNetworkV2GatewayInterfaceExists(n string, gatewayInterface *gateway_interfaces.GatewayInterface) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL network client: %w", err)
		}

		found, err := gateway_interfaces.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return fmt.Errorf("error getting ECL Gateway interface: %w", err)
		}

		*gatewayInterface = *found

		return nil
	}
}

func testAccNetworkV2GatewayInterfaceConfig(rName, nameSuffix, description string) string {
	return fmt.Sprintf(`
resource "ecl_network_network_v2" "test" {
    name = %[1]q
}

resource "ecl_network_subnet_v2" "test" {
    name = %[1]q
    cidr = "192.168.200.0/29"
    enable_dhcp = false
    no_gateway = true
    network_id = ecl_network_network_v2.test.id
}

data "ecl_network_internet_service_v2" "test" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "test" {
    name = %[1]q
    description = "test"
    internet_service_id = data.ecl_network_internet_service_v2.test.id
    qos_option_id = %[4]q
}

resource "ecl_network_gateway_interface_v2" "test" {
    description = %[3]q
    gw_vipv4 = "192.168.200.1"
    internet_gw_id = ecl_network_internet_gateway_v2.test.id
    name = "%[1]s-%[2]s"
    netmask = 29
    network_id = ecl_network_network_v2.test.id
    primary_ipv4 = "192.168.200.2"
    secondary_ipv4 = "192.168.200.3"
    service_type = "internet"
    vrid = 1
    depends_on = ["ecl_network_subnet_v2.test"]
}
`, rName, nameSuffix, description, OS_QOS_OPTION_ID_10M)
}
