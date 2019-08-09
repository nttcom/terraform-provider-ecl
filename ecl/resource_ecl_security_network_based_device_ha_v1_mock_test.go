package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_ha"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfCreateHA = "FGHA_809F858574E94699952D0D7E7C58C81B"
const SoIDOfUpdateHA = "FGHA_809F858574E94699952D0D7E7C58C81C"
const SoIDOfDeleteHA = "FGHA_F2349100C7D24EF3ACD6B9A9F91FD220"

const ProcessIDOfUpdateInterfaceHA = 85385

const expectedNewHADeviceHostName1 = "CES12085"
const expectedNewHADeviceHostName2 = "CES12086"
const expectedNewHADeviceUUID = "12768064-e7c9-44d1-b01d-e66f138a278e"

func TestMockedAccSecurityV1NetworkBasedDeviceHABasic(t *testing.T) {
	var hd1, hd2 security.HADevice

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	// TODO
	// mc.Register(t, "ha_device", "/ecl-api/devices", testMockSecurityV1NetworkBasedDeviceHAListDevicesAfterCreate)
	// mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID), testMockSecurityV1NetworkBasedDeviceHAListDevicesAfterCreate)

	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV1NetworkBasedDeviceHAListBeforeCreate)
	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV1NetworkBasedDeviceHACreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedDeviceHAGetProcessingAfterCreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedDeviceHAGetCompleteActiveAfterCreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV1NetworkBasedDeviceHAListAfterCreate)

	// TODO
	// mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV1NetworkBasedDeviceHAUpdate)
	// mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedDeviceHAGetProcessingAfterUpdate)
	// mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedDeviceHAGetCompleteActiveAfterUpdate)
	// mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV1NetworkBasedDeviceHAListAfterUpdate)

	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV1NetworkBasedDeviceHADelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedDeviceHAProcessingAfterDelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV1NetworkBasedDeviceHAGetDeleteComplete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV1NetworkBasedDeviceHAListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedDeviceHADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV1NetworkBasedDeviceHABasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v1.ha_1", &hd1, &hd2),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"operating_mode", "FW_HA"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"host_1_az_group", "zone1-groupa"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"host_2_az_group", "zone1-groupb"),
				),
			},
			// resource.TestStep{
			// 	Config: testMockedAccSecurityV1NetworkBasedDeviceHAUpdate,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckSecurityV1NetworkBasedDeviceHAExists(
			// 			"ecl_security_network_based_device_ha_v1.ha_1", &hd1, &hd2),
			// 		resource.TestCheckResourceAttr(
			// 			"ecl_security_network_based_device_ha_v1.ha_1",
			// 			"locale", "en"),
			// 		resource.TestCheckResourceAttr(
			// 			"ecl_security_network_based_device_ha_v1.ha_1",
			// 			"operating_mode", "UTM"),
			// 		resource.TestCheckResourceAttr(
			// 			"ecl_security_network_based_device_ha_v1.ha_1",
			// 			"license_kind", "08"),
			// 		resource.TestCheckResourceAttr(
			// 			"ecl_security_network_based_device_ha_v1.ha_1",
			// 			"az_group", "zone1-groupb"),
			// 	),
			// },
		},
	})
}

var testMockedAccSecurityV1NetworkBasedDeviceHABasic = fmt.Sprintf(`
resource "ecl_security_network_based_device_ha_v1" "ha_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW_HA"
	license_kind = "02"

	host_1_az_group = "zone1-groupa"
	host_2_az_group = "zone1-groupb"

	ha_link_1 {
		network_id = "DummyNetwork1"
		subnet_id = "DummySubnet1"
		host_1_ip_address = "192.168.1.3"
		host_2_ip_address = "192.168.1.4"
	}

	ha_link_2 {
		network_id = "DummyNetwork12"
		subnet_id = "DummySubnet2"
		host_1_ip_address = "192.168.2.3"
		host_2_ip_address = "192.168.2.4"
	}

}
`,
	OS_TENANT_ID,
)

var testMockedAccSecurityV1NetworkBasedDeviceHAUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_device_ha_v1" "ha_1" {
	tenant_id = "%s"
	locale = "en"
	operating_mode = "UTM"
	license_kind = "08"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)

var testMockedAccSecurityV1NetworkBasedDeviceHAUpdateInterface = fmt.Sprintf(`

resource "ecl_security_network_based_device_ha_v1" "ha_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "zone1-groupb"

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
)

var testMockSecurityV1NetworkBasedDeviceHAListBeforeCreate = `
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
            "rows": [
                {
                    "cell": ["false", "1", "1902F60E", "CES12083", "UTM_HA", "02", "ha", "zone1-groupa", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.3", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.3"],
                    "id": 1
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12084", "UTM_HA", "02", "ha", "zone1-groupb", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.4", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.4"],
                    "id": 2
                }
            ]
        }
expectedStatus:
    - ""
