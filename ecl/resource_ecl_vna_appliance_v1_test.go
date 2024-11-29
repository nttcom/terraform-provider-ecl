package ecl

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v3/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/subnets"
	"github.com/nttcom/eclcloud/v3/ecl/vna/v1/appliances"
)

var MaxLengthString = repeatedString("a", 255)

// Test process -> PASSED
// 1. create vna
// 2. set each metadata by max length strings
// 3. set those values as blank
func TestAccVNAV1Appliance_updateMetaBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n networks.Network
	var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.subnet_id", &sn.ID),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceUpdateMetaBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", MaxLengthString),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", MaxLengthString),
					testAccCheckVNAV1ApplianceTag(&vna, "k1", MaxLengthString),
					testAccCheckVNAV1ApplianceTag(&vna, "k2", MaxLengthString),

					testAccCheckVNAV1ApplianceInterfaceTag(&vna.Interfaces.Interface1, "interfaceK1", MaxLengthString),
					testAccCheckVNAV1ApplianceInterfaceTag(&vna.Interfaces.Interface1, "interfaceK2", MaxLengthString),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceUpdateMetaBasic2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", ""),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", ""),
					testAccCheckVNAV1ApplianeTagLengthIsZERO(&vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", ""),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", ""),

					// TODO interface tag can not be set as blank object if set some key-value pairs once.
					// testAccCheckVNAV1ApplianceInterfaceTagLengthIsZERO(&vna.Interfaces.Interface1),
				),
			},
		},
	})
}

// Test process -> PASSED
// 1. create vna
// 2. update name and description metadata
func TestAccVNAV1Appliance_updateMetaWithoutInterface(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n networks.Network
	var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.subnet_id", &sn.ID),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceUpdateMetaWithoutInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", MaxLengthString),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", MaxLengthString),
				),
			},
		},
	})
}

// Test process -> Passed
// 1. create vna
// 2. connect interface2 with network-2
// 3. disconnect interface2 from network-2
func TestAccVNAV1Appliance_connectAndDisconnectInterface(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n, n2 networks.Network
	var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceConnectInterface2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_2_info.0.network_id", &n2.ID),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceDisconnectInterface2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_2_info.0.network_id", ""),
				),
			},
		},
	})
}

// Test process -> Passed
// 1. create vna
// 2. set allowed address pairs which has type of "VRRP" and VRID=123
// 3. set(change and over write) allowed address pairs which has type of "" and VRID is "null"
// 4. unset allowed address pairs and check if length is correctly set as 0
func TestAccVNAV1Appliance_updateAllowedAddressPairBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n networks.Network

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceUpdateAllowedAddressPairVRRP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.ip_address", "192.168.1.200"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.type", "vrrp"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.vrid", "123"),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceUpdateAllowedAddressPairNoType,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.ip_address", "192.168.1.200"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.type", ""),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.vrid", "null"),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceUpdateNoAllowedAddressPair,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AllowedAddressPairLength(&vna, 1, 0),
				),
			},
		},
	})
}

func TestAccVNAV1Appliance_updateFixedIPBasic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n, n2, n3, n4 networks.Network
	var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.subnet_id", &sn.ID),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceUpdateFixedIPBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_4", &n4),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_2_info.0.network_id", &n2.ID),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.network_id", &n3.ID),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_4_info.0.network_id", &n4.ID),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 1, "192.168.1.50"),

					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 3, "192.168.3.50"),

					testAccCheckVNAV1FixedIPLength(&vna, 4, 0),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceRemoveFixedIP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_2_info.0.network_id", &n2.ID),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.network_id", &n3.ID),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_4_info.0.network_id", &n4.ID),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 1, "192.168.1.50"),
					testAccCheckVNAV1FixedIPLength(&vna, 2, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 3, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 4, 0),
				),
			},
		},
	})
}

