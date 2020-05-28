package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccNetworkV2InternetGatewayDataSource_basic(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListNameQuery)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayPost)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetBasic)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingCreate)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingDelete)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetDeleted)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayDelete)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListNameQuery)

	mc.StartServer(t)

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

func TestMockedAccNetworkV2InternetGatewayDataSource_queries(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListNameQuery)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayPost)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetBasic)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingCreate)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingDelete)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetDeleted)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayDelete)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListNameQuery)

	mc.StartServer(t)

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

var testMockNetworkV2InternetGatewayListNameQuery = `
request:
    method: GET
    Query:
      name:
        - Terraform_Test_Internet_Gateway_01
response: 
    code: 200
    body: >
        {
            "internet_gateways": [
                {
                    "description": "test_internet_gateway",
                    "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                    "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                    "name": "Terraform_Test_Internet_Gateway_01",
                    "qos_option_id": "e497bbc3-1127-4490-a51d-93582c40ab40",
                    "status": "ACTIVE",
                    "tenant_id": "01234567890123456789abcdefabcdef"
                }
            ]
        }
`
