package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v3/network_based_device_single"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfCreate = "FGS_809F858574E94699952D0D7E7C58C81B"
const SoIDOfUpdate = "FGS_809F858574E94699952D0D7E7C58C81C"
const SoIDOfDelete = "FGS_F2349100C7D24EF3ACD6B9A9F91FD220"

const ProcessIDOfUpdateInterface = 85385

const expectedNewSingleDeviceHostName = "CES11811"
const expectedNewSingleDeviceUUID = "12768064-e7c9-44d1-b01d-e66f138a278e"

func TestMockedAccSecurityV3NetworkBasedDeviceSingle_basic(t *testing.T) {
	var sd security.SingleDevice

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "single_device", "/ecl-api/devices", testMockSecurityV3NetworkBasedDeviceSingleListDevicesAfterCreate)
	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV3NetworkBasedDeviceSingleListDeviceInterfaces)

	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListBeforeCreate)
	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV3NetworkBasedDeviceSingleCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetProcessingAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetCompleteActiveAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListAfterCreate)

	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV3NetworkBasedDeviceSingleUpdate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetProcessingAfterUpdate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetCompleteActiveAfterUpdate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListAfterUpdate)

	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV3NetworkBasedDeviceSingleDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleProcessingAfterDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetDeleteComplete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"operating_mode", "FW"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"az_group", OS_DEFAULT_ZONE),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceSingleUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"operating_mode", "UTM"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1",
						"az_group", OS_DEFAULT_ZONE),
				),
			},
		},
	})
}

func TestMockedAccSecurityV3NetworkBasedDeviceSingleUpdateInterface(t *testing.T) {
	var sd security.SingleDevice

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListBeforeCreate)
	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV3NetworkBasedDeviceSingleCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetProcessingAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetCompleteActiveAfterCreate)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListAfterCreate)
	mc.Register(t, "single_device", "/ecl-api/devices", testMockSecurityV3NetworkBasedDeviceSingleListDevicesAfterCreate)
	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV3NetworkBasedDeviceSingleListDeviceInterfacesAfterCreate)

	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/ports/utm/%s", expectedNewSingleDeviceHostName), testMockSecurityV3NetworkBasedDeviceSingleUpdateInterface)
	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/process/%d/status", ProcessIDOfUpdateInterface), testMockSecurityV3NetworkBasedDeviceSingleGetProcessingAfterUpdateInterface)
	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/process/%d/status", ProcessIDOfUpdateInterface), testMockSecurityV3NetworkBasedDeviceSingleGetCompleteActiveAfterUpdateInterface)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListAfterInterfaceUpdate)
	mc.Register(t, "single_device", "/ecl-api/devices", testMockSecurityV3NetworkBasedDeviceSingleListDevicesAfterUpdate)
	mc.Register(t, "single_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewSingleDeviceUUID), testMockSecurityV3NetworkBasedDeviceSingleListDeviceInterfacesAfterUpdate)

	mc.Register(t, "single_device", "/API/SoEntryFGS", testMockSecurityV3NetworkBasedDeviceSingleDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleProcessingAfterDelete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceSingleGetDeleteComplete)
	mc.Register(t, "single_device", "/API/ScreenEventFGSDeviceGet", testMockSecurityV3NetworkBasedDeviceSingleListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceSingleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceSingleUpdateInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceSingleExists(
						"ecl_security_network_based_device_single_v3.device_1", &sd),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.network_id", "dummyNetwork1"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.subnet_id", "dummySubnet1"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.0.comment", "port 0 comment"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.ip_address", "192.168.2.50"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.network_id", "dummyNetwork2"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.subnet_id", "dummySubnet2"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_single_v3.device_1", "port.3.comment", "port 3 comment"),
				),
			},
		},
	})
}

