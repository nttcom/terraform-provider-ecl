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
	r := "ecl_network_load_balancer_v2.lb_test1"
	var n1 networks.Network
	var sn1 subnets.Subnet
	var n2 networks.Network
	var sn2 subnets.Subnet
	var lb load_balancers.LoadBalancer
	var lbIF1 load_balancer_interfaces.LoadBalancerInterface
	var lbIF2 load_balancer_interfaces.LoadBalancerInterface
	var lbIF3 load_balancer_interfaces.LoadBalancerInterface
	var lbIF4 load_balancer_interfaces.LoadBalancerInterface
	var lbIF5 load_balancer_interfaces.LoadBalancerInterface
	var lbIF6 load_balancer_interfaces.LoadBalancerInterface
	var lbIF7 load_balancer_interfaces.LoadBalancerInterface
	var lbIF8 load_balancer_interfaces.LoadBalancerInterface
	var lbSyslog1 load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2 load_balancer_syslog_servers.LoadBalancerSyslogServer
	syslog1Key := "4120068917"
	syslog2Key := "1121660642"

	var lbUpdated load_balancers.LoadBalancer
	var lbIF1Updated load_balancer_interfaces.LoadBalancerInterface
	var lbIF2Updated load_balancer_interfaces.LoadBalancerInterface
	var lbSyslog1Updated load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2Updated load_balancer_syslog_servers.LoadBalancerSyslogServer
	syslog1UpdateKey := "1046271650"
	syslog2UpdateKey := "4036206399"

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
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttrPtr(r, "id", &lb.ID),
					resource.TestCheckResourceAttrSet(r, "admin_password"),
					resource.TestCheckResourceAttrSet(r, "admin_username"),
					resource.TestCheckResourceAttr(r, "name", "lb_test1"),
					resource.TestCheckResourceAttr(r, "availability_zone", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr(r, "description", "load_balancer_test1_description"),
					resource.TestCheckResourceAttrSet(r, "load_balancer_plan_id"),
					resource.TestCheckResourceAttr(r, "default_gateway", "192.168.151.1"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),
					resource.TestCheckResourceAttr(r, "tenant_id", OS_TENANT_ID),
					resource.TestCheckResourceAttrSet(r, "user_password"),
					resource.TestCheckResourceAttrSet(r, "user_username"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.description", "lb_test1_interface1_description"),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.name", "lb_test1_interface1"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_address", "192.168.151.31"),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.0.protocol", "vrrp"),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.0.vrid", "20"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.1.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "2", &lbIF3),
					resource.TestCheckResourceAttrPtr(r, "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr(r, "interfaces.2.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.2.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.2.name", "Interface 1/3"),
					resource.TestCheckResourceAttr(r, "interfaces.2.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.2.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.2.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr(r, "interfaces.2.status", "DOWN"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "3", &lbIF4),
					resource.TestCheckResourceAttrPtr(r, "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr(r, "interfaces.3.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.3.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.3.name", "Interface 1/4"),
					resource.TestCheckResourceAttr(r, "interfaces.3.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.3.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.3.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr(r, "interfaces.3.status", "DOWN"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1Key, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1Key), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1Key), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1Key), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1Key), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1Key), OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog2Key, &lbSyslog2),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog2Key), &lbSyslog2.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog2Key), "DISABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog2Key), "DISABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog2Key), "YYYYMMDD"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog2Key), "lb_test1_syslog2_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog2Key), "192.168.151.22"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog2Key), "LOCAL1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog2Key), "DEBUG"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog2Key), "lb_test1_syslog2"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog2Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog2Key), "0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog2Key), "NONE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog2Key), "GMT_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog2Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog2Key), "YES"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog2Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog2Key), OS_TENANT_ID),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerUpdateBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists(r, &lbUpdated),

					resource.TestCheckResourceAttrPtr(r, "id", &lb.ID),
					resource.TestCheckResourceAttrSet(r, "admin_password"),
					resource.TestCheckResourceAttrSet(r, "admin_username"),
					resource.TestCheckResourceAttr(r, "name", "lb_test1_update"),
					resource.TestCheckResourceAttr(r, "availability_zone", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr(r, "description", "load_balancer_test1_description_update"),
					resource.TestCheckResourceAttrSet(r, "load_balancer_plan_id"),
					resource.TestCheckResourceAttr(r, "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),
					resource.TestCheckResourceAttr(r, "tenant_id", OS_TENANT_ID),
					resource.TestCheckResourceAttrSet(r, "user_password"),
					resource.TestCheckResourceAttrSet(r, "user_username"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "8"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.description", "lb_test1_interface1_description_update"),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.12"),
					resource.TestCheckResourceAttr(r, "interfaces.0.name", "lb_test1_interface1_update"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_address", "192.168.151.32"),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.0.protocol", "vrrp"),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.0.vrid", "30"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.description", "lb_test1_interface2_description_update"),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.12"),
					resource.TestCheckResourceAttr(r, "interfaces.1.name", "lb_test1_interface2_update"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "2", &lbIF3),
					resource.TestCheckResourceAttrPtr(r, "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr(r, "interfaces.2.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.2.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.2.name", "Interface 1/3"),
					resource.TestCheckResourceAttr(r, "interfaces.2.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.2.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.2.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr(r, "interfaces.2.status", "DOWN"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "3", &lbIF4),
					resource.TestCheckResourceAttrPtr(r, "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr(r, "interfaces.3.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.3.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.3.name", "Interface 1/4"),
					resource.TestCheckResourceAttr(r, "interfaces.3.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.3.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.3.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr(r, "interfaces.3.status", "DOWN"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "4", &lbIF5),
					resource.TestCheckResourceAttrPtr(r, "interfaces.4.id", &lbIF5.ID),
					resource.TestCheckResourceAttr(r, "interfaces.4.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.4.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.4.name", "Interface 1/5"),
					resource.TestCheckResourceAttr(r, "interfaces.4.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.4.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.4.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.4.slot_number", "5"),
					resource.TestCheckResourceAttr(r, "interfaces.4.status", "DOWN"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "5", &lbIF6),
					resource.TestCheckResourceAttrPtr(r, "interfaces.5.id", &lbIF6.ID),
					resource.TestCheckResourceAttr(r, "interfaces.5.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.5.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.5.name", "Interface 1/6"),
					resource.TestCheckResourceAttr(r, "interfaces.5.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.5.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.5.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.5.slot_number", "6"),
					resource.TestCheckResourceAttr(r, "interfaces.5.status", "DOWN"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "6", &lbIF7),
					resource.TestCheckResourceAttrPtr(r, "interfaces.6.id", &lbIF7.ID),
					resource.TestCheckResourceAttr(r, "interfaces.6.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.6.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.6.name", "Interface 1/7"),
					resource.TestCheckResourceAttr(r, "interfaces.6.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.6.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.6.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.6.slot_number", "7"),
					resource.TestCheckResourceAttr(r, "interfaces.6.status", "DOWN"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "7", &lbIF8),
					resource.TestCheckResourceAttrPtr(r, "interfaces.7.id", &lbIF8.ID),
					resource.TestCheckResourceAttr(r, "interfaces.7.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.7.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.7.name", "Interface 1/8"),
					resource.TestCheckResourceAttr(r, "interfaces.7.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.7.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.7.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.7.slot_number", "8"),
					resource.TestCheckResourceAttr(r, "interfaces.7.status", "DOWN"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1UpdateKey, &lbSyslog1Updated),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1UpdateKey), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1UpdateKey), "DISABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1UpdateKey), "DISABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1UpdateKey), "YYYYMMDD"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1UpdateKey), "lb_test1_syslog1_description_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1UpdateKey), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1UpdateKey), "LOCAL1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1UpdateKey), "DEBUG"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1UpdateKey), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1UpdateKey), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1UpdateKey), "0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1UpdateKey), "NONE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1UpdateKey), "GMT_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1UpdateKey), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1UpdateKey), "YES"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1UpdateKey), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1UpdateKey), OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog2UpdateKey, &lbSyslog2Updated),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog2UpdateKey), &lbSyslog2.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog2UpdateKey), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog2UpdateKey), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog2UpdateKey), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog2UpdateKey), "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog2UpdateKey), "192.168.151.22"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog2UpdateKey), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog2UpdateKey), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog2UpdateKey), "lb_test1_syslog2"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog2UpdateKey), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog2UpdateKey), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog2UpdateKey), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog2UpdateKey), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog2UpdateKey), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog2UpdateKey), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog2UpdateKey), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog2UpdateKey), OS_TENANT_ID),
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

	r := "ecl_network_load_balancer_v2.lb_test1"
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
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttrSet(r, "load_balancer_plan_id"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "8"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "2", &lbIF3),
					resource.TestCheckResourceAttrPtr(r, "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr(r, "interfaces.2.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr(r, "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr(r, "interfaces.2.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "3", &lbIF4),
					resource.TestCheckResourceAttrPtr(r, "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr(r, "interfaces.3.ip_address", "192.168.154.11"),
					resource.TestCheckResourceAttr(r, "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr(r, "interfaces.3.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "4", &lbIF5),
					resource.TestCheckResourceAttrPtr(r, "interfaces.4.id", &lbIF5.ID),
					resource.TestCheckResourceAttr(r, "interfaces.4.ip_address", "192.168.155.11"),
					resource.TestCheckResourceAttr(r, "interfaces.4.slot_number", "5"),
					resource.TestCheckResourceAttr(r, "interfaces.4.status", "ACTIVE"),
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
					testAccCheckNetworkV2LoadBalancerExists(r, &lbUpdated),

					resource.TestCheckResourceAttrSet(r, "load_balancer_plan_id"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "2", &lbIF3Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr(r, "interfaces.2.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr(r, "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr(r, "interfaces.2.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "3", &lbIF4Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr(r, "interfaces.3.ip_address", "192.168.154.11"),
					resource.TestCheckResourceAttr(r, "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr(r, "interfaces.3.status", "ACTIVE"),
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

	r := "ecl_network_load_balancer_v2.lb_test1"
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
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttrSet(r, "load_balancer_plan_id"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "2", &lbIF3),
					resource.TestCheckResourceAttrPtr(r, "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr(r, "interfaces.2.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr(r, "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr(r, "interfaces.2.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "3", &lbIF4),
					resource.TestCheckResourceAttrPtr(r, "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr(r, "interfaces.3.ip_address", "192.168.154.11"),
					resource.TestCheckResourceAttr(r, "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr(r, "interfaces.3.status", "ACTIVE"),
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
					testAccCheckNetworkV2LoadBalancerExists(r, &lbUpdated),

					resource.TestCheckResourceAttrSet(r, "load_balancer_plan_id"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "8"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "2", &lbIF3Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.2.id", &lbIF3.ID),
					resource.TestCheckResourceAttr(r, "interfaces.2.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr(r, "interfaces.2.slot_number", "3"),
					resource.TestCheckResourceAttr(r, "interfaces.2.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "3", &lbIF4Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.3.id", &lbIF4.ID),
					resource.TestCheckResourceAttr(r, "interfaces.3.ip_address", "192.168.154.11"),
					resource.TestCheckResourceAttr(r, "interfaces.3.slot_number", "4"),
					resource.TestCheckResourceAttr(r, "interfaces.3.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "4", &lbIF5),
					resource.TestCheckResourceAttrPtr(r, "interfaces.4.id", &lbIF5.ID),
					resource.TestCheckResourceAttr(r, "interfaces.4.ip_address", "192.168.155.11"),
					resource.TestCheckResourceAttr(r, "interfaces.4.slot_number", "5"),
					resource.TestCheckResourceAttr(r, "interfaces.4.status", "ACTIVE"),
				),
			},
		},
	})
}

// TestAccNetworkV2LoadBalancer_forceNew tests that ForceNew attribute works functionally.
// Step 0: Create Load Balancer with 1 Interface.
// Step 1: Update Load Balancer availability_zone to force destroying/recreating.
func TestAccNetworkV2LoadBalancer_forceNew(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	r := "ecl_network_load_balancer_v2.lb_test1"
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
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttrPtr(r, "id", &lb.ID),
					resource.TestCheckResourceAttr(r, "name", "lb_test1"),
					resource.TestCheckResourceAttr(r, "availability_zone", OS_DEFAULT_ZONE),
					resource.TestCheckResourceAttr(r, "description", "load_balancer_test1_description"),
					resource.TestCheckResourceAttrSet(r, "load_balancer_plan_id"),
					resource.TestCheckResourceAttr(r, "default_gateway", ""),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),
					resource.TestCheckResourceAttr(r, "tenant_id", OS_TENANT_ID),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),
					resource.TestCheckResourceAttr(r, "syslog_servers.#", "0"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerMinimumUpdateAZ,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2LoadBalancerExists(r, &lbUpdated),
					testAccCheckNetworkV2LoadBalancerDoesNotExist(&lb),

					resource.TestCheckResourceAttrPtr(r, "id", &lbUpdated.ID),
					resource.TestCheckResourceAttr(r, "name", "lb_test1"),
					resource.TestCheckResourceAttr(r, "availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr(r, "description", "load_balancer_test1_description"),
					resource.TestCheckResourceAttrSet(r, "load_balancer_plan_id"),
					resource.TestCheckResourceAttr(r, "default_gateway", ""),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),
					resource.TestCheckResourceAttr(r, "tenant_id", OS_TENANT_ID),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),
					resource.TestCheckResourceAttr(r, "syslog_servers.#", "0"),
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

	r := "ecl_network_load_balancer_v2.lb_test1"
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
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "default_gateway", "192.168.151.1"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.description", "lb_test1_interface1_description"),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.name", "lb_test1_interface1"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_address", "192.168.151.31"),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.0.protocol", "vrrp"),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.0.vrid", "20"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces1And2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists(r, &lbUpdated),

					resource.TestCheckResourceAttr(r, "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.description", "lb_test1_interface1_description"),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.name", "lb_test1_interface1"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.1.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces3And2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_3", &n3),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_3_1", &sn3),
					testAccCheckNetworkV2LoadBalancerExists(r, &lbUpdated),

					resource.TestCheckResourceAttr(r, "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.description", "lb_test1_interface3_description"),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.153.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.name", "lb_test1_interface3"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.network_id", &n3.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.1.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists(r, &lbUpdated),

					resource.TestCheckResourceAttr(r, "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.description", ""),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.0.name", "Interface 1/1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.network_id", ""),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "DOWN"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "1", &lbIF2Updated),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.id", &lbIF2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr(r, "interfaces.1.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.1.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.1.network_id", &n2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.1.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.1.slot_number", "2"),
					resource.TestCheckResourceAttr(r, "interfaces.1.status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "0"),
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

	r := "ecl_network_load_balancer_v2.lb_test1"
	var n1 networks.Network
	var sn1 subnets.Subnet
	var lb load_balancers.LoadBalancer
	var lbSyslog1 load_balancer_syslog_servers.LoadBalancerSyslogServer
	var lbSyslog2 load_balancer_syslog_servers.LoadBalancerSyslogServer
	syslog1Key := "4120068917"
	syslog2Key := "1121660642"
	syslog2Update1Key := "4036206399"
	syslog2UpdateForceNew1Key := "4204560330"
	syslog2UpdateForceNew2Key := "1155946326"
	syslog2UpdateForceNew3Key := "2914665767"

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
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "1"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1Key, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1Key), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1Key), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1Key), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1Key), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1Key), OS_TENANT_ID),
				),
			},
			{
				// Step 1
				Config: testAccNetworkV2LoadBalancerModifySyslogServerAdd,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1Key, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1Key), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1Key), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1Key), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1Key), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1Key), OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog2Key, &lbSyslog2),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog2Key), &lbSyslog2.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog2Key), "DISABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog2Key), "DISABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog2Key), "YYYYMMDD"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog2Key), "lb_test1_syslog2_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog2Key), "192.168.151.22"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog2Key), "LOCAL1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog2Key), "DEBUG"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog2Key), "lb_test1_syslog2"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog2Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog2Key), "0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog2Key), "NONE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog2Key), "GMT_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog2Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog2Key), "YES"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog2Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog2Key), OS_TENANT_ID),
				),
			},
			{
				// Step 2
				Config: testAccNetworkV2LoadBalancerModifySyslogServerUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1Key, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1Key), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1Key), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1Key), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1Key), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1Key), OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog2Update1Key, &lbSyslog2Updated),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog2Update1Key), &lbSyslog2.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog2Update1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog2Update1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog2Update1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog2Update1Key), "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog2Update1Key), "192.168.151.22"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog2Update1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog2Update1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog2Update1Key), "lb_test1_syslog2"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog2Update1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog2Update1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog2Update1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog2Update1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog2Update1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog2Update1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog2Update1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog2Update1Key), OS_TENANT_ID),
				),
			},
			{
				// Step 3
				Config: testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1Key, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1Key), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1Key), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1Key), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1Key), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1Key), OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(&lbSyslog2Updated),
					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog2UpdateForceNew1Key, &lbSyslog2UpdatedForceNew1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog2UpdateForceNew1Key), &lbSyslog2UpdatedForceNew1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog2UpdateForceNew1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog2UpdateForceNew1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog2UpdateForceNew1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog2UpdateForceNew1Key), "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog2UpdateForceNew1Key), "192.168.151.22"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog2UpdateForceNew1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog2UpdateForceNew1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog2UpdateForceNew1Key), "lb_test1_syslog2_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog2UpdateForceNew1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog2UpdateForceNew1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog2UpdateForceNew1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog2UpdateForceNew1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog2UpdateForceNew1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog2UpdateForceNew1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog2UpdateForceNew1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog2UpdateForceNew1Key), OS_TENANT_ID),
				),
			},
			{
				// Step 4
				Config: testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1Key, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1Key), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1Key), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1Key), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1Key), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1Key), OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(&lbSyslog2UpdatedForceNew1),
					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog2UpdateForceNew2Key, &lbSyslog2UpdatedForceNew2),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog2UpdateForceNew2Key), &lbSyslog2UpdatedForceNew2.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog2UpdateForceNew2Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog2UpdateForceNew2Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog2UpdateForceNew2Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog2UpdateForceNew2Key), "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog2UpdateForceNew2Key), "192.168.151.22"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog2UpdateForceNew2Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog2UpdateForceNew2Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog2UpdateForceNew2Key), "lb_test1_syslog2_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog2UpdateForceNew2Key), "1514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog2UpdateForceNew2Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog2UpdateForceNew2Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog2UpdateForceNew2Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog2UpdateForceNew2Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog2UpdateForceNew2Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog2UpdateForceNew2Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog2UpdateForceNew2Key), OS_TENANT_ID),
				),
			},
			{
				// Step 5
				Config: testAccNetworkV2LoadBalancerModifySyslogServerUpdateForceNew3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "2"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1Key, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1Key), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1Key), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1Key), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1Key), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1Key), OS_TENANT_ID),

					testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(&lbSyslog2UpdatedForceNew2),
					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog2UpdateForceNew3Key, &lbSyslog2UpdatedForceNew3),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog2UpdateForceNew3Key), &lbSyslog2UpdatedForceNew3.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog2UpdateForceNew3Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog2UpdateForceNew3Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog2UpdateForceNew3Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog2UpdateForceNew3Key), "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog2UpdateForceNew3Key), "192.168.151.23"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog2UpdateForceNew3Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog2UpdateForceNew3Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog2UpdateForceNew3Key), "lb_test1_syslog2_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog2UpdateForceNew3Key), "1514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog2UpdateForceNew3Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog2UpdateForceNew3Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog2UpdateForceNew3Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog2UpdateForceNew3Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog2UpdateForceNew3Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog2UpdateForceNew3Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog2UpdateForceNew3Key), OS_TENANT_ID),
				),
			},
			{
				// Step 6
				Config: testAccNetworkV2LoadBalancerModifySyslogServerDelete,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "1"),

					testAccCheckNetworkV2LoadBalancerSyslogServerDoesNotExist(&lbSyslog1),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog2UpdateForceNew3Key, &lbSyslog2UpdatedForceNew4),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog2UpdateForceNew3Key), &lbSyslog2UpdatedForceNew4.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog2UpdateForceNew3Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog2UpdateForceNew3Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog2UpdateForceNew3Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog2UpdateForceNew3Key), "lb_test1_syslog2_description_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog2UpdateForceNew3Key), "192.168.151.23"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog2UpdateForceNew3Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog2UpdateForceNew3Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog2UpdateForceNew3Key), "lb_test1_syslog2_update"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog2UpdateForceNew3Key), "1514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog2UpdateForceNew3Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog2UpdateForceNew3Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog2UpdateForceNew3Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog2UpdateForceNew3Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog2UpdateForceNew3Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog2UpdateForceNew3Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog2UpdateForceNew3Key), OS_TENANT_ID),
				),
			},
		},
	})
}

