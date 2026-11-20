package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v4/host_based"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfCreateHostBasedV4 = "FGHA_809F858574E94699952D0D7E7C58C81B"
const SoIDOfUpdateM1HostBasedV4 = "FGHA_809F858574E94699952D0D7E7C58C81C"
const SoIDOfUpdateM2HostBasedV4 = "FGHA_809F858574E94699952D0D7E7C58C81C"
const SoIDOfDeleteHostBasedV4 = "FGHA_F2349100C7D24EF3ACD6B9A9F91FD220"

func TestMockedAccSecurityV4HostBased_basic(t *testing.T) {
	var hs security.HostBasedSecurity

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "host_based", "/API/v4/SoEntryHBS", testMockSecurityV4HostBasedCreate)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderProgressRate", testMockSecurityV4HostBasedGetProcessingAfterCreate)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderProgressRate", testMockSecurityV4HostBasedGetCompleteActiveAfterCreate)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderInfoGet", testMockSecurityV4HostBasedGetAfterCreate)

	mc.Register(t, "host_based", "/API/v4/SoEntryHBS", testMockSecurityV4HostBasedUpdateM1)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderProgressRate", testMockSecurityV4HostBasedGetProcessingAfterUpdateM1)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderProgressRate", testMockSecurityV4HostBasedGetCompleteActiveAfterUpdateM1)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderInfoGet", testMockSecurityV4HostBasedGetAfterUpdatedM1)

	mc.Register(t, "host_based", "/API/v4/SoEntryHBS", testMockSecurityV4HostBasedUpdateM2)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderProgressRate", testMockSecurityV4HostBasedGetProcessingAfterUpdateM2)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderProgressRate", testMockSecurityV4HostBasedGetCompleteActiveAfterUpdateM2)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderInfoGet", testMockSecurityV4HostBasedGetAfterUpdatedM2)

	mc.Register(t, "host_based", "/API/v4/SoEntryHBS", testMockSecurityV4HostBasedDelete)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderProgressRate", testMockSecurityV4HostBasedProcessingAfterDelete)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderProgressRate", testMockSecurityV4HostBasedGetDeleteComplete)
	mc.Register(t, "host_based", "/API/v4/ScreenEventHBSOrderInfoGet", testMockSecurityV4HostBasedGetAfterDeleted)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV4HostBasedDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV4HostBasedBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV4HostBasedExists(
						"ecl_security_host_based_v4.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"service_order_service", "Managed Anti-Virus"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"max_agent_value", "1"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"mail_address", "hoge@example.com"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV4HostBasedUpdateM1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV4HostBasedExists(
						"ecl_security_host_based_v4.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"service_order_service", "Managed Virtual Patch"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"max_agent_value", "1"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"mail_address", "hoge@example.com"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV4HostBasedUpdateM2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV4HostBasedExists(
						"ecl_security_host_based_v4.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"service_order_service", "Managed Virtual Patch"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"max_agent_value", "2"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"mail_address", "hoge@example.com"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v4.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
		},
	})
}

var testMockedAccSecurityV4HostBasedBasic = fmt.Sprintf(`
resource "ecl_security_host_based_v4" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service = "Managed Anti-Virus"
	max_agent_value = 1
	mail_address = "hoge@example.com"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`, OS_TENANT_ID)

var testMockedAccSecurityV4HostBasedUpdateM1 = fmt.Sprintf(`
resource "ecl_security_host_based_v4" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service = "Managed Virtual Patch"
	max_agent_value = 1
	mail_address = "hoge@example.com"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`, OS_TENANT_ID)

var testMockedAccSecurityV4HostBasedUpdateM2 = fmt.Sprintf(`
resource "ecl_security_host_based_v4" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service ="Managed Virtual Patch"
	max_agent_value = 2
	mail_address = "hoge@example.com"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`, OS_TENANT_ID)

var testMockSecurityV4HostBasedCreate = fmt.Sprintf(`
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
    - ""
newStatus: Creating
`,
	SoIDOfCreateHostBasedV4,
)

