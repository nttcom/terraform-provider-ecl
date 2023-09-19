package internet_services

import "github.com/nttcom/eclcloud/v4"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("internet_services", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("internet_services")
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}
