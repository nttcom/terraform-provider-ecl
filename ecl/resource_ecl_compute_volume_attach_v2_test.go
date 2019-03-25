package ecl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/compute/v2/extensions/volumeattach"
	"github.com/nttcom/eclcloud/ecl/compute/v2/servers"
	"github.com/nttcom/eclcloud/ecl/computevolume/v2/volumes"
)

func TestAccComputeVolumeV2AttachBasic(t *testing.T) {
	var va volumes.Attachment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2AttachDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2AttachBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeVolumeV2AttachExists("ecl_compute_volume_attach_v2.va_1", &va),
				),
			},
		},
	})
}

func TestAccComputeVolumeV2AttachTimeout(t *testing.T) {
	var va volumes.Attachment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVolumeV2AttachDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeVolumeV2AttachTimeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeVolumeV2AttachExists("ecl_compute_volume_attach_v2.va_1", &va),
				),
			},
		},
	})
}

func testAccCheckComputeVolumeV2AttachDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)

	computeClient, err := config.computeV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL compute client: %s", err)
	}

	computeVolumeClient, err := config.computeVolumeV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL compute volume client: %s", err)
	}

	// Cofirmation for deletion for attachment
	log.Printf("[DEBUG] Confirming volume attachment deletion.")
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_compute_volume_attach_v2" {
			continue
		}

		_, serverID, err := parseAttachmentID(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("%s", err)
		}

		_, err = volumeattach.List(computeClient, serverID).AllPages()
		if err == nil {
			return fmt.Errorf("Volume attachment still exists")
		}
	}

	// Confirmation for deletion for instance
	log.Printf("[DEBUG] Confirming instance deletion.")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_compute_instance_v2" {
			continue
		}

		_, err := servers.Get(computeClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Compute Instance still exists")
		}
	}

	// Confirmation for deletion for volume
	log.Printf("[DEBUG] Confirming volume deletion.")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_compute_volume_v2" {
			continue
		}

		_, err := volumes.Get(computeVolumeClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Compute volume still exists")
		}

	}

	return nil
}

func testAccCheckComputeVolumeV2AttachExists(n string, va *volumes.Attachment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.computeVolumeV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL compute volume client: %s", err)
		}

		attachmentID, _, err := parseAttachmentID(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("%s", err)
		}

		volumeID := attachmentID
		volume, err := volumes.Get(client, volumeID).Extract()
		if err != nil {
			return err
		}

		var found bool
		found = (volume.Status == "in-use")

		if !found {
			return fmt.Errorf("Volume Attachment not found")
		}

		return nil
	}
}

const testCreateNetworkForAttachTargetInstance = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  plane = "data"
}

resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  network_id = "${ecl_network_network_v2.network_1.id}"
  cidr = "192.168.1.0/24"
  gateway_ip = "192.168.1.1"
  allocation_pools = {
    start = "192.168.1.100"
    end = "192.168.1.200"
  }
}
`

var testAccComputeVolumeV2AttachBasic = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}

resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1"
  size = 15
}

resource "ecl_compute_volume_attach_v2" "va_1" {
  volume_id = "${ecl_compute_volume_v2.volume_1.id}"
  server_id = "${ecl_compute_instance_v2.instance_1.id}"
  device = "/dev/vdb"
}
`, testCreateNetworkForAttachTargetInstance)

var testAccComputeVolumeV2AttachTimeout = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}

resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1"
  size = 15
}

resource "ecl_compute_volume_attach_v2" "va_1" {
  volume_id = "${ecl_compute_volume_v2.volume_1.id}"
  server_id = "${ecl_compute_instance_v2.instance_1.id}"
  device = "/dev/vdb"
  
  timeouts {
    create = "5m"
  }
}
`, testCreateNetworkForAttachTargetInstance)
