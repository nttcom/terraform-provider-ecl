package ecl

import (
	"fmt"
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_interfaces"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_syslog_servers"
	"github.com/nttcom/eclcloud/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/ecl/network/v2/subnets"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancers"
)

// TestAccNetworkV2LoadBalancer_basic tests basic behavior of Load Balancer creation and update requests.
// Step 0: Create Load Balancer with 2 connected interfaces (One with VIP configurations and another one without)
//           and 2 syslog servers.
// Step 1: Update Load Balancer and all sub resources as much as possible without recreating resources.
func TestAccNetworkV2LoadBalancer_basic(t *testing.T) {
	// Equivalent test exists in mocked test, so skip this test in short mode.
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	testAccNetworkV2LoadBalancerBasicSteps(t)
}

func testAccNetworkV2LoadBalancerBasicSteps(t *testing.T) {
	var n1 networks.Network
	var sn1 subnets.Subnet
	var n2 networks.Network
	var sn2 subnets.Subnet
	var lb load_balancers.LoadBalancer
	var lbIF1 load_balancer_interfaces.LoadBalancerInterface
	var lbIF2 load_balancer_interfaces.LoadBalancerInterface
	var lbSyslog1 load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2 load_balancer_syslog_servers.LoadBalancerSyslogServer

	var lbUpdated load_balancers.LoadBalancer
	var lbIF1Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF2Updated load_balancer_interfaces.LoadBalancerInterface
	var lbSyslog1Updated load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2Updated load_balancer_syslog_servers.LoadBalancerSyslogServer

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "id", &lb.ID),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "admin_password"),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "admin_username"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "name", "lb_test1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "availability_zone", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "description", "load_balancer_test1_description"),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "load_balancer_plan_id"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", "192.168.151.1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "tenant_id", OS_TENANT_ID),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "user_password"),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "user_username"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "2"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.description", "lb_test1_interface1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.name", "lb_test1_interface1"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_address", "192.168.151.31"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.0.protocol", "vrrp"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.0.vrid", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbIF2),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbSyslog2),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.id", &lbSyslog2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.acl_logging", "DISABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.appflow_logging", "DISABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.date_format", "YYYYMMDD"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.description", "lb_test1_syslog2_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.ip_address", "192.168.151.22"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_facility", "LOCAL1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_level", "DEBUG"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.name", "lb_test1_syslog2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.priority", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tcp_logging", "NONE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.time_zone", "GMT_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.user_configurable_log_messages", "YES"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tenant_id", OS_TENANT_ID),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerUpdateBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lbUpdated),

					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "id", &lb.ID),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "admin_password"),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "admin_username"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "name", "lb_test1_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "availability_zone", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "description", "load_balancer_test1_description_update"),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "load_balancer_plan_id"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "tenant_id", OS_TENANT_ID),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "user_password"),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "user_username"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "2"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.description", "lb_test1_interface1_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.12"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.name", "lb_test1_interface1_update"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_address", "192.168.151.32"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.0.protocol", "vrrp"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.0.vrid", "30"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.description", "lb_test1_interface2_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.ip_address", "192.168.152.12"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.name", "lb_test1_interface2_update"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "DISABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "DISABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "YYYYMMDD"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "DEBUG"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "NONE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "GMT_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "YES"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbSyslog2Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.id", &lbSyslog2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.description", "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.ip_address", "192.168.151.22"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.name", "lb_test1_syslog2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tenant_id", OS_TENANT_ID),
				),
			},
		},
	})
}

