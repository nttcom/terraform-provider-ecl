package bootfromvolume

import "github.com/nttcom/eclcloud/v4"

func createURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
