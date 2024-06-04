/*
Package appliances contains functionality for working with
ECL Commnon Function Gateway resources.

Example to List VirtualNetworkAppliances

	listOpts := virtual_network_appliances.ListOpts{
		TenantID: "a99e9b4e620e4db09a2dfb6e42a01e66",
	}

	allPages, err := virtual_network_appliances.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allVirtualNetworkAppliances, err := virtual_network_appliances.ExtractVirtualNetworkAppliances(allPages)
	if err != nil {
		panic(err)
	}

	for _, virtual_network_appliances := range allVirtualNetworkAppliances {
		fmt.Printf("%+v", virtual_network_appliances)
	}

Example to Create a virtual_network_appliances

	createOpts := virtual_network_appliances.CreateOpts{
		Name:         "network_1",
	}

	virtual_network_appliances, err := virtual_network_appliances.Create(networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a virtual_network_appliances

	virtualNetworkApplianceID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"

	updateOpts := virtual_network_appliances.UpdateOpts{
		Name: "new_name",
	}

	virtual_network_appliances, err := virtual_network_appliances.Update(networkClient, virtualNetworkApplianceID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a virtual_network_appliances

	virtualNetworkApplianceID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	err := virtual_network_appliances.Delete(networkClient, virtualNetworkApplianceID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package appliances
