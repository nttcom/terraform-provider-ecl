package ecl

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-ecl/ecl/testhelper/mock"
	"testing"
)

// TestMockedAccNetworkV2LoadBalancer_basic tests basic behavior of Load Balancer creation and update requests.
// Step 0: Create Load Balancer with 2 connected interfaces (One with VIP configurations and one without)
//           and 2 syslog servers.
// Step 1: Update Load Balancer and all sub resources as much as possible without recreating resources.
func TestMockedAccNetworkV2LoadBalancer_basic(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	// Mock registration: keystone
	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)

	// Mock registration: Network
	mc.Register(t, "network1", "/v2.0/networks", testMockNetworkV2LoadBalancerNetwork1Post)
	mc.Register(t, "network1", "/v2.0/networks/89b67a81-1a2a-4adf-b096-3f78f0d62831", testMockNetworkV2LoadBalancerNetwork1GetAfterCreate)
	mc.Register(t, "network1", "/v2.0/networks/89b67a81-1a2a-4adf-b096-3f78f0d62831", testMockNetworkV2LoadBalancerDelete)
	mc.Register(t, "network1", "/v2.0/networks/89b67a81-1a2a-4adf-b096-3f78f0d62831", testMockNetworkV2LoadBalancerGetDeleted)
	mc.Register(t, "network2", "/v2.0/networks", testMockNetworkV2LoadBalancerNetwork2Post)
	mc.Register(t, "network2", "/v2.0/networks/77d3282f-2974-4902-840f-c0d2d5b0fa34", testMockNetworkV2LoadBalancerNetwork2GetAfterCreate)
	mc.Register(t, "network2", "/v2.0/networks/77d3282f-2974-4902-840f-c0d2d5b0fa34", testMockNetworkV2LoadBalancerDelete)
	mc.Register(t, "network2", "/v2.0/networks/77d3282f-2974-4902-840f-c0d2d5b0fa34", testMockNetworkV2LoadBalancerGetDeleted)

	// Mock registration: Subnet
	mc.Register(t, "subnet1", "/v2.0/subnets", testMockNetworkV2LoadBalancerSubnet1Post)
	mc.Register(t, "subnet1", "/v2.0/subnets/f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678", testMockNetworkV2LoadBalancerSubnet1GetAfterCreate)
	mc.Register(t, "subnet1", "/v2.0/subnets/f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678", testMockNetworkV2LoadBalancerDelete)
	mc.Register(t, "subnet1", "/v2.0/subnets/f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678", testMockNetworkV2LoadBalancerGetDeleted)
	mc.Register(t, "subnet2", "/v2.0/subnets", testMockNetworkV2LoadBalancerSubnet2Post)
	mc.Register(t, "subnet2", "/v2.0/subnets/deedc8e7-543b-4c36-b1fa-d94ba71e7707", testMockNetworkV2LoadBalancerSubnet2GetAfterCreate)
	mc.Register(t, "subnet2", "/v2.0/subnets/deedc8e7-543b-4c36-b1fa-d94ba71e7707", testMockNetworkV2LoadBalancerDelete)
	mc.Register(t, "subnet2", "/v2.0/subnets/deedc8e7-543b-4c36-b1fa-d94ba71e7707", testMockNetworkV2LoadBalancerGetDeleted)

	// Mock registration: Load Balancer Plan
	mc.Register(t, "load_balancer_plans", "/v2.0/load_balancer_plans", testMockNetworkV2LoadBalancerPlanList)
	mc.Register(t, "load_balancer_plans", "/v2.0/load_balancer_plans/ed306566-646d-4132-a96a-3a984da9a4ca", testMockNetworkV2LoadBalancerPlanGet4IF)
	mc.Register(t, "load_balancer_plans", "/v2.0/load_balancer_plans/c17c4f48-8aad-4083-81d5-5a4f68a71ea0", testMockNetworkV2LoadBalancerPlanGet8IF)

	// Mock registration: Load Balancer Step 0
	mc.Register(t, "load_balancers", "/v2.0/load_balancers", testMockNetworkV2LoadBalancerPost)
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerGetAfterCreate)
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerPut1) // Update default_gateway
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerGetAfterUpdate1)

	// Mock registration: Load Balancer Step 1
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerPut2) // Initialize default_gateway before updating Load Balancer Interface
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerGetAfterUpdate2)
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerPut3) // Update load_balancer_plan_id, default_gateway, ...
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerGetAfterUpdate3)
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerDelete)
	mc.Register(t, "load_balancers", "/v2.0/load_balancers/", testMockNetworkV2LoadBalancerGetDeleted)

	// Mock registration: Load Balancer Interface Step 1
	mc.Register(t, "load_balancer_interfaces1", "/v2.0/load_balancer_interfaces/0439e533-3f67-47c0-919c-8c9b698257e9", testMockNetworkV2LoadBalancerInterface1GetAfterUpdate2)
	mc.Register(t, "load_balancer_interfaces1", "/v2.0/load_balancer_interfaces/0439e533-3f67-47c0-919c-8c9b698257e9", testMockNetworkV2LoadBalancerInterface1Put2)
	mc.Register(t, "load_balancer_interfaces2", "/v2.0/load_balancer_interfaces/c44889e9-89f6-413d-b186-307dde40d125", testMockNetworkV2LoadBalancerInterface2GetAfterUpdate2)
	mc.Register(t, "load_balancer_interfaces2", "/v2.0/load_balancer_interfaces/c44889e9-89f6-413d-b186-307dde40d125", testMockNetworkV2LoadBalancerInterface2Put2)

	// Mock registration: Load Balancer Interface Step 0
	mc.Register(t, "load_balancer_interfaces1", "/v2.0/load_balancer_interfaces/0439e533-3f67-47c0-919c-8c9b698257e9", testMockNetworkV2LoadBalancerInterface1GetAfterUpdate1)
	mc.Register(t, "load_balancer_interfaces1", "/v2.0/load_balancer_interfaces/0439e533-3f67-47c0-919c-8c9b698257e9", testMockNetworkV2LoadBalancerInterface1Put1)
	mc.Register(t, "load_balancer_interfaces2", "/v2.0/load_balancer_interfaces/c44889e9-89f6-413d-b186-307dde40d125", testMockNetworkV2LoadBalancerInterface2GetAfterUpdate1)
	mc.Register(t, "load_balancer_interfaces2", "/v2.0/load_balancer_interfaces/c44889e9-89f6-413d-b186-307dde40d125", testMockNetworkV2LoadBalancerInterface2Put1)

	// Mock registration: Load Balancer Interface Initial GET
	mc.Register(t, "load_balancer_interfaces1", "/v2.0/load_balancer_interfaces/0439e533-3f67-47c0-919c-8c9b698257e9", testMockNetworkV2LoadBalancerInterface1GetInit)
	mc.Register(t, "load_balancer_interfaces2", "/v2.0/load_balancer_interfaces/c44889e9-89f6-413d-b186-307dde40d125", testMockNetworkV2LoadBalancerInterface2GetInit)
	mc.Register(t, "load_balancer_interfaces3", "/v2.0/load_balancer_interfaces/494dbd21-e182-46f5-8c57-42db2d756fb2", testMockNetworkV2LoadBalancerInterface3GetInit)
	mc.Register(t, "load_balancer_interfaces4", "/v2.0/load_balancer_interfaces/dfe5e5f2-1f13-443e-9a9a-966501a8dd75", testMockNetworkV2LoadBalancerInterface4GetInit)

	// Mock registration: Load Balancer Syslog Server Step 0
	mc.Register(t, "load_balancer_syslog_servers1", "/v2.0/load_balancer_syslog_servers", testMockNetworkV2LoadBalancerSyslogServer1Post)
	mc.Register(t, "load_balancer_syslog_servers1", "/v2.0/load_balancer_syslog_servers/079eba0a-95a1-4f31-8979-7a409c3da148", testMockNetworkV2LoadBalancerSyslogServer1GetAfterCreate)
	mc.Register(t, "load_balancer_syslog_servers2", "/v2.0/load_balancer_syslog_servers", testMockNetworkV2LoadBalancerSyslogServer2Post)
	mc.Register(t, "load_balancer_syslog_servers2", "/v2.0/load_balancer_syslog_servers/c4410f15-4a35-47fd-a659-21a1eabd11cb", testMockNetworkV2LoadBalancerSyslogServer2GetAfterCreate)

	// Mock registration: Load Balancer Syslog Server Step 1
	mc.Register(t, "load_balancer_syslog_servers1", "/v2.0/load_balancer_syslog_servers/079eba0a-95a1-4f31-8979-7a409c3da148", testMockNetworkV2LoadBalancerSyslogServer1Put1)
	mc.Register(t, "load_balancer_syslog_servers1", "/v2.0/load_balancer_syslog_servers/079eba0a-95a1-4f31-8979-7a409c3da148", testMockNetworkV2LoadBalancerSyslogServer1GetAfterUpdate1)
	mc.Register(t, "load_balancer_syslog_servers1", "/v2.0/load_balancer_syslog_servers/079eba0a-95a1-4f31-8979-7a409c3da148", testMockNetworkV2LoadBalancerSyslogServerDelete)
	mc.Register(t, "load_balancer_syslog_servers1", "/v2.0/load_balancer_syslog_servers/079eba0a-95a1-4f31-8979-7a409c3da148", testMockNetworkV2LoadBalancerGetDeleted)
	mc.Register(t, "load_balancer_syslog_servers2", "/v2.0/load_balancer_syslog_servers/c4410f15-4a35-47fd-a659-21a1eabd11cb", testMockNetworkV2LoadBalancerSyslogServer2Put1)
	mc.Register(t, "load_balancer_syslog_servers2", "/v2.0/load_balancer_syslog_servers/c4410f15-4a35-47fd-a659-21a1eabd11cb", testMockNetworkV2LoadBalancerSyslogServer2GetAfterUpdate1)
	mc.Register(t, "load_balancer_syslog_servers2", "/v2.0/load_balancer_syslog_servers/c4410f15-4a35-47fd-a659-21a1eabd11cb", testMockNetworkV2LoadBalancerSyslogServerDelete)
	mc.Register(t, "load_balancer_syslog_servers2", "/v2.0/load_balancer_syslog_servers/c4410f15-4a35-47fd-a659-21a1eabd11cb", testMockNetworkV2LoadBalancerGetDeleted)

	mc.StartServer(t)

	testAccNetworkV2LoadBalancerBasicSteps(t)
}

