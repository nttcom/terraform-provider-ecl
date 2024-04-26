package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1RouteImport(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesCreate)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowAfterCreate)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesDelete)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Route,
			},
			{
				ResourceName:      "ecl_mlb_route_v1.route",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
