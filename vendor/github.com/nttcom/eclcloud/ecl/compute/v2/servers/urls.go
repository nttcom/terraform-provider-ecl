package servers

import "github.com/nttcom/eclcloud"

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("servers")
}

func listURL(client *eclcloud.ServiceClient) string {
	return createURL(client)
}

func listDetailURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("servers", "detail")
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}

func getURL(client *eclcloud.ServiceClient, id string) string {
	return deleteURL(client, id)
}

func updateURL(client *eclcloud.ServiceClient, id string) string {
	return deleteURL(client, id)
}

func actionURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}

func metadatumURL(client *eclcloud.ServiceClient, id, key string) string {
	return client.ServiceURL("servers", id, "metadata", key)
}

func metadataURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "metadata")
}

func listAddressesURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "ips")
}

func listAddressesByNetworkURL(client *eclcloud.ServiceClient, id, network string) string {
	return client.ServiceURL("servers", id, "ips", network)
}

func passwordURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "os-server-password")
}
