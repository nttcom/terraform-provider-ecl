package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1OperationDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "operations", "/v1.0/operations", testMockMLBV1OperationsListResourceIDQuery)
	mc.Register(t, "operations", "/v1.0/operations", testMockMLBV1OperationsListResourceTypeQuery)
	mc.Register(t, "operations", "/v1.0/operations", testMockMLBV1OperationsListStatusQuery)
	mc.Register(t, "operations", "/v1.0/operations", testMockMLBV1OperationsListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1OperationDataSourceQueryResourceID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "resource_id", "4d5215ed-38bb-48ed-879a-fdb9ca58522f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "resource_type", "ECL::ManagedLoadBalancer::LoadBalancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "request_id", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "request_types.0", "Action::apply-configurations"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "reception_datetime", "2019-08-24 14:15:22"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "commit_datetime", "2019-08-24 14:30:44"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "warning", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "error", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1OperationDataSourceQueryResourceType,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "resource_id", "4d5215ed-38bb-48ed-879a-fdb9ca58522f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "resource_type", "ECL::ManagedLoadBalancer::LoadBalancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "request_id", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "request_types.0", "Action::apply-configurations"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "reception_datetime", "2019-08-24 14:15:22"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "commit_datetime", "2019-08-24 14:30:44"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "warning", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "error", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1OperationDataSourceQueryStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "resource_id", "4d5215ed-38bb-48ed-879a-fdb9ca58522f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "resource_type", "ECL::ManagedLoadBalancer::LoadBalancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "request_id", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "request_types.0", "Action::apply-configurations"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "reception_datetime", "2019-08-24 14:15:22"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "commit_datetime", "2019-08-24 14:30:44"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "warning", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "error", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
			{
				Config: testAccMLBV1OperationDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "resource_id", "4d5215ed-38bb-48ed-879a-fdb9ca58522f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "resource_type", "ECL::ManagedLoadBalancer::LoadBalancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "request_id", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "request_types.0", "Action::apply-configurations"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "reception_datetime", "2019-08-24 14:15:22"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "commit_datetime", "2019-08-24 14:30:44"),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "warning", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "error", ""),
					resource.TestCheckResourceAttr("data.ecl_mlb_operation_v1.operation_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
				),
			},
		},
	})
}

var testAccMLBV1OperationDataSourceQueryResourceID = fmt.Sprintf(`
data "ecl_mlb_operation_v1" "operation_1" {
  resource_id = "4d5215ed-38bb-48ed-879a-fdb9ca58522f"
}
`)

var testMockMLBV1OperationsListResourceIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    resource_id:
      - 4d5215ed-38bb-48ed-879a-fdb9ca58522f
response:
  code: 200
  body: >
    {
      "operations": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "resource_id": "4d5215ed-38bb-48ed-879a-fdb9ca58522f",
          "resource_type": "ECL::ManagedLoadBalancer::LoadBalancer",
          "request_id": "",
          "request_types": [
            "Action::apply-configurations"
          ],
          "status": "COMPLETE",
          "reception_datetime": "2019-08-24 14:15:22",
          "commit_datetime": "2019-08-24 14:30:44",
          "warning": "",
          "error": "",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0"
        }
      ]
    }
`)

var testAccMLBV1OperationDataSourceQueryResourceType = fmt.Sprintf(`
data "ecl_mlb_operation_v1" "operation_1" {
  resource_type = "ECL::ManagedLoadBalancer::LoadBalancer"
}
`)

var testMockMLBV1OperationsListResourceTypeQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    resource_type:
      - ECL::ManagedLoadBalancer::LoadBalancer
response:
  code: 200
  body: >
    {
      "operations": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "resource_id": "4d5215ed-38bb-48ed-879a-fdb9ca58522f",
          "resource_type": "ECL::ManagedLoadBalancer::LoadBalancer",
          "request_id": "",
          "request_types": [
            "Action::apply-configurations"
          ],
          "status": "COMPLETE",
          "reception_datetime": "2019-08-24 14:15:22",
          "commit_datetime": "2019-08-24 14:30:44",
          "warning": "",
          "error": "",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0"
        }
      ]
    }
`)

var testAccMLBV1OperationDataSourceQueryStatus = fmt.Sprintf(`
data "ecl_mlb_operation_v1" "operation_1" {
  status = "COMPLETE"
}
`)

var testMockMLBV1OperationsListStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    status:
      - COMPLETE
response:
  code: 200
  body: >
    {
      "operations": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "resource_id": "4d5215ed-38bb-48ed-879a-fdb9ca58522f",
          "resource_type": "ECL::ManagedLoadBalancer::LoadBalancer",
          "request_id": "",
          "request_types": [
            "Action::apply-configurations"
          ],
          "status": "COMPLETE",
          "reception_datetime": "2019-08-24 14:15:22",
          "commit_datetime": "2019-08-24 14:30:44",
          "warning": "",
          "error": "",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0"
        }
      ]
    }
`)

var testAccMLBV1OperationDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_operation_v1" "operation_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1OperationsListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
response:
  code: 200
  body: >
    {
      "operations": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "resource_id": "4d5215ed-38bb-48ed-879a-fdb9ca58522f",
          "resource_type": "ECL::ManagedLoadBalancer::LoadBalancer",
          "request_id": "",
          "request_types": [
            "Action::apply-configurations"
          ],
          "status": "COMPLETE",
          "reception_datetime": "2019-08-24 14:15:22",
          "commit_datetime": "2019-08-24 14:30:44",
          "warning": "",
          "error": "",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0"
        }
      ]
    }
`)
