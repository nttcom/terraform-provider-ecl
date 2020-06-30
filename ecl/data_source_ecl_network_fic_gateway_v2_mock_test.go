package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccNetworkV2FICGatewayDataSource_basic(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "fic_gateway", "/v2.0/fic_gateways", testMockNetworkV2FICGatewayListNameQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "description", ""),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "fic_service_id", "66a63898-32a5-4b9d-8925-f52be1d84764"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "id", "fc546cf7-1956-436b-a9b4-edc917e397cf"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "name", "F032000001492"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "qos_option_id", "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "status", "ACTIVE"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "tenant_id", OS_TENANT_ID),
				),
			},
		},
	})
}

func TestMockedAccNetworkV2FICGatewayDataSource_queries(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "fic_gateway", "/v2.0/fic_gateways", testMockNetworkV2FICGatewayListDescriptionQuery)
	mc.Register(t, "fic_gateway", "/v2.0/fic_gateways", testMockNetworkV2FICGatewayListFICServiceIDQuery)
	mc.Register(t, "fic_gateway", "/v2.0/fic_gateways", testMockNetworkV2FICGatewayListIDQuery)
	mc.Register(t, "fic_gateway", "/v2.0/fic_gateways", testMockNetworkV2FICGatewayListNameQuery)
	mc.Register(t, "fic_gateway", "/v2.0/fic_gateways", testMockNetworkV2FICGatewayListQoSOptionQuery)
	mc.Register(t, "fic_gateway", "/v2.0/fic_gateways", testMockNetworkV2FICGatewayListStatusQuery)
	mc.Register(t, "fic_gateway", "/v2.0/fic_gateways", testMockNetworkV2FICGatewayListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceFICServiceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceID2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceQoSOptionID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceStatus,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2FICGatewayDataSourceTenantID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2FICGatewayDataSourceID("data.ecl_network_fic_gateway_v2.fic_gateway_1"),
				),
			},
		},
	})
}

var testMockNetworkV2FICGatewayListDescriptionQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      description:
        - " "
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": "",
                    "fic_service_id": "66a63898-32a5-4b9d-8925-f52be1d84764",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492",
                    "qos_option_id": "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd",
                    "status": "ACTIVE",
                    "tenant_id": "%s"
                }
            ]
        }
`,
	OS_TENANT_ID)

var testMockNetworkV2FICGatewayListFICServiceIDQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      fic_service_id:
        - 66a63898-32a5-4b9d-8925-f52be1d84764
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": "",
                    "fic_service_id": "66a63898-32a5-4b9d-8925-f52be1d84764",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492",
                    "qos_option_id": "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd",
                    "status": "ACTIVE",
                    "tenant_id": "%s"
                }
            ]
        }
`,
	OS_TENANT_ID)

var testMockNetworkV2FICGatewayListIDQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      id:
        - fc546cf7-1956-436b-a9b4-edc917e397cf
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": "",
                    "fic_service_id": "66a63898-32a5-4b9d-8925-f52be1d84764",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492",
                    "qos_option_id": "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd",
                    "status": "ACTIVE",
                    "tenant_id": "%s"
                }
            ]
        }
`,
	OS_TENANT_ID)

var testMockNetworkV2FICGatewayListNameQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      name:
        - F032000001492
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": "",
                    "fic_service_id": "66a63898-32a5-4b9d-8925-f52be1d84764",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492",
                    "qos_option_id": "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd",
                    "status": "ACTIVE",
                    "tenant_id": "%s"
                }
            ]
        }
`,
	OS_TENANT_ID)

var testMockNetworkV2FICGatewayListQoSOptionQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      qos_option_id:
        - d384d7f5-22aa-46e5-8cf5-759e87c7b2fd
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": "",
                    "fic_service_id": "66a63898-32a5-4b9d-8925-f52be1d84764",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492",
                    "qos_option_id": "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd",
                    "status": "ACTIVE",
                    "tenant_id": "%s"
                }
            ]
        }
`,
	OS_TENANT_ID)

var testMockNetworkV2FICGatewayListStatusQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      status:
        - ACTIVE
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": "",
                    "fic_service_id": "66a63898-32a5-4b9d-8925-f52be1d84764",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492",
                    "qos_option_id": "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd",
                    "status": "ACTIVE",
                    "tenant_id": "%s"
                }
            ]
        }
`,
	OS_TENANT_ID)

var testMockNetworkV2FICGatewayListTenantIDQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      tenant_id:
        - %s
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": "",
                    "fic_service_id": "66a63898-32a5-4b9d-8925-f52be1d84764",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492",
                    "qos_option_id": "d384d7f5-22aa-46e5-8cf5-759e87c7b2fd",
                    "status": "ACTIVE",
                    "tenant_id": "%s"
                }
            ]
        }
`,
	OS_TENANT_ID, OS_TENANT_ID)
