package ecl

import (
	"fmt"
	"testing"

	"github.com/nttcom/eclcloud"

	"github.com/nttcom/eclcloud/ecl/dedicated_hypervisor/v1/servers"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDedicatedHypervisorV1Server_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var server servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckDedicatedHypervisor(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDedicatedHypervisorV1ServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDedicatedHypervisorV1ServerBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDedicatedHypervisorV1ServerExists("ecl_dedicated_hypervisor_server_v1.server_1", &server),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "name", "server1"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "description", "ESXi Dedicated Hypervisor"),
					resource.TestCheckResourceAttrSet("ecl_dedicated_hypervisor_server_v1.server_1", "networks.0.uuid"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.0.fixed_ip", "192.168.1.10"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.0.plane", "data"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.0.segmentation_id", "4"),
					resource.TestCheckResourceAttrSet("ecl_dedicated_hypervisor_server_v1.server_1", "networks.1.uuid"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.1.fixed_ip", "192.168.1.11"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.1.plane", "data"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "networks.1.segmentation_id", "4"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "admin_pass", "aabbccddeeff"),
					resource.TestCheckResourceAttrSet("ecl_dedicated_hypervisor_server_v1.server_1", "image_ref"),
					resource.TestCheckResourceAttrSet("ecl_dedicated_hypervisor_server_v1.server_1", "flavor_ref"),
					resource.TestCheckResourceAttrSet("ecl_dedicated_hypervisor_server_v1.server_1", "availability_zone"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "metadata.k1", "v1"),
					resource.TestCheckResourceAttr("ecl_dedicated_hypervisor_server_v1.server_1", "metadata.k2", "v2"),
					resource.TestCheckResourceAttrSet("ecl_dedicated_hypervisor_server_v1.server_1", "baremetal_server_id"),
				),
			},
		},
	})
}

func testAccCheckDedicatedHypervisorV1ServerDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.dedicatedHypervisorV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL Dedicated Hypervisor client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_dedicated_hypervisor_server_v1" {
			continue
		}

		if _, err := servers.Get(client, rs.Primary.ID).Extract(); err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				continue
			}

			return fmt.Errorf("error getting ECL Dedicated Hypervisor server: %s", err)
		}

		return fmt.Errorf("license still exists")
	}

	return nil
}

func testAccCheckDedicatedHypervisorV1ServerExists(n string, server *servers.Server) resource.TestCheckFunc {
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

		found, err := servers.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("server not found")
		}

		*server = *found

		return nil
	}
}

var testAccDedicatedHypervisorV1ServerBasic = fmt.Sprintf(`
data "ecl_baremetal_flavor_v2" "gp1" {
    name = "General Purpose 1 v2"
}

data "ecl_imagestorages_image_v2" "esxi" {
    name = "vSphere_ESXi-6.0.u2_64_dedicated-hypervisor_01"
}

data "ecl_baremetal_availability_zone_v2" "groupa" {
    zone_name = "%s"
}

resource "ecl_network_network_v2" "network_1" {
    name = "dedicated_hypervisor_network"
    plane = "data"
}

resource "ecl_network_subnet_v2" "subnet_1" {
    name = "dedicated_hypervisor_subnet"
    network_id = "${ecl_network_network_v2.network_1.id}"
    cidr = "192.168.1.0/24"
    gateway_ip = "192.168.1.1"
    allocation_pools {
        start = "192.168.1.100"
        end = "192.168.1.200"
    }
}

resource "ecl_dedicated_hypervisor_server_v1" "server_1" {
    depends_on = [
        "ecl_network_subnet_v2.subnet_1"
    ]

    name = "server1"
    description = "ESXi Dedicated Hypervisor"
    networks {
        uuid = "${ecl_network_network_v2.network_1.id}"
        fixed_ip = "192.168.1.10"
        plane = "data"
        segmentation_id = 4
    }
    networks {
        uuid = "${ecl_network_network_v2.network_1.id}"
        fixed_ip = "192.168.1.11"
        plane = "data"
        segmentation_id = 4
    }
    admin_pass = "aabbccddeeff"
    image_ref = "${data.ecl_imagestorages_image_v2.esxi.id}"
    flavor_ref = "${data.ecl_baremetal_flavor_v2.gp1.id}"
    availability_zone = "${data.ecl_baremetal_availability_zone_v2.groupa.zone_name}"
    metadata = {
        k1 = "v1"
        k2 = "v2"
    }
}
`,
	OS_BAREMETAL_ZONE,
)
