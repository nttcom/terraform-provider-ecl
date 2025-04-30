package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1LoadBalancerResource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersCreate)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterCreate)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersUpdateAttributes)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowBeforeUpdateConfigurations)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1LoadBalancersUpdateConfigurations)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterUpdateConfigurations)
	// Staged configurations of the load balancer and related resources are applied here
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterApplyConfigurations)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowBeforeCreateConfigurations)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1LoadBalancersCreateConfigurations)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterCreateConfigurations)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersDelete)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterDeleteProcessing)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterDeleteCompleted)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1LoadBalancer,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "name", "load_balancer"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "description", "description"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "tags.key", "value"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerUpdateBeforeApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "name", "load_balancer-update"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.#", "2"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.1.ip_address", "192.168.1.6"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.1.port", "514"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.1.protocol", "udp"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.#", "2"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.1.network_id", "58e6d72b-f5e7-4b83-b306-06989ff78a84"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.1.virtual_ip_address", "192.168.1.1"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.1.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.1.reserved_fixed_ips.0.ip_address", "192.168.1.2"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.1.reserved_fixed_ips.1.ip_address", "192.168.1.3"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.1.reserved_fixed_ips.2.ip_address", "192.168.1.4"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.1.reserved_fixed_ips.3.ip_address", "192.168.1.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerUpdateAfterApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "name", "load_balancer-update"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.ip_address", "192.168.1.6"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.network_id", "58e6d72b-f5e7-4b83-b306-06989ff78a84"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.virtual_ip_address", "192.168.1.1"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.1.2"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.1.3"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.1.4"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_v1.load_balancer", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.1.5"),
				),
			},
		},
	})
}

var testAccMLBV1LoadBalancer = fmt.Sprintf(`
resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  name = "load_balancer"
  description = "description"
  tags = {
    key = "value"
  }
  plan_id = "00713021-9aea-41da-9a88-87760c08fa72"
  syslog_servers {
    ip_address = "192.168.0.6"
    port = 514
    protocol = "udp"
  }
  interfaces {
    network_id = "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"
    virtual_ip_address = "192.168.0.1"
    reserved_fixed_ips {
      ip_address = "192.168.0.2"
    }
    reserved_fixed_ips {
      ip_address = "192.168.0.3"
    }
    reserved_fixed_ips {
      ip_address = "192.168.0.4"
    }
    reserved_fixed_ips {
      ip_address = "192.168.0.5"
    }
  }
}
`)

var testAccMLBV1LoadBalancerUpdateBeforeApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  name = "load_balancer-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  plan_id = "00713021-9aea-41da-9a88-87760c08fa72"
  syslog_servers {
    ip_address = "192.168.0.6"
    port = 514
    protocol = "udp"
  }
  syslog_servers {
    ip_address = "192.168.1.6"
    port = 514
    protocol = "udp"
  }
  interfaces {
    network_id = "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"
    virtual_ip_address = "192.168.0.1"
    reserved_fixed_ips {
      ip_address = "192.168.0.2"
    }
    reserved_fixed_ips {
      ip_address = "192.168.0.3"
    }
    reserved_fixed_ips {
      ip_address = "192.168.0.4"
    }
    reserved_fixed_ips {
      ip_address = "192.168.0.5"
    }
  }
  interfaces {
    network_id = "58e6d72b-f5e7-4b83-b306-06989ff78a84"
    virtual_ip_address = "192.168.1.1"
    reserved_fixed_ips {
      ip_address = "192.168.1.2"
    }
    reserved_fixed_ips {
      ip_address = "192.168.1.3"
    }
    reserved_fixed_ips {
      ip_address = "192.168.1.4"
    }
    reserved_fixed_ips {
      ip_address = "192.168.1.5"
    }
  }
}
`)

var testAccMLBV1LoadBalancerUpdateAfterApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_load_balancer_v1" "load_balancer" {
  name = "load_balancer-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  plan_id = "00713021-9aea-41da-9a88-87760c08fa72"
  syslog_servers {
    ip_address = "192.168.1.6"
    port = 514
    protocol = "udp"
  }
  interfaces {
    network_id = "58e6d72b-f5e7-4b83-b306-06989ff78a84"
    virtual_ip_address = "192.168.1.1"
    reserved_fixed_ips {
      ip_address = "192.168.1.2"
    }
    reserved_fixed_ips {
      ip_address = "192.168.1.3"
    }
    reserved_fixed_ips {
      ip_address = "192.168.1.4"
    }
    reserved_fixed_ips {
      ip_address = "192.168.1.5"
    }
  }
}
`)

