package users

import "github.com/nttcom/eclcloud"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("users")
}

func getURL(client *eclcloud.ServiceClient, name string) string {
	return client.ServiceURL("users", name)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("users")
}

func deleteURL(client *eclcloud.ServiceClient, name string) string {
	return client.ServiceURL("users", name)
}

func updateURL(client *eclcloud.ServiceClient, name string) string {
	return client.ServiceURL("users", name)
}
