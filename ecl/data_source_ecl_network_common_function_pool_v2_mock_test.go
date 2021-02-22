package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccNetworkV2CommonFunctionPoolsDataSource_name(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "common_function_pools", "/v2.0/common_function_pools", testMockNetworkV2CommonFunctionPoolsListNameQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testMockedAccPreCheckCommonFunctionPool(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testMockedAccNetworkV2CommonFunctionPoolDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_common_function_pool_v2.common_function_pool_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "description", "test Common Function Pool"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "id", "fc546cf7-1956-436b-a9b4-edc917e397cf"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "name", "F032000001492"),
				),
			},
		},
	})
}

func TestMockedAccNetworkV2CommonFunctionPoolsDataSource_description(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "common_function_pools", "/v2.0/common_function_pools", testMockNetworkV2CommonFunctionPoolsListDescriptionQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testMockedAccPreCheckCommonFunctionPool(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testMockedAccNetworkV2CommonFunctionPoolDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_common_function_pool_v2.common_function_pool_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "description", "test Common Function Pool"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "id", "fc546cf7-1956-436b-a9b4-edc917e397cf"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "name", "F032000001492"),
				),
			},
		},
	})
}

func TestMockedAccNetworkV2CommonFunctionPoolsDataSource_ID(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "common_function_pools", "/v2.0/common_function_pools", testMockNetworkV2CommonFunctionPoolsListIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testMockedAccPreCheckCommonFunctionPool(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testMockedAccNetworkV2CommonFunctionPoolDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_common_function_pool_v2.common_function_pool_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "description", "test Common Function Pool"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "id", "fc546cf7-1956-436b-a9b4-edc917e397cf"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "name", "F032000001492"),
				),
			},
		},
	})
}

var testMockNetworkV2CommonFunctionPoolsListNameQuery = fmt.Sprintf(`
request:
    method: GET
    query:
        name:
            - F032000001492
response:
    code: 200
    body: >
        {
            "common_function_pools": [
                {
                    "description": "test Common Function Pool",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492"
                }
            ]
        }
`)

var testMockNetworkV2CommonFunctionPoolsListDescriptionQuery = fmt.Sprintf(`
request:
    method: GET
    query:
        description:
            - test Common Function Pool
response:
    code: 200
    body: >
        {
            "common_function_pools": [
                {
                    "description": "test Common Function Pool",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492"
                }
            ]
        }
`)

var testMockNetworkV2CommonFunctionPoolsListIDQuery = fmt.Sprintf(`
request:
    method: GET
    query:
        id:
            - fc546cf7-1956-436b-a9b4-edc917e397cf
response:
    code: 200
    body: >
        {
            "common_function_pools": [
                {
                    "description": "test Common Function Pool",
                    "id": "fc546cf7-1956-436b-a9b4-edc917e397cf",
                    "name": "F032000001492"
                }
            ]
        }
`)

var testMockedAccNetworkV2CommonFunctionPoolDataSourceName = `
data "ecl_network_common_function_pool_v2" "common_function_pool_1" {
    name = "F032000001492"
}
`

var testMockedAccNetworkV2CommonFunctionPoolDataSourceDescription = `
data "ecl_network_common_function_pool_v2" "common_function_pool_1" {
    description = "test Common Function Pool"
}
`

var testMockedAccNetworkV2CommonFunctionPoolDataSourceID = `
data "ecl_network_common_function_pool_v2" "common_function_pool_1" {
    id = "fc546cf7-1956-436b-a9b4-edc917e397cf"
}
`
