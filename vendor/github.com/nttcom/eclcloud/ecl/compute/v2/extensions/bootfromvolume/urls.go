package bootfromvolume

import "github.com/nttcom/eclcloud"

func createURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("os-volumes_boot")
}
