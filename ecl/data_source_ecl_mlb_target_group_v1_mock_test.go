package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1TargetGroupDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsListNameQuery)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsListDescriptionQuery)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsListConfigurationStatusQuery)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsListOperationStatusQuery)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsListLoadBalancerIDQuery)
	mc.Register(t, "target_groups", "/v1.0/target_groups", testMockMLBV1TargetGroupsListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1TargetGroupDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "name", "target_group"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.ip_address", "192.168.0.7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.weight", "1"),
				),
			},
			{
				Config: testAccMLBV1TargetGroupDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "name", "target_group"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.ip_address", "192.168.0.7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.weight", "1"),
				),
			},
			{
				Config: testAccMLBV1TargetGroupDataSourceQueryConfigurationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "name", "target_group"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.ip_address", "192.168.0.7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.weight", "1"),
				),
			},
			{
				Config: testAccMLBV1TargetGroupDataSourceQueryOperationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "name", "target_group"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.ip_address", "192.168.0.7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.weight", "1"),
				),
			},
			{
				Config: testAccMLBV1TargetGroupDataSourceQueryLoadBalancerID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "name", "target_group"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.ip_address", "192.168.0.7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.weight", "1"),
				),
			},
			{
				Config: testAccMLBV1TargetGroupDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "name", "target_group"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.ip_address", "192.168.0.7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_target_group_v1.target_group_1", "members.0.weight", "1"),
				),
			},
		},
	})
}

var testAccMLBV1TargetGroupDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_target_group_v1" "target_group_1" {
  name = "target_group"
}
`)

var testMockMLBV1TargetGroupsListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - target_group
response:
  code: 200
  body: >
    {
      "target_groups": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "target_group",
          "description": "description",
          "tags": {
            "key": "value"
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
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1TargetGroupDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_target_group_v1" "target_group_1" {
  description = "description"
}
`)

var testMockMLBV1TargetGroupsListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "target_groups": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "target_group",
          "description": "description",
          "tags": {
            "key": "value"
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
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1TargetGroupDataSourceQueryConfigurationStatus = fmt.Sprintf(`
data "ecl_mlb_target_group_v1" "target_group_1" {
  configuration_status = "ACTIVE"
}
`)

var testMockMLBV1TargetGroupsListConfigurationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    configuration_status:
      - ACTIVE
response:
  code: 200
  body: >
    {
      "target_groups": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "target_group",
          "description": "description",
          "tags": {
            "key": "value"
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
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1TargetGroupDataSourceQueryOperationStatus = fmt.Sprintf(`
data "ecl_mlb_target_group_v1" "target_group_1" {
  operation_status = "COMPLETE"
}
`)

var testMockMLBV1TargetGroupsListOperationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    operation_status:
      - COMPLETE
response:
  code: 200
  body: >
    {
      "target_groups": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "target_group",
          "description": "description",
          "tags": {
            "key": "value"
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
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1TargetGroupDataSourceQueryLoadBalancerID = fmt.Sprintf(`
data "ecl_mlb_target_group_v1" "target_group_1" {
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1TargetGroupsListLoadBalancerIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 67fea379-cff0-4191-9175-de7d6941a040
response:
  code: 200
  body: >
    {
      "target_groups": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "target_group",
          "description": "description",
          "tags": {
            "key": "value"
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
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1TargetGroupDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_target_group_v1" "target_group_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1TargetGroupsListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
response:
  code: 200
  body: >
    {
      "target_groups": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "target_group",
          "description": "description",
          "tags": {
            "key": "value"
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
            }
          ]
        }
      ]
    }
`)
