package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestAccSecurityV2NetworkBasedDeviceHAImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	// TODO
	mc.Register(t, "ha_device", "/ecl-api/devices", testMockSecurityV2NetworkBasedDeviceHAListDevicesAfterCreate)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID1), testMockSecurityV2NetworkBasedDeviceHAListDeviceInterfaces)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID2), testMockSecurityV2NetworkBasedDeviceHAListDeviceInterfaces)

	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV2NetworkBasedDeviceHAListBeforeCreate)
	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV2NetworkBasedDeviceHACreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV2NetworkBasedDeviceHAGetProcessingAfterCreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV2NetworkBasedDeviceHAGetCompleteActiveAfterCreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV2NetworkBasedDeviceHAListAfterCreate)

	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV2NetworkBasedDeviceHADelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV2NetworkBasedDeviceHAProcessingAfterDelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV2NetworkBasedDeviceHAGetDeleteComplete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV2NetworkBasedDeviceHAListAfterDelete)

	mc.StartServer(t)

	resourceName := "ecl_security_network_based_device_ha_v2.ha_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV2NetworkBasedDeviceHADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV2NetworkBasedDeviceHABasic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ha_link_1.#", "ha_link_1.0.host_1_ip_address",
					"ha_link_1.0.host_2_ip_address", "ha_link_1.0.network_id", "ha_link_1.0.subnet_id",
					"ha_link_2.#", "ha_link_2.0.host_1_ip_address", "ha_link_2.0.host_2_ip_address",
					"ha_link_2.0.network_id", "ha_link_2.0.subnet_id", "locale", "tenant_id"},
			},
		},
	})
}