// TestAccNetworkV2LoadBalancer_updatePlanDecrease tests changing both Load Balancer Plan and Interfaces at the same time.
// Step 0: Create Load Balancer with 8IF Load Balancer Plan and 5 connected interfaces.
// Step 1: Update Load Balancer to 4IF Load Balancer Plan and disconnect 1 interface.
func TestAccNetworkV2LoadBalancer_updatePlanDecrease(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var n1 networks.Network
	var sn1 subnets.Subnet
	var n2 networks.Network
	var sn2 subnets.Subnet
	var n3 networks.Network
	var sn3 subnets.Subnet
	var n4 networks.Network
	var sn4 subnets.Subnet
	var n5 networks.Network
	var sn5 subnets.Subnet
	var lb load_balancers.LoadBalancer
	var lbIF1 load_balancer_interfaces.LoadBalancerInterface
	var lbIF2 load_balancer_interfaces.LoadBalancerInterface
	var lbIF3 load_balancer_interfaces.LoadBalancerInterface
	var lbIF4 load_balancer_interfaces.LoadBalancerInterface
	var lbIF5 load_balancer_interfaces.LoadBalancerInterface

	var lbUpdated load_balancers.LoadBalancer
	var lbIF1Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF2Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF3Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF4Updated load_balancer_interfaces.LoadBalancerInterface

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerUpdatePlan5IF,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_3_1", &sn3),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_4", &n4),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_4_1", &sn4),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_5", &n5),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_5_1", &sn5),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "load_balancer_plan_id"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "5"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbIF2),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "2", &lbIF3),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "3", &lbIF4),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.ip_address", "192.168.154.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "4", &lbIF5),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.4.id", &lbIF5.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.4.ip_address", "192.168.155.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.4.slot_number", "5"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.4.status", "ACTIVE"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerUpdatePlan4IF,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_3_1", &sn3),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_4", &n4),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_4_1", &sn4),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lbUpdated),

					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "load_balancer_plan_id"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "2", &lbIF3Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "3", &lbIF4Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.ip_address", "192.168.154.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.status", "ACTIVE"),
				),
			},
		},
	})
}

// TestAccNetworkV2LoadBalancer_updatePlanIncrease tests changing both Load Balancer Plan and Interfaces at the same time.
// Step 0: Create Load Balancer with 4IF Load Balancer Plan and 4 connected interfaces.
// Step 1: Update Load Balancer to 8IF Load Balancer Plan and connect new 1 interface.
func TestAccNetworkV2LoadBalancer_updatePlanIncrease(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var n1 networks.Network
	var sn1 subnets.Subnet
	var n2 networks.Network
	var sn2 subnets.Subnet
	var n3 networks.Network
	var sn3 subnets.Subnet
	var n4 networks.Network
	var sn4 subnets.Subnet
	var n5 networks.Network
	var sn5 subnets.Subnet
	var lb load_balancers.LoadBalancer
	var lbIF1 load_balancer_interfaces.LoadBalancerInterface
	var lbIF2 load_balancer_interfaces.LoadBalancerInterface
	var lbIF3 load_balancer_interfaces.LoadBalancerInterface
	var lbIF4 load_balancer_interfaces.LoadBalancerInterface
	var lbIF5 load_balancer_interfaces.LoadBalancerInterface

	var lbUpdated load_balancers.LoadBalancer
	var lbIF1Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF2Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF3Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF4Updated load_balancer_interfaces.LoadBalancerInterface

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerUpdatePlan4IF,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_3_1", &sn3),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_4", &n4),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_4_1", &sn4),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "load_balancer_plan_id"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbIF2),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "2", &lbIF3),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "3", &lbIF4),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.ip_address", "192.168.154.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.status", "ACTIVE"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerUpdatePlan5IF,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_3_1", &sn3),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_4", &n4),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_4_1", &sn4),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_5", &n5),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_5_1", &sn5),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lbUpdated),

					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "load_balancer_plan_id"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "5"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "2", &lbIF3Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.2.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "3", &lbIF4Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.ip_address", "192.168.154.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.3.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "4", &lbIF5),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.4.id", &lbIF5.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.4.ip_address", "192.168.155.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.4.slot_number", "5"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.4.status", "ACTIVE"),
				),
			},
		},
	})
}

// TestAccNetworkV2LoadBalancer_forceNew tests that ForceNew attribute works functionally.
// Step 0: Create Load Balancer.
// Step 1: Update Load Balancer availability_zone to force destroying/recreating.
func TestAccNetworkV2LoadBalancer_forceNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var lb load_balancers.LoadBalancer
	var lbUpdated load_balancers.LoadBalancer

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerMinimumInit,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "id", &lb.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "name", "lb_test1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "availability_zone", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "description", "load_balancer_test1_description"),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "load_balancer_plan_id"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "tenant_id", OS_TENANT_ID),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "0"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerMinimumUpdateAZ,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lbUpdated),
					testAccCheckNetworkV2LoadBalancerDoesNotExist(&lb),

					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "id", &lbUpdated.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "name", "lb_test1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "description", "load_balancer_test1_description"),
					resource.TestCheckResourceAttrSet("ecl_network_load_balancer_v2.lb_test1", "load_balancer_plan_id"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "tenant_id", OS_TENANT_ID),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "0"),
				),
			},
		},
	})
}

