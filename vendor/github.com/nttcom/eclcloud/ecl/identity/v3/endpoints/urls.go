package endpoints

import "github.com/nttcom/eclcloud"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("endpoints")
}

func endpointURL(client *eclcloud.ServiceClient, endpointID string) string {
	return client.ServiceURL("endpoints", endpointID)
}
