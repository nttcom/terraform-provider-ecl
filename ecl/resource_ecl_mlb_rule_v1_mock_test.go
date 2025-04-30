package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1RuleResource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesCreate)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowAfterCreate)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesUpdateAttributes)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowBeforeUpdateConfigurations)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1RulesUpdateConfigurations)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowAfterUpdateConfigurations)
	// Staged configurations of the load balancer and related resources are applied here
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowAfterApplyConfigurations)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowBeforeCreateConfigurations)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08/staged", testMockMLBV1RulesCreateConfigurations)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowAfterCreateConfigurations)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesDelete)
	mc.Register(t, "rules", "/v1.0/rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", testMockMLBV1RulesShowAfterDelete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1Rule,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "name", "rule"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "description", "description"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "tags.key", "value"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "priority", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "backup_target_group_id", "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.0.path_patterns.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleUpdateBeforeApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "name", "rule-update"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "priority", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "backup_target_group_id", "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.0.path_patterns.#", "2"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.0.path_patterns.0", "^/statics/"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.0.path_patterns.1", "^/assets/"),
				),
			},
			{
				Config: testAccMLBV1RuleUpdateAfterApplyConfigurations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "name", "rule-update"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "description", "description-update"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "tags.key-update", "value-update"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "priority", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "backup_target_group_id", "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.0.path_patterns.#", "1"),
					resource.TestCheckResourceAttr("ecl_mlb_rule_v1.rule", "conditions.0.path_patterns.0", "^/assets/"),
				),
			},
		},
	})
}

var testAccMLBV1Rule = fmt.Sprintf(`
resource "ecl_mlb_rule_v1" "rule" {
  name = "rule"
  description = "description"
  tags = {
    key = "value"
  }
  priority = 1
  target_group_id = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
  backup_target_group_id = "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"
  policy_id = "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"
  conditions {
    path_patterns = ["^/statics/"]
  }
}
`)

var testAccMLBV1RuleUpdateBeforeApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_rule_v1" "rule" {
  name = "rule-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  priority = 1
  target_group_id = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
  backup_target_group_id = "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"
  policy_id = "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"
  conditions {
    path_patterns = ["^/statics/", "^/assets/"]
  }
}
`)

var testAccMLBV1RuleUpdateAfterApplyConfigurations = fmt.Sprintf(`
resource "ecl_mlb_rule_v1" "rule" {
  name = "rule-update"
  description = "description-update"
  tags = {
    key-update = "value-update"
  }
  priority = 1
  target_group_id = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
  backup_target_group_id = "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"
  policy_id = "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"
  conditions {
    path_patterns = ["^/assets/"]
  }
}
`)

var testMockMLBV1RulesCreate = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"rule":{"backup_target_group_id":"dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52","conditions":{"path_patterns":["^/statics/"]},"description":"description","name":"rule","policy_id":"fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4","priority":1,"tags":{"key":"value"},"target_group_id":"29527a3c-9e5d-48b7-868f-6442c7d21a95"}}
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": null,
        "target_group_id": null,
        "backup_target_group_id": null,
        "conditions": null
      }
    }
newStatus: Created
`)

var testMockMLBV1RulesShowAfterCreate = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule",
        "description": "description",
        "tags": {
          "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": null,
        "target_group_id": null,
        "backup_target_group_id": null,
        "conditions": null,
        "current": null,
        "staged": {
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
          "conditions": {
            "path_patterns": ["^/statics/"]
          }
        }
      }
    }
expectedStatus:
  - Created
`)

var testMockMLBV1RulesShowBeforeUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": null,
        "target_group_id": null,
        "backup_target_group_id": null,
        "conditions": null
      }
    }
expectedStatus:
  - AttributesUpdated
`)

var testMockMLBV1RulesShowAfterUpdateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": null,
        "target_group_id": null,
        "backup_target_group_id": null,
        "conditions": null,
        "current": null,
        "staged": {
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
          "conditions": {
            "path_patterns": ["^/statics/", "^/assets/"]
          }
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedBeforeApply
newStatus: ConfigurationsApplied
`)

var testMockMLBV1RulesShowAfterApplyConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
          "path_patterns": ["^/statics/", "^/assets/"]
        },
        "current": {
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
          "conditions": {
            "path_patterns": ["^/statics/", "^/assets/"]
          }
        },
        "staged": null
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1RulesShowBeforeCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
          "path_patterns": ["^/statics/", "^/assets/"]
        }
      }
    }
expectedStatus:
  - ConfigurationsApplied
`)

var testMockMLBV1RulesShowAfterCreateConfigurations = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "UPDATE_STAGED",
        "operation_status": "COMPLETE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
          "path_patterns": ["^/statics/", "^/assets/"]
        },
        "current": {
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
          "conditions": {
            "path_patterns": ["^/statics/", "^/assets/"]
          }
        },
        "staged": {
          "priority": null,
          "target_group_id": null,
          "backup_target_group_id": null,
          "conditions": {
            "path_patterns": ["^/assets/"]
          }
        }
      }
    }
expectedStatus:
  - ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1RulesShowAfterDelete = fmt.Sprintf(`
request:
  method: GET
  query:
    changes:
      - true
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "DELETE_STAGED",
        "operation_status": "COMPLETE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
          "path_patterns": ["^/statics/", "^/assets/"]
        },
        "current": {
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
          "conditions": {
            "path_patterns": ["^/statics/", "^/assets/"]
          }
        },
        "staged": null
      }
    }
expectedStatus:
  - Deleted
`)

var testMockMLBV1RulesUpdateAttributes = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"rule":{"description":"description-update","name":"rule-update","tags":{"key-update":"value-update"}}}
response:
  code: 200
  body: >
    {
      "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule-update",
        "description": "description-update",
        "tags": {
          "key-update": "value-update"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": null,
        "target_group_id": null,
        "backup_target_group_id": null,
        "conditions": null
      }
    }
expectedStatus:
  - Created
newStatus: AttributesUpdated
`)

var testMockMLBV1RulesCreateConfigurations = fmt.Sprintf(`
request:
  method: POST
  body: >
    {"rule":{"conditions":{"path_patterns":["^/assets/"]}}}
response:
  code: 200
  body: >
    {
      "rule": {
        "priority": null,
        "target_group_id": null,
        "backup_target_group_id": null,
        "conditions": {
          "path_patterns": ["^/assets/"]
        }
      }
    }
expectedStatus:
  - ConfigurationsApplied
newStatus: ConfigurationsUpdatedAfterApply
`)

var testMockMLBV1RulesUpdateConfigurations = fmt.Sprintf(`
request:
  method: PATCH
  body: >
    {"rule":{"conditions":{"path_patterns":["^/statics/","^/assets/"]}}}
response:
  code: 200
  body: >
    {
      "rule": {
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
          "path_patterns": ["^/statics/", "^/assets/"]
        }
      }
    }
expectedStatus:
  - AttributesUpdated
newStatus: ConfigurationsUpdatedBeforeApply
`)

var testMockMLBV1RulesDelete = fmt.Sprintf(`
request:
  method: DELETE
response:
  code: 204
expectedStatus:
  - Created
  - ConfigurationsUpdatedAfterApply
newStatus: Deleted
`)