var testMockNetworkV2LoadBalancerNetwork1Post = fmt.Sprintf(`
request:
    method: POST
    body: '{"network":{"name":"network_1","plane":"data"}}'
response:
    code: 201
    body: >
        {
            "network": {
              "admin_state_up": true,
              "description": "",
              "id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
              "name": "network_1",
              "plane": "data",
              "shared": false,
              "status": "PENDING_CREATE",
              "subnets": [],
              "tenant_id": "%s"
            }
          }
newStatus: Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerNetwork1GetAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "network": {
              "admin_state_up": true,
              "description": "",
              "id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
              "name": "network_1",
              "plane": "data",
              "shared": false,
              "status": "ACTIVE",
              "subnets": [],
              "tenant_id": "%s"
            }
          }
expectedStatus:
    - Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSubnet1Post = fmt.Sprintf(`
request:
    method: POST
    body: '{"subnet":{"allocation_pools":[{"end":"192.168.151.200","start":"192.168.151.100"}],"cidr":"192.168.151.0/24","gateway_ip":"192.168.151.1","ip_version":4,"name":"subnet_1_1","network_id":"89b67a81-1a2a-4adf-b096-3f78f0d62831"}}'
response:
    code: 201
    body: >
        {
            "subnet": {
              "allocation_pools": [
                {
                    "end": "192.168.151.200",
                    "start": "192.168.151.100"
                }
              ],
              "cidr": "192.168.151.0/24",
              "description": "",
              "enable_dhcp": true,
              "gateway_ip": "192.168.151.1",
              "host_routes": [],
              "id": "f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678",
              "ip_version": 4,
              "name": "subnet_1_1",
              "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
              "ntp_servers": [],
              "status": "PENDING_CREATE",
              "tags": {},
              "tenant_id": "%s"
            }
        }
newStatus: Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSubnet1GetAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "subnet": {
              "allocation_pools": [
                {
                    "end": "192.168.151.200",
                    "start": "192.168.151.100"
                }
              ],
              "cidr": "192.168.151.0/24",
              "description": "",
              "enable_dhcp": true,
              "gateway_ip": "192.168.151.1",
              "host_routes": [],
              "id": "f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678",
              "ip_version": 4,
              "name": "subnet_1_1",
              "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
              "ntp_servers": [],
              "status": "ACTIVE",
              "tags": {},
              "tenant_id": "%s"
            }
        }
