package approval_requests

import "github.com/nttcom/eclcloud"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("approval-requests")
}

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("approval-requests", id)
}

func updateURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("approval-requests", id)
}