// TestAccNetworkV2LoadBalancer_modifyInterfaces tests interface connect/replace/disconnect operations.
// Step 0: Create Load Balancer with 1 interface -> network: [1]
//			This interface has VIP configuration.
// Step 1: connect 1 new network to Load Balancer interface -> network: [1, 2]
//			Also remove VIP configuration from the existing interface.
// Step 2: replace 1 network -> network: [3, 2]
// Step 3: disconnect 1 network -> network: [2]
func TestAccNetworkV2LoadBalancer_modifyInterfaces(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var n1 networks.Network
	var sn1 subnets.Subnet
	var n2 networks.Network
	var sn2 subnets.Subnet
	var n3 networks.Network
	var sn3 subnets.Subnet
	var lb load_balancers.LoadBalancer
	var lbIF1 load_balancer_interfaces.LoadBalancerInterface
	var lbIF2 load_balancer_interfaces.LoadBalancerInterface

	var lbUpdated load_balancers.LoadBalancer
	var lbIF1Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF2Updated load_balancer_interfaces.LoadBalancerInterface

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerInterfaces1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", "192.168.151.1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "1"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.description", "lb_test1_interface1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.name", "lb_test1_interface1"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_address", "192.168.151.31"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.0.protocol", "vrrp"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.0.vrid", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces1And2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lbUpdated),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "2"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.description", "lb_test1_interface1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.name", "lb_test1_interface1"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_address", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbIF2),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.status", "ACTIVE"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces3And2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_3_1", &sn3),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lbUpdated),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "2"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.description", "lb_test1_interface3_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.name", "lb_test1_interface3"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.network_id", &n3.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_address", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.1.status", "ACTIVE"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lbUpdated),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "1"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.network_id", &n2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_address", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "0"),
				),
			},
		},
	})
}

// TestAccNetworkV2LoadBalancer_modifySyslogServers tests syslog server create/update/recreate/delete operations.
// Step 0: Create Load Balancer with 1 interface and 1 syslog server -> syslog server: [1]
// Step 1: Add 1 syslog server -> syslog server: [1, 2]
// Step 2: Update syslog server's fields -> syslog server: [1, 2] (2 is simply updated, not recreated)
// Step 3: Update syslog server's name -> syslog server: [1, 2] (2 is recreated)
// Step 4: Update syslog server's port_number -> syslog server: [1, 2] (2 is recreated)
// Step 5: Update syslog server's ip_address -> syslog server: [1, 2] (2 is recreated)
// Step 6: Delete one syslog server -> syslog server: [2]
func TestAccNetworkV2LoadBalancer_modifySyslogServers(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var n1 networks.Network
	var sn1 subnets.Subnet
	var lb load_balancers.LoadBalancer
	var lbSyslog1 load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2 load_balancer_syslog_servers.LoadBalancerSyslogServer

	var lbSyslog2Updated load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2UpdatedForceNew1 load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2UpdatedForceNew2 load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2UpdatedForceNew3 load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2UpdatedForceNew4 load_balancer_syslog_servers.LoadBalancerSyslogServer

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				// Step 0
				Config: testAccNetworkV2LoadBalancerModifySyslogServerInit,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "1"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),
				),
			},
			{
				// Step 1
				Config: testAccNetworkV2LoadBalancerModifySyslogServerAdd,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbSyslog2),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.id", &lbSyslog2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.acl_logging", "DISABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.appflow_logging", "DISABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.date_format", "YYYYMMDD"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.description", "lb_test1_syslog2_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.ip_address", "192.168.151.22"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_facility", "LOCAL1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_level", "DEBUG"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.name", "lb_test1_syslog2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.priority", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tcp_logging", "NONE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.time_zone", "GMT_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.user_configurable_log_messages", "YES"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tenant_id", OS_TENANT_ID),
				),
			},
			{
				// Step 2
				Config: testAccNetworkV2LoadBalancerModifySyslogServerUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbSyslog2Updated),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.id", &lbSyslog2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.description", "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.ip_address", "192.168.151.22"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.name", "lb_test1_syslog2"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tenant_id", OS_TENANT_ID),
				),
			},
			{
				// Step 3
				Config: testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(&lbSyslog2Updated),
					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbSyslog2UpdatedForceNew1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.id", &lbSyslog2UpdatedForceNew1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.description", "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.ip_address", "192.168.151.22"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.name", "lb_test1_syslog2_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tenant_id", OS_TENANT_ID),
				),
			},
			{
				// Step 4
				Config: testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(&lbSyslog2UpdatedForceNew1),
					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbSyslog2UpdatedForceNew2),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.id", &lbSyslog2UpdatedForceNew2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.description", "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.ip_address", "192.168.151.22"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.name", "lb_test1_syslog2_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.port_number", "1514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tenant_id", OS_TENANT_ID),
				),
			},
			{
				// Step 5
				Config: testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(&lbSyslog2UpdatedForceNew2),
					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "1", &lbSyslog2UpdatedForceNew3),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.id", &lbSyslog2UpdatedForceNew3.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.description", "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.ip_address", "192.168.151.23"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.name", "lb_test1_syslog2_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.port_number", "1514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.1.tenant_id", OS_TENANT_ID),
				),
			},
			{
				// Step 6
				Config: testAccNetworkV2LoadBalancerModifySyslogServerDelete,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "1"),

					testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(&lbSyslog1),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog2UpdatedForceNew4),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog2UpdatedForceNew4.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.23"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog2_update"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "1514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),
				),
			},
		},
	})
}

