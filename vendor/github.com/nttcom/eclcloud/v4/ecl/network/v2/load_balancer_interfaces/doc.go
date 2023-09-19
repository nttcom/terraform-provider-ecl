/*
Package load_balancer_interfaces contains functionality for working with
ECL Load Balancer Interface resources.

Example to List Load Balancer Interfaces

	listOpts := load_balancer_interfaces.ListOpts{
		Status: "ACTIVE",
	}

	allPages, err := load_balancer_interfaces.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allLoadBalancerInterfaces, err := load_balancer_interfaces.ExtractLoadBalancerInterfaces(allPages)
	if err != nil {
		panic(err)
	}

	for _, loadBalancerInterface := range allLoadBalancerInterfaces {
		fmt.Printf("%+v\n", loadBalancerInterface)
	}


Example to Show Load Balancer Interface

	loadBalancerInterfaceID := "f44e063c-5fea-45b8-9124-956995eafe2a"

	loadBalancerInterface, err := load_balancer_interfaces.Get(networkClient, loadBalancerInterfaceID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancerInterface)


Example to Update Load Balancer Interface

	loadBalancerInterfaceID := "f44e063c-5fea-45b8-9124-956995eafe2a"

	updateOpts := load_balancer_interfaces.UpdateOpts{
		Name:           "new_name",
	}

	loadBalancerInterface, err := load_balancer_interfaces.Update(networkClient, loadBalancerInterfaceID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}
*/
package load_balancer_interfaces
