package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	// "github.com/hashicorp/terraform/terraform"

	// "github.com/nttcom/eclcloud/ecl/network/v2/common_function_gateways"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/single_devices"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfCreation = "FGS_809F858574E94699952D0D7E7C58C81B"
const SoIDOfDeletion = "FGS_F2349100C7D24EF3ACD6B9A9F91FD220"

func TestMockedAccSecurityV1NetworkBasedSingleDeviceBasic(t *testing.T) {
	var sd single_devices.SingleDevice

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV1NetworkBasedSingleDeviceListBeforeCreate)
	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV1NetworkBasedSingleDeviceCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedSingleDeviceGetProcessingAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedSingleDeviceGetCompleteActiveAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV1NetworkBasedSingleDeviceListAfterCreate)
	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV1NetworkBasedSingleDeviceDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedSingleDeviceProcessingAfterDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedSingleDeviceGetDeleteComplete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV1NetworkBasedSingleDeviceListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedSingleDeviceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV1NetworkBasedSingleDeviceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedSingleDeviceExists(
						"ecl_security_network_based_single_device_v1.single_device_1", &sd),
				),
			},
		},
	})
}

var testMockedAccSecurityV1NetworkBasedSingleDeviceBasic = fmt.Sprintf(`
resource "ecl_security_network_based_single_device_v1" "single_device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)

var testMockSecurityV1NetworkBasedSingleDeviceListBeforeCreate = `
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

var testMockSecurityV1NetworkBasedSingleDeviceListAfterCreate = `
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

var testMockSecurityV1NetworkBasedSingleDeviceListAfterDelete = `
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

var testMockSecurityV1NetworkBasedSingleDeviceCreate = fmt.Sprintf(`
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
	SoIDOfCreation,
)

var testMockSecurityV1NetworkBasedSingleDeviceGetProcessingAfterCreate = `
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

var testMockSecurityV1NetworkBasedSingleDeviceGetCompleteActiveAfterCreate = `
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

var testMockSecurityV1NetworkBasedSingleDeviceDelete = fmt.Sprintf(`
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
newStatus: Deleted
`,
	SoIDOfDeletion,
)

var testMockSecurityV1NetworkBasedSingleDeviceProcessingAfterDelete = `
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

var testMockSecurityV1NetworkBasedSingleDeviceGetDeleteComplete = `
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