// TestAccNetworkV2LoadBalancer_modifyInterfacesWithIPs tests simultaneous updating interface and other IP addresses.
// Step 0: Create Load Balancer without sub resources.
//	-> network: [], default_gateway: null, syslog server IP: []
// Step 1: Connect 1 new network to interface, set default_gateway and create 1 syslog server
//	-> network: [1], default_gateway: in network 1, syslog server IP: [in network 1]
// Step 2: Replace network and update other IPs to new network
//	-> network: [2], default_gateway: in network 2, syslog server IP: [in network 2]
// Step 3: Disconnect network and release other IPs
//	-> network: [], default_gateway: null, syslog server IP: []
func TestAccNetworkV2LoadBalancer_modifyInterfacesWithIPs(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var lb, lbUpdated load_balancers.LoadBalancer
	var lbIF1 load_balancer_interfaces.LoadBalancerInterface
	var lbSyslog1 load_balancer_syslog_servers.LoadBalancerSyslogServer
	var n1, n2 networks.Network
	var sn1, sn2 subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerMinimumInit,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "0"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces1WithSyslogServer,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", "192.168.151.1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "1"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.description", "lb_test1_interface1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.name", "lb_test1_interface1"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_address", "192.168.151.31"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.0.protocol", "vrrp"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.0.vrid", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "1"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.151.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces2WithSyslogServer,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lbUpdated),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "1"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbIF1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.network_id", &n2.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_address", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.0.status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.#", "1"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists("ecl_network_load_balancer_v2.lb_test1", "0", &lbSyslog1),
					resource.TestCheckResourceAttrPtr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.id", &lbSyslog1.ID),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.acl_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.appflow_logging", "ENABLED"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.date_format", "MMDDYYYY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.description", "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.ip_address", "192.168.152.21"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_facility", "LOCAL0"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.log_level", "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.name", "lb_test1_syslog1"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.port_number", "514"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.priority", "20"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tcp_logging", "ALL"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.time_zone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.transport_type", "UDP"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.user_configurable_log_messages", "NO"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.status", "ACTIVE"),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "syslog_servers.0.tenant_id", OS_TENANT_ID),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerMinimumInit,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerExists("ecl_network_load_balancer_v2.lb_test1", &lb),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "default_gateway", ""),
					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "status", "ACTIVE"),

					resource.TestCheckResourceAttr("ecl_network_load_balancer_v2.lb_test1", "interfaces.#", "0"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2LoadBalancerDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_load_balancer_v2" {
			continue
		}

		_, err := load_balancers.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("load balancer still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2LoadBalancerExists(n string, lb *load_balancers.LoadBalancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL network client: %s", err)
		}

		found, err := load_balancers.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("load balancer not found")
		}

		*lb = *found

		return nil
	}
}

