package ecl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v2/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/v2/ecl/network/v2/subnets"
	security "github.com/nttcom/eclcloud/v2/ecl/security_order/v2/network_based_device_ha"
)

func TestAccSecurityV2NetworkBasedDeviceHA_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var hd1, hd2 security.HADevice

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV2NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV2NetworkBasedDeviceHABasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV2NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v2.ha_1", &hd1, &hd2),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"operating_mode", "FW_HA"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"host_1_az_group", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"host_2_az_group", OS_COMPUTE_ZONE_HA),
				),
			},
			resource.TestStep{
				Config: testAccSecurityV2NetworkBasedDeviceHAUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV2NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v2.ha_1", &hd1, &hd2),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"locale", "en"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"operating_mode", "UTM_HA"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"license_kind", "08"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"host_1_az_group", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"host_2_az_group", OS_COMPUTE_ZONE_HA),
				),
			},
		},
	})
}

func TestAccSecurityV2NetworkBasedDeviceHA_updateInterface(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var hd1, hd2 security.HADevice
	var n1, n2, n3, n4 networks.Network
	var sn1, sn2, sn3, sn4 subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSecurity(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityV2NetworkBasedDeviceSingleDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSecurityV2NetworkBasedDeviceHABasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV2NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v2.ha_1", &hd1, &hd2),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn1),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2", &sn2),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"locale", "ja"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"operating_mode", "FW_HA"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"license_kind", "02"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"host_1_az_group", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"host_2_az_group", OS_COMPUTE_ZONE_HA),

					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"ha_link_1.0.network_id", &n1.ID),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"ha_link_1.0.subnet_id", &sn1.ID),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"ha_link_1.0.host_1_ip_address", "192.168.1.3"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"ha_link_1.0.host_2_ip_address", "192.168.1.4"),

					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"ha_link_2.0.network_id", &n2.ID),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"ha_link_2.0.subnet_id", &sn2.ID),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"ha_link_2.0.host_1_ip_address", "192.168.2.3"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1",
						"ha_link_2.0.host_2_ip_address", "192.168.2.4"),
				),
			},

			resource.TestStep{
				Config: testAccSecurityV2NetworkBasedDeviceHAIntrfaceUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV2NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v2.ha_1", &hd1, &hd2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn1),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2", &sn2),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.user_network_1", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.user_subnet_1", &sn3),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.user_network_2", &n4),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.user_subnet_2", &sn4),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.vrrp_ip_address", "10.0.0.50"),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.network_id", &n3.ID),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.subnet_id", &sn3.ID),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.comment", "port 0 comment"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.host_1_ip_address", "10.0.0.51"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.host_1_ip_address_prefix", "24"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.host_2_ip_address", "10.0.0.52"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.host_2_ip_address_prefix", "24"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.vrrp_ip_address", "10.0.1.50"),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.network_id", &n4.ID),
					resource.TestCheckResourceAttrPtr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.subnet_id", &sn4.ID),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.mtu", "1500"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.comment", "port 3 comment"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.host_1_ip_address", "10.0.1.51"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.host_1_ip_address_prefix", "24"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.host_2_ip_address", "10.0.1.52"),
					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.host_2_ip_address_prefix", "24"),
				),
			},

			resource.TestStep{
				Config: testAccSecurityV2NetworkBasedDeviceHAIntrfaceDisconnect,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityV2NetworkBasedDeviceHAExists(
						"ecl_security_network_based_device_ha_v2.ha_1", &hd1, &hd2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn1),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2", &sn2),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.user_network_1", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.user_subnet_1", &sn3),

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.user_network_2", &n4),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.user_subnet_2", &sn4),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.0.enable", "false"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.1.enable", "false"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.2.enable", "false"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.3.enable", "false"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.4.enable", "false"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.5.enable", "false"),

					resource.TestCheckResourceAttr(
						"ecl_security_network_based_device_ha_v2.ha_1", "port.6.enable", "false"),
				),
			},
		},
	})
}

func testAccCheckSecurityV2NetworkBasedDeviceHAExists(n string, hd1, hd2 *security.HADevice) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.securityOrderV2Client(OS_REGION_NAME)
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
func testAccCheckSecurityV2NetworkBasedDeviceHADestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.securityOrderV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_security_network_based_device_ha_v2" {
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

const testAccSecurityV2NetworkBasedDeviceHANetworkSubnet1 = `
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
const testAccSecurityV2NetworkBasedDeviceHANetworkSubnet2 = `
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

