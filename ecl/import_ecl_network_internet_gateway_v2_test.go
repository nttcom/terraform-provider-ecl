package ecl

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestAccNetworkV2InternetGatewayImportBasic(t *testing.T) {
	resourceName := "ecl_network_internet_gateway_v2.internet_gateway_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckInternetGateway(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2InternetGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestMockedAccNetworkV2InternetGatewayImportBasic(t *testing.T) {
	resourceName := "ecl_network_internet_gateway_v2.internet_gateway_1"

	testPrecheckMockEnv(t)

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, OS_REGION_NAME, mc.Endpoint())
	err := mc.Register("keystone", "/v3/auth/tokens", postKeystone)
	err = testSetupMockInternetGatewayBasic(mc)
	if err != nil {
		t.Errorf("Failed to setup mock: %s", err)
	}

	mc.StartServer()
	os.Setenv("OS_AUTH_URL", mc.Endpoint()+"v3/")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckInternetGateway(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2InternetGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2InternetGatewayBasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
