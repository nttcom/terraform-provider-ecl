package ecl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	// "github.com/hashicorp/terraform/terraform"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func TestMockedAccVNAV1ApplianceBasic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	// region := "RegionOne"
	// if OS_REGION_NAME != "" {
	// 	region = OS_REGION_NAME
	// }
	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	log.Printf("[MYDEBUG] postKeystoneResponse: %+v", postKeystoneResponse)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
				),
			},
		},
	})
}

var testMockedAccVNAV1ApplianceBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "%s"

	interface_1_meta  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "dummyNetworkID"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	lifecycle {
		ignore_changes = [
			"default_gateway",
		]
	}
}`,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceDelete = `
request:
    method: DELETE
response:
    code: 204
expectedStatus:
    - Created
newStatus: Deleted
`

var testMockVNAV1AppliancePost = `
request:
    method: POST
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "zone1-groupb",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "1ad60fdc-a31f-4dc3-b0af-353c630c7708",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Passw0rd",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "initial"
            }
        }
newStatus: Created
`

var testMockVNAV1ApplianceProcessingAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "zone1-groupb",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "1ad60fdc-a31f-4dc3-b0af-353c630c7708",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "initial"
            }
        }
expectedStatus:
    - Created
counter:
    max: 3
`

var testMockVNAV1ApplianceGetActiveAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "zone1-groupb",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "1ad60fdc-a31f-4dc3-b0af-353c630c7708",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
counter:
    min: 4
`

var testMockVNAV1ApplianceGetDeleteComplete = `
request:
    method: GET
response:
    code: 404
expectedStatus:
    - Deleted
counter:
    min: 4
`

var testMockVNAV1ApplianceGetProcessingAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "zone1-groupb",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "1ad60fdc-a31f-4dc3-b0af-353c630c7708",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`
