package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1PolicyImport(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesCreate)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowAfterCreate)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesDelete)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Policy,
			},
			{
				ResourceName:      "ecl_mlb_policy_v1.policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
