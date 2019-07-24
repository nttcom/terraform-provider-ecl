package ecl

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	// "github.com/hashicorp/terraform/terraform"

	// "github.com/nttcom/eclcloud/ecl/network/v2/common_function_gateways"
	"github.com/nttcom/eclcloud/ecl/security_order/v1/single_devices"
)

func TestMockedAccSecurityV1NetworkBasedSingleDeviceBasic(t *testing.T) {
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

var testMockedAccSecurityV1NetworkBasedSingleDeviceBasic = fmt.Sprintf(`
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
