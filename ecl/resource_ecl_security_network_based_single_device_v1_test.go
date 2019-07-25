package ecl

import (
	"fmt"
	"log"
	// "strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	// "github.com/nttcom/eclcloud/ecl/network/v2/common_function_gateways"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/single_devices"
)

func TestAccSecurityV1NetworkBasedSingleDeviceBasic(t *testing.T) {
	var sd single_devices.SingleDevice

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedSingleDeviceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedSingleDeviceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedSingleDeviceExists(
						"ecl_security_network_based_single_device_v1.single_device_1", &sd),
				),
			},
		},
	})
}

func testAccCheckSecurityV1NetworkBasedSingleDeviceExists(n string, sd *single_devices.SingleDevice) resource.TestCheckFunc {
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

		found, err := getSingleDeviceByHostName(client, rs.Primary.ID)
		// found, err := single_devices.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		log.Printf("[MYDEBUG] found: %#v", found)
		log.Printf("[MYDEBUG] rs.Primary: %#v", rs.Primary)

		if found.Cell[2] != rs.Primary.ID {
			return fmt.Errorf("Security single device not found")
		}

		*sd = found

		return nil
	}
}

func testAccCheckSecurityV1NetworkBasedSingleDeviceDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_network_based_single_device_v1" {
			continue
		}
		// _, err := common_function_gateways.Get(client, rs.Primary.ID).Extract()

		_, err := getSingleDeviceByHostName(client, rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Common Function Gateway still exists")
		}

	}

	return nil
}

var testAccSecurityV1NetworkBasedSingleDeviceBasic = fmt.Sprintf(`
resource "ecl_security_network_based_single_device_v1" "single_device_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW"
	license_kind = "02"
	az_group = "zone1-groupb"
}
`,
	OS_TENANT_ID,
)
