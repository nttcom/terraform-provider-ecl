package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccNetworkV2FICGatewayDataSource_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "description", OS_FIC_GW_DESCRIPTION),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "fic_service_id", OS_FIC_SERVICE_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "fic_gateway_id", OS_FIC_GW_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "name", OS_FIC_GW_NAME),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "qos_option_id", OS_FIC_GW_QOS_OPTION_ID),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "status", OS_FIC_GW_STATUS),
					resource.TestCheckResourceAttr(
						"data.ecl_network_fic_gateway_v2.fic_gateway_1", "tenant_id", OS_TENANT_ID),
				),
			},
		},
	})
}

func TestAccNetworkV2FICGatewayDataSource_queries(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceFICServiceID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceID2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceQoSOptionID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFICGateway(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2FICGatewayDataSourceTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_network_fic_gateway_v2.fic_gateway_1", "id"),
				),
			},
		},
	})

}

var testAccNetworkV2FICGatewayDataSourceBasic = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	name = %q
}
`,
	OS_FIC_GW_NAME)

var testAccNetworkV2FICGatewayDataSourceDescription = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	description = %q
}
`,
	OS_FIC_GW_DESCRIPTION)

var testAccNetworkV2FICGatewayDataSourceFICServiceID = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	fic_service_id = %q
}
`,
	OS_FIC_SERVICE_ID)

var testAccNetworkV2FICGatewayDataSourceID = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	fic_gateway_id = %q
}
`,
	OS_FIC_GW_ID)

var testAccNetworkV2FICGatewayDataSourceID2 = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_0" {
	name = %q
}

data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	fic_gateway_id = "${data.ecl_network_fic_gateway_v2.fic_gateway_0.id}"
}
`,
	OS_FIC_GW_NAME)

var testAccNetworkV2FICGatewayDataSourceName = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	name = %q
}
`,
	OS_FIC_GW_NAME)

var testAccNetworkV2FICGatewayDataSourceQoSOptionID = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	qos_option_id = %q
}
`,
	OS_FIC_GW_QOS_OPTION_ID)

var testAccNetworkV2FICGatewayDataSourceStatus = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	status = %q
}
`,
	OS_FIC_GW_STATUS)

var testAccNetworkV2FICGatewayDataSourceTenantID = fmt.Sprintf(`
data "ecl_network_fic_gateway_v2" "fic_gateway_1" {
	tenant_id = %q
}
`,
	OS_TENANT_ID)