var testMockedAccSecurityV3NetworkBasedDeviceSingleBasic = fmt.Sprintf(`
resource "ecl_security_network_based_device_single_v3" "device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "%s"
}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testMockedAccSecurityV3NetworkBasedDeviceSingleUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_device_single_v3" "device_1" {
	tenant_id = "%s"
	locale = "en"
	operating_mode = "UTM"
	license_kind = "08"
	az_group = "%s"
}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testMockedAccSecurityV3NetworkBasedDeviceSingleUpdateInterface = fmt.Sprintf(`

resource "ecl_security_network_based_device_single_v3" "device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
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

  port {
    enable = "false"
  }

  port {
    enable = "false"
  }
  
  port {
      enable = "true"
      ip_address = "192.168.2.50"
      ip_address_prefix = 24
      network_id = "dummyNetwork2"
      subnet_id = "dummySubnet2"
      mtu = "1500"
      comment = "port 3 comment"
  }

  port {
    enable = "false"
  }
  port {
    enable = "false"
  }
  port {
    enable = "false"
  }

}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedDeviceSingleListBeforeCreate = fmt.Sprintf(`
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
            	"cell": ["false", "1", "CES11810", "FW", "02", "standalone", "%s", "jp4_zone1"]
            }]
        }
expectedStatus:
    - ""
newStatus: PreCreate
`,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedDeviceSingleListAfterCreate = fmt.Sprintf(`
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
                    "cell": ["false", "1", "CES11810", "FW", "02", "standalone", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "FW", "02", "standalone", "%s", "jp4_zone1"]
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

var testMockSecurityV3NetworkBasedDeviceSingleListDeviceInterfaces = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "devices": []
        }
expectedStatus:
    - Updated
    - Created
`

var testMockSecurityV3NetworkBasedDeviceSingleListDevicesAfterCreate = fmt.Sprintf(`
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
              "os_server_name": "UTM-CES11878",
              "os_availability_zone": "%s",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "FW",
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

var testMockSecurityV3NetworkBasedDeviceSingleListDevicesAfterUpdate = fmt.Sprintf(`
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
                "os_server_name": "UTM-CES11878",
                "os_availability_zone": "%s",
                "os_admin_username": "jp4_sdp_mss_utm_admin",
                "msa_device_type": "FW",
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

var testMockSecurityV3NetworkBasedDeviceSingleListDeviceInterfacesAfterCreate = `
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

var testMockSecurityV3NetworkBasedDeviceSingleListDeviceInterfacesAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "device_interfaces": [
            {
                "os_ip_address": "192.168.1.50",
                "msa_port_id": "port4",
                "os_port_name": "port4-CES11892",
                "os_port_id": "82ebe045-9c9a-4088-8b33-cb0d590079aa",
                "os_network_id": "dummyNetwork1",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:05:ff:66",
                "os_subnet_id": "dummySubnet1"
            },
            {
                "os_ip_address": "192.168.2.50",
                "msa_port_id": "port7",
                "os_port_name": "port7-CES11892",
                "os_port_id": "82ebe045-9c9a-4088-8b33-cb0d590079aa",
                "os_network_id": "dummyNetwork2",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:05:ff:67",
                "os_subnet_id": "dummySubnet2"
            }
          ]
        }
expectedStatus:
    - InterfaceIsUpdated
`

var testMockSecurityV3NetworkBasedDeviceSingleListAfterDelete = fmt.Sprintf(`
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
                "cell": ["false", "1", "CES11810", "FW", "02", "standalone", "%s", "jp4_zone1"]
            }]
        }
expectedStatus:
    - Deleted
`,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedDeviceSingleCreate = fmt.Sprintf(`
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

var testMockSecurityV3NetworkBasedDeviceSingleGetProcessingAfterCreate = `
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

var testMockSecurityV3NetworkBasedDeviceSingleGetCompleteActiveAfterCreate = `
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

var testMockSecurityV3NetworkBasedDeviceSingleDelete = fmt.Sprintf(`
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
    - InterfaceIsUpdated
newStatus: Deleted
`,
	SoIDOfDelete,
)

var testMockSecurityV3NetworkBasedDeviceSingleProcessingAfterDelete = `
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

var testMockSecurityV3NetworkBasedDeviceSingleGetDeleteComplete = `
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

var testMockSecurityV3NetworkBasedDeviceSingleUpdate = fmt.Sprintf(`
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

var testMockSecurityV3NetworkBasedDeviceSingleGetProcessingAfterUpdate = `
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

var testMockSecurityV3NetworkBasedDeviceSingleGetCompleteActiveAfterUpdate = `
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

var testMockSecurityV3NetworkBasedDeviceSingleListAfterUpdate = fmt.Sprintf(`
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
                    "cell": ["false", "1", "CES11810", "FW", "02", "standalone", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "UTM", "08", "standalone", "%s", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - Updated
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedDeviceSingleListAfterInterfaceUpdate = fmt.Sprintf(`
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
                    "cell": ["false", "1", "CES11810", "FW", "02", "standalone", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "FW", "02", "standalone", "%s", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)

var testMockSecurityV3NetworkBasedDeviceSingleUpdateInterface = fmt.Sprintf(`
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

var testMockSecurityV3NetworkBasedDeviceSingleGetProcessingAfterUpdateInterface = `
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

var testMockSecurityV3NetworkBasedDeviceSingleGetCompleteActiveAfterUpdateInterface = `
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

var testMockSecurityV3NetworkBasedDeviceSingleListAfterUpdateInterface = fmt.Sprintf(`
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
                    "cell": ["false", "1", "CES11810", "FW", "02", "standalone", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES11811", "UTM", "08", "standalone", "%s", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE,
)
