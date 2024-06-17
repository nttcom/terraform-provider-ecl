package ports

import (
	"fmt"

	"github.com/nttcom/eclcloud/v3"
)

func updateURL(client *eclcloud.ServiceClient, deviceType string, hostName string) string {
	url := fmt.Sprintf("ecl-api/ports/%s/%s", deviceType, hostName)
	return client.ServiceURL(url)
}
