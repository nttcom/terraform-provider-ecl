/*
Package load_balancer_plans contains functionality for working with
ECL Load Balancer Plan resources.

Example to List Load Balancer Plans

	listOpts := load_balancer_plans.ListOpts{
		Description: "general",
	}

	allPages, err := load_balancer_plans.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allLoadBalancerPlans, err := load_balancer_plans.ExtractLoadBalancerPlans(allPages)
	if err != nil {
		panic(err)
	}

	for _, loadBalancerPlan := range allLoadBalancerPlans {
		fmt.Printf("%+v\n", loadBalancerPlan)
	}

Example to Show Load Balancer Plan

	loadBalancerPlanID := "a46eeb5a-bc0a-40fa-b455-e5dc13b1220a"

	loadBalancerPlan, err := load_balancer_plans.Get(networkClient, loadBalancerPlanID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancerPlan)

*/
package load_balancer_plans
