package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"

	"github.com/nttcom/eclcloud/ecl/baremetal/v2/servers"
)

func TestMockedBaremetalV2ServerBasic(t *testing.T) {
	if OS_REGION_NAME != "RegionOne" {
		t.Skipf("skip this test in %s region", OS_REGION_NAME)
	}

	var server servers.Server

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint())
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "image", "/v2/01234567890123456789abcdefabcdef/images/detail", testMockImageV2ImageList)
	mc.Register(t, "flavor", "/v2/01234567890123456789abcdefabcdef/flavors/detail", testMockBaremetalV2FlavorList)
	mc.Register(t, "baremetal_server", "/v2/01234567890123456789abcdefabcdef/servers", testMockBaremetalV2ServerCreate)
	mc.Register(t, "baremetal_server", "/v2/01234567890123456789abcdefabcdef/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931", testMockBaremetalV2ServerGetAfterCreate)
	mc.Register(t, "baremetal_server", "/v2/01234567890123456789abcdefabcdef/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931", testMockBaremetalV2ServerDelete)
	mc.Register(t, "baremetal_server", "/v2/01234567890123456789abcdefabcdef/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931", testMockBaremetalV2ServerGetAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBaremetalV2ServerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockBaremetalV2ServerBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaremetalV2ServerExists("ecl_baremetal_server_v2.server_1", &server),
					resource.TestCheckResourceAttr(
						"ecl_baremetal_server_v2.server_1", "name", "server1"),
				),
			},
		},
	})
}

const testMockBaremetalV2ServerBasic = `
resource "ecl_baremetal_server_v2" "server_1" {
    name = "server1"
    image_id = "image_id"
    flavor_id = "flavor_id"
    image_name = "image_name"
    flavor_name = "flavor_name"
    user_data = "user_data"
    availability_zone = "zone1"
    key_pair = "key_pair"
    admin_pass = "aabbccddeeff"
    metadata = {
        k1 = "v1"
        k2 = "v2"
    }
    networks {
        uuid = "6a9a64f6-45e7-46b2-b0e8-0a850896ff55"
        fixed_ip = "192.168.0.1"
        plane = "data"
    }
    raid_arrays {
        primary_storage = true
        partitions {
            lvm = true
            partition_label = "primary-part1"
        }
        partitions {
            lvm = false
            size = "100G"
            partition_label = "var"
        }
    }
    raid_arrays {
        raid_card_hardware_id = "raid_card_uuid"
        disk_hardware_ids = ["disk1_uuid", "disk2_uuid", "disk3_uuid", "disk4_uuid"]
        partitions {
            lvm = true
            partition_label = "secondary-part1"
        }
        laid_level = 10
    }
    lvm_volume_groups {
        vg_label = "VG_root"
        physical_volume_partition_labels = ["primary-part1", "secondary-part1"]
        logical_volumes {
            lv_label = "LV_root"
            size = "300G"
        }
        logical_volumes {
            lv_label = "LV_swap"
            size = "2G"
        }
    }
    filesystems {
        label = "LV_root"
        mount_point =  "/"
        fs_type = "xfs"
    }
    filesystems {
        label = "var"
        mount_point = "/var"
        fs_type = "xfs"
    }
    filesystems {
        label = "LV_swap"
        fs_type = "swap"
    }
    personality {
        path = "/home/big/banner.txt"
        contents = "ZWNobyAiS3VtYSBQZXJzb25hbGl0eSIgPj4gL2hvbWUvYmlnL3BlcnNvbmFsaXR5"
    }
}
`

var testMockBaremetalV2ServerCreate = `
request:
    method: POST
response:
    code: 200
    body: >
        {
            "server": {
                "id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
                "links": [
                    {
                        "href": "http://openstack.example.com/v2/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
                        "rel": "self"
                    },
                    {
                        "href": "http://openstack.example.com/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
                        "rel": "bookmark"
                    }
                ],
                "adminPass": "aabbccddeeff"
            }
        }
newStatus: Created
`