expectedStatus:
    - Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerNetwork2Post = fmt.Sprintf(`
request:
    method: POST
    body: '{"network":{"name":"network_2","plane":"data"}}'
response:
    code: 201
    body: >
        {
            "network": {
              "admin_state_up": true,
              "description": "",
              "id": "77d3282f-2974-4902-840f-c0d2d5b0fa34",
              "name": "network_2",
              "plane": "data",
              "shared": false,
              "status": "PENDING_CREATE",
              "subnets": [],
              "tenant_id": "%s"
            }
        }
newStatus: Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerNetwork2GetAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "network": {
              "admin_state_up": true,
              "description": "",
              "id": "77d3282f-2974-4902-840f-c0d2d5b0fa34",
              "name": "network_2",
              "plane": "data",
              "shared": false,
              "status": "ACTIVE",
              "subnets": [],
              "tenant_id": "%s"
            }
        }
expectedStatus:
    - Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSubnet2Post = fmt.Sprintf(`
request:
    method: POST
    body: '{"subnet":{"allocation_pools":[{"end":"192.168.152.200","start":"192.168.152.100"}],"cidr":"192.168.152.0/24","gateway_ip":"192.168.152.1","ip_version":4,"name":"subnet_2_1","network_id":"77d3282f-2974-4902-840f-c0d2d5b0fa34"}}'
response:
    code: 201
    body: >
        {
            "subnet": {
              "allocation_pools": [
                {
                    "end": "192.168.152.200",
                    "start": "192.168.152.100"
                }
              ],
              "cidr": "192.168.152.0/24",
              "description": "",
              "enable_dhcp": true,
              "gateway_ip": "192.168.152.1",
              "host_routes": [],
              "id": "deedc8e7-543b-4c36-b1fa-d94ba71e7707",
              "ip_version": 4,
              "name": "subnet_2_1",
              "network_id": "77d3282f-2974-4902-840f-c0d2d5b0fa34",
              "ntp_servers": [],
              "status": "PENDING_CREATE",
              "tags": {},
              "tenant_id": "%s"
            }
        }
newStatus: Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSubnet2GetAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "subnet": {
              "allocation_pools": [
                {
                    "end": "192.168.152.200",
                    "start": "192.168.152.100"
                }
              ],
              "cidr": "192.168.152.0/24",
              "description": "",
              "enable_dhcp": true,
              "gateway_ip": "192.168.152.1",
              "host_routes": [],
              "id": "deedc8e7-543b-4c36-b1fa-d94ba71e7707",
              "ip_version": 4,
              "name": "subnet_2_1",
              "network_id": "77d3282f-2974-4902-840f-c0d2d5b0fa34",
              "ntp_servers": [],
              "status": "ACTIVE",
              "tags": {},
              "tenant_id": "%s"
            }
        }
expectedStatus:
    - Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerPlanList = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_plans": [
            {
              "description": "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_200Mbps_2CPU-8GB-4IF",
              "enabled": true,
              "id": "ed306566-646d-4132-a96a-3a984da9a4ca",
              "maximum_syslog_servers": 10,
              "model": {
                "edition": "Standard",
                "size": "200"
              },
              "name": "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_200Mbps_2CPU-8GB-4IF",
              "vendor": "citrix",
              "version": "12.1-55.18"
            },
            {
              "description": "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_1000Mbps_4CPU-16GB-8IF",
              "enabled": true,
              "id": "c17c4f48-8aad-4083-81d5-5a4f68a71ea0",
              "maximum_syslog_servers": 10,
              "model": {
                "edition": "Standard",
                "size": "1000"
              },
              "name": "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_1000Mbps_4CPU-16GB-8IF",
              "vendor": "citrix",
              "version": "12.1-55.18"
            }
          ]
        }
`

var testMockNetworkV2LoadBalancerPlanGet4IF = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "load_balancer_plan": {
                "description": "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_200Mbps_2CPU-8GB-4IF",
                "enabled": true,
                "id": "ed306566-646d-4132-a96a-3a984da9a4ca",
                "maximum_syslog_servers": 10,
                "model": {
                  "edition": "Standard",
                  "size": "200"
                },
                "name": "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_200Mbps_2CPU-8GB-4IF",
                "vendor": "citrix",
                "version": "12.1-55.18"
            }
        }
`

var testMockNetworkV2LoadBalancerPlanGet8IF = `
request:
    method: GET
response:
    code: 200
    body: >
        {
            "load_balancer_plan": {
                "description": "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_1000Mbps_4CPU-16GB-8IF",
                "enabled": true,
                "id": "c17c4f48-8aad-4083-81d5-5a4f68a71ea0",
                "maximum_syslog_servers": 10,
                "model": {
                  "edition": "Standard",
                  "size": "1000"
                },
                "name": "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_1000Mbps_4CPU-16GB-8IF",
                "vendor": "citrix",
                "version": "12.1-55.18"
            }
        }
`

