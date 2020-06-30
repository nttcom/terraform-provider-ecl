package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2FICGatewayDataSource_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "description", ""),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "fic_service_id", "66a63898-32a5-4b9d-8925-f52be1d84764"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "id", "fc546cf7-1956-436b-a9b4-edc917e397cf"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "name", "F032000001492"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "qos_option_id", "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "status", "ACTIVE"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "tenant_id", OS_TENANT_ID),
				),
			},
		},
	})
}

func TestAccNetworkV2FICGatewayDataSource_queries(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceFICServiceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceID2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceQoSOptionID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceStatus,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceTenantID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2FICGatewayDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find fic gateway data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("FIC gateway data source ID not set")
		}

		return nil
	}
}

var testAccNetworkV2FICGatewayDataSourceBasic = `
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	name = "F032000001492"
}
`

var testAccNetworkV2FICGatewayDataSourceDescription = `
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	description = " "
}
`

var testAccNetworkV2FICGatewayDataSourceFICServiceID = `
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	fic_service_id = "66a63898-32a5-4b9d-8925-f52be1d84764"
}
`
var testAccNetworkV2FICGatewayDataSourceID = `
data "ecl_network_fic_gateway_v2" "fic_gateway_2" {
	fic_gateway_id = "fc546cf7-1956-436b-a9b4-edc917e397cf"
}
`

var testAccNetworkV2FICGatewayDataSourceID2 = `
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	name = "F032000001492"
}

data "ecl_network_fic_gateway_v2" "fic_gateway_2" {
	fic_gateway_id = "${data.ecl_network_fic_gateway_v2.fic_gateway_1.id}"
}
`

var testAccNetworkV2FICGatewayDataSourceName = `
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	name = "F032000001492"
}
`

var testAccNetworkV2FICGatewayDataSourceQoSOptionID = `
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	qos_option_id = "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd"
}
`

var testAccNetworkV2FICGatewayDataSourceStatus = `
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	status = "ACTIVE"
}
`

var testAccNetworkV2FICGatewayDataSourceTenantID = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	tenant_id = "%s"
}
`,
	OS_TENANT_ID)
