package load_balancer_interfaces

import "github.com/nttcom/eclcloud/v3"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("load_balancer_interfaces", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("load_balancer_interfaces")
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func updateURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
