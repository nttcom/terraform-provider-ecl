package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1RuleImport(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesCreate)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowAfterCreate)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesDelete)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Rule,
			},
			{
				ResourceName:      "ecl_mlb_rule_v1.rule",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
