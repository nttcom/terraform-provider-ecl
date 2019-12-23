package tokens

import "github.com/nttcom/eclcloud"

func tokenURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
