package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2InternetGatewayDataSource_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckInternetGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceInternetGateway,
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayDataSourceID("data.ecl_network_internet_gateway_v2.internet_gateway_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_internet_gateway_v2.internet_gateway_1", "name", "Terraform_Test_Internet_Gateway_01"),
				),
			},
		},
	})
}

func TestAccNetworkV2InternetGatewayDataSource_queries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckInternetGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceInternetGateway,
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceInternetServiceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayDataSourceID("data.ecl_network_internet_gateway_v2.internet_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceQoSOptionID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayDataSourceID("data.ecl_network_internet_gateway_v2.internet_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayDataSourceID("data.ecl_network_internet_gateway_v2.internet_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayDataSourceID("data.ecl_network_internet_gateway_v2.internet_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceStatus,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayDataSourceID("data.ecl_network_internet_gateway_v2.internet_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayDataSourceID("data.ecl_network_internet_gateway_v2.internet_gateway_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2InternetGatewayDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find internet gateway data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Internet gateway data source ID not set")
		}

		return nil
	}
}

var testAccNetworkV2InternetGatewayDataSourceInternetGateway = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	name = "Terraform_Test_Internet_Gateway_01"
	description = "test_internet_gateway"
	internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
	qos_option_id = "%s"
}
`,
	OS_QOS_OPTION_ID_10M)

var testAccNetworkV2InternetGatewayDataSourceBasic = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	name = "${ecl_network_internet_gateway_v2.internet_gateway_1.name}"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceInternetServiceID = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
}
`,
	testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceQoSOptionID = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	qos_option_id = "%s"
}
`,
	testAccNetworkV2InternetGatewayDataSourceInternetGateway,
	OS_QOS_OPTION_ID_10M)

var testAccNetworkV2InternetGatewayDataSourceName = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
  name = "Terraform_Test_Internet_Gateway_01"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceDescription = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	description = "test_internet_gateway"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceID = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
  internet_gateway_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceStatus = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
  status = "ACTIVE"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)
