package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1LoadBalancerActionResource_ApplyConfigurations(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowBeforeActionCreateStaged)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08/action", testMockMLBV1LoadBalancersActionApplyConfigurations)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterActionProcessing)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterActionCompleted)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1LoadBalancerActionApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "load_balancer_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "apply_configurations", "true"),
					resource.TestCheckNoResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "system_update"),
				),
			},
		},
	})
}

func TestMockedAccMLBV1LoadBalancerActionResource_SystemUpdate(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowBeforeActionCreateStaged)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08/action", testMockMLBV1LoadBalancersActionSystemUpdate)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterActionProcessing)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterActionCompleted)
	mc.Register(t, "system_updates", "/v1.0/system_updates/31746df7-92f9-4b5e-ad05-59f6684a54eb", testMockMLBV1SystemUpdatesShow)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1LoadBalancerActionSystemUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "load_balancer_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "system_update.system_update_id", "31746df7-92f9-4b5e-ad05-59f6684a54eb"),
					resource.TestCheckNoResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "apply_configurations"),
				),
			},
		},
	})
}

func TestMockedAccMLBV1LoadBalancerActionResource_ApplyConfigurationsAndSystemUpdate(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowBeforeActionCreateStaged)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08/action", testMockMLBV1LoadBalancersActionApplyConfigurationsAndSystemUpdate)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterActionProcessing)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowAfterActionCompleted)
	mc.Register(t, "system_updates", "/v1.0/system_updates/31746df7-92f9-4b5e-ad05-59f6684a54eb", testMockMLBV1SystemUpdatesShow)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1LoadBalancerActionApplyConfigurationsAndSystemUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "load_balancer_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "apply_configurations", "true"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "system_update.system_update_id", "31746df7-92f9-4b5e-ad05-59f6684a54eb"),
				),
			},
		},
	})
}

func TestMockedAccMLBV1LoadBalancerActionResource_NoApplyConfigurations(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowBeforeActionActive)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListEmpty)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListEmpty)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListEmpty)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListEmpty)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListEmpty)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsListEmpty)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1LoadBalancerActionApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "load_balancer_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "apply_configurations", "true"),
					resource.TestCheckNoResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "system_update"),
				),
			},
		},
	})
}

func TestMockedAccMLBV1LoadBalancerActionResource_NoSystemUpdate(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1LoadBalancersShowBeforeActionActive)
	mc.Register(t, "system_updates", "/v1.0/system_updates/31746df7-92f9-4b5e-ad05-59f6684a54eb", testMockMLBV1SystemUpdatesShow)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1LoadBalancerActionSystemUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "load_balancer_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "system_update.system_update_id", "31746df7-92f9-4b5e-ad05-59f6684a54eb"),
					resource.TestCheckNoResourceAttr("ecl_mlb_load_balancer_action_v1.load_balancer_action", "apply_configurations"),
				),
			},
		},
	})
}

var testAccMLBV1LoadBalancerActionApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_load_balancer_action_v1" "load_balancer_action" {
  load_balancer_id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
  apply_configurations = true
}
`)

var testAccMLBV1LoadBalancerActionSystemUpdate = fmt.Sprintf(`
resource "ecl_mlb_load_balancer_action_v1" "load_balancer_action" {
  load_balancer_id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
  system_update = {
    system_update_id = "31746df7-92f9-4b5e-ad05-59f6684a54eb"
  }
}
`)

var testAccMLBV1LoadBalancerActionApplyConfigurationsAndSystemUpdate = fmt.Sprintf(`
resource "ecl_mlb_load_balancer_action_v1" "load_balancer_action" {
  load_balancer_id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
  apply_configurations = true
  system_update = {
    system_update_id = "31746df7-92f9-4b5e-ad05-59f6684a54eb"
  }
}
`)

var testMockMLBV1LoadBalancersShowBeforeActionCreateStaged = fmt.Sprintf(`
request:
  method: GET
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
expectedStatus:
  - ~
`)

var testMockMLBV1LoadBalancersShowBeforeActionActive = fmt.Sprintf(`
request:
  method: GET
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
        "configuration_status": "ACTIVE",
        "monitoring_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "primary_availability_zone": "zone1_groupa",
        "secondary_availability_zone": "zone1_groupb",
        "active_availability_zone": "zone1_groupa",
        "revision": 2,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
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
expectedStatus:
  - ~
`)

var testMockMLBV1LoadBalancersShowAfterActionProcessing = fmt.Sprintf(`
request:
  method: GET
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
        "operation_status": "PROCESSING",
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
  - Performed
counter:
  max: 3
`)

var testMockMLBV1LoadBalancersShowAfterActionCompleted = fmt.Sprintf(`
request:
  method: GET
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
        "configuration_status": "ACTIVE",
        "monitoring_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "primary_availability_zone": "zone1_groupa",
        "secondary_availability_zone": "zone1_groupb",
        "active_availability_zone": "zone1_groupa",
        "revision": 2,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
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
expectedStatus:
  - Performed
counter:
  min: 4
`)

var testMockMLBV1LoadBalancersActionApplyConfigurations = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"apply-configurations":null}
response:
  code: 204
newStatus: Performed
`)

var testMockMLBV1LoadBalancersActionSystemUpdate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"system-update":{"system_update_id":"31746df7-92f9-4b5e-ad05-59f6684a54eb"}}
response:
  code: 204
newStatus: Performed
`)

var testMockMLBV1LoadBalancersActionApplyConfigurationsAndSystemUpdate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"apply-configurations":null,"system-update":{"system_update_id":"31746df7-92f9-4b5e-ad05-59f6684a54eb"}}
response:
  code: 204
newStatus: Performed
`)

var testMockMLBV1SystemUpdatesShow = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "system_update": {
        "id": "31746df7-92f9-4b5e-ad05-59f6684a54eb",
        "name": "security_update_202210",
        "description": "description",
        "href": "https://sdpf.ntt.com/news/2022100301/",
        "publish_datetime": "2022-10-03 00:00:00",
        "limit_datetime": "2022-10-11 12:59:59",
        "current_revision": 1,
        "next_revision": 2,
        "applicable": true
      }
    }
`)

var testMockMLBV1HealthMonitorsListEmpty = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 497f6eca-6276-4993-bfeb-53cbbbba6f08
response:
  code: 200
  body: >
    {
      "health_monitors": []
    }
`)

var testMockMLBV1ListenersListEmpty = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 497f6eca-6276-4993-bfeb-53cbbbba6f08
response:
  code: 200
  body: >
    {
      "listeners": []
    }
`)

var testMockMLBV1PoliciesListEmpty = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 497f6eca-6276-4993-bfeb-53cbbbba6f08
response:
  code: 200
  body: >
    {
      "policies": []
    }
`)

var testMockMLBV1RoutesListEmpty = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 497f6eca-6276-4993-bfeb-53cbbbba6f08
response:
  code: 200
  body: >
    {
      "routes": []
    }
`)

var testMockMLBV1RulesListEmpty = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 497f6eca-6276-4993-bfeb-53cbbbba6f08
response:
  code: 200
  body: >
    {
      "rules": []
    }
`)

var testMockMLBV1TargetGroupsListEmpty = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 497f6eca-6276-4993-bfeb-53cbbbba6f08
response:
  code: 200
  body: >
    {
      "target_groups": []
    }
`)
