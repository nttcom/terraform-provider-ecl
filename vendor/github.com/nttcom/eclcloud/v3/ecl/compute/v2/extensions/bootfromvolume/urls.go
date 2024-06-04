package bootfromvolume

import "github.com/nttcom/eclcloud/v3"

func createURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
