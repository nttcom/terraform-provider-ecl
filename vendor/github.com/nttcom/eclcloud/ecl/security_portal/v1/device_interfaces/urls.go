package device_interfaces

import (
	"fmt"

	"github.com/nttcom/eclcloud"
)

func listURL(client *eclcloud.ServiceClient, serverUUID string) string {
	url := fmt.Sprintf("ecl-api/devices/%s/interfaces", serverUUID)
	return client.ServiceURL(url)
}