func TestAccVNAV1Appliance_createInterfaceDiscontinuity(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n3, n8 networks.Network
	var sn3, sn8 subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceInterfaceDiscontinuity,
				Check: resource.ComposeTestCheckFunc(

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_8", &n8),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_3", &sn3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_8", &sn8),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "username", "root"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.name", "interface_3"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.description", "interface_3_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.network_id", &n3.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_fixed_ips.0.ip_address", "192.168.3.50"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_3_fixed_ips.0.subnet_id", &sn3.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_8_info.0.name", "interface_8"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_8_info.0.description", "interface_8_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_8_info.0.network_id", &n8.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_8_fixed_ips.0.ip_address", "192.168.8.50"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_8_fixed_ips.0.subnet_id", &sn8.ID),
					testAccCheckVNAV1FixedIPLength(&vna, 1, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 2, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 3, 1),
					testAccCheckVNAV1FixedIPLength(&vna, 4, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 5, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 6, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 7, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 8, 1),
				),
			},
		},
	})
}

func TestAccVNAV1Appliance_createNoInterface(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceNoInterface,
				Check: resource.ComposeTestCheckFunc(

					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "username", "root"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					testAccCheckVNAV1FixedIPLength(&vna, 1, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 2, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 3, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 4, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 5, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 6, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 7, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 8, 0),
				),
			},
		},
	})
}

func TestAccVNAV1Appliance_createFixedIPsEmpty(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n networks.Network
	var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceFixedIPsEmpty,
				Check: resource.ComposeTestCheckFunc(

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "username", "root"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					testAccCheckVNAV1FixedIPLength(&vna, 1, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 2, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 3, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 4, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 5, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 6, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 7, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 8, 0),
				),
			},
		},
	})
}

func TestAccVNAV1Appliance_createNoFixedIPs(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n networks.Network
	var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceNoFixedIPs,
				Check: resource.ComposeTestCheckFunc(

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "username", "root"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					testAccCheckVNAV1FixedIPLength(&vna, 1, 1),
					testAccCheckVNAV1FixedIPLength(&vna, 2, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 3, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 4, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 5, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 6, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 7, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 8, 0),
				),
			},
		},
	})
}

func TestAccVNAV1Appliance_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var vna appliances.Appliance
	var n networks.Network
	var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(

					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.subnet_id", &sn.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "initial_config.format", "set"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "initial_config.data", "c2V0IGludGVyZmFjZXMgZ2UtMC8wLzAgZGVzY3JpcHRpb24gc2FtcGxl"),
					testAccCheckVNAV1FixedIPLength(&vna, 1, 1),
					testAccCheckVNAV1FixedIPLength(&vna, 2, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 3, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 4, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 5, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 6, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 7, 0),
					testAccCheckVNAV1FixedIPLength(&vna, 8, 0),
				),
			},
		},
	})
}

func testAccCheckVNAV1ApplianceInterfaceTag(
	vnaIF *appliances.InterfaceInResponse, k string, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if vnaIF.Tags == nil {
			return fmt.Errorf("No tag")
		}

		for key, value := range vnaIF.Tags {
			if k != key {
				continue
			}

			if v == value {
				return nil
			}

			return fmt.Errorf("Bad value for %s: %s", k, value)
		}

		return fmt.Errorf("Tag not found: %s", k)
	}
}

func testAccCheckVNAV1ApplianceInterfaceTagLengthIsZERO(
	vnaIF *appliances.InterfaceInResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(vnaIF.Tags) != 0 {
			return fmt.Errorf("Interface Tag length is not ZERO")
		}

		return nil
	}
}

func testAccCheckVNAV1ApplianceTag(
	vna *appliances.Appliance, k string, v string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if vna.Tags == nil {
			return fmt.Errorf("No tag")
		}

		for key, value := range vna.Tags {
			if k != key {
				continue
			}

			if v == value {
				return nil
			}

			return fmt.Errorf("Bad value for %s: %s", k, value)
		}

		return fmt.Errorf("Tag not found: %s", k)
	}
}

func testAccCheckVNAV1ApplianeTagLengthIsZERO(ap *appliances.Appliance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(ap.Tags) != 0 {
			return fmt.Errorf("Tag length is not ZERO")
		}
		return nil
	}
}

