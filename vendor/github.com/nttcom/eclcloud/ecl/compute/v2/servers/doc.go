/*
Package servers provides information and interaction with the server API
resource in the Enterprise Cloud Compute service.

A server is a virtual machine instance in the compute system. In order for
one to be provisioned, a valid flavor and image are required.

Example to List Servers

	listOpts := servers.ListOpts{
		AllTenants: true,
	}

	allPages, err := servers.List(computeClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}

	for _, server := range allServers {
		fmt.Printf("%+v\n", server)
	}

Example to Get a Server

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	server, err := servers.Get(client, serverID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", server)

Example to Create a Server

	createOpts := servers.CreateOpts{
		Name:      "server_name",
		ImageRef:  "image-uuid",
		FlavorRef: "flavor-uuid",
	}

	result := servers.Create(computeClient, createOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Update a Server

	name := "update_name"
	updateOpts := servers.UpdateOpts{Name: &name}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	result := servers.Update(client, serverID, updateOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete a Server

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	result := servers.Delete(computeClient, serverID)
	if result.Err != nil {
		panic(err)
	}

Example to Show Metadata a server

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	metadata, err := servers.Metadata(client, serverID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", metadata)

Example to Show details for a Metadata item by key for a Server

	key := "key"

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	metadatum, err := servers.Metadatum(client, serverID, key).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", metadatum)

Example to Create Metadata a Server

	createMetadatumOpts := servers.MetadatumOpts{"key": "value"}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	result := servers.CreateMetadatum(client, serverID, createMetadatumOpts)
	if err != nil {
		panic(result.Err)
	}

Example to Update Metadata a Server

	updateMetadataOpts := servers.MetadataOpts{"key": "update"}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	result := servers.UpdateMetadata(client, serverID, updateMetadataOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete Metadata a Server

	key := "key"

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	result := servers.DeleteMetadatum(client, serverID, key)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Reset Metadata a Server

	resetMetadataOpts := servers.MetadataOpts{"key2": "val2"}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	result := servers.ResetMetadata(client, serverID, resetMetadataOpts)
	if result.Err != nil {
		panic(nil)
	}

Example to Resize a Server

	resizeOpts := servers.ResizeOpts{
		FlavorRef: "flavor-uuid",
	}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	result := servers.Resize(computeClient, serverID, resizeOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Snapshot a Server

	snapshotOpts := servers.CreateImageOpts{
		Name: "snapshot_name",
	}

	serverID := "d9072956-1560-487c-97f2-18bdf65ec749"

	result := servers.CreateImage(computeClient, serverID, snapshotOpts)
	if result.Err != nil {
		panic(result.Err)
	}

*/
package servers
