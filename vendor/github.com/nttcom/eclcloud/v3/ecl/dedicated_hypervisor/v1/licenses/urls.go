package licenses

import "github.com/nttcom/eclcloud/v3"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("licenses")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("licenses")
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("licenses", id)
}
