package ecl

import (
	"fmt"
	"os"
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

	testPrecheckMockEnv(t)

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, OS_REGION_NAME, mc.Endpoint())
	err := mc.Register("keystone", "/v3/auth/tokens", postKeystone)
	err = testSetupMockInternetGatewayDatasourceBasic(mc)
	if err != nil {
		t.Errorf("Failed to setup mock: %s", err)
	}

	mc.StartServer()
	os.Setenv("OS_AUTH_URL", mc.Endpoint()+"v3/")

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

	testPrecheckMockEnv(t)

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, OS_REGION_NAME, mc.Endpoint())
	err := mc.Register("keystone", "/v3/auth/tokens", postKeystone)
	err = testSetupMockInternetGatewayDatasourceTestQueries(mc)
	if err != nil {
		t.Errorf("Failed to setup mock: %s", err)
	}

	mc.StartServer()
	os.Setenv("OS_AUTH_URL", mc.Endpoint()+"v3/")

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

func testSetupMockInternetGatewayDatasourceBasic(mc *mock.MockController) error {
	err := testSetupMockInternetGatewayBasic(mc)
	err = mc.Register("internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListNamequery)

	// latest error match
	if err != nil {
		return err
	}
	return nil
}

func testSetupMockInternetGatewayDatasourceTestQueries(mc *mock.MockController) error {
	err := testSetupMockInternetGatewayBasic(mc)
	err = mc.Register("internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListNamequery)
	err = mc.Register("internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListDescriptionquery)
	err = mc.Register("internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListIDquery)
	err = mc.Register("internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListInternetServiceIDquery)
	err = mc.Register("internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListQoSOptionIDquery)
	err = mc.Register("internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayListStatusquery)

	// latest error match
	if err != nil {
		return err
	}
	return nil
}

var testAccNetworkV2InternetGatewayDataSourceInternetGateway = fmt.Sprintf(`
data "ecl_network_internet_service_v2" "internet_service_1" {
    name = "Internet-Service-01"
}

resource "ecl_network_internet_gateway_v2" "internet_gateway_1" {
    name = "Terraform_Test_Internet_Gateway_01"
    description = "test_internet_gateway"
    internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
    qos_option_id = "%s"
}
`,
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
    internet_service_id = "${data.ecl_network_internet_service_v2.internet_service_1.id}"
}
`,
	testAccNetworkV2InternetGatewayDataSourceInternetGateway)

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

var testMockNetworkV2InternetGatewayListNamequery = `
request:
    method: GET
    query:
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

var testMockNetworkV2InternetGatewayListDescriptionquery = `
request:
    method: GET
    query:
      description:
        - test_internet_gateway
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

var testMockNetworkV2InternetGatewayListIDquery = `
request:
    method: GET
    query:
      id:
        - 3e71cf00-ddb5-4eb5-9ed0-ed4c481f6d61
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

var testMockNetworkV2InternetGatewayListInternetServiceIDquery = `
request:
    method: GET
    query:
      internet_service_id:
        - a7791c79-19b0-4eb6-9a8f-ea739b44e8d5
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

var testMockNetworkV2InternetGatewayListQoSOptionIDquery = `
request:
    method: GET
    query:
      qos_option_id:
        - e497bbc3-1127-4490-a51d-93582c40ab40
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

var testMockNetworkV2InternetGatewayListStatusquery = `
request:
    method: GET
    query:
      status:
        - ACTIVE
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
