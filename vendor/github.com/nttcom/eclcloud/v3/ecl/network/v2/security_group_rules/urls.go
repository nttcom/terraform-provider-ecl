package security_group_rules

import "github.com/nttcom/eclcloud/v3"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("security-group-rules", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("security-group-rules")
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

func deleteURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
