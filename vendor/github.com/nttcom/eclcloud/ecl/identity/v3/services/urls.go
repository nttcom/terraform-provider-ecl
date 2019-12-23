package services

import "github.com/nttcom/eclcloud"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("services")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("services")
}

func serviceURL(client *eclcloud.ServiceClient, serviceID string) string {
	return client.ServiceURL("services", serviceID)
}

func updateURL(client *eclcloud.ServiceClient, serviceID string) string {
	return client.ServiceURL("services", serviceID)
}
