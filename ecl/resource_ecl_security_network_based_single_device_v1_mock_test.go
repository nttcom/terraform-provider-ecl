package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_firewall_utm_single"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfCreate = "FGS_809F858574E94699952D0D7E7C58C81B"
const SoIDOfUpdate = "FGS_809F858574E94699952D0D7E7C58C81C"
const SoIDOfDelete = "FGS_F2349100C7D24EF3ACD6B9A9F91FD220"

func TestMockedAccSecurityV1NetworkBasedFirewallUTMBasic(t *testing.T) {
	var sd security.SingleFirewallUTM

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV1NetworkBasedFirewallUTMListBeforeCreate)
	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV1NetworkBasedFirewallUTMCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedFirewallUTMGetProcessingAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedFirewallUTMGetCompleteActiveAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV1NetworkBasedFirewallUTMListAfterCreate)

	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV1NetworkBasedFirewallUTMUpdate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedFirewallUTMGetProcessingAfterUpdate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedFirewallUTMGetCompleteActiveAfterUpdate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV1NetworkBasedFirewallUTMListAfterUpdate)

	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV1NetworkBasedFirewallUTMDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedFirewallUTMProcessingAfterDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedFirewallUTMGetDeleteComplete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV1NetworkBasedFirewallUTMListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedFirewallUTMDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV1NetworkBasedFirewallUTMBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedFirewallUTMExists(
						"ecl_security_network_based_firewall_utm_single_v1.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"operating_mode", "FW"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"az_group", "zone1-groupb"),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV1NetworkBasedFirewallUTMUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedFirewallUTMExists(
						"ecl_security_network_based_firewall_utm_single_v1.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"operating_mode", "UTM"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"az_group", "zone1-groupb"),
				),
			},
		},
	})
}

var testMockedAccSecurityV1NetworkBasedFirewallUTMBasic = fmt.Sprintf(`
resource "ecl_security_network_based_firewall_utm_single_v1" "device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)

var testMockedAccSecurityV1NetworkBasedFirewallUTMUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_firewall_utm_single_v1" "device_1" {
	tenant_id = "%s"
	locale = "en"
	operating_mode = "UTM"
	license_kind = "08"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)

var testMockSecurityV1NetworkBasedFirewallUTMListBeforeCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-01",
            "message": "Successful completion",
            "records": 1,
            "rows": [{
            	"id": 1,
            	"cell": ["false", "1", "CES11810", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
            }]
        }
expectedStatus:
    - ""
newStatus: PreCreate
`

var testMockSecurityV1NetworkBasedFirewallUTMListAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-01",
            "message": "Successful completion",
            "records": 2,
            "rows": [
                {
                    "id": 1,
                    "cell": ["false", "1", "CES11810", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - Created
`

var testMockSecurityV1NetworkBasedFirewallUTMListAfterDelete = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-01",
            "message": "Successful completion",
            "records": 1,
            "rows": [{
            	"id": 1,
            	"cell": ["false", "1", "CES11810", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
            }]
        }
expectedStatus:
    - Deleted
`

var testMockSecurityV1NetworkBasedFirewallUTMCreate = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-02",
            "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
            "soId": "%s"
        }
expectedStatus:
    - PreCreate
newStatus: Creating
`,
	SoIDOfCreate,
)

var testMockSecurityV1NetworkBasedFirewallUTMGetProcessingAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-05",
            "message": "We accepted the order. Please wait",
            "progressRate": 45
        }
expectedStatus:
    - Creating
counter:
    max: 3
`

var testMockSecurityV1NetworkBasedFirewallUTMGetCompleteActiveAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-03",
            "message": "Order processing ends normally.",
            "progressRate": 100
        }
expectedStatus:
    - Creating
newStatus: Created
counter:
    min: 4
`

var testMockSecurityV1NetworkBasedFirewallUTMDelete = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-02",
            "message": "We accepted the order. You can check the status with ProgressRate.",
            "soId": "%s"
        }
expectedStatus:
    - Created
    - Updated
newStatus: Deleted
`,
	SoIDOfDelete,
)

var testMockSecurityV1NetworkBasedFirewallUTMProcessingAfterDelete = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-03",
            "message": "Order processing ends normally.",
            "progressRate": 55
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`

var testMockSecurityV1NetworkBasedFirewallUTMGetDeleteComplete = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-03",
            "message": "Order processing ends normally.",
            "progressRate": 100
        }
expectedStatus:
    - Deleted
counter:
    min: 4
`

var testMockSecurityV1NetworkBasedFirewallUTMUpdate = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-02",
            "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
            "soId": "%s"
        }
expectedStatus:
    - Created
newStatus: Updating
`,
	SoIDOfUpdate,
)

var testMockSecurityV1NetworkBasedFirewallUTMGetProcessingAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-05",
            "message": "We accepted the order. Please wait",
            "progressRate": 45
        }
expectedStatus:
    - Updating
counter:
    max: 3
`

var testMockSecurityV1NetworkBasedFirewallUTMGetCompleteActiveAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-03",
            "message": "Order processing ends normally.",
            "progressRate": 100
        }
expectedStatus:
    - Updating
newStatus: Updated
counter:
    min: 4
`

var testMockSecurityV1NetworkBasedFirewallUTMListAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-01",
            "message": "Successful completion",
            "records": 2,
            "rows": [
                {
                    "id": 1,
                    "cell": ["false", "1", "CES11810", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "UTM", "08", "standalone", "zone1-groupb", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - Updated
`
