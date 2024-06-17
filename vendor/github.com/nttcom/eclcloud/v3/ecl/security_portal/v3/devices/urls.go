package devices

import (
	"github.com/nttcom/eclcloud/v3"
)

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("ecl-api/devices")
}
