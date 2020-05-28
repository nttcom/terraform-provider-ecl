package ecl

import (
	"fmt"
	"strings"
	"testing"

	"crypto/sha1"
	"encoding/hex"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/compute/v2/extensions/volumeattach"
	"github.com/nttcom/eclcloud/ecl/compute/v2/servers"
	"github.com/nttcom/eclcloud/ecl/computevolume/v2/volumes"
	"github.com/nttcom/eclcloud/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/ecl/network/v2/ports"
	"github.com/nttcom/eclcloud/pagination"
)

func TestAccComputeV2Instance_basic(t *testing.T) {
	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "name", "i"),
					testAccCheckComputeV2InstanceMetadata(&instance, "foo", "bar"),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "all_metadata.foo", "bar"),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "config_drive", "false"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "name", strings.Repeat("a", 255)),
					testAccCheckComputeV2InstanceMetadata(&instance, "foo", "bar"),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "all_metadata.foo", "bar"),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_resize(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceResizeBase,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "flavor_id", "1CPU-2GB"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceResizeUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "flavor_id", "1CPU-4GB"),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_userData(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceUserData,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "user_data", encodedUserData()),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_configDrive(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceEnabledConfigDrive,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "config_drive", "true"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceDisabledConfigDrive,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "config_drive", "false"),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_initialStateActive(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceStateActive,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "power_state", "active"),
					testAccCheckComputeV2InstanceState(&instance, "active"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceStateShutoff,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "power_state", "shutoff"),
					testAccCheckComputeV2InstanceState(&instance, "shutoff"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceStateActive,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "power_state", "active"),
					testAccCheckComputeV2InstanceState(&instance, "active"),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_initialStateShutoff(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceStateShutoff,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "power_state", "shutoff"),
					testAccCheckComputeV2InstanceState(&instance, "shutoff"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceStateActive,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "power_state", "active"),
					testAccCheckComputeV2InstanceState(&instance, "active"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceStateShutoff,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "power_state", "shutoff"),
					testAccCheckComputeV2InstanceState(&instance, "shutoff"),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_bootFromVolumeWhichHasImageSource(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckDefaultZone(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceBootFromVolumeWhichHasImageSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					testAccCheckComputeV2InstanceBootVolumeAttachment(&instance),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_blockDeviceExistingVolume(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server
	var volume volumes.Volume

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckDefaultZone(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceBlockDeviceExistingVolume,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					testAccCheckBlockStorageV2VolumeExists(
						"ecl_compute_volume_v2.volume_1", &volume),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_keyPairForceNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance1 servers.Server
	var instance2 servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceKeyPairForceNew1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists(
						"ecl_compute_instance_v2.instance_1", &instance1),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "key_pair", "kp_1"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceKeyPairForceNew2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists(
						"ecl_compute_instance_v2.instance_1", &instance2),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "key_pair", "kp_2"),
					testAccCheckComputeV2InstanceInstanceIDsDoNotMatch(&instance1, &instance2),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_bootFromVolumeForceNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance1 servers.Server
	var instance2 servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckDefaultZone(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceBootFromVolumeForceNew1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists(
						"ecl_compute_instance_v2.instance_1", &instance1),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceBootFromVolumeForceNew2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists(
						"ecl_compute_instance_v2.instance_1", &instance2),
					testAccCheckComputeV2InstanceInstanceIDsDoNotMatch(&instance1, &instance2),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_blockDeviceNewVolume(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckDefaultZone(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceBlockDeviceNewVolume,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_accessIPv4(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceAccessIPv4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "access_ip_v4", "192.168.1.50"),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_changeFixedIP(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance1 servers.Server
	var instance2 servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceChangeFixedIP1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists(
						"ecl_compute_instance_v2.instance_1", &instance1),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "access_ip_v4", "192.168.1.10"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceChangeFixedIP2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists(
						"ecl_compute_instance_v2.instance_1", &instance2),
					testAccCheckComputeV2InstanceInstanceIDsDoNotMatch(&instance1, &instance2),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "access_ip_v4", "192.168.1.20"),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_stopBeforeDestroy(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceStopBeforeDestroy,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_metadataRemove(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceMetadataRemove1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					testAccCheckComputeV2InstanceMetadata(&instance, "k1", "v1"),
					testAccCheckComputeV2InstanceMetadata(&instance, "k2", "v2"),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "all_metadata.k1", "v1"),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "all_metadata.k2", "v2"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceMetadataRemove2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					testAccCheckComputeV2InstanceMetadata(&instance, "k3", "v3"),
					testAccCheckComputeV2InstanceMetadata(&instance, "k1", "v1"),
					testAccCheckComputeV2InstanceNoMetadataKey(&instance, "k2"),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "all_metadata.k3", "v3"),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "all_metadata.k1", "v1"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2InstanceMetadataRemove3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					testAccCheckComputeV2InstanceMetadataLengthIsZERO(&instance),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_timeout(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceTimeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_networkNameToID(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server
	var network networks.Network
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceNetworkNameToID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &network),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "network.0.name", "network_1"),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_multipleNICs(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceMultipleNICs(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_connectToCreatedPort(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server
	var port ports.Port

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceConnectToCreatedPort,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					testAccCheckNetworkV2PortExists("ecl_network_port_v2.port_1", &port),
					resource.TestCheckResourceAttrPtr(
						"ecl_compute_instance_v2.instance_1", "network.0.port", &port.ID),
					resource.TestCheckResourceAttrPtr(
						"ecl_compute_instance_v2.instance_1", "network.0.mac", &port.MACAddress),
				),
			},
		},
	})
}

func TestAccComputeV2Instance_connectToNetworkByName(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var instance servers.Server

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeV2InstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2InstanceConnectToNetworkByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2InstanceExists("ecl_compute_instance_v2.instance_1", &instance),
					resource.TestCheckResourceAttr(
						"ecl_compute_instance_v2.instance_1", "network.0.name", "network_1"),
				),
			},
		},
	})
}

