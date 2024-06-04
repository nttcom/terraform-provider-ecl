package zones

import "github.com/nttcom/eclcloud/v3"

func baseURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("zones")
}

func zoneURL(c *eclcloud.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID)
}
