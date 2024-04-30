package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1ListenerDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListNameQuery)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListDescriptionQuery)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListConfigurationStatusQuery)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListOperationStatusQuery)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListIPAddressQuery)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListPortQuery)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListProtocolQuery)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListLoadBalancerIDQuery)
	mc.Register(t, "listeners", "/v1.0/listeners", testMockMLBV1ListenersListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1ListenerDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
			{
				Config: testAccMLBV1ListenerDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
			{
				Config: testAccMLBV1ListenerDataSourceQueryConfigurationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
			{
				Config: testAccMLBV1ListenerDataSourceQueryOperationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
			{
				Config: testAccMLBV1ListenerDataSourceQueryIPAddress,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
			{
				Config: testAccMLBV1ListenerDataSourceQueryPort,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
			{
				Config: testAccMLBV1ListenerDataSourceQueryProtocol,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
			{
				Config: testAccMLBV1ListenerDataSourceQueryLoadBalancerID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
			{
				Config: testAccMLBV1ListenerDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "name", "listener"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "ip_address", "10.0.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "port", "443"),
					resource.TestCheckResourceAttr("data.ecl_mlb_listener_v1.listener_1", "protocol", "https"),
				),
			},
		},
	})
}

var testAccMLBV1ListenerDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  name = "listener"
}
`)

var testMockMLBV1ListenersListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - listener
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)

var testAccMLBV1ListenerDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  description = "description"
}
`)

var testMockMLBV1ListenersListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)

var testAccMLBV1ListenerDataSourceQueryConfigurationStatus = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  configuration_status = "ACTIVE"
}
`)

var testMockMLBV1ListenersListConfigurationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    configuration_status:
      - ACTIVE
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)

var testAccMLBV1ListenerDataSourceQueryOperationStatus = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  operation_status = "COMPLETE"
}
`)

var testMockMLBV1ListenersListOperationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    operation_status:
      - COMPLETE
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)

var testAccMLBV1ListenerDataSourceQueryIPAddress = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  ip_address = "10.0.0.1"
}
`)

var testMockMLBV1ListenersListIPAddressQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    ip_address:
      - 10.0.0.1
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)

var testAccMLBV1ListenerDataSourceQueryPort = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  port = "443"
}
`)

var testMockMLBV1ListenersListPortQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    port:
      - 443
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)

var testAccMLBV1ListenerDataSourceQueryProtocol = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  protocol = "https"
}
`)

var testMockMLBV1ListenersListProtocolQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    protocol:
      - https
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)

var testAccMLBV1ListenerDataSourceQueryLoadBalancerID = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1ListenersListLoadBalancerIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 67fea379-cff0-4191-9175-de7d6941a040
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)

var testAccMLBV1ListenerDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_listener_v1" "listener_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1ListenersListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
response:
  code: 200
  body: >
    {
      "listeners": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "listener",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ip_address": "10.0.0.1",
          "port": 443,
          "protocol": "https"
        }
      ]
    }
`)