var testMockNetworkV2LoadBalancerPost = fmt.Sprintf(`
request:
    method: POST
response:
    code: 201
    body: >
        {
          "load_balancer": {
            "admin_password": "FEPtzZst",
            "admin_username": "user-admin",
            "availability_zone": "zone1_groupa",
            "default_gateway": null,
            "description": "load_balancer_test1_description",
            "id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "interfaces": [
              {
                "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
                "ip_address": "",
                "name": "Interface 1/1",
                "network_id": "",
                "slot_number": 1,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "c44889e9-89f6-413d-b186-307dde40d125",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 2,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 3,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 4,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              }
            ],
            "load_balancer_plan_id": "ed306566-646d-4132-a96a-3a984da9a4ca",
            "name": "lb_test1",
            "status": "PENDING_CREATE",
            "syslog_servers": null,
            "tenant_id": "%s",
            "user_password": "t43h4NiU",
            "user_username": "user-read"
          }
        }
newStatus: Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerGetAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer": {
            "admin_username": "user-admin",
            "availability_zone": "zone1_groupa",
            "default_gateway": null,
            "description": "load_balancer_test1_description",
            "id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "interfaces": [
              {
                "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
                "ip_address": "",
                "name": "Interface 1/1",
                "network_id": "",
                "slot_number": 1,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "c44889e9-89f6-413d-b186-307dde40d125",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 2,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 3,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 4,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              }
            ],
            "load_balancer_plan_id": "ed306566-646d-4132-a96a-3a984da9a4ca",
            "name": "lb_test1",
            "status": "ACTIVE",
            "syslog_servers": null,
            "tenant_id": "%s",
            "user_username": "user-read"
          }
        }
expectedStatus:
    - Created
`, OS_TENANT_ID)

var testMockNetworkV2LoadBalancerInterface1GetInit = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface1_description",
            "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
            "ip_address": "192.168.151.11",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "Interface 1/2",
            "network_id": null,
            "slot_number": 1,
            "status": "DOWN",
            "tenant_id": "%s",
            "virtual_ip_address": null,
            "virtual_ip_properties": null
          }
        }
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface1Put1 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface1_description",
            "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
            "ip_address": "192.168.151.11",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "lb_test1_interface1",
            "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
            "slot_number": 1,
            "status": "PENDING_UPDATE",
            "tenant_id": "%s",
            "virtual_ip_address": "192.168.151.31",
            "virtual_ip_properties": {
              "protocol": "vrrp",
              "vrid": 20
            }
          }
        }
counter:
    max: 1
newStatus: Updated1
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface1GetAfterUpdate1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface1_description",
            "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
            "ip_address": "192.168.151.11",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "lb_test1_interface1",
            "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
            "slot_number": 1,
            "status": "ACTIVE",
            "tenant_id": "%s",
            "virtual_ip_address": "192.168.151.31",
            "virtual_ip_properties": {
              "protocol": "vrrp",
              "vrid": 20
            }
          }
        }
expectedStatus:
    - Updated1
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface1Put2 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface1_description_update",
            "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
            "ip_address": "192.168.151.12",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "lb_test1_interface1_update",
            "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
            "slot_number": 1,
            "status": "PENDING_UPDATE",
            "tenant_id": "%s",
            "virtual_ip_address": "192.168.151.32",
            "virtual_ip_properties": {
              "protocol": "vrrp",
              "vrid": 30
            }
          }
        }
expectedStatus:
    - Updated1
newStatus: Updated2
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface1GetAfterUpdate2 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface1_description_update",
            "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
            "ip_address": "192.168.151.12",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "lb_test1_interface1_update",
            "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
            "slot_number": 1,
            "status": "ACTIVE",
            "tenant_id": "%s",
            "virtual_ip_address": "192.168.151.32",
            "virtual_ip_properties": {
              "protocol": "vrrp",
              "vrid": 30
            }
          }
        }
expectedStatus:
    - Updated2
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerPut1 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
          "load_balancer": {
            "admin_username": "user-admin",
            "availability_zone": "zone1_groupa",
            "default_gateway": "192.168.151.1",
            "description": "load_balancer_test1_description",
            "id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "interfaces": [
              {
                "description": "lb_test1_interface1_description",
                "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
                "ip_address": "192.168.151.11",
                "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                "name": "lb_test1_interface1",
                "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
                "slot_number": 1,
                "status": "ACTIVE",
                "tenant_id": "%s",
                "virtual_ip_address": "192.168.151.31",
                "virtual_ip_properties": {
                  "protocol": "vrrp",
                  "vrid": 20
                }
              },
              {
                "id": "c44889e9-89f6-413d-b186-307dde40d125",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 2,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 3,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 4,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              }
            ],
            "load_balancer_plan_id": "ed306566-646d-4132-a96a-3a984da9a4ca",
            "name": "lb_test1",
            "status": "PENDING_UPDATE",
            "syslog_servers": [
                {
                            "acl_logging": "ENABLED",
                            "appflow_logging": "ENABLED",
                            "date_format": "MMDDYYYY",
                            "description": "lb_test1_syslog1_description",
                            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
                            "ip_address": "192.168.151.21",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL0",
                            "log_level": "ALERT|CRITICAL|EMERGENCY",
                            "name": "lb_test1_syslog1",
                            "port_number": 514,
                            "priority": 20,
                            "status": "ACTIVE",
                            "tcp_logging": "ALL",
                            "tenant_id": "%s",
                            "time_zone": "LOCAL_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "NO"
                },
                {
                            "acl_logging": "DISABLED",
                            "appflow_logging": "DISABLED",
                            "date_format": "YYYYMMDD",
                            "description": "lb_test1_syslog2_description",
                            "id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
                            "ip_address": "192.168.151.22",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL1",
                            "log_level": "DEBUG",
                            "name": "lb_test1_syslog2",
                            "port_number": 514,
                            "priority": 0,
                            "status": "ACTIVE",
                            "tcp_logging": "NONE",
                            "tenant_id": "%s",
                            "time_zone": "GMT_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "YES"
                }
            ],
            "tenant_id": "%s",
            "user_username": "user-read"
          }
        }
expectedStatus:
    - Created
