package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"

	"github.com/nttcom/eclcloud/v3/ecl/vna/v1/appliances"
)

func TestMockedAccVNAV1Appliance_updateMetaBasic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceMetaPatch1)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterMetaUpdate1)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterMetaUpdate1)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceMetaPatch2)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterMetaUpdate2)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterMetaUpdate2)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
				),
			},
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceUpdateMetaBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description-update"),
					testAccCheckVNAV1ApplianceTag(&vna, "k1", "v1"),
					testAccCheckVNAV1ApplianceTag(&vna, "k2", "v2"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description-update"),
					testAccCheckVNAV1ApplianceInterfaceTag(&vna.Interfaces.Interface1, "interfaceK1", "interfaceV1"),
					testAccCheckVNAV1ApplianceInterfaceTag(&vna.Interfaces.Interface1, "interfaceK2", "interfaceV2"),
				),
			},
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceUpdateMetaBasic2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description-update"),

					testAccCheckVNAV1ApplianeTagLengthIsZERO(&vna),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1-update"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description-update"),

					// TODO
					// Even this test is successfully passed, actual resource test will be failed.
					testAccCheckVNAV1ApplianceInterfaceTagLengthIsZERO(&vna.Interfaces.Interface1),
				),
			},
		},
	})
}

func TestMockedAccVNAV1Appliance_updateAllowedAddressPairBasic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceAllowedAddressPairPatch)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterAllowedAddressPairUpdate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterAllowedAddressPairUpdate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
				),
			},
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceUpdateAllowedAddressPairBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.ip_address", "192.168.1.200"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.type", "vrrp"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_allowed_address_pairs.0.vrid", "123"),
				),
			},
		},
	})
}

func TestMockedAccVNAV1ApplianceUpdateFixedIP_basic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceFixedIPPatch)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterFixedIPUpdate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterFixedIPUpdate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
				),
			},
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceUpdateFixedIPBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "name", "appliance_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),

					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", "dummyNetworkID"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_2_info.0.network_id", "dummyNetworkID2"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.network_id", "dummyNetworkID3"),

					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 1, "192.168.1.50"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 2, "192.168.2.101"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 3, "192.168.3.50"),
					testAccCheckVNAV1InterfaceHasIPAddress(&vna, 3, "192.168.3.60"),
				),
			},
		},
	})
}

func TestMockedAccVNAV1Appliance_createInterfaceDiscontinuity(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1ApplianceInterfaceDiscontinuityPost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterInterfaceDiscontinuityCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterInterfaceDiscontinuityCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterInterfaceDiscontinuityDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceInterfaceDiscontinuity,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.name", "interface_3"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.description", "interface_3_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_info.0.network_id", "dummyNetworkID"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_3_fixed_ips.0.ip_address", "192.168.1.53"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_8_info.0.name", "interface_8"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_8_info.0.description", "interface_8_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_8_info.0.network_id", "dummyNetworkID"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_8_fixed_ips.0.ip_address", "192.168.1.58"),
				),
			},
		},
	})
}

func TestMockedAccVNAV1Appliance_createNoInterface(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1ApplianceNoInterfacePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterNoInterfaceCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterNoInterfaceCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterNoInterfaceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceNoInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
				),
			},
		},
	})
}

func TestMockedAccVNAV1Appliance_createFixedIPsEmpty(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1ApplianceFixedIPsEmptyPost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterFixedIPsEmptyCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterFixedIPsEmptyCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterFixedIPsEmptyDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceFixedIPsEmpty,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", "dummyNetworkID"),
				),
			},
		},
	})
}

func TestMockedAccVNAV1Appliance_createNoFixedIPs(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1ApplianceNoFixedIPsPost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterNoFixedIPsCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterNoFixedIPsCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterNoFixedIPsDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceNoFixedIPs,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", "dummyNetworkID"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
				),
			},
		},
	})
}

func TestMockedAccVNAV1Appliance_basic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1AppliancePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "description", "appliance_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "virtual_network_appliance_plan_id", OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.name", "interface_1"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.description", "interface_1_description"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_info.0.network_id", "dummyNetworkID"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "interface_1_fixed_ips.0.ip_address", "192.168.1.50"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "initial_config.format", "set"),
					resource.TestCheckResourceAttr("ecl_vna_appliance_v1.appliance_1", "initial_config.data", "c2V0IGludGVyZmFjZXMgZ2UtMC8wLzAgZGVzY3JpcHRpb24gc2FtcGxl"),
				),
			},
		},
	})
}

