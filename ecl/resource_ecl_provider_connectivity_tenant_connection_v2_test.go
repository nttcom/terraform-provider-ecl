package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/provider_connectivity/v2/tenant_connections"
)

func TestAccProviderConnectivityV2TenantConnection_basic(t *testing.T) {
	var connection tenant_connections.TenantConnection

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProviderConnectivityV2TenantConnectionAttachmentComputeServer,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProviderConnectivityV2TenantConnectionExists("ecl_provider_connectivity_tenant_connection_v2.connection_1", &connection),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "test_name1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "test_desc1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.test_tags1", "test1"),
					resource.TestCheckResourceAttrSet("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_connection_request_id"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_type", "ECL::Compute::Server"),
					resource.TestCheckResourceAttrSet("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_id"),
				),
			},
			{
				Config: testAccProviderConnectivityV2TenantConnectionUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "updated_name"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "updated_desc"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.k2", "v2"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProviderConnectivityV2TenantConnectionAttachmentBaremetalServer,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProviderConnectivityV2TenantConnectionExists("ecl_provider_connectivity_tenant_connection_v2.connection_1", &connection),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "test_name1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "test_desc1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.test_tags1", "test1"),
					resource.TestCheckResourceAttrSet("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_connection_request_id"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_type", "ECL::Baremetal::Server"),
					resource.TestCheckResourceAttrSet("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_id"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProviderConnectivityV2TenantConnectionAttachmentVna,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProviderConnectivityV2TenantConnectionExists("ecl_provider_connectivity_tenant_connection_v2.connection_1", &connection),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "name", "test_name1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "description", "test_desc1"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tags.test_tags1", "test1"),
					resource.TestCheckResourceAttrSet("ecl_provider_connectivity_tenant_connection_v2.connection_1", "tenant_connection_request_id"),
					resource.TestCheckResourceAttr("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_type", "ECL::VirtualNetworkAppliance::VSRX"),
					resource.TestCheckResourceAttrSet("ecl_provider_connectivity_tenant_connection_v2.connection_1", "device_id"),
				),
			},
		},
	})
}

func TestAccProviderConnectivityV2TenantConnection_conflictAttachmentOpts(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccProviderConnectivityV2TenantConnectionConflictAttachmentOpts,
				ExpectError: regexp.MustCompile("\"attachment_opts_server\": conflicts with attachment_opts_vna"),
			},
			resource.TestStep{
				Config:      testAccProviderConnectivityV2TenantConnectionConflictAttachmentOpts,
				ExpectError: regexp.MustCompile("\"attachment_opts_vna\": conflicts with attachment_opts_server"),
			},
		},
	})
}

func TestAccProviderConnectivityV2TenantConnection_multipleListAttachmentOpts(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckTenantConnection(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccProviderConnectivityV2TenantConnectionAttachmentOptsServerMultipleList,
				ExpectError: regexp.MustCompile("Too many attachment_opts_server blocks: No more than 1 \"attachment_opts_server\" blocks are allowed"),
			},
			resource.TestStep{
				Config:      testAccProviderConnectivityV2TenantConnectionAttachmentOptsVnaMultipleList,
				ExpectError: regexp.MustCompile("Too many attachment_opts_vna blocks: No more than 1 \"attachment_opts_vna\" blocks are allowed"),
			},
		},
	})
}

func TestAccProviderConnectivityV2TenantConnection_NotSpecifyDeviceInterfaceID(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccProviderConnectivityV2TenantConnectionAttachmentBaremetalNotSpecifyDeviceInterfaceID,
				ExpectError: regexp.MustCompile("device_type is For ECL::Baremetal::Server or ECL::VirtualNetworkAppliance::VSRX, " +
					"device_interface_id is required"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProviderConnectivityV2TenantConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccProviderConnectivityV2TenantConnectionAttachmentVnaNotSpecifyDeviceInterfaceID,
				ExpectError: regexp.MustCompile("device_type is For ECL::Baremetal::Server or ECL::VirtualNetworkAppliance::VSRX, " +
					"device_interface_id is required"),
			},
		},
	})
}

