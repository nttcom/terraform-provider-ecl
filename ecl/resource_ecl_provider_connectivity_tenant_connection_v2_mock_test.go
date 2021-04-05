package ecl

import (
	"fmt"
	"testing"

	"github.com/nttcom/eclcloud/v2/ecl/provider_connectivity/v2/tenant_connections"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestMockedAccProviderConnectivityV2TenantConnection_basic(t *testing.T) {
	var connection tenant_connections.TenantConnection

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections", testMockProviderConnectivityV2TenantConnectionCreateComputeServer)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionGetAfterCreateComputeServer)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionUpdateComputeServer)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionGetAfterUpdateComputeServer)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionDelete)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/06969a7c-9fc0-11ea-b509-525403060400", testMockProviderConnectivityV2TenantConnectionGetAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testMockAccProviderConnectivityV2TenantConnectionComputeServerConfig,
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
				Config: testMockAccProviderConnectivityV2TenantConnectionComputeSeverUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "updated_name"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "updated_desc"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.k2", "v2"),
				),
			},
		},
	})
}

var testMockAccProviderConnectivityV2TenantConnectionComputeServerConfig = `
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

var testMockAccProviderConnectivityV2TenantConnectionComputeSeverUpdateConfig = `
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

var testMockProviderConnectivityV2TenantConnectionCreateComputeServer = `
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

var testMockProviderConnectivityV2TenantConnectionGetAfterCreateComputeServer = `
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

var testMockProviderConnectivityV2TenantConnectionUpdateComputeServer = `
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

var testMockProviderConnectivityV2TenantConnectionGetAfterUpdateComputeServer = `
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
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections", testMockProviderConnectivityV2TenantConnectionCreateBaremetalServer)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/76ea7594-9ff9-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionGetAfterCreateBaremetalServer)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/76ea7594-9ff9-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionDelete)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/76ea7594-9ff9-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionGetAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testMockAccProviderConnectivityV2TenantConnectionBaremetalServerConfig,
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

var testMockAccProviderConnectivityV2TenantConnectionBaremetalServerConfig = `
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

var testMockProviderConnectivityV2TenantConnectionCreateBaremetalServer = `
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

var testMockProviderConnectivityV2TenantConnectionGetAfterCreateBaremetalServer = `
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

func TestMockedAccProviderConnectivityV2TenantConnection_vna(t *testing.T) {
	var connection tenant_connections.TenantConnection

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections", testMockProviderConnectivityV2TenantConnectionCreateVna)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/a27e1006-a00a-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionGetAfterCreateVna)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/a27e1006-a00a-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionDelete)
	mc.Register(t, "provider-connectivity", "/v2.0/tenant_connections/a27e1006-a00a-11ea-9e55-525403060300", testMockProviderConnectivityV2TenantConnectionGetAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testMockAccProviderConnectivityV2TenantConnectionVnaConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProviderConnectivityV2TenantConnectionExists("ecl_provider_connectivity_tenant_connection_v2.connection_1", &connection),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "id", "a27e1006-a00a-11ea-9e55-525403060300"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_connection_request_id", "94e0f9cc-a00a-11ea-9ada-525403060500"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "test_name1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "test_desc1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.test_tags1", "test1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_id", "2c76532f048849aab41c1bff2ec8b996"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_id_other", "7af0c902bd51424f8b2c85f5320ab181"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "network_id", "045d234e-203f-4e35-affb-a1d3ba0380e0"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_type", "ECL::VirtualNetworkAppliance::VSRX"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_id", "baab608e-7f3c-4b67-82b9-1e6fa7befc95"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_interface_id", "interface_2"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "port_id", ""),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "status", "active"),
				),
			},
		},
	})
}

var testMockAccProviderConnectivityV2TenantConnectionVnaConfig = `
resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "94e0f9cc-a00a-11ea-9ada-525403060500"
	device_type = "ECL::VirtualNetworkAppliance::VSRX"
	device_id = "baab608e-7f3c-4b67-82b9-1e6fa7befc95"
 	device_interface_id = "interface_2"
	attachment_opts_vna {
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
}
`

var testMockProviderConnectivityV2TenantConnectionCreateVna = `
request:
    method: POST
response:
    code: 200
    body: >
        {
            "tenant_connection": {
                "description": "test_desc1",
                "description_other": "",
                "device_id": "baab608e-7f3c-4b67-82b9-1e6fa7befc95",
                "device_interface_id": "interface_2",
                "device_type": "ECL::VirtualNetworkAppliance::VSRX",
                "id": "a27e1006-a00a-11ea-9e55-525403060300",
                "name": "test_name1",
                "name_other": "",
                "network_id": "045d234e-203f-4e35-affb-a1d3ba0380e0",
                "port_id": "",
                "status": "creating",
                "tags": {
                    "test_tags1": "test1"
                },
                "tags_other": "{}",
                "tenant_connection_request_id": "94e0f9cc-a00a-11ea-9ada-525403060500",
                "tenant_id": "2c76532f048849aab41c1bff2ec8b996",
                "tenant_id_other": "7af0c902bd51424f8b2c85f5320ab181"
            }
        }
newStatus: Created
`

var testMockProviderConnectivityV2TenantConnectionGetAfterCreateVna = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "tenant_connection": {
                "description": "test_desc1",
                "description_other": "",
                "device_id": "baab608e-7f3c-4b67-82b9-1e6fa7befc95",
                "device_interface_id": "interface_2",
                "device_type": "ECL::VirtualNetworkAppliance::VSRX",
                "id": "a27e1006-a00a-11ea-9e55-525403060300",
                "name": "test_name1",
                "name_other": "",
                "network_id": "045d234e-203f-4e35-affb-a1d3ba0380e0",
                "port_id": "",
                "status": "active",
                "tags": {
                    "test_tags1": "test1"
                },
                "tags_other": {},
                "tenant_connection_request_id": "94e0f9cc-a00a-11ea-9ada-525403060500",
                "tenant_id": "2c76532f048849aab41c1bff2ec8b996",
                "tenant_id_other": "7af0c902bd51424f8b2c85f5320ab181"
            }
        }
expectedStatus:
    - Created
`
