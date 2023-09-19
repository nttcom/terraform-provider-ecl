package networks

import "github.com/nttcom/eclcloud/v4"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("networks", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("networks")
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
