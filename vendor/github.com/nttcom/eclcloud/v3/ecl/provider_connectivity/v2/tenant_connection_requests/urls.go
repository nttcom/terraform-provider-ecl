package tenant_connection_requests

import "github.com/nttcom/eclcloud/v3"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("tenant_connection_requests")
}

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("tenant_connection_requests", id)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("tenant_connection_requests")
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("tenant_connection_requests", id)
}

func updateURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("tenant_connection_requests", id)
}
