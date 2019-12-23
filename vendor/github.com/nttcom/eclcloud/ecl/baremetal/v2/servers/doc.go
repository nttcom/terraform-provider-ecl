/*
Package servers contains functionality for working with
ECL Baremetal Server resources.

Example to create server

	createOpts := servers.CreateOpts{
		Name: "server-test-1",
		Networks: []servers.CreateOptsNetwork{
			{
				UUID:    "d32019d3-bc6e-4319-9c1d-6722fc136a22",
				FixedIP: "10.0.0.100",
			},
		},
		AdminPass:        "aabbccddeeff",
		ImageRef:         "b5660a6e-4b46-4be3-9707-6b47221b454f",
		FlavorRef:        "05184ba3-00ba-4fbc-b7a2-03b62b884931",
		AvailabilityZone: "zone1-groupa",
		UserData:         "IyEvYmluL2Jhc2gKZWNobyAiS3VtYSBQb3N0IEluc3RhbGwgU2NyaXB0IiA+PiAvaG9tZS9iaWcvcG9zdC1pbnN0YWxsLXNjcmlwdA==",
		RaidArrays: []servers.CreateOptsRaidArray{
			{
				PrimaryStorage: true,
				Partitions: []map[string]interface{}{
					{
						"lvm":             true,
						"partition_label": "primary-part1",
					},
					{
						"lvm":             false,
						"size":            "100G",
						"partition_label": "var",
					},
				},
			},
			{
				RaidCardHardwareID: "raid_card_uuid",
				DiskHardwareIDs: []string{
					"disk1_uuid",
					"disk2_uuid",
					"disk3_uuid",
					"disk4_uuid",
				},
				Partitions: []map[string]interface{}{
					{
						"lvm":             true,
						"partition_label": "secondary-part1",
					},
				},
			},
		},
		LVMVolumeGroups: []servers.CreateOptsLVMVolumeGroup{
			{
				VGLabel: "VG_root",
				PhysicalVolumePartitionLabels: []string{
					"primary-part1",
					"secondary-part1",
				},
				LogicalVolumes: []map[string]string{
					{
						"size":     "300G",
						"lv_label": "LV_root",
					},
					{
						"size":     "2G",
						"lv_label": "LV_swap",
					},
				},
			},
		},
		Filesystems: []servers.CreateOptsFilesystem{
			{
				Label:      "LV_root",
				FSType:     "xfs",
				MountPoint: "/",
			},
			{
				Label:      "var",
				FSType:     "xfs",
				MountPoint: "/var",
			},
			{
				Label:  "LV_swap",
				FSType: "swap",
			},
		},
		Metadata: map[string]string{
			"foo": "bar",
		},
	}
	server, err := servers.Create(client, createOpts).Extract()

Example to list servers

	listOpts := servers.ListOpts{
		Status: "ACTIVE",
	}

	allPages, err := servers.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}

	for _, server := range allServers {
		fmt.Printf("%+v", server)
	}

Example to delete server

	err = servers.Delete(client, "server-id"").ExtractErr()
	if err != nil {
		panic(err)
	}

*/
package servers
