package availabilityzones

import "github.com/nttcom/eclcloud/v2"

func listURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
