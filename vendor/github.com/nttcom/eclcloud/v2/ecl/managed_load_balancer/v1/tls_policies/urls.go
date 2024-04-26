package tls_policies

import (
	"github.com/nttcom/eclcloud/v2"
)

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("tls_policies")
}

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("tls_policies", id)
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func showURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