func TestMockedAccVNAV1ApplianceSimple_basic(t *testing.T) {
	var vna appliances.Appliance

	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystoneResponse := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystoneResponse)

	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances", testMockVNAV1ApplianceSimplePost)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterSimpleCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetActiveAfterSimpleCreate)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceProcessingAfterDelete)
	mc.Register(t, "virtual_network_appliance", "/v1.0/virtual_network_appliances/", testMockVNAV1ApplianceGetDeleteComplete)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckVNA(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVNAV1ApplianceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testMockedAccVNAV1ApplianceSimpleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceExists("ecl_vna_appliance_v1.appliance_1", &vna),
				),
			},
		},
	})
}

var testMockedAccVNAV1ApplianceBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "dummyNetworkID"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	initial_config = {
		format = "set"
		data = "c2V0IGludGVyZmFjZXMgZ2UtMC8wLzAgZGVzY3JpcHRpb24gc2FtcGxl"
	}

	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceInterfaceDiscontinuity = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	interface_3_info  {
		name = "interface_3"
		description = "interface_3_description"
		network_id = "dummyNetworkID"
	}

	interface_3_fixed_ips {
		ip_address = "192.168.1.53"
	}

	interface_8_info  {
		name = "interface_8"
		description = "interface_8_description"
		network_id = "dummyNetworkID"
	}

	interface_8_fixed_ips {
		ip_address = "192.168.1.58"
	}
	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceNoInterface = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceFixedIPsEmpty = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "dummyNetworkID"
	}

	interface_1_no_fixed_ips = "true"

	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceNoFixedIPs = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "dummyNetworkID"
	}

	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceUpdateAllowedAddressPairBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "dummyNetworkID"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_1_allowed_address_pairs {
		ip_address = "192.168.1.200"
		mac_address = "aa:bb:cc:dd:ee:f1"
		type = "vrrp"
		vrid = "123"
	}

	interface_1_allowed_address_pairs {
		ip_address = "192.168.1.201"
		mac_address = "aa:bb:cc:dd:ee:f2"
		type = ""
		vrid = "null"
	}

	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceUpdateMetaBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1-update"
	description = "appliance_1_description-update"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	tags = {
		k1 = "v1"
		k2 = "v2"
	}

	interface_1_info  {
		name = "interface_1-update"
		description = "interface_1_description-update"
		network_id = "dummyNetworkID"
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
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceUpdateMetaBasic2 = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1-update"
	description = "appliance_1_description-update"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	interface_1_info  {
		name = "interface_1-update"
		description = "interface_1_description-update"
		network_id = "dummyNetworkID"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceUpdateFixedIPBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
		network_id = "dummyNetworkID"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	interface_2_info  {
		network_id = "dummyNetworkID2"
	}

	interface_3_info  {
		network_id = "dummyNetworkID3"
	}

	interface_3_fixed_ips {
		ip_address = "192.168.3.50"
	}

	interface_3_fixed_ips {
		ip_address = "192.168.3.60"
	}

	interface_4_info  {
		network_id = "dummyNetworkID4"
	}

	interface_4_no_fixed_ips = "true"

	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceDelete = `
request:
    method: DELETE
response:
    code: 204
expectedStatus:
    - Created
    - Updated
    - Updated1
    - Updated2
newStatus: Deleted
`

var testMockVNAV1ApplianceMetaPatch1 = fmt.Sprintf(`
request:
    method: PATCH
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1-update",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {
                    "k1": "v1",
                    "k2": "v2"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
newStatus: Updated1
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceMetaPatch2 = fmt.Sprintf(`
request:
    method: PATCH
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1-update",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated1
newStatus: Updated2
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceFixedIPPatch = fmt.Sprintf(`
request:
    method: PATCH
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "%s",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "%s",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "PROCESSING",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID3",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.3.50",
        					"subnet_id": ""
        				}, {
        					"ip_address": "192.168.3.60",
        					"subnet_id": ""
        				}]
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID4",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID2",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Created
newStatus: Updated
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID)

var testMockVNAV1ApplianceAllowedAddressPairPatch = fmt.Sprintf(`
request:
    method: PATCH
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "%s",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "%s",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "PROCESSING",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [
                            {
                                "ip_address": "192.168.1.200",
                                "mac_address": "aa:bb:cc:dd:ee:f1",
                                "type": "vrrp",
                                "vrid": 123
                            },
                            {
                                "ip_address": "192.168.1.201",
                                "mac_address": "aa:bb:cc:dd:ee:f2",
                                "type": "",
                                "vrid": null
                            }
                        ],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
- Created
newStatus: Updated
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1AppliancePost = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Passw0rd",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
newStatus: Created
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
expectedStatus:
    - Created
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetDeleteComplete = `
request:
    method: GET
response:
    code: 404
expectedStatus:
    - Deleted
counter:
    min: 4
`

var testMockVNAV1ApplianceProcessingAfterDelete = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {
                    "k1": "v1"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterMetaUpdate1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1-update",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {
                    "k1": "v1",
                    "k2": "v2"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated1
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterMetaUpdate1 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {
                            "interfaceK1": "interfaceV1",
                            "interfaceK2": "interfaceV2"
                        },
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1-update",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {
                    "k1": "v1",
                    "k2": "v2"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated1
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterMetaUpdate2 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1-update",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated2
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterMetaUpdate2 = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description-update",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description-update",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1-update",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1-update",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Updated2
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterFixedIPUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "%s",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "%s",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "PROCESSING",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID3",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.3.50",
        					"subnet_id": ""
        				}, {
        					"ip_address": "192.168.3.60",
        					"subnet_id": ""
        				}]
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID4",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID2",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Updated
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterFixedIPUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "%s",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "%s",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "COMPLETE",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID3",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.3.50",
        					"subnet_id": "670c1b56-4b9c-4bde-b7ac-1f2e09391d81"
        				}, {
        					"ip_address": "192.168.3.60",
        					"subnet_id": "670c1b56-4b9c-4bde-b7ac-1f2e09391d81"
        				}]
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID4",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID2",
        				"allowed_address_pairs": [],
        				"fixed_ips": [{
        					"ip_address": "192.168.2.101",
        					"subnet_id": "4be82753-9dc5-4065-a4c0-46abe02bb93a"
        				}]
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Updated
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterAllowedAddressPairUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "%s",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "%s",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "PROCESSING",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [
                            {
                                "ip_address": "192.168.1.200",
                                "mac_address": "aa:bb:cc:dd:ee:f1",
                                "type": "vrrp",
                                "vrid": 123
                            },
                            {
                                "ip_address": "192.168.1.201",
                                "mac_address": "aa:bb:cc:dd:ee:f2",
                                "type": "",
                                "vrid": null
                            }
                        ],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Updated
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterAllowedAddressPairUpdate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
        	"virtual_network_appliance": {
        		"id": "45db3e66-31af-45a6-8ad2-d01521726145",
        		"name": "appliance_1",
        		"description": "appliance_1_description",
        		"tags": {},
        		"appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
        		"availability_zone": "%s",
        		"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
        		"virtual_network_appliance_plan_id": "%s",
        		"os_monitoring_status": "ACTIVE",
        		"os_login_status": "ACTIVE",
        		"vm_status": "ACTIVE",
        		"operation_status": "COMPLETE",
        		"interfaces": {
        			"interface_7": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_8": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_3": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_6": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_1": {
        				"name": "interface_1",
        				"description": "interface_1_description",
        				"tags": {},
        				"updatable": true,
        				"network_id": "dummyNetworkID",
        				"allowed_address_pairs": [
                            {
                                "ip_address": "192.168.1.200",
                                "mac_address": "aa:bb:cc:dd:ee:f1",
                                "type": "vrrp",
                                "vrid": 123
                            },
                            {
                                "ip_address": "192.168.1.201",
                                "mac_address": "aa:bb:cc:dd:ee:f2",
                                "type": "",
                                "vrid": null
                            }
                        ],
        				"fixed_ips": [{
        					"ip_address": "192.168.1.50",
        					"subnet_id": "b9e8b310-774b-4a39-a9ef-fada5dee252c"
        				}]
        			},
        			"interface_4": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_2": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			},
        			"interface_5": {
        				"name": "",
        				"description": "",
        				"tags": {},
        				"updatable": true,
        				"network_id": "",
        				"allowed_address_pairs": [],
        				"fixed_ips": []
        			}
        		}
        	}
        }
