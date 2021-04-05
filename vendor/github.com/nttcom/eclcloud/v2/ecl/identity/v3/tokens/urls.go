package tokens

import "github.com/nttcom/eclcloud/v2"

func tokenURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
