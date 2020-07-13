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
			{
				Config: testAccNetworkV2FICGatewayDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "description", OS_FIC_GW_DESCRIPTION),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "fic_service_id", OS_FIC_SERVICE_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "fic_gateway_id", OS_FIC_GW_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "name", OS_FIC_GW_NAME),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "qos_option_id", OS_FIC_GW_QOS_OPTION_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "status", OS_FIC_GW_STATUS),
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
			{
				Config: testAccNetworkV2FICGatewayDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceFICServiceID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceID2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceQoSOptionID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
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
        - %[1]q
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": %[1]q,
                    "fic_service_id": %[2]q,
                    "id": %[3]q,
                    "name": %[4]q,
                    "qos_option_id": %[5]q,
                    "status": %[6]q,
                    "tenant_id": %[7]q
                }
            ]
        }
`,
	OS_FIC_GW_DESCRIPTION, OS_FIC_SERVICE_ID, OS_FIC_GW_ID, OS_FIC_GW_NAME, OS_FIC_GW_QOS_OPTION_ID, OS_FIC_GW_STATUS, OS_TENANT_ID)

var testMockNetworkV2FICGatewayListFICServiceIDQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      fic_service_id:
        - %[2]q
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": %[1]q,
                    "fic_service_id": %[2]q,
                    "id": %[3]q,
                    "name": %[4]q,
                    "qos_option_id": %[5]q,
                    "status": %[6]q,
                    "tenant_id": %[7]q
                }
            ]
        }
`,
	OS_FIC_GW_DESCRIPTION, OS_FIC_SERVICE_ID, OS_FIC_GW_ID, OS_FIC_GW_NAME, OS_FIC_GW_QOS_OPTION_ID, OS_FIC_GW_STATUS, OS_TENANT_ID)

var testMockNetworkV2FICGatewayListIDQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      id:
        - %[3]q
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": %[1]q,
                    "fic_service_id": %[2]q,
                    "id": %[3]q,
                    "name": %[4]q,
                    "qos_option_id": %[5]q,
                    "status": %[6]q,
                    "tenant_id": %[7]q
                }
            ]
        }
`,
	OS_FIC_GW_DESCRIPTION, OS_FIC_SERVICE_ID, OS_FIC_GW_ID, OS_FIC_GW_NAME, OS_FIC_GW_QOS_OPTION_ID, OS_FIC_GW_STATUS, OS_TENANT_ID)

var testMockNetworkV2FICGatewayListNameQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      name:
        - %[4]q
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": %[1]q,
                    "fic_service_id": %[2]q,
                    "id": %[3]q,
                    "name": %[4]q,
                    "qos_option_id": %[5]q,
                    "status": %[6]q,
                    "tenant_id": %[7]q
                }
            ]
        }
`,
	OS_FIC_GW_DESCRIPTION, OS_FIC_SERVICE_ID, OS_FIC_GW_ID, OS_FIC_GW_NAME, OS_FIC_GW_QOS_OPTION_ID, OS_FIC_GW_STATUS, OS_TENANT_ID)

var testMockNetworkV2FICGatewayListQoSOptionQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      qos_option_id:
        - %[5]q
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": %[1]q,
                    "fic_service_id": %[2]q,
                    "id": %[3]q,
                    "name": %[4]q,
                    "qos_option_id": %[5]q,
                    "status": %[6]q,
                    "tenant_id": %[7]q
                }
            ]
        }
`,
	OS_FIC_GW_DESCRIPTION, OS_FIC_SERVICE_ID, OS_FIC_GW_ID, OS_FIC_GW_NAME, OS_FIC_GW_QOS_OPTION_ID, OS_FIC_GW_STATUS, OS_TENANT_ID)

var testMockNetworkV2FICGatewayListStatusQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      status:
        - %[6]q
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": %[1]q,
                    "fic_service_id": %[2]q,
                    "id": %[3]q,
                    "name": %[4]q,
                    "qos_option_id": %[5]q,
                    "status": %[6]q,
                    "tenant_id": %[7]q
                }
            ]
        }
`,
	OS_FIC_GW_DESCRIPTION, OS_FIC_SERVICE_ID, OS_FIC_GW_ID, OS_FIC_GW_NAME, OS_FIC_GW_QOS_OPTION_ID, OS_FIC_GW_STATUS, OS_TENANT_ID)

var testMockNetworkV2FICGatewayListTenantIDQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      tenant_id:
        - %[7]q
response:
    code: 200
    body: >
        {
            "fic_gateways": [
                {
                    "description": %[1]q,
                    "fic_service_id": %[2]q,
                    "id": %[3]q,
                    "name": %[4]q,
                    "qos_option_id": %[5]q,
                    "status": %[6]q,
                    "tenant_id": %[7]q
                }
            ]
        }
`,
	OS_FIC_GW_DESCRIPTION, OS_FIC_SERVICE_ID, OS_FIC_GW_ID, OS_FIC_GW_NAME, OS_FIC_GW_QOS_OPTION_ID, OS_FIC_GW_STATUS, OS_TENANT_ID)
