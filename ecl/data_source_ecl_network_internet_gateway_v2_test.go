package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestAccNetworkV2InternetGatewayDataSourceBasic(t *testing.T) {
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

func TestMockedAccNetworkV2InternetGatewayDataSourceBasic(t *testing.T) {

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
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

func TestAccNetworkV2InternetGatewayDataSourceTestQueries(t *testing.T) {
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

func TestMockedAccNetworkV2InternetGatewayDataSourceTestQueries(t *testing.T) {

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
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

func testAccCheckNetworkV2InternetGatewayDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find internet gateway data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Internet gateway data source ID not set")
		}

		return nil
	}
}

var testAccNetworkV2InternetGatewayDataSourceInternetGateway = fmt.Sprintf(`
resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	name = "Terraform_Test_Internet_Gateway_01"
	description = "test_internet_gateway"
	internet_service_id = "%s"
	qos_option_id = "%s"
}
`,
	OS_INTERNET_SERVICE_ID,
	OS_QOS_OPTION_ID_10M)

var testAccNetworkV2InternetGatewayDataSourceBasic = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	name = "${ecl_network_internet_gateway_v2.internet_gateway_1.name}"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceInternetServiceID = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	internet_service_id = "%s"
}
`,
	testAccNetworkV2InternetGatewayDataSourceInternetGateway,
	OS_INTERNET_SERVICE_ID)

var testAccNetworkV2InternetGatewayDataSourceQoSOptionID = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	qos_option_id = "%s"
}
`,
	testAccNetworkV2InternetGatewayDataSourceInternetGateway,
	OS_QOS_OPTION_ID_10M)

var testAccNetworkV2InternetGatewayDataSourceName = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
  name = "Terraform_Test_Internet_Gateway_01"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceDescription = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
	description = "test_internet_gateway"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceID = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
  internet_gateway_id = "${ecl_network_internet_gateway_v2.internet_gateway_1.id}"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

var testAccNetworkV2InternetGatewayDataSourceStatus = fmt.Sprintf(`
%s

data "ecl_network_internet_gateway_v2" "internet_gateway_1" {
  status = "ACTIVE"
}
`, testAccNetworkV2InternetGatewayDataSourceInternetGateway)

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
                    "internet_service_id": "5536154d-9a00-4b11-81fb-b185c9111d90",
                    "name": "Terraform_Test_Internet_Gateway_01",
                    "qos_option_id": "e497bbc3-1127-4490-a51d-93582c40ab40",
                    "status": "ACTIVE",
                    "tenant_id": "01234567890123456789abcdefabcdef"
                }
            ]
        }
`
