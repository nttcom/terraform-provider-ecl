package fic_gateways

import "github.com/nttcom/eclcloud"

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("fic_gateways", id)
}

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("fic_gateways")
}
