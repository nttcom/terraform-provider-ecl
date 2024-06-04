/*
Package flavors contains functionality for working with
ECL Baremetal Server's flavor resources.

Example to list flavors

	listOpts := flavors.ListOpts{
		TenantID: "a99e9b4e620e4db09a2dfb6e42a01e66",
	}

	allPages, err := flavors.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		panic(err)
	}

	for _, flavor := range allFlavors {
		fmt.Printf("%+v", flavor)
	}
*/
package flavors