func testAccCheckProviderConnectivityV2TenantConnectionDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.providerConnectivityV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL Provider Connectivity client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_provider_connectivity_tenant_connection_v2" {
			continue
		}

		if _, err := tenant_connections.Get(client, rs.Primary.ID).Extract(); err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				continue
			}

			return fmt.Errorf("error getting Tenent Connection: %s", err)
		}

		return fmt.Errorf("tenent connection still exists")
	}

	return nil
}

func testAccCheckProviderConnectivityV2TenantConnectionExists(n string, request *tenant_connections.TenantConnection) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.providerConnectivityV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL Provider Connectivity client: %s", err)
		}

		found, err := tenant_connections.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Tenent Connection not found")
		}

		*request = *found

		return nil
	}
}

const oppositeTenantNetwork = `
resource "ecl_network_network_v2" "network_1" {
	provider = "ecl_accepter"
	name = "network_1"
}

resource "ecl_network_subnet_v2" "subnet_1" {
	provider = "ecl_accepter"
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

const tenantNetwork = `
resource "ecl_network_network_v2" "network_2" {
	name = "network_2"
}

resource "ecl_network_subnet_v2" "subnet_2" {
	name = "subnet_2"
	cidr = "192.168.2.0/24"
	network_id = "${ecl_network_network_v2.network_2.id}"
	gateway_ip = "192.168.2.1"
	allocation_pools {
		start = "192.168.2.100"
		end = "192.168.2.200"
	}
}
`

const attachmentComputeServer = `
resource "ecl_compute_instance_v2" "instance_1" {
  name = "i"
  image_name = "Ubuntu-18.04.1_64_virtual-server_02"
  flavor_id = "1CPU-2GB"
  metadata = {
    foo = "bar"
  }
  network {
    uuid = "${ecl_network_network_v2.network_2.id}"
  }
  depends_on = ["ecl_network_subnet_v2.subnet_2"]
}
`

const attachmentBaremetalServer = `
data "ecl_imagestorages_image_v2" "centos" {
    name = "CentOS-7.3-1611_64_baremetal-server_01"
}

data "ecl_baremetal_flavor_v2" "gp2" {
	name = "General Purpose 2 v1"
}

data "ecl_baremetal_availability_zone_v2" "groupa" {
    zone_name = "groupa"
}

