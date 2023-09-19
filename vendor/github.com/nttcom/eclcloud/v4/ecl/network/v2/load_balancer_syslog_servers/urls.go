package load_balancer_syslog_servers

import "github.com/nttcom/eclcloud/v4"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("load_balancer_syslog_servers", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("load_balancer_syslog_servers")
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
