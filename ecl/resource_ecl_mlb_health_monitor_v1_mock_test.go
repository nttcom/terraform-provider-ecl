package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1HealthMonitorResource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsCreate)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowAfterCreate)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsUpdateAttributes)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowBeforeUpdateConfigurations)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1HealthMonitorsUpdateConfigurations)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowAfterUpdateConfigurations)
	// Staged configurations of the load balancer and related resources are applied here
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowAfterApplyConfigurations)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowBeforeCreateConfigurations)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1HealthMonitorsCreateConfigurations)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowAfterCreateConfigurations)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsDelete)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1HealthMonitorsShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1HealthMonitor,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "name", "health_monitor"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "description", "description"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "tags.key", "value"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "port", "80"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "protocol", "http"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "interval", "5"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "retry", "3"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "timeout", "5"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "path", "/health"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "http_status_code", "200-299"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorUpdateBeforeApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "name", "health_monitor-update"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "port", "0"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "protocol", "icmp"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "interval", "5"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "retry", "3"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "timeout", "5"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "path", ""),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "http_status_code", ""),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorUpdateAfterApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "name", "health_monitor-update"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "port", "80"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "protocol", "http"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "interval", "5"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "retry", "3"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "timeout", "5"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "path", "/health"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "http_status_code", "200-299"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_health_monitor_v1.health_monitor", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
		},
	})
}

var testAccMLBV1HealthMonitor = fmt.Sprintf(`
resource "ecl_mlb_health_monitor_v1" "health_monitor" {
  name = "health_monitor"
  description = "description"
  tags = {
    key = "value"
  }
  port = 80
  protocol = "http"
  interval = 5
  retry = 3
  timeout = 5
  path = "/health"
  http_status_code = "200-299"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testAccMLBV1HealthMonitorUpdateBeforeApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_health_monitor_v1" "health_monitor" {
  name = "health_monitor-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  port = 0
  protocol = "icmp"
  interval = 5
  retry = 3
  timeout = 5
  path = ""
  http_status_code = ""
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testAccMLBV1HealthMonitorUpdateAfterApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_health_monitor_v1" "health_monitor" {
  name = "health_monitor-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  port = 80
  protocol = "http"
  interval = 5
  retry = 3
  timeout = 5
  path = "/health"
  http_status_code = "200-299"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1HealthMonitorsCreate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"health_monitor":{"description":"description","http_status_code":"200-299","interval":5,"load_balancer_id":"67fea379-cff0-4191-9175-de7d6941a040","name":"health_monitor","path":"/health","port":80,"protocol":"http","retry":3,"tags":{"key":"value"},"timeout":5}}
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": null,
        "protocol": null,
        "interval": null,
        "retry": null,
        "timeout": null,
        "path": null,
        "http_status_code": null
      }
    }
newStatus: Created
`)

var testMockMLBV1HealthMonitorsShowAfterCreate = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": null,
        "protocol": null,
        "interval": null,
        "retry": null,
        "timeout": null,
        "path": null,
        "http_status_code": null,
        "current": null,
        "staged": {
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      }
    }
expectedStatus:
  - Created
`)

var testMockMLBV1HealthMonitorsShowBeforeUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": null,
        "protocol": null,
        "interval": null,
        "retry": null,
        "timeout": null,
        "path": null,
        "http_status_code": null
      }
    }
expectedStatus:
  - AttributesUpdated
`)

var testMockMLBV1HealthMonitorsShowAfterUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": null,
        "protocol": null,
        "interval": null,
        "retry": null,
        "timeout": null,
        "path": null,
        "http_status_code": null,
        "current": null,
        "staged": {
          "port": 0,
          "protocol": "icmp",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "",
          "http_status_code": ""
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedBeforeApply
newStatus: ConfigurationsApplied
`)

var testMockMLBV1HealthMonitorsShowAfterApplyConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": 0,
        "protocol": "icmp",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "",
        "http_status_code": "",
        "current": {
          "port": 0,
          "protocol": "icmp",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "",
          "http_status_code": ""
        },
        "staged": null
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1HealthMonitorsShowBeforeCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": 0,
        "protocol": "icmp",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "",
        "http_status_code": ""
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1HealthMonitorsShowAfterCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "UPDATE_STAGED",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": 0,
        "protocol": "icmp",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "",
        "http_status_code": "",
        "current": {
          "port": 0,
          "protocol": "icmp",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "",
          "http_status_code": ""
        },
        "staged": {
          "port": 80,
          "protocol": "http",
          "interval": null,
          "retry": null,
          "timeout": null,
          "path": "/health",
          "http_status_code": "200-299"
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1HealthMonitorsShowAfterDelete = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "DELETE_STAGED",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": 0,
        "protocol": "icmp",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "",
        "http_status_code": "",
        "current": {
          "port": 0,
          "protocol": "icmp",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "",
          "http_status_code": ""
        },
        "staged": null
      }
    }
expectedStatus:
  - Deleted
`)

var testMockMLBV1HealthMonitorsUpdateAttributes = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"health_monitor":{"description":"description-update","name":"health_monitor-update","tags":{"key-update":"value-update"}}}
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": null,
        "protocol": null,
        "interval": null,
        "retry": null,
        "timeout": null,
        "path": null,
        "http_status_code": null
      }
    }
expectedStatus:
  - Created
newStatus: AttributesUpdated
`)

var testMockMLBV1HealthMonitorsCreateConfigurations = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"health_monitor":{"http_status_code":"200-299","path":"/health","port":80,"protocol":"http"}}
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "port": 80,
        "protocol": "http",
        "interval": null,
        "retry": null,
        "timeout": null,
        "path": "/health",
        "http_status_code": "200-299"
      }
    }
expectedStatus:
  - ConfigurationsApplied
newStatus: ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1HealthMonitorsUpdateConfigurations = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"health_monitor":{"http_status_code":"","path":"","port":0,"protocol":"icmp"}}
response:
  code: 200
  body: >
    {
      "health_monitor": {
        "port": 0,
        "protocol": "icmp",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "",
        "http_status_code": ""
      }
    }
expectedStatus:
  - AttributesUpdated
newStatus: ConfigurationsUpdatedBeforeApply
`)

var testMockMLBV1HealthMonitorsDelete = fmt.Sprintf(`
request:
  method: DELETE
response:
  code: 204
expectedStatus:
  - Created
  - ConfigurationsUpdatedAfterApply
newStatus: Deleted
`)
