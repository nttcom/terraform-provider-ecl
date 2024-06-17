package plans

import (
	"github.com/nttcom/eclcloud/v3"
)

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("plans")
}

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("plans", id)
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func showURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
