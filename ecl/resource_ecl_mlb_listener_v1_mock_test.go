package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1ListenerResource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersCreate)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowAfterCreate)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersUpdateAttributes)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowBeforeUpdateConfigurations)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1ListenersUpdateConfigurations)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowAfterUpdateConfigurations)
	// Staged configurations of the load balancer and related resources are applied here
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowAfterApplyConfigurations)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowBeforeCreateConfigurations)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1ListenersCreateConfigurations)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowAfterCreateConfigurations)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersDelete)
	mc.Register(t, "listeners", "/v1.0/listeners/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1ListenersShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Listener,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "name", "listener"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "description", "description"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "tags.key", "value"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "port", "80"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "protocol", "http"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1ListenerUpdateBeforeApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "name", "listener-update"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "port", "443"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "protocol", "https"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1ListenerUpdateAfterApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "name", "listener-update"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "port", "80"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "protocol", "http"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_listener_v1.listener", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
		},
	})
}

var testAccMLBV1Listener = fmt.Sprintf(`
resource "ecl_mlb_listener_v1" "listener" {
  name = "listener"
  description = "description"
  tags = {
    key = "value"
  }
  ip_address = "10.0.0.1"
  port = 80
  protocol = "http"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testAccMLBV1ListenerUpdateBeforeApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_listener_v1" "listener" {
  name = "listener-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  ip_address = "10.0.0.1"
  port = 443
  protocol = "https"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testAccMLBV1ListenerUpdateAfterApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_listener_v1" "listener" {
  name = "listener-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  ip_address = "10.0.0.1"
  port = 80
  protocol = "http"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1ListenersCreate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"listener":{"description":"description","ip_address":"10.0.0.1","load_balancer_id":"67fea379-cff0-4191-9175-de7d6941a040","name":"listener","port":80,"protocol":"http","tags":{"key":"value"}}}
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": null,
        "port": null,
        "protocol": null
      }
    }
newStatus: Created
`)

var testMockMLBV1ListenersShowAfterCreate = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": null,
        "port": null,
        "protocol": null,
        "current": null,
        "staged": {
          "ip_address": "10.0.0.1",
          "port": 80,
          "protocol": "http"
        }
      }
    }
expectedStatus:
  - Created
`)

var testMockMLBV1ListenersShowBeforeUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": null,
        "port": null,
        "protocol": null
      }
    }
expectedStatus:
  - AttributesUpdated
`)

var testMockMLBV1ListenersShowAfterUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": null,
        "port": null,
        "protocol": null,
        "current": null,
        "staged": {
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedBeforeApply
newStatus: ConfigurationsApplied
`)

var testMockMLBV1ListenersShowAfterApplyConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https",
        "current": {
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        },
        "staged": null
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1ListenersShowBeforeCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https"
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1ListenersShowAfterCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "UPDATE_STAGED",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https",
        "current": {
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        },
        "staged": {
          "ip_address": null,
          "port": 80,
          "protocol": "http"
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1ListenersShowAfterDelete = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "DELETE_STAGED",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https",
        "current": {
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        },
        "staged": null
      }
    }
expectedStatus:
  - Deleted
`)

var testMockMLBV1ListenersUpdateAttributes = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"listener":{"description":"description-update","name":"listener-update","tags":{"key-update":"value-update"}}}
response:
  code: 200
  body: >
    {
      "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": null,
        "port": null,
        "protocol": null
      }
    }
expectedStatus:
  - Created
newStatus: AttributesUpdated
`)

var testMockMLBV1ListenersCreateConfigurations = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"listener":{"port":80,"protocol":"http"}}
response:
  code: 200
  body: >
    {
      "listener": {
        "ip_address": null,
        "port": 80,
        "protocol": "http"
      }
    }
expectedStatus:
  - ConfigurationsApplied
newStatus: ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1ListenersUpdateConfigurations = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"listener":{"port":443,"protocol":"https"}}
response:
  code: 200
  body: >
    {
      "listener": {
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https"
      }
    }
expectedStatus:
  - AttributesUpdated
newStatus: ConfigurationsUpdatedBeforeApply
`)

var testMockMLBV1ListenersDelete = fmt.Sprintf(`
request:
  method: DELETE
response:
  code: 204
expectedStatus:
  - Created
  - ConfigurationsUpdatedAfterApply
newStatus: Deleted
`)