func testAccCheckComputeV2InstanceDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	computeClient, err := config.computeV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL compute client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_compute_instance_v2" {
			continue
		}

		server, err := servers.Get(computeClient, rs.Primary.ID).Extract()
		if err == nil {
			if server.Status != "SOFT_DELETED" {
				return fmt.Errorf("Instance still exists")
			}
		}
	}

	return nil
}

func testAccCheckComputeV2InstanceExists(n string, instance *servers.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		computeClient, err := config.computeV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL compute client: %s", err)
		}

		found, err := servers.Get(computeClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}

		*instance = *found

		return nil
	}
}

func testAccCheckComputeV2InstanceDoesNotExist(n string, instance *servers.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		computeClient, err := config.computeV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL compute client: %s", err)
		}

		_, err = servers.Get(computeClient, instance.ID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return nil
			}
			return err
		}

		return fmt.Errorf("Instance still exists")
	}
}

func testAccCheckComputeV2InstanceMetadataLengthIsZERO(instance *servers.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(instance.Metadata) != 0 {
			return fmt.Errorf("Metadata length is not ZERO")
		}
		return nil
	}
}

func testAccCheckComputeV2InstanceMetadata(
	instance *servers.Server, k string, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if instance.Metadata == nil {
			return fmt.Errorf("No metadata")
		}

		for key, value := range instance.Metadata {
			if k != key {
				continue
			}

			if v == value {
				return nil
			}

			return fmt.Errorf("Bad value for %s: %s", k, value)
		}

		return fmt.Errorf("Metadata not found: %s", k)
	}
}

func testAccCheckComputeV2InstanceNoMetadataKey(
	instance *servers.Server, k string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if instance.Metadata == nil {
			return nil
		}

		for key := range instance.Metadata {
			if k == key {
				return fmt.Errorf("Metadata found: %s", k)
			}
		}

		return nil
	}
}

func testAccCheckComputeV2InstanceBootVolumeAttachment(
	instance *servers.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var attachments []volumeattach.VolumeAttachment

		config := testAccProvider.Meta().(*Config)
		computeClient, err := config.computeV2Client(OS_REGION_NAME)
		if err != nil {
			return err
		}

		err = volumeattach.List(computeClient, instance.ID).EachPage(
			func(page pagination.Page) (bool, error) {

				actual, err := volumeattach.ExtractVolumeAttachments(page)
				if err != nil {
					return false, fmt.Errorf("Unable to lookup attachment: %s", err)
				}

				attachments = actual
				return true, nil
			})

		if len(attachments) == 1 {
			return nil
		}

		return fmt.Errorf("No attached volume found")
	}
}

func testAccCheckComputeV2InstanceInstanceIDsDoNotMatch(
	instance1, instance2 *servers.Server) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if instance1.ID == instance2.ID {
			return fmt.Errorf("Instance was not recreated")
		}

		return nil
	}
}

