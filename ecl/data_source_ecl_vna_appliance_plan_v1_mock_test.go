package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccVNAV1AppliancePlanDataSource_basic(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "virtual_network_appliance_plans", "/v1.0/virtual_network_appliance_plans", testMockVNAV1AppliancePlansListNameQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVNAV1AppliancePlanDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AppliancePlanDataSourceID("data.ecl_vna_appliance_plan_v1.appliance_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "name", testAccVNAV1AppliancePlanDataSourcePlanName),
				),
			},
		},
	})
}

func TestMockedAccVNAV1AppliancePlanDataSource_queries(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "virtual_network_appliance_plans", "/v1.0/virtual_network_appliance_plans", testMockVNAV1AppliancePlansListNameQuery)
	mc.Register(t, "virtual_network_appliance_plans", "/v1.0/virtual_network_appliance_plans", testMockVNAV1AppliancePlansListIDQuery)
	mc.Register(t, "virtual_network_appliance_plans", "/v1.0/virtual_network_appliance_plans", testMockVNAV1AppliancePlansListDescriptionQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVNAV1AppliancePlanDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AppliancePlanDataSourceID("data.ecl_vna_appliance_plan_v1.appliance_plan_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_vna_appliance_plan_v1.appliance_plan_1", "name", testAccVNAV1AppliancePlanDataSourcePlanName),
				),
			},
			{
				Config: testAccVNAV1AppliancePlanDataSourceQueryID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AppliancePlanDataSourceID("data.ecl_vna_appliance_plan_v1.appliance_plan_2"),
				),
			},
			{
				Config: testAccVNAV1AppliancePlanDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1AppliancePlanDataSourceID("data.ecl_vna_appliance_plan_v1.appliance_plan_1"),
				),
			},
		},
	})
}

var testMockVNAV1AppliancePlansListNameQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      name:
        - %q
response:
    code: 200
    body: >
        {
          "virtual_network_appliance_plans": [
            {
              "id": "37556569-87f2-4699-b5ff-bf38e7cbf8a7",
              "name": %q,
              "description": "virtual_network_appliance_plans_description",
              "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
              "version": "",
              "flavor": "2CPU-8GB",
              "number_of_interfaces": 8,
              "enabled": true,
              "max_number_of_aap": 1,
              "licenses": [
                {
                  "license_type": "STD"
                }
              ],
              "availability_zones": [
                {
                  "availability_zone": "zone1_groupa",
                  "available": true,
                  "rank": 1
                },
                {
                  "availability_zone": "zone1_groupb",
                  "available": false,
                  "rank": 2
                }
              ]
            }
          ]
        }
`, testAccVNAV1AppliancePlanDataSourcePlanName, testAccVNAV1AppliancePlanDataSourcePlanName,
)
var testMockVNAV1AppliancePlansListIDQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      id:
        - 37556569-87f2-4699-b5ff-bf38e7cbf8a7
response:
    code: 200
    body: >
        {
          "virtual_network_appliance_plans": [
            {
              "id": "37556569-87f2-4699-b5ff-bf38e7cbf8a7",
              "name": %q,
              "description": "virtual_network_appliance_plans_description",
              "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
              "version": "",
              "flavor": "2CPU-8GB",
              "number_of_interfaces": 8,
              "enabled": true,
              "max_number_of_aap": 1,
              "licenses": [
                {
                  "license_type": "STD"
                }
              ],
              "availability_zones": [
                {
                  "availability_zone": "zone1_groupa",
                  "available": true,
                  "rank": 1
                },
                {
                  "availability_zone": "zone1_groupb",
                  "available": false,
                  "rank": 2
                }
              ]
            }
          ]
        }
`, testAccVNAV1AppliancePlanDataSourcePlanName,
)

var testMockVNAV1AppliancePlansListDescriptionQuery = fmt.Sprintf(`
request:
    method: GET
    query:
      id:
        - virtual_network_appliance_plans_description
response:
    code: 200
    body: >
        {
          "virtual_network_appliance_plans": [
            {
              "id": "37556569-87f2-4699-b5ff-bf38e7cbf8a7",
              "name": %q,
              "description": "virtual_network_appliance_plans_description",
              "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
              "version": "",
              "flavor": "2CPU-8GB",
              "number_of_interfaces": 8,
              "enabled": true,
              "max_number_of_aap": 1,
              "licenses": [
                {
                  "license_type": "STD"
                }
              ],
              "availability_zones": [
                {
                  "availability_zone": "zone1_groupa",
                  "available": true,
                  "rank": 1
                },
                {
                  "availability_zone": "zone1_groupb",
                  "available": false,
                  "rank": 2
                }
              ]
            }
          ]
        }
`, testAccVNAV1AppliancePlanDataSourcePlanName,
)
