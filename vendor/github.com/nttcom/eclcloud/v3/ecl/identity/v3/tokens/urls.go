package tokens

import "github.com/nttcom/eclcloud/v3"

func tokenURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
