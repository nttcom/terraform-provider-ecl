package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1TargetGroupImport(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsCreate)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowAfterCreate)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsDelete)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1TargetGroup,
			},
			{
				ResourceName:      "ecl_mlb_target_group_v1.target_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
