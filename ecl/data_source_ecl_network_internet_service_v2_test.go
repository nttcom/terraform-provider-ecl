package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2InternetServiceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckInternetService(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2InternetServiceDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetServiceDataSourceID("data.ecl_network_internet_service_v2.internet_service_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_internet_service_v2.internet_service_1", "name", "Internet-Service-01"),
				),
			},
		},
	})
}

func TestAccNetworkV2InternetServiceDataSourceTestQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckInternetService(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2InternetServiceDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetServiceDataSourceID("data.ecl_network_internet_service_v2.internet_service_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetServiceDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetServiceDataSourceID("data.ecl_network_internet_service_v2.internet_service_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetServiceDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetServiceDataSourceID("data.ecl_network_internet_service_v2.internet_service_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetServiceDataSourceMinimalSubmaskLength,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetServiceDataSourceID("data.ecl_network_internet_service_v2.internet_service_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2InternetServiceDataSourceZone,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetServiceDataSourceID("data.ecl_network_internet_service_v2.internet_service_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2InternetServiceDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find internet service data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Internet service data source ID not set")
		}

		return nil
	}
}

func testAccReturnMinimalSubmaskLength(region string) int {
	var minimalSubmaskLength int
	switch region {
	case "jp1":
	case "jp3":
	case "jp4":
	case "jp5":
	case "lab3ec":
		minimalSubmaskLength = 26
		break
	default:
		minimalSubmaskLength = 28
	}
	return minimalSubmaskLength
}

var testAccNetworkV2InternetServiceDataSourceBasic = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
    name = "Internet-Service-01"
}
`)

var testAccNetworkV2InternetServiceDataSourceID = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
  internet_service_id = "%s"
}
`,
	OS_INTERNET_SERVICE_ID)

var testAccNetworkV2InternetServiceDataSourceName = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
  name = "Internet-Service-01"
}
`)
var testAccNetworkV2InternetServiceDataSourceDescription = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
    description = ""
}
`)

var testAccNetworkV2InternetServiceDataSourceMinimalSubmaskLength = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
  minimal_submask_length = %d
}
`,
	testAccReturnMinimalSubmaskLength(OS_REGION_NAME))

var testAccNetworkV2InternetServiceDataSourceZone = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
  zone = "%s"
}
`,
	OS_INTERNET_SERVICE_ZONE_NAME)
