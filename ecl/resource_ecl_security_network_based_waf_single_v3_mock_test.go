package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v3/network_based_device_single"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfWAFUpdate = "FGWAF_809F858574E94699952D0D7E7C58C81C"
const SoIDOfWAFDelete = "FGWAF_F2349100C7D24EF3ACD6B9A9F91FD220"

func TestMockedAccSecurityV3NetworkBasedWAFSingle_basic(t *testing.T) {
	var sd security.SingleDevice

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "waf", "/ecl-api/devices", testMockSecurityV3NetworkBasedWAFSingleListDevicesAfterCreate)
	mc.Register(t, "waf", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV3NetworkBasedWAFSingleListDevicesAfterCreate)

	mc.Register(t, "waf", "/API/ScreenEventFGWAFDeviceGet", testMockSecurityV3NetworkBasedWAFSingleListBeforeCreate)
	mc.Register(t, "waf", "/API/SoEntryFGWAF", testMockSecurityV3NetworkBasedWAFSingleCreate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleGetProcessingAfterCreate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleGetCompleteActiveAfterCreate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFDeviceGet", testMockSecurityV3NetworkBasedWAFSingleListAfterCreate)

	mc.Register(t, "waf", "/API/SoEntryFGWAF", testMockSecurityV3NetworkBasedWAFSingleUpdate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleGetProcessingAfterUpdate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleGetCompleteActiveAfterUpdate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFDeviceGet", testMockSecurityV3NetworkBasedWAFSingleListAfterUpdate)

	mc.Register(t, "waf", "/API/SoEntryFGWAF", testMockSecurityV3NetworkBasedWAFSingleDelete)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleProcessingAfterDelete)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleGetDeleteComplete)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFDeviceGet", testMockSecurityV3NetworkBasedWAFSingleListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedWAFSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedWAFSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v3.waf_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "operating_mode", "WAF"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "az_group", OS_DEFAULT_ZONE),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedWAFSingleUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v3.waf_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "operating_mode", "WAF"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "az_group", OS_DEFAULT_ZONE),
				),
			},
		},
	})
}

func TestMockedAccSecurityV3NetworkBasedWAFSingleUpdateInterface(t *testing.T) {
	var sd security.SingleDevice

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "waf", "/API/ScreenEventFGWAFDeviceGet", testMockSecurityV3NetworkBasedWAFSingleListBeforeCreate)
	mc.Register(t, "waf", "/API/SoEntryFGWAF", testMockSecurityV3NetworkBasedWAFSingleCreate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleGetProcessingAfterCreate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleGetCompleteActiveAfterCreate)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFDeviceGet", testMockSecurityV3NetworkBasedWAFSingleListAfterCreate)
	mc.Register(t, "waf", "/ecl-api/devices", testMockSecurityV3NetworkBasedWAFSingleListDevicesAfterCreate)
	mc.Register(t, "waf", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV3NetworkBasedWAFSingleListDeviceInterfacesAfterCreate)

	mc.Register(t, "waf", fmt.Sprintf("/ecl-api/ports/waf/%s", expectedNewSingleDeviceHostName), testMockSecurityV3NetworkBasedWAFSingleUpdateInterface)
	mc.Register(t, "waf", fmt.Sprintf("/ecl-api/process/%d/status", ProcessIDOfUpdateInterface), testMockSecurityV3NetworkBasedWAFSingleGetProcessingAfterUpdateInterface)
	mc.Register(t, "waf", fmt.Sprintf("/ecl-api/process/%d/status", ProcessIDOfUpdateInterface), testMockSecurityV3NetworkBasedWAFSingleGetCompleteActiveAfterUpdateInterface)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFDeviceGet", testMockSecurityV3NetworkBasedWAFSingleListAfterInterfaceUpdate)
	mc.Register(t, "waf", "/ecl-api/devices", testMockSecurityV3NetworkBasedWAFSingleListDevicesAfterUpdate)
	mc.Register(t, "waf", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV3NetworkBasedWAFSingleListDeviceInterfacesAfterUpdate)

	mc.Register(t, "waf", "/API/SoEntryFGWAF", testMockSecurityV3NetworkBasedWAFSingleDelete)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleProcessingAfterDelete)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFOrderProgressRate", testMockSecurityV3NetworkBasedWAFSingleGetDeleteComplete)
	mc.Register(t, "waf", "/API/ScreenEventFGWAFDeviceGet", testMockSecurityV3NetworkBasedWAFSingleListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedWAFSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedWAFSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v3.waf_1", &sd),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedWAFSingleUpdateInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedWAFSingleExists(
						"ecl_security_network_based_waf_single_v3.waf_1", &sd),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "port.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "port.0.network_id", "dummyNetwork1"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "port.0.subnet_id", "dummySubnet1"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "port.0.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_waf_single_v3.waf_1", "port.0.comment", "port 0 comment"),

					// resource.TestCheckResourceAttr(
					// 	"ecl_security_network_based_waf_single_v3.waf_1", "port.3.ip_address", "192.168.2.50"),
					// resource.TestCheckResourceAttr(
					// 	"ecl_security_network_based_waf_single_v3.waf_1", "port.3.network_id", "dummyNetwork2"),
					// resource.TestCheckResourceAttr(
					// 	"ecl_security_network_based_waf_single_v3.waf_1", "port.3.subnet_id", "dummySubnet2"),
					// resource.TestCheckResourceAttr(
					// 	"ecl_security_network_based_waf_single_v3.waf_1", "port.3.mtu", "1500"),
					// resource.TestCheckResourceAttr(
					// 	"ecl_security_network_based_waf_single_v3.waf_1", "port.3.comment", "port 3 comment"),
				),
			},
		},
	})
}