// In current VNA implementation, order of each element of fixed_ips in response
// will sometimes changed.
// This function checks whether VNA has specified IP address
// regardless of position of ip address element.
func testAccCheckVNAV1InterfaceHasIPAddress(
	vna *appliances.Appliance,
	slotNumber int,
	expectedAddress string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualFixedIPs := getFixedIPsBySlotNumber(vna, slotNumber)
		for _, fixedIP := range actualFixedIPs {
			if fixedIP.IPAddress == expectedAddress {
				log.Printf("[DEBUG] Comparison between fixedIP.IPAddress <=> expectedAddress : %s <=> %s", fixedIP.IPAddress, expectedAddress)
				return nil
			}
		}
		return fmt.Errorf("Virtual Network Appliance does not have expected IP address: %s", expectedAddress)
	}
}

func testAccCheckVNAV1AllowedAddressPairLength(
	vna *appliances.Appliance,
	slotNumber int,
	expectedLength int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualAAPs := getAllowedAddressPairsBySlotNumber(vna, slotNumber)

		actualLength := len(actualAAPs)
		if actualLength != expectedLength {
			return fmt.Errorf(
				"Length of Allowed Address Pairs list is different. expected %d, actual %d",
				expectedLength,
				actualLength,
			)
		}
		return nil
	}
}

func testAccCheckVNAV1FixedIPLength(
	vna *appliances.Appliance,
	slotNumber int,
	expectedLength int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualFixedIPs := getFixedIPsBySlotNumber(vna, slotNumber)

		actualLength := len(actualFixedIPs)
		if actualLength != expectedLength {
			return fmt.Errorf(
				"Length of FixedIPs list is different. expected %d, actual %d",
				expectedLength,
				actualLength,
			)
		}
		return nil
	}
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

const testAccVNAV1ApplianceSingleNetworkAndSubnetPair2 = `
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

const testAccVNAV1ApplianceSingleNetworkAndSubnetPair3 = `
resource "ecl_network_network_v2" "network_3" {
	name = "network_3"
}

resource "ecl_network_subnet_v2" "subnet_3" {
	name = "subnet_3"
	cidr = "192.168.3.0/24"
	network_id = "${ecl_network_network_v2.network_3.id}"
	gateway_ip = "192.168.3.1"
	allocation_pools {
		start = "192.168.3.100"
		end = "192.168.3.200"
	}
}
`
const testAccVNAV1ApplianceSingleNetworkAndSubnetPair4 = `
resource "ecl_network_network_v2" "network_4" {
	name = "network_4"
}

resource "ecl_network_subnet_v2" "subnet_4" {
	name = "subnet_3"
	cidr = "192.168.4.0/24"
	network_id = "${ecl_network_network_v2.network_4.id}"
	gateway_ip = "192.168.4.1"
	allocation_pools {
		start = "192.168.4.100"
		end = "192.168.4.200"
	}
}
`
const testAccVNAV1ApplianceSingleNetworkAndSubnetPair8 = `
resource "ecl_network_network_v2" "network_8" {
	name = "network_8"
}

resource "ecl_network_subnet_v2" "subnet_8" {
	name = "subnet_8"
	cidr = "192.168.8.0/24"
	network_id = "${ecl_network_network_v2.network_8.id}"
	gateway_ip = "192.168.8.1"
	allocation_pools {
		start = "192.168.8.100"
		end = "192.168.8.200"
	}
}
`

var testAccVNAV1ApplianceBasic = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	initial_config = {
		"format" = "set"
		"data" = "c2V0IGludGVyZmFjZXMgZ2UtMC8wLzAgZGVzY3JpcHRpb24gc2FtcGxl"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceInterfaceDiscontinuity = fmt.Sprintf(`
