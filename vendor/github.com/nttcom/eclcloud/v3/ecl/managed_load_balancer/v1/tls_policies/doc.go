/*
Package tls_policies contains functionality for working with ECL Managed Load Balancer resources.

Example to list tls policies

	listOpts := tls_policies.ListOpts{}

	allPages, err := tls_policies.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allTLSPolicies, err := tls_policies.ExtractTLSPolicies(allPages)
	if err != nil {
		panic(err)
	}

	for _, tLSPolicy := range allTLSPolicies {
		fmt.Printf("%+v\n", tLSPolicy)
	}

Example to show a tls policy

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	tLSPolicy, err := tls_policies.Show(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", tLSPolicy)
*/
package tls_policies