var testMockBaremetalV2ServerGetAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "server": {
            "OS-EXT-STS:power_state": "RUNNING",
            "OS-EXT-STS:task_state": "None",
            "OS-EXT-STS:vm_state": "ACTIVE",
            "OS-EXT-AZ:availability_zone": "zone1-groupa",
            "created": "2012-09-07T16:56:37Z",
            "flavor": {
                "id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
                "links": [{
                    "href": "http://openstack.example.com/openstack/flavors/05184ba3-00ba-4fbc-b7a2-03b62b884931",
                    "rel": "bookmark"
                }]
            },
            "hostId": "16d193736a5cfdb60c697ca27ad071d6126fa13baeb670fc9d10645e",
            "id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
            "image": {
                "id": "70a599e0-31e7-49b7-b260-868f441e862b",
                "links": [{
                    "href": "http://openstack.example.com/openstack/images/70a599e0-31e7-49b7-b260-868f441e862b",
                    "rel": "bookmark"
                }]
            },
            "links": [{
                "href": "http://openstack.example.com/v2/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
                "rel": "self"
            },
            {
                "href": "http://openstack.example.com/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
                "rel": "bookmark"
            }],
            "metadata": {
                "My Server Name": "Apache1"
            },
            "name": "server1",
            "progress": 0,
            "status": "ACTIVE",
            "tenant_id": "openstack",
            "updated": "2012-09-07T16:56:37Z",
            "user_id": "fake",
            "raid_arrays": [
                {
                "primary_storage": true,
                "raid_card_hardware_id": "raid_card_uuid",
                "disk_hardware_ids": [
                    "disk0_uuid",
                    "disk1_uuid",
                    "disk2_uuid",
                    "disk3_uuid"
                ],
                "partitions": [
                    {
                    "lvm": true,
                    "partition_label": "primary-part1"
                    },
                    {
                    "lvm": false,
                    "size": 100,
                    "partition_label": "var"
                    }
                ]
                },
                {
                "primary_storage": false,
                "raid_card_hardware_id": "raid_card_uuid",
                "internal_disk_ids": [
                    "disk4_uuid",
                    "disk5_uuid",
                    "disk6_uuid",
                    "disk7_uuid"
                ],
                "raid_level": 10,
                "partitions": [
                    {
                    "lvm": true,
                    "partition_label": "secondary-part1"
                    }
                ]
                }
            ],
            "lvm_volume_groups": [
                {
                "vg_label": "VG_root",
                "physical_volume_partition_labels": [
                    "primary-part1",
                    "secondary-part1"
                ],
                "logical_volumes": [
                    {
                    "lv_label": "LV_root"
                    },
                    {
                    "size": 2,
                    "lv_label": "LV_swap"
                    }
                ]
                }
            ],
            "filesystems": [
                {
                "label": "LV_root",
                "mount_point": "/",
                "fs_type": "xfs"
                },
                {
                "label": "var",
                "mount_point": "/var",
                "fs_type": "xfs"
                },
                {
                "label": "LV_swap",
                "fs_type": "swap"
                }
            ],
            "nic_physical_ports": [
                {
                "id": "39285bf9-12fb-4064-b98b-a552efc51cfc",
                "mac_addr": "0a:31:c1:d5:6d:9c",
                "network_physical_port_id": "38268d94-584a-4f14-96ff-732a68aa7301",
                "plane": "data",
                "attached_ports": [
                    {
                    "port_id": "61b7da1e-9571-4d63-b779-e003a56b8105",
                    "network_id": "9aa93722-1ec4-4912-b813-b975c21460a5",
                    "fixed_ips": [
                        {
                        "subnet_id": "0419bbde-2b82-4107-9d8a-6bba76e364af",
                        "ip_address": "192.168.10.2"
                        }
                    ]
                    }
                ],
                "hardware_id": "c1e1546d-3063-46d0-8895-c6350eb691ff"
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
                "etc": true
            },
            "media_attachments": [
                {
                "image": {
                    "id": "3339fd5f-ec06-4ef8-9337-c1c70218a748",
                    "links": [
                    {
                        "href": "http://openstack.example.com/openstack/images/3339fd5f-ec06-4ef8-9337-c1c70218a748",
                        "rel": "bookmark"
                    }
                    ]
                }
                }
            ]
            }
        }
expectedStatus:
    - Created
`

var testMockBaremetalV2ServerDelete = `
request:
    method: DELETE
response:
    code: 202
newStatus: Deleted
`

var testMockBaremetalV2ServerGetAfterDelete = `
request:
    method: GET
response:
    code: 404
expectedStatus:
    - Deleted
`

var testMockImageV2ImageList = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "images": [
                {
                    "status": "<image_status>",
                    "name": "image_name",
                    "tags": [
                        "<tag>"
                    ],
                    "container_format": "<container_format>",
                    "created_at": "<created_time>",
                    "disk_format": "<disk_format>",
                    "locations": [
                        {
                            "url": "<location_url>",
                            "metadata": "<metadata>"
                        }
                    ],
                    "direct_url": "<direct_url>",
                    "<extra_key>": "<extra_value>",
                    "updated_at": "<updated_time>",
                    "visibility": "<visibility>",
                    "self": "<self>",
                    "min_disk": "<minimum_disk_size>",
                    "protected": "<protected_flag>",
                    "id": "image_id",
                    "file": "<file>",
                    "checksum": "<checksum>",
                    "owner": "<owner>",
                    "size": "<size>",
                    "min_ram": "<minimum_ram_size>",
                    "schema": "<image_schema>"
                }
            ],
            "schema": "<images_schema>",
            "first": "<first>"
        }
`

var testMockBaremetalV2FlavorList = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "flavors": [
                {
                    "disk": 500,
                    "id": "flavor_id",
                    "links": [
                        {
                            "href": "<href>",
                            "rel": "<rel>"
                        },
                        {
                            "href": "<href>",
                            "rel": "<rel>"
                        }
                    ],
                    "name": "flavor_name",
                    "ram": 1024,
                    "vcpus": 4
                },
                {
                    "disk": 100,
                    "id": "flavor_id_2",
                    "links": [
                        {
                            "href": "<href>",
                            "rel": "<rel>"
                        },
                        {
                            "href": "<href>",
                            "rel": "<rel>"
                        }
                    ],
                    "name": "flavor_name_2",
                    "ram": 512,
                    "vcpus": 1
                }
            ]
        }
`
