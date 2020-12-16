package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"

	"github.com/nttcom/eclcloud/ecl/network/v2/internet_gateways"
)

func TestMockedAccNetworkV2InternetGateway_basic(t *testing.T) {
	var internetGateway internet_gateways.InternetGateway

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListNameQuery)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayPost)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetBasic)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingCreate)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingUpdate1)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingUpdate2)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingDelete)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetUpdated1)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetUpdated2)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetDeleted)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayPut1)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayPut2)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckInternetGateway(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2InternetGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2InternetGatewayExists("ecl_network_internet_gateway_v2.internet_gateway_1", &internetGateway),
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

var testMockNetworkV2InternetGatewayPost = fmt.Sprintf(`
request:
    method: POST
response:
    code: 201
    body: >
        {
            "internet_gateway": {
                "description": "test_internet_gateway",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "Terraform_Test_Internet_Gateway_01",
                "qos_option_id": "%s",
                "status": "PENDING_CREATE",
                "tenant_id": "01234567890123456789abcdefabcdef"
            }
        }
newStatus: Created
`,
	OS_QOS_OPTION_ID_10M,
)

var testMockNetworkV2InternetGatewayGetBasic = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "test_internet_gateway",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "Terraform_Test_Internet_Gateway_01",
                "qos_option_id": "%s",
                "status": "ACTIVE",
                "tenant_id": "01234567890123456789abcdefabcdef"
            }
        }
expectedStatus:
    - Created
counter:
    min: 4
`,
	OS_QOS_OPTION_ID_10M,
)

var testMockNetworkV2InternetGatewayGetPendingCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "test_internet_gateway",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "Terraform_Test_Internet_Gateway_01",
                "qos_option_id": "%s",
                "status": "PENDING_CREATE",
                "tenant_id": "01234567890123456789abcdefabcdef"
            }
        }
expectedStatus:
    - Created
counter:
    max: 3
`,
	OS_QOS_OPTION_ID_10M,
)

var testMockNetworkV2InternetGatewayGetUpdated1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "qos_option_id": "%s",
                "status": "ACTIVE",
                "tenant_id": "01234567890123456789abcdefabcdef"
            }
        }
expectedStatus:
    - Updated1
counter:
    min: 4
`,
	OS_QOS_OPTION_ID_100M,
)

var testMockNetworkV2InternetGatewayGetUpdated2 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "",
                "qos_option_id": "%s",
                "status": "ACTIVE",
                "tenant_id": "01234567890123456789abcdefabcdef"
            }
        }
expectedStatus:
    - Updated2
counter:
    min: 4
`,
	OS_QOS_OPTION_ID_10M,
)

var testMockNetworkV2InternetGatewayGetPendingUpdate1 = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "qos_option_id": "4861fe30-d941-4199-8a20-eef1b2625a92",
                "status": "PENDING_UPDATE",
                "tenant_id": "01234567890123456789abcdefabcdef"
            }
        }
expectedStatus:
    - Updated1
counter:
    max: 3
`

var testMockNetworkV2InternetGatewayGetPendingUpdate2 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "",
                "qos_option_id": "%s",
                "status": "PENDING_UPDATE",
                "tenant_id": "01234567890123456789abcdefabcdef"
            }
        }
expectedStatus:
    - Updated2
counter:
    max: 3
`,
	OS_QOS_OPTION_ID_10M,
)

var testMockNetworkV2InternetGatewayGetDeleted = `
request:
    method: GET
response:
    code: 404
expectedStatus:
    - Deleted
counter:
    min: 4
`

var testMockNetworkV2InternetGatewayGetPendingDelete = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "test_internet_gateway2",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "Terraform_Test_Internet_Gateway_01",
                "qos_option_id": "%s",
                "status": "PENDING_DELETE",
                "tenant_id": "01234567890123456789abcdefabcdef"
            }
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`,
	OS_QOS_OPTION_ID_10M,
)

var testMockNetworkV2InternetGatewayPut1 = `
request:
    method: PUT
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "qos_option_id": "4861fe30-d941-4199-8a20-eef1b2625a92",
                "status": "PENDING_UPDATE",
                "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
            }
        }
expectedStatus:
    - Created
newStatus: Updated1
`

var testMockNetworkV2InternetGatewayPut2 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
            "internet_gateway": {
                "description": "",
                "id": "3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61",
                "internet_service_id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
                "name": "",
                "qos_option_id": "%s",
                "status": "PENDING_UPDATE",
                "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
            }
        }
expectedStatus:
    - Updated1
newStatus: Updated2
`,
	OS_QOS_OPTION_ID_10M,
)

var testMockNetworkV2InternetGatewayDelete = `
request:
    method: DELETE
response:
    code: 204
expectedStatus:
    - Created
    - Updated1
    - Updated2
newStatus: Deleted
`
