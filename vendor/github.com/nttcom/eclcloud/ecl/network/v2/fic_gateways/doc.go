/*
Package fic_gateways provides information of several service
in the Enterprise Cloud Compute service

Example to List FIC Gateways

	listOpts := fic_gateways.ListOpts{
		Status: "ACTIVE",
	}

	allPages, err := fic_gateways.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allFICGateways, err := fic_gateways.ExtractFICGateways(allPages)
	if err != nil {
		panic(err)
	}

	for _, ficGateway := range allFICGateways {
		fmt.Printf("%+v", ficGateway)
	}

Example to Show FIC Gateway

	id := "02dc9a22-129c-4b12-9936-4080f6a7ae44"
	ficGateway, err := fic_gateways.Get(client, id).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Print(ficGateway)

*/
package fic_gateways
