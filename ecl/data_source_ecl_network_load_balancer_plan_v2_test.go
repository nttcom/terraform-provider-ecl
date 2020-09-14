package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccNetworkV2LoadBalancerPlanDataSourcePlanName = "Citrix_NetScaler_VPX_12.1-55.18_Standard_Edition_1000Mbps_4CPU-16GB-8IF"

func TestAccNetworkV2LoadBalancerPlanDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerPlanDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerPlanDataSourceID("data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "name", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "description", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "maximum_syslog_servers", "8"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.edition", "Standard"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.size", "1000"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "vendor", "citrix"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "version", "12.1-55.18"),
				),
			},
		},
	})
}

func TestAccNetworkV2LoadBalancerPlanDataSource_queries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerPlanDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerPlanDataSourceID("data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "name", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "description", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "maximum_syslog_servers", "8"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.edition", "Standard"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.size", "1000"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "vendor", "citrix"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "version", "12.1-55.18"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerPlanDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerPlanDataSourceID("data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "name", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "description", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "maximum_syslog_servers", "8"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.edition", "Standard"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.size", "1000"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "vendor", "citrix"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "version", "12.1-55.18"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerPlanDataSourceQueryEnabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerPlanDataSourceID("data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "name", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "description", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "maximum_syslog_servers", "8"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.edition", "Standard"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.size", "1000"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "vendor", "citrix"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "version", "12.1-55.18"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerPlanDataSourceQueryMaximumSyslogServers,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerPlanDataSourceID("data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "name", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "description", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "maximum_syslog_servers", "8"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.edition", "Standard"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.size", "1000"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "vendor", "citrix"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "version", "12.1-55.18"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerPlanDataSourceQueryModelEdition,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerPlanDataSourceID("data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "name", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "description", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "maximum_syslog_servers", "8"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.edition", "Standard"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.size", "1000"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "vendor", "citrix"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "version", "12.1-55.18"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerPlanDataSourceQueryVendor,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerPlanDataSourceID("data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "name", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "description", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "maximum_syslog_servers", "8"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.edition", "Standard"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.size", "1000"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "vendor", "citrix"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "version", "12.1-55.18"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerPlanDataSourceQueryVersion,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerPlanDataSourceID("data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "name", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "description", testAccNetworkV2LoadBalancerPlanDataSourcePlanName),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "maximum_syslog_servers", "8"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.edition", "Standard"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "model.0.size", "1000"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "vendor", "citrix"),
					resource.TestCheckResourceAttr(
						"data.ecl_network_load_balancer_plan_v2.load_balancer_plan_1", "version", "12.1-55.18"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2LoadBalancerPlanDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("can't find load_balancer_plan data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("load balancer plan data source ID not set")
		}

		return nil
	}
}

const testAccNetworkV2LoadBalancerPlanDataSourceBasic = `
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  enabled = true
  model {
    size = "1000"
  }
}
`

var testAccNetworkV2LoadBalancerPlanDataSourceQueryName = fmt.Sprintf(`
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  name = %q
}
`, testAccNetworkV2LoadBalancerPlanDataSourcePlanName,
)

var testAccNetworkV2LoadBalancerPlanDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  description = %q
}
`, testAccNetworkV2LoadBalancerPlanDataSourcePlanName,
)

var testAccNetworkV2LoadBalancerPlanDataSourceQueryEnabled = fmt.Sprintf(`
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  name = %q
  enabled = true
}
`, testAccNetworkV2LoadBalancerPlanDataSourcePlanName,
)

var testAccNetworkV2LoadBalancerPlanDataSourceQueryMaximumSyslogServers = fmt.Sprintf(`
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  name = %q
  maximum_syslog_servers = 8
}
`, testAccNetworkV2LoadBalancerPlanDataSourcePlanName,
)

var testAccNetworkV2LoadBalancerPlanDataSourceQueryModelEdition = fmt.Sprintf(`
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  name = %q
  model {
    edition = "Standard"
  }
}
`, testAccNetworkV2LoadBalancerPlanDataSourcePlanName,
)

var testAccNetworkV2LoadBalancerPlanDataSourceQueryVendor = fmt.Sprintf(`
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  name = %q
  vendor = "citrix"
}
`, testAccNetworkV2LoadBalancerPlanDataSourcePlanName,
)

const testAccNetworkV2LoadBalancerPlanDataSourceQueryVersion = `
data "ecl_network_load_balancer_plan_v2" "load_balancer_plan_1" {
  model {
    size = "1000"
  }
  version = "12.1-55.18"
}
`
