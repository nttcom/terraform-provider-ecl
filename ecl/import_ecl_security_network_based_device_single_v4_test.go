package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestAccSecurityV4NetworkBasedDeviceSingleImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "single_device", "/ecl-api/v4/devices", testMockSecurityV4NetworkBasedDeviceSingleListDevicesAfterCreate)
	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/v4/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV4NetworkBasedDeviceSingleListDeviceInterfaces)

	mc.Register(t, "single_device", "/API/v4/ScreenEventFGSDeviceGet", testMockSecurityV4NetworkBasedDeviceSingleListBeforeCreate)
	mc.Register(t, "single_device", "/API/v4/SoEntryFGS", testMockSecurityV4NetworkBasedDeviceSingleCreate)
	mc.Register(t, "single_device", "/API/v4/ScreenEventFGSOrderProgressRate", testMockSecurityV4NetworkBasedDeviceSingleGetProcessingAfterCreate)
	mc.Register(t, "single_device", "/API/v4/ScreenEventFGSOrderProgressRate", testMockSecurityV4NetworkBasedDeviceSingleGetCompleteActiveAfterCreate)
	mc.Register(t, "single_device", "/API/v4/ScreenEventFGSDeviceGet", testMockSecurityV4NetworkBasedDeviceSingleListAfterCreate)

	mc.Register(t, "single_device", "/API/v4/SoEntryFGS", testMockSecurityV4NetworkBasedDeviceSingleDelete)
	mc.Register(t, "single_device", "/API/v4/ScreenEventFGSOrderProgressRate", testMockSecurityV4NetworkBasedDeviceSingleProcessingAfterDelete)
	mc.Register(t, "single_device", "/API/v4/ScreenEventFGSOrderProgressRate", testMockSecurityV4NetworkBasedDeviceSingleGetDeleteComplete)
	mc.Register(t, "single_device", "/API/v4/ScreenEventFGSDeviceGet", testMockSecurityV4NetworkBasedDeviceSingleListAfterDelete)

	mc.StartServer(t)

	resourceName := "ecl_security_network_based_device_single_v4.device_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV4NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV4NetworkBasedDeviceSingleBasic,
			},

			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"locale", "tenant_id"},
			},
		},
	})
}
