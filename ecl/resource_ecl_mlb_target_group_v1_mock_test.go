package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1TargetGroupResource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsCreate)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowAfterCreate)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsUpdateAttributes)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowBeforeUpdateConfigurations)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1TargetGroupsUpdateConfigurations)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowAfterUpdateConfigurations)
	// Staged configurations of the load balancer and related resources are applied here
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowAfterApplyConfigurations)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowBeforeCreateConfigurations)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1TargetGroupsCreateConfigurations)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowAfterCreateConfigurations)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsDelete)
	mc.Register(t, "target_groups", "/v1.0/target_groups/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1TargetGroupsShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1TargetGroup,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "name", "target_group"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "description", "description"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "tags.key", "value"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.ip_address", "192.168.0.7"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.port", "80"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.weight", "1"),
				),
			},
			{
				Config: testAccMLBV1TargetGroupUpdateBeforeApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "name", "target_group-update"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.#", "2"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.ip_address", "192.168.0.7"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.port", "80"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.weight", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.1.ip_address", "192.168.0.8"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.1.port", "80"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.1.weight", "1"),
				),
			},
			{
				Config: testAccMLBV1TargetGroupUpdateAfterApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "name", "target_group-update"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.ip_address", "192.168.0.8"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.port", "80"),
					resource.TestCheckResourceAttr("ecl_mlb_target_group_v1.target_group", "members.0.weight", "1"),
				),
			},
		},
	})
}

var testAccMLBV1TargetGroup = fmt.Sprintf(`
resource "ecl_mlb_target_group_v1" "target_group" {
  name = "target_group"
  description = "description"
  tags = {
    key = "value"
  }
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
  members {
	ip_address = "192.168.0.7"
	port = 80
	weight = 1
  }
}
`)

var testAccMLBV1TargetGroupUpdateBeforeApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_target_group_v1" "target_group" {
  name = "target_group-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
  members {
	ip_address = "192.168.0.7"
	port = 80
	weight = 1
  }
  members {
	ip_address = "192.168.0.8"
	port = 80
	weight = 1
  }
}
`)

var testAccMLBV1TargetGroupUpdateAfterApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_target_group_v1" "target_group" {
  name = "target_group-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
  members {
	ip_address = "192.168.0.8"
	port = 80
	weight = 1
  }
}
`)

var testMockMLBV1TargetGroupsCreate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"target_group":{"description":"description","load_balancer_id":"67fea379-cff0-4191-9175-de7d6941a040","members":[{"ip_address":"192.168.0.7","port":80,"weight":1}],"name":"target_group","tags":{"key":"value"}}}
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": null
      }
    }
newStatus: Created
`)

var testMockMLBV1TargetGroupsShowAfterCreate = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": null,
        "current": null,
        "staged": {
          "members": [
            {
              "ip_address": "192.168.0.7",
              "port": 80,
              "weight": 1
            }
          ]
        }
      }
    }
expectedStatus:
  - Created
`)

var testMockMLBV1TargetGroupsShowBeforeUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": null
      }
    }
expectedStatus:
  - AttributesUpdated
`)

var testMockMLBV1TargetGroupsShowAfterUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": null,
        "current": null,
        "staged": {
          "members": [
            {
              "ip_address": "192.168.0.7",
              "port": 80,
              "weight": 1
            },
            {
              "ip_address": "192.168.0.8",
              "port": 80,
              "weight": 1
            }
          ]
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedBeforeApply
newStatus: ConfigurationsApplied
`)

var testMockMLBV1TargetGroupsShowAfterApplyConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": [
          {
            "ip_address": "192.168.0.7",
            "port": 80,
            "weight": 1
          },
          {
            "ip_address": "192.168.0.8",
            "port": 80,
            "weight": 1
          }
        ],
        "current": {
          "members": [
            {
              "ip_address": "192.168.0.7",
              "port": 80,
              "weight": 1
            },
            {
              "ip_address": "192.168.0.8",
              "port": 80,
              "weight": 1
            }
          ]
        },
        "staged": null
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1TargetGroupsShowBeforeCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": [
          {
            "ip_address": "192.168.0.7",
            "port": 80,
            "weight": 1
          },
          {
            "ip_address": "192.168.0.8",
            "port": 80,
            "weight": 1
          }
        ]
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1TargetGroupsShowAfterCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "UPDATE_STAGED",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": [
          {
            "ip_address": "192.168.0.7",
            "port": 80,
            "weight": 1
          },
          {
            "ip_address": "192.168.0.8",
            "port": 80,
            "weight": 1
          }
        ],
        "current": {
          "members": [
            {
              "ip_address": "192.168.0.7",
              "port": 80,
              "weight": 1
            },
            {
              "ip_address": "192.168.0.8",
              "port": 80,
              "weight": 1
            }
          ]
        },
        "staged": {
          "members": [
            {
              "ip_address": "192.168.0.8",
              "port": 80,
              "weight": 1
            }
          ]
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1TargetGroupsShowAfterDelete = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "DELETE_STAGED",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": [
          {
            "ip_address": "192.168.0.7",
            "port": 80,
            "weight": 1
          },
          {
            "ip_address": "192.168.0.8",
            "port": 80,
            "weight": 1
          }
        ],
        "current": {
          "members": [
            {
              "ip_address": "192.168.0.7",
              "port": 80,
              "weight": 1
            },
            {
              "ip_address": "192.168.0.8",
              "port": 80,
              "weight": 1
            }
          ]
        },
        "staged": null
      }
    }
expectedStatus:
  - Deleted
`)

var testMockMLBV1TargetGroupsUpdateAttributes = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"target_group":{"description":"description-update","name":"target_group-update","tags":{"key-update":"value-update"}}}
response:
  code: 200
  body: >
    {
      "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": null
      }
    }
expectedStatus:
  - Created
newStatus: AttributesUpdated
`)

var testMockMLBV1TargetGroupsCreateConfigurations = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"target_group":{"members":[{"ip_address":"192.168.0.8","port":80,"weight":1}]}}
response:
  code: 200
  body: >
    {
      "target_group": {
        "members": [
          {
            "ip_address": "192.168.0.8",
            "port": 80,
            "weight": 1
          }
        ]
      }
    }
expectedStatus:
  - ConfigurationsApplied
newStatus: ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1TargetGroupsUpdateConfigurations = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"target_group":{"members":[{"ip_address":"192.168.0.7","port":80,"weight":1},{"ip_address":"192.168.0.8","port":80,"weight":1}]}}
response:
  code: 200
  body: >
    {
      "target_group": {
        "members": [
          {
            "ip_address": "192.168.0.7",
            "port": 80,
            "weight": 1
          },
          {
            "ip_address": "192.168.0.8",
            "port": 80,
            "weight": 1
          }
        ]
      }
    }
expectedStatus:
  - AttributesUpdated
newStatus: ConfigurationsUpdatedBeforeApply
`)

var testMockMLBV1TargetGroupsDelete = fmt.Sprintf(`
request:
  method: DELETE
response:
  code: 204
expectedStatus:
  - Created
  - ConfigurationsUpdatedAfterApply
newStatus: Deleted
`)