func testAccCheckComputeV2InstanceState(
	instance *servers.Server, state string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if strings.ToLower(instance.Status) != state {
			return fmt.Errorf("Instance state is not match")
		}

		return nil
	}
}

const testImageDataSource = `
data "ecl_imagestorages_image_v2" "image_1" {
	name = "Ubuntu-18.04.1_64_virtual-server_02"
}
`

const testCreateNetworkForInstance = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
  plane = "data"
}
 resource "ecl_network_subnet_v2" "subnet_1" {
  name = "subnet_1"
  network_id = "${ecl_network_network_v2.network_1.id}"
  cidr = "192.168.1.0/24"
  gateway_ip = "192.168.1.1"
  allocation_pools {
    start = "192.168.1.100"
    end = "192.168.1.200"
  }
}`

var testAccComputeV2InstanceBasic = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "i"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  metadata = {
    foo = "bar"
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceUpdate = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "%s"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  metadata = {
    foo = "bar"
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance, strings.Repeat("a", 255))

func encodedUserData() string {
	s := "#!/bin/sh\necho 'HOGE'"
	hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

var testAccComputeV2InstanceUserData = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  user_data = "#!/bin/sh\necho 'HOGE'"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceEnabledConfigDrive = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  config_drive = true
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceDisabledConfigDrive = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  config_drive = false
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceResizeBase = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceResizeUpdate = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-4GB"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceBootFromVolumeWhichHasImageSource = fmt.Sprintf(`
%s

%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  flavor_id = "1CPU-2GB"
  availability_zone = "%s"
  block_device {
    uuid = "${data.ecl_imagestorages_image_v2.image_1.id}"
    source_type = "image"
    volume_size = 15
    boot_index = 0
    destination_type = "volume"
    delete_on_termination = true
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testImageDataSource,
	testCreateNetworkForInstance,
	OS_DEFAULT_ZONE)

