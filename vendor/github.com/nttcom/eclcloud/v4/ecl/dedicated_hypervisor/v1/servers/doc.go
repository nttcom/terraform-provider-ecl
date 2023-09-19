/*
Package servers manages and retrieves servers in the Enterprise Cloud Dedicated Hypervisor Service.

Example to List servers

	listOpts := servers.ListOpts{
		Limit: 10,
	}

	allPages, err := servers.List(dhClient, listOpts).AllPages()
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

Example to List servers details

	listOpts := servers.ListOpts{
		Limit: 10,
	}

	allPages, err := servers.ListDetails(dhClient, listOpts).AllPages()
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

Example to Get a server

	serverID := "f42dbc37-4642-4628-8b47-50bf95d8fdd5"

	result := servers.Get(dhClient, serverID)
	if result.Err != nil {
		panic(result.Err)
	}

	server, err := result.Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", server)

Example to Create a server

	createOpts := servers.CreateOpts{
		Name: "test",
		Networks: []servers.Network{
			{
				UUID:           "94055904-6b2c-4839-a14a-c61c93a8bc48",
				Plane:          "data",
				SegmentationID: 6,
			},
			{
				UUID:           "94055904-6b2c-4839-a14a-c61c93a8bc48",
				Plane:          "data",
				SegmentationID: 6,
			},
		},
		ImageRef:  "dfd25820-b368-4012-997b-29a6d0cf8518",
		FlavorRef: "a830b61c-3155-4a61-b7ed-c450862845e6",
	}

	result := servers.Create(dhClient, createOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete a server

	serverID := "f42dbc37-4642-4628-8b47-50bf95d8fdd5"

	result := servers.Delete(dhClient, serverID)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Add license to a server

	serverID := "f42dbc37-4642-4628-8b47-50bf95d8fdd5"

	addLicenseOpts := servers.AddLicenseOpts{
		VmName: "Alice",
		LicenseTypes: []string{
			"Windows Server",
			"SQL Server Standard 2014",
		},
	}

	result := servers.AddLicense(dhClient, serverID, addLicenseOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Get result for add license to a server

	serverID := "f42dbc37-4642-4628-8b47-50bf95d8fdd5"

	getAddLicenseResultOpts := servers.GetAddLicenseResultOpts{
		JobID: AddLicenseJob.JobID,
	}

	result := servers.GetAddLicenseResult(dhClient, serverID, getAddLicenseResultOpts)
	if result.Err != nil {
		panic(result.Err)
	}

	job, err := result.Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", job)
*/
package servers
