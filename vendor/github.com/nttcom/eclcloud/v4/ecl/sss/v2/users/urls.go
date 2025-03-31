package users

import "github.com/nttcom/eclcloud/v4"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("users")
}

func getURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("users")
}

func deleteURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

func updateURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}