var testMockMLBV1LoadBalancersCreate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"load_balancer":{"description":"description","interfaces":[{"network_id":"d6797cf4-42b9-4cad-8591-9dd91c3f0fc3","reserved_fixed_ips":[{"ip_address":"192.168.0.2"},{"ip_address":"192.168.0.3"},{"ip_address":"192.168.0.4"},{"ip_address":"192.168.0.5"}],"virtual_ip_address":"192.168.0.1"}],"name":"load_balancer","plan_id":"00713021-9aea-41da-9a88-87760c08fa72","syslog_servers":[{"ip_address":"192.168.0.6","port":514,"protocol":"udp"}],"tags":{"key":"value"}}}
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "monitoring_status": "INITIAL",
        "operation_status": "NONE",
        "primary_availability_zone": null,
        "secondary_availability_zone": null,
        "active_availability_zone": "UNDEFINED",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": null,
        "interfaces": null
      }
    }
newStatus: Created
`)

var testMockMLBV1LoadBalancersShowAfterCreate = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "monitoring_status": "INITIAL",
        "operation_status": "NONE",
        "primary_availability_zone": null,
        "secondary_availability_zone": null,
        "active_availability_zone": "UNDEFINED",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": null,
        "interfaces": null,
        "current": null,
        "staged": {
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      }
    }
expectedStatus:
  - Created
`)

var testMockMLBV1LoadBalancersShowBeforeUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "monitoring_status": "INITIAL",
        "operation_status": "NONE",
        "primary_availability_zone": null,
        "secondary_availability_zone": null,
        "active_availability_zone": "UNDEFINED",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": null,
        "interfaces": null
      }
    }
expectedStatus:
  - AttributesUpdated
`)

var testMockMLBV1LoadBalancersShowAfterUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "monitoring_status": "INITIAL",
        "operation_status": "NONE",
        "primary_availability_zone": null,
        "secondary_availability_zone": null,
        "active_availability_zone": "UNDEFINED",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": null,
        "interfaces": null,
        "current": null,
        "staged": {
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            },
            {
              "ip_address": "192.168.1.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            },
            {
              "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
              "virtual_ip_address": "192.168.1.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.1.2"
                },
                {
                  "ip_address": "192.168.1.3"
                },
                {
                  "ip_address": "192.168.1.4"
                },
                {
                  "ip_address": "192.168.1.5"
                }
              ]
            }
          ]
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedBeforeApply
newStatus: ConfigurationsApplied
`)