newStatus: Updated1
`,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerGetAfterUpdate1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer": {
            "admin_username": "user-admin",
            "availability_zone": "zone1_groupa",
            "default_gateway": "192.168.151.1",
            "description": "load_balancer_test1_description",
            "id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "interfaces": [
              {
                "description": "lb_test1_interface1_description",
                "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
                "ip_address": "192.168.151.11",
                "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                "name": "lb_test1_interface1",
                "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
                "slot_number": 1,
                "status": "ACTIVE",
                "tenant_id": "%s",
                "virtual_ip_address": "192.168.151.31",
                "virtual_ip_properties": {
                  "protocol": "vrrp",
                  "vrid": 20
                }
              },
              {
                "id": "c44889e9-89f6-413d-b186-307dde40d125",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 2,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 3,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 4,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              }
            ],
            "load_balancer_plan_id": "ed306566-646d-4132-a96a-3a984da9a4ca",
            "name": "lb_test1",
            "status": "ACTIVE",
            "syslog_servers": [
                {
                            "acl_logging": "ENABLED",
                            "appflow_logging": "ENABLED",
                            "date_format": "MMDDYYYY",
                            "description": "lb_test1_syslog1_description",
                            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
                            "ip_address": "192.168.151.21",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL0",
                            "log_level": "ALERT|CRITICAL|EMERGENCY",
                            "name": "lb_test1_syslog1",
                            "port_number": 514,
                            "priority": 20,
                            "status": "ACTIVE",
                            "tcp_logging": "ALL",
                            "tenant_id": "%s",
                            "time_zone": "LOCAL_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "NO"
                },
                {
                            "acl_logging": "DISABLED",
                            "appflow_logging": "DISABLED",
                            "date_format": "YYYYMMDD",
                            "description": "lb_test1_syslog2_description",
                            "id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
                            "ip_address": "192.168.151.22",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL1",
                            "log_level": "DEBUG",
                            "name": "lb_test1_syslog2",
                            "port_number": 514,
                            "priority": 0,
                            "status": "ACTIVE",
                            "tcp_logging": "NONE",
                            "tenant_id": "%s",
                            "time_zone": "GMT_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "YES"
                }
            ],
            "tenant_id": "%s",
            "user_username": "user-read"
          }
        }
expectedStatus:
    - Updated1
`, OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerPut2 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
          "load_balancer": {
            "admin_username": "user-admin",
            "availability_zone": "zone1_groupa",
            "default_gateway": null,
            "description": "load_balancer_test1_description",
            "id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "interfaces": [
              {
                "description": "lb_test1_interface1_description",
                "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
                "ip_address": "192.168.151.11",
                "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                "name": "lb_test1_interface1",
                "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
                "slot_number": 1,
                "status": "ACTIVE",
                "tenant_id": "%s",
                "virtual_ip_address": "192.168.151.31",
                "virtual_ip_properties": {
                  "protocol": "vrrp",
                  "vrid": 20
                }
              },
              {
                "id": "c44889e9-89f6-413d-b186-307dde40d125",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 2,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 3,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 4,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              }
            ],
            "load_balancer_plan_id": "ed306566-646d-4132-a96a-3a984da9a4ca",
            "name": "lb_test1",
            "status": "PENDING_UPDATE",
            "syslog_servers": [
                {
                            "acl_logging": "ENABLED",
                            "appflow_logging": "ENABLED",
                            "date_format": "MMDDYYYY",
                            "description": "lb_test1_syslog1_description",
                            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
                            "ip_address": "192.168.151.21",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL0",
                            "log_level": "ALERT|CRITICAL|EMERGENCY",
                            "name": "lb_test1_syslog1",
                            "port_number": 514,
                            "priority": 20,
                            "status": "ACTIVE",
                            "tcp_logging": "ALL",
                            "tenant_id": "%s",
                            "time_zone": "LOCAL_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "NO"
                },
                {
                            "acl_logging": "DISABLED",
                            "appflow_logging": "DISABLED",
                            "date_format": "YYYYMMDD",
                            "description": "lb_test1_syslog2_description",
                            "id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
                            "ip_address": "192.168.151.22",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL1",
                            "log_level": "DEBUG",
                            "name": "lb_test1_syslog2",
                            "port_number": 514,
                            "priority": 0,
                            "status": "ACTIVE",
                            "tcp_logging": "NONE",
                            "tenant_id": "%s",
                            "time_zone": "GMT_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "YES"
                }
            ],
            "tenant_id": "%s",
            "user_username": "user-read"
          }
        }
expectedStatus:
    - Updated1