func testAccCheckNetworkV2LoadBalancerDoesNotExist(loadBalancer *load_balancers.LoadBalancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL network client: %s", err)
		}

		_, err = load_balancer_syslog_servers.Get(networkClient, loadBalancer.ID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return nil
			}
			return err
		}

		return fmt.Errorf("instance still exists")
	}
}

func testAccCheckNetworkV2LoadBalancerInterfaceExists(n string, i string, lbIF *load_balancer_interfaces.LoadBalancerInterface) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		key := fmt.Sprintf("interfaces.%s.id", i)
		value := rs.Primary.Attributes[key]
		if value == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL network client: %s", err)
		}

		found, err := load_balancer_interfaces.Get(networkClient, value).Extract()
		if err != nil {
			return err
		}

		if found.ID != value {
			return fmt.Errorf("load balancer interface not found")
		}

		*lbIF = *found

		return nil
	}
}

func testAccCheckNetworkV2LoadBalancerSyslogServerExists(n string, i string, lbSyslog *load_balancer_syslog_servers.LoadBalancerSyslogServer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		key := fmt.Sprintf("syslog_servers.%s.id", i)
		value := rs.Primary.Attributes[key]
		if value == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL network client: %s", err)
		}

		found, err := load_balancer_syslog_servers.Get(networkClient, value).Extract()
		if err != nil {
			return err
		}

		if found.ID != value {
			return fmt.Errorf("load balancer syslog server not found")
		}

		*lbSyslog = *found

		return nil
	}
}

func testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(syslogServer *load_balancer_syslog_servers.LoadBalancerSyslogServer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL network client: %s", err)
		}

		_, err = load_balancer_syslog_servers.Get(networkClient, syslogServer.ID).Extract()
		if err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				return nil
			}
			return err
		}

		return fmt.Errorf("instance still exists")
	}
}

var testAccNetworkV2LoadBalancerBasic = fmt.Sprintf(`
%s

%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerInterface2,
	testAccNetworkV2LoadBalancerSyslogServer1,
	testAccNetworkV2LoadBalancerSyslogServer2,
)

var testAccNetworkV2LoadBalancerUpdateBasic = fmt.Sprintf(`
%s

%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1_update"
  availability_zone = "%s"
  description = "load_balancer_test1_description_update"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerPlan8IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway2,
	testAccNetworkV2LoadBalancerInterface1UpdateBasic,
	testAccNetworkV2LoadBalancerInterface2UpdateBasic,
	testAccNetworkV2LoadBalancerSyslogServer1UpdateBasic,
	testAccNetworkV2LoadBalancerSyslogServer2UpdateBasic,
)

var testAccNetworkV2LoadBalancerUpdatePlan5IF = fmt.Sprintf(`
%s

%s

%s

%s

%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair3,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair4,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair5,
	testAccNetworkV2LoadBalancerPlan8IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerInterface2,
	testAccNetworkV2LoadBalancerInterface3,
	testAccNetworkV2LoadBalancerInterface4,
	testAccNetworkV2LoadBalancerInterface5,
)

var testAccNetworkV2LoadBalancerUpdatePlan4IF = fmt.Sprintf(`
%s

%s

%s

%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair3,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair4,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerInterface2,
	testAccNetworkV2LoadBalancerInterface3,
	testAccNetworkV2LoadBalancerInterface4,
)

var testAccNetworkV2LoadBalancerMinimumInit = fmt.Sprintf(`

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
}
`,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	"",
)

var testAccNetworkV2LoadBalancerMinimumUpdateAZ = fmt.Sprintf(`

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
}
`,
	testAccNetworkV2LoadBalancerPlan4IF,
	"zone1_groupb",
	"",
)

var testAccNetworkV2LoadBalancerInterfaces1 = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
)

var testAccNetworkV2LoadBalancerInterfaces1WithSyslogServer = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerSyslogServer1,
)

