package ecl

import (
	"fmt"
	"testing"

	"github.com/nttcom/eclcloud/ecl/provider_connectivity/v2/tenant_connections"

	"github.com/terraform-providers/terraform-provider-ecl/ecl/testhelper/mock"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestMockedAccProviderConnectivityV2TenantConnection_basic(t *testing.T) {
	var connection tenant_connections.TenantConnection

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections", testMockProviderConnectivityV2TenantConnectionCreateServer)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionGetAfterCreateServer)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionUpdate)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionGetAfterUpdate)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionDelete)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionGetAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testMockAccProviderConnectivityV2TenantConnectionServerConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProviderConnectivityV2TenantConnectionExists("ecl_provider_connectivity_tenant_connection_v2.connection_1", &connection),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "id", "06969a7c-9fc0-11ea-b509-525403060400"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_connection_request_id", "fb1aae9a-9fbf-11ea-9e55-525403060300"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "test_name1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "test_desc1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.test_tags1", "test1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_id", "2c76532f048849aab41c1bff2ec8b996"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_id_other", "7af0c902bd51424f8b2c85f5320ab181"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "network_id", "69a7f763-adc7-4587-a43c-334942322f35"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_type", "ECL::Compute::Server"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_id", "46f97549-b3be-4a8a-a000-84f76d17f355"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_interface_id", ""),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "port_id", "bfc5e3f5-40cb-4f6a-a4a5-dc6e8a9e2f96"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "status", "active"),
				),
			},
			{
				Config: testMockAccProviderConnectivityV2TenantConnectionSeverUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "updated_name"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "updated_desc"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.k2", "v2"),
				),
			},
		},
	})
}

var testMockAccProviderConnectivityV2TenantConnectionServerConfig = `
resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "fb1aae9a-9fbf-11ea-9e55-525403060300"
	device_type = "ECL::Compute::Server"
	device_id = "46f97549-b3be-4a8a-a000-84f76d17f355"
	attachment_opts_compute {
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
}
`

var testMockAccProviderConnectivityV2TenantConnectionSeverUpdateConfig = `
resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	name = "updated_name"
	description = "updated_desc"
	tags = {
		"k2" = "v2"
	}
	tenant_connection_request_id = "fb1aae9a-9fbf-11ea-9e55-525403060300"
	device_type = "ECL::Compute::Server"
	device_id = "46f97549-b3be-4a8a-a000-84f76d17f355"
	attachment_opts_compute {
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
}
`

var testMockProviderConnectivityV2TenantConnectionCreateServer = `
request:
    method: POST
response:
    code: 200
    body: >
        {
            "tenant_connection": {
                "description": "test_desc1",
                "description_other": "",
                "device_id": "46f97549-b3be-4a8a-a000-84f76d17f355",
                "device_interface_id": "",
                "device_type": "ECL::Compute::Server",
                "id": "06969a7c-9fc0-11ea-b509-525403060400",
                "name": "test_name1",
                "name_other": "",
                "network_id": "69a7f763-adc7-4587-a43c-334942322f35",
                "port_id": "",
                "status": "creating",
                "tags": {
                    "test_tags1": "test1"
                },
                "tags_other": {},
                "tenant_connection_request_id": "fb1aae9a-9fbf-11ea-9e55-525403060300",
                "tenant_id": "2c76532f048849aab41c1bff2ec8b996",
                "tenant_id_other": "7af0c902bd51424f8b2c85f5320ab181"
            }
        }
newStatus: Created
`

var testMockProviderConnectivityV2TenantConnectionGetAfterCreateServer = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "tenant_connection": {
                "description": "test_desc1",
                "description_other": "",
                "device_id": "46f97549-b3be-4a8a-a000-84f76d17f355",
                "device_interface_id": "",
                "device_type": "ECL::Compute::Server",
                "id": "06969a7c-9fc0-11ea-b509-525403060400",
                "name": "test_name1",
                "name_other": "",
                "network_id": "69a7f763-adc7-4587-a43c-334942322f35",
                "port_id": "bfc5e3f5-40cb-4f6a-a4a5-dc6e8a9e2f96",
                "status": "active",
                "tags": {
                    "test_tags1": "test1"
                },
                "tags_other": {},
                "tenant_connection_request_id": "fb1aae9a-9fbf-11ea-9e55-525403060300",
                "tenant_id": "2c76532f048849aab41c1bff2ec8b996",
                "tenant_id_other": "7af0c902bd51424f8b2c85f5320ab181"
            }
        }
expectedStatus:
    - Created
`

var testMockProviderConnectivityV2TenantConnectionUpdate = `
request:
    method: PUT
response:
    code: 200
    body: >
        {
            "tenant_connection": {
                "description": "test_desc1",
                "description_other": "",
                "device_id": "46f97549-b3be-4a8a-a000-84f76d17f355",
                "device_interface_id": "",
                "device_type": "ECL::Compute::Server",
                "id": "06969a7c-9fc0-11ea-b509-525403060400",
                "name": "test_name1",
                "name_other": "",
                "network_id": "69a7f763-adc7-4587-a43c-334942322f35",
                "port_id": "bfc5e3f5-40cb-4f6a-a4a5-dc6e8a9e2f96",
                "status": "active",
                "tags": {
                    "test_tags1": "test1"
                },
                "tags_other": {},
                "tenant_connection_request_id": "fb1aae9a-9fbf-11ea-9e55-525403060300",
                "tenant_id": "2c76532f048849aab41c1bff2ec8b996",
                "tenant_id_other": "7af0c902bd51424f8b2c85f5320ab181"
            }
        }
