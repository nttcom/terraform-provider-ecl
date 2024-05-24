package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1HealthMonitorImport(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsCreate)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowAfterCreate)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsDelete)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1HealthMonitor,
			},
			{
				ResourceName:      "ecl_mlb_health_monitor_v1.health_monitor",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
