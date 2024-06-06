/*
Package load_balancers contains functionality for working with
ECL Load Balancer resources.

Example to List Load Balancers

	listOpts := load_balancers.ListOpts{
		Status: "ACTIVE",
	}

	allPages, err := load_balancers.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allLoadBalancers, err := load_balancers.ExtractLoadBalancers(allPages)
	if err != nil {
		panic(err)
	}

	for _, loadBalancer := range allLoadBalancers {
		fmt.Printf("%+v\n", loadBalancer)
	}


Example to Show Load Balancer

	loadBalancerID := "f44e063c-5fea-45b8-9124-956995eafe2a"

	loadBalancer, err := load_balancers.Get(networkClient, loadBalancerID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancer)


Example to Create a Load Balancer

	createOpts := load_balancers.CreateOpts{
		AvailabilityZone:   "zone1-groupa",
		Description:        "Load Balancer 1",
		LoadBalancerPlanID: "69bf1e91-73f6-41d5-84c4-91de21a9af05",
		Name:               "abcdefghijklmnopqrstuvwxyz",
		TenantID:           "5cc454d62d8c4a0595134b2632bf2263",
	}

	loadBalancer, err := load_balancers.Create(networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a Load Balancer

	loadBalancerID := "f44e063c-5fea-45b8-9124-956995eafe2a"
	name := "new_name"

	updateOpts := load_balancers.UpdateOpts{
		Name:           &name,
	}

	loadBalancer, err := load_balancers.Update(networkClient, loadBalancerID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Load Balancer

	loadBalancerID := "165fb257-2365-4c05-b368-a7bed21bb927"
	err := load_balancers.Delete(networkClient, loadBalancerID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package load_balancers
