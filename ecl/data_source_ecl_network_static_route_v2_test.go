package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2StaticRouteDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckStaticRoute(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceStaticRoute,
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteDataSourceID("data.ecl_network_static_route_v2.static_route_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_static_route_v2.static_route_1", "name", "Terraform_Test_Static_Route_01"),
				),
			},
		},
	})
}

func TestAccNetworkV2StaticRouteDataSourceTestQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckStaticRoute(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceStaticRoute,
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteDataSourceID("data.ecl_network_static_route_v2.static_route_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteDataSourceID("data.ecl_network_static_route_v2.static_route_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceDestination,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteDataSourceID("data.ecl_network_static_route_v2.static_route_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceInetGwID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteDataSourceID("data.ecl_network_static_route_v2.static_route_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceNexthop,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteDataSourceID("data.ecl_network_static_route_v2.static_route_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceServiceType,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteDataSourceID("data.ecl_network_static_route_v2.static_route_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2StaticRouteDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteDataSourceID("data.ecl_network_static_route_v2.static_route_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2StaticRouteDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find static route data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Static route data source ID not set")
		}

		return nil
	}
}

var testAccNetworkV2StaticRouteDataSourceStaticRoute = fmt.Sprintf(`
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
`,
	OS_QOS_OPTION_ID_10M)

var testAccNetworkV2StaticRouteDataSourceBasic = fmt.Sprintf(`
%s

data "ecl_network_static_route_v2" "static_route_1" {
    name = "${ecl_network_static_route_v2.static_route_1.name}"
}
`, testAccNetworkV2StaticRouteDataSourceStaticRoute)

var testAccNetworkV2StaticRouteDataSourceName = fmt.Sprintf(`
%s

data "ecl_network_static_route_v2" "static_route_1" {
  name = "Terraform_Test_Static_Route_01"
}
`, testAccNetworkV2StaticRouteDataSourceStaticRoute)

var testAccNetworkV2StaticRouteDataSourceDescription = fmt.Sprintf(`
%s

data "ecl_network_static_route_v2" "static_route_1" {
    description = "test_static_route1"
}
`, testAccNetworkV2StaticRouteDataSourceStaticRoute)

var testAccNetworkV2StaticRouteDataSourceDestination = fmt.Sprintf(`
%s

data "ecl_network_static_route_v2" "static_route_1" {
    destination = "${ecl_network_public_ip_v2.public_ip_1.cidr}"
}
`, testAccNetworkV2StaticRouteDataSourceStaticRoute)

var testAccNetworkV2StaticRouteDataSourceInetGwID = fmt.Sprintf(`
%s

data "ecl_network_static_route_v2" "static_route_1" {
	internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
}
`, testAccNetworkV2StaticRouteDataSourceStaticRoute)

var testAccNetworkV2StaticRouteDataSourceNexthop = fmt.Sprintf(`
%s

data "ecl_network_static_route_v2" "static_route_1" {
    nexthop = "192.168.200.1"
}
`, testAccNetworkV2StaticRouteDataSourceStaticRoute)

var testAccNetworkV2StaticRouteDataSourceServiceType = fmt.Sprintf(`
%s

data "ecl_network_static_route_v2" "static_route_1" {
    service_type = "internet"
}
`, testAccNetworkV2StaticRouteDataSourceStaticRoute)

var testAccNetworkV2StaticRouteDataSourceID = fmt.Sprintf(`
%s

data "ecl_network_static_route_v2" "static_route_1" {
  static_route_id = "${ecl_network_static_route_v2.static_route_1.id}"
}
`, testAccNetworkV2StaticRouteDataSourceStaticRoute)