// TestAccNetworkV2LoadBalancer_modifyInterfacesWithIPs tests simultaneous updating interface and other IP addresses.
// Step 0: Connect 1 new network to interface, set default_gateway and create 1 syslog server
//	-> network: [1], default_gateway: in network 1, syslog server IP: [in network 1]
// Step 1: Replace network and update other IPs to new network
//	-> network: [2], default_gateway: in network 2, syslog server IP: [in network 2]
func TestAccNetworkV2LoadBalancer_modifyInterfacesWithIPs(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	r := "ecl_network_load_balancer_v2.lb_test1"
	var lb, lbUpdated load_balancers.LoadBalancer
	var lbIF1 load_balancer_interfaces.LoadBalancerInterface
	var lbSyslog1 load_balancer_syslog_servers.LoadBalancerSyslogServer
	syslog1Key := "4120068917"
	syslog1UpdateKey := "1491927015"
	var n1, n2 networks.Network
	var sn1, sn2 subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2LoadBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2LoadBalancerInterfaces1WithSyslogServer,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_1", &n1),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_1_1", &sn1),
					testAccCheckNetworkV2LoadBalancerExists(r, &lb),

					resource.TestCheckResourceAttr(r, "default_gateway", "192.168.151.1"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.description", "lb_test1_interface1_description"),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.151.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.name", "lb_test1_interface1"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.network_id", &n1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_address", "192.168.151.31"),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.0.protocol", "vrrp"),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.0.vrid", "20"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "1"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1Key, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1Key), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1Key), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1Key), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1Key), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1Key), "192.168.151.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1Key), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1Key), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1Key), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1Key), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1Key), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1Key), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1Key), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1Key), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1Key), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1Key), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1Key), OS_TENANT_ID),
				),
			},
			{
				Config: testAccNetworkV2LoadBalancerInterfaces2WithSyslogServer,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2NetworkExists("ecl_network_network_v2.network_2", &n2),
					testAccCheckNetworkV2SubnetExists("ecl_network_subnet_v2.subnet_2_1", &sn2),
					testAccCheckNetworkV2LoadBalancerExists(r, &lbUpdated),

					resource.TestCheckResourceAttr(r, "default_gateway", "192.168.152.1"),
					resource.TestCheckResourceAttr(r, "status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "interfaces.#", "4"),

					testAccCheckNetworkV2LoadBalancerInterfaceExists(r, "0", &lbIF1),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.id", &lbIF1.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.description", "lb_test1_interface2_description"),
					resource.TestCheckResourceAttr(r, "interfaces.0.ip_address", "192.168.152.11"),
					resource.TestCheckResourceAttr(r, "interfaces.0.name", "lb_test1_interface2"),
					resource.TestCheckResourceAttrPtr(r, "interfaces.0.network_id", &n2.ID),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_address", ""),
					resource.TestCheckResourceAttr(r, "interfaces.0.virtual_ip_properties.#", "0"),
					resource.TestCheckResourceAttr(r, "interfaces.0.slot_number", "1"),
					resource.TestCheckResourceAttr(r, "interfaces.0.status", "ACTIVE"),

					resource.TestCheckResourceAttr(r, "syslog_servers.#", "1"),

					testAccCheckNetworkV2LoadBalancerSyslogServerExists(r, syslog1UpdateKey, &lbSyslog1),
					resource.TestCheckResourceAttrPtr(r, fmt.Sprintf("syslog_servers.%s.id", syslog1UpdateKey), &lbSyslog1.ID),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.acl_logging", syslog1UpdateKey), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.appflow_logging", syslog1UpdateKey), "ENABLED"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.date_format", syslog1UpdateKey), "MMDDYYYY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.description", syslog1UpdateKey), "lb_test1_syslog1_description"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.ip_address", syslog1UpdateKey), "192.168.152.21"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_facility", syslog1UpdateKey), "LOCAL0"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.log_level", syslog1UpdateKey), "ALERT|CRITICAL|EMERGENCY"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.name", syslog1UpdateKey), "lb_test1_syslog1"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.port_number", syslog1UpdateKey), "514"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.priority", syslog1UpdateKey), "20"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tcp_logging", syslog1UpdateKey), "ALL"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.time_zone", syslog1UpdateKey), "LOCAL_TIME"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.transport_type", syslog1UpdateKey), "UDP"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.user_configurable_log_messages", syslog1UpdateKey), "NO"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.status", syslog1UpdateKey), "ACTIVE"),
					resource.TestCheckResourceAttr(r, fmt.Sprintf("syslog_servers.%s.tenant_id", syslog1UpdateKey), OS_TENANT_ID),
				),
			},
		},
	})
}

