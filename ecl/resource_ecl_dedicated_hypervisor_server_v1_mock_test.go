package ecl

import (
	"fmt"
	"os"
	"testing"

	"github.com/nttcom/eclcloud/ecl/dedicated_hypervisor/v1/servers"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedDedicatedHypervisorV1ServerBasic(t *testing.T) {
	if region := os.Getenv("OS_REGION_NAME"); region != "RegionOne" {
		t.Skipf("skip this test in %s region", region)
	}

	var server servers.Server

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "dedicated_hypervisor", "/v1.0/1bc271e7a8af4d988ff91612f5b122f8/servers", testMockDedicatedHypervisorV1ServerCreate)
	mc.Register(t, "dedicated_hypervisor", "/v1.0/1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5", testMockDedicatedHypervisorV1ServerGetAfterCreate)
	mc.Register(t, "dedicated_hypervisor", "/v1.0/1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5", testMockDedicatedHypervisorV1ServerDelete)
	mc.Register(t, "dedicated_hypervisor", "/v1.0/1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5", testMockDedicatedHypervisorV1ServerGetAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDedicatedHypervisorV1ServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testMockDedicatedHypervisorV1ServerBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDedicatedHypervisorV1ServerExists("ecl_dedicated_hypervisor_server_v1.server_1", &server),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "name", "server1"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "description", "ESXi Dedicated Hypervisor"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.0.uuid", "94055904-6b2c-4839-a14a-c61c93a8bc48"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.0.fixed_ip", "192.168.1.10"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.0.plane", "data"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.0.segmentation_id", "4"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.1.uuid", "94055904-6b2c-4839-a14a-c61c93a8bc48"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.1.fixed_ip", "192.168.1.11"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.1.plane", "data"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.1.segmentation_id", "4"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "admin_pass", "aabbccddeeff"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "image_ref", "dfd25820-b368-4012-997b-29a6d0cf8518"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "flavor_ref", "a830b61c-3155-4a61-b7ed-c450862845e6"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "availability_zone", "groupb"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "metadata.k1", "v1"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "metadata.k2", "v2"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "baremetal_server_id", "24ebe7b8-ecfb-4d9f-a66b-c0120534fc90"),
				),
			},
		},
	})
}

const testMockDedicatedHypervisorV1ServerBasic = `
resource "ecl_dedicated_hypervisor_server_v1" "server_1" {
    name = "server1"
    description = "ESXi Dedicated Hypervisor"
    networks {
        uuid = "94055904-6b2c-4839-a14a-c61c93a8bc48"
        fixed_ip = "192.168.1.10"
        plane = "data"
        segmentation_id = 4
    }
    networks {
        uuid = "94055904-6b2c-4839-a14a-c61c93a8bc48"
        fixed_ip = "192.168.1.11"
        plane = "data"
        segmentation_id = 4
    }
    admin_pass = "aabbccddeeff"
    image_ref = "dfd25820-b368-4012-997b-29a6d0cf8518"
    flavor_ref = "a830b61c-3155-4a61-b7ed-c450862845e6"
    availability_zone = "groupb"
    metadata = {
        k1 = "v1"
        k2 = "v2"
    }
}
`

var testMockDedicatedHypervisorV1ServerCreate = `
request:
    method: POST
response:
    code: 200
    body: >
        {
            "server": {
                "id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
                "links": [
                    {
                        "href": "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//v2/1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5",
                        "rel": "self"
                    },
                    {
                        "href": "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5",
                        "rel": "bookmark"
                    }
                ],
            "adminPass": "aabbccddeeff"
            }
        }
newStatus: Created
`

