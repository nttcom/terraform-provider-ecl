/*
Package users contains user management functionality on SSS.

Example to List users

	listOpts := users.ListOpts{}

	allPages, err := users.List(client, listOpts).AllPages()
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

	id := "ecid0000000001"
	user, err := users.Get(client, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", user)

Example to Create a user

	createOpts := users.CreateOpts{
		LoginID:        "sample",
		MailAddress:    "example@example.com",
		Password:       "Passw0rd",
		NotifyPassword: "true",
	}

	user, err := users.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a user

	userID := "ecid0000000001"
	loginID := "login-id-update"
	mailAddress := "update@example.com"
	newPassword := "NewPassw0rd"

	updateOpts := users.UpdateOpts{
		LoginID:     &loginID,
		MailAddress: &mailAddress,
		NewPassword: &newPassword,
	}

	result := users.Update(client, userID, updateOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete a user

	userID := "ecid0000000001"
	res := users.Delete(client, userID)
	if res.Err != nil {
		panic(res.Err)
	}

*/
package users