var testMockedAccSecurityV3NetworkBasedWAFSingleBasic = fmt.Sprintf(`
resource "ecl_security_network_based_waf_single_v3" "waf_1" {
	tenant_id = "%s"
	locale = "ja"
	license_kind = "02"
	az_group = "%s"
}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testMockedAccSecurityV3NetworkBasedWAFSingleUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_waf_single_v3" "waf_1" {
	tenant_id = "%s"
	locale = "en"
	license_kind = "08"
	az_group = "%s"
}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testMockedAccSecurityV3NetworkBasedWAFSingleUpdateInterface = fmt.Sprintf(`

resource "ecl_security_network_based_waf_single_v3" "waf_1" {
	tenant_id = "%s"
	locale = "ja"
	license_kind = "02"
	az_group = "%s"

  port {
      enable = "true"
      ip_address = "192.168.1.50"
      ip_address_prefix = 24
      network_id = "dummyNetwork1"
      subnet_id = "dummySubnet1"
      mtu = "1500"
      comment = "port 0 comment"
  }
}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedWAFSingleListBeforeCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-01",
            "message": "Successful completion",
            "records": 1,
            "rows": [{
            	"id": 1,
            	"cell": ["false", "1", "CES11810", "WAF", "02", "%s", "jp4_zone1"]
            }]
        }
expectedStatus:
    - ""
newStatus: PreCreate
`,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedWAFSingleListAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-01",
            "message": "Successful completion",
            "records": 2,
            "rows": [
                {
                    "id": 1,
                    "cell": ["false", "1", "CES11810", "WAF", "02", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "WAF", "02", "%s", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - Created
    - Updating
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedWAFSingleListDevicesAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "devices": [
            {
              "msa_device_id": "CES11810",
              "os_server_id": "392a90bf-2c1b-45fd-8221-096894fff39d",
              "os_server_name": "WAF-CES11878",
              "os_availability_zone": "%s",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "WAF",
              "os_server_status": "ACTIVE"
            },
            {
              "msa_device_id": "CES11811",
              "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278e",
              "os_server_name": "WAF-CES11816",
              "os_availability_zone": "%s",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "WAF",
              "os_server_status": "ACTIVE"
            }
          ]
        }
expectedStatus:
    - Updated
    - Created
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedWAFSingleListDevicesAfterUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "devices": [
            {
              "msa_device_id": "CES11810",
              "os_server_id": "392a90bf-2c1b-45fd-8221-096894fff39d",
              "os_server_name": "WAF-CES11878",
              "os_availability_zone": "%s",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "WAF",
              "os_server_status": "ACTIVE"
            },
            {
              "msa_device_id": "CES11811",
              "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278e",
              "os_server_name": "WAF-CES11816",
              "os_availability_zone": "%s",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "WAF",
              "os_server_status": "ACTIVE"
            }
          ]
        }
expectedStatus:
    - InterfaceIsUpdated
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedWAFSingleListDeviceInterfacesAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "device_interfaces": []
        }
expectedStatus:
    - Created
    - Updating
    - Updated
`