var testAccNetworkV2LoadBalancerInterfaces1And2 = fmt.Sprintf(`
%s

%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway2,
	testAccNetworkV2LoadBalancerInterface1WithoutVIP,
	testAccNetworkV2LoadBalancerInterface2,
)

var testAccNetworkV2LoadBalancerInterfaces3And2 = fmt.Sprintf(`
%s

%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair3,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway2,
	testAccNetworkV2LoadBalancerInterface3,
	testAccNetworkV2LoadBalancerInterface2,
)

var testAccNetworkV2LoadBalancerInterfaces2 = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway2,
	testAccNetworkV2LoadBalancerInterface2,
)

var testAccNetworkV2LoadBalancerInterfaces2WithSyslogServer = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway2,
	testAccNetworkV2LoadBalancerInterface2,
	testAccNetworkV2LoadBalancerSyslogServer1InInterface2,
)

var testAccNetworkV2LoadBalancerModifySyslogServerInit = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerSyslogServer1,
)

var testAccNetworkV2LoadBalancerModifySyslogServerAdd = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerSyslogServer1,
	testAccNetworkV2LoadBalancerSyslogServer2,
)

var testAccNetworkV2LoadBalancerModifySyslogServerUpdate = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerSyslogServer1,
	testAccNetworkV2LoadBalancerSyslogServer2UpdateBasic,
)

var testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew1 = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerSyslogServer1,
	testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew1,
)

var testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew2 = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerSyslogServer1,
	testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew2,
)

var testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew3 = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerSyslogServer1,
	testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew3,
)

