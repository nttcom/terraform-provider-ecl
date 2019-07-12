package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func TestMockedAccVNAV1ApplianceUpdateMetaBasic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceMetaPatch1)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterMetaUpdate1)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterMetaUpdate1)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceMetaPatch2)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterMetaUpdate2)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterMetaUpdate2)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
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
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					testAccCheckVNAV1ApplianceTag(&vna, "k1", "v1"),
				),
			},
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceUpdateMetaBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description-update"),
					testAccCheckVNAV1ApplianceTag(&vna, "k1", "v1"),
					testAccCheckVNAV1ApplianceTag(&vna, "k2", "v2"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description-update"),
					testAccCheckVNAV1ApplianceInterfaceTag(&vna.Interfaces.Interface1, "interfaceK1", "interfaceV1"),
					testAccCheckVNAV1ApplianceInterfaceTag(&vna.Interfaces.Interface1, "interfaceK2", "interfaceV2"),
				),
			},
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceUpdateMetaBasic2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description-update"),

					testAccCheckVNAV1ApplianeTagLengthIsZERO(&vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description-update"),

					testAccCheckVNAV1ApplianceInterfaceTagLengthIsZERO(&vna.Interfaces.Interface1),
				),
			},
		},
	})
}

func TestMockedAccVNAV1ApplianceUpdateAllowedAddressPairBasic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceAllowedAddressPairPatch)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterAllowedAddressPairUpdate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterAllowedAddressPairUpdate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
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
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceUpdateAllowedAddressPairBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.ip_address", "192.168.1.200"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.type", "vrrp"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.vrid", "123"),
				),
			},
		},
	})
}

func TestMockedAccVNAV1ApplianceUpdateFixedIPBasic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceFixedIPPatch)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterFixedIPUpdate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterFixedIPUpdate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
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
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceUpdateFixedIPBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					// Check network id in interface metadata part
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", "dummyNetworkID"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_2_info.0.network_id", "dummyNetworkID2"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.network_id", "dummyNetworkID3"),
					// Check fixed_ips part
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 1, "192.168.1.50"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 2, "192.168.2.101"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 3, "192.168.3.50"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 3, "192.168.3.60"),
				),
			},
		},
	})
}

func TestMockedAccVNAV1ApplianceBasic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
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
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", "dummyNetworkID"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
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

    tags = {
        k1 = "v1"
    }

    interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
        network_id = "dummyNetworkID"
        tags = {
            interfaceK1 = "interfaceV1"
            interfaceK2 = "interfaceV2"
        }
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

var testMockedAccVNAV1ApplianceUpdateAllowedAddressPairBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "%s"

    tags = {
        k1 = "v1"
    }

    interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
        network_id = "dummyNetworkID"
        tags = {
            interfaceK1 = "interfaceV1"
            interfaceK2 = "interfaceV2"
        }
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_1_allowed_address_pairs {
		ip_address = "192.168.1.200"
		mac_address = "aa:bb:cc:dd:ee:f1"
		type = "vrrp"
		vrid = "123"
	}

	interface_1_allowed_address_pairs {
		ip_address = "192.168.1.201"
		mac_address = "aa:bb:cc:dd:ee:f2"
		type = ""
		vrid = "null"
	}

    lifecycle {
		ignore_changes = [
			"default_gateway",
		]
	}
}`,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceUpdateMetaBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1-update"
	description = "appliance_1_description-update"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "%s"

    tags = {
        k1 = "v1"
        k2 = "v2"
    }

    interface_1_info  {
		name = "interface_1-update"
		description = "interface_1_description-update"
        network_id = "dummyNetworkID"
        tags = {
            interfaceK1 = "interfaceV1"
            interfaceK2 = "interfaceV2"
        }
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

var testMockedAccVNAV1ApplianceUpdateMetaBasic2 = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1-update"
	description = "appliance_1_description-update"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "%s"

    interface_1_info  {
		name = "interface_1-update"
		description = "interface_1_description-update"
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

var testMockedAccVNAV1ApplianceUpdateFixedIPBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "%s"

    tags = {
        k1 = "v1"
    }

    interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
        network_id = "dummyNetworkID"
        tags = {
            interfaceK1 = "interfaceV1"
            interfaceK2 = "interfaceV2"
        }
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

    interface_2_info  {
        network_id = "dummyNetworkID2"
	}

    interface_3_info  {
        network_id = "dummyNetworkID3"
	}

    interface_3_fixed_ips {
		ip_address = "192.168.3.50"
	}

    interface_3_fixed_ips {
		ip_address = "192.168.3.60"
	}

    interface_4_info  {
        network_id = "dummyNetworkID4"
	}

    interface_4_no_fixed_ips = "true"

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
    - Updated
    - Updated1
    - Updated2
newStatus: Deleted
`

var testMockVNAV1ApplianceMetaPatch1 = `
request:
    method: PATCH
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "zone1-groupb",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
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
                "name": "appliance_1-update",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {
                    "k1": "v1",
                    "k2": "v2"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
newStatus: Updated1
`

