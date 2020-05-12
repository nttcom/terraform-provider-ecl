package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/ports"
)

func TestAccNetworkV2NetworkDataSource_queries(t *testing.T) {
	var networkName = fmt.Sprintf("ACPTTEST%s-network", acctest.RandString(5))
	var networkDescription = fmt.Sprintf("ACPTTEST%s-network-description", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceNetwork(networkName, networkDescription),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceNetworkID(networkName, networkDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", networkName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceDescription(networkName, networkDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", networkName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceName(networkName, networkDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", networkName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourcePlane(networkName, networkDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", networkName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "plane", "data"),
				),
			},
		},
	})
}

func TestAccNetworkV2NetworkDataSource_createResource(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var port ports.Port
	var networkName = fmt.Sprintf("ACPTTEST%s-network", acctest.RandString(5))
	var networkDescription = fmt.Sprintf("ACPTTEST%s-network-description", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceNetwork(networkName, networkDescription),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceCreateResource(networkName, networkDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port", &port),
				),
			},
		},
	})
}

func testAccCheckNetworkV2NetworkDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find network data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network data source ID not set")
		}

		return nil
	}
}

func testAccNetworkV2NetworkDataSourceNetwork(networkName, networkDescription string) string {
	return fmt.Sprintf(`
		resource "ecl_network_network_v2" "net" {
			name = "%s"
			description = "%s"
			admin_state_up = "true"
			plane = "data"
			tags = {
				keyword1 = "value1"
			}
		}

		resource "ecl_network_subnet_v2" "subnet" {
			name = "tf_test_subnet"
			cidr = "192.168.199.0/24"
			no_gateway = true
			network_id = "${ecl_network_network_v2.net.id}"
		}`, networkName, networkDescription)
}

func testAccNetworkV2NetworkDataSourceNetworkID(networkName, networkDescription string) string {
	return fmt.Sprintf(`
		%s
	
		data "ecl_network_network_v2" "net" {
			network_id = "${ecl_network_network_v2.net.id}"
		}`, testAccNetworkV2NetworkDataSourceNetwork(networkName, networkDescription))
}

func testAccNetworkV2NetworkDataSourceDescription(networkName, networkDescription string) string {
	return fmt.Sprintf(`
		%s

		data "ecl_network_network_v2" "net" {
			description = "%s"
		}`, testAccNetworkV2NetworkDataSourceNetwork(networkName, networkDescription), networkDescription)
}

func testAccNetworkV2NetworkDataSourceName(networkName, networkDescription string) string {
	return fmt.Sprintf(`
		%s
	
		data "ecl_network_network_v2" "net" {
			name = "${ecl_network_network_v2.net.name}"
		}`, testAccNetworkV2NetworkDataSourceNetwork(networkName, networkDescription))
}

// 1. Firstly, this test case creates two network which has different plane, "data" and "storage".
//    These two networks have same but random name which is given as argument of this function.
// 2. Afterward, this case try to extract network by using above "random name" and specified plane,
//    so that Terraform can extract only one network as data source.
func testAccNetworkV2NetworkDataSourcePlane(networkName, networkDescription string) string {
	return fmt.Sprintf(`
		%s
	
		resource "ecl_network_network_v2" "net2" {
			name = "%s"
			description = "%s"
			admin_state_up = "true"
			plane = "storage"
			tags = {
				keyword1 = "value1"
			}
		}

		data "ecl_network_network_v2" "net" {
			name = "%s"
			plane = "${ecl_network_network_v2.net.plane}"
		}`,
		// fot top of Sprintf
		testAccNetworkV2NetworkDataSourceNetwork(networkName, networkDescription),
		// for network_v2.net2 section
		networkName,
		networkDescription,
		// for data section
		networkName,
	)
}

func testAccNetworkV2NetworkDataSourceCreateResource(networkName, networkDescription string) string {
	return fmt.Sprintf(`
		%s

		resource "ecl_network_port_v2" "port" {
			network_id = "${data.ecl_network_network_v2.net.id}"
		}

		data "ecl_network_network_v2" "net" { 
			name = "${ecl_network_network_v2.net.name}"
		}`, testAccNetworkV2NetworkDataSourceNetwork(networkName, networkDescription))
}
