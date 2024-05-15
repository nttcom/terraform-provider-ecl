package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1RuleDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListNameQuery)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListDescriptionQuery)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListConfigurationStatusQuery)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListOperationStatusQuery)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListPriorityQuery)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListTargetGroupIDQuery)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListPolicyIDQuery)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListLoadBalancerIDQuery)
	mc.Register(t, "rules", "/v1.0/rules", testMockMLBV1RulesListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1RuleDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleDataSourceQueryConfigurationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleDataSourceQueryOperationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleDataSourceQueryPriority,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleDataSourceQueryTargetGroupID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleDataSourceQueryPolicyID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleDataSourceQueryLoadBalancerID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
			{
				Config: testAccMLBV1RuleDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "name", "rule"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "policy_id", "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "priority", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "target_group_id", "29527a3c-9e5d-48b7-868f-6442c7d21a95"),
					resource.TestCheckResourceAttr("data.ecl_mlb_rule_v1.rule_1", "conditions.0.path_patterns.0", "^/statics/"),
				),
			},
		},
	})
}

var testAccMLBV1RuleDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  name = "rule"
}
`)

var testMockMLBV1RulesListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - rule
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)

var testAccMLBV1RuleDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  description = "description"
}
`)

var testMockMLBV1RulesListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)

var testAccMLBV1RuleDataSourceQueryConfigurationStatus = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  configuration_status = "ACTIVE"
}
`)

var testMockMLBV1RulesListConfigurationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    configuration_status:
      - ACTIVE
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)

var testAccMLBV1RuleDataSourceQueryOperationStatus = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  operation_status = "COMPLETE"
}
`)

var testMockMLBV1RulesListOperationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    operation_status:
      - COMPLETE
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)

var testAccMLBV1RuleDataSourceQueryPriority = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  priority = "1"
}
`)

var testMockMLBV1RulesListPriorityQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    priority:
      - 1
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)

var testAccMLBV1RuleDataSourceQueryTargetGroupID = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  target_group_id = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
}
`)

var testMockMLBV1RulesListTargetGroupIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    target_group_id:
      - 29527a3c-9e5d-48b7-868f-6442c7d21a95
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)

var testAccMLBV1RuleDataSourceQueryPolicyID = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  policy_id = "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"
}
`)

var testMockMLBV1RulesListPolicyIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    policy_id:
      - fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)

var testAccMLBV1RuleDataSourceQueryLoadBalancerID = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1RulesListLoadBalancerIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 67fea379-cff0-4191-9175-de7d6941a040
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)

var testAccMLBV1RuleDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_rule_v1" "rule_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1RulesListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
response:
  code: 200
  body: >
    {
      "rules": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "rule",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "priority": 1,
          "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
          "conditions": {
            "path_patterns": [
              "^/statics/"
            ]
          }
        }
      ]
    }
`)