newStatus: Updated2
`,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerGetAfterUpdate2 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer": {
            "admin_username": "user-admin",
            "availability_zone": "zone1_groupa",
            "default_gateway": null,
            "description": "load_balancer_test1_description",
            "id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "interfaces": [
              {
                "description": "lb_test1_interface1_description",
                "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
                "ip_address": "192.168.151.11",
                "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                "name": "lb_test1_interface1",
                "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
                "slot_number": 1,
                "status": "ACTIVE",
                "tenant_id": "%s",
                "virtual_ip_address": "192.168.151.31",
                "virtual_ip_properties": {
                  "protocol": "vrrp",
                  "vrid": 20
                }
              },
              {
                "id": "c44889e9-89f6-413d-b186-307dde40d125",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 2,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 3,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 4,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              }
            ],
            "load_balancer_plan_id": "ed306566-646d-4132-a96a-3a984da9a4ca",
            "name": "lb_test1",
            "status": "ACTIVE",
            "syslog_servers": [
                {
                            "acl_logging": "ENABLED",
                            "appflow_logging": "ENABLED",
                            "date_format": "MMDDYYYY",
                            "description": "lb_test1_syslog1_description",
                            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
                            "ip_address": "192.168.151.21",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL0",
                            "log_level": "ALERT|CRITICAL|EMERGENCY",
                            "name": "lb_test1_syslog1",
                            "port_number": 514,
                            "priority": 20,
                            "status": "ACTIVE",
                            "tcp_logging": "ALL",
                            "tenant_id": "%s",
                            "time_zone": "LOCAL_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "NO"
                },
                {
                            "acl_logging": "DISABLED",
                            "appflow_logging": "DISABLED",
                            "date_format": "YYYYMMDD",
                            "description": "lb_test1_syslog2_description",
                            "id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
                            "ip_address": "192.168.151.22",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL1",
                            "log_level": "DEBUG",
                            "name": "lb_test1_syslog2",
                            "port_number": 514,
                            "priority": 0,
                            "status": "ACTIVE",
                            "tcp_logging": "NONE",
                            "tenant_id": "%s",
                            "time_zone": "GMT_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "YES"
                }
            ],
            "tenant_id": "%s",
            "user_username": "user-read"
          }
        }
expectedStatus:
    - Updated2
`, OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerPut3 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
          "load_balancer": {
            "admin_username": "user-admin",
            "availability_zone": "zone1_groupa",
            "default_gateway": "192.168.152.1",
            "description": "load_balancer_test1_description_update",
            "id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "interfaces": [
              {
                "description": "lb_test1_interface1_description",
                "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
                "ip_address": "192.168.151.11",
                "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                "name": "lb_test1_interface1",
                "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
                "slot_number": 1,
                "status": "ACTIVE",
                "tenant_id": "%s",
                "virtual_ip_address": "192.168.151.31",
                "virtual_ip_properties": {
                  "protocol": "vrrp",
                  "vrid": 20
                }
              },
              {
                "id": "c44889e9-89f6-413d-b186-307dde40d125",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 2,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 3,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 4,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              }
            ],
            "load_balancer_plan_id": "c17c4f48-8aad-4083-81d5-5a4f68a71ea0",
            "name": "lb_test1_update",
            "status": "PENDING_UPDATE",
            "syslog_servers": [
                {
                            "acl_logging": "ENABLED",
                            "appflow_logging": "ENABLED",
                            "date_format": "MMDDYYYY",
                            "description": "lb_test1_syslog1_description",
                            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
                            "ip_address": "192.168.151.21",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL0",
                            "log_level": "ALERT|CRITICAL|EMERGENCY",
                            "name": "lb_test1_syslog1",
                            "port_number": 514,
                            "priority": 20,
                            "status": "ACTIVE",
                            "tcp_logging": "ALL",
                            "tenant_id": "%s",
                            "time_zone": "LOCAL_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "NO"
                },
                {
                            "acl_logging": "DISABLED",
                            "appflow_logging": "DISABLED",
                            "date_format": "YYYYMMDD",
                            "description": "lb_test1_syslog2_description",
                            "id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
                            "ip_address": "192.168.151.22",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL1",
                            "log_level": "DEBUG",
                            "name": "lb_test1_syslog2",
                            "port_number": 514,
                            "priority": 0,
                            "status": "ACTIVE",
                            "tcp_logging": "NONE",
                            "tenant_id": "%s",
                            "time_zone": "GMT_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "YES"
                }
            ],
            "tenant_id": "%s",
            "user_username": "user-read"
          }
        }
expectedStatus:
    - Updated2