var testAccComputeV2InstanceBlockDeviceExistingVolume = fmt.Sprintf(`
%s

%s

resource "ecl_compute_volume_v2" "volume_1" {
  name = "volume_1"
  size = 15
  availability_zone = "%s"
}

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_id = "${data.ecl_imagestorages_image_v2.image_1.id}"
  flavor_id = "1CPU-2GB"
  availability_zone = "%s"
  block_device {
    uuid = "${data.ecl_imagestorages_image_v2.image_1.id}"
    source_type = "image"
    destination_type = "local"
    boot_index = 0
    delete_on_termination = true
  }
  block_device {
    uuid = "${ecl_compute_volume_v2.volume_1.id}"
    source_type = "volume"
    destination_type = "volume"
    boot_index = 1
    delete_on_termination = true
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testImageDataSource,
	testCreateNetworkForInstance,
	OS_DEFAULT_ZONE,
	OS_DEFAULT_ZONE)

var testKeyPairForInstance = `
resource "ecl_compute_keypair_v2" "kp_1" {
  name = "kp_1"
}

resource "ecl_compute_keypair_v2" "kp_2" {
  name = "kp_2"
}`

var testAccComputeV2InstanceKeyPairForceNew1 = fmt.Sprintf(`
%s

%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  key_pair = "${ecl_compute_keypair_v2.kp_1.name}"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`,
	testCreateNetworkForInstance,
	testKeyPairForInstance)

var testAccComputeV2InstanceKeyPairForceNew2 = fmt.Sprintf(`
%s

%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  key_pair = "${ecl_compute_keypair_v2.kp_2.name}"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`,
	testCreateNetworkForInstance,
	testKeyPairForInstance,
)

var testAccComputeV2InstanceBootFromVolumeForceNew1 = fmt.Sprintf(`
%s

%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  flavor_id = "1CPU-2GB"
  availability_zone = "%s"
  block_device {
    uuid = "${data.ecl_imagestorages_image_v2.image_1.id}"
    source_type = "image"
    volume_size = 15
    boot_index = 0
    destination_type = "volume"
    delete_on_termination = true
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testImageDataSource,
	testCreateNetworkForInstance,
	OS_DEFAULT_ZONE,
)

var testAccComputeV2InstanceBootFromVolumeForceNew2 = fmt.Sprintf(`
%s

%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  flavor_id = "1CPU-2GB"
  availability_zone = "%s"
  block_device {
    uuid = "${data.ecl_imagestorages_image_v2.image_1.id}"
    source_type = "image"
    volume_size = 40
    boot_index = 0
    destination_type = "volume"
    delete_on_termination = true
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testImageDataSource,
	testCreateNetworkForInstance,
	OS_DEFAULT_ZONE,
)

var testAccComputeV2InstanceBlockDeviceNewVolume = fmt.Sprintf(`
%s

%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_id = "${data.ecl_imagestorages_image_v2.image_1.id}"
  flavor_id = "1CPU-2GB"
  availability_zone = "%s"
  block_device {
    uuid = "${data.ecl_imagestorages_image_v2.image_1.id}"
    source_type = "image"
    destination_type = "local"
    boot_index = 0
    delete_on_termination = true
  }
  block_device {
    source_type = "blank"
    destination_type = "volume"
    volume_size = 15
    boot_index = 1
    delete_on_termination = true
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`,
	testImageDataSource,
	testCreateNetworkForInstance,
	OS_DEFAULT_ZONE,
)

var testAccComputeV2InstanceAccessIPv4 = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  depends_on = ["ecl_network_subnet_v2.subnet_1"]

  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"

  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
    fixed_ip_v4 = "192.168.1.50"
  }
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceChangeFixedIP1 = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
    fixed_ip_v4 = "192.168.1.10"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceChangeFixedIP2 = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
    fixed_ip_v4 = "192.168.1.20"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceStopBeforeDestroy = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  stop_before_destroy = true
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceMetadataRemove1 = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  metadata = {
    k1 = "v1"
    k2 = "v2"
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceMetadataRemove2 = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  metadata = {
    k3 = "v3"
    k1 = "v1"
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceMetadataRemove3 = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"

  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceTimeout = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"

  timeouts {
    create = "10m"
  }
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceNetworkNameToID = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"

  network {
    name = "${ecl_network_network_v2.network_1.name}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceStateActive = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  power_state = "active"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceStateShutoff = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  power_state = "shutoff"
  network {
    uuid = "${ecl_network_network_v2.network_1.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_1"]
}
`, testCreateNetworkForInstance)

func testAccComputeV2InstanceMultipleNICs() string {
	var result string
	nicMax := 8

	templateForNetwork := `
	resource "ecl_network_network_v2" "network_%d" {
	 name = "network_%d"
	}
	  
	resource "ecl_network_subnet_v2" "subnet_%d" {
	  name = "subnet_%d"
	  network_id = "${ecl_network_network_v2.network_%d.id}"
	  cidr = "192.168.%d.0/24"
	  enable_dhcp = true
	  gateway_ip = "192.168.%d.1"
	}`

	for i := 0; i < nicMax; i++ {
		result += fmt.Sprintf(templateForNetwork, i, i, i, i, i, i, i)
	}

	result += `
	resource "ecl_compute_instance_v2" "instance_1" {
	  name = "instance_1"
	  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
	  flavor_id = "1CPU-2GB"
	  depends_on = [`

	for i := 0; i < nicMax; i++ {
		result += fmt.Sprintf(`"ecl_network_subnet_v2.subnet_%d",`, i)
	}

	result += `]`

	for i := 0; i < nicMax; i++ {
		result += fmt.Sprintf(`
		network {
		  uuid = "${ecl_network_network_v2.network_%d.id}"
		}`, i)
	}

	result += `
    }`

	return result
}

var testAccComputeV2InstanceConnectToCreatedPort = fmt.Sprintf(`
%s

resource "ecl_network_port_v2" "port_1" {
  network_id = "${ecl_network_network_v2.network_1.id}"
  fixed_ip {
    subnet_id = "${ecl_network_subnet_v2.subnet_1.id}"
    ip_address = "192.168.1.50"
  }
}

resource "ecl_compute_instance_v2" "instance_1" {
  depends_on = ["ecl_network_subnet_v2.subnet_1"]

  name = "instance_1"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"

  network {
    port = "${ecl_network_port_v2.port_1.id}"
  }
}
`, testCreateNetworkForInstance)

var testAccComputeV2InstanceConnectToNetworkByName = fmt.Sprintf(`
%s

resource "ecl_compute_instance_v2" "instance_1" {
  depends_on = ["ecl_network_subnet_v2.subnet_1"]

  name = "instance_1"
  flavor_id = "1CPU-2GB"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"

  network {
    name = "${ecl_network_network_v2.network_1.name}"
  }
}
`, testCreateNetworkForInstance)
