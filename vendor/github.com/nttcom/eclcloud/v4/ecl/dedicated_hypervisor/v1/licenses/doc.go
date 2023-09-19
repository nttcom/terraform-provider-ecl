/*
Package licenses manages and retrieves license in the Enterprise Cloud Dedicated Hypervisor Service.

Example to List Licenses

	listOpts := licenses.ListOpts{
		LicenseType: "vCenter Server 6.x Standard",
	}

	allPages, err := licenses.List(dhClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allLicenses, err := licenses.ExtractLicenses(allPages)
	if err != nil {
		panic(err)
	}

	for _, license := range allLicenses {
		fmt.Printf("%+v\n", license)
	}

Example to Create a License

	createOpts := licenses.CreateOpts{
		LicenseType: "vCenter Server 6.x Standard",
	}

	result := licenses.Create(dhClient, createOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete a license

	licenseID := "02471b45-3de0-4fc8-8469-a7cc52c378df"

	result := licenses.Delete(dhClient, licenseID)
	if result.Err != nil {
		panic(result.Err)
	}
*/
package licenses
