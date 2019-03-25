package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/ports"
)

func TestAccNetworkV2NetworkDataSourceTestQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceNetwork,
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceNetworkID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourcePlane,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceStatus,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceTenantID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkDataSourceID("data.ecl_network_network_v2.net"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "name", "tf_test_network"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_network_v2.net", "admin_state_up", "true"),
				),
			},
		},
	})
}

func TestAccNetworkV2NetworkDataSourceCreateResource(t *testing.T) {
	var port ports.Port

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceNetwork,
			},
			resource.TestStep{
				Config: testAccNetworkV2NetworkDataSourceCreateResource,
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

const testAccNetworkV2NetworkDataSourceNetwork = `
resource "ecl_network_network_v2" "net" {
		name = "tf_test_network"
		description = "tf_test_network_description"
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
}
`

var testAccNetworkV2NetworkDataSourceNetworkID = fmt.Sprintf(`
%s

data "ecl_network_network_v2" "net" {
	network_id = "${ecl_network_network_v2.net.id}"
}
`, testAccNetworkV2NetworkDataSourceNetwork)

var testAccNetworkV2NetworkDataSourceDescription = fmt.Sprintf(`
%s

data "ecl_network_network_v2" "net" {
	description = "${ecl_network_network_v2.net.description}"
}
`, testAccNetworkV2NetworkDataSourceNetwork)

var testAccNetworkV2NetworkDataSourceName = fmt.Sprintf(`
%s

data "ecl_network_network_v2" "net" {
	name = "${ecl_network_network_v2.net.name}"
}
`, testAccNetworkV2NetworkDataSourceNetwork)

var testAccNetworkV2NetworkDataSourcePlane = fmt.Sprintf(`
%s

data "ecl_network_network_v2" "net" {
	plane = "${ecl_network_network_v2.net.plane}"
}
`, testAccNetworkV2NetworkDataSourceNetwork)

var testAccNetworkV2NetworkDataSourceStatus = fmt.Sprintf(`
%s

data "ecl_network_network_v2" "net" {
	status = "${ecl_network_network_v2.net.status}"
}
`, testAccNetworkV2NetworkDataSourceNetwork)

var testAccNetworkV2NetworkDataSourceTenantID = fmt.Sprintf(`
%s

data "ecl_network_network_v2" "net" {
	tenant_id = "${ecl_network_network_v2.net.tenant_id}"
}
`, testAccNetworkV2NetworkDataSourceNetwork)

var testAccNetworkV2NetworkDataSourceCreateResource = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port" {
	network_id = "${data.ecl_network_network_v2.net.id}"
}

data "ecl_network_network_v2" "net" {
	name = "${ecl_network_network_v2.net.name}"
}
`, testAccNetworkV2NetworkDataSourceNetwork)
