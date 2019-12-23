/*
Package users manages and retrieves users in the Enterprise Cloud Remote Console Access Service.

Example to List users

	allPages, err := users.List(rcaClient).AllPages()
	if err != nil {
		panic(err)
	}

	allUsers, err := users.ExtractUsers(allPages)
	if err != nil {
		panic(err)
	}

	for _, user := range allUsers {
		fmt.Printf("%+v\n", user)
	}

Example to Get a user

	username := "02471b45-3de0-4fc8-8469-a7cc52c378df"

	user, err := users.Get(rcaClient, username).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

Example to Create a user

	createOpts := users.CreateOpts{
		Password: "dummy_passw@rd",
	}

	result := users.Create(rcaClient, createOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Update a user

	username := "02471b45-3de0-4fc8-8469-a7cc52c378df"
	updateOpts := users.UpdateOpts{
		Password: "dummy_passw@rd",
	}

	result := users.Update(rcaClient, username, updateOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete a user

	username := "02471b45-3de0-4fc8-8469-a7cc52c378df"

	result := users.Delete(rcaClient, username)
	if result.Err != nil {
		panic(result.Err)
	}

*/
package users
