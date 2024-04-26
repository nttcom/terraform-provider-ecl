package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1ListenerImport(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersCreate)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowAfterCreate)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersDelete)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Listener,
			},
			{
				ResourceName:      "ecl_mlb_listener_v1.listener",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
