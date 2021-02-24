package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkV2CommonFunctionPoolDataSource_name(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommonFunctionPool(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2CommonFunctionPoolDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_common_function_pool_v2.common_function_pool_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "description", OS_COMMON_FUNCTION_POOL_DESCRIPTION),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "id", OS_COMMON_FUNCTION_POOL_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "name", OS_COMMON_FUNCTION_POOL_NAME),
				),
			},
		},
	})
}

func TestAccNetworkV2CommonFunctionPoolDataSource_ID(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommonFunctionPool(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2CommonFunctionPoolDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_common_function_pool_v2.common_function_pool_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "description", OS_COMMON_FUNCTION_POOL_DESCRIPTION),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "id", OS_COMMON_FUNCTION_POOL_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_network_common_function_pool_v2.common_function_pool_1", "name", OS_COMMON_FUNCTION_POOL_NAME),
				),
			},
		},
	})
}

var testAccNetworkV2CommonFunctionPoolDataSourceName = fmt.Sprintf(`
data "ecl_network_common_function_pool_v2" "common_function_pool_1" {
	name = %q
}
`, OS_COMMON_FUNCTION_POOL_NAME)

var testAccNetworkV2CommonFunctionPoolDataSourceID = fmt.Sprintf(`
data "ecl_network_common_function_pool_v2" "common_function_pool_1" {
	id = %q
}
`, OS_COMMON_FUNCTION_POOL_ID)
