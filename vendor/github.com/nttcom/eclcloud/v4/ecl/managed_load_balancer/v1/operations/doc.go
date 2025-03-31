/*
Package operations contains functionality for working with ECL Managed Load Balancer resources.

Example to list operations

	listOpts := operations.ListOpts{}

	allPages, err := operations.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allOperations, err := operations.ExtractOperations(allPages)
	if err != nil {
		panic(err)
	}

	for _, operation := range allOperations {
		fmt.Printf("%+v\n", operation)
	}

Example to show a operation

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	operation, err := operations.Show(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", operation)
*/
package operations
