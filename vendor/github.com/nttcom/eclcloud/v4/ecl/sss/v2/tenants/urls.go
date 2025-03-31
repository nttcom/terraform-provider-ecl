package tenants

import "github.com/nttcom/eclcloud/v4"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}

func getURL(client *eclcloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}
