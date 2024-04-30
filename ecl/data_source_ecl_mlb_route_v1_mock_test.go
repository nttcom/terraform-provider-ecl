package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1RouteDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListNameQuery)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListDescriptionQuery)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListConfigurationStatusQuery)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListOperationStatusQuery)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListDestinationCidrQuery)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListNextHopIPAddressQuery)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListLoadBalancerIDQuery)
	mc.Register(t, "routes", "/v1.0/routes", testMockMLBV1RoutesListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1RouteDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "name", "route"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "next_hop_ip_address", "192.168.0.254"),
				),
			},
			{
				Config: testAccMLBV1RouteDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "name", "route"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "next_hop_ip_address", "192.168.0.254"),
				),
			},
			{
				Config: testAccMLBV1RouteDataSourceQueryConfigurationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "name", "route"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "next_hop_ip_address", "192.168.0.254"),
				),
			},
			{
				Config: testAccMLBV1RouteDataSourceQueryOperationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "name", "route"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "next_hop_ip_address", "192.168.0.254"),
				),
			},
			{
				Config: testAccMLBV1RouteDataSourceQueryDestinationCidr,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "name", "route"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "next_hop_ip_address", "192.168.0.254"),
				),
			},
			{
				Config: testAccMLBV1RouteDataSourceQueryNextHopIPAddress,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "name", "route"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "next_hop_ip_address", "192.168.0.254"),
				),
			},
			{
				Config: testAccMLBV1RouteDataSourceQueryLoadBalancerID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "name", "route"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "next_hop_ip_address", "192.168.0.254"),
				),
			},
			{
				Config: testAccMLBV1RouteDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "name", "route"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "destination_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_route_v1.route_1", "next_hop_ip_address", "192.168.0.254"),
				),
			},
		},
	})
}

var testAccMLBV1RouteDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_route_v1" "route_1" {
  name = "route"
}
`)

var testMockMLBV1RoutesListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - route
response:
  code: 200
  body: >
    {
      "routes": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "route",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "destination_cidr": "172.16.0.0/24",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "next_hop_ip_address": "192.168.0.254"
        }
      ]
    }
`)

var testAccMLBV1RouteDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_route_v1" "route_1" {
  description = "description"
}
`)

var testMockMLBV1RoutesListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "routes": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "route",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "destination_cidr": "172.16.0.0/24",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "next_hop_ip_address": "192.168.0.254"
        }
      ]
    }
`)

var testAccMLBV1RouteDataSourceQueryConfigurationStatus = fmt.Sprintf(`
data "ecl_mlb_route_v1" "route_1" {
  configuration_status = "ACTIVE"
}
`)

var testMockMLBV1RoutesListConfigurationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    configuration_status:
      - ACTIVE
response:
  code: 200
  body: >
    {
      "routes": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "route",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "destination_cidr": "172.16.0.0/24",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "next_hop_ip_address": "192.168.0.254"
        }
      ]
    }
`)

var testAccMLBV1RouteDataSourceQueryOperationStatus = fmt.Sprintf(`
data "ecl_mlb_route_v1" "route_1" {
  operation_status = "COMPLETE"
}
`)

var testMockMLBV1RoutesListOperationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    operation_status:
      - COMPLETE
response:
  code: 200
  body: >
    {
      "routes": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "route",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "destination_cidr": "172.16.0.0/24",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "next_hop_ip_address": "192.168.0.254"
        }
      ]
    }
`)

var testAccMLBV1RouteDataSourceQueryDestinationCidr = fmt.Sprintf(`
data "ecl_mlb_route_v1" "route_1" {
  destination_cidr = "172.16.0.0/24"
}
`)

var testMockMLBV1RoutesListDestinationCidrQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    destination_cidr:
      - 172.16.0.0/24
response:
  code: 200
  body: >
    {
      "routes": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "route",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "destination_cidr": "172.16.0.0/24",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "next_hop_ip_address": "192.168.0.254"
        }
      ]
    }
`)

var testAccMLBV1RouteDataSourceQueryNextHopIPAddress = fmt.Sprintf(`
data "ecl_mlb_route_v1" "route_1" {
  next_hop_ip_address = "192.168.0.254"
}
`)

var testMockMLBV1RoutesListNextHopIPAddressQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    next_hop_ip_address:
      - 192.168.0.254
response:
  code: 200
  body: >
    {
      "routes": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "route",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "destination_cidr": "172.16.0.0/24",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "next_hop_ip_address": "192.168.0.254"
        }
      ]
    }
`)

var testAccMLBV1RouteDataSourceQueryLoadBalancerID = fmt.Sprintf(`
data "ecl_mlb_route_v1" "route_1" {
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1RoutesListLoadBalancerIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 67fea379-cff0-4191-9175-de7d6941a040
response:
  code: 200
  body: >
    {
      "routes": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "route",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "destination_cidr": "172.16.0.0/24",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "next_hop_ip_address": "192.168.0.254"
        }
      ]
    }
`)

var testAccMLBV1RouteDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_route_v1" "route_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1RoutesListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
response:
  code: 200
  body: >
    {
      "routes": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "route",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "destination_cidr": "172.16.0.0/24",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "next_hop_ip_address": "192.168.0.254"
        }
      ]
    }
`)
