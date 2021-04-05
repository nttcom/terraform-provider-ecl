package appliances

import (
	"github.com/nttcom/eclcloud/v2"
)

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("virtual_network_appliances", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("virtual_network_appliances")
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
