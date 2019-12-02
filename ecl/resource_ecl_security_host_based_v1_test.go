package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/host_based"
)

func TestAccSecurityV1HostBasedBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var hs security.HostBasedSecurity

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurityHostBased(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1HostBasedDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV1HostBasedBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1HostBasedExists(
						"ecl_security_host_based_v1.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"service_order_service", "Managed Anti-Virus"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"max_agent_value", "1"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"mail_address", OS_MAIL_ADDRESS),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV1HostBasedUpdateM1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1HostBasedExists(
						"ecl_security_host_based_v1.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"service_order_service", "Managed Virtual Patch"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"max_agent_value", "1"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"mail_address", OS_MAIL_ADDRESS),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV1HostBasedUpdateM2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1HostBasedExists(
						"ecl_security_host_based_v1.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"service_order_service", "Managed Virtual Patch"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"max_agent_value", "2"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"mail_address", OS_MAIL_ADDRESS),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV1HostBasedUpdateM2Again,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1HostBasedExists(
						"ecl_security_host_based_v1.host_1", &hs),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"service_order_service", "Managed Virtual Patch"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"max_agent_value", "3"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"mail_address", OS_MAIL_ADDRESS),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"dsm_lang", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_host_based_v1.host_1",
						"time_zone", "Asia/Tokyo"),
				),
			},
		},
	})
}

func testAccCheckSecurityV1HostBasedExists(n string, hs *security.HostBasedSecurity) resource.TestCheckFunc {
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

		getOpts := security.GetOpts{
			TenantID: OS_TENANT_ID,
		}
		found, err := security.Get(client, getOpts).Extract()
		if err != nil {
			return err
		}

		if found.ServiceOrderService == "" {
			return fmt.Errorf("Host based security not found")
		}

		*hs = *found

		return nil
	}
}

func testAccCheckSecurityV1HostBasedDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_host_based_v1" {
			continue
		}

		getOpts := security.GetOpts{
			TenantID: OS_TENANT_ID,
		}
		hs, err := security.Get(client, getOpts).Extract()
		if err != nil {
			return err
		}

		if hs.TenantFlg {
			return fmt.Errorf("Security single device still exists")
		}
	}

	return nil
}

var testAccSecurityV1HostBasedBasic = fmt.Sprintf(`
resource "ecl_security_host_based_v1" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service = "Managed Anti-Virus"
	max_agent_value = 1
	mail_address = "%s"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`,
	OS_TENANT_ID,
	OS_MAIL_ADDRESS,
)

var testAccSecurityV1HostBasedUpdateM1 = fmt.Sprintf(`
resource "ecl_security_host_based_v1" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service = "Managed Virtual Patch"
	max_agent_value = 1
	mail_address = "%s"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`,
	OS_TENANT_ID,
	OS_MAIL_ADDRESS,
)

var testAccSecurityV1HostBasedUpdateM2 = fmt.Sprintf(`
resource "ecl_security_host_based_v1" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service = "Managed Virtual Patch"
	max_agent_value = 2
	mail_address = "%s"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`,
	OS_TENANT_ID,
	OS_MAIL_ADDRESS,
)

var testAccSecurityV1HostBasedUpdateM2Again = fmt.Sprintf(`
resource "ecl_security_host_based_v1" "host_1" {
	tenant_id = "%s"
	locale = "ja"
	service_order_service = "Managed Virtual Patch"
	max_agent_value = 3
	mail_address = "%s"
	dsm_lang = "ja"
	time_zone = "Asia/Tokyo"
}
`,
	OS_TENANT_ID,
	OS_MAIL_ADDRESS,
)
