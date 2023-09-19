package availabilityzones

import "github.com/nttcom/eclcloud/v4"

func listURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
