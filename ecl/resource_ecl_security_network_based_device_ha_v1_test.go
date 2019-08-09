package ecl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_ha"
)

func TestAccSecurityV1NetworkBasedDeviceHABasic(t *testing.T) {
	var hd1, hd2 security.HADevice

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV1NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedDeviceHABasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v1.ha_1", &hd1, &hd2),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"operating_mode", "FW_HA"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"host_1_az_group", "zone1-groupa"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"host_2_az_group", "zone1-groupb"),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV1NetworkBasedDeviceHAUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV1NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v1.ha_1", &hd1, &hd2),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"operating_mode", "UTM_HA"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"host_1_az_group", "zone1-groupa"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v1.ha_1",
						"host_2_az_group", "zone1-groupb"),
				),
			},
		},
	})
}

func testAccCheckSecurityV1NetworkBasedDeviceHAExists(n string, hd1, hd2 *security.HADevice) resource.TestCheckFunc {
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

		ids := strings.Split(rs.Primary.ID, "/")
		id1 := ids[0]
		id2 := ids[1]

		found1, err := getHADeviceByHostName(client, id1)
		if err != nil {
			return err
		}

		if found1.Cell[3] != id1 {
			return fmt.Errorf("Security single device-1 not found")
		}
		*hd1 = found1

		found2, err := getHADeviceByHostName(client, id2)
		if err != nil {
			return err
		}

		if found2.Cell[3] != id2 {
			return fmt.Errorf("Security single device-2 not found")
		}
		*hd2 = found2

		return nil
	}
}
func testAccCheckSecurityV1NetworkBasedDeviceHADestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_network_based_device_ha_v1" {
			continue
		}

		ids := strings.Split(rs.Primary.ID, "/")
		id1 := ids[0]
		id2 := ids[1]

		_, err := getHADeviceByHostName(client, id1)
		if err == nil {
			return fmt.Errorf("Security single device-1 still exists")
		}

		_, err = getHADeviceByHostName(client, id2)
		if err == nil {
			return fmt.Errorf("Security single device-2 still exists")
		}

	}

	return nil
}

const testAccSecurityV1NetworkBasedDeviceHANetworkSubnet1 = `
resource "ecl_network_network_v2" "network_1" {
	name = "network_1_for_utm_ha"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "subnet_1_for_utm_ha"
	cidr = "192.168.1.0/29"
	network_id = "${ecl_network_network_v2.network_1.id}"
	no_gateway = "true"
}
`
const testAccSecurityV1NetworkBasedDeviceHANetworkSubnet2 = `
resource "ecl_network_network_v2" "network_2" {
	name = "network_2_for_utm_ha"
}

resource "ecl_network_subnet_v2" "subnet_2" {
	name = "subnet_2_for_utm_ha"
	cidr = "192.168.2.0/29"
	network_id = "${ecl_network_network_v2.network_2.id}"
	no_gateway = "true"
}
`

var testAccSecurityV1NetworkBasedDeviceHABasic = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_device_ha_v1" "ha_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW_HA"
	license_kind = "02"

	host_1_az_group = "zone1-groupa"
	host_2_az_group = "zone1-groupb"
  
	ha_link_1 {
		network_id = "${ecl_network_network_v2.network_1.id}"
		subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
		host_1_ip_address = "192.168.1.3"
		host_2_ip_address = "192.168.1.4"
	}

	ha_link_2 {
		network_id = "${ecl_network_network_v2.network_2.id}"
		subnet_id = "${ecl_network_subnet_v2.subnet_2.id}"
		host_1_ip_address = "192.168.2.3"
		host_2_ip_address = "192.168.2.4"
	}

}
`,
	testAccSecurityV1NetworkBasedDeviceHANetworkSubnet1,
	testAccSecurityV1NetworkBasedDeviceHANetworkSubnet2,
	OS_TENANT_ID,
)

var testAccSecurityV1NetworkBasedDeviceHAUpdate = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_device_ha_v1" "ha_1" {
	tenant_id = "%s"
	locale = "en"
	operating_mode = "UTM_HA"
	license_kind = "08"

	host_1_az_group = "zone1-groupa"
	host_2_az_group = "zone1-groupb"
  
	ha_link_1 {
		network_id = "${ecl_network_network_v2.network_1.id}"
		subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
		host_1_ip_address = "192.168.1.3"
		host_2_ip_address = "192.168.1.4"
	}

	ha_link_2 {
		network_id = "${ecl_network_network_v2.network_2.id}"
		subnet_id = "${ecl_network_subnet_v2.subnet_2.id}"
		host_1_ip_address = "192.168.2.3"
		host_2_ip_address = "192.168.2.4"
	}

}
`,
	testAccSecurityV1NetworkBasedDeviceHANetworkSubnet1,
	testAccSecurityV1NetworkBasedDeviceHANetworkSubnet2,
	OS_TENANT_ID,
)
