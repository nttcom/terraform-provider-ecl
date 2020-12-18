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
	mc.Register(t, "qos_options", "/v2.0/qos_optionss", testMockNetworkV2QosOptionsListNameQuery)

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
						"data.ecl_network_qos_options_v2.qos_options_1", "name", "20M-GA-AZURE"),
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
	mc.Register(t, "qos_options", "/v2.0/qos_optionss", testMockNetworkV2QosOptionsListNameQuery)
	mc.Register(t, "qos_options", "/v2.0/qos_optionss", testMockNetworkV2QosOptionsListIDQuery)
	mc.Register(t, "qos_options", "/v2.0/qos_optionss", testMockNetworkV2QosOptionsListDescriptionQuery)

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
        - 20M-GA-AZURE
response: 
    code: 200
    body: >
        {
            "aws_service_id"	  : null,
            "azure_service_id"	  : "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
            "bandwidth"			  : "20",
            "description"		  : "20M-guarantee-menu-for-azure",
            "fic_service_id"	  : null,
            "gcp_service_id"	  : null,
            "id"				  : "a6b91294-8870-4f2c-b9e9-a899acada723",
            "interdc_service_id"  : null,
            "internet_service_id" : null,
            "name"				  : "20M-GA-AZURE",
            "qos_type"			  : "guarantee",
            "service_type"		  : "azure",
            "status"			  : "ACTIVE",
            "vpn_service_id"	  : null
        }
`

var testMockNetworkV2QosOptionsListIDQuery = `
request:
    method: GET
    query:
      id:
        - d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9
response: 
    code: 200
    body: >
        {
            "aws_service_id"	  : null,
            "azure_service_id"	  : "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
            "bandwidth"			  : "20",
            "description"		  : "20M-guarantee-menu-for-azure",
            "fic_service_id"	  : null,
            "gcp_service_id"	  : null,
            "id"				  : "a6b91294-8870-4f2c-b9e9-a899acada723",
            "interdc_service_id"  : null,
            "internet_service_id" : null,
            "name"				  : "20M-GA-AZURE",
            "qos_type"			  : "guarantee",
            "service_type"		  : "azure",
            "status"			  : "ACTIVE",
            "vpn_service_id"	  : null
        }
`

var testMockNetworkV2QosOptionsListDescriptionQuery = `
request:
    method: GET
    query:
      description:
        - "20M-guarantee-menu-for-azure"
response: 
    code: 200
    body: >
        {
            "aws_service_id"	  : null,
            "azure_service_id"	  : "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
            "bandwidth"			  : "20",
            "description"		  : "20M-guarantee-menu-for-azure",
            "fic_service_id"	  : null,
            "gcp_service_id"	  : null,
            "id"				  : "a6b91294-8870-4f2c-b9e9-a899acada723",
            "interdc_service_id"  : null,
            "internet_service_id" : null,
            "name"				  : "20M-GA-AZURE",
            "qos_type"			  : "guarantee",
            "service_type"		  : "azure",
            "status"			  : "ACTIVE",
            "vpn_service_id"	  : null
        }
`
