package network_based_device_single

import (
	"fmt"

	"github.com/nttcom/eclcloud/v2"
)

func getURLPartFromDeviceType(deviceType string) string {
	if deviceType == "WAF" {
		return "WAF"
	}
	return "S"
}

func listURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := getURLPartFromDeviceType(deviceType)
	url := fmt.Sprintf("API/ScreenEventFG%sDeviceGet", part)
	return client.ServiceURL(url)
}

func createURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := getURLPartFromDeviceType(deviceType)
	url := fmt.Sprintf("API/SoEntryFG%s", part)
	return client.ServiceURL(url)
}

func deleteURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := getURLPartFromDeviceType(deviceType)
	url := fmt.Sprintf("API/SoEntryFG%s", part)
	return client.ServiceURL(url)
}

func updateURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := getURLPartFromDeviceType(deviceType)
	url := fmt.Sprintf("API/SoEntryFG%s", part)
	return client.ServiceURL(url)
}