var testMockSecurityV4HostBasedGetProcessingAfterCreate = `
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

var testMockSecurityV4HostBasedGetCompleteActiveAfterCreate = `
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

var testMockSecurityV4HostBasedGetAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "message": "正常終了",
            "code": "DEP-01",
            "region": "jp1",
            "tenant_name": "SOTestTenant",
            "tenant_description": "SOTest用テナント",
            "contract_id": "econ0000002279",
            "service_order_service": "Managed Anti-Virus",
            "max_agent_value": 1,
            "time_zone": "Asia/Tokyo",
            "mailaddress": "hoge@example.com",
            "dsm_lang": "ja",
            "tenant_flg": true
        }
expectedStatus:
    - Created
`)

var testMockSecurityV4HostBasedGetAfterUpdatedM1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "message": "正常終了",
            "code": "DEP-01",
            "region": "jp1",
            "tenant_name": "SOTestTenant",
            "tenant_description": "SOTest用テナント",
            "contract_id": "econ0000002279",
            "service_order_service": "Managed Virtual Patch",
            "max_agent_value": 1,
            "time_zone": "Asia/Tokyo",
            "mailaddress": "hoge@example.com",
            "dsm_lang": "ja",
            "tenant_flg": true
        }
expectedStatus:
    - UpdatedM1
`)

var testMockSecurityV4HostBasedGetAfterUpdatedM2 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "message": "正常終了",
            "code": "DEP-01",
            "region": "jp1",
            "tenant_name": "SOTestTenant",
            "tenant_description": "SOTest用テナント",
            "contract_id": "econ0000002279",
            "service_order_service": "Managed Virtual Patch",
            "max_agent_value": 2,
            "time_zone": "Asia/Tokyo",
            "mailaddress": "hoge@example.com",
            "dsm_lang": "ja",
            "tenant_flg": true
        }
expectedStatus:
    - UpdatedM2
`)

var testMockSecurityV4HostBasedGetAfterDeleted = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "message": "",
            "code": "DEP-01",
            "region": "",
            "tenant_name": "",
            "tenant_description": "",
            "contract_id": "",
            "service_order_service": "",
            "max_agent_value": "",
            "time_zone": "",
            "mailaddress": "",
            "dsm_lang": "",
            "tenant_flg": false
        }
expectedStatus:
    - Deleted
`)

var testMockSecurityV4HostBasedUpdateM1 = fmt.Sprintf(`
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
newStatus: UpdatingM1
`,
	SoIDOfUpdateM1HostBasedV4,
)

var testMockSecurityV4HostBasedGetProcessingAfterUpdateM1 = `
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
    - UpdatingM1
counter:
    max: 3
`

var testMockSecurityV4HostBasedGetCompleteActiveAfterUpdateM1 = `
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
    - UpdatingM1
newStatus: UpdatedM1
counter:
    min: 4
`

var testMockSecurityV4HostBasedUpdateM2 = fmt.Sprintf(`
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
    - UpdatedM1
newStatus: UpdatingM2
`,
	SoIDOfUpdateM2HostBasedV4,
)

var testMockSecurityV4HostBasedGetProcessingAfterUpdateM2 = `
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
    - UpdatingM2
counter:
    max: 3
`

var testMockSecurityV4HostBasedGetCompleteActiveAfterUpdateM2 = `
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
    - UpdatingM2
newStatus: UpdatedM2
counter:
    min: 4
`

var testMockSecurityV4HostBasedDelete = fmt.Sprintf(`
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
    - UpdatedM1
    - UpdatedM2
newStatus: Deleted
`,
	SoIDOfDeleteHostBasedV4,
)

var testMockSecurityV4HostBasedProcessingAfterDelete = `
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

var testMockSecurityV4HostBasedGetDeleteComplete = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-03",
            "message": "Order processing ends normally.",
            "progressRate": 70
        }
expectedStatus:
    - Deleted
counter:
    min: 4
`
