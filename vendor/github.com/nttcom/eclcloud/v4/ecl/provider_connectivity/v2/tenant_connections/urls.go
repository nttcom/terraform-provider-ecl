package tenant_connections

import "github.com/nttcom/eclcloud/v4"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("tenant_connections")
}

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("tenant_connections", id)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("tenant_connections")
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("tenant_connections", id)
}

func updateURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("tenant_connections", id)
}