newStatus: Updated
`

var testMockProviderConnectivityV2TenantConnectionGetAfterUpdate = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "tenant_connection": {
                "description": "updated_desc",
                "description_other": "",
                "device_id": "46f97549-b3be-4a8a-a000-84f76d17f355",
                "device_interface_id": "",
                "device_type": "ECL::Compute::Server",
                "id": "06969a7c-9fc0-11ea-b509-525403060400",
                "name": "updated_name",
                "name_other": "",
                "network_id": "69a7f763-adc7-4587-a43c-334942322f35",
                "port_id": "bfc5e3f5-40cb-4f6a-a4a5-dc6e8a9e2f96",
                "status": "active",
                "tags": {
                    "k2": "v2"
                },
                "tags_other": {},
                "tenant_connection_request_id": "fb1aae9a-9fbf-11ea-9e55-525403060300",
                "tenant_id": "2c76532f048849aab41c1bff2ec8b996",
                "tenant_id_other": "7af0c902bd51424f8b2c85f5320ab181"
            }
        }
expectedStatus:
    - Updated
`

var testMockProviderConnectivityV2TenantConnectionDelete = `
request:
    method: DELETE
response:
    code: 204
newStatus: Deleted
`

var testMockProviderConnectivityV2TenantConnectionGetAfterDelete = `
request:
    method: GET
response:
    code: 404
expectedStatus:
    - Deleted
`

func TestMockedAccProviderConnectivityV2TenantConnection_baremetal(t *testing.T) {
	var connection tenant_connections.TenantConnection

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections", testMockProviderConnectivityV2TenantConnectionCreateBaremetal)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/76ea7594-9ff9-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionGetAfterCreateBaremetal)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/76ea7594-9ff9-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionDelete)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/76ea7594-9ff9-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionGetAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testMockAccProviderConnectivityV2TenantConnectionBaremetalConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProviderConnectivityV2TenantConnectionExists("ecl_provider_connectivity_tenant_connection_v2.connection_1", &connection),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "id", "76ea7594-9ff9-11ea-9e55-525403060300"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_connection_request_id", "276a0de2-9ff7-11ea-9e55-525403060300"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "test_name1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "test_desc1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.test_tags1", "test1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_id", "2c76532f048849aab41c1bff2ec8b996"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_id_other", "7af0c902bd51424f8b2c85f5320ab181"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "network_id", "1b0d0417-9757-43b2-9885-a3625162be08"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_type", "ECL::Baremetal::Server"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_id", "b9cc6ca1-7b17-47b7-bb75-705a806204a4"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_interface_id", "1a76a0f6-ac2d-47e0-97ef-7222af0077b1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "port_id", "16253ba3-93eb-4629-a998-2e1d6ed978ed"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "status", "active"),
				),
			},
		},
	})
}

var testMockAccProviderConnectivityV2TenantConnectionBaremetalConfig = `
resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "276a0de2-9ff7-11ea-9e55-525403060300"
	device_type = "ECL::Baremetal::Server"
	device_id = "b9cc6ca1-7b17-47b7-bb75-705a806204a4"
	device_interface_id = "1a76a0f6-ac2d-47e0-97ef-7222af0077b1"
	attachment_opts_baremetal {
		segmentation_type = "flat"
		segmentation_id = 10
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
}
`

var testMockProviderConnectivityV2TenantConnectionCreateBaremetal = `
request:
    method: POST
response:
    code: 200
    body: >
        {
            "tenant_connection": {
                "description": "test_desc1",
                "description_other": "",
                "device_id": "b9cc6ca1-7b17-47b7-bb75-705a806204a4",
                "device_interface_id": "1a76a0f6-ac2d-47e0-97ef-7222af0077b1",
                "device_type": "ECL::Baremetal::Server",
                "id": "76ea7594-9ff9-11ea-9e55-525403060300",
                "name": "test_name1",
                "name_other": "",
                "network_id": "1b0d0417-9757-43b2-9885-a3625162be08",
                "port_id": "",
                "status": "creating",
                "tags": {
                    "test_tags1": "test1"
                },
                "tags_other": "{}",
                "tenant_connection_request_id": "276a0de2-9ff7-11ea-9e55-525403060300",
                "tenant_id": "2c76532f048849aab41c1bff2ec8b996",
                "tenant_id_other": "7af0c902bd51424f8b2c85f5320ab181"
            }
        }
newStatus: Created
`

var testMockProviderConnectivityV2TenantConnectionGetAfterCreateBaremetal = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "tenant_connection": {
                "description": "test_desc1",
                "description_other": "",
                "device_id": "b9cc6ca1-7b17-47b7-bb75-705a806204a4",
                "device_interface_id": "1a76a0f6-ac2d-47e0-97ef-7222af0077b1",
                "device_type": "ECL::Baremetal::Server",
                "id": "76ea7594-9ff9-11ea-9e55-525403060300",
                "name": "test_name1",
                "name_other": "",
                "network_id": "1b0d0417-9757-43b2-9885-a3625162be08",
                "port_id": "16253ba3-93eb-4629-a998-2e1d6ed978ed",
                "status": "active",
                "tags": {
                    "test_tags1": "test1"
                },
                "tags_other": {},
                "tenant_connection_request_id": "276a0de2-9ff7-11ea-9e55-525403060300",
                "tenant_id": "2c76532f048849aab41c1bff2ec8b996",
                "tenant_id_other": "7af0c902bd51424f8b2c85f5320ab181"
            }
        }
expectedStatus:
    - Created
`
