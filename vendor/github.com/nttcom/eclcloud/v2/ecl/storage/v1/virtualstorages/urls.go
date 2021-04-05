package virtualstorages

import "github.com/nttcom/eclcloud/v2"

func createURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("virtual_storages")
}

func listURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("virtual_storages", "detail")
}

func deleteURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("virtual_storages", id)
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *eclcloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}