%s
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.3.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = [
		"ecl_network_subnet_v2.subnet_3",
		"ecl_network_subnet_v2.subnet_8"
	]
	tags = {
		k1 = "v1"
	}

	interface_3_info  {
		name = "interface_3"
		description = "interface_3_description"
		network_id = "${ecl_network_network_v2.network_3.id}"
	}

	interface_3_fixed_ips {
		ip_address = "192.168.3.50"
	}

	interface_8_info  {
		name = "interface_8"
		description = "interface_8_description"
		network_id = "${ecl_network_network_v2.network_8.id}"
	}

	interface_8_fixed_ips {
		ip_address = "192.168.8.50"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair3,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair8,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceNoInterface = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	tags = {
		k1 = "v1"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceFixedIPsEmpty = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_no_fixed_ips = "true"

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceNoFixedIPs = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateFixedIPBasic = fmt.Sprintf(`
%s
%s
%s
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = [
		"ecl_network_subnet_v2.subnet_1",
		"ecl_network_subnet_v2.subnet_2",
		"ecl_network_subnet_v2.subnet_3",
		"ecl_network_subnet_v2.subnet_4"
	]

	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_2_info  {
		network_id = "${ecl_network_network_v2.network_2.id}"
	}

	interface_3_info  {
		network_id = "${ecl_network_network_v2.network_3.id}"
	}

	interface_3_fixed_ips {
		ip_address = "192.168.3.50"
	}

	interface_4_info {
		network_id = "${ecl_network_network_v2.network_4.id}"
	}

	interface_4_no_fixed_ips = "true"

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair2,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair3,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair4,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)
var testAccVNAV1ApplianceRemoveFixedIP = fmt.Sprintf(`
%s
%s
%s
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = [
		"ecl_network_subnet_v2.subnet_1",
		"ecl_network_subnet_v2.subnet_2",
		"ecl_network_subnet_v2.subnet_3",
		"ecl_network_subnet_v2.subnet_4"
	]

	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_2_info  {
		network_id = "${ecl_network_network_v2.network_2.id}"
	}

	interface_2_no_fixed_ips = "true"

	interface_3_info  {
		network_id = "${ecl_network_network_v2.network_3.id}"
	}

	interface_3_no_fixed_ips = "true"

	interface_4_info {
		network_id = "${ecl_network_network_v2.network_4.id}"
	}

	interface_4_no_fixed_ips = "true"

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair2,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair3,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair4,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateMetaBasic = fmt.Sprintf(`
%[1]s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "%[3]s"
	description = "%[3]s"
	default_gateway = "192.168.1.1"
	availability_zone = "%[4]s"
	virtual_network_appliance_plan_id = "%[2]s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tags = {
		k1 = "%[3]s"
		k2 = "%[3]s"
	}

	interface_1_info  {
		name = "%[3]s"
		description = "%[3]s"
		network_id = "${ecl_network_network_v2.network_1.id}"
		tags = {
			interfaceK1 = "%[3]s"
			interfaceK2 = "%[3]s"
		}
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
	MaxLengthString,
	OS_DEFAULT_ZONE,
)

var testAccVNAV1ApplianceUpdateMetaBasic2 = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = ""
	description = ""
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]

	interface_1_info  {
		name = ""
		description = ""
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateMetaWithoutInterface = fmt.Sprintf(`
%[1]s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "%[3]s"
	description = "%[3]s"
	default_gateway = "192.168.1.1"
	availability_zone = "%[4]s"
	virtual_network_appliance_plan_id = "%[2]s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
	MaxLengthString,
	OS_DEFAULT_ZONE,
)

var testAccVNAV1ApplianceUpdateAllowedAddressPairVRRP = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_1_allowed_address_pairs {
		ip_address = "192.168.1.200"
		mac_address = ""
		type = "vrrp"
		vrid = "123"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateAllowedAddressPairNoType = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_1_allowed_address_pairs {
		ip_address = "192.168.1.200"
		mac_address = "aa:bb:cc:dd:ee:f1"
		type = ""
		vrid = "null"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateNoAllowedAddressPair = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_1_no_allowed_address_pairs = "true"

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceConnectInterface2 = fmt.Sprintf(`
%s
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = [
		"ecl_network_subnet_v2.subnet_1",
		"ecl_network_subnet_v2.subnet_2"
	]

	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_2_info  {
		network_id = "${ecl_network_network_v2.network_2.id}"
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair2,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceDisconnectInterface2 = fmt.Sprintf(`
%s
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	depends_on = [
		"ecl_network_subnet_v2.subnet_1",
		"ecl_network_subnet_v2.subnet_2"
	]

	tags = {
		k1 = "v1"
	}

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "${ecl_network_network_v2.network_1.id}"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_2_info  {
		network_id = ""
	}

	lifecycle {
		ignore_changes = [
			"username",
			"password",
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair2,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)
