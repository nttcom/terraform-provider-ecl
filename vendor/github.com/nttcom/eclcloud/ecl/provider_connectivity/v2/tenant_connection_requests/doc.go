/*
Package tenant_connection_requests manages and retrieves Tenant Connection Request in the Enterprise Cloud Provider Connectivity Service.

Example to List Tenant Connection Request

	allPages, err := tenant_connection_requests.List(tcrClient).AllPages()
	if err != nil {
		panic(err)
	}

	allTenantConnectionRequests, err := tenant_connection_requests.ExtractTenantConnectionRequests(allPages)
	if err != nil {
		panic(err)
	}

	for _, tenantConnectionRequest := range allTenantConnectionRequests {
		fmt.Printf("%+v\n", tenantConnectionRequest)
	}

Example to Get a Tenant Connection Request

	tenant_connection_request_id := "85a1dc30-2e48-11ea-9e55-525403060300"

	tenantConnectionRequest, err := tenant_connection_requests.Get(tcrClient, tenant_connection_request_id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", tenantConnectionRequest)

Example to Create a Tenant Connection Request

	createOpts := tenant_connection_requests.CreateOpts{
		TenantIDOther: "7e91b19b9baa423793ee74a8e1ff2be1",
		NetworkID: "c4d5fc41-b7e8-4f19-96f4-85299e54373c",
		Name: "create_test_name",
		Description: "create_test_desc",
		Tags: map[string]string{"foo", "bar"},
	}

	result := tenant_connection_requests.Create(tcrClient, createOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Update a Tenant Connection Request

	tenant_connection_request_id := "85a1dc30-2e48-11ea-9e55-525403060300"
	updateOpts := tenant_connection_requests.UpdateOpts{
		Name: "update_test_name",
		Description: "update_test_desc",
		Tags: map[string]string{
			"keyword1": "value1",
			"keyword2": "value2",
		},
		NameOther: "update_test_name_other",
		DescriptionOther: "update_test_desc_other",
		TagsOther: map[string]string{
			"keyword1": "value1",
			"keyword2": "value2",
		},
	}

	result := tenant_connection_requests.Update(tcrClient, tenant_connection_request_id, updateOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete a Tenant Connection Request

	tenant_connection_request_id := "85a1dc30-2e48-11ea-9e55-525403060300"

	result := tenant_connection_requests.Delete(tcrClient, tenant_connection_request_id)
	if result.Err != nil {
		panic(result.Err)
	}

*/
package tenant_connection_requests
