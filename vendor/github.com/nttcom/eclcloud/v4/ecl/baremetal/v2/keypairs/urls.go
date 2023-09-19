package keypairs

import "github.com/nttcom/eclcloud/v4"

const resourcePath = "os-keypairs"

func resourceURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *eclcloud.ServiceClient) string {
	return resourceURL(c)
}

func createURL(c *eclcloud.ServiceClient) string {
	return resourceURL(c)
}

func getURL(c *eclcloud.ServiceClient, name string) string {
	return c.ServiceURL(resourcePath, name)
}

func deleteURL(c *eclcloud.ServiceClient, name string) string {
	return getURL(c, name)
}
