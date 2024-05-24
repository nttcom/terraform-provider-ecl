package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1RouteResource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesCreate)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowAfterCreate)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesUpdateAttributes)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowBeforeUpdateConfigurations)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1RoutesUpdateConfigurations)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowAfterUpdateConfigurations)
	// Staged configurations of the load balancer and related resources are applied here
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowAfterApplyConfigurations)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowBeforeCreateConfigurations)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1RoutesCreateConfigurations)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowAfterCreateConfigurations)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesDelete)
	mc.Register(t, "routes", "/v1.0/routes/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RoutesShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Route,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "name", "route"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "description", "description"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "tags.key", "value"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "next_hop_ip_address", "192.168.0.254"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1RouteUpdateBeforeApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "name", "route-update"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "next_hop_ip_address", "192.168.1.254"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1RouteUpdateAfterApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "name", "route-update"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "next_hop_ip_address", "192.168.0.254"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_route_v1.route", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
		},
	})
}

var testAccMLBV1Route = fmt.Sprintf(`
resource "ecl_mlb_route_v1" "route" {
  name = "route"
  description = "description"
  tags = {
    key = "value"
  }
  destination_cidr = "172.16.0.0/24"
  next_hop_ip_address = "192.168.0.254"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testAccMLBV1RouteUpdateBeforeApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_route_v1" "route" {
  name = "route-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  destination_cidr = "172.16.0.0/24"
  next_hop_ip_address = "192.168.1.254"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testAccMLBV1RouteUpdateAfterApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_route_v1" "route" {
  name = "route-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  destination_cidr = "172.16.0.0/24"
  next_hop_ip_address = "192.168.0.254"
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1RoutesCreate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"route":{"description":"description","destination_cidr":"172.16.0.0/24","load_balancer_id":"67fea379-cff0-4191-9175-de7d6941a040","name":"route","next_hop_ip_address":"192.168.0.254","tags":{"key":"value"}}}
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": null
      }
    }
newStatus: Created
`)

var testMockMLBV1RoutesShowAfterCreate = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": null,
        "current": null,
        "staged": {
          "next_hop_ip_address": "192.168.0.254"
        }
      }
    }
expectedStatus:
  - Created
`)

var testMockMLBV1RoutesShowBeforeUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": null
      }
    }
expectedStatus:
  - AttributesUpdated
`)

var testMockMLBV1RoutesShowAfterUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": null,
        "current": null,
        "staged": {
          "next_hop_ip_address": "192.168.1.254"
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedBeforeApply
newStatus: ConfigurationsApplied
`)

var testMockMLBV1RoutesShowAfterApplyConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": "192.168.1.254",
        "current": {
          "next_hop_ip_address": "192.168.1.254"
        },
        "staged": null
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1RoutesShowBeforeCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": "192.168.1.254"
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1RoutesShowAfterCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "UPDATE_STAGED",
        "operation_status": "COMPLETE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": "192.168.1.254",
        "current": {
          "next_hop_ip_address": "192.168.1.254"
        },
        "staged": {
          "next_hop_ip_address": "192.168.0.254"
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1RoutesShowAfterDelete = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "DELETE_STAGED",
        "operation_status": "COMPLETE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": "192.168.1.254",
        "current": {
          "next_hop_ip_address": "192.168.1.254"
        },
        "staged": null
      }
    }
expectedStatus:
  - Deleted
`)

var testMockMLBV1RoutesUpdateAttributes = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"route":{"description":"description-update","name":"route-update","tags":{"key-update":"value-update"}}}
response:
  code: 200
  body: >
    {
      "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": null
      }
    }
expectedStatus:
  - Created
newStatus: AttributesUpdated
`)

var testMockMLBV1RoutesCreateConfigurations = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"route":{"next_hop_ip_address":"192.168.0.254"}}
response:
  code: 200
  body: >
    {
      "route": {
        "next_hop_ip_address": "192.168.0.254"
      }
    }
expectedStatus:
  - ConfigurationsApplied
newStatus: ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1RoutesUpdateConfigurations = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"route":{"next_hop_ip_address":"192.168.1.254"}}
response:
  code: 200
  body: >
    {
      "route": {
        "next_hop_ip_address": "192.168.1.254"
      }
    }
expectedStatus:
  - AttributesUpdated
newStatus: ConfigurationsUpdatedBeforeApply
`)

var testMockMLBV1RoutesDelete = fmt.Sprintf(`
request:
  method: DELETE
response:
  code: 204
expectedStatus:
  - Created
  - ConfigurationsUpdatedAfterApply
newStatus: Deleted
`)
