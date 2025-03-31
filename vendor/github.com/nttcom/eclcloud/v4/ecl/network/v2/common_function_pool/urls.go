package common_function_pool

import (
	"github.com/nttcom/eclcloud/v4"
)

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("common_function_pools", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("common_function_pools")
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}
