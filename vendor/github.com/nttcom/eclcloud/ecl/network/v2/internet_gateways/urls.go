package internet_gateways

import "github.com/nttcom/eclcloud"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("internet_gateways", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("internet_gateways")
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
