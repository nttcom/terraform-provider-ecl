package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestAccSecurityV3NetworkBasedDeviceSingleImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "single_device", "/ecl-api/devices", testMockSecurityV3NetworkBasedDeviceSingleListDevicesAfterCreate)
	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV3NetworkBasedDeviceSingleListDeviceInterfaces)

	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListBeforeCreate)
	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV3NetworkBasedDeviceSingleCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetProcessingAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetCompleteActiveAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListAfterCreate)

	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV3NetworkBasedDeviceSingleDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleProcessingAfterDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetDeleteComplete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListAfterDelete)

	mc.StartServer(t)

	resourceName := "ecl_security_network_based_device_single_v3.device_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceSingleBasic,
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
