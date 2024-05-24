package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1LoadBalancerImport(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersCreate)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterCreate)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersDelete)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1LoadBalancer,
			},
			{
				ResourceName:      "ecl_mlb_load_balancer_v1.load_balancer",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

var testMockMLBV1LoadBalancersShowAfterDelete = fmt.Sprintf(`
request:
  method: GET
response:
  code: 404
expectedStatus:
  - Deleted
`)
