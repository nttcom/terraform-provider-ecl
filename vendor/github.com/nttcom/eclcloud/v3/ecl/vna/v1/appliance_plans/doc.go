/*
Package appliance_plans contains functionality for working with
ECL Virtual Network Appliance Plan resources.

Example to List Virtual Network Appliance Plans

	listOpts := appliance_plans.ListOpts{
		Description: "general",
	}

	allPages, err := appliance_plans.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allVirtualNetworkAppliancePlans, err := appliance_plans.ExtractVirtualNetworkAppliancePlans(allPages)
	if err != nil {
		panic(err)
	}

	for _, virtualNetworkAppliancePlan := range allVirtualNetworkAppliancePlans {
		fmt.Printf("%+v\n", virtualNetworkAppliancePlan)
	}

Example to Show Virtual Network Appliance Plan

	virtualNetworkAppliancePlanID := "37556569-87f2-4699-b5ff-bf38e7cbf8a7"

	virtualNetworkAppliancePlan, err := appliance_plans.Get(networkClient, virtualNetworkAppliancePlanID, nil).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", virtualNetworkAppliancePlan)

*/
package appliance_plans
