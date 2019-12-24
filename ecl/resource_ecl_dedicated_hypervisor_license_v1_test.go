package ecl

import (
	"fmt"
	"testing"

	"github.com/nttcom/eclcloud/ecl/dedicated_hypervisor/v1/licenses"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDedicatedHypervisorV1License_basic(t *testing.T) {
	var license licenses.License

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckDedicatedHypervisor(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDedicatedHypervisorV1LicenseDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDedicatedHypervisorV1LicenseBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDedicatedHypervisorV1LicenseExists("ecl_dedicated_hypervisor_license_v1.license_1", &license),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_license_v1.license_1", "license_type", "vCenter Server 6.x Standard"),
					resource.TestCheckResourceAttrSet("ecl_dedicated_hypervisor_license_v1.license_1", "key"),
					resource.TestCheckResourceAttrSet("ecl_dedicated_hypervisor_license_v1.license_1", "assigned_from"),
					resource.TestCheckNoResourceAttr("ecl_dedicated_hypervisor_license_v1.license_1", "expires_at"),
				),
			},
		},
	})
}

func testAccCheckDedicatedHypervisorV1LicenseDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.dedicatedHypervisorV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_dedicated_hypervisor_license_v1" {
			continue
		}

		if _, err := getLicense(client, rs.Primary.ID); err != nil {
			if err == licenseNotFoundError {
				continue
			}
			return err
		}

		return fmt.Errorf("license still exists")
	}

	return nil
}

func testAccCheckDedicatedHypervisorV1LicenseExists(n string, license *licenses.License) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.dedicatedHypervisorV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL Dedicated Hyperviosr client: %s", err)
		}

		found, err := getLicense(client, rs.Primary.ID)
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("license not found")
		}

		*license = *found

		return nil
	}
}

const testAccDedicatedHypervisorV1LicenseBasic = `
resource "ecl_dedicated_hypervisor_license_v1" "license_1" {
    license_type = "vCenter Server 6.x Standard"
}
`
