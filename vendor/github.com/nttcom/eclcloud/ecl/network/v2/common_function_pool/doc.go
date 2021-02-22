/*
Package common_function_pool contains functionality for working with
ECL Common Function Pool resources.

Example to List Common Function Pools

	listOpts := common_function_pool.ListOpts{
		Description: "general",
	}

	allPages, err := common_function_pool.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allCommonFunctionPools, err := common_function_pool.ExtractCommonFunctionPools(allPages)
	if err != nil {
		panic(err)
	}

	for _, commonFunctionPool := range allCommonFunctionPools {
		fmt.Printf("%+v\n", commonFunctionPool)
	}

Example to Show Common Function Pool

	commonFunctionPoolID := "c57066cc-9553-43a6-90de-asfdfesfffff"

	commonFunctionPool, err := common_function_pool.Get(networkClient, commonFunctionPoolID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", commonFunctionPool)

Example to look for Common Function Pool's ID by its name

	commonFunctionPoolName := "CF_Pool1"

	commonFunctionPoolID, err := common_function_pool.IDFromName(networkClient, commonFunctionPoolName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", commonFunctionPoolID)
*/
package common_function_pool
