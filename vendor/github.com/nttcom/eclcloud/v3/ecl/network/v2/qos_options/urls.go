package qos_options

import "github.com/nttcom/eclcloud/v3"

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("qos_options", id)
}

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("qos_options")
}
