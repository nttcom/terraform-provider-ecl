package ha_ports

import (
	"fmt"

	"github.com/nttcom/eclcloud"
)

func updateURL(client *eclcloud.ServiceClient, hostName string) string {
	url := fmt.Sprintf("ecl-api/ports/utm/ha/%s", hostName)
	return client.ServiceURL(url)
}