newStatus: Updated3
`,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerGetAfterUpdate3 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer": {
            "admin_username": "user-admin",
            "availability_zone": "zone1_groupa",
            "default_gateway": "192.168.152.1",
            "description": "load_balancer_test1_description_update",
            "id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "interfaces": [
              {
                "description": "lb_test1_interface1_description",
                "id": "0439e533-3f67-47c0-919c-8c9b698257e9",
                "ip_address": "192.168.151.11",
                "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                "name": "lb_test1_interface1",
                "network_id": "89b67a81-1a2a-4adf-b096-3f78f0d62831",
                "slot_number": 1,
                "status": "ACTIVE",
                "tenant_id": "%s",
                "virtual_ip_address": "192.168.151.31",
                "virtual_ip_properties": {
                  "protocol": "vrrp",
                  "vrid": 20
                }
              },
              {
                "id": "c44889e9-89f6-413d-b186-307dde40d125",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 2,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 3,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              },
              {
                "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
                "ip_address": "",
                "name": "Interface 1/2",
                "network_id": "",
                "slot_number": 4,
                "status": "DOWN",
                "type": "user",
                "virtual_ip_address": null,
                "virtual_ip_properties": null
              }
            ],
            "load_balancer_plan_id": "c17c4f48-8aad-4083-81d5-5a4f68a71ea0",
            "name": "lb_test1_update",
            "status": "ACTIVE",
            "syslog_servers": [
                {
                            "acl_logging": "ENABLED",
                            "appflow_logging": "ENABLED",
                            "date_format": "MMDDYYYY",
                            "description": "lb_test1_syslog1_description",
                            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
                            "ip_address": "192.168.151.21",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL0",
                            "log_level": "ALERT|CRITICAL|EMERGENCY",
                            "name": "lb_test1_syslog1",
                            "port_number": 514,
                            "priority": 20,
                            "status": "ACTIVE",
                            "tcp_logging": "ALL",
                            "tenant_id": "%s",
                            "time_zone": "LOCAL_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "NO"
                },
                {
                            "acl_logging": "DISABLED",
                            "appflow_logging": "DISABLED",
                            "date_format": "YYYYMMDD",
                            "description": "lb_test1_syslog2_description",
                            "id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
                            "ip_address": "192.168.151.22",
                            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
                            "log_facility": "LOCAL1",
                            "log_level": "DEBUG",
                            "name": "lb_test1_syslog2",
                            "port_number": 514,
                            "priority": 0,
                            "status": "ACTIVE",
                            "tcp_logging": "NONE",
                            "tenant_id": "%s",
                            "time_zone": "GMT_TIME",
                            "transport_type": "UDP",
                            "user_configurable_log_messages": "YES"
                }
            ],
            "tenant_id": "%s",
            "user_username": "user-read"
          }
        }
expectedStatus:
    - Updated3
`, OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSyslogServer1Post = fmt.Sprintf(`
request:
    method: POST
    body: '{"load_balancer_syslog_server":{"acl_logging":"ENABLED","appflow_logging":"ENABLED","date_format":"MMDDYYYY","description":"lb_test1_syslog1_description","ip_address":"192.168.151.21","load_balancer_id":"caa214f9-84f3-402b-b3e6-f0e2ebe28cf9","log_facility":"LOCAL0","log_level":"ALERT|CRITICAL|EMERGENCY","name":"lb_test1_syslog1","port_number":514,"priority":20,"tcp_logging":"ALL","time_zone":"LOCAL_TIME","transport_type":"UDP","user_configurable_log_messages":"NO"}}'
response:
    code: 201
    body: >
        {
          "load_balancer_syslog_server": {
            "acl_logging": "ENABLED",
            "appflow_logging": "ENABLED",
            "date_format": "MMDDYYYY",
            "description": "lb_test1_syslog1_description",
            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
            "ip_address": "192.168.151.21",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "log_facility": "LOCAL0",
            "log_level": "ALERT|CRITICAL|EMERGENCY",
            "name": "lb_test1_syslog1",
            "port_number": 514,
            "priority": 20,
            "status": "PENDING_CREATE",
            "tcp_logging": "ALL",
            "tenant_id": "%s",
            "time_zone": "LOCAL_TIME",
            "transport_type": "UDP",
            "user_configurable_log_messages": "NO"
          }
        }
counter:
    max: 1
newStatus: Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSyslogServer1GetAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_syslog_server": {
            "acl_logging": "ENABLED",
            "appflow_logging": "ENABLED",
            "date_format": "MMDDYYYY",
            "description": "lb_test1_syslog1_description",
            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
            "ip_address": "192.168.151.21",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "log_facility": "LOCAL0",
            "log_level": "ALERT|CRITICAL|EMERGENCY",
            "name": "lb_test1_syslog1",
            "port_number": 514,
            "priority": 20,
            "status": "ACTIVE",
            "tcp_logging": "ALL",
            "tenant_id": "%s",
            "time_zone": "LOCAL_TIME",
            "transport_type": "UDP",
            "user_configurable_log_messages": "NO"
          }
        }
expectedStatus:
    - Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSyslogServer1Put1 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
          "load_balancer_syslog_server": {
            "acl_logging": "DISABLED",
            "appflow_logging": "DISABLED",
            "date_format": "YYYYMMDD",
            "description": "lb_test1_syslog1_description_update",
            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
            "ip_address": "192.168.151.21",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "log_facility": "LOCAL1",
            "log_level": "DEBUG",
            "name": "lb_test1_syslog1",
            "port_number": 514,
            "priority": 0,
            "status": "PENDING_UPDATE",
            "tcp_logging": "NONE",
            "tenant_id": "%s",
            "time_zone": "GMT_TIME",
            "transport_type": "UDP",
            "user_configurable_log_messages": "YES"
          }
        }
counter:
    max: 1
expectedStatus:
    - Created
newStatus: Updated1
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSyslogServer1GetAfterUpdate1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_syslog_server": {
            "acl_logging": "DISABLED",
            "appflow_logging": "DISABLED",
            "date_format": "YYYYMMDD",
            "description": "lb_test1_syslog1_description_update",
            "id": "079eba0a-95a1-4f31-8979-7a409c3da148",
            "ip_address": "192.168.151.21",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "log_facility": "LOCAL1",
            "log_level": "DEBUG",
            "name": "lb_test1_syslog1",
            "port_number": 514,
            "priority": 0,
            "status": "ACTIVE",
            "tcp_logging": "NONE",
            "tenant_id": "%s",
            "time_zone": "GMT_TIME",
            "transport_type": "UDP",
            "user_configurable_log_messages": "YES"
          }
        }
expectedStatus:
    - Updated1
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSyslogServer2Post = fmt.Sprintf(`
request:
    method: POST
    body: '{"load_balancer_syslog_server":{"acl_logging":"DISABLED","appflow_logging":"DISABLED","date_format":"YYYYMMDD","description":"lb_test1_syslog2_description","ip_address":"192.168.151.22","load_balancer_id":"caa214f9-84f3-402b-b3e6-f0e2ebe28cf9","log_facility":"LOCAL1","log_level":"DEBUG","name":"lb_test1_syslog2","port_number":514,"priority":0,"tcp_logging":"NONE","time_zone":"GMT_TIME","transport_type":"UDP","user_configurable_log_messages":"YES"}}'
response:
    code: 201
    body: >
        {
        	"load_balancer_syslog_server": {
        		"acl_logging": "DISABLED",
        		"appflow_logging": "DISABLED",
        		"date_format": "YYYYMMDD",
        		"description": "lb_test1_syslog2_description",
        		"id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
        		"ip_address": "192.168.151.22",
        		"load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
        		"log_facility": "LOCAL1",
        		"log_level": "DEBUG",
        		"name": "lb_test1_syslog2",
        		"port_number": 514,
        		"priority": 0,
        		"status": "PENDING_CREATE",
        		"tcp_logging": "NONE",
        		"tenant_id": "%s",
        		"time_zone": "GMT_TIME",
        		"transport_type": "UDP",
        		"user_configurable_log_messages": "YES"
        	}
        }
counter:
    max: 1
newStatus: Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSyslogServer2GetAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"load_balancer_syslog_server": {
        		"acl_logging": "DISABLED",
        		"appflow_logging": "DISABLED",
        		"date_format": "YYYYMMDD",
        		"description": "lb_test1_syslog2_description",
        		"id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
        		"ip_address": "192.168.151.22",
        		"load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
        		"log_facility": "LOCAL1",
        		"log_level": "DEBUG",
        		"name": "lb_test1_syslog2",
        		"port_number": 514,
        		"priority": 0,
        		"status": "ACTIVE",
        		"tcp_logging": "NONE",
        		"tenant_id": "%s",
        		"time_zone": "GMT_TIME",
        		"transport_type": "UDP",
        		"user_configurable_log_messages": "YES"
        	}
        }
