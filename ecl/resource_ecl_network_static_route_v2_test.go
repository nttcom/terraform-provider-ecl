package ecl

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nttcom/eclcloud"

	"github.com/nttcom/eclcloud/ecl/network/v2/static_routes"
)

func TestAccNetworkV2StaticRoute_internet(t *testing.T) {
	var staticRoute static_routes.StaticRoute
	rName := acctest.RandomWithPrefix("tf-acc-test")
	resourceName := "ecl_network_static_route_v2.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStaticRouteInternet(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2StaticRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2StaticRouteInternetConfig(rName, "create", "created"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteExists(resourceName, &staticRoute),
					resource.TestCheckResourceAttr(resourceName, "name", rName+"-create"),
					resource.TestCheckResourceAttr(resourceName, "description", "created"),
					resource.TestCheckResourceAttrSet(resourceName, "destination"),
					resource.TestCheckResourceAttrSet(resourceName, "internet_gw_id"),
					resource.TestCheckResourceAttr(resourceName, "nexthop", "192.168.200.1"),
					resource.TestCheckResourceAttr(resourceName, "service_type", "internet"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "aws_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "azure_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "fic_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "gcp_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "interdc_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "vpn_gw_id", ""),
				),
			},
			{
				Config: testAccNetworkV2StaticRouteInternetConfig(rName, "update", "name and description are updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteExists(resourceName, &staticRoute),
					resource.TestCheckResourceAttr(resourceName, "name", rName+"-update"),
					resource.TestCheckResourceAttr(resourceName, "description", "name and description are updated"),
					resource.TestCheckResourceAttrSet(resourceName, "destination"),
					resource.TestCheckResourceAttrSet(resourceName, "internet_gw_id"),
					resource.TestCheckResourceAttr(resourceName, "nexthop", "192.168.200.1"),
					resource.TestCheckResourceAttr(resourceName, "service_type", "internet"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "aws_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "azure_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "fic_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "gcp_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "interdc_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "vpn_gw_id", ""),
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

func TestAccNetworkV2StaticRoute_fic(t *testing.T) {
	var staticRoute static_routes.StaticRoute
	rName := acctest.RandomWithPrefix("tf-acc-test")
	resourceName := "ecl_network_static_route_v2.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckStaticRouteFIC(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2StaticRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2StaticRouteFICConfig(rName, "create", "created"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteExists(resourceName, &staticRoute),
					resource.TestCheckResourceAttr(resourceName, "name", rName+"-create"),
					resource.TestCheckResourceAttr(resourceName, "description", "created"),
					resource.TestCheckResourceAttrSet(resourceName, "destination"),
					resource.TestCheckResourceAttrSet(resourceName, "fic_gw_id"),
					resource.TestCheckResourceAttr(resourceName, "nexthop", "192.168.200.1"),
					resource.TestCheckResourceAttr(resourceName, "service_type", "fic"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "aws_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "azure_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "gcp_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "interdc_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "internet_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "vpn_gw_id", ""),
				),
			},
			{
				Config: testAccNetworkV2StaticRouteFICConfig(rName, "update", "name and description are updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2StaticRouteExists(resourceName, &staticRoute),
					resource.TestCheckResourceAttr(resourceName, "name", rName+"-update"),
					resource.TestCheckResourceAttr(resourceName, "description", "name and description are updated"),
					resource.TestCheckResourceAttrSet(resourceName, "destination"),
					resource.TestCheckResourceAttrSet(resourceName, "fic_gw_id"),
					resource.TestCheckResourceAttr(resourceName, "nexthop", "192.168.200.1"),
					resource.TestCheckResourceAttr(resourceName, "service_type", "fic"),
					resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "aws_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "azure_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "gcp_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "interdc_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "internet_gw_id", ""),
					resource.TestCheckResourceAttr(resourceName, "vpn_gw_id", ""),
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

func testAccCheckNetworkV2StaticRouteDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_static_route_v2" {
			continue
		}

		if result := static_routes.Get(networkClient, rs.Primary.ID); result.Err != nil {
			var e eclcloud.ErrDefault404
			if errors.As(result.Err, &e) {
				return nil
			}
			return fmt.Errorf("error getting ECL Staic route: %w", result.Err)
		}

		return fmt.Errorf("static route (%s) still exists", rs.Primary.ID)
	}

	return nil
}

func testAccCheckNetworkV2StaticRouteExists(n string, staticRoute *static_routes.StaticRoute) resource.TestCheckFunc {
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

		found, err := static_routes.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return fmt.Errorf("error getting ECL Static route: %w", err)
		}

		*staticRoute = *found

		return nil
	}
}

