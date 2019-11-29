package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccNetworkV2InternetServiceDataSourceBasic(t *testing.T) {
	if OS_REGION_NAME != "RegionOne" {
		t.Skipf("skip this test in %s region", OS_REGION_NAME)
	}

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

func TestMockedAccNetworkV2InternetServiceDataSourceTestQueries(t *testing.T) {
	if OS_REGION_NAME != "RegionOne" {
		t.Skipf("skip this test in %s region", OS_REGION_NAME)
	}

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
