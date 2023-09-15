package network_based_device_ha

import (
	"github.com/nttcom/eclcloud/v4"
)

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/ScreenEventFGHADeviceGet")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryFGHA")
}

func deleteURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryFGHA")
}

func updateURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryFGHA")
}