func testAccCheckNetworkV2LoadBalancerDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL network client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_load_balancer_v2" {
			continue
		}

		if _, err := load_balancers.Get(networkClient, rs.Primary.ID).Extract(); err == nil {
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
			return fmt.Errorf("error creating ECL network client: %w", err)
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
			return fmt.Errorf("error creating ECL network client: %w", err)
		}

		if _, err := load_balancer_syslog_servers.Get(networkClient, loadBalancer.ID).Extract(); err != nil {
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
			return fmt.Errorf("error creating ECL network client: %w", err)
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
			return fmt.Errorf("error creating ECL network client: %w", err)
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
			return fmt.Errorf("error creating ECL network client: %w", err)
		}

		if _, err = load_balancer_syslog_servers.Get(networkClient, syslogServer.ID).Extract(); err != nil {
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
	"",
	testAccNetworkV2LoadBalancerInterface1,
)

var testAccNetworkV2LoadBalancerMinimumWithDisconnectedIF = fmt.Sprintf(`

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
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerInterface1DefaultValue,
	"",
)

var testAccNetworkV2LoadBalancerMinimumUpdateAZ = fmt.Sprintf(`

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
	"zone1_groupb",
	"",
	testAccNetworkV2LoadBalancerInterface1,
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
  depends_on = ["ecl_network_subnet_v2.subnet_1_1"]
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
	testAccNetworkV2LoadBalancerInterface3inSlot1,
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
  depends_on = ["ecl_network_subnet_v2.subnet_2_1"]
}
`,
	testAccNetworkV2LoadBalancerSingleNetworkAndSubnetPair2,
	testAccNetworkV2LoadBalancerPlan4IF,
	OS_DEFAULT_ZONE,
	testAccNetworkV2LoadBalancerDefaultGateway2,
	testAccNetworkV2LoadBalancerInterface2InSlot1,
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
    slot_number = 1
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
    slot_number = 1
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
    slot_number = 1
    description = "lb_test1_interface1_description"
    ip_address = "192.168.151.11"
    name = "lb_test1_interface1"
    network_id = "${ecl_network_network_v2.network_1.id}"
}
`