var testMockMLBV1LoadBalancersShowAfterApplyConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "monitoring_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "primary_availability_zone": "zone1_groupa",
        "secondary_availability_zone": "zone1_groupb",
        "active_availability_zone": "zone1_groupa",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": [
          {
            "ip_address": "192.168.0.6",
            "port": 514,
            "protocol": "udp"
          },
          {
            "ip_address": "192.168.1.6",
            "port": 514,
            "protocol": "udp"
          }
        ],
        "interfaces": [
          {
            "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
            "virtual_ip_address": "192.168.0.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.0.2"
              },
              {
                "ip_address": "192.168.0.3"
              },
              {
                "ip_address": "192.168.0.4"
              },
              {
                "ip_address": "192.168.0.5"
              }
            ]
          },
          {
            "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
            "virtual_ip_address": "192.168.1.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.1.2"
              },
              {
                "ip_address": "192.168.1.3"
              },
              {
                "ip_address": "192.168.1.4"
              },
              {
                "ip_address": "192.168.1.5"
              }
            ]
          }
        ],
        "current": {
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            },
            {
              "ip_address": "192.168.1.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            },
            {
              "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
              "virtual_ip_address": "192.168.1.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.1.2"
                },
                {
                  "ip_address": "192.168.1.3"
                },
                {
                  "ip_address": "192.168.1.4"
                },
                {
                  "ip_address": "192.168.1.5"
                }
              ]
            }
          ]
        },
        "staged": null
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1LoadBalancersShowBeforeCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "monitoring_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "primary_availability_zone": "zone1_groupa",
        "secondary_availability_zone": "zone1_groupb",
        "active_availability_zone": "zone1_groupa",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": [
          {
            "ip_address": "192.168.0.6",
            "port": 514,
            "protocol": "udp"
          },
          {
            "ip_address": "192.168.1.6",
            "port": 514,
            "protocol": "udp"
          }
        ],
        "interfaces": [
          {
            "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
            "virtual_ip_address": "192.168.0.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.0.2"
              },
              {
                "ip_address": "192.168.0.3"
              },
              {
                "ip_address": "192.168.0.4"
              },
              {
                "ip_address": "192.168.0.5"
              }
            ]
          },
          {
            "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
            "virtual_ip_address": "192.168.1.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.1.2"
              },
              {
                "ip_address": "192.168.1.3"
              },
              {
                "ip_address": "192.168.1.4"
              },
              {
                "ip_address": "192.168.1.5"
              }
            ]
          }
        ]
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1LoadBalancersShowAfterCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "UPDATE_STAGED",
        "monitoring_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "primary_availability_zone": "zone1_groupa",
        "secondary_availability_zone": "zone1_groupb",
        "active_availability_zone": "zone1_groupa",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": [
          {
            "ip_address": "192.168.0.6",
            "port": 514,
            "protocol": "udp"
          },
          {
            "ip_address": "192.168.1.6",
            "port": 514,
            "protocol": "udp"
          }
        ],
        "interfaces": [
          {
            "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
            "virtual_ip_address": "192.168.0.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.0.2"
              },
              {
                "ip_address": "192.168.0.3"
              },
              {
                "ip_address": "192.168.0.4"
              },
              {
                "ip_address": "192.168.0.5"
              }
            ]
          },
          {
            "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
            "virtual_ip_address": "192.168.1.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.1.2"
              },
              {
                "ip_address": "192.168.1.3"
              },
              {
                "ip_address": "192.168.1.4"
              },
              {
                "ip_address": "192.168.1.5"
              }
            ]
          }
        ],
        "current": {
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            },
            {
              "ip_address": "192.168.1.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            },
            {
              "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
              "virtual_ip_address": "192.168.1.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.1.2"
                },
                {
                  "ip_address": "192.168.1.3"
                },
                {
                  "ip_address": "192.168.1.4"
                },
                {
                  "ip_address": "192.168.1.5"
                }
              ]
            }
          ]
        },
        "staged": {
          "syslog_servers": [
            {
              "ip_address": "192.168.1.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
              "virtual_ip_address": "192.168.1.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.1.2"
                },
                {
                  "ip_address": "192.168.1.3"
                },
                {
                  "ip_address": "192.168.1.4"
                },
                {
                  "ip_address": "192.168.1.5"
                }
              ]
            }
          ]
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1LoadBalancersShowAfterDeleteProcessing = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "UPDATE_STAGED",
        "monitoring_status": "ACTIVE",
        "operation_status": "PROCESSING",
        "primary_availability_zone": "zone1_groupa",
        "secondary_availability_zone": "zone1_groupb",
        "active_availability_zone": "zone1_groupa",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": [
          {
            "ip_address": "192.168.0.6",
            "port": 514,
            "protocol": "udp"
          },
          {
            "ip_address": "192.168.1.6",
            "port": 514,
            "protocol": "udp"
          }
        ],
        "interfaces": [
          {
            "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
            "virtual_ip_address": "192.168.0.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.0.2"
              },
              {
                "ip_address": "192.168.0.3"
              },
              {
                "ip_address": "192.168.0.4"
              },
              {
                "ip_address": "192.168.0.5"
              }
            ]
          },
          {
            "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
            "virtual_ip_address": "192.168.1.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.1.2"
              },
              {
                "ip_address": "192.168.1.3"
              },
              {
                "ip_address": "192.168.1.4"
              },
              {
                "ip_address": "192.168.1.5"
              }
            ]
          }
        ]
      }
    }
expectedStatus:
  - Deleted
counter:
  max: 3
