package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1LoadBalancerDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListNameQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListDescriptionQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListConfigurationStatusQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListMonitoringStatusQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListOperationStatusQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListPrimaryAvailabilityZoneQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListSecondaryAvailabilityZoneQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListActiveAvailabilityZoneQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListRevisionQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListPlanIDQuery)
	mc.Register(t, "load_balancers", "/v1.0/load_balancers", testMockMLBV1LoadBalancersListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryConfigurationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryMonitoringStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryOperationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryPrimaryAvailabilityZone,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQuerySecondaryAvailabilityZone,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryActiveAvailabilityZone,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryRevision,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryPlanID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
			{
				Config: testAccMLBV1LoadBalancerDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "name", "load_balancer"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "monitoring_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "primary_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "secondary_availability_zone", "zone1_groupb"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "active_availability_zone", "zone1_groupa"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "revision", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_id", "00713021-9aea-41da-9a88-87760c08fa72"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "plan_name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.ip_address", "192.168.0.6"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.port", "514"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "syslog_servers.0.protocol", "udp"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.network_id", "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.virtual_ip_address", "192.168.0.1"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.#", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.0.ip_address", "192.168.0.2"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.1.ip_address", "192.168.0.3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.2.ip_address", "192.168.0.4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_load_balancer_v1.load_balancer_1", "interfaces.0.reserved_fixed_ips.3.ip_address", "192.168.0.5"),
				),
			},
		},
	})
}

var testAccMLBV1LoadBalancerDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  name = "load_balancer"
}
`)

var testMockMLBV1LoadBalancersListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - load_balancer
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  description = "description"
}
`)

var testMockMLBV1LoadBalancersListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryConfigurationStatus = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  configuration_status = "ACTIVE"
}
`)

var testMockMLBV1LoadBalancersListConfigurationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    configuration_status:
      - ACTIVE
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryMonitoringStatus = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  monitoring_status = "ACTIVE"
}
`)

var testMockMLBV1LoadBalancersListMonitoringStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    monitoring_status:
      - ACTIVE
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryOperationStatus = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  operation_status = "COMPLETE"
}
`)

var testMockMLBV1LoadBalancersListOperationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    operation_status:
      - COMPLETE
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryPrimaryAvailabilityZone = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  primary_availability_zone = "zone1_groupa"
}
`)

var testMockMLBV1LoadBalancersListPrimaryAvailabilityZoneQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    primary_availability_zone:
      - zone1_groupa
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQuerySecondaryAvailabilityZone = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  secondary_availability_zone = "zone1_groupb"
}
`)

var testMockMLBV1LoadBalancersListSecondaryAvailabilityZoneQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    secondary_availability_zone:
      - zone1_groupb
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryActiveAvailabilityZone = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  active_availability_zone = "zone1_groupa"
}
`)

var testMockMLBV1LoadBalancersListActiveAvailabilityZoneQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    active_availability_zone:
      - zone1_groupa
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryRevision = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  revision = "1"
}
`)

var testMockMLBV1LoadBalancersListRevisionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    revision:
      - 1
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryPlanID = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  plan_id = "00713021-9aea-41da-9a88-87760c08fa72"
}
`)

var testMockMLBV1LoadBalancersListPlanIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    plan_id:
      - 00713021-9aea-41da-9a88-87760c08fa72
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)

var testAccMLBV1LoadBalancerDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_load_balancer_v1" "load_balancer_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1LoadBalancersListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
response:
  code: 200
  body: >
    {
      "load_balancers": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "load_balancer",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "monitoring_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "primary_availability_zone": "zone1_groupa",
          "secondary_availability_zone": "zone1_groupb",
          "active_availability_zone": "zone1_groupa",
          "revision": 1,
          "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
          "plan_name": "50M_HA_4IF",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "syslog_servers": [
            {
              "ip_address": "192.168.0.6",
              "port": 514,
              "protocol": "udp"
            }
          ],
          "interfaces": [
            {
              "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
              "virtual_ip_address": "192.168.0.1",
              "reserved_fixed_ips": [
                {
                  "ip_address": "192.168.0.2"
                },
                {
                  "ip_address": "192.168.0.3"
                },
                {
                  "ip_address": "192.168.0.4"
                },
                {
                  "ip_address": "192.168.0.5"
                }
              ]
            }
          ]
        }
      ]
    }
`)