const testAccNetworkV2LoadBalancerInterface1DefaultValue = `
interfaces {
    slot_number = 1
    name = "Interface 1/1"
}
`

const testAccNetworkV2LoadBalancerInterface2 = `
interfaces {
    slot_number = 2
    description = "lb_test1_interface2_description"
    ip_address = "192.168.152.11"
    name = "lb_test1_interface2"
    network_id = "${ecl_network_network_v2.network_2.id}"
}
`

const testAccNetworkV2LoadBalancerInterface2InSlot1 = `
interfaces {
    slot_number = 1
    description = "lb_test1_interface2_description"
    ip_address = "192.168.152.11"
    name = "lb_test1_interface2"
    network_id = "${ecl_network_network_v2.network_2.id}"
}
`

const testAccNetworkV2LoadBalancerInterface2UpdateBasic = `
interfaces {
    slot_number = 2
    description = "lb_test1_interface2_description_update"
    ip_address = "192.168.152.12"
    name = "lb_test1_interface2_update"
    network_id = "${ecl_network_network_v2.network_2.id}"
}
`

const testAccNetworkV2LoadBalancerInterface3 = `
interfaces {
    slot_number = 3
    description = "lb_test1_interface3_description"
    ip_address = "192.168.153.11"
    name = "lb_test1_interface3"
    network_id = "${ecl_network_network_v2.network_3.id}"
}
`

