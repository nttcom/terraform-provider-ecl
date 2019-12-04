package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/static_routes"
)

func TestAccNetworkV2StaticRouteBasic(t *testing.T) {
	var staticRoute static_routes.StaticRoute

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStaticRoute(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2StaticRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteExists("ecl_network_static_route_v2.static_route_1", &staticRoute),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_static_route_v2.static_route_1", "name", stringMaxLength),
					resource.TestCheckResourceAttr(
						"ecl_network_static_route_v2.static_route_1", "description", stringMaxLength),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteUpdate2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_static_route_v2.static_route_1", "name", ""),
					resource.TestCheckResourceAttr(
						"ecl_network_static_route_v2.static_route_1", "description", ""),
				),
			},
		},
	})
}

func TestAccNetworkV2StaticRouteMultiGateway(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStaticRoute(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2StaticRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccNetworkV2StaticRouteMultiGateway,
				ExpectError: regexp.MustCompile("\"internet_gw_id\": conflicts with vpn_gw_id"),
			},
			resource.TestStep{
				Config:      testAccNetworkV2StaticRouteMultiGateway,
				ExpectError: regexp.MustCompile("\"vpn_gw_id\": conflicts with internet_gw_id"),
			},
		},
	})
}

func testAccCheckNetworkV2StaticRouteDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_static_route_v2" {
			continue
		}

		_, err := static_routes.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Static route still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2StaticRouteExists(n string, staticRoute *static_routes.StaticRoute) resource.TestCheckFunc {
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

		found, err := static_routes.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Static route not found")
		}

		*staticRoute = *found

		return nil
	}
}

var testAccNetworkV2StaticRoute = fmt.Sprintf(`
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
    description = "test_gateway_interface1"
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

resource "ecl_network_public_ip_v2" "public_ip_1" {
  name = "Terraform_Test_Public_IP_01"
  description = "test_public_ip1"
  internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
  submask_length = 32
}
`,
	OS_QOS_OPTION_ID_10M,
)

var testAccNetworkV2StaticRouteBasic = fmt.Sprintf(`
%s

resource "ecl_network_static_route_v2" "static_route_1" {
    description = "test_static_route1"
    destination = "${ecl_network_public_ip_v2.public_ip_1.cidr}"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    name = "Terraform_Test_Static_Route_01"
    nexthop = "192.168.200.1"
    service_type = "internet"
	depends_on = ["ecl_network_gateway_interface_v2.gateway_interface_1",
				  "ecl_network_public_ip_v2.public_ip_1"]
}
`, testAccNetworkV2StaticRoute)

var testAccNetworkV2StaticRouteUpdate = fmt.Sprintf(`
%s

resource "ecl_network_static_route_v2" "static_route_1" {
    description = "%s"
    destination = "${ecl_network_public_ip_v2.public_ip_1.cidr}"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    name = "%s"
    nexthop = "192.168.200.1"
    service_type = "internet"
	depends_on = ["ecl_network_gateway_interface_v2.gateway_interface_1",
				  "ecl_network_public_ip_v2.public_ip_1"]
}
`,
	testAccNetworkV2StaticRoute,
	stringMaxLength,
	stringMaxLength)

var testAccNetworkV2StaticRouteUpdate2 = fmt.Sprintf(`
%s

resource "ecl_network_static_route_v2" "static_route_1" {
    description = ""
    destination = "${ecl_network_public_ip_v2.public_ip_1.cidr}"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    name = ""
    nexthop = "192.168.200.1"
    service_type = "internet"
	depends_on = ["ecl_network_gateway_interface_v2.gateway_interface_1",
				  "ecl_network_public_ip_v2.public_ip_1"]
}
`, testAccNetworkV2StaticRoute)

var testAccNetworkV2StaticRouteMultiGateway = fmt.Sprintf(`
%s

resource "ecl_network_static_route_v2" "static_route_1" {
    description = "test_static_route1"
    destination = "${ecl_network_public_ip_v2.public_ip_1.cidr}"
    internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
    name = "Terraform_Test_Static_Route_01"
    nexthop = "192.168.200.1"
	service_type = "internet"
	vpn_gw_id = "dummy_id"
	depends_on = ["ecl_network_gateway_interface_v2.gateway_interface_1",
				  "ecl_network_public_ip_v2.public_ip_1"]
}
`, testAccNetworkV2StaticRoute)
