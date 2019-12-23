/*
Package keypairs provides the ability to manage key pairs.

Example to List Key Pairs

	allPages, err := keypairs.List(client).AllPages()
	if err != nil {
		panic(err)
	}

	allKeyPairs, err := keypairs.ExtractKeyPairs(allPages)
	if err != nil {
		panic(err)
	}

	for _, kp := range allKeyPairs {
		fmt.Printf("%+v\n", kp)
	}

Example to Create a Key Pair

	createOpts := keypairs.CreateOpts{
		Name: "keypair-name",
	}

	keypair, err := keypairs.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", keypair)

Example to Import a Key Pair

	createOpts := keypairs.CreateOpts{
		Name:      "keypair-name",
		PublicKey: "public-key",
	}

	keypair, err := keypairs.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Key Pair

	err := keypairs.Delete(client, "keypair-name").ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package keypairs
