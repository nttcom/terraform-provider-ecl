package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2PublicIPDataSourceBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckPublicIP(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourcePublicIP,
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PublicIPDataSourceID("data.ecl_network_public_ip_v2.public_ip_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_public_ip_v2.public_ip_1", "name", "Terraform_Test_Public_IP_01"),
				),
			},
		},
	})
}

func TestAccNetworkV2PublicIPDataSourceTestQueries(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckPublicIP(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourcePublicIP,
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PublicIPDataSourceID("data.ecl_network_public_ip_v2.public_ip_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PublicIPDataSourceID("data.ecl_network_public_ip_v2.public_ip_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourceInternetGwID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PublicIPDataSourceID("data.ecl_network_public_ip_v2.public_ip_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourceStatus,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PublicIPDataSourceID("data.ecl_network_public_ip_v2.public_ip_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourceSubmaskLength,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PublicIPDataSourceID("data.ecl_network_public_ip_v2.public_ip_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2PublicIPDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2PublicIPDataSourceID("data.ecl_network_public_ip_v2.public_ip_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2PublicIPDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find public IP data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Public IP data source ID not set")
		}

		return nil
	}
}

var testAccNetworkV2PublicIPDataSourcePublicIP = fmt.Sprintf(`
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
	internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
	submask_length = 32
}
`,
	OS_QOS_OPTION_ID_10M)

var testAccNetworkV2PublicIPDataSourceBasic = fmt.Sprintf(`
%s

data "ecl_network_public_ip_v2" "public_ip_1" {
	name = "${ecl_network_public_ip_v2.public_ip_1.name}"
}
`, testAccNetworkV2PublicIPDataSourcePublicIP)

var testAccNetworkV2PublicIPDataSourceName = fmt.Sprintf(`
%s

data "ecl_network_public_ip_v2" "public_ip_1" {
  name = "Terraform_Test_Public_IP_01"
}
`, testAccNetworkV2PublicIPDataSourcePublicIP)

var testAccNetworkV2PublicIPDataSourceDescription = fmt.Sprintf(`
%s

data "ecl_network_public_ip_v2" "public_ip_1" {
	description = "test_public_ip"
}
`, testAccNetworkV2PublicIPDataSourcePublicIP)

var testAccNetworkV2PublicIPDataSourceInternetGwID = fmt.Sprintf(`
%s

data "ecl_network_public_ip_v2" "public_ip_1" {
	internet_gw_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
}
`, testAccNetworkV2PublicIPDataSourcePublicIP)

var testAccNetworkV2PublicIPDataSourceSubmaskLength = fmt.Sprintf(`
%s

data "ecl_network_public_ip_v2" "public_ip_1" {
	submask_length = 32
}
`, testAccNetworkV2PublicIPDataSourcePublicIP)

var testAccNetworkV2PublicIPDataSourceID = fmt.Sprintf(`
%s

data "ecl_network_public_ip_v2" "public_ip_1" {
  public_ip_id = "${ecl_network_public_ip_v2.public_ip_1.id}"
}
`, testAccNetworkV2PublicIPDataSourcePublicIP)

var testAccNetworkV2PublicIPDataSourceStatus = fmt.Sprintf(`
%s

data "ecl_network_public_ip_v2" "public_ip_1" {
  status = "ACTIVE"
}
`, testAccNetworkV2PublicIPDataSourcePublicIP)
