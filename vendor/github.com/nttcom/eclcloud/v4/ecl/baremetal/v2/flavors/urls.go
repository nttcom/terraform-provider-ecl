package flavors

import (
	"github.com/nttcom/eclcloud/v4"
)

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("flavors", "detail")
}
