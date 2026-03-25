/*
Package security_groups contains functionality for working with ECL Security Group resources.

Security Groups provide a way to define network access rules to control
inbound and outbound traffic to instances.

Example to List Security Groups

	listOpts := security_groups.ListOpts{
		TenantID: "tenant-id",
	}

	allPages, err := security_groups.List(networkClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allSecurityGroups, err := security_groups.ExtractSecurityGroups(allPages)
	if err != nil {
		panic(err)
	}

	for _, sg := range allSecurityGroups {
		fmt.Printf("%+v\n", sg)
	}

Example to Create a Security Group

	createOpts := security_groups.CreateOpts{
		Name:        "example-security-group",
		Description: "Example security group",
	}

	sg, err := security_groups.Create(networkClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a Security Group

	securityGroupID := "security-group-id"

	name := "updated-name"
	description := "updated description"
	updateOpts := security_groups.UpdateOpts{
		Name:        &name,
		Description: &description,
	}

	sg, err := security_groups.Update(networkClient, securityGroupID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Security Group

	securityGroupID := "security-group-id"
	err := security_groups.Delete(networkClient, securityGroupID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package security_groups
