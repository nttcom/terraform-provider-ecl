package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1HealthMonitorDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListNameQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListDescriptionQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListConfigurationStatusQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListOperationStatusQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListPortQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListProtocolQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListIntervalQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListRetryQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListTimeoutQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListPathQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListHttpStatusCodeQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListLoadBalancerIDQuery)
	mc.Register(t, "health_monitors", "/v1.0/health_monitors", testMockMLBV1HealthMonitorsListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryConfigurationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryOperationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryPort,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryProtocol,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryInterval,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryRetry,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryTimeout,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryPath,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryHttpStatusCode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryLoadBalancerID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
			{
				Config: testAccMLBV1HealthMonitorDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "name", "health_monitor"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "port", "80"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "protocol", "http"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "interval", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "retry", "3"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "timeout", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "path", "/health"),
					resource.TestCheckResourceAttr("data.ecl_mlb_health_monitor_v1.health_monitor_1", "http_status_code", "200-299"),
				),
			},
		},
	})
}

var testAccMLBV1HealthMonitorDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  name = "health_monitor"
}
`)

var testMockMLBV1HealthMonitorsListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - health_monitor
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  description = "description"
}
`)

var testMockMLBV1HealthMonitorsListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryConfigurationStatus = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  configuration_status = "ACTIVE"
}
`)

var testMockMLBV1HealthMonitorsListConfigurationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    configuration_status:
      - ACTIVE
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryOperationStatus = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  operation_status = "COMPLETE"
}
`)

var testMockMLBV1HealthMonitorsListOperationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    operation_status:
      - COMPLETE
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryPort = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  port = "80"
}
`)

var testMockMLBV1HealthMonitorsListPortQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    port:
      - 80
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryProtocol = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  protocol = "http"
}
`)

var testMockMLBV1HealthMonitorsListProtocolQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    protocol:
      - http
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryInterval = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  interval = "5"
}
`)

var testMockMLBV1HealthMonitorsListIntervalQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    interval:
      - 5
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryRetry = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  retry = "3"
}
`)

var testMockMLBV1HealthMonitorsListRetryQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    retry:
      - 3
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryTimeout = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  timeout = "5"
}
`)

var testMockMLBV1HealthMonitorsListTimeoutQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    timeout:
      - 5
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryPath = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  path = "/health"
}
`)

var testMockMLBV1HealthMonitorsListPathQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    path:
      - /health
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryHttpStatusCode = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  http_status_code = "200-299"
}
`)

var testMockMLBV1HealthMonitorsListHttpStatusCodeQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    http_status_code:
      - 200-299
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryLoadBalancerID = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1HealthMonitorsListLoadBalancerIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 67fea379-cff0-4191-9175-de7d6941a040
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)

var testAccMLBV1HealthMonitorDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_health_monitor_v1" "health_monitor_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1HealthMonitorsListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
response:
  code: 200
  body: >
    {
      "health_monitors": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "health_monitor",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "port": 80,
          "protocol": "http",
          "interval": 5,
          "retry": 3,
          "timeout": 5,
          "path": "/health",
          "http_status_code": "200-299"
        }
      ]
    }
`)
