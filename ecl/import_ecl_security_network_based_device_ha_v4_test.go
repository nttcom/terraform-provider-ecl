package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestAccSecurityV4NetworkBasedDeviceHAImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	// TODO
	mc.Register(t, "ha_device", "/ecl-api/v4/devices", testMockSecurityV4NetworkBasedDeviceHAListDevicesAfterCreate)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/v4/devices/%s/interfaces", expectedNewHADeviceUUID1), testMockSecurityV4NetworkBasedDeviceHAListDeviceInterfaces)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/v4/devices/%s/interfaces", expectedNewHADeviceUUID2), testMockSecurityV4NetworkBasedDeviceHAListDeviceInterfaces)

	mc.Register(t, "ha_device", "/API/v4/ScreenEventFGHADeviceGet", testMockSecurityV4NetworkBasedDeviceHAListBeforeCreate)
	mc.Register(t, "ha_device", "/API/v4/SoEntryFGHA", testMockSecurityV4NetworkBasedDeviceHACreate)
	mc.Register(t, "ha_device", "/API/v4/ScreenEventFGSOrderProgressRate", testMockSecurityV4NetworkBasedDeviceHAGetProcessingAfterCreate)
	mc.Register(t, "ha_device", "/API/v4/ScreenEventFGSOrderProgressRate", testMockSecurityV4NetworkBasedDeviceHAGetCompleteActiveAfterCreate)
	mc.Register(t, "ha_device", "/API/v4/ScreenEventFGHADeviceGet", testMockSecurityV4NetworkBasedDeviceHAListAfterCreate)

	mc.Register(t, "ha_device", "/API/v4/SoEntryFGHA", testMockSecurityV4NetworkBasedDeviceHADelete)
	mc.Register(t, "ha_device", "/API/v4/ScreenEventFGSOrderProgressRate", testMockSecurityV4NetworkBasedDeviceHAProcessingAfterDelete)
	mc.Register(t, "ha_device", "/API/v4/ScreenEventFGSOrderProgressRate", testMockSecurityV4NetworkBasedDeviceHAGetDeleteComplete)
	mc.Register(t, "ha_device", "/API/v4/ScreenEventFGHADeviceGet", testMockSecurityV4NetworkBasedDeviceHAListAfterDelete)

	mc.StartServer(t)

	resourceName := "ecl_security_network_based_device_ha_v4.ha_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV4NetworkBasedDeviceHADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV4NetworkBasedDeviceHABasic,
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
