package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccNetworkV2QosOptionsDataSource_basic(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "qos_options", "/v2.0/qos_options", testMockNetworkV2QosOptionsListNameQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2QosOptionsDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2QosOptionsDataSourceID("data.ecl_network_qos_options_v2.qos_options_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_qos_options_v2.qos_options_1", "name", "10Mbps-BestEffort"),
				),
			},
		},
	})
}

func TestMockedAccNetworkV2QosOptionsDataSource_queries(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "qos_options", "/v2.0/qos_options", testMockNetworkV2QosOptionsListNameQuery)
	mc.Register(t, "qos_options", "/v2.0/qos_options", testMockNetworkV2QosOptionsListIDQuery)
	mc.Register(t, "qos_options", "/v2.0/qos_options", testMockNetworkV2QosOptionsListDescriptionQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetworkV2QosOptionsDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2QosOptionsDataSourceID("data.ecl_network_qos_options_v2.qos_options_1"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2QosOptionsDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2QosOptionsDataSourceID("data.ecl_network_qos_options_v2.qos_options_2"),
				),
			},
			resource.TestStep{
				Config: testAccNetworkV2QosOptionsDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2QosOptionsDataSourceID("data.ecl_network_qos_options_v2.qos_options_1"),
				),
			},
		},
	})
}

var testMockNetworkV2QosOptionsListNameQuery = `
request:
    method: GET
    query:
      name:
        - 10Mbps-BestEffort
response: 
    code: 200
    body: >
        {
            "qos_options": [
                {
                    "aws_service_id": null,
                    "azure_service_id": null,
                    "bandwidth": "10",
                    "description": "10M-besteffort-menu",
                    "fic_service_id": null,
                    "gcp_service_id": null,
                    "id": "a6b91294-8870-4f2c-b9e9-a899acada723",
                    "interdc_service_id": null,
                    "internet_service_id": null,
                    "name": "10Mbps-BestEffort",
                    "qos_type": "besteffort",
                    "service_type": "internet",
                    "status": "ACTIVE",
                    "vpn_service_id": null
                }
            ]
        }
`

var testMockNetworkV2QosOptionsListIDQuery = `
request:
    method: GET
    query:
      id:
        - a6b91294-8870-4f2c-b9e9-a899acada723
response: 
    code: 200
    body: >
        {
            "qos_options": [
                {
                    "aws_service_id": null,
                    "azure_service_id": null,
                    "bandwidth": "10",
                    "description": "10M-besteffort-menu",
                    "fic_service_id": null,
                    "gcp_service_id": null,
                    "id": "a6b91294-8870-4f2c-b9e9-a899acada723",
                    "interdc_service_id": null,
                    "internet_service_id": null,
                    "name": "10Mbps-BestEffort",
                    "qos_type": "besteffort",
                    "service_type": "internet",
                    "status": "ACTIVE",
                    "vpn_service_id": null
                }
            ]
        }
`

var testMockNetworkV2QosOptionsListDescriptionQuery = `
request:
    method: GET
    query:
      description:
        - "10m-besteffort-menu"
response: 
    code: 200
    body: >
        {
            "qos_options": [
                {
                    "aws_service_id": null,
                    "azure_service_id": null,
                    "bandwidth": "10",
                    "description": "10m-besteffort-menu",
                    "fic_service_id": null,
                    "gcp_service_id": null,
                    "id": "a6b91294-8870-4f2c-b9e9-a899acada723",
                    "interdc_service_id": null,
                    "internet_service_id": null,
                    "name": "10mbps-besteffort",
                    "qos_type": "besteffort",
                    "service_type": "internet",
                    "status": "active",
                    "vpn_service_id": null
                }
            ]
        }
`