const testAccNetworkV2LoadBalancerInterface3inSlot1 = `
interfaces {
    slot_number = 1
    description = "lb_test1_interface3_description"
    ip_address = "192.168.153.11"
    name = "lb_test1_interface3"
    network_id = "${ecl_network_network_v2.network_3.id}"
}
`

const testAccNetworkV2LoadBalancerInterface4 = `
interfaces {
    slot_number = 4
    description = "lb_test1_interface4_description"
    ip_address = "192.168.154.11"
    name = "lb_test1_interface4"
    network_id = "${ecl_network_network_v2.network_4.id}"
}
`

const testAccNetworkV2LoadBalancerInterface5 = `
interfaces {
    slot_number = 5
    description = "lb_test1_interface5_description"
    ip_address = "192.168.155.11"
    name = "lb_test1_interface5"
    network_id = "${ecl_network_network_v2.network_5.id}"
}
`

var testAccNetworkV2LoadBalancerSyslogServer1 = fmt.Sprintf(`
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
    tenant_id = %q
}
`, OS_TENANT_ID)

var testAccNetworkV2LoadBalancerSyslogServer1InInterface2 = fmt.Sprintf(`
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
    tenant_id = %q
}
`, OS_TENANT_ID)

var testAccNetworkV2LoadBalancerSyslogServer1UpdateBasic = fmt.Sprintf(`
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
    tenant_id = %q
}
`, OS_TENANT_ID)

var testAccNetworkV2LoadBalancerSyslogServer2 = fmt.Sprintf(`
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
    tenant_id = %q
}
`, OS_TENANT_ID)

var testAccNetworkV2LoadBalancerSyslogServer2UpdateBasic = fmt.Sprintf(`
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
    tenant_id = %q
}
`, OS_TENANT_ID)

var testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew1 = fmt.Sprintf(`
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
    tenant_id = %q
}
`, OS_TENANT_ID)

var testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew2 = fmt.Sprintf(`
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
    tenant_id = %q
}
`, OS_TENANT_ID)

var testAccNetworkV2LoadBalancerSyslogServer2UpdateForceNew3 = fmt.Sprintf(`
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
    tenant_id = %q
}
`, OS_TENANT_ID)
