package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestAccSecurityV3NetworkBasedDeviceHAImport_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	// TODO
	mc.Register(t, "ha_device", "/ecl-api/devices", testMockSecurityV3NetworkBasedDeviceHAListDevicesAfterCreate)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID1), testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfaces)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID2), testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfaces)

	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListBeforeCreate)
	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV3NetworkBasedDeviceHACreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetProcessingAfterCreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetCompleteActiveAfterCreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListAfterCreate)

	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV3NetworkBasedDeviceHADelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAProcessingAfterDelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetDeleteComplete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListAfterDelete)

	mc.StartServer(t)

	resourceName := "ecl_security_network_based_device_ha_v3.ha_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedDeviceHADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceHABasic,
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
