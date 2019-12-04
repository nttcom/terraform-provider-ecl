package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/public_ips"
)

func TestAccNetworkV2PublicIPBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var publicIP public_ips.PublicIP

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckPublicIP(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2PublicIPDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PublicIPBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PublicIPExists("ecl_network_public_ip_v2.public_ip_1", &publicIP),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_public_ip_v2.public_ip_1", "name", stringMaxLength),
					resource.TestCheckResourceAttr(
						"ecl_network_public_ip_v2.public_ip_1", "description", stringMaxLength),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPUpdate2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ecl_network_public_ip_v2.public_ip_1", "name", ""),
					resource.TestCheckResourceAttr(
						"ecl_network_public_ip_v2.public_ip_1", "description", ""),
				),
			},
		},
	})
}

func testAccCheckNetworkV2PublicIPDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_public_ip_v2" {
			continue
		}

		_, err := public_ips.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Internet gateway still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2PublicIPExists(n string, public_ip *public_ips.PublicIP) resource.TestCheckFunc {
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

		found, err := public_ips.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Internet gateway not found")
		}

		*public_ip = *found

		return nil
	}
}

var testAccNetworkV2PublicIPBasic = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	name = "Terraform_Test_Internet_Gateway_01"
	description = "test_internet_gateway"
	internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
	qos_option_id = "%s"
}

resource "ecl_network_public_ip_v2" "public_ip_1" {
    name = "Terraform_Test_Public_IP_01"
    description = "test_public_ip"
    internet_gw_id = "%s"
    submask_length = 32
}
`,
	OS_QOS_OPTION_ID_10M,
	"${ecl_network_internet_gateway_v2.internet_gateway_1.id}",
)

var testAccNetworkV2PublicIPUpdate = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	name = "Terraform_Test_Internet_Gateway_01"
	description = "test_internet_gateway"
	internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
	qos_option_id = "%s"
}

resource "ecl_network_public_ip_v2" "public_ip_1" {
	name = "%s"
	description = "%s"
	internet_gw_id = "%s"
	submask_length = 32
}
`,
	OS_QOS_OPTION_ID_10M,
	stringMaxLength,
	stringMaxLength,
	"${ecl_network_internet_gateway_v2.internet_gateway_1.id}")

var testAccNetworkV2PublicIPUpdate2 = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
	name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	name = "Terraform_Test_Internet_Gateway_01"
	description = "test_internet_gateway"
	internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
	qos_option_id = "%s"
}

resource "ecl_network_public_ip_v2" "public_ip_1" {
	name = ""
	description = ""
	internet_gw_id = "%s"
 	submask_length = 32
}
`,
	OS_QOS_OPTION_ID_10M,
	"${ecl_network_internet_gateway_v2.internet_gateway_1.id}",
)
