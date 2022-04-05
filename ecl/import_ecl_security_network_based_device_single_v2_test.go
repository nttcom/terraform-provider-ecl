package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestAccSecurityV2NetworkBasedDeviceSingleImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "single_device", "/ecl-api/devices", testMockSecurityV2NetworkBasedDeviceSingleListDevicesAfterCreate)
	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV2NetworkBasedDeviceSingleListDeviceInterfaces)

	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV2NetworkBasedDeviceSingleListBeforeCreate)
	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV2NetworkBasedDeviceSingleCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV2NetworkBasedDeviceSingleGetProcessingAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV2NetworkBasedDeviceSingleGetCompleteActiveAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV2NetworkBasedDeviceSingleListAfterCreate)

	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV2NetworkBasedDeviceSingleDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV2NetworkBasedDeviceSingleProcessingAfterDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV2NetworkBasedDeviceSingleGetDeleteComplete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV2NetworkBasedDeviceSingleListAfterDelete)

	mc.StartServer(t)

	resourceName := "ecl_security_network_based_device_single_v2.device_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV2NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV2NetworkBasedDeviceSingleBasic,
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