var testMockVNAV1ApplianceMetaPatch2 = `
request:
    method: PATCH
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "zone1-groupb",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
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
                "name": "appliance_1-update",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated1
newStatus: Updated2
`

var testMockVNAV1ApplianceFixedIPPatch = `
request:
    method: PATCH
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {
        			"k1": "v1"
        		},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "zone1-groupb",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "PROCESSING",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID3",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.3.50",
        					"subnet_id": ""
        				}, {
        					"ip_address": "192.168.3.60",
        					"subnet_id": ""
        				}]
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID4",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID2",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Created
newStatus: Updated
`

var testMockVNAV1ApplianceAllowedAddressPairPatch = `
request:
    method: PATCH
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {
        			"k1": "v1"
        		},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "zone1-groupb",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "PROCESSING",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [
                            {
                                "ip_address": "192.168.1.200",
                                "mac_address": "aa:bb:cc:dd:ee:f1",
                                "type": "vrrp",
                                "vrid": 123
                            },
                            {
                                "ip_address": "192.168.1.201",
                                "mac_address": "aa:bb:cc:dd:ee:f2",
                                "type": "",
                                "vrid": null
                            }
                        ],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
- Created
newStatus: Updated
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
                        "network_id": "dummyNetworkID",
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
                "tags": {
                    "k1": "v1"
                },
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
                        "network_id": "dummyNetworkID",
                        "tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
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
                "tags": {
                    "k1": "v1"
                },
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
                        "network_id": "dummyNetworkID",
                        "tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
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
                "tags": {
                    "k1": "v1"
                },
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

var testMockVNAV1ApplianceProcessingAfterDelete = `
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
                        "network_id": "dummyNetworkID",
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
                "tags": {
                    "k1": "v1"
                },
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

var testMockVNAV1ApplianceProcessingAfterMetaUpdate1 = `
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
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
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
                "name": "appliance_1-update",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {
                    "k1": "v1",
                    "k2": "v2"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated1
counter:
    max: 3
`

var testMockVNAV1ApplianceGetActiveAfterMetaUpdate1 = `
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
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
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
                "name": "appliance_1-update",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {
                    "k1": "v1",
                    "k2": "v2"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated1
counter:
    min: 4
`
var testMockVNAV1ApplianceProcessingAfterMetaUpdate2 = `
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
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
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
                "name": "appliance_1-update",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated2
counter:
    max: 3
`

var testMockVNAV1ApplianceGetActiveAfterMetaUpdate2 = `
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
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
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
                "name": "appliance_1-update",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated2
counter:
    min: 4
`

var testMockVNAV1ApplianceProcessingAfterFixedIPUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {
        			"k1": "v1"
        		},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "zone1-groupb",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "PROCESSING",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID3",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.3.50",
        					"subnet_id": ""
        				}, {
        					"ip_address": "192.168.3.60",
        					"subnet_id": ""
        				}]
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID4",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID2",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Updated
counter:
    max: 3
`

var testMockVNAV1ApplianceGetActiveAfterFixedIPUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {
        			"k1": "v1"
        		},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "zone1-groupb",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "COMPLETE",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID3",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.3.50",
        					"subnet_id": "670c1b56-4b9c-4bde-b7ac-1f2e09391d81"
        				}, {
        					"ip_address": "192.168.3.60",
        					"subnet_id": "670c1b56-4b9c-4bde-b7ac-1f2e09391d81"
        				}]
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID4",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID2",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.2.101",
        					"subnet_id": "4be82753-9dc5-4065-a4c0-46abe02bb93a"
        				}]
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Updated
counter:
    min: 4
`

var testMockVNAV1ApplianceProcessingAfterAllowedAddressPairUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {
        			"k1": "v1"
        		},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "zone1-groupb",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "PROCESSING",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [
                            {
                                "ip_address": "192.168.1.200",
                                "mac_address": "aa:bb:cc:dd:ee:f1",
                                "type": "vrrp",
                                "vrid": 123
                            },
                            {
                                "ip_address": "192.168.1.201",
                                "mac_address": "aa:bb:cc:dd:ee:f2",
                                "type": "",
                                "vrid": null
                            }
                        ],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Updated
counter:
    max: 3
`

var testMockVNAV1ApplianceGetActiveAfterAllowedAddressPairUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {
        			"k1": "v1"
        		},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "zone1-groupb",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "6589b37a-cf82-4918-96fe-255683f78e76",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "COMPLETE",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [
                            {
                                "ip_address": "192.168.1.200",
                                "mac_address": "aa:bb:cc:dd:ee:f1",
                                "type": "vrrp",
                                "vrid": 123
                            },
                            {
                                "ip_address": "192.168.1.201",
                                "mac_address": "aa:bb:cc:dd:ee:f2",
                                "type": "",
                                "vrid": null
                            }
                        ],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Updated
counter:
    min: 4
`
