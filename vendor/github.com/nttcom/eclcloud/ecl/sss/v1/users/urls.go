package users

import "github.com/nttcom/eclcloud"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("users")
}

func getURL(client *eclcloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("users", tenantID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("users")
}

func deleteURL(client *eclcloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("users", tenantID)
}

func updateURL(client *eclcloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("users", tenantID)
}
