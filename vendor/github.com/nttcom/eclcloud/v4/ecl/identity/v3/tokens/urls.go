package tokens

import "github.com/nttcom/eclcloud/v4"

func tokenURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