const testAccSecurityV2NetworkBasedDeviceUserNetworkSubnet1 = `
resource "ecl_network_network_v2" "user_network_1" {
	name = "network_1_for_utm_user"
}

resource "ecl_network_subnet_v2" "user_subnet_1" {
	name = "subnet_1_for_utm_user"
	cidr = "10.0.0.0/24"
	network_id = "${ecl_network_network_v2.user_network_1.id}"
	gateway_ip = "10.0.0.1"
	allocation_pools {
		start = "10.0.0.100"
		end = "10.0.0.200"
	}
}
`
const testAccSecurityV2NetworkBasedDeviceUserNetworkSubnet2 = `
resource "ecl_network_network_v2" "user_network_2" {
	name = "network_1_for_utm_user"
}

resource "ecl_network_subnet_v2" "user_subnet_2" {
	name = "subnet_1_for_utm_user"
	cidr = "10.0.1.0/24"
	network_id = "${ecl_network_network_v2.user_network_2.id}"
	gateway_ip = "10.0.1.1"
	allocation_pools {
		start = "10.0.1.100"
		end = "10.0.1.200"
	}
}
`

var testAccSecurityV2NetworkBasedDeviceHABasic = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_device_ha_v2" "ha_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW_HA"
	license_kind = "02"

	host_1_az_group = "%s"
	host_2_az_group = "%s"
  
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
	testAccSecurityV2NetworkBasedDeviceHANetworkSubnet1,
	testAccSecurityV2NetworkBasedDeviceHANetworkSubnet2,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testAccSecurityV2NetworkBasedDeviceHAUpdate = fmt.Sprintf(`
%s
%s

resource "ecl_security_network_based_device_ha_v2" "ha_1" {
	tenant_id = "%s"
	locale = "en"
	operating_mode = "UTM_HA"
	license_kind = "08"

	host_1_az_group = "%s"
	host_2_az_group = "%s"
  
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
	testAccSecurityV2NetworkBasedDeviceHANetworkSubnet1,
	testAccSecurityV2NetworkBasedDeviceHANetworkSubnet2,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testAccSecurityV2NetworkBasedDeviceHAIntrfaceUpdate = fmt.Sprintf(`
%s
%s
%s
%s

resource "ecl_security_network_based_device_ha_v2" "ha_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW_HA"
	license_kind = "02"

	host_1_az_group = "%s"
	host_2_az_group = "%s"
  
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

	port {
		enable = "true"

		network_id = "${ecl_network_network_v2.user_network_1.id}"
		subnet_id = "${ecl_network_subnet_v2.user_subnet_1.id}"
		mtu = "1500"
		comment = "port 0 comment"
		enable_ping = "true"

		host_1_ip_address = "10.0.0.51"
		host_1_ip_address_prefix = 24

		host_2_ip_address = "10.0.0.52"
		host_2_ip_address_prefix = 24

		vrrp_ip_address = "10.0.0.50"
		vrrp_grp_id = "11"
		vrrp_id = "50"
		preempt = "true"
	}

	port {
	  enable = "false"
	}

	port {
	  enable = "false"
	}

	port {
		enable = "true"

		network_id = "${ecl_network_network_v2.user_network_2.id}"
		subnet_id = "${ecl_network_subnet_v2.user_subnet_2.id}"
		mtu = "1500"
		comment = "port 3 comment"
		enable_ping = "true"

		host_1_ip_address = "10.0.1.51"
		host_1_ip_address_prefix = 24

		host_2_ip_address = "10.0.1.52"
		host_2_ip_address_prefix = 24

		vrrp_ip_address = "10.0.1.50"
		vrrp_grp_id = "11"
		vrrp_id = "60"
		preempt = "true"
	}

	port {
	    enable = "false"
	}

	port {
	    enable = "false"
	}

	port {
	    enable = "false"
	}
}
`,
	testAccSecurityV2NetworkBasedDeviceHANetworkSubnet1,
	testAccSecurityV2NetworkBasedDeviceHANetworkSubnet2,
	testAccSecurityV2NetworkBasedDeviceUserNetworkSubnet1,
	testAccSecurityV2NetworkBasedDeviceUserNetworkSubnet2,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)

var testAccSecurityV2NetworkBasedDeviceHAIntrfaceDisconnect = fmt.Sprintf(`
%s
%s
%s
%s

resource "ecl_security_network_based_device_ha_v2" "ha_1" {
	tenant_id = "%s"
	locale = "ja"
	operating_mode = "FW_HA"
	license_kind = "02"

	host_1_az_group = "%s"
	host_2_az_group = "%s"
  
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

	port {
		enable = "false"
	}

	port {
		enable = "false"
	}

	port {
		enable = "false"
	}

	port {
		enable = "false"
	}

	port {
		enable = "false"
	}

	port {
		enable = "false"
	}

	port {
		enable = "false"
	}
}
`,
	testAccSecurityV2NetworkBasedDeviceHANetworkSubnet1,
	testAccSecurityV2NetworkBasedDeviceHANetworkSubnet2,
	testAccSecurityV2NetworkBasedDeviceUserNetworkSubnet1,
	testAccSecurityV2NetworkBasedDeviceUserNetworkSubnet2,
	OS_TENANT_ID,
	OS_DEFAULT_ZONE,
	OS_COMPUTE_ZONE_HA,
)
