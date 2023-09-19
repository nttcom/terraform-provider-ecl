/*
Package projects manages and retrieves Projects in the ECL SSS Service.

Example to List Tenants

	listOpts := tenants.ListOpts{
		Enabled: eclcloud.Enabled,
	}

	allPages, err := tenants.List(identityClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allProjects, err := tenants.ExtractProjects(allPages)
	if err != nil {
		panic(err)
	}

	for _, tenant := range allTenants {
		fmt.Printf("%+v\n", tenant)
	}

Example to Create a Tenant

	createOpts := projects.CreateOpts{
		Name:        "tenant_name",
		Description: "Tenant Description"
	}

	tenant, err := tenants.Create(identityClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a Tenant

	tenantID := "966b3c7d36a24facaf20b7e458bf2192"

	updateOpts := tenants.UpdateOpts{
		Description: "Tenant Description - New",
	}

	tenant, err := tenants.Update(identityClient, tenantID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Tenant

	tenantID := "966b3c7d36a24facaf20b7e458bf2192"
	err := projects.Delete(identityClient, projectID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package tenants
