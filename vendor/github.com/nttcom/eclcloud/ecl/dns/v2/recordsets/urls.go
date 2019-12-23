package recordsets

import "github.com/nttcom/eclcloud"

func baseURL(c *eclcloud.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID, "recordsets")
}

func rrsetURL(c *eclcloud.ServiceClient, zoneID string, rrsetID string) string {
	return c.ServiceURL("zones", zoneID, "recordsets", rrsetID)
}