var testMockSecurityV3NetworkBasedWAFSingleListDeviceInterfacesAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "device_interfaces": [
            {
              "os_ip_address": "192.168.1.50",
              "msa_port_id": "port2",
              "os_port_name": "port2-CES11892",
              "os_port_id": "82ebe045-9c9a-4088-8b33-cb0d590079aa",
              "os_network_id": "dummyNetwork1",
              "os_port_status": "ACTIVE",
              "os_mac_address": "fa:16:3e:05:ff:66",
              "os_subnet_id": "dummySubnet1"
            }
          ]
        }
expectedStatus:
    - InterfaceIsUpdated
`

var testMockSecurityV3NetworkBasedWAFSingleListAfterDelete = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-01",
            "message": "Successful completion",
            "records": 1,
            "rows": [{
            	"id": 1,
            	"cell": ["false", "1", "CES11810", "WAF", "02", "%s", "jp4_zone1"]
            }]
        }
expectedStatus:
    - Deleted
`,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedWAFSingleCreate = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-02",
            "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
            "soId": "%s"
        }
expectedStatus:
    - PreCreate
newStatus: Creating
`,
	SoIDOfCreate,
)

var testMockSecurityV3NetworkBasedWAFSingleGetProcessingAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-05",
            "message": "We accepted the order. Please wait",
            "progressRate": 45
        }
expectedStatus:
    - Creating
counter:
    max: 3
`

var testMockSecurityV3NetworkBasedWAFSingleGetCompleteActiveAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-03",
            "message": "Order processing ends normally.",
            "progressRate": 100
        }
expectedStatus:
    - Creating
newStatus: Created
counter:
    min: 4
`

var testMockSecurityV3NetworkBasedWAFSingleDelete = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-02",
            "message": "We accepted the order. You can check the status with ProgressRate.",
            "soId": "%s"
        }
expectedStatus:
    - Created
    - Updated
    - InterfaceIsUpdated
newStatus: Deleted
`,
	SoIDOfWAFDelete,
)

var testMockSecurityV3NetworkBasedWAFSingleProcessingAfterDelete = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-03",
            "message": "Order processing ends normally.",
            "progressRate": 55
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`

var testMockSecurityV3NetworkBasedWAFSingleGetDeleteComplete = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-03",
            "message": "Order processing ends normally.",
            "progressRate": 100
        }
expectedStatus:
    - Deleted
counter:
    min: 4
`

var testMockSecurityV3NetworkBasedWAFSingleUpdate = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-02",
            "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
            "soId": "%s"
        }
expectedStatus:
    - Created
newStatus: Updating
`,
	SoIDOfWAFUpdate,
)

var testMockSecurityV3NetworkBasedWAFSingleGetProcessingAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-05",
            "message": "We accepted the order. Please wait",
            "progressRate": 45
        }
expectedStatus:
    - Updating
counter:
    max: 3
`

var testMockSecurityV3NetworkBasedWAFSingleGetCompleteActiveAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-03",
            "message": "Order processing ends normally.",
            "progressRate": 100
        }
expectedStatus:
    - Updating
newStatus: Updated
counter:
    min: 4
