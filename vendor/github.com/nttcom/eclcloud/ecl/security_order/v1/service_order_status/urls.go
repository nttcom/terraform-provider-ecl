package service_order_status

import (
	"fmt"

	"github.com/nttcom/eclcloud"
)

func getURL(client *eclcloud.ServiceClient, deviceType string) string {
	var part string
	switch deviceType {
	case "WAF":
		part = "FGWAF"
		break
	case "HostBased":
		part = "HBS"
		break
	default:
		part = "FGS"
	}

	url := fmt.Sprintf("API/ScreenEvent%sOrderProgressRate", part)
	return client.ServiceURL(url)
}
