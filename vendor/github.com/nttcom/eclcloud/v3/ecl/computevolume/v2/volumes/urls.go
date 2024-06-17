package volumes

import "github.com/nttcom/eclcloud/v3"

func createURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("volumes")
}

func listURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("volumes", "detail")
}

func deleteURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id)
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *eclcloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}
