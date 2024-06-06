package servers

import "github.com/nttcom/eclcloud/v3"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("servers")
}

func listDetailsURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("servers", "detail")
}

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("servers")
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}

func actionURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
