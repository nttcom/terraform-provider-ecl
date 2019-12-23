package ports

import "github.com/nttcom/eclcloud"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("ports", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("ports")
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
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
