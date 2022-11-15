/*
Generated by https://github.com/tamac-io/openapi-to-terraform-rb
*/
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

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1ListenerDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_mlb_listener_v1.listener_1", "id"),
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
