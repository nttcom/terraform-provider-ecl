package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1PlanDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListNameQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListDescriptionQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListBandwidthQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListRedundancyQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfInterfacesQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfHealthMonitorsQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfListenersQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfPoliciesQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfRoutesQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfTargetGroupsQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfMembersQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfRulesQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListMaxNumberOfConditionsQuery)
	mc.Register(t, "plans", "/v1.0/plans", testMockMLBV1PlansListEnabledQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1PlanDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryBandwidth,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryRedundancy,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfInterfaces,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfHealthMonitors,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfListeners,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfPolicies,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfRoutes,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfTargetGroups,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfMembers,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfRules,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryMaxNumberOfConditions,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
			{
				Config: testAccMLBV1PlanDataSourceQueryEnabled,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "name", "50M_HA_4IF"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "bandwidth", "50M"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "redundancy", "HA"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_interfaces", "4"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_health_monitors", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_listeners", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_policies", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_routes", "25"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_target_groups", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_members", "100"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_rules", "50"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "max_number_of_conditions", "5"),
					resource.TestCheckResourceAttr("data.ecl_mlb_plan_v1.plan_1", "enabled", "true"),
				),
			},
		},
	})
}

var testAccMLBV1PlanDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  name = "50M_HA_4IF"
}
`)

var testMockMLBV1PlansListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - 50M_HA_4IF
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  description = "description"
}
`)

var testMockMLBV1PlansListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryBandwidth = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  bandwidth = "50M"
}
`)

var testMockMLBV1PlansListBandwidthQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    bandwidth:
      - 50M
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryRedundancy = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  redundancy = "HA"
}
`)

var testMockMLBV1PlansListRedundancyQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    redundancy:
      - HA
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfInterfaces = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_interfaces = "4"
}
`)

var testMockMLBV1PlansListMaxNumberOfInterfacesQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_interfaces:
      - 4
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfHealthMonitors = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_health_monitors = "50"
}
`)

var testMockMLBV1PlansListMaxNumberOfHealthMonitorsQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_health_monitors:
      - 50
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfListeners = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_listeners = "50"
}
`)

var testMockMLBV1PlansListMaxNumberOfListenersQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_listeners:
      - 50
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfPolicies = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_policies = "50"
}
`)

var testMockMLBV1PlansListMaxNumberOfPoliciesQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_policies:
      - 50
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfRoutes = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_routes = "25"
}
`)

var testMockMLBV1PlansListMaxNumberOfRoutesQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_routes:
      - 25
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfTargetGroups = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_target_groups = "50"
}
`)

var testMockMLBV1PlansListMaxNumberOfTargetGroupsQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_target_groups:
      - 50
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfMembers = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_members = "100"
}
`)

var testMockMLBV1PlansListMaxNumberOfMembersQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_members:
      - 100
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfRules = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_rules = "50"
}
`)

var testMockMLBV1PlansListMaxNumberOfRulesQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_rules:
      - 50
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryMaxNumberOfConditions = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  max_number_of_conditions = "5"
}
`)

var testMockMLBV1PlansListMaxNumberOfConditionsQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    max_number_of_conditions:
      - 5
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)

var testAccMLBV1PlanDataSourceQueryEnabled = fmt.Sprintf(`
data "ecl_mlb_plan_v1" "plan_1" {
  enabled = "true"
}
`)

var testMockMLBV1PlansListEnabledQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    enabled:
      - true
response:
  code: 200
  body: >
    {
      "plans": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "50M_HA_4IF",
          "description": "description",
          "bandwidth": "50M",
          "redundancy": "HA",
          "max_number_of_interfaces": 4,
          "max_number_of_health_monitors": 50,
          "max_number_of_listeners": 50,
          "max_number_of_policies": 50,
          "max_number_of_routes": 25,
          "max_number_of_target_groups": 50,
          "max_number_of_members": 100,
          "max_number_of_rules": 50,
          "max_number_of_conditions": 5,
          "enabled": true
        }
      ]
    }
`)
