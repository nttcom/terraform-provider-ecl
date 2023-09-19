package images

import "github.com/nttcom/eclcloud/v4"

func listDetailURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("images", "detail")
}

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}
