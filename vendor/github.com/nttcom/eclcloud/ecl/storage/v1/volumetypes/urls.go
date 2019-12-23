package volumetypes

import (
	"github.com/nttcom/eclcloud"
)

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("volume_types", id)
}

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("volume_types", "detail")
}
