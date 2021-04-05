package bootfromvolume

import "github.com/nttcom/eclcloud/v2"

func createURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
