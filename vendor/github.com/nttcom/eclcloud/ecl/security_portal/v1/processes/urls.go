package processes

import (
	"fmt"

	"github.com/nttcom/eclcloud"
)

func getURL(client *eclcloud.ServiceClient, processID string) string {
	url := fmt.Sprintf("ecl-api/process/%s/status", processID)
	return client.ServiceURL(url)
}
