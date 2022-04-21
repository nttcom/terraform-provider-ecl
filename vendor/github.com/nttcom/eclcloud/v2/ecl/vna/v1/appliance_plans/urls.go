package appliance_plans

import "github.com/nttcom/eclcloud/v2"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("virtual_network_appliance_plans", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("virtual_network_appliance_plans")
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