`)

var testMockMLBV1LoadBalancersShowAfterDeleteCompleted = fmt.Sprintf(`
request:
  method: GET
response:
  code: 404
expectedStatus:
  - Deleted
counter:
  min: 4
`)

var testMockMLBV1LoadBalancersUpdateAttributes = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"load_balancer":{"description":"description-update","name":"load_balancer-update","tags":{"key-update":"value-update"}}}
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "monitoring_status": "INITIAL",
        "operation_status": "NONE",
        "primary_availability_zone": null,
        "secondary_availability_zone": null,
        "active_availability_zone": "UNDEFINED",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": null,
        "interfaces": null
      }
    }
expectedStatus:
  - Created
newStatus: AttributesUpdated
`)

var testMockMLBV1LoadBalancersCreateConfigurations = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"load_balancer":{"interfaces":[{"network_id":"58e6d72b-f5e7-4b83-b306-06989ff78a84","reserved_fixed_ips":[{"ip_address":"192.168.1.2"},{"ip_address":"192.168.1.3"},{"ip_address":"192.168.1.4"},{"ip_address":"192.168.1.5"}],"virtual_ip_address":"192.168.1.1"}],"syslog_servers":[{"ip_address":"192.168.1.6","port":514,"protocol":"udp"}]}}
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "syslog_servers": [
          {
            "ip_address": "192.168.1.6",
            "port": 514,
            "protocol": "udp"
          }
        ],
        "interfaces": [
          {
            "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
            "virtual_ip_address": "192.168.1.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.1.2"
              },
              {
                "ip_address": "192.168.1.3"
              },
              {
                "ip_address": "192.168.1.4"
              },
              {
                "ip_address": "192.168.1.5"
              }
            ]
          }
        ]
      }
    }
expectedStatus:
  - ConfigurationsApplied
newStatus: ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1LoadBalancersUpdateConfigurations = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"load_balancer":{"interfaces":[{"network_id":"d6797cf4-42b9-4cad-8591-9dd91c3f0fc3","reserved_fixed_ips":[{"ip_address":"192.168.0.2"},{"ip_address":"192.168.0.3"},{"ip_address":"192.168.0.4"},{"ip_address":"192.168.0.5"}],"virtual_ip_address":"192.168.0.1"},{"network_id":"58e6d72b-f5e7-4b83-b306-06989ff78a84","reserved_fixed_ips":[{"ip_address":"192.168.1.2"},{"ip_address":"192.168.1.3"},{"ip_address":"192.168.1.4"},{"ip_address":"192.168.1.5"}],"virtual_ip_address":"192.168.1.1"}],"syslog_servers":[{"ip_address":"192.168.0.6","port":514,"protocol":"udp"},{"ip_address":"192.168.1.6","port":514,"protocol":"udp"}]}}
response:
  code: 200
  body: >
    {
      "load_balancer": {
        "syslog_servers": [
          {
            "ip_address": "192.168.0.6",
            "port": 514,
            "protocol": "udp"
          },
          {
            "ip_address": "192.168.1.6",
            "port": 514,
            "protocol": "udp"
          }
        ],
        "interfaces": [
          {
            "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
            "virtual_ip_address": "192.168.0.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.0.2"
              },
              {
                "ip_address": "192.168.0.3"
              },
              {
                "ip_address": "192.168.0.4"
              },
              {
                "ip_address": "192.168.0.5"
              }
            ]
          },
          {
            "network_id": "58e6d72b-f5e7-4b83-b306-06989ff78a84",
            "virtual_ip_address": "192.168.1.1",
            "reserved_fixed_ips": [
              {
                "ip_address": "192.168.1.2"
              },
              {
                "ip_address": "192.168.1.3"
              },
              {
                "ip_address": "192.168.1.4"
              },
              {
                "ip_address": "192.168.1.5"
              }
            ]
          }
        ]
      }
    }
expectedStatus:
  - AttributesUpdated
newStatus: ConfigurationsUpdatedBeforeApply
`)

var testMockMLBV1LoadBalancersDelete = fmt.Sprintf(`
request:
  method: DELETE
response:
  code: 204
expectedStatus:
  - Created
  - ConfigurationsUpdatedAfterApply
newStatus: Deleted
`)
