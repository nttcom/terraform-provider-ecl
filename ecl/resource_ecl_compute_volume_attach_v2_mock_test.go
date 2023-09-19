package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"

	"github.com/nttcom/eclcloud/v4/ecl/computevolume/v2/volumes"
)

func TestMockedVolumeV2Attach_basic(t *testing.T) {
	var va volumes.Attachment

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)
	mc.Register(t, "attach", "/v2/01234567890123456789abcdefabcdef/servers/9fd11843-2eda-4d46-9a95-0631ad65ad8e/os-volume_attachments", testMockAttachV2Create)
	mc.Register(t, "attach", "/v2/01234567890123456789abcdefabcdef/servers/9fd11843-2eda-4d46-9a95-0631ad65ad8e/os-volume_attachments", testMockAttachV2ListAfterCreate)
	mc.Register(t, "attach", "/v2/01234567890123456789abcdefabcdef/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63", testMockVolumeV2ShowStatusInuse)
	mc.Register(t, "attach", "/v2/01234567890123456789abcdefabcdef/servers/9fd11843-2eda-4d46-9a95-0631ad65ad8e/os-volume_attachments/5be9b6b8-2713-40a7-8c40-0737717e7b63", testMockAttachV2Delete)
	mc.Register(t, "attach", "/v2/01234567890123456789abcdefabcdef/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63", testMockVolumeV2ShowStatusDetaching)
	mc.Register(t, "attach", "/v2/01234567890123456789abcdefabcdef/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63", testMockVolumeV2ShowStatusAvailable)
	mc.Register(t, "attach", "/v2/01234567890123456789abcdefabcdef/servers/9fd11843-2eda-4d46-9a95-0631ad65ad8e/os-volume_attachments", testMockAttachV2ListAfterDelete)
	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckRequiredEnvVars(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2AttachDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockVolumeV2AttachBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeVolumeV2AttachExists("ecl_compute_volume_attach_v2.va_1", &va),
				),
			},
		},
	})
}

var testMockVolumeV2AttachBasic = `
resource "ecl_compute_volume_attach_v2" "va_1" {
    volume_id = "5be9b6b8-2713-40a7-8c40-0737717e7b63"
    server_id = "9fd11843-2eda-4d46-9a95-0631ad65ad8e"
    device = "/dev/vdb"
}
`

var testMockAttachV2Create = `
request:
    method: POST
response:
    code: 200
    body: >
        {
            "volumeAttachment": {
                "device": "/dev/vdb",
                "id": "5be9b6b8-2713-40a7-8c40-0737717e7b63",
                "serverId": "9fd11843-2eda-4d46-9a95-0631ad65ad8e",
                "volumeId": "5be9b6b8-2713-40a7-8c40-0737717e7b63"
            }
        }
newStatus: Created
`

var testMockAttachV2ListAfterCreate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "volumeAttachments": [
                {
                    "device": "/dev/vdb",
                    "id": "5be9b6b8-2713-40a7-8c40-0737717e7b63",
                    "serverId": "9fd11843-2eda-4d46-9a95-0631ad65ad8e",
                    "volumeId": "5be9b6b8-2713-40a7-8c40-0737717e7b63"
                }
            ]
        }
expectedStatus:
    - Created
`

var testMockAttachV2ListAfterDelete = `
request:
    method: GET
response:
    code: 404
    body: >
        {
            "itemNotFound": {
                "code": 404,
                "message": "Instance 590edec6-2b94-4a3b-b9d4-a5fe4e9f2733 could not be found."
            }
        }
expectedStatus:
    - Available
`

var testMockAttachV2Delete = `
request:
    method: DELETE
response:
    code: 202
