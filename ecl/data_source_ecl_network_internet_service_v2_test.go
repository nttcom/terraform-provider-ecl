package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
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

func TestMockedAccNetworkV2InternetServiceDataSourceBasic(t *testing.T) {

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListNameQuery)

	mc.StartServer(t)

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

func TestMockedAccNetworkV2InternetServiceDataSourceTestQueries(t *testing.T) {

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListNameQuery)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListIDQuery)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListDescriptionQuery)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListMinimalSubmaskLengthQuery)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListZoneQuery)

	mc.StartServer(t)

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
	minimal_submask_length := 26
	if region == "lab4ec" {
		minimal_submask_length = 28
	}

	return minimal_submask_length
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

var testMockNetworkV2InternetServiceListNameQuery = `
request:
    method: GET
    Query:
      name:
        - Internet-Service-01
response: 
    code: 200
    body: >
        {
            "internet_services": [
                {
                    "description": "Internet-Service-01",
                    "id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                    "minimal_submask_length": 26,
                    "name": "Internet-Service-01",
                    "zone": "jp1-zone1"
                }
            ]
        }
`

var testMockNetworkV2InternetServiceListIDQuery = `
request:
    method: GET
    Query:
      id:
        - a7791c79-19b0-4eb6-9a8f-ea739b44e8d5
response: 
    code: 200
    body: >
        {
            "internet_services": [
                {
                    "description": "Internet-Service-01",
                    "id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                    "minimal_submask_length": 26,
                    "name": "Internet-Service-01",
                    "zone": "jp1-zone1"
                }
            ]
        }
`

var testMockNetworkV2InternetServiceListDescriptionQuery = `
request:
    method: GET
    Query:
      decsription:
        - Internet-Service-01
response: 
    code: 200
    body: >
        {
            "internet_services": [
                {
                    "description": "Internet-Service-01",
                    "id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                    "minimal_submask_length": 26,
                    "name": "Internet-Service-01",
                    "zone": "jp1-zone1"
                }
            ]
        }
`

var testMockNetworkV2InternetServiceListMinimalSubmaskLengthQuery = `
request:
    method: GET
    Query:
      minimal_submask_length:
        - 26
response: 
    code: 200
    body: >
        {
            "internet_services": [
                {
                    "description": "Internet-Service-01",
                    "id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                    "minimal_submask_length": 26,
                    "name": "Internet-Service-01",
                    "zone": "jp1-zone1"
                }
            ]
        }
`

var testMockNetworkV2InternetServiceListZoneQuery = `
request:
    method: GET
    Query:
      zone:
        - jp1-zone1
response: 
    code: 200
    body: >
        {
            "internet_services": [
                {
                    "description": "Internet-Service-01",
                    "id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                    "minimal_submask_length": 26,
                    "name": "Internet-Service-01",
                    "zone": "jp1-zone1"
                }
            ]
        }
`
