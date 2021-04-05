package devices

import (
	"github.com/nttcom/eclcloud/v2"
)

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("ecl-api/devices")
}
