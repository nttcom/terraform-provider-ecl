package ecl

import (
	"fmt"
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

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "internet_service", "/v2.0/internet_services", testMockNetworkV2InternetServiceListNameQuery)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways", testMockNetworkV2InternetGatewayPost)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetBasic)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingCreate)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetPendingDelete)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayGetDeleted)
	mc.Register(t, "internet_gateway", "/v2.0/internet_gateways/", testMockNetworkV2InternetGatewayDelete)

	mc.StartServer(t)

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