resource "ecl_baremetal_server_v2" "server_1" {
    depends_on = [
        "ecl_network_subnet_v2.subnet_2"
    ]

    name = "server1"
    image_id = "${data.ecl_imagestorages_image_v2.centos.id}"
    flavor_id = "${data.ecl_baremetal_flavor_v2.gp2.id}"
    user_data = "user_data"
    availability_zone = "${data.ecl_baremetal_availability_zone_v2.groupa.zone_name}"
    admin_pass = "password"
    metadata = {
        k1 = "v1"
        k2 = "v2"
    }
    networks {
        uuid = "${ecl_network_network_v2.network_2.id}"
        fixed_ip = "192.168.2.10"
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
`

var attachmentVna = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.2.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_2"]
	interface_1_info {
		name = "interface_1"
		network_id = "${ecl_network_network_v2.network_2.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.2.10"
	}
}
`,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID)

var testAccProviderConnectivityV2TenantConnectionAttachmentComputeServer = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
		"ecl_compute_instance_v2.instance_1",
		"ecl_sss_approval_request_v1.approval_1"
	]
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::Compute::Server"
	device_id = "${ecl_compute_instance_v2.instance_1.id}"
	attachment_opts_server {
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentComputeServer,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionUpdate = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
		"ecl_compute_instance_v2.instance_1",
		"ecl_sss_approval_request_v1.approval_1"
	]
	name = "updated_name"
	description = "updated_desc"
	tags = {
		"k2" = "v2"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::Compute::Server"
	device_id = "${ecl_compute_instance_v2.instance_1.id}"
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentComputeServer,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionAttachmentBaremetalServer = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
		"ecl_baremetal_server_v2.server_1",
		"ecl_sss_approval_request_v1.approval_1"
	]
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::Baremetal::Server"
	device_id = "${ecl_baremetal_server_v2.server_1.id}"
	device_interface_id = "${ecl_baremetal_server_v2.server_1.nic_physical_ports.0.network_physical_port_id}"
	attachment_opts_server {
		segmentation_type = "flat"
		segmentation_id = 10
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentBaremetalServer,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionAttachmentVna = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
			"ecl_vna_appliance_v1.appliance_1",
			"ecl_sss_approval_request_v1.approval_1"
		]
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::VirtualNetworkAppliance::VSRX"
	device_id = "${ecl_vna_appliance_v1.appliance_1.id}"
 	device_interface_id = "interface_2"
	attachment_opts_vna {
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentVna,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionConflictAttachmentOpts = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
		"ecl_compute_instance_v2.instance_1",
		"ecl_sss_approval_request_v1.approval_1"
	]
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::Compute::Server"
	device_id = "${ecl_compute_instance_v2.instance_1.id}"
	attachment_opts_server {
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
	attachment_opts_vna {
		fixed_ips {
			ip_address = "192.168.1.2"
		}
	}
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentComputeServer,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionAttachmentOptsServerMultipleList = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
		"ecl_compute_instance_v2.instance_1",
		"ecl_sss_approval_request_v1.approval_1"
	]
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::Compute::Server"
	device_id = "${ecl_compute_instance_v2.instance_1.id}"
	attachment_opts_server {
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
	attachment_opts_server {
		fixed_ips {
			ip_address = "192.168.1.2"
		}
	}
	
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentComputeServer,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionAttachmentOptsVnaMultipleList = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
			"ecl_vna_appliance_v1.appliance_1",
			"ecl_sss_approval_request_v1.approval_1"
		]
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::VirtualNetworkAppliance::VSRX"
	device_id = "${ecl_vna_appliance_v1.appliance_1.id}"
 	device_interface_id = "interface_2"
	attachment_opts_vna {
		fixed_ips {
			ip_address = "192.168.1.1"
		}
	}
	attachment_opts_vna {
		fixed_ips {
			ip_address = "192.168.1.2"
		}
	}
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentVna,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionAttachmentBaremetalNotSpecifyDeviceInterfaceID = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
		"ecl_baremetal_server_v2.server_1",
		"ecl_sss_approval_request_v1.approval_1"
	]
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::Baremetal::Server"
	device_id = "${ecl_baremetal_server_v2.server_1.id}"
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentBaremetalServer,
	OS_ACCEPTER_TENANT_ID)

var testAccProviderConnectivityV2TenantConnectionAttachmentVnaNotSpecifyDeviceInterfaceID = fmt.Sprintf(`
%s

%s

%s

resource "ecl_provider_connectivity_tenant_connection_request_v2" "request_1" {
	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tenant_id_other = "%s"
	network_id = "${ecl_network_network_v2.network_1.id}"
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
}

resource "ecl_sss_approval_request_v1" "approval_1" {
	depends_on = ["ecl_provider_connectivity_tenant_connection_request_v2.request_1"]
	request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.approval_request_id}"
	status = "approved"
}

resource "ecl_provider_connectivity_tenant_connection_v2" "connection_1" {
	depends_on = [
		"ecl_vna_appliance_v1.appliance_1",
		"ecl_sss_approval_request_v1.approval_1"
	]
	name = "test_name1"
	description = "test_desc1"
	tags = {
		"test_tags1" = "test1"
	}
	tenant_connection_request_id = "${ecl_provider_connectivity_tenant_connection_request_v2.request_1.id}"
	device_type = "ECL::VirtualNetworkAppliance::VSRX"
	device_id = "${ecl_vna_appliance_v1.appliance_1.id}"
}
`,
	oppositeTenantNetwork,
	tenantNetwork,
	attachmentVna,
	OS_ACCEPTER_TENANT_ID)
