package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v3/network_based_device_ha"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

const SoIDOfCreateHA = "FGHA_809F858574E94699952D0D7E7C58C81B"
const SoIDOfUpdateHA = "FGHA_809F858574E94699952D0D7E7C58C81C"
const SoIDOfDeleteHA = "FGHA_F2349100C7D24EF3ACD6B9A9F91FD220"

const ProcessIDOfUpdateInterfaceHA = 85385

const expectedNewHADeviceHostName1 = "CES12085"
const expectedNewHADeviceUUID1 = "12768064-e7c9-44d1-b01d-e66f138a278e"
const expectedNewHADeviceUUID2 = "12768064-e7c9-44d1-b01d-e66f138a278f"

func TestMockedAccSecurityV3NetworkBasedDeviceHA_basic(t *testing.T) {
	var hd1, hd2 security.HADevice

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

	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV3NetworkBasedDeviceHAUpdate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetProcessingAfterUpdate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetCompleteActiveAfterUpdate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListAfterUpdate)

	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV3NetworkBasedDeviceHADelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAProcessingAfterDelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetDeleteComplete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedDeviceHADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceHABasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v3.ha_1", &hd1, &hd2),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"operating_mode", "FW_HA"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"host_1_az_group", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"host_2_az_group", OS_COMPUTE_ZONE_HA),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceHAUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v3.ha_1", &hd1, &hd2),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"operating_mode", "UTM_HA"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"host_1_az_group", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1",
						"host_2_az_group", OS_COMPUTE_ZONE_HA),
				),
			},
		},
	})
}

func TestMockedAccSecurityV3NetworkBasedDeviceHAUpdateInterface(t *testing.T) {
	var hd1, hd2 security.HADevice

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListBeforeCreate)
	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV3NetworkBasedDeviceHACreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetProcessingAfterCreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetCompleteActiveAfterCreate)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListAfterCreate)
	mc.Register(t, "ha_device", "/ecl-api/devices", testMockSecurityV3NetworkBasedDeviceHAListDevicesAfterCreate)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID1), testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfacesAfterCreate)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID2), testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfacesAfterCreate)

	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/ports/utm/ha/%s", expectedNewHADeviceHostName1), testMockSecurityV3NetworkBasedDeviceHAUpdateInterface)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/process/%d/status", ProcessIDOfUpdateInterfaceHA), testMockSecurityV3NetworkBasedDeviceHAGetProcessingAfterUpdateInterface)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/process/%d/status", ProcessIDOfUpdateInterfaceHA), testMockSecurityV3NetworkBasedDeviceHAGetCompleteActiveAfterUpdateInterface)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListAfterInterfaceUpdate)
	mc.Register(t, "ha_device", "/ecl-api/devices", testMockSecurityV3NetworkBasedDeviceHAListDevicesAfterUpdate)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID1), testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfacesOfHost1AfterUpdate)
	mc.Register(t, "ha_device", fmt.Sprintf("/ecl-api/devices/%s/interfaces", expectedNewHADeviceUUID2), testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfacesOfHost2AfterUpdate)

	mc.Register(t, "ha_device", "/API/SoEntryFGHA", testMockSecurityV3NetworkBasedDeviceHADelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAProcessingAfterDelete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGSOrderProgressRate", testMockSecurityV3NetworkBasedDeviceHAGetDeleteComplete)
	mc.Register(t, "ha_device", "/API/ScreenEventFGHADeviceGet", testMockSecurityV3NetworkBasedDeviceHAListAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV3NetworkBasedDeviceHADestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceHABasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v3.ha_1", &hd1, &hd2),
				),
			},
			resource.TestStep{
				Config: testMockedAccSecurityV3NetworkBasedDeviceHAUpdateInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV3NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v3.ha_1", &hd1, &hd2),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.vrrp_ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.network_id", "dummyNetwork-port1"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.subnet_id", "dummySubnet-port1"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.comment", "port 0 comment"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.host_1_ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.host_1_ip_address_prefix", "24"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.host_2_ip_address", "10.0.0.2"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.0.host_2_ip_address_prefix", "24"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.vrrp_ip_address", "192.168.1.51"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.network_id", "dummyNetwork-port2"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.subnet_id", "dummySubnet-port2"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.comment", "port 3 comment"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.host_1_ip_address", "10.0.0.3"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.host_1_ip_address_prefix", "24"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.host_2_ip_address", "10.0.0.4"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v3.ha_1", "port.3.host_2_ip_address_prefix", "24"),
				),
			},
		},
	})
}