newStatus: PreCreate
`

var testMockSecurityV1NetworkBasedDeviceHAListAfterCreate = `
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
            "rows": [
                {
                    "cell": ["false", "1", "1902F60E", "CES12083", "UTM_HA", "02", "ha", "zone1-groupa", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.3", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.3"],
                    "id": 1
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12084", "UTM_HA", "02", "ha", "zone1-groupb", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.4", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.4"],
                    "id": 2
                },
                {
                    "cell": ["false", "1", "1902F60E", "CES12085", "FW_HA", "02", "ha", "zone1-groupa", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.3", "dummyNetworkID2", "dummySubnetID2", "192.168.2.3"],
                    "id": 3
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12086", "FW_HA", "02", "ha", "zone1-groupb", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.4", "dummyNetworkID2", "dummySubnetID2", "192.168.2.4"],
                    "id": 4
                }
            ]
        }
expectedStatus:
    - Created
    - Updating
`

var testMockSecurityV1NetworkBasedDeviceHAListDevicesAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "devices": [
            {
              "msa_device_id": "CES777",
              "os_server_id": "392a90bf-2c1b-45fd-8221-096894fff39d",
              "os_server_name": "UTM-CES11878",
              "os_availability_zone": "zone1-groupb",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "FW",
              "os_server_status": "ACTIVE"
            },
            {
              "msa_device_id": "CES888",
              "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278e",
              "os_server_name": "WAF-CES11816",
              "os_availability_zone": "zone1-groupb",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "WAF",
              "os_server_status": "ACTIVE"
            }
          ]
        }
expectedStatus:
    - Updated
    - Created
`

var testMockSecurityV1NetworkBasedDeviceHAListDevicesAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "devices": [
            {
              "msa_device_id": "CES777",
              "os_server_id": "392a90bf-2c1b-45fd-8221-096894fff39d",
              "os_server_name": "UTM-CES11878",
              "os_availability_zone": "zone1-groupb",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "FW",
              "os_server_status": "ACTIVE"
            },
            {
              "msa_device_id": "CES888",
              "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278e",
              "os_server_name": "WAF-CES11816",
              "os_availability_zone": "zone1-groupb",
              "os_admin_username": "jp4_sdp_mss_utm_admin",
              "msa_device_type": "WAF",
              "os_server_status": "ACTIVE"
            }
          ]
        }
expectedStatus:
    - InterfaceIsUpdated
`

var testMockSecurityV1NetworkBasedDeviceHAListDeviceInterfacesAfterCreate = `
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

var testMockSecurityV1NetworkBasedDeviceHAListDeviceInterfacesAfterUpdate = `
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

var testMockSecurityV1NetworkBasedDeviceHAListAfterDelete = `
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
            "rows": [
                {
                    "cell": ["false", "1", "1902F60E", "CES12083", "UTM_HA", "02", "ha", "zone1-groupa", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.3", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.3"],
                    "id": 1
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12084", "UTM_HA", "02", "ha", "zone1-groupb", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.4", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.4"],
                    "id": 2
                }
            ]
        }
expectedStatus:
    - Deleted
`

var testMockSecurityV1NetworkBasedDeviceHACreate = fmt.Sprintf(`
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
	SoIDOfCreateHA,
)

var testMockSecurityV1NetworkBasedDeviceHAGetProcessingAfterCreate = `
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

var testMockSecurityV1NetworkBasedDeviceHAGetCompleteActiveAfterCreate = `
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

var testMockSecurityV1NetworkBasedDeviceHADelete = fmt.Sprintf(`
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
	SoIDOfDeleteHA,
)

var testMockSecurityV1NetworkBasedDeviceHAProcessingAfterDelete = `
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

var testMockSecurityV1NetworkBasedDeviceHAGetDeleteComplete = `
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

var testMockSecurityV1NetworkBasedDeviceHAUpdate = fmt.Sprintf(`
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
	SoIDOfUpdateHA,
)

var testMockSecurityV1NetworkBasedDeviceHAGetProcessingAfterUpdate = `
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

var testMockSecurityV1NetworkBasedDeviceHAGetCompleteActiveAfterUpdate = `
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

var testMockSecurityV1NetworkBasedDeviceHAListAfterUpdate = `
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
                    "cell": ["false", "1", "CES777", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES888", "UTM", "08", "standalone", "zone1-groupb", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - Updated
`

var testMockSecurityV1NetworkBasedDeviceHAListAfterInterfaceUpdate = `
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
                    "cell": ["false", "1", "CES777", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES888", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`

var testMockSecurityV1NetworkBasedDeviceHAUpdateInterface = fmt.Sprintf(`
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
	ProcessIDOfUpdateInterfaceHA,
)

var testMockSecurityV1NetworkBasedDeviceHAGetProcessingAfterUpdateInterface = `
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

var testMockSecurityV1NetworkBasedDeviceHAGetCompleteActiveAfterUpdateInterface = `
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

var testMockSecurityV1NetworkBasedDeviceHAListAfterUpdateInterface = `
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
                    "cell": ["false", "1", "CES777", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES888", "UTM", "08", "standalone", "zone1-groupb", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`