expectedStatus:
    - Updated
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockedAccVNAV1ApplianceSimpleBasic = fmt.Sprintf(`
resource "ecl_vna_appliance_v1" "appliance_1" {
	name = "appliance_1"
	description = "appliance_1_description"
	default_gateway = "192.168.1.1"
	availability_zone = "%s"
	virtual_network_appliance_plan_id = "%s"

	interface_1_info  {
		name = "interface_1"
		description = "interface_1_description"
        network_id = "dummyNetworkID"
	}

	interface_1_fixed_ips {
		ip_address = "192.168.1.50"
	}

	lifecycle {
		ignore_changes = [
			"default_gateway",
			"username",
			"password",
		]
	}
}`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceSimplePost = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "7c51af83-43de-4eed-9362-2abf685dcb43",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "YK8kWrwSiG0O",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
newStatus: Created
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterSimpleCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
expectedStatus:
    - Created
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterSimpleCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": "b7581b52-9d1c-4978-ba0e-2ecd7c16ee70"
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_2": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_4": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_5": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_6": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_7": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "",
                        "fixed_ips": [],
                        "name": "",
                        "network_id": "",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceInterfaceDiscontinuityPost = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "interface_3_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.53",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_3",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "interface_8_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.58",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_8",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Passw0rd",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
newStatus: Created
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterInterfaceDiscontinuityCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "interface_3_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.53",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_3",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "interface_8_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.58",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_8",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
expectedStatus:
    - Created
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterInterfaceDiscontinuityCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "interface_3_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.53",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_3",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "interface_8_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.58",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_8",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterInterfaceDiscontinuityDelete = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_3": {
                        "allowed_address_pairs": [],
                        "description": "interface_3_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.53",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_3",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    },
                    "interface_8": {
                        "allowed_address_pairs": [],
                        "description": "interface_8_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.58",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_8",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {
                    "k1": "v1"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceNoInterfacePost = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Passw0rd",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
newStatus: Created
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterNoInterfaceCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
expectedStatus:
    - Created
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterNoInterfaceCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "name": "appliance_1",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterNoInterfaceDelete = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {
                    "k1": "v1"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceFixedIPsEmptyPost = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Passw0rd",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
newStatus: Created
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterFixedIPsEmptyCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
expectedStatus:
    - Created
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterFixedIPsEmptyCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterFixedIPsEmptyDelete = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {
                    "k1": "v1"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceNoFixedIPsPost = fmt.Sprintf(`
request:
    method: POST
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Passw0rd",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
newStatus: Created
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterNoFixedIPsCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "initial",
                "os_monitoring_status": "initial",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "initial"
            }
        }
expectedStatus:
    - Created
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceGetActiveAfterNoFixedIPsCreate = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "COMPLETE",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {},
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Created
counter:
    min: 4
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)

var testMockVNAV1ApplianceProcessingAfterNoFixedIPsDelete = fmt.Sprintf(`
request:
    method: GET
response:
    code: 200
    body: >
        {
            "virtual_network_appliance": {
                "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
                "availability_zone": "%s",
                "default_gateway": "192.168.1.1",
                "description": "appliance_1_description",
                "id": "45db3e66-31af-45a6-8ad2-d01521726145",
                "interfaces": {
                    "interface_1": {
                        "allowed_address_pairs": [],
                        "description": "interface_1_description",
                        "fixed_ips": [
                            {
                                "ip_address": "192.168.1.50",
                                "subnet_id": ""
                            }
                        ],
                        "name": "interface_1",
                        "network_id": "dummyNetworkID",
                        "tags": {},
                        "updatable": true
                    }
                },
                "name": "appliance_1",
                "operation_status": "PROCESSING",
                "os_login_status": "ACTIVE",
                "os_monitoring_status": "ACTIVE",
                "password": "Undxlri8Bo6m",
                "tags": {
                    "k1": "v1"
                },
                "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
                "username": "root",
                "virtual_network_appliance_plan_id": "%s",
                "vm_status": "ACTIVE"
            }
        }
expectedStatus:
    - Deleted
counter:
    max: 3
`,
	OS_DEFAULT_ZONE,
	OS_VIRTUAL_NETWORK_APPLIANCE_PLAN_ID,
)