`

var testMockSecurityV3NetworkBasedWAFSingleListAfterUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-01",
            "message": "Successful completion",
            "records": 2,
            "rows": [
                {
                    "id": 1,
                    "cell": ["false", "1", "CES11810", "WAF", "02", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "WAF", "08", "%s", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - Updated
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedWAFSingleListAfterInterfaceUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-01",
            "message": "Successful completion",
            "records": 2,
            "rows": [
                {
                    "id": 1,
                    "cell": ["false", "1", "CES11810", "WAF", "02", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "WAF", "02", "%s", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedWAFSingleUpdateInterface = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
            "message": "The process launch request has been accepted",
            "processId": %d
        }
expectedStatus:
    - Created
newStatus: InterfaceIsUpdating
`,
	ProcessIDOfUpdateInterface,
)

var testMockSecurityV3NetworkBasedWAFSingleGetProcessingAfterUpdateInterface = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "processInstance": {
            "processId": {
              "id": 85385,
              "lastExecNumber": 1,
              "name": "ntt/FortiVA_Port_Management/Process_Manage_UTM_Interfaces/Process_Manage_UTM_Interfaces",
              "submissionType": "RUN"
            },
            "serviceId": {
              "id": 19382,
              "name": "FortiVA_Port_Management",
              "serviceReference": "PORT_MNGT_CES11892",
              "state": null
            },
            "status": {
              "comment": "Ping Monitoring started for the device 11892.",
              "duration": 0,
              "endingDate": "2019-07-26 04:34:56.0",
              "execNumber": 1,
              "processInstanceId": 85385,
              "processName": "ntt/FortiVA_Port_Management/Process_Manage_UTM_Interfaces/Process_Manage_UTM_Interfaces",
              "startingDate": "2019-07-26 04:24:45.0",
              "status": "RUNNING",
              "taskStatusList": [
                {
                  "comment": "IP Address inputs verified successfully.",
                  "endingDate": "2019-07-26 04:24:48.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:24:45.0",
                  "status": "ENDED",
                  "taskId": 1,
                  "taskName": "Verify IP Address, MTU Inputs"
                },
                {
                  "comment": "Ping Monitoring stopped for the device 11892.",
                  "endingDate": "2019-07-26 04:26:49.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:24:48.0",
                  "status": "ENDED",
                  "taskId": 2,
                  "taskName": "Stop Ping Monitoring"
                },
                {
                  "comment": "Openstack Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3 stopped successfully.\nServer Status : SHUTOFF\nTask State : -\nPower State : Shutdown\n",
                  "endingDate": "2019-07-26 04:27:03.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:26:49.0",
                  "status": "ENDED",
                  "taskId": 3,
                  "taskName": "Stop the UTM"
                },
                {
                  "comment": "IP Address 100.76.96.230 is now unreachable from MSA.\nPING Status : Destination Host Unreachable\n",
                  "endingDate": "2019-07-26 04:27:13.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:27:03.0",
                  "status": "ENDED",
                  "taskId": 4,
                  "taskName": "Wait for UTM Ping unreachability from MSA"
                },
                {
                  "comment": "Ports deleted successfully.",
                  "endingDate": "2019-07-26 04:28:29.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:27:13.0",
                  "status": "ENDED",
                  "taskId": 5,
                  "taskName": "Delete Ports"
                },
                {
                  "comment": "Ports created successfully.\nPort Id : 34c7389d-1428-4f98-a37c-9c2e32aab255\nPort Id : 3d09053b-fad8-45c4-bf71-501c0fc2b58a\nPort Id : 0262d90c-6056-4308-8b76-8e851f0132f5\nPort Id : 5fcabdf2-8a20-4337-bd10-02f5c5000ca1\nPort Id : 53211b09-f82b-40d5-bf5b-7289a298cbdf\nPort Id : 9ce2d3b7-7ae0-400d-8e41-16dc9b94f95e\nPort Id : a36493fe-43d2-4dc1-a39e-c96898e9c0be\n",
                  "endingDate": "2019-07-26 04:29:50.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:28:29.0",
                  "status": "ENDED",
                  "taskId": 6,
                  "taskName": "Create Ports"
                },
                {
                  "comment": "Ports attached successfully to the Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3.",
                  "endingDate": "2019-07-26 04:31:33.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:29:50.0",
                  "status": "ENDED",
                  "taskId": 7,
                  "taskName": "Attach Ports"
                },
                {
                  "comment": "Openstack Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3 started successfully.\nServer Status : ACTIVE\nTask State : -\nPower State : Running\n",
                  "endingDate": "2019-07-26 04:31:47.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:31:33.0",
                  "status": "ENDED",
                  "taskId": 8,
                  "taskName": "Start the UTM"
                },
                {
                  "comment": "IP Address 100.76.96.230 is now reachable from MSA.\nPING Status : OK\n",
                  "endingDate": "2019-07-26 04:32:30.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:31:47.0",
                  "status": "ENDED",
                  "taskId": 9,
                  "taskName": "Wait for UTM Ping reachability from MSA"
                },
                {
                  "comment": "OK LICENSE IS VALID",
                  "endingDate": "2019-07-26 04:32:56.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:32:30.0",
                  "status": "ENDED",
                  "taskId": 10,
                  "taskName": "Verify License Validity"
                },
                {
                  "comment": "Ports updated successfully on Fortigate Device 11892.\n",
                  "endingDate": "2019-07-26 04:33:17.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:32:56.0",
                  "status": "ENDED",
                  "taskId": 11,
                  "taskName": "Update UTM"
                },
                {
                  "comment": "Device 11892 Backup completed successfully.\nBackup Status : ENDED\nBackup Message : BACKUP  processed\n\nBackup Revision Id : 209408\n",
                  "endingDate": "2019-07-26 04:33:28.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:33:17.0",
                  "status": "ENDED",
                  "taskId": 12,
                  "taskName": "Device Backup"
                },
                {
                  "comment": "Ping Monitoring started for the device 11892.",
                  "endingDate": "2019-07-26 04:34:56.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:33:28.0",
                  "status": "ENDED",
                  "taskId": 13,
                  "taskName": "Start Ping Monitoring"
                }
              ]
            }
          }
        }