func testAccNetworkV2StaticRouteInternetConfig(rName, nameSuffix, description string) string {
	return fmt.Sprintf(`
resource "ecl_network_network_v2" "test" {
    name = %[1]q
}

resource "ecl_network_subnet_v2" "test" {
    name = %[1]q
    cidr = "192.168.200.0/29"
    enable_dhcp = false
    no_gateway = true
    network_id = "${ecl_network_network_v2.test.id}"
}

data "ecl_network_internet_service_v2" "test" {
    name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "test" {
    name = %[1]q
    description = "test_internet_gateway"
    internet_service_id = "${data.ecl_network_internet_service_v2.test.id}"
    qos_option_id = %[4]q
}

resource "ecl_network_gateway_interface_v2" "test" {
    description = "test_gateway_interface"
    gw_vipv4 = "192.168.200.1"
    internet_gw_id = "${ecl_network_internet_gateway_v2.test.id}"
    name = %[1]q
    netmask = 29
    network_id = "${ecl_network_network_v2.test.id}"
    primary_ipv4 = "192.168.200.2"
    secondary_ipv4 = "192.168.200.3"
    service_type = "internet"
    vrid = 1
    depends_on = ["ecl_network_subnet_v2.test"]
}

resource "ecl_network_public_ip_v2" "test" {
    name = %[1]q
    description = "test_public_ip"
    internet_gw_id = "${ecl_network_internet_gateway_v2.test.id}"
    submask_length = 32
    depends_on = ["ecl_network_gateway_interface_v2.test"]
}

resource "ecl_network_static_route_v2" "test" {
    description = %[3]q
    destination = "${ecl_network_public_ip_v2.test.cidr}"
    internet_gw_id = "${ecl_network_internet_gateway_v2.test.id}"
    name = "%[1]s-%[2]s"
    nexthop = "192.168.200.1"
    service_type = "internet"
    depends_on = ["ecl_network_gateway_interface_v2.test",
                  "ecl_network_public_ip_v2.test"]
}
`, rName, nameSuffix, description, OS_QOS_OPTION_ID_10M)
}

func testAccNetworkV2StaticRouteFICConfig(rName, nameSuffix, description string) string {
	return fmt.Sprintf(`
resource "ecl_network_network_v2" "test" {
    name = %[1]q
}

resource "ecl_network_subnet_v2" "test" {
    name = %[1]q
    cidr = "192.168.200.0/29"
    enable_dhcp = false
    no_gateway = true
    network_id = "${ecl_network_network_v2.test.id}"
}

data "ecl_network_internet_service_v2" "test" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "test" {
    name = %[1]q
    description = "test_internet_gateway"
    internet_service_id = "${data.ecl_network_internet_service_v2.test.id}"
    qos_option_id = %[4]q
}

resource "ecl_network_gateway_interface_v2" "test" {
    description = "test_gateway_interface"
    gw_vipv4 = "192.168.200.1"
    fic_gw_id = %[5]q
    name = %[1]q
    netmask = 29
    network_id = "${ecl_network_network_v2.test.id}"
    primary_ipv4 = "192.168.200.2"
    secondary_ipv4 = "192.168.200.3"
    service_type = "fic"
    vrid = 1
    depends_on = ["ecl_network_subnet_v2.test"]
}

resource "ecl_network_public_ip_v2" "test" {
    name = %[1]q
    description = "test_public_ip"
    internet_gw_id = "${ecl_network_internet_gateway_v2.test.id}"
    submask_length = 32
    depends_on = ["ecl_network_gateway_interface_v2.test"]
}

resource "ecl_network_static_route_v2" "test" {
    description = %[3]q
    destination = "${ecl_network_public_ip_v2.test.cidr}"
    fic_gw_id = %[5]q
    name = "%[1]s-%[2]s"
    nexthop = "192.168.200.1"
    service_type = "fic"
    depends_on = ["ecl_network_gateway_interface_v2.test",
                  "ecl_network_public_ip_v2.test"]
}
`, rName, nameSuffix, description, OS_QOS_OPTION_ID_10M, OS_FIC_GW_ID)
}