newStatus: Deleted
`

var testMockVolumeV2ShowStatusInuse = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "volume": {
                "attachments": [
                    {
                        "attached_at": "2022-07-12T05:36:04.000000",
                        "attachment_id": "6517902a-d593-4d2f-8e71-12eeb40bbb96",
                        "device": "/dev/vdb",
                        "host_name": null,
                        "id": "5be9b6b8-2713-40a7-8c40-0737717e7b63",
                        "server_id": "9fd11843-2eda-4d46-9a95-0631ad65ad8e",
                        "volume_id": "5be9b6b8-2713-40a7-8c40-0737717e7b63"
                    }
                ],
                "availability_zone": "zone1_groupa",
                "bootable": "false",
                "consistencygroup_id": null,
                "created_at": "2022-07-12T05:35:31.000000",
                "description": null,
                "encrypted": false,
                "id": "5be9b6b8-2713-40a7-8c40-0737717e7b63",
                "links": [
                    {
                        "href": "https://cinder-lab3ec-ecl.lab.api.ntt.com/v2/e4bbd9450deb4745986c7382b36ae50e/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63",
                        "rel": "self"
                    },
                    {
                        "href": "https://cinder-lab3ec-ecl.lab.api.ntt.com/e4bbd9450deb4745986c7382b36ae50e/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63",
                        "rel": "bookmark"
                    }
                ],
                "metadata": {
                    "attached_mode": "rw",
                    "readonly": "False"
                },
                "multiattach": false,
                "name": "volume_1",
                "os-vol-tenant-attr:tenant_id": "e4bbd9450deb4745986c7382b36ae50e",
                "replication_status": "disabled",
                "size": 15,
                "snapshot_id": null,
                "source_volid": null,
                "status": "in-use",
                "updated_at": "2022-07-12T05:36:05.000000",
                "user_id": "f98266e4910042dd87920cc44b22a477",
                "volume_type": "nfsdriver"
            }
        }
expectedStatus:
    - Created
`

var testMockVolumeV2ShowStatusDetaching = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "volume": {
                "attachments": [
                    {
                        "attached_at": "2022-07-12T05:36:04.000000",
                        "attachment_id": "6517902a-d593-4d2f-8e71-12eeb40bbb96",
                        "device": "/dev/vdb",
                        "host_name": null,
                        "id": "5be9b6b8-2713-40a7-8c40-0737717e7b63",
                        "server_id": "9fd11843-2eda-4d46-9a95-0631ad65ad8e",
                        "volume_id": "5be9b6b8-2713-40a7-8c40-0737717e7b63"
                    }
                ],
                "availability_zone": "zone1_groupa",
                "bootable": "false",
                "consistencygroup_id": null,
                "created_at": "2022-07-12T05:35:31.000000",
                "description": null,
                "encrypted": false,
                "id": "5be9b6b8-2713-40a7-8c40-0737717e7b63",
                "links": [
                    {
                        "href": "https://cinder-lab3ec-ecl.lab.api.ntt.com/v2/e4bbd9450deb4745986c7382b36ae50e/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63",
                        "rel": "self"
                    },
                    {
                        "href": "https://cinder-lab3ec-ecl.lab.api.ntt.com/e4bbd9450deb4745986c7382b36ae50e/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63",
                        "rel": "bookmark"
                    }
                ],
                "metadata": {
                    "attached_mode": "rw",
                    "readonly": "False"
                },
                "multiattach": false,
                "name": "volume_1",
                "os-vol-tenant-attr:tenant_id": "e4bbd9450deb4745986c7382b36ae50e",
                "replication_status": "disabled",
                "size": 15,
                "snapshot_id": null,
                "source_volid": null,
                "status": "detaching",
                "updated_at": "2022-07-12T05:36:05.000000",
                "user_id": "f98266e4910042dd87920cc44b22a477",
                "volume_type": "nfsdriver"
            }
        }
newStatus: Available
expectedStatus:
    - Deleted
`

var testMockVolumeV2ShowStatusAvailable = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "volume": {
                "attachments": [
                ],
                "availability_zone": "zone1_groupa",
                "bootable": "false",
                "consistencygroup_id": null,
                "created_at": "2022-07-12T05:35:31.000000",
                "description": null,
                "encrypted": false,
                "id": "5be9b6b8-2713-40a7-8c40-0737717e7b63",
                "links": [
                    {
                        "href": "https://cinder-lab3ec-ecl.lab.api.ntt.com/v2/e4bbd9450deb4745986c7382b36ae50e/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63",
                        "rel": "self"
                    },
                    {
                        "href": "https://cinder-lab3ec-ecl.lab.api.ntt.com/e4bbd9450deb4745986c7382b36ae50e/volumes/5be9b6b8-2713-40a7-8c40-0737717e7b63",
                        "rel": "bookmark"
                    }
                ],
                "metadata": {
                    "readonly": "False"
                },
                "multiattach": false,
                "name": "volume_1",
                "os-vol-tenant-attr:tenant_id": "e4bbd9450deb4745986c7382b36ae50e",
                "replication_status": "disabled",
                "size": 15,
                "snapshot_id": null,
                "source_volid": null,
                "status": "available",
                "updated_at": "2022-07-12T05:36:25.000000",
                "user_id": "f98266e4910042dd87920cc44b22a477",
                "volume_type": "nfsdriver"
            }
        }
expectedStatus:
    - Available
`
