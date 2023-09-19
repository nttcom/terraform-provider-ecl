/*
Package availabilityzones provides the ability to get lists and detailed
availability zone information and to extend a server result with
availability zone information.

Example of Get Availability Zone Information

	allPages, err := availabilityzones.List(client).AllPages()
	if err != nil {
		panic(err)
	}

	availabilityZoneInfo, err := availabilityzones.ExtractAvailabilityZones(allPages)
	if err != nil {
		panic(err)
	}

	for _, zoneInfo := range availabilityZoneInfo {
  		fmt.Printf("%+v\n", zoneInfo)
	}
*/
package availabilityzones
