package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1PolicyResource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesCreate)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowAfterCreate)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesUpdateAttributes)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowBeforeUpdateConfigurations)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1PoliciesUpdateConfigurations)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowAfterUpdateConfigurations)
	// Staged configurations of the load balancer and related resources are applied here
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowAfterApplyConfigurations)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowBeforeCreateConfigurations)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1PoliciesCreateConfigurations)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowAfterCreateConfigurations)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesDelete)
	mc.Register(t, "policies", "/v1.0/policies/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1PoliciesShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Policy,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "name", "policy"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "description", "description"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tags.key", "value"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "persistence", "cookie"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "source_nat", "enable"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1PolicyUpdateBeforeApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "name", "policy-update"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "algorithm", "least-connection"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "persistence", "source-ip"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "idle_timeout", "120"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "sorry_page_url", ""),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "source_nat", "disable"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "certificate_id", ""),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tls_policy_id", ""),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1PolicyUpdateAfterApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "name", "policy-update"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "persistence", "cookie"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "source_nat", "enable"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_policy_v1.policy", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
		},
	})
}

var testAccMLBV1Policy = fmt.Sprintf(`
resource "ecl_mlb_policy_v1" "policy" {
  name = "policy"
  description = "description"
  tags = {
    key = "value"
  }
  algorithm = "round-robin"
  persistence = "cookie"
  idle_timeout = 600
  sorry_page_url = "https://example.com/sorry"
  source_nat = "enable"
  certificate_id = "f57a98fe-d63e-4048-93a0-51fe163f30d7"
  health_monitor_id = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
  listener_id = "68633f4f-f52a-402f-8572-b8173418904f"
  default_target_group_id = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
  tls_policy_id = "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testAccMLBV1PolicyUpdateBeforeApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_policy_v1" "policy" {
  name = "policy-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  algorithm = "least-connection"
  persistence = "source-ip"
  idle_timeout = 120
  sorry_page_url = ""
  source_nat = "disable"
  certificate_id = ""
  health_monitor_id = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
  listener_id = "68633f4f-f52a-402f-8572-b8173418904f"
  default_target_group_id = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
  tls_policy_id = ""
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testAccMLBV1PolicyUpdateAfterApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_policy_v1" "policy" {
  name = "policy-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  algorithm = "round-robin"
  persistence = "cookie"
  idle_timeout = 600
  sorry_page_url = "https://example.com/sorry"
  source_nat = "enable"
  certificate_id = "f57a98fe-d63e-4048-93a0-51fe163f30d7"
  health_monitor_id = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
  listener_id = "68633f4f-f52a-402f-8572-b8173418904f"
  default_target_group_id = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
  tls_policy_id = "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1PoliciesCreate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"policy":{"algorithm":"round-robin","certificate_id":"f57a98fe-d63e-4048-93a0-51fe163f30d7","default_target_group_id":"a44c4072-ed90-4b50-a33a-6b38fb10c7db","description":"description","health_monitor_id":"dd7a96d6-4e66-4666-baca-a8555f0c472c","idle_timeout":600,"listener_id":"68633f4f-f52a-402f-8572-b8173418904f","load_balancer_id":"67fea379-cff0-4191-9175-de7d6941a040","name":"policy","persistence":"cookie","sorry_page_url":"https://example.com/sorry","source_nat":"enable","tags":{"key":"value"},"tls_policy_id":"4ba79662-f2a1-41a4-a3d9-595799bbcd86"}}
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": null,
        "persistence": null,
        "idle_timeout": null,
        "sorry_page_url": null,
        "source_nat": null,
        "certificate_id": null,
        "health_monitor_id": null,
        "listener_id": null,
        "default_target_group_id": null,
        "tls_policy_id": null
      }
    }
newStatus: Created
`)

var testMockMLBV1PoliciesShowAfterCreate = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": null,
        "persistence": null,
        "idle_timeout": null,
        "sorry_page_url": null,
        "source_nat": null,
        "certificate_id": null,
        "health_monitor_id": null,
        "listener_id": null,
        "default_target_group_id": null,
        "tls_policy_id": null,
        "current": null,
        "staged": {
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      }
    }
expectedStatus:
  - Created
`)

var testMockMLBV1PoliciesShowBeforeUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": null,
        "persistence": null,
        "idle_timeout": null,
        "sorry_page_url": null,
        "source_nat": null,
        "certificate_id": null,
        "health_monitor_id": null,
        "listener_id": null,
        "default_target_group_id": null,
        "tls_policy_id": null
      }
    }
expectedStatus:
  - AttributesUpdated
`)

