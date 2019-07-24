package ecl

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	// "github.com/nttcom/eclcloud/ecl/network/v2/common_function_gateways"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/single_devices"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfCreation = "FGS_809F858574E94699952D0D7E7C58C81B"
const SoIDOfDeletion = "FGS_F2349100C7D24EF3ACD6B9A9F91FD220"

func TestAccSecurityV1NetworkBasedSingleDeviceBasic(t *testing.T) {
	var sd single_devices.SingleDevice

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV1NetworkBasedSingleDevicePost)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedSingleDeviceGetProcessingAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedSingleDeviceGetCompleteActiveAfterCreate)
	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV1NetworkBasedSingleDeviceDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedSingleDeviceProcessingAfterDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedSingleDeviceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedSingleDeviceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedSingleDeviceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedSingleDeviceExists(
						"ecl_security_network_based_single_device_v1.single_device_1", &sd),
				),
			},
		},
	})
}

func testAccCheckSecurityV1NetworkBasedSingleDeviceExists(n string, sd *single_devices.SingleDevice) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.securityOrderV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL security client: %s", err)
		}

		found, err := getSingleDeviceByHostName(client, rs.Primary.ID)
		// found, err := single_devices.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if strconv.Itoa(found.ID) != rs.Primary.ID {
			return fmt.Errorf("Security single device not found")
		}

		*sd = found

		return nil
	}
}

func testAccCheckSecurityV1NetworkBasedSingleDeviceDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_network_based_single_device_v1" {
			continue
		}
		// _, err := common_function_gateways.Get(client, rs.Primary.ID).Extract()

		_, err := getSingleDeviceByHostName(client, rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Common Function Gateway still exists")
		}

	}

	return nil
}

var testAccSecurityV1NetworkBasedSingleDeviceBasic = fmt.Sprintf(`
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

var testMockSecurityV1NetworkBasedSingleDevicePost = fmt.Sprintf(`
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
newStatus: Created
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
	- Created
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
    - Created
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
    - Deleted
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
