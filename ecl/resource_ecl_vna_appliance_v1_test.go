package ecl

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/ecl/network/v2/subnets"
	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

func TestAccVNAV1ApplianceConnectAndDisconnectInterface(t *testing.T) {
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
					// Create resource reference
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceConnectInterface2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1InterfaceHasConnectionWithNetwork(&vna, 2, &n2),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceDisconnectInterface2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1InterfaceHasNoConnection(&vna, 2),
				),
			},
		},
	})
}

// Test process
// 1. create vna
// 2. set allowed address pairs which has type of "VRRP" and VRID=123
// 3. set(change and over write) allowed address pairs which has type of "" and VRID is "null"
// 4. unset allowed address pairs and check if length is correctly set as 0
func TestAccVNAV1ApplianceUpdateAllowedAddressPairBasic(t *testing.T) {
	var vna appliances.Appliance
	var n networks.Network
	// var sn subnets.Subnet

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
					testAccCheckVNAV1AllowedAddressPairs(
						&vna, 1,
						"192.168.1.200", "vrrp", "123",
					),
				),
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceUpdateAllowedAddressPairNoType,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AllowedAddressPairs(
						&vna, 1,
						"192.168.1.200", "", "null",
					),
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

func TestAccVNAV1ApplianceUpdateFixedIPBasic(t *testing.T) {
	var vna appliances.Appliance
	var n, n2, n3 networks.Network
	var sn subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					// Create resource reference
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					// Check about meta
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
					// Create resource reference
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					// Check network id in interface metadata part
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_2_info.0.network_id", &n2.ID),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.network_id", &n3.ID),
					// Check fixed_ips part
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 1, "192.168.1.50"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 2, "192.168.2.101"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 3, "192.168.3.50"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 3, "192.168.3.60"),
					testAccCheckVNAV1FixedIPLength(&vna, 4, 0),
				),
			},
		},
	})
}

func TestAccVNAV1ApplianceUpdateMetaBasic(t *testing.T) {
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
					// Create resource reference
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					// Check about meta
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
					// Create resource reference
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					// Check about meta
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description-update"),
					testAccCheckVNAV1ApplianceTag(&vna, "k1", "v1"),
					testAccCheckVNAV1ApplianceTag(&vna, "k2", "v2"),
					// Check interface meta
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description-update"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.subnet_id", &sn.ID),
					testAccCheckVNAV1ApplianceInterfaceTag(&vna.Interfaces.Interface1, "interfaceK1", "interfaceV1"),
					testAccCheckVNAV1ApplianceInterfaceTag(&vna.Interfaces.Interface1, "interfaceK2", "interfaceV2"),
				),
			},
		},
	})
}

func TestAccVNAV1ApplianceBasic(t *testing.T) {
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
					// Create resource reference
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1", &sn),
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					// Check about meta
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", &n.ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttrPtr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.subnet_id", &sn.ID),
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

func testAccCheckVNAV1InterfaceHasNoConnection(
	vna *appliances.Appliance,
	slotNumber int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualInterface := getInterfaceBySlotNumber(vna, slotNumber)
		actualNetworkID := actualInterface.NetworkID

		if actualNetworkID == "" {
			return nil
		}

		return fmt.Errorf(
			"Virtual Network Appliance has connection with %s on contrally to expectation",
			actualNetworkID,
		)
	}
}

func testAccCheckVNAV1InterfaceHasConnectionWithNetwork(
	vna *appliances.Appliance,
	slotNumber int,
	expectedNetwork *networks.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualInterface := getInterfaceBySlotNumber(vna, slotNumber)
		actualNetworkID := actualInterface.NetworkID

		if expectedNetwork.ID == actualNetworkID {
			return nil
		}
		return fmt.Errorf(
			"Virtual Network Appliance does not have connection with %s. Actually is connected with %s",
			expectedNetwork.ID, actualNetworkID,
		)
	}
}

func testAccCheckVNAV1InterfaceHasIPAddress(
	vna *appliances.Appliance,
	slotNumber int,
	expectedAddress string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actualFixedIPs := getFixedIPsBySlotNumber(vna, slotNumber)
		for _, fixedIP := range actualFixedIPs {
			if fixedIP.IPAddress == expectedAddress {
				return nil
			}
		}
		return fmt.Errorf("Virtual Network Appliance does not have expected IP address: %s", expectedAddress)
	}
}

