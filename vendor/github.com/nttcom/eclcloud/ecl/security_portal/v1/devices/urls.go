package devices

import (
	"github.com/nttcom/eclcloud"
)

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("ecl-api/devices")
}
