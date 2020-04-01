package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nttcom/eclcloud/ecl/baremetal/v2/servers"
)

func TestAccBaremetalV2Server_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var server servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckBaremetal(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBaremetalV2ServerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBaremetalV2ServerBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaremetalV2ServerExists("ecl_baremetal_server_v2.server_1", &server),
				),
			},
		},
	})
}

func testAccCheckBaremetalV2ServerDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.baremetalV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_baremetal_server_v2" {
			continue
		}

		_, err := servers.Get(client, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Server still exists")
		}
	}

	return nil
}

func testAccCheckBaremetalV2ServerExists(n string, server *servers.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		baremetalClient, err := config.baremetalV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL baremetal client: %s", err)
		}

		found, err := servers.Get(baremetalClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Baremetal server not found")
		}

		*server = *found

		return nil
	}
}

var testAccBaremetalV2ServerBasic = fmt.Sprintf(`
data "ecl_imagestorages_image_v2" "centos" {
    name = "CentOS-7.3-1611_64_baremetal-server_01"
}

data "ecl_baremetal_flavor_v2" "gp2" {
    name = "General Purpose 2 v1"
}

data "ecl_baremetal_availability_zone_v2" "groupa" {
    zone_name = "%s"
}

resource "ecl_network_network_v2" "network_1" {
    name = "baremetal_network"
    plane = "data"
}

resource "ecl_network_subnet_v2" "subnet_1" {
    name = "baremetal_subnet"
    network_id = "${ecl_network_network_v2.network_1.id}"
    cidr = "192.168.1.0/24"
    gateway_ip = "192.168.1.1"
    allocation_pools {
        start = "192.168.1.100"
        end = "192.168.1.200"
    }
}

resource "ecl_baremetal_keypair_v2" "keypair_1" {
    name = "keypair1"
}

resource "ecl_baremetal_server_v2" "server_1" {
    depends_on = [
        "ecl_network_subnet_v2.subnet_1"
    ]

    name = "server1"
    image_id = "${data.ecl_imagestorages_image_v2.centos.id}"
    flavor_id = "${data.ecl_baremetal_flavor_v2.gp2.id}"
    user_data = "user_data"
    availability_zone = "${data.ecl_baremetal_availability_zone_v2.groupa.zone_name}"
    key_pair = "${ecl_baremetal_keypair_v2.keypair_1.name}"
    admin_pass = "password"
    metadata = {
        k1 = "v1"
        k2 = "v2"
    }
    networks {
        uuid = "${ecl_network_network_v2.network_1.id}"
        fixed_ip = "192.168.1.10"
        plane = "data"
    }
    raid_arrays {
        primary_storage = true
        partitions {
            lvm = true
            partition_label = "primary-part1"
        }
        partitions {
            lvm = false
            size = "100G"
            partition_label = "var"
        }
    }
    lvm_volume_groups {
        vg_label = "VG_root"
        physical_volume_partition_labels = ["primary-part1"]
        logical_volumes {
            lv_label = "LV_root"
            size = "300G"
        }
        logical_volumes {
            lv_label = "LV_swap"
            size = "2G"
        }
    }
    filesystems {
        label = "LV_root"
        mount_point =  "/"
        fs_type = "xfs"
    }
    filesystems {
        label = "var"
        mount_point = "/var"
        fs_type = "xfs"
    }
    filesystems {
        label = "LV_swap"
        fs_type = "swap"
    }
    personality {
        path = "/home/big/banner.txt"
        contents = "ZWNobyAiS3VtYSBQZXJzb25hbGl0eSIgPj4gL2hvbWUvYmlnL3BlcnNvbmFsaXR5"
    }
}
`,
	OS_BAREMETAL_AVAILABLE_ZONE,
)