var testMockDedicatedHypervisorV1ServerGetAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "server": {
                "id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
                "name": "server1",
                "imageRef": "dfd25820-b368-4012-997b-29a6d0cf8518",
                "description": "ESXi Dedicated Hypervisor",
                "status": "ACTIVE",
                "hypervisor_type": "vsphere_esxi",
                "baremetal_server": {
                    "OS-EXT-STS:power_state": "RUNNING",
                    "OS-EXT-STS:task_state": "None",
                    "OS-EXT-STS:vm_state": "ACTIVE",
                    "OS-EXT-AZ:availability_zone": "groupb",
                    "progress": 100,
                    "created": "2019-10-10T04:11:41Z",
                    "flavor": {
                        "id": "a830b61c-3155-4a61-b7ed-c450862845e6",
                        "links": [
                            {
                                "href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/a830b61c-3155-4a61-b7ed-c450862845e6",
                                "rel": "bookmark"
                            }
                        ]
                    },
                    "id": "24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
                    "image": {
                        "id": "112a26a0-ff25-4513-afe1-407e41b0a48b",
                        "links": [
                            {
                                "href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/112a26a0-ff25-4513-afe1-407e41b0a48b",
                                "rel": "bookmark"
                            }
                        ]
                    },
                    "links": [
                        {
                            "href": "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
                            "rel": "self"
                        },
                        {
                            "href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
                            "rel": "bookmark"
                        }
                    ],
                    "metadata": {
                        "k1": "v1",
                        "k2": "v2"
                    },
                    "name": "test",
                    "status": "ACTIVE",
                    "tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
                    "updated": "2019-10-10T04:14:08Z",
                    "user_id": "55891ce6a3cb4bb0833514667d67288c",
                    "raid_arrays": [
                        {
                            "primary_storage": true,
                            "partitions": null,
                            "raid_card_hardware_id": "bdfb75d1-194d-426d-b288-f588dfa5ac49",
                            "disk_hardware_ids": [
                                "76649053-863e-4533-86e3-f194a79485a6",
                                "a25827e3-67da-47be-ba96-849ab4685a1d"
                            ]
                        }
                    ],
                    "lvm_volume_groups": null,
                    "filesystems": null,
                    "nic_physical_ports": [
                        {
                            "id": "a2f63380-6c77-4cd5-8868-e3556ffd35ce",
                            "mac_addr": "48:DF:37:90:B4:58",
                            "plane": "DATA",
                            "network_physical_port_id": "d8e40a51-f1e2-4681-8953-9fe1e9992c42",
                            "hardware_id": "be2d30d6-f891-4200-b827-95f229fb8c6b",
                            "attached_ports": [
                                {
                                    "network_id": "94055904-6b2c-4839-a14a-c61c93a8bc48",
                                    "port_id": "30fc1c27-fb5f-4955-94d0-a56cd28d09e8",
                                    "fixed_ips": [
                                        {
                                            "subnet_id": "acd41997-5ebb-4ff2-8cd2-22cae6cf2883",
                                            "ip_address": "192.168.1.10"
                                        }
                                    ]
                                },
                                {
                                    "network_id": "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
                                    "port_id": "aa6c61f4-db8a-44c7-a91c-7e636dac1dc6",
                                    "fixed_ips": [
                                        {
                                            "subnet_id": "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
                                            "ip_address": "169.254.0.9"
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "id": "b01dfdb0-f247-47d8-8224-c257aa3265e9",
                            "mac_addr": "48:DF:37:90:B4:50",
                            "plane": "STORAGE",
                            "network_physical_port_id": "00dfea92-5c5b-4860-aa05-efef6c2bb2af",
                            "hardware_id": "be2d30d6-f891-4200-b827-95f229fb8c6b",
                            "attached_ports": []
                        },
                        {
                            "id": "f4355e8e-39fc-48bd-a283-a2dbef8a2e32",
                            "mac_addr": "48:DF:37:82:B0:A0",
                            "plane": "STORAGE",
                            "network_physical_port_id": "cf798cc0-c869-45d5-a5a7-bcc578a300b0",
                            "hardware_id": "84c74a86-7045-4284-80f9-0e7aff5d27ad",
                            "attached_ports": []
                        },
                        {
                            "id": "5ef177fd-888c-4fae-9925-a8920beb07cb",
                            "mac_addr": "48:DF:37:82:B0:A8",
                            "plane": "DATA",
                            "network_physical_port_id": "2bbbb516-c75a-42b2-8a46-9cb5f26c219e",
                            "hardware_id": "84c74a86-7045-4284-80f9-0e7aff5d27ad",
                            "attached_ports": [
                                {
                                    "network_id": "94055904-6b2c-4839-a14a-c61c93a8bc48",
                                    "port_id": "4e329a01-2cf4-4028-9259-03b7aa145cb6",
                                    "fixed_ips": [
                                        {
                                            "subnet_id": "acd41997-5ebb-4ff2-8cd2-22cae6cf2883",
                                            "ip_address": "192.168.1.11"
                                        }
                                    ]
                                },
                                {
                                    "network_id": "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
                                    "port_id": "a256b4a1-3ae3-4102-a14e-987ae1610f97",
                                    "fixed_ips": [
                                        {
                                            "subnet_id": "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
                                            "ip_address": "169.254.0.10"
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "chassis-status": {
                        "chassis-power": true,
                        "power-supply": true,
                        "cpu": true,
                        "memory": true,
                        "fan": true,
                        "disk": 0,
                        "nic": true,
                        "system-board": true,
                        "etc": true,
                        "console": true
                    },
                    "media_attachments": [],
                    "managed_by_service": "dedicated-hypervisor",
                    "managed_service_resource_id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5"
                }
            }
        }
expectedStatus:
    - Created
`

var testMockDedicatedHypervisorV1ServerDelete = `
request:
    method: DELETE
response:
    code: 202
newStatus: Deleted
`

var testMockDedicatedHypervisorV1ServerGetAfterDelete = `
request:
    method: GET
response:
    code: 404
expectedStatus:
    - Deleted
`
