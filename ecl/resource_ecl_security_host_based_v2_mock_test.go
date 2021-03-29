package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	security "github.com/nttcom/eclcloud/v2/ecl/security_order/v2/host_based"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfCreateHostBased = "FGHA_809F858574E94699952D0D7E7C58C81B"
const SoIDOfUpdateM1HostBased = "FGHA_809F858574E94699952D0D7E7C58C81C"
const SoIDOfUpdateM2HostBased = "FGHA_809F858574E94699952D0D7E7C58C81C"
const SoIDOfDeleteHostBased = "FGHA_F2349100C7D24EF3ACD6B9A9F91FD220"

func TestMockedAccSecurityV2HostBased_basic(t *testing.T) {
	var hs security.HostBasedSecurity

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "host_based", "/API/SoEntryHBS", testMockSecurityV2HostBasedCreate)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderProgressRate", testMockSecurityV2HostBasedGetProcessingAfterCreate)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderProgressRate", testMockSecurityV2HostBasedGetCompleteActiveAfterCreate)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderInfoGet", testMockSecurityV2HostBasedGetAfterCreate)

	mc.Register(t, "host_based", "/API/SoEntryHBS", testMockSecurityV2HostBasedUpdateM1)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderProgressRate", testMockSecurityV2HostBasedGetProcessingAfterUpdateM1)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderProgressRate", testMockSecurityV2HostBasedGetCompleteActiveAfterUpdateM1)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderInfoGet", testMockSecurityV2HostBasedGetAfterUpdatedM1)

	mc.Register(t, "host_based", "/API/SoEntryHBS", testMockSecurityV2HostBasedUpdateM2)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderProgressRate", testMockSecurityV2HostBasedGetProcessingAfterUpdateM2)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderProgressRate", testMockSecurityV2HostBasedGetCompleteActiveAfterUpdateM2)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderInfoGet", testMockSecurityV2HostBasedGetAfterUpdatedM2)

	mc.Register(t, "host_based", "/API/SoEntryHBS", testMockSecurityV2HostBasedDelete)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderProgressRate", testMockSecurityV2HostBasedProcessingAfterDelete)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderProgressRate", testMockSecurityV2HostBasedGetDeleteComplete)
	mc.Register(t, "host_based", "/API/ScreenEventHBSOrderInfoGet", testMockSecurityV2HostBasedGetAfterDeleted)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV2HostBasedDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV2HostBasedBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV2HostBasedExists(
						"ecl_security_host_based_v2.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"service_order_service", "Managed Anti-Virus"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"max_agent_value", "1"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"mail_address", "hoge@example.com"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV2HostBasedUpdateM1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV2HostBasedExists(
						"ecl_security_host_based_v2.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"service_order_service", "Managed Virtual Patch"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"max_agent_value", "1"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"mail_address", "hoge@example.com"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV2HostBasedUpdateM2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV2HostBasedExists(
						"ecl_security_host_based_v2.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"service_order_service", "Managed Virtual Patch"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"max_agent_value", "2"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"mail_address", "hoge@example.com"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v2.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
		},
	})
}

var testMockedAccSecurityV2HostBasedBasic = fmt.Sprintf(`
resource "ecl_security_host_based_v2" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service = "Managed Anti-Virus"
	max_agent_value = 1
	mail_address = "hoge@example.com"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`, OS_TENANT_ID)

var testMockedAccSecurityV2HostBasedUpdateM1 = fmt.Sprintf(`
resource "ecl_security_host_based_v2" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service = "Managed Virtual Patch"
	max_agent_value = 1
	mail_address = "hoge@example.com"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`, OS_TENANT_ID)

var testMockedAccSecurityV2HostBasedUpdateM2 = fmt.Sprintf(`
resource "ecl_security_host_based_v2" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service ="Managed Virtual Patch"
	max_agent_value = 2
	mail_address = "hoge@example.com"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`, OS_TENANT_ID)

var testMockSecurityV2HostBasedCreate = fmt.Sprintf(`
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
	SoIDOfCreateHostBased,
)

var testMockSecurityV2HostBasedGetProcessingAfterCreate = `
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

var testMockSecurityV2HostBasedGetCompleteActiveAfterCreate = `
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

var testMockSecurityV2HostBasedGetAfterCreate = fmt.Sprintf(`
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

var testMockSecurityV2HostBasedGetAfterUpdatedM1 = fmt.Sprintf(`
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

var testMockSecurityV2HostBasedGetAfterUpdatedM2 = fmt.Sprintf(`
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

var testMockSecurityV2HostBasedGetAfterDeleted = fmt.Sprintf(`
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

var testMockSecurityV2HostBasedUpdateM1 = fmt.Sprintf(`
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
	SoIDOfUpdateM1HostBased,
)

var testMockSecurityV2HostBasedGetProcessingAfterUpdateM1 = `
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

var testMockSecurityV2HostBasedGetCompleteActiveAfterUpdateM1 = `
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

var testMockSecurityV2HostBasedUpdateM2 = fmt.Sprintf(`
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
	SoIDOfUpdateM2HostBased,
)

var testMockSecurityV2HostBasedGetProcessingAfterUpdateM2 = `
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

var testMockSecurityV2HostBasedGetCompleteActiveAfterUpdateM2 = `
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

var testMockSecurityV2HostBasedDelete = fmt.Sprintf(`
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
	SoIDOfDeleteHostBased,
)

var testMockSecurityV2HostBasedProcessingAfterDelete = `
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

var testMockSecurityV2HostBasedGetDeleteComplete = `
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
