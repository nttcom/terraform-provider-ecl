/*
Package tenant_connections manages and retrieves Tenant Connection in the Enterprise Cloud Provider Connectivity Service.

Example to List Tenant Connection

	allPages, err := tenant_connections.List(tcClient).AllPages()
	if err != nil {
		panic(err)
	}

	allTenantConnections, err := tenant_connections.ExtractTenantConnections(allPages)
	if err != nil {
		panic(err)
	}

	for _, tenantConnection := range allTenantConnections {
		fmt.Printf("%+v\n", tenantConnection)
	}

Example to Get a Tenant Connection

	tenant_connection_id := "ea5d975c-bd31-11e7-bcac-0050569c850d"

	tenantConnection, err := tenant_connections.Get(tcClient, tenant_connection_id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", tenantConnection)

Example to Create a Tenant Connection

	createOpts := tenant_connections.CreateOpts{
		Name:                      "create_test_name",
		Description:               "create_test_desc",
		Tags: map[string]string{
			"test_tags": "test",
		},
		TenantConnectionRequestID: "21b344d8-be11-11e7-bf3c-0050569c850d",
		DeviceType:                "ECL::VirtualNetworkAppliance::VSRX",
		DeviceID:                  "c291f4c4-a680-4db0-8b88-7e579f0aaa37",
		DeviceInterfaceID:		   "interface_2",
		AttachmentOpts: tenant_connections.Vna{
			FixedIPs: []tenant_connections.VnaFixedIPs{
				IPAddress: "192.168.1.3",
			},
		},
	}

	result := tenant_connections.Create(tcClient, createOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Update a Tenant Connection

	tenant_connection_id := "ea5d975c-bd31-11e7-bcac-0050569c850d"

	updateOpts := tenant_connections.UpdateOpts{
		Name: "test_name",
		Description: "test_desc",
		Tags: map[string]string{
			"test_tags": "test",
		},
		NameOther: "test_name_other",
		DescriptionOther: "test_desc_other",
		TagsOther: map[string]string{
			"test_tags_other": "test_other",
		},
	}

	result := tenant_connections.Update(tcClient, tenant_connection_id, updateOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete a Tenant Connection

	tenant_connection_id := "ea5d975c-bd31-11e7-bcac-0050569c850d"

	result := tenant_connections.Delete(tcClient, tenant_connection_id)
	if result.Err != nil {
		panic(result.Err)
	}

*/
package tenant_connections
