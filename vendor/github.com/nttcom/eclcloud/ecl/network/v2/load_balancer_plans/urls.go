package load_balancer_plans

import "github.com/nttcom/eclcloud"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("load_balancer_plans", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("load_balancer_plans")
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
