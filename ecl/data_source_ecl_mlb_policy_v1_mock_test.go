package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1PolicyDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListNameQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListDescriptionQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListConfigurationStatusQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListOperationStatusQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListAlgorithmQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListPersistenceQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListIdleTimeoutQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListSorryPageUrlQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListSourceNatQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListCertificateIDQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListHealthMonitorIDQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListListenerIDQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListDefaultTargetGroupIDQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListTLSPolicyIDQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListLoadBalancerIDQuery)
	mc.Register(t, "policies", "/v1.0/policies", testMockMLBV1PoliciesListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1PolicyDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryConfigurationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryOperationStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryAlgorithm,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryPersistence,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryIdleTimeout,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQuerySorryPageUrl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQuerySourceNat,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryCertificateID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryHealthMonitorID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryListenerID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryDefaultTargetGroupID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryTLSPolicyID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryLoadBalancerID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
			{
				Config: testAccMLBV1PolicyDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "name", "policy"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "configuration_status", "ACTIVE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "operation_status", "COMPLETE"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "load_balancer_id", "67fea379-cff0-4191-9175-de7d6941a040"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "persistence", "cookie"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "idle_timeout", "600"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "sorry_page_url", "https://example.com/sorry"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "source_nat", "enable"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "certificate_id", "f57a98fe-d63e-4048-93a0-51fe163f30d7"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "health_monitor_id", "dd7a96d6-4e66-4666-baca-a8555f0c472c"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "listener_id", "68633f4f-f52a-402f-8572-b8173418904f"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "default_target_group_id", "a44c4072-ed90-4b50-a33a-6b38fb10c7db"),
					resource.TestCheckResourceAttr("data.ecl_mlb_policy_v1.policy_1", "tls_policy_id", "4ba79662-f2a1-41a4-a3d9-595799bbcd86"),
				),
			},
		},
	})
}

var testAccMLBV1PolicyDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  name = "policy"
}
`)

var testMockMLBV1PoliciesListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - policy
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  description = "description"
}
`)

var testMockMLBV1PoliciesListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryConfigurationStatus = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  configuration_status = "ACTIVE"
}
`)

var testMockMLBV1PoliciesListConfigurationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    configuration_status:
      - ACTIVE
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryOperationStatus = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  operation_status = "COMPLETE"
}
`)

var testMockMLBV1PoliciesListOperationStatusQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    operation_status:
      - COMPLETE
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryAlgorithm = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  algorithm = "round-robin"
}
`)

var testMockMLBV1PoliciesListAlgorithmQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    algorithm:
      - round-robin
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryPersistence = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  persistence = "cookie"
}
`)

var testMockMLBV1PoliciesListPersistenceQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    persistence:
      - cookie
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryIdleTimeout = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  idle_timeout = "600"
}
`)

var testMockMLBV1PoliciesListIdleTimeoutQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    idle_timeout:
      - 600
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQuerySorryPageUrl = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  sorry_page_url = "https://example.com/sorry"
}
`)

var testMockMLBV1PoliciesListSorryPageUrlQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    sorry_page_url:
      - https://example.com/sorry
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQuerySourceNat = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  source_nat = "enable"
}
`)

var testMockMLBV1PoliciesListSourceNatQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    source_nat:
      - enable
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryCertificateID = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  certificate_id = "f57a98fe-d63e-4048-93a0-51fe163f30d7"
}
`)

var testMockMLBV1PoliciesListCertificateIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    certificate_id:
      - f57a98fe-d63e-4048-93a0-51fe163f30d7
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryHealthMonitorID = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  health_monitor_id = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
}
`)

var testMockMLBV1PoliciesListHealthMonitorIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    health_monitor_id:
      - dd7a96d6-4e66-4666-baca-a8555f0c472c
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryListenerID = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  listener_id = "68633f4f-f52a-402f-8572-b8173418904f"
}
`)

var testMockMLBV1PoliciesListListenerIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    listener_id:
      - 68633f4f-f52a-402f-8572-b8173418904f
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryDefaultTargetGroupID = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  default_target_group_id = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
}
`)

var testMockMLBV1PoliciesListDefaultTargetGroupIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    default_target_group_id:
      - a44c4072-ed90-4b50-a33a-6b38fb10c7db
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryTLSPolicyID = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  tls_policy_id = "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
}
`)

var testMockMLBV1PoliciesListTLSPolicyIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tls_policy_id:
      - 4ba79662-f2a1-41a4-a3d9-595799bbcd86
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryLoadBalancerID = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  load_balancer_id = "67fea379-cff0-4191-9175-de7d6941a040"
}
`)

var testMockMLBV1PoliciesListLoadBalancerIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    load_balancer_id:
      - 67fea379-cff0-4191-9175-de7d6941a040
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)

var testAccMLBV1PolicyDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_policy_v1" "policy_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1PoliciesListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
response:
  code: 200
  body: >
    {
      "policies": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "policy",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "configuration_status": "ACTIVE",
          "operation_status": "COMPLETE",
          "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "algorithm": "round-robin",
          "persistence": "cookie",
          "idle_timeout": 600,
          "sorry_page_url": "https://example.com/sorry",
          "source_nat": "enable",
          "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
          "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
          "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
          "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
          "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
      ]
    }
`)
