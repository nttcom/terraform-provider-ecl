package volumeattach

import "github.com/nttcom/eclcloud/v3"

const resourcePath = "os-volume_attachments"

func resourceURL(c *eclcloud.ServiceClient, serverID string) string {
	return c.ServiceURL("servers", serverID, resourcePath)
}

func listURL(c *eclcloud.ServiceClient, serverID string) string {
	return resourceURL(c, serverID)
}

func createURL(c *eclcloud.ServiceClient, serverID string) string {
	return resourceURL(c, serverID)
}

func getURL(c *eclcloud.ServiceClient, serverID, aID string) string {
	return c.ServiceURL("servers", serverID, resourcePath, aID)
}

func deleteURL(c *eclcloud.ServiceClient, serverID, aID string) string {
	return getURL(c, serverID, aID)
}
