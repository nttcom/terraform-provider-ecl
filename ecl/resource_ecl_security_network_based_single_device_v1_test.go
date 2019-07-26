package ecl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_firewall_utm_single"
)

func TestAccSecurityV1NetworkBasedFirewallUTMBasic(t *testing.T) {
	var sd security.SingleFirewallUTM

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedFirewallUTMDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedFirewallUTMBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedFirewallUTMExists(
						"ecl_security_network_based_firewall_utm_single_v1.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"operating_mode", "FW"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"az_group", "zone1-groupb"),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedFirewallUTMUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedFirewallUTMExists(
						"ecl_security_network_based_firewall_utm_single_v1.device_1", &sd),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"operating_mode", "UTM"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_firewall_utm_single_v1.device_1",
						"az_group", "zone1-groupb"),
				),
			},
		},
	})
}

func testAccCheckSecurityV1NetworkBasedFirewallUTMExists(n string, sd *security.SingleFirewallUTM) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.securityOrderV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL security client: %s", err)
		}

		found, err := getSingleFirewallUTMByHostName(client, rs.Primary.ID)
		if err != nil {
			return err
		}

		if found.Cell[2] != rs.Primary.ID {
			return fmt.Errorf("Security single device not found")
		}

		*sd = found

		return nil
	}
}

func testAccCheckSecurityV1NetworkBasedFirewallUTMDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_network_based_firewall_utm_single_v1" {
			continue
		}

		_, err := getSingleFirewallUTMByHostName(client, rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Common Function Gateway still exists")
		}

	}

	return nil
}

var testAccSecurityV1NetworkBasedFirewallUTMBasic = fmt.Sprintf(`
resource "ecl_security_network_based_firewall_utm_single_v1" "device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)

var testAccSecurityV1NetworkBasedFirewallUTMUpdate = fmt.Sprintf(`
resource "ecl_security_network_based_firewall_utm_single_v1" "device_1" {
	tenant_id = "%s"
	locale = "en"
	operating_mode = "UTM"
	license_kind = "08"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)
