package availabilityzones

import "github.com/nttcom/eclcloud"

func listURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
