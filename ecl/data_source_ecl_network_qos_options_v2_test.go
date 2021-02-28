package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2QosOptionsDataSource_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2QosOptionsDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2QosOptionsDataSourceID("data.ecl_network_qos_options_v2.qos_options_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_qos_options_v2.qos_options_1", "name", "10Mbps-BestEffort"),
				),
			},
		},
	})
}

func TestAccNetworkV2QosOptionsDataSource_queries(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2QosOptionsDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2QosOptionsDataSourceID("data.ecl_network_qos_options_v2.qos_options_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2QosOptionsDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2QosOptionsDataSourceID("data.ecl_network_qos_options_v2.qos_options_2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2QosOptionsDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2QosOptionsDataSourceID("data.ecl_network_qos_options_v2.qos_options_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2QosOptionsDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find qos options data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Qos Options data source ID not set")
		}

		return nil
	}
}

var testAccNetworkV2QosOptionsDataSourceBasic = `
data "ecl_network_qos_options_v2" "qos_options_1" {
    name = "10Mbps-BestEffort"
}
`

var testAccNetworkV2QosOptionsDataSourceID = `
data "ecl_network_qos_options_v2" "qos_options_1" {
    name = "10Mbps-BestEffort"
}

data "ecl_network_qos_options_v2" "qos_options_2" {
  qos_options_id = "${data.ecl_network_qos_options_v2.qos_options_1.id}"
}
`

var testAccNetworkV2QosOptionsDataSourceName = `
data "ecl_network_qos_options_v2" "qos_options_1" {
  name = "10Mbps-BestEffort"
}
`

var testAccNetworkV2QosOptionsDataSourceDescription = `
data "ecl_network_qos_options_v2" "qos_options_1" {
    description = "10M-besteffort-menu"
}
`
