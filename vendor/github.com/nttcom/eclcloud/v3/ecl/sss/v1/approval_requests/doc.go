/*
Package approval_requests manages and retrieves approval requests in the Enterprise Cloud.

Example to List approval requests

	allPages, err := approval_requests.List(client).AllPages()
	if err != nil {
		panic(err)
	}

	allApprovalRequests, err := approval_requests.ExtractApprovalRequests(allPages)
	if err != nil {
		panic(err)
	}

	for _, approvalRequest := range allApprovalRequests {
		fmt.Printf("%+v\n", approvalRequest)
	}

Example to Get an approval requests

	requestID := "02471b45-3de0-4fc8-8469-a7cc52c378df"

	approvalRequest, err := approval_requests.Get(client, requestID).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", approvalRequest)

Example to Update an approval request

	requestID := "02471b45-3de0-4fc8-8469-a7cc52c378df"
	updateOpts := approval_requests.UpdateOpts{
		Status: "approved",
	}

	result := approval_requests.Update(client, requestID, updateOpts)
	if result.Err != nil {
		panic(result.Err)
	}

*/
package approval_requests
