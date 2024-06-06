package host_based

import (
	"github.com/nttcom/eclcloud/v3"
)

func getURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/ScreenEventHBSOrderInfoGet")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryHBS")
}

func deleteURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryHBS")
}

func updateURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryHBS")
}
