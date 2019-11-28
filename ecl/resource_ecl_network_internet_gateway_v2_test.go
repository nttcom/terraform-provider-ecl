package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nttcom/eclcloud/ecl/network/v2/internet_gateways"
)

func TestAccNetworkV2InternetGatewayBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var internet_gateway internet_gateways.InternetGateway

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckInternetGateway(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2InternetGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayExists("ecl_network_internet_gateway_v2.internet_gateway_1", &internet_gateway),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_internet_gateway_v2.internet_gateway_1", "name", stringMaxLength),
					resource.TestCheckResourceAttr(
						"ecl_network_internet_gateway_v2.internet_gateway_1", "description", stringMaxLength),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayUpdate2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_internet_gateway_v2.internet_gateway_1", "name", ""),
					resource.TestCheckResourceAttr(
						"ecl_network_internet_gateway_v2.internet_gateway_1", "description", ""),
				),
			},
		},
	})
}

func testAccCheckNetworkV2InternetGatewayDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_internet_gateway_v2" {
			continue
		}

		_, err := internet_gateways.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Internet gateway still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2InternetGatewayExists(n string, internet_gateway *internet_gateways.InternetGateway) resource.TestCheckFunc {
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

		found, err := internet_gateways.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Internet gateway not found")
		}

		*internet_gateway = *found

		return nil
	}
}

var testAccNetworkV2InternetGatewayBasic = fmt.Sprintf(`
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

var testAccNetworkV2InternetGatewayUpdate = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
    name = "%s"
    description = "%s"
    internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
    qos_option_id = "%s"
}
`,
	stringMaxLength,
	stringMaxLength,
	OS_QOS_OPTION_ID_100M)

var testAccNetworkV2InternetGatewayUpdate2 = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
    name = ""
    description = ""
    internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
    qos_option_id = "%s"
}
`,
	OS_QOS_OPTION_ID_10M)