expectedStatus:
    - Created
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSyslogServer2Put1 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
        	"load_balancer_syslog_server": {
        		"acl_logging": "ENABLED",
        		"appflow_logging": "ENABLED",
        		"date_format": "MMDDYYYY",
        		"description": "lb_test1_syslog2_description_update",
        		"id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
        		"ip_address": "192.168.151.22",
        		"load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
        		"log_facility": "LOCAL0",
        		"log_level": "ALERT|CRITICAL|EMERGENCY",
        		"name": "lb_test1_syslog2",
        		"port_number": 514,
        		"priority": 20,
        		"status": "PENDING_UPDATE",
        		"tcp_logging": "ALL",
        		"tenant_id": "%s",
        		"time_zone": "LOCAL_TIME",
        		"transport_type": "UDP",
        		"user_configurable_log_messages": "NO"
        	}
        }
counter:
    max: 1
expectedStatus:
    - Created
newStatus: Updated1
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerSyslogServer2GetAfterUpdate1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"load_balancer_syslog_server": {
        		"acl_logging": "ENABLED",
        		"appflow_logging": "ENABLED",
        		"date_format": "MMDDYYYY",
        		"description": "lb_test1_syslog2_description_update",
        		"id": "c4410f15-4a35-47fd-a659-21a1eabd11cb",
        		"ip_address": "192.168.151.22",
        		"load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
        		"log_facility": "LOCAL0",
        		"log_level": "ALERT|CRITICAL|EMERGENCY",
        		"name": "lb_test1_syslog2",
        		"port_number": 514,
        		"priority": 20,
        		"status": "ACTIVE",
        		"tcp_logging": "ALL",
        		"tenant_id": "%s",
        		"time_zone": "LOCAL_TIME",
        		"transport_type": "UDP",
        		"user_configurable_log_messages": "NO"
        	}
        }
expectedStatus:
    - Updated1
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface2GetInit = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "id": "c44889e9-89f6-413d-b186-307dde40d125",
            "ip_address": "",
            "name": "Interface 1/2",
            "network_id": "",
            "slot_number": 2,
            "status": "DOWN",
            "type": "user",
            "virtual_ip_address": null,
            "virtual_ip_properties": null
          }
        }
`

var testMockNetworkV2LoadBalancerInterface2Put1 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface2_description",
            "id": "c44889e9-89f6-413d-b186-307dde40d125",
            "ip_address": "192.168.152.11",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "lb_test1_interface2",
            "network_id": "77d3282f-2974-4902-840f-c0d2d5b0fa34",
            "slot_number": 2,
            "status": "PENDING_UPDATE",
            "tenant_id": "%s",
            "virtual_ip_address": null,
            "virtual_ip_properties": null
          }
        }
counter:
    max: 1
newStatus: Updated1
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface2GetAfterUpdate1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface2_description",
            "id": "c44889e9-89f6-413d-b186-307dde40d125",
            "ip_address": "192.168.152.11",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "lb_test1_interface2",
            "network_id": "77d3282f-2974-4902-840f-c0d2d5b0fa34",
            "slot_number": 2,
            "status": "ACTIVE",
            "tenant_id": "%s",
            "virtual_ip_address": null,
            "virtual_ip_properties": null
          }
        }
expectedStatus:
    - Updated1
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface2Put2 = fmt.Sprintf(`
request:
    method: PUT
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface2_description_update",
            "id": "c44889e9-89f6-413d-b186-307dde40d125",
            "ip_address": "192.168.152.12",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "lb_test1_interface2_update",
            "network_id": "77d3282f-2974-4902-840f-c0d2d5b0fa34",
            "slot_number": 2,
            "status": "PENDING_UPDATE",
            "tenant_id": "%s",
            "virtual_ip_address": null,
            "virtual_ip_properties": null
          }
        }
expectedStatus:
    - Updated1
newStatus: Updated2
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface2GetAfterUpdate2 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "description": "lb_test1_interface2_description_update",
            "id": "c44889e9-89f6-413d-b186-307dde40d125",
            "ip_address": "192.168.152.12",
            "load_balancer_id": "caa214f9-84f3-402b-b3e6-f0e2ebe28cf9",
            "name": "lb_test1_interface2_update",
            "network_id": "77d3282f-2974-4902-840f-c0d2d5b0fa34",
            "slot_number": 2,
            "status": "ACTIVE",
            "tenant_id": "%s",
            "virtual_ip_address": null,
            "virtual_ip_properties": null
          }
        }
expectedStatus:
    - Updated2
`,
	OS_TENANT_ID,
)

var testMockNetworkV2LoadBalancerInterface3GetInit = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "id": "494dbd21-e182-46f5-8c57-42db2d756fb2",
            "ip_address": "",
            "name": "Interface 1/2",
            "network_id": "",
            "slot_number": 3,
            "status": "DOWN",
            "type": "user",
            "virtual_ip_address": null,
            "virtual_ip_properties": null
          }
        }
`

var testMockNetworkV2LoadBalancerInterface4GetInit = `
request:
    method: GET
response:
    code: 200
    body: >
        {
          "load_balancer_interface": {
            "id": "dfe5e5f2-1f13-443e-9a9a-966501a8dd75",
            "ip_address": "",
            "name": "Interface 1/2",
            "network_id": "",
            "slot_number": 4,
            "status": "DOWN",
            "type": "user",
            "virtual_ip_address": null,
            "virtual_ip_properties": null
          }
        }
`

var testMockNetworkV2LoadBalancerSyslogServerDelete = `
request:
    method: DELETE
response:
    code: 204
expectedStatus:
    - Created
    - Updated
    - Updated1
newStatus: Deleted
`

var testMockNetworkV2LoadBalancerDelete = `
request:
    method: DELETE
response:
    code: 204
expectedStatus:
    - Created
    - Updated
    - Updated1
    - Updated2
    - Updated3
newStatus: Deleted
`

var testMockNetworkV2LoadBalancerGetDeleted = `
request:
    method: GET
response:
    code: 404
expectedStatus:
    - Deleted
`