var testMockedAccSecurityV3NetworkBasedDeviceHABasic = fmt.Sprintf(`
resource "ecl_security_network_based_device_ha_v3" "ha_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW_HA"
	license_kind = "02"

	host_1_az_group = "%s"
	host_2_az_group = "%s"

	ha_link_1 {
		network_id = "DummyNetwork1"
		subnet_id = "DummySubnet1"
		host_1_ip_address = "192.168.1.3"
		host_2_ip_address = "192.168.1.4"
	}

	ha_link_2 {
		network_id = "DummyNetwork2"
		subnet_id = "DummySubnet2"
		host_1_ip_address = "192.168.2.3"
		host_2_ip_address = "192.168.2.4"
	}
}
`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testMockedAccSecurityV3NetworkBasedDeviceHAUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_device_ha_v3" "ha_1" {
	tenant_id = "%s"
	locale = "en"
	operating_mode = "UTM_HA"
	license_kind = "08"

	host_1_az_group = "%s"
	host_2_az_group = "%s"

	ha_link_1 {
		network_id = "DummyNetwork1"
		subnet_id = "DummySubnet1"
		host_1_ip_address = "192.168.1.3"
		host_2_ip_address = "192.168.1.4"
	}

	ha_link_2 {
		network_id = "DummyNetwork2"
		subnet_id = "DummySubnet2"
		host_1_ip_address = "192.168.2.3"
		host_2_ip_address = "192.168.2.4"
	}
}`,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testMockedAccSecurityV3NetworkBasedDeviceHAUpdateInterface = fmt.Sprintf(`

resource "ecl_security_network_based_device_ha_v3" "ha_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW_HA"
	license_kind = "02"

	host_1_az_group = "%s"
	host_2_az_group = "%s"

	ha_link_1 {
		network_id = "DummyNetwork1"
		subnet_id = "DummySubnet1"
		host_1_ip_address = "192.168.1.3"
		host_2_ip_address = "192.168.1.4"
	}

	ha_link_2 {
		network_id = "DummyNetwork2"
		subnet_id = "DummySubnet2"
		host_1_ip_address = "192.168.2.3"
		host_2_ip_address = "192.168.2.4"
	}

	port {
	    enable = "true"

	    network_id = "dummyNetwork-port1"
	    subnet_id = "dummySubnet-port1"
	    mtu = "1500"
		comment = "port 0 comment"
		enable_ping = "true"

		host_1_ip_address = "10.0.0.1"
		host_1_ip_address_prefix = 24

		host_2_ip_address = "10.0.0.2"
		host_2_ip_address_prefix = 24

		vrrp_ip_address = "192.168.1.50"
		vrrp_grp_id = "1"
		vrrp_id = "1"
		preempt = "true"
	}

	port {
	  enable = "false"
	}

	port {
	  enable = "false"
	}

	port {
	    enable = "true"

	    network_id = "dummyNetwork-port2"
	    subnet_id = "dummySubnet-port2"
	    mtu = "1500"
		comment = "port 3 comment"
		enable_ping = "true"

		host_1_ip_address = "10.0.0.3"
		host_1_ip_address_prefix = 24

		host_2_ip_address = "10.0.0.4"
		host_2_ip_address_prefix = 24

		vrrp_ip_address = "192.168.1.51"
		vrrp_grp_id = "1"
		vrrp_id = "2"
		preempt = "true"
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
	OS_COMPUTE_ZONE_HA,
)

var testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfaces = `
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

var testMockSecurityV3NetworkBasedDeviceHAListBeforeCreate = fmt.Sprintf(`
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
                    "cell": ["false", "1", "1902F60E", "CES12083", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.3", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.3"],
                    "id": 1
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12084", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.4", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.4"],
                    "id": 2
                }
            ]
        }
expectedStatus:
    - ""
newStatus: PreCreate
`,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testMockSecurityV3NetworkBasedDeviceHAListAfterCreate = fmt.Sprintf(`
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
                    "cell": ["false", "1", "1902F60E", "CES12083", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.3", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.3"],
                    "id": 1
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12084", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.4", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.4"],
                    "id": 2
                },
                {
                    "cell": ["false", "1", "1902F60E", "CES12085", "FW_HA", "02", "ha", "%s", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.3", "dummyNetworkID2", "dummySubnetID2", "192.168.2.3"],
                    "id": 3
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12086", "FW_HA", "02", "ha", "%s", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.4", "dummyNetworkID2", "dummySubnetID2", "192.168.2.4"],
                    "id": 4
                }
            ]
        }