var testMockMLBV1PoliciesShowAfterUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": null,
        "persistence": null,
        "idle_timeout": null,
        "sorry_page_url": null,
        "source_nat": null,
        "certificate_id": null,
        "health_monitor_id": null,
        "listener_id": null,
        "default_target_group_id": null,
        "tls_policy_id": null,
        "current": null,
        "staged": {
          "algorithm": "least-connection",
          "persistence": "source-ip",
          "idle_timeout": 120,
          "sorry_page_url": "",
          "source_nat": "disable",
          "certificate_id": "",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": ""
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedBeforeApply
newStatus: ConfigurationsApplied
`)

var testMockMLBV1PoliciesShowAfterApplyConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": "least-connection",
        "persistence": "source-ip",
        "idle_timeout": 120,
        "sorry_page_url": "",
        "source_nat": "disable",
        "certificate_id": "",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "tls_policy_id": "",
        "current": {
          "algorithm": "least-connection",
          "persistence": "source-ip",
          "idle_timeout": 120,
          "sorry_page_url": "",
          "source_nat": "disable",
          "certificate_id": "",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": ""
        },
        "staged": null
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1PoliciesShowBeforeCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": "least-connection",
        "persistence": "source-ip",
        "idle_timeout": 120,
        "sorry_page_url": "",
        "source_nat": "disable",
        "certificate_id": "",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "tls_policy_id": ""
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1PoliciesShowAfterCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "UPDATE_STAGED",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": "least-connection",
        "persistence": "source-ip",
        "idle_timeout": 120,
        "sorry_page_url": "",
        "source_nat": "disable",
        "certificate_id": "",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "tls_policy_id": "",
        "current": {
          "algorithm": "least-connection",
          "persistence": "source-ip",
          "idle_timeout": 120,
          "sorry_page_url": "",
          "source_nat": "disable",
          "certificate_id": "",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": ""
        },
        "staged": {
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": null,
          "listener_id": null,
          "default_target_group_id": null,
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1PoliciesShowAfterDelete = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "DELETE_STAGED",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": "least-connection",
        "persistence": "source-ip",
        "idle_timeout": 120,
        "sorry_page_url": "",
        "source_nat": "disable",
        "certificate_id": "",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "tls_policy_id": "",
        "current": {
          "algorithm": "least-connection",
          "persistence": "source-ip",
          "idle_timeout": 120,
          "sorry_page_url": "",
          "source_nat": "disable",
          "certificate_id": "",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": ""
        },
        "staged": null
      }
    }
expectedStatus:
  - Deleted
`)

var testMockMLBV1PoliciesUpdateAttributes = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"policy":{"description":"description-update","name":"policy-update","tags":{"key-update":"value-update"}}}
response:
  code: 200
  body: >
    {
      "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": null,
        "persistence": null,
        "idle_timeout": null,
        "sorry_page_url": null,
        "source_nat": null,
        "certificate_id": null,
        "health_monitor_id": null,
        "listener_id": null,
        "default_target_group_id": null,
        "tls_policy_id": null
      }
    }
expectedStatus:
  - Created
newStatus: AttributesUpdated
`)

var testMockMLBV1PoliciesCreateConfigurations = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"policy":{"algorithm":"round-robin","certificate_id":"f57a98fe-d63e-4048-93a0-51fe163f30d7","idle_timeout":600,"persistence":"cookie","sorry_page_url":"https://example.com/sorry","source_nat":"enable","tls_policy_id":"4ba79662-f2a1-41a4-a3d9-595799bbcd86"}}
response:
  code: 200
  body: >
    {
      "policy": {
        "algorithm": "round-robin",
        "persistence": "cookie",
        "idle_timeout": 600,
        "sorry_page_url": "https://example.com/sorry",
        "source_nat": "enable",
        "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
        "health_monitor_id": null,
        "listener_id": null,
        "default_target_group_id": null,
        "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
      }
    }
expectedStatus:
  - ConfigurationsApplied
newStatus: ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1PoliciesUpdateConfigurations = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"policy":{"algorithm":"least-connection","certificate_id":"","idle_timeout":120,"persistence":"source-ip","sorry_page_url":"","source_nat":"disable","tls_policy_id":""}}
response:
  code: 200
  body: >
    {
      "policy": {
        "algorithm": "least-connection",
        "persistence": "source-ip",
        "idle_timeout": 120,
        "sorry_page_url": "",
        "source_nat": "disable",
        "certificate_id": "",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "tls_policy_id": ""
      }
    }
expectedStatus:
  - AttributesUpdated
newStatus: ConfigurationsUpdatedBeforeApply
`)

var testMockMLBV1PoliciesDelete = fmt.Sprintf(`
request:
  method: DELETE
response:
  code: 204
expectedStatus:
  - Created
  - ConfigurationsUpdatedAfterApply
newStatus: Deleted
`)
