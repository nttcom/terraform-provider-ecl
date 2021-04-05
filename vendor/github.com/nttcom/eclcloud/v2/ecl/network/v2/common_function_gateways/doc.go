/*
Package common_function_gateways contains functionality for working with
ECL Commnon Function Gateway resources.

Example to List CommonFunctionGateways

	listOpts := common_function_gateways.ListOpts{
		TenantID: "a99e9b4e620e4db09a2dfb6e42a01e66",
	}

	allPages, err := common_function_gateways.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allCommonFunctionGateways, err := common_function_gateways.ExtractCommonFunctionGateways(allPages)
	if err != nil {
		panic(err)
	}

	for _, common_function_gateways := range allCommonFunctionGateways {
		fmt.Printf("%+v", common_function_gateways)
	}

Example to Create a common_function_gateways

	createOpts := common_function_gateways.CreateOpts{
		Name:         "network_1",
	}

	common_function_gateways, err := common_function_gateways.Create(networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a common_function_gateways

	commonFunctionGatewayID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"

	updateOpts := common_function_gateways.UpdateOpts{
		Name: "new_name",
	}

	common_function_gateways, err := common_function_gateways.Update(networkClient, commonFunctionGatewayID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a common_function_gateways

	commonFunctionGatewayID := "484cda0e-106f-4f4b-bb3f-d413710bbe78"
	err := common_function_gateways.Delete(networkClient, commonFunctionGatewayID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package common_function_gateways