expectedStatus:
    - Created
    - Updating
`,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testMockSecurityV3NetworkBasedDeviceHAListDevicesAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "devices": [
            {
                "msa_device_id": "CES12085",
                "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278e",
                "os_server_name": "UTM-CES11878",
                "os_availability_zone": "%s",
                "os_admin_username": "jp4_sdp_mss_utm_admin",
                "msa_device_type": "FW",
                "os_server_status": "ACTIVE"
            },
            {
                "msa_device_id": "CES12086",
                "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278f",
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
	OS_COMPUTE_ZONE_HA,
)

var testMockSecurityV3NetworkBasedDeviceHAListDevicesAfterUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "devices": [
            {
                "msa_device_id": "CES12085",
                "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278e",
                "os_server_name": "UTM-CES11878",
                "os_availability_zone": "%s",
                "os_admin_username": "jp4_sdp_mss_utm_admin",
                "msa_device_type": "FW",
                "os_server_status": "ACTIVE"
            },
            {
                "msa_device_id": "CES12086",
                "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278f",
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
	OS_COMPUTE_ZONE_HA,
	OS_COMPUTE_ZONE_HA,
)

var testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfacesAfterCreate = `
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

var testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfacesOfHost1AfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "device_interfaces": [
              {
                "os_ip_address": "192.168.1.3",
                "msa_port_id": "port2",
                "os_port_name": "HA-Port01-First-CES12136",
                "os_port_id": "a6471bfc-8d65-4fab-bdca-476cef33c0de",
                "os_network_id": "e51d6712-93b7-40e2-97be-3de06df02ffc",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:62:41:e1",
                "os_subnet_id": "1e925039-14d0-42b2-941a-69d180fcccdd"
              },
              {
                "os_ip_address": "10.0.0.1",
                "msa_port_id": "port4",
                "os_port_name": "port4-CES12136",
                "os_port_id": "e233ef27-6ccd-4be9-8da5-1316788ecf6b",
                "os_network_id": "dummyNetwork-port1",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:79:0a:4a",
                "os_subnet_id": "dummySubnet-port1"
              },
              {
                "os_ip_address": "10.0.0.3",
                "msa_port_id": "port7",
                "os_port_name": "port4-CES12136",
                "os_port_id": "e233ef27-6ccd-4be9-8da5-1316788ecf6b",
                "os_network_id": "a16627e2-1043-4cbb-b032-8a68221c9e60",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:79:0a:4a",
                "os_subnet_id": "ea3085cd-8cbb-4b7b-871f-2a6514fabd33"
              },
              {
                "os_ip_address": "192.168.2.3",
                "msa_port_id": "port3",
                "os_port_name": "HA-Port02-First-CES12136",
                "os_port_id": "e4cea571-fd7b-4798-bda4-07f739e6f9ce",
                "os_network_id": "dummyNetwork-port2",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:c8:0a:10",
                "os_subnet_id": "dummySubnet-port2"
              }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`

var testMockSecurityV3NetworkBasedDeviceHAListDeviceInterfacesOfHost2AfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "device_interfaces": [
              {
                "os_ip_address": "192.168.1.3",
                "msa_port_id": "port2",
                "os_port_name": "HA-Port01-First-CES12136",
                "os_port_id": "a6471bfc-8d65-4fab-bdca-476cef33c0de",
                "os_network_id": "e51d6712-93b7-40e2-97be-3de06df02ffc",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:62:41:e1",
                "os_subnet_id": "1e925039-14d0-42b2-941a-69d180fcccdd"
              },
              {
                "os_ip_address": "10.0.0.2",
                "msa_port_id": "port4",
                "os_port_name": "port4-CES12136",
                "os_port_id": "e233ef27-6ccd-4be9-8da5-1316788ecf6b",
                "os_network_id": "dummyNetwork-port1",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:79:0a:4a",
                "os_subnet_id": "dummySubnet-port1"
              },
              {
                "os_ip_address": "10.0.0.4",
                "msa_port_id": "port7",
                "os_port_name": "port4-CES12136",
                "os_port_id": "e233ef27-6ccd-4be9-8da5-1316788ecf6b",
                "os_network_id": "a16627e2-1043-4cbb-b032-8a68221c9e60",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:79:0a:4a",
                "os_subnet_id": "ea3085cd-8cbb-4b7b-871f-2a6514fabd33"
              },
              {
                "os_ip_address": "192.168.2.3",
                "msa_port_id": "port3",
                "os_port_name": "HA-Port02-First-CES12136",
                "os_port_id": "e4cea571-fd7b-4798-bda4-07f739e6f9ce",
                "os_network_id": "dummyNetwork-port2",
                "os_port_status": "ACTIVE",
                "os_mac_address": "fa:16:3e:c8:0a:10",
                "os_subnet_id": "dummySubnet-port2"
              }
          ]
        }
expectedStatus:
    - InterfaceIsUpdated
`

var testMockSecurityV3NetworkBasedDeviceHAListAfterDelete = fmt.Sprintf(`
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
                    "cell": ["false", "1", "1902F60E", "CES12083", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.3", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.3"],
                    "id": 1
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12084", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.4", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.4"],
                    "id": 2
                }
            ]
        }
expectedStatus:
    - Deleted
`,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testMockSecurityV3NetworkBasedDeviceHACreate = fmt.Sprintf(`
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

var testMockSecurityV3NetworkBasedDeviceHAGetProcessingAfterCreate = `
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

var testMockSecurityV3NetworkBasedDeviceHAGetCompleteActiveAfterCreate = `
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

var testMockSecurityV3NetworkBasedDeviceHADelete = fmt.Sprintf(`
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

var testMockSecurityV3NetworkBasedDeviceHAProcessingAfterDelete = `
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

var testMockSecurityV3NetworkBasedDeviceHAGetDeleteComplete = `
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

var testMockSecurityV3NetworkBasedDeviceHAUpdate = fmt.Sprintf(`
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

var testMockSecurityV3NetworkBasedDeviceHAGetProcessingAfterUpdate = `
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

var testMockSecurityV3NetworkBasedDeviceHAGetCompleteActiveAfterUpdate = `
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

var testMockSecurityV3NetworkBasedDeviceHAListAfterUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "status": 1,
            "code": "FOV-01",
            "message": "Successful completion",
            "records": 4,
            "rows": [
                {
                    "cell": ["false", "1", "1902F60E", "CES12083", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.3", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.3"],
                    "id": 1
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12084", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.4", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.4"],
                    "id": 2
                },
                {
                    "cell": ["false", "1", "1902F60E", "CES12085", "UTM_HA", "08", "ha", "%s", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.3", "dummyNetworkID2", "dummySubnetID2", "192.168.2.3"],
                    "id": 3
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12086", "UTM_HA", "08", "ha", "%s", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.4", "dummyNetworkID2", "dummySubnetID2", "192.168.2.4"],
                    "id": 4
                }
            ]
        }
expectedStatus:
    - Updated
`,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testMockSecurityV3NetworkBasedDeviceHAListAfterInterfaceUpdate = fmt.Sprintf(`
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
                    "cell": ["false", "1", "1902F60E", "CES12083", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.3", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.3"],
                    "id": 1
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12084", "UTM_HA", "02", "ha", "%s", "jp4_zone1", "56a5f5b0-dceb-47d9-8e75-8e16dc08d83f", "6b3ee9c8-0f28-41ff-a443-a3122cf89f1f", "192.168.1.4", "bfb4dcb7-f8fd-4ae8-9023-9c648e56b455", "085ea95a-a04b-4eb4-bdfc-124445fb5cec", "192.168.2.4"],
                    "id": 2
                },
                {
                    "cell": ["false", "1", "1902F60E", "CES12085", "FW_HA", "02", "ha", "%s", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.3", "dummyNetworkID2", "dummySubnetID2", "192.168.2.3"],
                    "id": 3
                }, 
                {
                    "cell": ["false", "2", "1902F60E", "CES12086", "FW_HA", "02", "ha", "%s", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.4", "dummyNetworkID2", "dummySubnetID2", "192.168.2.4"],
                    "id": 4
                }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testMockSecurityV3NetworkBasedDeviceHAUpdateInterface = fmt.Sprintf(`
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

var testMockSecurityV3NetworkBasedDeviceHAGetProcessingAfterUpdateInterface = `
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

var testMockSecurityV3NetworkBasedDeviceHAGetCompleteActiveAfterUpdateInterface = `
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

var testMockSecurityV3NetworkBasedDeviceHAListAfterUpdateInterface = fmt.Sprintf(`
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
                    "cell": ["false", "1", "CES777", "FW", "02", "standalone", "%s", "jp4_zone1"]
                },
                {
                    "id": 2,
                    "cell": ["false", "1", "CES888", "UTM", "08", "standalone", "%s", "jp4_zone1"]
                }
            ]
        }
expectedStatus:
    - InterfaceIsUpdated
`,
	OS_COMPUTE_ZONE_HA,
	OS_COMPUTE_ZONE_HA,
)
