package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func TestAccVNAV1ApplianceBasic(t *testing.T) {
	var vna appliances.Appliance

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
				),
			},
		},
	})
}

func testAccCheckVNAV1ApplianceDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	vnaClient, err := config.vnaV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL networking client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_vna_appliance_v1" {
			continue
		}

		_, err := appliances.Get(vnaClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Virtual Network Appliance still exists")
		}
	}

	return nil
}

func testAccCheckVNAV1ApplianceExists(n string, vna *appliances.Appliance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		vnaClient, err := config.vnaV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL virtual network appliance client: %s", err)
		}

		found, err := appliances.Get(vnaClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Virtual Network Appliance not found")
		}

		*vna = *found

		return nil
	}
}

const testAccVNAV1ApplianceSingleNetworkAndSubnetPair = `
resource "ecl_network_network_v2" "network_1" {
	name = "network_1"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	name = "subnet_1"
	cidr = "192.168.1.0/24"
	network_id = "${ecl_network_network_v2.network_1.id}"
	gateway_ip = "192.168.1.1"
	allocation_pools {
		start = "192.168.1.100"
		end = "192.168.1.200"
	}
}
`

var testAccVNAV1ApplianceBasic = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "%s"
	depends_on = ["ecl_network_subnet_v2.subnet_1"]

	interfaces {
		slot_number = 1
		name = "interface_1"
		network_id = "${ecl_network_network_v2.network_1.id}"
		fixed_ips {
			ip_address = "192.168.1.50"
		}
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)
