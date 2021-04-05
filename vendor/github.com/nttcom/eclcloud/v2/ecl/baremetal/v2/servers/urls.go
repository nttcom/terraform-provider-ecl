package servers

import (
	"github.com/nttcom/eclcloud/v2"
)

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("servers", "detail")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("servers")
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}