func testAccCheckVNAV1AllowedAddressPairs(
	vna *appliances.Appliance,
	slotNumber int,
	expectedIPAddress string,
	expectedType string,
	expectedVRID string,
) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		var thisIPAddress, thisType, thisVRID string
		actualAllowedAddressPairs := getAllowedAddressPairsBySlotNumber(vna, slotNumber)
		log.Printf("[DEBUG] Start checking this Allowed Address Pair: %#v", actualAllowedAddressPairs)
		// var success bool

		for _, aap := range actualAllowedAddressPairs {
			log.Printf("[DEBUG] Start checking each element of Allowed Address Pair: %#v", aap)
			thisIPAddress = aap.IPAddress

			log.Printf("[MYDEBUG] aap.VRID is : %#v", aap.VRID)
			if aap.VRID == interface{}(nil) {
				thisVRID = "null"
				log.Printf("[MYDEBUG] thisVRID(if) %#v", thisVRID)
			} else {
				thisVRID = strconv.Itoa(int(aap.VRID.(float64)))
				log.Printf("[MYDEBUG] thisVRID(else) %#v", thisVRID)
			}
			thisType = aap.Type

			log.Printf(
				"[MYDEBUG] aap actual - IP, VRID, Type: %s %s %s",
				thisIPAddress, thisVRID, thisType,
			)

			// MACAddress is auto assigned value in case type = "vrrp"
			// so is not possible to use as one of assertion conditions.
			if thisIPAddress == expectedIPAddress &&
				thisVRID == expectedVRID &&
				thisType == expectedType {
				log.Printf(
					"[DEBUG] actual and expected allowed address pairs are completely matched. "+
						"thisIPAddress <=> expectedIPAddress = %s <=> %s -- "+
						"thisType <=> expectedType = %s <=> %s -- "+
						"thisVRID <=> expectedVRID = %s <=> %s",
					thisIPAddress, expectedIPAddress,
					thisType, expectedType,
					thisVRID, expectedVRID,
				)
				return nil
			}
		}

		return fmt.Errorf(
			"Virtual Network Appliance does not have expected allowed address pairs: %s %s %s",
			thisIPAddress, thisVRID, thisType,
		)
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
		log.Printf("[MYDEBUG] VNA in existence check: %#v", vna)

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

var testAccVNAV1ApplianceBasic = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
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

	lifecycle {
		ignore_changes = [
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
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
	availability_zone = "zone1-groupb"
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

    interface_3_fixed_ips {
		ip_address = "192.168.3.60"
	}

	interface_4_info {
		network_id = "${ecl_network_network_v2.network_4.id}"
	}

	interface_4_no_fixed_ips = "true"

	lifecycle {
		ignore_changes = [
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair2,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair3,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair4,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateMetaBasic = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1-update"
	description = "appliance_1_description-update"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
	virtual_network_appliance_plan_id = "%s"

	depends_on = ["ecl_network_subnet_v2.subnet_1"]
    tags = {
        k1 = "v1"
        k2 = "v2"
    }

	interface_1_info  {
		name = "interface_1-update"
		description = "interface_1_description-update"
		network_id = "${ecl_network_network_v2.network_1.id}"
		tags = {
			interfaceK1 = "interfaceV1"
			interfaceK2 = "interfaceV2"
		}
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	lifecycle {
		ignore_changes = [
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateAllowedAddressPairVRRP = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
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
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateAllowedAddressPairNoType = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
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
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceUpdateNoAllowedAddressPair = fmt.Sprintf(`
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
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
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceConnectInterface2 = fmt.Sprintf(`
%s
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
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
		name = "interface_2"
		description = "interface_2_description"
		network_id = "${ecl_network_network_v2.network_2.id}"
	}
	
	lifecycle {
		ignore_changes = [
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair2,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testAccVNAV1ApplianceDisconnectInterface2 = fmt.Sprintf(`
%s
%s

resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "zone1-groupb"
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
		name = "interface_2"
		description = "interface_2_description"
		network_id = ""
	}
	
	lifecycle {
		ignore_changes = [
			"default_gateway",
		]
	}
}`,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair,
	testAccVNAV1ApplianceSingleNetworkAndSubnetPair2,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)