expectedStatus:
    - InterfaceIsUpdating
counter:
    max: 3
`

var testMockSecurityV3NetworkBasedWAFSingleGetCompleteActiveAfterUpdateInterface = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "processInstance": {
            "processId": {
              "id": 85385,
              "lastExecNumber": 1,
              "name": "ntt/FortiVA_Port_Management/Process_Manage_UTM_Interfaces/Process_Manage_UTM_Interfaces",
              "submissionType": "RUN"
            },
            "serviceId": {
              "id": 19382,
              "name": "FortiVA_Port_Management",
              "serviceReference": "PORT_MNGT_CES11892",
              "state": null
            },
            "status": {
              "comment": "Ping Monitoring started for the device 11892.",
              "duration": 0,
              "endingDate": "2019-07-26 04:34:56.0",
              "execNumber": 1,
              "processInstanceId": 85385,
              "processName": "ntt/FortiVA_Port_Management/Process_Manage_UTM_Interfaces/Process_Manage_UTM_Interfaces",
              "startingDate": "2019-07-26 04:24:45.0",
              "status": "ENDED",
              "taskStatusList": [
                {
                  "comment": "IP Address inputs verified successfully.",
                  "endingDate": "2019-07-26 04:24:48.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:24:45.0",
                  "status": "ENDED",
                  "taskId": 1,
                  "taskName": "Verify IP Address, MTU Inputs"
                },
                {
                  "comment": "Ping Monitoring stopped for the device 11892.",
                  "endingDate": "2019-07-26 04:26:49.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:24:48.0",
                  "status": "ENDED",
                  "taskId": 2,
                  "taskName": "Stop Ping Monitoring"
                },
                {
                  "comment": "Openstack Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3 stopped successfully.\nServer Status : SHUTOFF\nTask State : -\nPower State : Shutdown\n",
                  "endingDate": "2019-07-26 04:27:03.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:26:49.0",
                  "status": "ENDED",
                  "taskId": 3,
                  "taskName": "Stop the UTM"
                },
                {
                  "comment": "IP Address 100.76.96.230 is now unreachable from MSA.\nPING Status : Destination Host Unreachable\n",
                  "endingDate": "2019-07-26 04:27:13.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:27:03.0",
                  "status": "ENDED",
                  "taskId": 4,
                  "taskName": "Wait for UTM Ping unreachability from MSA"
                },
                {
                  "comment": "Ports deleted successfully.",
                  "endingDate": "2019-07-26 04:28:29.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:27:13.0",
                  "status": "ENDED",
                  "taskId": 5,
                  "taskName": "Delete Ports"
                },
                {
                  "comment": "Ports created successfully.\nPort Id : 34c7389d-1428-4f98-a37c-9c2e32aab255\nPort Id : 3d09053b-fad8-45c4-bf71-501c0fc2b58a\nPort Id : 0262d90c-6056-4308-8b76-8e851f0132f5\nPort Id : 5fcabdf2-8a20-4337-bd10-02f5c5000ca1\nPort Id : 53211b09-f82b-40d5-bf5b-7289a298cbdf\nPort Id : 9ce2d3b7-7ae0-400d-8e41-16dc9b94f95e\nPort Id : a36493fe-43d2-4dc1-a39e-c96898e9c0be\n",
                  "endingDate": "2019-07-26 04:29:50.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:28:29.0",
                  "status": "ENDED",
                  "taskId": 6,
                  "taskName": "Create Ports"
                },
                {
                  "comment": "Ports attached successfully to the Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3.",
                  "endingDate": "2019-07-26 04:31:33.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:29:50.0",
                  "status": "ENDED",
                  "taskId": 7,
                  "taskName": "Attach Ports"
                },
                {
                  "comment": "Openstack Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3 started successfully.\nServer Status : ACTIVE\nTask State : -\nPower State : Running\n",
                  "endingDate": "2019-07-26 04:31:47.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:31:33.0",
                  "status": "ENDED",
                  "taskId": 8,
                  "taskName": "Start the UTM"
                },
                {
                  "comment": "IP Address 100.76.96.230 is now reachable from MSA.\nPING Status : OK\n",
                  "endingDate": "2019-07-26 04:32:30.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:31:47.0",
                  "status": "ENDED",
                  "taskId": 9,
                  "taskName": "Wait for UTM Ping reachability from MSA"
                },
                {
                  "comment": "OK LICENSE IS VALID",
                  "endingDate": "2019-07-26 04:32:56.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:32:30.0",
                  "status": "ENDED",
                  "taskId": 10,
                  "taskName": "Verify License Validity"
                },
                {
                  "comment": "Ports updated successfully on Fortigate Device 11892.\n",
                  "endingDate": "2019-07-26 04:33:17.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:32:56.0",
                  "status": "ENDED",
                  "taskId": 11,
                  "taskName": "Update UTM"
                },
                {
                  "comment": "Device 11892 Backup completed successfully.\nBackup Status : ENDED\nBackup Message : BACKUP  processed\n\nBackup Revision Id : 209408\n",
                  "endingDate": "2019-07-26 04:33:28.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:33:17.0",
                  "status": "ENDED",
                  "taskId": 12,
                  "taskName": "Device Backup"
                },
                {
                  "comment": "Ping Monitoring started for the device 11892.",
                  "endingDate": "2019-07-26 04:34:56.0",
                  "execNumber": 1,
                  "newParameters": {},
                  "processInstanceId": 85385,
                  "startingDate": "2019-07-26 04:33:28.0",
                  "status": "ENDED",
                  "taskId": 13,
                  "taskName": "Start Ping Monitoring"
                }
              ]
            }
          }
        }
expectedStatus:
    - InterfaceIsUpdating
newStatus: InterfaceIsUpdated
counter:
    min: 4
`

var testMockSecurityV3NetworkBasedWAFSingleListAfterUpdateInterface = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOW-01",
            "message": "Successful completion",
            "records": 2,
            "rows": [
                {
                    "id": 1,
                    "cell": ["false", "1", "CES11810", "WAF", "02", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "WAF", "08", "%s", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)