var testAccNetworkV2LoadBalancerModifySyslogServerDelete = fmt.Sprintf(`
%s

%s

resource "ecl_network_load_balancer_v2" "lb_test1" {
  name = "lb_test1"
  availability_zone = "%s"
  description = "load_balancer_test1_description"
  load_balancer_plan_id = data.ecl_network_load_balancer_plan_v2.lb_plan1.id
  %s
  %s
  %s
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway1,
	testAccNetworkV2LoadBalancerInterface1,
	testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew3,
)

const testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair1 = `
resource "ecl_network_network_v2" "network_1" {
  name = "network_1"
}

resource "ecl_network_subnet_v2" "subnet_1_1" {
  name       = "subnet_1_1"
  cidr       = "192.168.151.0/24"
  gateway_ip = "192.168.151.1"
  network_id = "${ecl_network_network_v2.network_1.id}"

  allocation_pools {
    start = "192.168.151.100"
    end   = "192.168.151.200"
  }
}
`

const testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2 = `
resource "ecl_network_network_v2" "network_2" {
  name = "network_2"
}

resource "ecl_network_subnet_v2" "subnet_2_1" {
  name       = "subnet_2_1"
  cidr       = "192.168.152.0/24"
  gateway_ip = "192.168.152.1"
  network_id = "${ecl_network_network_v2.network_2.id}"

  allocation_pools {
    start = "192.168.152.100"
    end   = "192.168.152.200"
  }
}
`

const testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair3 = `
resource "ecl_network_network_v2" "network_3" {
  name = "network_3"
}

resource "ecl_network_subnet_v2" "subnet_3_1" {
  name       = "subnet_3_1"
  cidr       = "192.168.153.0/24"
  gateway_ip = "192.168.153.1"
  network_id = "${ecl_network_network_v2.network_3.id}"

  allocation_pools {
    start = "192.168.153.100"
    end   = "192.168.153.200"
  }
}
`

const testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair4 = `
resource "ecl_network_network_v2" "network_4" {
  name = "network_4"
}

resource "ecl_network_subnet_v2" "subnet_4_1" {
  name       = "subnet_4_1"
  cidr       = "192.168.154.0/24"
  gateway_ip = "192.168.154.1"
  network_id = "${ecl_network_network_v2.network_4.id}"

  allocation_pools {
    start = "192.168.154.100"
    end   = "192.168.154.200"
  }
}
`

const testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair5 = `
resource "ecl_network_network_v2" "network_5" {
  name = "network_5"
}

resource "ecl_network_subnet_v2" "subnet_5_1" {
  name       = "subnet_5_1"
  cidr       = "192.168.155.0/24"
  gateway_ip = "192.168.155.1"
  network_id = "${ecl_network_network_v2.network_5.id}"

  allocation_pools {
    start = "192.168.155.100"
    end   = "192.168.155.200"
  }
}
`

const testAccNetworkV2LoadBalancerPlan4IF = `
data "ecl_network_load_balancer_plan_v2" "lb_plan1" {
  enabled = true
  model {
    size = "200"
  }
}
`

const testAccNetworkV2LoadBalancerPlan8IF = `
data "ecl_network_load_balancer_plan_v2" "lb_plan1" {
  enabled = true
  model {
    size = "1000"
  }
}
`

const testAccNetworkV2LoadBalancerDefaultGateway1 = `default_gateway = "192.168.151.1"`
const testAccNetworkV2LoadBalancerDefaultGateway2 = `default_gateway = "192.168.152.1"`

const testAccNetworkV2LoadBalancerInterface1 = `
interfaces {
    description = "lb_test1_interface1_description"
    ip_address = "192.168.151.11"
    name = "lb_test1_interface1"
    network_id = "${ecl_network_network_v2.network_1.id}"
    virtual_ip_address = "192.168.151.31"
    virtual_ip_properties {
        protocol = "vrrp"
        vrid = 20
    }
}
`

const testAccNetworkV2LoadBalancerInterface1UpdateBasic = `
interfaces {
    description = "lb_test1_interface1_description_update"
    ip_address = "192.168.151.12"
    name = "lb_test1_interface1_update"
    network_id = "${ecl_network_network_v2.network_1.id}"
    virtual_ip_address = "192.168.151.32"
    virtual_ip_properties {
        protocol = "vrrp"
        vrid = 30
    }
}
`

const testAccNetworkV2LoadBalancerInterface1WithoutVIP = `
interfaces {
    description = "lb_test1_interface1_description"
    ip_address = "192.168.151.11"
    name = "lb_test1_interface1"
    network_id = "${ecl_network_network_v2.network_1.id}"
}
`

const testAccNetworkV2LoadBalancerInterface2 = `
interfaces {
    description = "lb_test1_interface2_description"
    ip_address = "192.168.152.11"
    name = "lb_test1_interface2"
    network_id = "${ecl_network_network_v2.network_2.id}"
}
`

const testAccNetworkV2LoadBalancerInterface2UpdateBasic = `
interfaces {
    description = "lb_test1_interface2_description_update"
    ip_address = "192.168.152.12"
    name = "lb_test1_interface2_update"
    network_id = "${ecl_network_network_v2.network_2.id}"
}
`

const testAccNetworkV2LoadBalancerInterface3 = `
interfaces {
    description = "lb_test1_interface3_description"
    ip_address = "192.168.153.11"
    name = "lb_test1_interface3"
    network_id = "${ecl_network_network_v2.network_3.id}"
}
`

const testAccNetworkV2LoadBalancerInterface4 = `
interfaces {
    description = "lb_test1_interface4_description"
    ip_address = "192.168.154.11"
    name = "lb_test1_interface4"
    network_id = "${ecl_network_network_v2.network_4.id}"
}
`

const testAccNetworkV2LoadBalancerInterface5 = `
interfaces {
    description = "lb_test1_interface5_description"
    ip_address = "192.168.155.11"
    name = "lb_test1_interface5"
    network_id = "${ecl_network_network_v2.network_5.id}"
}
`

const testAccNetworkV2LoadBalancerSyslogServer1 = `
syslog_servers {
    acl_logging = "ENABLED"
    appflow_logging = "ENABLED"
    date_format = "MMDDYYYY"
    description = "lb_test1_syslog1_description"
    ip_address = "192.168.151.21"
    log_facility = "LOCAL0"
    log_level = "ALERT|CRITICAL|EMERGENCY"
    name = "lb_test1_syslog1"
    port_number = "514"
    priority = "20"
    tcp_logging = "ALL"
    time_zone = "LOCAL_TIME"
    transport_type = "UDP"
    user_configurable_log_messages = "NO"
}
`

const testAccNetworkV2LoadBalancerSyslogServer1InInterface2 = `
syslog_servers {
    acl_logging = "ENABLED"
    appflow_logging = "ENABLED"
    date_format = "MMDDYYYY"
    description = "lb_test1_syslog1_description"
    ip_address = "192.168.152.21"
    log_facility = "LOCAL0"
    log_level = "ALERT|CRITICAL|EMERGENCY"
    name = "lb_test1_syslog1"
    port_number = "514"
    priority = "20"
    tcp_logging = "ALL"
    time_zone = "LOCAL_TIME"
    transport_type = "UDP"
    user_configurable_log_messages = "NO"
}
`

const testAccNetworkV2LoadBalancerSyslogServer1UpdateBasic = `
syslog_servers {
    acl_logging = "DISABLED"
    appflow_logging = "DISABLED"
    date_format = "YYYYMMDD"
    description = "lb_test1_syslog1_description_update"
    ip_address = "192.168.151.21"
    log_facility = "LOCAL1"
    log_level = "DEBUG"
    name = "lb_test1_syslog1"
    port_number = "514"
    priority = "0"
    tcp_logging = "NONE"
    time_zone = "GMT_TIME"
    transport_type = "UDP"
    user_configurable_log_messages = "YES"
}
`

const testAccNetworkV2LoadBalancerSyslogServer2 = `
syslog_servers {
    acl_logging = "DISABLED"
    appflow_logging = "DISABLED"
    date_format = "YYYYMMDD"
    description = "lb_test1_syslog2_description"
    ip_address = "192.168.151.22"
    log_facility = "LOCAL1"
    log_level = "DEBUG"
    name = "lb_test1_syslog2"
    port_number = "514"
    priority = "0"
    tcp_logging = "NONE"
    time_zone = "GMT_TIME"
    transport_type = "UDP"
    user_configurable_log_messages = "YES"
}
`

const testAccNetworkV2LoadBalancerSyslogServer2UpdateBasic = `
syslog_servers {
    acl_logging = "ENABLED"
    appflow_logging = "ENABLED"
    date_format = "MMDDYYYY"
    description = "lb_test1_syslog2_description_update"
    ip_address = "192.168.151.22"
    log_facility = "LOCAL0"
    log_level = "ALERT|CRITICAL|EMERGENCY"
    name = "lb_test1_syslog2"
    port_number = "514"
    priority = "20"
    tcp_logging = "ALL"
    time_zone = "LOCAL_TIME"
    transport_type = "UDP"
    user_configurable_log_messages = "NO"
}
`

const testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew1 = `
syslog_servers {
    acl_logging = "ENABLED"
    appflow_logging = "ENABLED"
    date_format = "MMDDYYYY"
    description = "lb_test1_syslog2_description_update"
    ip_address = "192.168.151.22"
    log_facility = "LOCAL0"
    log_level = "ALERT|CRITICAL|EMERGENCY"
    name = "lb_test1_syslog2_update"
    port_number = "514"
    priority = "20"
    tcp_logging = "ALL"
    time_zone = "LOCAL_TIME"
    transport_type = "UDP"
    user_configurable_log_messages = "NO"
}
`

const testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew2 = `
syslog_servers {
    acl_logging = "ENABLED"
    appflow_logging = "ENABLED"
    date_format = "MMDDYYYY"
    description = "lb_test1_syslog2_description_update"
    ip_address = "192.168.151.22"
    log_facility = "LOCAL0"
    log_level = "ALERT|CRITICAL|EMERGENCY"
    name = "lb_test1_syslog2_update"
    port_number = "1514"
    priority = "20"
    tcp_logging = "ALL"
    time_zone = "LOCAL_TIME"
    transport_type = "UDP"
    user_configurable_log_messages = "NO"
}
`

const testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew3 = `
syslog_servers {
    acl_logging = "ENABLED"
    appflow_logging = "ENABLED"
    date_format = "MMDDYYYY"
    description = "lb_test1_syslog2_description_update"
    ip_address = "192.168.151.23"
    log_facility = "LOCAL0"
    log_level = "ALERT|CRITICAL|EMERGENCY"
    name = "lb_test1_syslog2_update"
    port_number = "1514"
    priority = "20"
    tcp_logging = "ALL"
    time_zone = "LOCAL_TIME"
    transport_type = "UDP"
    user_configurable_log_messages = "NO"
}
`
